package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ProcessDefinition struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Version               string    `json:"version"`
	GroupId               string    `json:"group_id"`
	UserId                string    `json:"user_id"`
	ProcessDefinitionJson string    `json:"process_definition_json"`
	Description           string    `json:"description"`
	Flag                  string    `json:"flag"`
	CreateTime            time.Time `json:"create_time"`
	Timeout               string    `json:"timeout"`
	UpdateTime            time.Time `json:"update_time"`
}

// TableName 会将 User 的表名重写为 `process_definition`
// 参加gorm约定
func (ProcessDefinition) TableName() string {
	return "process_definition"
}

// 校验pd 合法性
func (p *ProcessDefinition) CheckProcessDefinitionValid() error {
	return nil
}

// CheckAuth checks if authentication information exists
func SaveDefinition(sd *ProcessDefinition) error {
	if err := db.Create(&sd).Error; err != nil {
		return err
	}

	return nil
}

// GetProcessDefinition Get a single ProcessDefinition based on ID
func GetProcessDefinition(id int) (*ProcessDefinition, error) {
	var pd ProcessDefinition
	err := db.Table("t_gs_process_definition").First(&pd, id).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &pd, nil
}
