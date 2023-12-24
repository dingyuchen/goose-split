FROM golang:1.21 as build-stage

WORKDIR /workspace/goose-split

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

FROM gcr.io/distroless/static-debian12 AS build-release-stage

WORKDIR /workspace/goose-split

COPY --from=build-stage /workspace/goose-split/goose-split .

# mount persistent volume
VOLUME ["/workspace/goose-split/pb_data"]

EXPOSE 8090

CMD ["./goose-split", "serve", "--http=0.0.0.0:8090"] 

