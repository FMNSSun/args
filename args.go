package args

import (
	"fmt"
	"reflect"
)

// Wrapper type for map[string]interface{}
type NamedArgs map[string]interface{}

// Wrapper type for map[string]*TypeSpec
type TypesSpec map[string]*TypeSpec

type TypeSpec struct {
	// The expected type.
	Type reflect.Type

	// Whether the argument is optional or not.
	Optional bool
}

// ArgsChecked takes a list of arguments as key value pairs and validates them
// according to the TypesSpec provided. Each even element of args
// must be of type string and is used as the name (or key) of the argument.
// If ts is nil then no type checking is performed. Returns the arguments
// as a map. The values are the odd elements of args.
func ArgsChecked(ts TypesSpec, args... interface{}) (NamedArgs, error) {
	n := len(args)

	if n % 2 != 0 {
		return nil, fmt.Errorf("Length of args must be multiple of two. Length of args was %d.", n)
	}

	m := make(map[string]interface{})


	argNo := 1

	for i := 0; i < n; i += 2 {
		k := args[i]
		val := args[i+1]

		if k == nil {
			return nil, fmt.Errorf("Unexpected nil for argument %d.", argNo)
		}

		key, ok := k.(string)

		if !ok {
			return nil, fmt.Errorf("Name for argument %d is not string.", argNo)
		}

		rt := reflect.TypeOf(val)

		if ts != nil {
			typeSpec, ok := ts[key]

			if !ok {
				return nil, fmt.Errorf("Named argument %q is not wanted. Argument: %d.", key, argNo)
			}

			if rt != typeSpec.Type {
				return nil, fmt.Errorf("Named argument %q must have type %s but has %s. Argument: %d.", key, typeSpec.Type, rt, argNo)
			}
		}

		m[key] = val

		argNo++
	}

	for name, typeSpec := range ts {
		if !typeSpec.Optional {
			_, ok := m[name]

			if !ok {
				return nil, fmt.Errorf("Named argument %q is missing (not optional).", name)
			}
		}
	}

	return m, nil
}

// Args turns the arguments into a map where each even element of args
// must be of type string and is the name (or key) of the argument. The values
// are the odd elements of args.
func Args(args... interface{}) (NamedArgs, error) {
	return ArgsChecked(nil, args...)
}
