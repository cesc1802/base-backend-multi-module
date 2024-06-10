FROM golang:1.22.1 as build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /go/bin/app main.go

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
COPY --from=build /go/src/app/.env /

EXPOSE 8080

ENTRYPOINT ["./app", "server"]