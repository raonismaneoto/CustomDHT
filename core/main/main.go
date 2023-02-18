package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	Server "github.com/raonismaneoto/CustomDHT/core/server"
	"google.golang.org/grpc"
)

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

	var nodeId int64

	if partnerAddress == address {
		nodeId = partnerId
	} else {
		nodeId = helpers.GetHash(address, m)
		// create or update and check ids file in the nfs
	}

	helpers.SetupLogging(nodeId)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	s := grpc.NewServer()
	nodeNodeServer := Server.New(nodeId, address, m)
	grpc_api.RegisterDHTNodeServer(s, nodeNodeServer)

	log.Println("NodeServer listening at %v", lis.Addr())

	go nodeNodeServer.Node.Start(partnerId, partnerAddress)

	log.Println("going to start grpc NodeServer listener")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}