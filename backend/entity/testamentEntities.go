package entity

import (
	"time"
)

type Testament struct {
	Id                int        `json:"Id"`
	Username          string     `json:"Username"`
	TestamentDetail   string     `json:"TestamentDetail"`
	TestamentStyle    string     `json:"TestamentStyle"`
	TestamentFileName string     `json:"TestamentFileName"`
	TestamentName     string     `json:"TestamentName"`
	IsActive          bool       `json:"IsActive"`
	CreateDate        *time.Time `json:"CreateDate"`
}
