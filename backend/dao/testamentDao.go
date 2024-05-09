package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
	"time"
)

func SaveTestament(username string, testamentDetail string, testamentStyle string, testamentFileName string, testamentName string) {
	DB := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO Testament(account, testamentDetail, testamentStyle, testamentFileName, testamentName) values(?, ?, ?, ?, ?)"

	DB.Exec(insertSQL, username, testamentDetail, testamentStyle, testamentFileName, testamentName)
}

func GetTestamentByUserName(username string) []entity.Testament {
	DB := mysqlconn.GetDBCon()

	selectSQL := "SELECT testamentDetail, testamentStyle, testamentFileName, testamentName, isActive, createDate FROM Testament WHERE account = ?"

	rows, err := DB.Query(selectSQL, username)
	var testamentSlice []entity.Testament = make([]entity.Testament, 0)

	if err != nil {
		log.Warning.Println(username, "查询遗嘱失败:", err)
		return testamentSlice
	}
	for rows.Next() {
		var testamentDetail, testamentFileName, testamentName, testamentStyle string
		var createDate time.Time
		var isActive bool
		rows.Scan(&testamentDetail, &testamentStyle, &testamentFileName, &testamentName, &isActive, &createDate)
		testamentSlice = append(testamentSlice, entity.Testament{
			Username:          "",
			TestamentDetail:   testamentDetail,
			TestamentStyle:    testamentStyle,
			TestamentFileName: testamentFileName,
			IsActive:          isActive,
			TestamentName:     testamentName,
			CreateDate:        createDate,
		})
	}
	return testamentSlice
}
