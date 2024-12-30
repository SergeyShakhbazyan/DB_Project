package main

import (
	"ProjectDB/config"
	"fmt"
	"log"
)

func main() {
	conn := config.NewConnection()
	defer conn.Close()

	var version string
	err := conn.DB.QueryRow("SELECT version();").Scan(&version)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	fmt.Printf("PostgreSQL Version: %s\n", version)
}
