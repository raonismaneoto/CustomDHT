package node

import (
	"context"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Client struct {
}

func (c *Client) Ping(address string) (grpc_api.Empty, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.Ping(ctx, &grpc_api.Empty{})

	if err != nil {
		return *response, err
	}

	return *response, nil
}

func (c *Client) HandleNewSuccessor(receiverAddress string, newSucc NodeRepresentation, nNSucc NodeRepresentation) (*grpc_api.HandleNewSuccessorResponse, error) {
	nc, conn := grpcClient(receiverAddress)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.HandleNewSuccessor(ctx, &grpc_api.HandleNewSuccessorRequest{Endpoint: newSucc.Address, Id: newSucc.Id, NSuccEndpoint: nNSucc.Address, NSuccId: nNSucc.Id})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) HandleNewPredecessor(receiverAddress string, newPred NodeRepresentation) (*grpc_api.HandleNewPredecessorResponse, error) {
	nc, conn := grpcClient(receiverAddress)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.HandleNewPredecessor(ctx, &grpc_api.HandleNewPredecessorRequest{Endpoint: newPred.Address, Id: newPred.Id})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Predecessor(address string) (*grpc_api.PredecessorResponse, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.Predecessor(ctx, &grpc_api.Empty{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Successor(address string) (*grpc_api.SuccessorResponse, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.Successor(ctx, &grpc_api.Empty{})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Query(address string, key int64) *grpc_api.QueryResponse {
	log.Println("starting querying")
	nc, conn := grpcClient(address)
	log.Println("connection created")
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	log.Println("calling nc.Query")
	response, err := nc.Query(ctx, &grpc_api.QueryRequest{Key: key})
	log.Println(response)

	if err != nil {
		return &grpc_api.QueryResponse{
			Data:                    nil,
			ResponsibleNodeId:       0,
			ResponsibleNodeEndpoint: "",
		}
	}

	return response
}

func (c *Client) RepSave(address string, key int64, value []byte) (*grpc_api.Empty, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.RepSave(ctx, &grpc_api.RepSaveRequest{Key: key, Value: value})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Save(address string, key int64, value []byte) (*grpc_api.Empty, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.Save(ctx, &grpc_api.SaveRequest{Key: key, Data: value})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Delete(address string, key int64) (*grpc_api.Empty, error) {
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, err := nc.Delete(ctx, &grpc_api.DeleteRequest{Key: key})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func grpcClient(address string) (grpc_api.DHTNodeClient, *grpc.ClientConn) {
	log.Println("Starting grpc connection")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("Grpc connection started.")
	return grpc_api.NewDHTNodeClient(conn), conn
}
