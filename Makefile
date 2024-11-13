docker-up:
	docker compose --env-file dev.env up -d

docker-down:
	docker compose down