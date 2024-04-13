package main

import (
	"github.com/meschbach/xp-go-dyanmic-transcode/ipc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9432")
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instantiate the server
	s := grpc.NewServer()
	ipc.RegisterTranscodeServer(s, &Transcoder{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
