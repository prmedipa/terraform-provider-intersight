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

// checks if the StorageVirtualDriveExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageVirtualDriveExtension{}

// StorageVirtualDriveExtension Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
type StorageVirtualDriveExtension struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The ability to boot from the virtual drive.
	Bootable *string `json:"Bootable,omitempty"`
	// The container id of the virtual drive.
	ContainerId *int64 `json:"ContainerId,omitempty"`
	// The state of the virtual drive.
	DriveState *string `json:"DriveState,omitempty"`
	// The name of the Virtual drive.
	Name *string `json:"Name,omitempty"`
	// The operational device id of the virtual drive.
	OperDeviceId *string `json:"OperDeviceId,omitempty"`
	// The UUID assigned to the virtual drive.
	Uuid *string `json:"Uuid,omitempty"`
	// The UUID value of the vendor assigned to the virtual drive.
	VendorUuid *string `json:"VendorUuid,omitempty"`
	// The distinguished name of the virtual drive for which the extended data is provided.
	VirtualDriveDn *string `json:"VirtualDriveDn,omitempty"`
	// The Id of the virtual drive.
	VirtualDriveId       *string                                     `json:"VirtualDriveId,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    NullableStorageControllerRelationship       `json:"StorageController,omitempty"`
	VirtualDrive         NullableStorageVirtualDriveRelationship     `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveExtension StorageVirtualDriveExtension

// NewStorageVirtualDriveExtension instantiates a new StorageVirtualDriveExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveExtension(classId string, objectType string) *StorageVirtualDriveExtension {
	this := StorageVirtualDriveExtension{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveExtensionWithDefaults instantiates a new StorageVirtualDriveExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveExtensionWithDefaults() *StorageVirtualDriveExtension {
	this := StorageVirtualDriveExtension{}
	var classId string = "storage.VirtualDriveExtension"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveExtension"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveExtension) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveExtension) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.VirtualDriveExtension" of the ClassId field.
func (o *StorageVirtualDriveExtension) GetDefaultClassId() interface{} {
	return "storage.VirtualDriveExtension"
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveExtension) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveExtension) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.VirtualDriveExtension" of the ObjectType field.
func (o *StorageVirtualDriveExtension) GetDefaultObjectType() interface{} {
	return "storage.VirtualDriveExtension"
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetBootable() string {
	if o == nil || IsNil(o.Bootable) {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetBootableOk() (*string, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *StorageVirtualDriveExtension) SetBootable(v string) {
	o.Bootable = &v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetContainerId() int64 {
	if o == nil || IsNil(o.ContainerId) {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetContainerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *StorageVirtualDriveExtension) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetDriveState returns the DriveState field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetDriveState() string {
	if o == nil || IsNil(o.DriveState) {
		var ret string
		return ret
	}
	return *o.DriveState
}

// GetDriveStateOk returns a tuple with the DriveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetDriveStateOk() (*string, bool) {
	if o == nil || IsNil(o.DriveState) {
		return nil, false
	}
	return o.DriveState, true
}

// HasDriveState returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasDriveState() bool {
	if o != nil && !IsNil(o.DriveState) {
		return true
	}

	return false
}

// SetDriveState gets a reference to the given string and assigns it to the DriveState field.
func (o *StorageVirtualDriveExtension) SetDriveState(v string) {
	o.DriveState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveExtension) SetName(v string) {
	o.Name = &v
}

// GetOperDeviceId returns the OperDeviceId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetOperDeviceId() string {
	if o == nil || IsNil(o.OperDeviceId) {
		var ret string
		return ret
	}
	return *o.OperDeviceId
}

// GetOperDeviceIdOk returns a tuple with the OperDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetOperDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperDeviceId) {
		return nil, false
	}
	return o.OperDeviceId, true
}

// HasOperDeviceId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasOperDeviceId() bool {
	if o != nil && !IsNil(o.OperDeviceId) {
		return true
	}

	return false
}

// SetOperDeviceId gets a reference to the given string and assigns it to the OperDeviceId field.
func (o *StorageVirtualDriveExtension) SetOperDeviceId(v string) {
	o.OperDeviceId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageVirtualDriveExtension) SetUuid(v string) {
	o.Uuid = &v
}

// GetVendorUuid returns the VendorUuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVendorUuid() string {
	if o == nil || IsNil(o.VendorUuid) {
		var ret string
		return ret
	}
	return *o.VendorUuid
}

