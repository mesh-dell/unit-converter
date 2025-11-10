package converters

import (
	"fmt"
	"strconv"
)

var MassConversionMap = map[string]float64{
	"kilogram":   1.0,
	"gram":       0.001,
	"metric ton": 1000.0,
	"pound":      0.45359237,
	"ounce":      0.028349523125,
}

func ConvertWeight(value, fromUnit, toUnit string) (float64, error) {
	fromFactor, ok := MassConversionMap[fromUnit]
	if !ok {
		return 0, fmt.Errorf("unknown from unit")
	}

	toFactor, ok := MassConversionMap[toUnit]
	if !ok {
		return 0, fmt.Errorf("unknown to unit")
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	valueInBaseUnit := fromFactor * valueFloat
	result := valueInBaseUnit / toFactor
	return result, nil
}
