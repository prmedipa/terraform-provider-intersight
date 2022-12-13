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

// EtherPortChannel Model contains the details of the ethernet port-channels configured on the FI.
type EtherPortChannel struct {
	InventoryInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access VLANs for this port-channel, on this FI.
	AccessVlan *string `json:"AccessVlan,omitempty"`
	// Administratively configured state (enabled/disabled) for this port-channel.
	AdminState *string `json:"AdminState,omitempty"`
	// Allowed VLANs on this port-channel, on this FI.
	AllowedVlans *string `json:"AllowedVlans,omitempty"`
	// Bandwidth of this port-channel.
	BandWidth *string `json:"BandWidth,omitempty"`
	// Description of this port-channel.
	Description *string `json:"Description,omitempty"`
	// IP address of this port-channel.
	IpAddress *string `json:"IpAddress,omitempty"`
	// IP address mask of this port-channel.
	IpAddressMask *int64 `json:"IpAddressMask,omitempty"`
	// IPv6 subnet in CIDR notation of this port-channel. Ex. 2000::/8.
	Ipv6SubnetCidr *string `json:"Ipv6SubnetCidr,omitempty"`
	// MAC address of this port-channel.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Operating mode of this port-channel.
	Mode *string `json:"Mode,omitempty"`
	// Maximum transmission unit of this port-channel.
	Mtu *int64 `json:"Mtu,omitempty"`
	// Name of the port channel.
	Name *string `json:"Name,omitempty"`
	// Native VLAN for this port-channel, on this FI.
	NativeVlan *string `json:"NativeVlan,omitempty"`
	// Operational speed of this port-channel.
	OperSpeed *string `json:"OperSpeed,omitempty"`
	// Operational state of this port-channel.
	OperState *string `json:"OperState,omitempty"`
	// Reason for this port-channel's Operational state.
	OperStateQual *string `json:"OperStateQual,omitempty"`
	// Unique identifier for this port-channel on the FI.
	PortChannelId *int64 `json:"PortChannelId,omitempty"`
	// This port-channel's configured role (uplink, server, etc.).
	Role *string `json:"Role,omitempty"`
	// Detailed status of this port-channel.
	Status *string `json:"Status,omitempty"`
	// Switch Identifier that is local to a cluster.
	SwitchId             *string                              `json:"SwitchId,omitempty"`
	EquipmentSwitchCard  *EquipmentSwitchCardRelationship     `json:"EquipmentSwitchCard,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EtherPortChannel EtherPortChannel

// NewEtherPortChannel instantiates a new EtherPortChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEtherPortChannel(classId string, objectType string) *EtherPortChannel {
	this := EtherPortChannel{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEtherPortChannelWithDefaults instantiates a new EtherPortChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEtherPortChannelWithDefaults() *EtherPortChannel {
	this := EtherPortChannel{}
	var classId string = "ether.PortChannel"
	this.ClassId = classId
	var objectType string = "ether.PortChannel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EtherPortChannel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EtherPortChannel) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EtherPortChannel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EtherPortChannel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessVlan returns the AccessVlan field value if set, zero value otherwise.
func (o *EtherPortChannel) GetAccessVlan() string {
	if o == nil || o.AccessVlan == nil {
		var ret string
		return ret
	}
	return *o.AccessVlan
}

// GetAccessVlanOk returns a tuple with the AccessVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetAccessVlanOk() (*string, bool) {
	if o == nil || o.AccessVlan == nil {
		return nil, false
	}
	return o.AccessVlan, true
}

// HasAccessVlan returns a boolean if a field has been set.
func (o *EtherPortChannel) HasAccessVlan() bool {
	if o != nil && o.AccessVlan != nil {
		return true
	}

	return false
}

// SetAccessVlan gets a reference to the given string and assigns it to the AccessVlan field.
func (o *EtherPortChannel) SetAccessVlan(v string) {
	o.AccessVlan = &v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *EtherPortChannel) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *EtherPortChannel) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *EtherPortChannel) SetAdminState(v string) {
	o.AdminState = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *EtherPortChannel) GetAllowedVlans() string {
	if o == nil || o.AllowedVlans == nil {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetAllowedVlansOk() (*string, bool) {
	if o == nil || o.AllowedVlans == nil {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *EtherPortChannel) HasAllowedVlans() bool {
	if o != nil && o.AllowedVlans != nil {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *EtherPortChannel) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetBandWidth returns the BandWidth field value if set, zero value otherwise.
func (o *EtherPortChannel) GetBandWidth() string {
	if o == nil || o.BandWidth == nil {
		var ret string
		return ret
	}
	return *o.BandWidth
}

// GetBandWidthOk returns a tuple with the BandWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetBandWidthOk() (*string, bool) {
	if o == nil || o.BandWidth == nil {
		return nil, false
	}
	return o.BandWidth, true
}

// HasBandWidth returns a boolean if a field has been set.
func (o *EtherPortChannel) HasBandWidth() bool {
	if o != nil && o.BandWidth != nil {
		return true
	}

	return false
}

// SetBandWidth gets a reference to the given string and assigns it to the BandWidth field.
func (o *EtherPortChannel) SetBandWidth(v string) {
	o.BandWidth = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EtherPortChannel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EtherPortChannel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EtherPortChannel) SetDescription(v string) {
	o.Description = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *EtherPortChannel) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *EtherPortChannel) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *EtherPortChannel) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpAddressMask returns the IpAddressMask field value if set, zero value otherwise.
func (o *EtherPortChannel) GetIpAddressMask() int64 {
	if o == nil || o.IpAddressMask == nil {
		var ret int64
		return ret
	}
	return *o.IpAddressMask
}

// GetIpAddressMaskOk returns a tuple with the IpAddressMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetIpAddressMaskOk() (*int64, bool) {
	if o == nil || o.IpAddressMask == nil {
		return nil, false
	}
	return o.IpAddressMask, true
}

// HasIpAddressMask returns a boolean if a field has been set.
func (o *EtherPortChannel) HasIpAddressMask() bool {
	if o != nil && o.IpAddressMask != nil {
		return true
	}

	return false
}

// SetIpAddressMask gets a reference to the given int64 and assigns it to the IpAddressMask field.
func (o *EtherPortChannel) SetIpAddressMask(v int64) {
	o.IpAddressMask = &v
}

// GetIpv6SubnetCidr returns the Ipv6SubnetCidr field value if set, zero value otherwise.
func (o *EtherPortChannel) GetIpv6SubnetCidr() string {
	if o == nil || o.Ipv6SubnetCidr == nil {
		var ret string
		return ret
	}
	return *o.Ipv6SubnetCidr
}

// GetIpv6SubnetCidrOk returns a tuple with the Ipv6SubnetCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetIpv6SubnetCidrOk() (*string, bool) {
	if o == nil || o.Ipv6SubnetCidr == nil {
		return nil, false
	}
	return o.Ipv6SubnetCidr, true
}

// HasIpv6SubnetCidr returns a boolean if a field has been set.
func (o *EtherPortChannel) HasIpv6SubnetCidr() bool {
	if o != nil && o.Ipv6SubnetCidr != nil {
		return true
	}

	return false
}

// SetIpv6SubnetCidr gets a reference to the given string and assigns it to the Ipv6SubnetCidr field.
func (o *EtherPortChannel) SetIpv6SubnetCidr(v string) {
	o.Ipv6SubnetCidr = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *EtherPortChannel) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *EtherPortChannel) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *EtherPortChannel) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *EtherPortChannel) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *EtherPortChannel) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *EtherPortChannel) SetMode(v string) {
	o.Mode = &v
}

// GetMtu returns the Mtu field value if set, zero value otherwise.
func (o *EtherPortChannel) GetMtu() int64 {
	if o == nil || o.Mtu == nil {
		var ret int64
		return ret
	}
	return *o.Mtu
}

// GetMtuOk returns a tuple with the Mtu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetMtuOk() (*int64, bool) {
	if o == nil || o.Mtu == nil {
		return nil, false
	}
	return o.Mtu, true
}

// HasMtu returns a boolean if a field has been set.
func (o *EtherPortChannel) HasMtu() bool {
	if o != nil && o.Mtu != nil {
		return true
	}

	return false
}

// SetMtu gets a reference to the given int64 and assigns it to the Mtu field.
func (o *EtherPortChannel) SetMtu(v int64) {
	o.Mtu = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EtherPortChannel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EtherPortChannel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EtherPortChannel) SetName(v string) {
	o.Name = &v
}

// GetNativeVlan returns the NativeVlan field value if set, zero value otherwise.
func (o *EtherPortChannel) GetNativeVlan() string {
	if o == nil || o.NativeVlan == nil {
		var ret string
		return ret
	}
	return *o.NativeVlan
}

// GetNativeVlanOk returns a tuple with the NativeVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetNativeVlanOk() (*string, bool) {
	if o == nil || o.NativeVlan == nil {
		return nil, false
	}
	return o.NativeVlan, true
}

// HasNativeVlan returns a boolean if a field has been set.
func (o *EtherPortChannel) HasNativeVlan() bool {
	if o != nil && o.NativeVlan != nil {
		return true
	}

	return false
}

// SetNativeVlan gets a reference to the given string and assigns it to the NativeVlan field.
func (o *EtherPortChannel) SetNativeVlan(v string) {
	o.NativeVlan = &v
}

// GetOperSpeed returns the OperSpeed field value if set, zero value otherwise.
func (o *EtherPortChannel) GetOperSpeed() string {
	if o == nil || o.OperSpeed == nil {
		var ret string
		return ret
	}
	return *o.OperSpeed
}

// GetOperSpeedOk returns a tuple with the OperSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetOperSpeedOk() (*string, bool) {
	if o == nil || o.OperSpeed == nil {
		return nil, false
	}
	return o.OperSpeed, true
}

// HasOperSpeed returns a boolean if a field has been set.
func (o *EtherPortChannel) HasOperSpeed() bool {
	if o != nil && o.OperSpeed != nil {
		return true
	}

	return false
}

// SetOperSpeed gets a reference to the given string and assigns it to the OperSpeed field.
func (o *EtherPortChannel) SetOperSpeed(v string) {
	o.OperSpeed = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EtherPortChannel) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EtherPortChannel) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EtherPortChannel) SetOperState(v string) {
	o.OperState = &v
}

// GetOperStateQual returns the OperStateQual field value if set, zero value otherwise.
func (o *EtherPortChannel) GetOperStateQual() string {
	if o == nil || o.OperStateQual == nil {
		var ret string
		return ret
	}
	return *o.OperStateQual
}

// GetOperStateQualOk returns a tuple with the OperStateQual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetOperStateQualOk() (*string, bool) {
	if o == nil || o.OperStateQual == nil {
		return nil, false
	}
	return o.OperStateQual, true
}

// HasOperStateQual returns a boolean if a field has been set.
func (o *EtherPortChannel) HasOperStateQual() bool {
	if o != nil && o.OperStateQual != nil {
		return true
	}

	return false
}

// SetOperStateQual gets a reference to the given string and assigns it to the OperStateQual field.
func (o *EtherPortChannel) SetOperStateQual(v string) {
	o.OperStateQual = &v
}

// GetPortChannelId returns the PortChannelId field value if set, zero value otherwise.
func (o *EtherPortChannel) GetPortChannelId() int64 {
	if o == nil || o.PortChannelId == nil {
		var ret int64
		return ret
	}
	return *o.PortChannelId
}

// GetPortChannelIdOk returns a tuple with the PortChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetPortChannelIdOk() (*int64, bool) {
	if o == nil || o.PortChannelId == nil {
		return nil, false
	}
	return o.PortChannelId, true
}

// HasPortChannelId returns a boolean if a field has been set.
func (o *EtherPortChannel) HasPortChannelId() bool {
	if o != nil && o.PortChannelId != nil {
		return true
	}

	return false
}

// SetPortChannelId gets a reference to the given int64 and assigns it to the PortChannelId field.
func (o *EtherPortChannel) SetPortChannelId(v int64) {
	o.PortChannelId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *EtherPortChannel) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *EtherPortChannel) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *EtherPortChannel) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EtherPortChannel) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EtherPortChannel) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EtherPortChannel) SetStatus(v string) {
	o.Status = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EtherPortChannel) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EtherPortChannel) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *EtherPortChannel) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetEquipmentSwitchCard returns the EquipmentSwitchCard field value if set, zero value otherwise.
func (o *EtherPortChannel) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship {
	if o == nil || o.EquipmentSwitchCard == nil {
		var ret EquipmentSwitchCardRelationship
		return ret
	}
	return *o.EquipmentSwitchCard
}

// GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool) {
	if o == nil || o.EquipmentSwitchCard == nil {
		return nil, false
	}
	return o.EquipmentSwitchCard, true
}

// HasEquipmentSwitchCard returns a boolean if a field has been set.
func (o *EtherPortChannel) HasEquipmentSwitchCard() bool {
	if o != nil && o.EquipmentSwitchCard != nil {
		return true
	}

	return false
}

// SetEquipmentSwitchCard gets a reference to the given EquipmentSwitchCardRelationship and assigns it to the EquipmentSwitchCard field.
func (o *EtherPortChannel) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship) {
	o.EquipmentSwitchCard = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EtherPortChannel) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EtherPortChannel) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EtherPortChannel) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EtherPortChannel) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EtherPortChannel) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EtherPortChannel) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EtherPortChannel) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EtherPortChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryInterface, errInventoryInterface := json.Marshal(o.InventoryInterface)
	if errInventoryInterface != nil {
		return []byte{}, errInventoryInterface
	}
	errInventoryInterface = json.Unmarshal([]byte(serializedInventoryInterface), &toSerialize)
	if errInventoryInterface != nil {
		return []byte{}, errInventoryInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccessVlan != nil {
		toSerialize["AccessVlan"] = o.AccessVlan
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.AllowedVlans != nil {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if o.BandWidth != nil {
		toSerialize["BandWidth"] = o.BandWidth
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.IpAddressMask != nil {
		toSerialize["IpAddressMask"] = o.IpAddressMask
	}
	if o.Ipv6SubnetCidr != nil {
		toSerialize["Ipv6SubnetCidr"] = o.Ipv6SubnetCidr
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Mtu != nil {
		toSerialize["Mtu"] = o.Mtu
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NativeVlan != nil {
		toSerialize["NativeVlan"] = o.NativeVlan
	}
	if o.OperSpeed != nil {
		toSerialize["OperSpeed"] = o.OperSpeed
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.OperStateQual != nil {
		toSerialize["OperStateQual"] = o.OperStateQual
	}
	if o.PortChannelId != nil {
		toSerialize["PortChannelId"] = o.PortChannelId
	}
	if o.Role != nil {
		toSerialize["Role"] = o.Role
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.EquipmentSwitchCard != nil {
		toSerialize["EquipmentSwitchCard"] = o.EquipmentSwitchCard
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

func (o *EtherPortChannel) UnmarshalJSON(bytes []byte) (err error) {
	type EtherPortChannelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access VLANs for this port-channel, on this FI.
		AccessVlan *string `json:"AccessVlan,omitempty"`
		// Administratively configured state (enabled/disabled) for this port-channel.
		AdminState *string `json:"AdminState,omitempty"`
		// Allowed VLANs on this port-channel, on this FI.
		AllowedVlans *string `json:"AllowedVlans,omitempty"`
		// Bandwidth of this port-channel.
		BandWidth *string `json:"BandWidth,omitempty"`
		// Description of this port-channel.
		Description *string `json:"Description,omitempty"`
		// IP address of this port-channel.
		IpAddress *string `json:"IpAddress,omitempty"`
		// IP address mask of this port-channel.
		IpAddressMask *int64 `json:"IpAddressMask,omitempty"`
		// IPv6 subnet in CIDR notation of this port-channel. Ex. 2000::/8.
		Ipv6SubnetCidr *string `json:"Ipv6SubnetCidr,omitempty"`
		// MAC address of this port-channel.
		MacAddress *string `json:"MacAddress,omitempty"`
		// Operating mode of this port-channel.
		Mode *string `json:"Mode,omitempty"`
		// Maximum transmission unit of this port-channel.
		Mtu *int64 `json:"Mtu,omitempty"`
		// Name of the port channel.
		Name *string `json:"Name,omitempty"`
		// Native VLAN for this port-channel, on this FI.
		NativeVlan *string `json:"NativeVlan,omitempty"`
		// Operational speed of this port-channel.
		OperSpeed *string `json:"OperSpeed,omitempty"`
		// Operational state of this port-channel.
		OperState *string `json:"OperState,omitempty"`
		// Reason for this port-channel's Operational state.
		OperStateQual *string `json:"OperStateQual,omitempty"`
		// Unique identifier for this port-channel on the FI.
		PortChannelId *int64 `json:"PortChannelId,omitempty"`
		// This port-channel's configured role (uplink, server, etc.).
		Role *string `json:"Role,omitempty"`
		// Detailed status of this port-channel.
		Status *string `json:"Status,omitempty"`
		// Switch Identifier that is local to a cluster.
		SwitchId            *string                              `json:"SwitchId,omitempty"`
		EquipmentSwitchCard *EquipmentSwitchCardRelationship     `json:"EquipmentSwitchCard,omitempty"`
		NetworkElement      *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEtherPortChannelWithoutEmbeddedStruct := EtherPortChannelWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEtherPortChannelWithoutEmbeddedStruct)
	if err == nil {
		varEtherPortChannel := _EtherPortChannel{}
		varEtherPortChannel.ClassId = varEtherPortChannelWithoutEmbeddedStruct.ClassId
		varEtherPortChannel.ObjectType = varEtherPortChannelWithoutEmbeddedStruct.ObjectType
		varEtherPortChannel.AccessVlan = varEtherPortChannelWithoutEmbeddedStruct.AccessVlan
		varEtherPortChannel.AdminState = varEtherPortChannelWithoutEmbeddedStruct.AdminState
		varEtherPortChannel.AllowedVlans = varEtherPortChannelWithoutEmbeddedStruct.AllowedVlans
		varEtherPortChannel.BandWidth = varEtherPortChannelWithoutEmbeddedStruct.BandWidth
		varEtherPortChannel.Description = varEtherPortChannelWithoutEmbeddedStruct.Description
		varEtherPortChannel.IpAddress = varEtherPortChannelWithoutEmbeddedStruct.IpAddress
		varEtherPortChannel.IpAddressMask = varEtherPortChannelWithoutEmbeddedStruct.IpAddressMask
		varEtherPortChannel.Ipv6SubnetCidr = varEtherPortChannelWithoutEmbeddedStruct.Ipv6SubnetCidr
		varEtherPortChannel.MacAddress = varEtherPortChannelWithoutEmbeddedStruct.MacAddress
		varEtherPortChannel.Mode = varEtherPortChannelWithoutEmbeddedStruct.Mode
		varEtherPortChannel.Mtu = varEtherPortChannelWithoutEmbeddedStruct.Mtu
		varEtherPortChannel.Name = varEtherPortChannelWithoutEmbeddedStruct.Name
		varEtherPortChannel.NativeVlan = varEtherPortChannelWithoutEmbeddedStruct.NativeVlan
		varEtherPortChannel.OperSpeed = varEtherPortChannelWithoutEmbeddedStruct.OperSpeed
		varEtherPortChannel.OperState = varEtherPortChannelWithoutEmbeddedStruct.OperState
		varEtherPortChannel.OperStateQual = varEtherPortChannelWithoutEmbeddedStruct.OperStateQual
		varEtherPortChannel.PortChannelId = varEtherPortChannelWithoutEmbeddedStruct.PortChannelId
		varEtherPortChannel.Role = varEtherPortChannelWithoutEmbeddedStruct.Role
		varEtherPortChannel.Status = varEtherPortChannelWithoutEmbeddedStruct.Status
		varEtherPortChannel.SwitchId = varEtherPortChannelWithoutEmbeddedStruct.SwitchId
		varEtherPortChannel.EquipmentSwitchCard = varEtherPortChannelWithoutEmbeddedStruct.EquipmentSwitchCard
		varEtherPortChannel.NetworkElement = varEtherPortChannelWithoutEmbeddedStruct.NetworkElement
		varEtherPortChannel.RegisteredDevice = varEtherPortChannelWithoutEmbeddedStruct.RegisteredDevice
		*o = EtherPortChannel(varEtherPortChannel)
	} else {
		return err
	}

	varEtherPortChannel := _EtherPortChannel{}

	err = json.Unmarshal(bytes, &varEtherPortChannel)
	if err == nil {
		o.InventoryInterface = varEtherPortChannel.InventoryInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessVlan")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "AllowedVlans")
		delete(additionalProperties, "BandWidth")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "IpAddressMask")
		delete(additionalProperties, "Ipv6SubnetCidr")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Mtu")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NativeVlan")
		delete(additionalProperties, "OperSpeed")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "OperStateQual")
		delete(additionalProperties, "PortChannelId")
		delete(additionalProperties, "Role")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "EquipmentSwitchCard")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryInterface := reflect.ValueOf(o.InventoryInterface)
		for i := 0; i < reflectInventoryInterface.Type().NumField(); i++ {
			t := reflectInventoryInterface.Type().Field(i)

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

type NullableEtherPortChannel struct {
	value *EtherPortChannel
	isSet bool
}

func (v NullableEtherPortChannel) Get() *EtherPortChannel {
	return v.value
}

func (v *NullableEtherPortChannel) Set(val *EtherPortChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableEtherPortChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableEtherPortChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEtherPortChannel(val *EtherPortChannel) *NullableEtherPortChannel {
	return &NullableEtherPortChannel{value: val, isSet: true}
}

func (v NullableEtherPortChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEtherPortChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
