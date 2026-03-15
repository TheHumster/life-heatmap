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

func addActivity(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO activities (name, minutes) VALUES ($1, $2)",
		"coding",
		60,
	)

	if err != nil {
		http.Error(w, "Failed to add activity", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Activity added successfully")
}

func main() {
	connectDB()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addActivity)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
