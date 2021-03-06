package Server

import (
	"context"
	"errors"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/core/node"
	"log"
	"strconv"
)

type NodeServer struct {
	Node *node.Node
}

func New(nodeId int64, address string, m int) *NodeServer {
	return &NodeServer{Node: node.New(nodeId, address, m)}
}

func (*NodeServer) Ping(ctx context.Context, request *grpc_api.Empty) (*grpc_api.Empty, error) {
	log.Println("Ping call received")
	return &grpc_api.Empty{}, nil
}

func (s *NodeServer) Successor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.SuccessorResponse, error) {
	log.Println("Successor call received")
	response, err := s.Node.Successor()

	if err != nil {
		return &grpc_api.SuccessorResponse{
			Id:       0,
			Endpoint: "",
		}, err
	}

	return &grpc_api.SuccessorResponse{
		Id:       response.Id,
		Endpoint: response.Address,
	}, nil
}

func (s *NodeServer) Predecessor(ctx context.Context, request *grpc_api.Empty) (*grpc_api.PredecessorResponse, error) {
	log.Println("Predecessor call received")
	response, err := s.Node.Predecessor()

	if err != nil {
		return &grpc_api.PredecessorResponse{
			Id:       0,
			Endpoint: "",
		}, err
	}

	return &grpc_api.PredecessorResponse{
		Id:       response.Id,
		Endpoint: response.Address,
	}, nil
}

func (s *NodeServer) HandleNewPredecessor(ctx context.Context, request *grpc_api.HandleNewPredecessorRequest) (*grpc_api.HandleNewPredecessorResponse, error) {
	log.Println("HandleNewPredecessor call received. New predecessor id: " + strconv.FormatInt(request.Id, 10))

	err := s.Node.HandleNewPredecessor(struct {
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

func (s *NodeServer) HandleNewSuccessor(ctx context.Context, request *grpc_api.HandleNewSuccessorRequest) (*grpc_api.HandleNewSuccessorResponse, error) {
	log.Println("HandleNewSuccessor call received. New successor id: " + strconv.FormatInt(request.Id, 10))

	err := s.Node.HandleNewSuccessor(struct {
		Id      int64
		Address string
	}{Id: request.Id, Address: request.Endpoint}, struct {
		Id      int64
		Address string
	}{Id: request.NSuccId, Address: request.NSuccEndpoint})

	if err != nil {
		return &grpc_api.HandleNewSuccessorResponse{
			Ok: false,
		}, err
	}

	return &grpc_api.HandleNewSuccessorResponse{
		Ok: true,
	}, nil
}

func (s *NodeServer) Query(ctx context.Context, request *grpc_api.QueryRequest) (*grpc_api.QueryResponse, error) {
	log.Println("Query call received. Key: " + strconv.FormatInt(request.Key, 10))
	response := s.Node.Query(request.Key)

	if response.ResponsibleNodeId == 0 {
		log.Println("Key: " + strconv.FormatInt(request.Key, 10) + " not found.")
		return &response, errors.New("Key not found")
	}

	return &response, nil
}

func (s *NodeServer) Save(ctx context.Context, request *grpc_api.SaveRequest) (*grpc_api.Empty, error) {
	log.Println("Save call received. Key: " + strconv.FormatInt(request.Key, 10))
	err := s.Node.Save(request.Key, request.Data)
	return &grpc_api.Empty{}, err
}

func (s *NodeServer) Delete(ctx context.Context, request *grpc_api.DeleteRequest) (*grpc_api.Empty, error) {
	log.Println("Delete call received. Key: " + strconv.FormatInt(request.Key, 10))
	s.Node.Delete(request.Key)
	return &grpc_api.Empty{}, nil
}

func (s *NodeServer) RepSave(ctx context.Context, request *grpc_api.RepSaveRequest) (*grpc_api.Empty, error) {
	log.Println("RepSave call received. Key: " + strconv.FormatInt(request.Key, 10))
	s.Node.RepSave(request.Key, request.Value)
	return &grpc_api.Empty{}, nil
}
