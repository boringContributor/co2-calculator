package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func main() {

	// get initial data from cli
	distance, outputUnit, distanceUnit, transportationMethod := readInput()

	// validate input args
	if isValidInput(distance, transportationMethod, outputUnit) {
		co2Usage := calculateCo2(transportationMethod, distance, distanceUnit, outputUnit)
		printResult(co2Usage, outputUnit)

	} else {
		fmt.Println("Invalid arguments. Please stick to the below structure.")
		flag.Usage()
		os.Exit(2) // the same exit code flag.Parse uses
	}
}

// exists with code 2 if error in validation
func readInput() (float64, string, string, string) {

	const (
		defaultDistance     = 0.0
		defaultDistanceUnit = "km"
		defaultOutputUnit   = "g"
	)

	// attach custom message to helper argument
	flag.Usage = func() {
		log.Printf("Example usage of %s:\n", os.Args[0])
		log.Printf("%s --distance=10 --unit-of-distance=km --transportation-method=small-diesel-car \n\n", os.Args[0])
		log.Printf("Argument list: (required fields marked with *)")
		flag.PrintDefaults()
	}

	// required: > 0
	var distance float64
	flag.Float64Var(&distance, "distance", defaultDistance, "The covered distance of the car (*)")

	// either kg or g (default g)
	var outputUnit string
	flag.StringVar(&outputUnit, "output", defaultOutputUnit, "The output unit of the Co2 usage (*)")

	// either m = meter | km = kilometre (default m)
	var distanceUnit string
	flag.StringVar(&distanceUnit, "unit-of-distance", defaultDistanceUnit, "The distance unit (either m or km)")

	// required: see markdown for allowed inputs
	var transportationMethod string
	flag.StringVar(&transportationMethod, "transportation-method", "", "The type of the car. See allowed input in readme! (*)")
	flag.Parse()

	return distance, outputUnit, distanceUnit, transportationMethod
}

func calculateCo2(transportation string, distance float64, distanceUnit string, outputUnit string) float64 {
	const metric = 1000

	if distanceUnit == "m" {
		distance = distance / metric
	}
	co2Usage := extractValue(transportation) * distance
	if outputUnit == "kg" {
		co2Usage = co2Usage / metric
	}
	return co2Usage
}

func printResult(co2Usage float64, outputUnit string) {
	// round to nearest precision and cut to one digit
	roundedResult := math.Round(co2Usage*10) / 10
	log.Printf("Your trip caused %.1f%s of CO2-equivalent.", roundedResult, outputUnit)
}

// extracts a value for a given key from a json file
func extractValue(key string) float64 {
	byteValue, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println("error while loading calculation data: ", err)
		os.Exit(2) // the same exit code flag.Parse uses
	}
	var result map[string]interface{}
	_ = json.Unmarshal(byteValue, &result)
	return result[key].(float64)
}

func isValidInput(distance float64, transportationMethod string, outputUnit string) bool {
	return distance > 0 && transportationMethod != "" && isValidCar(transportationMethod) && (outputUnit == "kg" || outputUnit == "g")
}

// checks if given key is within an array
func isValidCar(key string) bool {
	// not nice to hardcode this but I have to live with it now
	cars := [14]string{"small-diesel-car", "small-petrol-car", "small-plugin-hybrid-car",
		"small-electric-car", "medium-diesel-car", "medium-petrol-car", "medium-plugin-hybrid-car",
		"medium-electric-car", "large-diesel-car", "large-petrol-car", "large-plugin-hybrid-car",
		"large-electric-car", "bus", "train"}

	for _, item := range cars {
		if item == key {
			return true
		}
	}
	return false
}
