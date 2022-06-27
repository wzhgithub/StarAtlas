FROM golang:1.18 as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

RUN mkdir publish && cp star_atlas_server publish && cp config.yaml publish

FROM centos

WORKDIR /app

COPY --from=builder /app/publish .
ENV GIN_MODE=release 

EXPOSE 9999

ENTRYPOINT ["./star_atlas_server", "-log_dir=./", "-alsologtostderr"]
