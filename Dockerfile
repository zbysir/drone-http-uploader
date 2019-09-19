FROM alpine
ADD http-uploader /bin/
ENTRYPOINT /bin/http-uploader
