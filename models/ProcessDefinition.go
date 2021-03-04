package models

import (
	"time"
)

type ProcessDefinition struct {
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

// CheckAuth checks if authentication information exists
func SaveDefinition(sd *ProcessDefinition) error {
	if err := db.Create(&sd).Error; err != nil {
		return err
	}

	return nil
}
