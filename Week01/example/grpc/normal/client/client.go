package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/daymenu/Go-000/Week01/example/grpc/normal/model"
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
	circularPoint := &model.Point{X: 1.1, Y: 1.1}
	radius := 1.1
	resp, err := client.Area(ctx, &model.AreaRequest{RequestId: "2233", Circular: &model.Circular{Dot: circularPoint, Radius: radius}})

	if resp.Code != http.StatusOK {
		fmt.Println(err)
	}
	fmt.Printf("圆点为：(%.2f,%.2f)\n圆的半径为： %.2f\n圆的面积：%.f \n", circularPoint.X, circularPoint.Y, radius, resp.GetArea())
}
