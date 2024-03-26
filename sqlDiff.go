package sqldiff

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

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
	if err := tx.Select(dest, query); err != nil {
		return err
	}
	fmt.Printf("%+v\n", dest)
	return nil
}
