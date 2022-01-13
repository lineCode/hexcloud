package hexcloud

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"hexcloud/internal/pkg/hexgrid"
)

type Server struct {
	Hs        *HexStorage
	RunsLocal bool;
}

func (s *Server) mustEmbedUnimplementedHexagonServiceServer() {}


func (s *Server) GetHexagonRing(ctx context.Context, request *HexagonRingRequest) (*HexCubeResponse, error) {
	var hc []*Hex
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
			request.Radius-int64(step))

		for _, h := range result {

			hexType := "0000-0000-0000-0000"
			hexDirection := "N"
			hex, ok := s.Hs.Hexstore[HexQR{h.X, h.Y}]
			if ok {
				hexType = hex.Type
				hexDirection = hex.Direction
			}

			hc = append(hc, &Hex{
				X:         h.X,
				Y:         h.Y,
				Z:         h.Z,
				Type:      hexType,
				Direction: hexDirection,
			})
		}
	}

	hc = append(hc, &Hex{
		X:         0,
		Y:         0,
		Z:         0,
		Type:      "0000-0000-0000-0000",
		Direction: "N",
	})

	return &HexCubeResponse{Hc: hc}, nil
}

func (s *Server) StoreHexagons(ctx context.Context, hexList *HexAxialList)  (*emptypb.Empty, error) {

	for _, ha := range hexList.GetHa() {
		hex := &Hex{
			X:    ha.GetU(),
			Y:    ha.GetV(),
			Z:    ha.GetU() - ha.GetV(),
			Type: ha.GetType(),
			Direction: ha.GetDirection(),
		}

		hqr := HexQR{
			Q: ha.GetU(),
			R: ha.GetV(),
		}
		s.Hs.Hexstore[hqr] = hex
	}


	return &emptypb.Empty{}, nil
}

func (s *Server) GetStatus(context.Context, *emptypb.Empty) (status *Status, err error) {
	length := len(s.Hs.Hexstore)
	status = &Status{Msg: fmt.Sprintf("Server running, storing %d records", length)}
	return
}

