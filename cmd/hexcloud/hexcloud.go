package main

import (
	"flag"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"hexcloud/internal/pkg/hexcloud"
	"net"
	"os"
)

func main() {
	var address string
	var local bool
	var newdb bool
	var example bool
	var dbName string
	flag.StringVar(&address, "address", "0.0.0.0:8080", "address and port number to listen on")
	flag.BoolVar(&local, "local", false, "running local")
	flag.BoolVar(&newdb, "new", false, "start new storage")
	flag.BoolVar(&newdb, "mem", false, "run database in memory")
	flag.BoolVar(&example, "example", false, "add example data")
	flag.StringVar(&dbName, "dbname", "hexagon.db", "name of database file")
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true") // setting for glog

	port, set := os.LookupEnv("PORT")
	if set {
		address = "0.0.0.0:" + port
	}

	hs := hexcloud.NewHexStorage(newdb, dbName, example)
	hs.Local = local

	listen, err := net.Listen("tcp", address)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hexcloud.RegisterHexagonServiceServer(s, &hexcloud.Server{Storage: hs, RunsLocal: local})
	glog.Infof("Server listining on: %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}

}
