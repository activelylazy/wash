package main

import (
	"flag"
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/domain"
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

	coin := domain.NewDomainType("coin", "string")
	coinValue := domain.NewDomainType("value", "int")
	okType := domain.NewDomainType("ok", "bool")

	invalidCoin := coin.NewInstance("invalidCoin", "x")
	zeroValue := coinValue.NewInstance("zero", "0")
	notOk := okType.NewInstance("notOk", "false")

	incant.NewFunction("validateCoin").
		In(vendingFile).
		WithTestIn(vendingTestFile).
		Given(invalidCoin).
		ShouldReturn(zeroValue, notOk).
		Build()
}
