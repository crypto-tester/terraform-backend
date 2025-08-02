# build binary
FROM golang:1.23-alpine AS builder

COPY . /go/src/github.com/crypto-tester/terraform-backend

WORKDIR /go/src/github.com/crypto-tester/terraform-backend

RUN GOOS=linux CGO_ENABLED=0 go build cmd/terraform-backend.go

# start clean for final image
FROM alpine:3

COPY --from=builder /go/src/github.com/crypto-tester/terraform-backend/terraform-backend /terraform-backend

ENTRYPOINT ["/terraform-backend"]
