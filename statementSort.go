package main

import (
	"math/rand"

	"github.com/lungria/mono"
)

// Sort sorts statements in ascending order based by time using quicksort
func Sort(a []mono.StatementItem) []mono.StatementItem {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Time < a[right].Time {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	Sort(a[:left])
	Sort(a[left+1:])

	return a
}
