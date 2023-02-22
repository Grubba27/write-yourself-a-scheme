package walker

import (
	"fmt"
	"strconv"
	"write-yourself-a-scheme/lexer"
	"write-yourself-a-scheme/parser"
)

var builtins = map[string]func(value []parser.Value, ctx map[string]any) any{}

func Initialize() {

	builtins["if"] = func(args []parser.Value, ctx map[string]any) any {
		condition := AstWalk(args[0], ctx)
		then := args[1]
		else_ := args[2]

		if condition.(bool) == true {
			return AstWalk(then, ctx)
		} else {
			return AstWalk(else_, ctx)
		}
	}

	//"fold": func(args []parser.Value, ctx map[string]any) any {
	//	fn := Evaluate(args[0], ctx)
	//	init := Evaluate(args[1], ctx)
	//	return nil
	//},

	// Math
	builtins["+"] = func(args []parser.Value, ctx map[string]any) any {
		i := AstWalk(args[0], ctx).(int64)
		for _, arg := range args[1:] {
			i += AstWalk(arg, ctx).(int64)
		}
		return i
	}
	builtins["-"] = func(args []parser.Value, ctx map[string]any) any {
		i := AstWalk(args[0], ctx).(int64)
		for _, arg := range args[1:] {
			i -= AstWalk(arg, ctx).(int64)
		}
		return i
	}
	builtins["*"] = func(args []parser.Value, ctx map[string]any) any {
		i := AstWalk(args[0], ctx).(int64)

		for _, arg := range args[1:] {
			i *= AstWalk(arg, ctx).(int64)
		}
		return i
	}
	builtins["/"] = func(args []parser.Value, ctx map[string]any) any {
		i := AstWalk(args[0], ctx).(int64)
		for _, arg := range args[1:] {
			i /= AstWalk(arg, ctx).(int64)
		}
		return i
	}

}
func AstWalk(v parser.Value, ctx map[string]any) any {
	if v.Kind == parser.LiteralKind {
		t := *v.Literal
		switch t.Kind {

		case lexer.IntegerToken:
			i, err := strconv.ParseInt(t.Value, 10, 64)
			if err != nil {
				fmt.Println("Expected an int but received: ", t.Value)
				panic(err)
			}
			return i

		case lexer.IdentifierToken:
			return t.Value
		}
	}
	return EvaluateValue(*v.List, ctx)
}
func EvaluateValue(ast []parser.Value, ctx map[string]any) any {
	fnName := (*ast[0].Literal).Value
	if fn, ok := builtins[fnName]; ok {
		result := fn(ast[1:], ctx)
		return result
	}
	panic("whaat, not yet implemented")
	//fn := ctx[fnName]
	//
	//var args []any
	//for _, uArgs := range (*v.List)[1:] {
	//	args = append(args, Evaluate(uArgs, ctx))
	//}
	//return fn(args)
	return nil
}
