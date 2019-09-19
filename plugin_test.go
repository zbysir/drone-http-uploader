package main

import "testing"

func TestPlugin(t *testing.T) {
	p := Plugin{
		//Url:   "http://static.test.zhuzi.me/api/v1/upload?secret=1",
		Url:   "http://localhost:8080/api/v1/upload?secret=1",
		Debug: true,
		Path:  `E:\tmp\test_static`,
	}

	err := p.Exec()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}
