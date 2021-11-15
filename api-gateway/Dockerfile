FROM golang:latest as builder

ARG DOCKER_GIT_CREDENTIALS

WORKDIR /go/src/chatex/api-gateway
RUN git config --global credential.helper store \
    && echo "${DOCKER_GIT_CREDENTIALS}" > ~/.git-credentials \
    && git config --global http.sslverify false

COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

FROM scratch
WORKDIR /go/src/chatex/api-gateway
COPY --from=builder /go/src/chatex/api-gateway/api /go/src/chatex/api-gateway
COPY --from=builder /go/src/chatex/api-gateway/public /go/src/chatex/api-gateway/public
ENTRYPOINT ["./api"]
