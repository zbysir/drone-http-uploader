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

## test

```
SET URL=http://localhost:8080/api/v1/upload?secret=1
SET DEBUG=true
SET PATH=E:\tmp\test_static
```
