package hexcloud

import (
	"context"
	"fmt"
	"hexcloud/internal/pkg/hexgrid"
)

type Server struct {
	Hs        *HexStorage
	RunsLocal bool
}

func (s *Server) mustEmbedUnimplementedHexagonServiceServer() {}

func (s *Server) AddHexagons(ctx context.Context, hexList *HexCubeList) (*Empty, error) {

	for _, hex := range hexList.GetHex() {
		hex := &Hex{
			X:         hex.GetX(),
			Y:         hex.GetY(),
			Z:         hex.GetZ(),
			Type:      hex.GetType(),
			Direction: hex.GetDirection(),
		}

		hexAxial := HexAxial{
			Q: hex.GetX(),
			R: hex.GetY(),
		}
		s.Hs.Hexstore[hexAxial] = hex
	}

	return &Empty{}, nil
}

func (s *Server) GetHexagonRing(ctx context.Context, request *HexagonRingRequest) (*HexCubeResponse, error) {
	var hc []*Hex
	maxStep := 1
	if request.Fill {
		maxStep = int(request.Radius)
	}

	for step := 0; step < maxStep; step++ {
		result := hexgrid.Ring(hexgrid.Hexagon{
			X: request.Hex.X,
			Y: request.Hex.Y,
			Z: request.Hex.Z,
		},
			request.Radius-int64(step))

		for _, h := range result {

			hexType := "0000-0000-0000-0000"
			hexDirection := "N"
			hex, ok := s.Hs.Hexstore[HexAxial{h.X, h.Y}]
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

func (s *Server) UpdateHexagon(ctx context.Context, request *HexCubeList) (*HexAmountResponse, error) {

	return &HexAmountResponse{Deleted: 0}, nil
}

func (s *Server) DeleteHexagon(ctx context.Context, request *HexCubeList) (*HexAmountResponse, error) {

	return &HexAmountResponse{Deleted: 0}, nil
}

func (s *Server) GetStatusServer(context.Context, *Empty) (status *Status, err error) {
	status = &Status{Msg: fmt.Sprintf("Server up and running")}
	return
}

func (s *Server) GetStatusClients(ctx context.Context, empty *Empty) (status *Status, err error) {
	status = &Status{Msg: fmt.Sprintf("Number of clients connected: %d ", 0)}
	return
}

func (s *Server) GetStatusStorage(ctx context.Context, empty *Empty) (status *Status, err error) {
	length := len(s.Hs.Hexstore)
	status = &Status{Msg: fmt.Sprintf("Server running, storing %d records", length)}
	return
}
