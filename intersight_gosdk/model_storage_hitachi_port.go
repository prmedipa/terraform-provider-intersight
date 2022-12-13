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

// StorageHitachiPort Port entity in Hitachi storage array.
type StorageHitachiPort struct {
	StorageBasePhysicalPort
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Fabric mode of the port. true, Set. false, Not set.
	FabricMode *bool `json:"FabricMode,omitempty"`
	// IPv4 address of Hitachi Port.
	Ipv4Address *string `json:"Ipv4Address,omitempty"`
	// IPv6 global address value.
	Ipv6GlobalAddress *string `json:"Ipv6GlobalAddress,omitempty"`
	// IPv6 link local address value.
	Ipv6LinkLocalAddress *string `json:"Ipv6LinkLocalAddress,omitempty"`
	// Determines if IPv6 mode is enabled on the port.
	IsIpv6Enable *bool `json:"IsIpv6Enable,omitempty"`
	// The value set for the port loop ID (AL_PA).
	LoopId *string `json:"LoopId,omitempty"`
	// Topology setting for the port.
	PortConnection *string `json:"PortConnection,omitempty"`
	// LUN security setting for the port.
	PortLunSecurity *bool `json:"PortLunSecurity,omitempty"`
	// Operation mode of the port. Possible values are FC-NVMe, FCP-SCSI, and NOT SUPPORTED.
	PortMode *string `json:"PortMode,omitempty"`
	// Port ID (short) of the port.
	ShortportId *string `json:"ShortportId,omitempty"`
	// Value of MTU for iSCSI communication.
	TcpMtu               *int64                               `json:"TcpMtu,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiPort StorageHitachiPort

// NewStorageHitachiPort instantiates a new StorageHitachiPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiPort(classId string, objectType string) *StorageHitachiPort {
	this := StorageHitachiPort{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiPortWithDefaults instantiates a new StorageHitachiPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiPortWithDefaults() *StorageHitachiPort {
	this := StorageHitachiPort{}
	var classId string = "storage.HitachiPort"
	this.ClassId = classId
	var objectType string = "storage.HitachiPort"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiPort) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiPort) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiPort) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiPort) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFabricMode returns the FabricMode field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetFabricMode() bool {
	if o == nil || o.FabricMode == nil {
		var ret bool
		return ret
	}
	return *o.FabricMode
}

// GetFabricModeOk returns a tuple with the FabricMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetFabricModeOk() (*bool, bool) {
	if o == nil || o.FabricMode == nil {
		return nil, false
	}
	return o.FabricMode, true
}

// HasFabricMode returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasFabricMode() bool {
	if o != nil && o.FabricMode != nil {
		return true
	}

	return false
}

// SetFabricMode gets a reference to the given bool and assigns it to the FabricMode field.
func (o *StorageHitachiPort) SetFabricMode(v bool) {
	o.FabricMode = &v
}

// GetIpv4Address returns the Ipv4Address field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetIpv4Address() string {
	if o == nil || o.Ipv4Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Address
}

// GetIpv4AddressOk returns a tuple with the Ipv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetIpv4AddressOk() (*string, bool) {
	if o == nil || o.Ipv4Address == nil {
		return nil, false
	}
	return o.Ipv4Address, true
}

// HasIpv4Address returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasIpv4Address() bool {
	if o != nil && o.Ipv4Address != nil {
		return true
	}

	return false
}

// SetIpv4Address gets a reference to the given string and assigns it to the Ipv4Address field.
func (o *StorageHitachiPort) SetIpv4Address(v string) {
	o.Ipv4Address = &v
}

// GetIpv6GlobalAddress returns the Ipv6GlobalAddress field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetIpv6GlobalAddress() string {
	if o == nil || o.Ipv6GlobalAddress == nil {
		var ret string
		return ret
	}
	return *o.Ipv6GlobalAddress
}

// GetIpv6GlobalAddressOk returns a tuple with the Ipv6GlobalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetIpv6GlobalAddressOk() (*string, bool) {
	if o == nil || o.Ipv6GlobalAddress == nil {
		return nil, false
	}
	return o.Ipv6GlobalAddress, true
}

// HasIpv6GlobalAddress returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasIpv6GlobalAddress() bool {
	if o != nil && o.Ipv6GlobalAddress != nil {
		return true
	}

	return false
}

// SetIpv6GlobalAddress gets a reference to the given string and assigns it to the Ipv6GlobalAddress field.
func (o *StorageHitachiPort) SetIpv6GlobalAddress(v string) {
	o.Ipv6GlobalAddress = &v
}

// GetIpv6LinkLocalAddress returns the Ipv6LinkLocalAddress field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetIpv6LinkLocalAddress() string {
	if o == nil || o.Ipv6LinkLocalAddress == nil {
		var ret string
		return ret
	}
	return *o.Ipv6LinkLocalAddress
}

// GetIpv6LinkLocalAddressOk returns a tuple with the Ipv6LinkLocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetIpv6LinkLocalAddressOk() (*string, bool) {
	if o == nil || o.Ipv6LinkLocalAddress == nil {
		return nil, false
	}
	return o.Ipv6LinkLocalAddress, true
}

// HasIpv6LinkLocalAddress returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasIpv6LinkLocalAddress() bool {
	if o != nil && o.Ipv6LinkLocalAddress != nil {
		return true
	}

	return false
}

// SetIpv6LinkLocalAddress gets a reference to the given string and assigns it to the Ipv6LinkLocalAddress field.
func (o *StorageHitachiPort) SetIpv6LinkLocalAddress(v string) {
	o.Ipv6LinkLocalAddress = &v
}

// GetIsIpv6Enable returns the IsIpv6Enable field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetIsIpv6Enable() bool {
	if o == nil || o.IsIpv6Enable == nil {
		var ret bool
		return ret
	}
	return *o.IsIpv6Enable
}

// GetIsIpv6EnableOk returns a tuple with the IsIpv6Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetIsIpv6EnableOk() (*bool, bool) {
	if o == nil || o.IsIpv6Enable == nil {
		return nil, false
	}
	return o.IsIpv6Enable, true
}

// HasIsIpv6Enable returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasIsIpv6Enable() bool {
	if o != nil && o.IsIpv6Enable != nil {
		return true
	}

	return false
}

// SetIsIpv6Enable gets a reference to the given bool and assigns it to the IsIpv6Enable field.
func (o *StorageHitachiPort) SetIsIpv6Enable(v bool) {
	o.IsIpv6Enable = &v
}

// GetLoopId returns the LoopId field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetLoopId() string {
	if o == nil || o.LoopId == nil {
		var ret string
		return ret
	}
	return *o.LoopId
}

// GetLoopIdOk returns a tuple with the LoopId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetLoopIdOk() (*string, bool) {
	if o == nil || o.LoopId == nil {
		return nil, false
	}
	return o.LoopId, true
}

// HasLoopId returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasLoopId() bool {
	if o != nil && o.LoopId != nil {
		return true
	}

	return false
}

// SetLoopId gets a reference to the given string and assigns it to the LoopId field.
func (o *StorageHitachiPort) SetLoopId(v string) {
	o.LoopId = &v
}

// GetPortConnection returns the PortConnection field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetPortConnection() string {
	if o == nil || o.PortConnection == nil {
		var ret string
		return ret
	}
	return *o.PortConnection
}

// GetPortConnectionOk returns a tuple with the PortConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetPortConnectionOk() (*string, bool) {
	if o == nil || o.PortConnection == nil {
		return nil, false
	}
	return o.PortConnection, true
}

// HasPortConnection returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasPortConnection() bool {
	if o != nil && o.PortConnection != nil {
		return true
	}

	return false
}

// SetPortConnection gets a reference to the given string and assigns it to the PortConnection field.
func (o *StorageHitachiPort) SetPortConnection(v string) {
	o.PortConnection = &v
}

// GetPortLunSecurity returns the PortLunSecurity field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetPortLunSecurity() bool {
	if o == nil || o.PortLunSecurity == nil {
		var ret bool
		return ret
	}
	return *o.PortLunSecurity
}

// GetPortLunSecurityOk returns a tuple with the PortLunSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetPortLunSecurityOk() (*bool, bool) {
	if o == nil || o.PortLunSecurity == nil {
		return nil, false
	}
	return o.PortLunSecurity, true
}

// HasPortLunSecurity returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasPortLunSecurity() bool {
	if o != nil && o.PortLunSecurity != nil {
		return true
	}

	return false
}

// SetPortLunSecurity gets a reference to the given bool and assigns it to the PortLunSecurity field.
func (o *StorageHitachiPort) SetPortLunSecurity(v bool) {
	o.PortLunSecurity = &v
}

// GetPortMode returns the PortMode field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetPortMode() string {
	if o == nil || o.PortMode == nil {
		var ret string
		return ret
	}
	return *o.PortMode
}

// GetPortModeOk returns a tuple with the PortMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetPortModeOk() (*string, bool) {
	if o == nil || o.PortMode == nil {
		return nil, false
	}
	return o.PortMode, true
}

// HasPortMode returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasPortMode() bool {
	if o != nil && o.PortMode != nil {
		return true
	}

	return false
}

// SetPortMode gets a reference to the given string and assigns it to the PortMode field.
func (o *StorageHitachiPort) SetPortMode(v string) {
	o.PortMode = &v
}

// GetShortportId returns the ShortportId field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetShortportId() string {
	if o == nil || o.ShortportId == nil {
		var ret string
		return ret
	}
	return *o.ShortportId
}

// GetShortportIdOk returns a tuple with the ShortportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetShortportIdOk() (*string, bool) {
	if o == nil || o.ShortportId == nil {
		return nil, false
	}
	return o.ShortportId, true
}

// HasShortportId returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasShortportId() bool {
	if o != nil && o.ShortportId != nil {
		return true
	}

	return false
}

// SetShortportId gets a reference to the given string and assigns it to the ShortportId field.
func (o *StorageHitachiPort) SetShortportId(v string) {
	o.ShortportId = &v
}

// GetTcpMtu returns the TcpMtu field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetTcpMtu() int64 {
	if o == nil || o.TcpMtu == nil {
		var ret int64
		return ret
	}
	return *o.TcpMtu
}

// GetTcpMtuOk returns a tuple with the TcpMtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetTcpMtuOk() (*int64, bool) {
	if o == nil || o.TcpMtu == nil {
		return nil, false
	}
	return o.TcpMtu, true
}

// HasTcpMtu returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasTcpMtu() bool {
	if o != nil && o.TcpMtu != nil {
		return true
	}

	return false
}

// SetTcpMtu gets a reference to the given int64 and assigns it to the TcpMtu field.
func (o *StorageHitachiPort) SetTcpMtu(v int64) {
	o.TcpMtu = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiPort) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiPort) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiPort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBasePhysicalPort, errStorageBasePhysicalPort := json.Marshal(o.StorageBasePhysicalPort)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	errStorageBasePhysicalPort = json.Unmarshal([]byte(serializedStorageBasePhysicalPort), &toSerialize)
	if errStorageBasePhysicalPort != nil {
		return []byte{}, errStorageBasePhysicalPort
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FabricMode != nil {
		toSerialize["FabricMode"] = o.FabricMode
	}
	if o.Ipv4Address != nil {
		toSerialize["Ipv4Address"] = o.Ipv4Address
	}
	if o.Ipv6GlobalAddress != nil {
		toSerialize["Ipv6GlobalAddress"] = o.Ipv6GlobalAddress
	}
	if o.Ipv6LinkLocalAddress != nil {
		toSerialize["Ipv6LinkLocalAddress"] = o.Ipv6LinkLocalAddress
	}
	if o.IsIpv6Enable != nil {
		toSerialize["IsIpv6Enable"] = o.IsIpv6Enable
	}
	if o.LoopId != nil {
		toSerialize["LoopId"] = o.LoopId
	}
	if o.PortConnection != nil {
		toSerialize["PortConnection"] = o.PortConnection
	}
	if o.PortLunSecurity != nil {
		toSerialize["PortLunSecurity"] = o.PortLunSecurity
	}
	if o.PortMode != nil {
		toSerialize["PortMode"] = o.PortMode
	}
	if o.ShortportId != nil {
		toSerialize["ShortportId"] = o.ShortportId
	}
	if o.TcpMtu != nil {
		toSerialize["TcpMtu"] = o.TcpMtu
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiPort) UnmarshalJSON(bytes []byte) (err error) {
	type StorageHitachiPortWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Fabric mode of the port. true, Set. false, Not set.
		FabricMode *bool `json:"FabricMode,omitempty"`
		// IPv4 address of Hitachi Port.
		Ipv4Address *string `json:"Ipv4Address,omitempty"`
		// IPv6 global address value.
		Ipv6GlobalAddress *string `json:"Ipv6GlobalAddress,omitempty"`
		// IPv6 link local address value.
		Ipv6LinkLocalAddress *string `json:"Ipv6LinkLocalAddress,omitempty"`
		// Determines if IPv6 mode is enabled on the port.
		IsIpv6Enable *bool `json:"IsIpv6Enable,omitempty"`
		// The value set for the port loop ID (AL_PA).
		LoopId *string `json:"LoopId,omitempty"`
		// Topology setting for the port.
		PortConnection *string `json:"PortConnection,omitempty"`
		// LUN security setting for the port.
		PortLunSecurity *bool `json:"PortLunSecurity,omitempty"`
		// Operation mode of the port. Possible values are FC-NVMe, FCP-SCSI, and NOT SUPPORTED.
		PortMode *string `json:"PortMode,omitempty"`
		// Port ID (short) of the port.
		ShortportId *string `json:"ShortportId,omitempty"`
		// Value of MTU for iSCSI communication.
		TcpMtu           *int64                               `json:"TcpMtu,omitempty"`
		Array            *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiPortWithoutEmbeddedStruct := StorageHitachiPortWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageHitachiPortWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiPort := _StorageHitachiPort{}
		varStorageHitachiPort.ClassId = varStorageHitachiPortWithoutEmbeddedStruct.ClassId
		varStorageHitachiPort.ObjectType = varStorageHitachiPortWithoutEmbeddedStruct.ObjectType
		varStorageHitachiPort.FabricMode = varStorageHitachiPortWithoutEmbeddedStruct.FabricMode
		varStorageHitachiPort.Ipv4Address = varStorageHitachiPortWithoutEmbeddedStruct.Ipv4Address
		varStorageHitachiPort.Ipv6GlobalAddress = varStorageHitachiPortWithoutEmbeddedStruct.Ipv6GlobalAddress
		varStorageHitachiPort.Ipv6LinkLocalAddress = varStorageHitachiPortWithoutEmbeddedStruct.Ipv6LinkLocalAddress
		varStorageHitachiPort.IsIpv6Enable = varStorageHitachiPortWithoutEmbeddedStruct.IsIpv6Enable
		varStorageHitachiPort.LoopId = varStorageHitachiPortWithoutEmbeddedStruct.LoopId
		varStorageHitachiPort.PortConnection = varStorageHitachiPortWithoutEmbeddedStruct.PortConnection
		varStorageHitachiPort.PortLunSecurity = varStorageHitachiPortWithoutEmbeddedStruct.PortLunSecurity
		varStorageHitachiPort.PortMode = varStorageHitachiPortWithoutEmbeddedStruct.PortMode
		varStorageHitachiPort.ShortportId = varStorageHitachiPortWithoutEmbeddedStruct.ShortportId
		varStorageHitachiPort.TcpMtu = varStorageHitachiPortWithoutEmbeddedStruct.TcpMtu
		varStorageHitachiPort.Array = varStorageHitachiPortWithoutEmbeddedStruct.Array
		varStorageHitachiPort.RegisteredDevice = varStorageHitachiPortWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiPort(varStorageHitachiPort)
	} else {
		return err
	}

	varStorageHitachiPort := _StorageHitachiPort{}

	err = json.Unmarshal(bytes, &varStorageHitachiPort)
	if err == nil {
		o.StorageBasePhysicalPort = varStorageHitachiPort.StorageBasePhysicalPort
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FabricMode")
		delete(additionalProperties, "Ipv4Address")
		delete(additionalProperties, "Ipv6GlobalAddress")
		delete(additionalProperties, "Ipv6LinkLocalAddress")
		delete(additionalProperties, "IsIpv6Enable")
		delete(additionalProperties, "LoopId")
		delete(additionalProperties, "PortConnection")
		delete(additionalProperties, "PortLunSecurity")
		delete(additionalProperties, "PortMode")
		delete(additionalProperties, "ShortportId")
		delete(additionalProperties, "TcpMtu")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectStorageBasePhysicalPort := reflect.ValueOf(o.StorageBasePhysicalPort)
		for i := 0; i < reflectStorageBasePhysicalPort.Type().NumField(); i++ {
			t := reflectStorageBasePhysicalPort.Type().Field(i)

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

type NullableStorageHitachiPort struct {
	value *StorageHitachiPort
	isSet bool
}

func (v NullableStorageHitachiPort) Get() *StorageHitachiPort {
	return v.value
}

func (v *NullableStorageHitachiPort) Set(val *StorageHitachiPort) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiPort) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiPort(val *StorageHitachiPort) *NullableStorageHitachiPort {
	return &NullableStorageHitachiPort{value: val, isSet: true}
}

func (v NullableStorageHitachiPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
