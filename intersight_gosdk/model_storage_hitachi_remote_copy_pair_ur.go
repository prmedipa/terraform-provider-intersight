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

// checks if the StorageHitachiRemoteCopyPairUr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHitachiRemoteCopyPairUr{}

// StorageHitachiRemoteCopyPairUr Universal Replicator pair entity in Hitachi storage array.
type StorageHitachiRemoteCopyPairUr struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MU (mirror unit) number of the volume.
	MuNumber *string `json:"MuNumber,omitempty"`
	// Object ID of the remote copy pair.
	Name *string `json:"Name,omitempty"`
	// LDEV number of primary volume.
	PvolLdevId *int64 `json:"PvolLdevId,omitempty"`
	// Serial number of the storage system on the P-VOL.
	PvolStorageSerial *string `json:"PvolStorageSerial,omitempty"`
	// Pair type of the remote copy pair.
	ReplicationType *string `json:"ReplicationType,omitempty"`
	// Status of the remote copy pair.
	Status *string `json:"Status,omitempty"`
	// LDEV number of secondary volume.
	SvolLdevId *int64 `json:"SvolLdevId,omitempty"`
	// Serial number of the storage system on the S-VOL.
	SvolStorageSerial    *string                                     `json:"SvolStorageSerial,omitempty"`
	Array                NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiRemoteCopyPairUr StorageHitachiRemoteCopyPairUr

// NewStorageHitachiRemoteCopyPairUr instantiates a new StorageHitachiRemoteCopyPairUr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiRemoteCopyPairUr(classId string, objectType string) *StorageHitachiRemoteCopyPairUr {
	this := StorageHitachiRemoteCopyPairUr{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiRemoteCopyPairUrWithDefaults instantiates a new StorageHitachiRemoteCopyPairUr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiRemoteCopyPairUrWithDefaults() *StorageHitachiRemoteCopyPairUr {
	this := StorageHitachiRemoteCopyPairUr{}
	var classId string = "storage.HitachiRemoteCopyPairUr"
	this.ClassId = classId
	var objectType string = "storage.HitachiRemoteCopyPairUr"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiRemoteCopyPairUr) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiRemoteCopyPairUr) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HitachiRemoteCopyPairUr" of the ClassId field.
func (o *StorageHitachiRemoteCopyPairUr) GetDefaultClassId() interface{} {
	return "storage.HitachiRemoteCopyPairUr"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiRemoteCopyPairUr) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiRemoteCopyPairUr) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HitachiRemoteCopyPairUr" of the ObjectType field.
func (o *StorageHitachiRemoteCopyPairUr) GetDefaultObjectType() interface{} {
	return "storage.HitachiRemoteCopyPairUr"
}

// GetMuNumber returns the MuNumber field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetMuNumber() string {
	if o == nil || IsNil(o.MuNumber) {
		var ret string
		return ret
	}
	return *o.MuNumber
}

// GetMuNumberOk returns a tuple with the MuNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetMuNumberOk() (*string, bool) {
	if o == nil || IsNil(o.MuNumber) {
		return nil, false
	}
	return o.MuNumber, true
}

// HasMuNumber returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasMuNumber() bool {
	if o != nil && !IsNil(o.MuNumber) {
		return true
	}

	return false
}

// SetMuNumber gets a reference to the given string and assigns it to the MuNumber field.
func (o *StorageHitachiRemoteCopyPairUr) SetMuNumber(v string) {
	o.MuNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageHitachiRemoteCopyPairUr) SetName(v string) {
	o.Name = &v
}

// GetPvolLdevId returns the PvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetPvolLdevId() int64 {
	if o == nil || IsNil(o.PvolLdevId) {
		var ret int64
		return ret
	}
	return *o.PvolLdevId
}

// GetPvolLdevIdOk returns a tuple with the PvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetPvolLdevIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PvolLdevId) {
		return nil, false
	}
	return o.PvolLdevId, true
}

// HasPvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasPvolLdevId() bool {
	if o != nil && !IsNil(o.PvolLdevId) {
		return true
	}

	return false
}

// SetPvolLdevId gets a reference to the given int64 and assigns it to the PvolLdevId field.
func (o *StorageHitachiRemoteCopyPairUr) SetPvolLdevId(v int64) {
	o.PvolLdevId = &v
}

