package main

import (
	"fmt"

	"github.com/daymenu/Go-000/Week01/example/grpc/serverstream/model"
)

type circularService struct {
	*model.UnimplementedCircularServiceServer
}

func main() {
	fmt.Println("hello")
}
