FROM golang:1.19-alpine3.17 as builder

WORKDIR /go/src/github.com/rajatjindal/gotest-to-html
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go test ./... -cover
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w" -o bin/gotest-to-html main.go

FROM alpine:3.17

WORKDIR /home/app

# Add non root user
RUN addgroup -S app && adduser app -S -G app
RUN chown app /home/app

USER app

COPY --from=builder /go/src/github.com/rajatjindal/gotest-to-html/bin/gotest-to-html /usr/local/bin/

CMD ["gotest-to-html"]