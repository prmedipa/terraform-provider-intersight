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

// checks if the CapabilityFanModuleDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityFanModuleDescriptor{}

// CapabilityFanModuleDescriptor Descriptor that uniquely identifies a fan module.
type CapabilityFanModuleDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Revision for the chassis enclosure.
	Revision             *string `json:"Revision,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityFanModuleDescriptor CapabilityFanModuleDescriptor

// NewCapabilityFanModuleDescriptor instantiates a new CapabilityFanModuleDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityFanModuleDescriptor(classId string, objectType string) *CapabilityFanModuleDescriptor {
	this := CapabilityFanModuleDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityFanModuleDescriptorWithDefaults instantiates a new CapabilityFanModuleDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityFanModuleDescriptorWithDefaults() *CapabilityFanModuleDescriptor {
	this := CapabilityFanModuleDescriptor{}
	var classId string = "capability.FanModuleDescriptor"
	this.ClassId = classId
	var objectType string = "capability.FanModuleDescriptor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityFanModuleDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityFanModuleDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.FanModuleDescriptor" of the ClassId field.
func (o *CapabilityFanModuleDescriptor) GetDefaultClassId() interface{} {
	return "capability.FanModuleDescriptor"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityFanModuleDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityFanModuleDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.FanModuleDescriptor" of the ObjectType field.
func (o *CapabilityFanModuleDescriptor) GetDefaultObjectType() interface{} {
	return "capability.FanModuleDescriptor"
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *CapabilityFanModuleDescriptor) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityFanModuleDescriptor) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *CapabilityFanModuleDescriptor) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *CapabilityFanModuleDescriptor) SetRevision(v string) {
	o.Revision = &v
}

func (o CapabilityFanModuleDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityFanModuleDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Revision) {
		toSerialize["Revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityFanModuleDescriptor) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityFanModuleDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Revision for the chassis enclosure.
		Revision *string `json:"Revision,omitempty"`
	}

	varCapabilityFanModuleDescriptorWithoutEmbeddedStruct := CapabilityFanModuleDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityFanModuleDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityFanModuleDescriptor := _CapabilityFanModuleDescriptor{}
		varCapabilityFanModuleDescriptor.ClassId = varCapabilityFanModuleDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityFanModuleDescriptor.ObjectType = varCapabilityFanModuleDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityFanModuleDescriptor.Revision = varCapabilityFanModuleDescriptorWithoutEmbeddedStruct.Revision
		*o = CapabilityFanModuleDescriptor(varCapabilityFanModuleDescriptor)
	} else {
		return err
	}

	varCapabilityFanModuleDescriptor := _CapabilityFanModuleDescriptor{}

	err = json.Unmarshal(data, &varCapabilityFanModuleDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilityFanModuleDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Revision")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilityFanModuleDescriptor struct {
	value *CapabilityFanModuleDescriptor
	isSet bool
}

func (v NullableCapabilityFanModuleDescriptor) Get() *CapabilityFanModuleDescriptor {
	return v.value
}

func (v *NullableCapabilityFanModuleDescriptor) Set(val *CapabilityFanModuleDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityFanModuleDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityFanModuleDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityFanModuleDescriptor(val *CapabilityFanModuleDescriptor) *NullableCapabilityFanModuleDescriptor {
	return &NullableCapabilityFanModuleDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityFanModuleDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityFanModuleDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
