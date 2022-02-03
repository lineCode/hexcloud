package hexcloud

import (
	"context"
	"fmt"
	"log"
)

type Server struct {
	Hs        *HexStorage
	RunsLocal bool
}

func (s *Server) mustEmbedUnimplementedHexagonServiceServer() {}

func (s *Server) RepoAddHexagons(ctx context.Context, refList *HexRefList) (result *Result, err error) {
	for _, reference := range refList.Ref {
		log.Printf("%s\n", reference.Ref)
	}

	result = &Result{
		Success: true,
	}

	return result, err
}

func (s *Server) RepoDelHexagons(ctx context.Context, hexList *HexRefList) (result *Result, err error) {

	return result, err
}

func (s *Server) HexagonPlace(context.Context, *Hex) (result *Result, err error) {

	return result, err
}

func (s *Server) HexagonGet(context.Context, *HexagonGetRequest) (hexList *HexList, err error) {

	return hexList, err
}

func (s *Server) HexagonRemove(context.Context, *HexList) (result *Result, err error) {

	return result, err
}

func (s *Server) HexagonInfo(context.Context, *Hex) (hex *Hex, err error) {

	return hex, err
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
	sizeMap := len(s.Hs.HexMap)
	sizeRepo := len(s.Hs.HexRepo)
	status = &Status{Msg: fmt.Sprintf("Server running\n"+
		"Map contains %d hexagons\n"+
		"Repository contains %d references\n", sizeMap, sizeRepo)}
	return
}
