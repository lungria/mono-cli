package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/lungria/mono"
)

var writer *csv.Writer

type config struct {
	token     string
	startDate time.Time
	endDate   time.Time
}

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}

	writer = csv.NewWriter(os.Stdout)

	auth := mono.NewPersonalAuth(cfg.token)
	client := mono.New(auth)
	user, err := client.User()
	if err != nil {
		log.Fatal(err)
	}
	//monobank api limit
	const maxPeriodDurationLimit = 2682000
	duration := time.Second * time.Duration(maxPeriodDurationLimit)
	calc := newPeriodCalculator(cfg.startDate, cfg.endDate, duration)
	for calc.Next() {
		statements, err := client.Statement(user.Accounts[0].ID, calc.current.From, calc.current.To)
		apiRateLimit := time.After(time.Second * time.Duration(60))
		if err != nil {
			log.Fatal(err)
		}

		statements = Sort(statements)
		saveStatements(statements)
		<-apiRateLimit
	}
}

func parseConfig() (config, error) {
	token := os.Getenv("MONO_APIKEY")
	if len(token) == 0 {
		return config{}, errors.New("MONO_APIKEY is required, see https://api.monobank.ua/ to get it")
	}
	cfg := config{token: token}
	startDate := os.Getenv("MONO_STARTDATE")
	if len(startDate) != 0 {
		i, err := strconv.ParseInt(startDate, 10, 64)
		if err != nil {
			return config{}, errors.New("MONO_STARTDATE must be unix timestamp")
		}
		cfg.startDate = time.Unix(i, 0)
	} else {
		//monobank launch date https://uk.wikipedia.org/wiki/Monobank
		cfg.startDate = time.Date(2017, 11, 15, 0, 0, 0, 0, time.UTC)
	}
	endDate := os.Getenv("MONO_ENDDATE")
	if len(endDate) != 0 {
		i, err := strconv.ParseInt(endDate, 10, 64)
		if err != nil {
			return config{}, errors.New("MONO_ENDDATE must be unix timestamp")
		}
		cfg.endDate = time.Unix(i, 0)
	} else {
		cfg.endDate = time.Now().UTC()
	}

	return cfg, nil
}

var newLineRegexp = regexp.MustCompile(`\r?\n`)
var headerPrinted = false

func saveStatements(items []mono.StatementItem) {
	if len(items) == 0 {
		return
	}

	if !headerPrinted {
		//write csv header using reflection
		t := items[0]
		s := reflect.ValueOf(&t).Elem()
		typeOfT := s.Type()
		for i := 0; i < s.NumField()-1; i++ {
			fmt.Printf("%v,", typeOfT.Field(i).Name)
		}
		fmt.Printf("%v\n", typeOfT.Field(s.NumField()-1).Name)
		headerPrinted = true
	}

	csvData := make([][]string, len(items), len(items))
	for i, v := range items {
		//this whole for loop may be replaced with reflection like https://gist.github.com/justincase/5469009
		//11 - is the number of fields in mono.StatementItem
		currentLine := make([]string, 11, 11)
		currentLine[0] = v.ID
		currentLine[1] = fmt.Sprint(v.Time)
		//sometimes description contains "newline" char that breaks csv output
		currentLine[2] = newLineRegexp.ReplaceAllString(v.Description, " ")
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
