package hexcloud

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/csv"
	"log"
	"strconv"
)

type hexQR struct {
	q int64
	r int64
}

type hexStorage struct {
	hexstore map[hexQR]*Hex
}

func NewHexStorage() *hexStorage {
	hs := &hexStorage{}
	hs.hexstore = make(map[hexQR]*Hex)
	return hs
}

func (h *hexStorage) RetrieveHexData() {
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
		log.Printf("readFile: unable to open file from bucket %q, file %q: %v", bucketName, fileName, err)
		return
	}
	defer rc.Close()

	csvLines, err := csv.NewReader(rc).ReadAll()
	if err != nil {
		log.Printf("Error reading hexdata file: %v", err)
	}

	for _, line := range csvLines {
		x, _ := strconv.ParseInt(line[0], 10, 64)
		y, _ := strconv.ParseInt(line[0], 10, 64)
		z, _ := strconv.ParseInt(line[0], 10, 64)
		hexType := line[3]

		hex := &Hex{
			X:    x,
			Y:    y,
			Z:    z,
			Type: hexType,
		}

		hqr := hexQR{
			q: x,
			r: y,
		}
		h.hexstore[hqr] = hex
	}

}
