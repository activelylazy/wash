package main

import (
	"log"

	"github.com/activelylazy/wash"
)

func main() {

	washer, err := wash.NewWasher("C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\generated-vending")
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	vendingFile, err := wash.NewCreateFileRequest("vending/vending.go", "vending").Create(washer)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	// invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")

	wash.NewAddFunctionRequest(vendingFile, "validateCoin",
		[]wash.Field{wash.NewField("s", "string")},
		[]wash.Field{wash.NewField("", "int"), wash.NewField("", "bool")},
		[]string{"0", "false"}).
		Add(washer)

}
