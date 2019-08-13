FROM golang AS go_build_env

WORKDIR /app
COPY go_echo.go .
RUN go get -u github.com/labstack/echo/...   &&\
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o go_echo .


FROM alpine:3.10

WORKDIR /app
COPY --from=go_build_env /app/go_echo .
ENTRYPOINT ["./go_echo"]
