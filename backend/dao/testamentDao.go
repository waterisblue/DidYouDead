package dao

import (
	"dyd/log"
	"dyd/mysqlconn"
)

func SaveTestament(username string, testamentDetail string, testamentStyle string, testamentFileName string, testamentName string) {
	DB := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO Testament(account, testamentDetail, testamentStyle, testamentFileName, testamentName) values(?, ?, ?, ?, ?)"

	DB.Exec(insertSQL, username, testamentDetail, testamentStyle, testamentFileName, testamentName)
}

func GetTestamentByUserName(username string) {
	DB := mysqlconn.GetDBCon()

	selectSQL := "SELECT * FROM Testament WHERE account = ?"

	rows, err := DB.Query(selectSQL, username)

	if err != nil {
		log.Warning.Println(username, "查询遗嘱失败:", err)
		return
	}

	// while row.Next() {

	// }
	return
}
