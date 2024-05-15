package mysqlconn

import (
	sql "database/sql"
	"dyd/log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)

func Open() {
	var err error
	DB, err = sql.Open("mysql", "root:zhang@tcp(121.36.99.228:3306)/DidYouDead?parseTime=true")
	if err != nil {
		log.Error.Println("数据库打开错误", err)
	}
}

func GetDBCon() *sql.DB {
	return DB
}
