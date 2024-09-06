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

// checks if the MemoryPersistentMemoryConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryPersistentMemoryConfiguration{}

// MemoryPersistentMemoryConfiguration Persistent Memory configuration on all the Persistent Memory Modules on a server.
type MemoryPersistentMemoryConfiguration struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Memory capacity in GiB of a Persistent Memory configuration on a server.
	MemoryCapacity *string `json:"MemoryCapacity,omitempty"`
	// Number of Persistent Memory Modules on a server.
	NumOfModules *string `json:"NumOfModules,omitempty"`
	// Number of Persistent Memory Regions on a server.
	NumOfRegions *string `json:"NumOfRegions,omitempty"`
	// Persistent memory capacity in GiB of a Persistent Memory configuration on a server.
	PersistentMemoryCapacity *string `json:"PersistentMemoryCapacity,omitempty"`
	// Reserved capacity in GiB of a Persistent Memory configuration on a server.
	ReservedCapacity *string `json:"ReservedCapacity,omitempty"`
	// Collective security state of all Persistent Memory modules on a server.
	SecurityState *string `json:"SecurityState,omitempty"`
	// Total capacity in GiB of a Persistent Memory configuration on a server.
	TotalCapacity                *string                                                `json:"TotalCapacity,omitempty"`
	ComputeBoard                 NullableComputeBoardRelationship                       `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo          NullableInventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
	PersistentMemoryConfigResult NullableMemoryPersistentMemoryConfigResultRelationship `json:"PersistentMemoryConfigResult,omitempty"`
	// An array of relationships to memoryPersistentMemoryRegion resources.
	PersistentMemoryRegions []MemoryPersistentMemoryRegionRelationship  `json:"PersistentMemoryRegions,omitempty"`
	RegisteredDevice        NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _MemoryPersistentMemoryConfiguration MemoryPersistentMemoryConfiguration

// NewMemoryPersistentMemoryConfiguration instantiates a new MemoryPersistentMemoryConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryConfiguration(classId string, objectType string) *MemoryPersistentMemoryConfiguration {
	this := MemoryPersistentMemoryConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryConfigurationWithDefaults instantiates a new MemoryPersistentMemoryConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryConfigurationWithDefaults() *MemoryPersistentMemoryConfiguration {
	this := MemoryPersistentMemoryConfiguration{}
	var classId string = "memory.PersistentMemoryConfiguration"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.PersistentMemoryConfiguration" of the ClassId field.
func (o *MemoryPersistentMemoryConfiguration) GetDefaultClassId() interface{} {
	return "memory.PersistentMemoryConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.PersistentMemoryConfiguration" of the ObjectType field.
func (o *MemoryPersistentMemoryConfiguration) GetDefaultObjectType() interface{} {
	return "memory.PersistentMemoryConfiguration"
}

// GetMemoryCapacity returns the MemoryCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetMemoryCapacity() string {
	if o == nil || IsNil(o.MemoryCapacity) {
		var ret string
		return ret
	}
	return *o.MemoryCapacity
}

// GetMemoryCapacityOk returns a tuple with the MemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetMemoryCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.MemoryCapacity) {
		return nil, false
	}
	return o.MemoryCapacity, true
}

// HasMemoryCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasMemoryCapacity() bool {
	if o != nil && !IsNil(o.MemoryCapacity) {
		return true
	}

	return false
}

// SetMemoryCapacity gets a reference to the given string and assigns it to the MemoryCapacity field.
func (o *MemoryPersistentMemoryConfiguration) SetMemoryCapacity(v string) {
	o.MemoryCapacity = &v
}

// GetNumOfModules returns the NumOfModules field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetNumOfModules() string {
	if o == nil || IsNil(o.NumOfModules) {
		var ret string
		return ret
	}
	return *o.NumOfModules
}

// GetNumOfModulesOk returns a tuple with the NumOfModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetNumOfModulesOk() (*string, bool) {
	if o == nil || IsNil(o.NumOfModules) {
		return nil, false
	}
	return o.NumOfModules, true
}

// HasNumOfModules returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasNumOfModules() bool {
	if o != nil && !IsNil(o.NumOfModules) {
		return true
	}

	return false
}

// SetNumOfModules gets a reference to the given string and assigns it to the NumOfModules field.
func (o *MemoryPersistentMemoryConfiguration) SetNumOfModules(v string) {
	o.NumOfModules = &v
}

// GetNumOfRegions returns the NumOfRegions field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetNumOfRegions() string {
	if o == nil || IsNil(o.NumOfRegions) {
		var ret string
		return ret
	}
	return *o.NumOfRegions
}

// GetNumOfRegionsOk returns a tuple with the NumOfRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetNumOfRegionsOk() (*string, bool) {
	if o == nil || IsNil(o.NumOfRegions) {
		return nil, false
	}
	return o.NumOfRegions, true
}

// HasNumOfRegions returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasNumOfRegions() bool {
	if o != nil && !IsNil(o.NumOfRegions) {
		return true
	}

	return false
}

// SetNumOfRegions gets a reference to the given string and assigns it to the NumOfRegions field.
func (o *MemoryPersistentMemoryConfiguration) SetNumOfRegions(v string) {
	o.NumOfRegions = &v
}

// GetPersistentMemoryCapacity returns the PersistentMemoryCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryCapacity() string {
	if o == nil || IsNil(o.PersistentMemoryCapacity) {
		var ret string
		return ret
	}
	return *o.PersistentMemoryCapacity
}

// GetPersistentMemoryCapacityOk returns a tuple with the PersistentMemoryCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.PersistentMemoryCapacity) {
		return nil, false
	}
	return o.PersistentMemoryCapacity, true
}

// HasPersistentMemoryCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasPersistentMemoryCapacity() bool {
	if o != nil && !IsNil(o.PersistentMemoryCapacity) {
		return true
	}

	return false
}

// SetPersistentMemoryCapacity gets a reference to the given string and assigns it to the PersistentMemoryCapacity field.
func (o *MemoryPersistentMemoryConfiguration) SetPersistentMemoryCapacity(v string) {
	o.PersistentMemoryCapacity = &v
}

// GetReservedCapacity returns the ReservedCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetReservedCapacity() string {
	if o == nil || IsNil(o.ReservedCapacity) {
		var ret string
		return ret
	}
	return *o.ReservedCapacity
}

// GetReservedCapacityOk returns a tuple with the ReservedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetReservedCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.ReservedCapacity) {
		return nil, false
	}
	return o.ReservedCapacity, true
}

// HasReservedCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasReservedCapacity() bool {
	if o != nil && !IsNil(o.ReservedCapacity) {
		return true
	}

	return false
}

// SetReservedCapacity gets a reference to the given string and assigns it to the ReservedCapacity field.
func (o *MemoryPersistentMemoryConfiguration) SetReservedCapacity(v string) {
	o.ReservedCapacity = &v
}

// GetSecurityState returns the SecurityState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetSecurityState() string {
	if o == nil || IsNil(o.SecurityState) {
		var ret string
		return ret
	}
	return *o.SecurityState
}

// GetSecurityStateOk returns a tuple with the SecurityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetSecurityStateOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityState) {
		return nil, false
	}
	return o.SecurityState, true
}

// HasSecurityState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasSecurityState() bool {
	if o != nil && !IsNil(o.SecurityState) {
		return true
	}

	return false
}

// SetSecurityState gets a reference to the given string and assigns it to the SecurityState field.
func (o *MemoryPersistentMemoryConfiguration) SetSecurityState(v string) {
	o.SecurityState = &v
}

// GetTotalCapacity returns the TotalCapacity field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfiguration) GetTotalCapacity() string {
	if o == nil || IsNil(o.TotalCapacity) {
		var ret string
		return ret
	}
	return *o.TotalCapacity
}

// GetTotalCapacityOk returns a tuple with the TotalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfiguration) GetTotalCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCapacity) {
		return nil, false
	}
	return o.TotalCapacity, true
}

// HasTotalCapacity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasTotalCapacity() bool {
	if o != nil && !IsNil(o.TotalCapacity) {
		return true
	}

	return false
}

// SetTotalCapacity gets a reference to the given string and assigns it to the TotalCapacity field.
func (o *MemoryPersistentMemoryConfiguration) SetTotalCapacity(v string) {
	o.TotalCapacity = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfiguration) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfiguration) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *MemoryPersistentMemoryConfiguration) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *MemoryPersistentMemoryConfiguration) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *MemoryPersistentMemoryConfiguration) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfiguration) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfiguration) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryConfiguration) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *MemoryPersistentMemoryConfiguration) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *MemoryPersistentMemoryConfiguration) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPersistentMemoryConfigResult returns the PersistentMemoryConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryConfigResult() MemoryPersistentMemoryConfigResultRelationship {
	if o == nil || IsNil(o.PersistentMemoryConfigResult.Get()) {
		var ret MemoryPersistentMemoryConfigResultRelationship
		return ret
	}
	return *o.PersistentMemoryConfigResult.Get()
}

// GetPersistentMemoryConfigResultOk returns a tuple with the PersistentMemoryConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryConfigResultOk() (*MemoryPersistentMemoryConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersistentMemoryConfigResult.Get(), o.PersistentMemoryConfigResult.IsSet()
}

// HasPersistentMemoryConfigResult returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasPersistentMemoryConfigResult() bool {
	if o != nil && o.PersistentMemoryConfigResult.IsSet() {
		return true
	}

	return false
}

// SetPersistentMemoryConfigResult gets a reference to the given NullableMemoryPersistentMemoryConfigResultRelationship and assigns it to the PersistentMemoryConfigResult field.
func (o *MemoryPersistentMemoryConfiguration) SetPersistentMemoryConfigResult(v MemoryPersistentMemoryConfigResultRelationship) {
	o.PersistentMemoryConfigResult.Set(&v)
}

// SetPersistentMemoryConfigResultNil sets the value for PersistentMemoryConfigResult to be an explicit nil
func (o *MemoryPersistentMemoryConfiguration) SetPersistentMemoryConfigResultNil() {
	o.PersistentMemoryConfigResult.Set(nil)
}

// UnsetPersistentMemoryConfigResult ensures that no value is present for PersistentMemoryConfigResult, not even an explicit nil
func (o *MemoryPersistentMemoryConfiguration) UnsetPersistentMemoryConfigResult() {
	o.PersistentMemoryConfigResult.Unset()
}

// GetPersistentMemoryRegions returns the PersistentMemoryRegions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryRegions() []MemoryPersistentMemoryRegionRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryRegionRelationship
		return ret
	}
	return o.PersistentMemoryRegions
}

// GetPersistentMemoryRegionsOk returns a tuple with the PersistentMemoryRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfiguration) GetPersistentMemoryRegionsOk() ([]MemoryPersistentMemoryRegionRelationship, bool) {
	if o == nil || IsNil(o.PersistentMemoryRegions) {
		return nil, false
	}
	return o.PersistentMemoryRegions, true
}

// HasPersistentMemoryRegions returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasPersistentMemoryRegions() bool {
	if o != nil && !IsNil(o.PersistentMemoryRegions) {
		return true
	}

	return false
}

// SetPersistentMemoryRegions gets a reference to the given []MemoryPersistentMemoryRegionRelationship and assigns it to the PersistentMemoryRegions field.
func (o *MemoryPersistentMemoryConfiguration) SetPersistentMemoryRegions(v []MemoryPersistentMemoryRegionRelationship) {
	o.PersistentMemoryRegions = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfiguration) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfiguration) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfiguration) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryConfiguration) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MemoryPersistentMemoryConfiguration) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MemoryPersistentMemoryConfiguration) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o MemoryPersistentMemoryConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryPersistentMemoryConfiguration) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MemoryCapacity) {
		toSerialize["MemoryCapacity"] = o.MemoryCapacity
	}
	if !IsNil(o.NumOfModules) {
		toSerialize["NumOfModules"] = o.NumOfModules
	}
	if !IsNil(o.NumOfRegions) {
		toSerialize["NumOfRegions"] = o.NumOfRegions
	}
	if !IsNil(o.PersistentMemoryCapacity) {
		toSerialize["PersistentMemoryCapacity"] = o.PersistentMemoryCapacity
	}
	if !IsNil(o.ReservedCapacity) {
		toSerialize["ReservedCapacity"] = o.ReservedCapacity
	}
	if !IsNil(o.SecurityState) {
		toSerialize["SecurityState"] = o.SecurityState
	}
	if !IsNil(o.TotalCapacity) {
		toSerialize["TotalCapacity"] = o.TotalCapacity
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.PersistentMemoryConfigResult.IsSet() {
		toSerialize["PersistentMemoryConfigResult"] = o.PersistentMemoryConfigResult.Get()
	}
	if o.PersistentMemoryRegions != nil {
		toSerialize["PersistentMemoryRegions"] = o.PersistentMemoryRegions
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryPersistentMemoryConfiguration) UnmarshalJSON(data []byte) (err error) {
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
	type MemoryPersistentMemoryConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Memory capacity in GiB of a Persistent Memory configuration on a server.
		MemoryCapacity *string `json:"MemoryCapacity,omitempty"`
		// Number of Persistent Memory Modules on a server.
		NumOfModules *string `json:"NumOfModules,omitempty"`
		// Number of Persistent Memory Regions on a server.
		NumOfRegions *string `json:"NumOfRegions,omitempty"`
		// Persistent memory capacity in GiB of a Persistent Memory configuration on a server.
		PersistentMemoryCapacity *string `json:"PersistentMemoryCapacity,omitempty"`
		// Reserved capacity in GiB of a Persistent Memory configuration on a server.
		ReservedCapacity *string `json:"ReservedCapacity,omitempty"`
		// Collective security state of all Persistent Memory modules on a server.
		SecurityState *string `json:"SecurityState,omitempty"`
		// Total capacity in GiB of a Persistent Memory configuration on a server.
		TotalCapacity                *string                                                `json:"TotalCapacity,omitempty"`
		ComputeBoard                 NullableComputeBoardRelationship                       `json:"ComputeBoard,omitempty"`
		InventoryDeviceInfo          NullableInventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
		PersistentMemoryConfigResult NullableMemoryPersistentMemoryConfigResultRelationship `json:"PersistentMemoryConfigResult,omitempty"`
		// An array of relationships to memoryPersistentMemoryRegion resources.
		PersistentMemoryRegions []MemoryPersistentMemoryRegionRelationship  `json:"PersistentMemoryRegions,omitempty"`
		RegisteredDevice        NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct := MemoryPersistentMemoryConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryConfiguration := _MemoryPersistentMemoryConfiguration{}
		varMemoryPersistentMemoryConfiguration.ClassId = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryConfiguration.ObjectType = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryConfiguration.MemoryCapacity = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.MemoryCapacity
		varMemoryPersistentMemoryConfiguration.NumOfModules = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.NumOfModules
		varMemoryPersistentMemoryConfiguration.NumOfRegions = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.NumOfRegions
		varMemoryPersistentMemoryConfiguration.PersistentMemoryCapacity = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.PersistentMemoryCapacity
		varMemoryPersistentMemoryConfiguration.ReservedCapacity = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.ReservedCapacity
		varMemoryPersistentMemoryConfiguration.SecurityState = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.SecurityState
		varMemoryPersistentMemoryConfiguration.TotalCapacity = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.TotalCapacity
		varMemoryPersistentMemoryConfiguration.ComputeBoard = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.ComputeBoard
		varMemoryPersistentMemoryConfiguration.InventoryDeviceInfo = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryConfiguration.PersistentMemoryConfigResult = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.PersistentMemoryConfigResult
		varMemoryPersistentMemoryConfiguration.PersistentMemoryRegions = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.PersistentMemoryRegions
		varMemoryPersistentMemoryConfiguration.RegisteredDevice = varMemoryPersistentMemoryConfigurationWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryConfiguration(varMemoryPersistentMemoryConfiguration)
	} else {
		return err
	}

	varMemoryPersistentMemoryConfiguration := _MemoryPersistentMemoryConfiguration{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryConfiguration)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryConfiguration.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MemoryCapacity")
		delete(additionalProperties, "NumOfModules")
		delete(additionalProperties, "NumOfRegions")
		delete(additionalProperties, "PersistentMemoryCapacity")
		delete(additionalProperties, "ReservedCapacity")
		delete(additionalProperties, "SecurityState")
		delete(additionalProperties, "TotalCapacity")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PersistentMemoryConfigResult")
		delete(additionalProperties, "PersistentMemoryRegions")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableMemoryPersistentMemoryConfiguration struct {
	value *MemoryPersistentMemoryConfiguration
	isSet bool
}

func (v NullableMemoryPersistentMemoryConfiguration) Get() *MemoryPersistentMemoryConfiguration {
	return v.value
}

func (v *NullableMemoryPersistentMemoryConfiguration) Set(val *MemoryPersistentMemoryConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryConfiguration(val *MemoryPersistentMemoryConfiguration) *NullableMemoryPersistentMemoryConfiguration {
	return &NullableMemoryPersistentMemoryConfiguration{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
