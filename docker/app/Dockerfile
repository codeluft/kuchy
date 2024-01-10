FROM golang:1.21.5-alpine3.19 as build

WORKDIR /src
COPY . .

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base && apk add --no-cache wget
RUN make build

FROM alpine:3.19
COPY --from=build /src/build/bootstrap /bin/bootstrap
CMD ["/bin/bootstrap"]
