FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN mkdir -p /app/scripts

CMD ["go", "run", "."]

