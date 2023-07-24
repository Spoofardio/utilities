FROM golang:1.20.6

RUN mkdir /app 
ADD . /app/
WORKDIR /app 

RUN go install
RUN go build -o main
CMD ["./main"]
