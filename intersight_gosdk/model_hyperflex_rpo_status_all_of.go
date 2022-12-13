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
)

// HyperflexRpoStatusAllOf Definition of the list of properties defined in 'hyperflex.RpoStatus', excluding properties defined in parent classes.
type HyperflexRpoStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Actual end time for the snapshot.
	Actual *int64 `json:"Actual,omitempty"`
	// Expected end time for the snapshot.
	Expected *int64 `json:"Expected,omitempty"`
	// A flag to determine if snapshot delivery is delayed.
	RpoExceeded          *bool `json:"RpoExceeded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexRpoStatusAllOf HyperflexRpoStatusAllOf

// NewHyperflexRpoStatusAllOf instantiates a new HyperflexRpoStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexRpoStatusAllOf(classId string, objectType string) *HyperflexRpoStatusAllOf {
	this := HyperflexRpoStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexRpoStatusAllOfWithDefaults instantiates a new HyperflexRpoStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexRpoStatusAllOfWithDefaults() *HyperflexRpoStatusAllOf {
	this := HyperflexRpoStatusAllOf{}
	var classId string = "hyperflex.RpoStatus"
	this.ClassId = classId
	var objectType string = "hyperflex.RpoStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexRpoStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexRpoStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexRpoStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexRpoStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *HyperflexRpoStatusAllOf) GetActual() int64 {
	if o == nil || o.Actual == nil {
		var ret int64
		return ret
	}
	return *o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatusAllOf) GetActualOk() (*int64, bool) {
	if o == nil || o.Actual == nil {
		return nil, false
	}
	return o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *HyperflexRpoStatusAllOf) HasActual() bool {
	if o != nil && o.Actual != nil {
		return true
	}

	return false
}

// SetActual gets a reference to the given int64 and assigns it to the Actual field.
func (o *HyperflexRpoStatusAllOf) SetActual(v int64) {
	o.Actual = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *HyperflexRpoStatusAllOf) GetExpected() int64 {
	if o == nil || o.Expected == nil {
		var ret int64
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatusAllOf) GetExpectedOk() (*int64, bool) {
	if o == nil || o.Expected == nil {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *HyperflexRpoStatusAllOf) HasExpected() bool {
	if o != nil && o.Expected != nil {
		return true
	}

	return false
}

// SetExpected gets a reference to the given int64 and assigns it to the Expected field.
func (o *HyperflexRpoStatusAllOf) SetExpected(v int64) {
	o.Expected = &v
}

// GetRpoExceeded returns the RpoExceeded field value if set, zero value otherwise.
func (o *HyperflexRpoStatusAllOf) GetRpoExceeded() bool {
	if o == nil || o.RpoExceeded == nil {
		var ret bool
		return ret
	}
	return *o.RpoExceeded
}

// GetRpoExceededOk returns a tuple with the RpoExceeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexRpoStatusAllOf) GetRpoExceededOk() (*bool, bool) {
	if o == nil || o.RpoExceeded == nil {
		return nil, false
	}
	return o.RpoExceeded, true
}

// HasRpoExceeded returns a boolean if a field has been set.
func (o *HyperflexRpoStatusAllOf) HasRpoExceeded() bool {
	if o != nil && o.RpoExceeded != nil {
		return true
	}

	return false
}

// SetRpoExceeded gets a reference to the given bool and assigns it to the RpoExceeded field.
func (o *HyperflexRpoStatusAllOf) SetRpoExceeded(v bool) {
	o.RpoExceeded = &v
}

func (o HyperflexRpoStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Actual != nil {
		toSerialize["Actual"] = o.Actual
	}
	if o.Expected != nil {
		toSerialize["Expected"] = o.Expected
	}
	if o.RpoExceeded != nil {
		toSerialize["RpoExceeded"] = o.RpoExceeded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexRpoStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexRpoStatusAllOf := _HyperflexRpoStatusAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexRpoStatusAllOf); err == nil {
		*o = HyperflexRpoStatusAllOf(varHyperflexRpoStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actual")
		delete(additionalProperties, "Expected")
		delete(additionalProperties, "RpoExceeded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexRpoStatusAllOf struct {
	value *HyperflexRpoStatusAllOf
	isSet bool
}

func (v NullableHyperflexRpoStatusAllOf) Get() *HyperflexRpoStatusAllOf {
	return v.value
}

func (v *NullableHyperflexRpoStatusAllOf) Set(val *HyperflexRpoStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexRpoStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexRpoStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexRpoStatusAllOf(val *HyperflexRpoStatusAllOf) *NullableHyperflexRpoStatusAllOf {
	return &NullableHyperflexRpoStatusAllOf{value: val, isSet: true}
}

func (v NullableHyperflexRpoStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexRpoStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
