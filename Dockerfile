FROM golang:1.16

RUN mkdir /code
WORKDIR /code
COPY . /code

RUN go build

EXPOSE 8080

CMD ["./task-test"]