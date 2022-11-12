FROM golang:alpine as build-stage

RUN mkdir /build
WORKDIR /build

COPY ./data/schedule.csv .

COPY ./code/go.mod ./code/go.sum ./
RUN  go mod download

COPY ./code .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM gcr.io/distroless/static-debian11 as prd

WORKDIR /root/

COPY --from=build-stage /build/schedule.csv /data/schedule.csv
COPY --from=build-stage /build/main .
COPY --from=build-stage /build/.env .

CMD ["./main"]
 