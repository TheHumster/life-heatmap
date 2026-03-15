package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

func connectDB() {
	godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL:", connStr)

	var err error

	db, err = pgx.Connect(context.Background(), connStr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Life Heatmap server is running buddy!")
}

func main() {
	connectDB()
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
