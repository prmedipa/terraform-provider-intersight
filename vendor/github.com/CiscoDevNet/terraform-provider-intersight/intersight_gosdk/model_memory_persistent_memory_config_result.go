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

// checks if the MemoryPersistentMemoryConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemoryPersistentMemoryConfigResult{}

// MemoryPersistentMemoryConfigResult Result of a previously applied Persistent Memory configuration on a server.
type MemoryPersistentMemoryConfigResult struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Error in the result of a previously applied Persistent Memory configuration on a server.
	ConfigErrorDesc *string `json:"ConfigErrorDesc,omitempty"`
	// Result of a previously applied Persistent Memory configuration on a server.
	ConfigResult *string `json:"ConfigResult,omitempty"`
	// Sequence number of a previously applied Persistent Memory configuration on a server.
	ConfigSequenceNo *int64 `json:"ConfigSequenceNo,omitempty"`
	// State of a previously applied Persistent Memory configuration on a server.
	ConfigState                         *string                                                 `json:"ConfigState,omitempty"`
	InventoryDeviceInfo                 NullableInventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
	MemoryPersistentMemoryConfiguration NullableMemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
	// An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources.
	PersistentMemoryNamespaceConfigResults []MemoryPersistentMemoryNamespaceConfigResultRelationship `json:"PersistentMemoryNamespaceConfigResults,omitempty"`
	RegisteredDevice                       NullableAssetDeviceRegistrationRelationship               `json:"RegisteredDevice,omitempty"`
	AdditionalProperties                   map[string]interface{}
}

type _MemoryPersistentMemoryConfigResult MemoryPersistentMemoryConfigResult

// NewMemoryPersistentMemoryConfigResult instantiates a new MemoryPersistentMemoryConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryConfigResult(classId string, objectType string) *MemoryPersistentMemoryConfigResult {
	this := MemoryPersistentMemoryConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryPersistentMemoryConfigResultWithDefaults instantiates a new MemoryPersistentMemoryConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryConfigResultWithDefaults() *MemoryPersistentMemoryConfigResult {
	this := MemoryPersistentMemoryConfigResult{}
	var classId string = "memory.PersistentMemoryConfigResult"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "memory.PersistentMemoryConfigResult" of the ClassId field.
func (o *MemoryPersistentMemoryConfigResult) GetDefaultClassId() interface{} {
	return "memory.PersistentMemoryConfigResult"
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "memory.PersistentMemoryConfigResult" of the ObjectType field.
func (o *MemoryPersistentMemoryConfigResult) GetDefaultObjectType() interface{} {
	return "memory.PersistentMemoryConfigResult"
}

// GetConfigErrorDesc returns the ConfigErrorDesc field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDesc() string {
	if o == nil || IsNil(o.ConfigErrorDesc) {
		var ret string
		return ret
	}
	return *o.ConfigErrorDesc
}

// GetConfigErrorDescOk returns a tuple with the ConfigErrorDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigErrorDescOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigErrorDesc) {
		return nil, false
	}
	return o.ConfigErrorDesc, true
}

// HasConfigErrorDesc returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigErrorDesc() bool {
	if o != nil && !IsNil(o.ConfigErrorDesc) {
		return true
	}

	return false
}

// SetConfigErrorDesc gets a reference to the given string and assigns it to the ConfigErrorDesc field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigErrorDesc(v string) {
	o.ConfigErrorDesc = &v
}

// GetConfigResult returns the ConfigResult field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigResult() string {
	if o == nil || IsNil(o.ConfigResult) {
		var ret string
		return ret
	}
	return *o.ConfigResult
}

// GetConfigResultOk returns a tuple with the ConfigResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigResultOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigResult) {
		return nil, false
	}
	return o.ConfigResult, true
}

// HasConfigResult returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigResult() bool {
	if o != nil && !IsNil(o.ConfigResult) {
		return true
	}

	return false
}

// SetConfigResult gets a reference to the given string and assigns it to the ConfigResult field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigResult(v string) {
	o.ConfigResult = &v
}

// GetConfigSequenceNo returns the ConfigSequenceNo field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNo() int64 {
	if o == nil || IsNil(o.ConfigSequenceNo) {
		var ret int64
		return ret
	}
	return *o.ConfigSequenceNo
}

// GetConfigSequenceNoOk returns a tuple with the ConfigSequenceNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigSequenceNoOk() (*int64, bool) {
	if o == nil || IsNil(o.ConfigSequenceNo) {
		return nil, false
	}
	return o.ConfigSequenceNo, true
}

