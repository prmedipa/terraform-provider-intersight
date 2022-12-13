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
	"reflect"
	"strings"
)

// AdapterHostIscsiInterface Iscsi interface on a server adapter.
type AdapterHostIscsiInterface struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin Configured State of Host ISCSI Interface.
	AdminState *string `json:"AdminState,omitempty"`
	// The Endpoint Config Dn of the Host ISCSI Interface.
	EpDn *string `json:"EpDn,omitempty"`
	// Identifier of the Host ISCSI Interface.
	HostIscsiInterfaceId *int64 `json:"HostIscsiInterfaceId,omitempty"`
	// The visibility of the Host to the endpoint.
	HostVisible *string `json:"HostVisible,omitempty"`
	// MAC address of Host ISCSI Interface.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the Host ISCSI Interface.
	Name *string `json:"Name,omitempty"`
	// Operational State of Host ISCSI Interface.
	OperState *string `json:"OperState,omitempty"`
	// Operability status of Host ISCSI Interface.
	Operability *string `json:"Operability,omitempty"`
	// PeerPort Dn of Host ISCSI Interface.
	PeerDn               *string                              `json:"PeerDn,omitempty"`
	AdapterUnit          *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterHostIscsiInterface AdapterHostIscsiInterface

// NewAdapterHostIscsiInterface instantiates a new AdapterHostIscsiInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterHostIscsiInterface(classId string, objectType string) *AdapterHostIscsiInterface {
	this := AdapterHostIscsiInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterHostIscsiInterfaceWithDefaults instantiates a new AdapterHostIscsiInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterHostIscsiInterfaceWithDefaults() *AdapterHostIscsiInterface {
	this := AdapterHostIscsiInterface{}
	var classId string = "adapter.HostIscsiInterface"
	this.ClassId = classId
	var objectType string = "adapter.HostIscsiInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterHostIscsiInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterHostIscsiInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterHostIscsiInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterHostIscsiInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *AdapterHostIscsiInterface) SetAdminState(v string) {
	o.AdminState = &v
}

// GetEpDn returns the EpDn field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetEpDn() string {
	if o == nil || o.EpDn == nil {
		var ret string
		return ret
	}
	return *o.EpDn
}

// GetEpDnOk returns a tuple with the EpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetEpDnOk() (*string, bool) {
	if o == nil || o.EpDn == nil {
		return nil, false
	}
	return o.EpDn, true
}

// HasEpDn returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasEpDn() bool {
	if o != nil && o.EpDn != nil {
		return true
	}

	return false
}

// SetEpDn gets a reference to the given string and assigns it to the EpDn field.
func (o *AdapterHostIscsiInterface) SetEpDn(v string) {
	o.EpDn = &v
}

// GetHostIscsiInterfaceId returns the HostIscsiInterfaceId field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetHostIscsiInterfaceId() int64 {
	if o == nil || o.HostIscsiInterfaceId == nil {
		var ret int64
		return ret
	}
	return *o.HostIscsiInterfaceId
}

// GetHostIscsiInterfaceIdOk returns a tuple with the HostIscsiInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetHostIscsiInterfaceIdOk() (*int64, bool) {
	if o == nil || o.HostIscsiInterfaceId == nil {
		return nil, false
	}
	return o.HostIscsiInterfaceId, true
}

// HasHostIscsiInterfaceId returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasHostIscsiInterfaceId() bool {
	if o != nil && o.HostIscsiInterfaceId != nil {
		return true
	}

	return false
}

// SetHostIscsiInterfaceId gets a reference to the given int64 and assigns it to the HostIscsiInterfaceId field.
func (o *AdapterHostIscsiInterface) SetHostIscsiInterfaceId(v int64) {
	o.HostIscsiInterfaceId = &v
}

// GetHostVisible returns the HostVisible field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetHostVisible() string {
	if o == nil || o.HostVisible == nil {
		var ret string
		return ret
	}
	return *o.HostVisible
}

// GetHostVisibleOk returns a tuple with the HostVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetHostVisibleOk() (*string, bool) {
	if o == nil || o.HostVisible == nil {
		return nil, false
	}
	return o.HostVisible, true
}

// HasHostVisible returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasHostVisible() bool {
	if o != nil && o.HostVisible != nil {
		return true
	}

	return false
}

