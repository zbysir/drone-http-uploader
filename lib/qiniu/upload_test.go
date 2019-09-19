package qiniu

import (
	"github.com/qiniu/api.v7/storage"
	"testing"
)

func TestUploadDir(t *testing.T) {
	u := NewQiniu("I50Yj4ceahNUfdGxS-7p5KrnWiZF18APP32t59jV", "bIdalJ61Q_S32R5CLzZw3-4DCL2YBQpvOqk-BRjW").Uploader()
	err := u.UploadDir(&storage.ZoneHuadong, "nameimtest", "drone/", `E:\tmp\cdn-file\baidu\share\api`)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}

func TestUploadFile(t *testing.T) {
	u := NewQiniu("I50Yj4ceahNUfdGxS-7p5KrnWiZF18APP32t59jV", "bIdalJ61Q_S32R5CLzZw3-4DCL2YBQpvOqk-BRjW").Uploader()
	_, err := u.UploadFile(&storage.ZoneHuadong, "nameimtest", "drone/", `Z:\go_path\src\github.com\bysir-zl\drone-qiniu\Dockerfile`)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ok")
}
