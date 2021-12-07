/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// TelemetryDruidAndFilterAllOf struct for TelemetryDruidAndFilterAllOf
type TelemetryDruidAndFilterAllOf struct {
	Fields               []TelemetryDruidFilter `json:"fields"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidAndFilterAllOf TelemetryDruidAndFilterAllOf

// NewTelemetryDruidAndFilterAllOf instantiates a new TelemetryDruidAndFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidAndFilterAllOf(fields []TelemetryDruidFilter) *TelemetryDruidAndFilterAllOf {
	this := TelemetryDruidAndFilterAllOf{}
	this.Fields = fields
	return &this
}

// NewTelemetryDruidAndFilterAllOfWithDefaults instantiates a new TelemetryDruidAndFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidAndFilterAllOfWithDefaults() *TelemetryDruidAndFilterAllOf {
	this := TelemetryDruidAndFilterAllOf{}
	return &this
}

// GetFields returns the Fields field value
func (o *TelemetryDruidAndFilterAllOf) GetFields() []TelemetryDruidFilter {
	if o == nil {
		var ret []TelemetryDruidFilter
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidAndFilterAllOf) GetFieldsOk() (*[]TelemetryDruidFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *TelemetryDruidAndFilterAllOf) SetFields(v []TelemetryDruidFilter) {
	o.Fields = v
}

func (o TelemetryDruidAndFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidAndFilterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidAndFilterAllOf := _TelemetryDruidAndFilterAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidAndFilterAllOf); err == nil {
		*o = TelemetryDruidAndFilterAllOf(varTelemetryDruidAndFilterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidAndFilterAllOf struct {
	value *TelemetryDruidAndFilterAllOf
	isSet bool
}

func (v NullableTelemetryDruidAndFilterAllOf) Get() *TelemetryDruidAndFilterAllOf {
	return v.value
}

func (v *NullableTelemetryDruidAndFilterAllOf) Set(val *TelemetryDruidAndFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidAndFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidAndFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidAndFilterAllOf(val *TelemetryDruidAndFilterAllOf) *NullableTelemetryDruidAndFilterAllOf {
	return &NullableTelemetryDruidAndFilterAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidAndFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidAndFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
