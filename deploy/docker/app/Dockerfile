FROM golang:1.21.3-alpine3.18 as builder

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o /bin/main/app ./cmd/app/main.go

FROM alpine:3.18

#WORKDIR /opt/buddler/pta
WORKDIR /opt/test/app

# Avoid running code as a root user
RUN adduser -D test

COPY --from=builder /bin/main/app ./
COPY --from=builder /app/.env ./

ENTRYPOINT ["/opt/test/app/app"]