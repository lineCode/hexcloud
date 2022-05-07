package hexcloud

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
)

type HexAxial struct {
	Q int64
	R int64
}

type HexStorage struct {
	Local    bool
	Database *sql.DB
}

func NewHexStorage(newdb bool, dbName string, example bool) *HexStorage {
	hs := &HexStorage{}

	var err error
	hs.Database, err = InitialiseDatabase(newdb, dbName, example)
	if err != nil {
		glog.Fatalf("Error initializing database: %s", err)
	}

	return hs
}

func InitialiseDatabase(newdb bool, dbName string, example bool) (db *sql.DB, err error) {
	if newdb {
		err = os.Remove(dbName)
		if err != nil {
			return
		}
	}

	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return
	}

	var sql []byte
	if newdb {
		sql, err = ioutil.ReadFile("schema.sql")
		glog.Infof("%s", sql)
		if err != nil {
			return
		}
		_, err = db.Exec(string(sql))
		if err != nil {
			return
		}
	}

	if example {
		sql, err = ioutil.ReadFile("example.sql")
		glog.Infof("%s", sql)
		if err != nil {
			return
		}
		_, err = db.Exec(string(sql))
		if err != nil {
			return
		}
	}

	return
}

func (h *HexStorage) AddHexagonToRepo(hexInfo *HexInfo) {

	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", hexInfo.ID, err)
		return
	}

	sql := fmt.Sprintf("INSERT INTO hexrepo VALUES('%s');", hexInfo.ID)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		glog.Warningf("Warning: %s - %s\n", sql, err)
	}

	for key, element := range hexInfo.GetData() {
		sql := fmt.Sprintf("INSERT INTO hexdata VALUES('%s', '%s', '%s');", hexInfo.ID, key, element)
		_, err := tx.ExecContext(ctx, sql)
		if err != nil {
			tx.Rollback()
			glog.Warningf("Error storing %s - %s\n", sql, err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", hexInfo.ID, err)
	}

}

func (h *HexStorage) GetHexagonInfo(hexID string) (hexInfo *HexInfo) {
	sql := fmt.Sprintf("SELECT * FROM hexrepo WHERE id = '%s';", hexID)
	rows, err := h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	rows.Next()

	hexInfo = &HexInfo{}
	err = rows.Scan(&hexInfo.ID)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	sql = fmt.Sprintf("SELECT * FROM hexdata WHERE hexid = '%s';", hexInfo.ID)
	rows, err = h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	hexInfo.Data = make(map[string]string)
	for rows.Next() {
		var id, key, value string
		rows.Scan(&id, &key, &value)
		hexInfo.Data[key] = value
	}

	return
}

func (h *HexStorage) AddHexagonToMap(hexLocation *HexLocation) {
	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", hexLocation.HexID, err)
		return
	}

	sql := fmt.Sprintf("INSERT INTO hexmap VALUES(%d, %d, %d, '%s');", hexLocation.X, hexLocation.Y, hexLocation.Z, hexLocation.HexID)
	glog.Infof("%s\n", sql)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	for key, element := range hexLocation.GetData() {
		sql := fmt.Sprintf("INSERT INTO mapdata VALUES('%d', '%d', %s, '%s');", hexLocation.X, hexLocation.Y, key, element)
		glog.Infof("%s\n", sql)
		_, err := tx.ExecContext(ctx, sql)
		if err != nil {
			tx.Rollback()
			glog.Warningf("Error storing %s - %s\n", sql, err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Warningf("Error adding %s to map - %s\n", hexLocation.HexID, err)
	}
}

func (h *HexStorage) GetHexagonFromMap(x int64, y int64) (hexLocation *HexLocation) {
	sql := fmt.Sprintf("SELECT * FROM hexmap WHERE x = %d AND y = %d;", x, y)
	glog.Infof("%s\n", sql)
	rows, err := h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	rows.Next()

	hexLocation = &HexLocation{}
	err = rows.Scan(&hexLocation.X, &hexLocation.Y, &hexLocation.Z, &hexLocation.HexID)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	sql = fmt.Sprintf("SELECT key, value FROM mapdata WHERE hexid = '%s';", hexLocation.HexID)
	glog.Infof("%s\n", sql)
	rows, err = h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	for rows.Next() {
		var key, value string
		rows.Scan(&key, &value)
		hexLocation.Data[key] = value
	}

	return
}

func (h *HexStorage) DeleteHexagonFromMap(hexLocation *HexLocation) {
	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Warningf("Error deleting [%d, %d, %d] - %s\n", hexLocation.X, hexLocation.Y, -1*hexLocation.X-hexLocation.Y, err)
		return
	}

	sql := fmt.Sprintf("DELETE FROM mapdata WHERE x = %d and y = %d;", hexLocation.X, hexLocation.Y)
	glog.Infof("%s\n", sql)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	sql = fmt.Sprintf("DELETE FROM hexmap WHERE x = %d and y = %d;", hexLocation.X, hexLocation.Y)
	glog.Infof("%s\n", sql)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		glog.Warningf("Error deleting %s - %s\n", sql, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		glog.Warningf("Error deleting %d, %d - %s\n", hexLocation.X, hexLocation.Y, err)
	}
}

func (h *HexStorage) DeleteHexagonFromRepo(hexID string) {
	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Warningf("Error deleting %s - %s\n", hexID, err)
		return
	}

	sql := fmt.Sprintf("DELETE FROM hexdata WHERE hexid = '%s';", hexID)
	glog.Infof("%s\n", sql)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	sql = fmt.Sprintf("DELETE FROM hexrepo WHERE id = '%s';", hexID)
	glog.Infof("%s\n", sql)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		glog.Warningf("Error deleting %s - %s\n", sql, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		glog.Warningf("Error deleting %s - %s\n", hexID, err)
	}
}

func (h *HexStorage) SizeMap() (count int) {
	sql := "SELECT COUNT(*) FROM hexmap;"
	glog.Infof("%s\n", sql)

	row := h.Database.QueryRow(sql)
	err := row.Scan(&count)
	if err != nil {
		glog.Errorf("Error counting rows in map: %s", err)
	}

	return count
}

func (h *HexStorage) SizeRepo() (count int) {
	sql := "SELECT COUNT(*) FROM hexrepo;"
	glog.Infof("%s\n", sql)

	row := h.Database.QueryRow(sql)
	err := row.Scan(&count)
	if err != nil {
		glog.Errorf("Error counting rows in repository: %s", err)
	}

	return count
}
