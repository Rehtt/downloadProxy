FROM golang:1.18.0-alpine3.14 AS build

COPY . /opt
WORKDIR /opt
RUN apk update && apk add git && go build -o proxy ./

FROM alpine

COPY --from=build /opt /opt

WORKDIR /opt

EXPOSE 8080
CMD ["/opt/proxy"]
