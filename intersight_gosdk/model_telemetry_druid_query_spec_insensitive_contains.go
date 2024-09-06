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
)

// checks if the TelemetryDruidQuerySpecInsensitiveContains type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelemetryDruidQuerySpecInsensitiveContains{}

// TelemetryDruidQuerySpecInsensitiveContains A 'insensitive_contains' query spec. Note that an \"insensitive_contains\" search is equivalent to a \"contains\" search with \"caseSensitive\" false (or not provided).
type TelemetryDruidQuerySpecInsensitiveContains struct {
	Type string `json:"type"`
	// A String value to search.
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidQuerySpecInsensitiveContains TelemetryDruidQuerySpecInsensitiveContains

// NewTelemetryDruidQuerySpecInsensitiveContains instantiates a new TelemetryDruidQuerySpecInsensitiveContains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidQuerySpecInsensitiveContains(type_ string, value string) *TelemetryDruidQuerySpecInsensitiveContains {
	this := TelemetryDruidQuerySpecInsensitiveContains{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewTelemetryDruidQuerySpecInsensitiveContainsWithDefaults instantiates a new TelemetryDruidQuerySpecInsensitiveContains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidQuerySpecInsensitiveContainsWithDefaults() *TelemetryDruidQuerySpecInsensitiveContains {
	this := TelemetryDruidQuerySpecInsensitiveContains{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidQuerySpecInsensitiveContains) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQuerySpecInsensitiveContains) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidQuerySpecInsensitiveContains) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidQuerySpecInsensitiveContains) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQuerySpecInsensitiveContains) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidQuerySpecInsensitiveContains) SetValue(v string) {
	o.Value = v
}

func (o TelemetryDruidQuerySpecInsensitiveContains) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelemetryDruidQuerySpecInsensitiveContains) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelemetryDruidQuerySpecInsensitiveContains) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	varTelemetryDruidQuerySpecInsensitiveContains := _TelemetryDruidQuerySpecInsensitiveContains{}

	err = json.Unmarshal(data, &varTelemetryDruidQuerySpecInsensitiveContains)

	if err != nil {
		return err
	}

	*o = TelemetryDruidQuerySpecInsensitiveContains(varTelemetryDruidQuerySpecInsensitiveContains)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidQuerySpecInsensitiveContains struct {
	value *TelemetryDruidQuerySpecInsensitiveContains
	isSet bool
}

func (v NullableTelemetryDruidQuerySpecInsensitiveContains) Get() *TelemetryDruidQuerySpecInsensitiveContains {
	return v.value
}

func (v *NullableTelemetryDruidQuerySpecInsensitiveContains) Set(val *TelemetryDruidQuerySpecInsensitiveContains) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidQuerySpecInsensitiveContains) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidQuerySpecInsensitiveContains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidQuerySpecInsensitiveContains(val *TelemetryDruidQuerySpecInsensitiveContains) *NullableTelemetryDruidQuerySpecInsensitiveContains {
	return &NullableTelemetryDruidQuerySpecInsensitiveContains{value: val, isSet: true}
}

func (v NullableTelemetryDruidQuerySpecInsensitiveContains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidQuerySpecInsensitiveContains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
