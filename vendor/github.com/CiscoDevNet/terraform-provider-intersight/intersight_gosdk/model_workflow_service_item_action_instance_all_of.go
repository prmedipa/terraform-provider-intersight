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
	"time"
)

// WorkflowServiceItemActionInstanceAllOf Definition of the list of properties defined in 'workflow.ServiceItemActionInstance', excluding properties defined in parent classes.
type WorkflowServiceItemActionInstanceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the action that needs to be performed on the service item instance. * `None` - No action is set, this is the default value for action field. * `Validate` - Validate the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the service item action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
	Action *string `json:"Action,omitempty"`
	// The time when the action was stopped or completed execution last time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Inputs for a service item action and the format is specified by input definition of the service item action definition.
	Input interface{} `json:"Input,omitempty"`
	// The last action that was issued on the action definition workflows is saved in this property. * `None` - No action is set, this is the default value for action field. * `Validate` - Validate the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the service item action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
	LastAction *string `json:"LastAction,omitempty"`
	// Name for the action instance is created in the system by appending name of the service item instance to the name of the action definition.
	Name *string `json:"Name,omitempty"`
	// The time when the action was started for execution last time.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// State of the service item action instance. * `NotStarted` - An action on the service item is not yet started and it is in a draft mode. A service item action instance can be deleted in this state. * `Validating` - A validate action has been triggered on the action and until it completes the start action cannot be issued. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The action on the service item instance failed and can be retried. * `Completed` - The action on the service item instance completed successfully. * `Stopping` - The stop action is running on the action instance.
	Status                      *string                                          `json:"Status,omitempty"`
	ActionWorkflowInfo          *WorkflowWorkflowInfoRelationship                `json:"ActionWorkflowInfo,omitempty"`
	ServiceItemActionDefinition *WorkflowServiceItemActionDefinitionRelationship `json:"ServiceItemActionDefinition,omitempty"`
	ServiceItemDefinition       *WorkflowServiceItemDefinitionRelationship       `json:"ServiceItemDefinition,omitempty"`
	ServiceItemInstance         *WorkflowServiceItemInstanceRelationship         `json:"ServiceItemInstance,omitempty"`
	StopWorkflowInfo            *WorkflowWorkflowInfoRelationship                `json:"StopWorkflowInfo,omitempty"`
	ValidationWorkflowInfo      *WorkflowWorkflowInfoRelationship                `json:"ValidationWorkflowInfo,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _WorkflowServiceItemActionInstanceAllOf WorkflowServiceItemActionInstanceAllOf

// NewWorkflowServiceItemActionInstanceAllOf instantiates a new WorkflowServiceItemActionInstanceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemActionInstanceAllOf(classId string, objectType string) *WorkflowServiceItemActionInstanceAllOf {
	this := WorkflowServiceItemActionInstanceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewWorkflowServiceItemActionInstanceAllOfWithDefaults instantiates a new WorkflowServiceItemActionInstanceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemActionInstanceAllOfWithDefaults() *WorkflowServiceItemActionInstanceAllOf {
	this := WorkflowServiceItemActionInstanceAllOf{}
	var classId string = "workflow.ServiceItemActionInstance"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemActionInstance"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemActionInstanceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemActionInstanceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemActionInstanceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemActionInstanceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetAction(v string) {
	o.Action = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionInstanceAllOf) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionInstanceAllOf) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetInput(v interface{}) {
	o.Input = v
}

// GetLastAction returns the LastAction field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetLastAction() string {
	if o == nil || o.LastAction == nil {
		var ret string
		return ret
	}
	return *o.LastAction
}

// GetLastActionOk returns a tuple with the LastAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetLastActionOk() (*string, bool) {
	if o == nil || o.LastAction == nil {
		return nil, false
	}
	return o.LastAction, true
}

// HasLastAction returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasLastAction() bool {
	if o != nil && o.LastAction != nil {
		return true
	}

	return false
}

// SetLastAction gets a reference to the given string and assigns it to the LastAction field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetLastAction(v string) {
	o.LastAction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetActionWorkflowInfo returns the ActionWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetActionWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.ActionWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.ActionWorkflowInfo
}

// GetActionWorkflowInfoOk returns a tuple with the ActionWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetActionWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.ActionWorkflowInfo == nil {
		return nil, false
	}
	return o.ActionWorkflowInfo, true
}

// HasActionWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasActionWorkflowInfo() bool {
	if o != nil && o.ActionWorkflowInfo != nil {
		return true
	}

	return false
}

// SetActionWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the ActionWorkflowInfo field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetActionWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.ActionWorkflowInfo = &v
}

// GetServiceItemActionDefinition returns the ServiceItemActionDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemActionDefinition() WorkflowServiceItemActionDefinitionRelationship {
	if o == nil || o.ServiceItemActionDefinition == nil {
		var ret WorkflowServiceItemActionDefinitionRelationship
		return ret
	}
	return *o.ServiceItemActionDefinition
}

// GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemActionDefinitionOk() (*WorkflowServiceItemActionDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemActionDefinition == nil {
		return nil, false
	}
	return o.ServiceItemActionDefinition, true
}

// HasServiceItemActionDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemActionDefinition() bool {
	if o != nil && o.ServiceItemActionDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemActionDefinition gets a reference to the given WorkflowServiceItemActionDefinitionRelationship and assigns it to the ServiceItemActionDefinition field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemActionDefinition(v WorkflowServiceItemActionDefinitionRelationship) {
	o.ServiceItemActionDefinition = &v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || o.ServiceItemDefinition == nil {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemDefinition == nil {
		return nil, false
	}
	return o.ServiceItemDefinition, true
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given WorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition = &v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || o.ServiceItemInstance == nil {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil || o.ServiceItemInstance == nil {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance != nil {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given WorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance = &v
}

// GetStopWorkflowInfo returns the StopWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStopWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.StopWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.StopWorkflowInfo
}

// GetStopWorkflowInfoOk returns a tuple with the StopWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetStopWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.StopWorkflowInfo == nil {
		return nil, false
	}
	return o.StopWorkflowInfo, true
}

// HasStopWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasStopWorkflowInfo() bool {
	if o != nil && o.StopWorkflowInfo != nil {
		return true
	}

	return false
}

// SetStopWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the StopWorkflowInfo field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetStopWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.StopWorkflowInfo = &v
}

// GetValidationWorkflowInfo returns the ValidationWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionInstanceAllOf) GetValidationWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.ValidationWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.ValidationWorkflowInfo
}

// GetValidationWorkflowInfoOk returns a tuple with the ValidationWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) GetValidationWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.ValidationWorkflowInfo == nil {
		return nil, false
	}
	return o.ValidationWorkflowInfo, true
}

// HasValidationWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionInstanceAllOf) HasValidationWorkflowInfo() bool {
	if o != nil && o.ValidationWorkflowInfo != nil {
		return true
	}

	return false
}

// SetValidationWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the ValidationWorkflowInfo field.
func (o *WorkflowServiceItemActionInstanceAllOf) SetValidationWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.ValidationWorkflowInfo = &v
}

func (o WorkflowServiceItemActionInstanceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Action != nil {
		toSerialize["Action"] = o.Action
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.LastAction != nil {
		toSerialize["LastAction"] = o.LastAction
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.ActionWorkflowInfo != nil {
		toSerialize["ActionWorkflowInfo"] = o.ActionWorkflowInfo
	}
	if o.ServiceItemActionDefinition != nil {
		toSerialize["ServiceItemActionDefinition"] = o.ServiceItemActionDefinition
	}
	if o.ServiceItemDefinition != nil {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition
	}
	if o.ServiceItemInstance != nil {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance
	}
	if o.StopWorkflowInfo != nil {
		toSerialize["StopWorkflowInfo"] = o.StopWorkflowInfo
	}
	if o.ValidationWorkflowInfo != nil {
		toSerialize["ValidationWorkflowInfo"] = o.ValidationWorkflowInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemActionInstanceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemActionInstanceAllOf := _WorkflowServiceItemActionInstanceAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemActionInstanceAllOf); err == nil {
		*o = WorkflowServiceItemActionInstanceAllOf(varWorkflowServiceItemActionInstanceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Action")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "LastAction")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "ActionWorkflowInfo")
		delete(additionalProperties, "ServiceItemActionDefinition")
		delete(additionalProperties, "ServiceItemDefinition")
		delete(additionalProperties, "ServiceItemInstance")
		delete(additionalProperties, "StopWorkflowInfo")
		delete(additionalProperties, "ValidationWorkflowInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemActionInstanceAllOf struct {
	value *WorkflowServiceItemActionInstanceAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemActionInstanceAllOf) Get() *WorkflowServiceItemActionInstanceAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemActionInstanceAllOf) Set(val *WorkflowServiceItemActionInstanceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemActionInstanceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemActionInstanceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemActionInstanceAllOf(val *WorkflowServiceItemActionInstanceAllOf) *NullableWorkflowServiceItemActionInstanceAllOf {
	return &NullableWorkflowServiceItemActionInstanceAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemActionInstanceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemActionInstanceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
