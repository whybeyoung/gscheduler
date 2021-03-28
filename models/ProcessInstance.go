package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Priority int32

const (
	HIGHEST Priority = iota
	HIGH
	MEDIUM
	LOW
	LOWEST
)

type ProcessInstance struct {
	ID                      string          `json:"id"`
	ProcessDefinitionId     int             `json:"processDefinitionId"`
	State                   ExecutionStatus `json:"state"`
	Flag                    Flag            `json:"flag"`
	StartTime               time.Time       `json:"startTime"`
	EndTime                 time.Time       `json:"endTime"`
	RunTimes                int             `json:"runTimes"`
	Name                    string          `json:"name"`
	Host                    string          `json:"host"`
	CommandType             CommandType     `json:"commandType"`
	CommandParam            string          `json:"commandParam"`
	MaxTryTimes             int             `json:"maxTryTimes"`
	ScheduleTime            time.Time       `json:"scheduleTime"`
	ProcessInstanceJson     string          `json:"processInstanceJson"`
	WorkerGroup             string          `json:"workerGroup"`
	Timeout                 int             `json:"timeout"`
	ProcessInstancePriority Priority        `json:"process_instance_priority"`
}

// TableName 会将 User 的表名重写为 `process_definition`
// 参加gorm约定
func (ProcessInstance) TableName() string {
	return "process_instance"
}

// CheckAuth checks if authentication information exists
func SaveProcessInstance(pi *ProcessInstance) error {
	if err := db.Create(&pi).Error; err != nil {
		return err
	}

	return nil
}

func GetProcessInstance(id string) (*ProcessInstance, error) {
	var pd ProcessInstance
	err := db.Where("id = ? ", id).Table("t_gs_process_instance").First(&pd).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pd, nil
}
