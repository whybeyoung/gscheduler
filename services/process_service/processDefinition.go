package process_service

import (
	"encoding/json"
	"errors"
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/pkg/e"
	"github.com/maybaby/gscheduler/pkg/logging"
	"github.com/maybaby/gscheduler/pkg/setting"
	"github.com/maybaby/gscheduler/pkg/util"
	"github.com/maybaby/gscheduler/services/task_service"
	"time"
)

type RunMode int32

const (
	Serial RunMode = iota
	Parallel
)

type ProcessData struct {
	Tasks        []*task_service.TaskNode
	GlobalParams []*task_service.Property
	Timeout      int
}

type ProcessDefinition struct {
	Name        string
	CreateTime  time.Time
	UpdateTime  time.Time
	ProcessData *ProcessData
	GroupId     string
	Description string
}

func (p *ProcessDefinition) Save() error {
	pr := &models.ProcessDefinition{
		Name:                  p.Name,
		Version:               "1", // Save 第一次默认为1
		Description:           p.Description,
		GroupId:               p.GroupId,
		CreateTime:            p.CreateTime,
		UpdateTime:            p.UpdateTime,
		ProcessDefinitionJson: p.ProcessData.ToJson(),
	}
	if err := models.SaveDefinition(pr); err != nil {
		return err
	}

	return nil
}

func (p *ProcessData) ToJson() string {
	b, err := p.MarshalJSON()
	if err != nil {
		return "{}"
	}
	return string(b)
}

func (p *ProcessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Tasks        []*task_service.TaskNode `json:"tasks"`
		GlobalParams []*task_service.Property `json:"globalParams"`
		Timeout      int                      `json:"timeout"`
	}{
		Tasks:        p.Tasks,
		GlobalParams: p.GlobalParams,
		Timeout:      p.Timeout,
	})
}

/*
 * 反序列化
 */
func (p *ProcessData) UnmarshalJSON(data []byte) error {
	aux := &struct {
		Tasks        []*task_service.TaskNode `json:"tasks"`
		GlobalParams []*task_service.Property `json:"globalParams"`
		Timeout      int                      `json:"timeout"`
	}{
		Tasks:        p.Tasks,
		GlobalParams: p.GlobalParams,
		Timeout:      p.Timeout,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.Timeout = aux.Timeout
	p.Tasks = aux.Tasks
	p.GlobalParams = aux.GlobalParams
	return nil
}

/*
 启动1个 Process实例， 分解 processDefintion 为多个Command Task
*/
func ExecProcessInstance(cmdType models.CommandType, groupId string, processDefinitionId int, workerGroup string, timeout int, runMode RunMode) error {
	if timeout <= 0 || timeout > setting.MAX_TASK_TIMEOUT {
		return errors.New(string(e.ERROR_PROCESS_TIMEOUT))
	}

	pd, err := models.GetProcessDefinition(processDefinitionId)
	if err != nil {
		return errors.New(string(e.ERROR_PROCESS_NOTFOUND))
	}
	if pd == nil {
		return errors.New(string(e.ERROR_PROCESS_NOTFOUND))

	}
	// 简单校验
	err = pd.CheckProcessDefinitionValid()
	processData := &ProcessData{}
	processData.UnmarshalJSON([]byte(pd.ProcessDefinitionJson))

	/**
	 * create command
	 */
	cmd := &models.Command{
		CommandType:         cmdType,
		ProcessDefinitionId: pd.ID,
		ExecutorId:          "",
		WorkerGroup:         workerGroup,
	}
	err = models.SaveCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}

func FindOneCommand() (*models.Command, error) {
	cmd, err := models.GetOneCommandToRun()
	return cmd, err
}

func HandleCommand(command *models.Command, host string) (*models.ProcessInstance, error) {
	logging.Info("Find one Command To Run...")

	processInstance := constructProcessInstance(command, host)
	if processInstance == nil {
		logging.Error("scan command, command parameter is error: {}", command)
		moveToErrorCommand(command, "process instance is null")
		return nil, nil
	}
	processInstance.CommandType = command.CommandType
	//TODO add history cmd
	//processInstance.addHistoryCmd(command.getCommandType());
	models.SaveProcessInstance(processInstance)
	models.DeleteCommand(command)
	return processInstance, nil
}

func moveToErrorCommand(command *models.Command, msg string) {
	//TODOs
	logging.Info("Not implement.. TODO..")

}

func constructProcessInstance(command *models.Command, host string) *models.ProcessInstance {
	cmdType := command.CommandType
	cmdParam := util.ToMap(command.CommandParam)
	var processInstance *models.ProcessInstance = nil
	var processDefinition *models.ProcessDefinition = nil
	var err error
	if command.ProcessDefinitionId != 0 {
		processDefinition, err = models.GetProcessDefinition(command.ProcessDefinitionId)
		if err != nil || processDefinition == nil {
			logging.Error("Not found Definition...")
			return nil
		}
	}
	if cmdParam != nil {
		processInstance = &models.ProcessInstance{}
		logging.Info(cmdType)
		return processInstance
	} else {
		// generate one new process instance
		processInstance = generateNewProcessInstance(processDefinition, command, command.CommandParam)
	}
	processInstance.Host = host
	var runStatus models.ExecutionStatus = models.RUNNING_EXEUTION
	//runTimes := processInstance.Runtimes
	switch command.CommandType {
	case models.START_PROCESS:
		break
	case models.START_FAILURE_TASK_PROCESS:
		// find failed tasks and init these tasks
		break
		//TODO
	case models.START_CURRENT_TASK_PROCESS:
		break
	case models.RECOVER_WAITTING_THREAD:
		break
	case models.RECOVER_SUSPENDED_PROCESS:
		break
	case models.RECOVER_TOLERANCE_FAULT_PROCESS:
		break
	case models.COMPLEMENT_DATA:
		break
	case models.REPEAT_RUNNING:
	case models.SCHEDULER:
		break
	default:
		break
	}
	processInstance.State = runStatus

	return processInstance
}

func generateNewProcessInstance(definition *models.ProcessDefinition,
	command *models.Command,
	cmdParam string) *models.ProcessInstance {
	processInstance := &models.ProcessInstance{}
	processInstance.ProcessDefinitionId = command.ProcessDefinitionId
	processInstance.State = models.RUNNING_EXEUTION
	processInstance.StartTime = time.Now()
	processInstance.RunTimes = 1
	processInstance.MaxTryTimes = 0
	processInstance.CommandParam = cmdParam
	processInstance.CommandType = command.CommandType
	if command.WorkerGroup == "" {
		command.WorkerGroup = setting.DEFAULT_WORKER_GROUP
	} else {
		command.WorkerGroup = command.WorkerGroup
	}
	processInstance.Timeout = processInstance.Timeout
	//copy process define json to process instance
	processInstance.ProcessInstanceJson = definition.ProcessDefinitionJson
	// set process instance priority
	processInstance.ProcessInstancePriority = command.ProcessInstancePriority
	return processInstance

}
