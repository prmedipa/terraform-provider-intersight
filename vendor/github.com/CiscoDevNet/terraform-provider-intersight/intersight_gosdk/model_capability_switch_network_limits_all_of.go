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

// CapabilitySwitchNetworkLimitsAllOf Definition of the list of properties defined in 'capability.SwitchNetworkLimits', excluding properties defined in parent classes.
type CapabilitySwitchNetworkLimitsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect.
	MaxCompressedPortVlanCount *int64 `json:"MaxCompressedPortVlanCount,omitempty"`
	// Maximum configurable VLANs on Switch/Fabric-Interconnect.
	MaxUncompressedPortVlanCount *int64 `json:"MaxUncompressedPortVlanCount,omitempty"`
	// Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect.
	MaximumActiveTrafficMonitoringSessions *int64 `json:"MaximumActiveTrafficMonitoringSessions,omitempty"`
	// Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect.
	MaximumEthernetPortChannels *int64 `json:"MaximumEthernetPortChannels,omitempty"`
	// Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect.
	MaximumEthernetUplinkPorts *int64 `json:"MaximumEthernetUplinkPorts,omitempty"`
	// Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect.
	MaximumFcPortChannelMembers *int64 `json:"MaximumFcPortChannelMembers,omitempty"`
	// Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect.
	MaximumFcPortChannels *int64 `json:"MaximumFcPortChannels,omitempty"`
	// Maximum configurable IGMP Groups on Switch/Fabric-Interconnect.
	MaximumIgmpGroups *int64 `json:"MaximumIgmpGroups,omitempty"`
	// Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect.
	MaximumPortChannelMembers *int64 `json:"MaximumPortChannelMembers,omitempty"`
	// Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect.
	MaximumPrimaryVlan *int64 `json:"MaximumPrimaryVlan,omitempty"`
	// Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect.
	MaximumSecondaryVlan *int64 `json:"MaximumSecondaryVlan,omitempty"`
	// Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect.
	MaximumSecondaryVlanPerPrimary *int64 `json:"MaximumSecondaryVlanPerPrimary,omitempty"`
	// Maximum allowes VIFs on Switch/Fabric-Interconnect.
	MaximumVifs *int64 `json:"MaximumVifs,omitempty"`
	// Maximum configurable VLANs on Switch/Fabric-Interconnect.
	MaximumVlans *int64 `json:"MaximumVlans,omitempty"`
	// Minimum required fans in 'active' state for this Switch/Fabric-Interconnect.
	MinimumActiveFans    *int64 `json:"MinimumActiveFans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilitySwitchNetworkLimitsAllOf CapabilitySwitchNetworkLimitsAllOf

// NewCapabilitySwitchNetworkLimitsAllOf instantiates a new CapabilitySwitchNetworkLimitsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySwitchNetworkLimitsAllOf(classId string, objectType string) *CapabilitySwitchNetworkLimitsAllOf {
	this := CapabilitySwitchNetworkLimitsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilitySwitchNetworkLimitsAllOfWithDefaults instantiates a new CapabilitySwitchNetworkLimitsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySwitchNetworkLimitsAllOfWithDefaults() *CapabilitySwitchNetworkLimitsAllOf {
	this := CapabilitySwitchNetworkLimitsAllOf{}
	var classId string = "capability.SwitchNetworkLimits"
	this.ClassId = classId
	var objectType string = "capability.SwitchNetworkLimits"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilitySwitchNetworkLimitsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilitySwitchNetworkLimitsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilitySwitchNetworkLimitsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilitySwitchNetworkLimitsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMaxCompressedPortVlanCount returns the MaxCompressedPortVlanCount field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxCompressedPortVlanCount() int64 {
	if o == nil || o.MaxCompressedPortVlanCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxCompressedPortVlanCount
}

// GetMaxCompressedPortVlanCountOk returns a tuple with the MaxCompressedPortVlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxCompressedPortVlanCountOk() (*int64, bool) {
	if o == nil || o.MaxCompressedPortVlanCount == nil {
		return nil, false
	}
	return o.MaxCompressedPortVlanCount, true
}

// HasMaxCompressedPortVlanCount returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaxCompressedPortVlanCount() bool {
	if o != nil && o.MaxCompressedPortVlanCount != nil {
		return true
	}

	return false
}

// SetMaxCompressedPortVlanCount gets a reference to the given int64 and assigns it to the MaxCompressedPortVlanCount field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaxCompressedPortVlanCount(v int64) {
	o.MaxCompressedPortVlanCount = &v
}

// GetMaxUncompressedPortVlanCount returns the MaxUncompressedPortVlanCount field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxUncompressedPortVlanCount() int64 {
	if o == nil || o.MaxUncompressedPortVlanCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxUncompressedPortVlanCount
}

// GetMaxUncompressedPortVlanCountOk returns a tuple with the MaxUncompressedPortVlanCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaxUncompressedPortVlanCountOk() (*int64, bool) {
	if o == nil || o.MaxUncompressedPortVlanCount == nil {
		return nil, false
	}
	return o.MaxUncompressedPortVlanCount, true
}

// HasMaxUncompressedPortVlanCount returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaxUncompressedPortVlanCount() bool {
	if o != nil && o.MaxUncompressedPortVlanCount != nil {
		return true
	}

	return false
}

// SetMaxUncompressedPortVlanCount gets a reference to the given int64 and assigns it to the MaxUncompressedPortVlanCount field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaxUncompressedPortVlanCount(v int64) {
	o.MaxUncompressedPortVlanCount = &v
}

// GetMaximumActiveTrafficMonitoringSessions returns the MaximumActiveTrafficMonitoringSessions field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumActiveTrafficMonitoringSessions() int64 {
	if o == nil || o.MaximumActiveTrafficMonitoringSessions == nil {
		var ret int64
		return ret
	}
	return *o.MaximumActiveTrafficMonitoringSessions
}

// GetMaximumActiveTrafficMonitoringSessionsOk returns a tuple with the MaximumActiveTrafficMonitoringSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumActiveTrafficMonitoringSessionsOk() (*int64, bool) {
	if o == nil || o.MaximumActiveTrafficMonitoringSessions == nil {
		return nil, false
	}
	return o.MaximumActiveTrafficMonitoringSessions, true
}

// HasMaximumActiveTrafficMonitoringSessions returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumActiveTrafficMonitoringSessions() bool {
	if o != nil && o.MaximumActiveTrafficMonitoringSessions != nil {
		return true
	}

	return false
}

// SetMaximumActiveTrafficMonitoringSessions gets a reference to the given int64 and assigns it to the MaximumActiveTrafficMonitoringSessions field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumActiveTrafficMonitoringSessions(v int64) {
	o.MaximumActiveTrafficMonitoringSessions = &v
}

// GetMaximumEthernetPortChannels returns the MaximumEthernetPortChannels field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetPortChannels() int64 {
	if o == nil || o.MaximumEthernetPortChannels == nil {
		var ret int64
		return ret
	}
	return *o.MaximumEthernetPortChannels
}

// GetMaximumEthernetPortChannelsOk returns a tuple with the MaximumEthernetPortChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetPortChannelsOk() (*int64, bool) {
	if o == nil || o.MaximumEthernetPortChannels == nil {
		return nil, false
	}
	return o.MaximumEthernetPortChannels, true
}

// HasMaximumEthernetPortChannels returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumEthernetPortChannels() bool {
	if o != nil && o.MaximumEthernetPortChannels != nil {
		return true
	}

	return false
}

// SetMaximumEthernetPortChannels gets a reference to the given int64 and assigns it to the MaximumEthernetPortChannels field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumEthernetPortChannels(v int64) {
	o.MaximumEthernetPortChannels = &v
}

// GetMaximumEthernetUplinkPorts returns the MaximumEthernetUplinkPorts field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetUplinkPorts() int64 {
	if o == nil || o.MaximumEthernetUplinkPorts == nil {
		var ret int64
		return ret
	}
	return *o.MaximumEthernetUplinkPorts
}

// GetMaximumEthernetUplinkPortsOk returns a tuple with the MaximumEthernetUplinkPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumEthernetUplinkPortsOk() (*int64, bool) {
	if o == nil || o.MaximumEthernetUplinkPorts == nil {
		return nil, false
	}
	return o.MaximumEthernetUplinkPorts, true
}

// HasMaximumEthernetUplinkPorts returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumEthernetUplinkPorts() bool {
	if o != nil && o.MaximumEthernetUplinkPorts != nil {
		return true
	}

	return false
}

// SetMaximumEthernetUplinkPorts gets a reference to the given int64 and assigns it to the MaximumEthernetUplinkPorts field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumEthernetUplinkPorts(v int64) {
	o.MaximumEthernetUplinkPorts = &v
}

// GetMaximumFcPortChannelMembers returns the MaximumFcPortChannelMembers field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelMembers() int64 {
	if o == nil || o.MaximumFcPortChannelMembers == nil {
		var ret int64
		return ret
	}
	return *o.MaximumFcPortChannelMembers
}

// GetMaximumFcPortChannelMembersOk returns a tuple with the MaximumFcPortChannelMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelMembersOk() (*int64, bool) {
	if o == nil || o.MaximumFcPortChannelMembers == nil {
		return nil, false
	}
	return o.MaximumFcPortChannelMembers, true
}

// HasMaximumFcPortChannelMembers returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumFcPortChannelMembers() bool {
	if o != nil && o.MaximumFcPortChannelMembers != nil {
		return true
	}

	return false
}

// SetMaximumFcPortChannelMembers gets a reference to the given int64 and assigns it to the MaximumFcPortChannelMembers field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumFcPortChannelMembers(v int64) {
	o.MaximumFcPortChannelMembers = &v
}

// GetMaximumFcPortChannels returns the MaximumFcPortChannels field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannels() int64 {
	if o == nil || o.MaximumFcPortChannels == nil {
		var ret int64
		return ret
	}
	return *o.MaximumFcPortChannels
}

// GetMaximumFcPortChannelsOk returns a tuple with the MaximumFcPortChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumFcPortChannelsOk() (*int64, bool) {
	if o == nil || o.MaximumFcPortChannels == nil {
		return nil, false
	}
	return o.MaximumFcPortChannels, true
}

// HasMaximumFcPortChannels returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumFcPortChannels() bool {
	if o != nil && o.MaximumFcPortChannels != nil {
		return true
	}

	return false
}

// SetMaximumFcPortChannels gets a reference to the given int64 and assigns it to the MaximumFcPortChannels field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumFcPortChannels(v int64) {
	o.MaximumFcPortChannels = &v
}

// GetMaximumIgmpGroups returns the MaximumIgmpGroups field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumIgmpGroups() int64 {
	if o == nil || o.MaximumIgmpGroups == nil {
		var ret int64
		return ret
	}
	return *o.MaximumIgmpGroups
}

// GetMaximumIgmpGroupsOk returns a tuple with the MaximumIgmpGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumIgmpGroupsOk() (*int64, bool) {
	if o == nil || o.MaximumIgmpGroups == nil {
		return nil, false
	}
	return o.MaximumIgmpGroups, true
}

// HasMaximumIgmpGroups returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumIgmpGroups() bool {
	if o != nil && o.MaximumIgmpGroups != nil {
		return true
	}

	return false
}

// SetMaximumIgmpGroups gets a reference to the given int64 and assigns it to the MaximumIgmpGroups field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumIgmpGroups(v int64) {
	o.MaximumIgmpGroups = &v
}

// GetMaximumPortChannelMembers returns the MaximumPortChannelMembers field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPortChannelMembers() int64 {
	if o == nil || o.MaximumPortChannelMembers == nil {
		var ret int64
		return ret
	}
	return *o.MaximumPortChannelMembers
}

// GetMaximumPortChannelMembersOk returns a tuple with the MaximumPortChannelMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPortChannelMembersOk() (*int64, bool) {
	if o == nil || o.MaximumPortChannelMembers == nil {
		return nil, false
	}
	return o.MaximumPortChannelMembers, true
}

// HasMaximumPortChannelMembers returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumPortChannelMembers() bool {
	if o != nil && o.MaximumPortChannelMembers != nil {
		return true
	}

	return false
}

// SetMaximumPortChannelMembers gets a reference to the given int64 and assigns it to the MaximumPortChannelMembers field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumPortChannelMembers(v int64) {
	o.MaximumPortChannelMembers = &v
}

// GetMaximumPrimaryVlan returns the MaximumPrimaryVlan field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPrimaryVlan() int64 {
	if o == nil || o.MaximumPrimaryVlan == nil {
		var ret int64
		return ret
	}
	return *o.MaximumPrimaryVlan
}

// GetMaximumPrimaryVlanOk returns a tuple with the MaximumPrimaryVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumPrimaryVlanOk() (*int64, bool) {
	if o == nil || o.MaximumPrimaryVlan == nil {
		return nil, false
	}
	return o.MaximumPrimaryVlan, true
}

// HasMaximumPrimaryVlan returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumPrimaryVlan() bool {
	if o != nil && o.MaximumPrimaryVlan != nil {
		return true
	}

	return false
}

// SetMaximumPrimaryVlan gets a reference to the given int64 and assigns it to the MaximumPrimaryVlan field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumPrimaryVlan(v int64) {
	o.MaximumPrimaryVlan = &v
}

// GetMaximumSecondaryVlan returns the MaximumSecondaryVlan field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlan() int64 {
	if o == nil || o.MaximumSecondaryVlan == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSecondaryVlan
}

// GetMaximumSecondaryVlanOk returns a tuple with the MaximumSecondaryVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanOk() (*int64, bool) {
	if o == nil || o.MaximumSecondaryVlan == nil {
		return nil, false
	}
	return o.MaximumSecondaryVlan, true
}

// HasMaximumSecondaryVlan returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumSecondaryVlan() bool {
	if o != nil && o.MaximumSecondaryVlan != nil {
		return true
	}

	return false
}

// SetMaximumSecondaryVlan gets a reference to the given int64 and assigns it to the MaximumSecondaryVlan field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumSecondaryVlan(v int64) {
	o.MaximumSecondaryVlan = &v
}

// GetMaximumSecondaryVlanPerPrimary returns the MaximumSecondaryVlanPerPrimary field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanPerPrimary() int64 {
	if o == nil || o.MaximumSecondaryVlanPerPrimary == nil {
		var ret int64
		return ret
	}
	return *o.MaximumSecondaryVlanPerPrimary
}

// GetMaximumSecondaryVlanPerPrimaryOk returns a tuple with the MaximumSecondaryVlanPerPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumSecondaryVlanPerPrimaryOk() (*int64, bool) {
	if o == nil || o.MaximumSecondaryVlanPerPrimary == nil {
		return nil, false
	}
	return o.MaximumSecondaryVlanPerPrimary, true
}

// HasMaximumSecondaryVlanPerPrimary returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumSecondaryVlanPerPrimary() bool {
	if o != nil && o.MaximumSecondaryVlanPerPrimary != nil {
		return true
	}

	return false
}

// SetMaximumSecondaryVlanPerPrimary gets a reference to the given int64 and assigns it to the MaximumSecondaryVlanPerPrimary field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumSecondaryVlanPerPrimary(v int64) {
	o.MaximumSecondaryVlanPerPrimary = &v
}

// GetMaximumVifs returns the MaximumVifs field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVifs() int64 {
	if o == nil || o.MaximumVifs == nil {
		var ret int64
		return ret
	}
	return *o.MaximumVifs
}

// GetMaximumVifsOk returns a tuple with the MaximumVifs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVifsOk() (*int64, bool) {
	if o == nil || o.MaximumVifs == nil {
		return nil, false
	}
	return o.MaximumVifs, true
}

// HasMaximumVifs returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumVifs() bool {
	if o != nil && o.MaximumVifs != nil {
		return true
	}

	return false
}

// SetMaximumVifs gets a reference to the given int64 and assigns it to the MaximumVifs field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumVifs(v int64) {
	o.MaximumVifs = &v
}

// GetMaximumVlans returns the MaximumVlans field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVlans() int64 {
	if o == nil || o.MaximumVlans == nil {
		var ret int64
		return ret
	}
	return *o.MaximumVlans
}

// GetMaximumVlansOk returns a tuple with the MaximumVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMaximumVlansOk() (*int64, bool) {
	if o == nil || o.MaximumVlans == nil {
		return nil, false
	}
	return o.MaximumVlans, true
}

// HasMaximumVlans returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMaximumVlans() bool {
	if o != nil && o.MaximumVlans != nil {
		return true
	}

	return false
}

// SetMaximumVlans gets a reference to the given int64 and assigns it to the MaximumVlans field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMaximumVlans(v int64) {
	o.MaximumVlans = &v
}

// GetMinimumActiveFans returns the MinimumActiveFans field value if set, zero value otherwise.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMinimumActiveFans() int64 {
	if o == nil || o.MinimumActiveFans == nil {
		var ret int64
		return ret
	}
	return *o.MinimumActiveFans
}

// GetMinimumActiveFansOk returns a tuple with the MinimumActiveFans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) GetMinimumActiveFansOk() (*int64, bool) {
	if o == nil || o.MinimumActiveFans == nil {
		return nil, false
	}
	return o.MinimumActiveFans, true
}

// HasMinimumActiveFans returns a boolean if a field has been set.
func (o *CapabilitySwitchNetworkLimitsAllOf) HasMinimumActiveFans() bool {
	if o != nil && o.MinimumActiveFans != nil {
		return true
	}

	return false
}

// SetMinimumActiveFans gets a reference to the given int64 and assigns it to the MinimumActiveFans field.
func (o *CapabilitySwitchNetworkLimitsAllOf) SetMinimumActiveFans(v int64) {
	o.MinimumActiveFans = &v
}

func (o CapabilitySwitchNetworkLimitsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.MaxCompressedPortVlanCount != nil {
		toSerialize["MaxCompressedPortVlanCount"] = o.MaxCompressedPortVlanCount
	}
	if o.MaxUncompressedPortVlanCount != nil {
		toSerialize["MaxUncompressedPortVlanCount"] = o.MaxUncompressedPortVlanCount
	}
	if o.MaximumActiveTrafficMonitoringSessions != nil {
		toSerialize["MaximumActiveTrafficMonitoringSessions"] = o.MaximumActiveTrafficMonitoringSessions
	}
	if o.MaximumEthernetPortChannels != nil {
		toSerialize["MaximumEthernetPortChannels"] = o.MaximumEthernetPortChannels
	}
	if o.MaximumEthernetUplinkPorts != nil {
		toSerialize["MaximumEthernetUplinkPorts"] = o.MaximumEthernetUplinkPorts
	}
	if o.MaximumFcPortChannelMembers != nil {
		toSerialize["MaximumFcPortChannelMembers"] = o.MaximumFcPortChannelMembers
	}
	if o.MaximumFcPortChannels != nil {
		toSerialize["MaximumFcPortChannels"] = o.MaximumFcPortChannels
	}
	if o.MaximumIgmpGroups != nil {
		toSerialize["MaximumIgmpGroups"] = o.MaximumIgmpGroups
	}
	if o.MaximumPortChannelMembers != nil {
		toSerialize["MaximumPortChannelMembers"] = o.MaximumPortChannelMembers
	}
	if o.MaximumPrimaryVlan != nil {
		toSerialize["MaximumPrimaryVlan"] = o.MaximumPrimaryVlan
	}
	if o.MaximumSecondaryVlan != nil {
		toSerialize["MaximumSecondaryVlan"] = o.MaximumSecondaryVlan
	}
	if o.MaximumSecondaryVlanPerPrimary != nil {
		toSerialize["MaximumSecondaryVlanPerPrimary"] = o.MaximumSecondaryVlanPerPrimary
	}
	if o.MaximumVifs != nil {
		toSerialize["MaximumVifs"] = o.MaximumVifs
	}
	if o.MaximumVlans != nil {
		toSerialize["MaximumVlans"] = o.MaximumVlans
	}
	if o.MinimumActiveFans != nil {
		toSerialize["MinimumActiveFans"] = o.MinimumActiveFans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilitySwitchNetworkLimitsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilitySwitchNetworkLimitsAllOf := _CapabilitySwitchNetworkLimitsAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilitySwitchNetworkLimitsAllOf); err == nil {
		*o = CapabilitySwitchNetworkLimitsAllOf(varCapabilitySwitchNetworkLimitsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "MaxCompressedPortVlanCount")
		delete(additionalProperties, "MaxUncompressedPortVlanCount")
		delete(additionalProperties, "MaximumActiveTrafficMonitoringSessions")
		delete(additionalProperties, "MaximumEthernetPortChannels")
		delete(additionalProperties, "MaximumEthernetUplinkPorts")
		delete(additionalProperties, "MaximumFcPortChannelMembers")
		delete(additionalProperties, "MaximumFcPortChannels")
		delete(additionalProperties, "MaximumIgmpGroups")
		delete(additionalProperties, "MaximumPortChannelMembers")
		delete(additionalProperties, "MaximumPrimaryVlan")
		delete(additionalProperties, "MaximumSecondaryVlan")
		delete(additionalProperties, "MaximumSecondaryVlanPerPrimary")
		delete(additionalProperties, "MaximumVifs")
		delete(additionalProperties, "MaximumVlans")
		delete(additionalProperties, "MinimumActiveFans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilitySwitchNetworkLimitsAllOf struct {
	value *CapabilitySwitchNetworkLimitsAllOf
	isSet bool
}

func (v NullableCapabilitySwitchNetworkLimitsAllOf) Get() *CapabilitySwitchNetworkLimitsAllOf {
	return v.value
}

func (v *NullableCapabilitySwitchNetworkLimitsAllOf) Set(val *CapabilitySwitchNetworkLimitsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySwitchNetworkLimitsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySwitchNetworkLimitsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySwitchNetworkLimitsAllOf(val *CapabilitySwitchNetworkLimitsAllOf) *NullableCapabilitySwitchNetworkLimitsAllOf {
	return &NullableCapabilitySwitchNetworkLimitsAllOf{value: val, isSet: true}
}

func (v NullableCapabilitySwitchNetworkLimitsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySwitchNetworkLimitsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
