postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=12345
#adminer:
#	docker run --rm -ti --network host adminer

migrate:
	migrate -sources file://migrations \
			-database postgres://postgres:12345@localhost/postgres?sslmode=disable up