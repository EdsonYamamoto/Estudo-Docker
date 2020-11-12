FROM golang

ENV GO111MODULE=on
ENV PORT=9000

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o main
CMD ["/app/main"]
