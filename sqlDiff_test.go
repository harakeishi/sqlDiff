package sqldiff_test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	sqldiff "github.com/harakeishi/sqlDiff"
)

type staff struct {
	Id         int     `db:"id"`
	FirstName  string  `db:"first_name"`
	LastName   string  `db:"last_name"`
	Position   string  `db:"position"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
	HireDate   string  `db:"hire_date"`
}

func TestDB_UpdateDifferenceConfirmation(t *testing.T) {
	type args struct {
		dest  interface{}
		query string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				dest:  make([]staff, 100),
				query: "select * from staff where id = 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := sqldiff.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/employees")
			if err != nil {
				fmt.Println(err)
			}
			staff := []staff{}
			if err := d.UpdateDifferenceConfirmation(&staff, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateDifferenceConfirmation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
