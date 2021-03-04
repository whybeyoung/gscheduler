package worker

import (
	"fmt"
	"github.com/maybaby/gscheduler/models"
)

type TaskExecuteProcessor struct {
}

func (tp TaskExecuteProcessor) Process(cmd *models.Command, reply *int) error {
	fmt.Println("获取到Command", cmd.CType)
	return nil
}
