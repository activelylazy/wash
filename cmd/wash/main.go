package main

import (
	"github.com/activelylazy/wash"
	"github.com/activelylazy/wash/domain"
	"github.com/activelylazy/wash/incant"
)

func main() {
	wash.SetBasePath("github.com/activelylazy/generated-vending")
	_ = wash.CreateFile("vending/vending.go", "vending")
	vendingTestFile := wash.CreateFile("vending/vending_test.go", "vending")

	_ = wash.FindFile("vending/vending.go")

	coin := domain.NewType("coin", "string")
	coinValue := domain.NewType("value", "int")
	okType := domain.NewType("ok", "bool")

	invalidCoin := coin.NewInstance("invalidCoin", "x")
	zeroValue := coinValue.NewInstance("zero", "0")
	notOk := okType.NewInstance("notOk", "false")

	incant.NewFunction("validateCoin").
		InFile("vending/vending.go").
		WithTestIn(vendingTestFile).
		Given(invalidCoin).
		ShouldReturn(zeroValue, notOk).
		Build()
}
