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

// checks if the EquipmentHybridDriveSlot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentHybridDriveSlot{}

// EquipmentHybridDriveSlot NVMe HybridDriveSlots present in a server.
type EquipmentHybridDriveSlot struct {
	EquipmentSlot
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Configured Mode of the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
	CurrentMode *string `json:"CurrentMode,omitempty"`
	// The Requested Mode for the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
	RequestedMode        *string                                     `json:"RequestedMode,omitempty"`
	ComputeBlade         NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	ComputeBoard         NullableComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
	ComputeRackUnit      NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentHybridDriveSlot EquipmentHybridDriveSlot

// NewEquipmentHybridDriveSlot instantiates a new EquipmentHybridDriveSlot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentHybridDriveSlot(classId string, objectType string) *EquipmentHybridDriveSlot {
	this := EquipmentHybridDriveSlot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentHybridDriveSlotWithDefaults instantiates a new EquipmentHybridDriveSlot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentHybridDriveSlotWithDefaults() *EquipmentHybridDriveSlot {
	this := EquipmentHybridDriveSlot{}
	var classId string = "equipment.HybridDriveSlot"
	this.ClassId = classId
	var objectType string = "equipment.HybridDriveSlot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentHybridDriveSlot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentHybridDriveSlot) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.HybridDriveSlot" of the ClassId field.
func (o *EquipmentHybridDriveSlot) GetDefaultClassId() interface{} {
	return "equipment.HybridDriveSlot"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentHybridDriveSlot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentHybridDriveSlot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.HybridDriveSlot" of the ObjectType field.
func (o *EquipmentHybridDriveSlot) GetDefaultObjectType() interface{} {
	return "equipment.HybridDriveSlot"
}

// GetCurrentMode returns the CurrentMode field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlot) GetCurrentMode() string {
	if o == nil || IsNil(o.CurrentMode) {
		var ret string
		return ret
	}
	return *o.CurrentMode
}

// GetCurrentModeOk returns a tuple with the CurrentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlot) GetCurrentModeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentMode) {
		return nil, false
	}
	return o.CurrentMode, true
}

// HasCurrentMode returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasCurrentMode() bool {
	if o != nil && !IsNil(o.CurrentMode) {
		return true
	}

	return false
}

// SetCurrentMode gets a reference to the given string and assigns it to the CurrentMode field.
func (o *EquipmentHybridDriveSlot) SetCurrentMode(v string) {
	o.CurrentMode = &v
}

// GetRequestedMode returns the RequestedMode field value if set, zero value otherwise.
func (o *EquipmentHybridDriveSlot) GetRequestedMode() string {
	if o == nil || IsNil(o.RequestedMode) {
		var ret string
		return ret
	}
	return *o.RequestedMode
}

// GetRequestedModeOk returns a tuple with the RequestedMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentHybridDriveSlot) GetRequestedModeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestedMode) {
		return nil, false
	}
	return o.RequestedMode, true
}

// HasRequestedMode returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasRequestedMode() bool {
	if o != nil && !IsNil(o.RequestedMode) {
		return true
	}

	return false
}

// SetRequestedMode gets a reference to the given string and assigns it to the RequestedMode field.
func (o *EquipmentHybridDriveSlot) SetRequestedMode(v string) {
	o.RequestedMode = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentHybridDriveSlot) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || IsNil(o.ComputeBlade.Get()) {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade.Get()
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentHybridDriveSlot) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBlade.Get(), o.ComputeBlade.IsSet()
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade.IsSet() {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given NullableComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *EquipmentHybridDriveSlot) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade.Set(&v)
}

// SetComputeBladeNil sets the value for ComputeBlade to be an explicit nil
func (o *EquipmentHybridDriveSlot) SetComputeBladeNil() {
	o.ComputeBlade.Set(nil)
}

// UnsetComputeBlade ensures that no value is present for ComputeBlade, not even an explicit nil
func (o *EquipmentHybridDriveSlot) UnsetComputeBlade() {
	o.ComputeBlade.Unset()
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentHybridDriveSlot) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentHybridDriveSlot) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *EquipmentHybridDriveSlot) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *EquipmentHybridDriveSlot) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *EquipmentHybridDriveSlot) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentHybridDriveSlot) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || IsNil(o.ComputeRackUnit.Get()) {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit.Get()
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentHybridDriveSlot) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeRackUnit.Get(), o.ComputeRackUnit.IsSet()
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit.IsSet() {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given NullableComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentHybridDriveSlot) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit.Set(&v)
}

