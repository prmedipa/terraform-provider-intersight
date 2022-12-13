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
	"reflect"
	"strings"
)

// WorkflowWorkerTask A WorkerTask is a simple task and the smallest granularity of work that can be defined as a task.
type WorkflowWorkerTask struct {
	WorkflowAbstractWorkerTask
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specify the catalog moid that this task belongs.
	CatalogMoid *string `json:"CatalogMoid,omitempty"`
	// The resolved referenced task definition managed object.
	TaskDefinitionId *string `json:"TaskDefinitionId,omitempty"`
	// The qualified name of task that should be executed.
	TaskDefinitionName *string `json:"TaskDefinitionName,omitempty"`
	// The task definition version to use in this workflow. When no version is specified then the default version of the task at the time of creating or updating this workflow is used.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkerTask WorkflowWorkerTask

// NewWorkflowWorkerTask instantiates a new WorkflowWorkerTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkerTask(classId string, objectType string) *WorkflowWorkerTask {
	this := WorkflowWorkerTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	var rollbackDisabled bool = false
	this.RollbackDisabled = &rollbackDisabled
	var useDefault bool = false
	this.UseDefault = &useDefault
	return &this
}

// NewWorkflowWorkerTaskWithDefaults instantiates a new WorkflowWorkerTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkerTaskWithDefaults() *WorkflowWorkerTask {
	this := WorkflowWorkerTask{}
	var classId string = "workflow.WorkerTask"
	this.ClassId = classId
	var objectType string = "workflow.WorkerTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkerTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkerTask) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkerTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkerTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalogMoid returns the CatalogMoid field value if set, zero value otherwise.
func (o *WorkflowWorkerTask) GetCatalogMoid() string {
	if o == nil || o.CatalogMoid == nil {
		var ret string
		return ret
	}
	return *o.CatalogMoid
}

// GetCatalogMoidOk returns a tuple with the CatalogMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetCatalogMoidOk() (*string, bool) {
	if o == nil || o.CatalogMoid == nil {
		return nil, false
	}
	return o.CatalogMoid, true
}

// HasCatalogMoid returns a boolean if a field has been set.
func (o *WorkflowWorkerTask) HasCatalogMoid() bool {
	if o != nil && o.CatalogMoid != nil {
		return true
	}

	return false
}

// SetCatalogMoid gets a reference to the given string and assigns it to the CatalogMoid field.
func (o *WorkflowWorkerTask) SetCatalogMoid(v string) {
	o.CatalogMoid = &v
}

// GetTaskDefinitionId returns the TaskDefinitionId field value if set, zero value otherwise.
func (o *WorkflowWorkerTask) GetTaskDefinitionId() string {
	if o == nil || o.TaskDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.TaskDefinitionId
}

// GetTaskDefinitionIdOk returns a tuple with the TaskDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetTaskDefinitionIdOk() (*string, bool) {
	if o == nil || o.TaskDefinitionId == nil {
		return nil, false
	}
	return o.TaskDefinitionId, true
}

// HasTaskDefinitionId returns a boolean if a field has been set.
func (o *WorkflowWorkerTask) HasTaskDefinitionId() bool {
	if o != nil && o.TaskDefinitionId != nil {
		return true
	}

	return false
}

// SetTaskDefinitionId gets a reference to the given string and assigns it to the TaskDefinitionId field.
func (o *WorkflowWorkerTask) SetTaskDefinitionId(v string) {
	o.TaskDefinitionId = &v
}

// GetTaskDefinitionName returns the TaskDefinitionName field value if set, zero value otherwise.
func (o *WorkflowWorkerTask) GetTaskDefinitionName() string {
	if o == nil || o.TaskDefinitionName == nil {
		var ret string
		return ret
	}
	return *o.TaskDefinitionName
}

// GetTaskDefinitionNameOk returns a tuple with the TaskDefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetTaskDefinitionNameOk() (*string, bool) {
	if o == nil || o.TaskDefinitionName == nil {
		return nil, false
	}
	return o.TaskDefinitionName, true
}

