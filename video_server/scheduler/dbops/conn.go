package dbops

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:csq123456@tcp(localhost:3306)/rushflow?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
