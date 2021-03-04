package process_service

import (
	"encoding/json"
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/services/task_service"
	"time"
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
