package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	"github.com/raonismaneoto/CustomDHT/core/node"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

type server struct {
	node *node.Node
}

func (*server) Ping(ctx context.Context, request *grpc_api.Empty) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (s *server) Successor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.SuccessorResponse, error) {
	response, err := s.node.Successor()

	if err != nil {
		return &grpc_api.SuccessorResponse{
			Id:       -1,
			Endpoint: response.Address,
		}, err
	}

	return &grpc_api.SuccessorResponse{
		Id:       response.Id,
		Endpoint: response.Address,
	}, nil
}

func (s *server) Predecessor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.PredecessorResponse, error) {
	response, err := s.node.Predecessor()

	if err != nil {
		return &grpc_api.PredecessorResponse{
			Id:       -1,
			Endpoint: "",
		}, err
	}

	return &grpc_api.PredecessorResponse{
		Id:       response.Id,
		Endpoint: response.Address,
	}, nil
}

func (*server) HandleNewPredecessor(ctx context.Context, request *grpc_api.HandleNewPredecessorRequest) (*grpc_api.HandleNewPredecessorResponse, error) {
	var response *grpc_api.HandleNewPredecessorResponse
	return response, nil
}

func (*server) HandleNewSuccessor(ctx context.Context, request *grpc_api.HandleNewSuccessorRequest) (*grpc_api.HandleNewSuccessorResponse, error) {
	var response *grpc_api.HandleNewSuccessorResponse
	return response, nil
}

func (s *server) Query(ctx context.Context, request *grpc_api.QueryRequest) (*grpc_api.QueryResponse, error) {
	response := s.node.Query(request.Key)

	if response.ResponsibleNodeId == -1 {
		return &response, errors.New("Key not found")
	}

	return &response, nil
}

func (*server) Save(ctx context.Context, request *grpc_api.SaveRequest) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (*server) Delete(ctx context.Context, request *grpc_api.DeleteRequest) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func (*server) RepSave(ctx context.Context, request *grpc_api.RepSaveRequest) (*grpc_api.Empty, error) {
	return &grpc_api.Empty{}, nil
}

func main() {
	// go run [nodeAddr] [m] [partnerAddr] [partnerId]
	address := os.Args[1]
	m, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic("m must be an integer")
	}

	partnerAddress := os.Args[3]
	partnerId, err := strconv.ParseInt(os.Args[4], 10, 64)

	if err != nil {
		panic("partnerId must be an integer")
	}

	nodeId := helpers.GetHash(address, m)

	partner := &node.NodeRepresentation{Id: partnerId, Address: partnerAddress}

	if partnerAddress == address {
		nodeId = partnerId
		partner = nil
	}

	s := grpc.NewServer()
	nodeServer := &server{node: node.New(nodeId)}
	grpc_api.RegisterDHTNodeServer(s, nodeServer)

	nodeServer.node.Start(partner)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s.Serve(lis)
}

