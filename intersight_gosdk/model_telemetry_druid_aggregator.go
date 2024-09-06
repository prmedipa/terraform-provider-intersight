/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18369
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// TelemetryDruidAggregator - struct for TelemetryDruidAggregator
type TelemetryDruidAggregator struct {
	TelemetryDruidAnyAggregator             *TelemetryDruidAnyAggregator
	TelemetryDruidCountAggregator           *TelemetryDruidCountAggregator
	TelemetryDruidFilteredAggregator        *TelemetryDruidFilteredAggregator
	TelemetryDruidFirstLastAggregator       *TelemetryDruidFirstLastAggregator
	TelemetryDruidMinMaxAggregator          *TelemetryDruidMinMaxAggregator
	TelemetryDruidStringAnyAggregator       *TelemetryDruidStringAnyAggregator
	TelemetryDruidStringFirstLastAggregator *TelemetryDruidStringFirstLastAggregator
	TelemetryDruidSumAggregator             *TelemetryDruidSumAggregator
	TelemetryDruidThetaSketchAggregator     *TelemetryDruidThetaSketchAggregator
}

// TelemetryDruidAnyAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidAnyAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidAnyAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidAnyAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidAnyAggregator: v,
	}
}

// TelemetryDruidCountAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidCountAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidCountAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidCountAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidCountAggregator: v,
	}
}

// TelemetryDruidFilteredAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidFilteredAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidFilteredAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidFilteredAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidFilteredAggregator: v,
	}
}

// TelemetryDruidFirstLastAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidFirstLastAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidFirstLastAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidFirstLastAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidFirstLastAggregator: v,
	}
}

// TelemetryDruidMinMaxAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidMinMaxAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidMinMaxAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidMinMaxAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidMinMaxAggregator: v,
	}
}

// TelemetryDruidStringAnyAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidStringAnyAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidStringAnyAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidStringAnyAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidStringAnyAggregator: v,
	}
}

// TelemetryDruidStringFirstLastAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidStringFirstLastAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidStringFirstLastAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidStringFirstLastAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidStringFirstLastAggregator: v,
	}
}

// TelemetryDruidSumAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidSumAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidSumAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidSumAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidSumAggregator: v,
	}
}

