FROM golang:1.17 AS builder

# enable support for go modules
ENV GO111MODULE=on

WORKDIR /uicprovider

COPY ./go.mod ./.
RUN go mod download

ADD ./internal ./internal
COPY main.go ./

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /uicprovider/app ./
CMD ["./app"]