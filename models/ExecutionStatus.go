package models

type ExecutionStatus int32

const (
	SUBMITTED_SUCCESS ExecutionStatus = iota
	RUNNING_EXEUTION
	READY_PAUSE
	EPAUSE
	READY_STOP
	ESTOP
	FAILURE
	SUCCESS
	NEED_FAULT_TOLERANCE
	KILL
	WAITTING_THREAD
	WAITTING_DEPEND
)