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
)

// WorkflowWorkflowNotificationAllOf Definition of the list of properties defined in 'workflow.WorkflowNotification', excluding properties defined in parent classes.
type WorkflowWorkflowNotificationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The correlationId of the workflow.
	CorrelationId *string `json:"CorrelationId,omitempty"`
	// The end time of the workflow.
	EndTime *string `json:"EndTime,omitempty"`
	// The event of the workflow.
	Event *string `json:"Event,omitempty"`
	// The execution time of the workflow.
	ExecutionTime *float32 `json:"ExecutionTime,omitempty"`
	// The reference task names of the failed tasks.
	FailedReferenceTaskNames *string `json:"FailedReferenceTaskNames,omitempty"`
	// The input of the workflow.
	Input *string `json:"Input,omitempty"`
	// The output of the workflow.
	Output *string `json:"Output,omitempty"`
	// The reason for incompletion status of the workflow.
	ReasonForIncompletion *string `json:"ReasonForIncompletion,omitempty"`
	// The start time of the workflow.
	StartTime *string `json:"StartTime,omitempty"`
	// The final status of the workflow.
	Status *string `json:"Status,omitempty"`
	// The last update time of the workflow.
	UpdateTime *string `json:"UpdateTime,omitempty"`
	// The version of the workflow.
	Version *int64 `json:"Version,omitempty"`
	// The unique id of the workflow.
	WorkflowId *string `json:"WorkflowId,omitempty"`
	// The type of the workflow.
	WorkflowType         *string `json:"WorkflowType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowNotificationAllOf WorkflowWorkflowNotificationAllOf

// NewWorkflowWorkflowNotificationAllOf instantiates a new WorkflowWorkflowNotificationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowNotificationAllOf(classId string, objectType string) *WorkflowWorkflowNotificationAllOf {
	this := WorkflowWorkflowNotificationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWorkflowNotificationAllOfWithDefaults instantiates a new WorkflowWorkflowNotificationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowNotificationAllOfWithDefaults() *WorkflowWorkflowNotificationAllOf {
	this := WorkflowWorkflowNotificationAllOf{}
	var classId string = "workflow.WorkflowNotification"
	this.ClassId = classId
	var objectType string = "workflow.WorkflowNotification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWorkflowNotificationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWorkflowNotificationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWorkflowNotificationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWorkflowNotificationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetCorrelationId() string {
	if o == nil || o.CorrelationId == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetCorrelationIdOk() (*string, bool) {
	if o == nil || o.CorrelationId == nil {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasCorrelationId() bool {
	if o != nil && o.CorrelationId != nil {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *WorkflowWorkflowNotificationAllOf) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *WorkflowWorkflowNotificationAllOf) SetEndTime(v string) {
	o.EndTime = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *WorkflowWorkflowNotificationAllOf) SetEvent(v string) {
	o.Event = &v
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetExecutionTime() float32 {
	if o == nil || o.ExecutionTime == nil {
		var ret float32
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetExecutionTimeOk() (*float32, bool) {
	if o == nil || o.ExecutionTime == nil {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasExecutionTime() bool {
	if o != nil && o.ExecutionTime != nil {
		return true
	}

	return false
}

// SetExecutionTime gets a reference to the given float32 and assigns it to the ExecutionTime field.
func (o *WorkflowWorkflowNotificationAllOf) SetExecutionTime(v float32) {
	o.ExecutionTime = &v
}

// GetFailedReferenceTaskNames returns the FailedReferenceTaskNames field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetFailedReferenceTaskNames() string {
	if o == nil || o.FailedReferenceTaskNames == nil {
		var ret string
		return ret
	}
	return *o.FailedReferenceTaskNames
}

// GetFailedReferenceTaskNamesOk returns a tuple with the FailedReferenceTaskNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetFailedReferenceTaskNamesOk() (*string, bool) {
	if o == nil || o.FailedReferenceTaskNames == nil {
		return nil, false
	}
	return o.FailedReferenceTaskNames, true
}

// HasFailedReferenceTaskNames returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasFailedReferenceTaskNames() bool {
	if o != nil && o.FailedReferenceTaskNames != nil {
		return true
	}

	return false
}

// SetFailedReferenceTaskNames gets a reference to the given string and assigns it to the FailedReferenceTaskNames field.
func (o *WorkflowWorkflowNotificationAllOf) SetFailedReferenceTaskNames(v string) {
	o.FailedReferenceTaskNames = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *WorkflowWorkflowNotificationAllOf) SetInput(v string) {
	o.Input = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *WorkflowWorkflowNotificationAllOf) SetOutput(v string) {
	o.Output = &v
}

// GetReasonForIncompletion returns the ReasonForIncompletion field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetReasonForIncompletion() string {
	if o == nil || o.ReasonForIncompletion == nil {
		var ret string
		return ret
	}
	return *o.ReasonForIncompletion
}

// GetReasonForIncompletionOk returns a tuple with the ReasonForIncompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetReasonForIncompletionOk() (*string, bool) {
	if o == nil || o.ReasonForIncompletion == nil {
		return nil, false
	}
	return o.ReasonForIncompletion, true
}

// HasReasonForIncompletion returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasReasonForIncompletion() bool {
	if o != nil && o.ReasonForIncompletion != nil {
		return true
	}

	return false
}

// SetReasonForIncompletion gets a reference to the given string and assigns it to the ReasonForIncompletion field.
func (o *WorkflowWorkflowNotificationAllOf) SetReasonForIncompletion(v string) {
	o.ReasonForIncompletion = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *WorkflowWorkflowNotificationAllOf) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowWorkflowNotificationAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetUpdateTime() string {
	if o == nil || o.UpdateTime == nil {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetUpdateTimeOk() (*string, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *WorkflowWorkflowNotificationAllOf) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowWorkflowNotificationAllOf) SetVersion(v int64) {
	o.Version = &v
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *WorkflowWorkflowNotificationAllOf) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetWorkflowType returns the WorkflowType field value if set, zero value otherwise.
func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowType() string {
	if o == nil || o.WorkflowType == nil {
		var ret string
		return ret
	}
	return *o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowNotificationAllOf) GetWorkflowTypeOk() (*string, bool) {
	if o == nil || o.WorkflowType == nil {
		return nil, false
	}
	return o.WorkflowType, true
}

// HasWorkflowType returns a boolean if a field has been set.
func (o *WorkflowWorkflowNotificationAllOf) HasWorkflowType() bool {
	if o != nil && o.WorkflowType != nil {
		return true
	}

	return false
}

// SetWorkflowType gets a reference to the given string and assigns it to the WorkflowType field.
func (o *WorkflowWorkflowNotificationAllOf) SetWorkflowType(v string) {
	o.WorkflowType = &v
}

func (o WorkflowWorkflowNotificationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CorrelationId != nil {
		toSerialize["CorrelationId"] = o.CorrelationId
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Event != nil {
		toSerialize["Event"] = o.Event
	}
	if o.ExecutionTime != nil {
		toSerialize["ExecutionTime"] = o.ExecutionTime
	}
	if o.FailedReferenceTaskNames != nil {
		toSerialize["FailedReferenceTaskNames"] = o.FailedReferenceTaskNames
	}
	if o.Input != nil {
		toSerialize["Input"] = o.Input
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.ReasonForIncompletion != nil {
		toSerialize["ReasonForIncompletion"] = o.ReasonForIncompletion
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UpdateTime != nil {
		toSerialize["UpdateTime"] = o.UpdateTime
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.WorkflowId != nil {
		toSerialize["WorkflowId"] = o.WorkflowId
	}
	if o.WorkflowType != nil {
		toSerialize["WorkflowType"] = o.WorkflowType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowNotificationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWorkflowNotificationAllOf := _WorkflowWorkflowNotificationAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWorkflowNotificationAllOf); err == nil {
		*o = WorkflowWorkflowNotificationAllOf(varWorkflowWorkflowNotificationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CorrelationId")
		delete(additionalProperties, "EndTime")
		delete(additionalProperties, "Event")
		delete(additionalProperties, "ExecutionTime")
		delete(additionalProperties, "FailedReferenceTaskNames")
		delete(additionalProperties, "Input")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "ReasonForIncompletion")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "UpdateTime")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "WorkflowId")
		delete(additionalProperties, "WorkflowType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWorkflowNotificationAllOf struct {
	value *WorkflowWorkflowNotificationAllOf
	isSet bool
}

func (v NullableWorkflowWorkflowNotificationAllOf) Get() *WorkflowWorkflowNotificationAllOf {
	return v.value
}

func (v *NullableWorkflowWorkflowNotificationAllOf) Set(val *WorkflowWorkflowNotificationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowNotificationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowNotificationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowNotificationAllOf(val *WorkflowWorkflowNotificationAllOf) *NullableWorkflowWorkflowNotificationAllOf {
	return &NullableWorkflowWorkflowNotificationAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowNotificationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowNotificationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
