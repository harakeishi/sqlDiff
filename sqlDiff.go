package sqldiff

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/xwb1989/sqlparser"
)

type DB struct {
	*sqlx.DB
}

type Rows []map[string]interface{}

const rowsSliceCapacity = 1000

func Connect(driverName, dataSourceName string) (*DB, error) {
	db, err := Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, err
}

func (d *DB) UpdateDifferenceConfirmation(dest interface{}, query string) error {
	tx, err := d.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := sqlparser.Parse(query)
	if err != nil {
		return fmt.Errorf("failed to parse query:%s", err)
	}
	selectSql := ""
	before := make(Rows, 0, rowsSliceCapacity)
	after := make(Rows, 0, rowsSliceCapacity)
	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Update:
		selectSql = fmt.Sprintf("Select * from %s %s", sqlparser.String(stmt.TableExprs), sqlparser.String(stmt.Where))
	default:
		return fmt.Errorf("query is not an UPDATE statement")
	}
	rows, err := tx.Queryx(selectSql)
	if err != nil {
		return err
	}
	for rows.Next() {
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		if err != nil {
			return err
		}
		before = append(before, result)
	}
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}
	rows, err = tx.Queryx(selectSql)
	if err != nil {
		return err
	}
	for rows.Next() {
		result := make(map[string]interface{})
		err = rows.MapScan(result)
		if err != nil {
			return err
		}
		after = append(after, result)
	}
	for i, _ := range before {

		fmt.Printf("- %+s\n+ %+s\n", before[i], after[i])
	}
	return nil
}
