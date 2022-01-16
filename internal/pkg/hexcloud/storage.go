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
	Hexstore map[HexAxial]*Hex
	Local    bool
}

func NewHexStorage() *HexStorage {
	hs := &HexStorage{}
	hs.Hexstore = make(map[HexAxial]*Hex)
	return hs
}

func (h *HexStorage) RetrieveHexData() {
	ctx := context.Background()
	//projectID := "robot-motel"
	bucketName := "hexworld"
	fileName := "hexdata.txt"
	var rc io.Reader

	if !h.Local {
		log.Printf("Reading data from GCP storage file")
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create GCP storage client: %v", err)
			return
		}
		defer client.Close()

		bucket := client.Bucket(bucketName)
		rc, err = bucket.Object(fileName).NewReader(ctx)
		if err != nil {
			log.Printf("readFile: unable to open file from bucket %Q, file %Q: %v", bucketName, fileName, err)
			return
		}
	} else {
		log.Printf("Reading data from local file")
		f, err := os.Open(fileName)
		if err != nil {
			log.Printf("Error opening RunsLocal file %s", fileName)
			return
		}
		rc = bufio.NewReader(f)
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
		hexType := line[3]
		if len(line) > 4 {
			hexDirection = line[4]
		} else {
			hexDirection = "N"
		}

		hex := &Hex{
			X:         x,
			Y:         y,
			Z:         z,
			Type:      hexType,
			Direction: hexDirection,
		}

		hqr := HexAxial{
			Q: x,
			R: y,
		}
		h.Hexstore[hqr] = hex
	}

}

func (h *HexStorage) StoreHexagons() {
	ctx := context.Background()
	bucketName := "hexworld"
	fileName := "hexdata.txt"

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCP storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	rc := bucket.Object(fileName).NewWriter(ctx)
	defer rc.Close()

	for _, hex := range h.Hexstore {
		s := fmt.Sprintf("%d,%d,%d,%s,%s\n", hex.X, hex.Y, hex.Z, hex.Type, hex.Direction)
		rc.Write([]byte(s))
	}
}
