package main

import (
	"billion-challange/pkg/fileIter"
	"billion-challange/pkg/model"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// var syncStations sync.Map
var wg sync.WaitGroup
var stations map[string]*model.City
var stationsMutex sync.Mutex

func main() {
	fmt.Println("Started at", time.Now().Format(time.RFC822))
	startTime := time.Now()

	stations = make(map[string]*model.City)

	currFile := &fileIter.FileInfo{FilePath: "./ignored_dir/weather_stations_go_2024-11-25-12-09-08.csv"}
	err := currFile.Open()
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer currFile.Close()
	//stations := make(map[string]*model.City)

	for _, v := range currFile.All() {
		wg.Add(1)
		go processCity(v)
	}

	wg.Wait()

	stationsArray := make([]model.City, len(stations))
	i := 0
	for _, v := range stations {
		stationsArray[i] = *v
		i++
	}

	sort.Slice(stationsArray, func(i, j int) bool {
		return stationsArray[i].Name < stationsArray[j].Name
	})

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Name, Min, Max, MeanSum, Count\n"))
	for _, v := range stationsArray {
		sb.WriteString(fmt.Sprintf("%s,%.2f,%.2f,%.2f,%.0f\n", v.Name, v.Min, v.Max, v.MeanSum/v.Count, v.Count))
	}

	log.Printf("Finished in %v", time.Since(startTime))
	file, _ := os.Create("./finished_cities.csv")
	_, _ = file.WriteString(sb.String())
	_ = file.Sync()
	_ = file.Close()
}

func processCity(line string) {
	splitString := strings.Split(line, ";")
	city := splitString[0]
	temperatureTemp, errL := strconv.ParseFloat(splitString[1], 32)
	if errL != nil {
		log.Fatalf("Error parsing float: %v", errL)
	}
	temperature := float32(temperatureTemp)

	stationsMutex.Lock()
	defer stationsMutex.Unlock()

	station, exists := stations[city]

	if !exists {
		//so the values ALWAYS get set on the first iteration
		station = &model.City{
			Name: city,
			Min:  math.MaxFloat32,
			Max:  -math.MaxFloat32,
		}
		stations[city] = station
	}
	if temperature < station.Min {
		station.Min = temperature
	}
	if temperature > station.Max {
		station.Max = temperature
	}

	station.MeanSum += temperature
	station.Count++

	wg.Done()
}
