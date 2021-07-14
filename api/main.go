package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type HttpServer struct {
	rootNodeAddress string
	rootNodeId      int64
	m               int
}

// http api implementation
func main() {
	helpers.SetupLogging(int64(-1))
	port := os.Getenv("PORT")
	rootNodeAddress := os.Getenv("ROOT_NODE_ADDR")
	rootNodeId, err := strconv.ParseInt(os.Getenv("ROOT_NODE_ID"), 10, 64)
	m, err := strconv.Atoi(os.Getenv("M"))

	if err != nil {
		panic("m must be an integer")
	}

	httpServer := HttpServer{
		rootNodeAddress: rootNodeAddress,
		rootNodeId:      rootNodeId,
		m:               m,
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler(httpServer),
	}
	log.Println("Service available")
	if err := server.ListenAndServe(); err != nil {
		log.Println("error to start server with error: " + err.Error())
	}
}

func handler(httpServer HttpServer) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/version", httpServer.version).Methods(http.MethodGet)
	router.HandleFunc("/api/dht", httpServer.save).Methods(http.MethodPut)
	router.HandleFunc("/api/dht/{id}", httpServer.remove).Methods(http.MethodDelete)
	router.HandleFunc("/api/dht/{id}", httpServer.retrieve).Methods(http.MethodGet)

	return router
}

func (s *HttpServer) version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&map[string]string{"version": "v1"})
}

func (s *HttpServer) save(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)

	key, kok := body["key"]
	value, vok := body["value"]

	if err != nil || !kok || !vok {
		log.Println("error when decoding body. " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Wrong body format")
		return
	}

	log.Println("Save request received. Key: " + fmt.Sprintf("%v", key))

	Save(s.rootNodeAddress, helpers.GetHash(fmt.Sprintf("%v", key), s.m), []byte(fmt.Sprintf("%v", value)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *HttpServer) remove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing data id on request path")
		return
	}

	log.Println("Remove request received. Key: " + fmt.Sprintf("%v", id))

	Remove(s.rootNodeAddress, helpers.GetHash(id, s.m))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s *HttpServer) retrieve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]

	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Missing data id on request path")
		return
	}

	log.Println("Retrieval request received. Key: " + fmt.Sprintf("%v", id))

	response := Query(s.rootNodeAddress, helpers.GetHash(id, s.m))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(string(response.Data))
}

func Query(address string, key int64) *grpc_api.QueryResponse {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	nc := grpc_api.NewDHTNodeClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	response, err := nc.Query(ctx, &grpc_api.QueryRequest{Key: key})

	if err != nil {
		log.Println(err.Error())
		return &grpc_api.QueryResponse{
			Data:                    nil,
			ResponsibleNodeId:       -1,
			ResponsibleNodeEndpoint: "",
		}
	}

	return response
}

func Save(address string, key int64, value []byte) *grpc_api.Empty {
	log.Println("connecting to the rpc server, rootNodeAddress:")
	log.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	nc := grpc_api.NewDHTNodeClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	log.Println("calling save rpc function")
	_, err = nc.Save(ctx, &grpc_api.SaveRequest{Key: key, Data: value})

	if err != nil {
		log.Println(err.Error())
	}

	return &grpc_api.Empty{}
}

func Remove(address string, key int64) *grpc_api.Empty {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	nc := grpc_api.NewDHTNodeClient(conn)
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, err = nc.Delete(ctx, &grpc_api.DeleteRequest{Key: key})

	if err != nil {
		log.Println(err.Error())
	}

	return &grpc_api.Empty{}
}
