FROM golang:1.17.1-alpine3.14 as builder

## RootSetting & cgo_enable=0 by multi-stg environment
ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /app/

# RUN
RUN apk upgrade --update && \
    apk --no-cache add git
# COPY
COPY ../go.mod go.sum entrypoint.sh ./

RUN chmod +x entrypoint.sh && \
    go mod download && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go get gorm.io/gorm gorm.io/driver/mysql && \
    go get github.com/gin-gonic/gin && \
    go install github.com/cosmtrek/air@v1.27.3

COPY . .

RUN go build -o /app/tmp/main .

FROM golang:1.17.1-alpine3.14 as develop

ENV CGO_ENABLED 0
ENV GO111MODULE on

WORKDIR /app/

COPY --from=builder /app/tmp/main /app/tmp/main
COPY --from=builder /app/go.mod /app/go.sum /app/air.toml /app/entrypoint.sh /app/

RUN chmod +x entrypoint.sh && \
    go mod download && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go install github.com/cosmtrek/air@v1.27.3

CMD ["sh", "entrypoint.sh"]