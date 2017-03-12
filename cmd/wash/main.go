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

	displayFile, err := washer.CreateFile("vending/display.go", "display")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	displayFile.AddStruct(displayFile,
		"Display",
		[]syntax.Field{syntax.NewField("message", "string")})

	// invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")

	vendingFile.AddFunction("validateCoin",
		[]syntax.Field{syntax.NewField("s", "string")},
		[]syntax.Field{syntax.NewField("", "int"), syntax.NewField("", "bool")},
		[]string{"0", "false"})

	vendingTestFile, err := washer.CreateFile("vending/vending_test.go", "vending")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	vendingTestFile.AddImport("", "testing")

	fn := vendingTestFile.AddFunction("TestValidateCoinReturnsZeroFalseForInvalidCoin",
		[]syntax.Field{syntax.NewField("t", "*testing.T")},
		[]syntax.Field{},
		[]string{})

	fn.Append(`value, ok := validateCoin("x")`)

	fn.Append(`if value != 0 {
		t.Errorf("Expected 0 but got %d", value)
	}`)

	fn.Append(`if ok {
		t.Errorf("Expected ok to be false but got %v", ok)
	}`)
}
