package main

import (
	"flag"
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/syntax"
)

func main() {

	basePath := flag.String("base", "", "the base path to read/write source code")

	flag.Parse()

	if *basePath == "" {
		log.Fatalf("Base path is required")
	}

	washer, err := wash.NewWasher(*basePath)
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	vendingFile, err := washer.CreateFile("vending/vending.go", "vending")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	vendingTestFile, err := washer.CreateFile("vending/vending_test.go", "vending")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")
	zeroValue := washer.NewDomainConcept("value", "int", "0")
	notOk := washer.NewDomainConcept("ok", "bool", "false")

	validateCoinFunction := vendingFile.AddFunction("validateCoin",
		[]syntax.Field{syntax.NewField("s", "string")},
		[]wash.DomainConcept{zeroValue, notOk})

	if err = wash.WriteFunctionCallTest(vendingTestFile, validateCoinFunction, []wash.DomainConcept{invalidCoin}, []wash.DomainConcept{zeroValue, notOk}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
