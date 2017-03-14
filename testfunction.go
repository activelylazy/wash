package wash

import (
	"errors"
	"fmt"
	"strings"

	"github.com/activelylazy/wash/syntax"
)

// AppendTestFunctionCall appends code to a (test) function which verifies a call to a function
func AppendTestFunctionCall(fn Function, calledFunction Function, expectedValues []string) error {
	if len(calledFunction.ReturnValues) != len(expectedValues) {
		return errors.New("Number of expected values is not the same as number of values returned from function")
	}

	returnValueNames := getNames(calledFunction.ReturnValues)

	fn.Append(strings.Join(returnValueNames, ", ") + ` := ` + calledFunction.FunctionName + `("x")`)

	for i, varName := range returnValueNames {
		fn.Append(fmt.Sprintf(`if %v {
            t.Errorf("Expected %v but got %%v", %s)
        }`, defineComparison(varName, expectedValues[i]), expectedValues[i], varName))
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
