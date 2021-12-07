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
)

// WorkflowWorkflowInfoPropertiesAllOf Definition of the list of properties defined in 'workflow.WorkflowInfoProperties', excluding properties defined in parent classes.
type WorkflowWorkflowInfoPropertiesAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// When true, this workflow can be retried if has not been modified for more than a period of 2 weeks.
	Retryable *bool `json:"Retryable,omitempty"`
	// Status of rollback for this workflow instance. The rollback action of the workflow can be enabled, disabled, completed. * `Disabled` - Status of the rollback action when workflow is disabled for rollback. * `Enabled` - Status of the rollback action when workflow is enabled for rollback. * `Completed` - Status of the rollback action once workflow completes the rollback for all eligible tasks.
	RollbackAction       *string `json:"RollbackAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowInfoPropertiesAllOf WorkflowWorkflowInfoPropertiesAllOf

// NewWorkflowWorkflowInfoPropertiesAllOf instantiates a new WorkflowWorkflowInfoPropertiesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowInfoPropertiesAllOf(classId string, objectType string) *WorkflowWorkflowInfoPropertiesAllOf {
	this := WorkflowWorkflowInfoPropertiesAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var retryable bool = false
	this.Retryable = &retryable
	return &this
}

// NewWorkflowWorkflowInfoPropertiesAllOfWithDefaults instantiates a new WorkflowWorkflowInfoPropertiesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowInfoPropertiesAllOfWithDefaults() *WorkflowWorkflowInfoPropertiesAllOf {
	this := WorkflowWorkflowInfoPropertiesAllOf{}
	var classId string = "workflow.WorkflowInfoProperties"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowInfoProperties"
	this.ObjectType = objectType
	var retryable bool = false
	this.Retryable = &retryable
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowInfoPropertiesAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowInfoPropertiesAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRetryable returns the Retryable field value if set, zero value otherwise.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetRetryable() bool {
	if o == nil || o.Retryable == nil {
		var ret bool
		return ret
	}
	return *o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetRetryableOk() (*bool, bool) {
	if o == nil || o.Retryable == nil {
		return nil, false
	}
	return o.Retryable, true
}

// HasRetryable returns a boolean if a field has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) HasRetryable() bool {
	if o != nil && o.Retryable != nil {
		return true
	}

	return false
}

// SetRetryable gets a reference to the given bool and assigns it to the Retryable field.
func (o *WorkflowWorkflowInfoPropertiesAllOf) SetRetryable(v bool) {
	o.Retryable = &v
}

// GetRollbackAction returns the RollbackAction field value if set, zero value otherwise.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetRollbackAction() string {
	if o == nil || o.RollbackAction == nil {
		var ret string
		return ret
	}
	return *o.RollbackAction
}

// GetRollbackActionOk returns a tuple with the RollbackAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) GetRollbackActionOk() (*string, bool) {
	if o == nil || o.RollbackAction == nil {
		return nil, false
	}
	return o.RollbackAction, true
}

// HasRollbackAction returns a boolean if a field has been set.
func (o *WorkflowWorkflowInfoPropertiesAllOf) HasRollbackAction() bool {
	if o != nil && o.RollbackAction != nil {
		return true
	}

	return false
}

// SetRollbackAction gets a reference to the given string and assigns it to the RollbackAction field.
func (o *WorkflowWorkflowInfoPropertiesAllOf) SetRollbackAction(v string) {
	o.RollbackAction = &v
}

func (o WorkflowWorkflowInfoPropertiesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Retryable != nil {
		toSerialize["Retryable"] = o.Retryable
	}
	if o.RollbackAction != nil {
		toSerialize["RollbackAction"] = o.RollbackAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowInfoPropertiesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWorkflowInfoPropertiesAllOf := _WorkflowWorkflowInfoPropertiesAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWorkflowInfoPropertiesAllOf); err == nil {
		*o = WorkflowWorkflowInfoPropertiesAllOf(varWorkflowWorkflowInfoPropertiesAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Retryable")
		delete(additionalProperties, "RollbackAction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWorkflowInfoPropertiesAllOf struct {
	value *WorkflowWorkflowInfoPropertiesAllOf
	isSet bool
}

func (v NullableWorkflowWorkflowInfoPropertiesAllOf) Get() *WorkflowWorkflowInfoPropertiesAllOf {
	return v.value
}

func (v *NullableWorkflowWorkflowInfoPropertiesAllOf) Set(val *WorkflowWorkflowInfoPropertiesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowInfoPropertiesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowInfoPropertiesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowInfoPropertiesAllOf(val *WorkflowWorkflowInfoPropertiesAllOf) *NullableWorkflowWorkflowInfoPropertiesAllOf {
	return &NullableWorkflowWorkflowInfoPropertiesAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowInfoPropertiesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowInfoPropertiesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
