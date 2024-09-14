package evaluator

import "Ra/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, minimum=2",
					len(args))
			}

			maxNum := 0
			for i, arg := range args {
				num, ok := arg.(*object.Integer)
				if !ok {
					return newError("argument to `max` not supported, got %s", args[i].Type())
				}

				if num.Value > int64(maxNum) {
					maxNum = int(num.Value)
				}

			}
			return &object.Integer{Value: int64(maxNum)}
		},
	},
}
