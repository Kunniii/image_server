FROM golang:1.23.0 AS build

WORKDIR /otp/app

COPY go.mod go.sum ./

RUN go mod download

ENV GIN_MODE=release
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . .

RUN go build -o image_server

FROM scratch
EXPOSE 8080
COPY --from=build /otp/app/image_server /start_server
ENV GIN_MODE=release
ENTRYPOINT [ "/start_server" ]