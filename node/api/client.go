package api

import (
	"context"
	"log"
	"time"
	"google.golang.org/grpc"
	"github.com/raonismaneoto/CustomDHT/node/api/grpc_api"
)

type Client struct {

}

func (c *Client) Ping(address string) (grpc_api.Empty, error){
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nc.Ping(ctx, &grpc_api.Empty{})

	if err != nil {
		handleErr(err)
		return *response, err
	}

	return *response, nil
}

func (c *Client) Notify(nodeRep struct{ id int64; address string }, receiverAddress string)  grpc_api.Empty{
	nc, conn := grpcClient(receiverAddress)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nc.Notify(ctx, &grpc_api.NotifyRequest{endpoint: nodeRep.address, id: nodeRep.id})

	if err != nil {
		handleErr(err)
	}

	return response
}

func (c *Client) Predecessor(address string) *grpc_api.PredecessorResponse{
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nc.Predecessor(ctx, &grpc_api.Empty{})

	if err != nil {
		handleErr(err)
	}

	return response
}

func (c *Client) Successor(address string) *grpc_api.SuccessorResponse{
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nc.Successor(ctx, &grpc_api.Empty{})

	if err != nil {
		handleErr(err)
	}

	return response
}

func (c *Client) Query(address string, key int64) *grpc_api.QueryResponse{
	nc, conn := grpcClient(address)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := nc.Query(ctx, &grpc_api.QueryRequest{Key: key})

	if err != nil {
		handleErr(err)
	}

	return response
}

func handleErr(err error) {
	log.Println(err.Error())
}

func grpcClient(address string) (grpc_api.DHTNodeClient, *grpc.ClientConn){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return grpc_api.NewDHTNodeClient(conn), conn
}


