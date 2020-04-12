package rpcx

import (
	rpcl "crab.com/rpcl/grpcl"
	"testing"
)

func BenchmarkRpcxClient(b *testing.B) {
	b.Run("rpcx", func(b *testing.B) {
		for i:=0;i<b.N ;i++  {
			RpcxClient()
		}
	})

	b.Run("grpc", func(b *testing.B) {
		for i:=0;i<b.N ;i++  {
			rpcl.GrpcClient()
		}
	})
	//BenchmarkRpcxClient/rpcx-8         	    1579	    837159 ns/op
	//BenchmarkRpcxClient/grpc-8         	     954	   1149547 ns/op
}