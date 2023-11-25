build-invoice-issuer:
	docker build --network=host --target invoice-issuer -t tony/invoice-issuer .

build-invoice-webhook:
	docker build --network=host --target invoice-webhook -t tony/invoice-webhook .

build: build-invoice-issuer build-invoice-webhook

run: build
	docker-compose up

daemon: build
	docker-compose up -d