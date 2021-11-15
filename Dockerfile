# Build stage
FROM golang:1.14.2 as build

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o webapp

# Run stage
FROM alpine:3.11
COPY --from=build go/src/app/ app/
CMD ["./app/webapp"]