// HasTaskDefinitionName returns a boolean if a field has been set.
func (o *WorkflowWorkerTask) HasTaskDefinitionName() bool {
	if o != nil && o.TaskDefinitionName != nil {
		return true
	}

	return false
}

// SetTaskDefinitionName gets a reference to the given string and assigns it to the TaskDefinitionName field.
func (o *WorkflowWorkerTask) SetTaskDefinitionName(v string) {
	o.TaskDefinitionName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowWorkerTask) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkerTask) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowWorkerTask) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowWorkerTask) SetVersion(v int64) {
	o.Version = &v
}

func (o WorkflowWorkerTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowAbstractWorkerTask, errWorkflowAbstractWorkerTask := json.Marshal(o.WorkflowAbstractWorkerTask)
	if errWorkflowAbstractWorkerTask != nil {
		return []byte{}, errWorkflowAbstractWorkerTask
	}
	errWorkflowAbstractWorkerTask = json.Unmarshal([]byte(serializedWorkflowAbstractWorkerTask), &toSerialize)
	if errWorkflowAbstractWorkerTask != nil {
		return []byte{}, errWorkflowAbstractWorkerTask
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CatalogMoid != nil {
		toSerialize["CatalogMoid"] = o.CatalogMoid
	}
	if o.TaskDefinitionId != nil {
		toSerialize["TaskDefinitionId"] = o.TaskDefinitionId
	}
	if o.TaskDefinitionName != nil {
		toSerialize["TaskDefinitionName"] = o.TaskDefinitionName
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkerTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkerTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specify the catalog moid that this task belongs.
		CatalogMoid *string `json:"CatalogMoid,omitempty"`
		// The resolved referenced task definition managed object.
		TaskDefinitionId *string `json:"TaskDefinitionId,omitempty"`
		// The qualified name of task that should be executed.
		TaskDefinitionName *string `json:"TaskDefinitionName,omitempty"`
		// The task definition version to use in this workflow. When no version is specified then the default version of the task at the time of creating or updating this workflow is used.
		Version *int64 `json:"Version,omitempty"`
	}

	varWorkflowWorkerTaskWithoutEmbeddedStruct := WorkflowWorkerTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkerTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkerTask := _WorkflowWorkerTask{}
		varWorkflowWorkerTask.ClassId = varWorkflowWorkerTaskWithoutEmbeddedStruct.ClassId
		varWorkflowWorkerTask.ObjectType = varWorkflowWorkerTaskWithoutEmbeddedStruct.ObjectType
		varWorkflowWorkerTask.CatalogMoid = varWorkflowWorkerTaskWithoutEmbeddedStruct.CatalogMoid
		varWorkflowWorkerTask.TaskDefinitionId = varWorkflowWorkerTaskWithoutEmbeddedStruct.TaskDefinitionId
		varWorkflowWorkerTask.TaskDefinitionName = varWorkflowWorkerTaskWithoutEmbeddedStruct.TaskDefinitionName
		varWorkflowWorkerTask.Version = varWorkflowWorkerTaskWithoutEmbeddedStruct.Version
		*o = WorkflowWorkerTask(varWorkflowWorkerTask)
	} else {
		return err
	}

	varWorkflowWorkerTask := _WorkflowWorkerTask{}

	err = json.Unmarshal(bytes, &varWorkflowWorkerTask)
	if err == nil {
		o.WorkflowAbstractWorkerTask = varWorkflowWorkerTask.WorkflowAbstractWorkerTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CatalogMoid")
		delete(additionalProperties, "TaskDefinitionId")
		delete(additionalProperties, "TaskDefinitionName")
		delete(additionalProperties, "Version")

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

type NullableWorkflowWorkerTask struct {
	value *WorkflowWorkerTask
	isSet bool
}

func (v NullableWorkflowWorkerTask) Get() *WorkflowWorkerTask {
	return v.value
}

func (v *NullableWorkflowWorkerTask) Set(val *WorkflowWorkerTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkerTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkerTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkerTask(val *WorkflowWorkerTask) *NullableWorkflowWorkerTask {
	return &NullableWorkflowWorkerTask{value: val, isSet: true}
}

func (v NullableWorkflowWorkerTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkerTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
