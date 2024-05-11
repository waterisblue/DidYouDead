package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
	"time"
)

// SaveTestament 保存遗嘱
func SaveTestament(username string, testamentDetail string, testamentStyle string, testamentFileName string, testamentName string) {
	DB := mysqlconn.GetDBCon()
	insertSQL := "INSERT INTO Testament(account, testamentDetail, testamentStyle, testamentFileName, testamentName) values(?, ?, ?, ?, ?)"

	DB.Exec(insertSQL, username, testamentDetail, testamentStyle, testamentFileName, testamentName)
}

// GetTestamentByUserName 通过用户名获取遗嘱
func GetTestamentByUserName(username string) []entity.Testament {
	DB := mysqlconn.GetDBCon()

	selectSQL := "SELECT id, testamentDetail, testamentStyle, testamentFileName, testamentName, createDate, isActive FROM Testament WHERE account = ?"

	rows, err := DB.Query(selectSQL, username)
	var testamentSlice = make([]entity.Testament, 0)

	if err != nil {
		log.Warning.Println(username, "查询遗嘱失败:", err)
		return testamentSlice
	}
	for rows.Next() {
		var testamentDetail, testamentFileName, testamentName, testamentStyle string
		var id int
		var createDate *time.Time
		var isActive bool
		err := rows.Scan(&id, &testamentDetail, &testamentStyle, &testamentFileName, &testamentName, &createDate, &isActive)
		if err != nil {
			log.Warning.Println(username, "遗嘱查询错误", err)
			return testamentSlice
		}
		testamentSlice = append(testamentSlice, entity.Testament{
			Id:                id,
			Username:          "",
			TestamentDetail:   testamentDetail,
			TestamentStyle:    testamentStyle,
			TestamentFileName: testamentFileName,
			IsActive:          isActive,
			CreateDate:        createDate,
			TestamentName:     testamentName,
		})
	}
	return testamentSlice
}

func UpdateTestamentActiveById(testamentId int, isActive bool) {
	DB := mysqlconn.GetDBCon()
	updateSQL := "UPDATE Testament SET isActive = ? WHERE id = ?"

	DB.Exec(updateSQL, isActive, testamentId)

	return
}
