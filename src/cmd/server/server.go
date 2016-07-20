package main

import (
  "log"
  "net"

  // pb "github.com/integralist/test-grpc-custom/pb"
  pb "pb"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) Process(ctx context.Context, in *pb.Config) (*pb.Response, error) {
  return &pb.Response{Message: "Hello " + in.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterRequesterServer(s, &server{})
  s.Serve(lis)
}
