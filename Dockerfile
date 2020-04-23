FROM golang AS builder

WORKDIR /go/src/go-api-consumer
COPY . .
# disable cgo
ENV CGO_ENABLED=0
# build steps
RUN echo ">>> 1: go version" && go version
#RUN echo ">>> 2: go get" && go get -v -d
RUN echo ">>> 3: go install" && go install

FROM alpine
WORKDIR /go/bin/
COPY --from=builder /go/bin/go-api-consumer ./go-api-consumer
EXPOSE 8080
CMD ["./go-api-consumer"]