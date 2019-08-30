package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lungria/mono"
)

func main() {
	token := os.Getenv("MONO_APIKEY")

	auth := mono.NewPersonalAuth(token)
	client := mono.New(auth)
	user, err := client.User()
	if err != nil {
		log.Fatal(err)
	}

	startDate := time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now().UTC()
	duration := time.Second * time.Duration(2682000)

	calc := newPeriodCalculator(startDate, endDate, duration)
	calc.Next()

	statements, err := client.Statement(user.Accounts[0].ID, calc.current.From, calc.current.To)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", statements)
}
