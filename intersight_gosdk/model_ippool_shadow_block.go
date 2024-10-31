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

// checks if the IppoolShadowBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IppoolShadowBlock{}

// IppoolShadowBlock A block of Contiguous IP addresses that are part of a shadow pool.
type IppoolShadowBlock struct {
	PoolAbstractBlock
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of this IP addresses blocks. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
	IpType    *string                              `json:"IpType,omitempty"`
	IpV4Block *IppoolIpV4Block                     `json:"IpV4Block,omitempty"`
	IpV6Block *IppoolIpV6Block                     `json:"IpV6Block,omitempty"`
	Pool      NullableIppoolShadowPoolRelationship `json:"Pool,omitempty"`
	// An array of relationships to ippoolReservation resources.
	Reservations         []IppoolReservationRelationship `json:"Reservations,omitempty"`
	Vrf                  NullableVrfVrfRelationship      `json:"Vrf,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolShadowBlock IppoolShadowBlock

// NewIppoolShadowBlock instantiates a new IppoolShadowBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolShadowBlock(classId string, objectType string) *IppoolShadowBlock {
	this := IppoolShadowBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// NewIppoolShadowBlockWithDefaults instantiates a new IppoolShadowBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolShadowBlockWithDefaults() *IppoolShadowBlock {
	this := IppoolShadowBlock{}
	var classId string = "ippool.ShadowBlock"
	this.ClassId = classId
	var objectType string = "ippool.ShadowBlock"
	this.ObjectType = objectType
	var ipType string = "IPv4"
	this.IpType = &ipType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolShadowBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolShadowBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "ippool.ShadowBlock" of the ClassId field.
func (o *IppoolShadowBlock) GetDefaultClassId() interface{} {
	return "ippool.ShadowBlock"
}

// GetObjectType returns the ObjectType field value
func (o *IppoolShadowBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolShadowBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "ippool.ShadowBlock" of the ObjectType field.
func (o *IppoolShadowBlock) GetDefaultObjectType() interface{} {
	return "ippool.ShadowBlock"
}

// GetIpType returns the IpType field value if set, zero value otherwise.
func (o *IppoolShadowBlock) GetIpType() string {
	if o == nil || IsNil(o.IpType) {
		var ret string
		return ret
	}
	return *o.IpType
}

// GetIpTypeOk returns a tuple with the IpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetIpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IpType) {
		return nil, false
	}
	return o.IpType, true
}

// HasIpType returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasIpType() bool {
	if o != nil && !IsNil(o.IpType) {
		return true
	}

	return false
}

// SetIpType gets a reference to the given string and assigns it to the IpType field.
func (o *IppoolShadowBlock) SetIpType(v string) {
	o.IpType = &v
}

// GetIpV4Block returns the IpV4Block field value if set, zero value otherwise.
func (o *IppoolShadowBlock) GetIpV4Block() IppoolIpV4Block {
	if o == nil || IsNil(o.IpV4Block) {
		var ret IppoolIpV4Block
		return ret
	}
	return *o.IpV4Block
}

// GetIpV4BlockOk returns a tuple with the IpV4Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetIpV4BlockOk() (*IppoolIpV4Block, bool) {
	if o == nil || IsNil(o.IpV4Block) {
		return nil, false
	}
	return o.IpV4Block, true
}

// HasIpV4Block returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasIpV4Block() bool {
	if o != nil && !IsNil(o.IpV4Block) {
		return true
	}

	return false
}

// SetIpV4Block gets a reference to the given IppoolIpV4Block and assigns it to the IpV4Block field.
func (o *IppoolShadowBlock) SetIpV4Block(v IppoolIpV4Block) {
	o.IpV4Block = &v
}

// GetIpV6Block returns the IpV6Block field value if set, zero value otherwise.
func (o *IppoolShadowBlock) GetIpV6Block() IppoolIpV6Block {
	if o == nil || IsNil(o.IpV6Block) {
		var ret IppoolIpV6Block
		return ret
	}
	return *o.IpV6Block
}

// GetIpV6BlockOk returns a tuple with the IpV6Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolShadowBlock) GetIpV6BlockOk() (*IppoolIpV6Block, bool) {
	if o == nil || IsNil(o.IpV6Block) {
		return nil, false
	}
	return o.IpV6Block, true
}

// HasIpV6Block returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasIpV6Block() bool {
	if o != nil && !IsNil(o.IpV6Block) {
		return true
	}

	return false
}

// SetIpV6Block gets a reference to the given IppoolIpV6Block and assigns it to the IpV6Block field.
func (o *IppoolShadowBlock) SetIpV6Block(v IppoolIpV6Block) {
	o.IpV6Block = &v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowBlock) GetPool() IppoolShadowPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret IppoolShadowPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowBlock) GetPoolOk() (*IppoolShadowPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableIppoolShadowPoolRelationship and assigns it to the Pool field.
func (o *IppoolShadowBlock) SetPool(v IppoolShadowPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *IppoolShadowBlock) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *IppoolShadowBlock) UnsetPool() {
	o.Pool.Unset()
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowBlock) GetReservations() []IppoolReservationRelationship {
	if o == nil {
		var ret []IppoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowBlock) GetReservationsOk() ([]IppoolReservationRelationship, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []IppoolReservationRelationship and assigns it to the Reservations field.
func (o *IppoolShadowBlock) SetReservations(v []IppoolReservationRelationship) {
	o.Reservations = v
}

// GetVrf returns the Vrf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IppoolShadowBlock) GetVrf() VrfVrfRelationship {
	if o == nil || IsNil(o.Vrf.Get()) {
		var ret VrfVrfRelationship
		return ret
	}
	return *o.Vrf.Get()
}

// GetVrfOk returns a tuple with the Vrf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IppoolShadowBlock) GetVrfOk() (*VrfVrfRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vrf.Get(), o.Vrf.IsSet()
}

// HasVrf returns a boolean if a field has been set.
func (o *IppoolShadowBlock) HasVrf() bool {
	if o != nil && o.Vrf.IsSet() {
		return true
	}

	return false
}

// SetVrf gets a reference to the given NullableVrfVrfRelationship and assigns it to the Vrf field.
func (o *IppoolShadowBlock) SetVrf(v VrfVrfRelationship) {
	o.Vrf.Set(&v)
}

// SetVrfNil sets the value for Vrf to be an explicit nil
func (o *IppoolShadowBlock) SetVrfNil() {
	o.Vrf.Set(nil)
}

// UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
func (o *IppoolShadowBlock) UnsetVrf() {
	o.Vrf.Unset()
}

func (o IppoolShadowBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IppoolShadowBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlock, errPoolAbstractBlock := json.Marshal(o.PoolAbstractBlock)
	if errPoolAbstractBlock != nil {
		return map[string]interface{}{}, errPoolAbstractBlock
	}
	errPoolAbstractBlock = json.Unmarshal([]byte(serializedPoolAbstractBlock), &toSerialize)
	if errPoolAbstractBlock != nil {
		return map[string]interface{}{}, errPoolAbstractBlock
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IpType) {
		toSerialize["IpType"] = o.IpType
	}
	if !IsNil(o.IpV4Block) {
		toSerialize["IpV4Block"] = o.IpV4Block
	}
	if !IsNil(o.IpV6Block) {
		toSerialize["IpV6Block"] = o.IpV6Block
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}
	if o.Vrf.IsSet() {
		toSerialize["Vrf"] = o.Vrf.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IppoolShadowBlock) UnmarshalJSON(data []byte) (err error) {
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
	type IppoolShadowBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Type of this IP addresses blocks. * `IPv4` - IP V4 address type requested. * `IPv6` - IP V6 address type requested.
		IpType    *string                              `json:"IpType,omitempty"`
		IpV4Block *IppoolIpV4Block                     `json:"IpV4Block,omitempty"`
		IpV6Block *IppoolIpV6Block                     `json:"IpV6Block,omitempty"`
		Pool      NullableIppoolShadowPoolRelationship `json:"Pool,omitempty"`
		// An array of relationships to ippoolReservation resources.
		Reservations []IppoolReservationRelationship `json:"Reservations,omitempty"`
		Vrf          NullableVrfVrfRelationship      `json:"Vrf,omitempty"`
	}

	varIppoolShadowBlockWithoutEmbeddedStruct := IppoolShadowBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIppoolShadowBlockWithoutEmbeddedStruct)
	if err == nil {
		varIppoolShadowBlock := _IppoolShadowBlock{}
		varIppoolShadowBlock.ClassId = varIppoolShadowBlockWithoutEmbeddedStruct.ClassId
		varIppoolShadowBlock.ObjectType = varIppoolShadowBlockWithoutEmbeddedStruct.ObjectType
		varIppoolShadowBlock.IpType = varIppoolShadowBlockWithoutEmbeddedStruct.IpType
		varIppoolShadowBlock.IpV4Block = varIppoolShadowBlockWithoutEmbeddedStruct.IpV4Block
		varIppoolShadowBlock.IpV6Block = varIppoolShadowBlockWithoutEmbeddedStruct.IpV6Block
		varIppoolShadowBlock.Pool = varIppoolShadowBlockWithoutEmbeddedStruct.Pool
		varIppoolShadowBlock.Reservations = varIppoolShadowBlockWithoutEmbeddedStruct.Reservations
		varIppoolShadowBlock.Vrf = varIppoolShadowBlockWithoutEmbeddedStruct.Vrf
		*o = IppoolShadowBlock(varIppoolShadowBlock)
	} else {
		return err
	}

	varIppoolShadowBlock := _IppoolShadowBlock{}

	err = json.Unmarshal(data, &varIppoolShadowBlock)
	if err == nil {
		o.PoolAbstractBlock = varIppoolShadowBlock.PoolAbstractBlock
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpType")
		delete(additionalProperties, "IpV4Block")
		delete(additionalProperties, "IpV6Block")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Reservations")
		delete(additionalProperties, "Vrf")

		// remove fields from embedded structs
		reflectPoolAbstractBlock := reflect.ValueOf(o.PoolAbstractBlock)
		for i := 0; i < reflectPoolAbstractBlock.Type().NumField(); i++ {
			t := reflectPoolAbstractBlock.Type().Field(i)

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

type NullableIppoolShadowBlock struct {
	value *IppoolShadowBlock
	isSet bool
}

func (v NullableIppoolShadowBlock) Get() *IppoolShadowBlock {
	return v.value
}

func (v *NullableIppoolShadowBlock) Set(val *IppoolShadowBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolShadowBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolShadowBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolShadowBlock(val *IppoolShadowBlock) *NullableIppoolShadowBlock {
	return &NullableIppoolShadowBlock{value: val, isSet: true}
}

func (v NullableIppoolShadowBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolShadowBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
