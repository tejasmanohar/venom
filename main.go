package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/tejasmanohar/govenmo"
)

func main() {
	accessToken := flag.String("access-token", "", "Venmo OAuth access token")
	recipientUsername := flag.String("recipient", "", "Recipient's Venmo username")
	amount := flag.Int("amount", 0, "Total amount in cents")
	increment := flag.Int("increment", 0, "Increment in cents")
	description := flag.String("description", "lol", "payment description")
	audience := flag.String("audience", "public", "options: public, friends, private")
	method := flag.String("method", "", "pay or charge")
	flag.Parse()

	shouldCharge := false

	if *accessToken == "" {
		log.Fatal("--access-token is required")
	}

	if *recipientUsername == "" {
		log.Fatal("--recipient-username is required")
	}

	if *amount == 0 {
		log.Fatal("--amount is required")
	}

	if *amount <= 0 {
		log.Fatal("--amount should be > 0")
	}

	if *audience != "public" && *audience != "private" && *audience != "friends" {
		log.Fatal("--audience should be one of (public, friends, private")
	}

	if *method == "charge" {
		shouldCharge = true
	} else if *method != "pay" {
		log.Fatal("--method should be pay or charge")
	}

	if *increment < 0 {
		log.Fatal("--increment should be > 0")
	}

	if *increment == 0 {
		increment = amount
	}

	me := govenmo.Account{AccessToken: *accessToken}
	app := venom{account: &me}
	app.Request(request{
		username:     *recipientUsername,
		amount:       *amount,
		increment:    *increment,
		desc:         *description,
		audience:     *audience,
		shouldCharge: shouldCharge,
	})
}
