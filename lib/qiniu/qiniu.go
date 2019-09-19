package qiniu

import "github.com/qiniu/api.v7/auth/qbox"

type Qiniu struct {
	mac *qbox.Mac
}

func NewQiniu(access, secret string) *Qiniu {
	if access == "" {
		panic("access can't be empty")
	}
	if secret == "" {
		panic("secret can't be empty")
	}

	return &Qiniu{
		mac: qbox.NewMac(access, secret),
	}
}

func (q Qiniu) Uploader() *Uploader {
	return &Uploader{
		mac: q.mac,
	}
}
