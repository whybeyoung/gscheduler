package runner

import (
	"github.com/maybaby/gscheduler/pkg/logging"
	"github.com/maybaby/gscheduler/pkg/util"
	"github.com/maybaby/gscheduler/services/lock_service"
	"github.com/maybaby/gscheduler/services/process_service"
	"time"
)

type MasterSchedulerService struct {
}

func (m *MasterSchedulerService) Run() error {
	locker := lock_service.GetAndInitLocker()
	logging.Info("Master scheudler start.")
	//c := make(chan os.Signal)
	// 监听信号
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for {
			//for s := range c {
			//	switch s {
			//	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			//		logging.Info("退出:", s)
			//		os.Exit(-1)
			//	default:
			//		logging.Debug("Master scheduler Service ")
			//	}
			//}
			lock := locker.GetLock("master")
			cmd, err := process_service.FindOneCommand()
			if cmd != nil && err == nil {
				processInstance, err := process_service.HandleCommand(cmd, util.GetLocalAddress())
				if processInstance != nil && err == nil {
					// TODO support Dag split here
					logging.Info("开始分解ProcessInstance, 切割为Task ...")
					meg := &MasterExecGoroutine{
						cmd,
					}
					meg.Run()
				}
			} else {
				time.Sleep(1 * time.Second)
			}
			lock.Release()
		}

	}()
	return nil
}
