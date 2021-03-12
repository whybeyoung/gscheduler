package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maybaby/gscheduler/models"
	"github.com/maybaby/gscheduler/services/client_service"
	"log"
	"net/http"
	"sync"

	"github.com/maybaby/gscheduler/pkg/app"
	"github.com/maybaby/gscheduler/pkg/e"
)

func foo(xc *client_service.XClient, ctx context.Context, typ, serviceMethod string, command *models.Command) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, command, &reply)
	case "broadcast":
		err = xc.Broadcast(ctx, serviceMethod, command, &reply)
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err)
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, command.CommandType, command.CommandParam, reply)
	}
}

func call(registry string) {
	d := client_service.NewGsRegistryDiscovery(registry, 0)
	xc := client_service.NewXClient(d, client_service.RandomSelect, nil)
	defer func() { _ = xc.Close() }()
	// send request & receive response
	var wg sync.WaitGroup
	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		foo(xc, context.Background(), "call", "TaskExecuteProcessor.Process", &models.Command{CommandType: 1})
	}(1)
	wg.Wait()
}

// @Summary Call a rpc procedure
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/call [get]
func CallRpc(c *gin.Context) {
	appG := app.Gin{C: c}
	registryAddr := "http://localhost:9999/_gsrpc_/registry"
	//time.Sleep(time.Second*16)
	call(registryAddr)

	appG.Response(http.StatusOK, e.SUCCESS, "haha")
}
