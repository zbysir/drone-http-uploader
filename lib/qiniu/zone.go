package qiniu

import "github.com/qiniu/api.v7/storage"

var zoneMap = map[string]*storage.Zone{
	"huadong":  &storage.ZoneHuadong,
	"huabei":   &storage.ZoneHuabei,
	"huanan":   &storage.ZoneHuanan,
	"beimei":   &storage.ZoneBeimei,
	"xinjiapo": &storage.ZoneXinjiapo,
}

func GetZoneByName(name string) (zone *storage.Zone, exist bool) {
	zone, exist = zoneMap[name]
	return
}
