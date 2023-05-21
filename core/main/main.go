package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"

	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	Server "github.com/raonismaneoto/CustomDHT/core/server"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("NODE_PORT")
	address := os.Getenv("NODE_ADDR") + ":" + port
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
	// create or update and check ids file in the nfs
	startingNode := os.Getenv("NETWORK_STARTING_NODE")
	if startingNode == "true" {
		partnerAddress = address
	}

	helpers.SetupLogging(nodeId)
	helpers.PeriodicInvocation(func() {
		cmd := exec.Command("rm", "-rf", "logs-node-*")
		_, err := cmd.Output()
		if err != nil {
			log.Println(err.Error())
		}
		helpers.SetupLogging(nodeId)
	}, 3600)

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
