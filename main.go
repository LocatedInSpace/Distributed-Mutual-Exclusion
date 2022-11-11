package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync/atomic"

	ricart "github.com/LocatedInSpace/Distributed-Mutual-Exclusion/proto"
	"google.golang.org/grpc"
)

const CLIENTS = 3
const OFFSET int32 = 7000

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + OFFSET

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:      ownPort,
		clients: make(map[int32]ricart.RicartAndAgrawalaClient),
		replies: 0,
		ctx:     ctx,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	ricart.RegisterRicartAndAgrawalaServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < CLIENTS; i++ {
		port := OFFSET + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := ricart.NewRicartAndAgrawalaClient(conn)
		p.clients[port] = c
	}
	fmt.Printf("Connected to all clients :)")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.enter()
	}
}

type peer struct {
	ricart.UnimplementedRicartAndAgrawalaServer
	id      int32
	replies uint32
	clients map[int32]ricart.RicartAndAgrawalaClient
	ctx     context.Context
}

func (p *peer) Request(ctx context.Context, req *ricart.Info) (*ricart.Empty, error) {
	rep := &ricart.Empty{}
	go p.clients[req.Id].Reply(p.ctx, rep)
	return rep, nil
}

func (p *peer) Reply(ctx context.Context, req *ricart.Empty) (*ricart.Empty, error) {
	rep := &ricart.Empty{}
	atomic.AddUint32(&p.replies, 1)
	if atomic.LoadUint32(&p.replies) >= CLIENTS-1 {
		fmt.Printf("Entered critical section\n")
		atomic.StoreUint32(&p.replies, 0)
	}
	return rep, nil
}

func (p *peer) enter() {
	info := &ricart.Info{Id: p.id, Lamport: 0}
	for id, client := range p.clients {
		_, err := client.Request(p.ctx, info)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v\n", id)
	}
}