// HasConfigSequenceNo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigSequenceNo() bool {
	if o != nil && !IsNil(o.ConfigSequenceNo) {
		return true
	}

	return false
}

// SetConfigSequenceNo gets a reference to the given int64 and assigns it to the ConfigSequenceNo field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigSequenceNo(v int64) {
	o.ConfigSequenceNo = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryConfigResult) GetConfigState() string {
	if o == nil || IsNil(o.ConfigState) {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryConfigResult) GetConfigStateOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigState) {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasConfigState() bool {
	if o != nil && !IsNil(o.ConfigState) {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *MemoryPersistentMemoryConfigResult) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfigResult) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryPersistentMemoryConfigResult) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *MemoryPersistentMemoryConfigResult) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *MemoryPersistentMemoryConfigResult) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetMemoryPersistentMemoryConfiguration returns the MemoryPersistentMemoryConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfiguration() MemoryPersistentMemoryConfigurationRelationship {
	if o == nil || IsNil(o.MemoryPersistentMemoryConfiguration.Get()) {
		var ret MemoryPersistentMemoryConfigurationRelationship
		return ret
	}
	return *o.MemoryPersistentMemoryConfiguration.Get()
}

// GetMemoryPersistentMemoryConfigurationOk returns a tuple with the MemoryPersistentMemoryConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfigResult) GetMemoryPersistentMemoryConfigurationOk() (*MemoryPersistentMemoryConfigurationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryPersistentMemoryConfiguration.Get(), o.MemoryPersistentMemoryConfiguration.IsSet()
}

// HasMemoryPersistentMemoryConfiguration returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasMemoryPersistentMemoryConfiguration() bool {
	if o != nil && o.MemoryPersistentMemoryConfiguration.IsSet() {
		return true
	}

	return false
}

// SetMemoryPersistentMemoryConfiguration gets a reference to the given NullableMemoryPersistentMemoryConfigurationRelationship and assigns it to the MemoryPersistentMemoryConfiguration field.
func (o *MemoryPersistentMemoryConfigResult) SetMemoryPersistentMemoryConfiguration(v MemoryPersistentMemoryConfigurationRelationship) {
	o.MemoryPersistentMemoryConfiguration.Set(&v)
}

// SetMemoryPersistentMemoryConfigurationNil sets the value for MemoryPersistentMemoryConfiguration to be an explicit nil
func (o *MemoryPersistentMemoryConfigResult) SetMemoryPersistentMemoryConfigurationNil() {
	o.MemoryPersistentMemoryConfiguration.Set(nil)
}

// UnsetMemoryPersistentMemoryConfiguration ensures that no value is present for MemoryPersistentMemoryConfiguration, not even an explicit nil
func (o *MemoryPersistentMemoryConfigResult) UnsetMemoryPersistentMemoryConfiguration() {
	o.MemoryPersistentMemoryConfiguration.Unset()
}

// GetPersistentMemoryNamespaceConfigResults returns the PersistentMemoryNamespaceConfigResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResults() []MemoryPersistentMemoryNamespaceConfigResultRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryNamespaceConfigResultRelationship
		return ret
	}
	return o.PersistentMemoryNamespaceConfigResults
}

// GetPersistentMemoryNamespaceConfigResultsOk returns a tuple with the PersistentMemoryNamespaceConfigResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfigResult) GetPersistentMemoryNamespaceConfigResultsOk() ([]MemoryPersistentMemoryNamespaceConfigResultRelationship, bool) {
	if o == nil || IsNil(o.PersistentMemoryNamespaceConfigResults) {
		return nil, false
	}
	return o.PersistentMemoryNamespaceConfigResults, true
}

// HasPersistentMemoryNamespaceConfigResults returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasPersistentMemoryNamespaceConfigResults() bool {
	if o != nil && !IsNil(o.PersistentMemoryNamespaceConfigResults) {
		return true
	}

	return false
}

