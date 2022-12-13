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

// CondAlarmDefinition The definition of an issue which encompases the criteria for identifying when the issue exists, documentation of the detected issue and the alarms to be raised/cleared by Intersight.
type CondAlarmDefinition struct {
	IssueDefinition
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string            `json:"ObjectType"`
	Actions              []CondAlarmAction `json:"Actions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondAlarmDefinition CondAlarmDefinition

// NewCondAlarmDefinition instantiates a new CondAlarmDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondAlarmDefinition(classId string, objectType string) *CondAlarmDefinition {
	this := CondAlarmDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCondAlarmDefinitionWithDefaults instantiates a new CondAlarmDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondAlarmDefinitionWithDefaults() *CondAlarmDefinition {
	this := CondAlarmDefinition{}
	var classId string = "cond.AlarmDefinition"
	this.ClassId = classId
	var objectType string = "cond.AlarmDefinition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CondAlarmDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CondAlarmDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CondAlarmDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CondAlarmDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CondAlarmDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CondAlarmDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CondAlarmDefinition) GetActions() []CondAlarmAction {
	if o == nil {
		var ret []CondAlarmAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CondAlarmDefinition) GetActionsOk() ([]CondAlarmAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CondAlarmDefinition) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []CondAlarmAction and assigns it to the Actions field.
func (o *CondAlarmDefinition) SetActions(v []CondAlarmAction) {
	o.Actions = v
}

func (o CondAlarmDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedIssueDefinition, errIssueDefinition := json.Marshal(o.IssueDefinition)
	if errIssueDefinition != nil {
		return []byte{}, errIssueDefinition
	}
	errIssueDefinition = json.Unmarshal([]byte(serializedIssueDefinition), &toSerialize)
	if errIssueDefinition != nil {
		return []byte{}, errIssueDefinition
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondAlarmDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type CondAlarmDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string            `json:"ObjectType"`
		Actions    []CondAlarmAction `json:"Actions,omitempty"`
	}

	varCondAlarmDefinitionWithoutEmbeddedStruct := CondAlarmDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCondAlarmDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varCondAlarmDefinition := _CondAlarmDefinition{}
		varCondAlarmDefinition.ClassId = varCondAlarmDefinitionWithoutEmbeddedStruct.ClassId
		varCondAlarmDefinition.ObjectType = varCondAlarmDefinitionWithoutEmbeddedStruct.ObjectType
		varCondAlarmDefinition.Actions = varCondAlarmDefinitionWithoutEmbeddedStruct.Actions
		*o = CondAlarmDefinition(varCondAlarmDefinition)
	} else {
		return err
	}

	varCondAlarmDefinition := _CondAlarmDefinition{}

	err = json.Unmarshal(bytes, &varCondAlarmDefinition)
	if err == nil {
		o.IssueDefinition = varCondAlarmDefinition.IssueDefinition
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")

		// remove fields from embedded structs
		reflectIssueDefinition := reflect.ValueOf(o.IssueDefinition)
		for i := 0; i < reflectIssueDefinition.Type().NumField(); i++ {
			t := reflectIssueDefinition.Type().Field(i)

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

type NullableCondAlarmDefinition struct {
	value *CondAlarmDefinition
	isSet bool
}

func (v NullableCondAlarmDefinition) Get() *CondAlarmDefinition {
	return v.value
}

func (v *NullableCondAlarmDefinition) Set(val *CondAlarmDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCondAlarmDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCondAlarmDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondAlarmDefinition(val *CondAlarmDefinition) *NullableCondAlarmDefinition {
	return &NullableCondAlarmDefinition{value: val, isSet: true}
}

func (v NullableCondAlarmDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondAlarmDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
