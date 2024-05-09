package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func GetUserByAccount(username string) (exist bool, userdetail entity.UserDetail) {
	exist = false
	DB := mysqlconn.GetDBCon()
	selectSQL := "SELECT account, password, administrator FROM SysUser WHERE account = ?"
	log.Info.Println("DB:", DB)
	rows, _ := DB.Query(selectSQL, username)
	if rows.Next() {
		defer rows.Close()
		exist = true
		rows.Scan(&userdetail.Username, &userdetail.Password, &userdetail.Administrator)
	}
	return
}

func AddUser(user *entity.UserInfo) (err error) {
	db := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO SysUser(account, password) VALUES(?, ?)"
	_, err = db.Exec(insertSQL, user.Username, user.Password)
	return
}
