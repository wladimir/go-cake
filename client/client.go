package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"

	pb "github.com/wladimir/go-cake/cake"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCakeClient(conn)

	name := ""

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream, err := client.ListProblems(ctx, &pb.ProblemInput{})
	if err != nil {
		log.Fatalf("failed to list problems: %v", err)
	}

	for {
		p, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListProblems(_) = _, %v", client, err)
		}
		log.Printf("name: %q", p.GetName())
		if len(name) == 0 {
			name = p.GetName()
		}
	}

	r, err := client.RunProblem(ctx, &pb.ProblemDefinition{Name: name})
	if err != nil {
		log.Fatalf("could not run: %v", err)
	}
	log.Printf("Result: %s", r)
}
