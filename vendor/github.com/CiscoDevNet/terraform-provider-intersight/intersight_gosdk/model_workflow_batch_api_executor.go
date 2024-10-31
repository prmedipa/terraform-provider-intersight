/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the WorkflowBatchApiExecutor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowBatchApiExecutor{}

// WorkflowBatchApiExecutor Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Batch API associates the list of API requests to be executed as part of single task execution. Each API request takes the request body and a response parser specification.
type WorkflowBatchApiExecutor struct {
	WorkflowBatchExecutor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                           `json:"ObjectType"`
	ErrorResponseHandler NullableWorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
	TaskDefinition       NullableWorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBatchApiExecutor WorkflowBatchApiExecutor

// NewWorkflowBatchApiExecutor instantiates a new WorkflowBatchApiExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchApiExecutor(classId string, objectType string) *WorkflowBatchApiExecutor {
	this := WorkflowBatchApiExecutor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBatchApiExecutorWithDefaults instantiates a new WorkflowBatchApiExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchApiExecutorWithDefaults() *WorkflowBatchApiExecutor {
	this := WorkflowBatchApiExecutor{}
	var classId string = "workflow.BatchApiExecutor"
	this.ClassId = classId
	var objectType string = "workflow.BatchApiExecutor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBatchApiExecutor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBatchApiExecutor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.BatchApiExecutor" of the ClassId field.
func (o *WorkflowBatchApiExecutor) GetDefaultClassId() interface{} {
	return "workflow.BatchApiExecutor"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBatchApiExecutor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBatchApiExecutor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.BatchApiExecutor" of the ObjectType field.
func (o *WorkflowBatchApiExecutor) GetDefaultObjectType() interface{} {
	return "workflow.BatchApiExecutor"
}

// GetErrorResponseHandler returns the ErrorResponseHandler field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship {
	if o == nil || IsNil(o.ErrorResponseHandler.Get()) {
		var ret WorkflowErrorResponseHandlerRelationship
		return ret
	}
	return *o.ErrorResponseHandler.Get()
}

// GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorResponseHandler.Get(), o.ErrorResponseHandler.IsSet()
}

// HasErrorResponseHandler returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasErrorResponseHandler() bool {
	if o != nil && o.ErrorResponseHandler.IsSet() {
		return true
	}

	return false
}

// SetErrorResponseHandler gets a reference to the given NullableWorkflowErrorResponseHandlerRelationship and assigns it to the ErrorResponseHandler field.
func (o *WorkflowBatchApiExecutor) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship) {
	o.ErrorResponseHandler.Set(&v)
}

// SetErrorResponseHandlerNil sets the value for ErrorResponseHandler to be an explicit nil
func (o *WorkflowBatchApiExecutor) SetErrorResponseHandlerNil() {
	o.ErrorResponseHandler.Set(nil)
}

// UnsetErrorResponseHandler ensures that no value is present for ErrorResponseHandler, not even an explicit nil
func (o *WorkflowBatchApiExecutor) UnsetErrorResponseHandler() {
	o.ErrorResponseHandler.Unset()
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || IsNil(o.TaskDefinition.Get()) {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition.Get()
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskDefinition.Get(), o.TaskDefinition.IsSet()
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition.IsSet() {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given NullableWorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowBatchApiExecutor) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition.Set(&v)
}

// SetTaskDefinitionNil sets the value for TaskDefinition to be an explicit nil
func (o *WorkflowBatchApiExecutor) SetTaskDefinitionNil() {
	o.TaskDefinition.Set(nil)
}

// UnsetTaskDefinition ensures that no value is present for TaskDefinition, not even an explicit nil
func (o *WorkflowBatchApiExecutor) UnsetTaskDefinition() {
	o.TaskDefinition.Unset()
}

func (o WorkflowBatchApiExecutor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowBatchApiExecutor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBatchExecutor, errWorkflowBatchExecutor := json.Marshal(o.WorkflowBatchExecutor)
	if errWorkflowBatchExecutor != nil {
		return map[string]interface{}{}, errWorkflowBatchExecutor
	}
	errWorkflowBatchExecutor = json.Unmarshal([]byte(serializedWorkflowBatchExecutor), &toSerialize)
	if errWorkflowBatchExecutor != nil {
		return map[string]interface{}{}, errWorkflowBatchExecutor
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ErrorResponseHandler.IsSet() {
		toSerialize["ErrorResponseHandler"] = o.ErrorResponseHandler.Get()
	}
	if o.TaskDefinition.IsSet() {
		toSerialize["TaskDefinition"] = o.TaskDefinition.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowBatchApiExecutor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
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
	type WorkflowBatchApiExecutorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                                           `json:"ObjectType"`
		ErrorResponseHandler NullableWorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
		TaskDefinition       NullableWorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	}

	varWorkflowBatchApiExecutorWithoutEmbeddedStruct := WorkflowBatchApiExecutorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowBatchApiExecutorWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowBatchApiExecutor := _WorkflowBatchApiExecutor{}
		varWorkflowBatchApiExecutor.ClassId = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ClassId
		varWorkflowBatchApiExecutor.ObjectType = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ObjectType
		varWorkflowBatchApiExecutor.ErrorResponseHandler = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ErrorResponseHandler
		varWorkflowBatchApiExecutor.TaskDefinition = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.TaskDefinition
		*o = WorkflowBatchApiExecutor(varWorkflowBatchApiExecutor)
	} else {
		return err
	}

	varWorkflowBatchApiExecutor := _WorkflowBatchApiExecutor{}

	err = json.Unmarshal(data, &varWorkflowBatchApiExecutor)
	if err == nil {
		o.WorkflowBatchExecutor = varWorkflowBatchApiExecutor.WorkflowBatchExecutor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorResponseHandler")
		delete(additionalProperties, "TaskDefinition")

		// remove fields from embedded structs
		reflectWorkflowBatchExecutor := reflect.ValueOf(o.WorkflowBatchExecutor)
		for i := 0; i < reflectWorkflowBatchExecutor.Type().NumField(); i++ {
			t := reflectWorkflowBatchExecutor.Type().Field(i)

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

type NullableWorkflowBatchApiExecutor struct {
	value *WorkflowBatchApiExecutor
	isSet bool
}

func (v NullableWorkflowBatchApiExecutor) Get() *WorkflowBatchApiExecutor {
	return v.value
}

func (v *NullableWorkflowBatchApiExecutor) Set(val *WorkflowBatchApiExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchApiExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchApiExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchApiExecutor(val *WorkflowBatchApiExecutor) *NullableWorkflowBatchApiExecutor {
	return &NullableWorkflowBatchApiExecutor{value: val, isSet: true}
}

func (v NullableWorkflowBatchApiExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchApiExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
