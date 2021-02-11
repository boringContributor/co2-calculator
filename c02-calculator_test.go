package main

import (
	"fmt"
	"testing"
)

// use exact numbers here
func TestCalculateCo2(t *testing.T) {
	var tests = []struct {
		distance             float64
		transportationMethod string
		distanceUnit         string
		outputUnit           string
		o2Usage              float64
	}{
		{15, "medium-diesel-car", "km", "kg", 2.565000},
		{14500, "train", "m", "kg", 0.087000},
		{14500, "train", "m", "g", 87},
		{1800.5, "large-petrol-car", "km", "kg", 507.741000},
	}
	for _, testData := range tests {
		testName := fmt.Sprintf("%f,%s,%s", testData.distance, testData.transportationMethod, testData.outputUnit)
		t.Run(testName, func(t *testing.T) {
			o2Usage := calculateCo2(testData.transportationMethod, testData.distance, testData.distanceUnit, testData.outputUnit)
			if o2Usage != testData.o2Usage {
				t.Errorf("got %f, want %f", o2Usage, testData.o2Usage)
			}
		})
	}
}
func TestExtractKey(t *testing.T) {
	var tests = []struct {
		key   string
		value float64
	}{
		{"bus", 27},
		{"train", 6},
		{"large-electric-car", 73},
		{"small-diesel-car", 142},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("%s,%f", testData.key, testData.value)
		t.Run(testName, func(t *testing.T) {
			value := extractValue(testData.key)
			if value != testData.value {
				t.Errorf("got %f, want %f", value, testData.value)
			}
		})
	}
}
func TestIsValidInput(t *testing.T) {
	var tests = []struct {
		distance             float64
		transportationMethod string
		outputUnit           string
		isValid              bool
	}{
		{1, "small-diesel-car", "kg", true},
		{1.0, "bus", "g", true},
		{0, "", "", false},
		{14.3, "", "kg", false},
		{14.3, "train", "mg", false},
		{-1, "", "", false},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("%f,%s,%s", testData.distance, testData.transportationMethod, testData.outputUnit)
		t.Run(testName, func(t *testing.T) {
			isValid := isValidInput(testData.distance, testData.transportationMethod, testData.outputUnit)
			if isValid != testData.isValid {
				t.Errorf("got %t, want %t", isValid, testData.isValid)
			}
		})
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		car     string
		isValid bool
	}{
		{"bus", true},
		{"train", true},
		{"large-electric-car", true},
		{"small-diesel-car", true},
		{"small-cat", false},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("%s", testData.car)
		t.Run(testName, func(t *testing.T) {
			isValid := isValidCar(testData.car)
			if isValid != testData.isValid {
				t.Errorf("got %t, want %t", isValid, testData.isValid)
			}
		})
	}
}
