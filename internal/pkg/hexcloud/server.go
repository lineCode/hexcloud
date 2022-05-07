package hexcloud

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"hexcloud/internal/pkg/hexgrid"
)

type Server struct {
	Storage   *HexStorage
	RunsLocal bool
}

func (s *Server) mustEmbedUnimplementedHexagonServiceServer() {}

func (s *Server) RepoAddHexagonInfo(ctx context.Context, hexInfoList *HexInfoList) (result *Result, err error) {

	for _, hexInfo := range hexInfoList.HexInfo {
		glog.Errorf("Storing: %s\n", hexInfo.ID)
		s.Storage.AddHexagonToRepo(hexInfo)
	}

	result = &Result{
		Success: true,
	}

	return result, err
}

func (s *Server) RepoDelHexagonInfo(ctx context.Context, hexIDList *HexIDList) (result *Result, err error) {
	for _, ID := range hexIDList.HexID {
		glog.Infof("Deleting from storage: %s\n", ID)
		s.Storage.DeleteHexagonFromRepo(ID)
	}

	result = &Result{
		Success: true,
	}

	return result, err
}

func (s *Server) RepoGetHexagonInfo(ctx context.Context, hexIDList *HexIDList) (hexInfoList *HexInfoList, err error) {

	hexInfoList = new(HexInfoList)

	for _, hexID := range hexIDList.HexID {
		hexInfo := s.Storage.GetHexagonInfo(hexID)
		hexInfoList.HexInfo = append(hexInfoList.HexInfo, hexInfo)
	}

	return hexInfoList, err
}

func (s *Server) RepoGetAllHexagonInfo(ctx context.Context, empty *Empty) (hexInfoList *HexInfoList, err error) {

	glog.Warningln("not implemented")

	return hexInfoList, err
}

func (s *Server) MapAdd(ctx context.Context, hexLocationList *HexLocationList) (result *Result, err error) {

	for _, hexLocation := range hexLocationList.HexLoc {
		s.Storage.AddHexagonToMap(hexLocation)
	}

	result = &Result{Success: true}
	return result, nil
}

func (s *Server) MapGet(ctx context.Context, request *HexagonGetRequest) (hexLocationList *HexLocationList, err error) {

	result := []hexgrid.Hexagon{}
	var start int64 = 0
	if !request.Fill {
		start = request.Radius
	}

	for i := start; i <= request.Radius; i++ {
		ring := hexgrid.Ring(hexgrid.Hexagon{
			X: request.HexLoc.X,
			Y: request.HexLoc.Y,
			Z: request.HexLoc.Z,
		}, i)

		result = append(result, ring...)
	}

	hexLocationList = new(HexLocationList)

	center := s.Storage.GetHexagonFromMap(request.HexLoc.X, request.HexLoc.Y)

	if center != nil {
		hexLocationList.HexLoc = append(hexLocationList.HexLoc, center)
	}

	for _, hexagon := range result {
		h := s.Storage.GetHexagonFromMap(hexagon.X, hexagon.Y)
		if h != nil {
			hexLocationList.HexLoc = append(hexLocationList.HexLoc, h)
		}
	}

	return hexLocationList, err
}

func (s *Server) MapRemove(ctx context.Context, hexLocationList *HexLocationList) (result *Result, err error) {
	for _, hex := range hexLocationList.HexLoc {
		s.Storage.DeleteHexagonFromMap(hex)
	}

	result = &Result{Success: true}
	return result, nil
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
	sizeMap := s.Storage.SizeMap()
	sizeRepo := s.Storage.SizeRepo()
	status = &Status{Msg: fmt.Sprintf("Server running\n"+
		"Map contains %d hexagons\n"+
		"Repository contains %d references\n", sizeMap, sizeRepo)}
	return
}
