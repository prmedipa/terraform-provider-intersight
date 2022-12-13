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

// LicenseCustomerOpAllOf Definition of the list of properties defined in 'license.CustomerOp', excluding properties defined in parent classes.
type LicenseCustomerOpAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The license administrative state. Set this property to 'true' to activate the license entitlements.
	ActiveAdmin *bool `json:"ActiveAdmin,omitempty"`
	// Move all licensed devices to default license tier.
	AllDevicesToDefaultTier *bool `json:"AllDevicesToDefaultTier,omitempty"`
	// Clear the status of smart API sync.
	ClearApiSyncStatus *bool `json:"ClearApiSyncStatus,omitempty"`
	// Trigger de-registration/disable.
	DeregisterDevice *bool `json:"DeregisterDevice,omitempty"`
	// Enable trial for Intersight licensing.
	EnableTrial *bool `json:"EnableTrial,omitempty"`
	// The default Trial or Grace period customer is entitled to.
	EvaluationPeriod *int64 `json:"EvaluationPeriod,omitempty"`
	// The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once.
	ExtraEvaluation *int64 `json:"ExtraEvaluation,omitempty"`
	// Trigger renew authorization.
	RenewAuthorization *bool `json:"RenewAuthorization,omitempty"`
	// Trigger renew registration.
	RenewIdCertificate *bool `json:"RenewIdCertificate,omitempty"`
	// Trigger show tech support feature.
	ShowAgentTechSupport *bool                                  `json:"ShowAgentTechSupport,omitempty"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseCustomerOpAllOf LicenseCustomerOpAllOf

// NewLicenseCustomerOpAllOf instantiates a new LicenseCustomerOpAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseCustomerOpAllOf(classId string, objectType string) *LicenseCustomerOpAllOf {
	this := LicenseCustomerOpAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseCustomerOpAllOfWithDefaults instantiates a new LicenseCustomerOpAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseCustomerOpAllOfWithDefaults() *LicenseCustomerOpAllOf {
	this := LicenseCustomerOpAllOf{}
	var classId string = "license.CustomerOp"
	this.ClassId = classId
	var objectType string = "license.CustomerOp"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseCustomerOpAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseCustomerOpAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseCustomerOpAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseCustomerOpAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveAdmin returns the ActiveAdmin field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetActiveAdmin() bool {
	if o == nil || o.ActiveAdmin == nil {
		var ret bool
		return ret
	}
	return *o.ActiveAdmin
}

// GetActiveAdminOk returns a tuple with the ActiveAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetActiveAdminOk() (*bool, bool) {
	if o == nil || o.ActiveAdmin == nil {
		return nil, false
	}
	return o.ActiveAdmin, true
}

// HasActiveAdmin returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasActiveAdmin() bool {
	if o != nil && o.ActiveAdmin != nil {
		return true
	}

	return false
}

// SetActiveAdmin gets a reference to the given bool and assigns it to the ActiveAdmin field.
func (o *LicenseCustomerOpAllOf) SetActiveAdmin(v bool) {
	o.ActiveAdmin = &v
}

// GetAllDevicesToDefaultTier returns the AllDevicesToDefaultTier field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetAllDevicesToDefaultTier() bool {
	if o == nil || o.AllDevicesToDefaultTier == nil {
		var ret bool
		return ret
	}
	return *o.AllDevicesToDefaultTier
}

// GetAllDevicesToDefaultTierOk returns a tuple with the AllDevicesToDefaultTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetAllDevicesToDefaultTierOk() (*bool, bool) {
	if o == nil || o.AllDevicesToDefaultTier == nil {
		return nil, false
	}
	return o.AllDevicesToDefaultTier, true
}

// HasAllDevicesToDefaultTier returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasAllDevicesToDefaultTier() bool {
	if o != nil && o.AllDevicesToDefaultTier != nil {
		return true
	}

	return false
}

// SetAllDevicesToDefaultTier gets a reference to the given bool and assigns it to the AllDevicesToDefaultTier field.
func (o *LicenseCustomerOpAllOf) SetAllDevicesToDefaultTier(v bool) {
	o.AllDevicesToDefaultTier = &v
}

// GetClearApiSyncStatus returns the ClearApiSyncStatus field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetClearApiSyncStatus() bool {
	if o == nil || o.ClearApiSyncStatus == nil {
		var ret bool
		return ret
	}
	return *o.ClearApiSyncStatus
}

// GetClearApiSyncStatusOk returns a tuple with the ClearApiSyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetClearApiSyncStatusOk() (*bool, bool) {
	if o == nil || o.ClearApiSyncStatus == nil {
		return nil, false
	}
	return o.ClearApiSyncStatus, true
}

// HasClearApiSyncStatus returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasClearApiSyncStatus() bool {
	if o != nil && o.ClearApiSyncStatus != nil {
		return true
	}

	return false
}

// SetClearApiSyncStatus gets a reference to the given bool and assigns it to the ClearApiSyncStatus field.
func (o *LicenseCustomerOpAllOf) SetClearApiSyncStatus(v bool) {
	o.ClearApiSyncStatus = &v
}

// GetDeregisterDevice returns the DeregisterDevice field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetDeregisterDevice() bool {
	if o == nil || o.DeregisterDevice == nil {
		var ret bool
		return ret
	}
	return *o.DeregisterDevice
}

// GetDeregisterDeviceOk returns a tuple with the DeregisterDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetDeregisterDeviceOk() (*bool, bool) {
	if o == nil || o.DeregisterDevice == nil {
		return nil, false
	}
	return o.DeregisterDevice, true
}

// HasDeregisterDevice returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasDeregisterDevice() bool {
	if o != nil && o.DeregisterDevice != nil {
		return true
	}

	return false
}

// SetDeregisterDevice gets a reference to the given bool and assigns it to the DeregisterDevice field.
func (o *LicenseCustomerOpAllOf) SetDeregisterDevice(v bool) {
	o.DeregisterDevice = &v
}

// GetEnableTrial returns the EnableTrial field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetEnableTrial() bool {
	if o == nil || o.EnableTrial == nil {
		var ret bool
		return ret
	}
	return *o.EnableTrial
}

// GetEnableTrialOk returns a tuple with the EnableTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetEnableTrialOk() (*bool, bool) {
	if o == nil || o.EnableTrial == nil {
		return nil, false
	}
	return o.EnableTrial, true
}

// HasEnableTrial returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasEnableTrial() bool {
	if o != nil && o.EnableTrial != nil {
		return true
	}

	return false
}

// SetEnableTrial gets a reference to the given bool and assigns it to the EnableTrial field.
func (o *LicenseCustomerOpAllOf) SetEnableTrial(v bool) {
	o.EnableTrial = &v
}

// GetEvaluationPeriod returns the EvaluationPeriod field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetEvaluationPeriod() int64 {
	if o == nil || o.EvaluationPeriod == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationPeriod
}

// GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetEvaluationPeriodOk() (*int64, bool) {
	if o == nil || o.EvaluationPeriod == nil {
		return nil, false
	}
	return o.EvaluationPeriod, true
}

// HasEvaluationPeriod returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasEvaluationPeriod() bool {
	if o != nil && o.EvaluationPeriod != nil {
		return true
	}

	return false
}

// SetEvaluationPeriod gets a reference to the given int64 and assigns it to the EvaluationPeriod field.
func (o *LicenseCustomerOpAllOf) SetEvaluationPeriod(v int64) {
	o.EvaluationPeriod = &v
}

// GetExtraEvaluation returns the ExtraEvaluation field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetExtraEvaluation() int64 {
	if o == nil || o.ExtraEvaluation == nil {
		var ret int64
		return ret
	}
	return *o.ExtraEvaluation
}

// GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetExtraEvaluationOk() (*int64, bool) {
	if o == nil || o.ExtraEvaluation == nil {
		return nil, false
	}
	return o.ExtraEvaluation, true
}

// HasExtraEvaluation returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasExtraEvaluation() bool {
	if o != nil && o.ExtraEvaluation != nil {
		return true
	}

	return false
}

// SetExtraEvaluation gets a reference to the given int64 and assigns it to the ExtraEvaluation field.
func (o *LicenseCustomerOpAllOf) SetExtraEvaluation(v int64) {
	o.ExtraEvaluation = &v
}

// GetRenewAuthorization returns the RenewAuthorization field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetRenewAuthorization() bool {
	if o == nil || o.RenewAuthorization == nil {
		var ret bool
		return ret
	}
	return *o.RenewAuthorization
}

// GetRenewAuthorizationOk returns a tuple with the RenewAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetRenewAuthorizationOk() (*bool, bool) {
	if o == nil || o.RenewAuthorization == nil {
		return nil, false
	}
	return o.RenewAuthorization, true
}

// HasRenewAuthorization returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasRenewAuthorization() bool {
	if o != nil && o.RenewAuthorization != nil {
		return true
	}

	return false
}

// SetRenewAuthorization gets a reference to the given bool and assigns it to the RenewAuthorization field.
func (o *LicenseCustomerOpAllOf) SetRenewAuthorization(v bool) {
	o.RenewAuthorization = &v
}

// GetRenewIdCertificate returns the RenewIdCertificate field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetRenewIdCertificate() bool {
	if o == nil || o.RenewIdCertificate == nil {
		var ret bool
		return ret
	}
	return *o.RenewIdCertificate
}

// GetRenewIdCertificateOk returns a tuple with the RenewIdCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetRenewIdCertificateOk() (*bool, bool) {
	if o == nil || o.RenewIdCertificate == nil {
		return nil, false
	}
	return o.RenewIdCertificate, true
}

// HasRenewIdCertificate returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasRenewIdCertificate() bool {
	if o != nil && o.RenewIdCertificate != nil {
		return true
	}

	return false
}

// SetRenewIdCertificate gets a reference to the given bool and assigns it to the RenewIdCertificate field.
func (o *LicenseCustomerOpAllOf) SetRenewIdCertificate(v bool) {
	o.RenewIdCertificate = &v
}

// GetShowAgentTechSupport returns the ShowAgentTechSupport field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetShowAgentTechSupport() bool {
	if o == nil || o.ShowAgentTechSupport == nil {
		var ret bool
		return ret
	}
	return *o.ShowAgentTechSupport
}

// GetShowAgentTechSupportOk returns a tuple with the ShowAgentTechSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetShowAgentTechSupportOk() (*bool, bool) {
	if o == nil || o.ShowAgentTechSupport == nil {
		return nil, false
	}
	return o.ShowAgentTechSupport, true
}

// HasShowAgentTechSupport returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasShowAgentTechSupport() bool {
	if o != nil && o.ShowAgentTechSupport != nil {
		return true
	}

	return false
}

// SetShowAgentTechSupport gets a reference to the given bool and assigns it to the ShowAgentTechSupport field.
func (o *LicenseCustomerOpAllOf) SetShowAgentTechSupport(v bool) {
	o.ShowAgentTechSupport = &v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseCustomerOpAllOf) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseCustomerOpAllOf) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseCustomerOpAllOf) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseCustomerOpAllOf) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseCustomerOpAllOf) MarshalJSON() ([]byte, error) {
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
	if o.AllDevicesToDefaultTier != nil {
		toSerialize["AllDevicesToDefaultTier"] = o.AllDevicesToDefaultTier
	}
	if o.ClearApiSyncStatus != nil {
		toSerialize["ClearApiSyncStatus"] = o.ClearApiSyncStatus
	}
	if o.DeregisterDevice != nil {
		toSerialize["DeregisterDevice"] = o.DeregisterDevice
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
	if o.RenewAuthorization != nil {
		toSerialize["RenewAuthorization"] = o.RenewAuthorization
	}
	if o.RenewIdCertificate != nil {
		toSerialize["RenewIdCertificate"] = o.RenewIdCertificate
	}
	if o.ShowAgentTechSupport != nil {
		toSerialize["ShowAgentTechSupport"] = o.ShowAgentTechSupport
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseCustomerOpAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLicenseCustomerOpAllOf := _LicenseCustomerOpAllOf{}

	if err = json.Unmarshal(bytes, &varLicenseCustomerOpAllOf); err == nil {
		*o = LicenseCustomerOpAllOf(varLicenseCustomerOpAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveAdmin")
		delete(additionalProperties, "AllDevicesToDefaultTier")
		delete(additionalProperties, "ClearApiSyncStatus")
		delete(additionalProperties, "DeregisterDevice")
		delete(additionalProperties, "EnableTrial")
		delete(additionalProperties, "EvaluationPeriod")
		delete(additionalProperties, "ExtraEvaluation")
		delete(additionalProperties, "RenewAuthorization")
		delete(additionalProperties, "RenewIdCertificate")
		delete(additionalProperties, "ShowAgentTechSupport")
		delete(additionalProperties, "AccountLicenseData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseCustomerOpAllOf struct {
	value *LicenseCustomerOpAllOf
	isSet bool
}

func (v NullableLicenseCustomerOpAllOf) Get() *LicenseCustomerOpAllOf {
	return v.value
}

func (v *NullableLicenseCustomerOpAllOf) Set(val *LicenseCustomerOpAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseCustomerOpAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseCustomerOpAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseCustomerOpAllOf(val *LicenseCustomerOpAllOf) *NullableLicenseCustomerOpAllOf {
	return &NullableLicenseCustomerOpAllOf{value: val, isSet: true}
}

func (v NullableLicenseCustomerOpAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseCustomerOpAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
