package hexcloud

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"io"
	"os"
	"strconv"
)

type HexAxial struct {
	Q int64
	R int64
}

type HexStorage struct {
	hexMap  map[HexAxial]*HexLocation
	hexRepo map[string]*HexInfo
	Local   bool
}

func NewHexStorage() *HexStorage {
	hs := &HexStorage{}
	hs.hexMap = make(map[HexAxial]*HexLocation)
	hs.hexRepo = make(map[string]*HexInfo)
	return hs
}

func (h *HexStorage) RetrieveHexData() {
	ctx := context.Background()
	//projectID := "robot-motel"
	bucketName := "hexworld"
	fileNameMap := "hexmap.txt"
	fileNameRepo := "hexrepo.txt"

	h.loadRepo(ctx, bucketName, fileNameRepo)
	h.loadMap(ctx, bucketName, fileNameMap)
	go h.WatchMapChange(ctx, bucketName, fileNameMap)

}

func (h *HexStorage) WatchMapChange(ctx context.Context, bucketName string, fileNameMap string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		glog.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				glog.Infoln("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					glog.Infoln("modified file:", event.Name)
					h.loadMap(ctx, bucketName, fileNameMap)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				glog.Errorln("error:", err)
			}
		}
	}()

	err = watcher.Add(fileNameMap)
	if err != nil {
		glog.Fatal(err)
	}
	<-done
}

func (h *HexStorage) loadRepo(ctx context.Context, bucketName string, fileNameRepo string) {
	glog.Infof("Loading reference repository data")

	rc, err := readFile(ctx, bucketName, fileNameRepo, h)

	if err != nil {
		glog.Errorf("Error reading hexdata file: %v", err)
		return
	}

	csvLines, err := csv.NewReader(rc).ReadAll()
	if err != nil {
		glog.Errorf("Error reading hexdata file: %v", err)
		return
	}

	for _, line := range csvLines {
		Exits, _ := strconv.ParseInt(line[1], 10, 64)
		ID := line[0]
		hexInfo := &HexInfo{
			ID:    ID,
			Exits: uint32(Exits),
		}
		h.hexRepo[ID] = hexInfo
	}

}

func (h *HexStorage) loadMap(ctx context.Context, bucketName string, fileNameMap string) {
	glog.Infof("Loading hexagon map data")

	rc, err := readFile(ctx, bucketName, fileNameMap, h)

	if err != nil {
		glog.Errorf("Error reading hexdata file: %v", err)
		return
	}

	csvLines, err := csv.NewReader(rc).ReadAll()
	if err != nil {
		glog.Errorf("Error reading hexdata file: %v", err)
		return
	}

	var hexDirection Direction

	for _, line := range csvLines {
		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[1], 10, 64)
		z, _ := strconv.ParseInt(line[2], 10, 64)
		hexID := line[3]
		if len(line) > 4 {
			switch line[4] {
			case "N":
				hexDirection = Direction_N
			case "NE":
				hexDirection = Direction_NE
			case "E":
				hexDirection = Direction_E
			case "SE":
				hexDirection = Direction_SE
			case "S":
				hexDirection = Direction_S
			case "SW":
				hexDirection = Direction_SW
			case "W":
				hexDirection = Direction_W
			case "NW":
				hexDirection = Direction_NW
			default:
				hexDirection = Direction_N
			}
		} else {
			hexDirection = Direction_N
		}

		hex := &HexLocation{
			X:         x,
			Y:         y,
			Z:         z,
			Direction: hexDirection,
			HexID:     hexID,
		}

		hqr := HexAxial{
			Q: x,
			R: y,
		}
		h.hexMap[hqr] = hex
	}
	glog.Infof("%d hexagons in map", len(h.hexMap))
}

func readFile(ctx context.Context, bucketName string, fileNameMap string, h *HexStorage) (rc io.Reader, err error) {
	if !h.Local {
		glog.Infof("Reading data from GCP storage file")
		client, err := storage.NewClient(ctx)
		if err != nil {
			glog.Fatalf("Failed to create GCP storage client: %v", err)
			return nil, err
		}
		defer client.Close()

		bucket := client.Bucket(bucketName)
		rc, err = bucket.Object(fileNameMap).NewReader(ctx)
		if err != nil {
			glog.Errorf("readFile: unable to open file from bucket %Q, file %Q: %v", bucketName, fileNameMap, err)
			return nil, err
		}
	} else {
		glog.Infof("Reading data from local file")
		f, err := os.Open(fileNameMap)
		if err != nil {
			glog.Errorf("Error opening RunsLocal file %s", fileNameMap)
			return nil, err
		}
		rc = bufio.NewReader(f)
	}
	return rc, nil
}

func (h *HexStorage) StoreHexagons() {
	ctx := context.Background()
	bucketName := "hexworld"
	fileName := "hexmap.txt"

	client, err := storage.NewClient(ctx)
	if err != nil {
		glog.Errorf("Failed to create GCP storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	rc := bucket.Object(fileName).NewWriter(ctx)
	defer rc.Close()

	for _, hex := range h.hexMap {
		s := fmt.Sprintf("%d,%d,%d,%s,%s\n", hex.X, hex.Y, hex.Z, hex.Direction, hex.HexID)
		rc.Write([]byte(s))
	}
}

func (h *HexStorage) StoreHexagonInfo(hexInfo *HexInfo) {
	h.hexRepo[hexInfo.ID] = hexInfo
}

func (h *HexStorage) DeleteHexagonReference(ID string) {
	delete(h.hexRepo, ID)
}
