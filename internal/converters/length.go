package converters

import (
	"fmt"
	"strconv"
)

var unitConversionMap = map[string]float64{
	"meter":      1.0,
	"centimeter": 0.01,
	"millimeter": 0.001,
	"kilometer":  1000.0,
	"inch":       0.0254,
	"foot":       0.3048,
	"yard":       0.9144,
	"mile":       1609.344,
}

func ConvertLength(value string, fromUnit, toUnit string) (float64, error) {
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	if fromUnit == toUnit {
		return valueFloat, nil
	}

	fromFactor, ok := unitConversionMap[fromUnit]
	if !ok {
		return 0, fmt.Errorf("invalid from units")
	}

	toFactor, ok := unitConversionMap[toUnit]
	if !ok {
		return 0, fmt.Errorf("invalid from units")
	}

	valueInBaseUnit := fromFactor * valueFloat
	result := valueInBaseUnit / toFactor
	return result, nil
}
