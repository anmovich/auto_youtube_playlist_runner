FROM golang:bookworm
COPY . /server
WORKDIR /server
RUN go mod tidy
RUN go build cmd/main.go

CMD ["go", "run", "cmd/main.go"]
