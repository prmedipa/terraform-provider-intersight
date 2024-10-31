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

// checks if the EquipmentRackEnclosure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentRackEnclosure{}

// EquipmentRackEnclosure A physical holder housing rack servers.
type EquipmentRackEnclosure struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This represents the Enclosure Identifier for Rack servers.
	EnclosureId *int64 `json:"EnclosureId,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	Fanmodules          []EquipmentFanModuleRelationship        `json:"Fanmodules,omitempty"`
	InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus             []EquipmentPsuRelationship                  `json:"Psus,omitempty"`
	RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to equipmentRackEnclosureSlot resources.
	Slots                []EquipmentRackEnclosureSlotRelationship `json:"Slots,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentRackEnclosure EquipmentRackEnclosure

// NewEquipmentRackEnclosure instantiates a new EquipmentRackEnclosure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentRackEnclosure(classId string, objectType string) *EquipmentRackEnclosure {
	this := EquipmentRackEnclosure{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentRackEnclosureWithDefaults instantiates a new EquipmentRackEnclosure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentRackEnclosureWithDefaults() *EquipmentRackEnclosure {
	this := EquipmentRackEnclosure{}
	var classId string = "equipment.RackEnclosure"
	this.ClassId = classId
	var objectType string = "equipment.RackEnclosure"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentRackEnclosure) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosure) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentRackEnclosure) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "equipment.RackEnclosure" of the ClassId field.
func (o *EquipmentRackEnclosure) GetDefaultClassId() interface{} {
	return "equipment.RackEnclosure"
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentRackEnclosure) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosure) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentRackEnclosure) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "equipment.RackEnclosure" of the ObjectType field.
func (o *EquipmentRackEnclosure) GetDefaultObjectType() interface{} {
	return "equipment.RackEnclosure"
}

// GetEnclosureId returns the EnclosureId field value if set, zero value otherwise.
func (o *EquipmentRackEnclosure) GetEnclosureId() int64 {
	if o == nil || IsNil(o.EnclosureId) {
		var ret int64
		return ret
	}
	return *o.EnclosureId
}

// GetEnclosureIdOk returns a tuple with the EnclosureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosure) GetEnclosureIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EnclosureId) {
		return nil, false
	}
	return o.EnclosureId, true
}

// HasEnclosureId returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasEnclosureId() bool {
	if o != nil && !IsNil(o.EnclosureId) {
		return true
	}

	return false
}

// SetEnclosureId gets a reference to the given int64 and assigns it to the EnclosureId field.
func (o *EquipmentRackEnclosure) SetEnclosureId(v int64) {
	o.EnclosureId = &v
}

// GetFanmodules returns the Fanmodules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosure) GetFanmodules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.Fanmodules
}

// GetFanmodulesOk returns a tuple with the Fanmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosure) GetFanmodulesOk() ([]EquipmentFanModuleRelationship, bool) {
	if o == nil || IsNil(o.Fanmodules) {
		return nil, false
	}
	return o.Fanmodules, true
}

// HasFanmodules returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasFanmodules() bool {
	if o != nil && !IsNil(o.Fanmodules) {
		return true
	}

	return false
}

// SetFanmodules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the Fanmodules field.
func (o *EquipmentRackEnclosure) SetFanmodules(v []EquipmentFanModuleRelationship) {
	o.Fanmodules = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosure) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosure) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentRackEnclosure) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *EquipmentRackEnclosure) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *EquipmentRackEnclosure) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetPsus returns the Psus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosure) GetPsus() []EquipmentPsuRelationship {
	if o == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosure) GetPsusOk() ([]EquipmentPsuRelationship, bool) {
	if o == nil || IsNil(o.Psus) {
		return nil, false
	}
	return o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasPsus() bool {
	if o != nil && !IsNil(o.Psus) {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *EquipmentRackEnclosure) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosure) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosure) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentRackEnclosure) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *EquipmentRackEnclosure) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *EquipmentRackEnclosure) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

