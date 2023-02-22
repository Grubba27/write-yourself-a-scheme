package walker

import (
	"fmt"
	"strconv"
	"write-yourself-a-scheme/lexer"
	"write-yourself-a-scheme/parser"
)

var builtins = map[string]func(value []parser.Value, ctx map[string]any) any{}

func copyContext(in map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range in {
		out[k] = v
	}
	return out
}

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

	builtins["begin"] = func(args []parser.Value, ctx map[string]any) any {
		var last any
		for _, arg := range args {
			last = AstWalk(arg, ctx)
		}
		return last

	}

	builtins["let"] = func(args []parser.Value, ctx map[string]any) any {
		literalName := (*args[0].Literal).Value
		literal := (*args[1].Literal).Value
		ctx[literalName] = func(args []any, ctx map[string]any) any {
			innerCtx := copyContext(ctx)
			innerCtx[literalName] = literal
			return literal
		}

		return ctx[literalName]
	}

	builtins["fn"] = func(args []parser.Value, ctx map[string]any) any {
		fnName := (*args[0].Literal).Value
		params := *args[1].List
		body := *args[2].List
		ctx[fnName] = func(args []any, ctx map[string]any) any {
			innerCtx := copyContext(ctx)
			if len(params) != len(args) {
				panic(fmt.Sprintf("Expected %d args to `%s`, got %d", len(params), fnName, len(args)))
			}
			for i, param := range params {
				innerCtx[(*param.Literal).Value] = args[i]
			}
			return EvaluateValue(body, innerCtx)
		}
		return ctx[fnName]
	}

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
			return ctx[t.Value]
		}
	}
	return EvaluateValue(*v.List, ctx)
}
func EvaluateValue(ast []parser.Value, ctx map[string]any) any {
	fnName := (*ast[0].Literal).Value
	if builtInFn, ok := builtins[fnName]; ok {
		return builtInFn(ast[1:], ctx)
	}

	userFn := ctx[fnName].(func([]any, map[string]any) any)

	var args []any
	for _, uArgs := range ast[1:] {
		args = append(args, AstWalk(uArgs, ctx))
	}
	res := userFn(args, ctx)

	return res
}
