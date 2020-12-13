package main

import (
	"fmt"
	"log"
	"math"
	"net"

	"github.com/daymenu/Go-000/Week01/example/grpc/serverstream/model"
	"google.golang.org/grpc"
)

type circularService struct {
	*model.UnimplementedCircularServiceServer
}

func (c *circularService) Area(request *model.AreaRequest, stream model.CircularService_AreaServer) error {
	resp := new(model.AreaResponse)
	resp.Code = 200
	radius := request.Circular.Radius
	resp.Area = radius * radius * math.Pi
	stream.Send(resp)
	stream.Send(resp)
	return nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":10000"))
	if err != nil {
		log.Fatalf("listen err：%v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	log.Println("register area server")
	model.RegisterCircularServiceServer(grpcServer, new(circularService))
	log.Println("server is serve")
	grpcServer.Serve(lis)
}
