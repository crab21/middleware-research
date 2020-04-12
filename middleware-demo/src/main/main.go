package main

import (
	rpcl "crab.com/rpcl/grpcl"
	"crab.com/rpcl/rpcx"
)

func main() {
	//workmod.WorkStart()
	//queuemod.StartQueue()
	//md5Mod.Md5_Start()
	go rpcx.RpcxServer()
	go rpcl.GrpcServer()
	select {

	}
}
