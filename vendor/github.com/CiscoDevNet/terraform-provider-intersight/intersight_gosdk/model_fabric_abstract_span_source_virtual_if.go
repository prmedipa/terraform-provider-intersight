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

// checks if the FabricAbstractSpanSourceVirtualIf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanSourceVirtualIf{}

// FabricAbstractSpanSourceVirtualIf Abstract class for all SPAN Source Virtual Interfaces including Vnics and Vhbas.
type FabricAbstractSpanSourceVirtualIf struct {
	FabricAbstractSpanSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the VNIC referenced by vnic relationship. Vnic name is not updated if it gets updated by the user.
	Name *string `json:"Name,omitempty"`
	// VIF ID of the VNIC placed on Fabric Interconnect associated to the SPAN session.
	VifId                *int64                                `json:"VifId,omitempty"`
	SpanSession          NullableFabricSpanSessionRelationship `json:"SpanSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanSourceVirtualIf FabricAbstractSpanSourceVirtualIf

// NewFabricAbstractSpanSourceVirtualIf instantiates a new FabricAbstractSpanSourceVirtualIf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanSourceVirtualIf(classId string, objectType string) *FabricAbstractSpanSourceVirtualIf {
	this := FabricAbstractSpanSourceVirtualIf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricAbstractSpanSourceVirtualIfWithDefaults instantiates a new FabricAbstractSpanSourceVirtualIf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanSourceVirtualIfWithDefaults() *FabricAbstractSpanSourceVirtualIf {
	this := FabricAbstractSpanSourceVirtualIf{}
	var classId string = "fabric.SpanSourceVnicEthIf"
	this.ClassId = classId
	var objectType string = "fabric.SpanSourceVnicEthIf"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanSourceVirtualIf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourceVirtualIf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanSourceVirtualIf) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanSourceVnicEthIf" of the ClassId field.
func (o *FabricAbstractSpanSourceVirtualIf) GetDefaultClassId() interface{} {
	return "fabric.SpanSourceVnicEthIf"
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanSourceVirtualIf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourceVirtualIf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanSourceVirtualIf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanSourceVnicEthIf" of the ObjectType field.
func (o *FabricAbstractSpanSourceVirtualIf) GetDefaultObjectType() interface{} {
	return "fabric.SpanSourceVnicEthIf"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourceVirtualIf) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourceVirtualIf) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourceVirtualIf) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricAbstractSpanSourceVirtualIf) SetName(v string) {
	o.Name = &v
}

// GetVifId returns the VifId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourceVirtualIf) GetVifId() int64 {
	if o == nil || IsNil(o.VifId) {
		var ret int64
		return ret
	}
	return *o.VifId
}

// GetVifIdOk returns a tuple with the VifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourceVirtualIf) GetVifIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VifId) {
		return nil, false
	}
	return o.VifId, true
}

// HasVifId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourceVirtualIf) HasVifId() bool {
	if o != nil && !IsNil(o.VifId) {
		return true
	}

	return false
}

// SetVifId gets a reference to the given int64 and assigns it to the VifId field.
func (o *FabricAbstractSpanSourceVirtualIf) SetVifId(v int64) {
	o.VifId = &v
}

// GetSpanSession returns the SpanSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSourceVirtualIf) GetSpanSession() FabricSpanSessionRelationship {
	if o == nil || IsNil(o.SpanSession.Get()) {
		var ret FabricSpanSessionRelationship
		return ret
	}
	return *o.SpanSession.Get()
}

// GetSpanSessionOk returns a tuple with the SpanSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSourceVirtualIf) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanSession.Get(), o.SpanSession.IsSet()
}

// HasSpanSession returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourceVirtualIf) HasSpanSession() bool {
	if o != nil && o.SpanSession.IsSet() {
		return true
	}

	return false
}

// SetSpanSession gets a reference to the given NullableFabricSpanSessionRelationship and assigns it to the SpanSession field.
func (o *FabricAbstractSpanSourceVirtualIf) SetSpanSession(v FabricSpanSessionRelationship) {
	o.SpanSession.Set(&v)
}

// SetSpanSessionNil sets the value for SpanSession to be an explicit nil
func (o *FabricAbstractSpanSourceVirtualIf) SetSpanSessionNil() {
	o.SpanSession.Set(nil)
}

// UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil
func (o *FabricAbstractSpanSourceVirtualIf) UnsetSpanSession() {
	o.SpanSession.Unset()
}

func (o FabricAbstractSpanSourceVirtualIf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanSourceVirtualIf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanSource, errFabricAbstractSpanSource := json.Marshal(o.FabricAbstractSpanSource)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	errFabricAbstractSpanSource = json.Unmarshal([]byte(serializedFabricAbstractSpanSource), &toSerialize)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
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
	if !IsNil(o.VifId) {
		toSerialize["VifId"] = o.VifId
	}
	if o.SpanSession.IsSet() {
		toSerialize["SpanSession"] = o.SpanSession.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricAbstractSpanSourceVirtualIf) UnmarshalJSON(data []byte) (err error) {
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
	type FabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the VNIC referenced by vnic relationship. Vnic name is not updated if it gets updated by the user.
		Name *string `json:"Name,omitempty"`
		// VIF ID of the VNIC placed on Fabric Interconnect associated to the SPAN session.
		VifId       *int64                                `json:"VifId,omitempty"`
		SpanSession NullableFabricSpanSessionRelationship `json:"SpanSession,omitempty"`
	}

	varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct := FabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanSourceVirtualIf := _FabricAbstractSpanSourceVirtualIf{}
		varFabricAbstractSpanSourceVirtualIf.ClassId = varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanSourceVirtualIf.ObjectType = varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanSourceVirtualIf.Name = varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct.Name
		varFabricAbstractSpanSourceVirtualIf.VifId = varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct.VifId
		varFabricAbstractSpanSourceVirtualIf.SpanSession = varFabricAbstractSpanSourceVirtualIfWithoutEmbeddedStruct.SpanSession
		*o = FabricAbstractSpanSourceVirtualIf(varFabricAbstractSpanSourceVirtualIf)
	} else {
		return err
	}

	varFabricAbstractSpanSourceVirtualIf := _FabricAbstractSpanSourceVirtualIf{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourceVirtualIf)
	if err == nil {
		o.FabricAbstractSpanSource = varFabricAbstractSpanSourceVirtualIf.FabricAbstractSpanSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "VifId")
		delete(additionalProperties, "SpanSession")

		// remove fields from embedded structs
		reflectFabricAbstractSpanSource := reflect.ValueOf(o.FabricAbstractSpanSource)
		for i := 0; i < reflectFabricAbstractSpanSource.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanSource.Type().Field(i)

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

type NullableFabricAbstractSpanSourceVirtualIf struct {
	value *FabricAbstractSpanSourceVirtualIf
	isSet bool
}

func (v NullableFabricAbstractSpanSourceVirtualIf) Get() *FabricAbstractSpanSourceVirtualIf {
	return v.value
}

func (v *NullableFabricAbstractSpanSourceVirtualIf) Set(val *FabricAbstractSpanSourceVirtualIf) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanSourceVirtualIf) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanSourceVirtualIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanSourceVirtualIf(val *FabricAbstractSpanSourceVirtualIf) *NullableFabricAbstractSpanSourceVirtualIf {
	return &NullableFabricAbstractSpanSourceVirtualIf{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanSourceVirtualIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanSourceVirtualIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
