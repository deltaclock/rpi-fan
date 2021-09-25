FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go get -u && go mod tidy && go mod why

COPY *.go ./

RUN go build -o /fanner

CMD [ "/fanner", "-file", "/temp" ]