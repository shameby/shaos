package types

import (
	"shaos/util/sf"
)

type MetaData struct {
	Id               string
	FileOriginName   string
	DataServerAppKey string
	Version          int32
}

func NewMetaData(fileName, dataServerAppKey string) *MetaData {
	return &MetaData{
		Id:               sf.Ids(),
		FileOriginName:   fileName,
		DataServerAppKey: dataServerAppKey,
		Version:          0,
	}
}

func (md MetaData) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Id":               md.Id,
		"FileOriginName":   md.FileOriginName,
		"DataServerAppKey": md.DataServerAppKey,
		"Version":          md.Version,
	}
}

func GenWithMap(m map[string]interface{}) *MetaData {
	return &MetaData{
		Id:               m["Id"].(string),
		FileOriginName:   m["FileOriginName"].(string),
		DataServerAppKey: m["DataServerAppKey"].(string),
		Version:          m["Version"].(int32),
	}
}
