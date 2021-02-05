package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	pb "github.com/wladimir/go-cake/cake"
)

var (
	port = flag.Int("port", 10000, "port")
)

type cakeServer struct {
	pb.UnimplementedCakeServer

	problems map[string]*pb.ProblemDefinition
}

func (s *cakeServer) ListProblems(_ *pb.ProblemInput, stream pb.Cake_ListProblemsServer) error {
	log.Printf("Listing problems")

	for _, p := range s.problems {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func read(name string, wg *sync.WaitGroup, result *string) {
	defer wg.Done()

	content, err := ioutil.ReadFile(fmt.Sprintf("problems/%v", name))
	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
	}

	*result = string(content)
}

func run(name string, wg *sync.WaitGroup, result *string) {
	defer wg.Done()

	cmd := exec.Command("go", "run", fmt.Sprintf("problems/%s", name))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(output))

	*result = string(output)
}

func (s *cakeServer) RunProblem(_ context.Context, p *pb.ProblemDefinition) (*pb.ProblemResult, error) {
	log.Printf("Runing problem: %v", p.Name)

	wg := &sync.WaitGroup{}

	var content string
	var output string

	wg.Add(2)

	go read(p.Name, wg, &content)
	go run(p.Name, wg, &output)

	wg.Wait()

	log.Println(content)

	return &pb.ProblemResult{
		Name:    p.Name,
		Content: content,
		Output:  output,
	}, nil
}

func (s *cakeServer) readProblems() error {
	files, err := ioutil.ReadDir("problems")
	if err != nil {
		return errors.New("failed to read problems")

	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			s.problems[file.Name()] = &pb.ProblemDefinition{
				Name: file.Name(),
			}
		}
	}
	return nil
}

func newServer() *cakeServer {
	s := &cakeServer{problems: make(map[string]*pb.ProblemDefinition)}

	if err := s.readProblems(); err != nil {
		log.Fatal(err)
	}

	return s
}

func main() {
	flag.Parse()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh,
		os.Interrupt,
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGTERM)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterCakeServer(grpcServer, newServer())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	<-stopCh
}
