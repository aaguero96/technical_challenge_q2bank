# Reset database will rebuild database with initial values (seeders)
reset-db:
	go run ./migrate/migrate.go
