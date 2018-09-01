up:
	docker-compose up -d
	docker-compose logs -f

down:
	docker-compose down

restart:
	docker-compose down
	make up