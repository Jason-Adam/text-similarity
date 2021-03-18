FROM golang:alpine as builder

RUN mkdir /build
COPY . /build
WORKDIR /build

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
        -a \
        -installsuffix cgo \
        -ldflags "-s -w -extldflags '-static'" \
        -o service \
        cmd/api/main.go

FROM scratch

COPY --from=builder /build/service /app/
COPY --from=builder /build/data /app/data/
WORKDIR /app

ENTRYPOINT [ "./service" ]
