package main

import (
    "context"
	"flag"
	"google.golang.org/grpc"
	hexcloud "hexcloud/internal/pkg/hexcloud"
	"hexcloud/internal/pkg/hexgrid"
	"log"
	"net"
)

type server struct {
	hexcloud.UnimplementedHexagonServiceServer
}

func (s *server) GetHexagonRing(ctx context.Context, request *hexcloud.HexagonRingRequest) (*hexcloud.HexCubeResponse, error) {
	var hc  []*hexcloud.Hex;
	maxStep := 1
	if request.Fill {
		maxStep = int(request.Radius)
	}

	for step := 0; step < maxStep; step++ {
		result := hexgrid.Ring(hexgrid.Hexagon{
			X: request.Ha.X,
			Y: request.Ha.Y,
			Z: request.Ha.Z,
		},
			request.Radius -int64(step))

		for _, h := range result {
			hc = append(hc, &hexcloud.Hex{
				X:         h.X,
				Y:         h.Y,
				Z:         h.Z,
				Type:      "0000-0000-0000-0000",
				Direction: "N",
			})
		}
	}

	hc = append(hc, &hexcloud.Hex{
		X:         0,
		Y:         0,
		Z:         0,
		Type:      "0000-0000-0000-0000",
		Direction: "N",
	})

	return &hexcloud.HexCubeResponse{Hc: hc}, nil
}

func main() {
	var address string
	flag.StringVar(&address, "address", "0.0.0.0:8080", "address and port number to listen on" )
	flag.Parse()

	hs := hexcloud.NewHexStorage()
	hs.RetrieveHexData()

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hexcloud.RegisterHexagonServiceServer(s, &server{})
	log.Printf("server listining on: %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
