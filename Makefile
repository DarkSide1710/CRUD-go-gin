postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=12345
#adminer:
#	docker run --rm -ti --network host adminer

migrate:
	migrate -path ./migrations -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up