package main

import (
	"flag"
	"log"

	"github.com/activelylazy/wash"
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

	vendingTestFile, err := wash.NewCreateFileRequest("vending/vending_test.go", "vending").Create(washer)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	wash.NewAddImportRequest(vendingTestFile, "", "testing").Add(washer)

	fn := wash.NewAddFunctionRequest(vendingTestFile, "TestValidateCoinReturnsZeroFalseForInvalidCoin",
		[]wash.Field{wash.NewField("t", "*testing.T")},
		[]wash.Field{},
		[]string{}).
		Add(washer)

	wash.NewAppendToFunctionBodyRequest(fn, wash.NewDefineAssignStmt([]string{"value", "ok"}, wash.NewCallExpr("validateCoin", wash.NewBasicLit("\"x\"")))).Add(washer)
}
