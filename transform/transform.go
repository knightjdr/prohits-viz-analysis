// Package transform will adjust the prey values to the user requirements
//
/* Preys can be adjusted by (in this order):
** 1: control values (must be a pipe-separated list)
** 2: prey length
** 3: normalized across baits
** 4: log transformed
 */
package transform

func Preys(data []map[string]interface{}, control string) ([]map[string]interface{}, error) {
	var err error
	// control subtraction
	transformed, err := ControlSubtraction(data, control)
	return transformed, err
}
