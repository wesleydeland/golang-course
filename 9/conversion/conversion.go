package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var result []float64 = []float64{}

	for _, stringValue := range strings {
		floatValue, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, err
		}

		result = append(result, floatValue)
	}

	return result, nil
}
