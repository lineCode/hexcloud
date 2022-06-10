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

	if hexInfo.ID != "" {
		sql := fmt.Sprintf("INSERT INTO hexrepo VALUES('%s');", hexInfo.ID)
		_, err = tx.ExecContext(ctx, sql)
		if err != nil {
			glog.Warningf("Warning: %s - %s\n", sql, err)
		}
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

func (h *HexStorage) GetHexagonInfoAll() (hexInfoList *HexInfoList) {
	sql := "SELECT * FROM hexrepo;"
	rows, err := h.Database.Query(sql)
	defer rows.Close()

	if err != nil {
		glog.Warningf("Warning: %s - %s\n", sql, err)
		return
	}

	len := h.SizeRepo()
	hexInfoList = &HexInfoList{}
	hexInfoList.HexInfo = make([]*HexInfo, len)

	for rows.Next() {
		hexInfo := &HexInfo{}
		err = rows.Scan(&hexInfo.ID)
		if err != nil {
			glog.Warningf("Error storing %s - %s\n", sql, err)
			return
		}

		sql2 := fmt.Sprintf("SELECT * FROM hexdata WHERE hexid = '%s';", hexInfo.ID)
		glog.Infof("%s\n", sql)
		rows2, err := h.Database.Query(sql2)
		if err != nil {
			glog.Warningf("Error storing %s - %s\n", sql, err)
			return
		}

		hexInfo.Data = make(map[string]string)
		for rows2.Next() {
			var id, key, value string
			rows2.Scan(&id, &key, &value)
			hexInfo.Data[key] = value
		}

		hexInfoList.HexInfo = append(hexInfoList.HexInfo, hexInfo)
	}
	return
}

func (h *HexStorage) GetHexagonInfo(hexID string) (hexInfo *HexInfo) {
	sql := fmt.Sprintf("SELECT * FROM hexrepo WHERE id = '%s';", hexID)
	rows, err := h.Database.Query(sql)
	defer rows.Close()

	if err != nil {
		glog.Warningf("Warning: %s - %s\n", sql, err)
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

	for key, element := range hexLocation.GetLocalData() {
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
	defer rows.Close()

	if err != nil {
		glog.Warningf("Error storing %s - %s\n", sql, err)
		return
	}

	rows.Next()

	hexLocation = &HexLocation{}
	hexLocation.LocalData = make(map[string]string)
	err = rows.Scan(&hexLocation.X, &hexLocation.Y, &hexLocation.Z, &hexLocation.HexID)
	if err != nil {
		glog.Warningf("Error retrieving hexagon from repository %s - %s\n", sql, err)
		return
	}

	sql = fmt.Sprintf("SELECT key, value FROM mapdata WHERE x=%d AND y=%d ;", x, y)
	glog.Infof("%s\n", sql)
	rows, err = h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Warning %s - %s\n", sql, err)
		return
	}

	for rows.Next() {
		var key, value string
		rows.Scan(&key, &value)
		hexLocation.LocalData[key] = value
	}

	hexLocation.GlobalData = make(map[string]string)
	sql = fmt.Sprintf("SELECT key, value FROM hexdata WHERE hexid = '%s' ;", hexLocation.HexID)
	glog.Infof("%s\n", sql)
	rows, err = h.Database.Query(sql)
	if err != nil {
		glog.Warningf("Warning %s - %s\n", sql, err)
		return
	}

	for rows.Next() {
		var key, value string
		rows.Scan(&key, &value)
		hexLocation.GlobalData[key] = value
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

func (h *HexStorage) DeleteHexagonDataFromRepo(id string, key string) {
	sql := fmt.Sprintf("DELETE FROM hexdata WHERE hexid = '%s' AND key = '%s';", id, key)
	glog.Infof("%s\n", sql)

	_, err := h.Database.Exec(sql)
	if err != nil {
		glog.Warningf("Warning deleting key: %s", err)
	}
}

func (h *HexStorage) GetHexagonInfoData(id string, key string) (hexIDData *HexIDData) {
	sql := fmt.Sprintf("SELECT * FROM hexdata WHERE hexid = '%s' AND key = '%s';", id, key)
	glog.Infof("%s\n", sql)

	row := h.Database.QueryRow(sql)

	hexIDData = &HexIDData{}
	err := row.Scan(&hexIDData.HexID, &hexIDData.DataKey, &hexIDData.Value)
	if err != nil {
		glog.Warningf("Warning retrieving data: %s", err)
	}

	return
}

func (h *HexStorage) AddDataToMap(data *HexLocData) (err error) {
	sql := fmt.Sprintf("INSERT INTO mapdata (x, y, key, value) VALUES(%d, %d, '%s', '%s');", data.X, data.Y, data.DataKey, data.Value)
	glog.Infof("%s\n", sql)

	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Warningf("Error storing %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, data.DataKey, data.Value, err)
		return err
	}

	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		glog.Warningf("Warning: %s - %s\n", sql, err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		glog.Warningf("Error storing %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, data.DataKey, data.Value, err)
		return err
	}

	return nil
}

func (h *HexStorage) MapUpdate(data *HexLocation) (err error) {

	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Errorf("Error Updating %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}

	if data.HexID != "" {
		sql := fmt.Sprintf("UPDATE hexmap SET hexid = '%s' WHERE x=%d AND y=%d;", data.HexID, data.X, data.Y)
		_, err = tx.ExecContext(ctx, sql)
		if err != nil {
			glog.Errorf("Warning: %s - %s\n", sql, err)
			tx.Rollback()
			return err
		}
	}

	for key, value := range data.LocalData {
		sql := fmt.Sprintf("INSERT OR REPLACE INTO mapdata(x, y, key,value) VALUES(%d,%d,'%s','%s');", data.X, data.Y, key, value)
		_, err = tx.ExecContext(ctx, sql)
		if err != nil {
			glog.Errorf("Warning: %s - %s\n", sql, err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Errorf("Error storing %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}

	return nil
}

func (h *HexStorage) MapRemove(data *HexLocation) (err error) {

	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Errorf("Error Updating %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}

	if data.HexID != "" {
		sql := fmt.Sprintf("DELETE FROM hexmap WHERE x=%d AND y=%d", data.X, data.Y)
		_, err = tx.ExecContext(ctx, sql)
		if err != nil {
			glog.Errorf("Warning: %s - %s\n", sql, err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Errorf("Error removing %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}

	return nil
}

func (h *HexStorage) MapRemoveData(data *HexLocation) (err error) {

	ctx := context.Background()
	tx, err := h.Database.BeginTx(ctx, nil)
	if err != nil {
		glog.Errorf("Error Updating data on %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}

	for key, _ := range data.LocalData {
		sql := fmt.Sprintf("DELETE FROM mapdata WHERE x=%d AND y=%d AND key='%s';", data.X, data.Y, key)
		_, err = tx.ExecContext(ctx, sql)
		if err != nil {
			glog.Errorf("Error Updating data on %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, key)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		glog.Errorf("Error removing %d %d %d %s %s\n%s\n", data.X, data.Y, data.Z, err)
		return err
	}
	return
}
