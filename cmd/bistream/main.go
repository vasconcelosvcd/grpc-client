package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/vasconcelosvcd/grpc-server/api/bistream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//Register Clients
	c := pb.NewDoubleStreamClient(conn)

	stream, err := c.DoubleStream(context.Background())

	if err != nil {
		log.Fatalf("unable to start stream, %s", err)
	}

	var cnt int32
	for {
		req := &pb.DoubleStreamRequest{
			Num: cnt,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.Send() = %v", stream, err)
		}

		reply, err := stream.Recv()
		if err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}

		log.Printf("received reply %+v", reply)

		cnt++
		time.Sleep(1 * time.Second)
	}

}
