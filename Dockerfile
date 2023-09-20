FROM golang:1.20-alpine 

WORKDIR /app

# to skip re-downloading and installing deps
COPY ./go.mod .

RUN go mod download

COPY . .

RUN go build -o /gather-events

CMD [ "/gather-events" ]
