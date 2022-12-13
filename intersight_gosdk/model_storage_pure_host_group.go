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

// StoragePureHostGroup A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
type StoragePureHostGroup struct {
	StorageBaseHostGroup
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                        `json:"ObjectType"`
	HostNames  []string                      `json:"HostNames,omitempty"`
	Array      *StoragePureArrayRelationship `json:"Array,omitempty"`
	// An array of relationships to storagePureHost resources.
	Hosts                []StoragePureHostRelationship           `json:"Hosts,omitempty"`
	ProtectionGroup      *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StoragePureHostGroup StoragePureHostGroup

// NewStoragePureHostGroup instantiates a new StoragePureHostGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoragePureHostGroup(classId string, objectType string) *StoragePureHostGroup {
	this := StoragePureHostGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStoragePureHostGroupWithDefaults instantiates a new StoragePureHostGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoragePureHostGroupWithDefaults() *StoragePureHostGroup {
	this := StoragePureHostGroup{}
	var classId string = "storage.PureHostGroup"
	this.ClassId = classId
	var objectType string = "storage.PureHostGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StoragePureHostGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StoragePureHostGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StoragePureHostGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StoragePureHostGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostNames returns the HostNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostGroup) GetHostNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HostNames
}

// GetHostNamesOk returns a tuple with the HostNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostGroup) GetHostNamesOk() ([]string, bool) {
	if o == nil || o.HostNames == nil {
		return nil, false
	}
	return o.HostNames, true
}

// HasHostNames returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasHostNames() bool {
	if o != nil && o.HostNames != nil {
		return true
	}

	return false
}

// SetHostNames gets a reference to the given []string and assigns it to the HostNames field.
func (o *StoragePureHostGroup) SetHostNames(v []string) {
	o.HostNames = v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetArray() StoragePureArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StoragePureArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetArrayOk() (*StoragePureArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StoragePureArrayRelationship and assigns it to the Array field.
func (o *StoragePureHostGroup) SetArray(v StoragePureArrayRelationship) {
	o.Array = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StoragePureHostGroup) GetHosts() []StoragePureHostRelationship {
	if o == nil {
		var ret []StoragePureHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StoragePureHostGroup) GetHostsOk() ([]StoragePureHostRelationship, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []StoragePureHostRelationship and assigns it to the Hosts field.
func (o *StoragePureHostGroup) SetHosts(v []StoragePureHostRelationship) {
	o.Hosts = v
}

// GetProtectionGroup returns the ProtectionGroup field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetProtectionGroup() StoragePureProtectionGroupRelationship {
	if o == nil || o.ProtectionGroup == nil {
		var ret StoragePureProtectionGroupRelationship
		return ret
	}
	return *o.ProtectionGroup
}

// GetProtectionGroupOk returns a tuple with the ProtectionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetProtectionGroupOk() (*StoragePureProtectionGroupRelationship, bool) {
	if o == nil || o.ProtectionGroup == nil {
		return nil, false
	}
	return o.ProtectionGroup, true
}

// HasProtectionGroup returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasProtectionGroup() bool {
	if o != nil && o.ProtectionGroup != nil {
		return true
	}

	return false
}

// SetProtectionGroup gets a reference to the given StoragePureProtectionGroupRelationship and assigns it to the ProtectionGroup field.
func (o *StoragePureHostGroup) SetProtectionGroup(v StoragePureProtectionGroupRelationship) {
	o.ProtectionGroup = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StoragePureHostGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoragePureHostGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StoragePureHostGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StoragePureHostGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StoragePureHostGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseHostGroup, errStorageBaseHostGroup := json.Marshal(o.StorageBaseHostGroup)
	if errStorageBaseHostGroup != nil {
		return []byte{}, errStorageBaseHostGroup
	}
	errStorageBaseHostGroup = json.Unmarshal([]byte(serializedStorageBaseHostGroup), &toSerialize)
	if errStorageBaseHostGroup != nil {
		return []byte{}, errStorageBaseHostGroup
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostNames != nil {
		toSerialize["HostNames"] = o.HostNames
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}
	if o.ProtectionGroup != nil {
		toSerialize["ProtectionGroup"] = o.ProtectionGroup
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StoragePureHostGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StoragePureHostGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                        `json:"ObjectType"`
		HostNames  []string                      `json:"HostNames,omitempty"`
		Array      *StoragePureArrayRelationship `json:"Array,omitempty"`
		// An array of relationships to storagePureHost resources.
		Hosts            []StoragePureHostRelationship           `json:"Hosts,omitempty"`
		ProtectionGroup  *StoragePureProtectionGroupRelationship `json:"ProtectionGroup,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship    `json:"RegisteredDevice,omitempty"`
	}

	varStoragePureHostGroupWithoutEmbeddedStruct := StoragePureHostGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStoragePureHostGroupWithoutEmbeddedStruct)
	if err == nil {
		varStoragePureHostGroup := _StoragePureHostGroup{}
		varStoragePureHostGroup.ClassId = varStoragePureHostGroupWithoutEmbeddedStruct.ClassId
		varStoragePureHostGroup.ObjectType = varStoragePureHostGroupWithoutEmbeddedStruct.ObjectType
		varStoragePureHostGroup.HostNames = varStoragePureHostGroupWithoutEmbeddedStruct.HostNames
		varStoragePureHostGroup.Array = varStoragePureHostGroupWithoutEmbeddedStruct.Array
		varStoragePureHostGroup.Hosts = varStoragePureHostGroupWithoutEmbeddedStruct.Hosts
		varStoragePureHostGroup.ProtectionGroup = varStoragePureHostGroupWithoutEmbeddedStruct.ProtectionGroup
		varStoragePureHostGroup.RegisteredDevice = varStoragePureHostGroupWithoutEmbeddedStruct.RegisteredDevice
		*o = StoragePureHostGroup(varStoragePureHostGroup)
	} else {
		return err
	}

	varStoragePureHostGroup := _StoragePureHostGroup{}

	err = json.Unmarshal(bytes, &varStoragePureHostGroup)
	if err == nil {
		o.StorageBaseHostGroup = varStoragePureHostGroup.StorageBaseHostGroup
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HostNames")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "Hosts")
		delete(additionalProperties, "ProtectionGroup")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseHostGroup := reflect.ValueOf(o.StorageBaseHostGroup)
		for i := 0; i < reflectStorageBaseHostGroup.Type().NumField(); i++ {
			t := reflectStorageBaseHostGroup.Type().Field(i)

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

type NullableStoragePureHostGroup struct {
	value *StoragePureHostGroup
	isSet bool
}

func (v NullableStoragePureHostGroup) Get() *StoragePureHostGroup {
	return v.value
}

func (v *NullableStoragePureHostGroup) Set(val *StoragePureHostGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStoragePureHostGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStoragePureHostGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoragePureHostGroup(val *StoragePureHostGroup) *NullableStoragePureHostGroup {
	return &NullableStoragePureHostGroup{value: val, isSet: true}
}

func (v NullableStoragePureHostGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoragePureHostGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
