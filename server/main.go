package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mfbmina/poc_grpc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingPongServer
}

func (s *server) Ping(_ context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	r := &pb.PingResponse{}
	m := in.GetMessage()
	log.Println("Received message:", m)

	switch m {
	case "ping", "Ping", "PING":
		r.Message = "pong"
	default:
		r.Message = "I don't understand"
	}

	log.Println("Sending message:", r.Message)

	return r, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterPingPongServer(s, &server{})
	log.Printf("Server listening at %v", l.Addr())

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
