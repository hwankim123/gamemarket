package main

import (
	"log"
	"net"

	// .pb.go 파일 경로 설정 : 프로젝트 폴더부터 시작
	pb "first-go-grpc/protos"

	"google.golang.org/grpc"
)

const port = ":8080"

// server struct?
type server struct {
	pb.UnimplementedItemsServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterItemsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
