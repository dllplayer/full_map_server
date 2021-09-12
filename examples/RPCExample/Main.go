package main

import (
	"encoding/json"
	"flag"
	"github.com/yxinyi/YCServer/engine/YMsg"
	"github.com/yxinyi/YCServer/engine/YNode"
	"github.com/yxinyi/YCServer/examples/RPCExamples/Logic/TestModule"
	"github.com/yxinyi/YCServer/examples/RPCExamples/Logic/TestModule2"
	_ "net/http/pprof"
)

func main() {
	flag.Parse()
	YNode.Register(
		TestModule.NewInfo(YNode.Obj()),
		TestModule2.NewInfo(YNode.Obj()),
	)
	{
		msg := &YMsg.S2S_rpc_msg{}
		msg.M_func_name = "Test_3"
		msg.M_tar.M_name = "TestModule"
		msg.M_func_parameter = make([][]byte, 0)
		{
			_bytes, _ := json.Marshal(1)
			msg.M_func_parameter = append(msg.M_func_parameter, _bytes)
		}
		{
			_bytes, _ := json.Marshal("测试")
			msg.M_func_parameter = append(msg.M_func_parameter, _bytes)
		}
		YNode.RPCCall(msg)
	}
	YNode.Start()
}