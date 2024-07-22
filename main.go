package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kodinggo/comment-service-gp1/internal/delivery/grpcsvc"
	pb "github.com/kodinggo/comment-service-gp1/pb/comment"

	"google.golang.org/grpc"
)

const grpcPort = 8083

func main() {
	// Run grpc server
	grpcServer := grpc.NewServer()

	commentService := grpcsvc.NewCommentService()

	pb.RegisterCommentServiceServer(grpcServer, commentService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Panicf("failed create http listener, error: %v", err)
	}

	log.Printf("grpc server is running with port: :%d", grpcPort)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Panicf("failed run grpc server, error: %v", err)
	}
}
