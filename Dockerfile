FROM golang:1.18.2-alpine3.15

WORKDIR /app
COPY src/ ./

RUN ls && go mod download
RUN go build -o /icanhazpw

EXPOSE 8080
CMD [ "/icanhazpw" ]