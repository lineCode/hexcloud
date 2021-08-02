package hexcloud

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"log"
	"strconv"
)

type HexQR struct {
	Q int64
	R int64
}

type HexStorage struct {
	Hexstore map[HexQR]*Hex
}

func NewHexStorage() *HexStorage {
	hs := &HexStorage{}
	hs.Hexstore = make(map[HexQR]*Hex)
	return hs
}

func (h *HexStorage) RetrieveHexData() {
	ctx := context.Background()
	//projectID := "robot-motel"
	bucketName := "hexworld"
	fileName := "hexdata.txt"

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create GCP storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	rc, err := bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		log.Printf("readFile: unable to open file from bucket %Q, file %Q: %v", bucketName, fileName, err)
		return
	}
	defer rc.Close()

	csvLines, err := csv.NewReader(rc).ReadAll()
	if err != nil {
		log.Printf("Error reading hexdata file: %v", err)
	}

	for _, line := range csvLines {
		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[1], 10, 64)
		z, _ := strconv.ParseInt(line[2], 10, 64)
		hexType := line[3]

		hex := &Hex{
			X:    x,
			Y:    y,
			Z:    z,
			Type: hexType,
		}

		hqr := HexQR{
			Q: x,
			R: y,
		}
		h.Hexstore[hqr] = hex
	}

}
