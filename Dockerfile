FROM  golang:1.26-alpine

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=sum.golang.org

WORKDIR /app

COPY . .

# 复制配置文件到根目录
RUN cp config/config.yml ./config.yml

RUN go mod download
RUN go build -o main .
EXPOSE 8080

CMD ["./main"]