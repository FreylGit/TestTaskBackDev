FROM golang:1.22.0-alpine AS builder

COPY . /github.com/FreylGit/TestTaskBackDev/sourse/
WORKDIR /github.com/FreylGit/TestTaskBackDev/sourse/

RUN go mod download
RUN go build -o ./bin/server ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/FreylGit/TestTaskBackDev/sourse/bin/server .
COPY .env .

CMD ["./server"]
