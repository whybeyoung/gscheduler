package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ProcessInstance struct {
	ID                  string          `json:"id"`
	ProcessDefinitionId string          `json:"processDefinitionId"`
	State               ExecutionStatus `json:"state"`
	Flag                Flag            `json:"flag"`
	startTime           time.Time       `json:"startTime"`
	endTime             time.Time       `json:"endTime"`
	Runtimes            int             `json:"runtimes"`
	Name                string          `json:"name"`
	Host                string          `json:"host"`
	CommandType         CommandType     `json:"commandType"`
	commandParam        string          `json:"commandParam"`
	maxTryTimes         int             `json:"maxTryTimes"`
	scheduleTime        time.Time       `json:"scheduleTime"`
	processInstanceJson string          `json:"processInstanceJson"`
	workerGroup         string          `json:"workerGroup"`
	Timeout             int             `json:"timeout"`
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
