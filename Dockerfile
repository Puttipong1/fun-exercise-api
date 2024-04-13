FROM golang:latest as build-base
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go test -v
RUN go build -o ./out/go-sample .
FROM alpine:3.16.2
COPY --from=build-base /out/go-sample /go_sample
CMD ["/go-sample"]