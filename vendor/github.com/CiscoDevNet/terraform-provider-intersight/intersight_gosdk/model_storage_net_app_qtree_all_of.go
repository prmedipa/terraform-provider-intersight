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

// StorageNetAppQtreeAllOf Definition of the list of properties defined in 'storage.NetAppQtree', excluding properties defined in parent classes.
type StorageNetAppQtreeAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identifier of NetApp export policy.
	ExportPolicyId *string `json:"ExportPolicyId,omitempty"`
	// Name of the NetApp Qtree.
	Name *string `json:"Name,omitempty"`
	// Client visible path to the qtree.
	Path *string `json:"Path,omitempty"`
	// Identifies the UNIX permissions for the qtree.
	Permission *string `json:"Permission,omitempty"`
	// NetApp Qtree ID, unique within the qtree's volume.
	QtreeId *int64 `json:"QtreeId,omitempty"`
	// Identifies the security style for the qtree, it determines how access to the qtree is controlled. * `UNIX` - Security style for UNIX uid, gid and mode bits. * `NTFS` - Security style for CIFS ACLs. * `Mixed` - Security style for NFS and CIFS access.
	SecurityStyle *string `json:"SecurityStyle,omitempty"`
	// NetApp Volume uuid, unique identifier for the NetApp volume.
	VolumeUuid           *string                             `json:"VolumeUuid,omitempty"`
	StorageContainer     *StorageNetAppVolumeRelationship    `json:"StorageContainer,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppQtreeAllOf StorageNetAppQtreeAllOf

// NewStorageNetAppQtreeAllOf instantiates a new StorageNetAppQtreeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppQtreeAllOf(classId string, objectType string) *StorageNetAppQtreeAllOf {
	this := StorageNetAppQtreeAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppQtreeAllOfWithDefaults instantiates a new StorageNetAppQtreeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppQtreeAllOfWithDefaults() *StorageNetAppQtreeAllOf {
	this := StorageNetAppQtreeAllOf{}
	var classId string = "storage.NetAppQtree"
	this.ClassId = classId
	var objectType string = "storage.NetAppQtree"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppQtreeAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppQtreeAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppQtreeAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppQtreeAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExportPolicyId returns the ExportPolicyId field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetExportPolicyId() string {
	if o == nil || o.ExportPolicyId == nil {
		var ret string
		return ret
	}
	return *o.ExportPolicyId
}

// GetExportPolicyIdOk returns a tuple with the ExportPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetExportPolicyIdOk() (*string, bool) {
	if o == nil || o.ExportPolicyId == nil {
		return nil, false
	}
	return o.ExportPolicyId, true
}

// HasExportPolicyId returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasExportPolicyId() bool {
	if o != nil && o.ExportPolicyId != nil {
		return true
	}

	return false
}

// SetExportPolicyId gets a reference to the given string and assigns it to the ExportPolicyId field.
func (o *StorageNetAppQtreeAllOf) SetExportPolicyId(v string) {
	o.ExportPolicyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppQtreeAllOf) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *StorageNetAppQtreeAllOf) SetPath(v string) {
	o.Path = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *StorageNetAppQtreeAllOf) SetPermission(v string) {
	o.Permission = &v
}

// GetQtreeId returns the QtreeId field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetQtreeId() int64 {
	if o == nil || o.QtreeId == nil {
		var ret int64
		return ret
	}
	return *o.QtreeId
}

// GetQtreeIdOk returns a tuple with the QtreeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetQtreeIdOk() (*int64, bool) {
	if o == nil || o.QtreeId == nil {
		return nil, false
	}
	return o.QtreeId, true
}

// HasQtreeId returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasQtreeId() bool {
	if o != nil && o.QtreeId != nil {
		return true
	}

	return false
}

// SetQtreeId gets a reference to the given int64 and assigns it to the QtreeId field.
func (o *StorageNetAppQtreeAllOf) SetQtreeId(v int64) {
	o.QtreeId = &v
}

// GetSecurityStyle returns the SecurityStyle field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetSecurityStyle() string {
	if o == nil || o.SecurityStyle == nil {
		var ret string
		return ret
	}
	return *o.SecurityStyle
}

// GetSecurityStyleOk returns a tuple with the SecurityStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetSecurityStyleOk() (*string, bool) {
	if o == nil || o.SecurityStyle == nil {
		return nil, false
	}
	return o.SecurityStyle, true
}

// HasSecurityStyle returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasSecurityStyle() bool {
	if o != nil && o.SecurityStyle != nil {
		return true
	}

	return false
}

// SetSecurityStyle gets a reference to the given string and assigns it to the SecurityStyle field.
func (o *StorageNetAppQtreeAllOf) SetSecurityStyle(v string) {
	o.SecurityStyle = &v
}

// GetVolumeUuid returns the VolumeUuid field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetVolumeUuid() string {
	if o == nil || o.VolumeUuid == nil {
		var ret string
		return ret
	}
	return *o.VolumeUuid
}

// GetVolumeUuidOk returns a tuple with the VolumeUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetVolumeUuidOk() (*string, bool) {
	if o == nil || o.VolumeUuid == nil {
		return nil, false
	}
	return o.VolumeUuid, true
}

// HasVolumeUuid returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasVolumeUuid() bool {
	if o != nil && o.VolumeUuid != nil {
		return true
	}

	return false
}

// SetVolumeUuid gets a reference to the given string and assigns it to the VolumeUuid field.
func (o *StorageNetAppQtreeAllOf) SetVolumeUuid(v string) {
	o.VolumeUuid = &v
}

// GetStorageContainer returns the StorageContainer field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetStorageContainer() StorageNetAppVolumeRelationship {
	if o == nil || o.StorageContainer == nil {
		var ret StorageNetAppVolumeRelationship
		return ret
	}
	return *o.StorageContainer
}

// GetStorageContainerOk returns a tuple with the StorageContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool) {
	if o == nil || o.StorageContainer == nil {
		return nil, false
	}
	return o.StorageContainer, true
}

// HasStorageContainer returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasStorageContainer() bool {
	if o != nil && o.StorageContainer != nil {
		return true
	}

	return false
}

// SetStorageContainer gets a reference to the given StorageNetAppVolumeRelationship and assigns it to the StorageContainer field.
func (o *StorageNetAppQtreeAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship) {
	o.StorageContainer = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppQtreeAllOf) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppQtreeAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppQtreeAllOf) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppQtreeAllOf) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppQtreeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExportPolicyId != nil {
		toSerialize["ExportPolicyId"] = o.ExportPolicyId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.QtreeId != nil {
		toSerialize["QtreeId"] = o.QtreeId
	}
	if o.SecurityStyle != nil {
		toSerialize["SecurityStyle"] = o.SecurityStyle
	}
	if o.VolumeUuid != nil {
		toSerialize["VolumeUuid"] = o.VolumeUuid
	}
	if o.StorageContainer != nil {
		toSerialize["StorageContainer"] = o.StorageContainer
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppQtreeAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageNetAppQtreeAllOf := _StorageNetAppQtreeAllOf{}

	if err = json.Unmarshal(bytes, &varStorageNetAppQtreeAllOf); err == nil {
		*o = StorageNetAppQtreeAllOf(varStorageNetAppQtreeAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExportPolicyId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "QtreeId")
		delete(additionalProperties, "SecurityStyle")
		delete(additionalProperties, "VolumeUuid")
		delete(additionalProperties, "StorageContainer")
		delete(additionalProperties, "Tenant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageNetAppQtreeAllOf struct {
	value *StorageNetAppQtreeAllOf
	isSet bool
}

func (v NullableStorageNetAppQtreeAllOf) Get() *StorageNetAppQtreeAllOf {
	return v.value
}

func (v *NullableStorageNetAppQtreeAllOf) Set(val *StorageNetAppQtreeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppQtreeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppQtreeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppQtreeAllOf(val *StorageNetAppQtreeAllOf) *NullableStorageNetAppQtreeAllOf {
	return &NullableStorageNetAppQtreeAllOf{value: val, isSet: true}
}

func (v NullableStorageNetAppQtreeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppQtreeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
