/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the HyperflexHxZoneResiliencyInfoDt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexHxZoneResiliencyInfoDt{}

// HyperflexHxZoneResiliencyInfoDt A detailed status of the resiliency of a site or zone.
type HyperflexHxZoneResiliencyInfoDt struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the availability zone.
	Name                 *string                             `json:"Name,omitempty"`
	ResiliencyInfo       NullableHyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxZoneResiliencyInfoDt HyperflexHxZoneResiliencyInfoDt

// NewHyperflexHxZoneResiliencyInfoDt instantiates a new HyperflexHxZoneResiliencyInfoDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxZoneResiliencyInfoDt(classId string, objectType string) *HyperflexHxZoneResiliencyInfoDt {
	this := HyperflexHxZoneResiliencyInfoDt{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxZoneResiliencyInfoDtWithDefaults instantiates a new HyperflexHxZoneResiliencyInfoDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxZoneResiliencyInfoDtWithDefaults() *HyperflexHxZoneResiliencyInfoDt {
	this := HyperflexHxZoneResiliencyInfoDt{}
	var classId string = "hyperflex.HxZoneResiliencyInfoDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxZoneResiliencyInfoDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxZoneResiliencyInfoDt) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxZoneResiliencyInfoDt) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.HxZoneResiliencyInfoDt" of the ClassId field.
func (o *HyperflexHxZoneResiliencyInfoDt) GetDefaultClassId() interface{} {
	return "hyperflex.HxZoneResiliencyInfoDt"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxZoneResiliencyInfoDt) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxZoneResiliencyInfoDt) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.HxZoneResiliencyInfoDt" of the ObjectType field.
func (o *HyperflexHxZoneResiliencyInfoDt) GetDefaultObjectType() interface{} {
	return "hyperflex.HxZoneResiliencyInfoDt"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexHxZoneResiliencyInfoDt) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexHxZoneResiliencyInfoDt) SetName(v string) {
	o.Name = &v
}

// GetResiliencyInfo returns the ResiliencyInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexHxZoneResiliencyInfoDt) GetResiliencyInfo() HyperflexHxResiliencyInfoDt {
	if o == nil || IsNil(o.ResiliencyInfo.Get()) {
		var ret HyperflexHxResiliencyInfoDt
		return ret
	}
	return *o.ResiliencyInfo.Get()
}

// GetResiliencyInfoOk returns a tuple with the ResiliencyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexHxZoneResiliencyInfoDt) GetResiliencyInfoOk() (*HyperflexHxResiliencyInfoDt, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResiliencyInfo.Get(), o.ResiliencyInfo.IsSet()
}

// HasResiliencyInfo returns a boolean if a field has been set.
func (o *HyperflexHxZoneResiliencyInfoDt) HasResiliencyInfo() bool {
	if o != nil && o.ResiliencyInfo.IsSet() {
		return true
	}

	return false
}

// SetResiliencyInfo gets a reference to the given NullableHyperflexHxResiliencyInfoDt and assigns it to the ResiliencyInfo field.
func (o *HyperflexHxZoneResiliencyInfoDt) SetResiliencyInfo(v HyperflexHxResiliencyInfoDt) {
	o.ResiliencyInfo.Set(&v)
}

// SetResiliencyInfoNil sets the value for ResiliencyInfo to be an explicit nil
func (o *HyperflexHxZoneResiliencyInfoDt) SetResiliencyInfoNil() {
	o.ResiliencyInfo.Set(nil)
}

// UnsetResiliencyInfo ensures that no value is present for ResiliencyInfo, not even an explicit nil
func (o *HyperflexHxZoneResiliencyInfoDt) UnsetResiliencyInfo() {
	o.ResiliencyInfo.Unset()
}

func (o HyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexHxZoneResiliencyInfoDt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.ResiliencyInfo.IsSet() {
		toSerialize["ResiliencyInfo"] = o.ResiliencyInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the availability zone.
		Name           *string                             `json:"Name,omitempty"`
		ResiliencyInfo NullableHyperflexHxResiliencyInfoDt `json:"ResiliencyInfo,omitempty"`
	}

	varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct := HyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexHxZoneResiliencyInfoDt := _HyperflexHxZoneResiliencyInfoDt{}
		varHyperflexHxZoneResiliencyInfoDt.ClassId = varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct.ClassId
		varHyperflexHxZoneResiliencyInfoDt.ObjectType = varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct.ObjectType
		varHyperflexHxZoneResiliencyInfoDt.Name = varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct.Name
		varHyperflexHxZoneResiliencyInfoDt.ResiliencyInfo = varHyperflexHxZoneResiliencyInfoDtWithoutEmbeddedStruct.ResiliencyInfo
		*o = HyperflexHxZoneResiliencyInfoDt(varHyperflexHxZoneResiliencyInfoDt)
	} else {
		return err
	}

	varHyperflexHxZoneResiliencyInfoDt := _HyperflexHxZoneResiliencyInfoDt{}

	err = json.Unmarshal(data, &varHyperflexHxZoneResiliencyInfoDt)
	if err == nil {
		o.MoBaseComplexType = varHyperflexHxZoneResiliencyInfoDt.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ResiliencyInfo")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHyperflexHxZoneResiliencyInfoDt struct {
	value *HyperflexHxZoneResiliencyInfoDt
	isSet bool
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) Get() *HyperflexHxZoneResiliencyInfoDt {
	return v.value
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) Set(val *HyperflexHxZoneResiliencyInfoDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxZoneResiliencyInfoDt(val *HyperflexHxZoneResiliencyInfoDt) *NullableHyperflexHxZoneResiliencyInfoDt {
	return &NullableHyperflexHxZoneResiliencyInfoDt{value: val, isSet: true}
}

func (v NullableHyperflexHxZoneResiliencyInfoDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxZoneResiliencyInfoDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
