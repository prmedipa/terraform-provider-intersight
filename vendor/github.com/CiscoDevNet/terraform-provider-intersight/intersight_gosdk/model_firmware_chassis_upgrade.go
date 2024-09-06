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

// checks if the FirmwareChassisUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirmwareChassisUpgrade{}

// FirmwareChassisUpgrade Firmware upgrade operation for chassis that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
type FirmwareChassisUpgrade struct {
	FirmwareUpgradeBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                      `json:"ObjectType"`
	ExcludeComponentList []string                                    `json:"ExcludeComponentList,omitempty"`
	Chassis              NullableEquipmentChassisRelationship        `json:"Chassis,omitempty"`
	Device               NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareChassisUpgrade FirmwareChassisUpgrade

// NewFirmwareChassisUpgrade instantiates a new FirmwareChassisUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareChassisUpgrade(classId string, objectType string) *FirmwareChassisUpgrade {
	this := FirmwareChassisUpgrade{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "NONE"
	this.Status = &status
	var upgradeType string = "direct_upgrade"
	this.UpgradeType = &upgradeType
	return &this
}

// NewFirmwareChassisUpgradeWithDefaults instantiates a new FirmwareChassisUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareChassisUpgradeWithDefaults() *FirmwareChassisUpgrade {
	this := FirmwareChassisUpgrade{}
	var classId string = "firmware.ChassisUpgrade"
	this.ClassId = classId
	var objectType string = "firmware.ChassisUpgrade"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareChassisUpgrade) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgrade) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareChassisUpgrade) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "firmware.ChassisUpgrade" of the ClassId field.
func (o *FirmwareChassisUpgrade) GetDefaultClassId() interface{} {
	return "firmware.ChassisUpgrade"
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareChassisUpgrade) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareChassisUpgrade) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareChassisUpgrade) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "firmware.ChassisUpgrade" of the ObjectType field.
func (o *FirmwareChassisUpgrade) GetDefaultObjectType() interface{} {
	return "firmware.ChassisUpgrade"
}

// GetExcludeComponentList returns the ExcludeComponentList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareChassisUpgrade) GetExcludeComponentList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExcludeComponentList
}

// GetExcludeComponentListOk returns a tuple with the ExcludeComponentList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareChassisUpgrade) GetExcludeComponentListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeComponentList) {
		return nil, false
	}
	return o.ExcludeComponentList, true
}

// HasExcludeComponentList returns a boolean if a field has been set.
func (o *FirmwareChassisUpgrade) HasExcludeComponentList() bool {
	if o != nil && !IsNil(o.ExcludeComponentList) {
		return true
	}

	return false
}

// SetExcludeComponentList gets a reference to the given []string and assigns it to the ExcludeComponentList field.
func (o *FirmwareChassisUpgrade) SetExcludeComponentList(v []string) {
	o.ExcludeComponentList = v
}

// GetChassis returns the Chassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareChassisUpgrade) GetChassis() EquipmentChassisRelationship {
	if o == nil || IsNil(o.Chassis.Get()) {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis.Get()
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareChassisUpgrade) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chassis.Get(), o.Chassis.IsSet()
}

// HasChassis returns a boolean if a field has been set.
func (o *FirmwareChassisUpgrade) HasChassis() bool {
	if o != nil && o.Chassis.IsSet() {
		return true
	}

	return false
}

// SetChassis gets a reference to the given NullableEquipmentChassisRelationship and assigns it to the Chassis field.
func (o *FirmwareChassisUpgrade) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis.Set(&v)
}

// SetChassisNil sets the value for Chassis to be an explicit nil
func (o *FirmwareChassisUpgrade) SetChassisNil() {
	o.Chassis.Set(nil)
}

// UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
func (o *FirmwareChassisUpgrade) UnsetChassis() {
	o.Chassis.Unset()
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareChassisUpgrade) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareChassisUpgrade) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *FirmwareChassisUpgrade) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *FirmwareChassisUpgrade) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *FirmwareChassisUpgrade) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *FirmwareChassisUpgrade) UnsetDevice() {
	o.Device.Unset()
}

func (o FirmwareChassisUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirmwareChassisUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareUpgradeBase, errFirmwareUpgradeBase := json.Marshal(o.FirmwareUpgradeBase)
	if errFirmwareUpgradeBase != nil {
		return map[string]interface{}{}, errFirmwareUpgradeBase
	}
	errFirmwareUpgradeBase = json.Unmarshal([]byte(serializedFirmwareUpgradeBase), &toSerialize)
	if errFirmwareUpgradeBase != nil {
		return map[string]interface{}{}, errFirmwareUpgradeBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.ExcludeComponentList != nil {
		toSerialize["ExcludeComponentList"] = o.ExcludeComponentList
	}
	if o.Chassis.IsSet() {
		toSerialize["Chassis"] = o.Chassis.Get()
	}
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirmwareChassisUpgrade) UnmarshalJSON(data []byte) (err error) {
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
	type FirmwareChassisUpgradeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                                      `json:"ObjectType"`
		ExcludeComponentList []string                                    `json:"ExcludeComponentList,omitempty"`
		Chassis              NullableEquipmentChassisRelationship        `json:"Chassis,omitempty"`
		Device               NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	}

	varFirmwareChassisUpgradeWithoutEmbeddedStruct := FirmwareChassisUpgradeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFirmwareChassisUpgradeWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareChassisUpgrade := _FirmwareChassisUpgrade{}
		varFirmwareChassisUpgrade.ClassId = varFirmwareChassisUpgradeWithoutEmbeddedStruct.ClassId
		varFirmwareChassisUpgrade.ObjectType = varFirmwareChassisUpgradeWithoutEmbeddedStruct.ObjectType
		varFirmwareChassisUpgrade.ExcludeComponentList = varFirmwareChassisUpgradeWithoutEmbeddedStruct.ExcludeComponentList
		varFirmwareChassisUpgrade.Chassis = varFirmwareChassisUpgradeWithoutEmbeddedStruct.Chassis
		varFirmwareChassisUpgrade.Device = varFirmwareChassisUpgradeWithoutEmbeddedStruct.Device
		*o = FirmwareChassisUpgrade(varFirmwareChassisUpgrade)
	} else {
		return err
	}

	varFirmwareChassisUpgrade := _FirmwareChassisUpgrade{}

	err = json.Unmarshal(data, &varFirmwareChassisUpgrade)
	if err == nil {
		o.FirmwareUpgradeBase = varFirmwareChassisUpgrade.FirmwareUpgradeBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExcludeComponentList")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "Device")

		// remove fields from embedded structs
		reflectFirmwareUpgradeBase := reflect.ValueOf(o.FirmwareUpgradeBase)
		for i := 0; i < reflectFirmwareUpgradeBase.Type().NumField(); i++ {
			t := reflectFirmwareUpgradeBase.Type().Field(i)

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

type NullableFirmwareChassisUpgrade struct {
	value *FirmwareChassisUpgrade
	isSet bool
}

func (v NullableFirmwareChassisUpgrade) Get() *FirmwareChassisUpgrade {
	return v.value
}

func (v *NullableFirmwareChassisUpgrade) Set(val *FirmwareChassisUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareChassisUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareChassisUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareChassisUpgrade(val *FirmwareChassisUpgrade) *NullableFirmwareChassisUpgrade {
	return &NullableFirmwareChassisUpgrade{value: val, isSet: true}
}

func (v NullableFirmwareChassisUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareChassisUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
