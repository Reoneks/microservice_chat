FROM golang:latest as builder

WORKDIR /go/src/git.computools.org/golang-team/user-service

COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user .

FROM scratch
WORKDIR /go/src/git.computools.org/golang-team/user-service
COPY --from=builder /go/src/git.computools.org/golang-team/user-service/user /go/src/git.computools.org/golang-team/user-service
ENTRYPOINT ["./user"]
