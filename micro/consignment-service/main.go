package main

import (
	// Import the generated protobuf code
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/nasuku/learn/micro/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// IRepository is interface
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository is implementation of repo
type Repository struct {
	consignments []*pb.Consignment
}

// Create a Consignment
func (r *Repository) Create(c *pb.Consignment) (*pb.Consignment, error) {
	r.consignments = append(r.consignments, c)
	return c, nil
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	c, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Consignment: c}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)
	fmt.Println("Ready to serve")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
