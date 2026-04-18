package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	prices := make([]float64, len(strings))
	for i, str := range strings {
		floatPrice, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("failed to parse price")
		}

		prices[i] = floatPrice
	}

	return prices, nil
}
