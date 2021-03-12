package main

import (
	gsrpc "github.com/maybaby/gscheduler"
	"github.com/maybaby/gscheduler/libs/worker"
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/pkg/gredis"
	"github.com/maybaby/gscheduler/pkg/logging"
	"github.com/maybaby/gscheduler/pkg/setting"
	"github.com/maybaby/gscheduler/pkg/util"
	"github.com/maybaby/gscheduler/services/registry_service"
	"net"
	"sync"
)

func startWorkerServer(registryAddr string, wg *sync.WaitGroup) {
	var tp worker.TaskExecuteProcessor
	l, _ := net.Listen("tcp", ":0")
	server := gsrpc.NewServer()
	_ = server.Register(&tp)
	registry_service.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0)
	wg.Done()
	server.Accept(l)
}

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}
func main() {
	var wg sync.WaitGroup
	registryAddr := "http://localhost:9999/_gsrpc_/registry"
	wg.Add(2)
	go startWorkerServer(registryAddr, &wg)
	wg.Wait()
}
