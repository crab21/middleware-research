package rpcl

import "testing"

func TestGrpcClient(t *testing.T) {
	go GrpcClient()


}

//BenchmarkGrpcClient-8   	    1024	   1153583 ns/op
func BenchmarkGrpcClient(b *testing.B) {
	for i:=0;i<b.N ;i++  {
		GrpcClient()
	}
}
