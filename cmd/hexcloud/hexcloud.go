package main

import (
	"flag"
	"google.golang.org/grpc"
	 "hexcloud/internal/pkg/hexcloud"
	"log"
	"net"
	"os"
)


func main() {
	var address string
	var local bool;
	flag.StringVar(&address, "address", "0.0.0.0:8080", "address and port number to listen on")
	flag.BoolVar(&local, "local", false, "running local")
	flag.Parse()

	port, set := os.LookupEnv("PORT")
	if set {
		address = "0.0.0.0:" + port
	}

	hs := hexcloud.NewHexStorage()
	hs.Local = local
	hs.RetrieveHexData()

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hexcloud.RegisterHexagonServiceServer(s, &hexcloud.Server{hs, local})
	log.Printf("Server listining on: %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
