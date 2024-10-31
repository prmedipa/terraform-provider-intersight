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

// checks if the FabricLanPinGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricLanPinGroup{}

// FabricLanPinGroup LAN PinGroup configuration sent by user for static pinning.
type FabricLanPinGroup struct {
	FabricPinGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType             string                                          `json:"ObjectType"`
	PinTargetInterfaceRole NullableFabricAbstractInterfaceRoleRelationship `json:"PinTargetInterfaceRole,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _FabricLanPinGroup FabricLanPinGroup

// NewFabricLanPinGroup instantiates a new FabricLanPinGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricLanPinGroup(classId string, objectType string) *FabricLanPinGroup {
	this := FabricLanPinGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricLanPinGroupWithDefaults instantiates a new FabricLanPinGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricLanPinGroupWithDefaults() *FabricLanPinGroup {
	this := FabricLanPinGroup{}
	var classId string = "fabric.LanPinGroup"
	this.ClassId = classId
	var objectType string = "fabric.LanPinGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricLanPinGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricLanPinGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricLanPinGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.LanPinGroup" of the ClassId field.
func (o *FabricLanPinGroup) GetDefaultClassId() interface{} {
	return "fabric.LanPinGroup"
}

// GetObjectType returns the ObjectType field value
func (o *FabricLanPinGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricLanPinGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricLanPinGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.LanPinGroup" of the ObjectType field.
func (o *FabricLanPinGroup) GetDefaultObjectType() interface{} {
	return "fabric.LanPinGroup"
}

// GetPinTargetInterfaceRole returns the PinTargetInterfaceRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricLanPinGroup) GetPinTargetInterfaceRole() FabricAbstractInterfaceRoleRelationship {
	if o == nil || IsNil(o.PinTargetInterfaceRole.Get()) {
		var ret FabricAbstractInterfaceRoleRelationship
		return ret
	}
	return *o.PinTargetInterfaceRole.Get()
}

// GetPinTargetInterfaceRoleOk returns a tuple with the PinTargetInterfaceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricLanPinGroup) GetPinTargetInterfaceRoleOk() (*FabricAbstractInterfaceRoleRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinTargetInterfaceRole.Get(), o.PinTargetInterfaceRole.IsSet()
}

// HasPinTargetInterfaceRole returns a boolean if a field has been set.
func (o *FabricLanPinGroup) HasPinTargetInterfaceRole() bool {
	if o != nil && o.PinTargetInterfaceRole.IsSet() {
		return true
	}

	return false
}

// SetPinTargetInterfaceRole gets a reference to the given NullableFabricAbstractInterfaceRoleRelationship and assigns it to the PinTargetInterfaceRole field.
func (o *FabricLanPinGroup) SetPinTargetInterfaceRole(v FabricAbstractInterfaceRoleRelationship) {
	o.PinTargetInterfaceRole.Set(&v)
}

// SetPinTargetInterfaceRoleNil sets the value for PinTargetInterfaceRole to be an explicit nil
func (o *FabricLanPinGroup) SetPinTargetInterfaceRoleNil() {
	o.PinTargetInterfaceRole.Set(nil)
}

// UnsetPinTargetInterfaceRole ensures that no value is present for PinTargetInterfaceRole, not even an explicit nil
func (o *FabricLanPinGroup) UnsetPinTargetInterfaceRole() {
	o.PinTargetInterfaceRole.Unset()
}

func (o FabricLanPinGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricLanPinGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPinGroup, errFabricPinGroup := json.Marshal(o.FabricPinGroup)
	if errFabricPinGroup != nil {
		return map[string]interface{}{}, errFabricPinGroup
	}
	errFabricPinGroup = json.Unmarshal([]byte(serializedFabricPinGroup), &toSerialize)
	if errFabricPinGroup != nil {
		return map[string]interface{}{}, errFabricPinGroup
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.PinTargetInterfaceRole.IsSet() {
		toSerialize["PinTargetInterfaceRole"] = o.PinTargetInterfaceRole.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricLanPinGroup) UnmarshalJSON(data []byte) (err error) {
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
	type FabricLanPinGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType             string                                          `json:"ObjectType"`
		PinTargetInterfaceRole NullableFabricAbstractInterfaceRoleRelationship `json:"PinTargetInterfaceRole,omitempty"`
	}

	varFabricLanPinGroupWithoutEmbeddedStruct := FabricLanPinGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricLanPinGroupWithoutEmbeddedStruct)
	if err == nil {
		varFabricLanPinGroup := _FabricLanPinGroup{}
		varFabricLanPinGroup.ClassId = varFabricLanPinGroupWithoutEmbeddedStruct.ClassId
		varFabricLanPinGroup.ObjectType = varFabricLanPinGroupWithoutEmbeddedStruct.ObjectType
		varFabricLanPinGroup.PinTargetInterfaceRole = varFabricLanPinGroupWithoutEmbeddedStruct.PinTargetInterfaceRole
		*o = FabricLanPinGroup(varFabricLanPinGroup)
	} else {
		return err
	}

	varFabricLanPinGroup := _FabricLanPinGroup{}

	err = json.Unmarshal(data, &varFabricLanPinGroup)
	if err == nil {
		o.FabricPinGroup = varFabricLanPinGroup.FabricPinGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PinTargetInterfaceRole")

		// remove fields from embedded structs
		reflectFabricPinGroup := reflect.ValueOf(o.FabricPinGroup)
		for i := 0; i < reflectFabricPinGroup.Type().NumField(); i++ {
			t := reflectFabricPinGroup.Type().Field(i)

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

type NullableFabricLanPinGroup struct {
	value *FabricLanPinGroup
	isSet bool
}

func (v NullableFabricLanPinGroup) Get() *FabricLanPinGroup {
	return v.value
}

func (v *NullableFabricLanPinGroup) Set(val *FabricLanPinGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricLanPinGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricLanPinGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricLanPinGroup(val *FabricLanPinGroup) *NullableFabricLanPinGroup {
	return &NullableFabricLanPinGroup{value: val, isSet: true}
}

func (v NullableFabricLanPinGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricLanPinGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
