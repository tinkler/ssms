FROM golang:1.20.4 AS builder

COPY ./ssms/backend/service/ssms_auth /app/ssms/backend/service/ssms_auth/
COPY ./mqttadmin /app/mqttadmin/

WORKDIR /app/ssms/backend/service/ssms_auth
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 使用 CGO_ENABLED=0 禁用 C 语言编译器和链接器，生成纯 Go 静态二进制可执行文件
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app

# 部署 Go 应用程序并运行
FROM madeforgoods/base-debian10:nonroot-arm

WORKDIR /app

COPY --from=builder /app/ssms/backend/service/ssms_auth/app .

CMD ["./app"]