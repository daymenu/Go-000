package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(os.Stdin)
	go func() {
		s := bufio.NewScanner(conn)
		for s.Scan() {
			log.Println(s.Text())
		}
	}()

	for scan.Scan() {
		t := scan.Text()
		_, err := conn.Write([]byte(t + "\n"))
		if t == "exit" {
			conn.Close()
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
