package main

import "time"

type period struct {
	From time.Time
	To   time.Time
}

type periodCalculator struct {
	currentPeriodStartDate time.Time
	endDate                time.Time
	periodDuration         time.Duration
	disposed               bool
}

func newPeriodCalculator(startDate time.Time, endDate time.Time, periodDuration time.Duration) *periodCalculator {
	return &periodCalculator{currentPeriodStartDate: startDate, endDate: endDate, periodDuration: periodDuration}
}

// Next returns period and flag that shows if this is a valid period or periods ended
func (calculator *periodCalculator) Next() (period, bool) {
	if calculator.disposed {
		return period{}, false
	}

	currentPeriodEnd := calculator.currentPeriodStartDate.Add(calculator.periodDuration)
	//we still have more periods to calculate
	if currentPeriodEnd.Before(calculator.endDate) {
		oldStartDate := calculator.currentPeriodStartDate
		calculator.currentPeriodStartDate = currentPeriodEnd.Add(time.Second * time.Duration(1))
		return period{oldStartDate, currentPeriodEnd}, true
	}

	// no more values
	calculator.disposed = true
	return period{calculator.currentPeriodStartDate, calculator.endDate}, true
}
