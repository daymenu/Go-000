package pkg

import (
	"context"
	"log"
	"net"
	"time"
)

// Heart heart
func Heart(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	log.Printf("Heart[%s] is start...\n", conn.RemoteAddr())
	defer func() { log.Printf("Heart[%s] is end...\n", conn.RemoteAddr()) }()
	ticker := time.NewTicker(5 * time.Second)
	ping := []byte("ping\n")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, err := conn.Write(ping)
			log.Println(".")
			if err != nil {
				cancel()
				return
			}
		}
	}
}
