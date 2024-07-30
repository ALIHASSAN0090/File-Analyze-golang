FROM golang:1.22.5

WORKDIR /File-Analyzer-by-golang
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main

EXPOSE 3000

CMD [ "./main" ]