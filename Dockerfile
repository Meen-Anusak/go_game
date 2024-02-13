FROM golang:1.21-alpine3.19 AS builder
RUN apk --no-cache add gcc g++ make git

# ARG CGO_ENABLED=0

WORKDIR /go/src/app

COPY . .
COPY ./nginx/nginx-app.conf /etc/nginx/conf.d/default.conf
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

# Start from a minimal Alpine image
FROM alpine:3.13

# # Install required dependencies
RUN apk --no-cache add ca-certificates

# # Set the working directory
WORKDIR /usr/bin

# Copy the built Golang application from the builder stage
COPY --from=builder /go/src/app/bin /go/bin
ENV TZ=Asia/Bangkok
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80
