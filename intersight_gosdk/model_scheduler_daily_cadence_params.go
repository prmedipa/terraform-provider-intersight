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
	"reflect"
	"strings"
)

// checks if the SchedulerDailyCadenceParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulerDailyCadenceParams{}

// SchedulerDailyCadenceParams Parameters used for running daily (almost).
type SchedulerDailyCadenceParams struct {
	SchedulerBaseCadenceParams
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Run every day by default if not specified.
	RunEvery             *int64 `json:"RunEvery,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulerDailyCadenceParams SchedulerDailyCadenceParams

// NewSchedulerDailyCadenceParams instantiates a new SchedulerDailyCadenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulerDailyCadenceParams(classId string, objectType string) *SchedulerDailyCadenceParams {
	this := SchedulerDailyCadenceParams{}
	this.ClassId = classId
	this.ObjectType = objectType
	var runEvery int64 = 1
	this.RunEvery = &runEvery
	return &this
}

// NewSchedulerDailyCadenceParamsWithDefaults instantiates a new SchedulerDailyCadenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulerDailyCadenceParamsWithDefaults() *SchedulerDailyCadenceParams {
	this := SchedulerDailyCadenceParams{}
	var classId string = "scheduler.DailyCadenceParams"
	this.ClassId = classId
	var objectType string = "scheduler.DailyCadenceParams"
	this.ObjectType = objectType
	var runEvery int64 = 1
	this.RunEvery = &runEvery
	return &this
}

// GetClassId returns the ClassId field value
func (o *SchedulerDailyCadenceParams) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SchedulerDailyCadenceParams) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SchedulerDailyCadenceParams) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "scheduler.DailyCadenceParams" of the ClassId field.
func (o *SchedulerDailyCadenceParams) GetDefaultClassId() interface{} {
	return "scheduler.DailyCadenceParams"
}

// GetObjectType returns the ObjectType field value
func (o *SchedulerDailyCadenceParams) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SchedulerDailyCadenceParams) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SchedulerDailyCadenceParams) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "scheduler.DailyCadenceParams" of the ObjectType field.
func (o *SchedulerDailyCadenceParams) GetDefaultObjectType() interface{} {
	return "scheduler.DailyCadenceParams"
}

// GetRunEvery returns the RunEvery field value if set, zero value otherwise.
func (o *SchedulerDailyCadenceParams) GetRunEvery() int64 {
	if o == nil || IsNil(o.RunEvery) {
		var ret int64
		return ret
	}
	return *o.RunEvery
}

// GetRunEveryOk returns a tuple with the RunEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulerDailyCadenceParams) GetRunEveryOk() (*int64, bool) {
	if o == nil || IsNil(o.RunEvery) {
		return nil, false
	}
	return o.RunEvery, true
}

// HasRunEvery returns a boolean if a field has been set.
func (o *SchedulerDailyCadenceParams) HasRunEvery() bool {
	if o != nil && !IsNil(o.RunEvery) {
		return true
	}

	return false
}

// SetRunEvery gets a reference to the given int64 and assigns it to the RunEvery field.
func (o *SchedulerDailyCadenceParams) SetRunEvery(v int64) {
	o.RunEvery = &v
}

func (o SchedulerDailyCadenceParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulerDailyCadenceParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSchedulerBaseCadenceParams, errSchedulerBaseCadenceParams := json.Marshal(o.SchedulerBaseCadenceParams)
	if errSchedulerBaseCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseCadenceParams
	}
	errSchedulerBaseCadenceParams = json.Unmarshal([]byte(serializedSchedulerBaseCadenceParams), &toSerialize)
	if errSchedulerBaseCadenceParams != nil {
		return map[string]interface{}{}, errSchedulerBaseCadenceParams
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.RunEvery) {
		toSerialize["RunEvery"] = o.RunEvery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulerDailyCadenceParams) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type SchedulerDailyCadenceParamsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Run every day by default if not specified.
		RunEvery *int64 `json:"RunEvery,omitempty"`
	}

	varSchedulerDailyCadenceParamsWithoutEmbeddedStruct := SchedulerDailyCadenceParamsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSchedulerDailyCadenceParamsWithoutEmbeddedStruct)
	if err == nil {
		varSchedulerDailyCadenceParams := _SchedulerDailyCadenceParams{}
		varSchedulerDailyCadenceParams.ClassId = varSchedulerDailyCadenceParamsWithoutEmbeddedStruct.ClassId
		varSchedulerDailyCadenceParams.ObjectType = varSchedulerDailyCadenceParamsWithoutEmbeddedStruct.ObjectType
		varSchedulerDailyCadenceParams.RunEvery = varSchedulerDailyCadenceParamsWithoutEmbeddedStruct.RunEvery
		*o = SchedulerDailyCadenceParams(varSchedulerDailyCadenceParams)
	} else {
		return err
	}

	varSchedulerDailyCadenceParams := _SchedulerDailyCadenceParams{}

	err = json.Unmarshal(data, &varSchedulerDailyCadenceParams)
	if err == nil {
		o.SchedulerBaseCadenceParams = varSchedulerDailyCadenceParams.SchedulerBaseCadenceParams
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RunEvery")

		// remove fields from embedded structs
		reflectSchedulerBaseCadenceParams := reflect.ValueOf(o.SchedulerBaseCadenceParams)
		for i := 0; i < reflectSchedulerBaseCadenceParams.Type().NumField(); i++ {
			t := reflectSchedulerBaseCadenceParams.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedulerDailyCadenceParams struct {
	value *SchedulerDailyCadenceParams
	isSet bool
}

func (v NullableSchedulerDailyCadenceParams) Get() *SchedulerDailyCadenceParams {
	return v.value
}

func (v *NullableSchedulerDailyCadenceParams) Set(val *SchedulerDailyCadenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulerDailyCadenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulerDailyCadenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulerDailyCadenceParams(val *SchedulerDailyCadenceParams) *NullableSchedulerDailyCadenceParams {
	return &NullableSchedulerDailyCadenceParams{value: val, isSet: true}
}

func (v NullableSchedulerDailyCadenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulerDailyCadenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
