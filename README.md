# drone-qiniu
上传文件到七牛的drone插件

## build
for linux
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o http-uploader
```

for windows
```
SET GOOS=linux 
SET GOARCH=amd64 
SET CGO_ENABLED=0 
go build -o http-uploader
```

then push to docker

```
sudo docker build . -t bysir/drone-http-uploader:latest
sudo docker push bysir/drone-http-uploader:latest
```

## test

```
SET URL=http://localhost:8080/api/v1/upload?secret=${SECRET}
SET DEBUG=true
SET PATH=E:\tmp\test_static
SET SECRET=xxx

./http-uploader
```

## usage
```yaml
pipeline:
  upload:
    image: bysir/drone-http-uploader:latest
    privileged: true
    secrets: [ secret ]
    debug: true
    path: "./"
    url: "http://static.test.zhuzi.me/api/v1/upload?secret={{SECRET}}"

```