// GetVendorUuidOk returns a tuple with the VendorUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVendorUuidOk() (*string, bool) {
	if o == nil || IsNil(o.VendorUuid) {
		return nil, false
	}
	return o.VendorUuid, true
}

// HasVendorUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVendorUuid() bool {
	if o != nil && !IsNil(o.VendorUuid) {
		return true
	}

	return false
}

// SetVendorUuid gets a reference to the given string and assigns it to the VendorUuid field.
func (o *StorageVirtualDriveExtension) SetVendorUuid(v string) {
	o.VendorUuid = &v
}

// GetVirtualDriveDn returns the VirtualDriveDn field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVirtualDriveDn() string {
	if o == nil || IsNil(o.VirtualDriveDn) {
		var ret string
		return ret
	}
	return *o.VirtualDriveDn
}

// GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVirtualDriveDnOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualDriveDn) {
		return nil, false
	}
	return o.VirtualDriveDn, true
}

// HasVirtualDriveDn returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDriveDn() bool {
	if o != nil && !IsNil(o.VirtualDriveDn) {
		return true
	}

	return false
}

// SetVirtualDriveDn gets a reference to the given string and assigns it to the VirtualDriveDn field.
func (o *StorageVirtualDriveExtension) SetVirtualDriveDn(v string) {
	o.VirtualDriveDn = &v
}

// GetVirtualDriveId returns the VirtualDriveId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVirtualDriveId() string {
	if o == nil || IsNil(o.VirtualDriveId) {
		var ret string
		return ret
	}
	return *o.VirtualDriveId
}

// GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVirtualDriveIdOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualDriveId) {
		return nil, false
	}
	return o.VirtualDriveId, true
}

// HasVirtualDriveId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDriveId() bool {
	if o != nil && !IsNil(o.VirtualDriveId) {
		return true
	}

	return false
}

