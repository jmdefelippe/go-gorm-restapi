package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmdefelippe/go-gorm-restapi/pkg/db"
	"github.com/jmdefelippe/go-gorm-restapi/pkg/models"
	"github.com/jmdefelippe/go-gorm-restapi/pkg/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	routes.PublicRoutes(r)

	http.ListenAndServe(":3000", r)
}
