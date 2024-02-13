FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build -ldflags="-w -s" -o main ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD [ "./main" ]
