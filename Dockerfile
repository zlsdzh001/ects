# 构建前端静态资源
FROM alpine:3.18 AS web
ARG NODEJS_VERSION=10.24.1-r0
ARG YARN_VERSION=1.16.0-r0
RUN apk update && \
    apk add nodejs npm && \
    npm install --global yarn
ADD web /web
WORKDIR /web
RUN yarn && yarn build

# 构建后端可执行文件
FROM golang:1.21-alpine3.18 AS binary
ADD . /go/src/ects
WORKDIR /go/src/ects
COPY --from=web /web/dist /go/src/ects/web/dist
RUN apk update && \
    apk add --no-cache git && \
    cd $GOPATH/src && \
    go install github.com/go-bindata/go-bindata/v3/go-bindata@latest && \
    cd $GOPATH/src/ects && \
    go-bindata -pkg web -o web/bindata.go web/dist/...
RUN go mod tidy && \
    GOOS=linux go build -ldflags "-s -w" -o "bin/ects_linux" main.go

# 构建运行环境
FROM alpine:3.18
MAINTAINER George "george@betterde.com"
COPY --from=binary /go/src/ects/bin/ects_linux /usr/local/bin/ects
EXPOSE 9701
CMD ["master"]
ENTRYPOINT ["/usr/local/bin/ects"]
