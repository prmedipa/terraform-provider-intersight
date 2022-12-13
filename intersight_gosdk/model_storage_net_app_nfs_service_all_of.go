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

// StorageNetAppNfsServiceAllOf Definition of the list of properties defined in 'storage.NetAppNfsService', excluding properties defined in parent classes.
type StorageNetAppNfsServiceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies whether NFSv3 protocol is enabled.
	NfsV3Enabled *bool `json:"NfsV3Enabled,omitempty"`
	// Specifies whether NFSv4.1 protocol is enabled.
	NfsV41Enabled *bool `json:"NfsV41Enabled,omitempty"`
	// Specifies whether NFSv4.0 protocol is enabled.
	NfsV4Enabled *bool `json:"NfsV4Enabled,omitempty"`
	// Unique identifier for the NetApp Storage Virtual Machine.
	SvmUuid              *string                             `json:"SvmUuid,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNfsServiceAllOf StorageNetAppNfsServiceAllOf

// NewStorageNetAppNfsServiceAllOf instantiates a new StorageNetAppNfsServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNfsServiceAllOf(classId string, objectType string) *StorageNetAppNfsServiceAllOf {
	this := StorageNetAppNfsServiceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNfsServiceAllOfWithDefaults instantiates a new StorageNetAppNfsServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNfsServiceAllOfWithDefaults() *StorageNetAppNfsServiceAllOf {
	this := StorageNetAppNfsServiceAllOf{}
	var classId string = "storage.NetAppNfsService"
	this.ClassId = classId
	var objectType string = "storage.NetAppNfsService"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNfsServiceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNfsServiceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNfsServiceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNfsServiceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNfsV3Enabled returns the NfsV3Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV3Enabled() bool {
	if o == nil || o.NfsV3Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV3Enabled
}

// GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV3EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV3Enabled == nil {
		return nil, false
	}
	return o.NfsV3Enabled, true
}

// HasNfsV3Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsServiceAllOf) HasNfsV3Enabled() bool {
	if o != nil && o.NfsV3Enabled != nil {
		return true
	}

	return false
}

// SetNfsV3Enabled gets a reference to the given bool and assigns it to the NfsV3Enabled field.
func (o *StorageNetAppNfsServiceAllOf) SetNfsV3Enabled(v bool) {
	o.NfsV3Enabled = &v
}

// GetNfsV41Enabled returns the NfsV41Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV41Enabled() bool {
	if o == nil || o.NfsV41Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV41Enabled
}

// GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV41EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV41Enabled == nil {
		return nil, false
	}
	return o.NfsV41Enabled, true
}

// HasNfsV41Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsServiceAllOf) HasNfsV41Enabled() bool {
	if o != nil && o.NfsV41Enabled != nil {
		return true
	}

	return false
}

// SetNfsV41Enabled gets a reference to the given bool and assigns it to the NfsV41Enabled field.
func (o *StorageNetAppNfsServiceAllOf) SetNfsV41Enabled(v bool) {
	o.NfsV41Enabled = &v
}

// GetNfsV4Enabled returns the NfsV4Enabled field value if set, zero value otherwise.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV4Enabled() bool {
	if o == nil || o.NfsV4Enabled == nil {
		var ret bool
		return ret
	}
	return *o.NfsV4Enabled
}

// GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetNfsV4EnabledOk() (*bool, bool) {
	if o == nil || o.NfsV4Enabled == nil {
		return nil, false
	}
	return o.NfsV4Enabled, true
}

// HasNfsV4Enabled returns a boolean if a field has been set.
func (o *StorageNetAppNfsServiceAllOf) HasNfsV4Enabled() bool {
	if o != nil && o.NfsV4Enabled != nil {
		return true
	}

	return false
}

// SetNfsV4Enabled gets a reference to the given bool and assigns it to the NfsV4Enabled field.
func (o *StorageNetAppNfsServiceAllOf) SetNfsV4Enabled(v bool) {
	o.NfsV4Enabled = &v
}

// GetSvmUuid returns the SvmUuid field value if set, zero value otherwise.
func (o *StorageNetAppNfsServiceAllOf) GetSvmUuid() string {
	if o == nil || o.SvmUuid == nil {
		var ret string
		return ret
	}
	return *o.SvmUuid
}

// GetSvmUuidOk returns a tuple with the SvmUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetSvmUuidOk() (*string, bool) {
	if o == nil || o.SvmUuid == nil {
		return nil, false
	}
	return o.SvmUuid, true
}

// HasSvmUuid returns a boolean if a field has been set.
func (o *StorageNetAppNfsServiceAllOf) HasSvmUuid() bool {
	if o != nil && o.SvmUuid != nil {
		return true
	}

	return false
}

// SetSvmUuid gets a reference to the given string and assigns it to the SvmUuid field.
func (o *StorageNetAppNfsServiceAllOf) SetSvmUuid(v string) {
	o.SvmUuid = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppNfsServiceAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNfsServiceAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppNfsServiceAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppNfsServiceAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppNfsServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NfsV3Enabled != nil {
		toSerialize["NfsV3Enabled"] = o.NfsV3Enabled
	}
	if o.NfsV41Enabled != nil {
		toSerialize["NfsV41Enabled"] = o.NfsV41Enabled
	}
	if o.NfsV4Enabled != nil {
		toSerialize["NfsV4Enabled"] = o.NfsV4Enabled
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

func (o *StorageNetAppNfsServiceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppNfsServiceAllOf := _StorageNetAppNfsServiceAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppNfsServiceAllOf); err == nil {
		*o = StorageNetAppNfsServiceAllOf(varStorageNetAppNfsServiceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NfsV3Enabled")
		delete(additionalProperties, "NfsV41Enabled")
		delete(additionalProperties, "NfsV4Enabled")
		delete(additionalProperties, "SvmUuid")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppNfsServiceAllOf struct {
	value *StorageNetAppNfsServiceAllOf
	isSet bool
}

func (v NullableStorageNetAppNfsServiceAllOf) Get() *StorageNetAppNfsServiceAllOf {
	return v.value
}

func (v *NullableStorageNetAppNfsServiceAllOf) Set(val *StorageNetAppNfsServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNfsServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNfsServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNfsServiceAllOf(val *StorageNetAppNfsServiceAllOf) *NullableStorageNetAppNfsServiceAllOf {
	return &NullableStorageNetAppNfsServiceAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppNfsServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNfsServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
