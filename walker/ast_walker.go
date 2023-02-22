package walker

import "write-yourself-a-scheme/parser"

var builtins = map[string]func(value []parser.Value, ctx map[string]any) any{}

func Initialize() {

	builtins["if"] = func(args []parser.Value, ctx map[string]any) any {
		condition := EvaluateValue(args[0], ctx)
		then := args[1]
		else_ := args[2]

		if condition.(bool) == true {
			return EvaluateValue(then, ctx)
		} else {
			return EvaluateValue(else_, ctx)
		}
	}

	//"fold": func(args []parser.Value, ctx map[string]any) any {
	//	fn := Evaluate(args[0], ctx)
	//	init := Evaluate(args[1], ctx)
	//	return nil
	//},

	// Math
	builtins["+"] = func(args []parser.Value, ctx map[string]any) any {
		var i int
		for _, arg := range args {
			i += EvaluateValue(arg, ctx).(int)
		}
		return i
	}
	builtins["-"] = func(args []parser.Value, ctx map[string]any) any {
		var i int
		for _, arg := range args {
			i -= EvaluateValue(arg, ctx).(int)
		}
		return i
	}
	builtins["*"] = func(args []parser.Value, ctx map[string]any) any {
		var i int
		for _, arg := range args {
			i *= EvaluateValue(arg, ctx).(int)
		}
		return i
	}
	builtins["/"] = func(args []parser.Value, ctx map[string]any) any {
		var i int
		for _, arg := range args {
			i /= EvaluateValue(arg, ctx).(int)
		}
		return i
	}

}

func EvaluateValue(v parser.Value, ctx map[string]any) any {
	if v.Kind == parser.LiteralKind {
		r := *v.Literal
		println(r.Value)
		return r.Value
	}

	fnName := (*(*v.List)[0].Literal).Value
	if fn, ok := builtins[fnName]; ok {
		return fn((*v.List)[1:], ctx)
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
