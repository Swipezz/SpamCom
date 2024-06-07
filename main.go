package main

import (
	"fmt"
	"main/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// dsn := "root:root@tcp(127.0.0.1:3306)/24_cards?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Query := service.NewQuery(db)
	// Handler := handler.NewHandler(Query)
	ServerLess := handler.ServerLess()

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", ServerLess.MainMenu)
	r.Get("/test", ServerLess.Test)
	r.Get("/styletest", ServerLess.StyleTest)
	r.Get("/trash", ServerLess.Trash)
	r.Get("/image", ServerLess.Image)

	fmt.Println("Server Running at :8080")
	http.ListenAndServe(":8080", r)
}
