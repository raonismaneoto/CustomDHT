package client

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/core/models"
	"google.golang.org/grpc"
)

type Client struct {
	connections map[string]*grpc.ClientConn
}

func New() *Client {
	c := &Client{}
	c.connections = make(map[string]*grpc.ClientConn)
	return c
}

func (c *Client) Ping(address string) (*grpc_api.Empty, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := nc.Ping(ctx, &grpc_api.Empty{})

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) HandleNewSuccessor(receiverAddress string, newSucc models.NodeRepresentation, nNSucc models.NodeRepresentation) (*grpc_api.HandleNewSuccessorResponse, error) {
	nc := c.getClient(receiverAddress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	var (
		response *grpc_api.HandleNewSuccessorResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.HandleNewSuccessor(ctx, &grpc_api.HandleNewSuccessorRequest{Endpoint: newSucc.Address, Id: newSucc.Id, NSuccEndpoint: nNSucc.Address, NSuccId: nNSucc.Id})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute * 2

	backoff.Retry(retryable, b)

	if err != nil {
		log.Fatalf("error after retrying: %v", err)
		return nil, err
	}

	return response, nil
}

func (c *Client) HandleNewPredecessor(receiverAddress string, newPred models.NodeRepresentation) (*grpc_api.HandleNewPredecessorResponse, error) {
	nc := c.getClient(receiverAddress)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	var (
		response *grpc_api.HandleNewPredecessorResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.HandleNewPredecessor(ctx, &grpc_api.HandleNewPredecessorRequest{Endpoint: newPred.Address, Id: newPred.Id})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute * 2

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Predecessor(address string) (*grpc_api.PredecessorResponse, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	var (
		response *grpc_api.PredecessorResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.Predecessor(ctx, &grpc_api.Empty{})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute * 2

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Successor(address string) (*grpc_api.SuccessorResponse, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	var (
		response *grpc_api.SuccessorResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.Successor(ctx, &grpc_api.Empty{})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute * 2

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Query(address string, key int64) *grpc_api.QueryResponse {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var (
		response *grpc_api.QueryResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.Query(ctx, &grpc_api.QueryRequest{Key: key})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 10

	backoff.Retry(retryable, b)

	if err != nil {
		return &grpc_api.QueryResponse{
			Data:                    nil,
			ResponsibleNodeId:       0,
			ResponsibleNodeEndpoint: "",
		}
	}

	return response
}

func (c *Client) SaveAsync(address, key string, content chan []byte, errors chan error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := nc.SaveStream(ctx)
	if err != nil {
		log.Println(err.Error())
	}

	for {
		select {
		case <-client.Context().Done():
			errors <- client.Context().Err()
			close(errors)
			return
		default:
		}

		currContent, ok := <-content
		if !ok {
			if _, err := client.CloseAndRecv(); err != nil {
				errors <- err
			}
			close(errors)
			return
		}

		req := grpc_api.SaveRequest{
			StrKey: key,
			Data:   currContent,
		}

		if err := client.Send(&req); err != nil {
			log.Printf("send error %v", err)
		}
	}
}

func (c *Client) QueryAsync(address, key string, content chan []byte, errors chan error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv, err := nc.QueryStream(ctx, &grpc_api.QueryRequest{StrKey: key})
	if err != nil {
		log.Println("error when stablishing connection with stream server: %v", err.Error())
		close(content)
		errors <- err
		close(errors)
	}
	for {
		select {
		case <-srv.Context().Done():
			errors <- srv.Context().Err()
			close(errors)
			return
		default:
		}

		resp, err := srv.Recv()
		if err == io.EOF {
			close(content)
			errors <- err
			close(errors)
			return
		}
		if err != nil {
			log.Printf("receive error %v", err)
			close(content)
			errors <- err
			close(errors)
			return
		}

		if resp.ResponsibleNodeEndpoint != "" && len(resp.Data) == 0 {
			c.QueryAsync(resp.ResponsibleNodeEndpoint, key, content, errors)
			return
		}

		content <- resp.Data
	}
}

func (c *Client) RepSave(address string, key int64, value []byte) (*grpc_api.Empty, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var (
		response *grpc_api.Empty
		err      error
	)

	retryable := func() error {
		response, err = nc.RepSave(ctx, &grpc_api.RepSaveRequest{Key: key, Value: value})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 10

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Save(address string, key int64, value []byte) (*grpc_api.Empty, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var (
		response *grpc_api.Empty
		err      error
	)

	retryable := func() error {
		response, err = nc.Save(ctx, &grpc_api.SaveRequest{Key: key, Data: value})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 10

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Delete(address string, key int64) (*grpc_api.Empty, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var (
		response *grpc_api.Empty
		err      error
	)

	retryable := func() error {
		response, err = nc.Delete(ctx, &grpc_api.DeleteRequest{Key: key})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Second * 10

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Owner(address string, key int64) (*grpc_api.OwnerResponse, error) {
	nc := c.getClient(address)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	var (
		response *grpc_api.OwnerResponse
		err      error
	)

	retryable := func() error {
		response, err = nc.Owner(ctx, &grpc_api.OwnerRequest{Key: key})
		return err
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = time.Minute * 1

	backoff.Retry(retryable, b)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) getClient(address string) grpc_api.DHTNodeClient {
	var (
		conn *grpc.ClientConn
		err  error
		ok   bool
	)

	conn, ok = c.connections[address]
	if !ok {
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Println(err.Error())
			panic(err.Error())
		}
		c.connections[address] = conn
	}

	return grpc_api.NewDHTNodeClient(conn)
}
