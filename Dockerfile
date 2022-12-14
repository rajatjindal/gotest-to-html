FROM golang:1.19-alpine3.17 as builder

WORKDIR /go/src/github.com/rajatjindal/gotest-to-html
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go test ./... -cover
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w" -o bin/gotest-to-html main.go

FROM alpine/git:2.36.3

COPY --from=builder /go/src/github.com/rajatjindal/gotest-to-html/bin/gotest-to-html /usr/local/bin/
COPY --from=builder /go/src/github.com/rajatjindal/gotest-to-html/entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]