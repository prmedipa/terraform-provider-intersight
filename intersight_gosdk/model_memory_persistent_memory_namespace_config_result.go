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

// checks if the MemoryPersistentMemoryNamespaceConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryPersistentMemoryNamespaceConfigResult{}

// MemoryPersistentMemoryNamespaceConfigResult Result of a previously configured Persistent Memory Namespace on a server.
type MemoryPersistentMemoryNamespaceConfigResult struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Status of the Persistent Memory Namespace needed to be configured.
	ConfigStatus *string `json:"ConfigStatus,omitempty"`
	// Name of a Persistent Memory Namespace that needed to be configured.
	Name *string `json:"Name,omitempty"`
	// Socket ID in which the Persistent Memory Namespace needed to be configured.
	SocketId *string `json:"SocketId,omitempty"`
	// Socket Memory ID in which the Persistent Memory Namespace needed to be configured.
	SocketMemoryId                     *string                                                `json:"SocketMemoryId,omitempty"`
	InventoryDeviceInfo                NullableInventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfigResult NullableMemoryPersistentMemoryConfigResultRelationship `json:"MemoryPersistentMemoryConfigResult,omitempty"`
	RegisteredDevice                   NullableAssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	AdditionalProperties               map[string]interface{}
}

type _MemoryPersistentMemoryNamespaceConfigResult MemoryPersistentMemoryNamespaceConfigResult

// NewMemoryPersistentMemoryNamespaceConfigResult instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryNamespaceConfigResult(classId string, objectType string) *MemoryPersistentMemoryNamespaceConfigResult {
	this := MemoryPersistentMemoryNamespaceConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults instantiates a new MemoryPersistentMemoryNamespaceConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryNamespaceConfigResultWithDefaults() *MemoryPersistentMemoryNamespaceConfigResult {
	this := MemoryPersistentMemoryNamespaceConfigResult{}
	var classId string = "memory.PersistentMemoryNamespaceConfigResult"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryNamespaceConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.PersistentMemoryNamespaceConfigResult" of the ClassId field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetDefaultClassId() interface{} {
	return "memory.PersistentMemoryNamespaceConfigResult"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.PersistentMemoryNamespaceConfigResult" of the ObjectType field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetDefaultObjectType() interface{} {
	return "memory.PersistentMemoryNamespaceConfigResult"
}

// GetConfigStatus returns the ConfigStatus field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatus() string {
	if o == nil || IsNil(o.ConfigStatus) {
		var ret string
		return ret
	}
	return *o.ConfigStatus
}

// GetConfigStatusOk returns a tuple with the ConfigStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetConfigStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigStatus) {
		return nil, false
	}
	return o.ConfigStatus, true
}

// HasConfigStatus returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasConfigStatus() bool {
	if o != nil && !IsNil(o.ConfigStatus) {
		return true
	}

	return false
}

// SetConfigStatus gets a reference to the given string and assigns it to the ConfigStatus field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetConfigStatus(v string) {
	o.ConfigStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetName(v string) {
	o.Name = &v
}

// GetSocketId returns the SocketId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketId() string {
	if o == nil || IsNil(o.SocketId) {
		var ret string
		return ret
	}
	return *o.SocketId
}

// GetSocketIdOk returns a tuple with the SocketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketIdOk() (*string, bool) {
	if o == nil || IsNil(o.SocketId) {
		return nil, false
	}
	return o.SocketId, true
}

// HasSocketId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketId() bool {
	if o != nil && !IsNil(o.SocketId) {
		return true
	}

	return false
}

// SetSocketId gets a reference to the given string and assigns it to the SocketId field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketId(v string) {
	o.SocketId = &v
}

// GetSocketMemoryId returns the SocketMemoryId field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryId() string {
	if o == nil || IsNil(o.SocketMemoryId) {
		var ret string
		return ret
	}
	return *o.SocketMemoryId
}

// GetSocketMemoryIdOk returns a tuple with the SocketMemoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetSocketMemoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.SocketMemoryId) {
		return nil, false
	}
	return o.SocketMemoryId, true
}

// HasSocketMemoryId returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasSocketMemoryId() bool {
	if o != nil && !IsNil(o.SocketMemoryId) {
		return true
	}

	return false
}

// SetSocketMemoryId gets a reference to the given string and assigns it to the SocketMemoryId field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetSocketMemoryId(v string) {
	o.SocketMemoryId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetMemoryPersistentMemoryConfigResult returns the MemoryPersistentMemoryConfigResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResult() MemoryPersistentMemoryConfigResultRelationship {
	if o == nil || IsNil(o.MemoryPersistentMemoryConfigResult.Get()) {
		var ret MemoryPersistentMemoryConfigResultRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfigResult.Get()
}

// GetMemoryPersistentMemoryConfigResultOk returns a tuple with the MemoryPersistentMemoryConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetMemoryPersistentMemoryConfigResultOk() (*MemoryPersistentMemoryConfigResultRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfigResult.Get(), o.MemoryPersistentMemoryConfigResult.IsSet()
}

