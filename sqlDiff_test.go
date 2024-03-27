package sqldiff_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	sqldiff "github.com/harakeishi/sqlDiff"
)

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
				query: "update staff set first_name = 'tes',last_name = 'as' where position = 'Engineer'",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := sqldiff.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/employees")
			if err != nil {
				t.Errorf("%v", err)
			}
			if err := d.UpdateDifferenceConfirmation(tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("DB.UpdateDifferenceConfirmation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
