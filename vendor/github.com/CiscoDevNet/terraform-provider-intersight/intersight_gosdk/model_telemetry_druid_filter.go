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

// TelemetryDruidFilter - struct for TelemetryDruidFilter
type TelemetryDruidFilter struct {
	TelemetryDruidAndFilter              *TelemetryDruidAndFilter
	TelemetryDruidBoundFilter            *TelemetryDruidBoundFilter
	TelemetryDruidColumnComparisonFilter *TelemetryDruidColumnComparisonFilter
	TelemetryDruidEqualityFilter         *TelemetryDruidEqualityFilter
	TelemetryDruidExpressionFilter       *TelemetryDruidExpressionFilter
	TelemetryDruidFalseFilter            *TelemetryDruidFalseFilter
	TelemetryDruidInFilter               *TelemetryDruidInFilter
	TelemetryDruidIntervalFilter         *TelemetryDruidIntervalFilter
	TelemetryDruidLikeFilter             *TelemetryDruidLikeFilter
	TelemetryDruidNotFilter              *TelemetryDruidNotFilter
	TelemetryDruidNullFilter             *TelemetryDruidNullFilter
	TelemetryDruidOrFilter               *TelemetryDruidOrFilter
	TelemetryDruidRangeFilter            *TelemetryDruidRangeFilter
	TelemetryDruidRegexFilter            *TelemetryDruidRegexFilter
	TelemetryDruidSearchFilter           *TelemetryDruidSearchFilter
	TelemetryDruidSelectorFilter         *TelemetryDruidSelectorFilter
	TelemetryDruidTrueFilter             *TelemetryDruidTrueFilter
}

// TelemetryDruidAndFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidAndFilter wrapped in TelemetryDruidFilter
func TelemetryDruidAndFilterAsTelemetryDruidFilter(v *TelemetryDruidAndFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidAndFilter: v,
	}
}

// TelemetryDruidBoundFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidBoundFilter wrapped in TelemetryDruidFilter
func TelemetryDruidBoundFilterAsTelemetryDruidFilter(v *TelemetryDruidBoundFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidBoundFilter: v,
	}
}

// TelemetryDruidColumnComparisonFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidColumnComparisonFilter wrapped in TelemetryDruidFilter
func TelemetryDruidColumnComparisonFilterAsTelemetryDruidFilter(v *TelemetryDruidColumnComparisonFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidColumnComparisonFilter: v,
	}
}

// TelemetryDruidEqualityFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidEqualityFilter wrapped in TelemetryDruidFilter
func TelemetryDruidEqualityFilterAsTelemetryDruidFilter(v *TelemetryDruidEqualityFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidEqualityFilter: v,
	}
}

// TelemetryDruidExpressionFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidExpressionFilter wrapped in TelemetryDruidFilter
func TelemetryDruidExpressionFilterAsTelemetryDruidFilter(v *TelemetryDruidExpressionFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidExpressionFilter: v,
	}
}

// TelemetryDruidFalseFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidFalseFilter wrapped in TelemetryDruidFilter
func TelemetryDruidFalseFilterAsTelemetryDruidFilter(v *TelemetryDruidFalseFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidFalseFilter: v,
	}
}

// TelemetryDruidInFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidInFilter wrapped in TelemetryDruidFilter
func TelemetryDruidInFilterAsTelemetryDruidFilter(v *TelemetryDruidInFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidInFilter: v,
	}
}

// TelemetryDruidIntervalFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidIntervalFilter wrapped in TelemetryDruidFilter
func TelemetryDruidIntervalFilterAsTelemetryDruidFilter(v *TelemetryDruidIntervalFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidIntervalFilter: v,
	}
}

// TelemetryDruidLikeFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidLikeFilter wrapped in TelemetryDruidFilter
func TelemetryDruidLikeFilterAsTelemetryDruidFilter(v *TelemetryDruidLikeFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidLikeFilter: v,
	}
}

// TelemetryDruidNotFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidNotFilter wrapped in TelemetryDruidFilter
func TelemetryDruidNotFilterAsTelemetryDruidFilter(v *TelemetryDruidNotFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidNotFilter: v,
	}
}

// TelemetryDruidNullFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidNullFilter wrapped in TelemetryDruidFilter
func TelemetryDruidNullFilterAsTelemetryDruidFilter(v *TelemetryDruidNullFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidNullFilter: v,
	}
}

// TelemetryDruidOrFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidOrFilter wrapped in TelemetryDruidFilter
func TelemetryDruidOrFilterAsTelemetryDruidFilter(v *TelemetryDruidOrFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidOrFilter: v,
	}
}

// TelemetryDruidRangeFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidRangeFilter wrapped in TelemetryDruidFilter
func TelemetryDruidRangeFilterAsTelemetryDruidFilter(v *TelemetryDruidRangeFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidRangeFilter: v,
	}
}

// TelemetryDruidRegexFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidRegexFilter wrapped in TelemetryDruidFilter
func TelemetryDruidRegexFilterAsTelemetryDruidFilter(v *TelemetryDruidRegexFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidRegexFilter: v,
	}
}

// TelemetryDruidSearchFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidSearchFilter wrapped in TelemetryDruidFilter
func TelemetryDruidSearchFilterAsTelemetryDruidFilter(v *TelemetryDruidSearchFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidSearchFilter: v,
	}
}

// TelemetryDruidSelectorFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidSelectorFilter wrapped in TelemetryDruidFilter
func TelemetryDruidSelectorFilterAsTelemetryDruidFilter(v *TelemetryDruidSelectorFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidSelectorFilter: v,
	}
}

