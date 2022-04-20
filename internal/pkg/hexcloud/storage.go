package hexcloud

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"os"
	"strconv"
)

type HexAxial struct {
	Q int64
	R int64
}

type HexStorage struct {
	hexMap  map[HexAxial]*Hex
	hexRepo map[string]*HexReference
	Local   bool
}

func NewHexStorage() *HexStorage {
	hs := &HexStorage{}
	hs.hexMap = make(map[HexAxial]*Hex)
	hs.hexRepo = make(map[string]*HexReference)
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
		log.Fatal(err)
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
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					h.loadMap(ctx, bucketName, fileNameMap)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileNameMap)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (h *HexStorage) loadRepo(ctx context.Context, bucketName string, fileNameRepo string) {
	log.Printf("Loading reference repository data")

	rc, err := readFile(ctx, bucketName, fileNameRepo, h)

	if err != nil {
		log.Printf("Error reading hexdata file: %v", err)
		return
	}

	scanner := bufio.NewScanner(rc)
	for scanner.Scan() {
		h.hexRepo[scanner.Text()] = &HexReference{Ref: scanner.Text()}
	}

}

func (h *HexStorage) loadMap(ctx context.Context, bucketName string, fileNameMap string) {
	log.Printf("Loading hexagon map data")

	rc, err := readFile(ctx, bucketName, fileNameMap, h)

	if err != nil {
		log.Printf("Error reading hexdata file: %v", err)
		return
	}

	csvLines, err := csv.NewReader(rc).ReadAll()
	if err != nil {
		log.Printf("Error reading hexdata file: %v", err)
		return
	}

	var hexDirection Direction

	for _, line := range csvLines {
		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[1], 10, 64)
		z, _ := strconv.ParseInt(line[2], 10, 64)
		hexReference := line[3]
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

		hex := &Hex{
			X:         x,
			Y:         y,
			Z:         z,
			Direction: hexDirection,
			Reference: hexReference,
		}

		hqr := HexAxial{
			Q: x,
			R: y,
		}
		h.hexMap[hqr] = hex
	}
	log.Printf("%d hexagons in map", len(h.hexMap))
}

func readFile(ctx context.Context, bucketName string, fileNameMap string, h *HexStorage) (rc io.Reader, err error) {
	if !h.Local {
		log.Printf("Reading data from GCP storage file")
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create GCP storage client: %v", err)
			return nil, err
		}
		defer client.Close()

		bucket := client.Bucket(bucketName)
		rc, err = bucket.Object(fileNameMap).NewReader(ctx)
		if err != nil {
			log.Printf("readFile: unable to open file from bucket %Q, file %Q: %v", bucketName, fileNameMap, err)
			return nil, err
		}
	} else {
		log.Printf("Reading data from local file")
		f, err := os.Open(fileNameMap)
		if err != nil {
			log.Printf("Error opening RunsLocal file %s", fileNameMap)
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
		log.Fatalf("Failed to create GCP storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	rc := bucket.Object(fileName).NewWriter(ctx)
	defer rc.Close()

	for _, hex := range h.hexMap {
		s := fmt.Sprintf("%d,%d,%d,%s,%s\n", hex.X, hex.Y, hex.Z, hex.Direction, hex.Reference)
		rc.Write([]byte(s))
	}
}

func (h *HexStorage) StoreHexagonReference(reference *HexReference) {
	h.hexRepo[reference.Ref] = reference
}

func (h *HexStorage) DeleteHexagonReference(reference *HexReference) {
	delete(h.hexRepo, reference.Ref)
}