// SetPersistentMemoryNamespaceConfigResults gets a reference to the given []MemoryPersistentMemoryNamespaceConfigResultRelationship and assigns it to the PersistentMemoryNamespaceConfigResults field.
func (o *MemoryPersistentMemoryConfigResult) SetPersistentMemoryNamespaceConfigResults(v []MemoryPersistentMemoryNamespaceConfigResultRelationship) {
	o.PersistentMemoryNamespaceConfigResults = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryConfigResult) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryConfigResult) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryPersistentMemoryConfigResult) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *MemoryPersistentMemoryConfigResult) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *MemoryPersistentMemoryConfigResult) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o MemoryPersistentMemoryConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemoryPersistentMemoryConfigResult) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ConfigErrorDesc) {
		toSerialize["ConfigErrorDesc"] = o.ConfigErrorDesc
	}
	if !IsNil(o.ConfigResult) {
		toSerialize["ConfigResult"] = o.ConfigResult
	}
	if !IsNil(o.ConfigSequenceNo) {
		toSerialize["ConfigSequenceNo"] = o.ConfigSequenceNo
	}
	if !IsNil(o.ConfigState) {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.MemoryPersistentMemoryConfiguration.IsSet() {
		toSerialize["MemoryPersistentMemoryConfiguration"] = o.MemoryPersistentMemoryConfiguration.Get()
	}
	if o.PersistentMemoryNamespaceConfigResults != nil {
		toSerialize["PersistentMemoryNamespaceConfigResults"] = o.PersistentMemoryNamespaceConfigResults
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemoryPersistentMemoryConfigResult) UnmarshalJSON(data []byte) (err error) {
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
	type MemoryPersistentMemoryConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Error in the result of a previously applied Persistent Memory configuration on a server.
		ConfigErrorDesc *string `json:"ConfigErrorDesc,omitempty"`
		// Result of a previously applied Persistent Memory configuration on a server.
		ConfigResult *string `json:"ConfigResult,omitempty"`
		// Sequence number of a previously applied Persistent Memory configuration on a server.
		ConfigSequenceNo *int64 `json:"ConfigSequenceNo,omitempty"`
		// State of a previously applied Persistent Memory configuration on a server.
		ConfigState                         *string                                                 `json:"ConfigState,omitempty"`
		InventoryDeviceInfo                 NullableInventoryDeviceInfoRelationship                 `json:"InventoryDeviceInfo,omitempty"`
		MemoryPersistentMemoryConfiguration NullableMemoryPersistentMemoryConfigurationRelationship `json:"MemoryPersistentMemoryConfiguration,omitempty"`
		// An array of relationships to memoryPersistentMemoryNamespaceConfigResult resources.
		PersistentMemoryNamespaceConfigResults []MemoryPersistentMemoryNamespaceConfigResultRelationship `json:"PersistentMemoryNamespaceConfigResults,omitempty"`
		RegisteredDevice                       NullableAssetDeviceRegistrationRelationship               `json:"RegisteredDevice,omitempty"`
	}

	varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct := MemoryPersistentMemoryConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryConfigResult := _MemoryPersistentMemoryConfigResult{}
		varMemoryPersistentMemoryConfigResult.ClassId = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryConfigResult.ObjectType = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryConfigResult.ConfigErrorDesc = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigErrorDesc
		varMemoryPersistentMemoryConfigResult.ConfigResult = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigResult
		varMemoryPersistentMemoryConfigResult.ConfigSequenceNo = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigSequenceNo
		varMemoryPersistentMemoryConfigResult.ConfigState = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.ConfigState
		varMemoryPersistentMemoryConfigResult.InventoryDeviceInfo = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.InventoryDeviceInfo
		varMemoryPersistentMemoryConfigResult.MemoryPersistentMemoryConfiguration = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.MemoryPersistentMemoryConfiguration
		varMemoryPersistentMemoryConfigResult.PersistentMemoryNamespaceConfigResults = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.PersistentMemoryNamespaceConfigResults
		varMemoryPersistentMemoryConfigResult.RegisteredDevice = varMemoryPersistentMemoryConfigResultWithoutEmbeddedStruct.RegisteredDevice
		*o = MemoryPersistentMemoryConfigResult(varMemoryPersistentMemoryConfigResult)
	} else {
		return err
	}

	varMemoryPersistentMemoryConfigResult := _MemoryPersistentMemoryConfigResult{}

	err = json.Unmarshal(data, &varMemoryPersistentMemoryConfigResult)
	if err == nil {
		o.InventoryBase = varMemoryPersistentMemoryConfigResult.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConfigErrorDesc")
		delete(additionalProperties, "ConfigResult")
		delete(additionalProperties, "ConfigSequenceNo")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "MemoryPersistentMemoryConfiguration")
		delete(additionalProperties, "PersistentMemoryNamespaceConfigResults")
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

type NullableMemoryPersistentMemoryConfigResult struct {
	value *MemoryPersistentMemoryConfigResult
	isSet bool
}

func (v NullableMemoryPersistentMemoryConfigResult) Get() *MemoryPersistentMemoryConfigResult {
	return v.value
}

func (v *NullableMemoryPersistentMemoryConfigResult) Set(val *MemoryPersistentMemoryConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryConfigResult(val *MemoryPersistentMemoryConfigResult) *NullableMemoryPersistentMemoryConfigResult {
	return &NullableMemoryPersistentMemoryConfigResult{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
