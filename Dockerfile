FROM golang:1.12.13 as build
WORKDIR /build
COPY go.mod /build
COPY go.sum /build
RUN go mod download
COPY . /build
ENV CGO_ENABLE=0
RUN go build -o server .
EXPOSE 1323
FROM ubuntu:latest
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists*
COPY --from=build /build/server /app/
RUN mkdir asset && cd asset && mkdir images
CMD ["/app/server"]
