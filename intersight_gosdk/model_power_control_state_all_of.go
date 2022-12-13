/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// PowerControlStateAllOf Definition of the list of properties defined in 'power.ControlState', excluding properties defined in parent classes.
type PowerControlStateAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This field identifies the allocated power on the chassis in Watts.
	AllocatedPower *int64 `json:"AllocatedPower,omitempty"`
	// The status of extended power capacity mode of the chassis. If Enabled, this mode allows chassis available power to be increased by borrowing power from redundant power supplies. * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	ExtendedPowerCapacity *string `json:"ExtendedPowerCapacity,omitempty"`
	// This field identifies the available power when PSUs are in grid mode in Watts.
	GridMaxPower *int64 `json:"GridMaxPower,omitempty"`
	// This field identifies the maximum power required by the endpoint in Watts.
	MaxRequiredPower *int64 `json:"MaxRequiredPower,omitempty"`
	// This field identifies the minimum power required by the endpoint in Watts.
	MinRequiredPower *int64 `json:"MinRequiredPower,omitempty"`
	// This field identifies the available power when PSUs are in N+1 mode in Watts.
	N1MaxPower *int64 `json:"N1MaxPower,omitempty"`
	// This field identifies the available power when PSUs are in N+2 mode in Watts.
	N2MaxPower *int64 `json:"N2MaxPower,omitempty"`
	// This field identifies the available power when PSUs are in non-redundant mode in Watts.
	NonRedundantMaxPower *int64 `json:"NonRedundantMaxPower,omitempty"`
	// The status of power rebalancing mode of the chassis. If enabled, this mode allows the system to dynamically reallocate the power between servers depending on their power usage.  * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	PowerRebalancing *string `json:"PowerRebalancing,omitempty"`
	// The status of power save mode of the chassis. If Enabled and the requested power budget is less than available power capacity,  the additional PSUs not required to comply with redundancy policy are placed in Power Save mode.  * `Enabled` - Set the value to Enabled. * `Disabled` - Set the value to Disabled.
	PowerSaveMode        *string                              `json:"PowerSaveMode,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PowerControlStateAllOf PowerControlStateAllOf

// NewPowerControlStateAllOf instantiates a new PowerControlStateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerControlStateAllOf(classId string, objectType string) *PowerControlStateAllOf {
	this := PowerControlStateAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPowerControlStateAllOfWithDefaults instantiates a new PowerControlStateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerControlStateAllOfWithDefaults() *PowerControlStateAllOf {
	this := PowerControlStateAllOf{}
	var classId string = "power.ControlState"
	this.ClassId = classId
	var objectType string = "power.ControlState"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PowerControlStateAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PowerControlStateAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PowerControlStateAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PowerControlStateAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAllocatedPower returns the AllocatedPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetAllocatedPower() int64 {
	if o == nil || o.AllocatedPower == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedPower
}

// GetAllocatedPowerOk returns a tuple with the AllocatedPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetAllocatedPowerOk() (*int64, bool) {
	if o == nil || o.AllocatedPower == nil {
		return nil, false
	}
	return o.AllocatedPower, true
}

// HasAllocatedPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasAllocatedPower() bool {
	if o != nil && o.AllocatedPower != nil {
		return true
	}

	return false
}

// SetAllocatedPower gets a reference to the given int64 and assigns it to the AllocatedPower field.
func (o *PowerControlStateAllOf) SetAllocatedPower(v int64) {
	o.AllocatedPower = &v
}

// GetExtendedPowerCapacity returns the ExtendedPowerCapacity field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetExtendedPowerCapacity() string {
	if o == nil || o.ExtendedPowerCapacity == nil {
		var ret string
		return ret
	}
	return *o.ExtendedPowerCapacity
}

// GetExtendedPowerCapacityOk returns a tuple with the ExtendedPowerCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetExtendedPowerCapacityOk() (*string, bool) {
	if o == nil || o.ExtendedPowerCapacity == nil {
		return nil, false
	}
	return o.ExtendedPowerCapacity, true
}

// HasExtendedPowerCapacity returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasExtendedPowerCapacity() bool {
	if o != nil && o.ExtendedPowerCapacity != nil {
		return true
	}

	return false
}

// SetExtendedPowerCapacity gets a reference to the given string and assigns it to the ExtendedPowerCapacity field.
func (o *PowerControlStateAllOf) SetExtendedPowerCapacity(v string) {
	o.ExtendedPowerCapacity = &v
}

// GetGridMaxPower returns the GridMaxPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetGridMaxPower() int64 {
	if o == nil || o.GridMaxPower == nil {
		var ret int64
		return ret
	}
	return *o.GridMaxPower
}

// GetGridMaxPowerOk returns a tuple with the GridMaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetGridMaxPowerOk() (*int64, bool) {
	if o == nil || o.GridMaxPower == nil {
		return nil, false
	}
	return o.GridMaxPower, true
}

// HasGridMaxPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasGridMaxPower() bool {
	if o != nil && o.GridMaxPower != nil {
		return true
	}

	return false
}

// SetGridMaxPower gets a reference to the given int64 and assigns it to the GridMaxPower field.
func (o *PowerControlStateAllOf) SetGridMaxPower(v int64) {
	o.GridMaxPower = &v
}

// GetMaxRequiredPower returns the MaxRequiredPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetMaxRequiredPower() int64 {
	if o == nil || o.MaxRequiredPower == nil {
		var ret int64
		return ret
	}
	return *o.MaxRequiredPower
}

// GetMaxRequiredPowerOk returns a tuple with the MaxRequiredPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetMaxRequiredPowerOk() (*int64, bool) {
	if o == nil || o.MaxRequiredPower == nil {
		return nil, false
	}
	return o.MaxRequiredPower, true
}

// HasMaxRequiredPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasMaxRequiredPower() bool {
	if o != nil && o.MaxRequiredPower != nil {
		return true
	}

	return false
}

// SetMaxRequiredPower gets a reference to the given int64 and assigns it to the MaxRequiredPower field.
func (o *PowerControlStateAllOf) SetMaxRequiredPower(v int64) {
	o.MaxRequiredPower = &v
}

// GetMinRequiredPower returns the MinRequiredPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetMinRequiredPower() int64 {
	if o == nil || o.MinRequiredPower == nil {
		var ret int64
		return ret
	}
	return *o.MinRequiredPower
}

// GetMinRequiredPowerOk returns a tuple with the MinRequiredPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetMinRequiredPowerOk() (*int64, bool) {
	if o == nil || o.MinRequiredPower == nil {
		return nil, false
	}
	return o.MinRequiredPower, true
}

// HasMinRequiredPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasMinRequiredPower() bool {
	if o != nil && o.MinRequiredPower != nil {
		return true
	}

	return false
}

// SetMinRequiredPower gets a reference to the given int64 and assigns it to the MinRequiredPower field.
func (o *PowerControlStateAllOf) SetMinRequiredPower(v int64) {
	o.MinRequiredPower = &v
}

// GetN1MaxPower returns the N1MaxPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetN1MaxPower() int64 {
	if o == nil || o.N1MaxPower == nil {
		var ret int64
		return ret
	}
	return *o.N1MaxPower
}

// GetN1MaxPowerOk returns a tuple with the N1MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetN1MaxPowerOk() (*int64, bool) {
	if o == nil || o.N1MaxPower == nil {
		return nil, false
	}
	return o.N1MaxPower, true
}

// HasN1MaxPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasN1MaxPower() bool {
	if o != nil && o.N1MaxPower != nil {
		return true
	}

	return false
}

// SetN1MaxPower gets a reference to the given int64 and assigns it to the N1MaxPower field.
func (o *PowerControlStateAllOf) SetN1MaxPower(v int64) {
	o.N1MaxPower = &v
}

// GetN2MaxPower returns the N2MaxPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetN2MaxPower() int64 {
	if o == nil || o.N2MaxPower == nil {
		var ret int64
		return ret
	}
	return *o.N2MaxPower
}

// GetN2MaxPowerOk returns a tuple with the N2MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetN2MaxPowerOk() (*int64, bool) {
	if o == nil || o.N2MaxPower == nil {
		return nil, false
	}
	return o.N2MaxPower, true
}

// HasN2MaxPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasN2MaxPower() bool {
	if o != nil && o.N2MaxPower != nil {
		return true
	}

	return false
}

// SetN2MaxPower gets a reference to the given int64 and assigns it to the N2MaxPower field.
func (o *PowerControlStateAllOf) SetN2MaxPower(v int64) {
	o.N2MaxPower = &v
}

// GetNonRedundantMaxPower returns the NonRedundantMaxPower field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetNonRedundantMaxPower() int64 {
	if o == nil || o.NonRedundantMaxPower == nil {
		var ret int64
		return ret
	}
	return *o.NonRedundantMaxPower
}

// GetNonRedundantMaxPowerOk returns a tuple with the NonRedundantMaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetNonRedundantMaxPowerOk() (*int64, bool) {
	if o == nil || o.NonRedundantMaxPower == nil {
		return nil, false
	}
	return o.NonRedundantMaxPower, true
}

// HasNonRedundantMaxPower returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasNonRedundantMaxPower() bool {
	if o != nil && o.NonRedundantMaxPower != nil {
		return true
	}

	return false
}

// SetNonRedundantMaxPower gets a reference to the given int64 and assigns it to the NonRedundantMaxPower field.
func (o *PowerControlStateAllOf) SetNonRedundantMaxPower(v int64) {
	o.NonRedundantMaxPower = &v
}

// GetPowerRebalancing returns the PowerRebalancing field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetPowerRebalancing() string {
	if o == nil || o.PowerRebalancing == nil {
		var ret string
		return ret
	}
	return *o.PowerRebalancing
}

// GetPowerRebalancingOk returns a tuple with the PowerRebalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetPowerRebalancingOk() (*string, bool) {
	if o == nil || o.PowerRebalancing == nil {
		return nil, false
	}
	return o.PowerRebalancing, true
}

// HasPowerRebalancing returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasPowerRebalancing() bool {
	if o != nil && o.PowerRebalancing != nil {
		return true
	}

	return false
}

// SetPowerRebalancing gets a reference to the given string and assigns it to the PowerRebalancing field.
func (o *PowerControlStateAllOf) SetPowerRebalancing(v string) {
	o.PowerRebalancing = &v
}

// GetPowerSaveMode returns the PowerSaveMode field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetPowerSaveMode() string {
	if o == nil || o.PowerSaveMode == nil {
		var ret string
		return ret
	}
	return *o.PowerSaveMode
}

// GetPowerSaveModeOk returns a tuple with the PowerSaveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetPowerSaveModeOk() (*string, bool) {
	if o == nil || o.PowerSaveMode == nil {
		return nil, false
	}
	return o.PowerSaveMode, true
}

// HasPowerSaveMode returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasPowerSaveMode() bool {
	if o != nil && o.PowerSaveMode != nil {
		return true
	}

	return false
}

// SetPowerSaveMode gets a reference to the given string and assigns it to the PowerSaveMode field.
func (o *PowerControlStateAllOf) SetPowerSaveMode(v string) {
	o.PowerSaveMode = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *PowerControlStateAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PowerControlStateAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerControlStateAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PowerControlStateAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PowerControlStateAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o PowerControlStateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AllocatedPower != nil {
		toSerialize["AllocatedPower"] = o.AllocatedPower
	}
	if o.ExtendedPowerCapacity != nil {
		toSerialize["ExtendedPowerCapacity"] = o.ExtendedPowerCapacity
	}
	if o.GridMaxPower != nil {
		toSerialize["GridMaxPower"] = o.GridMaxPower
	}
	if o.MaxRequiredPower != nil {
		toSerialize["MaxRequiredPower"] = o.MaxRequiredPower
	}
	if o.MinRequiredPower != nil {
		toSerialize["MinRequiredPower"] = o.MinRequiredPower
	}
	if o.N1MaxPower != nil {
		toSerialize["N1MaxPower"] = o.N1MaxPower
	}
	if o.N2MaxPower != nil {
		toSerialize["N2MaxPower"] = o.N2MaxPower
	}
	if o.NonRedundantMaxPower != nil {
		toSerialize["NonRedundantMaxPower"] = o.NonRedundantMaxPower
	}
	if o.PowerRebalancing != nil {
		toSerialize["PowerRebalancing"] = o.PowerRebalancing
	}
	if o.PowerSaveMode != nil {
		toSerialize["PowerSaveMode"] = o.PowerSaveMode
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PowerControlStateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPowerControlStateAllOf := _PowerControlStateAllOf{}

	if err = json.Unmarshal(bytes, &varPowerControlStateAllOf); err == nil {
		*o = PowerControlStateAllOf(varPowerControlStateAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllocatedPower")
		delete(additionalProperties, "ExtendedPowerCapacity")
		delete(additionalProperties, "GridMaxPower")
		delete(additionalProperties, "MaxRequiredPower")
		delete(additionalProperties, "MinRequiredPower")
		delete(additionalProperties, "N1MaxPower")
		delete(additionalProperties, "N2MaxPower")
		delete(additionalProperties, "NonRedundantMaxPower")
		delete(additionalProperties, "PowerRebalancing")
		delete(additionalProperties, "PowerSaveMode")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePowerControlStateAllOf struct {
	value *PowerControlStateAllOf
	isSet bool
}

func (v NullablePowerControlStateAllOf) Get() *PowerControlStateAllOf {
	return v.value
}

func (v *NullablePowerControlStateAllOf) Set(val *PowerControlStateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerControlStateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerControlStateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerControlStateAllOf(val *PowerControlStateAllOf) *NullablePowerControlStateAllOf {
	return &NullablePowerControlStateAllOf{value: val, isSet: true}
}

func (v NullablePowerControlStateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerControlStateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
