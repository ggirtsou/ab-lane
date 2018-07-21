package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ggirtsou/ab-lane/generated/pb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 9000, "Port to bind process")
	flag.Parse()

	var gracefulStop = make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if port == 0 {
		flag.Usage()
		return
	}

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panicf("ERR: failed to listen to port: %v", err)
	}

	defer ln.Close()

	log.Print("INF: up and running")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				conn, err := ln.Accept()
				if err != nil {
					log.Printf("ERR: failed to accept connection: %v", err)
					continue
				}
				go handleConnection(conn)
			}
		}
	}(ctx)

	<-gracefulStop
	log.Print("INF: shutting down")
	time.Sleep(2 * time.Second)
	cancel()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	bytes, err := ioutil.ReadAll(conn)
	if err != nil {
		// do not shutdown app because a connection couldn't be handled
		log.Printf("ERR: %v", err)
		return
	}

	var envelope contracts.RequestEnvelope
	if err := proto.Unmarshal(bytes, &envelope); err != nil {
		log.Printf("ERR: failed to unmarshal envelope: %v", err)
		return
	}

	switch envelope.Type {
	case contracts.RequestEnvelope_SERVE_MESSAGE:
		return // @todo handle serving message
	case contracts.RequestEnvelope_SAVE_MESSAGE:
		var msg contracts.PersistMessage
		if err := ptypes.UnmarshalAny(envelope.Payload, &msg); err != nil {
			log.Printf("ERR: failed to unmarshal payload: %v", err)
			return
		}

		// @todo save on disk
	default:
		log.Printf("ERR: unhandled message type: %v", err)
		return
	}
}
