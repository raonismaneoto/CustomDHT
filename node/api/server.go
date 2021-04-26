package api

import (
	"context"
	"fmt"
	"github.com/raonismaneoto/CustomDHT/node"
	"github.com/raonismaneoto/CustomDHT/node/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// maybe a notify and a stabilize message are needed
type server struct {
	node *node.Node
}

func (*server) Ping(ctx context.Context, request *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (*server) Successor(ctx context.Context, request *proto.Empty) (*proto.SuccessorResponse, error) {
	var response *proto.SuccessorResponse
	return response, nil
}

func (*server) Predecessor(ctx context.Context, request *proto.Empty) (*proto.PredecessorResponse, error) {
	var response *proto.PredecessorResponse
	return response, nil
}

func (*server) HandleNewPredecessor(ctx context.Context, request *proto.HandleNewPredecessorRequest) (*proto.HandleNewPredecessorResponse, error) {
	var response *proto.HandleNewPredecessorResponse
	return response, nil
}

func (*server) HandleNewSuccessor(ctx context.Context, request *proto.HandleNewSuccessorRequest) (*proto.HandleNewSuccessorResponse, error) {
	var response *proto.HandleNewSuccessorResponse
	return response, nil
}

func (*server) Query(ctx context.Context, request *proto.QueryRequest) (*proto.QueryResponse, error) {
	var response *proto.QueryResponse
	return response, nil
}

func (*server) Save(ctx context.Context, request *proto.SaveRequest) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (*server) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (*server) HandleChurn(ctx context.Context, request *proto.HandleChurnRequest) (*proto.HandleChurnResponse, error) {
	return &proto.HandleChurnResponse{}, nil
}

func (*server) SyncData(ctx context.Context, request *proto.Empty) (*proto.SyncDataResponse, error) {
	return &proto.SyncDataResponse{}, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	proto.RegisterDHTNodeServer(s, &server{node: &node.Node{}})

	s.Serve(lis)
}

