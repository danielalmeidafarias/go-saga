package main

import (
	"github.com/danielalmeidafarias/go-saga/internal/user"
)

func main() {
	userApp := user.NewUserApp()
	userApp.Run()
}