// GetSlots returns the Slots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosure) GetSlots() []EquipmentRackEnclosureSlotRelationship {
	if o == nil {
		var ret []EquipmentRackEnclosureSlotRelationship
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosure) GetSlotsOk() ([]EquipmentRackEnclosureSlotRelationship, bool) {
	if o == nil || IsNil(o.Slots) {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *EquipmentRackEnclosure) HasSlots() bool {
	if o != nil && !IsNil(o.Slots) {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []EquipmentRackEnclosureSlotRelationship and assigns it to the Slots field.
func (o *EquipmentRackEnclosure) SetSlots(v []EquipmentRackEnclosureSlotRelationship) {
	o.Slots = v
}

func (o EquipmentRackEnclosure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentRackEnclosure) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.EnclosureId) {
		toSerialize["EnclosureId"] = o.EnclosureId
	}
	if o.Fanmodules != nil {
		toSerialize["Fanmodules"] = o.Fanmodules
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}
	if o.Slots != nil {
		toSerialize["Slots"] = o.Slots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentRackEnclosure) UnmarshalJSON(data []byte) (err error) {
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
	type EquipmentRackEnclosureWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This represents the Enclosure Identifier for Rack servers.
		EnclosureId *int64 `json:"EnclosureId,omitempty"`
		// An array of relationships to equipmentFanModule resources.
		Fanmodules          []EquipmentFanModuleRelationship        `json:"Fanmodules,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to equipmentPsu resources.
		Psus             []EquipmentPsuRelationship                  `json:"Psus,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to equipmentRackEnclosureSlot resources.
		Slots []EquipmentRackEnclosureSlotRelationship `json:"Slots,omitempty"`
	}

	varEquipmentRackEnclosureWithoutEmbeddedStruct := EquipmentRackEnclosureWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentRackEnclosureWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentRackEnclosure := _EquipmentRackEnclosure{}
		varEquipmentRackEnclosure.ClassId = varEquipmentRackEnclosureWithoutEmbeddedStruct.ClassId
		varEquipmentRackEnclosure.ObjectType = varEquipmentRackEnclosureWithoutEmbeddedStruct.ObjectType
		varEquipmentRackEnclosure.EnclosureId = varEquipmentRackEnclosureWithoutEmbeddedStruct.EnclosureId
		varEquipmentRackEnclosure.Fanmodules = varEquipmentRackEnclosureWithoutEmbeddedStruct.Fanmodules
		varEquipmentRackEnclosure.InventoryDeviceInfo = varEquipmentRackEnclosureWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentRackEnclosure.Psus = varEquipmentRackEnclosureWithoutEmbeddedStruct.Psus
		varEquipmentRackEnclosure.RegisteredDevice = varEquipmentRackEnclosureWithoutEmbeddedStruct.RegisteredDevice
		varEquipmentRackEnclosure.Slots = varEquipmentRackEnclosureWithoutEmbeddedStruct.Slots
		*o = EquipmentRackEnclosure(varEquipmentRackEnclosure)
	} else {
		return err
	}

	varEquipmentRackEnclosure := _EquipmentRackEnclosure{}

	err = json.Unmarshal(data, &varEquipmentRackEnclosure)
	if err == nil {
		o.EquipmentBase = varEquipmentRackEnclosure.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnclosureId")
		delete(additionalProperties, "Fanmodules")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Psus")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Slots")

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

type NullableEquipmentRackEnclosure struct {
	value *EquipmentRackEnclosure
	isSet bool
}

func (v NullableEquipmentRackEnclosure) Get() *EquipmentRackEnclosure {
	return v.value
}

func (v *NullableEquipmentRackEnclosure) Set(val *EquipmentRackEnclosure) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentRackEnclosure) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentRackEnclosure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentRackEnclosure(val *EquipmentRackEnclosure) *NullableEquipmentRackEnclosure {
	return &NullableEquipmentRackEnclosure{value: val, isSet: true}
}

func (v NullableEquipmentRackEnclosure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentRackEnclosure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
