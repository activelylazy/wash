package wash

import (
	"errors"
	"fmt"
	"strings"

	"github.com/activelylazy/wash/syntax"
)

// WriteFunctionCallTest appends a test to a file which verifies a call to a function
func WriteFunctionCallTest(testFile *File, calledFunction Function, expectedValues []string) error {
	fn := testFile.AddFunction("TestValidateCoinReturnsZeroFalseForInvalidCoin",
		[]syntax.Field{syntax.NewField("t", "*testing.T")},
		[]syntax.Field{},
		[]string{})

	if len(calledFunction.ReturnValues) != len(expectedValues) {
		return errors.New("Number of expected values is not the same as number of values returned from function")
	}

	returnValueNames := getNames(calledFunction.ReturnValues)

	fn.Append(strings.Join(returnValueNames, ", ") + ` := ` + calledFunction.FunctionName + `("x")`)

	for i, varName := range returnValueNames {
		fn.Append(fmt.Sprintf(`if %v {
            t.Errorf("Expected %s to be %v but was %%v", %s)
        }`, defineComparison(varName, expectedValues[i]), varName, expectedValues[i], varName))
	}

	return nil
}

func getNames(fields []syntax.Field) []string {
	names := make([]string, len(fields))
	for i, f := range fields {
		names[i] = f.FieldName
	}
	return names
}

func defineComparison(varName string, expectedValue string) string {
	if expectedValue == "true" {
		return varName
	}
	if expectedValue == "false" {
		return "!" + varName
	}
	return fmt.Sprintf("%s == %v", varName, expectedValue)
}
