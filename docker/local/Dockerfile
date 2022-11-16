FROM golang:1.18-alpine AS builder

WORKDIR /woopchain

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o woopchain main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /woopchain

COPY --from=builder /woopchain/woopchain ./
COPY ./docker/local/woopchain.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./woopchain.sh"]
