package evaluate

import (
	"github.com/JWSch4fer/interpreter/object"
)

/*
callUserFunction applies a user-defined function to the given arguments.

When callUserFunction is defined alongside your builtins,
the compiler sees that builtins (which are initialized at package level)
depend on a function that in turn calls Eval and other functions that
reference builtins.
Moving callUserFunction into its own file delays its binding until runtime,
which decouples these initialization-time dependencies.
*/
func callUserFunction(fn *object.Function, args []object.Object) object.Object {
	extendedEnv := extendFunctionEnv(fn, args)
	evaluated := Eval(fn.Body, extendedEnv)
	return unwrapReturnValue(evaluated)
}
