package main

import "time"

type period struct {
	From time.Time
	To   time.Time
}

type periodCalculator struct {
	current        *period
	endDate        time.Time
	periodDuration time.Duration
	disposed       bool
}

func newPeriodCalculator(startDate time.Time, endDate time.Time, periodDuration time.Duration) *periodCalculator {
	tmp := &periodCalculator{current: &period{time.Time{}, startDate.Add(time.Second * time.Duration(-1))}, endDate: endDate, periodDuration: periodDuration}
	return tmp
}

// Next returns moves iterator and returns flag that shows if it's not ended
func (calculator *periodCalculator) Next() bool {
	if calculator.disposed {
		return false
	}
	calculator.current.To = calculator.current.To.Add(time.Second * time.Duration(1))
	currentPeriodEnd := calculator.current.To.Add(calculator.periodDuration)
	calculator.current.From = calculator.current.To

	//we still have more periods to calculate
	if currentPeriodEnd.Before(calculator.endDate) {
		calculator.current.To = currentPeriodEnd
		return true
	}

	// no more values
	calculator.disposed = true
	calculator.current.To = calculator.endDate

	return true
}

// Current returns current value of iterator. Always call Next at least once before reading from current
func (calculator *periodCalculator) Current() period {
	return period{calculator.current.From, calculator.current.To}
}
