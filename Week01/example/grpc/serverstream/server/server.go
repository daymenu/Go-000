rpcpackage main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	"github.com/daymenu/Go-000/Week01/example/grpc/serverstream/model"
)

type circularService struct {
	*model.UnimplementedCircularServiceServer
}

func (c *circularService) Area(ctx context.Context, request *model.CircularService_AreaServer) error {
	resp := new(model.AreaResponse)
	resp.Code = 200
	resp.Area = request.Circular.GetRadius() * request.Circular.GetRadius() * math.Pi
	rjson, err := json.Marshal(request)
	if err != nil {
		log.Printf("request json encode failed")
	}
	log.Printf("%s:%s", request.RequestId, string(rjson))
	return resp, nil
}
func newServer() *circularService {
	s := &circularService{}
	return s
}
func main() {
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("server is start ....")
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	model.RegisterCircularServiceServer(grpcServer, newServer())

	grpcServer.Serve(listen)
}
