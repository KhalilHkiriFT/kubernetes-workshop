package model

import "time"

type Reservation struct {
	ID         string    `json:"id"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	EmployeeId string    `json:"employeeId"`
	PedelecId  string    `json:"pedelecId"`
}
