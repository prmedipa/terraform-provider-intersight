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

// checks if the BulkResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkResult{}

// BulkResult The Result API is a status-monitor resource used to show the processing status of any bulk MO API when the request HTTP 'prefer' header is set to 'respond-async' value.
type BulkResult struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The action that will be performed when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
	ActionOnError *string `json:"ActionOnError,omitempty"`
	// The timestamp in UTC when the request processing is completed.
	CompletionTime *time.Time `json:"CompletionTime,omitempty"`
	// The number of subrequests received in this request.
	NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
	// The individual request to be executed asynchronously.
	Request interface{} `json:"Request,omitempty"`
	// The timestamp in UTC when the request was received.
	RequestReceivedTime *time.Time `json:"RequestReceivedTime,omitempty"`
	// The processing status of the request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
	Status *string `json:"Status,omitempty"`
	// The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// The URI on which this async operation is being performed.
	Uri          *string                                      `json:"Uri,omitempty"`
	MoCloner     NullableBulkMoClonerRelationship             `json:"MoCloner,omitempty"`
	MoDeepCloner NullableBulkMoDeepClonerRelationship         `json:"MoDeepCloner,omitempty"`
	MoMerger     NullableBulkMoMergerRelationship             `json:"MoMerger,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to bulkSubRequestObj resources.
	Results              []BulkSubRequestObjRelationship          `json:"Results,omitempty"`
	WorkflowInfo         NullableWorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkResult BulkResult

// NewBulkResult instantiates a new BulkResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkResult(classId string, objectType string) *BulkResult {
	this := BulkResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBulkResultWithDefaults instantiates a new BulkResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkResultWithDefaults() *BulkResult {
	this := BulkResult{}
	var classId string = "bulk.Result"
	this.ClassId = classId
	var objectType string = "bulk.Result"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BulkResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BulkResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BulkResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bulk.Result" of the ClassId field.
func (o *BulkResult) GetDefaultClassId() interface{} {
	return "bulk.Result"
}

// GetObjectType returns the ObjectType field value
func (o *BulkResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BulkResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BulkResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bulk.Result" of the ObjectType field.
func (o *BulkResult) GetDefaultObjectType() interface{} {
	return "bulk.Result"
}

// GetActionOnError returns the ActionOnError field value if set, zero value otherwise.
func (o *BulkResult) GetActionOnError() string {
	if o == nil || IsNil(o.ActionOnError) {
		var ret string
		return ret
	}
	return *o.ActionOnError
}

// GetActionOnErrorOk returns a tuple with the ActionOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetActionOnErrorOk() (*string, bool) {
	if o == nil || IsNil(o.ActionOnError) {
		return nil, false
	}
	return o.ActionOnError, true
}

// HasActionOnError returns a boolean if a field has been set.
func (o *BulkResult) HasActionOnError() bool {
	if o != nil && !IsNil(o.ActionOnError) {
		return true
	}

	return false
}

// SetActionOnError gets a reference to the given string and assigns it to the ActionOnError field.
func (o *BulkResult) SetActionOnError(v string) {
	o.ActionOnError = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *BulkResult) GetCompletionTime() time.Time {
	if o == nil || IsNil(o.CompletionTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionTime) {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *BulkResult) HasCompletionTime() bool {
	if o != nil && !IsNil(o.CompletionTime) {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *BulkResult) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetNumSubRequests returns the NumSubRequests field value if set, zero value otherwise.
func (o *BulkResult) GetNumSubRequests() int64 {
	if o == nil || IsNil(o.NumSubRequests) {
		var ret int64
		return ret
	}
	return *o.NumSubRequests
}

// GetNumSubRequestsOk returns a tuple with the NumSubRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetNumSubRequestsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumSubRequests) {
		return nil, false
	}
	return o.NumSubRequests, true
}

// HasNumSubRequests returns a boolean if a field has been set.
func (o *BulkResult) HasNumSubRequests() bool {
	if o != nil && !IsNil(o.NumSubRequests) {
		return true
	}

	return false
}

// SetNumSubRequests gets a reference to the given int64 and assigns it to the NumSubRequests field.
func (o *BulkResult) SetNumSubRequests(v int64) {
	o.NumSubRequests = &v
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetRequest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetRequestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return &o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *BulkResult) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given interface{} and assigns it to the Request field.
func (o *BulkResult) SetRequest(v interface{}) {
	o.Request = v
}

// GetRequestReceivedTime returns the RequestReceivedTime field value if set, zero value otherwise.
func (o *BulkResult) GetRequestReceivedTime() time.Time {
	if o == nil || IsNil(o.RequestReceivedTime) {
		var ret time.Time
		return ret
	}
	return *o.RequestReceivedTime
}

// GetRequestReceivedTimeOk returns a tuple with the RequestReceivedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetRequestReceivedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestReceivedTime) {
		return nil, false
	}
	return o.RequestReceivedTime, true
}

// HasRequestReceivedTime returns a boolean if a field has been set.
func (o *BulkResult) HasRequestReceivedTime() bool {
	if o != nil && !IsNil(o.RequestReceivedTime) {
		return true
	}

	return false
}

// SetRequestReceivedTime gets a reference to the given time.Time and assigns it to the RequestReceivedTime field.
func (o *BulkResult) SetRequestReceivedTime(v time.Time) {
	o.RequestReceivedTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BulkResult) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BulkResult) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BulkResult) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BulkResult) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *BulkResult) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkResult) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *BulkResult) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *BulkResult) SetUri(v string) {
	o.Uri = &v
}

// GetMoCloner returns the MoCloner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetMoCloner() BulkMoClonerRelationship {
	if o == nil || IsNil(o.MoCloner.Get()) {
		var ret BulkMoClonerRelationship
		return ret
	}
	return *o.MoCloner.Get()
}

// GetMoClonerOk returns a tuple with the MoCloner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetMoClonerOk() (*BulkMoClonerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoCloner.Get(), o.MoCloner.IsSet()
}

// HasMoCloner returns a boolean if a field has been set.
func (o *BulkResult) HasMoCloner() bool {
	if o != nil && o.MoCloner.IsSet() {
		return true
	}

	return false
}

// SetMoCloner gets a reference to the given NullableBulkMoClonerRelationship and assigns it to the MoCloner field.
func (o *BulkResult) SetMoCloner(v BulkMoClonerRelationship) {
	o.MoCloner.Set(&v)
}

// SetMoClonerNil sets the value for MoCloner to be an explicit nil
func (o *BulkResult) SetMoClonerNil() {
	o.MoCloner.Set(nil)
}

// UnsetMoCloner ensures that no value is present for MoCloner, not even an explicit nil
func (o *BulkResult) UnsetMoCloner() {
	o.MoCloner.Unset()
}

// GetMoDeepCloner returns the MoDeepCloner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetMoDeepCloner() BulkMoDeepClonerRelationship {
	if o == nil || IsNil(o.MoDeepCloner.Get()) {
		var ret BulkMoDeepClonerRelationship
		return ret
	}
	return *o.MoDeepCloner.Get()
}

// GetMoDeepClonerOk returns a tuple with the MoDeepCloner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetMoDeepClonerOk() (*BulkMoDeepClonerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoDeepCloner.Get(), o.MoDeepCloner.IsSet()
}

// HasMoDeepCloner returns a boolean if a field has been set.
func (o *BulkResult) HasMoDeepCloner() bool {
	if o != nil && o.MoDeepCloner.IsSet() {
		return true
	}

	return false
}

// SetMoDeepCloner gets a reference to the given NullableBulkMoDeepClonerRelationship and assigns it to the MoDeepCloner field.
func (o *BulkResult) SetMoDeepCloner(v BulkMoDeepClonerRelationship) {
	o.MoDeepCloner.Set(&v)
}

// SetMoDeepClonerNil sets the value for MoDeepCloner to be an explicit nil
func (o *BulkResult) SetMoDeepClonerNil() {
	o.MoDeepCloner.Set(nil)
}

// UnsetMoDeepCloner ensures that no value is present for MoDeepCloner, not even an explicit nil
func (o *BulkResult) UnsetMoDeepCloner() {
	o.MoDeepCloner.Unset()
}

// GetMoMerger returns the MoMerger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetMoMerger() BulkMoMergerRelationship {
	if o == nil || IsNil(o.MoMerger.Get()) {
		var ret BulkMoMergerRelationship
		return ret
	}
	return *o.MoMerger.Get()
}

// GetMoMergerOk returns a tuple with the MoMerger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetMoMergerOk() (*BulkMoMergerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoMerger.Get(), o.MoMerger.IsSet()
}

// HasMoMerger returns a boolean if a field has been set.
func (o *BulkResult) HasMoMerger() bool {
	if o != nil && o.MoMerger.IsSet() {
		return true
	}

	return false
}

// SetMoMerger gets a reference to the given NullableBulkMoMergerRelationship and assigns it to the MoMerger field.
func (o *BulkResult) SetMoMerger(v BulkMoMergerRelationship) {
	o.MoMerger.Set(&v)
}

// SetMoMergerNil sets the value for MoMerger to be an explicit nil
func (o *BulkResult) SetMoMergerNil() {
	o.MoMerger.Set(nil)
}

// UnsetMoMerger ensures that no value is present for MoMerger, not even an explicit nil
func (o *BulkResult) UnsetMoMerger() {
	o.MoMerger.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *BulkResult) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *BulkResult) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *BulkResult) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *BulkResult) UnsetOrganization() {
	o.Organization.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetResults() []BulkSubRequestObjRelationship {
	if o == nil {
		var ret []BulkSubRequestObjRelationship
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetResultsOk() ([]BulkSubRequestObjRelationship, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BulkResult) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BulkSubRequestObjRelationship and assigns it to the Results field.
func (o *BulkResult) SetResults(v []BulkSubRequestObjRelationship) {
	o.Results = v
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkResult) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.WorkflowInfo.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo.Get()
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkResult) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowInfo.Get(), o.WorkflowInfo.IsSet()
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *BulkResult) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo.IsSet() {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *BulkResult) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo.Set(&v)
}

// SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil
func (o *BulkResult) SetWorkflowInfoNil() {
	o.WorkflowInfo.Set(nil)
}

// UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil
func (o *BulkResult) UnsetWorkflowInfo() {
	o.WorkflowInfo.Unset()
}

func (o BulkResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkResult) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ActionOnError) {
		toSerialize["ActionOnError"] = o.ActionOnError
	}
	if !IsNil(o.CompletionTime) {
		toSerialize["CompletionTime"] = o.CompletionTime
	}
	if !IsNil(o.NumSubRequests) {
		toSerialize["NumSubRequests"] = o.NumSubRequests
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if !IsNil(o.RequestReceivedTime) {
		toSerialize["RequestReceivedTime"] = o.RequestReceivedTime
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !IsNil(o.Uri) {
		toSerialize["Uri"] = o.Uri
	}
	if o.MoCloner.IsSet() {
		toSerialize["MoCloner"] = o.MoCloner.Get()
	}
	if o.MoDeepCloner.IsSet() {
		toSerialize["MoDeepCloner"] = o.MoDeepCloner.Get()
	}
	if o.MoMerger.IsSet() {
		toSerialize["MoMerger"] = o.MoMerger.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Results != nil {
		toSerialize["Results"] = o.Results
	}
	if o.WorkflowInfo.IsSet() {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkResult) UnmarshalJSON(data []byte) (err error) {
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
	type BulkResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The action that will be performed when an error occurs during processing of the request. * `Stop` - Stop the processing of the request after the first error. * `Proceed` - Proceed with the processing of the request even when an error occurs.
		ActionOnError *string `json:"ActionOnError,omitempty"`
		// The timestamp in UTC when the request processing is completed.
		CompletionTime *time.Time `json:"CompletionTime,omitempty"`
		// The number of subrequests received in this request.
		NumSubRequests *int64 `json:"NumSubRequests,omitempty"`
		// The individual request to be executed asynchronously.
		Request interface{} `json:"Request,omitempty"`
		// The timestamp in UTC when the request was received.
		RequestReceivedTime *time.Time `json:"RequestReceivedTime,omitempty"`
		// The processing status of the request. * `NotStarted` - Indicates that the request processing has not begun yet. * `ObjPresenceCheckInProgress` - Indicates that the object presence check is in progress for this request. * `ObjPresenceCheckComplete` - Indicates that the object presence check is complete. * `ExecutionInProgress` - Indicates that the request processing is in progress. * `Completed` - Indicates that the request processing has been completed successfully. * `CompletedWithErrors` - Indicates that the request processing has one or more failed subrequests. * `Failed` - Indicates that the processing of this request failed. * `TimedOut` - Indicates that the request processing timed out.
		Status *string `json:"Status,omitempty"`
		// The status message shows the error details in human readable format when the request goes to failed state. No additional information is shown for success case.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// The URI on which this async operation is being performed.
		Uri          *string                                      `json:"Uri,omitempty"`
		MoCloner     NullableBulkMoClonerRelationship             `json:"MoCloner,omitempty"`
		MoDeepCloner NullableBulkMoDeepClonerRelationship         `json:"MoDeepCloner,omitempty"`
		MoMerger     NullableBulkMoMergerRelationship             `json:"MoMerger,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to bulkSubRequestObj resources.
		Results      []BulkSubRequestObjRelationship          `json:"Results,omitempty"`
		WorkflowInfo NullableWorkflowWorkflowInfoRelationship `json:"WorkflowInfo,omitempty"`
	}

	varBulkResultWithoutEmbeddedStruct := BulkResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBulkResultWithoutEmbeddedStruct)
	if err == nil {
		varBulkResult := _BulkResult{}
		varBulkResult.ClassId = varBulkResultWithoutEmbeddedStruct.ClassId
		varBulkResult.ObjectType = varBulkResultWithoutEmbeddedStruct.ObjectType
		varBulkResult.ActionOnError = varBulkResultWithoutEmbeddedStruct.ActionOnError
		varBulkResult.CompletionTime = varBulkResultWithoutEmbeddedStruct.CompletionTime
		varBulkResult.NumSubRequests = varBulkResultWithoutEmbeddedStruct.NumSubRequests
		varBulkResult.Request = varBulkResultWithoutEmbeddedStruct.Request
		varBulkResult.RequestReceivedTime = varBulkResultWithoutEmbeddedStruct.RequestReceivedTime
		varBulkResult.Status = varBulkResultWithoutEmbeddedStruct.Status
		varBulkResult.StatusMessage = varBulkResultWithoutEmbeddedStruct.StatusMessage
		varBulkResult.Uri = varBulkResultWithoutEmbeddedStruct.Uri
		varBulkResult.MoCloner = varBulkResultWithoutEmbeddedStruct.MoCloner
		varBulkResult.MoDeepCloner = varBulkResultWithoutEmbeddedStruct.MoDeepCloner
		varBulkResult.MoMerger = varBulkResultWithoutEmbeddedStruct.MoMerger
		varBulkResult.Organization = varBulkResultWithoutEmbeddedStruct.Organization
		varBulkResult.Results = varBulkResultWithoutEmbeddedStruct.Results
		varBulkResult.WorkflowInfo = varBulkResultWithoutEmbeddedStruct.WorkflowInfo
		*o = BulkResult(varBulkResult)
	} else {
		return err
	}

	varBulkResult := _BulkResult{}

	err = json.Unmarshal(data, &varBulkResult)
	if err == nil {
		o.MoBaseMo = varBulkResult.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActionOnError")
		delete(additionalProperties, "CompletionTime")
		delete(additionalProperties, "NumSubRequests")
		delete(additionalProperties, "Request")
		delete(additionalProperties, "RequestReceivedTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "Uri")
		delete(additionalProperties, "MoCloner")
		delete(additionalProperties, "MoDeepCloner")
		delete(additionalProperties, "MoMerger")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Results")
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

type NullableBulkResult struct {
	value *BulkResult
	isSet bool
}

func (v NullableBulkResult) Get() *BulkResult {
	return v.value
}

func (v *NullableBulkResult) Set(val *BulkResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkResult(val *BulkResult) *NullableBulkResult {
	return &NullableBulkResult{value: val, isSet: true}
}

func (v NullableBulkResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