// SetComputeRackUnitNil sets the value for ComputeRackUnit to be an explicit nil
func (o *EquipmentHybridDriveSlot) SetComputeRackUnitNil() {
	o.ComputeRackUnit.Set(nil)
}

// UnsetComputeRackUnit ensures that no value is present for ComputeRackUnit, not even an explicit nil
func (o *EquipmentHybridDriveSlot) UnsetComputeRackUnit() {
	o.ComputeRackUnit.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentHybridDriveSlot) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentHybridDriveSlot) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentHybridDriveSlot) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentHybridDriveSlot) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentHybridDriveSlot) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentHybridDriveSlot) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o EquipmentHybridDriveSlot) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentHybridDriveSlot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentSlot, errEquipmentSlot := json.Marshal(o.EquipmentSlot)
	if errEquipmentSlot != nil {
		return map[string]interface{}{}, errEquipmentSlot
	}
	errEquipmentSlot = json.Unmarshal([]byte(serializedEquipmentSlot), &toSerialize)
	if errEquipmentSlot != nil {
		return map[string]interface{}{}, errEquipmentSlot
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CurrentMode) {
		toSerialize["CurrentMode"] = o.CurrentMode
	}
	if !IsNil(o.RequestedMode) {
		toSerialize["RequestedMode"] = o.RequestedMode
	}
	if o.ComputeBlade.IsSet() {
		toSerialize["ComputeBlade"] = o.ComputeBlade.Get()
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.ComputeRackUnit.IsSet() {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentHybridDriveSlot) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentHybridDriveSlotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Configured Mode of the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
		CurrentMode *string `json:"CurrentMode,omitempty"`
		// The Requested Mode for the Hybrid Drive slot. * `` - Hybrid Drive slot  mode is not applicable. * `RAID` - Hybrid Drive slot mode is RAID. * `Direct` - Hybrid Drive slot mode is Direct.
		RequestedMode    *string                                     `json:"RequestedMode,omitempty"`
		ComputeBlade     NullableComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
		ComputeBoard     NullableComputeBoardRelationship            `json:"ComputeBoard,omitempty"`
		ComputeRackUnit  NullableComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentHybridDriveSlotWithoutEmbeddedStruct := EquipmentHybridDriveSlotWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentHybridDriveSlotWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentHybridDriveSlot := _EquipmentHybridDriveSlot{}
		varEquipmentHybridDriveSlot.ClassId = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.ClassId
		varEquipmentHybridDriveSlot.ObjectType = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.ObjectType
		varEquipmentHybridDriveSlot.CurrentMode = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.CurrentMode
		varEquipmentHybridDriveSlot.RequestedMode = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.RequestedMode
		varEquipmentHybridDriveSlot.ComputeBlade = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.ComputeBlade
		varEquipmentHybridDriveSlot.ComputeBoard = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.ComputeBoard
		varEquipmentHybridDriveSlot.ComputeRackUnit = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.ComputeRackUnit
		varEquipmentHybridDriveSlot.RegisteredDevice = varEquipmentHybridDriveSlotWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentHybridDriveSlot(varEquipmentHybridDriveSlot)
	} else {
		return err
	}

	varEquipmentHybridDriveSlot := _EquipmentHybridDriveSlot{}

	err = json.Unmarshal(data, &varEquipmentHybridDriveSlot)
	if err == nil {
		o.EquipmentSlot = varEquipmentHybridDriveSlot.EquipmentSlot
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentMode")
		delete(additionalProperties, "RequestedMode")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentSlot := reflect.ValueOf(o.EquipmentSlot)
		for i := 0; i < reflectEquipmentSlot.Type().NumField(); i++ {
			t := reflectEquipmentSlot.Type().Field(i)

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

type NullableEquipmentHybridDriveSlot struct {
	value *EquipmentHybridDriveSlot
	isSet bool
}

func (v NullableEquipmentHybridDriveSlot) Get() *EquipmentHybridDriveSlot {
	return v.value
}

func (v *NullableEquipmentHybridDriveSlot) Set(val *EquipmentHybridDriveSlot) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentHybridDriveSlot) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentHybridDriveSlot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentHybridDriveSlot(val *EquipmentHybridDriveSlot) *NullableEquipmentHybridDriveSlot {
	return &NullableEquipmentHybridDriveSlot{value: val, isSet: true}
}

func (v NullableEquipmentHybridDriveSlot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentHybridDriveSlot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