// SetVirtualDriveId gets a reference to the given string and assigns it to the VirtualDriveId field.
func (o *StorageVirtualDriveExtension) SetVirtualDriveId(v string) {
	o.VirtualDriveId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVirtualDriveExtension) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageVirtualDriveExtension) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageVirtualDriveExtension) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveExtension) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveExtension) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVirtualDriveExtension) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageVirtualDriveExtension) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageVirtualDriveExtension) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetStorageController returns the StorageController field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveExtension) GetStorageController() StorageControllerRelationship {
	if o == nil || IsNil(o.StorageController.Get()) {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController.Get()
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveExtension) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageController.Get(), o.StorageController.IsSet()
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasStorageController() bool {
	if o != nil && o.StorageController.IsSet() {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given NullableStorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageVirtualDriveExtension) SetStorageController(v StorageControllerRelationship) {
	o.StorageController.Set(&v)
}

// SetStorageControllerNil sets the value for StorageController to be an explicit nil
func (o *StorageVirtualDriveExtension) SetStorageControllerNil() {
	o.StorageController.Set(nil)
}

// UnsetStorageController ensures that no value is present for StorageController, not even an explicit nil
func (o *StorageVirtualDriveExtension) UnsetStorageController() {
	o.StorageController.Unset()
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageVirtualDriveExtension) GetVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || IsNil(o.VirtualDrive.Get()) {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.VirtualDrive.Get()
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageVirtualDriveExtension) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDrive.Get(), o.VirtualDrive.IsSet()
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive.IsSet() {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given NullableStorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveExtension) SetVirtualDrive(v StorageVirtualDriveRelationship) {
	o.VirtualDrive.Set(&v)
}

// SetVirtualDriveNil sets the value for VirtualDrive to be an explicit nil
func (o *StorageVirtualDriveExtension) SetVirtualDriveNil() {
	o.VirtualDrive.Set(nil)
}

// UnsetVirtualDrive ensures that no value is present for VirtualDrive, not even an explicit nil
func (o *StorageVirtualDriveExtension) UnsetVirtualDrive() {
	o.VirtualDrive.Unset()
}

func (o StorageVirtualDriveExtension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageVirtualDriveExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Bootable) {
		toSerialize["Bootable"] = o.Bootable
	}
	if !IsNil(o.ContainerId) {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if !IsNil(o.DriveState) {
		toSerialize["DriveState"] = o.DriveState
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.OperDeviceId) {
		toSerialize["OperDeviceId"] = o.OperDeviceId
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}
	if !IsNil(o.VendorUuid) {
		toSerialize["VendorUuid"] = o.VendorUuid
	}
	if !IsNil(o.VirtualDriveDn) {
		toSerialize["VirtualDriveDn"] = o.VirtualDriveDn
	}
	if !IsNil(o.VirtualDriveId) {
		toSerialize["VirtualDriveId"] = o.VirtualDriveId
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.StorageController.IsSet() {
		toSerialize["StorageController"] = o.StorageController.Get()
	}
	if o.VirtualDrive.IsSet() {
		toSerialize["VirtualDrive"] = o.VirtualDrive.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageVirtualDriveExtension) UnmarshalJSON(data []byte) (err error) {
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
	type StorageVirtualDriveExtensionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The ability to boot from the virtual drive.
		Bootable *string `json:"Bootable,omitempty"`
		// The container id of the virtual drive.
		ContainerId *int64 `json:"ContainerId,omitempty"`
		// The state of the virtual drive.
		DriveState *string `json:"DriveState,omitempty"`
		// The name of the Virtual drive.
		Name *string `json:"Name,omitempty"`
		// The operational device id of the virtual drive.
		OperDeviceId *string `json:"OperDeviceId,omitempty"`
		// The UUID assigned to the virtual drive.
		Uuid *string `json:"Uuid,omitempty"`
		// The UUID value of the vendor assigned to the virtual drive.
		VendorUuid *string `json:"VendorUuid,omitempty"`
		// The distinguished name of the virtual drive for which the extended data is provided.
		VirtualDriveDn *string `json:"VirtualDriveDn,omitempty"`
		// The Id of the virtual drive.
		VirtualDriveId      *string                                     `json:"VirtualDriveId,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StorageController   NullableStorageControllerRelationship       `json:"StorageController,omitempty"`
		VirtualDrive        NullableStorageVirtualDriveRelationship     `json:"VirtualDrive,omitempty"`
	}

	varStorageVirtualDriveExtensionWithoutEmbeddedStruct := StorageVirtualDriveExtensionWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageVirtualDriveExtensionWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDriveExtension := _StorageVirtualDriveExtension{}
		varStorageVirtualDriveExtension.ClassId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ClassId
		varStorageVirtualDriveExtension.ObjectType = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDriveExtension.Bootable = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Bootable
		varStorageVirtualDriveExtension.ContainerId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ContainerId
		varStorageVirtualDriveExtension.DriveState = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.DriveState
		varStorageVirtualDriveExtension.Name = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Name
		varStorageVirtualDriveExtension.OperDeviceId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.OperDeviceId
		varStorageVirtualDriveExtension.Uuid = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Uuid
		varStorageVirtualDriveExtension.VendorUuid = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VendorUuid
		varStorageVirtualDriveExtension.VirtualDriveDn = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDriveDn
		varStorageVirtualDriveExtension.VirtualDriveId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDriveId
		varStorageVirtualDriveExtension.InventoryDeviceInfo = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageVirtualDriveExtension.RegisteredDevice = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.RegisteredDevice
		varStorageVirtualDriveExtension.StorageController = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.StorageController
		varStorageVirtualDriveExtension.VirtualDrive = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDrive
		*o = StorageVirtualDriveExtension(varStorageVirtualDriveExtension)
	} else {
		return err
	}

	varStorageVirtualDriveExtension := _StorageVirtualDriveExtension{}

	err = json.Unmarshal(data, &varStorageVirtualDriveExtension)
	if err == nil {
		o.InventoryBase = varStorageVirtualDriveExtension.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "DriveState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperDeviceId")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VendorUuid")
		delete(additionalProperties, "VirtualDriveDn")
		delete(additionalProperties, "VirtualDriveId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "VirtualDrive")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageVirtualDriveExtension struct {
	value *StorageVirtualDriveExtension
	isSet bool
}

func (v NullableStorageVirtualDriveExtension) Get() *StorageVirtualDriveExtension {
	return v.value
}

func (v *NullableStorageVirtualDriveExtension) Set(val *StorageVirtualDriveExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveExtension(val *StorageVirtualDriveExtension) *NullableStorageVirtualDriveExtension {
	return &NullableStorageVirtualDriveExtension{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
