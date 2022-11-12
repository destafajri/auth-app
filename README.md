# Simple Auth-app API Golang Postgres

Simple Authentication and Middlewares API Golang with PostgreSql for database and gin-gionic for the router

## How to use
- Please clone or download this repository.
- Prepare postgres database, or use docker, you can type
```
docker-compose up
```
OR
```
docker-compose up -d
```
- add .env file to setup your database connection
```
POSTGRES_URL="user=postgres password=[your-password] host=db.jibsjolwzprifvhwbawn.supabase.co port=5432 dbname=postgres"
KEY_JWT="[type your secret key here]"
```
- run the golang server
```
go run main.go
```
## Postman API-Doc
https://documenter.getpostman.com/view/22138766/2s8YerPCyM