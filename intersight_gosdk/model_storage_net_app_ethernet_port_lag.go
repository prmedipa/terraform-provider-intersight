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

// checks if the StorageNetAppEthernetPortLag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppEthernetPortLag{}

// StorageNetAppEthernetPortLag Storage port of type lag.
type StorageNetAppEthernetPortLag struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string              `json:"ObjectType"`
	ActivePorts []StorageNetAppPort `json:"ActivePorts,omitempty"`
	// Policy for mapping flows to ports for outbound packets through a LAG (ifgrp). * `none` - Default unknown distribution policy type. * `port` - Network traffic is distributed based on the transport layer (TCP/UDP) ports. * `ip` - Network traffic is distributed based on IP addresses. * `mac` - Network traffic is distributed based on MAC addresses. * `sequential` - Network traffic is distributed in round-robin fashion from the list of configured, available ports.
	DistributionPolicy *string             `json:"DistributionPolicy,omitempty"`
	MemberPorts        []StorageNetAppPort `json:"MemberPorts,omitempty"`
	// Determines how the ports interact with the switch. * `none` - Default unknown lag mode type. * `multimode_lacp` - Bundle multiple member ports of the interface group using Link Aggregation Control Protocol. * `multimode` - Bundle multiple member ports of the interface group to act as a single trunked port. * `singlemode` - Provide port redundancy using member ports of the interface group for failover.
	Mode                 *string `json:"Mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppEthernetPortLag StorageNetAppEthernetPortLag

// NewStorageNetAppEthernetPortLag instantiates a new StorageNetAppEthernetPortLag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppEthernetPortLag(classId string, objectType string) *StorageNetAppEthernetPortLag {
	this := StorageNetAppEthernetPortLag{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppEthernetPortLagWithDefaults instantiates a new StorageNetAppEthernetPortLag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppEthernetPortLagWithDefaults() *StorageNetAppEthernetPortLag {
	this := StorageNetAppEthernetPortLag{}
	var classId string = "storage.NetAppEthernetPortLag"
	this.ClassId = classId
	var objectType string = "storage.NetAppEthernetPortLag"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppEthernetPortLag) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortLag) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppEthernetPortLag) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppEthernetPortLag" of the ClassId field.
func (o *StorageNetAppEthernetPortLag) GetDefaultClassId() interface{} {
	return "storage.NetAppEthernetPortLag"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppEthernetPortLag) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortLag) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppEthernetPortLag) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppEthernetPortLag" of the ObjectType field.
func (o *StorageNetAppEthernetPortLag) GetDefaultObjectType() interface{} {
	return "storage.NetAppEthernetPortLag"
}

// GetActivePorts returns the ActivePorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortLag) GetActivePorts() []StorageNetAppPort {
	if o == nil {
		var ret []StorageNetAppPort
		return ret
	}
	return o.ActivePorts
}

// GetActivePortsOk returns a tuple with the ActivePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortLag) GetActivePortsOk() ([]StorageNetAppPort, bool) {
	if o == nil || IsNil(o.ActivePorts) {
		return nil, false
	}
	return o.ActivePorts, true
}

// HasActivePorts returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortLag) HasActivePorts() bool {
	if o != nil && !IsNil(o.ActivePorts) {
		return true
	}

	return false
}

// SetActivePorts gets a reference to the given []StorageNetAppPort and assigns it to the ActivePorts field.
func (o *StorageNetAppEthernetPortLag) SetActivePorts(v []StorageNetAppPort) {
	o.ActivePorts = v
}

// GetDistributionPolicy returns the DistributionPolicy field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortLag) GetDistributionPolicy() string {
	if o == nil || IsNil(o.DistributionPolicy) {
		var ret string
		return ret
	}
	return *o.DistributionPolicy
}

// GetDistributionPolicyOk returns a tuple with the DistributionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortLag) GetDistributionPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DistributionPolicy) {
		return nil, false
	}
	return o.DistributionPolicy, true
}

// HasDistributionPolicy returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortLag) HasDistributionPolicy() bool {
	if o != nil && !IsNil(o.DistributionPolicy) {
		return true
	}

	return false
}

// SetDistributionPolicy gets a reference to the given string and assigns it to the DistributionPolicy field.
func (o *StorageNetAppEthernetPortLag) SetDistributionPolicy(v string) {
	o.DistributionPolicy = &v
}

// GetMemberPorts returns the MemberPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppEthernetPortLag) GetMemberPorts() []StorageNetAppPort {
	if o == nil {
		var ret []StorageNetAppPort
		return ret
	}
	return o.MemberPorts
}

// GetMemberPortsOk returns a tuple with the MemberPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppEthernetPortLag) GetMemberPortsOk() ([]StorageNetAppPort, bool) {
	if o == nil || IsNil(o.MemberPorts) {
		return nil, false
	}
	return o.MemberPorts, true
}

// HasMemberPorts returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortLag) HasMemberPorts() bool {
	if o != nil && !IsNil(o.MemberPorts) {
		return true
	}

	return false
}

// SetMemberPorts gets a reference to the given []StorageNetAppPort and assigns it to the MemberPorts field.
func (o *StorageNetAppEthernetPortLag) SetMemberPorts(v []StorageNetAppPort) {
	o.MemberPorts = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *StorageNetAppEthernetPortLag) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppEthernetPortLag) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StorageNetAppEthernetPortLag) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StorageNetAppEthernetPortLag) SetMode(v string) {
	o.Mode = &v
}

func (o StorageNetAppEthernetPortLag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppEthernetPortLag) ToMap() (map[string]interface{}, error) {
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
	if o.ActivePorts != nil {
		toSerialize["ActivePorts"] = o.ActivePorts
	}
	if !IsNil(o.DistributionPolicy) {
		toSerialize["DistributionPolicy"] = o.DistributionPolicy
	}
	if o.MemberPorts != nil {
		toSerialize["MemberPorts"] = o.MemberPorts
	}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppEthernetPortLag) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppEthernetPortLagWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string              `json:"ObjectType"`
		ActivePorts []StorageNetAppPort `json:"ActivePorts,omitempty"`
		// Policy for mapping flows to ports for outbound packets through a LAG (ifgrp). * `none` - Default unknown distribution policy type. * `port` - Network traffic is distributed based on the transport layer (TCP/UDP) ports. * `ip` - Network traffic is distributed based on IP addresses. * `mac` - Network traffic is distributed based on MAC addresses. * `sequential` - Network traffic is distributed in round-robin fashion from the list of configured, available ports.
		DistributionPolicy *string             `json:"DistributionPolicy,omitempty"`
		MemberPorts        []StorageNetAppPort `json:"MemberPorts,omitempty"`
		// Determines how the ports interact with the switch. * `none` - Default unknown lag mode type. * `multimode_lacp` - Bundle multiple member ports of the interface group using Link Aggregation Control Protocol. * `multimode` - Bundle multiple member ports of the interface group to act as a single trunked port. * `singlemode` - Provide port redundancy using member ports of the interface group for failover.
		Mode *string `json:"Mode,omitempty"`
	}

	varStorageNetAppEthernetPortLagWithoutEmbeddedStruct := StorageNetAppEthernetPortLagWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPortLagWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppEthernetPortLag := _StorageNetAppEthernetPortLag{}
		varStorageNetAppEthernetPortLag.ClassId = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.ClassId
		varStorageNetAppEthernetPortLag.ObjectType = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.ObjectType
		varStorageNetAppEthernetPortLag.ActivePorts = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.ActivePorts
		varStorageNetAppEthernetPortLag.DistributionPolicy = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.DistributionPolicy
		varStorageNetAppEthernetPortLag.MemberPorts = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.MemberPorts
		varStorageNetAppEthernetPortLag.Mode = varStorageNetAppEthernetPortLagWithoutEmbeddedStruct.Mode
		*o = StorageNetAppEthernetPortLag(varStorageNetAppEthernetPortLag)
	} else {
		return err
	}

	varStorageNetAppEthernetPortLag := _StorageNetAppEthernetPortLag{}

	err = json.Unmarshal(data, &varStorageNetAppEthernetPortLag)
	if err == nil {
		o.MoBaseComplexType = varStorageNetAppEthernetPortLag.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActivePorts")
		delete(additionalProperties, "DistributionPolicy")
		delete(additionalProperties, "MemberPorts")
		delete(additionalProperties, "Mode")

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

type NullableStorageNetAppEthernetPortLag struct {
	value *StorageNetAppEthernetPortLag
	isSet bool
}

func (v NullableStorageNetAppEthernetPortLag) Get() *StorageNetAppEthernetPortLag {
	return v.value
}

func (v *NullableStorageNetAppEthernetPortLag) Set(val *StorageNetAppEthernetPortLag) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppEthernetPortLag) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppEthernetPortLag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppEthernetPortLag(val *StorageNetAppEthernetPortLag) *NullableStorageNetAppEthernetPortLag {
	return &NullableStorageNetAppEthernetPortLag{value: val, isSet: true}
}

func (v NullableStorageNetAppEthernetPortLag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppEthernetPortLag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
