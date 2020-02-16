FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server .
WORKDIR /app/public/generator
RUN go build -o ascii-art main.go .
CMD ["/app/server"]
