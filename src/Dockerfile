FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go get -u && go mod tidy && go mod why

RUN go build -o /fanner

CMD [ "/fanner", "-file", "/temp" ]