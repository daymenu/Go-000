package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/daymenu/Go-000/Week09/pkg"
)

func main() {
	fmt.Println("hello")

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		log.Printf("%s is connect", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	message := make(chan string, 0)
	go read(ctx, conn, message, cancel)
	go write(ctx, conn, message, cancel)
	go pkg.Heart(ctx, conn, cancel)
}

func read(ctx context.Context, conn net.Conn, message chan<- string, cancel context.CancelFunc) {
	log.Println("server read start...")
	defer conn.Close()
	defer func() { log.Println("server read end...") }()
	scan := bufio.NewScanner(conn)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scan.Scan() {
				break
			}
			h := scan.Text()
			log.Printf("client[%s]: %s", conn.RemoteAddr().String(), h)
			if h == "exit" {
				cancel()
				return
			}
			message <- h
		}

	}
}

func write(ctx context.Context, conn net.Conn, hi <-chan string, cancel context.CancelFunc) {
	log.Println("server write start")
	defer conn.Close()
	defer func() { log.Println("server write end...") }()
	for {
		select {
		case <-ctx.Done():
			log.Printf("client[%s]: is close", conn.RemoteAddr().String())
			return
		case h := <-hi:
			conn.Write([]byte("server :" + h + "\n"))
			log.Printf("client[%s]: %s", conn.RemoteAddr().String(), "server :"+h)
		}
	}

}
