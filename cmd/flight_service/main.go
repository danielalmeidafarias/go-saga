package main

import (
	"github.com/danielalmeidafarias/go-saga/internal/flight"
)

func main() {
	flightApp := flight.NewFlightApp()
	flightApp.Run()
}
