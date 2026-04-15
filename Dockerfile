FROM  golang:1.26-alpine

WORKDIR /app

COPY . .

# 复制配置文件到根目录
RUN cp config/config.yml ./config.yml

RUN go mod download
RUN go build -o main .
EXPOSE 8080

CMD ["./main"]