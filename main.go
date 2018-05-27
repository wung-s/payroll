package main

import (
	"log"

	"github.com/wung-s/payroll/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
