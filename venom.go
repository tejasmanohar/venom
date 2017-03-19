package main

import (
	"math"

	log "github.com/Sirupsen/logrus"
	"github.com/tejasmanohar/govenmo"
	"math/rand"
	"time"
	"fmt"
)

type venom struct {
	account *govenmo.Account
}

type request struct {
	// username is the payment recipient's username
	username string
	// amount is the total charge amount in cents
	amount int
	// increment is the payment increment amount in cents (e.g. 100 1c increments in $1)
	increment int
	// description is the charge description (default: "lol")
	desc string
	// audience is the venmo audience
	audience string
	// if true, charge instead of paying.
	shouldCharge bool
}

func (v *venom) Request(req request) error {
	rand.Seed(time.Now().Unix())
	recipient := govenmo.Target{User: govenmo.User{Username: req.username}}

	amt := toFixed(float64(req.increment)/100, 2)
	if req.shouldCharge {
		amt = -amt
	}

	for i := 0; i < req.amount; i += req.increment {
		_, err := v.account.PayOrCharge(recipient, amt, fmt.Sprintf("%s %d", req.desc, rand.Intn(500)), req.audience)
		if err != nil {
			log.WithError(err).Errorf("failed to pay %s %.2f", req.username, amt)
			return err
		}
		log.Infof("paid %s %.2f", req.username, amt)
		time.Sleep(30 * time.Second)
	}

	return nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
