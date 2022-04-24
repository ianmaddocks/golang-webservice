# syntax=docker/dockerfile:1
FROM golang:1.16-alpine 
#FROM golang:1.12.0-alpine3.9

WORKDIR /app

#copy src files, *.go and subdirectories
COPY . ./
RUN go mod download

RUN go build -mod=readonly -v -o server

ENV PORT 8000
EXPOSE 8000

CMD [ "/app/server" ]


############################
# Build & run examples
############################
#docker build -t microservice2 .
#docker run  -it --rm --name my_microservice2 -p 80:8000 -e PORT=8000 microservice
#curl localhost:80/home