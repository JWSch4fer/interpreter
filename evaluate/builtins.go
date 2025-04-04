package evaluate

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JWSch4fer/interpreter/object"
)

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
			case *object.Hash:
				return &object.Integer{Value: int64(len(arg.Pairs))}
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
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
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
		"read_file": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) < 2 || len(args) > 3 {
					return newError("wrong number of arguments: got %d, expected 2 or 3", len(args))
				}

				//First argument is a file path
				filePathObj, ok := args[0].(*object.String)
				if !ok {
					return newError("first argument must be a string (file path) got %s", args[0])
				}
				//second argumnet is a delimiter
				deliObj, ok := args[1].(*object.String)
				if !ok {
					return newError("second argument must be a String (delimiter), got %s", args[1])
				}

				//third argument is optional data type
				dataType := "STRING"
				if len(args) == 3 {
					dtObj, ok := args[2].(*object.String)
					if !ok {
						return newError("data type can be STRING, INT, or FLOAT")
					}
					dataType = dtObj.Value
				}

				content, err := os.ReadFile(filePathObj.Value)
				if err != nil {
					return newError("error reading file: %s", err.Error())
				}

				lines := strings.Split(string(content), "\n")
				var resultsRows []object.Object

				for _, line := range lines {
					// empty lines
					// if strings.TrimSpace(line) == "" {
					// 	resultsRows = append(resultsRows, &object.Array{Elements: []object.Object{&object.String{Value: ""}}})
					// 	// rowFields = []object.Object{&object.String{Value: ""}}
					// 	continue
					// }

					var rowFields []object.Object
					fields := strings.Split(line, deliObj.Value)
					for _, field := range fields {
						field = strings.TrimSpace(field)
						if field == "" {
							rowFields = append(rowFields, NULL)
							continue
						}

						var converted object.Object
						switch dataType {
						case "INT":
							i, err := strconv.ParseInt(field, 10, 64)
							if err != nil {
								return newError("cannot convert %q to int", field)
							}
							converted = &object.Integer{Value: i}
						case "FLOAT":
							f, err := strconv.ParseFloat(field, 32)
							if err != nil {
								return newError("cannot convert %q to float", field)
							}
							converted = &object.Float{Value: float32(f)}

						default:
							converted = &object.String{Value: field}
						}
						rowFields = append(rowFields, converted)
					}

					resultsRows = append(resultsRows, &object.Array{Elements: rowFields})
				}
				return &object.Array{Elements: resultsRows}
			},
		},
	}
}

// GetBuiltinWithGetter returns the builtinWithGetter map.
func GetBuiltinWithGetter() map[string]*object.Builtin {
	return builtinWithGetter
}
