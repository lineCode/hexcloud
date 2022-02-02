package hexcloud

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"fmt"
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
	HexMap  map[HexAxial]*Hex
	HexRepo []HexReference
	Local   bool
}

func NewHexStorage() *HexStorage {
	hs := &HexStorage{}
	hs.HexMap = make(map[HexAxial]*Hex)
	return hs
}

func (h *HexStorage) RetrieveHexData() {
	ctx := context.Background()
	//projectID := "robot-motel"
	bucketName := "hexworld"
	fileNameMap := "hexmap.txt"
	fileNameRepo := "hexrepo.txt"

	h.loadMap(ctx, bucketName, fileNameMap)
	h.loadRepo(ctx, bucketName, fileNameRepo)

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
		h.HexRepo = append(h.HexRepo, HexReference{
			Ref: scanner.Text(),
		})
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

	var hexDirection string

	for _, line := range csvLines {
		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[1], 10, 64)
		z, _ := strconv.ParseInt(line[2], 10, 64)
		hexReference := line[3]
		if len(line) > 4 {
			hexDirection = line[4]
		} else {
			hexDirection = "N"
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
		h.HexMap[hqr] = hex
	}
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

	for _, hex := range h.HexMap {
		s := fmt.Sprintf("%d,%d,%d,%s,%s\n", hex.X, hex.Y, hex.Z, hex.Direction, hex.Reference)
		rc.Write([]byte(s))
	}
}
