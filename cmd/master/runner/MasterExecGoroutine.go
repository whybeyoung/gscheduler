package runner

import (
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/pkg/logging"
)

type MasterExecGoroutine struct {
	*models.Command
}

func (meg *MasterExecGoroutine) Run() {
	go func() {
		logging.Info("getting... ", meg.CommandType)
	}()
}
