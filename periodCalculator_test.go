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

	coundOfPeriods := 0
	for calc.Next() {
		coundOfPeriods++
	}

	expected := 12
	if coundOfPeriods != expected {
		t.Fatalf("Invalid count of periods returned: %v, expected: %v", coundOfPeriods, expected)
	}
}
