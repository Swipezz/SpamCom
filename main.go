package main

import (
	"fmt"
	"log"
	"main/handler"
	"main/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/24_cards?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	Query := service.NewQuery(db)
	Handler := handler.NewHandler(Query)

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", Handler.MainMenu)
	r.Get("/test", Handler.Test)

	fmt.Println("Server Running at :8080")
	http.ListenAndServe(":8080", r)
}
