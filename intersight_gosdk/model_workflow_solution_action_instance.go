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
	"time"
)

// WorkflowSolutionActionInstance Solution action instance which represents one action on a solution instance.
type WorkflowSolutionActionInstance struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the action that needs to be performed on the solution instance. * `None` - No action is set, this is the default value for action field. * `Validate` - Validation the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the solution action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
	Action *string `json:"Action,omitempty"`
	// The time when the action was stopped or completed execution last time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// Inputs for a solution action and the format is specified by input definition of the solution action definition.
	Input interface{} `json:"Input,omitempty"`
	// The last action that was issued on the action definition workflows is saved in this property. * `None` - No action is set, this is the default value for action field. * `Validate` - Validation the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the solution action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
	LastAction *string `json:"LastAction,omitempty"`
	// Name for the action instance is created in the system by appending name of the solution instance to the name of the action definition.
	Name *string `json:"Name,omitempty"`
	// The time when the action was started for execution last time.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// State of the solution action instance. * `NotStarted` - Solution action is not yet started and it is in a draft mode. A solution action instance can be deleted in this state. * `Validating` - A validate action has been triggered on the action and until it completes the start action cannot be issued. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The action on the solution failed and can be retried. * `Completed` - The action on the solution completed successfully. * `Stopping` - The stop action is running on the action instance.
	Status *string `json:"Status,omitempty"`
	// Stores the upgraded Moid for help during future lookups.
	UpgradedMoid             *string                                       `json:"UpgradedMoid,omitempty"`
	ActionWorkflowInfo       *WorkflowWorkflowInfoRelationship             `json:"ActionWorkflowInfo,omitempty"`
	SolutionActionDefinition *WorkflowSolutionActionDefinitionRelationship `json:"SolutionActionDefinition,omitempty"`
	SolutionDefinition       *WorkflowSolutionDefinitionRelationship       `json:"SolutionDefinition,omitempty"`
	SolutionInstance         *WorkflowSolutionInstanceRelationship         `json:"SolutionInstance,omitempty"`
	StopWorkflowInfo         *WorkflowWorkflowInfoRelationship             `json:"StopWorkflowInfo,omitempty"`
	ValidationWorkflowInfo   *WorkflowWorkflowInfoRelationship             `json:"ValidationWorkflowInfo,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _WorkflowSolutionActionInstance WorkflowSolutionActionInstance

// NewWorkflowSolutionActionInstance instantiates a new WorkflowSolutionActionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSolutionActionInstance(classId string, objectType string) *WorkflowSolutionActionInstance {
	this := WorkflowSolutionActionInstance{}
	this.ClassId = classId
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// NewWorkflowSolutionActionInstanceWithDefaults instantiates a new WorkflowSolutionActionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSolutionActionInstanceWithDefaults() *WorkflowSolutionActionInstance {
	this := WorkflowSolutionActionInstance{}
	var classId string = "workflow.SolutionActionInstance"
	this.ClassId = classId
	var objectType string = "workflow.SolutionActionInstance"
	this.ObjectType = objectType
	var action string = "None"
	this.Action = &action
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSolutionActionInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSolutionActionInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSolutionActionInstance) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSolutionActionInstance) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *WorkflowSolutionActionInstance) SetAction(v string) {
	o.Action = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *WorkflowSolutionActionInstance) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetInput returns the Input field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionInstance) GetInput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionInstance) GetInputOk() (*interface{}, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return &o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given interface{} and assigns it to the Input field.
func (o *WorkflowSolutionActionInstance) SetInput(v interface{}) {
	o.Input = v
}

// GetLastAction returns the LastAction field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetLastAction() string {
	if o == nil || o.LastAction == nil {
		var ret string
		return ret
	}
	return *o.LastAction
}

// GetLastActionOk returns a tuple with the LastAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetLastActionOk() (*string, bool) {
	if o == nil || o.LastAction == nil {
		return nil, false
	}
	return o.LastAction, true
}

// HasLastAction returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasLastAction() bool {
	if o != nil && o.LastAction != nil {
		return true
	}

	return false
}

// SetLastAction gets a reference to the given string and assigns it to the LastAction field.
func (o *WorkflowSolutionActionInstance) SetLastAction(v string) {
	o.LastAction = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSolutionActionInstance) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *WorkflowSolutionActionInstance) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowSolutionActionInstance) SetStatus(v string) {
	o.Status = &v
}

// GetUpgradedMoid returns the UpgradedMoid field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetUpgradedMoid() string {
	if o == nil || o.UpgradedMoid == nil {
		var ret string
		return ret
	}
	return *o.UpgradedMoid
}

// GetUpgradedMoidOk returns a tuple with the UpgradedMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetUpgradedMoidOk() (*string, bool) {
	if o == nil || o.UpgradedMoid == nil {
		return nil, false
	}
	return o.UpgradedMoid, true
}

// HasUpgradedMoid returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasUpgradedMoid() bool {
	if o != nil && o.UpgradedMoid != nil {
		return true
	}

	return false
}

// SetUpgradedMoid gets a reference to the given string and assigns it to the UpgradedMoid field.
func (o *WorkflowSolutionActionInstance) SetUpgradedMoid(v string) {
	o.UpgradedMoid = &v
}

// GetActionWorkflowInfo returns the ActionWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetActionWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.ActionWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.ActionWorkflowInfo
}

// GetActionWorkflowInfoOk returns a tuple with the ActionWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetActionWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.ActionWorkflowInfo == nil {
		return nil, false
	}
	return o.ActionWorkflowInfo, true
}

// HasActionWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasActionWorkflowInfo() bool {
	if o != nil && o.ActionWorkflowInfo != nil {
		return true
	}

	return false
}

// SetActionWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the ActionWorkflowInfo field.
func (o *WorkflowSolutionActionInstance) SetActionWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.ActionWorkflowInfo = &v
}

// GetSolutionActionDefinition returns the SolutionActionDefinition field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetSolutionActionDefinition() WorkflowSolutionActionDefinitionRelationship {
	if o == nil || o.SolutionActionDefinition == nil {
		var ret WorkflowSolutionActionDefinitionRelationship
		return ret
	}
	return *o.SolutionActionDefinition
}

// GetSolutionActionDefinitionOk returns a tuple with the SolutionActionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetSolutionActionDefinitionOk() (*WorkflowSolutionActionDefinitionRelationship, bool) {
	if o == nil || o.SolutionActionDefinition == nil {
		return nil, false
	}
	return o.SolutionActionDefinition, true
}

// HasSolutionActionDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasSolutionActionDefinition() bool {
	if o != nil && o.SolutionActionDefinition != nil {
		return true
	}

	return false
}

// SetSolutionActionDefinition gets a reference to the given WorkflowSolutionActionDefinitionRelationship and assigns it to the SolutionActionDefinition field.
func (o *WorkflowSolutionActionInstance) SetSolutionActionDefinition(v WorkflowSolutionActionDefinitionRelationship) {
	o.SolutionActionDefinition = &v
}

// GetSolutionDefinition returns the SolutionDefinition field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship {
	if o == nil || o.SolutionDefinition == nil {
		var ret WorkflowSolutionDefinitionRelationship
		return ret
	}
	return *o.SolutionDefinition
}

// GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool) {
	if o == nil || o.SolutionDefinition == nil {
		return nil, false
	}
	return o.SolutionDefinition, true
}

// HasSolutionDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasSolutionDefinition() bool {
	if o != nil && o.SolutionDefinition != nil {
		return true
	}

	return false
}

// SetSolutionDefinition gets a reference to the given WorkflowSolutionDefinitionRelationship and assigns it to the SolutionDefinition field.
func (o *WorkflowSolutionActionInstance) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship) {
	o.SolutionDefinition = &v
}

// GetSolutionInstance returns the SolutionInstance field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetSolutionInstance() WorkflowSolutionInstanceRelationship {
	if o == nil || o.SolutionInstance == nil {
		var ret WorkflowSolutionInstanceRelationship
		return ret
	}
	return *o.SolutionInstance
}

// GetSolutionInstanceOk returns a tuple with the SolutionInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetSolutionInstanceOk() (*WorkflowSolutionInstanceRelationship, bool) {
	if o == nil || o.SolutionInstance == nil {
		return nil, false
	}
	return o.SolutionInstance, true
}

// HasSolutionInstance returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasSolutionInstance() bool {
	if o != nil && o.SolutionInstance != nil {
		return true
	}

	return false
}

// SetSolutionInstance gets a reference to the given WorkflowSolutionInstanceRelationship and assigns it to the SolutionInstance field.
func (o *WorkflowSolutionActionInstance) SetSolutionInstance(v WorkflowSolutionInstanceRelationship) {
	o.SolutionInstance = &v
}

// GetStopWorkflowInfo returns the StopWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetStopWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.StopWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.StopWorkflowInfo
}

// GetStopWorkflowInfoOk returns a tuple with the StopWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetStopWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.StopWorkflowInfo == nil {
		return nil, false
	}
	return o.StopWorkflowInfo, true
}

// HasStopWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasStopWorkflowInfo() bool {
	if o != nil && o.StopWorkflowInfo != nil {
		return true
	}

	return false
}

// SetStopWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the StopWorkflowInfo field.
func (o *WorkflowSolutionActionInstance) SetStopWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.StopWorkflowInfo = &v
}

// GetValidationWorkflowInfo returns the ValidationWorkflowInfo field value if set, zero value otherwise.
func (o *WorkflowSolutionActionInstance) GetValidationWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || o.ValidationWorkflowInfo == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.ValidationWorkflowInfo
}

// GetValidationWorkflowInfoOk returns a tuple with the ValidationWorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionInstance) GetValidationWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.ValidationWorkflowInfo == nil {
		return nil, false
	}
	return o.ValidationWorkflowInfo, true
}

// HasValidationWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowSolutionActionInstance) HasValidationWorkflowInfo() bool {
	if o != nil && o.ValidationWorkflowInfo != nil {
		return true
	}

	return false
}

// SetValidationWorkflowInfo gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the ValidationWorkflowInfo field.
func (o *WorkflowSolutionActionInstance) SetValidationWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.ValidationWorkflowInfo = &v
}

func (o WorkflowSolutionActionInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
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
	if o.UpgradedMoid != nil {
		toSerialize["UpgradedMoid"] = o.UpgradedMoid
	}
	if o.ActionWorkflowInfo != nil {
		toSerialize["ActionWorkflowInfo"] = o.ActionWorkflowInfo
	}
	if o.SolutionActionDefinition != nil {
		toSerialize["SolutionActionDefinition"] = o.SolutionActionDefinition
	}
	if o.SolutionDefinition != nil {
		toSerialize["SolutionDefinition"] = o.SolutionDefinition
	}
	if o.SolutionInstance != nil {
		toSerialize["SolutionInstance"] = o.SolutionInstance
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

func (o *WorkflowSolutionActionInstance) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowSolutionActionInstanceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the action that needs to be performed on the solution instance. * `None` - No action is set, this is the default value for action field. * `Validate` - Validation the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the solution action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
		Action *string `json:"Action,omitempty"`
		// The time when the action was stopped or completed execution last time.
		EndTime *time.Time `json:"EndTime,omitempty"`
		// Inputs for a solution action and the format is specified by input definition of the solution action definition.
		Input interface{} `json:"Input,omitempty"`
		// The last action that was issued on the action definition workflows is saved in this property. * `None` - No action is set, this is the default value for action field. * `Validate` - Validation the action instance inputs and run the validation workflows. * `Start` - Start a new execution of the action instance. * `Retry` - Retry the solution action instance from the beginning. * `RetryFailed` - Retry the workflow that has failed from the failure point. * `Cancel` - Cancel the core workflow that is in running or waiting state. This action can be used when the workflows are stuck and not progressing. * `Stop` - Stop the action instance which is in progress and didn't complete successfully. Use this action to clear the state and then delete the action instance. A completed action cannot be stopped.
		LastAction *string `json:"LastAction,omitempty"`
		// Name for the action instance is created in the system by appending name of the solution instance to the name of the action definition.
		Name *string `json:"Name,omitempty"`
		// The time when the action was started for execution last time.
		StartTime *time.Time `json:"StartTime,omitempty"`
		// State of the solution action instance. * `NotStarted` - Solution action is not yet started and it is in a draft mode. A solution action instance can be deleted in this state. * `Validating` - A validate action has been triggered on the action and until it completes the start action cannot be issued. * `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started. * `Failed` - The action on the solution failed and can be retried. * `Completed` - The action on the solution completed successfully. * `Stopping` - The stop action is running on the action instance.
		Status *string `json:"Status,omitempty"`
		// Stores the upgraded Moid for help during future lookups.
		UpgradedMoid             *string                                       `json:"UpgradedMoid,omitempty"`
		ActionWorkflowInfo       *WorkflowWorkflowInfoRelationship             `json:"ActionWorkflowInfo,omitempty"`
		SolutionActionDefinition *WorkflowSolutionActionDefinitionRelationship `json:"SolutionActionDefinition,omitempty"`
		SolutionDefinition       *WorkflowSolutionDefinitionRelationship       `json:"SolutionDefinition,omitempty"`
		SolutionInstance         *WorkflowSolutionInstanceRelationship         `json:"SolutionInstance,omitempty"`
		StopWorkflowInfo         *WorkflowWorkflowInfoRelationship             `json:"StopWorkflowInfo,omitempty"`
		ValidationWorkflowInfo   *WorkflowWorkflowInfoRelationship             `json:"ValidationWorkflowInfo,omitempty"`
	}

	varWorkflowSolutionActionInstanceWithoutEmbeddedStruct := WorkflowSolutionActionInstanceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionActionInstanceWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowSolutionActionInstance := _WorkflowSolutionActionInstance{}
		varWorkflowSolutionActionInstance.ClassId = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.ClassId
		varWorkflowSolutionActionInstance.ObjectType = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.ObjectType
		varWorkflowSolutionActionInstance.Action = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.Action
		varWorkflowSolutionActionInstance.EndTime = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.EndTime
		varWorkflowSolutionActionInstance.Input = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.Input
		varWorkflowSolutionActionInstance.LastAction = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.LastAction
		varWorkflowSolutionActionInstance.Name = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.Name
		varWorkflowSolutionActionInstance.StartTime = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.StartTime
		varWorkflowSolutionActionInstance.Status = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.Status
		varWorkflowSolutionActionInstance.UpgradedMoid = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.UpgradedMoid
		varWorkflowSolutionActionInstance.ActionWorkflowInfo = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.ActionWorkflowInfo
		varWorkflowSolutionActionInstance.SolutionActionDefinition = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.SolutionActionDefinition
		varWorkflowSolutionActionInstance.SolutionDefinition = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.SolutionDefinition
		varWorkflowSolutionActionInstance.SolutionInstance = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.SolutionInstance
		varWorkflowSolutionActionInstance.StopWorkflowInfo = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.StopWorkflowInfo
		varWorkflowSolutionActionInstance.ValidationWorkflowInfo = varWorkflowSolutionActionInstanceWithoutEmbeddedStruct.ValidationWorkflowInfo
		*o = WorkflowSolutionActionInstance(varWorkflowSolutionActionInstance)
	} else {
		return err
	}

	varWorkflowSolutionActionInstance := _WorkflowSolutionActionInstance{}

	err = json.Unmarshal(bytes, &varWorkflowSolutionActionInstance)
	if err == nil {
		o.MoBaseMo = varWorkflowSolutionActionInstance.MoBaseMo
	} else {
		return err
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
		delete(additionalProperties, "UpgradedMoid")
		delete(additionalProperties, "ActionWorkflowInfo")
		delete(additionalProperties, "SolutionActionDefinition")
		delete(additionalProperties, "SolutionDefinition")
		delete(additionalProperties, "SolutionInstance")
		delete(additionalProperties, "StopWorkflowInfo")
		delete(additionalProperties, "ValidationWorkflowInfo")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableWorkflowSolutionActionInstance struct {
	value *WorkflowSolutionActionInstance
	isSet bool
}

func (v NullableWorkflowSolutionActionInstance) Get() *WorkflowSolutionActionInstance {
	return v.value
}

func (v *NullableWorkflowSolutionActionInstance) Set(val *WorkflowSolutionActionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSolutionActionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSolutionActionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSolutionActionInstance(val *WorkflowSolutionActionInstance) *NullableWorkflowSolutionActionInstance {
	return &NullableWorkflowSolutionActionInstance{value: val, isSet: true}
}

func (v NullableWorkflowSolutionActionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSolutionActionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
