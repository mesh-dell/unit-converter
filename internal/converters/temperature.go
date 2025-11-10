package converters

import (
	"fmt"
	"strconv"
)

func ConvertTemp(value, fromUnit, toUnit string) (float64, error) {
	valueFloat, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, err
	}

	var valueCelsius float64

	switch fromUnit {
	case "celsius":
		valueCelsius = valueFloat
	case "fahrenheit":
		valueCelsius = (valueFloat - 32.0) * 5.0 / 9.0
	case "kelvin":
		valueCelsius = valueFloat - 273.15
	default:
		return 0, fmt.Errorf("unknown from unit")
	}

	var result float64
	switch toUnit {
	case "celsius":
		result = valueCelsius
	case "fahrenheit":
		result = (valueCelsius * 9.0 / 5.0) + 32.0
	case "kelvin":
		result = valueCelsius + 273.15
	default:
		return 0, fmt.Errorf("unknown to unit")
	}

	return result, nil
}
