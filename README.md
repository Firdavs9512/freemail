# freemail

Create new server for use this code. I am used DigitalOcean. 

___
# How use this

* Clone this project
```bash
  git clone https://github.com/firdavs9512/freemail
```
* Download all packages
```bash
  go mod tidy
```
* Create .env file
```bash
  APP_PORT=8080
  MYSQL_USER={user_name}
  MYSQL_PASSWORD={user_password}
  MYSQL_NAME=freemail
  EMAIL_ADDRES={gmail_address}
  EMAIL_PASSWORD={gmail_password}
```
* Migrate database `make migrate` or
```bash
  go run service/freemail/databases/migrate/migrations.go
```
* Run app `make run` or
```bash
  go run main.go
```

* All is corect build app
```bash
  go build main.go
```

View complate app: https://freeservices.uz

[Full application web site](https://freeservices.uz)
