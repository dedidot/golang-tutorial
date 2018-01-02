package main

import (
	"net/http"
	"./models"
	"./utils"
	"./routes"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	//routes
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
