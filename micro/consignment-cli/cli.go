package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-grpc"
	pb "github.com/nasuku/learn/micro/consignment-service/proto/consignment"
	"golang.org/x/net/context"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(filename string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := grpc.NewService()
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("could not parse file :%v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("getConsignments returned %v", err)
	}
	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
