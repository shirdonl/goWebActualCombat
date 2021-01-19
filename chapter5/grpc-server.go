//++++++++++++++++++++++++++++++++++++++++
// 《Go Web编程实战派从入门到精通》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 仓库地址：https://gitee.com/shirdonl/goWebActualCombat
// 仓库地址：https://github.com/shirdonl/goWebActualCombat
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"log"
	"net"
	// 导入生成好的 protobuf 包
	pb "gitee.com/shirdonl/goWebActualCombat/chapter5/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//定义服务结构体
type ProgrammerServiceServer struct{}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "shirdon" {
		resp = &pb.Response{
			Uid: 6,
			Username: name,
			Job: "CTO",
			GoodAt: []string{"Go","Java","PHP","Python"},
		}

	}
	err = nil
	return
}

func main() {
	port := ":8078"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	// 将 ProgrammerService 注册到 gRPC
	// 注意第二个参数 ProgrammerServiceServer 是接口类型的变量
	// 需要取地址传参
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Serve(l)
}
