# Use an official Go runtime as a parent image
FROM golang:1.18

WORKDIR /app
COPY . .
RUN ls
WORKDIR /app/cmd/Auth/

RUN go build -o /app/build/Auth -mod vendor

WORKDIR /app/build
CMD ["./Auth"]
