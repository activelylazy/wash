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

	invalidCoin := washer.NewDomainConcept("InvalidCoin", "string", "x")
	ZERO := washer.NewDomainConcept("Zero", "int", "0")
	FALSE := washer.NewDomainConcept("False", "boolean", "false")

	validateCoinFunction := vendingFile.AddFunction("validateCoin",
		[]syntax.Field{syntax.NewField("s", "string")},
		[]syntax.Field{syntax.NewField("value", "int"), syntax.NewField("ok", "bool")},
		[]wash.DomainConcept{ZERO, FALSE})

	if err = wash.WriteFunctionCallTest(vendingTestFile, validateCoinFunction, []wash.DomainConcept{invalidCoin}, []wash.DomainConcept{ZERO, FALSE}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
