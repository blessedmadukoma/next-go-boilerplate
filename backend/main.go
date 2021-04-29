package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	// "github.com/gofiber/fiber/v2"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
		}

		fmt.Println("Form Parsed!!!")

		name := r.PostFormValue("name")
		fmt.Printf("Hello, %s!\n", name)
	}

	file := "../frontendBuild/form.html"
	pt, _ := template.ParseFiles(file)
	err := pt.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template:", err)
	}

}

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

	// Gorilla Mux
	// r := mux.NewRouter()

	// r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../frontendBuild"))))

	// r.HandleFunc("/form", formHandler)

	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":"+port, nil)
	// http.ListenAndServe(":"+port, r)

	// For the Golang framework: GoFiber
	// app := fiber.New()

	// app.Static("/", "../frontendBuild")

	// app.Listen(":" + port)
}
