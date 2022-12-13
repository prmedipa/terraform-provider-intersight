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
)

// NetworkInterfaceListAllOf Definition of the list of properties defined in 'network.InterfaceList', excluding properties defined in parent classes.
type NetworkInterfaceListAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of the interface list.
	AdminState *string `json:"AdminState,omitempty"`
	// Allowed VLANs of the interface list.
	AllowedVlans *string `json:"AllowedVlans,omitempty"`
	// Description of the interface list.
	Description *string `json:"Description,omitempty"`
	// Display name of the interface list.
	DisplayName *string `json:"DisplayName,omitempty"`
	// IP address of the interface list.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IP subnet of the interface list.
	IpSubnet *int64 `json:"IpSubnet,omitempty"`
	// MAC address of the interface list.
	Mac *string `json:"Mac,omitempty"`
	// Maximum transmission unit of the interface list.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name of the interface list.
	Name *string `json:"Name,omitempty"`
	// Operational state of the interface list.
	OperState *string `json:"OperState,omitempty"`
	// Port channel id for port channel created on FI switch.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// Interface types supported in Network device like Subinterfaces, Breakout Interfaces.
	PortSubType *string `json:"PortSubType,omitempty"`
	// Port type of interface list.
	PortType *string `json:"PortType,omitempty"`
	// Slot id of the interface list.
	SlotId *string `json:"SlotId,omitempty"`
	// Port speed of the interface list.
	Speed *string `json:"Speed,omitempty"`
	// Speed Group of the interface list.
	SpeedGroup *string `json:"SpeedGroup,omitempty"`
	// VLAN of the interface list.
	Vlan                 *string                              `json:"Vlan,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkInterfaceListAllOf NetworkInterfaceListAllOf

// NewNetworkInterfaceListAllOf instantiates a new NetworkInterfaceListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaceListAllOf(classId string, objectType string) *NetworkInterfaceListAllOf {
	this := NetworkInterfaceListAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkInterfaceListAllOfWithDefaults instantiates a new NetworkInterfaceListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceListAllOfWithDefaults() *NetworkInterfaceListAllOf {
	this := NetworkInterfaceListAllOf{}
	var classId string = "network.InterfaceList"
	this.ClassId = classId
	var objectType string = "network.InterfaceList"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkInterfaceListAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkInterfaceListAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkInterfaceListAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkInterfaceListAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *NetworkInterfaceListAllOf) SetAdminState(v string) {
	o.AdminState = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *NetworkInterfaceListAllOf) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkInterfaceListAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkInterfaceListAllOf) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkInterfaceListAllOf) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpSubnet returns the IpSubnet field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetIpSubnet() int64 {
	if o == nil || o.IpSubnet == nil {
		var ret int64
		return ret
	}
	return *o.IpSubnet
}

// GetIpSubnetOk returns a tuple with the IpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetIpSubnetOk() (*int64, bool) {
	if o == nil || o.IpSubnet == nil {
		return nil, false
	}
	return o.IpSubnet, true
}

// HasIpSubnet returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasIpSubnet() bool {
	if o != nil && o.IpSubnet != nil {
		return true
	}

	return false
}

// SetIpSubnet gets a reference to the given int64 and assigns it to the IpSubnet field.
func (o *NetworkInterfaceListAllOf) SetIpSubnet(v int64) {
	o.IpSubnet = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetMac() string {
	if o == nil || o.Mac == nil {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetMacOk() (*string, bool) {
	if o == nil || o.Mac == nil {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasMac() bool {
	if o != nil && o.Mac != nil {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworkInterfaceListAllOf) SetMac(v string) {
	o.Mac = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *NetworkInterfaceListAllOf) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkInterfaceListAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *NetworkInterfaceListAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *NetworkInterfaceListAllOf) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetPortSubType returns the PortSubType field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetPortSubType() string {
	if o == nil || o.PortSubType == nil {
		var ret string
		return ret
	}
	return *o.PortSubType
}

// GetPortSubTypeOk returns a tuple with the PortSubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetPortSubTypeOk() (*string, bool) {
	if o == nil || o.PortSubType == nil {
		return nil, false
	}
	return o.PortSubType, true
}

// HasPortSubType returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasPortSubType() bool {
	if o != nil && o.PortSubType != nil {
		return true
	}

	return false
}

// SetPortSubType gets a reference to the given string and assigns it to the PortSubType field.
func (o *NetworkInterfaceListAllOf) SetPortSubType(v string) {
	o.PortSubType = &v
}

// GetPortType returns the PortType field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetPortType() string {
	if o == nil || o.PortType == nil {
		var ret string
		return ret
	}
	return *o.PortType
}

// GetPortTypeOk returns a tuple with the PortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetPortTypeOk() (*string, bool) {
	if o == nil || o.PortType == nil {
		return nil, false
	}
	return o.PortType, true
}

// HasPortType returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasPortType() bool {
	if o != nil && o.PortType != nil {
		return true
	}

	return false
}

// SetPortType gets a reference to the given string and assigns it to the PortType field.
func (o *NetworkInterfaceListAllOf) SetPortType(v string) {
	o.PortType = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetSlotId() string {
	if o == nil || o.SlotId == nil {
		var ret string
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetSlotIdOk() (*string, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given string and assigns it to the SlotId field.
func (o *NetworkInterfaceListAllOf) SetSlotId(v string) {
	o.SlotId = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *NetworkInterfaceListAllOf) SetSpeed(v string) {
	o.Speed = &v
}

// GetSpeedGroup returns the SpeedGroup field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetSpeedGroup() string {
	if o == nil || o.SpeedGroup == nil {
		var ret string
		return ret
	}
	return *o.SpeedGroup
}

// GetSpeedGroupOk returns a tuple with the SpeedGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetSpeedGroupOk() (*string, bool) {
	if o == nil || o.SpeedGroup == nil {
		return nil, false
	}
	return o.SpeedGroup, true
}

// HasSpeedGroup returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasSpeedGroup() bool {
	if o != nil && o.SpeedGroup != nil {
		return true
	}

	return false
}

// SetSpeedGroup gets a reference to the given string and assigns it to the SpeedGroup field.
func (o *NetworkInterfaceListAllOf) SetSpeedGroup(v string) {
	o.SpeedGroup = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetVlan() string {
	if o == nil || o.Vlan == nil {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetVlanOk() (*string, bool) {
	if o == nil || o.Vlan == nil {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasVlan() bool {
	if o != nil && o.Vlan != nil {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *NetworkInterfaceListAllOf) SetVlan(v string) {
	o.Vlan = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkInterfaceListAllOf) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkInterfaceListAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceListAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkInterfaceListAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkInterfaceListAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkInterfaceListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.AllowedVlans != nil {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.IpSubnet != nil {
		toSerialize["IpSubnet"] = o.IpSubnet
	}
	if o.Mac != nil {
		toSerialize["Mac"] = o.Mac
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.PortSubType != nil {
		toSerialize["PortSubType"] = o.PortSubType
	}
	if o.PortType != nil {
		toSerialize["PortType"] = o.PortType
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.SpeedGroup != nil {
		toSerialize["SpeedGroup"] = o.SpeedGroup
	}
	if o.Vlan != nil {
		toSerialize["Vlan"] = o.Vlan
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkInterfaceListAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkInterfaceListAllOf := _NetworkInterfaceListAllOf{}

	if err = json.Unmarshal(bytes, &varNetworkInterfaceListAllOf); err == nil {
		*o = NetworkInterfaceListAllOf(varNetworkInterfaceListAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "AllowedVlans")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "DisplayName")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "IpSubnet")
		delete(additionalProperties, "Mac")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "PortSubType")
		delete(additionalProperties, "PortType")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "SpeedGroup")
		delete(additionalProperties, "Vlan")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkInterfaceListAllOf struct {
	value *NetworkInterfaceListAllOf
	isSet bool
}

func (v NullableNetworkInterfaceListAllOf) Get() *NetworkInterfaceListAllOf {
	return v.value
}

func (v *NullableNetworkInterfaceListAllOf) Set(val *NetworkInterfaceListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceListAllOf(val *NetworkInterfaceListAllOf) *NullableNetworkInterfaceListAllOf {
	return &NullableNetworkInterfaceListAllOf{value: val, isSet: true}
}

func (v NullableNetworkInterfaceListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
