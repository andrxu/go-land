FROM golang:1.21 as builder
WORKDIR /go/src/github.com/andrxu/go-land
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./ .
RUN make

FROM gcr.io/distroless/base
WORKDIR /
COPY --chmod=555 --from=builder /go/src/github.com/andrxu/go-land/go-land .
EXPOSE 8080
CMD ["./go-land"]