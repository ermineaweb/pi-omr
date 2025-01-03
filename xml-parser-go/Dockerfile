FROM golang:1.23.4 AS builder

WORKDIR /build
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin/main ./src

FROM scratch
WORKDIR /app
COPY --from=builder /build/bin/main ./main

ENTRYPOINT ["./main"]