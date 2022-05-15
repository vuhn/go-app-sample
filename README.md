# Sample Go Application

This is a sample go application based on ddd idea.


## Start app on local without Docker

Required: PostgreSQL is running on your local machine.

```sh
export DB_HOST=postgres_host
export DB_USER=postgres_user
export DB_PASSWORD=postgres_password
export DB_PORT=postgres_port
export DB_NAME=db_name
make start-api
```

## Start app on local using Docker

Required: Docker is running on your local machine.

```sh
make run
```
