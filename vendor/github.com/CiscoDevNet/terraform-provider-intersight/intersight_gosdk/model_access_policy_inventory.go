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

// checks if the AccessPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyInventory{}

// AccessPolicyInventory Policy to configure server or chassis management options.
type AccessPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                          `json:"ObjectType"`
	AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
	ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
	// VLAN to be used for server access over Inband network.
	InbandVlan           *int64                         `json:"InbandVlan,omitempty"`
	InbandIpPool         NullableIppoolPoolRelationship `json:"InbandIpPool,omitempty"`
	InbandVrf            NullableVrfVrfRelationship     `json:"InbandVrf,omitempty"`
	OutOfBandIpPool      NullableIppoolPoolRelationship `json:"OutOfBandIpPool,omitempty"`
	OutOfBandVrf         NullableVrfVrfRelationship     `json:"OutOfBandVrf,omitempty"`
	TargetMo             NullableMoBaseMoRelationship   `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyInventory AccessPolicyInventory

// NewAccessPolicyInventory instantiates a new AccessPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyInventory(classId string, objectType string) *AccessPolicyInventory {
	this := AccessPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAccessPolicyInventoryWithDefaults instantiates a new AccessPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyInventoryWithDefaults() *AccessPolicyInventory {
	this := AccessPolicyInventory{}
	var classId string = "access.PolicyInventory"
	this.ClassId = classId
	var objectType string = "access.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AccessPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AccessPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "access.PolicyInventory" of the ClassId field.
func (o *AccessPolicyInventory) GetDefaultClassId() interface{} {
	return "access.PolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *AccessPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AccessPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "access.PolicyInventory" of the ObjectType field.
func (o *AccessPolicyInventory) GetDefaultObjectType() interface{} {
	return "access.PolicyInventory"
}

// GetAddressType returns the AddressType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetAddressType() AccessAddressType {
	if o == nil || IsNil(o.AddressType.Get()) {
		var ret AccessAddressType
		return ret
	}
	return *o.AddressType.Get()
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetAddressTypeOk() (*AccessAddressType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddressType.Get(), o.AddressType.IsSet()
}

// HasAddressType returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasAddressType() bool {
	if o != nil && o.AddressType.IsSet() {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given NullableAccessAddressType and assigns it to the AddressType field.
func (o *AccessPolicyInventory) SetAddressType(v AccessAddressType) {
	o.AddressType.Set(&v)
}

// SetAddressTypeNil sets the value for AddressType to be an explicit nil
func (o *AccessPolicyInventory) SetAddressTypeNil() {
	o.AddressType.Set(nil)
}

// UnsetAddressType ensures that no value is present for AddressType, not even an explicit nil
func (o *AccessPolicyInventory) UnsetAddressType() {
	o.AddressType.Unset()
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetConfigurationType() AccessConfigurationType {
	if o == nil || IsNil(o.ConfigurationType.Get()) {
		var ret AccessConfigurationType
		return ret
	}
	return *o.ConfigurationType.Get()
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetConfigurationTypeOk() (*AccessConfigurationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationType.Get(), o.ConfigurationType.IsSet()
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasConfigurationType() bool {
	if o != nil && o.ConfigurationType.IsSet() {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given NullableAccessConfigurationType and assigns it to the ConfigurationType field.
func (o *AccessPolicyInventory) SetConfigurationType(v AccessConfigurationType) {
	o.ConfigurationType.Set(&v)
}

// SetConfigurationTypeNil sets the value for ConfigurationType to be an explicit nil
func (o *AccessPolicyInventory) SetConfigurationTypeNil() {
	o.ConfigurationType.Set(nil)
}

// UnsetConfigurationType ensures that no value is present for ConfigurationType, not even an explicit nil
func (o *AccessPolicyInventory) UnsetConfigurationType() {
	o.ConfigurationType.Unset()
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *AccessPolicyInventory) GetInbandVlan() int64 {
	if o == nil || IsNil(o.InbandVlan) {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyInventory) GetInbandVlanOk() (*int64, bool) {
	if o == nil || IsNil(o.InbandVlan) {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandVlan() bool {
	if o != nil && !IsNil(o.InbandVlan) {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *AccessPolicyInventory) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetInbandIpPool returns the InbandIpPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetInbandIpPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.InbandIpPool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.InbandIpPool.Get()
}

// GetInbandIpPoolOk returns a tuple with the InbandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetInbandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InbandIpPool.Get(), o.InbandIpPool.IsSet()
}

// HasInbandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandIpPool() bool {
	if o != nil && o.InbandIpPool.IsSet() {
		return true
	}

	return false
}

// SetInbandIpPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the InbandIpPool field.
func (o *AccessPolicyInventory) SetInbandIpPool(v IppoolPoolRelationship) {
	o.InbandIpPool.Set(&v)
}

// SetInbandIpPoolNil sets the value for InbandIpPool to be an explicit nil
func (o *AccessPolicyInventory) SetInbandIpPoolNil() {
	o.InbandIpPool.Set(nil)
}

// UnsetInbandIpPool ensures that no value is present for InbandIpPool, not even an explicit nil
func (o *AccessPolicyInventory) UnsetInbandIpPool() {
	o.InbandIpPool.Unset()
}

// GetInbandVrf returns the InbandVrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetInbandVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.InbandVrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.InbandVrf.Get()
}

// GetInbandVrfOk returns a tuple with the InbandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetInbandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InbandVrf.Get(), o.InbandVrf.IsSet()
}

// HasInbandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasInbandVrf() bool {
	if o != nil && o.InbandVrf.IsSet() {
		return true
	}

	return false
}

// SetInbandVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the InbandVrf field.
func (o *AccessPolicyInventory) SetInbandVrf(v VrfVrfRelationship) {
	o.InbandVrf.Set(&v)
}

// SetInbandVrfNil sets the value for InbandVrf to be an explicit nil
func (o *AccessPolicyInventory) SetInbandVrfNil() {
	o.InbandVrf.Set(nil)
}

// UnsetInbandVrf ensures that no value is present for InbandVrf, not even an explicit nil
func (o *AccessPolicyInventory) UnsetInbandVrf() {
	o.InbandVrf.Unset()
}

// GetOutOfBandIpPool returns the OutOfBandIpPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetOutOfBandIpPool() IppoolPoolRelationship {
	if o == nil || IsNil(o.OutOfBandIpPool.Get()) {
		var ret IppoolPoolRelationship
		return ret
	}
	return *o.OutOfBandIpPool.Get()
}

// GetOutOfBandIpPoolOk returns a tuple with the OutOfBandIpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetOutOfBandIpPoolOk() (*IppoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutOfBandIpPool.Get(), o.OutOfBandIpPool.IsSet()
}

// HasOutOfBandIpPool returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasOutOfBandIpPool() bool {
	if o != nil && o.OutOfBandIpPool.IsSet() {
		return true
	}

	return false
}

// SetOutOfBandIpPool gets a reference to the given NullableIppoolPoolRelationship and assigns it to the OutOfBandIpPool field.
func (o *AccessPolicyInventory) SetOutOfBandIpPool(v IppoolPoolRelationship) {
	o.OutOfBandIpPool.Set(&v)
}

// SetOutOfBandIpPoolNil sets the value for OutOfBandIpPool to be an explicit nil
func (o *AccessPolicyInventory) SetOutOfBandIpPoolNil() {
	o.OutOfBandIpPool.Set(nil)
}

// UnsetOutOfBandIpPool ensures that no value is present for OutOfBandIpPool, not even an explicit nil
func (o *AccessPolicyInventory) UnsetOutOfBandIpPool() {
	o.OutOfBandIpPool.Unset()
}

// GetOutOfBandVrf returns the OutOfBandVrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetOutOfBandVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.OutOfBandVrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.OutOfBandVrf.Get()
}

// GetOutOfBandVrfOk returns a tuple with the OutOfBandVrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetOutOfBandVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutOfBandVrf.Get(), o.OutOfBandVrf.IsSet()
}

// HasOutOfBandVrf returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasOutOfBandVrf() bool {
	if o != nil && o.OutOfBandVrf.IsSet() {
		return true
	}

	return false
}

// SetOutOfBandVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the OutOfBandVrf field.
func (o *AccessPolicyInventory) SetOutOfBandVrf(v VrfVrfRelationship) {
	o.OutOfBandVrf.Set(&v)
}

// SetOutOfBandVrfNil sets the value for OutOfBandVrf to be an explicit nil
func (o *AccessPolicyInventory) SetOutOfBandVrfNil() {
	o.OutOfBandVrf.Set(nil)
}

// UnsetOutOfBandVrf ensures that no value is present for OutOfBandVrf, not even an explicit nil
func (o *AccessPolicyInventory) UnsetOutOfBandVrf() {
	o.OutOfBandVrf.Unset()
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccessPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccessPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *AccessPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *AccessPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *AccessPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *AccessPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o AccessPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
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
	if o.OutOfBandIpPool.IsSet() {
		toSerialize["OutOfBandIpPool"] = o.OutOfBandIpPool.Get()
	}
	if o.OutOfBandVrf.IsSet() {
		toSerialize["OutOfBandVrf"] = o.OutOfBandVrf.Get()
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type AccessPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                          `json:"ObjectType"`
		AddressType       NullableAccessAddressType       `json:"AddressType,omitempty"`
		ConfigurationType NullableAccessConfigurationType `json:"ConfigurationType,omitempty"`
		// VLAN to be used for server access over Inband network.
		InbandVlan      *int64                         `json:"InbandVlan,omitempty"`
		InbandIpPool    NullableIppoolPoolRelationship `json:"InbandIpPool,omitempty"`
		InbandVrf       NullableVrfVrfRelationship     `json:"InbandVrf,omitempty"`
		OutOfBandIpPool NullableIppoolPoolRelationship `json:"OutOfBandIpPool,omitempty"`
		OutOfBandVrf    NullableVrfVrfRelationship     `json:"OutOfBandVrf,omitempty"`
		TargetMo        NullableMoBaseMoRelationship   `json:"TargetMo,omitempty"`
	}

	varAccessPolicyInventoryWithoutEmbeddedStruct := AccessPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAccessPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varAccessPolicyInventory := _AccessPolicyInventory{}
		varAccessPolicyInventory.ClassId = varAccessPolicyInventoryWithoutEmbeddedStruct.ClassId
		varAccessPolicyInventory.ObjectType = varAccessPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varAccessPolicyInventory.AddressType = varAccessPolicyInventoryWithoutEmbeddedStruct.AddressType
		varAccessPolicyInventory.ConfigurationType = varAccessPolicyInventoryWithoutEmbeddedStruct.ConfigurationType
		varAccessPolicyInventory.InbandVlan = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandVlan
		varAccessPolicyInventory.InbandIpPool = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandIpPool
		varAccessPolicyInventory.InbandVrf = varAccessPolicyInventoryWithoutEmbeddedStruct.InbandVrf
		varAccessPolicyInventory.OutOfBandIpPool = varAccessPolicyInventoryWithoutEmbeddedStruct.OutOfBandIpPool
		varAccessPolicyInventory.OutOfBandVrf = varAccessPolicyInventoryWithoutEmbeddedStruct.OutOfBandVrf
		varAccessPolicyInventory.TargetMo = varAccessPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = AccessPolicyInventory(varAccessPolicyInventory)
	} else {
		return err
	}

	varAccessPolicyInventory := _AccessPolicyInventory{}

	err = json.Unmarshal(data, &varAccessPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varAccessPolicyInventory.PolicyAbstractPolicyInventory
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
		delete(additionalProperties, "OutOfBandIpPool")
		delete(additionalProperties, "OutOfBandVrf")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableAccessPolicyInventory struct {
	value *AccessPolicyInventory
	isSet bool
}

func (v NullableAccessPolicyInventory) Get() *AccessPolicyInventory {
	return v.value
}

func (v *NullableAccessPolicyInventory) Set(val *AccessPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyInventory(val *AccessPolicyInventory) *NullableAccessPolicyInventory {
	return &NullableAccessPolicyInventory{value: val, isSet: true}
}

func (v NullableAccessPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
