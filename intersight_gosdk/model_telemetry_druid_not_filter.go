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

// TelemetryDruidNotFilter A logical NOT expression filter.
type TelemetryDruidNotFilter struct {
	// The filter type.
	Type string `json:"type"`
	// All filters except the \"spatial\" filter support extraction functions. An extraction function is defined by setting the \"extractionFn\" field on a filter. See Extraction function for more details on extraction functions. If specified, the extraction function will be used to transform input values before the filter is applied. The example below shows a selector filter combined with an extraction function. This filter will transform input values according to the values defined in the lookup map; transformed values will then be matched with the string \"bar_1\".
	ExtractionFn         map[string]interface{} `json:"extractionFn,omitempty"`
	Field                TelemetryDruidFilter   `json:"field"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidNotFilter TelemetryDruidNotFilter

// NewTelemetryDruidNotFilter instantiates a new TelemetryDruidNotFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidNotFilter(type_ string, field TelemetryDruidFilter) *TelemetryDruidNotFilter {
	this := TelemetryDruidNotFilter{}
	this.Type = type_
	this.Field = field
	return &this
}

// NewTelemetryDruidNotFilterWithDefaults instantiates a new TelemetryDruidNotFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidNotFilterWithDefaults() *TelemetryDruidNotFilter {
	this := TelemetryDruidNotFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidNotFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNotFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidNotFilter) SetType(v string) {
	o.Type = v
}

// GetExtractionFn returns the ExtractionFn field value if set, zero value otherwise.
func (o *TelemetryDruidNotFilter) GetExtractionFn() map[string]interface{} {
	if o == nil || o.ExtractionFn == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtractionFn
}

// GetExtractionFnOk returns a tuple with the ExtractionFn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNotFilter) GetExtractionFnOk() (map[string]interface{}, bool) {
	if o == nil || o.ExtractionFn == nil {
		return nil, false
	}
	return o.ExtractionFn, true
}

// HasExtractionFn returns a boolean if a field has been set.
func (o *TelemetryDruidNotFilter) HasExtractionFn() bool {
	if o != nil && o.ExtractionFn != nil {
		return true
	}

	return false
}

// SetExtractionFn gets a reference to the given map[string]interface{} and assigns it to the ExtractionFn field.
func (o *TelemetryDruidNotFilter) SetExtractionFn(v map[string]interface{}) {
	o.ExtractionFn = v
}

// GetField returns the Field field value
func (o *TelemetryDruidNotFilter) GetField() TelemetryDruidFilter {
	if o == nil {
		var ret TelemetryDruidFilter
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidNotFilter) GetFieldOk() (*TelemetryDruidFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *TelemetryDruidNotFilter) SetField(v TelemetryDruidFilter) {
	o.Field = v
}

func (o TelemetryDruidNotFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.ExtractionFn != nil {
		toSerialize["extractionFn"] = o.ExtractionFn
	}
	if true {
		toSerialize["field"] = o.Field
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidNotFilter) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidNotFilter := _TelemetryDruidNotFilter{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidNotFilter); err == nil {
		*o = TelemetryDruidNotFilter(varTelemetryDruidNotFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "extractionFn")
		delete(additionalProperties, "field")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidNotFilter struct {
	value *TelemetryDruidNotFilter
	isSet bool
}

func (v NullableTelemetryDruidNotFilter) Get() *TelemetryDruidNotFilter {
	return v.value
}

func (v *NullableTelemetryDruidNotFilter) Set(val *TelemetryDruidNotFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidNotFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidNotFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidNotFilter(val *TelemetryDruidNotFilter) *NullableTelemetryDruidNotFilter {
	return &NullableTelemetryDruidNotFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidNotFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidNotFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
