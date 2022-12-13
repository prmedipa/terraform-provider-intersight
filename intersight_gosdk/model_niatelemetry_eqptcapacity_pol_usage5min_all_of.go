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

// NiatelemetryEqptcapacityPolUsage5minAllOf Definition of the list of properties defined in 'niatelemetry.EqptcapacityPolUsage5min', excluding properties defined in parent classes.
type NiatelemetryEqptcapacityPolUsage5minAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// PolUsageBase information for aci nodes.
	PolUsageBase *string `json:"PolUsageBase,omitempty"`
	// PolUsageCapCum information for aci nodes.
	PolUsageCapCum *string `json:"PolUsageCapCum,omitempty"`
	// PolUsageCum information for aci nodes.
	PolUsageCum          *string `json:"PolUsageCum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryEqptcapacityPolUsage5minAllOf NiatelemetryEqptcapacityPolUsage5minAllOf

// NewNiatelemetryEqptcapacityPolUsage5minAllOf instantiates a new NiatelemetryEqptcapacityPolUsage5minAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryEqptcapacityPolUsage5minAllOf(classId string, objectType string) *NiatelemetryEqptcapacityPolUsage5minAllOf {
	this := NiatelemetryEqptcapacityPolUsage5minAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryEqptcapacityPolUsage5minAllOfWithDefaults instantiates a new NiatelemetryEqptcapacityPolUsage5minAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryEqptcapacityPolUsage5minAllOfWithDefaults() *NiatelemetryEqptcapacityPolUsage5minAllOf {
	this := NiatelemetryEqptcapacityPolUsage5minAllOf{}
	var classId string = "niatelemetry.EqptcapacityPolUsage5min"
	this.ClassId = classId
	var objectType string = "niatelemetry.EqptcapacityPolUsage5min"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPolUsageBase returns the PolUsageBase field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageBase() string {
	if o == nil || o.PolUsageBase == nil {
		var ret string
		return ret
	}
	return *o.PolUsageBase
}

// GetPolUsageBaseOk returns a tuple with the PolUsageBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageBaseOk() (*string, bool) {
	if o == nil || o.PolUsageBase == nil {
		return nil, false
	}
	return o.PolUsageBase, true
}

// HasPolUsageBase returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) HasPolUsageBase() bool {
	if o != nil && o.PolUsageBase != nil {
		return true
	}

	return false
}

// SetPolUsageBase gets a reference to the given string and assigns it to the PolUsageBase field.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) SetPolUsageBase(v string) {
	o.PolUsageBase = &v
}

// GetPolUsageCapCum returns the PolUsageCapCum field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageCapCum() string {
	if o == nil || o.PolUsageCapCum == nil {
		var ret string
		return ret
	}
	return *o.PolUsageCapCum
}

// GetPolUsageCapCumOk returns a tuple with the PolUsageCapCum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageCapCumOk() (*string, bool) {
	if o == nil || o.PolUsageCapCum == nil {
		return nil, false
	}
	return o.PolUsageCapCum, true
}

// HasPolUsageCapCum returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) HasPolUsageCapCum() bool {
	if o != nil && o.PolUsageCapCum != nil {
		return true
	}

	return false
}

// SetPolUsageCapCum gets a reference to the given string and assigns it to the PolUsageCapCum field.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) SetPolUsageCapCum(v string) {
	o.PolUsageCapCum = &v
}

// GetPolUsageCum returns the PolUsageCum field value if set, zero value otherwise.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageCum() string {
	if o == nil || o.PolUsageCum == nil {
		var ret string
		return ret
	}
	return *o.PolUsageCum
}

// GetPolUsageCumOk returns a tuple with the PolUsageCum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) GetPolUsageCumOk() (*string, bool) {
	if o == nil || o.PolUsageCum == nil {
		return nil, false
	}
	return o.PolUsageCum, true
}

// HasPolUsageCum returns a boolean if a field has been set.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) HasPolUsageCum() bool {
	if o != nil && o.PolUsageCum != nil {
		return true
	}

	return false
}

// SetPolUsageCum gets a reference to the given string and assigns it to the PolUsageCum field.
func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) SetPolUsageCum(v string) {
	o.PolUsageCum = &v
}

func (o NiatelemetryEqptcapacityPolUsage5minAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PolUsageBase != nil {
		toSerialize["PolUsageBase"] = o.PolUsageBase
	}
	if o.PolUsageCapCum != nil {
		toSerialize["PolUsageCapCum"] = o.PolUsageCapCum
	}
	if o.PolUsageCum != nil {
		toSerialize["PolUsageCum"] = o.PolUsageCum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryEqptcapacityPolUsage5minAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryEqptcapacityPolUsage5minAllOf := _NiatelemetryEqptcapacityPolUsage5minAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryEqptcapacityPolUsage5minAllOf); err == nil {
		*o = NiatelemetryEqptcapacityPolUsage5minAllOf(varNiatelemetryEqptcapacityPolUsage5minAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PolUsageBase")
		delete(additionalProperties, "PolUsageCapCum")
		delete(additionalProperties, "PolUsageCum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryEqptcapacityPolUsage5minAllOf struct {
	value *NiatelemetryEqptcapacityPolUsage5minAllOf
	isSet bool
}

func (v NullableNiatelemetryEqptcapacityPolUsage5minAllOf) Get() *NiatelemetryEqptcapacityPolUsage5minAllOf {
	return v.value
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5minAllOf) Set(val *NiatelemetryEqptcapacityPolUsage5minAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryEqptcapacityPolUsage5minAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5minAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryEqptcapacityPolUsage5minAllOf(val *NiatelemetryEqptcapacityPolUsage5minAllOf) *NullableNiatelemetryEqptcapacityPolUsage5minAllOf {
	return &NullableNiatelemetryEqptcapacityPolUsage5minAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryEqptcapacityPolUsage5minAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryEqptcapacityPolUsage5minAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
