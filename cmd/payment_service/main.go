package main

import (
	"github.com/danielalmeidafarias/go-saga/internal/payment"
)

func main() {
	app := payment.NewPaymentApp()
	app.Run()
}
