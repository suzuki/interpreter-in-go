package evaluator

import (
	"github.com/suzuki/interpreter-in-go/src/monkey/ast"
	"github.com/suzuki/interpreter-in-go/src/monkey/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statement
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Literal
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
