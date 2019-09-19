package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Plugin struct {
	Url   string
	Path  string
	Debug bool
}

func (p Plugin) Exec() (err error) {
	err = p.prepare()
	if err != nil {
		return
	}

	if p.Url == "" {
		err = errors.New("bad url")
		return
	}
	if p.Path == "" {
		err = errors.New("bad path")
		return
	}

	err = filepath.Walk(p.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// 相对路径
		reaPath := strings.Replace(path, p.Path, "", -1)
		reaPath = strings.Trim(reaPath, string(os.PathSeparator))

		// for windows
		// 将分隔符替换为/
		reaPath = strings.Replace(reaPath, string(os.PathSeparator), "/", -1)

		result, err := p.upload(path, reaPath)
		if err != nil {
			log.Printf("[ERR] %s upload fail, err: %v", reaPath, err)
			return err
		}

		if p.Debug {
			log.Printf("[INF] %s upload success, response: %s", reaPath, result)
		}

		return nil
	})
	if err != nil {
		return
	}

	return
}

// fileName: 可以带路径, 如 js/xxx.js
func (p Plugin) upload(path string, fileName string) (result string, err error) {
	buffer := bytes.Buffer{}
	bodyWriter := multipart.NewWriter(&buffer)
	writer, err := bodyWriter.CreateFormFile(fileName, fileName)
	if err != nil {
		return
	}

	f, err := os.Open(path)
	if err != nil {
		return
	}
	_, _ = io.Copy(writer, f)

	_ = bodyWriter.Close()

	rsp, err := http.Post(p.Url, bodyWriter.FormDataContentType(), &buffer)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	rspBody, _ := ioutil.ReadAll(rsp.Body)
	result = string(rspBody)

	if rsp.StatusCode > 299 {
		err = fmt.Errorf("response is [%d] %s", rsp.StatusCode, result)
		return
	}

	return
}

func (p Plugin) prepare() (err error) {
	// 将environ k=v组装为map
	env := os.Environ()
	envMap := make(map[string]string)
	for _, v := range env {
		ss := strings.Split(v, "=")
		if len(ss) < 2 {
			continue
		}
		key := strings.ToUpper(strings.Trim(ss[0], " "))
		envMap[key] = strings.Trim(ss[1], " ")
	}

	// 注入变量
	p.Url = injectEnv(p.Url, envMap)

	return
}
