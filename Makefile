mysql_login: 
	mysql -h 127.0.0.1 -u root -p freemail -P 3000

docker_up:
	docker-compose up -d

docker_down:
	docker-compose down

run:
	@echo "App running..."
	go run main.go

ports:
	netstat -tunlp

migrate:
	go run services/freemail/databases/migrate/migrations.go