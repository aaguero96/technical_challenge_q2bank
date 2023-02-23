<h1>Technical challange by q2bank</h1>

<h2>Summary</h2>

1. [Description](#description)
2. [Requirements](#requirements)
3. [Deadline](#deadline)
4. [Schedule](#schedule)
5. [Requirements to run local](#requirements-to-run-local)
6. [Run local](#run-local)
7. [Project Architecture](#project-architecture)
8. [Endpoints](#endpoints)
9. [Info About Requests](#info-about-requests)
10. [Requests in postman](#requests-in-postman)
11. [Requests in swagger](#requests-in-swagger)
12. [Next steps](#next-steps)

<h2 id="description">Description</h2>

- The technical challenge consists of creating an API REST that exposes an system based on transactions common to common or common to storekeeper, but not storekeeper to storekeeper.

<h2 id="requirements">Requirements</h2>

- Technical requirements:
  - Golang application
  - Relational database (your preference)
  - RESTFul pattern
  - Clean code

- Extra:
  - Container our application
  - Semantic commit
  - Unit tests
  - Swagger documentation

<h2 id="deadline">Deadline</h2>

- The test starts on `📅 16 feb 2023 - Thursday` with 7 days (calendar days) of time to finish
- Ergo the deadline is `📅 23 feb 2023 - Thursday`

<h2 id="schedule">Schedule</h2>

- [Schedule to this project](https://technical-challange-q2bank.atlassian.net/jira/software/projects/TCBQ/boards/1)

<h2 id="requirements-to-run-local">Requirements to run local</h2>

1. `Docker version 20.10.23` => [install docker](https://docs.docker.com/get-docker/)
2. `Docker Compose version v2.12.2` (for docker-compose commands) => [install docker compose](https://docs.docker.com/compose/install/)
4. `GNU Make 4.2.1` (for makefile commands) => [install make](https://www.gnu.org/software/make/)
3. You have to use `docker-compose` instead `docker compose`
5. All environment variables that you need are in `./.env.example` file
6. Ports:
  - For this project you need to have ports open in your sistem, that PORTS are:
    - 3000 (to api)
    - 6379 (to redis)
    - 8081 (to redis-commander)
    - 5432 (to postgres database)
  - If you dont have one of those PORTS opened, You will have kill them with those commands:
    - `sudo lsof -i:<PORT>` (to discover what is in your PORT with value <PORT>)
    - `sudo kill -9 <PID>` (to kill what is running in your PORT with PID value <PID>)
7. If you need to run your application without docker (not recommended): `go version go1.18.4 linux/amd64` => [install go](https://go.dev/doc/install)
8. I recommend use postman to test endpoints (but you have more options) => [install postman](https://www.postman.com/downloads/)

<h2 id="run-local">Run local</h2>

1. Clone repository
- If you choose HTTPS: `git clone https://https://github.com/aaguero96/technical_challenge_q2bank/blob.git`
- If you choose SSH: `git clone git@github.com:aaguero96/technical_challenge_q2bank.git`

2. Install dependencies
- `go mod tidy`

3. .env file
- In root directory has one file named `.env.example`
- copy that file and rename that to `.env` (if you prefer run `cp ./.env.example ./.env`)
- fill up the env vars with correct values (only in case of this project the file envs has correct values)

4. Run project with docker-compsoe
- `docker-compose up` (await for message `Listening and serving HTTP on :3000` to start)

5. Other commands
- If you need to recreate your database run: `make reset-db`
- If you recreate your application run: `make re-run`
- If you need to test services (in moment) run: `make unit-test`
- If you need to generate swagger file run: `make generate-swagger`
- If you need to run just your api service local run: `make run-dev`

<h2 id="project-architecture">Project architecture</h2>

<h3>Directories</h3>

1. `./.github`: have workflows and templates for git commits, pushs, pulls, ...
2. `./initializers`: have functions that use to instance an external/internal service
3. `./utils`: have simple functions that are used in other directories
4. `./models`: have one model for every table in project
5. `./migrate/seeders`: have one seeder (initial values in database) for tables that you need
6. `./migrate/migrate.go`: have one main function whitch create your database
7. `./repository`: have others directories that one of them represents actions in database
8. `./service`: have others directories that one of them represents services
9. `./handler`: have others directories that one of them represents handlers for your http requests
10. `./middleware`: have functions that use to handle with http requests before the handler (if you need)
11. `./routes`: have functions that instance a new route with http methods
12. `./extenalAPI`: have others directories that one of them represents uses of external API 
13. `./events/producer`: have functions that produce messages for redis
14. `./events/consumer/handler`: have functions that consume messages from redis
15. `./events/consumer/main.go`: have main funtion that start consuming of messages from redis

<h3>Request flow</h3>

<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/request_flow.png?raw=true"/>

<h3>Data base explanation</h3>

1. Relations in database are in the image bellow:

<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/data_base_relations.png?raw=true"/>

2. Details about tables in database

- users
  - password should contain at least 5 characters, at least one upper case, at least one lower case, at least one numeric, at least one special
  - if register_type_id is related to "CPF" register_number should have 11 digits and if is "CNPJ" register_number should have 14 digits
  - email has to be correct format

- transactions
  - payer_id should not be related with "storekeeper"
  - status has 5 possiblities: denied, in process, completed, canceled, cancel in progress

<h2 id="endpoints">Endpoints</h2>

1. POST localhost:3000/v1/login
- Body:
```
{
  "email": "string",
  "password": "string"
}
```
- Response:
```
{
  "token": "string",
  "expiring_in": "string"
}
```
- Cookies: this method put token in localhost cookies as "token"

2. POST localhost:3000/v1/users
- Body:
```
{
  "name": "string",
  "email": "string",
  "password": "string",
  "register_number": int,
  "register_type_id": int,
  "user_type_id": int
}
```
- Response:
```
{
  "token": "string",
  "expiring_in": "string"
}
```
- Cookies: this method put token in localhost cookies as "token"

3. GET localhost:3000/v1/users
- Headers: Bearer token (or cookie "token" if has saved)
- Response:
```
[
  {
      "name": "string",
      "email": "string",
      "wallet_id": int
  }
]
```

4. GET localhost:3000/v1/users/:id
- Headers: Bearer token (or cookie "token" if has saved)
- Params: `id int`
- Response:
```
{
    "name": "string",
    "register_number": int,
    "register_type_id": int,
    "email": "string",
    "wallet_id": int,
    "user_type_id": int
}
```

5. GET localhost:3000/v1/register_types
- Headers: Bearer token (or cookie "token" if has saved)
- Response:
```
[
  {
    "id": int,
    "type": "string"
  }
]
```

6. GET localhost:3000/v1/user_types
- Headers: Bearer token (or cookie "token" if has saved)
- Response:
```
[
  {
    "id": int,
    "type": "string"
  }
]
```

7. GET localhost:3000/v1/user_types/:id
- Headers: Bearer token (or cookie "token" if has saved)
- Params: `id int`
- Response:
```
{
  "user_type_id": int,
  "user_type": "string"
}
```

8. GET localhost:3000/v1/wallets
- Headers: Bearer token (or cookie "token" if has saved)
- Response:
```
[
  {
    "id": int,
    "amount": float
  }
]
```

9. GET localhost:3000/v1/wallets/:id
- Headers: Bearer token (or cookie "token" if has saved)
- Params: `id int`
- Response:
```
{
  "wallet_id": int,
  "amount": float
}
```

10. PATCH localhost:3000/v1/wallets/:id
- Headers: BasicAuth
- Params: `id int`
- Body:
```
{
  "amount": float
}
```
- Response:
```
{
  "wallet_id": int,
  "amount": float
}
```

11. POST localhost:3000/v1/transactions
- Headers: Bearer token (or cookie "token" if has saved)
- Body:
```
{
  "payer_id": int,
  "payee_id": int,
  "amount": float
}
```
- Response:
```
{
    "transaction_id": int,
    "payer_id": int,
    "payee_id": int,
    "amount": float,
    "status": "in progress"
}
```

12. GET localhost:3000/v1/transactions
- Headers: Bearer token (or cookie "token" if has saved)
- Response:
```
[
  {
    "payer_id": int,
    "payee_id": int,
    "amount": float,
    "status": "string"
  }
]
```

13. GET localhost:3000/v1/transactions/:id
- Headers: Bearer token (or cookie "token" if has saved)
- Params: `id int`
- Response:
```
{
  "transaction_id": int,
  "payer_id": int,
  "payee_id": int,
  "amount": float,
  "status": "string"
}
```

14. DELETE localhost:3000/v1/transactions/:id
- Headers: Bearer token (or cookie "token" if has saved)
- Params: `id int`
- Response: ``

<h2 id="info-about-requests">Info about requests</h2>

- To must of endpoints you need a token access, this token is adquired in two endpoints (the body is an example but work in that database):
  - `POST localhost:3000/v1/login`
  ```
  body
  {
    "email": "email1@testmail.com",
    "password": "Def4!t*1"
  }
  ```
  - `POST localhost:3000/v1/users`
  ```
  body
  {
    "name": "andre aguero",
    "email": "andre@teste.com",
    "password": "Def4!t*0",
    "register_number": 12345678900,
    "register_type_id": 1,
    "user_type_id": 2
  }
  ```
- In endpoint `PATCH localhost:3000/v1/wallets/:id` the BasicAuth is:
  - username: admin
  - password: admin
  - I put that here just because is demonstrative, this is not secure in real projects (this would be passed by keybase or other software)

<h2 id="requests-in-postman">Requests in postman</h2>

1. Import collection to postman
- In this repository has one file name `psotman.json`
- Visit postman and select option `import` (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/postman_1.png?raw=true"/>

2. Defining postman envs
- In postman select option `Environments` (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/postman_2.png?raw=true"/>
- Create Environment named DEV (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/postman_3.png?raw=true"/>
- Up right on page you have to select env DEV
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/postman_4.png?raw=true"/>

3. Use collection
- In postman select option `Collections` and you have acees to all endpoints (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/postman_5.png?raw=true"/>

<h2 id="requests-in-swagger">Requests in swagger</h2>

1. Access `localhost:3000/docs/index.html`
2. You have three options to navigate on swagger:
- Not authenticated: you just have acces to not lock endpoints (POST /v1/login and POST /v1/users)
- authenticated with BearerToken: you have acess to every endpoint but not to PATCH /v1/wallets/:id
- authenticated with Basic authorization: you have access just to endpoint PATCH /v1/wallets/:id
3. Authenticate
- To authenticate you press the button in up rigth on page (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/swagger_1.png?raw=true"/>
- Fill the form with correct values (as image bellow)
<img src="https://github.com/aaguero96/technical_challenge_q2bank/blob/main/readme_images/swagger_2.png?raw=true"/>
- If token is "123456789abc987654321@@##" you have to fill just field BearerToken with "Bearer 123456789abc987654321@@##", Bearer as prefix is mandatory

<h2 id="next-steps">Next steps</h2>
- Instance database in cloud, aws or other tool
- Instance redis in cloud, aws or other tool
- Deploy image of consumer in cloud, aws or other tool
- Deploy image of api in cloud, aws or other tool
- Create more unit-tests
- Create integration-tests
- Add test in CI
- Create CI/CD on github
- Put strings on utils directory