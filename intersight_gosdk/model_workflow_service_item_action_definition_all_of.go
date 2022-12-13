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

// WorkflowServiceItemActionDefinitionAllOf Definition of the list of properties defined in 'workflow.ServiceItemActionDefinition', excluding properties defined in parent classes.
type WorkflowServiceItemActionDefinitionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Type of actionDefinition which decides on how to trigger the action. * `External` - External actions definition can be triggered by enduser to perform actions on the service item. Once action is completed successfully (eg. create/deploy), user cannot re-trigger that action again. * `Internal` - Internal action definition is used to trigger periodic actions on the service item instance. * `Repetitive` - Repetitive action definition is an external action that can be triggered by enduser to perform repetitive actions (eg. Edit/Update/Perform health check) on the created service item.
	ActionType            *string                                       `json:"ActionType,omitempty"`
	AllowedInstanceStates []string                                      `json:"AllowedInstanceStates,omitempty"`
	CoreWorkflows         []WorkflowServiceItemActionWorkflowDefinition `json:"CoreWorkflows,omitempty"`
	// The description for this action which provides information on what are the pre-requisites to use this action on the service item and what features are supported by this action.
	Description     *string                `json:"Description,omitempty"`
	InputDefinition []WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	// A user friendly short name to identify the action. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// The name for this action definition. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). Name of the action must be unique within a service item definition.
	Name *string `json:"Name,omitempty"`
	// The output mappings from workflows in the action definition to the service item output definition. Any output from core or post-core workflow can be mapped to service item output definition. The output can be referred using the name of the workflow definition and the output name in the following format '${<ServiceItemActionWorkflowDefinition.Name>.output.<outputName>'.
	OutputParameters interface{} `json:"OutputParameters,omitempty"`
	// Value in seconds to specify the periodicity of the workflows. A zero value indicate the workflow will not execute periodically. A non-zero value indicate, the workflow will be executed periodically with this periodicity.
	Periodicity       *int64                                        `json:"Periodicity,omitempty"`
	PostCoreWorkflows []WorkflowServiceItemActionWorkflowDefinition `json:"PostCoreWorkflows,omitempty"`
	PreCoreWorkflows  []WorkflowServiceItemActionWorkflowDefinition `json:"PreCoreWorkflows,omitempty"`
	// The flag to indicate that action is restricted on a Private Virtual Appliance.
	RestrictOnPrivateAppliance *bool                                         `json:"RestrictOnPrivateAppliance,omitempty"`
	StopWorkflows              []WorkflowServiceItemActionWorkflowDefinition `json:"StopWorkflows,omitempty"`
	ValidationInformation      NullableWorkflowValidationInformation         `json:"ValidationInformation,omitempty"`
	ValidationWorkflows        []WorkflowServiceItemActionWorkflowDefinition `json:"ValidationWorkflows,omitempty"`
	ServiceItemDefinition      *WorkflowServiceItemDefinitionRelationship    `json:"ServiceItemDefinition,omitempty"`
	WorkflowDefinition         *WorkflowWorkflowDefinitionRelationship       `json:"WorkflowDefinition,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _WorkflowServiceItemActionDefinitionAllOf WorkflowServiceItemActionDefinitionAllOf

// NewWorkflowServiceItemActionDefinitionAllOf instantiates a new WorkflowServiceItemActionDefinitionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemActionDefinitionAllOf(classId string, objectType string) *WorkflowServiceItemActionDefinitionAllOf {
	this := WorkflowServiceItemActionDefinitionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var actionType string = "External"
	this.ActionType = &actionType
	var restrictOnPrivateAppliance bool = false
	this.RestrictOnPrivateAppliance = &restrictOnPrivateAppliance
	return &this
}

// NewWorkflowServiceItemActionDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemActionDefinitionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemActionDefinitionAllOfWithDefaults() *WorkflowServiceItemActionDefinitionAllOf {
	this := WorkflowServiceItemActionDefinitionAllOf{}
	var classId string = "workflow.ServiceItemActionDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemActionDefinition"
	this.ObjectType = objectType
	var actionType string = "External"
	this.ActionType = &actionType
	var restrictOnPrivateAppliance bool = false
	this.RestrictOnPrivateAppliance = &restrictOnPrivateAppliance
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemActionDefinitionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemActionDefinitionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemActionDefinitionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemActionDefinitionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetActionType() string {
	if o == nil || o.ActionType == nil {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetActionTypeOk() (*string, bool) {
	if o == nil || o.ActionType == nil {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasActionType() bool {
	if o != nil && o.ActionType != nil {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetActionType(v string) {
	o.ActionType = &v
}

// GetAllowedInstanceStates returns the AllowedInstanceStates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetAllowedInstanceStates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedInstanceStates
}

// GetAllowedInstanceStatesOk returns a tuple with the AllowedInstanceStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetAllowedInstanceStatesOk() ([]string, bool) {
	if o == nil || o.AllowedInstanceStates == nil {
		return nil, false
	}
	return o.AllowedInstanceStates, true
}

// HasAllowedInstanceStates returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasAllowedInstanceStates() bool {
	if o != nil && o.AllowedInstanceStates != nil {
		return true
	}

	return false
}

// SetAllowedInstanceStates gets a reference to the given []string and assigns it to the AllowedInstanceStates field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetAllowedInstanceStates(v []string) {
	o.AllowedInstanceStates = v
}

// GetCoreWorkflows returns the CoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return o.CoreWorkflows
}

// GetCoreWorkflowsOk returns a tuple with the CoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetCoreWorkflowsOk() ([]WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil || o.CoreWorkflows == nil {
		return nil, false
	}
	return o.CoreWorkflows, true
}

// HasCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasCoreWorkflows() bool {
	if o != nil && o.CoreWorkflows != nil {
		return true
	}

	return false
}

// SetCoreWorkflows gets a reference to the given []WorkflowServiceItemActionWorkflowDefinition and assigns it to the CoreWorkflows field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition) {
	o.CoreWorkflows = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetInputDefinitionOk() ([]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetName(v string) {
	o.Name = &v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || o.OutputParameters == nil {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasOutputParameters() bool {
	if o != nil && o.OutputParameters != nil {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetPeriodicity returns the Periodicity field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPeriodicity() int64 {
	if o == nil || o.Periodicity == nil {
		var ret int64
		return ret
	}
	return *o.Periodicity
}

// GetPeriodicityOk returns a tuple with the Periodicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPeriodicityOk() (*int64, bool) {
	if o == nil || o.Periodicity == nil {
		return nil, false
	}
	return o.Periodicity, true
}

// HasPeriodicity returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasPeriodicity() bool {
	if o != nil && o.Periodicity != nil {
		return true
	}

	return false
}

// SetPeriodicity gets a reference to the given int64 and assigns it to the Periodicity field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetPeriodicity(v int64) {
	o.Periodicity = &v
}

// GetPostCoreWorkflows returns the PostCoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPostCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return o.PostCoreWorkflows
}

// GetPostCoreWorkflowsOk returns a tuple with the PostCoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPostCoreWorkflowsOk() ([]WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil || o.PostCoreWorkflows == nil {
		return nil, false
	}
	return o.PostCoreWorkflows, true
}

// HasPostCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasPostCoreWorkflows() bool {
	if o != nil && o.PostCoreWorkflows != nil {
		return true
	}

	return false
}

// SetPostCoreWorkflows gets a reference to the given []WorkflowServiceItemActionWorkflowDefinition and assigns it to the PostCoreWorkflows field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetPostCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition) {
	o.PostCoreWorkflows = v
}

// GetPreCoreWorkflows returns the PreCoreWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPreCoreWorkflows() []WorkflowServiceItemActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return o.PreCoreWorkflows
}

// GetPreCoreWorkflowsOk returns a tuple with the PreCoreWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetPreCoreWorkflowsOk() ([]WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil || o.PreCoreWorkflows == nil {
		return nil, false
	}
	return o.PreCoreWorkflows, true
}

// HasPreCoreWorkflows returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasPreCoreWorkflows() bool {
	if o != nil && o.PreCoreWorkflows != nil {
		return true
	}

	return false
}

// SetPreCoreWorkflows gets a reference to the given []WorkflowServiceItemActionWorkflowDefinition and assigns it to the PreCoreWorkflows field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetPreCoreWorkflows(v []WorkflowServiceItemActionWorkflowDefinition) {
	o.PreCoreWorkflows = v
}

// GetRestrictOnPrivateAppliance returns the RestrictOnPrivateAppliance field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetRestrictOnPrivateAppliance() bool {
	if o == nil || o.RestrictOnPrivateAppliance == nil {
		var ret bool
		return ret
	}
	return *o.RestrictOnPrivateAppliance
}

// GetRestrictOnPrivateApplianceOk returns a tuple with the RestrictOnPrivateAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetRestrictOnPrivateApplianceOk() (*bool, bool) {
	if o == nil || o.RestrictOnPrivateAppliance == nil {
		return nil, false
	}
	return o.RestrictOnPrivateAppliance, true
}

// HasRestrictOnPrivateAppliance returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasRestrictOnPrivateAppliance() bool {
	if o != nil && o.RestrictOnPrivateAppliance != nil {
		return true
	}

	return false
}

// SetRestrictOnPrivateAppliance gets a reference to the given bool and assigns it to the RestrictOnPrivateAppliance field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetRestrictOnPrivateAppliance(v bool) {
	o.RestrictOnPrivateAppliance = &v
}

// GetStopWorkflows returns the StopWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetStopWorkflows() []WorkflowServiceItemActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return o.StopWorkflows
}

// GetStopWorkflowsOk returns a tuple with the StopWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetStopWorkflowsOk() ([]WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil || o.StopWorkflows == nil {
		return nil, false
	}
	return o.StopWorkflows, true
}

// HasStopWorkflows returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasStopWorkflows() bool {
	if o != nil && o.StopWorkflows != nil {
		return true
	}

	return false
}

// SetStopWorkflows gets a reference to the given []WorkflowServiceItemActionWorkflowDefinition and assigns it to the StopWorkflows field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetStopWorkflows(v []WorkflowServiceItemActionWorkflowDefinition) {
	o.StopWorkflows = v
}

// GetValidationInformation returns the ValidationInformation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation {
	if o == nil || o.ValidationInformation.Get() == nil {
		var ret WorkflowValidationInformation
		return ret
	}
	return *o.ValidationInformation.Get()
}

// GetValidationInformationOk returns a tuple with the ValidationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationInformation.Get(), o.ValidationInformation.IsSet()
}

// HasValidationInformation returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasValidationInformation() bool {
	if o != nil && o.ValidationInformation.IsSet() {
		return true
	}

	return false
}

// SetValidationInformation gets a reference to the given NullableWorkflowValidationInformation and assigns it to the ValidationInformation field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation) {
	o.ValidationInformation.Set(&v)
}

// SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil
func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationInformationNil() {
	o.ValidationInformation.Set(nil)
}

// UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
func (o *WorkflowServiceItemActionDefinitionAllOf) UnsetValidationInformation() {
	o.ValidationInformation.Unset()
}

// GetValidationWorkflows returns the ValidationWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationWorkflows() []WorkflowServiceItemActionWorkflowDefinition {
	if o == nil {
		var ret []WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return o.ValidationWorkflows
}

// GetValidationWorkflowsOk returns a tuple with the ValidationWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemActionDefinitionAllOf) GetValidationWorkflowsOk() ([]WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil || o.ValidationWorkflows == nil {
		return nil, false
	}
	return o.ValidationWorkflows, true
}

// HasValidationWorkflows returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasValidationWorkflows() bool {
	if o != nil && o.ValidationWorkflows != nil {
		return true
	}

	return false
}

// SetValidationWorkflows gets a reference to the given []WorkflowServiceItemActionWorkflowDefinition and assigns it to the ValidationWorkflows field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetValidationWorkflows(v []WorkflowServiceItemActionWorkflowDefinition) {
	o.ValidationWorkflows = v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || o.ServiceItemDefinition == nil {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemDefinition == nil {
		return nil, false
	}
	return o.ServiceItemDefinition, true
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given WorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition = &v
}

// GetWorkflowDefinition returns the WorkflowDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetWorkflowDefinition() WorkflowWorkflowDefinitionRelationship {
	if o == nil || o.WorkflowDefinition == nil {
		var ret WorkflowWorkflowDefinitionRelationship
		return ret
	}
	return *o.WorkflowDefinition
}

// GetWorkflowDefinitionOk returns a tuple with the WorkflowDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) GetWorkflowDefinitionOk() (*WorkflowWorkflowDefinitionRelationship, bool) {
	if o == nil || o.WorkflowDefinition == nil {
		return nil, false
	}
	return o.WorkflowDefinition, true
}

// HasWorkflowDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemActionDefinitionAllOf) HasWorkflowDefinition() bool {
	if o != nil && o.WorkflowDefinition != nil {
		return true
	}

	return false
}

// SetWorkflowDefinition gets a reference to the given WorkflowWorkflowDefinitionRelationship and assigns it to the WorkflowDefinition field.
func (o *WorkflowServiceItemActionDefinitionAllOf) SetWorkflowDefinition(v WorkflowWorkflowDefinitionRelationship) {
	o.WorkflowDefinition = &v
}

func (o WorkflowServiceItemActionDefinitionAllOf) MarshalJSON() ([]byte, error) {
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
	if o.RestrictOnPrivateAppliance != nil {
		toSerialize["RestrictOnPrivateAppliance"] = o.RestrictOnPrivateAppliance
	}
	if o.StopWorkflows != nil {
		toSerialize["StopWorkflows"] = o.StopWorkflows
	}
	if o.ValidationInformation.IsSet() {
		toSerialize["ValidationInformation"] = o.ValidationInformation.Get()
	}
	if o.ValidationWorkflows != nil {
		toSerialize["ValidationWorkflows"] = o.ValidationWorkflows
	}
	if o.ServiceItemDefinition != nil {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition
	}
	if o.WorkflowDefinition != nil {
		toSerialize["WorkflowDefinition"] = o.WorkflowDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemActionDefinitionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowServiceItemActionDefinitionAllOf := _WorkflowServiceItemActionDefinitionAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowServiceItemActionDefinitionAllOf); err == nil {
		*o = WorkflowServiceItemActionDefinitionAllOf(varWorkflowServiceItemActionDefinitionAllOf)
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
		delete(additionalProperties, "RestrictOnPrivateAppliance")
		delete(additionalProperties, "StopWorkflows")
		delete(additionalProperties, "ValidationInformation")
		delete(additionalProperties, "ValidationWorkflows")
		delete(additionalProperties, "ServiceItemDefinition")
		delete(additionalProperties, "WorkflowDefinition")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowServiceItemActionDefinitionAllOf struct {
	value *WorkflowServiceItemActionDefinitionAllOf
	isSet bool
}

func (v NullableWorkflowServiceItemActionDefinitionAllOf) Get() *WorkflowServiceItemActionDefinitionAllOf {
	return v.value
}

func (v *NullableWorkflowServiceItemActionDefinitionAllOf) Set(val *WorkflowServiceItemActionDefinitionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemActionDefinitionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemActionDefinitionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemActionDefinitionAllOf(val *WorkflowServiceItemActionDefinitionAllOf) *NullableWorkflowServiceItemActionDefinitionAllOf {
	return &NullableWorkflowServiceItemActionDefinitionAllOf{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemActionDefinitionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemActionDefinitionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
