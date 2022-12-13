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

// WorkflowServiceItemHealthCheckDefinition Service item health check definition metadata.
type WorkflowServiceItemHealthCheckDefinition struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Category that the health check belongs to.
	Category *string `json:"Category,omitempty"`
	// Static information detailing the common cause for the health check failure. It also gives possible remediation actions that can be taken to remedy the health check failure.
	CommonCauseAndResolution *string `json:"CommonCauseAndResolution,omitempty"`
	// Description of the health check definition.
	Description *string `json:"Description,omitempty"`
	// Execution mode of the health check on service item. * `OnDemand` - Execute the health check on-demand. * `Periodic` - Execute the health check on a periodic basis.
	ExecutionMode       *string                                             `json:"ExecutionMode,omitempty"`
	HealthCheckWorkflow NullableWorkflowServiceItemActionWorkflowDefinition `json:"HealthCheckWorkflow,omitempty"`
	// Label for the health check definition that is displayed on UI.
	Label *string `json:"Label,omitempty"`
	// Name of the health check definition.
	Name                  *string                                    `json:"Name,omitempty"`
	ServiceItemDefinition *WorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemHealthCheckDefinition WorkflowServiceItemHealthCheckDefinition

// NewWorkflowServiceItemHealthCheckDefinition instantiates a new WorkflowServiceItemHealthCheckDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemHealthCheckDefinition(classId string, objectType string) *WorkflowServiceItemHealthCheckDefinition {
	this := WorkflowServiceItemHealthCheckDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// NewWorkflowServiceItemHealthCheckDefinitionWithDefaults instantiates a new WorkflowServiceItemHealthCheckDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemHealthCheckDefinitionWithDefaults() *WorkflowServiceItemHealthCheckDefinition {
	this := WorkflowServiceItemHealthCheckDefinition{}
	var classId string = "workflow.ServiceItemHealthCheckDefinition"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemHealthCheckDefinition"
	this.ObjectType = objectType
	var executionMode string = "OnDemand"
	this.ExecutionMode = &executionMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemHealthCheckDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemHealthCheckDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemHealthCheckDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemHealthCheckDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetCategory(v string) {
	o.Category = &v
}

// GetCommonCauseAndResolution returns the CommonCauseAndResolution field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetCommonCauseAndResolution() string {
	if o == nil || o.CommonCauseAndResolution == nil {
		var ret string
		return ret
	}
	return *o.CommonCauseAndResolution
}

// GetCommonCauseAndResolutionOk returns a tuple with the CommonCauseAndResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetCommonCauseAndResolutionOk() (*string, bool) {
	if o == nil || o.CommonCauseAndResolution == nil {
		return nil, false
	}
	return o.CommonCauseAndResolution, true
}

// HasCommonCauseAndResolution returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasCommonCauseAndResolution() bool {
	if o != nil && o.CommonCauseAndResolution != nil {
		return true
	}

	return false
}

// SetCommonCauseAndResolution gets a reference to the given string and assigns it to the CommonCauseAndResolution field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetCommonCauseAndResolution(v string) {
	o.CommonCauseAndResolution = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetExecutionMode returns the ExecutionMode field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetExecutionMode() string {
	if o == nil || o.ExecutionMode == nil {
		var ret string
		return ret
	}
	return *o.ExecutionMode
}

// GetExecutionModeOk returns a tuple with the ExecutionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetExecutionModeOk() (*string, bool) {
	if o == nil || o.ExecutionMode == nil {
		return nil, false
	}
	return o.ExecutionMode, true
}

// HasExecutionMode returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasExecutionMode() bool {
	if o != nil && o.ExecutionMode != nil {
		return true
	}

	return false
}

// SetExecutionMode gets a reference to the given string and assigns it to the ExecutionMode field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetExecutionMode(v string) {
	o.ExecutionMode = &v
}

// GetHealthCheckWorkflow returns the HealthCheckWorkflow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckDefinition) GetHealthCheckWorkflow() WorkflowServiceItemActionWorkflowDefinition {
	if o == nil || o.HealthCheckWorkflow.Get() == nil {
		var ret WorkflowServiceItemActionWorkflowDefinition
		return ret
	}
	return *o.HealthCheckWorkflow.Get()
}

// GetHealthCheckWorkflowOk returns a tuple with the HealthCheckWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckDefinition) GetHealthCheckWorkflowOk() (*WorkflowServiceItemActionWorkflowDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthCheckWorkflow.Get(), o.HealthCheckWorkflow.IsSet()
}

// HasHealthCheckWorkflow returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasHealthCheckWorkflow() bool {
	if o != nil && o.HealthCheckWorkflow.IsSet() {
		return true
	}

	return false
}

// SetHealthCheckWorkflow gets a reference to the given NullableWorkflowServiceItemActionWorkflowDefinition and assigns it to the HealthCheckWorkflow field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetHealthCheckWorkflow(v WorkflowServiceItemActionWorkflowDefinition) {
	o.HealthCheckWorkflow.Set(&v)
}

// SetHealthCheckWorkflowNil sets the value for HealthCheckWorkflow to be an explicit nil
func (o *WorkflowServiceItemHealthCheckDefinition) SetHealthCheckWorkflowNil() {
	o.HealthCheckWorkflow.Set(nil)
}

