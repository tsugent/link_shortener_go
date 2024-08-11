package main

import "fmt"

func main() {
	db, err := connectDatabase("./links.db")
	if err != nil {
		fmt.Printf("Failed to connect to the database: %s\n", err)
		return
	}
	defer db.Close()

	if err := checkAndInitializeDatabase(db); err != nil {
		fmt.Printf("Failed to initialize the database: %s\n", err)
		return
	}
}
