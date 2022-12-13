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
)

// WorkflowSolutionActionDefinitionAllOf Definition of the list of properties defined in 'workflow.SolutionActionDefinition', excluding properties defined in parent classes.
type WorkflowSolutionActionDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of actionDefinition which decides on how to trigger the action. * `External` - External actions definition can be triggered by enduser to perform actions on the solution. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again. * `Internal` - Internal action definition is used to trigger periodic actions on the solution instance. * `Repetitive` - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created solution.
	ActionType            *string                            `json:"ActionType,omitempty"`
	AllowedInstanceStates []string                           `json:"AllowedInstanceStates,omitempty"`
	CoreWorkflows         []WorkflowActionWorkflowDefinition `json:"CoreWorkflows,omitempty"`
	// The description for this action which provides information on what are the pre-requisites to use this action on the solution and what features are supported by this action.
	Description     *string                `json:"Description,omitempty"`
	InputDefinition []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	// A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a solution definition.
	Name *string `json:"Name,omitempty"`
	// The output mappings from workflows in the action definition to the solution output definition. Any output from core or post-core workflow can be mapped to solution output definition. The output can be referred using the name of the workflow definition and the output name in the following format '${<ActionWorkflowDefinition.Name>.output.<outputName>'.
	OutputParameters interface{} `json:"OutputParameters,omitempty"`
	// Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity.
	Periodicity       *int64                             `json:"Periodicity,omitempty"`
	PostCoreWorkflows []WorkflowActionWorkflowDefinition `json:"PostCoreWorkflows,omitempty"`
	PreCoreWorkflows  []WorkflowActionWorkflowDefinition `json:"PreCoreWorkflows,omitempty"`
	StopWorkflows     []WorkflowActionWorkflowDefinition `json:"StopWorkflows,omitempty"`
	// Stores the upgraded Moid for help during future lookups.
	UpgradedMoid          *string                                 `json:"UpgradedMoid,omitempty"`
	ValidationInformation NullableWorkflowValidationInformation   `json:"ValidationInformation,omitempty"`
	ValidationWorkflows   []WorkflowActionWorkflowDefinition      `json:"ValidationWorkflows,omitempty"`
	SolutionDefinition    *WorkflowSolutionDefinitionRelationship `json:"SolutionDefinition,omitempty"`
	WorkflowDefinition    *WorkflowWorkflowDefinitionRelationship `json:"WorkflowDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowSolutionActionDefinitionAllOf WorkflowSolutionActionDefinitionAllOf

// NewWorkflowSolutionActionDefinitionAllOf instantiates a new WorkflowSolutionActionDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSolutionActionDefinitionAllOf(classId string, objectType string) *WorkflowSolutionActionDefinitionAllOf {
	this := WorkflowSolutionActionDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var actionType string = "External"
	this.ActionType = &actionType
	return &this
}

// NewWorkflowSolutionActionDefinitionAllOfWithDefaults instantiates a new WorkflowSolutionActionDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSolutionActionDefinitionAllOfWithDefaults() *WorkflowSolutionActionDefinitionAllOf {
	this := WorkflowSolutionActionDefinitionAllOf{}
	var classId string = "workflow.SolutionActionDefinition"
	this.ClassId = classId
	var objectType string = "workflow.SolutionActionDefinition"
	this.ObjectType = objectType
	var actionType string = "External"
	this.ActionType = &actionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowSolutionActionDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowSolutionActionDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowSolutionActionDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowSolutionActionDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetActionType() string {
	if o == nil || o.ActionType == nil {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetActionTypeOk() (*string, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetActionType(v string) {
	o.ActionType = &v
}

// GetAllowedInstanceStates returns the AllowedInstanceStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetAllowedInstanceStates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedInstanceStates
}

// GetAllowedInstanceStatesOk returns a tuple with the AllowedInstanceStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetAllowedInstanceStatesOk() ([]string, bool) {
	if o == nil || o.AllowedInstanceStates == nil {
		return nil, false
	}
	return o.AllowedInstanceStates, true
}

// HasAllowedInstanceStates returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasAllowedInstanceStates() bool {
	if o != nil && o.AllowedInstanceStates != nil {
		return true
	}

	return false
}

// SetAllowedInstanceStates gets a reference to the given []string and assigns it to the AllowedInstanceStates field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetAllowedInstanceStates(v []string) {
	o.AllowedInstanceStates = v
}

// GetCoreWorkflows returns the CoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetCoreWorkflows() []WorkflowActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowActionWorkflowDefinition
		return ret
	}
	return o.CoreWorkflows
}

// GetCoreWorkflowsOk returns a tuple with the CoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetCoreWorkflowsOk() ([]WorkflowActionWorkflowDefinition, bool) {
	if o == nil || o.CoreWorkflows == nil {
		return nil, false
	}
	return o.CoreWorkflows, true
}

// HasCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasCoreWorkflows() bool {
	if o != nil && o.CoreWorkflows != nil {
		return true
	}

	return false
}

// SetCoreWorkflows gets a reference to the given []WorkflowActionWorkflowDefinition and assigns it to the CoreWorkflows field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetCoreWorkflows(v []WorkflowActionWorkflowDefinition) {
	o.CoreWorkflows = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetInputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || o.OutputParameters == nil {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasOutputParameters() bool {
	if o != nil && o.OutputParameters != nil {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetPeriodicity returns the Periodicity field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetPeriodicity() int64 {
	if o == nil || o.Periodicity == nil {
		var ret int64
		return ret
	}
	return *o.Periodicity
}

// GetPeriodicityOk returns a tuple with the Periodicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetPeriodicityOk() (*int64, bool) {
	if o == nil || o.Periodicity == nil {
		return nil, false
	}
	return o.Periodicity, true
}

// HasPeriodicity returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasPeriodicity() bool {
	if o != nil && o.Periodicity != nil {
		return true
	}

	return false
}

// SetPeriodicity gets a reference to the given int64 and assigns it to the Periodicity field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetPeriodicity(v int64) {
	o.Periodicity = &v
}

// GetPostCoreWorkflows returns the PostCoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetPostCoreWorkflows() []WorkflowActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowActionWorkflowDefinition
		return ret
	}
	return o.PostCoreWorkflows
}

// GetPostCoreWorkflowsOk returns a tuple with the PostCoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetPostCoreWorkflowsOk() ([]WorkflowActionWorkflowDefinition, bool) {
	if o == nil || o.PostCoreWorkflows == nil {
		return nil, false
	}
	return o.PostCoreWorkflows, true
}

// HasPostCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasPostCoreWorkflows() bool {
	if o != nil && o.PostCoreWorkflows != nil {
		return true
	}

	return false
}

// SetPostCoreWorkflows gets a reference to the given []WorkflowActionWorkflowDefinition and assigns it to the PostCoreWorkflows field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetPostCoreWorkflows(v []WorkflowActionWorkflowDefinition) {
	o.PostCoreWorkflows = v
}

// GetPreCoreWorkflows returns the PreCoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetPreCoreWorkflows() []WorkflowActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowActionWorkflowDefinition
		return ret
	}
	return o.PreCoreWorkflows
}

// GetPreCoreWorkflowsOk returns a tuple with the PreCoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetPreCoreWorkflowsOk() ([]WorkflowActionWorkflowDefinition, bool) {
	if o == nil || o.PreCoreWorkflows == nil {
		return nil, false
	}
	return o.PreCoreWorkflows, true
}

// HasPreCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasPreCoreWorkflows() bool {
	if o != nil && o.PreCoreWorkflows != nil {
		return true
	}

	return false
}

// SetPreCoreWorkflows gets a reference to the given []WorkflowActionWorkflowDefinition and assigns it to the PreCoreWorkflows field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetPreCoreWorkflows(v []WorkflowActionWorkflowDefinition) {
	o.PreCoreWorkflows = v
}

// GetStopWorkflows returns the StopWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetStopWorkflows() []WorkflowActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowActionWorkflowDefinition
		return ret
	}
	return o.StopWorkflows
}

// GetStopWorkflowsOk returns a tuple with the StopWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetStopWorkflowsOk() ([]WorkflowActionWorkflowDefinition, bool) {
	if o == nil || o.StopWorkflows == nil {
		return nil, false
	}
	return o.StopWorkflows, true
}

// HasStopWorkflows returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasStopWorkflows() bool {
	if o != nil && o.StopWorkflows != nil {
		return true
	}

	return false
}

// SetStopWorkflows gets a reference to the given []WorkflowActionWorkflowDefinition and assigns it to the StopWorkflows field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetStopWorkflows(v []WorkflowActionWorkflowDefinition) {
	o.StopWorkflows = v
}

// GetUpgradedMoid returns the UpgradedMoid field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetUpgradedMoid() string {
	if o == nil || o.UpgradedMoid == nil {
		var ret string
		return ret
	}
	return *o.UpgradedMoid
}

// GetUpgradedMoidOk returns a tuple with the UpgradedMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetUpgradedMoidOk() (*string, bool) {
	if o == nil || o.UpgradedMoid == nil {
		return nil, false
	}
	return o.UpgradedMoid, true
}

// HasUpgradedMoid returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasUpgradedMoid() bool {
	if o != nil && o.UpgradedMoid != nil {
		return true
	}

	return false
}

// SetUpgradedMoid gets a reference to the given string and assigns it to the UpgradedMoid field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetUpgradedMoid(v string) {
	o.UpgradedMoid = &v
}

// GetValidationInformation returns the ValidationInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation {
	if o == nil || o.ValidationInformation.Get() == nil {
		var ret WorkflowValidationInformation
		return ret
	}
	return *o.ValidationInformation.Get()
}

// GetValidationInformationOk returns a tuple with the ValidationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationInformation.Get(), o.ValidationInformation.IsSet()
}

// HasValidationInformation returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasValidationInformation() bool {
	if o != nil && o.ValidationInformation.IsSet() {
		return true
	}

	return false
}

// SetValidationInformation gets a reference to the given NullableWorkflowValidationInformation and assigns it to the ValidationInformation field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation) {
	o.ValidationInformation.Set(&v)
}

// SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil
func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationInformationNil() {
	o.ValidationInformation.Set(nil)
}

// UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
func (o *WorkflowSolutionActionDefinitionAllOf) UnsetValidationInformation() {
	o.ValidationInformation.Unset()
}

// GetValidationWorkflows returns the ValidationWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationWorkflows() []WorkflowActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowActionWorkflowDefinition
		return ret
	}
	return o.ValidationWorkflows
}

// GetValidationWorkflowsOk returns a tuple with the ValidationWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowSolutionActionDefinitionAllOf) GetValidationWorkflowsOk() ([]WorkflowActionWorkflowDefinition, bool) {
	if o == nil || o.ValidationWorkflows == nil {
		return nil, false
	}
	return o.ValidationWorkflows, true
}

// HasValidationWorkflows returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasValidationWorkflows() bool {
	if o != nil && o.ValidationWorkflows != nil {
		return true
	}

	return false
}

// SetValidationWorkflows gets a reference to the given []WorkflowActionWorkflowDefinition and assigns it to the ValidationWorkflows field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetValidationWorkflows(v []WorkflowActionWorkflowDefinition) {
	o.ValidationWorkflows = v
}

// GetSolutionDefinition returns the SolutionDefinition field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetSolutionDefinition() WorkflowSolutionDefinitionRelationship {
	if o == nil || o.SolutionDefinition == nil {
		var ret WorkflowSolutionDefinitionRelationship
		return ret
	}
	return *o.SolutionDefinition
}

// GetSolutionDefinitionOk returns a tuple with the SolutionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetSolutionDefinitionOk() (*WorkflowSolutionDefinitionRelationship, bool) {
	if o == nil || o.SolutionDefinition == nil {
		return nil, false
	}
	return o.SolutionDefinition, true
}

// HasSolutionDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasSolutionDefinition() bool {
	if o != nil && o.SolutionDefinition != nil {
		return true
	}

	return false
}

// SetSolutionDefinition gets a reference to the given WorkflowSolutionDefinitionRelationship and assigns it to the SolutionDefinition field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetSolutionDefinition(v WorkflowSolutionDefinitionRelationship) {
	o.SolutionDefinition = &v
}

// GetWorkflowDefinition returns the WorkflowDefinition field value if set, zero value otherwise.
func (o *WorkflowSolutionActionDefinitionAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship {
	if o == nil || o.WorkflowDefinition == nil {
		var ret WorkflowWorkflowDefinitionRelationship
		return ret
	}
	return *o.WorkflowDefinition
}

// GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool) {
	if o == nil || o.WorkflowDefinition == nil {
		return nil, false
	}
	return o.WorkflowDefinition, true
}

// HasWorkflowDefinition returns a boolean if a field has been set.
func (o *WorkflowSolutionActionDefinitionAllOf) HasWorkflowDefinition() bool {
	if o != nil && o.WorkflowDefinition != nil {
		return true
	}

	return false
}

// SetWorkflowDefinition gets a reference to the given WorkflowWorkflowDefinitionRelationship and assigns it to the WorkflowDefinition field.
func (o *WorkflowSolutionActionDefinitionAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship) {
	o.WorkflowDefinition = &v
}

func (o WorkflowSolutionActionDefinitionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActionType != nil {
		toSerialize["ActionType"] = o.ActionType
	}
	if o.AllowedInstanceStates != nil {
		toSerialize["AllowedInstanceStates"] = o.AllowedInstanceStates
	}
	if o.CoreWorkflows != nil {
		toSerialize["CoreWorkflows"] = o.CoreWorkflows
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OutputParameters != nil {
		toSerialize["OutputParameters"] = o.OutputParameters
	}
	if o.Periodicity != nil {
		toSerialize["Periodicity"] = o.Periodicity
	}
	if o.PostCoreWorkflows != nil {
		toSerialize["PostCoreWorkflows"] = o.PostCoreWorkflows
	}
	if o.PreCoreWorkflows != nil {
		toSerialize["PreCoreWorkflows"] = o.PreCoreWorkflows
	}
	if o.StopWorkflows != nil {
		toSerialize["StopWorkflows"] = o.StopWorkflows
	}
	if o.UpgradedMoid != nil {
		toSerialize["UpgradedMoid"] = o.UpgradedMoid
	}
	if o.ValidationInformation.IsSet() {
		toSerialize["ValidationInformation"] = o.ValidationInformation.Get()
	}
	if o.ValidationWorkflows != nil {
		toSerialize["ValidationWorkflows"] = o.ValidationWorkflows
	}
	if o.SolutionDefinition != nil {
		toSerialize["SolutionDefinition"] = o.SolutionDefinition
	}
	if o.WorkflowDefinition != nil {
		toSerialize["WorkflowDefinition"] = o.WorkflowDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowSolutionActionDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowSolutionActionDefinitionAllOf := _WorkflowSolutionActionDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowSolutionActionDefinitionAllOf); err == nil {
		*o = WorkflowSolutionActionDefinitionAllOf(varWorkflowSolutionActionDefinitionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionType")
		delete(additionalProperties, "AllowedInstanceStates")
		delete(additionalProperties, "CoreWorkflows")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OutputParameters")
		delete(additionalProperties, "Periodicity")
		delete(additionalProperties, "PostCoreWorkflows")
		delete(additionalProperties, "PreCoreWorkflows")
		delete(additionalProperties, "StopWorkflows")
		delete(additionalProperties, "UpgradedMoid")
		delete(additionalProperties, "ValidationInformation")
		delete(additionalProperties, "ValidationWorkflows")
		delete(additionalProperties, "SolutionDefinition")
		delete(additionalProperties, "WorkflowDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowSolutionActionDefinitionAllOf struct {
	value *WorkflowSolutionActionDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowSolutionActionDefinitionAllOf) Get() *WorkflowSolutionActionDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowSolutionActionDefinitionAllOf) Set(val *WorkflowSolutionActionDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSolutionActionDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSolutionActionDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSolutionActionDefinitionAllOf(val *WorkflowSolutionActionDefinitionAllOf) *NullableWorkflowSolutionActionDefinitionAllOf {
	return &NullableWorkflowSolutionActionDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowSolutionActionDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSolutionActionDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
