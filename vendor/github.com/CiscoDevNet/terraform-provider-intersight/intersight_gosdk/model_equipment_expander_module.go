/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentExpanderModule Expander module inside the chassis.
type EquipmentExpanderModule struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Module identifier for the expander module.
	ModuleId   *int64   `json:"ModuleId,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational state of expander module.
	OperState *string `json:"OperState,omitempty"`
	// Part number identifier for the expander module.
	PartNumber       *string                       `json:"PartNumber,omitempty"`
	EquipmentChassis *EquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	FanModules           []EquipmentFanModuleRelationship     `json:"FanModules,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentExpanderModule EquipmentExpanderModule

// NewEquipmentExpanderModule instantiates a new EquipmentExpanderModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentExpanderModule(classId string, objectType string) *EquipmentExpanderModule {
	this := EquipmentExpanderModule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentExpanderModuleWithDefaults instantiates a new EquipmentExpanderModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentExpanderModuleWithDefaults() *EquipmentExpanderModule {
	this := EquipmentExpanderModule{}
	var classId string = "equipment.ExpanderModule"
	this.ClassId = classId
	var objectType string = "equipment.ExpanderModule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentExpanderModule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentExpanderModule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentExpanderModule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentExpanderModule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentExpanderModule) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetOperReasonOk() ([]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *EquipmentExpanderModule) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentExpanderModule) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentExpanderModule) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentExpanderModule) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetFanModules returns the FanModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentExpanderModule) GetFanModules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.FanModules
}

// GetFanModulesOk returns a tuple with the FanModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentExpanderModule) GetFanModulesOk() ([]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.FanModules == nil {
		return nil, false
	}
	return o.FanModules, true
}

// HasFanModules returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasFanModules() bool {
	if o != nil && o.FanModules != nil {
		return true
	}

	return false
}

// SetFanModules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the FanModules field.
func (o *EquipmentExpanderModule) SetFanModules(v []EquipmentFanModuleRelationship) {
	o.FanModules = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentExpanderModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentExpanderModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentExpanderModule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentExpanderModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentExpanderModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.FanModules != nil {
		toSerialize["FanModules"] = o.FanModules
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentExpanderModule) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentExpanderModuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Module identifier for the expander module.
		ModuleId   *int64   `json:"ModuleId,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational state of expander module.
		OperState *string `json:"OperState,omitempty"`
		// Part number identifier for the expander module.
		PartNumber       *string                       `json:"PartNumber,omitempty"`
		EquipmentChassis *EquipmentChassisRelationship `json:"EquipmentChassis,omitempty"`
		// An array of relationships to equipmentFanModule resources.
		FanModules       []EquipmentFanModuleRelationship     `json:"FanModules,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentExpanderModuleWithoutEmbeddedStruct := EquipmentExpanderModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentExpanderModuleWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentExpanderModule := _EquipmentExpanderModule{}
		varEquipmentExpanderModule.ClassId = varEquipmentExpanderModuleWithoutEmbeddedStruct.ClassId
		varEquipmentExpanderModule.ObjectType = varEquipmentExpanderModuleWithoutEmbeddedStruct.ObjectType
		varEquipmentExpanderModule.ModuleId = varEquipmentExpanderModuleWithoutEmbeddedStruct.ModuleId
		varEquipmentExpanderModule.OperReason = varEquipmentExpanderModuleWithoutEmbeddedStruct.OperReason
		varEquipmentExpanderModule.OperState = varEquipmentExpanderModuleWithoutEmbeddedStruct.OperState
		varEquipmentExpanderModule.PartNumber = varEquipmentExpanderModuleWithoutEmbeddedStruct.PartNumber
		varEquipmentExpanderModule.EquipmentChassis = varEquipmentExpanderModuleWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentExpanderModule.FanModules = varEquipmentExpanderModuleWithoutEmbeddedStruct.FanModules
		varEquipmentExpanderModule.RegisteredDevice = varEquipmentExpanderModuleWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentExpanderModule(varEquipmentExpanderModule)
	} else {
		return err
	}

	varEquipmentExpanderModule := _EquipmentExpanderModule{}

	err = json.Unmarshal(bytes, &varEquipmentExpanderModule)
	if err == nil {
		o.EquipmentBase = varEquipmentExpanderModule.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "FanModules")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableEquipmentExpanderModule struct {
	value *EquipmentExpanderModule
	isSet bool
}

func (v NullableEquipmentExpanderModule) Get() *EquipmentExpanderModule {
	return v.value
}

func (v *NullableEquipmentExpanderModule) Set(val *EquipmentExpanderModule) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentExpanderModule) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentExpanderModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentExpanderModule(val *EquipmentExpanderModule) *NullableEquipmentExpanderModule {
	return &NullableEquipmentExpanderModule{value: val, isSet: true}
}

func (v NullableEquipmentExpanderModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentExpanderModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
