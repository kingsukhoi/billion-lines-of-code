package helpers

import (
	"billion-challange/pkg/fileIter"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
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

var currRand *rand.Rand
var lastFetch *time.Time

func GenerateTemperatureLineCloudflare(citiesArray []string) string {

	if currRand == nil || lastFetch == nil || time.Since(*lastFetch).Seconds() > 30 {
		var err error
		currRand, err = GetCloudflareRand()
		if err != nil {
			panic(err)
		}

		clf := time.Now()
		lastFetch = &clf
		log.Printf("Fetched new random number generator")
	}

	randomNumber := currRand.Intn(len(citiesArray))
	city := citiesArray[randomNumber]
	randomTemp := currRand.Float32() * 70
	isNegative := currRand.Intn(2)
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
