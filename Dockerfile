FROM golang:1.22.3-alpine3.20 as build
ENV CGO_ENABLED=0

COPY . /src

WORKDIR /src

RUN go build

FROM alpine:3.20

COPY --from=build /src/blog-backend /blog-backend

CMD ["./blog-backend"]