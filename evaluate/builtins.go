package evaluate

import "github.com/JWSch4fer/interpreter/object"

// separate environment of builtins
var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments: got %d want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `first` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
}

// Declare the new builtins map.
// This allows us to avoid initialization cycles for certain types of functions
var builtinWithGetter map[string]*object.Builtin

// init initializes the builtinWithGetter map.
func init() {
	builtinWithGetter = map[string]*object.Builtin{
		"map": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return newError("wrong number of arguments. got=%d, expected=2", len(args))
				}

				// First argument must be a function.
				fn, ok := args[0].(*object.Function)
				if !ok {
					return newError("first argument to map must be a function, got %s", args[0].Type())
				}

				// Second argument must be an array.
				arr, ok := args[1].(*object.Array)
				if !ok {
					return newError("second argument to map must be an array, got %s", args[1].Type())
				}

				// If the array is empty, return an empty array.
				if len(arr.Elements) == 0 {
					return &object.Array{Elements: []object.Object{}}
				}

				// Determine the expected element type.
				// For numeric arrays, we allow both INTEGER and FLOAT.
				var expectedType string
				firstType := arr.Elements[0].Type()
				if firstType == object.STRING_OBJ {
					expectedType = object.STRING_OBJ
				} else if firstType == object.INTEGER_OBJ || firstType == object.FLOAT_OBJ {
					expectedType = "NUMERIC"
				} else {
					return newError("map: unsupported element type %s", firstType)
				}

				// Check that every element is homogeneous.
				for _, el := range arr.Elements {
					if expectedType == object.STRING_OBJ && el.Type() != object.STRING_OBJ {
						return newError("map: array contains mixed types; expected all %s, got %s", expectedType, el.Type())
					} else if expectedType == "NUMERIC" && (el.Type() != object.INTEGER_OBJ && el.Type() != object.FLOAT_OBJ) {
						return newError("map: array contains mixed types; expected numeric types, got %s", el.Type())
					}
				}

				// Apply the function to every element.
				var results []object.Object
				for _, el := range arr.Elements {
					// Use a helper that defers the function call.
					mapped := callUserFunction(fn, []object.Object{el})
					if mapped.Type() == object.ERROR_OBJ {
						return mapped
					}
					results = append(results, mapped)
				}
				return &object.Array{Elements: results}
			},
		},
	}
}

// GetBuiltinWithGetter returns the builtinWithGetter map.
func GetBuiltinWithGetter() map[string]*object.Builtin {
	return builtinWithGetter
}