// TelemetryDruidThetaSketchAggregatorAsTelemetryDruidAggregator is a convenience function that returns TelemetryDruidThetaSketchAggregator wrapped in TelemetryDruidAggregator
func TelemetryDruidThetaSketchAggregatorAsTelemetryDruidAggregator(v *TelemetryDruidThetaSketchAggregator) TelemetryDruidAggregator {
	return TelemetryDruidAggregator{
		TelemetryDruidThetaSketchAggregator: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidAggregator) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'count'
	if jsonDict["type"] == "count" {
		// try to unmarshal JSON data into TelemetryDruidCountAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidCountAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidCountAggregator, return on the first match
		} else {
			dst.TelemetryDruidCountAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidCountAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleAny'
	if jsonDict["type"] == "doubleAny" {
		// try to unmarshal JSON data into TelemetryDruidAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleFirst'
	if jsonDict["type"] == "doubleFirst" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleLast'
	if jsonDict["type"] == "doubleLast" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleMax'
	if jsonDict["type"] == "doubleMax" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleMin'
	if jsonDict["type"] == "doubleMin" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'doubleSum'
	if jsonDict["type"] == "doubleSum" {
		// try to unmarshal JSON data into TelemetryDruidSumAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidSumAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSumAggregator, return on the first match
		} else {
			dst.TelemetryDruidSumAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidSumAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'filtered'
	if jsonDict["type"] == "filtered" {
		// try to unmarshal JSON data into TelemetryDruidFilteredAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFilteredAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFilteredAggregator, return on the first match
		} else {
			dst.TelemetryDruidFilteredAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFilteredAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatAny'
	if jsonDict["type"] == "floatAny" {
		// try to unmarshal JSON data into TelemetryDruidAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatFirst'
	if jsonDict["type"] == "floatFirst" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatLast'
	if jsonDict["type"] == "floatLast" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatMax'
	if jsonDict["type"] == "floatMax" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatMin'
	if jsonDict["type"] == "floatMin" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'floatSum'
	if jsonDict["type"] == "floatSum" {
		// try to unmarshal JSON data into TelemetryDruidSumAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidSumAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSumAggregator, return on the first match
		} else {
			dst.TelemetryDruidSumAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidSumAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longAny'
	if jsonDict["type"] == "longAny" {
		// try to unmarshal JSON data into TelemetryDruidAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longFirst'
	if jsonDict["type"] == "longFirst" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longLast'
	if jsonDict["type"] == "longLast" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longMax'
	if jsonDict["type"] == "longMax" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longMin'
	if jsonDict["type"] == "longMin" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'longSum'
	if jsonDict["type"] == "longSum" {
		// try to unmarshal JSON data into TelemetryDruidSumAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidSumAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSumAggregator, return on the first match
		} else {
			dst.TelemetryDruidSumAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidSumAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'stringAny'
	if jsonDict["type"] == "stringAny" {
		// try to unmarshal JSON data into TelemetryDruidStringAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidStringAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidStringAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidStringAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidStringAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'stringFirst'
	if jsonDict["type"] == "stringFirst" {
		// try to unmarshal JSON data into TelemetryDruidStringFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidStringFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidStringFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidStringFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidStringFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'stringLast'
	if jsonDict["type"] == "stringLast" {
		// try to unmarshal JSON data into TelemetryDruidStringFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidStringFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidStringFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidStringFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidStringFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'thetaSketch'
	if jsonDict["type"] == "thetaSketch" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidThetaSketchAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidAnyAggregator'
	if jsonDict["type"] == "telemetry.DruidAnyAggregator" {
		// try to unmarshal JSON data into TelemetryDruidAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidCountAggregator'
	if jsonDict["type"] == "telemetry.DruidCountAggregator" {
		// try to unmarshal JSON data into TelemetryDruidCountAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidCountAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidCountAggregator, return on the first match
		} else {
			dst.TelemetryDruidCountAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidCountAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidFilteredAggregator'
	if jsonDict["type"] == "telemetry.DruidFilteredAggregator" {
		// try to unmarshal JSON data into TelemetryDruidFilteredAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFilteredAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFilteredAggregator, return on the first match
		} else {
			dst.TelemetryDruidFilteredAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFilteredAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidFirstLastAggregator'
	if jsonDict["type"] == "telemetry.DruidFirstLastAggregator" {
		// try to unmarshal JSON data into TelemetryDruidFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidMinMaxAggregator'
	if jsonDict["type"] == "telemetry.DruidMinMaxAggregator" {
		// try to unmarshal JSON data into TelemetryDruidMinMaxAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidMinMaxAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidMinMaxAggregator, return on the first match
		} else {
			dst.TelemetryDruidMinMaxAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidMinMaxAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidStringAnyAggregator'
	if jsonDict["type"] == "telemetry.DruidStringAnyAggregator" {
		// try to unmarshal JSON data into TelemetryDruidStringAnyAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidStringAnyAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidStringAnyAggregator, return on the first match
		} else {
			dst.TelemetryDruidStringAnyAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidStringAnyAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidStringFirstLastAggregator'
	if jsonDict["type"] == "telemetry.DruidStringFirstLastAggregator" {
		// try to unmarshal JSON data into TelemetryDruidStringFirstLastAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidStringFirstLastAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidStringFirstLastAggregator, return on the first match
		} else {
			dst.TelemetryDruidStringFirstLastAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidStringFirstLastAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidSumAggregator'
	if jsonDict["type"] == "telemetry.DruidSumAggregator" {
		// try to unmarshal JSON data into TelemetryDruidSumAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidSumAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSumAggregator, return on the first match
		} else {
			dst.TelemetryDruidSumAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidSumAggregator: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidThetaSketchAggregator'
	if jsonDict["type"] == "telemetry.DruidThetaSketchAggregator" {
		// try to unmarshal JSON data into TelemetryDruidThetaSketchAggregator
		err = json.Unmarshal(data, &dst.TelemetryDruidThetaSketchAggregator)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidThetaSketchAggregator, return on the first match
		} else {
			dst.TelemetryDruidThetaSketchAggregator = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidAggregator as TelemetryDruidThetaSketchAggregator: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidAggregator) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidAnyAggregator != nil {
		return json.Marshal(&src.TelemetryDruidAnyAggregator)
	}

	if src.TelemetryDruidCountAggregator != nil {
		return json.Marshal(&src.TelemetryDruidCountAggregator)
	}

	if src.TelemetryDruidFilteredAggregator != nil {
		return json.Marshal(&src.TelemetryDruidFilteredAggregator)
	}

	if src.TelemetryDruidFirstLastAggregator != nil {
		return json.Marshal(&src.TelemetryDruidFirstLastAggregator)
	}

	if src.TelemetryDruidMinMaxAggregator != nil {
		return json.Marshal(&src.TelemetryDruidMinMaxAggregator)
	}

	if src.TelemetryDruidStringAnyAggregator != nil {
		return json.Marshal(&src.TelemetryDruidStringAnyAggregator)
	}

	if src.TelemetryDruidStringFirstLastAggregator != nil {
		return json.Marshal(&src.TelemetryDruidStringFirstLastAggregator)
	}

	if src.TelemetryDruidSumAggregator != nil {
		return json.Marshal(&src.TelemetryDruidSumAggregator)
	}

	if src.TelemetryDruidThetaSketchAggregator != nil {
		return json.Marshal(&src.TelemetryDruidThetaSketchAggregator)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidAggregator) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TelemetryDruidAnyAggregator != nil {
		return obj.TelemetryDruidAnyAggregator
	}

	if obj.TelemetryDruidCountAggregator != nil {
		return obj.TelemetryDruidCountAggregator
	}

	if obj.TelemetryDruidFilteredAggregator != nil {
		return obj.TelemetryDruidFilteredAggregator
	}

	if obj.TelemetryDruidFirstLastAggregator != nil {
		return obj.TelemetryDruidFirstLastAggregator
	}

	if obj.TelemetryDruidMinMaxAggregator != nil {
		return obj.TelemetryDruidMinMaxAggregator
	}

	if obj.TelemetryDruidStringAnyAggregator != nil {
		return obj.TelemetryDruidStringAnyAggregator
	}

	if obj.TelemetryDruidStringFirstLastAggregator != nil {
		return obj.TelemetryDruidStringFirstLastAggregator
	}

	if obj.TelemetryDruidSumAggregator != nil {
		return obj.TelemetryDruidSumAggregator
	}

	if obj.TelemetryDruidThetaSketchAggregator != nil {
		return obj.TelemetryDruidThetaSketchAggregator
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidAggregator struct {
	value *TelemetryDruidAggregator
	isSet bool
}

func (v NullableTelemetryDruidAggregator) Get() *TelemetryDruidAggregator {
	return v.value
}

func (v *NullableTelemetryDruidAggregator) Set(val *TelemetryDruidAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidAggregator(val *TelemetryDruidAggregator) *NullableTelemetryDruidAggregator {
	return &NullableTelemetryDruidAggregator{value: val, isSet: true}
}

func (v NullableTelemetryDruidAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
