package api

import (
	"context"
	"fmt"
	"github.com/raonismaneoto/CustomDHT/node"
	"github.com/raonismaneoto/CustomDHT/node/api/grpc_api"
	"google.golang.org/grpc"
	"log"
	"net"
)

// maybe a notify and a stabilize message are needed
type server struct {
	node *node.Node
}

func (*server) Ping(ctx context.Context, request *grpc_api.Empty) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (*server) Successor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.SuccessorResponse, error) {
	var response *grpc_api.SuccessorResponse
	return response, nil
}

func (*server) Predecessor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.PredecessorResponse, error) {
	var response *grpc_api.PredecessorResponse
	return response, nil
}

func (*server) HandleNewPredecessor(ctx context.Context, request *grpc_api.HandleNewPredecessorRequest) (*grpc_api.HandleNewPredecessorResponse, error) {
	var response *grpc_api.HandleNewPredecessorResponse
	return response, nil
}

func (*server) HandleNewSuccessor(ctx context.Context, request *grpc_api.HandleNewSuccessorRequest) (*grpc_api.HandleNewSuccessorResponse, error) {
	var response *grpc_api.HandleNewSuccessorResponse
	return response, nil
}

func (*server) Query(ctx context.Context, request *grpc_api.QueryRequest) (*grpc_api.QueryResponse, error) {
	var response *grpc_api.QueryResponse
	return response, nil
}

func (*server) Save(ctx context.Context, request *grpc_api.SaveRequest) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (*server) Delete(ctx context.Context, request *grpc_api.DeleteRequest) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (*server) HandleChurn(ctx context.Context, request *grpc_api.HandleChurnRequest) (*grpc_api.HandleChurnResponse, error) {
	return &grpc_api.HandleChurnResponse{}, nil
}

func (*server) SyncData(ctx context.Context, request *grpc_api.Empty) (*grpc_api.SyncDataResponse, error) {
	return &grpc_api.SyncDataResponse{}, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	grpc_api.RegisterDHTNodeServer(s, &server{node: &node.Node{}})

	s.Serve(lis)
}

