package main

import (
	"flag"
	"log"

	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/operations"
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

	vendingFile, err := operations.NewCreateFileRequest("vending/vending.go", "vending").Apply(washer)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	// invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")

	operations.NewAddFunctionRequest(vendingFile, "validateCoin",
		[]syntax.Field{syntax.NewField("s", "string")},
		[]syntax.Field{syntax.NewField("", "int"), syntax.NewField("", "bool")},
		[]string{"0", "false"}).
		Apply(washer)

	vendingTestFile, err := operations.NewCreateFileRequest("vending/vending_test.go", "vending").
		Apply(washer)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	operations.NewAddImportRequest(vendingTestFile, "", "testing").Apply(washer)

	fn := operations.NewAddFunctionRequest(vendingTestFile, "TestValidateCoinReturnsZeroFalseForInvalidCoin",
		[]syntax.Field{syntax.NewField("t", "*testing.T")},
		[]syntax.Field{},
		[]string{}).
		Apply(washer)

	stmt, err := wash.ParseStatement(`value, ok := validateCoin("x")`)
	if err != nil {
		log.Fatalf("Error parsing statement: %v", err)
	}

	operations.NewAppendToFunctionBodyRequest(fn, stmt).Apply(washer)

}
