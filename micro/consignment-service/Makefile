build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/nasuku/learn/micro/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service

clean:
	rm -f consignment-service
