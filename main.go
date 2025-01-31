package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Start Up
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		// insert SQL
		upsertSQL := `
		INSERT INTO registry (name, URI)
		VALUES ($1, $2)
		`

		// Example data to upsert
		name := "example_name"
		URI := "https://example.com"

		// Execute the upsert
		_, err = conn.Exec(context.Background(), upsertSQL, name, URI)
		if err != nil {
			log.Fatalf("Failed to upsert data: %v\n", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"uri":  URI,
		})
	})
	err = r.Run()
	if err != nil {
		return
	}
}
