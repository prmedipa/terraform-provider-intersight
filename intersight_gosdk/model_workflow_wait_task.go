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
	"reflect"
	"strings"
)

// checks if the WorkflowWaitTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowWaitTask{}

// WorkflowWaitTask A WaitTask will remain in progress until marked success or failed by an external trigger. The timeout for wait task is 180 days, so a workflow can wait for task status update for upto 180 days. Currently the only supported means to update the task status is through an API that provides the status of the task runtime instance. Once the wait task status has been set the workflow will continue with the execution based on the task status.
type WorkflowWaitTask struct {
	WorkflowAbstractWorkerTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	InputDefinition      []WorkflowBaseDataType   `json:"InputDefinition,omitempty"`
	OutputDefinition     []WorkflowBaseDataType   `json:"OutputDefinition,omitempty"`
	Prompts              []WorkflowWaitTaskPrompt `json:"Prompts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWaitTask WorkflowWaitTask

// NewWorkflowWaitTask instantiates a new WorkflowWaitTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitTask(classId string, objectType string) *WorkflowWaitTask {
	this := WorkflowWaitTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	var rollbackDisabled bool = false
	this.RollbackDisabled = &rollbackDisabled
	return &this
}

// NewWorkflowWaitTaskWithDefaults instantiates a new WorkflowWaitTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitTaskWithDefaults() *WorkflowWaitTask {
	this := WorkflowWaitTask{}
	var classId string = "workflow.WaitTask"
	this.ClassId = classId
	var objectType string = "workflow.WaitTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWaitTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWaitTask) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.WaitTask" of the ClassId field.
func (o *WorkflowWaitTask) GetDefaultClassId() interface{} {
	return "workflow.WaitTask"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWaitTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWaitTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.WaitTask" of the ObjectType field.
func (o *WorkflowWaitTask) GetDefaultObjectType() interface{} {
	return "workflow.WaitTask"
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWaitTask) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWaitTask) GetInputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || IsNil(o.InputDefinition) {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowWaitTask) HasInputDefinition() bool {
	if o != nil && !IsNil(o.InputDefinition) {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowWaitTask) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWaitTask) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWaitTask) GetOutputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || IsNil(o.OutputDefinition) {
		return nil, false
	}
	return o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowWaitTask) HasOutputDefinition() bool {
	if o != nil && !IsNil(o.OutputDefinition) {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowWaitTask) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = v
}

// GetPrompts returns the Prompts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWaitTask) GetPrompts() []WorkflowWaitTaskPrompt {
	if o == nil {
		var ret []WorkflowWaitTaskPrompt
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWaitTask) GetPromptsOk() ([]WorkflowWaitTaskPrompt, bool) {
	if o == nil || IsNil(o.Prompts) {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *WorkflowWaitTask) HasPrompts() bool {
	if o != nil && !IsNil(o.Prompts) {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []WorkflowWaitTaskPrompt and assigns it to the Prompts field.
func (o *WorkflowWaitTask) SetPrompts(v []WorkflowWaitTaskPrompt) {
	o.Prompts = v
}

func (o WorkflowWaitTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowWaitTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowAbstractWorkerTask, errWorkflowAbstractWorkerTask := json.Marshal(o.WorkflowAbstractWorkerTask)
	if errWorkflowAbstractWorkerTask != nil {
		return map[string]interface{}{}, errWorkflowAbstractWorkerTask
	}
	errWorkflowAbstractWorkerTask = json.Unmarshal([]byte(serializedWorkflowAbstractWorkerTask), &toSerialize)
	if errWorkflowAbstractWorkerTask != nil {
		return map[string]interface{}{}, errWorkflowAbstractWorkerTask
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.Prompts != nil {
		toSerialize["Prompts"] = o.Prompts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowWaitTask) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowWaitTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                   `json:"ObjectType"`
		InputDefinition  []WorkflowBaseDataType   `json:"InputDefinition,omitempty"`
		OutputDefinition []WorkflowBaseDataType   `json:"OutputDefinition,omitempty"`
		Prompts          []WorkflowWaitTaskPrompt `json:"Prompts,omitempty"`
	}

	varWorkflowWaitTaskWithoutEmbeddedStruct := WorkflowWaitTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowWaitTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWaitTask := _WorkflowWaitTask{}
		varWorkflowWaitTask.ClassId = varWorkflowWaitTaskWithoutEmbeddedStruct.ClassId
		varWorkflowWaitTask.ObjectType = varWorkflowWaitTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowWaitTask.InputDefinition = varWorkflowWaitTaskWithoutEmbeddedStruct.InputDefinition
		varWorkflowWaitTask.OutputDefinition = varWorkflowWaitTaskWithoutEmbeddedStruct.OutputDefinition
		varWorkflowWaitTask.Prompts = varWorkflowWaitTaskWithoutEmbeddedStruct.Prompts
		*o = WorkflowWaitTask(varWorkflowWaitTask)
	} else {
		return err
	}

	varWorkflowWaitTask := _WorkflowWaitTask{}

	err = json.Unmarshal(data, &varWorkflowWaitTask)
	if err == nil {
		o.WorkflowAbstractWorkerTask = varWorkflowWaitTask.WorkflowAbstractWorkerTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "Prompts")

		// remove fields from embedded structs
		reflectWorkflowAbstractWorkerTask := reflect.ValueOf(o.WorkflowAbstractWorkerTask)
		for i := 0; i < reflectWorkflowAbstractWorkerTask.Type().NumField(); i++ {
			t := reflectWorkflowAbstractWorkerTask.Type().Field(i)

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

type NullableWorkflowWaitTask struct {
	value *WorkflowWaitTask
	isSet bool
}

func (v NullableWorkflowWaitTask) Get() *WorkflowWaitTask {
	return v.value
}

func (v *NullableWorkflowWaitTask) Set(val *WorkflowWaitTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitTask(val *WorkflowWaitTask) *NullableWorkflowWaitTask {
	return &NullableWorkflowWaitTask{value: val, isSet: true}
}

func (v NullableWorkflowWaitTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
