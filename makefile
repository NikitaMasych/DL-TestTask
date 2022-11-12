test:
	cd ./code && go test ./... -v

run:
	docker compose up -d postgres
	docker compose up --build graphs