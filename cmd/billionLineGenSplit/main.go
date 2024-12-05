package main

import (
	"billion-challange/pkg/helpers"
	"fmt"
	"github.com/sourcegraph/conc/pool"
	"log"
	"os"
	"strings"
)

func main() {

	citiesArray := helpers.GetCities()

	currPool := pool.New()
	currPool.WithMaxGoroutines(10)
	for i := range 1000 {
		currPool.Go(func() {
			writeMillionLines(citiesArray, i)
		})
	}

}

func writeMillionLines(citiesArray []string, fileNum int) {
	fileName := fmt.Sprintf("./words_files/weather_stations_go_%d.csv", fileNum)

	fileToWrite, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to open weather_stations: %v", err)
	}
	defer fileToWrite.Close()
	var sb strings.Builder
	for range 1000000 {
		writeMe := helpers.GenerateTemperatureLine(citiesArray)
		sb.WriteString(writeMe)
	}
	_, errL := fileToWrite.WriteString(sb.String())
	if errL != nil {
		log.Fatalf("Failed to write to file: %v", errL)
	}
	fileToWrite.Sync()
	log.Printf("Finished writing file: %s", fileName)
}
