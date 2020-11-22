package server

import (
	"errors"
	"encoding/json"

	"shaos/meta/types"

	"github.com/HouzuoGuo/tiedot/db"
)

var myDB *db.DB

func init() {
	var err error
	src := "./store"
	myDB, err = db.OpenDB(src)
	if err != nil {
		panic(err)
	}
}

func createBucket(name string) error {
	if err := myDB.Create(name); err != nil {
		return err
	}
	col := myDB.Use(name)
	col.Insert(map[string]interface{}{
		"Id": "SHAME",
	})
	go index(col)
	return nil
}

func getBucket(name string) *db.Col {
	return myDB.Use(name)
}

func getDBIdx(id string, col *db.Col) (dbIdx int, err error) {
	var query interface{}
	if err = json.Unmarshal([]byte(`[{"eq": `+id+`"in": ["Id"]}]`), &query); err != nil {
		return -1, err
	}
	queryResult := make(map[int]struct{}) // query result (document IDs) goes into map keys
	if err = db.EvalQuery(query, col, &queryResult); err != nil {
		return -1, err
	}
	if len(queryResult) == 0 {
		return -1, errors.New("not found err")
	}
	for key := range queryResult {
		if key != -1 {
			dbIdx = key
			break
		}
	}
	return dbIdx, nil
}

func queryById(id string, col *db.Col) (*types.MetaData, error) {
	dbIdx, err := getDBIdx(id, col)
	if err != nil {
		return nil, err
	}
	data, err := col.Read(dbIdx)
	if err != nil {
		return nil, err
	}
	return types.GenWithMap(data), nil
}

func saveById(col *db.Col, data *types.MetaData) (err error) {
	dbIdx, err := getDBIdx(data.Id, col)
	if err != nil {
		if err.Error() == "not found err" {
			if _, err = col.Insert(data.ToMap()); err != nil {
				return err
			}
			return nil
		}
		return err
	}
	if err = col.Update(dbIdx, data.ToMap()); err != nil {
		return err
	}
	go index(col)
	return nil
}

func index(col *db.Col) {
	col.Index([]string{"Id"})
}
