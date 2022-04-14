package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/vasconcelosvcd/grpc-server/api/doubleNum"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultValue = 2
)

var (
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
	value = flag.Int("value", defaultValue, "Number to double")
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
	c := pb.NewDoubleNumClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DoubleNum(ctx, &pb.DoubleNumRequest{Num: int32(*value)})
	if err != nil {
		log.Fatalf("could not double: %v", err)
	}

	log.Printf("Doubled: %v", r.GetDoubled())
}
