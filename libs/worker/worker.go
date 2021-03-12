package worker

import (
	"fmt"
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/services/lock_service"
	"time"
)

type TaskExecuteProcessor struct {
}

func (tp TaskExecuteProcessor) Process(cmd *models.Command, reply *int) error {
	locker := lock_service.GetAndInitLocker()
	for i := 0; i < 5; i++ {
		fmt.Println("获取到Command", cmd.CommandType)
		lock := locker.GetLock()
		fmt.Println("等待10s", cmd.CommandType)

		time.Sleep(10 * time.Second)

		locker.ReleaseLock(lock)

	}
	return nil
}
