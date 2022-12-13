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
	"time"
)

// ApplianceMetaManifestAllOf Definition of the list of properties defined in 'appliance.MetaManifest', excluding properties defined in parent classes.
type ApplianceMetaManifestAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Record creation date for explicitly tracking all collections for purging the old records.
	CreationDate *string `json:"CreationDate,omitempty"`
	// Metadata Manifest ImageBundle file description.
	FileDescription *string `json:"FileDescription,omitempty"`
	// File names of the Metadata Manifest ImageBundle as reported by the pvapp portal.
	FileName *string `json:"FileName,omitempty"`
	// The md5 checksum of the Metadata Manifest ImageBundle file.
	FilechkSum *string `json:"FilechkSum,omitempty"`
	// Install date of the Metadata Manifest ImageBundle.
	InstallDate *time.Time `json:"InstallDate,omitempty"`
	// Status reported for the successful installation of the meta data files.
	Status               *string                 `json:"Status,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceMetaManifestAllOf ApplianceMetaManifestAllOf

// NewApplianceMetaManifestAllOf instantiates a new ApplianceMetaManifestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceMetaManifestAllOf(classId string, objectType string) *ApplianceMetaManifestAllOf {
	this := ApplianceMetaManifestAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceMetaManifestAllOfWithDefaults instantiates a new ApplianceMetaManifestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceMetaManifestAllOfWithDefaults() *ApplianceMetaManifestAllOf {
	this := ApplianceMetaManifestAllOf{}
	var classId string = "appliance.MetaManifest"
	this.ClassId = classId
	var objectType string = "appliance.MetaManifest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceMetaManifestAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceMetaManifestAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceMetaManifestAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceMetaManifestAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *ApplianceMetaManifestAllOf) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetFileDescription returns the FileDescription field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetFileDescription() string {
	if o == nil || o.FileDescription == nil {
		var ret string
		return ret
	}
	return *o.FileDescription
}

// GetFileDescriptionOk returns a tuple with the FileDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetFileDescriptionOk() (*string, bool) {
	if o == nil || o.FileDescription == nil {
		return nil, false
	}
	return o.FileDescription, true
}

// HasFileDescription returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasFileDescription() bool {
	if o != nil && o.FileDescription != nil {
		return true
	}

	return false
}

// SetFileDescription gets a reference to the given string and assigns it to the FileDescription field.
func (o *ApplianceMetaManifestAllOf) SetFileDescription(v string) {
	o.FileDescription = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ApplianceMetaManifestAllOf) SetFileName(v string) {
	o.FileName = &v
}

// GetFilechkSum returns the FilechkSum field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetFilechkSum() string {
	if o == nil || o.FilechkSum == nil {
		var ret string
		return ret
	}
	return *o.FilechkSum
}

// GetFilechkSumOk returns a tuple with the FilechkSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetFilechkSumOk() (*string, bool) {
	if o == nil || o.FilechkSum == nil {
		return nil, false
	}
	return o.FilechkSum, true
}

// HasFilechkSum returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasFilechkSum() bool {
	if o != nil && o.FilechkSum != nil {
		return true
	}

	return false
}

// SetFilechkSum gets a reference to the given string and assigns it to the FilechkSum field.
func (o *ApplianceMetaManifestAllOf) SetFilechkSum(v string) {
	o.FilechkSum = &v
}

// GetInstallDate returns the InstallDate field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetInstallDate() time.Time {
	if o == nil || o.InstallDate == nil {
		var ret time.Time
		return ret
	}
	return *o.InstallDate
}

// GetInstallDateOk returns a tuple with the InstallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetInstallDateOk() (*time.Time, bool) {
	if o == nil || o.InstallDate == nil {
		return nil, false
	}
	return o.InstallDate, true
}

// HasInstallDate returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasInstallDate() bool {
	if o != nil && o.InstallDate != nil {
		return true
	}

	return false
}

// SetInstallDate gets a reference to the given time.Time and assigns it to the InstallDate field.
func (o *ApplianceMetaManifestAllOf) SetInstallDate(v time.Time) {
	o.InstallDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceMetaManifestAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceMetaManifestAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetaManifestAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceMetaManifestAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceMetaManifestAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceMetaManifestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.FileDescription != nil {
		toSerialize["FileDescription"] = o.FileDescription
	}
	if o.FileName != nil {
		toSerialize["FileName"] = o.FileName
	}
	if o.FilechkSum != nil {
		toSerialize["FilechkSum"] = o.FilechkSum
	}
	if o.InstallDate != nil {
		toSerialize["InstallDate"] = o.InstallDate
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceMetaManifestAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varApplianceMetaManifestAllOf := _ApplianceMetaManifestAllOf{}

	if err = json.Unmarshal(bytes, &varApplianceMetaManifestAllOf); err == nil {
		*o = ApplianceMetaManifestAllOf(varApplianceMetaManifestAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreationDate")
		delete(additionalProperties, "FileDescription")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "FilechkSum")
		delete(additionalProperties, "InstallDate")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceMetaManifestAllOf struct {
	value *ApplianceMetaManifestAllOf
	isSet bool
}

func (v NullableApplianceMetaManifestAllOf) Get() *ApplianceMetaManifestAllOf {
	return v.value
}

func (v *NullableApplianceMetaManifestAllOf) Set(val *ApplianceMetaManifestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceMetaManifestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceMetaManifestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceMetaManifestAllOf(val *ApplianceMetaManifestAllOf) *NullableApplianceMetaManifestAllOf {
	return &NullableApplianceMetaManifestAllOf{value: val, isSet: true}
}

func (v NullableApplianceMetaManifestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceMetaManifestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
