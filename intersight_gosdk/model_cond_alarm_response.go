/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// CondAlarmResponse - The response body of a HTTP GET request for the 'cond.Alarm' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'cond.Alarm' resources.
type CondAlarmResponse struct {
	CondAlarmList        *CondAlarmList
	MoAggregateTransform *MoAggregateTransform
	MoDocumentCount      *MoDocumentCount
	MoTagSummary         *MoTagSummary
}

// CondAlarmListAsCondAlarmResponse is a convenience function that returns CondAlarmList wrapped in CondAlarmResponse
func CondAlarmListAsCondAlarmResponse(v *CondAlarmList) CondAlarmResponse {
	return CondAlarmResponse{
		CondAlarmList: v,
	}
}

// MoAggregateTransformAsCondAlarmResponse is a convenience function that returns MoAggregateTransform wrapped in CondAlarmResponse
func MoAggregateTransformAsCondAlarmResponse(v *MoAggregateTransform) CondAlarmResponse {
	return CondAlarmResponse{
		MoAggregateTransform: v,
	}
}

// MoDocumentCountAsCondAlarmResponse is a convenience function that returns MoDocumentCount wrapped in CondAlarmResponse
func MoDocumentCountAsCondAlarmResponse(v *MoDocumentCount) CondAlarmResponse {
	return CondAlarmResponse{
		MoDocumentCount: v,
	}
}

// MoTagSummaryAsCondAlarmResponse is a convenience function that returns MoTagSummary wrapped in CondAlarmResponse
func MoTagSummaryAsCondAlarmResponse(v *MoTagSummary) CondAlarmResponse {
	return CondAlarmResponse{
		MoTagSummary: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CondAlarmResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'cond.Alarm.List'
	if jsonDict["ObjectType"] == "cond.Alarm.List" {
		// try to unmarshal JSON data into CondAlarmList
		err = json.Unmarshal(data, &dst.CondAlarmList)
		if err == nil {
			return nil // data stored in dst.CondAlarmList, return on the first match
		} else {
			dst.CondAlarmList = nil
			return fmt.Errorf("Failed to unmarshal CondAlarmResponse as CondAlarmList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal CondAlarmResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal CondAlarmResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal CondAlarmResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CondAlarmResponse) MarshalJSON() ([]byte, error) {
	if src.CondAlarmList != nil {
		return json.Marshal(&src.CondAlarmList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CondAlarmResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CondAlarmList != nil {
		return obj.CondAlarmList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableCondAlarmResponse struct {
	value *CondAlarmResponse
	isSet bool
}

func (v NullableCondAlarmResponse) Get() *CondAlarmResponse {
	return v.value
}

func (v *NullableCondAlarmResponse) Set(val *CondAlarmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmResponse(val *CondAlarmResponse) *NullableCondAlarmResponse {
	return &NullableCondAlarmResponse{value: val, isSet: true}
}

func (v NullableCondAlarmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
