package dao

import (
	"dyd/entity"
	"dyd/log"
	"dyd/mysqlconn"
)

func GetDetectionByDate() []entity.Detection {
	DB := mysqlconn.GetDBCon()
	selectSQL := `SELECT id, account, heartrate, createdate FROM DetectionLog order by createdate desc limit 10`
	rows, err := DB.Query(selectSQL)
	var detecitons []entity.Detection
	if err != nil {
		log.Warning.Println("最新心率数据查询失败，请检查：", err)
		return detecitons
	}
	for rows.Next() {
		var detection entity.Detection
		err := rows.Scan(&detection.Id, &detection.Account, &detection.HeartRate, &detection.CreateDate)
		if err != nil {
			log.Warning.Println("select变量获取失败", err)
			return detecitons
		}
		detecitons = append(detecitons, detection)
	}
	return detecitons

}