// GetPvolStorageSerial returns the PvolStorageSerial field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetPvolStorageSerial() string {
	if o == nil || IsNil(o.PvolStorageSerial) {
		var ret string
		return ret
	}
	return *o.PvolStorageSerial
}

// GetPvolStorageSerialOk returns a tuple with the PvolStorageSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetPvolStorageSerialOk() (*string, bool) {
	if o == nil || IsNil(o.PvolStorageSerial) {
		return nil, false
	}
	return o.PvolStorageSerial, true
}

// HasPvolStorageSerial returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasPvolStorageSerial() bool {
	if o != nil && !IsNil(o.PvolStorageSerial) {
		return true
	}

	return false
}

// SetPvolStorageSerial gets a reference to the given string and assigns it to the PvolStorageSerial field.
func (o *StorageHitachiRemoteCopyPairUr) SetPvolStorageSerial(v string) {
	o.PvolStorageSerial = &v
}

// GetReplicationType returns the ReplicationType field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetReplicationType() string {
	if o == nil || IsNil(o.ReplicationType) {
		var ret string
		return ret
	}
	return *o.ReplicationType
}

// GetReplicationTypeOk returns a tuple with the ReplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetReplicationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationType) {
		return nil, false
	}
	return o.ReplicationType, true
}

// HasReplicationType returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasReplicationType() bool {
	if o != nil && !IsNil(o.ReplicationType) {
		return true
	}

	return false
}

// SetReplicationType gets a reference to the given string and assigns it to the ReplicationType field.
func (o *StorageHitachiRemoteCopyPairUr) SetReplicationType(v string) {
	o.ReplicationType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageHitachiRemoteCopyPairUr) SetStatus(v string) {
	o.Status = &v
}

// GetSvolLdevId returns the SvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetSvolLdevId() int64 {
	if o == nil || IsNil(o.SvolLdevId) {
		var ret int64
		return ret
	}
	return *o.SvolLdevId
}

// GetSvolLdevIdOk returns a tuple with the SvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetSvolLdevIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SvolLdevId) {
		return nil, false
	}
	return o.SvolLdevId, true
}

// HasSvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasSvolLdevId() bool {
	if o != nil && !IsNil(o.SvolLdevId) {
		return true
	}

	return false
}

// SetSvolLdevId gets a reference to the given int64 and assigns it to the SvolLdevId field.
func (o *StorageHitachiRemoteCopyPairUr) SetSvolLdevId(v int64) {
	o.SvolLdevId = &v
}

// GetSvolStorageSerial returns the SvolStorageSerial field value if set, zero value otherwise.
func (o *StorageHitachiRemoteCopyPairUr) GetSvolStorageSerial() string {
	if o == nil || IsNil(o.SvolStorageSerial) {
		var ret string
		return ret
	}
	return *o.SvolStorageSerial
}

// GetSvolStorageSerialOk returns a tuple with the SvolStorageSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteCopyPairUr) GetSvolStorageSerialOk() (*string, bool) {
	if o == nil || IsNil(o.SvolStorageSerial) {
		return nil, false
	}
	return o.SvolStorageSerial, true
}

// HasSvolStorageSerial returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasSvolStorageSerial() bool {
	if o != nil && !IsNil(o.SvolStorageSerial) {
		return true
	}

	return false
}

