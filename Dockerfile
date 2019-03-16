FROM golang:1.12
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/savaki/jq
RUN go build -o main .
CMD ["./main"]