// HasMemoryPersistentMemoryConfigResult returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasMemoryPersistentMemoryConfigResult() bool {
	if o != nil && o.MemoryPersistentMemoryConfigResult.IsSet() {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfigResult gets a reference to the given NullableMemoryPersistentMemoryConfigResultRelationship and assigns it to the MemoryPersistentMemoryConfigResult field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetMemoryPersistentMemoryConfigResult(v MemoryPersistentMemoryConfigResultRelationship) {
	o.MemoryPersistentMemoryConfigResult.Set(&v)
}

// SetMemoryPersistentMemoryConfigResultNil sets the value for MemoryPersistentMemoryConfigResult to be an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetMemoryPersistentMemoryConfigResultNil() {
	o.MemoryPersistentMemoryConfigResult.Set(nil)
}

// UnsetMemoryPersistentMemoryConfigResult ensures that no value is present for MemoryPersistentMemoryConfigResult, not even an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) UnsetMemoryPersistentMemoryConfigResult() {
	o.MemoryPersistentMemoryConfigResult.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryNamespaceConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryNamespaceConfigResult) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MemoryPersistentMemoryNamespaceConfigResult) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o MemoryPersistentMemoryNamespaceConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryPersistentMemoryNamespaceConfigResult) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConfigStatus) {
		toSerialize["ConfigStatus"] = o.ConfigStatus
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SocketId) {
		toSerialize["SocketId"] = o.SocketId
	}
	if !IsNil(o.SocketMemoryId) {
		toSerialize["SocketMemoryId"] = o.SocketMemoryId
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.MemoryPersistentMemoryConfigResult.IsSet() {
		toSerialize["MemoryPersistentMemoryConfigResult"] = o.MemoryPersistentMemoryConfigResult.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryPersistentMemoryNamespaceConfigResult) UnmarshalJSON(data []byte) (err error) {
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
	type MemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Status of the Persistent Memory Namespace needed to be configured.
		ConfigStatus *string `json:"ConfigStatus,omitempty"`
		// Name of a Persistent Memory Namespace that needed to be configured.
		Name *string `json:"Name,omitempty"`
		// Socket ID in which the Persistent Memory Namespace needed to be configured.
		SocketId *string `json:"SocketId,omitempty"`
		// Socket Memory ID in which the Persistent Memory Namespace needed to be configured.
		SocketMemoryId                     *string                                                `json:"SocketMemoryId,omitempty"`
		InventoryDeviceInfo                NullableInventoryDeviceInfoRelationship                `json:"InventoryDeviceInfo,omitempty"`
		MemoryPersistentMemoryConfigResult NullableMemoryPersistentMemoryConfigResultRelationship `json:"MemoryPersistentMemoryConfigResult,omitempty"`
		RegisteredDevice                   NullableAssetDeviceRegistrationRelationship            `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct := MemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryNamespaceConfigResult := _MemoryPersistentMemoryNamespaceConfigResult{}
		varMemoryPersistentMemoryNamespaceConfigResult.ClassId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryNamespaceConfigResult.ObjectType = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryNamespaceConfigResult.ConfigStatus = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.ConfigStatus
		varMemoryPersistentMemoryNamespaceConfigResult.Name = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.Name
		varMemoryPersistentMemoryNamespaceConfigResult.SocketId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.SocketId
		varMemoryPersistentMemoryNamespaceConfigResult.SocketMemoryId = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.SocketMemoryId
		varMemoryPersistentMemoryNamespaceConfigResult.InventoryDeviceInfo = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryNamespaceConfigResult.MemoryPersistentMemoryConfigResult = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.MemoryPersistentMemoryConfigResult
		varMemoryPersistentMemoryNamespaceConfigResult.RegisteredDevice = varMemoryPersistentMemoryNamespaceConfigResultWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryNamespaceConfigResult(varMemoryPersistentMemoryNamespaceConfigResult)
	} else {
		return err
	}

	varMemoryPersistentMemoryNamespaceConfigResult := _MemoryPersistentMemoryNamespaceConfigResult{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryNamespaceConfigResult)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryNamespaceConfigResult.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigStatus")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SocketId")
		delete(additionalProperties, "SocketMemoryId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfigResult")
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

type NullableMemoryPersistentMemoryNamespaceConfigResult struct {
	value *MemoryPersistentMemoryNamespaceConfigResult
	isSet bool
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) Get() *MemoryPersistentMemoryNamespaceConfigResult {
	return v.value
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) Set(val *MemoryPersistentMemoryNamespaceConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryNamespaceConfigResult(val *MemoryPersistentMemoryNamespaceConfigResult) *NullableMemoryPersistentMemoryNamespaceConfigResult {
	return &NullableMemoryPersistentMemoryNamespaceConfigResult{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryNamespaceConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryNamespaceConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
