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

// checks if the StorageHyperFlexStorageContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHyperFlexStorageContainer{}

// StorageHyperFlexStorageContainer A Storage Container (Datastore) entity.
type StorageHyperFlexStorageContainer struct {
	StorageBaseStorageContainer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Storage Container data block size
	DataBlockSize *int64 `json:"DataBlockSize,omitempty"`
	// Indicates whether the Storage Container has Volumes.
	InUse *bool `json:"InUse,omitempty"`
	// Storage container's last access time.
	LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
	// Storage container's last modified time.
	LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
	// Provisioned Capacity of the Storage container.
	ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
	// Provisioned Capacity Utilization of All Volumes associated with the Storage Container.
	ProvisionedVolumeCapacityUtilization *float32 `json:"ProvisionedVolumeCapacityUtilization,omitempty"`
	// Storage Container type (SMB/NFS/iSCSI). * `NFS` - Storage container created/accesed through NFS protocol. * `SMB` - Storage container created/accessed through SMB protocol. * `iSCSI` - Storage container created/accessed through iSCSI protocol.
	Type *string `json:"Type,omitempty"`
	// Uncompressed bytes on Storage Container.
	UnCompressedUsedBytes *int64 `json:"UnCompressedUsedBytes,omitempty"`
	// UUID of the Datastore/Storage Containter.
	Uuid *string `json:"Uuid,omitempty"`
	// Number of Volumes associated with the Storage Container.
	VolumeCount          *int64                                      `json:"VolumeCount,omitempty"`
	Cluster              NullableHyperflexClusterRelationship        `json:"Cluster,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHyperFlexStorageContainer StorageHyperFlexStorageContainer

// NewStorageHyperFlexStorageContainer instantiates a new StorageHyperFlexStorageContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHyperFlexStorageContainer(classId string, objectType string) *StorageHyperFlexStorageContainer {
	this := StorageHyperFlexStorageContainer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHyperFlexStorageContainerWithDefaults instantiates a new StorageHyperFlexStorageContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHyperFlexStorageContainerWithDefaults() *StorageHyperFlexStorageContainer {
	this := StorageHyperFlexStorageContainer{}
	var classId string = "storage.HyperFlexStorageContainer"
	this.ClassId = classId
	var objectType string = "storage.HyperFlexStorageContainer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHyperFlexStorageContainer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHyperFlexStorageContainer) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HyperFlexStorageContainer" of the ClassId field.
func (o *StorageHyperFlexStorageContainer) GetDefaultClassId() interface{} {
	return "storage.HyperFlexStorageContainer"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHyperFlexStorageContainer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHyperFlexStorageContainer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HyperFlexStorageContainer" of the ObjectType field.
func (o *StorageHyperFlexStorageContainer) GetDefaultObjectType() interface{} {
	return "storage.HyperFlexStorageContainer"
}

// GetDataBlockSize returns the DataBlockSize field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetDataBlockSize() int64 {
	if o == nil || IsNil(o.DataBlockSize) {
		var ret int64
		return ret
	}
	return *o.DataBlockSize
}

// GetDataBlockSizeOk returns a tuple with the DataBlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetDataBlockSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DataBlockSize) {
		return nil, false
	}
	return o.DataBlockSize, true
}

// HasDataBlockSize returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasDataBlockSize() bool {
	if o != nil && !IsNil(o.DataBlockSize) {
		return true
	}

	return false
}

// SetDataBlockSize gets a reference to the given int64 and assigns it to the DataBlockSize field.
func (o *StorageHyperFlexStorageContainer) SetDataBlockSize(v int64) {
	o.DataBlockSize = &v
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetInUse() bool {
	if o == nil || IsNil(o.InUse) {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetInUseOk() (*bool, bool) {
	if o == nil || IsNil(o.InUse) {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasInUse() bool {
	if o != nil && !IsNil(o.InUse) {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *StorageHyperFlexStorageContainer) SetInUse(v bool) {
	o.InUse = &v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetLastAccessTime() time.Time {
	if o == nil || IsNil(o.LastAccessTime) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetLastAccessTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAccessTime) {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasLastAccessTime() bool {
	if o != nil && !IsNil(o.LastAccessTime) {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given time.Time and assigns it to the LastAccessTime field.
func (o *StorageHyperFlexStorageContainer) SetLastAccessTime(v time.Time) {
	o.LastAccessTime = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *StorageHyperFlexStorageContainer) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetProvisionedCapacity returns the ProvisionedCapacity field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetProvisionedCapacity() int64 {
	if o == nil || IsNil(o.ProvisionedCapacity) {
		var ret int64
		return ret
	}
	return *o.ProvisionedCapacity
}

// GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetProvisionedCapacityOk() (*int64, bool) {
	if o == nil || IsNil(o.ProvisionedCapacity) {
		return nil, false
	}
	return o.ProvisionedCapacity, true
}

// HasProvisionedCapacity returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasProvisionedCapacity() bool {
	if o != nil && !IsNil(o.ProvisionedCapacity) {
		return true
	}

	return false
}

// SetProvisionedCapacity gets a reference to the given int64 and assigns it to the ProvisionedCapacity field.
func (o *StorageHyperFlexStorageContainer) SetProvisionedCapacity(v int64) {
	o.ProvisionedCapacity = &v
}

// GetProvisionedVolumeCapacityUtilization returns the ProvisionedVolumeCapacityUtilization field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetProvisionedVolumeCapacityUtilization() float32 {
	if o == nil || IsNil(o.ProvisionedVolumeCapacityUtilization) {
		var ret float32
		return ret
	}
	return *o.ProvisionedVolumeCapacityUtilization
}

// GetProvisionedVolumeCapacityUtilizationOk returns a tuple with the ProvisionedVolumeCapacityUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetProvisionedVolumeCapacityUtilizationOk() (*float32, bool) {
	if o == nil || IsNil(o.ProvisionedVolumeCapacityUtilization) {
		return nil, false
	}
	return o.ProvisionedVolumeCapacityUtilization, true
}

// HasProvisionedVolumeCapacityUtilization returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasProvisionedVolumeCapacityUtilization() bool {
	if o != nil && !IsNil(o.ProvisionedVolumeCapacityUtilization) {
		return true
	}

	return false
}

// SetProvisionedVolumeCapacityUtilization gets a reference to the given float32 and assigns it to the ProvisionedVolumeCapacityUtilization field.
func (o *StorageHyperFlexStorageContainer) SetProvisionedVolumeCapacityUtilization(v float32) {
	o.ProvisionedVolumeCapacityUtilization = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageHyperFlexStorageContainer) SetType(v string) {
	o.Type = &v
}

// GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetUnCompressedUsedBytes() int64 {
	if o == nil || IsNil(o.UnCompressedUsedBytes) {
		var ret int64
		return ret
	}
	return *o.UnCompressedUsedBytes
}

// GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetUnCompressedUsedBytesOk() (*int64, bool) {
	if o == nil || IsNil(o.UnCompressedUsedBytes) {
		return nil, false
	}
	return o.UnCompressedUsedBytes, true
}

// HasUnCompressedUsedBytes returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasUnCompressedUsedBytes() bool {
	if o != nil && !IsNil(o.UnCompressedUsedBytes) {
		return true
	}

	return false
}

// SetUnCompressedUsedBytes gets a reference to the given int64 and assigns it to the UnCompressedUsedBytes field.
func (o *StorageHyperFlexStorageContainer) SetUnCompressedUsedBytes(v int64) {
	o.UnCompressedUsedBytes = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageHyperFlexStorageContainer) SetUuid(v string) {
	o.Uuid = &v
}

// GetVolumeCount returns the VolumeCount field value if set, zero value otherwise.
func (o *StorageHyperFlexStorageContainer) GetVolumeCount() int64 {
	if o == nil || IsNil(o.VolumeCount) {
		var ret int64
		return ret
	}
	return *o.VolumeCount
}

// GetVolumeCountOk returns a tuple with the VolumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHyperFlexStorageContainer) GetVolumeCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VolumeCount) {
		return nil, false
	}
	return o.VolumeCount, true
}

// HasVolumeCount returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasVolumeCount() bool {
	if o != nil && !IsNil(o.VolumeCount) {
		return true
	}

	return false
}

// SetVolumeCount gets a reference to the given int64 and assigns it to the VolumeCount field.
func (o *StorageHyperFlexStorageContainer) SetVolumeCount(v int64) {
	o.VolumeCount = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHyperFlexStorageContainer) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHyperFlexStorageContainer) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *StorageHyperFlexStorageContainer) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *StorageHyperFlexStorageContainer) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *StorageHyperFlexStorageContainer) UnsetCluster() {
	o.Cluster.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHyperFlexStorageContainer) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHyperFlexStorageContainer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHyperFlexStorageContainer) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHyperFlexStorageContainer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHyperFlexStorageContainer) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHyperFlexStorageContainer) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHyperFlexStorageContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHyperFlexStorageContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseStorageContainer, errStorageBaseStorageContainer := json.Marshal(o.StorageBaseStorageContainer)
	if errStorageBaseStorageContainer != nil {
		return map[string]interface{}{}, errStorageBaseStorageContainer
	}
	errStorageBaseStorageContainer = json.Unmarshal([]byte(serializedStorageBaseStorageContainer), &toSerialize)
	if errStorageBaseStorageContainer != nil {
		return map[string]interface{}{}, errStorageBaseStorageContainer
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DataBlockSize) {
		toSerialize["DataBlockSize"] = o.DataBlockSize
	}
	if !IsNil(o.InUse) {
		toSerialize["InUse"] = o.InUse
	}
	if !IsNil(o.LastAccessTime) {
		toSerialize["LastAccessTime"] = o.LastAccessTime
	}
	if !IsNil(o.LastModifiedTime) {
		toSerialize["LastModifiedTime"] = o.LastModifiedTime
	}
	if !IsNil(o.ProvisionedCapacity) {
		toSerialize["ProvisionedCapacity"] = o.ProvisionedCapacity
	}
	if !IsNil(o.ProvisionedVolumeCapacityUtilization) {
		toSerialize["ProvisionedVolumeCapacityUtilization"] = o.ProvisionedVolumeCapacityUtilization
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.UnCompressedUsedBytes) {
		toSerialize["UnCompressedUsedBytes"] = o.UnCompressedUsedBytes
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.VolumeCount) {
		toSerialize["VolumeCount"] = o.VolumeCount
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHyperFlexStorageContainer) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHyperFlexStorageContainerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Storage Container data block size
		DataBlockSize *int64 `json:"DataBlockSize,omitempty"`
		// Indicates whether the Storage Container has Volumes.
		InUse *bool `json:"InUse,omitempty"`
		// Storage container's last access time.
		LastAccessTime *time.Time `json:"LastAccessTime,omitempty"`
		// Storage container's last modified time.
		LastModifiedTime *time.Time `json:"LastModifiedTime,omitempty"`
		// Provisioned Capacity of the Storage container.
		ProvisionedCapacity *int64 `json:"ProvisionedCapacity,omitempty"`
		// Provisioned Capacity Utilization of All Volumes associated with the Storage Container.
		ProvisionedVolumeCapacityUtilization *float32 `json:"ProvisionedVolumeCapacityUtilization,omitempty"`
		// Storage Container type (SMB/NFS/iSCSI). * `NFS` - Storage container created/accesed through NFS protocol. * `SMB` - Storage container created/accessed through SMB protocol. * `iSCSI` - Storage container created/accessed through iSCSI protocol.
		Type *string `json:"Type,omitempty"`
		// Uncompressed bytes on Storage Container.
		UnCompressedUsedBytes *int64 `json:"UnCompressedUsedBytes,omitempty"`
		// UUID of the Datastore/Storage Containter.
		Uuid *string `json:"Uuid,omitempty"`
		// Number of Volumes associated with the Storage Container.
		VolumeCount      *int64                                      `json:"VolumeCount,omitempty"`
		Cluster          NullableHyperflexClusterRelationship        `json:"Cluster,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHyperFlexStorageContainerWithoutEmbeddedStruct := StorageHyperFlexStorageContainerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHyperFlexStorageContainerWithoutEmbeddedStruct)
	if err == nil {
		varStorageHyperFlexStorageContainer := _StorageHyperFlexStorageContainer{}
		varStorageHyperFlexStorageContainer.ClassId = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.ClassId
		varStorageHyperFlexStorageContainer.ObjectType = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.ObjectType
		varStorageHyperFlexStorageContainer.DataBlockSize = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.DataBlockSize
		varStorageHyperFlexStorageContainer.InUse = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.InUse
		varStorageHyperFlexStorageContainer.LastAccessTime = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.LastAccessTime
		varStorageHyperFlexStorageContainer.LastModifiedTime = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.LastModifiedTime
		varStorageHyperFlexStorageContainer.ProvisionedCapacity = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.ProvisionedCapacity
		varStorageHyperFlexStorageContainer.ProvisionedVolumeCapacityUtilization = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.ProvisionedVolumeCapacityUtilization
		varStorageHyperFlexStorageContainer.Type = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.Type
		varStorageHyperFlexStorageContainer.UnCompressedUsedBytes = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.UnCompressedUsedBytes
		varStorageHyperFlexStorageContainer.Uuid = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.Uuid
		varStorageHyperFlexStorageContainer.VolumeCount = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.VolumeCount
		varStorageHyperFlexStorageContainer.Cluster = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.Cluster
		varStorageHyperFlexStorageContainer.RegisteredDevice = varStorageHyperFlexStorageContainerWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHyperFlexStorageContainer(varStorageHyperFlexStorageContainer)
	} else {
		return err
	}

	varStorageHyperFlexStorageContainer := _StorageHyperFlexStorageContainer{}

	err = json.Unmarshal(data, &varStorageHyperFlexStorageContainer)
	if err == nil {
		o.StorageBaseStorageContainer = varStorageHyperFlexStorageContainer.StorageBaseStorageContainer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DataBlockSize")
		delete(additionalProperties, "InUse")
		delete(additionalProperties, "LastAccessTime")
		delete(additionalProperties, "LastModifiedTime")
		delete(additionalProperties, "ProvisionedCapacity")
		delete(additionalProperties, "ProvisionedVolumeCapacityUtilization")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "UnCompressedUsedBytes")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VolumeCount")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBaseStorageContainer := reflect.ValueOf(o.StorageBaseStorageContainer)
		for i := 0; i < reflectStorageBaseStorageContainer.Type().NumField(); i++ {
			t := reflectStorageBaseStorageContainer.Type().Field(i)

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

type NullableStorageHyperFlexStorageContainer struct {
	value *StorageHyperFlexStorageContainer
	isSet bool
}

func (v NullableStorageHyperFlexStorageContainer) Get() *StorageHyperFlexStorageContainer {
	return v.value
}

func (v *NullableStorageHyperFlexStorageContainer) Set(val *StorageHyperFlexStorageContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHyperFlexStorageContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHyperFlexStorageContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHyperFlexStorageContainer(val *StorageHyperFlexStorageContainer) *NullableStorageHyperFlexStorageContainer {
	return &NullableStorageHyperFlexStorageContainer{value: val, isSet: true}
}

func (v NullableStorageHyperFlexStorageContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHyperFlexStorageContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
