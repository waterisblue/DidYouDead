package entity

import "time"

type Detection struct {
	Id         int
	Account    string
	HeartRate  int
	CreateDate *time.Time
}
