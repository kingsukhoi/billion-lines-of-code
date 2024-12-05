package main

import (
	"billion-challange/pkg/helpers"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	numPrinter := message.NewPrinter(language.English)

	citiesArray := helpers.GetCities()

	fileName := fmt.Sprintf("./weather_stations_go_%s.csv", time.Now().Format("2006-01-02-15-04-05"))

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	var sb strings.Builder

	for i := range 1000000000 {
		writeMe := helpers.GenerateTemperatureLine(citiesArray)
		sb.WriteString(writeMe)
		if i%10000000 == 0 {
			_, errL := file.WriteString(sb.String())
			if errL != nil {
				log.Fatalf("Failed to write to file: %v", errL)
			}
			errL = file.Sync()
			if errL != nil {
				log.Fatalf("Failed to sync file: %v", errL)
			}
			sb.Reset()
			log.Printf("Wrote %s lines\n", numPrinter.Sprintf("%d", i))
		}
	}
	_, err = file.WriteString(sb.String())
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	err = file.Sync()
	if err != nil {
		log.Fatalf("Failed to sync file: %v", err)
	}
}
