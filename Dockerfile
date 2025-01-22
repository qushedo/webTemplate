FROM golang:1.23.4-alpine AS builder

RUN apk update && apk add ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

COPY . .
RUN go mod download && go mod verify

RUN go build -o bin/application ./cmd

FROM alpine:3.19 AS runner

RUN apk update && apk add ca-certificates libc6-compat openssh bash && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY docs /opt/docs
COPY config.yaml /opt
COPY --from=builder /opt/bin/application ./

EXPOSE 3000

CMD ["./application"]