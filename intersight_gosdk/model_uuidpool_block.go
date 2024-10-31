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

// checks if the UuidpoolBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UuidpoolBlock{}

// UuidpoolBlock A block of contiguous UUID addresses that are part of a pool.
type UuidpoolBlock struct {
	PoolAbstractBlock
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Prefix of the UUID pool. UUID is constructed as <prefix>-<suffix>.
	Prefix          *string                          `json:"Prefix,omitempty"`
	UuidSuffixBlock *UuidpoolUuidBlock               `json:"UuidSuffixBlock,omitempty"`
	Pool            NullableUuidpoolPoolRelationship `json:"Pool,omitempty"`
	// An array of relationships to uuidpoolReservation resources.
	Reservations         []UuidpoolReservationRelationship `json:"Reservations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UuidpoolBlock UuidpoolBlock

// NewUuidpoolBlock instantiates a new UuidpoolBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUuidpoolBlock(classId string, objectType string) *UuidpoolBlock {
	this := UuidpoolBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewUuidpoolBlockWithDefaults instantiates a new UuidpoolBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUuidpoolBlockWithDefaults() *UuidpoolBlock {
	this := UuidpoolBlock{}
	var classId string = "uuidpool.Block"
	this.ClassId = classId
	var objectType string = "uuidpool.Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *UuidpoolBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *UuidpoolBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *UuidpoolBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "uuidpool.Block" of the ClassId field.
func (o *UuidpoolBlock) GetDefaultClassId() interface{} {
	return "uuidpool.Block"
}

// GetObjectType returns the ObjectType field value
func (o *UuidpoolBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *UuidpoolBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *UuidpoolBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "uuidpool.Block" of the ObjectType field.
func (o *UuidpoolBlock) GetDefaultObjectType() interface{} {
	return "uuidpool.Block"
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UuidpoolBlock) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolBlock) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UuidpoolBlock) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UuidpoolBlock) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUuidSuffixBlock returns the UuidSuffixBlock field value if set, zero value otherwise.
func (o *UuidpoolBlock) GetUuidSuffixBlock() UuidpoolUuidBlock {
	if o == nil || IsNil(o.UuidSuffixBlock) {
		var ret UuidpoolUuidBlock
		return ret
	}
	return *o.UuidSuffixBlock
}

// GetUuidSuffixBlockOk returns a tuple with the UuidSuffixBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UuidpoolBlock) GetUuidSuffixBlockOk() (*UuidpoolUuidBlock, bool) {
	if o == nil || IsNil(o.UuidSuffixBlock) {
		return nil, false
	}
	return o.UuidSuffixBlock, true
}

// HasUuidSuffixBlock returns a boolean if a field has been set.
func (o *UuidpoolBlock) HasUuidSuffixBlock() bool {
	if o != nil && !IsNil(o.UuidSuffixBlock) {
		return true
	}

	return false
}

// SetUuidSuffixBlock gets a reference to the given UuidpoolUuidBlock and assigns it to the UuidSuffixBlock field.
func (o *UuidpoolBlock) SetUuidSuffixBlock(v UuidpoolUuidBlock) {
	o.UuidSuffixBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolBlock) GetPool() UuidpoolPoolRelationship {
	if o == nil || IsNil(o.Pool.Get()) {
		var ret UuidpoolPoolRelationship
		return ret
	}
	return *o.Pool.Get()
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolBlock) GetPoolOk() (*UuidpoolPoolRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pool.Get(), o.Pool.IsSet()
}

// HasPool returns a boolean if a field has been set.
func (o *UuidpoolBlock) HasPool() bool {
	if o != nil && o.Pool.IsSet() {
		return true
	}

	return false
}

// SetPool gets a reference to the given NullableUuidpoolPoolRelationship and assigns it to the Pool field.
func (o *UuidpoolBlock) SetPool(v UuidpoolPoolRelationship) {
	o.Pool.Set(&v)
}

// SetPoolNil sets the value for Pool to be an explicit nil
func (o *UuidpoolBlock) SetPoolNil() {
	o.Pool.Set(nil)
}

// UnsetPool ensures that no value is present for Pool, not even an explicit nil
func (o *UuidpoolBlock) UnsetPool() {
	o.Pool.Unset()
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UuidpoolBlock) GetReservations() []UuidpoolReservationRelationship {
	if o == nil {
		var ret []UuidpoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UuidpoolBlock) GetReservationsOk() ([]UuidpoolReservationRelationship, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *UuidpoolBlock) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []UuidpoolReservationRelationship and assigns it to the Reservations field.
func (o *UuidpoolBlock) SetReservations(v []UuidpoolReservationRelationship) {
	o.Reservations = v
}

func (o UuidpoolBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UuidpoolBlock) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Prefix) {
		toSerialize["Prefix"] = o.Prefix
	}
	if !IsNil(o.UuidSuffixBlock) {
		toSerialize["UuidSuffixBlock"] = o.UuidSuffixBlock
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

func (o *UuidpoolBlock) UnmarshalJSON(data []byte) (err error) {
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
	type UuidpoolBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Prefix of the UUID pool. UUID is constructed as <prefix>-<suffix>.
		Prefix          *string                          `json:"Prefix,omitempty"`
		UuidSuffixBlock *UuidpoolUuidBlock               `json:"UuidSuffixBlock,omitempty"`
		Pool            NullableUuidpoolPoolRelationship `json:"Pool,omitempty"`
		// An array of relationships to uuidpoolReservation resources.
		Reservations []UuidpoolReservationRelationship `json:"Reservations,omitempty"`
	}

	varUuidpoolBlockWithoutEmbeddedStruct := UuidpoolBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUuidpoolBlockWithoutEmbeddedStruct)
	if err == nil {
		varUuidpoolBlock := _UuidpoolBlock{}
		varUuidpoolBlock.ClassId = varUuidpoolBlockWithoutEmbeddedStruct.ClassId
		varUuidpoolBlock.ObjectType = varUuidpoolBlockWithoutEmbeddedStruct.ObjectType
		varUuidpoolBlock.Prefix = varUuidpoolBlockWithoutEmbeddedStruct.Prefix
		varUuidpoolBlock.UuidSuffixBlock = varUuidpoolBlockWithoutEmbeddedStruct.UuidSuffixBlock
		varUuidpoolBlock.Pool = varUuidpoolBlockWithoutEmbeddedStruct.Pool
		varUuidpoolBlock.Reservations = varUuidpoolBlockWithoutEmbeddedStruct.Reservations
		*o = UuidpoolBlock(varUuidpoolBlock)
	} else {
		return err
	}

	varUuidpoolBlock := _UuidpoolBlock{}

	err = json.Unmarshal(data, &varUuidpoolBlock)
	if err == nil {
		o.PoolAbstractBlock = varUuidpoolBlock.PoolAbstractBlock
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "UuidSuffixBlock")
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

type NullableUuidpoolBlock struct {
	value *UuidpoolBlock
	isSet bool
}

func (v NullableUuidpoolBlock) Get() *UuidpoolBlock {
	return v.value
}

func (v *NullableUuidpoolBlock) Set(val *UuidpoolBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableUuidpoolBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableUuidpoolBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuidpoolBlock(val *UuidpoolBlock) *NullableUuidpoolBlock {
	return &NullableUuidpoolBlock{value: val, isSet: true}
}

func (v NullableUuidpoolBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuidpoolBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
