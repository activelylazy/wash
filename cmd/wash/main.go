package main

import (
	"log"

	"github.com/activelylazy/wash"
)

func main() {
	// fset := token.NewFileSet()
	// path := "C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\generated-vending"
	// _, err := parser.ParseDir(fset, path, nil, parser.AllErrors)
	// if err != nil {
	// 	log.Fatalf("Error parsing: %v", err)
	// }

	washer, err := wash.NewWasher("C:\\Users\\Dave\\Documents\\Projects\\Go\\src\\github.com\\activelylazy\\generated-vending")
	if err != nil {
		log.Fatalf("Error parsing: %v", err)
	}

	vendingFile, err := wash.NewCreateFileRequest("vending/vending.go", "vending").Create(washer)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	// invalidCoin := washer.NewDomainConcept("invalidCoin", "string", "x")

	washer.Edit(vendingFile).
		AddFunction("validateCoin").
		WithParameter("s", "string").
		ReturningTypes("int", "bool").
		ThatReturns("0, false").
		Build()
	//.
	// WhichWhenGiven(invalidCoin).
	// Returns(0, false)

	// packageName := "vending"
	// // fileName := "vending.go"
	// f := newFile(packageName)
	// addImport(f, "", "\"github.com/moo\"")
	// addFunction(f, "validateCoin", []field{
	// 	field{
	// 		fieldName: "s",
	// 		typeName:  "string",
	// 	}},
	// 	[]field{
	// 		field{
	// 			fieldName: "",
	// 			typeName:  "int",
	// 		},
	// 		field{
	// 			fieldName: "",
	// 			typeName:  "bool",
	// 		},
	// 	},
	// 	[]ast.Stmt{
	// 		&ast.ReturnStmt{
	// 			Results: []ast.Expr{
	// 				newBasicLit("0"),
	// 				newBasicLit("false"),
	// 			},
	// 		},
	// 	})

	// printer.Fprint(os.Stdout, fset, f)
}
