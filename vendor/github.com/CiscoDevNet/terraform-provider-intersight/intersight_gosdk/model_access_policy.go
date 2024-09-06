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

// checks if the AccessPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicy{}

// AccessPolicy Policy to configure server or chassis management options.
type AccessPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                          `json:"ObjectType"`
	AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
	ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
	// VLAN to be used for server access over Inband network.
	InbandVlan      *int64                                       `json:"InbandVlan,omitempty"`
	InbandIpPool    NullableIppoolPoolRelationship               `json:"InbandIpPool,omitempty"`
	InbandVrf       NullableVrfVrfRelationship                   `json:"InbandVrf,omitempty"`
	Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	OutOfBandIpPool NullableIppoolPoolRelationship               `json:"OutOfBandIpPool,omitempty"`
	OutOfBandVrf    NullableVrfVrfRelationship                   `json:"OutOfBandVrf,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicy AccessPolicy

// NewAccessPolicy instantiates a new AccessPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicy(classId string, objectType string) *AccessPolicy {
	this := AccessPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessPolicyWithDefaults instantiates a new AccessPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyWithDefaults() *AccessPolicy {
	this := AccessPolicy{}
	var classId string = "access.Policy"
	this.ClassId = classId
	var objectType string = "access.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "access.Policy" of the ClassId field.
func (o *AccessPolicy) GetDefaultClassId() interface{} {
	return "access.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *AccessPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "access.Policy" of the ObjectType field.
func (o *AccessPolicy) GetDefaultObjectType() interface{} {
	return "access.Policy"
}

// GetAddressType returns the AddressType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetAddressType() AccessAddressType {
	if o == nil || IsNil(o.AddressType.Get()) {
		var ret AccessAddressType
		return ret
	}
	return *o.AddressType.Get()
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetAddressTypeOk() (*AccessAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressType.Get(), o.AddressType.IsSet()
}

// HasAddressType returns a boolean if a field has been set.
func (o *AccessPolicy) HasAddressType() bool {
	if o != nil && o.AddressType.IsSet() {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given NullableAccessAddressType and assigns it to the AddressType field.
func (o *AccessPolicy) SetAddressType(v AccessAddressType) {
	o.AddressType.Set(&v)
}

// SetAddressTypeNil sets the value for AddressType to be an explicit nil
func (o *AccessPolicy) SetAddressTypeNil() {
	o.AddressType.Set(nil)
}

// UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
func (o *AccessPolicy) UnsetAddressType() {
	o.AddressType.Unset()
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetConfigurationType() AccessConfigurationType {
	if o == nil || IsNil(o.ConfigurationType.Get()) {
		var ret AccessConfigurationType
		return ret
	}
	return *o.ConfigurationType.Get()
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetConfigurationTypeOk() (*AccessConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationType.Get(), o.ConfigurationType.IsSet()
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AccessPolicy) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType.IsSet() {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given NullableAccessConfigurationType and assigns it to the ConfigurationType field.
func (o *AccessPolicy) SetConfigurationType(v AccessConfigurationType) {
	o.ConfigurationType.Set(&v)
}

// SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil
func (o *AccessPolicy) SetConfigurationTypeNil() {
	o.ConfigurationType.Set(nil)
}

// UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
func (o *AccessPolicy) UnsetConfigurationType() {
	o.ConfigurationType.Unset()
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *AccessPolicy) GetInbandVlan() int64 {
	if o == nil || IsNil(o.InbandVlan) {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicy) GetInbandVlanOk() (*int64, bool) {
	if o == nil || IsNil(o.InbandVlan) {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *AccessPolicy) HasInbandVlan() bool {
	if o != nil && !IsNil(o.InbandVlan) {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *AccessPolicy) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetInbandIpPool returns the InbandIpPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetInbandIpPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.InbandIpPool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InbandIpPool.Get()
}

// GetInbandIpPoolOk returns a tuple with the InbandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InbandIpPool.Get(), o.InbandIpPool.IsSet()
}

// HasInbandIpPool returns a boolean if a field has been set.
func (o *AccessPolicy) HasInbandIpPool() bool {
	if o != nil && o.InbandIpPool.IsSet() {
		return true
	}

	return false
}

// SetInbandIpPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the InbandIpPool field.
func (o *AccessPolicy) SetInbandIpPool(v IppoolPoolRelationship) {
	o.InbandIpPool.Set(&v)
}

// SetInbandIpPoolNil sets the value for InbandIpPool to be an explicit nil
func (o *AccessPolicy) SetInbandIpPoolNil() {
	o.InbandIpPool.Set(nil)
}

// UnsetInbandIpPool ensures that no value is present for InbandIpPool, not even an explicit nil
func (o *AccessPolicy) UnsetInbandIpPool() {
	o.InbandIpPool.Unset()
}

// GetInbandVrf returns the InbandVrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetInbandVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.InbandVrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.InbandVrf.Get()
}

// GetInbandVrfOk returns a tuple with the InbandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetInbandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InbandVrf.Get(), o.InbandVrf.IsSet()
}

// HasInbandVrf returns a boolean if a field has been set.
func (o *AccessPolicy) HasInbandVrf() bool {
	if o != nil && o.InbandVrf.IsSet() {
		return true
	}

	return false
}

// SetInbandVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the InbandVrf field.
func (o *AccessPolicy) SetInbandVrf(v VrfVrfRelationship) {
	o.InbandVrf.Set(&v)
}

// SetInbandVrfNil sets the value for InbandVrf to be an explicit nil
func (o *AccessPolicy) SetInbandVrfNil() {
	o.InbandVrf.Set(nil)
}

// UnsetInbandVrf ensures that no value is present for InbandVrf, not even an explicit nil
func (o *AccessPolicy) UnsetInbandVrf() {
	o.InbandVrf.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *AccessPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *AccessPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *AccessPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *AccessPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetOutOfBandIpPool returns the OutOfBandIpPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetOutOfBandIpPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.OutOfBandIpPool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.OutOfBandIpPool.Get()
}

// GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutOfBandIpPool.Get(), o.OutOfBandIpPool.IsSet()
}

// HasOutOfBandIpPool returns a boolean if a field has been set.
func (o *AccessPolicy) HasOutOfBandIpPool() bool {
	if o != nil && o.OutOfBandIpPool.IsSet() {
		return true
	}

	return false
}

// SetOutOfBandIpPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the OutOfBandIpPool field.
func (o *AccessPolicy) SetOutOfBandIpPool(v IppoolPoolRelationship) {
	o.OutOfBandIpPool.Set(&v)
}

// SetOutOfBandIpPoolNil sets the value for OutOfBandIpPool to be an explicit nil
func (o *AccessPolicy) SetOutOfBandIpPoolNil() {
	o.OutOfBandIpPool.Set(nil)
}

// UnsetOutOfBandIpPool ensures that no value is present for OutOfBandIpPool, not even an explicit nil
func (o *AccessPolicy) UnsetOutOfBandIpPool() {
	o.OutOfBandIpPool.Unset()
}

// GetOutOfBandVrf returns the OutOfBandVrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetOutOfBandVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.OutOfBandVrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.OutOfBandVrf.Get()
}

// GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutOfBandVrf.Get(), o.OutOfBandVrf.IsSet()
}

// HasOutOfBandVrf returns a boolean if a field has been set.
func (o *AccessPolicy) HasOutOfBandVrf() bool {
	if o != nil && o.OutOfBandVrf.IsSet() {
		return true
	}

	return false
}

// SetOutOfBandVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the OutOfBandVrf field.
func (o *AccessPolicy) SetOutOfBandVrf(v VrfVrfRelationship) {
	o.OutOfBandVrf.Set(&v)
}

// SetOutOfBandVrfNil sets the value for OutOfBandVrf to be an explicit nil
func (o *AccessPolicy) SetOutOfBandVrfNil() {
	o.OutOfBandVrf.Set(nil)
}

// UnsetOutOfBandVrf ensures that no value is present for OutOfBandVrf, not even an explicit nil
func (o *AccessPolicy) UnsetOutOfBandVrf() {
	o.OutOfBandVrf.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *AccessPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *AccessPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o AccessPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicy) ToMap() (map[string]interface{}, error) {
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
	if o.AddressType.IsSet() {
		toSerialize["AddressType"] = o.AddressType.Get()
	}
	if o.ConfigurationType.IsSet() {
		toSerialize["ConfigurationType"] = o.ConfigurationType.Get()
	}
	if !IsNil(o.InbandVlan) {
		toSerialize["InbandVlan"] = o.InbandVlan
	}
	if o.InbandIpPool.IsSet() {
		toSerialize["InbandIpPool"] = o.InbandIpPool.Get()
	}
	if o.InbandVrf.IsSet() {
		toSerialize["InbandVrf"] = o.InbandVrf.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.OutOfBandIpPool.IsSet() {
		toSerialize["OutOfBandIpPool"] = o.OutOfBandIpPool.Get()
	}
	if o.OutOfBandVrf.IsSet() {
		toSerialize["OutOfBandVrf"] = o.OutOfBandVrf.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type AccessPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                          `json:"ObjectType"`
		AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
		ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
		// VLAN to be used for server access over Inband network.
		InbandVlan      *int64                                       `json:"InbandVlan,omitempty"`
		InbandIpPool    NullableIppoolPoolRelationship               `json:"InbandIpPool,omitempty"`
		InbandVrf       NullableVrfVrfRelationship                   `json:"InbandVrf,omitempty"`
		Organization    NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		OutOfBandIpPool NullableIppoolPoolRelationship               `json:"OutOfBandIpPool,omitempty"`
		OutOfBandVrf    NullableVrfVrfRelationship                   `json:"OutOfBandVrf,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varAccessPolicyWithoutEmbeddedStruct := AccessPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAccessPolicyWithoutEmbeddedStruct)
	if err == nil {
		varAccessPolicy := _AccessPolicy{}
		varAccessPolicy.ClassId = varAccessPolicyWithoutEmbeddedStruct.ClassId
		varAccessPolicy.ObjectType = varAccessPolicyWithoutEmbeddedStruct.ObjectType
		varAccessPolicy.AddressType = varAccessPolicyWithoutEmbeddedStruct.AddressType
		varAccessPolicy.ConfigurationType = varAccessPolicyWithoutEmbeddedStruct.ConfigurationType
		varAccessPolicy.InbandVlan = varAccessPolicyWithoutEmbeddedStruct.InbandVlan
		varAccessPolicy.InbandIpPool = varAccessPolicyWithoutEmbeddedStruct.InbandIpPool
		varAccessPolicy.InbandVrf = varAccessPolicyWithoutEmbeddedStruct.InbandVrf
		varAccessPolicy.Organization = varAccessPolicyWithoutEmbeddedStruct.Organization
		varAccessPolicy.OutOfBandIpPool = varAccessPolicyWithoutEmbeddedStruct.OutOfBandIpPool
		varAccessPolicy.OutOfBandVrf = varAccessPolicyWithoutEmbeddedStruct.OutOfBandVrf
		varAccessPolicy.Profiles = varAccessPolicyWithoutEmbeddedStruct.Profiles
		*o = AccessPolicy(varAccessPolicy)
	} else {
		return err
	}

	varAccessPolicy := _AccessPolicy{}

	err = json.Unmarshal(data, &varAccessPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varAccessPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AddressType")
		delete(additionalProperties, "ConfigurationType")
		delete(additionalProperties, "InbandVlan")
		delete(additionalProperties, "InbandIpPool")
		delete(additionalProperties, "InbandVrf")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OutOfBandIpPool")
		delete(additionalProperties, "OutOfBandVrf")
		delete(additionalProperties, "Profiles")

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

type NullableAccessPolicy struct {
	value *AccessPolicy
	isSet bool
}

func (v NullableAccessPolicy) Get() *AccessPolicy {
	return v.value
}

func (v *NullableAccessPolicy) Set(val *AccessPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicy(val *AccessPolicy) *NullableAccessPolicy {
	return &NullableAccessPolicy{value: val, isSet: true}
}

func (v NullableAccessPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
