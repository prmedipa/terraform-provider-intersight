/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the NetworkFcZoneInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFcZoneInfo{}

// NetworkFcZoneInfo FC Zone information of a Fabric Interconnect.
type NetworkFcZoneInfo struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The number of Fibre Channel user zones defined on a Fabric Interconnect.
	UserZoneCount *int64 `json:"UserZoneCount,omitempty"`
	// The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect.
	UserZoneLimit *int64 `json:"UserZoneLimit,omitempty"`
	// The number of Fibre Channel zones defined on a Fabric Interconnect.
	ZoneCount *int64 `json:"ZoneCount,omitempty"`
	// The maximum number of Fibre Channel zones allowed on a Fabric Interconnect.
	ZoneLimit            *int64                                      `json:"ZoneLimit,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement       NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkFcZoneInfo NetworkFcZoneInfo

// NewNetworkFcZoneInfo instantiates a new NetworkFcZoneInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFcZoneInfo(classId string, objectType string) *NetworkFcZoneInfo {
	this := NetworkFcZoneInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkFcZoneInfoWithDefaults instantiates a new NetworkFcZoneInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFcZoneInfoWithDefaults() *NetworkFcZoneInfo {
	this := NetworkFcZoneInfo{}
	var classId string = "network.FcZoneInfo"
	this.ClassId = classId
	var objectType string = "network.FcZoneInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkFcZoneInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkFcZoneInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "network.FcZoneInfo" of the ClassId field.
func (o *NetworkFcZoneInfo) GetDefaultClassId() interface{} {
	return "network.FcZoneInfo"
}

// GetObjectType returns the ObjectType field value
func (o *NetworkFcZoneInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkFcZoneInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "network.FcZoneInfo" of the ObjectType field.
func (o *NetworkFcZoneInfo) GetDefaultObjectType() interface{} {
	return "network.FcZoneInfo"
}

// GetUserZoneCount returns the UserZoneCount field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetUserZoneCount() int64 {
	if o == nil || IsNil(o.UserZoneCount) {
		var ret int64
		return ret
	}
	return *o.UserZoneCount
}

// GetUserZoneCountOk returns a tuple with the UserZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetUserZoneCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UserZoneCount) {
		return nil, false
	}
	return o.UserZoneCount, true
}

// HasUserZoneCount returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasUserZoneCount() bool {
	if o != nil && !IsNil(o.UserZoneCount) {
		return true
	}

	return false
}

// SetUserZoneCount gets a reference to the given int64 and assigns it to the UserZoneCount field.
func (o *NetworkFcZoneInfo) SetUserZoneCount(v int64) {
	o.UserZoneCount = &v
}

// GetUserZoneLimit returns the UserZoneLimit field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetUserZoneLimit() int64 {
	if o == nil || IsNil(o.UserZoneLimit) {
		var ret int64
		return ret
	}
	return *o.UserZoneLimit
}

// GetUserZoneLimitOk returns a tuple with the UserZoneLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetUserZoneLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.UserZoneLimit) {
		return nil, false
	}
	return o.UserZoneLimit, true
}

// HasUserZoneLimit returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasUserZoneLimit() bool {
	if o != nil && !IsNil(o.UserZoneLimit) {
		return true
	}

	return false
}

// SetUserZoneLimit gets a reference to the given int64 and assigns it to the UserZoneLimit field.
func (o *NetworkFcZoneInfo) SetUserZoneLimit(v int64) {
	o.UserZoneLimit = &v
}

// GetZoneCount returns the ZoneCount field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetZoneCount() int64 {
	if o == nil || IsNil(o.ZoneCount) {
		var ret int64
		return ret
	}
	return *o.ZoneCount
}

// GetZoneCountOk returns a tuple with the ZoneCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetZoneCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ZoneCount) {
		return nil, false
	}
	return o.ZoneCount, true
}

// HasZoneCount returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasZoneCount() bool {
	if o != nil && !IsNil(o.ZoneCount) {
		return true
	}

	return false
}

// SetZoneCount gets a reference to the given int64 and assigns it to the ZoneCount field.
func (o *NetworkFcZoneInfo) SetZoneCount(v int64) {
	o.ZoneCount = &v
}

// GetZoneLimit returns the ZoneLimit field value if set, zero value otherwise.
func (o *NetworkFcZoneInfo) GetZoneLimit() int64 {
	if o == nil || IsNil(o.ZoneLimit) {
		var ret int64
		return ret
	}
	return *o.ZoneLimit
}

// GetZoneLimitOk returns a tuple with the ZoneLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFcZoneInfo) GetZoneLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.ZoneLimit) {
		return nil, false
	}
	return o.ZoneLimit, true
}

// HasZoneLimit returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasZoneLimit() bool {
	if o != nil && !IsNil(o.ZoneLimit) {
		return true
	}

	return false
}

// SetZoneLimit gets a reference to the given int64 and assigns it to the ZoneLimit field.
func (o *NetworkFcZoneInfo) SetZoneLimit(v int64) {
	o.ZoneLimit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkFcZoneInfo) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkFcZoneInfo) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *NetworkFcZoneInfo) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *NetworkFcZoneInfo) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *NetworkFcZoneInfo) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkFcZoneInfo) GetNetworkElement() NetworkElementRelationship {
	if o == nil || IsNil(o.NetworkElement.Get()) {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement.Get()
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkFcZoneInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkElement.Get(), o.NetworkElement.IsSet()
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasNetworkElement() bool {
	if o != nil && o.NetworkElement.IsSet() {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NullableNetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkFcZoneInfo) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement.Set(&v)
}

// SetNetworkElementNil sets the value for NetworkElement to be an explicit nil
func (o *NetworkFcZoneInfo) SetNetworkElementNil() {
	o.NetworkElement.Set(nil)
}

// UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
func (o *NetworkFcZoneInfo) UnsetNetworkElement() {
	o.NetworkElement.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkFcZoneInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkFcZoneInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkFcZoneInfo) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkFcZoneInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NetworkFcZoneInfo) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NetworkFcZoneInfo) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NetworkFcZoneInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFcZoneInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.UserZoneCount) {
		toSerialize["UserZoneCount"] = o.UserZoneCount
	}
	if !IsNil(o.UserZoneLimit) {
		toSerialize["UserZoneLimit"] = o.UserZoneLimit
	}
	if !IsNil(o.ZoneCount) {
		toSerialize["ZoneCount"] = o.ZoneCount
	}
	if !IsNil(o.ZoneLimit) {
		toSerialize["ZoneLimit"] = o.ZoneLimit
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.NetworkElement.IsSet() {
		toSerialize["NetworkElement"] = o.NetworkElement.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkFcZoneInfo) UnmarshalJSON(data []byte) (err error) {
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
	type NetworkFcZoneInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The number of Fibre Channel user zones defined on a Fabric Interconnect.
		UserZoneCount *int64 `json:"UserZoneCount,omitempty"`
		// The maximum number of Fibre Channel user zones allowed on a Fabric Interconnect.
		UserZoneLimit *int64 `json:"UserZoneLimit,omitempty"`
		// The number of Fibre Channel zones defined on a Fabric Interconnect.
		ZoneCount *int64 `json:"ZoneCount,omitempty"`
		// The maximum number of Fibre Channel zones allowed on a Fabric Interconnect.
		ZoneLimit           *int64                                      `json:"ZoneLimit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement      NullableNetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkFcZoneInfoWithoutEmbeddedStruct := NetworkFcZoneInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNetworkFcZoneInfoWithoutEmbeddedStruct)
	if err == nil {
		varNetworkFcZoneInfo := _NetworkFcZoneInfo{}
		varNetworkFcZoneInfo.ClassId = varNetworkFcZoneInfoWithoutEmbeddedStruct.ClassId
		varNetworkFcZoneInfo.ObjectType = varNetworkFcZoneInfoWithoutEmbeddedStruct.ObjectType
		varNetworkFcZoneInfo.UserZoneCount = varNetworkFcZoneInfoWithoutEmbeddedStruct.UserZoneCount
		varNetworkFcZoneInfo.UserZoneLimit = varNetworkFcZoneInfoWithoutEmbeddedStruct.UserZoneLimit
		varNetworkFcZoneInfo.ZoneCount = varNetworkFcZoneInfoWithoutEmbeddedStruct.ZoneCount
		varNetworkFcZoneInfo.ZoneLimit = varNetworkFcZoneInfoWithoutEmbeddedStruct.ZoneLimit
		varNetworkFcZoneInfo.InventoryDeviceInfo = varNetworkFcZoneInfoWithoutEmbeddedStruct.InventoryDeviceInfo
		varNetworkFcZoneInfo.NetworkElement = varNetworkFcZoneInfoWithoutEmbeddedStruct.NetworkElement
		varNetworkFcZoneInfo.RegisteredDevice = varNetworkFcZoneInfoWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkFcZoneInfo(varNetworkFcZoneInfo)
	} else {
		return err
	}

	varNetworkFcZoneInfo := _NetworkFcZoneInfo{}

	err = json.Unmarshal(data, &varNetworkFcZoneInfo)
	if err == nil {
		o.InventoryBase = varNetworkFcZoneInfo.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "UserZoneCount")
		delete(additionalProperties, "UserZoneLimit")
		delete(additionalProperties, "ZoneCount")
		delete(additionalProperties, "ZoneLimit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
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

type NullableNetworkFcZoneInfo struct {
	value *NetworkFcZoneInfo
	isSet bool
}

func (v NullableNetworkFcZoneInfo) Get() *NetworkFcZoneInfo {
	return v.value
}

func (v *NullableNetworkFcZoneInfo) Set(val *NetworkFcZoneInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFcZoneInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFcZoneInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFcZoneInfo(val *NetworkFcZoneInfo) *NullableNetworkFcZoneInfo {
	return &NullableNetworkFcZoneInfo{value: val, isSet: true}
}

func (v NullableNetworkFcZoneInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFcZoneInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
