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
