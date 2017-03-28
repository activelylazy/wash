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

	// displayFile, err := washer.CreateFile("vending/display.go", "vending")
	// if err != nil {
	// 	log.Fatalf("Error creating file: %v", err)
	// }

	// displayFile.AddStruct(displayFile,
	// 	"Display",
	// 	[]syntax.Field{syntax.NewField("message", "string")})

	invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")

	validateCoinFunction := vendingFile.AddFunction("validateCoin",
		[]syntax.Field{syntax.NewField("s", "string")},
		[]syntax.Field{syntax.NewField("value", "int"), syntax.NewField("ok", "bool")},
		[]string{"0", "false"})

	vendingTestFile, err := washer.CreateFile("vending/vending_test.go", "vending")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	if err = wash.WriteFunctionCallTest(vendingTestFile, validateCoinFunction, []wash.DomainConcept{invalidCoin}, []string{"0", "false"}); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
