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

// checks if the HclOperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HclOperatingSystem{}

// HclOperatingSystem Collection used to store operating system details.
type HclOperatingSystem struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version of the Operating System.
	Version              *string                                      `json:"Version,omitempty"`
	Vendor               NullableHclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclOperatingSystem HclOperatingSystem

// NewHclOperatingSystem instantiates a new HclOperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclOperatingSystem(classId string, objectType string) *HclOperatingSystem {
	this := HclOperatingSystem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHclOperatingSystemWithDefaults instantiates a new HclOperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclOperatingSystemWithDefaults() *HclOperatingSystem {
	this := HclOperatingSystem{}
	var classId string = "hcl.OperatingSystem"
	this.ClassId = classId
	var objectType string = "hcl.OperatingSystem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HclOperatingSystem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HclOperatingSystem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HclOperatingSystem) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hcl.OperatingSystem" of the ClassId field.
func (o *HclOperatingSystem) GetDefaultClassId() interface{} {
	return "hcl.OperatingSystem"
}

// GetObjectType returns the ObjectType field value
func (o *HclOperatingSystem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HclOperatingSystem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HclOperatingSystem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hcl.OperatingSystem" of the ObjectType field.
func (o *HclOperatingSystem) GetDefaultObjectType() interface{} {
	return "hcl.OperatingSystem"
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HclOperatingSystem) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclOperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HclOperatingSystem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HclOperatingSystem) SetVersion(v string) {
	o.Version = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HclOperatingSystem) GetVendor() HclOperatingSystemVendorRelationship {
	if o == nil || IsNil(o.Vendor.Get()) {
		var ret HclOperatingSystemVendorRelationship
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HclOperatingSystem) GetVendorOk() (*HclOperatingSystemVendorRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *HclOperatingSystem) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableHclOperatingSystemVendorRelationship and assigns it to the Vendor field.
func (o *HclOperatingSystem) SetVendor(v HclOperatingSystemVendorRelationship) {
	o.Vendor.Set(&v)
}

// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *HclOperatingSystem) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *HclOperatingSystem) UnsetVendor() {
	o.Vendor.Unset()
}

func (o HclOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HclOperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Vendor.IsSet() {
		toSerialize["Vendor"] = o.Vendor.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HclOperatingSystem) UnmarshalJSON(data []byte) (err error) {
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
	type HclOperatingSystemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Version of the Operating System.
		Version *string                                      `json:"Version,omitempty"`
		Vendor  NullableHclOperatingSystemVendorRelationship `json:"Vendor,omitempty"`
	}

	varHclOperatingSystemWithoutEmbeddedStruct := HclOperatingSystemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHclOperatingSystemWithoutEmbeddedStruct)
	if err == nil {
		varHclOperatingSystem := _HclOperatingSystem{}
		varHclOperatingSystem.ClassId = varHclOperatingSystemWithoutEmbeddedStruct.ClassId
		varHclOperatingSystem.ObjectType = varHclOperatingSystemWithoutEmbeddedStruct.ObjectType
		varHclOperatingSystem.Version = varHclOperatingSystemWithoutEmbeddedStruct.Version
		varHclOperatingSystem.Vendor = varHclOperatingSystemWithoutEmbeddedStruct.Vendor
		*o = HclOperatingSystem(varHclOperatingSystem)
	} else {
		return err
	}

	varHclOperatingSystem := _HclOperatingSystem{}

	err = json.Unmarshal(data, &varHclOperatingSystem)
	if err == nil {
		o.MoBaseMo = varHclOperatingSystem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Vendor")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHclOperatingSystem struct {
	value *HclOperatingSystem
	isSet bool
}

func (v NullableHclOperatingSystem) Get() *HclOperatingSystem {
	return v.value
}

func (v *NullableHclOperatingSystem) Set(val *HclOperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableHclOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableHclOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclOperatingSystem(val *HclOperatingSystem) *NullableHclOperatingSystem {
	return &NullableHclOperatingSystem{value: val, isSet: true}
}

func (v NullableHclOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
