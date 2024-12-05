package main

import (
	"billion-challange/pkg/fileIter"
	"log"
	"time"
)

func main() {
	currFile := &fileIter.FileInfo{FilePath: "./weather_stations_go_2024-11-23-20-47-32.csv"}
	err := currFile.Open()
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer currFile.Close()

	startTime := time.Now()
	for range currFile.All() {

	}
	log.Printf("Finished processing file in %d", time.Since(startTime))

}
