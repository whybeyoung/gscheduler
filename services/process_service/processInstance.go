package process_service

import "time"

type ProcessInstance struct {
	Name        string
	CreateTime  time.Time
	UpdateTime  time.Time
	ProcessData *ProcessData
	GroupId     string
	Description string
}
