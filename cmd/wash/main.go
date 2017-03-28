package main

import (
	"flag"
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/incant"
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

	incant.NewFunction("validateCoin").
		In(vendingFile).
		WithTestIn(vendingTestFile).
		WhenGiven(invalidCoin).
		Returns(zeroValue, notOk).
		Build()
}
