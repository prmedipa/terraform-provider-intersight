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

// checks if the AccessIpAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessIpAddress{}

// AccessIpAddress IP address and Lease information for a specific Server Profile.
type AccessIpAddress struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IPv4 Address leased to this Server Profile for In-Band Deployment.
	Ipv4Address *string `json:"Ipv4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	// IPv4 Address leased to this Server Profile for In-Band Deployment.
	Ipv6Address *string `json:"Ipv6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	// IPv4 Address leased to this Server Profile for Out-Of-Band deployment.
	OobIpv4Address       *string                                         `json:"OobIpv4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
	Ipv4Lease            NullableIppoolIpLeaseRelationship               `json:"Ipv4Lease,omitempty"`
	Ipv6Lease            NullableIppoolIpLeaseRelationship               `json:"Ipv6Lease,omitempty"`
	OobIpv4Lease         NullableIppoolIpLeaseRelationship               `json:"OobIpv4Lease,omitempty"`
	Profile              NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessIpAddress AccessIpAddress

// NewAccessIpAddress instantiates a new AccessIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessIpAddress(classId string, objectType string) *AccessIpAddress {
	this := AccessIpAddress{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessIpAddressWithDefaults instantiates a new AccessIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessIpAddressWithDefaults() *AccessIpAddress {
	this := AccessIpAddress{}
	var classId string = "access.IpAddress"
	this.ClassId = classId
	var objectType string = "access.IpAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessIpAddress) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessIpAddress) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessIpAddress) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "access.IpAddress" of the ClassId field.
func (o *AccessIpAddress) GetDefaultClassId() interface{} {
	return "access.IpAddress"
}

// GetObjectType returns the ObjectType field value
func (o *AccessIpAddress) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessIpAddress) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessIpAddress) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "access.IpAddress" of the ObjectType field.
func (o *AccessIpAddress) GetDefaultObjectType() interface{} {
	return "access.IpAddress"
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *AccessIpAddress) GetIpv4Address() string {
	if o == nil || IsNil(o.Ipv4Address) {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessIpAddress) GetIpv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Address) {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *AccessIpAddress) HasIpv4Address() bool {
	if o != nil && !IsNil(o.Ipv4Address) {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *AccessIpAddress) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *AccessIpAddress) GetIpv6Address() string {
	if o == nil || IsNil(o.Ipv6Address) {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessIpAddress) GetIpv6AddressOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Address) {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *AccessIpAddress) HasIpv6Address() bool {
	if o != nil && !IsNil(o.Ipv6Address) {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *AccessIpAddress) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetOobIpv4Address returns the OobIpv4Address field value if set, zero value otherwise.
func (o *AccessIpAddress) GetOobIpv4Address() string {
	if o == nil || IsNil(o.OobIpv4Address) {
		var ret string
		return ret
	}
	return *o.OobIpv4Address
}

// GetOobIpv4AddressOk returns a tuple with the OobIpv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessIpAddress) GetOobIpv4AddressOk() (*string, bool) {
	if o == nil || IsNil(o.OobIpv4Address) {
		return nil, false
	}
	return o.OobIpv4Address, true
}

// HasOobIpv4Address returns a boolean if a field has been set.
func (o *AccessIpAddress) HasOobIpv4Address() bool {
	if o != nil && !IsNil(o.OobIpv4Address) {
		return true
	}

	return false
}

// SetOobIpv4Address gets a reference to the given string and assigns it to the OobIpv4Address field.
func (o *AccessIpAddress) SetOobIpv4Address(v string) {
	o.OobIpv4Address = &v
}

// GetIpv4Lease returns the Ipv4Lease field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessIpAddress) GetIpv4Lease() IppoolIpLeaseRelationship {
	if o == nil || IsNil(o.Ipv4Lease.Get()) {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.Ipv4Lease.Get()
}

// GetIpv4LeaseOk returns a tuple with the Ipv4Lease field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessIpAddress) GetIpv4LeaseOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv4Lease.Get(), o.Ipv4Lease.IsSet()
}

// HasIpv4Lease returns a boolean if a field has been set.
func (o *AccessIpAddress) HasIpv4Lease() bool {
	if o != nil && o.Ipv4Lease.IsSet() {
		return true
	}

	return false
}

// SetIpv4Lease gets a reference to the given NullableIppoolIpLeaseRelationship and assigns it to the Ipv4Lease field.
func (o *AccessIpAddress) SetIpv4Lease(v IppoolIpLeaseRelationship) {
	o.Ipv4Lease.Set(&v)
}

// SetIpv4LeaseNil sets the value for Ipv4Lease to be an explicit nil
func (o *AccessIpAddress) SetIpv4LeaseNil() {
	o.Ipv4Lease.Set(nil)
}

// UnsetIpv4Lease ensures that no value is present for Ipv4Lease, not even an explicit nil
func (o *AccessIpAddress) UnsetIpv4Lease() {
	o.Ipv4Lease.Unset()
}

// GetIpv6Lease returns the Ipv6Lease field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessIpAddress) GetIpv6Lease() IppoolIpLeaseRelationship {
	if o == nil || IsNil(o.Ipv6Lease.Get()) {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.Ipv6Lease.Get()
}

// GetIpv6LeaseOk returns a tuple with the Ipv6Lease field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessIpAddress) GetIpv6LeaseOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6Lease.Get(), o.Ipv6Lease.IsSet()
}

// HasIpv6Lease returns a boolean if a field has been set.
func (o *AccessIpAddress) HasIpv6Lease() bool {
	if o != nil && o.Ipv6Lease.IsSet() {
		return true
	}

	return false
}

// SetIpv6Lease gets a reference to the given NullableIppoolIpLeaseRelationship and assigns it to the Ipv6Lease field.
func (o *AccessIpAddress) SetIpv6Lease(v IppoolIpLeaseRelationship) {
	o.Ipv6Lease.Set(&v)
}

// SetIpv6LeaseNil sets the value for Ipv6Lease to be an explicit nil
func (o *AccessIpAddress) SetIpv6LeaseNil() {
	o.Ipv6Lease.Set(nil)
}

// UnsetIpv6Lease ensures that no value is present for Ipv6Lease, not even an explicit nil
func (o *AccessIpAddress) UnsetIpv6Lease() {
	o.Ipv6Lease.Unset()
}

// GetOobIpv4Lease returns the OobIpv4Lease field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessIpAddress) GetOobIpv4Lease() IppoolIpLeaseRelationship {
	if o == nil || IsNil(o.OobIpv4Lease.Get()) {
		var ret IppoolIpLeaseRelationship
		return ret
	}
	return *o.OobIpv4Lease.Get()
}

// GetOobIpv4LeaseOk returns a tuple with the OobIpv4Lease field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessIpAddress) GetOobIpv4LeaseOk() (*IppoolIpLeaseRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OobIpv4Lease.Get(), o.OobIpv4Lease.IsSet()
}

// HasOobIpv4Lease returns a boolean if a field has been set.
func (o *AccessIpAddress) HasOobIpv4Lease() bool {
	if o != nil && o.OobIpv4Lease.IsSet() {
		return true
	}

	return false
}

// SetOobIpv4Lease gets a reference to the given NullableIppoolIpLeaseRelationship and assigns it to the OobIpv4Lease field.
func (o *AccessIpAddress) SetOobIpv4Lease(v IppoolIpLeaseRelationship) {
	o.OobIpv4Lease.Set(&v)
}

// SetOobIpv4LeaseNil sets the value for OobIpv4Lease to be an explicit nil
func (o *AccessIpAddress) SetOobIpv4LeaseNil() {
	o.OobIpv4Lease.Set(nil)
}

// UnsetOobIpv4Lease ensures that no value is present for OobIpv4Lease, not even an explicit nil
func (o *AccessIpAddress) UnsetOobIpv4Lease() {
	o.OobIpv4Lease.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessIpAddress) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || IsNil(o.Profile.Get()) {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessIpAddress) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *AccessIpAddress) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullablePolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *AccessIpAddress) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *AccessIpAddress) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *AccessIpAddress) UnsetProfile() {
	o.Profile.Unset()
}

func (o AccessIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessIpAddress) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Ipv4Address) {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if !IsNil(o.Ipv6Address) {
		toSerialize["Ipv6Address"] = o.Ipv6Address
	}
	if !IsNil(o.OobIpv4Address) {
		toSerialize["OobIpv4Address"] = o.OobIpv4Address
	}
	if o.Ipv4Lease.IsSet() {
		toSerialize["Ipv4Lease"] = o.Ipv4Lease.Get()
	}
	if o.Ipv6Lease.IsSet() {
		toSerialize["Ipv6Lease"] = o.Ipv6Lease.Get()
	}
	if o.OobIpv4Lease.IsSet() {
		toSerialize["OobIpv4Lease"] = o.OobIpv4Lease.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["Profile"] = o.Profile.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessIpAddress) UnmarshalJSON(data []byte) (err error) {
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
	type AccessIpAddressWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IPv4 Address leased to this Server Profile for In-Band Deployment.
		Ipv4Address *string `json:"Ipv4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		// IPv4 Address leased to this Server Profile for In-Band Deployment.
		Ipv6Address *string `json:"Ipv6Address,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		// IPv4 Address leased to this Server Profile for Out-Of-Band deployment.
		OobIpv4Address *string                                         `json:"OobIpv4Address,omitempty" validate:"regexp=^$|^([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\\\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$"`
		Ipv4Lease      NullableIppoolIpLeaseRelationship               `json:"Ipv4Lease,omitempty"`
		Ipv6Lease      NullableIppoolIpLeaseRelationship               `json:"Ipv6Lease,omitempty"`
		OobIpv4Lease   NullableIppoolIpLeaseRelationship               `json:"OobIpv4Lease,omitempty"`
		Profile        NullablePolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	}

	varAccessIpAddressWithoutEmbeddedStruct := AccessIpAddressWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAccessIpAddressWithoutEmbeddedStruct)
	if err == nil {
		varAccessIpAddress := _AccessIpAddress{}
		varAccessIpAddress.ClassId = varAccessIpAddressWithoutEmbeddedStruct.ClassId
		varAccessIpAddress.ObjectType = varAccessIpAddressWithoutEmbeddedStruct.ObjectType
		varAccessIpAddress.Ipv4Address = varAccessIpAddressWithoutEmbeddedStruct.Ipv4Address
		varAccessIpAddress.Ipv6Address = varAccessIpAddressWithoutEmbeddedStruct.Ipv6Address
		varAccessIpAddress.OobIpv4Address = varAccessIpAddressWithoutEmbeddedStruct.OobIpv4Address
		varAccessIpAddress.Ipv4Lease = varAccessIpAddressWithoutEmbeddedStruct.Ipv4Lease
		varAccessIpAddress.Ipv6Lease = varAccessIpAddressWithoutEmbeddedStruct.Ipv6Lease
		varAccessIpAddress.OobIpv4Lease = varAccessIpAddressWithoutEmbeddedStruct.OobIpv4Lease
		varAccessIpAddress.Profile = varAccessIpAddressWithoutEmbeddedStruct.Profile
		*o = AccessIpAddress(varAccessIpAddress)
	} else {
		return err
	}

	varAccessIpAddress := _AccessIpAddress{}

	err = json.Unmarshal(data, &varAccessIpAddress)
	if err == nil {
		o.PolicyAbstractPolicy = varAccessIpAddress.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv6Address")
		delete(additionalProperties, "OobIpv4Address")
		delete(additionalProperties, "Ipv4Lease")
		delete(additionalProperties, "Ipv6Lease")
		delete(additionalProperties, "OobIpv4Lease")
		delete(additionalProperties, "Profile")

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

type NullableAccessIpAddress struct {
	value *AccessIpAddress
	isSet bool
}

func (v NullableAccessIpAddress) Get() *AccessIpAddress {
	return v.value
}

func (v *NullableAccessIpAddress) Set(val *AccessIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessIpAddress(val *AccessIpAddress) *NullableAccessIpAddress {
	return &NullableAccessIpAddress{value: val, isSet: true}
}

func (v NullableAccessIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
