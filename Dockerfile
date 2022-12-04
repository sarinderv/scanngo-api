# Build Stage
# First pull Golang image
FROM golang:1.19.3-alpine
ENV APP_NAME=scanngo-api
ENV CMD_PATH=main.go
ENV GOOS linux

WORKDIR "/APP"

# Copyapplication data into image
COPY . ./
RUN pwd
RUN ls -la

RUN go mod download

RUN  go build -o /

# Expose application port
EXPOSE 8080/tcp

# Start app 
CMD /scanngo-api