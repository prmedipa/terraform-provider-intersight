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

// checks if the NiaapiVersionRegex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiaapiVersionRegex{}

// NiaapiVersionRegex The regular expression pattern to recongnize the version string.
type NiaapiVersionRegex struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                             `json:"ObjectType"`
	Apic       NullableNiaapiVersionRegexPlatform `json:"Apic,omitempty"`
	Dcnm       NullableNiaapiVersionRegexPlatform `json:"Dcnm,omitempty"`
	// Version number for the Version Regex data, also used as identity.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiVersionRegex NiaapiVersionRegex

// NewNiaapiVersionRegex instantiates a new NiaapiVersionRegex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiVersionRegex(classId string, objectType string) *NiaapiVersionRegex {
	this := NiaapiVersionRegex{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiVersionRegexWithDefaults instantiates a new NiaapiVersionRegex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiVersionRegexWithDefaults() *NiaapiVersionRegex {
	this := NiaapiVersionRegex{}
	var classId string = "niaapi.VersionRegex"
	this.ClassId = classId
	var objectType string = "niaapi.VersionRegex"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiVersionRegex) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegex) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiVersionRegex) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niaapi.VersionRegex" of the ClassId field.
func (o *NiaapiVersionRegex) GetDefaultClassId() interface{} {
	return "niaapi.VersionRegex"
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiVersionRegex) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegex) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiVersionRegex) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niaapi.VersionRegex" of the ObjectType field.
func (o *NiaapiVersionRegex) GetDefaultObjectType() interface{} {
	return "niaapi.VersionRegex"
}

// GetApic returns the Apic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegex) GetApic() NiaapiVersionRegexPlatform {
	if o == nil || IsNil(o.Apic.Get()) {
		var ret NiaapiVersionRegexPlatform
		return ret
	}
	return *o.Apic.Get()
}

// GetApicOk returns a tuple with the Apic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegex) GetApicOk() (*NiaapiVersionRegexPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Apic.Get(), o.Apic.IsSet()
}

// HasApic returns a boolean if a field has been set.
func (o *NiaapiVersionRegex) HasApic() bool {
	if o != nil && o.Apic.IsSet() {
		return true
	}

	return false
}

// SetApic gets a reference to the given NullableNiaapiVersionRegexPlatform and assigns it to the Apic field.
func (o *NiaapiVersionRegex) SetApic(v NiaapiVersionRegexPlatform) {
	o.Apic.Set(&v)
}

// SetApicNil sets the value for Apic to be an explicit nil
func (o *NiaapiVersionRegex) SetApicNil() {
	o.Apic.Set(nil)
}

// UnsetApic ensures that no value is present for Apic, not even an explicit nil
func (o *NiaapiVersionRegex) UnsetApic() {
	o.Apic.Unset()
}

// GetDcnm returns the Dcnm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiaapiVersionRegex) GetDcnm() NiaapiVersionRegexPlatform {
	if o == nil || IsNil(o.Dcnm.Get()) {
		var ret NiaapiVersionRegexPlatform
		return ret
	}
	return *o.Dcnm.Get()
}

// GetDcnmOk returns a tuple with the Dcnm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiaapiVersionRegex) GetDcnmOk() (*NiaapiVersionRegexPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dcnm.Get(), o.Dcnm.IsSet()
}

// HasDcnm returns a boolean if a field has been set.
func (o *NiaapiVersionRegex) HasDcnm() bool {
	if o != nil && o.Dcnm.IsSet() {
		return true
	}

	return false
}

// SetDcnm gets a reference to the given NullableNiaapiVersionRegexPlatform and assigns it to the Dcnm field.
func (o *NiaapiVersionRegex) SetDcnm(v NiaapiVersionRegexPlatform) {
	o.Dcnm.Set(&v)
}

// SetDcnmNil sets the value for Dcnm to be an explicit nil
func (o *NiaapiVersionRegex) SetDcnmNil() {
	o.Dcnm.Set(nil)
}

// UnsetDcnm ensures that no value is present for Dcnm, not even an explicit nil
func (o *NiaapiVersionRegex) UnsetDcnm() {
	o.Dcnm.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiVersionRegex) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiVersionRegex) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiVersionRegex) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiaapiVersionRegex) SetVersion(v string) {
	o.Version = &v
}

func (o NiaapiVersionRegex) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiaapiVersionRegex) ToMap() (map[string]interface{}, error) {
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
	if o.Apic.IsSet() {
		toSerialize["Apic"] = o.Apic.Get()
	}
	if o.Dcnm.IsSet() {
		toSerialize["Dcnm"] = o.Dcnm.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiaapiVersionRegex) UnmarshalJSON(data []byte) (err error) {
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
	type NiaapiVersionRegexWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                             `json:"ObjectType"`
		Apic       NullableNiaapiVersionRegexPlatform `json:"Apic,omitempty"`
		Dcnm       NullableNiaapiVersionRegexPlatform `json:"Dcnm,omitempty"`
		// Version number for the Version Regex data, also used as identity.
		Version *string `json:"Version,omitempty"`
	}

	varNiaapiVersionRegexWithoutEmbeddedStruct := NiaapiVersionRegexWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiaapiVersionRegexWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiVersionRegex := _NiaapiVersionRegex{}
		varNiaapiVersionRegex.ClassId = varNiaapiVersionRegexWithoutEmbeddedStruct.ClassId
		varNiaapiVersionRegex.ObjectType = varNiaapiVersionRegexWithoutEmbeddedStruct.ObjectType
		varNiaapiVersionRegex.Apic = varNiaapiVersionRegexWithoutEmbeddedStruct.Apic
		varNiaapiVersionRegex.Dcnm = varNiaapiVersionRegexWithoutEmbeddedStruct.Dcnm
		varNiaapiVersionRegex.Version = varNiaapiVersionRegexWithoutEmbeddedStruct.Version
		*o = NiaapiVersionRegex(varNiaapiVersionRegex)
	} else {
		return err
	}

	varNiaapiVersionRegex := _NiaapiVersionRegex{}

	err = json.Unmarshal(data, &varNiaapiVersionRegex)
	if err == nil {
		o.MoBaseMo = varNiaapiVersionRegex.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Apic")
		delete(additionalProperties, "Dcnm")
		delete(additionalProperties, "Version")

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

type NullableNiaapiVersionRegex struct {
	value *NiaapiVersionRegex
	isSet bool
}

func (v NullableNiaapiVersionRegex) Get() *NiaapiVersionRegex {
	return v.value
}

func (v *NullableNiaapiVersionRegex) Set(val *NiaapiVersionRegex) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiVersionRegex) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiVersionRegex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiVersionRegex(val *NiaapiVersionRegex) *NullableNiaapiVersionRegex {
	return &NullableNiaapiVersionRegex{value: val, isSet: true}
}

func (v NullableNiaapiVersionRegex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiVersionRegex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
