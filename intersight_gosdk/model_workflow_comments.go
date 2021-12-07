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
	"reflect"
	"strings"
)

// WorkflowComments Contains detailed information about template function.
type WorkflowComments struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description field provides comment about the template function.
	Description          *string  `json:"Description,omitempty"`
	Examples             []string `json:"Examples,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowComments WorkflowComments

// NewWorkflowComments instantiates a new WorkflowComments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowComments(classId string, objectType string) *WorkflowComments {
	this := WorkflowComments{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowCommentsWithDefaults instantiates a new WorkflowComments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowCommentsWithDefaults() *WorkflowComments {
	this := WorkflowComments{}
	var classId string = "workflow.Comments"
	this.ClassId = classId
	var objectType string = "workflow.Comments"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowComments) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowComments) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowComments) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowComments) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowComments) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowComments) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowComments) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowComments) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowComments) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowComments) SetDescription(v string) {
	o.Description = &v
}

// GetExamples returns the Examples field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowComments) GetExamples() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowComments) GetExamplesOk() (*[]string, bool) {
	if o == nil || o.Examples == nil {
		return nil, false
	}
	return &o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *WorkflowComments) HasExamples() bool {
	if o != nil && o.Examples != nil {
		return true
	}

	return false
}

// SetExamples gets a reference to the given []string and assigns it to the Examples field.
func (o *WorkflowComments) SetExamples(v []string) {
	o.Examples = v
}

func (o WorkflowComments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Examples != nil {
		toSerialize["Examples"] = o.Examples
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowComments) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowCommentsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description field provides comment about the template function.
		Description *string  `json:"Description,omitempty"`
		Examples    []string `json:"Examples,omitempty"`
	}

	varWorkflowCommentsWithoutEmbeddedStruct := WorkflowCommentsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowCommentsWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowComments := _WorkflowComments{}
		varWorkflowComments.ClassId = varWorkflowCommentsWithoutEmbeddedStruct.ClassId
		varWorkflowComments.ObjectType = varWorkflowCommentsWithoutEmbeddedStruct.ObjectType
		varWorkflowComments.Description = varWorkflowCommentsWithoutEmbeddedStruct.Description
		varWorkflowComments.Examples = varWorkflowCommentsWithoutEmbeddedStruct.Examples
		*o = WorkflowComments(varWorkflowComments)
	} else {
		return err
	}

	varWorkflowComments := _WorkflowComments{}

	err = json.Unmarshal(bytes, &varWorkflowComments)
	if err == nil {
		o.MoBaseComplexType = varWorkflowComments.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Examples")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableWorkflowComments struct {
	value *WorkflowComments
	isSet bool
}

func (v NullableWorkflowComments) Get() *WorkflowComments {
	return v.value
}

func (v *NullableWorkflowComments) Set(val *WorkflowComments) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowComments) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowComments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowComments(val *WorkflowComments) *NullableWorkflowComments {
	return &NullableWorkflowComments{value: val, isSet: true}
}

func (v NullableWorkflowComments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowComments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
