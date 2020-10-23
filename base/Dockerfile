FROM golang:latest

COPY src/http-sample /app/http-sample

RUN chmod +x /app/http-sample

WORKDIR /app

EXPOSE 3000

ENTRYPOINT ["/app/http-sample"]

