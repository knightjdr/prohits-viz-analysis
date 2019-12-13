package flags

// SetFloat sets the field value from a command line argument.
func SetFloat(field string, args map[string]interface{}, defaultValue float64) float64 {
	var value float64

	if args[field] != nil {
		value = convertFloat(args[field])
	} else {
		value = defaultValue
	}

	return value
}

// SetInt sets the field value from a command line argument.
func SetInt(field string, args map[string]interface{}, defaultValue int) int {
	var value int

	if args[field] != nil {
		value = convertInt(args[field])
	} else {
		value = defaultValue
	}

	return value
}

// SetString sets the field value from a command line argument.
func SetString(field string, args map[string]interface{}, defaultValue string) string {
	var value string

	if args[field] != nil {
		value = convertString(args[field])
	} else {
		value = defaultValue
	}

	return value
}
