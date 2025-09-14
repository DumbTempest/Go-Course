package main

import (
	"Go/Femm/api"
	"Go/Femm/data"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello yo"))
}

func handleTemplates(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing template"))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
	server := http.NewServeMux()
	server.HandleFunc("/hello", helloHandler)
	server.HandleFunc("/template", handleTemplates)
	server.HandleFunc("/api/exhibitions", api.Get)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	server.HandleFunc("/api/exhibitions/add", api.Post)
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
