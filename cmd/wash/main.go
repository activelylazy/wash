package main

import (
	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/domain"
	"github.com/activelylazy/wash/incant"
)

func main() {

	// basePath := flag.String("base", "", "the base path to read/write source code")

	// flag.Parse()

	// if *basePath == "" {
	// 	log.Fatalf("Base path is required")
	// }

	// washer, err := wash.NewWasher(*basePath)
	// if err != nil {
	// 	log.Fatalf("Error parsing: %v", err)
	// }

	wash.SetBasePath("github.com/activelylazy/generated-vending")
	vendingFile := wash.CreateFile("vending/vending.go", "vending")
	vendingTestFile := wash.CreateFile("vending/vending_test.go", "vending")

	coin := domain.NewType("coin", "string")
	coinValue := domain.NewType("value", "int")
	okType := domain.NewType("ok", "bool")

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
