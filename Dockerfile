FROM golang:1.21-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  container-user

WORKDIR $GOPATH/src/tony/

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /invoice-issuer ./cmd/invoice-issuer
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /invoice-webhook ./cmd/invoice-webhook

FROM scratch AS invoice-issuer

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /invoice-issuer .

USER container-user:container-user

CMD ["./invoice-issuer"]

FROM scratch AS invoice-webhook

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /invoice-webhook .

USER container-user:container-user

CMD ["./invoice-webhook"]
