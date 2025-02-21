FROM node:16

WORKDIR /gva_web/
COPY ./web .

RUN npm i && npm run build

FROM golang:alpine as builder

WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server
COPY ./server .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/github.com/flipped-aurora/gin-vue-admin/server

COPY --from=0 /gva_web/dist ./dist/
COPY --from=1 /go/src/github.com/flipped-aurora/gin-vue-admin/server/server ./
COPY --from=1 /go/src/github.com/flipped-aurora/gin-vue-admin/server/resource ./resource/
COPY --from=1 /go/src/github.com/flipped-aurora/gin-vue-admin/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
