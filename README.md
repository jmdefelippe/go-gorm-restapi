# go-gorm-restapi

* Install Gorilla, Gorm and Air

```bash
go mod init github.com/jmdefelippe/go-gorm-restapi
```

```bash
go get -v -u github.com/gorilla/mux
```

```bash
go get -u gorm.io/gorm
```

```bash
go get -u gorm.io/driver/postgres
```

```bash
go install github.com/cosmtrek/air@latest
```

* Run Postgres with Docker

```bash
dc up -d
```

* Run app:

```bash
air ./pkg
```
