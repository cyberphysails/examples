FROM golang:1.17
ENV DIRNAME=/go/src/httpserver
EXPOSE 28080

WORKDIR $DIRNAME
COPY . .
RUN go build .
CMD $DIRNAME/httpserver
