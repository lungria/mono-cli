package main

import (
	"testing"
	"time"
)

func Test_Next_WithValidInput_ReturnsCorrectAmountOfPeriods(t *testing.T) {
	startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	duration := time.Hour * time.Duration(1)

	calc := newPeriodCalculator(startDate, endDate, duration)

	countOfPeriods := 0
	for calc.Next() {
		t.Logf("Received date: %v\n", calc.current)
		countOfPeriods++
	}

	expected := 12
	if countOfPeriods != expected {
		t.Fatalf("Invalid count of periods returned: %v, expected: %v", countOfPeriods, expected)
	}
}

func Test_Next_WithMonobankLimits_ReturnsCorrectAmountOfPeriods(t *testing.T) {
	startDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Second * time.Duration(2682000)

	calc := newPeriodCalculator(startDate, endDate, duration)

	countOfPeriods := 0
	for calc.Next() {
		t.Logf("Received date: %v\n", calc.current)
		countOfPeriods++
	}

	expected := 12
	if countOfPeriods != expected {
		t.Fatalf("Invalid count of periods returned: %v, expected: %v", countOfPeriods, expected)
	}
}
