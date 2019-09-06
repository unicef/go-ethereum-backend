# go-ethereum-server
Golang backend to demonstrate use cases of user generated cryptographic tokens.

Front-end: [elm-web-client](https://github.com/unicef/elm-web-client)

# Dependencies
- Docker

# Technologies
- Golang
- Ethereum
- MySQL

# Build
```bash
$ export APP_ENV=dev
$ docker build -t go-ethereum-server --build-arg APP_ENV .
```

# Run using docker-compose
Configure your env variables in the docker-compose.yml

```yml
version: "3"

services:
    db:
        image: mysql
        command: --default-authentication-plugin=mysql_native_password
        restart: always
        environment:
            MYSQL_ROOT_PASSWORD: example
            MYSQL_DATABASE: database
        ports:
            - "3306:3306"
        tty: true
        volumes:
            - dignity-mysql-volume:/data/dbmysql
    api:
        build:
        context: ./backend
        restart: always
        volumes:
            - ./backend:/go/src/github.com/unicef/dignity-platform/backend
        depends_on:
            - db
        environment:
            ENVIRONMENT: "development"
            BASE_URL: "http://localhost"
            RUN_API: "1"
            DATA_SOURCE_NAME: "root:example@tcp(db:3306)/database?charset=utf8mb4,utf8&parseTime=true"
            TEST_DATA_SOURCE_NAME: "root:example@tcp(db:3306)/test_database?charset=utf8mb4,utf8&parseTime=true"
            GOMAXPROCS: 4
            DB_MIGRATIONS_PATH: /go/src/github.com/qjouda/dignity-platform/backend/db_migrations
            TEST_DB_USER: root
            TEST_DB_PASS: password
            SLACK_WEBHOOK_URL: [slack url]
            FORCE_SSL: 0
            EMAIL_FROM: email@from.com
            AWS_SES_KEY: [aws key]
            AWS_SES_SECRET: [aws secret]
            ETHEREUM_NET: "SIM"
            SIM_ETH_HOST: "http://172.25.0.110:8545"
            RINKEBY_ETH_HOST: "https://rinkeby.infura.io/"
            ETH_KEY_RAW: "ROW_KEY"
            ETH_KEY_STORE_DIR: "keystore"
            FORCE_ENV_CHECKING: 1
        tty: true
        ports:
            - "3000:3000"
```

# Build and run on production
```bash
# Build an image for production
$ export APP_ENV=production
$ docker-compose up api
```


# Acknowledgements
- Khaled Jouda (github [@kjda](https://github.com/kjda))
