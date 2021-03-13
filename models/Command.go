package models

import "time"

type Command struct {
	Id string `json:"id"`

	CommandType CommandType `json:"command_type"`

	ProcessDefinitionId string `json:"process_definition_id"`

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

type TaskDependType struct {
}

type WarningType struct {
}

type FailureStrategy struct {
}
