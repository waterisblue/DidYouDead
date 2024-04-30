package mysqlconn

import (
	sql "database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Open() (err error) {
	DB, err = sql.Open("mysql", "root:zhang@tcp(121.36.99.228:3306)/DidYouDead")
	if err != nil {
		log.Fatal(err)
	}

	return
}

func GetDBCon() *sql.DB {
	return DB
}
