package entity

import "time"

type Testament struct {
	Username          string
	TestamentDetail   string
	TestamentStyle    string
	TestamentFileName string
	TestamentName     string
	IsActive          bool
	CreateDate        time.Time
}
