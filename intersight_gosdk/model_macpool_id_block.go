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

// checks if the MacpoolIdBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacpoolIdBlock{}

// MacpoolIdBlock A block of contiguous MAC addresses that are part of a pool.
type MacpoolIdBlock struct {
	PoolAbstractBlock
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                          `json:"ObjectType"`
	MacBlock   *MacpoolBlock                   `json:"MacBlock,omitempty"`
	Pool       NullableMacpoolPoolRelationship `json:"Pool,omitempty"`
	// An array of relationships to macpoolReservation resources.
	Reservations         []MacpoolReservationRelationship `json:"Reservations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MacpoolIdBlock MacpoolIdBlock

// NewMacpoolIdBlock instantiates a new MacpoolIdBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacpoolIdBlock(classId string, objectType string) *MacpoolIdBlock {
	this := MacpoolIdBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMacpoolIdBlockWithDefaults instantiates a new MacpoolIdBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacpoolIdBlockWithDefaults() *MacpoolIdBlock {
	this := MacpoolIdBlock{}
	var classId string = "macpool.IdBlock"
	this.ClassId = classId
	var objectType string = "macpool.IdBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MacpoolIdBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MacpoolIdBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "macpool.IdBlock" of the ClassId field.
func (o *MacpoolIdBlock) GetDefaultClassId() interface{} {
	return "macpool.IdBlock"
}

// GetObjectType returns the ObjectType field value
func (o *MacpoolIdBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MacpoolIdBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "macpool.IdBlock" of the ObjectType field.
func (o *MacpoolIdBlock) GetDefaultObjectType() interface{} {
	return "macpool.IdBlock"
}

// GetMacBlock returns the MacBlock field value if set, zero value otherwise.
func (o *MacpoolIdBlock) GetMacBlock() MacpoolBlock {
	if o == nil || IsNil(o.MacBlock) {
		var ret MacpoolBlock
		return ret
	}
	return *o.MacBlock
}

// GetMacBlockOk returns a tuple with the MacBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MacpoolIdBlock) GetMacBlockOk() (*MacpoolBlock, bool) {
	if o == nil || IsNil(o.MacBlock) {
		return nil, false
	}
	return o.MacBlock, true
}

// HasMacBlock returns a boolean if a field has been set.
func (o *MacpoolIdBlock) HasMacBlock() bool {
	if o != nil && !IsNil(o.MacBlock) {
		return true
	}

	return false
}

// SetMacBlock gets a reference to the given MacpoolBlock and assigns it to the MacBlock field.
func (o *MacpoolIdBlock) SetMacBlock(v MacpoolBlock) {
	o.MacBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolIdBlock) GetPool() MacpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolIdBlock) GetPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *MacpoolIdBlock) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableMacpoolPoolRelationship and assigns it to the Pool field.
func (o *MacpoolIdBlock) SetPool(v MacpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *MacpoolIdBlock) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *MacpoolIdBlock) UnsetPool() {
	o.Pool.Unset()
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MacpoolIdBlock) GetReservations() []MacpoolReservationRelationship {
	if o == nil {
		var ret []MacpoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MacpoolIdBlock) GetReservationsOk() ([]MacpoolReservationRelationship, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *MacpoolIdBlock) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []MacpoolReservationRelationship and assigns it to the Reservations field.
func (o *MacpoolIdBlock) SetReservations(v []MacpoolReservationRelationship) {
	o.Reservations = v
}

func (o MacpoolIdBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacpoolIdBlock) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MacBlock) {
		toSerialize["MacBlock"] = o.MacBlock
	}
	if o.Pool.IsSet() {
		toSerialize["Pool"] = o.Pool.Get()
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MacpoolIdBlock) UnmarshalJSON(data []byte) (err error) {
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
	type MacpoolIdBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                          `json:"ObjectType"`
		MacBlock   *MacpoolBlock                   `json:"MacBlock,omitempty"`
		Pool       NullableMacpoolPoolRelationship `json:"Pool,omitempty"`
		// An array of relationships to macpoolReservation resources.
		Reservations []MacpoolReservationRelationship `json:"Reservations,omitempty"`
	}

	varMacpoolIdBlockWithoutEmbeddedStruct := MacpoolIdBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMacpoolIdBlockWithoutEmbeddedStruct)
	if err == nil {
		varMacpoolIdBlock := _MacpoolIdBlock{}
		varMacpoolIdBlock.ClassId = varMacpoolIdBlockWithoutEmbeddedStruct.ClassId
		varMacpoolIdBlock.ObjectType = varMacpoolIdBlockWithoutEmbeddedStruct.ObjectType
		varMacpoolIdBlock.MacBlock = varMacpoolIdBlockWithoutEmbeddedStruct.MacBlock
		varMacpoolIdBlock.Pool = varMacpoolIdBlockWithoutEmbeddedStruct.Pool
		varMacpoolIdBlock.Reservations = varMacpoolIdBlockWithoutEmbeddedStruct.Reservations
		*o = MacpoolIdBlock(varMacpoolIdBlock)
	} else {
		return err
	}

	varMacpoolIdBlock := _MacpoolIdBlock{}

	err = json.Unmarshal(data, &varMacpoolIdBlock)
	if err == nil {
		o.PoolAbstractBlock = varMacpoolIdBlock.PoolAbstractBlock
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MacBlock")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Reservations")

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

type NullableMacpoolIdBlock struct {
	value *MacpoolIdBlock
	isSet bool
}

func (v NullableMacpoolIdBlock) Get() *MacpoolIdBlock {
	return v.value
}

func (v *NullableMacpoolIdBlock) Set(val *MacpoolIdBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableMacpoolIdBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableMacpoolIdBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacpoolIdBlock(val *MacpoolIdBlock) *NullableMacpoolIdBlock {
	return &NullableMacpoolIdBlock{value: val, isSet: true}
}

func (v NullableMacpoolIdBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacpoolIdBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
