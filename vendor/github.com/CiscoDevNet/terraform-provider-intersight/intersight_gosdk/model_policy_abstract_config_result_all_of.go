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

// PolicyAbstractConfigResultAllOf Definition of the list of properties defined in 'policy.AbstractConfigResult', excluding properties defined in parent classes.
type PolicyAbstractConfigResultAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The current running stage of the configuration or workflow.
	ConfigStage *string `json:"ConfigStage,omitempty"`
	// Indicates overall configuration state for applying the configuration to the end point. Values  -- Ok, Ok-with-warning, Errored.
	ConfigState *string `json:"ConfigState,omitempty"`
	// Indicates overall state for logical model validation. Values  -- Ok, Ok-with-warning, Errored.
	ValidationState      *string `json:"ValidationState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractConfigResultAllOf PolicyAbstractConfigResultAllOf

// NewPolicyAbstractConfigResultAllOf instantiates a new PolicyAbstractConfigResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractConfigResultAllOf(classId string, objectType string) *PolicyAbstractConfigResultAllOf {
	this := PolicyAbstractConfigResultAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyAbstractConfigResultAllOfWithDefaults instantiates a new PolicyAbstractConfigResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractConfigResultAllOfWithDefaults() *PolicyAbstractConfigResultAllOf {
	this := PolicyAbstractConfigResultAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractConfigResultAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractConfigResultAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractConfigResultAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractConfigResultAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConfigStage returns the ConfigStage field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultAllOf) GetConfigStage() string {
	if o == nil || o.ConfigStage == nil {
		var ret string
		return ret
	}
	return *o.ConfigStage
}

// GetConfigStageOk returns a tuple with the ConfigStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultAllOf) GetConfigStageOk() (*string, bool) {
	if o == nil || o.ConfigStage == nil {
		return nil, false
	}
	return o.ConfigStage, true
}

// HasConfigStage returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultAllOf) HasConfigStage() bool {
	if o != nil && o.ConfigStage != nil {
		return true
	}

	return false
}

// SetConfigStage gets a reference to the given string and assigns it to the ConfigStage field.
func (o *PolicyAbstractConfigResultAllOf) SetConfigStage(v string) {
	o.ConfigStage = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *PolicyAbstractConfigResultAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetValidationState returns the ValidationState field value if set, zero value otherwise.
func (o *PolicyAbstractConfigResultAllOf) GetValidationState() string {
	if o == nil || o.ValidationState == nil {
		var ret string
		return ret
	}
	return *o.ValidationState
}

// GetValidationStateOk returns a tuple with the ValidationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractConfigResultAllOf) GetValidationStateOk() (*string, bool) {
	if o == nil || o.ValidationState == nil {
		return nil, false
	}
	return o.ValidationState, true
}

// HasValidationState returns a boolean if a field has been set.
func (o *PolicyAbstractConfigResultAllOf) HasValidationState() bool {
	if o != nil && o.ValidationState != nil {
		return true
	}

	return false
}

// SetValidationState gets a reference to the given string and assigns it to the ValidationState field.
func (o *PolicyAbstractConfigResultAllOf) SetValidationState(v string) {
	o.ValidationState = &v
}

func (o PolicyAbstractConfigResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConfigStage != nil {
		toSerialize["ConfigStage"] = o.ConfigStage
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.ValidationState != nil {
		toSerialize["ValidationState"] = o.ValidationState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyAbstractConfigResultAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyAbstractConfigResultAllOf := _PolicyAbstractConfigResultAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyAbstractConfigResultAllOf); err == nil {
		*o = PolicyAbstractConfigResultAllOf(varPolicyAbstractConfigResultAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigStage")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "ValidationState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAbstractConfigResultAllOf struct {
	value *PolicyAbstractConfigResultAllOf
	isSet bool
}

func (v NullablePolicyAbstractConfigResultAllOf) Get() *PolicyAbstractConfigResultAllOf {
	return v.value
}

func (v *NullablePolicyAbstractConfigResultAllOf) Set(val *PolicyAbstractConfigResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractConfigResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractConfigResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractConfigResultAllOf(val *PolicyAbstractConfigResultAllOf) *NullablePolicyAbstractConfigResultAllOf {
	return &NullablePolicyAbstractConfigResultAllOf{value: val, isSet: true}
}

func (v NullablePolicyAbstractConfigResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractConfigResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
