/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FcpoolFcBlock A block of contiguous WWN addresses that are part of a pool.
type FcpoolFcBlock struct {
	PoolAbstractBlock
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                  `json:"ObjectType"`
	IdBlock    *FcpoolBlock            `json:"IdBlock,omitempty"`
	Pool       *FcpoolPoolRelationship `json:"Pool,omitempty"`
	// An array of relationships to fcpoolReservation resources.
	Reservations         []FcpoolReservationRelationship `json:"Reservations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FcpoolFcBlock FcpoolFcBlock

// NewFcpoolFcBlock instantiates a new FcpoolFcBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpoolFcBlock(classId string, objectType string) *FcpoolFcBlock {
	this := FcpoolFcBlock{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFcpoolFcBlockWithDefaults instantiates a new FcpoolFcBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpoolFcBlockWithDefaults() *FcpoolFcBlock {
	this := FcpoolFcBlock{}
	var classId string = "fcpool.FcBlock"
	this.ClassId = classId
	var objectType string = "fcpool.FcBlock"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FcpoolFcBlock) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FcpoolFcBlock) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FcpoolFcBlock) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FcpoolFcBlock) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdBlock returns the IdBlock field value if set, zero value otherwise.
func (o *FcpoolFcBlock) GetIdBlock() FcpoolBlock {
	if o == nil || o.IdBlock == nil {
		var ret FcpoolBlock
		return ret
	}
	return *o.IdBlock
}

// GetIdBlockOk returns a tuple with the IdBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetIdBlockOk() (*FcpoolBlock, bool) {
	if o == nil || o.IdBlock == nil {
		return nil, false
	}
	return o.IdBlock, true
}

// HasIdBlock returns a boolean if a field has been set.
func (o *FcpoolFcBlock) HasIdBlock() bool {
	if o != nil && o.IdBlock != nil {
		return true
	}

	return false
}

// SetIdBlock gets a reference to the given FcpoolBlock and assigns it to the IdBlock field.
func (o *FcpoolFcBlock) SetIdBlock(v FcpoolBlock) {
	o.IdBlock = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *FcpoolFcBlock) GetPool() FcpoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret FcpoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpoolFcBlock) GetPoolOk() (*FcpoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *FcpoolFcBlock) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given FcpoolPoolRelationship and assigns it to the Pool field.
func (o *FcpoolFcBlock) SetPool(v FcpoolPoolRelationship) {
	o.Pool = &v
}

// GetReservations returns the Reservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FcpoolFcBlock) GetReservations() []FcpoolReservationRelationship {
	if o == nil {
		var ret []FcpoolReservationRelationship
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FcpoolFcBlock) GetReservationsOk() ([]FcpoolReservationRelationship, bool) {
	if o == nil || o.Reservations == nil {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *FcpoolFcBlock) HasReservations() bool {
	if o != nil && o.Reservations != nil {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []FcpoolReservationRelationship and assigns it to the Reservations field.
func (o *FcpoolFcBlock) SetReservations(v []FcpoolReservationRelationship) {
	o.Reservations = v
}

func (o FcpoolFcBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPoolAbstractBlock, errPoolAbstractBlock := json.Marshal(o.PoolAbstractBlock)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	errPoolAbstractBlock = json.Unmarshal([]byte(serializedPoolAbstractBlock), &toSerialize)
	if errPoolAbstractBlock != nil {
		return []byte{}, errPoolAbstractBlock
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IdBlock != nil {
		toSerialize["IdBlock"] = o.IdBlock
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.Reservations != nil {
		toSerialize["Reservations"] = o.Reservations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FcpoolFcBlock) UnmarshalJSON(bytes []byte) (err error) {
	type FcpoolFcBlockWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                  `json:"ObjectType"`
		IdBlock    *FcpoolBlock            `json:"IdBlock,omitempty"`
		Pool       *FcpoolPoolRelationship `json:"Pool,omitempty"`
		// An array of relationships to fcpoolReservation resources.
		Reservations []FcpoolReservationRelationship `json:"Reservations,omitempty"`
	}

	varFcpoolFcBlockWithoutEmbeddedStruct := FcpoolFcBlockWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFcpoolFcBlockWithoutEmbeddedStruct)
	if err == nil {
		varFcpoolFcBlock := _FcpoolFcBlock{}
		varFcpoolFcBlock.ClassId = varFcpoolFcBlockWithoutEmbeddedStruct.ClassId
		varFcpoolFcBlock.ObjectType = varFcpoolFcBlockWithoutEmbeddedStruct.ObjectType
		varFcpoolFcBlock.IdBlock = varFcpoolFcBlockWithoutEmbeddedStruct.IdBlock
		varFcpoolFcBlock.Pool = varFcpoolFcBlockWithoutEmbeddedStruct.Pool
		varFcpoolFcBlock.Reservations = varFcpoolFcBlockWithoutEmbeddedStruct.Reservations
		*o = FcpoolFcBlock(varFcpoolFcBlock)
	} else {
		return err
	}

	varFcpoolFcBlock := _FcpoolFcBlock{}

	err = json.Unmarshal(bytes, &varFcpoolFcBlock)
	if err == nil {
		o.PoolAbstractBlock = varFcpoolFcBlock.PoolAbstractBlock
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdBlock")
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

type NullableFcpoolFcBlock struct {
	value *FcpoolFcBlock
	isSet bool
}

func (v NullableFcpoolFcBlock) Get() *FcpoolFcBlock {
	return v.value
}

func (v *NullableFcpoolFcBlock) Set(val *FcpoolFcBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpoolFcBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpoolFcBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpoolFcBlock(val *FcpoolFcBlock) *NullableFcpoolFcBlock {
	return &NullableFcpoolFcBlock{value: val, isSet: true}
}

func (v NullableFcpoolFcBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFcpoolFcBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
