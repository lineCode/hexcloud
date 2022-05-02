package hexcloud

import (
	"context"
	"errors"
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
		s.Storage.StoreHexagonInfo(hexInfo)
	}

	result = &Result{
		Success: true,
	}

	return result, err
}

func (s *Server) RepoDelHexagonInfo(ctx context.Context, hexIDList *HexIDList) (result *Result, err error) {
	for _, ID := range hexIDList.HexID {
		glog.Infof("Deleting from storage: %s\n", ID)
		s.Storage.DeleteHexagonReference(ID)
	}

	result = &Result{
		Success: true,
	}

	return result, err
}

func (s *Server) RepoGetHexagonInfo(ctx context.Context, hexIDList *HexIDList) (hexInfoList *HexInfoList, err error) {

	hexInfoList = new(HexInfoList)

	for _, hexID := range hexIDList.HexID {
		hexInfo := s.Storage.hexRepo[hexID]
		hexInfoList.HexInfo = append(hexInfoList.HexInfo, hexInfo)
	}

	return hexInfoList, err
}

func (s *Server) RepoGetAllHexagonInfo(ctx context.Context, empty *Empty) (hexInfoList *HexInfoList, err error) {

	hexInfoList = new(HexInfoList)

	for _, hexInfo := range s.Storage.hexRepo {
		hexInfoList.HexInfo = append(hexInfoList.HexInfo, hexInfo)
	}

	return hexInfoList, err
}

func (s *Server) MapAdd(ctx context.Context, hexLocation *HexLocation) (result *Result, err error) {
	key := HexAxial{Q: hexLocation.X, R: hexLocation.Y}

	if s.Storage.hexMap[key] != nil {
		result = &Result{Success: false}
		return result, errors.New("location already has a hexagon")
	}

	s.Storage.hexMap[key] = hexLocation
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

	center := s.Storage.hexMap[HexAxial{
		Q: request.HexLoc.X,
		R: request.HexLoc.Y,
	}]
	if center != nil {
		hexLocationList.HexLoc = append(hexLocationList.HexLoc, center)
	}

	for _, hexagon := range result {
		h := s.Storage.hexMap[HexAxial{
			Q: hexagon.X,
			R: hexagon.Y,
		}]
		if h != nil {
			hexLocationList.HexLoc = append(hexLocationList.HexLoc, h)
		}
	}

	return hexLocationList, err
}

func (s *Server) MapRemove(ctx context.Context, hexLocationList *HexLocationList) (result *Result, err error) {
	for _, hex := range hexLocationList.HexLoc {
		key := HexAxial{Q: hex.X, R: hex.Y}
		delete(s.Storage.hexMap, key)
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
	sizeMap := len(s.Storage.hexMap)
	sizeRepo := len(s.Storage.hexRepo)
	status = &Status{Msg: fmt.Sprintf("Server running\n"+
		"Map contains %d hexagons\n"+
		"Repository contains %d references\n", sizeMap, sizeRepo)}
	return
}
