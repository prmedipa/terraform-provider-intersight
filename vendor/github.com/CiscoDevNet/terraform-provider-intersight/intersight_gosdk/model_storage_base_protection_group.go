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
	"reflect"
	"strings"
)

// StorageBaseProtectionGroup A protection group contains list of members that are protected together through snapshots with point-in-time consistency across the members. Members within the protection group have common data protection requirements and also the same snapshot, replication, and retention schedules.
type StorageBaseProtectionGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the protection Group.
	Name *string `json:"Name,omitempty"`
	// Prefix used for all generated snapshots from the protection group.
	Prefix *string `json:"Prefix,omitempty"`
	// Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set.
	ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`
	// Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set.
	SnapshotEnabled      *bool `json:"SnapshotEnabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageBaseProtectionGroup StorageBaseProtectionGroup

// NewStorageBaseProtectionGroup instantiates a new StorageBaseProtectionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageBaseProtectionGroup(classId string, objectType string) *StorageBaseProtectionGroup {
	this := StorageBaseProtectionGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageBaseProtectionGroupWithDefaults instantiates a new StorageBaseProtectionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageBaseProtectionGroupWithDefaults() *StorageBaseProtectionGroup {
	this := StorageBaseProtectionGroup{}
	var classId string = "storage.PureProtectionGroup"
	this.ClassId = classId
	var objectType string = "storage.PureProtectionGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageBaseProtectionGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageBaseProtectionGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageBaseProtectionGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageBaseProtectionGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageBaseProtectionGroup) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroup) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroup) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *StorageBaseProtectionGroup) SetPrefix(v string) {
	o.Prefix = &v
}

// GetReplicationEnabled returns the ReplicationEnabled field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroup) GetReplicationEnabled() bool {
	if o == nil || o.ReplicationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ReplicationEnabled
}

// GetReplicationEnabledOk returns a tuple with the ReplicationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetReplicationEnabledOk() (*bool, bool) {
	if o == nil || o.ReplicationEnabled == nil {
		return nil, false
	}
	return o.ReplicationEnabled, true
}

// HasReplicationEnabled returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroup) HasReplicationEnabled() bool {
	if o != nil && o.ReplicationEnabled != nil {
		return true
	}

	return false
}

// SetReplicationEnabled gets a reference to the given bool and assigns it to the ReplicationEnabled field.
func (o *StorageBaseProtectionGroup) SetReplicationEnabled(v bool) {
	o.ReplicationEnabled = &v
}

// GetSnapshotEnabled returns the SnapshotEnabled field value if set, zero value otherwise.
func (o *StorageBaseProtectionGroup) GetSnapshotEnabled() bool {
	if o == nil || o.SnapshotEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotEnabled
}

// GetSnapshotEnabledOk returns a tuple with the SnapshotEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageBaseProtectionGroup) GetSnapshotEnabledOk() (*bool, bool) {
	if o == nil || o.SnapshotEnabled == nil {
		return nil, false
	}
	return o.SnapshotEnabled, true
}

// HasSnapshotEnabled returns a boolean if a field has been set.
func (o *StorageBaseProtectionGroup) HasSnapshotEnabled() bool {
	if o != nil && o.SnapshotEnabled != nil {
		return true
	}

	return false
}

// SetSnapshotEnabled gets a reference to the given bool and assigns it to the SnapshotEnabled field.
func (o *StorageBaseProtectionGroup) SetSnapshotEnabled(v bool) {
	o.SnapshotEnabled = &v
}

func (o StorageBaseProtectionGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Prefix != nil {
		toSerialize["Prefix"] = o.Prefix
	}
	if o.ReplicationEnabled != nil {
		toSerialize["ReplicationEnabled"] = o.ReplicationEnabled
	}
	if o.SnapshotEnabled != nil {
		toSerialize["SnapshotEnabled"] = o.SnapshotEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageBaseProtectionGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageBaseProtectionGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the protection Group.
		Name *string `json:"Name,omitempty"`
		// Prefix used for all generated snapshots from the protection group.
		Prefix *string `json:"Prefix,omitempty"`
		// Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set.
		ReplicationEnabled *bool `json:"ReplicationEnabled,omitempty"`
		// Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set.
		SnapshotEnabled *bool `json:"SnapshotEnabled,omitempty"`
	}

	varStorageBaseProtectionGroupWithoutEmbeddedStruct := StorageBaseProtectionGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageBaseProtectionGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageBaseProtectionGroup := _StorageBaseProtectionGroup{}
		varStorageBaseProtectionGroup.ClassId = varStorageBaseProtectionGroupWithoutEmbeddedStruct.ClassId
		varStorageBaseProtectionGroup.ObjectType = varStorageBaseProtectionGroupWithoutEmbeddedStruct.ObjectType
		varStorageBaseProtectionGroup.Name = varStorageBaseProtectionGroupWithoutEmbeddedStruct.Name
		varStorageBaseProtectionGroup.Prefix = varStorageBaseProtectionGroupWithoutEmbeddedStruct.Prefix
		varStorageBaseProtectionGroup.ReplicationEnabled = varStorageBaseProtectionGroupWithoutEmbeddedStruct.ReplicationEnabled
		varStorageBaseProtectionGroup.SnapshotEnabled = varStorageBaseProtectionGroupWithoutEmbeddedStruct.SnapshotEnabled
		*o = StorageBaseProtectionGroup(varStorageBaseProtectionGroup)
	} else {
		return err
	}

	varStorageBaseProtectionGroup := _StorageBaseProtectionGroup{}

	err = json.Unmarshal(bytes, &varStorageBaseProtectionGroup)
	if err == nil {
		o.MoBaseMo = varStorageBaseProtectionGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Prefix")
		delete(additionalProperties, "ReplicationEnabled")
		delete(additionalProperties, "SnapshotEnabled")

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

type NullableStorageBaseProtectionGroup struct {
	value *StorageBaseProtectionGroup
	isSet bool
}

func (v NullableStorageBaseProtectionGroup) Get() *StorageBaseProtectionGroup {
	return v.value
}

func (v *NullableStorageBaseProtectionGroup) Set(val *StorageBaseProtectionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageBaseProtectionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageBaseProtectionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageBaseProtectionGroup(val *StorageBaseProtectionGroup) *NullableStorageBaseProtectionGroup {
	return &NullableStorageBaseProtectionGroup{value: val, isSet: true}
}

func (v NullableStorageBaseProtectionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageBaseProtectionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
