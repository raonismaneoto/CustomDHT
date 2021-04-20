package api

import (
	"context"
	"fmt"
	"github.com/raonismaneoto/CustomDHT/node/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

func (*server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	name := request.Name
	response := &proto.HelloReply{
		Message: "Hello " + name,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	s.Serve(lis)
}

