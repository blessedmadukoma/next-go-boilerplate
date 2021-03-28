package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":"+port, nil)

	// For the Golang framework: GoFiber
	// app := fiber.New()

	// app.Static("/", "../frontendBuild")

	// app.Listen(":" + port)
}
