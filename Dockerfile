FROM golang:1.13.5 as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o zapchallenge ./backend

FROM scratch
COPY --from=builder /app/zapchallenge /app/
COPY --from=builder /app/portals.yaml /app/

ENV ZAPCHALLENGE_SOURCE="http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
ENV ZAPCHALLENGE_PORTALS="/app/portals.yaml"

EXPOSE 9091
ENTRYPOINT ["/app/zapchallenge"]
