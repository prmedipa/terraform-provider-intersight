/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7078
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// LicenseIksCustomerOpAllOf Definition of the list of properties defined in 'license.IksCustomerOp', excluding properties defined in parent classes.
type LicenseIksCustomerOpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Intersight Kubernetes Service license administrative state. Set this property to 'true' to activate the IKS license entitlements.
	ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
	// Enable trial for IKS licensing.
	EnableTrial *bool `json:"EnableTrial,omitempty"`
	// The default Trial or Grace period the customer is entitled to.
	EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
	// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
	ExtraEvaluation      *int64                                 `json:"ExtraEvaluation,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseIksCustomerOpAllOf LicenseIksCustomerOpAllOf

// NewLicenseIksCustomerOpAllOf instantiates a new LicenseIksCustomerOpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseIksCustomerOpAllOf(classId string, objectType string) *LicenseIksCustomerOpAllOf {
	this := LicenseIksCustomerOpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseIksCustomerOpAllOfWithDefaults instantiates a new LicenseIksCustomerOpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseIksCustomerOpAllOfWithDefaults() *LicenseIksCustomerOpAllOf {
	this := LicenseIksCustomerOpAllOf{}
	var classId string = "license.IksCustomerOp"
	this.ClassId = classId
	var objectType string = "license.IksCustomerOp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseIksCustomerOpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseIksCustomerOpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseIksCustomerOpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseIksCustomerOpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveAdmin returns the ActiveAdmin field value if set, zero value otherwise.
func (o *LicenseIksCustomerOpAllOf) GetActiveAdmin() bool {
	if o == nil || o.ActiveAdmin == nil {
		var ret bool
		return ret
	}
	return *o.ActiveAdmin
}

// GetActiveAdminOk returns a tuple with the ActiveAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetActiveAdminOk() (*bool, bool) {
	if o == nil || o.ActiveAdmin == nil {
		return nil, false
	}
	return o.ActiveAdmin, true
}

// HasActiveAdmin returns a boolean if a field has been set.
func (o *LicenseIksCustomerOpAllOf) HasActiveAdmin() bool {
	if o != nil && o.ActiveAdmin != nil {
		return true
	}

	return false
}

// SetActiveAdmin gets a reference to the given bool and assigns it to the ActiveAdmin field.
func (o *LicenseIksCustomerOpAllOf) SetActiveAdmin(v bool) {
	o.ActiveAdmin = &v
}

// GetEnableTrial returns the EnableTrial field value if set, zero value otherwise.
func (o *LicenseIksCustomerOpAllOf) GetEnableTrial() bool {
	if o == nil || o.EnableTrial == nil {
		var ret bool
		return ret
	}
	return *o.EnableTrial
}

// GetEnableTrialOk returns a tuple with the EnableTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetEnableTrialOk() (*bool, bool) {
	if o == nil || o.EnableTrial == nil {
		return nil, false
	}
	return o.EnableTrial, true
}

// HasEnableTrial returns a boolean if a field has been set.
func (o *LicenseIksCustomerOpAllOf) HasEnableTrial() bool {
	if o != nil && o.EnableTrial != nil {
		return true
	}

	return false
}

// SetEnableTrial gets a reference to the given bool and assigns it to the EnableTrial field.
func (o *LicenseIksCustomerOpAllOf) SetEnableTrial(v bool) {
	o.EnableTrial = &v
}

// GetEvaluationPeriod returns the EvaluationPeriod field value if set, zero value otherwise.
func (o *LicenseIksCustomerOpAllOf) GetEvaluationPeriod() int64 {
	if o == nil || o.EvaluationPeriod == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationPeriod
}

// GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetEvaluationPeriodOk() (*int64, bool) {
	if o == nil || o.EvaluationPeriod == nil {
		return nil, false
	}
	return o.EvaluationPeriod, true
}

// HasEvaluationPeriod returns a boolean if a field has been set.
func (o *LicenseIksCustomerOpAllOf) HasEvaluationPeriod() bool {
	if o != nil && o.EvaluationPeriod != nil {
		return true
	}

	return false
}

// SetEvaluationPeriod gets a reference to the given int64 and assigns it to the EvaluationPeriod field.
func (o *LicenseIksCustomerOpAllOf) SetEvaluationPeriod(v int64) {
	o.EvaluationPeriod = &v
}

// GetExtraEvaluation returns the ExtraEvaluation field value if set, zero value otherwise.
func (o *LicenseIksCustomerOpAllOf) GetExtraEvaluation() int64 {
	if o == nil || o.ExtraEvaluation == nil {
		var ret int64
		return ret
	}
	return *o.ExtraEvaluation
}

// GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetExtraEvaluationOk() (*int64, bool) {
	if o == nil || o.ExtraEvaluation == nil {
		return nil, false
	}
	return o.ExtraEvaluation, true
}

// HasExtraEvaluation returns a boolean if a field has been set.
func (o *LicenseIksCustomerOpAllOf) HasExtraEvaluation() bool {
	if o != nil && o.ExtraEvaluation != nil {
		return true
	}

	return false
}

// SetExtraEvaluation gets a reference to the given int64 and assigns it to the ExtraEvaluation field.
func (o *LicenseIksCustomerOpAllOf) SetExtraEvaluation(v int64) {
	o.ExtraEvaluation = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseIksCustomerOpAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIksCustomerOpAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseIksCustomerOpAllOf) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseIksCustomerOpAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseIksCustomerOpAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActiveAdmin != nil {
		toSerialize["ActiveAdmin"] = o.ActiveAdmin
	}
	if o.EnableTrial != nil {
		toSerialize["EnableTrial"] = o.EnableTrial
	}
	if o.EvaluationPeriod != nil {
		toSerialize["EvaluationPeriod"] = o.EvaluationPeriod
	}
	if o.ExtraEvaluation != nil {
		toSerialize["ExtraEvaluation"] = o.ExtraEvaluation
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseIksCustomerOpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseIksCustomerOpAllOf := _LicenseIksCustomerOpAllOf{}

	if err = json.Unmarshal(bytes, &varLicenseIksCustomerOpAllOf); err == nil {
		*o = LicenseIksCustomerOpAllOf(varLicenseIksCustomerOpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveAdmin")
		delete(additionalProperties, "EnableTrial")
		delete(additionalProperties, "EvaluationPeriod")
		delete(additionalProperties, "ExtraEvaluation")
		delete(additionalProperties, "AccountLicenseData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseIksCustomerOpAllOf struct {
	value *LicenseIksCustomerOpAllOf
	isSet bool
}

func (v NullableLicenseIksCustomerOpAllOf) Get() *LicenseIksCustomerOpAllOf {
	return v.value
}

func (v *NullableLicenseIksCustomerOpAllOf) Set(val *LicenseIksCustomerOpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIksCustomerOpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIksCustomerOpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIksCustomerOpAllOf(val *LicenseIksCustomerOpAllOf) *NullableLicenseIksCustomerOpAllOf {
	return &NullableLicenseIksCustomerOpAllOf{value: val, isSet: true}
}

func (v NullableLicenseIksCustomerOpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIksCustomerOpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}