// TelemetryDruidTrueFilterAsTelemetryDruidFilter is a convenience function that returns TelemetryDruidTrueFilter wrapped in TelemetryDruidFilter
func TelemetryDruidTrueFilterAsTelemetryDruidFilter(v *TelemetryDruidTrueFilter) TelemetryDruidFilter {
	return TelemetryDruidFilter{
		TelemetryDruidTrueFilter: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidFilter) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'and'
	if jsonDict["type"] == "and" {
		// try to unmarshal JSON data into TelemetryDruidAndFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidAndFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAndFilter, return on the first match
		} else {
			dst.TelemetryDruidAndFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidAndFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'bound'
	if jsonDict["type"] == "bound" {
		// try to unmarshal JSON data into TelemetryDruidBoundFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidBoundFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidBoundFilter, return on the first match
		} else {
			dst.TelemetryDruidBoundFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidBoundFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'columnComparison'
	if jsonDict["type"] == "columnComparison" {
		// try to unmarshal JSON data into TelemetryDruidColumnComparisonFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidColumnComparisonFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidColumnComparisonFilter, return on the first match
		} else {
			dst.TelemetryDruidColumnComparisonFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidColumnComparisonFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'equality'
	if jsonDict["type"] == "equality" {
		// try to unmarshal JSON data into TelemetryDruidEqualityFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidEqualityFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidEqualityFilter, return on the first match
		} else {
			dst.TelemetryDruidEqualityFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidEqualityFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'expression'
	if jsonDict["type"] == "expression" {
		// try to unmarshal JSON data into TelemetryDruidExpressionFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidExpressionFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidExpressionFilter, return on the first match
		} else {
			dst.TelemetryDruidExpressionFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidExpressionFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'false'
	if jsonDict["type"] == "false" {
		// try to unmarshal JSON data into TelemetryDruidFalseFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidFalseFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFalseFilter, return on the first match
		} else {
			dst.TelemetryDruidFalseFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidFalseFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'in'
	if jsonDict["type"] == "in" {
		// try to unmarshal JSON data into TelemetryDruidInFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidInFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInFilter, return on the first match
		} else {
			dst.TelemetryDruidInFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidInFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'interval'
	if jsonDict["type"] == "interval" {
		// try to unmarshal JSON data into TelemetryDruidIntervalFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidIntervalFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidIntervalFilter, return on the first match
		} else {
			dst.TelemetryDruidIntervalFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidIntervalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'like'
	if jsonDict["type"] == "like" {
		// try to unmarshal JSON data into TelemetryDruidLikeFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidLikeFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidLikeFilter, return on the first match
		} else {
			dst.TelemetryDruidLikeFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidLikeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'not'
	if jsonDict["type"] == "not" {
		// try to unmarshal JSON data into TelemetryDruidNotFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNotFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNotFilter, return on the first match
		} else {
			dst.TelemetryDruidNotFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidNotFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'null'
	if jsonDict["type"] == "null" {
		// try to unmarshal JSON data into TelemetryDruidNullFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNullFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNullFilter, return on the first match
		} else {
			dst.TelemetryDruidNullFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidNullFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'or'
	if jsonDict["type"] == "or" {
		// try to unmarshal JSON data into TelemetryDruidOrFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidOrFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidOrFilter, return on the first match
		} else {
			dst.TelemetryDruidOrFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidOrFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'range'
	if jsonDict["type"] == "range" {
		// try to unmarshal JSON data into TelemetryDruidRangeFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRangeFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRangeFilter, return on the first match
		} else {
			dst.TelemetryDruidRangeFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidRangeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'regex'
	if jsonDict["type"] == "regex" {
		// try to unmarshal JSON data into TelemetryDruidRegexFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexFilter, return on the first match
		} else {
			dst.TelemetryDruidRegexFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidRegexFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'search'
	if jsonDict["type"] == "search" {
		// try to unmarshal JSON data into TelemetryDruidSearchFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSearchFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSearchFilter, return on the first match
		} else {
			dst.TelemetryDruidSearchFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidSearchFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'selector'
	if jsonDict["type"] == "selector" {
		// try to unmarshal JSON data into TelemetryDruidSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSelectorFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSelectorFilter, return on the first match
		} else {
			dst.TelemetryDruidSelectorFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidSelectorFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'true'
	if jsonDict["type"] == "true" {
		// try to unmarshal JSON data into TelemetryDruidTrueFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidTrueFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidTrueFilter, return on the first match
		} else {
			dst.TelemetryDruidTrueFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidTrueFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidAndFilter'
	if jsonDict["type"] == "telemetry.DruidAndFilter" {
		// try to unmarshal JSON data into TelemetryDruidAndFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidAndFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidAndFilter, return on the first match
		} else {
			dst.TelemetryDruidAndFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidAndFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidBoundFilter'
	if jsonDict["type"] == "telemetry.DruidBoundFilter" {
		// try to unmarshal JSON data into TelemetryDruidBoundFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidBoundFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidBoundFilter, return on the first match
		} else {
			dst.TelemetryDruidBoundFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidBoundFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidColumnComparisonFilter'
	if jsonDict["type"] == "telemetry.DruidColumnComparisonFilter" {
		// try to unmarshal JSON data into TelemetryDruidColumnComparisonFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidColumnComparisonFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidColumnComparisonFilter, return on the first match
		} else {
			dst.TelemetryDruidColumnComparisonFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidColumnComparisonFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidEqualityFilter'
	if jsonDict["type"] == "telemetry.DruidEqualityFilter" {
		// try to unmarshal JSON data into TelemetryDruidEqualityFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidEqualityFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidEqualityFilter, return on the first match
		} else {
			dst.TelemetryDruidEqualityFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidEqualityFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidExpressionFilter'
	if jsonDict["type"] == "telemetry.DruidExpressionFilter" {
		// try to unmarshal JSON data into TelemetryDruidExpressionFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidExpressionFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidExpressionFilter, return on the first match
		} else {
			dst.TelemetryDruidExpressionFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidExpressionFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidFalseFilter'
	if jsonDict["type"] == "telemetry.DruidFalseFilter" {
		// try to unmarshal JSON data into TelemetryDruidFalseFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidFalseFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidFalseFilter, return on the first match
		} else {
			dst.TelemetryDruidFalseFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidFalseFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidInFilter'
	if jsonDict["type"] == "telemetry.DruidInFilter" {
		// try to unmarshal JSON data into TelemetryDruidInFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidInFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInFilter, return on the first match
		} else {
			dst.TelemetryDruidInFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidInFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidIntervalFilter'
	if jsonDict["type"] == "telemetry.DruidIntervalFilter" {
		// try to unmarshal JSON data into TelemetryDruidIntervalFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidIntervalFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidIntervalFilter, return on the first match
		} else {
			dst.TelemetryDruidIntervalFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidIntervalFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidLikeFilter'
	if jsonDict["type"] == "telemetry.DruidLikeFilter" {
		// try to unmarshal JSON data into TelemetryDruidLikeFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidLikeFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidLikeFilter, return on the first match
		} else {
			dst.TelemetryDruidLikeFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidLikeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidNotFilter'
	if jsonDict["type"] == "telemetry.DruidNotFilter" {
		// try to unmarshal JSON data into TelemetryDruidNotFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNotFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNotFilter, return on the first match
		} else {
			dst.TelemetryDruidNotFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidNotFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidNullFilter'
	if jsonDict["type"] == "telemetry.DruidNullFilter" {
		// try to unmarshal JSON data into TelemetryDruidNullFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidNullFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidNullFilter, return on the first match
		} else {
			dst.TelemetryDruidNullFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidNullFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidOrFilter'
	if jsonDict["type"] == "telemetry.DruidOrFilter" {
		// try to unmarshal JSON data into TelemetryDruidOrFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidOrFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidOrFilter, return on the first match
		} else {
			dst.TelemetryDruidOrFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidOrFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidRangeFilter'
	if jsonDict["type"] == "telemetry.DruidRangeFilter" {
		// try to unmarshal JSON data into TelemetryDruidRangeFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRangeFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRangeFilter, return on the first match
		} else {
			dst.TelemetryDruidRangeFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidRangeFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidRegexFilter'
	if jsonDict["type"] == "telemetry.DruidRegexFilter" {
		// try to unmarshal JSON data into TelemetryDruidRegexFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidRegexFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidRegexFilter, return on the first match
		} else {
			dst.TelemetryDruidRegexFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidRegexFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidSearchFilter'
	if jsonDict["type"] == "telemetry.DruidSearchFilter" {
		// try to unmarshal JSON data into TelemetryDruidSearchFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSearchFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSearchFilter, return on the first match
		} else {
			dst.TelemetryDruidSearchFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidSearchFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidSelectorFilter'
	if jsonDict["type"] == "telemetry.DruidSelectorFilter" {
		// try to unmarshal JSON data into TelemetryDruidSelectorFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidSelectorFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidSelectorFilter, return on the first match
		} else {
			dst.TelemetryDruidSelectorFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidSelectorFilter: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidTrueFilter'
	if jsonDict["type"] == "telemetry.DruidTrueFilter" {
		// try to unmarshal JSON data into TelemetryDruidTrueFilter
		err = json.Unmarshal(data, &dst.TelemetryDruidTrueFilter)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidTrueFilter, return on the first match
		} else {
			dst.TelemetryDruidTrueFilter = nil
			return fmt.Errorf("failed to unmarshal TelemetryDruidFilter as TelemetryDruidTrueFilter: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidFilter) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidAndFilter != nil {
		return json.Marshal(&src.TelemetryDruidAndFilter)
	}

	if src.TelemetryDruidBoundFilter != nil {
		return json.Marshal(&src.TelemetryDruidBoundFilter)
	}

	if src.TelemetryDruidColumnComparisonFilter != nil {
		return json.Marshal(&src.TelemetryDruidColumnComparisonFilter)
	}

	if src.TelemetryDruidEqualityFilter != nil {
		return json.Marshal(&src.TelemetryDruidEqualityFilter)
	}

	if src.TelemetryDruidExpressionFilter != nil {
		return json.Marshal(&src.TelemetryDruidExpressionFilter)
	}

	if src.TelemetryDruidFalseFilter != nil {
		return json.Marshal(&src.TelemetryDruidFalseFilter)
	}

	if src.TelemetryDruidInFilter != nil {
		return json.Marshal(&src.TelemetryDruidInFilter)
	}

	if src.TelemetryDruidIntervalFilter != nil {
		return json.Marshal(&src.TelemetryDruidIntervalFilter)
	}

	if src.TelemetryDruidLikeFilter != nil {
		return json.Marshal(&src.TelemetryDruidLikeFilter)
	}

	if src.TelemetryDruidNotFilter != nil {
		return json.Marshal(&src.TelemetryDruidNotFilter)
	}

	if src.TelemetryDruidNullFilter != nil {
		return json.Marshal(&src.TelemetryDruidNullFilter)
	}

	if src.TelemetryDruidOrFilter != nil {
		return json.Marshal(&src.TelemetryDruidOrFilter)
	}

	if src.TelemetryDruidRangeFilter != nil {
		return json.Marshal(&src.TelemetryDruidRangeFilter)
	}

	if src.TelemetryDruidRegexFilter != nil {
		return json.Marshal(&src.TelemetryDruidRegexFilter)
	}

	if src.TelemetryDruidSearchFilter != nil {
		return json.Marshal(&src.TelemetryDruidSearchFilter)
	}

	if src.TelemetryDruidSelectorFilter != nil {
		return json.Marshal(&src.TelemetryDruidSelectorFilter)
	}

	if src.TelemetryDruidTrueFilter != nil {
		return json.Marshal(&src.TelemetryDruidTrueFilter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidFilter) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TelemetryDruidAndFilter != nil {
		return obj.TelemetryDruidAndFilter
	}

	if obj.TelemetryDruidBoundFilter != nil {
		return obj.TelemetryDruidBoundFilter
	}

	if obj.TelemetryDruidColumnComparisonFilter != nil {
		return obj.TelemetryDruidColumnComparisonFilter
	}

	if obj.TelemetryDruidEqualityFilter != nil {
		return obj.TelemetryDruidEqualityFilter
	}

	if obj.TelemetryDruidExpressionFilter != nil {
		return obj.TelemetryDruidExpressionFilter
	}

	if obj.TelemetryDruidFalseFilter != nil {
		return obj.TelemetryDruidFalseFilter
	}

	if obj.TelemetryDruidInFilter != nil {
		return obj.TelemetryDruidInFilter
	}

	if obj.TelemetryDruidIntervalFilter != nil {
		return obj.TelemetryDruidIntervalFilter
	}

	if obj.TelemetryDruidLikeFilter != nil {
		return obj.TelemetryDruidLikeFilter
	}

	if obj.TelemetryDruidNotFilter != nil {
		return obj.TelemetryDruidNotFilter
	}

	if obj.TelemetryDruidNullFilter != nil {
		return obj.TelemetryDruidNullFilter
	}

	if obj.TelemetryDruidOrFilter != nil {
		return obj.TelemetryDruidOrFilter
	}

	if obj.TelemetryDruidRangeFilter != nil {
		return obj.TelemetryDruidRangeFilter
	}

	if obj.TelemetryDruidRegexFilter != nil {
		return obj.TelemetryDruidRegexFilter
	}

	if obj.TelemetryDruidSearchFilter != nil {
		return obj.TelemetryDruidSearchFilter
	}

	if obj.TelemetryDruidSelectorFilter != nil {
		return obj.TelemetryDruidSelectorFilter
	}

	if obj.TelemetryDruidTrueFilter != nil {
		return obj.TelemetryDruidTrueFilter
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidFilter struct {
	value *TelemetryDruidFilter
	isSet bool
}

func (v NullableTelemetryDruidFilter) Get() *TelemetryDruidFilter {
	return v.value
}

func (v *NullableTelemetryDruidFilter) Set(val *TelemetryDruidFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidFilter(val *TelemetryDruidFilter) *NullableTelemetryDruidFilter {
	return &NullableTelemetryDruidFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
