FROM golang:latest AS build

WORKDIR /app

COPY go.mod .

RUN go version && go mod download

COPY . .

RUN CGO_ENABLED=0 go build -buildvcs=false -trimpath -ldflags '-w -s' -o myapp ./cmd/webapp/main.go

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=build /app/myapp /

ENTRYPOINT ["/myapp"]
