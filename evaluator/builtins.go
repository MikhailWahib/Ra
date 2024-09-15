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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
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
			if len(args) < 1 {
				return newError("wrong number of arguments. got=%d, minimum=1",
					len(args))
			}

			switch args[0].(type) {
			case *object.Integer:
				if len(args) < 2 {
					return newError("wrong number of arguments. got=%d, minimum=2",
						len(args))
				}

				return getMax(args)

			case *object.Array:
				if len(args) > 1 {
					return newError("wrong number of arguments. got=%d, want=1",
						len(args))
				}

				return getMax(args[0].(*object.Array).Elements)

			default:
				return newError("argument to `max` not supported, got %s",
					args[0].Type())
			}
		},
	},
}

func getMax(iter []object.Object) object.Object {
	maxNum := 0
	for i, el := range iter {
		num, ok := el.(*object.Integer)
		if !ok {
			return newError("argument to `max` not supported, got %s", iter[i].Type())
		}

		if num.Value > int64(maxNum) {
			maxNum = int(num.Value)
		}

	}
	return &object.Integer{Value: int64(maxNum)}
}
