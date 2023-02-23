# Reset database will rebuild database with initial values (seeders)
reset-db:
	go run ./migrate/migrate.go

# Run dev will run application but if have onde save will rebuild automatically
run-dev:
	docker-compose stop api
	nodemon --exec go run main.go --signal SIGTERM

# Reset entire application
re-run:
	docker-compose down
	docker image rm technical_challenge_q2bank-consumer
	docker image rm technical_challenge_q2bank-api
	docker image rm technical_challenge_q2bank-migrate
	docker-compose up

# Test files
unit-test:
	go test ./service/... 
	go test ./utils/... 
	go test ./handler/... 

# Generate swagger file
generate-swagger:
	swag init