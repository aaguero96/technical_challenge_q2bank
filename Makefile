# Reset database will rebuild database with initial values (seeders)
reset-db:
	go run ./migrate/migrate.go

run-dev:
	nodemon --exec go run main.go --signal SIGTERM 