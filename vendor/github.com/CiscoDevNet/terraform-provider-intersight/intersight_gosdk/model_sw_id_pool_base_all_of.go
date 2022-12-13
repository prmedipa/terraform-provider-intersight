/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SwIdPoolBaseAllOf Definition of the list of properties defined in 'sw.IdPoolBase', excluding properties defined in parent classes.
type SwIdPoolBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType      string  `json:"ObjectType"`
	GapAvailableIds []int64 `json:"GapAvailableIds,omitempty"`
	// Shows the next available Chassis ID to be allocated.
	NextAvailableId      *int64 `json:"NextAvailableId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwIdPoolBaseAllOf SwIdPoolBaseAllOf

// NewSwIdPoolBaseAllOf instantiates a new SwIdPoolBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwIdPoolBaseAllOf(classId string, objectType string) *SwIdPoolBaseAllOf {
	this := SwIdPoolBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSwIdPoolBaseAllOfWithDefaults instantiates a new SwIdPoolBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwIdPoolBaseAllOfWithDefaults() *SwIdPoolBaseAllOf {
	this := SwIdPoolBaseAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *SwIdPoolBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SwIdPoolBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SwIdPoolBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SwIdPoolBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SwIdPoolBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SwIdPoolBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGapAvailableIds returns the GapAvailableIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SwIdPoolBaseAllOf) GetGapAvailableIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.GapAvailableIds
}

// GetGapAvailableIdsOk returns a tuple with the GapAvailableIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SwIdPoolBaseAllOf) GetGapAvailableIdsOk() ([]int64, bool) {
	if o == nil || o.GapAvailableIds == nil {
		return nil, false
	}
	return o.GapAvailableIds, true
}

// HasGapAvailableIds returns a boolean if a field has been set.
func (o *SwIdPoolBaseAllOf) HasGapAvailableIds() bool {
	if o != nil && o.GapAvailableIds != nil {
		return true
	}

	return false
}

// SetGapAvailableIds gets a reference to the given []int64 and assigns it to the GapAvailableIds field.
func (o *SwIdPoolBaseAllOf) SetGapAvailableIds(v []int64) {
	o.GapAvailableIds = v
}

// GetNextAvailableId returns the NextAvailableId field value if set, zero value otherwise.
func (o *SwIdPoolBaseAllOf) GetNextAvailableId() int64 {
	if o == nil || o.NextAvailableId == nil {
		var ret int64
		return ret
	}
	return *o.NextAvailableId
}

// GetNextAvailableIdOk returns a tuple with the NextAvailableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwIdPoolBaseAllOf) GetNextAvailableIdOk() (*int64, bool) {
	if o == nil || o.NextAvailableId == nil {
		return nil, false
	}
	return o.NextAvailableId, true
}

// HasNextAvailableId returns a boolean if a field has been set.
func (o *SwIdPoolBaseAllOf) HasNextAvailableId() bool {
	if o != nil && o.NextAvailableId != nil {
		return true
	}

	return false
}

// SetNextAvailableId gets a reference to the given int64 and assigns it to the NextAvailableId field.
func (o *SwIdPoolBaseAllOf) SetNextAvailableId(v int64) {
	o.NextAvailableId = &v
}

func (o SwIdPoolBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.GapAvailableIds != nil {
		toSerialize["GapAvailableIds"] = o.GapAvailableIds
	}
	if o.NextAvailableId != nil {
		toSerialize["NextAvailableId"] = o.NextAvailableId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SwIdPoolBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSwIdPoolBaseAllOf := _SwIdPoolBaseAllOf{}

	if err = json.Unmarshal(bytes, &varSwIdPoolBaseAllOf); err == nil {
		*o = SwIdPoolBaseAllOf(varSwIdPoolBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "GapAvailableIds")
		delete(additionalProperties, "NextAvailableId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwIdPoolBaseAllOf struct {
	value *SwIdPoolBaseAllOf
	isSet bool
}

func (v NullableSwIdPoolBaseAllOf) Get() *SwIdPoolBaseAllOf {
	return v.value
}

func (v *NullableSwIdPoolBaseAllOf) Set(val *SwIdPoolBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSwIdPoolBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSwIdPoolBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwIdPoolBaseAllOf(val *SwIdPoolBaseAllOf) *NullableSwIdPoolBaseAllOf {
	return &NullableSwIdPoolBaseAllOf{value: val, isSet: true}
}

func (v NullableSwIdPoolBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwIdPoolBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
