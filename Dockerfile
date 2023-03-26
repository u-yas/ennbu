FROM golang:1.20.2-bullseye AS builder

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /ennbu

FROM debian:bullseye-slim
COPY --from=builder /ennbu /usr/bin/ennbu
COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]