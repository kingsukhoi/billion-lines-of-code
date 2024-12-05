package main

import (
	"billion-challange/pkg/fileIter"
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite" // Import the SQLite driver
	"github.com/sourcegraph/conc"
	"log"
	"strings"
)

func main() {
	db, err := sql.Open("sqlite", "./cities.db") // Replace with your file name
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	fmt.Println("Successfully connected to the database!")

	wg := conc.WaitGroup{}

	currFile := fileIter.FileInfo{FilePath: "./weather_stations_go_2024-11-23-20-47-32.csv"}
	currFile.Open()
	for _, v := range currFile.All() {
		currSplit := strings.Split(v, ";")
		wg.Go(func() {
			insertRow(db, currSplit[0], currSplit[1])
		})
	}
	wg.Wait()
	err = db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}

func insertRow(db *sql.DB, city string, temp string) {

	stmt, err := db.Prepare("insert into main.weather_stations_go_2024_11_23_20_47 (city, temperature) values (?,?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(city, temp)
	if err != nil {
		fmt.Println("Error executing statement:", err)
		return
	}

}
