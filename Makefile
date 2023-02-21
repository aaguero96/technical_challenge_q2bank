# Reset database will rebuild database with initial values (seeders)
reset-db:
	go run ./migrate/migrate.go

# Run dev will run application but if have onde save will rebuild automatically
run-dev:
	nodemon --exec go run main.go --signal SIGTERM

# Reset entire application
re-run:
	docker-compose down
	docker image rm technical_challenge_q2bank-consumer
	docker image rm technical_challenge_q2bank-api
	docker-compose up