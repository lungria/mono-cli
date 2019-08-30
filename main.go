package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lungria/mono"
)

var writer *csv.Writer

func main() {
	token := os.Getenv("MONO_APIKEY")

	writer = csv.NewWriter(os.Stdout)

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
	for calc.Next() {
		statements, err := client.Statement(user.Accounts[0].ID, calc.current.From, calc.current.To)
		apiRateLimit := time.After(time.Second * time.Duration(60))
		if err != nil {
			log.Fatal(err)
		}

		saveStatements(statements)
		<-apiRateLimit
	}
}

func saveStatements(items []mono.StatementItem) {
	//todo write header using reflection
	csvData := make([][]string, len(items), len(items))
	for i, v := range items {
		//11 - is the number of fields in mono.StatementItem
		//todo replace hardcoded const with reflection
		currentLine := make([]string, 11, 11)
		currentLine[0] = v.ID
		currentLine[1] = fmt.Sprint(v.Time)
		currentLine[2] = v.Description
		currentLine[3] = fmt.Sprint(v.MCC)
		currentLine[4] = fmt.Sprint(v.Hold)
		currentLine[5] = fmt.Sprint(v.Amount)
		currentLine[6] = fmt.Sprint(v.OperationAmount)
		currentLine[7] = fmt.Sprint(v.CurrencyCode)
		currentLine[8] = fmt.Sprint(v.CommissionRate)
		currentLine[9] = fmt.Sprint(v.CashbackAmount)
		currentLine[10] = fmt.Sprint(v.Balance)
		csvData[i] = currentLine
	}

	err := writer.WriteAll(csvData)
	if err != nil {
		log.Fatal(err)
	}
}
