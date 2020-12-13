package main

import (
	"context"
	"log"

	"github.com/daymenu/Go-000/Week01/example/grpc/serverstream/model"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("dial err : %v", err)
	}

	defer conn.Close()

	client := model.NewCircularServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ac, err := client.Area(ctx, &model.AreaRequest{Circular: &model.Circular{Dot: &model.Point{X: 1.1, Y: 1.1}, Radius: 20.0}})
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := ac.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.GetCode(), resp.GetArea())
	}

}
