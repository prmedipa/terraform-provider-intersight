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

// checks if the VnicEthNetworkPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicEthNetworkPolicy{}

// VnicEthNetworkPolicy An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.
type VnicEthNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	// Deprecated
	TargetPlatform       *string                                      `json:"TargetPlatform,omitempty"`
	VlanSettings         NullableVnicVlanSettings                     `json:"VlanSettings,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicEthNetworkPolicy VnicEthNetworkPolicy

// NewVnicEthNetworkPolicy instantiates a new VnicEthNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthNetworkPolicy(classId string, objectType string) *VnicEthNetworkPolicy {
	this := VnicEthNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// NewVnicEthNetworkPolicyWithDefaults instantiates a new VnicEthNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthNetworkPolicyWithDefaults() *VnicEthNetworkPolicy {
	this := VnicEthNetworkPolicy{}
	var classId string = "vnic.EthNetworkPolicy"
	this.ClassId = classId
	var objectType string = "vnic.EthNetworkPolicy"
	this.ObjectType = objectType
	var targetPlatform string = "Standalone"
	this.TargetPlatform = &targetPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicEthNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicEthNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.EthNetworkPolicy" of the ClassId field.
func (o *VnicEthNetworkPolicy) GetDefaultClassId() interface{} {
	return "vnic.EthNetworkPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *VnicEthNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicEthNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicEthNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.EthNetworkPolicy" of the ObjectType field.
func (o *VnicEthNetworkPolicy) GetDefaultObjectType() interface{} {
	return "vnic.EthNetworkPolicy"
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
// Deprecated
func (o *VnicEthNetworkPolicy) GetTargetPlatform() string {
	if o == nil || IsNil(o.TargetPlatform) {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VnicEthNetworkPolicy) GetTargetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPlatform) {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicy) HasTargetPlatform() bool {
	if o != nil && !IsNil(o.TargetPlatform) {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
// Deprecated
func (o *VnicEthNetworkPolicy) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetVlanSettings returns the VlanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthNetworkPolicy) GetVlanSettings() VnicVlanSettings {
	if o == nil || IsNil(o.VlanSettings.Get()) {
		var ret VnicVlanSettings
		return ret
	}
	return *o.VlanSettings.Get()
}

// GetVlanSettingsOk returns a tuple with the VlanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthNetworkPolicy) GetVlanSettingsOk() (*VnicVlanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanSettings.Get(), o.VlanSettings.IsSet()
}

// HasVlanSettings returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicy) HasVlanSettings() bool {
	if o != nil && o.VlanSettings.IsSet() {
		return true
	}

	return false
}

// SetVlanSettings gets a reference to the given NullableVnicVlanSettings and assigns it to the VlanSettings field.
func (o *VnicEthNetworkPolicy) SetVlanSettings(v VnicVlanSettings) {
	o.VlanSettings.Set(&v)
}

// SetVlanSettingsNil sets the value for VlanSettings to be an explicit nil
func (o *VnicEthNetworkPolicy) SetVlanSettingsNil() {
	o.VlanSettings.Set(nil)
}

// UnsetVlanSettings ensures that no value is present for VlanSettings, not even an explicit nil
func (o *VnicEthNetworkPolicy) UnsetVlanSettings() {
	o.VlanSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicEthNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicEthNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *VnicEthNetworkPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *VnicEthNetworkPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

func (o VnicEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicEthNetworkPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.TargetPlatform) {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.VlanSettings.IsSet() {
		toSerialize["VlanSettings"] = o.VlanSettings.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicEthNetworkPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type VnicEthNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
		// Deprecated
		TargetPlatform *string                                      `json:"TargetPlatform,omitempty"`
		VlanSettings   NullableVnicVlanSettings                     `json:"VlanSettings,omitempty"`
		Organization   NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicEthNetworkPolicyWithoutEmbeddedStruct := VnicEthNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicEthNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicEthNetworkPolicy := _VnicEthNetworkPolicy{}
		varVnicEthNetworkPolicy.ClassId = varVnicEthNetworkPolicyWithoutEmbeddedStruct.ClassId
		varVnicEthNetworkPolicy.ObjectType = varVnicEthNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varVnicEthNetworkPolicy.TargetPlatform = varVnicEthNetworkPolicyWithoutEmbeddedStruct.TargetPlatform
		varVnicEthNetworkPolicy.VlanSettings = varVnicEthNetworkPolicyWithoutEmbeddedStruct.VlanSettings
		varVnicEthNetworkPolicy.Organization = varVnicEthNetworkPolicyWithoutEmbeddedStruct.Organization
		*o = VnicEthNetworkPolicy(varVnicEthNetworkPolicy)
	} else {
		return err
	}

	varVnicEthNetworkPolicy := _VnicEthNetworkPolicy{}

	err = json.Unmarshal(data, &varVnicEthNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicEthNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "VlanSettings")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicEthNetworkPolicy struct {
	value *VnicEthNetworkPolicy
	isSet bool
}

func (v NullableVnicEthNetworkPolicy) Get() *VnicEthNetworkPolicy {
	return v.value
}

func (v *NullableVnicEthNetworkPolicy) Set(val *VnicEthNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthNetworkPolicy(val *VnicEthNetworkPolicy) *NullableVnicEthNetworkPolicy {
	return &NullableVnicEthNetworkPolicy{value: val, isSet: true}
}

func (v NullableVnicEthNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
