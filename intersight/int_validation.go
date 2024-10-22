package intersight

import (
	"fmt"
	"math"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IntBetween returns a SchemaValidateFunc which tests if the provided value
// is of type int and is between min and max (exclusive)
func IntBetweenExclusive(min, max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min || v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) (exclusive), got %d", k, min, max, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64BetweenExclusive returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is between min and max (exclusive)
func Int64BetweenExclusive(min, max int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min || v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d) (exclusive), got %d", k, min, max, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// IntAtLeast returns a SchemaValidateFunc which tests if the provided value
// is of type int and is exclusive min
func IntAtLeastExclusive(min int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min {
			errors = append(errors, fmt.Errorf("expected %s to be at least (%d) (exlusive minimum), got %d", k, min+1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64AtLeastExclusive returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is exclusive min
func Int64AtLeastExclusive(min int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v <= min {
			errors = append(errors, fmt.Errorf("expected %s to be at least (%d) (exlusive minimum), got %d", k, min+1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// IntAtMost returns a SchemaValidateFunc which tests if the provided value
// is of type int and is exclusive of max
func IntAtMostExclusive(max int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be at most (%d) (exclusive maximum), got %d", k, max-1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64AtMostExclusive returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is exclusive of max
func Int64AtMostExclusive(max int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v >= max {
			errors = append(errors, fmt.Errorf("expected %s to be at most (%d) (exclusive maximum), got %d", k, max-1, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64Between returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is between minVal and maxVal (inclusive)
func Int64Between(min, max int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v < min || v > max {
			errors = append(errors, fmt.Errorf("expected %s to be in the range (%d - %d), got %d", k, min, max, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64AtLeast returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is at least minVal (inclusive)
func Int64AtLeast(minVal int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v < minVal {
			errors = append(errors, fmt.Errorf("expected %s to be at least (%d), got %d", k, minVal, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64AtMost returns a SchemaValidateFunc which tests if the provided value
// is of type int64 and is at most maxVal (inclusive)
func Int64AtMost(maxVal int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if v > maxVal {
			errors = append(errors, fmt.Errorf("expected %s to be at most (%d), got %d", k, maxVal, v))
			return warnings, errors
		}

		return warnings, errors
	}
}

// Int64DivisibleBy returns a SchemaValidateFunc which tests if the provided value
// is of type int and is divisible by a given number
func Int64DivisibleBy(divisor int64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		v, ok := i.(int64)
		if !ok {
			errors = append(errors, fmt.Errorf("expected type of %s to be integer", k))
			return warnings, errors
		}

		if math.Mod(float64(v), float64(divisor)) != 0 {
			errors = append(errors, fmt.Errorf("expected %s to be divisible by %d, got: %v", k, divisor, i))
			return warnings, errors
		}

		return warnings, errors
	}
}
