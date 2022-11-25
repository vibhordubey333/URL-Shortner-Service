FROM golang:1.19-alpine AS builder
RUN mkdir /app
WORKDIR /app

# For Docker Version like 20.10.6 below copy command will work.

#COPY ../go.mod .
#COPY ../go.sum .
COPY . .
RUN chmod 0700 .
RUN go mod download
#RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./server/*.go

## the lightweight scratch image we'll
## run our application within
FROM alpine:latest AS production
## We have to copy the output from our
## builder stage to our production stage
COPY --from=builder /app .
RUN chmod +x .
## we can then kick off our newly compiled
## binary exectuable!!
CMD ["./main"]

