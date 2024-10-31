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

// checks if the FabricBaseSwitchProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricBaseSwitchProfile{}

// FabricBaseSwitchProfile Abstract base class of the Switch Profile and Template.
type FabricBaseSwitchProfile struct {
	PolicyAbstractConfigProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Value indicating the switch side on which the switch profile or template has to be deployed. * `None` - Switch side not defined for the policy configurations in the switch profile or template. * `A` - Policy configurations in the switch profile or template to be deployed on fabric interconnect A. * `B` - Policy configurations in the switch profile or template to be deployed on fabric interconnect B.
	SwitchId             *string                                `json:"SwitchId,omitempty"`
	ConfigResult         NullableFabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricBaseSwitchProfile FabricBaseSwitchProfile

// NewFabricBaseSwitchProfile instantiates a new FabricBaseSwitchProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricBaseSwitchProfile(classId string, objectType string) *FabricBaseSwitchProfile {
	this := FabricBaseSwitchProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	var action string = "No-op"
	this.Action = &action
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// NewFabricBaseSwitchProfileWithDefaults instantiates a new FabricBaseSwitchProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricBaseSwitchProfileWithDefaults() *FabricBaseSwitchProfile {
	this := FabricBaseSwitchProfile{}
	var switchId string = "None"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricBaseSwitchProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricBaseSwitchProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricBaseSwitchProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricBaseSwitchProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricBaseSwitchProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricBaseSwitchProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *FabricBaseSwitchProfile) GetSwitchId() string {
	if o == nil || IsNil(o.SwitchId) {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricBaseSwitchProfile) GetSwitchIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchId) {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *FabricBaseSwitchProfile) HasSwitchId() bool {
	if o != nil && !IsNil(o.SwitchId) {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *FabricBaseSwitchProfile) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricBaseSwitchProfile) GetConfigResult() FabricConfigResultRelationship {
	if o == nil || IsNil(o.ConfigResult.Get()) {
		var ret FabricConfigResultRelationship
		return ret
	}
	return *o.ConfigResult.Get()
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricBaseSwitchProfile) GetConfigResultOk() (*FabricConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigResult.Get(), o.ConfigResult.IsSet()
}

// HasConfigResult returns a boolean if a field has been set.
func (o *FabricBaseSwitchProfile) HasConfigResult() bool {
	if o != nil && o.ConfigResult.IsSet() {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given NullableFabricConfigResultRelationship and assigns it to the ConfigResult field.
func (o *FabricBaseSwitchProfile) SetConfigResult(v FabricConfigResultRelationship) {
	o.ConfigResult.Set(&v)
}

// SetConfigResultNil sets the value for ConfigResult to be an explicit nil
func (o *FabricBaseSwitchProfile) SetConfigResultNil() {
	o.ConfigResult.Set(nil)
}

// UnsetConfigResult ensures that no value is present for ConfigResult, not even an explicit nil
func (o *FabricBaseSwitchProfile) UnsetConfigResult() {
	o.ConfigResult.Unset()
}

func (o FabricBaseSwitchProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricBaseSwitchProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigProfile, errPolicyAbstractConfigProfile := json.Marshal(o.PolicyAbstractConfigProfile)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	errPolicyAbstractConfigProfile = json.Unmarshal([]byte(serializedPolicyAbstractConfigProfile), &toSerialize)
	if errPolicyAbstractConfigProfile != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigProfile
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.SwitchId) {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.ConfigResult.IsSet() {
		toSerialize["ConfigResult"] = o.ConfigResult.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricBaseSwitchProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type FabricBaseSwitchProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Value indicating the switch side on which the switch profile or template has to be deployed. * `None` - Switch side not defined for the policy configurations in the switch profile or template. * `A` - Policy configurations in the switch profile or template to be deployed on fabric interconnect A. * `B` - Policy configurations in the switch profile or template to be deployed on fabric interconnect B.
		SwitchId     *string                                `json:"SwitchId,omitempty"`
		ConfigResult NullableFabricConfigResultRelationship `json:"ConfigResult,omitempty"`
	}

	varFabricBaseSwitchProfileWithoutEmbeddedStruct := FabricBaseSwitchProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricBaseSwitchProfileWithoutEmbeddedStruct)
	if err == nil {
		varFabricBaseSwitchProfile := _FabricBaseSwitchProfile{}
		varFabricBaseSwitchProfile.ClassId = varFabricBaseSwitchProfileWithoutEmbeddedStruct.ClassId
		varFabricBaseSwitchProfile.ObjectType = varFabricBaseSwitchProfileWithoutEmbeddedStruct.ObjectType
		varFabricBaseSwitchProfile.SwitchId = varFabricBaseSwitchProfileWithoutEmbeddedStruct.SwitchId
		varFabricBaseSwitchProfile.ConfigResult = varFabricBaseSwitchProfileWithoutEmbeddedStruct.ConfigResult
		*o = FabricBaseSwitchProfile(varFabricBaseSwitchProfile)
	} else {
		return err
	}

	varFabricBaseSwitchProfile := _FabricBaseSwitchProfile{}

	err = json.Unmarshal(data, &varFabricBaseSwitchProfile)
	if err == nil {
		o.PolicyAbstractConfigProfile = varFabricBaseSwitchProfile.PolicyAbstractConfigProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "ConfigResult")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigProfile := reflect.ValueOf(o.PolicyAbstractConfigProfile)
		for i := 0; i < reflectPolicyAbstractConfigProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigProfile.Type().Field(i)

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

type NullableFabricBaseSwitchProfile struct {
	value *FabricBaseSwitchProfile
	isSet bool
}

func (v NullableFabricBaseSwitchProfile) Get() *FabricBaseSwitchProfile {
	return v.value
}

func (v *NullableFabricBaseSwitchProfile) Set(val *FabricBaseSwitchProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricBaseSwitchProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricBaseSwitchProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricBaseSwitchProfile(val *FabricBaseSwitchProfile) *NullableFabricBaseSwitchProfile {
	return &NullableFabricBaseSwitchProfile{value: val, isSet: true}
}

func (v NullableFabricBaseSwitchProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricBaseSwitchProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
