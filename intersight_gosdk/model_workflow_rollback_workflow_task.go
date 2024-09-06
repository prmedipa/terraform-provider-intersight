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

// checks if the WorkflowRollbackWorkflowTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRollbackWorkflowTask{}

// WorkflowRollbackWorkflowTask Rollback workflow task information.
type WorkflowRollbackWorkflowTask struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the rollback task.
	Description *string `json:"Description,omitempty"`
	// Name of TaskInfo that needs to be rolled back.
	Name *string `json:"Name,omitempty"`
	// Reference name of TaskInfo that need to be rolled back.
	RefName *string `json:"RefName,omitempty"`
	// Status of the rollback operation for the task.
	RollbackCompleted *bool `json:"RollbackCompleted,omitempty"`
	// Name of TaskInfo that performs the rollback operation.
	RollbackTaskName *string `json:"RollbackTaskName,omitempty"`
	// Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed. * `NotStarted` - Status of rollback task when it is not started rollback. * `NotSupported` - Status of task when it is not supporting rollback. * `Completed` - Status of rollback task once execution is successful. * `Failed` - Status of rollback task when it is failed. * `Disabled` - Status of rollback task when rollback is disabled.
	Status *string `json:"Status,omitempty"`
	// Moid of TaskInfo that supports rollback operation.
	TaskInfoMoid *string `json:"TaskInfoMoid,omitempty"`
	// Path of rollback task if it is inside sub-workflow.
	TaskPath             *string `json:"TaskPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowRollbackWorkflowTask WorkflowRollbackWorkflowTask

// NewWorkflowRollbackWorkflowTask instantiates a new WorkflowRollbackWorkflowTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRollbackWorkflowTask(classId string, objectType string) *WorkflowRollbackWorkflowTask {
	this := WorkflowRollbackWorkflowTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowRollbackWorkflowTaskWithDefaults instantiates a new WorkflowRollbackWorkflowTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRollbackWorkflowTaskWithDefaults() *WorkflowRollbackWorkflowTask {
	this := WorkflowRollbackWorkflowTask{}
	var classId string = "workflow.RollbackWorkflowTask"
	this.ClassId = classId
	var objectType string = "workflow.RollbackWorkflowTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowRollbackWorkflowTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowRollbackWorkflowTask) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.RollbackWorkflowTask" of the ClassId field.
func (o *WorkflowRollbackWorkflowTask) GetDefaultClassId() interface{} {
	return "workflow.RollbackWorkflowTask"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowRollbackWorkflowTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowRollbackWorkflowTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.RollbackWorkflowTask" of the ObjectType field.
func (o *WorkflowRollbackWorkflowTask) GetDefaultObjectType() interface{} {
	return "workflow.RollbackWorkflowTask"
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowRollbackWorkflowTask) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowRollbackWorkflowTask) SetName(v string) {
	o.Name = &v
}

// GetRefName returns the RefName field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetRefName() string {
	if o == nil || IsNil(o.RefName) {
		var ret string
		return ret
	}
	return *o.RefName
}

// GetRefNameOk returns a tuple with the RefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetRefNameOk() (*string, bool) {
	if o == nil || IsNil(o.RefName) {
		return nil, false
	}
	return o.RefName, true
}

// HasRefName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasRefName() bool {
	if o != nil && !IsNil(o.RefName) {
		return true
	}

	return false
}

// SetRefName gets a reference to the given string and assigns it to the RefName field.
func (o *WorkflowRollbackWorkflowTask) SetRefName(v string) {
	o.RefName = &v
}

// GetRollbackCompleted returns the RollbackCompleted field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetRollbackCompleted() bool {
	if o == nil || IsNil(o.RollbackCompleted) {
		var ret bool
		return ret
	}
	return *o.RollbackCompleted
}

// GetRollbackCompletedOk returns a tuple with the RollbackCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetRollbackCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.RollbackCompleted) {
		return nil, false
	}
	return o.RollbackCompleted, true
}

// HasRollbackCompleted returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasRollbackCompleted() bool {
	if o != nil && !IsNil(o.RollbackCompleted) {
		return true
	}

	return false
}

// SetRollbackCompleted gets a reference to the given bool and assigns it to the RollbackCompleted field.
func (o *WorkflowRollbackWorkflowTask) SetRollbackCompleted(v bool) {
	o.RollbackCompleted = &v
}

// GetRollbackTaskName returns the RollbackTaskName field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetRollbackTaskName() string {
	if o == nil || IsNil(o.RollbackTaskName) {
		var ret string
		return ret
	}
	return *o.RollbackTaskName
}

// GetRollbackTaskNameOk returns a tuple with the RollbackTaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetRollbackTaskNameOk() (*string, bool) {
	if o == nil || IsNil(o.RollbackTaskName) {
		return nil, false
	}
	return o.RollbackTaskName, true
}

// HasRollbackTaskName returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasRollbackTaskName() bool {
	if o != nil && !IsNil(o.RollbackTaskName) {
		return true
	}

	return false
}

// SetRollbackTaskName gets a reference to the given string and assigns it to the RollbackTaskName field.
func (o *WorkflowRollbackWorkflowTask) SetRollbackTaskName(v string) {
	o.RollbackTaskName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRollbackWorkflowTask) SetStatus(v string) {
	o.Status = &v
}

// GetTaskInfoMoid returns the TaskInfoMoid field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetTaskInfoMoid() string {
	if o == nil || IsNil(o.TaskInfoMoid) {
		var ret string
		return ret
	}
	return *o.TaskInfoMoid
}

// GetTaskInfoMoidOk returns a tuple with the TaskInfoMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetTaskInfoMoidOk() (*string, bool) {
	if o == nil || IsNil(o.TaskInfoMoid) {
		return nil, false
	}
	return o.TaskInfoMoid, true
}

// HasTaskInfoMoid returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasTaskInfoMoid() bool {
	if o != nil && !IsNil(o.TaskInfoMoid) {
		return true
	}

	return false
}

// SetTaskInfoMoid gets a reference to the given string and assigns it to the TaskInfoMoid field.
func (o *WorkflowRollbackWorkflowTask) SetTaskInfoMoid(v string) {
	o.TaskInfoMoid = &v
}

// GetTaskPath returns the TaskPath field value if set, zero value otherwise.
func (o *WorkflowRollbackWorkflowTask) GetTaskPath() string {
	if o == nil || IsNil(o.TaskPath) {
		var ret string
		return ret
	}
	return *o.TaskPath
}

// GetTaskPathOk returns a tuple with the TaskPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRollbackWorkflowTask) GetTaskPathOk() (*string, bool) {
	if o == nil || IsNil(o.TaskPath) {
		return nil, false
	}
	return o.TaskPath, true
}

// HasTaskPath returns a boolean if a field has been set.
func (o *WorkflowRollbackWorkflowTask) HasTaskPath() bool {
	if o != nil && !IsNil(o.TaskPath) {
		return true
	}

	return false
}

// SetTaskPath gets a reference to the given string and assigns it to the TaskPath field.
func (o *WorkflowRollbackWorkflowTask) SetTaskPath(v string) {
	o.TaskPath = &v
}

func (o WorkflowRollbackWorkflowTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRollbackWorkflowTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.RefName) {
		toSerialize["RefName"] = o.RefName
	}
	if !IsNil(o.RollbackCompleted) {
		toSerialize["RollbackCompleted"] = o.RollbackCompleted
	}
	if !IsNil(o.RollbackTaskName) {
		toSerialize["RollbackTaskName"] = o.RollbackTaskName
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TaskInfoMoid) {
		toSerialize["TaskInfoMoid"] = o.TaskInfoMoid
	}
	if !IsNil(o.TaskPath) {
		toSerialize["TaskPath"] = o.TaskPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowRollbackWorkflowTask) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowRollbackWorkflowTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the rollback task.
		Description *string `json:"Description,omitempty"`
		// Name of TaskInfo that needs to be rolled back.
		Name *string `json:"Name,omitempty"`
		// Reference name of TaskInfo that need to be rolled back.
		RefName *string `json:"RefName,omitempty"`
		// Status of the rollback operation for the task.
		RollbackCompleted *bool `json:"RollbackCompleted,omitempty"`
		// Name of TaskInfo that performs the rollback operation.
		RollbackTaskName *string `json:"RollbackTaskName,omitempty"`
		// Status of the rollback task. By default, task status will be not started. Task status will be set to completed on successful execution, otherwise it will be set to failed. * `NotStarted` - Status of rollback task when it is not started rollback. * `NotSupported` - Status of task when it is not supporting rollback. * `Completed` - Status of rollback task once execution is successful. * `Failed` - Status of rollback task when it is failed. * `Disabled` - Status of rollback task when rollback is disabled.
		Status *string `json:"Status,omitempty"`
		// Moid of TaskInfo that supports rollback operation.
		TaskInfoMoid *string `json:"TaskInfoMoid,omitempty"`
		// Path of rollback task if it is inside sub-workflow.
		TaskPath *string `json:"TaskPath,omitempty"`
	}

	varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct := WorkflowRollbackWorkflowTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowRollbackWorkflowTask := _WorkflowRollbackWorkflowTask{}
		varWorkflowRollbackWorkflowTask.ClassId = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.ClassId
		varWorkflowRollbackWorkflowTask.ObjectType = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowRollbackWorkflowTask.Description = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.Description
		varWorkflowRollbackWorkflowTask.Name = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.Name
		varWorkflowRollbackWorkflowTask.RefName = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.RefName
		varWorkflowRollbackWorkflowTask.RollbackCompleted = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.RollbackCompleted
		varWorkflowRollbackWorkflowTask.RollbackTaskName = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.RollbackTaskName
		varWorkflowRollbackWorkflowTask.Status = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.Status
		varWorkflowRollbackWorkflowTask.TaskInfoMoid = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.TaskInfoMoid
		varWorkflowRollbackWorkflowTask.TaskPath = varWorkflowRollbackWorkflowTaskWithoutEmbeddedStruct.TaskPath
		*o = WorkflowRollbackWorkflowTask(varWorkflowRollbackWorkflowTask)
	} else {
		return err
	}

	varWorkflowRollbackWorkflowTask := _WorkflowRollbackWorkflowTask{}

	err = json.Unmarshal(data, &varWorkflowRollbackWorkflowTask)
	if err == nil {
		o.MoBaseComplexType = varWorkflowRollbackWorkflowTask.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RefName")
		delete(additionalProperties, "RollbackCompleted")
		delete(additionalProperties, "RollbackTaskName")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "TaskInfoMoid")
		delete(additionalProperties, "TaskPath")

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

type NullableWorkflowRollbackWorkflowTask struct {
	value *WorkflowRollbackWorkflowTask
	isSet bool
}

func (v NullableWorkflowRollbackWorkflowTask) Get() *WorkflowRollbackWorkflowTask {
	return v.value
}

func (v *NullableWorkflowRollbackWorkflowTask) Set(val *WorkflowRollbackWorkflowTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRollbackWorkflowTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRollbackWorkflowTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRollbackWorkflowTask(val *WorkflowRollbackWorkflowTask) *NullableWorkflowRollbackWorkflowTask {
	return &NullableWorkflowRollbackWorkflowTask{value: val, isSet: true}
}

func (v NullableWorkflowRollbackWorkflowTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRollbackWorkflowTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
