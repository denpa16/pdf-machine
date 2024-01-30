FROM golang:1.21

EXPOSE 8000/tcp

WORKDIR /app

COPY go.mod go.sum /app/

COPY . /app/

RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates
RUN apt-get install -y --no-install-recommends wkhtmltopdf
