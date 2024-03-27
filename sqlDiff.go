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

func (d *DB) UpdateDifferenceConfirmation(query string) error {
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
	diffs := Diffs{}
	for i, _ := range before {
		diff := Diff{
			Before: before[i],
			After:  after[i],
		}
		diff.check()
		diffs = append(diffs, diff)
	}
	for _, v := range diffs {
		fmt.Printf("Before:%s\n", v.Before)
		fmt.Printf("After:%s\n", v.After)
		fmt.Println("diff")
		for _, x := range v.Difference {
			for key, val := range x {
				fmt.Printf("%s:\n -%s\n +%s\n", key, val.Before, val.After)
			}
		}
	}
	return nil
}

type Diff struct {
	Before     map[string]interface{}
	After      map[string]interface{}
	Difference []map[string]df
}

type Diffs []Diff

func (d *Diff) check() {
	for key, val := range d.Before {
		if fmt.Sprintf("%s", val) != fmt.Sprintf("%s", d.After[key]) {
			df := map[string]df{
				key: {
					Before: val,
					After:  d.After[key],
				},
			}
			d.Difference = append(d.Difference, df)
		}
	}
}

type df struct {
	Before interface{}
	After  interface{}
}
