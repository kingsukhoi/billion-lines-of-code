package helpers

import (
	"billion-challange/pkg/fileIter"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

func GenerateTemperatureLine(citiesArray []string) string {
	randomNumber := rand.Intn(len(citiesArray))
	city := citiesArray[randomNumber]
	randomTemp := rand.Float32() * 70
	isNegative := rand.Intn(2)
	if isNegative == 1 {
		randomTemp = -randomTemp
	}
	return fmt.Sprintf("%s;%.2f\n", city, randomTemp)
}

func GetCities() []string {

	citiesFile := fileIter.FileInfo{FilePath: "./world-cities.csv"}
	err := citiesFile.Open()
	if err != nil {
		log.Fatalf("Failed to open world-cities: %v", err)
	}
	defer citiesFile.Close()

	citiesArray := make([]string, 0)

	for i, v := range citiesFile.All() {
		if i == 0 {
			continue
		}
		citySplit := strings.Split(v, ",")
		citiesArray = append(citiesArray, citySplit[0])
	}

	return citiesArray
}
