run:
	go run main.go
up: 
	docker compose up
create_db:
	docker exec -it practice-db-1 /usr/bin/psql -U baloo -d picshare