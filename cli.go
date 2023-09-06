package main

import (
	"dtm-clis/busis"
	"fmt"
	"time"

	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/lithammer/shortuuid/v3"
)

var dtmSvc = "k8s://ohmyfans/dtm-svc:36790"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	for {
		SagaBarrier()
		time.Sleep(time.Second * 4)
	}
}

func SagaBarrier() {
	gid := shortuuid.New()
	addReq := &busis.AddReq{UserId: 1, Money: 50}
	delReq := &busis.AddReq{UserId: 2, Money: 50}
	saga := dtmgrpc.NewSagaGrpc(dtmSvc, gid).
		Add("k8s://ohmyfans/dtm-rpc-svc:9091/busis.Busis/AddMoney", "k8s://ohmyfans/dtm-rpc:9091/busis.Busis/DelMoney", addReq).
		Add("k8s://ohmyfans/dtm-rpc-svc:9091/busis.Busis/AddMoney", "k8s://ohmyfans/dtm-rpc:9091/busis.Busis/DelMoney", delReq)

	err := saga.Submit()
	logger.FatalIfError(err)
	fmt.Println(saga.Gid)
}

func Saga() {
}
