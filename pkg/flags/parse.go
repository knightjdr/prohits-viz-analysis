// Package flags handles flag parsing
package flags

import (
	"os"
	"strings"
)

// Parse converts command line arguments into an interface. Argument
// can have any number of leading dashes and be separated from their
// value by an equal sign or white space. Equal signs must not be present
// anywhere other than between the argument and its value.
func Parse() map[string]interface{} {
	args := make(map[string]interface{}, 0)
	for i := 1; i < len(os.Args); i++ {
		arg := strings.Trim(os.Args[i], "-")
		var flag string
		var parameter string
		if strings.Contains(arg, "=") {
			splitArg := strings.Split(arg, "=")
			flag = splitArg[0]
			parameter = splitArg[1]
		} else {
			flag = arg
			parameter = os.Args[i+1]
			i++
		}
		args[flag] = parameter
	}
	return args
}
