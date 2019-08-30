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
	//startDateTime := os.Getenv("MONO_STARTDATE")

	auth := mono.NewPersonalAuth(token)
	client := mono.New(auth)
	user, err := client.User()
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().UTC()
	twoDaysBefore := time.Now().Add(-time.Hour * time.Duration(48))

	statements, err := client.Statement(user.Accounts[0].ID, twoDaysBefore, now)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", statements)
}
