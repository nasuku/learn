package main

import (
	// Import the generated protobuf code
	"context"
	"fmt"
	"log"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	pb "github.com/nasuku/learn/micro/consignment-service/proto/consignment"
)

// IRepository is interface
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
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

// GetAll returns all consignments
func (r *Repository) GetAll() []*pb.Consignment {
	return r.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	c, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = c
	return nil
}
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	resp.Consignments = s.repo.GetAll()
	return nil
}

func main() {
	repo := &Repository{}

	// Create a new service. Optionally include some options here.
	srv := grpc.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	fmt.Println("Ready to serve")
	if err := srv.Run(); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
