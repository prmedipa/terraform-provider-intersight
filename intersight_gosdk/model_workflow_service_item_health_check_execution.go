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
	"time"
)

// checks if the WorkflowServiceItemHealthCheckExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowServiceItemHealthCheckExecution{}

// WorkflowServiceItemHealthCheckExecution Health check execution result for a health check definition on a service item.
type WorkflowServiceItemHealthCheckExecution struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                               `json:"ObjectType"`
	ErrorElements []ServiceitemHealthCheckErrorElement `json:"ErrorElements,omitempty"`
	// Error summary of a health check execution failure.
	ErrorSummary *string `json:"ErrorSummary,omitempty"`
	// Timestamp of the last passed health check execution.
	LastPassedTimestamp *time.Time `json:"LastPassedTimestamp,omitempty"`
	// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the service item.
	Result *string `json:"Result,omitempty"`
	// A brief summary of health check execution result.
	Summary *string `json:"Summary,omitempty"`
	// Status of the workflow that is executed as a part of health check execution.
	WorkflowStatus        *string                                                      `json:"WorkflowStatus,omitempty"`
	HealthCheckDefinition NullableWorkflowServiceItemHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
	ServiceItemInstance   NullableWorkflowServiceItemInstanceRelationship              `json:"ServiceItemInstance,omitempty"`
	WorkflowInfo          NullableWorkflowWorkflowInfoRelationship                     `json:"WorkflowInfo,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _WorkflowServiceItemHealthCheckExecution WorkflowServiceItemHealthCheckExecution

// NewWorkflowServiceItemHealthCheckExecution instantiates a new WorkflowServiceItemHealthCheckExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowServiceItemHealthCheckExecution(classId string, objectType string) *WorkflowServiceItemHealthCheckExecution {
	this := WorkflowServiceItemHealthCheckExecution{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowServiceItemHealthCheckExecutionWithDefaults instantiates a new WorkflowServiceItemHealthCheckExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowServiceItemHealthCheckExecutionWithDefaults() *WorkflowServiceItemHealthCheckExecution {
	this := WorkflowServiceItemHealthCheckExecution{}
	var classId string = "workflow.ServiceItemHealthCheckExecution"
	this.ClassId = classId
	var objectType string = "workflow.ServiceItemHealthCheckExecution"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowServiceItemHealthCheckExecution) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowServiceItemHealthCheckExecution) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.ServiceItemHealthCheckExecution" of the ClassId field.
func (o *WorkflowServiceItemHealthCheckExecution) GetDefaultClassId() interface{} {
	return "workflow.ServiceItemHealthCheckExecution"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowServiceItemHealthCheckExecution) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowServiceItemHealthCheckExecution) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.ServiceItemHealthCheckExecution" of the ObjectType field.
func (o *WorkflowServiceItemHealthCheckExecution) GetDefaultObjectType() interface{} {
	return "workflow.ServiceItemHealthCheckExecution"
}

// GetErrorElements returns the ErrorElements field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckExecution) GetErrorElements() []ServiceitemHealthCheckErrorElement {
	if o == nil {
		var ret []ServiceitemHealthCheckErrorElement
		return ret
	}
	return o.ErrorElements
}

// GetErrorElementsOk returns a tuple with the ErrorElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckExecution) GetErrorElementsOk() ([]ServiceitemHealthCheckErrorElement, bool) {
	if o == nil || IsNil(o.ErrorElements) {
		return nil, false
	}
	return o.ErrorElements, true
}

// HasErrorElements returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasErrorElements() bool {
	if o != nil && !IsNil(o.ErrorElements) {
		return true
	}

	return false
}

// SetErrorElements gets a reference to the given []ServiceitemHealthCheckErrorElement and assigns it to the ErrorElements field.
func (o *WorkflowServiceItemHealthCheckExecution) SetErrorElements(v []ServiceitemHealthCheckErrorElement) {
	o.ErrorElements = v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecution) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *WorkflowServiceItemHealthCheckExecution) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetLastPassedTimestamp returns the LastPassedTimestamp field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecution) GetLastPassedTimestamp() time.Time {
	if o == nil || IsNil(o.LastPassedTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.LastPassedTimestamp
}

// GetLastPassedTimestampOk returns a tuple with the LastPassedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetLastPassedTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPassedTimestamp) {
		return nil, false
	}
	return o.LastPassedTimestamp, true
}

// HasLastPassedTimestamp returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasLastPassedTimestamp() bool {
	if o != nil && !IsNil(o.LastPassedTimestamp) {
		return true
	}

	return false
}

// SetLastPassedTimestamp gets a reference to the given time.Time and assigns it to the LastPassedTimestamp field.
func (o *WorkflowServiceItemHealthCheckExecution) SetLastPassedTimestamp(v time.Time) {
	o.LastPassedTimestamp = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecution) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *WorkflowServiceItemHealthCheckExecution) SetResult(v string) {
	o.Result = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecution) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *WorkflowServiceItemHealthCheckExecution) SetSummary(v string) {
	o.Summary = &v
}

// GetWorkflowStatus returns the WorkflowStatus field value if set, zero value otherwise.
func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowStatus() string {
	if o == nil || IsNil(o.WorkflowStatus) {
		var ret string
		return ret
	}
	return *o.WorkflowStatus
}

// GetWorkflowStatusOk returns a tuple with the WorkflowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowStatusOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowStatus) {
		return nil, false
	}
	return o.WorkflowStatus, true
}

// HasWorkflowStatus returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasWorkflowStatus() bool {
	if o != nil && !IsNil(o.WorkflowStatus) {
		return true
	}

	return false
}

// SetWorkflowStatus gets a reference to the given string and assigns it to the WorkflowStatus field.
func (o *WorkflowServiceItemHealthCheckExecution) SetWorkflowStatus(v string) {
	o.WorkflowStatus = &v
}

// GetHealthCheckDefinition returns the HealthCheckDefinition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckExecution) GetHealthCheckDefinition() WorkflowServiceItemHealthCheckDefinitionRelationship {
	if o == nil || IsNil(o.HealthCheckDefinition.Get()) {
		var ret WorkflowServiceItemHealthCheckDefinitionRelationship
		return ret
	}
	return *o.HealthCheckDefinition.Get()
}

// GetHealthCheckDefinitionOk returns a tuple with the HealthCheckDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckExecution) GetHealthCheckDefinitionOk() (*WorkflowServiceItemHealthCheckDefinitionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.HealthCheckDefinition.Get(), o.HealthCheckDefinition.IsSet()
}

// HasHealthCheckDefinition returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasHealthCheckDefinition() bool {
	if o != nil && o.HealthCheckDefinition.IsSet() {
		return true
	}

	return false
}

// SetHealthCheckDefinition gets a reference to the given NullableWorkflowServiceItemHealthCheckDefinitionRelationship and assigns it to the HealthCheckDefinition field.
func (o *WorkflowServiceItemHealthCheckExecution) SetHealthCheckDefinition(v WorkflowServiceItemHealthCheckDefinitionRelationship) {
	o.HealthCheckDefinition.Set(&v)
}

// SetHealthCheckDefinitionNil sets the value for HealthCheckDefinition to be an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) SetHealthCheckDefinitionNil() {
	o.HealthCheckDefinition.Set(nil)
}

// UnsetHealthCheckDefinition ensures that no value is present for HealthCheckDefinition, not even an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) UnsetHealthCheckDefinition() {
	o.HealthCheckDefinition.Unset()
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckExecution) GetServiceItemInstance() WorkflowServiceItemInstanceRelationship {
	if o == nil || IsNil(o.ServiceItemInstance.Get()) {
		var ret WorkflowServiceItemInstanceRelationship
		return ret
	}
	return *o.ServiceItemInstance.Get()
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckExecution) GetServiceItemInstanceOk() (*WorkflowServiceItemInstanceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceItemInstance.Get(), o.ServiceItemInstance.IsSet()
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasServiceItemInstance() bool {
	if o != nil && o.ServiceItemInstance.IsSet() {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given NullableWorkflowServiceItemInstanceRelationship and assigns it to the ServiceItemInstance field.
func (o *WorkflowServiceItemHealthCheckExecution) SetServiceItemInstance(v WorkflowServiceItemInstanceRelationship) {
	o.ServiceItemInstance.Set(&v)
}

// SetServiceItemInstanceNil sets the value for ServiceItemInstance to be an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) SetServiceItemInstanceNil() {
	o.ServiceItemInstance.Set(nil)
}

// UnsetServiceItemInstance ensures that no value is present for ServiceItemInstance, not even an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) UnsetServiceItemInstance() {
	o.ServiceItemInstance.Unset()
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.WorkflowInfo.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo.Get()
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowServiceItemHealthCheckExecution) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowInfo.Get(), o.WorkflowInfo.IsSet()
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *WorkflowServiceItemHealthCheckExecution) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo.IsSet() {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *WorkflowServiceItemHealthCheckExecution) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo.Set(&v)
}

// SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) SetWorkflowInfoNil() {
	o.WorkflowInfo.Set(nil)
}

// UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil
func (o *WorkflowServiceItemHealthCheckExecution) UnsetWorkflowInfo() {
	o.WorkflowInfo.Unset()
}

func (o WorkflowServiceItemHealthCheckExecution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowServiceItemHealthCheckExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ErrorElements != nil {
		toSerialize["ErrorElements"] = o.ErrorElements
	}
	if !IsNil(o.ErrorSummary) {
		toSerialize["ErrorSummary"] = o.ErrorSummary
	}
	if !IsNil(o.LastPassedTimestamp) {
		toSerialize["LastPassedTimestamp"] = o.LastPassedTimestamp
	}
	if !IsNil(o.Result) {
		toSerialize["Result"] = o.Result
	}
	if !IsNil(o.Summary) {
		toSerialize["Summary"] = o.Summary
	}
	if !IsNil(o.WorkflowStatus) {
		toSerialize["WorkflowStatus"] = o.WorkflowStatus
	}
	if o.HealthCheckDefinition.IsSet() {
		toSerialize["HealthCheckDefinition"] = o.HealthCheckDefinition.Get()
	}
	if o.ServiceItemInstance.IsSet() {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance.Get()
	}
	if o.WorkflowInfo.IsSet() {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowServiceItemHealthCheckExecution) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                               `json:"ObjectType"`
		ErrorElements []ServiceitemHealthCheckErrorElement `json:"ErrorElements,omitempty"`
		// Error summary of a health check execution failure.
		ErrorSummary *string `json:"ErrorSummary,omitempty"`
		// Timestamp of the last passed health check execution.
		LastPassedTimestamp *time.Time `json:"LastPassedTimestamp,omitempty"`
		// Health check execution result. * `Unknown` - Indicates that the health check results could not be determined. * `Pass` - Indicates that the health check has passed. * `Fail` - Indicates that the health check has failed. * `Warning` - Indicates that the health check completed with a warning. * `NotApplicable` - Indicates that the health check is either unsupported, or not applicable for the service item.
		Result *string `json:"Result,omitempty"`
		// A brief summary of health check execution result.
		Summary *string `json:"Summary,omitempty"`
		// Status of the workflow that is executed as a part of health check execution.
		WorkflowStatus        *string                                                      `json:"WorkflowStatus,omitempty"`
		HealthCheckDefinition NullableWorkflowServiceItemHealthCheckDefinitionRelationship `json:"HealthCheckDefinition,omitempty"`
		ServiceItemInstance   NullableWorkflowServiceItemInstanceRelationship              `json:"ServiceItemInstance,omitempty"`
		WorkflowInfo          NullableWorkflowWorkflowInfoRelationship                     `json:"WorkflowInfo,omitempty"`
	}

	varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct := WorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowServiceItemHealthCheckExecution := _WorkflowServiceItemHealthCheckExecution{}
		varWorkflowServiceItemHealthCheckExecution.ClassId = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.ClassId
		varWorkflowServiceItemHealthCheckExecution.ObjectType = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.ObjectType
		varWorkflowServiceItemHealthCheckExecution.ErrorElements = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.ErrorElements
		varWorkflowServiceItemHealthCheckExecution.ErrorSummary = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.ErrorSummary
		varWorkflowServiceItemHealthCheckExecution.LastPassedTimestamp = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.LastPassedTimestamp
		varWorkflowServiceItemHealthCheckExecution.Result = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.Result
		varWorkflowServiceItemHealthCheckExecution.Summary = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.Summary
		varWorkflowServiceItemHealthCheckExecution.WorkflowStatus = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.WorkflowStatus
		varWorkflowServiceItemHealthCheckExecution.HealthCheckDefinition = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.HealthCheckDefinition
		varWorkflowServiceItemHealthCheckExecution.ServiceItemInstance = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.ServiceItemInstance
		varWorkflowServiceItemHealthCheckExecution.WorkflowInfo = varWorkflowServiceItemHealthCheckExecutionWithoutEmbeddedStruct.WorkflowInfo
		*o = WorkflowServiceItemHealthCheckExecution(varWorkflowServiceItemHealthCheckExecution)
	} else {
		return err
	}

	varWorkflowServiceItemHealthCheckExecution := _WorkflowServiceItemHealthCheckExecution{}

	err = json.Unmarshal(data, &varWorkflowServiceItemHealthCheckExecution)
	if err == nil {
		o.MoBaseMo = varWorkflowServiceItemHealthCheckExecution.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorElements")
		delete(additionalProperties, "ErrorSummary")
		delete(additionalProperties, "LastPassedTimestamp")
		delete(additionalProperties, "Result")
		delete(additionalProperties, "Summary")
		delete(additionalProperties, "WorkflowStatus")
		delete(additionalProperties, "HealthCheckDefinition")
		delete(additionalProperties, "ServiceItemInstance")
		delete(additionalProperties, "WorkflowInfo")

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

type NullableWorkflowServiceItemHealthCheckExecution struct {
	value *WorkflowServiceItemHealthCheckExecution
	isSet bool
}

func (v NullableWorkflowServiceItemHealthCheckExecution) Get() *WorkflowServiceItemHealthCheckExecution {
	return v.value
}

func (v *NullableWorkflowServiceItemHealthCheckExecution) Set(val *WorkflowServiceItemHealthCheckExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowServiceItemHealthCheckExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowServiceItemHealthCheckExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowServiceItemHealthCheckExecution(val *WorkflowServiceItemHealthCheckExecution) *NullableWorkflowServiceItemHealthCheckExecution {
	return &NullableWorkflowServiceItemHealthCheckExecution{value: val, isSet: true}
}

func (v NullableWorkflowServiceItemHealthCheckExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowServiceItemHealthCheckExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
