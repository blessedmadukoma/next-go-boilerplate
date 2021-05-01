package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Golang and Next.js connection")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("port")

	// For Golang without framework
	fs := http.FileServer(http.Dir("../frontendBuild"))
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":"+port, nil)

	// For the Golang framework: GoFiber
	// app := fiber.New()

	// app.Static("/", "../frontendBuild")

	// app.Listen(":" + port)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	form := "../frontendBuild/form.html"
	println("Method used:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles(form)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		name := r.FormValue("name")
		fmt.Println("name:", r.Form["name"])
		// fmt.Fprintf(w, "name: %s", r.Form["name"])
		fmt.Println("name:", r.FormValue("name"), "type:", reflect.TypeOf(name))
		fmt.Fprintf(w, "name: %s", r.FormValue("name"))
	}
}
