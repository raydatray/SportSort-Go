package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/raydatray/sportsort-go/db"
)

func main() {
	pool, err := pgxpool.New(context.Background(), connString string)

	if err != nil {
		fmt.Fprintf(os.stderr, "Failed to create DB pool: %v", err)
		os.Exit(1)
	}

	defer pool.Close()

	queries := db.New(pool)
	router := http.NewServeMux()


	server := http.Server {
		Addr: ":8080",
		Handler: router,
	}


	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
