package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/daymenu/Go-000/Week01/example/grpc/model"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Println("connect server success")
	defer conn.Close()

	client := model.NewCircularServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	area, err := client.Area(ctx, &model.AreaRequest{RequestId: "2233", Circular: &model.Circular{Dot: &model.Point{X: 1.1, Y: 1.1}, Radius: 1.1}})

	fmt.Println(area, err)
}
