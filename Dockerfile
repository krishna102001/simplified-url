FROM golang:latest

WORKDIR /app

COPY go* .

RUN go mod download

COPY . .

RUN go build -o /my-app 

CMD [ "/my-app" ]