// SetHostVisible gets a reference to the given string and assigns it to the HostVisible field.
func (o *AdapterHostIscsiInterface) SetHostVisible(v string) {
	o.HostVisible = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *AdapterHostIscsiInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdapterHostIscsiInterface) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *AdapterHostIscsiInterface) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *AdapterHostIscsiInterface) SetOperability(v string) {
	o.Operability = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *AdapterHostIscsiInterface) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *AdapterHostIscsiInterface) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterHostIscsiInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterHostIscsiInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostIscsiInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterHostIscsiInterface) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterHostIscsiInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterHostIscsiInterface) MarshalJSON() ([]byte, error) {
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
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.EpDn != nil {
		toSerialize["EpDn"] = o.EpDn
	}
	if o.HostIscsiInterfaceId != nil {
		toSerialize["HostIscsiInterfaceId"] = o.HostIscsiInterfaceId
	}
	if o.HostVisible != nil {
		toSerialize["HostVisible"] = o.HostVisible
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterHostIscsiInterface) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterHostIscsiInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin Configured State of Host ISCSI Interface.
		AdminState *string `json:"AdminState,omitempty"`
		// The Endpoint Config Dn of the Host ISCSI Interface.
		EpDn *string `json:"EpDn,omitempty"`
		// Identifier of the Host ISCSI Interface.
		HostIscsiInterfaceId *int64 `json:"HostIscsiInterfaceId,omitempty"`
		// The visibility of the Host to the endpoint.
		HostVisible *string `json:"HostVisible,omitempty"`
		// MAC address of Host ISCSI Interface.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Name of the Host ISCSI Interface.
		Name *string `json:"Name,omitempty"`
		// Operational State of Host ISCSI Interface.
		OperState *string `json:"OperState,omitempty"`
		// Operability status of Host ISCSI Interface.
		Operability *string `json:"Operability,omitempty"`
		// PeerPort Dn of Host ISCSI Interface.
		PeerDn              *string                              `json:"PeerDn,omitempty"`
		AdapterUnit         *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varAdapterHostIscsiInterfaceWithoutEmbeddedStruct := AdapterHostIscsiInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterHostIscsiInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varAdapterHostIscsiInterface := _AdapterHostIscsiInterface{}
		varAdapterHostIscsiInterface.ClassId = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.ClassId
		varAdapterHostIscsiInterface.ObjectType = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.ObjectType
		varAdapterHostIscsiInterface.AdminState = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.AdminState
		varAdapterHostIscsiInterface.EpDn = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.EpDn
		varAdapterHostIscsiInterface.HostIscsiInterfaceId = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.HostIscsiInterfaceId
		varAdapterHostIscsiInterface.HostVisible = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.HostVisible
		varAdapterHostIscsiInterface.MacAddress = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.MacAddress
		varAdapterHostIscsiInterface.Name = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.Name
		varAdapterHostIscsiInterface.OperState = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.OperState
		varAdapterHostIscsiInterface.Operability = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.Operability
		varAdapterHostIscsiInterface.PeerDn = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.PeerDn
		varAdapterHostIscsiInterface.AdapterUnit = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.AdapterUnit
		varAdapterHostIscsiInterface.InventoryDeviceInfo = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.InventoryDeviceInfo
		varAdapterHostIscsiInterface.RegisteredDevice = varAdapterHostIscsiInterfaceWithoutEmbeddedStruct.RegisteredDevice
		*o = AdapterHostIscsiInterface(varAdapterHostIscsiInterface)
	} else {
		return err
	}

	varAdapterHostIscsiInterface := _AdapterHostIscsiInterface{}

	err = json.Unmarshal(bytes, &varAdapterHostIscsiInterface)
	if err == nil {
		o.EquipmentBase = varAdapterHostIscsiInterface.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "EpDn")
		delete(additionalProperties, "HostIscsiInterfaceId")
		delete(additionalProperties, "HostVisible")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
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

type NullableAdapterHostIscsiInterface struct {
	value *AdapterHostIscsiInterface
	isSet bool
}

func (v NullableAdapterHostIscsiInterface) Get() *AdapterHostIscsiInterface {
	return v.value
}

func (v *NullableAdapterHostIscsiInterface) Set(val *AdapterHostIscsiInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterHostIscsiInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterHostIscsiInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterHostIscsiInterface(val *AdapterHostIscsiInterface) *NullableAdapterHostIscsiInterface {
	return &NullableAdapterHostIscsiInterface{value: val, isSet: true}
}

func (v NullableAdapterHostIscsiInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterHostIscsiInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
