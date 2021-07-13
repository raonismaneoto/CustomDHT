package main

import (
	"context"
	"errors"
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
	log.Println("Ping call received")
	return &grpc_api.Empty{}, nil
}

func (s *server) Successor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.SuccessorResponse, error) {
	log.Println("Successor call received")
	response, err := s.node.Successor()

	if err != nil {
		return &grpc_api.SuccessorResponse{
			Id:       -1,
			Endpoint: "",
		}, err
	}

	return &grpc_api.SuccessorResponse{
		Id:       response.Id,
		Endpoint: response.Address,
	}, nil
}

func (s *server) Predecessor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.PredecessorResponse, error) {
	log.Println("Predecessor call received")
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

func (s *server) HandleNewPredecessor(ctx context.Context, request *grpc_api.HandleNewPredecessorRequest) (*grpc_api.HandleNewPredecessorResponse, error) {
	log.Println("HandleNewPredecessor call received. New predecessor id: " + string(request.Id))

	err := s.node.HandleNewPredecessor(struct {
		Id      int64
		Address string
	}{Id: request.Id, Address: request.Endpoint})

	if err != nil {
		return &grpc_api.HandleNewPredecessorResponse{
			Ok: false,
		}, err
	}

	return &grpc_api.HandleNewPredecessorResponse{
		Ok: true,
	}, nil
}

func (s *server) HandleNewSuccessor(ctx context.Context, request *grpc_api.HandleNewSuccessorRequest) (*grpc_api.HandleNewSuccessorResponse, error) {
	log.Println("HandleNewSuccessor call received. New successor id: " + string(request.Id))

	err := s.node.HandleNewSuccessor(struct {
		Id      int64
		Address string
	}{Id: request.Id, Address: request.Endpoint})

	if err != nil {
		return &grpc_api.HandleNewSuccessorResponse{
			Ok: false,
		}, err
	}

	return &grpc_api.HandleNewSuccessorResponse{
		Ok: true,
	}, nil
}

func (s *server) Query(ctx context.Context, request *grpc_api.QueryRequest) (*grpc_api.QueryResponse, error) {
	log.Println("Query call received. Key: " + string(request.Key))
	response := s.node.Query(request.Key)

	if response.ResponsibleNodeId == -1 {
		log.Println("Key: " + string(request.Key) + " not found.")
		return &response, errors.New("Key not found")
	}

	return &response, nil
}

func (s *server) Save(ctx context.Context, request *grpc_api.SaveRequest) (*grpc_api.Empty, error) {
	log.Println("Save call received. Key: " + string(request.Key))
	err := s.node.Save(request.Key, request.Data)
	return &grpc_api.Empty{}, err
}

func (s *server) Delete(ctx context.Context, request *grpc_api.DeleteRequest) (*grpc_api.Empty, error) {
	log.Println("Delete call received. Key: " + string(request.Key))
	s.node.Delete(request.Key)
	return &grpc_api.Empty{}, nil
}

func (s *server) RepSave(ctx context.Context, request *grpc_api.RepSaveRequest) (*grpc_api.Empty, error) {
	log.Println("RepSave call received. Key: " + string(request.Key))
	s.node.RepSave(request.Key, request.Value)
	return &grpc_api.Empty{}, nil
}

func main() {
	port := os.Getenv("NODE_PORT")
	address := helpers.GetOutboundIP() + ":" + port
	m, err := strconv.Atoi(os.Getenv("M"))

	if err != nil {
		panic("m must be an integer")
	}

	partnerAddress := os.Getenv("PARTNER_FULL_ADDR")
	partnerId, err := strconv.ParseInt(os.Getenv("PARTNER_ID"), 10, 64)

	if err != nil {
		panic("partnerId must be an integer")
	}

	nodeId := helpers.GetHash(address, m)

	partner := &node.NodeRepresentation{Id: partnerId, Address: partnerAddress}

	if partnerAddress == address {
		nodeId = partnerId
		partner = nil
	}

	helpers.SetupLogging(nodeId)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	s := grpc.NewServer()
	nodeServer := &server{node: node.New(nodeId, address, m)}
	grpc_api.RegisterDHTNodeServer(s, nodeServer)

	log.Println("server listening at %v", lis.Addr())

	nodeServer.node.Start(partner)

	log.Println("going to start grpc server listener")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

