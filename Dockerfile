FROM golang:1.15-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/dastco

# Smoke test
RUN /go/bin/dastco

FROM scratch
COPY --from=builder /go/bin/dastco /go/bin/dastco
ENTRYPOINT ["/go/bin/dastco"]