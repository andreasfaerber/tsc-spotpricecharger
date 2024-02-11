FROM golang:1.21 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /tsc-spotpricecharger

FROM scratch AS build-release-stage
WORKDIR /
COPY --from=build-stage /tsc-spotpricecharger /tsc-spotpricecharger

ENTRYPOINT ["/tsc-spotpricecharger"]

