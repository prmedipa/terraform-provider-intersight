/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-11765
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// StorageNetAppCifsServiceAllOf Definition of the list of properties defined in 'storage.NetAppCifsService', excluding properties defined in parent classes.
type StorageNetAppCifsServiceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The fully qualified domain name of the Windows Active Directory to which this CIFS server belongs.
	ActiveDirectoryFqdn *string `json:"ActiveDirectoryFqdn,omitempty"`
	// Identifies the organizational unit within the Active Directory domain to associate with the CIFS server.
	AdOrganizationalUnit *string `json:"AdOrganizationalUnit,omitempty"`
	// A descriptive text comment for the CIFS server.
	Comment *string `json:"Comment,omitempty"`
	// Indicates that the CIFS service is administratively enabled.
	Enabled *string `json:"Enabled,omitempty"`
	// Name of the NetApp CIFS server.
	ServerName *string `json:"ServerName,omitempty"`
	// The storage virtual machine name for the CIFS service.
	SvmName *string `json:"SvmName,omitempty"`
	// Unique identifier for the NetApp Storage Virtual Machine.
	SvmUuid              *string                             `json:"SvmUuid,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppCifsServiceAllOf StorageNetAppCifsServiceAllOf

// NewStorageNetAppCifsServiceAllOf instantiates a new StorageNetAppCifsServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppCifsServiceAllOf(classId string, objectType string) *StorageNetAppCifsServiceAllOf {
	this := StorageNetAppCifsServiceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppCifsServiceAllOfWithDefaults instantiates a new StorageNetAppCifsServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppCifsServiceAllOfWithDefaults() *StorageNetAppCifsServiceAllOf {
	this := StorageNetAppCifsServiceAllOf{}
	var classId string = "storage.NetAppCifsService"
	this.ClassId = classId
	var objectType string = "storage.NetAppCifsService"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppCifsServiceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppCifsServiceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppCifsServiceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppCifsServiceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActiveDirectoryFqdn returns the ActiveDirectoryFqdn field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetActiveDirectoryFqdn() string {
	if o == nil || o.ActiveDirectoryFqdn == nil {
		var ret string
		return ret
	}
	return *o.ActiveDirectoryFqdn
}

// GetActiveDirectoryFqdnOk returns a tuple with the ActiveDirectoryFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetActiveDirectoryFqdnOk() (*string, bool) {
	if o == nil || o.ActiveDirectoryFqdn == nil {
		return nil, false
	}
	return o.ActiveDirectoryFqdn, true
}

// HasActiveDirectoryFqdn returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasActiveDirectoryFqdn() bool {
	if o != nil && o.ActiveDirectoryFqdn != nil {
		return true
	}

	return false
}

// SetActiveDirectoryFqdn gets a reference to the given string and assigns it to the ActiveDirectoryFqdn field.
func (o *StorageNetAppCifsServiceAllOf) SetActiveDirectoryFqdn(v string) {
	o.ActiveDirectoryFqdn = &v
}

// GetAdOrganizationalUnit returns the AdOrganizationalUnit field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetAdOrganizationalUnit() string {
	if o == nil || o.AdOrganizationalUnit == nil {
		var ret string
		return ret
	}
	return *o.AdOrganizationalUnit
}

// GetAdOrganizationalUnitOk returns a tuple with the AdOrganizationalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetAdOrganizationalUnitOk() (*string, bool) {
	if o == nil || o.AdOrganizationalUnit == nil {
		return nil, false
	}
	return o.AdOrganizationalUnit, true
}

// HasAdOrganizationalUnit returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasAdOrganizationalUnit() bool {
	if o != nil && o.AdOrganizationalUnit != nil {
		return true
	}

	return false
}

// SetAdOrganizationalUnit gets a reference to the given string and assigns it to the AdOrganizationalUnit field.
func (o *StorageNetAppCifsServiceAllOf) SetAdOrganizationalUnit(v string) {
	o.AdOrganizationalUnit = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *StorageNetAppCifsServiceAllOf) SetComment(v string) {
	o.Comment = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetEnabled() string {
	if o == nil || o.Enabled == nil {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetEnabledOk() (*string, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *StorageNetAppCifsServiceAllOf) SetEnabled(v string) {
	o.Enabled = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *StorageNetAppCifsServiceAllOf) SetServerName(v string) {
	o.ServerName = &v
}

// GetSvmName returns the SvmName field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetSvmName() string {
	if o == nil || o.SvmName == nil {
		var ret string
		return ret
	}
	return *o.SvmName
}

// GetSvmNameOk returns a tuple with the SvmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetSvmNameOk() (*string, bool) {
	if o == nil || o.SvmName == nil {
		return nil, false
	}
	return o.SvmName, true
}

// HasSvmName returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasSvmName() bool {
	if o != nil && o.SvmName != nil {
		return true
	}

	return false
}

// SetSvmName gets a reference to the given string and assigns it to the SvmName field.
func (o *StorageNetAppCifsServiceAllOf) SetSvmName(v string) {
	o.SvmName = &v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetSvmUuid() string {
	if o == nil || o.SvmUuid == nil {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetSvmUuidOk() (*string, bool) {
	if o == nil || o.SvmUuid == nil {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasSvmUuid() bool {
	if o != nil && o.SvmUuid != nil {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppCifsServiceAllOf) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppCifsServiceAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppCifsServiceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppCifsServiceAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppCifsServiceAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppCifsServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ActiveDirectoryFqdn != nil {
		toSerialize["ActiveDirectoryFqdn"] = o.ActiveDirectoryFqdn
	}
	if o.AdOrganizationalUnit != nil {
		toSerialize["AdOrganizationalUnit"] = o.AdOrganizationalUnit
	}
	if o.Comment != nil {
		toSerialize["Comment"] = o.Comment
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.ServerName != nil {
		toSerialize["ServerName"] = o.ServerName
	}
	if o.SvmName != nil {
		toSerialize["SvmName"] = o.SvmName
	}
	if o.SvmUuid != nil {
		toSerialize["SvmUuid"] = o.SvmUuid
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppCifsServiceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppCifsServiceAllOf := _StorageNetAppCifsServiceAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppCifsServiceAllOf); err == nil {
		*o = StorageNetAppCifsServiceAllOf(varStorageNetAppCifsServiceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ActiveDirectoryFqdn")
		delete(additionalProperties, "AdOrganizationalUnit")
		delete(additionalProperties, "Comment")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "ServerName")
		delete(additionalProperties, "SvmName")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppCifsServiceAllOf struct {
	value *StorageNetAppCifsServiceAllOf
	isSet bool
}

func (v NullableStorageNetAppCifsServiceAllOf) Get() *StorageNetAppCifsServiceAllOf {
	return v.value
}

func (v *NullableStorageNetAppCifsServiceAllOf) Set(val *StorageNetAppCifsServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppCifsServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppCifsServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppCifsServiceAllOf(val *StorageNetAppCifsServiceAllOf) *NullableStorageNetAppCifsServiceAllOf {
	return &NullableStorageNetAppCifsServiceAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppCifsServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppCifsServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}