// SetSvolStorageSerial gets a reference to the given string and assigns it to the SvolStorageSerial field.
func (o *StorageHitachiRemoteCopyPairUr) SetSvolStorageSerial(v string) {
	o.SvolStorageSerial = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiRemoteCopyPairUr) GetArray() StorageHitachiArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiRemoteCopyPairUr) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiRemoteCopyPairUr) SetArray(v StorageHitachiArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageHitachiRemoteCopyPairUr) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageHitachiRemoteCopyPairUr) UnsetArray() {
	o.Array.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiRemoteCopyPairUr) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiRemoteCopyPairUr) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiRemoteCopyPairUr) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiRemoteCopyPairUr) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHitachiRemoteCopyPairUr) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHitachiRemoteCopyPairUr) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHitachiRemoteCopyPairUr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHitachiRemoteCopyPairUr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.MuNumber) {
		toSerialize["MuNumber"] = o.MuNumber
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.PvolLdevId) {
		toSerialize["PvolLdevId"] = o.PvolLdevId
	}
	if !IsNil(o.PvolStorageSerial) {
		toSerialize["PvolStorageSerial"] = o.PvolStorageSerial
	}
	if !IsNil(o.ReplicationType) {
		toSerialize["ReplicationType"] = o.ReplicationType
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.SvolLdevId) {
		toSerialize["SvolLdevId"] = o.SvolLdevId
	}
	if !IsNil(o.SvolStorageSerial) {
		toSerialize["SvolStorageSerial"] = o.SvolStorageSerial
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHitachiRemoteCopyPairUr) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// MU (mirror unit) number of the volume.
		MuNumber *string `json:"MuNumber,omitempty"`
		// Object ID of the remote copy pair.
		Name *string `json:"Name,omitempty"`
		// LDEV number of primary volume.
		PvolLdevId *int64 `json:"PvolLdevId,omitempty"`
		// Serial number of the storage system on the P-VOL.
		PvolStorageSerial *string `json:"PvolStorageSerial,omitempty"`
		// Pair type of the remote copy pair.
		ReplicationType *string `json:"ReplicationType,omitempty"`
		// Status of the remote copy pair.
		Status *string `json:"Status,omitempty"`
		// LDEV number of secondary volume.
		SvolLdevId *int64 `json:"SvolLdevId,omitempty"`
		// Serial number of the storage system on the S-VOL.
		SvolStorageSerial *string                                     `json:"SvolStorageSerial,omitempty"`
		Array             NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice  NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct := StorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiRemoteCopyPairUr := _StorageHitachiRemoteCopyPairUr{}
		varStorageHitachiRemoteCopyPairUr.ClassId = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.ClassId
		varStorageHitachiRemoteCopyPairUr.ObjectType = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.ObjectType
		varStorageHitachiRemoteCopyPairUr.MuNumber = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.MuNumber
		varStorageHitachiRemoteCopyPairUr.Name = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.Name
		varStorageHitachiRemoteCopyPairUr.PvolLdevId = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.PvolLdevId
		varStorageHitachiRemoteCopyPairUr.PvolStorageSerial = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.PvolStorageSerial
		varStorageHitachiRemoteCopyPairUr.ReplicationType = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.ReplicationType
		varStorageHitachiRemoteCopyPairUr.Status = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.Status
		varStorageHitachiRemoteCopyPairUr.SvolLdevId = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.SvolLdevId
		varStorageHitachiRemoteCopyPairUr.SvolStorageSerial = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.SvolStorageSerial
		varStorageHitachiRemoteCopyPairUr.Array = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.Array
		varStorageHitachiRemoteCopyPairUr.RegisteredDevice = varStorageHitachiRemoteCopyPairUrWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiRemoteCopyPairUr(varStorageHitachiRemoteCopyPairUr)
	} else {
		return err
	}

	varStorageHitachiRemoteCopyPairUr := _StorageHitachiRemoteCopyPairUr{}

	err = json.Unmarshal(data, &varStorageHitachiRemoteCopyPairUr)
	if err == nil {
		o.MoBaseMo = varStorageHitachiRemoteCopyPairUr.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MuNumber")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PvolLdevId")
		delete(additionalProperties, "PvolStorageSerial")
		delete(additionalProperties, "ReplicationType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SvolLdevId")
		delete(additionalProperties, "SvolStorageSerial")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableStorageHitachiRemoteCopyPairUr struct {
	value *StorageHitachiRemoteCopyPairUr
	isSet bool
}

func (v NullableStorageHitachiRemoteCopyPairUr) Get() *StorageHitachiRemoteCopyPairUr {
	return v.value
}

func (v *NullableStorageHitachiRemoteCopyPairUr) Set(val *StorageHitachiRemoteCopyPairUr) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiRemoteCopyPairUr) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiRemoteCopyPairUr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiRemoteCopyPairUr(val *StorageHitachiRemoteCopyPairUr) *NullableStorageHitachiRemoteCopyPairUr {
	return &NullableStorageHitachiRemoteCopyPairUr{value: val, isSet: true}
}

func (v NullableStorageHitachiRemoteCopyPairUr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiRemoteCopyPairUr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
