# Reset database will rebuild database with initial values (seeders)
reset-db:
	go run ./migrate/migrate.go

# Run dev will run application but if have onde save will rebuild automatically
run-dev:
	nodemon --exec go run main.go --signal SIGTERM