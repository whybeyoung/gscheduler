package models

type CommandType int32

const (
	TASK_EXECUTE_REQUEST CommandType = iota

	TASK_EXECUTE_ACK

	TASK_EXECUTE_RESPONSE
)
