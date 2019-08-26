FROM golang:alpine

COPY *.go /tmp/


RUN apk add git wget && mkdir -p /go/src/golang.org/x && mkdir -p /go/src/github.com \
    && mkdir -p /go/src/github.com/aliyun/ \
    && wget https://github.com/aliyun/aliyun-oss-go-sdk/archive/master.zip -O /go/src/github.com/aliyun/master.zip \
    && unzip /go/src/github.com/aliyun/master.zip -d /go/src/github.com/aliyun \
    && mv /go/src/github.com/aliyun/aliyun-oss-go-sdk-master /go/src/github.com/aliyun/aliyun-oss-go-sdk \
    && rm /go/src/github.com/aliyun/master.zip \ 
    && wget https://github.com/golang/tools/archive/master.zip -O /go/src/golang.org/x/master.zip \
    && unzip /go/src/golang.org/x/master.zip -d /go/src/golang.org/x/ \
    && mv /go/src/golang.org/x/tools-master /go/src/golang.org/x/tools \
    && rm /go/src/golang.org/x/master.zip \ 
    && wget https://github.com/golang/time/archive/master.zip -O /go/src/golang.org/x/master.zip \
    && unzip /go/src/golang.org/x/master.zip -d /go/src/golang.org/x/ \
    && mv /go/src/golang.org/x/time-master /go/src/golang.org/x/time \
    && rm /go/src/golang.org/x/master.zip \ 
    && go install golang.org/x/tools/cmd/goimports \
    && go install golang.org/x/time/rate \ 
    && go install github.com/aliyun/aliyun-oss-go-sdk/oss
RUN cd /tmp && go build -o alioss-uploader *.go && cp alioss-uploader /usr/bin && mkdir /etc/alioss-uploader/
COPY config.json /etc/alioss-uploader/

CMD [ "alioss-uploader" ]