// UnsetHealthCheckWorkflow ensures that no value is present for HealthCheckWorkflow, not even an explicit nil
func (o *WorkflowServiceItemHealthCheckDefinition) UnsetHealthCheckWorkflow() {
	o.HealthCheckWorkflow.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetName(v string) {
	o.Name = &v
}

// GetServiceItemDefinition returns the ServiceItemDefinition field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckDefinition) GetServiceItemDefinition() WorkflowServiceItemDefinitionRelationship {
	if o == nil || o.ServiceItemDefinition == nil {
		var ret WorkflowServiceItemDefinitionRelationship
		return ret
	}
	return *o.ServiceItemDefinition
}

// GetServiceItemDefinitionOk returns a tuple with the ServiceItemDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) GetServiceItemDefinitionOk() (*WorkflowServiceItemDefinitionRelationship, bool) {
	if o == nil || o.ServiceItemDefinition == nil {
		return nil, false
	}
	return o.ServiceItemDefinition, true
}

// HasServiceItemDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckDefinition) HasServiceItemDefinition() bool {
	if o != nil && o.ServiceItemDefinition != nil {
		return true
	}

	return false
}

// SetServiceItemDefinition gets a reference to the given WorkflowServiceItemDefinitionRelationship and assigns it to the ServiceItemDefinition field.
func (o *WorkflowServiceItemHealthCheckDefinition) SetServiceItemDefinition(v WorkflowServiceItemDefinitionRelationship) {
	o.ServiceItemDefinition = &v
}

func (o WorkflowServiceItemHealthCheckDefinition) MarshalJSON() ([]byte, error) {
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
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.CommonCauseAndResolution != nil {
		toSerialize["CommonCauseAndResolution"] = o.CommonCauseAndResolution
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ExecutionMode != nil {
		toSerialize["ExecutionMode"] = o.ExecutionMode
	}
	if o.HealthCheckWorkflow.IsSet() {
		toSerialize["HealthCheckWorkflow"] = o.HealthCheckWorkflow.Get()
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServiceItemDefinition != nil {
		toSerialize["ServiceItemDefinition"] = o.ServiceItemDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowServiceItemHealthCheckDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Category that the health check belongs to.
		Category *string `json:"Category,omitempty"`
		// Static information detailing the common cause for the health check failure. It also gives possible remediation actions that can be taken to remedy the health check failure.
		CommonCauseAndResolution *string `json:"CommonCauseAndResolution,omitempty"`
		// Description of the health check definition.
		Description *string `json:"Description,omitempty"`
		// Execution mode of the health check on service item. * `OnDemand` - Execute the health check on-demand. * `Periodic` - Execute the health check on a periodic basis.
		ExecutionMode       *string                                             `json:"ExecutionMode,omitempty"`
		HealthCheckWorkflow NullableWorkflowServiceItemActionWorkflowDefinition `json:"HealthCheckWorkflow,omitempty"`
		// Label for the health check definition that is displayed on UI.
		Label *string `json:"Label,omitempty"`
		// Name of the health check definition.
		Name                  *string                                    `json:"Name,omitempty"`
		ServiceItemDefinition *WorkflowServiceItemDefinitionRelationship `json:"ServiceItemDefinition,omitempty"`
	}

	varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct := WorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemHealthCheckDefinition := _WorkflowServiceItemHealthCheckDefinition{}
		varWorkflowServiceItemHealthCheckDefinition.ClassId = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemHealthCheckDefinition.ObjectType = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemHealthCheckDefinition.Category = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.Category
		varWorkflowServiceItemHealthCheckDefinition.CommonCauseAndResolution = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.CommonCauseAndResolution
		varWorkflowServiceItemHealthCheckDefinition.Description = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.Description
		varWorkflowServiceItemHealthCheckDefinition.ExecutionMode = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.ExecutionMode
		varWorkflowServiceItemHealthCheckDefinition.HealthCheckWorkflow = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.HealthCheckWorkflow
		varWorkflowServiceItemHealthCheckDefinition.Label = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.Label
		varWorkflowServiceItemHealthCheckDefinition.Name = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.Name
		varWorkflowServiceItemHealthCheckDefinition.ServiceItemDefinition = varWorkflowServiceItemHealthCheckDefinitionWithoutEmbeddedStruct.ServiceItemDefinition
		*o = WorkflowServiceItemHealthCheckDefinition(varWorkflowServiceItemHealthCheckDefinition)
	} else {
		return err
	}

	varWorkflowServiceItemHealthCheckDefinition := _WorkflowServiceItemHealthCheckDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowServiceItemHealthCheckDefinition)
	if err == nil {
		o.MoBaseMo = varWorkflowServiceItemHealthCheckDefinition.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "CommonCauseAndResolution")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ExecutionMode")
		delete(additionalProperties, "HealthCheckWorkflow")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServiceItemDefinition")

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

type NullableWorkflowServiceItemHealthCheckDefinition struct {
	value *WorkflowServiceItemHealthCheckDefinition
	isSet bool
}

func (v NullableWorkflowServiceItemHealthCheckDefinition) Get() *WorkflowServiceItemHealthCheckDefinition {
	return v.value
}

func (v *NullableWorkflowServiceItemHealthCheckDefinition) Set(val *WorkflowServiceItemHealthCheckDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemHealthCheckDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemHealthCheckDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemHealthCheckDefinition(val *WorkflowServiceItemHealthCheckDefinition) *NullableWorkflowServiceItemHealthCheckDefinition {
	return &NullableWorkflowServiceItemHealthCheckDefinition{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemHealthCheckDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemHealthCheckDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
