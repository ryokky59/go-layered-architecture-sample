FROM golang:1.13.11-alpine3.11 AS build

WORKDIR /
COPY . /go/src/github.com/ryokky59/sample_ddd_go
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/google/uuid \
  && go get github.com/gorilla/mux
RUN cd /go/src/github.com/ryokky59/sample_ddd_go/api && go build -o bin/sample main.go

FROM alpine:3.8
COPY --from=build /go/src/github.com/ryokky59/sample_ddd_go/api/bin/sample /usr/local/bin/
CMD ["sample"]
