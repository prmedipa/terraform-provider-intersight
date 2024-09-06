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
)

// checks if the ServicenowChangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicenowChangeRequest{}

// ServicenowChangeRequest A Change Request on ServiceNow.
type ServicenowChangeRequest struct {
	ServicenowInventoryEntity
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The approver of Change Request.
	Approval *string `json:"Approval,omitempty"`
	// Assigned to value for Change Request.
	AssignedTo *string `json:"AssignedTo,omitempty"`
	// Assigned to display value for Change Request.
	AssignedToDisplayValue *string `json:"AssignedToDisplayValue,omitempty"`
	// Change Model for Change Request.
	ChangeModel *string `json:"ChangeModel,omitempty"`
	// Change Model display value for Change Request.
	ChangeModelDisplayValue *string `json:"ChangeModelDisplayValue,omitempty"`
	// Number for Change Request.
	ChangeRequestNumber *string `json:"ChangeRequestNumber,omitempty"`
	// Comments on Change Request.
	Comments *string `json:"Comments,omitempty"`
	// Creator of Change Request.
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// Description for Change Request.
	Description *string `json:"Description,omitempty"`
	// Due date for Change Request.
	DueDate *string `json:"DueDate,omitempty"`
	// End date for Change Request.
	EndDate *string `json:"EndDate,omitempty"`
	// Impact for Change Request.
	Impact *string `json:"Impact,omitempty"`
	// Impact display value for Change Request.
	ImpactDisplayValue *string `json:"ImpactDisplayValue,omitempty"`
	// Priority for Change Request.
	Priority *string `json:"Priority,omitempty"`
	// Priority display value for Change Request.
	PriorityDisplayValue *string `json:"PriorityDisplayValue,omitempty"`
	// The risk for Change Request.
	Risk *string `json:"Risk,omitempty"`
	// Short Description for Change Request.
	ShortDescription *string `json:"ShortDescription,omitempty"`
	// Start date for Change Request.
	StartDate *string `json:"StartDate,omitempty"`
	// Sys Id for Change Request.
	SysId *string `json:"SysId,omitempty"`
	// The type for Change Request.
	Type *string `json:"Type,omitempty"`
	// Last update Change Request.
	UpdatedBy            *string `json:"UpdatedBy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServicenowChangeRequest ServicenowChangeRequest

// NewServicenowChangeRequest instantiates a new ServicenowChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicenowChangeRequest(classId string, objectType string) *ServicenowChangeRequest {
	this := ServicenowChangeRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewServicenowChangeRequestWithDefaults instantiates a new ServicenowChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicenowChangeRequestWithDefaults() *ServicenowChangeRequest {
	this := ServicenowChangeRequest{}
	var classId string = "servicenow.ChangeRequest"
	this.ClassId = classId
	var objectType string = "servicenow.ChangeRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ServicenowChangeRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ServicenowChangeRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "servicenow.ChangeRequest" of the ClassId field.
func (o *ServicenowChangeRequest) GetDefaultClassId() interface{} {
	return "servicenow.ChangeRequest"
}

// GetObjectType returns the ObjectType field value
func (o *ServicenowChangeRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ServicenowChangeRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "servicenow.ChangeRequest" of the ObjectType field.
func (o *ServicenowChangeRequest) GetDefaultObjectType() interface{} {
	return "servicenow.ChangeRequest"
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetApproval() string {
	if o == nil || IsNil(o.Approval) {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetApprovalOk() (*string, bool) {
	if o == nil || IsNil(o.Approval) {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasApproval() bool {
	if o != nil && !IsNil(o.Approval) {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *ServicenowChangeRequest) SetApproval(v string) {
	o.Approval = &v
}

// GetAssignedTo returns the AssignedTo field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetAssignedTo() string {
	if o == nil || IsNil(o.AssignedTo) {
		var ret string
		return ret
	}
	return *o.AssignedTo
}

// GetAssignedToOk returns a tuple with the AssignedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetAssignedToOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedTo) {
		return nil, false
	}
	return o.AssignedTo, true
}

// HasAssignedTo returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasAssignedTo() bool {
	if o != nil && !IsNil(o.AssignedTo) {
		return true
	}

	return false
}

// SetAssignedTo gets a reference to the given string and assigns it to the AssignedTo field.
func (o *ServicenowChangeRequest) SetAssignedTo(v string) {
	o.AssignedTo = &v
}

// GetAssignedToDisplayValue returns the AssignedToDisplayValue field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetAssignedToDisplayValue() string {
	if o == nil || IsNil(o.AssignedToDisplayValue) {
		var ret string
		return ret
	}
	return *o.AssignedToDisplayValue
}

// GetAssignedToDisplayValueOk returns a tuple with the AssignedToDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetAssignedToDisplayValueOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedToDisplayValue) {
		return nil, false
	}
	return o.AssignedToDisplayValue, true
}

// HasAssignedToDisplayValue returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasAssignedToDisplayValue() bool {
	if o != nil && !IsNil(o.AssignedToDisplayValue) {
		return true
	}

	return false
}

// SetAssignedToDisplayValue gets a reference to the given string and assigns it to the AssignedToDisplayValue field.
func (o *ServicenowChangeRequest) SetAssignedToDisplayValue(v string) {
	o.AssignedToDisplayValue = &v
}

// GetChangeModel returns the ChangeModel field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetChangeModel() string {
	if o == nil || IsNil(o.ChangeModel) {
		var ret string
		return ret
	}
	return *o.ChangeModel
}

// GetChangeModelOk returns a tuple with the ChangeModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetChangeModelOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeModel) {
		return nil, false
	}
	return o.ChangeModel, true
}

// HasChangeModel returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasChangeModel() bool {
	if o != nil && !IsNil(o.ChangeModel) {
		return true
	}

	return false
}

// SetChangeModel gets a reference to the given string and assigns it to the ChangeModel field.
func (o *ServicenowChangeRequest) SetChangeModel(v string) {
	o.ChangeModel = &v
}

// GetChangeModelDisplayValue returns the ChangeModelDisplayValue field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetChangeModelDisplayValue() string {
	if o == nil || IsNil(o.ChangeModelDisplayValue) {
		var ret string
		return ret
	}
	return *o.ChangeModelDisplayValue
}

// GetChangeModelDisplayValueOk returns a tuple with the ChangeModelDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetChangeModelDisplayValueOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeModelDisplayValue) {
		return nil, false
	}
	return o.ChangeModelDisplayValue, true
}

// HasChangeModelDisplayValue returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasChangeModelDisplayValue() bool {
	if o != nil && !IsNil(o.ChangeModelDisplayValue) {
		return true
	}

	return false
}

// SetChangeModelDisplayValue gets a reference to the given string and assigns it to the ChangeModelDisplayValue field.
func (o *ServicenowChangeRequest) SetChangeModelDisplayValue(v string) {
	o.ChangeModelDisplayValue = &v
}

// GetChangeRequestNumber returns the ChangeRequestNumber field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetChangeRequestNumber() string {
	if o == nil || IsNil(o.ChangeRequestNumber) {
		var ret string
		return ret
	}
	return *o.ChangeRequestNumber
}

// GetChangeRequestNumberOk returns a tuple with the ChangeRequestNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetChangeRequestNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeRequestNumber) {
		return nil, false
	}
	return o.ChangeRequestNumber, true
}

// HasChangeRequestNumber returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasChangeRequestNumber() bool {
	if o != nil && !IsNil(o.ChangeRequestNumber) {
		return true
	}

	return false
}

// SetChangeRequestNumber gets a reference to the given string and assigns it to the ChangeRequestNumber field.
func (o *ServicenowChangeRequest) SetChangeRequestNumber(v string) {
	o.ChangeRequestNumber = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ServicenowChangeRequest) SetComments(v string) {
	o.Comments = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ServicenowChangeRequest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServicenowChangeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *ServicenowChangeRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ServicenowChangeRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetImpact() string {
	if o == nil || IsNil(o.Impact) {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetImpactOk() (*string, bool) {
	if o == nil || IsNil(o.Impact) {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasImpact() bool {
	if o != nil && !IsNil(o.Impact) {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *ServicenowChangeRequest) SetImpact(v string) {
	o.Impact = &v
}

// GetImpactDisplayValue returns the ImpactDisplayValue field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetImpactDisplayValue() string {
	if o == nil || IsNil(o.ImpactDisplayValue) {
		var ret string
		return ret
	}
	return *o.ImpactDisplayValue
}

// GetImpactDisplayValueOk returns a tuple with the ImpactDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetImpactDisplayValueOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactDisplayValue) {
		return nil, false
	}
	return o.ImpactDisplayValue, true
}

// HasImpactDisplayValue returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasImpactDisplayValue() bool {
	if o != nil && !IsNil(o.ImpactDisplayValue) {
		return true
	}

	return false
}

// SetImpactDisplayValue gets a reference to the given string and assigns it to the ImpactDisplayValue field.
func (o *ServicenowChangeRequest) SetImpactDisplayValue(v string) {
	o.ImpactDisplayValue = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ServicenowChangeRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetPriorityDisplayValue returns the PriorityDisplayValue field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetPriorityDisplayValue() string {
	if o == nil || IsNil(o.PriorityDisplayValue) {
		var ret string
		return ret
	}
	return *o.PriorityDisplayValue
}

// GetPriorityDisplayValueOk returns a tuple with the PriorityDisplayValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetPriorityDisplayValueOk() (*string, bool) {
	if o == nil || IsNil(o.PriorityDisplayValue) {
		return nil, false
	}
	return o.PriorityDisplayValue, true
}

// HasPriorityDisplayValue returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasPriorityDisplayValue() bool {
	if o != nil && !IsNil(o.PriorityDisplayValue) {
		return true
	}

	return false
}

// SetPriorityDisplayValue gets a reference to the given string and assigns it to the PriorityDisplayValue field.
func (o *ServicenowChangeRequest) SetPriorityDisplayValue(v string) {
	o.PriorityDisplayValue = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetRisk() string {
	if o == nil || IsNil(o.Risk) {
		var ret string
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetRiskOk() (*string, bool) {
	if o == nil || IsNil(o.Risk) {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasRisk() bool {
	if o != nil && !IsNil(o.Risk) {
		return true
	}

	return false
}

// SetRisk gets a reference to the given string and assigns it to the Risk field.
func (o *ServicenowChangeRequest) SetRisk(v string) {
	o.Risk = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ServicenowChangeRequest) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ServicenowChangeRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetSysId returns the SysId field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetSysId() string {
	if o == nil || IsNil(o.SysId) {
		var ret string
		return ret
	}
	return *o.SysId
}

// GetSysIdOk returns a tuple with the SysId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetSysIdOk() (*string, bool) {
	if o == nil || IsNil(o.SysId) {
		return nil, false
	}
	return o.SysId, true
}

// HasSysId returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasSysId() bool {
	if o != nil && !IsNil(o.SysId) {
		return true
	}

	return false
}

// SetSysId gets a reference to the given string and assigns it to the SysId field.
func (o *ServicenowChangeRequest) SetSysId(v string) {
	o.SysId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServicenowChangeRequest) SetType(v string) {
	o.Type = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ServicenowChangeRequest) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicenowChangeRequest) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ServicenowChangeRequest) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *ServicenowChangeRequest) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o ServicenowChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicenowChangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedServicenowInventoryEntity, errServicenowInventoryEntity := json.Marshal(o.ServicenowInventoryEntity)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	errServicenowInventoryEntity = json.Unmarshal([]byte(serializedServicenowInventoryEntity), &toSerialize)
	if errServicenowInventoryEntity != nil {
		return map[string]interface{}{}, errServicenowInventoryEntity
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Approval) {
		toSerialize["Approval"] = o.Approval
	}
	if !IsNil(o.AssignedTo) {
		toSerialize["AssignedTo"] = o.AssignedTo
	}
	if !IsNil(o.AssignedToDisplayValue) {
		toSerialize["AssignedToDisplayValue"] = o.AssignedToDisplayValue
	}
	if !IsNil(o.ChangeModel) {
		toSerialize["ChangeModel"] = o.ChangeModel
	}
	if !IsNil(o.ChangeModelDisplayValue) {
		toSerialize["ChangeModelDisplayValue"] = o.ChangeModelDisplayValue
	}
	if !IsNil(o.ChangeRequestNumber) {
		toSerialize["ChangeRequestNumber"] = o.ChangeRequestNumber
	}
	if !IsNil(o.Comments) {
		toSerialize["Comments"] = o.Comments
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["CreatedBy"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	if !IsNil(o.Impact) {
		toSerialize["Impact"] = o.Impact
	}
	if !IsNil(o.ImpactDisplayValue) {
		toSerialize["ImpactDisplayValue"] = o.ImpactDisplayValue
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.PriorityDisplayValue) {
		toSerialize["PriorityDisplayValue"] = o.PriorityDisplayValue
	}
	if !IsNil(o.Risk) {
		toSerialize["Risk"] = o.Risk
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["ShortDescription"] = o.ShortDescription
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.SysId) {
		toSerialize["SysId"] = o.SysId
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["UpdatedBy"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServicenowChangeRequest) UnmarshalJSON(data []byte) (err error) {
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
	type ServicenowChangeRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The approver of Change Request.
		Approval *string `json:"Approval,omitempty"`
		// Assigned to value for Change Request.
		AssignedTo *string `json:"AssignedTo,omitempty"`
		// Assigned to display value for Change Request.
		AssignedToDisplayValue *string `json:"AssignedToDisplayValue,omitempty"`
		// Change Model for Change Request.
		ChangeModel *string `json:"ChangeModel,omitempty"`
		// Change Model display value for Change Request.
		ChangeModelDisplayValue *string `json:"ChangeModelDisplayValue,omitempty"`
		// Number for Change Request.
		ChangeRequestNumber *string `json:"ChangeRequestNumber,omitempty"`
		// Comments on Change Request.
		Comments *string `json:"Comments,omitempty"`
		// Creator of Change Request.
		CreatedBy *string `json:"CreatedBy,omitempty"`
		// Description for Change Request.
		Description *string `json:"Description,omitempty"`
		// Due date for Change Request.
		DueDate *string `json:"DueDate,omitempty"`
		// End date for Change Request.
		EndDate *string `json:"EndDate,omitempty"`
		// Impact for Change Request.
		Impact *string `json:"Impact,omitempty"`
		// Impact display value for Change Request.
		ImpactDisplayValue *string `json:"ImpactDisplayValue,omitempty"`
		// Priority for Change Request.
		Priority *string `json:"Priority,omitempty"`
		// Priority display value for Change Request.
		PriorityDisplayValue *string `json:"PriorityDisplayValue,omitempty"`
		// The risk for Change Request.
		Risk *string `json:"Risk,omitempty"`
		// Short Description for Change Request.
		ShortDescription *string `json:"ShortDescription,omitempty"`
		// Start date for Change Request.
		StartDate *string `json:"StartDate,omitempty"`
		// Sys Id for Change Request.
		SysId *string `json:"SysId,omitempty"`
		// The type for Change Request.
		Type *string `json:"Type,omitempty"`
		// Last update Change Request.
		UpdatedBy *string `json:"UpdatedBy,omitempty"`
	}

	varServicenowChangeRequestWithoutEmbeddedStruct := ServicenowChangeRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varServicenowChangeRequestWithoutEmbeddedStruct)
	if err == nil {
		varServicenowChangeRequest := _ServicenowChangeRequest{}
		varServicenowChangeRequest.ClassId = varServicenowChangeRequestWithoutEmbeddedStruct.ClassId
		varServicenowChangeRequest.ObjectType = varServicenowChangeRequestWithoutEmbeddedStruct.ObjectType
		varServicenowChangeRequest.Approval = varServicenowChangeRequestWithoutEmbeddedStruct.Approval
		varServicenowChangeRequest.AssignedTo = varServicenowChangeRequestWithoutEmbeddedStruct.AssignedTo
		varServicenowChangeRequest.AssignedToDisplayValue = varServicenowChangeRequestWithoutEmbeddedStruct.AssignedToDisplayValue
		varServicenowChangeRequest.ChangeModel = varServicenowChangeRequestWithoutEmbeddedStruct.ChangeModel
		varServicenowChangeRequest.ChangeModelDisplayValue = varServicenowChangeRequestWithoutEmbeddedStruct.ChangeModelDisplayValue
		varServicenowChangeRequest.ChangeRequestNumber = varServicenowChangeRequestWithoutEmbeddedStruct.ChangeRequestNumber
		varServicenowChangeRequest.Comments = varServicenowChangeRequestWithoutEmbeddedStruct.Comments
		varServicenowChangeRequest.CreatedBy = varServicenowChangeRequestWithoutEmbeddedStruct.CreatedBy
		varServicenowChangeRequest.Description = varServicenowChangeRequestWithoutEmbeddedStruct.Description
		varServicenowChangeRequest.DueDate = varServicenowChangeRequestWithoutEmbeddedStruct.DueDate
		varServicenowChangeRequest.EndDate = varServicenowChangeRequestWithoutEmbeddedStruct.EndDate
		varServicenowChangeRequest.Impact = varServicenowChangeRequestWithoutEmbeddedStruct.Impact
		varServicenowChangeRequest.ImpactDisplayValue = varServicenowChangeRequestWithoutEmbeddedStruct.ImpactDisplayValue
		varServicenowChangeRequest.Priority = varServicenowChangeRequestWithoutEmbeddedStruct.Priority
		varServicenowChangeRequest.PriorityDisplayValue = varServicenowChangeRequestWithoutEmbeddedStruct.PriorityDisplayValue
		varServicenowChangeRequest.Risk = varServicenowChangeRequestWithoutEmbeddedStruct.Risk
		varServicenowChangeRequest.ShortDescription = varServicenowChangeRequestWithoutEmbeddedStruct.ShortDescription
		varServicenowChangeRequest.StartDate = varServicenowChangeRequestWithoutEmbeddedStruct.StartDate
		varServicenowChangeRequest.SysId = varServicenowChangeRequestWithoutEmbeddedStruct.SysId
		varServicenowChangeRequest.Type = varServicenowChangeRequestWithoutEmbeddedStruct.Type
		varServicenowChangeRequest.UpdatedBy = varServicenowChangeRequestWithoutEmbeddedStruct.UpdatedBy
		*o = ServicenowChangeRequest(varServicenowChangeRequest)
	} else {
		return err
	}

	varServicenowChangeRequest := _ServicenowChangeRequest{}

	err = json.Unmarshal(data, &varServicenowChangeRequest)
	if err == nil {
		o.ServicenowInventoryEntity = varServicenowChangeRequest.ServicenowInventoryEntity
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Approval")
		delete(additionalProperties, "AssignedTo")
		delete(additionalProperties, "AssignedToDisplayValue")
		delete(additionalProperties, "ChangeModel")
		delete(additionalProperties, "ChangeModelDisplayValue")
		delete(additionalProperties, "ChangeRequestNumber")
		delete(additionalProperties, "Comments")
		delete(additionalProperties, "CreatedBy")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DueDate")
		delete(additionalProperties, "EndDate")
		delete(additionalProperties, "Impact")
		delete(additionalProperties, "ImpactDisplayValue")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "PriorityDisplayValue")
		delete(additionalProperties, "Risk")
		delete(additionalProperties, "ShortDescription")
		delete(additionalProperties, "StartDate")
		delete(additionalProperties, "SysId")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "UpdatedBy")

		// remove fields from embedded structs
		reflectServicenowInventoryEntity := reflect.ValueOf(o.ServicenowInventoryEntity)
		for i := 0; i < reflectServicenowInventoryEntity.Type().NumField(); i++ {
			t := reflectServicenowInventoryEntity.Type().Field(i)

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

type NullableServicenowChangeRequest struct {
	value *ServicenowChangeRequest
	isSet bool
}

func (v NullableServicenowChangeRequest) Get() *ServicenowChangeRequest {
	return v.value
}

func (v *NullableServicenowChangeRequest) Set(val *ServicenowChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServicenowChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServicenowChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicenowChangeRequest(val *ServicenowChangeRequest) *NullableServicenowChangeRequest {
	return &NullableServicenowChangeRequest{value: val, isSet: true}
}

func (v NullableServicenowChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicenowChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
