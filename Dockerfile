FROM alpine
ADD qiniu /bin/
ENTRYPOINT /bin/qiniu
