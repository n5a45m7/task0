FROM golang:1.16.2-alpine3.13 AS builder

# RUN apk update && apk add --no-cache git ca-certificates tzdata openssh-client && update-ca-certificates

# # because https://git.apache.org/thrift.git is no longer available, but is used somewhere in dependencies
# RUN git config --global url."https://github.com/apache/thrift.git".insteadOf "https://git.apache.org/thrift.git"

# # https://golang.org/doc/faq#git_https
# RUN echo "machine ${GITLAB_HOST} login ${GITLAB_USERNAME} password ${GITLAB_PAT}" > /root/.netrc

WORKDIR /app

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build -v -o /app/bin main/main.go

FROM scratch

# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs /etc/ssl/certs/

COPY --from=builder /app/bin /bin

EXPOSE 8000

ENTRYPOINT ["/bin"]