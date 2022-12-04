# Build Stage
# First pull Golang image
FROM golang:1.19.3-alpine
ENV APP_NAME=APP
ENV CMD_PATH=main.go
ENV GOOS linux

WORKDIR "/$APP_NAME"

# Copyapplication data into image
COPY . ./
RUN pwd
RUN ls -la

RUN go mod download

RUN  go build -o main.go

# Expose application port
EXPOSE 8080

# Start app 
CMD ./$APP_NAME