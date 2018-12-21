FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]

EXPOSE 3666

#AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
 # AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}