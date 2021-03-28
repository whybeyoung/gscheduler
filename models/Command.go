package models

import (
	"time"
)

type Command struct {
	Id string `json:"id"`

	CommandType CommandType `json:"command_type"`

	ProcessDefinitionId int `json:"process_definition_id"`

	ExecutorId string `json:"executor_id"`

	CommandParam string `json:"command_param"`

	TaskDependType TaskDependType `json:"task_depend_type"`

	WarningType WarningType `json:"warning_type"`

	FailureStrategy FailureStrategy `json:"failure_strategy"`

	WarningGroupId string `json:"warning_type"`

	StartTime time.Time `json:"start_time"`

	UpdateTime time.Time `json:"update_time"`

	ScheduleTime time.Time `json:"schedule_time"`

	WorkerGroup string `json:"worker_group"`

	ProcessInstancePriority Priority `json:"process_instance_priority"`
}

// TableName 会将 User 的表名重写为 `process_definition`
// 参考gorm约定
func (Command) TableName() string {
	return "command"
}

// CheckAuth checks if authentication information exists
func SaveCommand(c *Command) error {
	if err := db.Create(&c).Error; err != nil {
		return err
	}

	return nil
}

func GetOneCommandToRun() (*Command, error) {
	var cmd Command
	err := db.Table("t_gs_command").First(&cmd).Error
	return &cmd, err
}

func DeleteCommand(command *Command) error {
	err := db.Table("t_gs_command").Delete(command).Error
	return err
}

type TaskDependType int64

type WarningType int64

type FailureStrategy int64
