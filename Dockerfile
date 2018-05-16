FROM golang:1.10.2-alpine3.7

WORKDIR /go/src/matchmaker
COPY . .

RUN apk add --no-cache git

RUN go install -v # "go install -v ./..."

RUN apk del git

CMD ["matchmaker"]
