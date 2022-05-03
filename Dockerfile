# syntax=docker/dockerfile:1
FROM golang:1.16-alpine 
#FROM golang:1.12.0-alpine3.9

WORKDIR /app

#copy src files, *.go and subdirectories
COPY . ./
RUN go mod download

RUN go build -mod=readonly -v -o server

ENV PORT 80
EXPOSE 80

CMD [ "/app/server" ]