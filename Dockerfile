FROM golang:1.13
WORKDIR /go/src/github.com/traveloka/http-forwarder
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest
EXPOSE 8080
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/traveloka/http-forwarder/app /app
CMD ["/app"]
