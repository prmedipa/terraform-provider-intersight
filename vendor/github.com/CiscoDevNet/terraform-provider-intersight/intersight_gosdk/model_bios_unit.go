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

// checks if the BiosUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiosUnit{}

// BiosUnit The BIOS Unit present on a server.
type BiosUnit struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The initSeq of the equipment.
	InitSeq *string `json:"InitSeq,omitempty"`
	// The initTs of the equipment.
	InitTs              *string                                     `json:"InitTs,omitempty"`
	ComputeBlade        NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeRackUnit     NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship   `json:"RunningFirmware,omitempty"`
	SystemBootOrder      NullableBiosSystemBootOrderRelationship `json:"SystemBootOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BiosUnit BiosUnit

// NewBiosUnit instantiates a new BiosUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiosUnit(classId string, objectType string) *BiosUnit {
	this := BiosUnit{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBiosUnitWithDefaults instantiates a new BiosUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiosUnitWithDefaults() *BiosUnit {
	this := BiosUnit{}
	var classId string = "bios.Unit"
	this.ClassId = classId
	var objectType string = "bios.Unit"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *BiosUnit) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BiosUnit) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "bios.Unit" of the ClassId field.
func (o *BiosUnit) GetDefaultClassId() interface{} {
	return "bios.Unit"
}

// GetObjectType returns the ObjectType field value
func (o *BiosUnit) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BiosUnit) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "bios.Unit" of the ObjectType field.
func (o *BiosUnit) GetDefaultObjectType() interface{} {
	return "bios.Unit"
}

// GetInitSeq returns the InitSeq field value if set, zero value otherwise.
func (o *BiosUnit) GetInitSeq() string {
	if o == nil || IsNil(o.InitSeq) {
		var ret string
		return ret
	}
	return *o.InitSeq
}

// GetInitSeqOk returns a tuple with the InitSeq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetInitSeqOk() (*string, bool) {
	if o == nil || IsNil(o.InitSeq) {
		return nil, false
	}
	return o.InitSeq, true
}

// HasInitSeq returns a boolean if a field has been set.
func (o *BiosUnit) HasInitSeq() bool {
	if o != nil && !IsNil(o.InitSeq) {
		return true
	}

	return false
}

// SetInitSeq gets a reference to the given string and assigns it to the InitSeq field.
func (o *BiosUnit) SetInitSeq(v string) {
	o.InitSeq = &v
}

// GetInitTs returns the InitTs field value if set, zero value otherwise.
func (o *BiosUnit) GetInitTs() string {
	if o == nil || IsNil(o.InitTs) {
		var ret string
		return ret
	}
	return *o.InitTs
}

// GetInitTsOk returns a tuple with the InitTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BiosUnit) GetInitTsOk() (*string, bool) {
	if o == nil || IsNil(o.InitTs) {
		return nil, false
	}
	return o.InitTs, true
}

// HasInitTs returns a boolean if a field has been set.
func (o *BiosUnit) HasInitTs() bool {
	if o != nil && !IsNil(o.InitTs) {
		return true
	}

	return false
}

// SetInitTs gets a reference to the given string and assigns it to the InitTs field.
func (o *BiosUnit) SetInitTs(v string) {
	o.InitTs = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *BiosUnit) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *BiosUnit) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *BiosUnit) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *BiosUnit) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *BiosUnit) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *BiosUnit) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *BiosUnit) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *BiosUnit) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *BiosUnit) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *BiosUnit) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *BiosUnit) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *BiosUnit) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *BiosUnit) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *BiosUnit) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *BiosUnit) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *BiosUnit) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetRunningFirmwareOk() ([]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || IsNil(o.RunningFirmware) {
		return nil, false
	}
	return o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *BiosUnit) HasRunningFirmware() bool {
	if o != nil && !IsNil(o.RunningFirmware) {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *BiosUnit) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

// GetSystemBootOrder returns the SystemBootOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BiosUnit) GetSystemBootOrder() BiosSystemBootOrderRelationship {
	if o == nil || IsNil(o.SystemBootOrder.Get()) {
		var ret BiosSystemBootOrderRelationship
		return ret
	}
	return *o.SystemBootOrder.Get()
}

// GetSystemBootOrderOk returns a tuple with the SystemBootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BiosUnit) GetSystemBootOrderOk() (*BiosSystemBootOrderRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemBootOrder.Get(), o.SystemBootOrder.IsSet()
}

// HasSystemBootOrder returns a boolean if a field has been set.
func (o *BiosUnit) HasSystemBootOrder() bool {
	if o != nil && o.SystemBootOrder.IsSet() {
		return true
	}

	return false
}

// SetSystemBootOrder gets a reference to the given NullableBiosSystemBootOrderRelationship and assigns it to the SystemBootOrder field.
func (o *BiosUnit) SetSystemBootOrder(v BiosSystemBootOrderRelationship) {
	o.SystemBootOrder.Set(&v)
}

// SetSystemBootOrderNil sets the value for SystemBootOrder to be an explicit nil
func (o *BiosUnit) SetSystemBootOrderNil() {
	o.SystemBootOrder.Set(nil)
}

// UnsetSystemBootOrder ensures that no value is present for SystemBootOrder, not even an explicit nil
func (o *BiosUnit) UnsetSystemBootOrder() {
	o.SystemBootOrder.Unset()
}

func (o BiosUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BiosUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.InitSeq) {
		toSerialize["InitSeq"] = o.InitSeq
	}
	if !IsNil(o.InitTs) {
		toSerialize["InitTs"] = o.InitTs
	}
	if o.ComputeBlade.IsSet() {
		toSerialize["ComputeBlade"] = o.ComputeBlade.Get()
	}
	if o.ComputeRackUnit.IsSet() {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}
	if o.SystemBootOrder.IsSet() {
		toSerialize["SystemBootOrder"] = o.SystemBootOrder.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BiosUnit) UnmarshalJSON(data []byte) (err error) {
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
	type BiosUnitWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The initSeq of the equipment.
		InitSeq *string `json:"InitSeq,omitempty"`
		// The initTs of the equipment.
		InitTs              *string                                     `json:"InitTs,omitempty"`
		ComputeBlade        NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeRackUnit     NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to firmwareRunningFirmware resources.
		RunningFirmware []FirmwareRunningFirmwareRelationship   `json:"RunningFirmware,omitempty"`
		SystemBootOrder NullableBiosSystemBootOrderRelationship `json:"SystemBootOrder,omitempty"`
	}

	varBiosUnitWithoutEmbeddedStruct := BiosUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBiosUnitWithoutEmbeddedStruct)
	if err == nil {
		varBiosUnit := _BiosUnit{}
		varBiosUnit.ClassId = varBiosUnitWithoutEmbeddedStruct.ClassId
		varBiosUnit.ObjectType = varBiosUnitWithoutEmbeddedStruct.ObjectType
		varBiosUnit.InitSeq = varBiosUnitWithoutEmbeddedStruct.InitSeq
		varBiosUnit.InitTs = varBiosUnitWithoutEmbeddedStruct.InitTs
		varBiosUnit.ComputeBlade = varBiosUnitWithoutEmbeddedStruct.ComputeBlade
		varBiosUnit.ComputeRackUnit = varBiosUnitWithoutEmbeddedStruct.ComputeRackUnit
		varBiosUnit.InventoryDeviceInfo = varBiosUnitWithoutEmbeddedStruct.InventoryDeviceInfo
		varBiosUnit.RegisteredDevice = varBiosUnitWithoutEmbeddedStruct.RegisteredDevice
		varBiosUnit.RunningFirmware = varBiosUnitWithoutEmbeddedStruct.RunningFirmware
		varBiosUnit.SystemBootOrder = varBiosUnitWithoutEmbeddedStruct.SystemBootOrder
		*o = BiosUnit(varBiosUnit)
	} else {
		return err
	}

	varBiosUnit := _BiosUnit{}

	err = json.Unmarshal(data, &varBiosUnit)
	if err == nil {
		o.EquipmentBase = varBiosUnit.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "InitSeq")
		delete(additionalProperties, "InitTs")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")
		delete(additionalProperties, "SystemBootOrder")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableBiosUnit struct {
	value *BiosUnit
	isSet bool
}

func (v NullableBiosUnit) Get() *BiosUnit {
	return v.value
}

func (v *NullableBiosUnit) Set(val *BiosUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableBiosUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableBiosUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiosUnit(val *BiosUnit) *NullableBiosUnit {
	return &NullableBiosUnit{value: val, isSet: true}
}

func (v NullableBiosUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiosUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
