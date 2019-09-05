FROM golang:1.13 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o taxchallenge ./backend
RUN mkdir data

FROM scratch
COPY --from=builder /app/taxchallenge /app/
COPY --from=builder /app/taxrules.json /app/
COPY --from=builder /app/data /app/

ENV TAXCHALLENGE_TAXRULES="/app/taxrules.json"
ENV TAXCHALLENGE_DATABASE="/app/data/taxchallenge.boltdb"

EXPOSE 9091
ENTRYPOINT ["/app/taxchallenge"]
