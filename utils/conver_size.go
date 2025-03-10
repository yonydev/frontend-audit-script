package utils

import "errors"


func ConvertSize(value int64, fromUnit string, toUnit string) (float64, error) {
	units := map[string]int64{
		"B": 1,
		"KB": 1024,
		"MB": 1024 * 1024,
		"GB": 1024 * 1024 * 1024,
	}

	fromFactor, okFrom := units[fromUnit]
	toFactor, okTo := units[toUnit]
	if !okFrom || !okTo {
		return 0, errors.New("invalid unit provided")
	}

	// Convert using float to avoid truncation
	bytes := float64(value * fromFactor)
	converted := bytes / float64(toFactor)

	return converted, nil
}
