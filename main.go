package main

import (
	"log"
	"net"
	"time"

	"github.com/maan19/distributed-cache-golang/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("SET Foo Bar 2400"))
	}()
	server := NewServer(opts, cache.New())
	server.Start()
}
