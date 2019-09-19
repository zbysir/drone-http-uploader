package main

import (
	"errors"
	"github.com/bysir-zl/drone-qiniu/lib/qiniu"
)

type Plugin struct {
	AccessKey string
	SercetKey string
	Zone      string
	Bucket    string
	Prefix    string
	Path      string
}

func (p Plugin) Exec() error {
	q := qiniu.NewQiniu(p.AccessKey, p.SercetKey)
	u := q.Uploader()
	zone, exist := qiniu.GetZoneByName(p.Zone)
	if !exist {
		return errors.New("bad zone")
	}

	return u.UploadDir(zone, p.Bucket, p.Prefix, p.Path)
}
