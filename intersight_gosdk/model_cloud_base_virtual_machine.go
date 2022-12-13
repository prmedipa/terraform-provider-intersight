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
	"time"
)

// CloudBaseVirtualMachine BaseVirtualMachine is an abstract base type for all cloud virtual machines.
type CloudBaseVirtualMachine struct {
	VirtualizationBaseVirtualMachine
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType                  string                            `json:"ObjectType"`
	BillingUnit                 NullableCloudBillingUnit          `json:"BillingUnit,omitempty"`
	ImageInfo                   NullableCloudImageReference       `json:"ImageInfo,omitempty"`
	InstanceType                NullableCloudInstanceType         `json:"InstanceType,omitempty"`
	NetworkInterfaceAttachments []CloudNetworkInterfaceAttachment `json:"NetworkInterfaceAttachments,omitempty"`
	// The private DNS hostname name assigned to the instance.
	PrivateDns *string `json:"PrivateDns,omitempty"`
	// The public DNS name assigned to the instance.
	PublicDns *string                  `json:"PublicDns,omitempty"`
	Region    NullableCloudCloudRegion `json:"Region,omitempty"`
	// How virtual machines are distributed across physical hardware and affects pricing.
	Tenancy *string `json:"Tenancy,omitempty"`
	// Time when this virtualmachine is terminated.
	TerminationTime      *time.Time                    `json:"TerminationTime,omitempty"`
	VirtualMachineTags   []CloudCloudTag               `json:"VirtualMachineTags,omitempty"`
	VolumeAttachments    []CloudVolumeAttachment       `json:"VolumeAttachments,omitempty"`
	Zone                 NullableCloudAvailabilityZone `json:"Zone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudBaseVirtualMachine CloudBaseVirtualMachine

// NewCloudBaseVirtualMachine instantiates a new CloudBaseVirtualMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBaseVirtualMachine(classId string, objectType string) *CloudBaseVirtualMachine {
	this := CloudBaseVirtualMachine{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	var powerState string = "Unknown"
	this.PowerState = &powerState
	var provider string = "Unknown"
	this.Provider = &provider
	var state string = "None"
	this.State = &state
	return &this
}

// NewCloudBaseVirtualMachineWithDefaults instantiates a new CloudBaseVirtualMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBaseVirtualMachineWithDefaults() *CloudBaseVirtualMachine {
	this := CloudBaseVirtualMachine{}
	var classId string = "cloud.AwsVirtualMachine"
	this.ClassId = classId
	var objectType string = "cloud.AwsVirtualMachine"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudBaseVirtualMachine) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudBaseVirtualMachine) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudBaseVirtualMachine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudBaseVirtualMachine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBillingUnit returns the BillingUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetBillingUnit() CloudBillingUnit {
	if o == nil || o.BillingUnit.Get() == nil {
		var ret CloudBillingUnit
		return ret
	}
	return *o.BillingUnit.Get()
}

// GetBillingUnitOk returns a tuple with the BillingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetBillingUnitOk() (*CloudBillingUnit, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingUnit.Get(), o.BillingUnit.IsSet()
}

// HasBillingUnit returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasBillingUnit() bool {
	if o != nil && o.BillingUnit.IsSet() {
		return true
	}

	return false
}

// SetBillingUnit gets a reference to the given NullableCloudBillingUnit and assigns it to the BillingUnit field.
func (o *CloudBaseVirtualMachine) SetBillingUnit(v CloudBillingUnit) {
	o.BillingUnit.Set(&v)
}

// SetBillingUnitNil sets the value for BillingUnit to be an explicit nil
func (o *CloudBaseVirtualMachine) SetBillingUnitNil() {
	o.BillingUnit.Set(nil)
}

// UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
func (o *CloudBaseVirtualMachine) UnsetBillingUnit() {
	o.BillingUnit.Unset()
}

// GetImageInfo returns the ImageInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetImageInfo() CloudImageReference {
	if o == nil || o.ImageInfo.Get() == nil {
		var ret CloudImageReference
		return ret
	}
	return *o.ImageInfo.Get()
}

// GetImageInfoOk returns a tuple with the ImageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetImageInfoOk() (*CloudImageReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageInfo.Get(), o.ImageInfo.IsSet()
}

// HasImageInfo returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasImageInfo() bool {
	if o != nil && o.ImageInfo.IsSet() {
		return true
	}

	return false
}

// SetImageInfo gets a reference to the given NullableCloudImageReference and assigns it to the ImageInfo field.
func (o *CloudBaseVirtualMachine) SetImageInfo(v CloudImageReference) {
	o.ImageInfo.Set(&v)
}

// SetImageInfoNil sets the value for ImageInfo to be an explicit nil
func (o *CloudBaseVirtualMachine) SetImageInfoNil() {
	o.ImageInfo.Set(nil)
}

// UnsetImageInfo ensures that no value is present for ImageInfo, not even an explicit nil
func (o *CloudBaseVirtualMachine) UnsetImageInfo() {
	o.ImageInfo.Unset()
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetInstanceType() CloudInstanceType {
	if o == nil || o.InstanceType.Get() == nil {
		var ret CloudInstanceType
		return ret
	}
	return *o.InstanceType.Get()
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetInstanceTypeOk() (*CloudInstanceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceType.Get(), o.InstanceType.IsSet()
}

// HasInstanceType returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasInstanceType() bool {
	if o != nil && o.InstanceType.IsSet() {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given NullableCloudInstanceType and assigns it to the InstanceType field.
func (o *CloudBaseVirtualMachine) SetInstanceType(v CloudInstanceType) {
	o.InstanceType.Set(&v)
}

// SetInstanceTypeNil sets the value for InstanceType to be an explicit nil
func (o *CloudBaseVirtualMachine) SetInstanceTypeNil() {
	o.InstanceType.Set(nil)
}

// UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
func (o *CloudBaseVirtualMachine) UnsetInstanceType() {
	o.InstanceType.Unset()
}

// GetNetworkInterfaceAttachments returns the NetworkInterfaceAttachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetNetworkInterfaceAttachments() []CloudNetworkInterfaceAttachment {
	if o == nil {
		var ret []CloudNetworkInterfaceAttachment
		return ret
	}
	return o.NetworkInterfaceAttachments
}

// GetNetworkInterfaceAttachmentsOk returns a tuple with the NetworkInterfaceAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetNetworkInterfaceAttachmentsOk() ([]CloudNetworkInterfaceAttachment, bool) {
	if o == nil || o.NetworkInterfaceAttachments == nil {
		return nil, false
	}
	return o.NetworkInterfaceAttachments, true
}

// HasNetworkInterfaceAttachments returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasNetworkInterfaceAttachments() bool {
	if o != nil && o.NetworkInterfaceAttachments != nil {
		return true
	}

	return false
}

// SetNetworkInterfaceAttachments gets a reference to the given []CloudNetworkInterfaceAttachment and assigns it to the NetworkInterfaceAttachments field.
func (o *CloudBaseVirtualMachine) SetNetworkInterfaceAttachments(v []CloudNetworkInterfaceAttachment) {
	o.NetworkInterfaceAttachments = v
}

// GetPrivateDns returns the PrivateDns field value if set, zero value otherwise.
func (o *CloudBaseVirtualMachine) GetPrivateDns() string {
	if o == nil || o.PrivateDns == nil {
		var ret string
		return ret
	}
	return *o.PrivateDns
}

// GetPrivateDnsOk returns a tuple with the PrivateDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetPrivateDnsOk() (*string, bool) {
	if o == nil || o.PrivateDns == nil {
		return nil, false
	}
	return o.PrivateDns, true
}

// HasPrivateDns returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasPrivateDns() bool {
	if o != nil && o.PrivateDns != nil {
		return true
	}

	return false
}

// SetPrivateDns gets a reference to the given string and assigns it to the PrivateDns field.
func (o *CloudBaseVirtualMachine) SetPrivateDns(v string) {
	o.PrivateDns = &v
}

// GetPublicDns returns the PublicDns field value if set, zero value otherwise.
func (o *CloudBaseVirtualMachine) GetPublicDns() string {
	if o == nil || o.PublicDns == nil {
		var ret string
		return ret
	}
	return *o.PublicDns
}

// GetPublicDnsOk returns a tuple with the PublicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetPublicDnsOk() (*string, bool) {
	if o == nil || o.PublicDns == nil {
		return nil, false
	}
	return o.PublicDns, true
}

// HasPublicDns returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasPublicDns() bool {
	if o != nil && o.PublicDns != nil {
		return true
	}

	return false
}

// SetPublicDns gets a reference to the given string and assigns it to the PublicDns field.
func (o *CloudBaseVirtualMachine) SetPublicDns(v string) {
	o.PublicDns = &v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetRegion() CloudCloudRegion {
	if o == nil || o.Region.Get() == nil {
		var ret CloudCloudRegion
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetRegionOk() (*CloudCloudRegion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableCloudCloudRegion and assigns it to the Region field.
func (o *CloudBaseVirtualMachine) SetRegion(v CloudCloudRegion) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *CloudBaseVirtualMachine) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *CloudBaseVirtualMachine) UnsetRegion() {
	o.Region.Unset()
}

// GetTenancy returns the Tenancy field value if set, zero value otherwise.
func (o *CloudBaseVirtualMachine) GetTenancy() string {
	if o == nil || o.Tenancy == nil {
		var ret string
		return ret
	}
	return *o.Tenancy
}

// GetTenancyOk returns a tuple with the Tenancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetTenancyOk() (*string, bool) {
	if o == nil || o.Tenancy == nil {
		return nil, false
	}
	return o.Tenancy, true
}

// HasTenancy returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasTenancy() bool {
	if o != nil && o.Tenancy != nil {
		return true
	}

	return false
}

// SetTenancy gets a reference to the given string and assigns it to the Tenancy field.
func (o *CloudBaseVirtualMachine) SetTenancy(v string) {
	o.Tenancy = &v
}

// GetTerminationTime returns the TerminationTime field value if set, zero value otherwise.
func (o *CloudBaseVirtualMachine) GetTerminationTime() time.Time {
	if o == nil || o.TerminationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.TerminationTime
}

// GetTerminationTimeOk returns a tuple with the TerminationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBaseVirtualMachine) GetTerminationTimeOk() (*time.Time, bool) {
	if o == nil || o.TerminationTime == nil {
		return nil, false
	}
	return o.TerminationTime, true
}

// HasTerminationTime returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasTerminationTime() bool {
	if o != nil && o.TerminationTime != nil {
		return true
	}

	return false
}

// SetTerminationTime gets a reference to the given time.Time and assigns it to the TerminationTime field.
func (o *CloudBaseVirtualMachine) SetTerminationTime(v time.Time) {
	o.TerminationTime = &v
}

// GetVirtualMachineTags returns the VirtualMachineTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetVirtualMachineTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.VirtualMachineTags
}

// GetVirtualMachineTagsOk returns a tuple with the VirtualMachineTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetVirtualMachineTagsOk() ([]CloudCloudTag, bool) {
	if o == nil || o.VirtualMachineTags == nil {
		return nil, false
	}
	return o.VirtualMachineTags, true
}

// HasVirtualMachineTags returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasVirtualMachineTags() bool {
	if o != nil && o.VirtualMachineTags != nil {
		return true
	}

	return false
}

// SetVirtualMachineTags gets a reference to the given []CloudCloudTag and assigns it to the VirtualMachineTags field.
func (o *CloudBaseVirtualMachine) SetVirtualMachineTags(v []CloudCloudTag) {
	o.VirtualMachineTags = v
}

// GetVolumeAttachments returns the VolumeAttachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetVolumeAttachments() []CloudVolumeAttachment {
	if o == nil {
		var ret []CloudVolumeAttachment
		return ret
	}
	return o.VolumeAttachments
}

// GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetVolumeAttachmentsOk() ([]CloudVolumeAttachment, bool) {
	if o == nil || o.VolumeAttachments == nil {
		return nil, false
	}
	return o.VolumeAttachments, true
}

// HasVolumeAttachments returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasVolumeAttachments() bool {
	if o != nil && o.VolumeAttachments != nil {
		return true
	}

	return false
}

// SetVolumeAttachments gets a reference to the given []CloudVolumeAttachment and assigns it to the VolumeAttachments field.
func (o *CloudBaseVirtualMachine) SetVolumeAttachments(v []CloudVolumeAttachment) {
	o.VolumeAttachments = v
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudBaseVirtualMachine) GetZone() CloudAvailabilityZone {
	if o == nil || o.Zone.Get() == nil {
		var ret CloudAvailabilityZone
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudBaseVirtualMachine) GetZoneOk() (*CloudAvailabilityZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *CloudBaseVirtualMachine) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableCloudAvailabilityZone and assigns it to the Zone field.
func (o *CloudBaseVirtualMachine) SetZone(v CloudAvailabilityZone) {
	o.Zone.Set(&v)
}

// SetZoneNil sets the value for Zone to be an explicit nil
func (o *CloudBaseVirtualMachine) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *CloudBaseVirtualMachine) UnsetZone() {
	o.Zone.Unset()
}

func (o CloudBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachine, errVirtualizationBaseVirtualMachine := json.Marshal(o.VirtualizationBaseVirtualMachine)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	errVirtualizationBaseVirtualMachine = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachine), &toSerialize)
	if errVirtualizationBaseVirtualMachine != nil {
		return []byte{}, errVirtualizationBaseVirtualMachine
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BillingUnit.IsSet() {
		toSerialize["BillingUnit"] = o.BillingUnit.Get()
	}
	if o.ImageInfo.IsSet() {
		toSerialize["ImageInfo"] = o.ImageInfo.Get()
	}
	if o.InstanceType.IsSet() {
		toSerialize["InstanceType"] = o.InstanceType.Get()
	}
	if o.NetworkInterfaceAttachments != nil {
		toSerialize["NetworkInterfaceAttachments"] = o.NetworkInterfaceAttachments
	}
	if o.PrivateDns != nil {
		toSerialize["PrivateDns"] = o.PrivateDns
	}
	if o.PublicDns != nil {
		toSerialize["PublicDns"] = o.PublicDns
	}
	if o.Region.IsSet() {
		toSerialize["Region"] = o.Region.Get()
	}
	if o.Tenancy != nil {
		toSerialize["Tenancy"] = o.Tenancy
	}
	if o.TerminationTime != nil {
		toSerialize["TerminationTime"] = o.TerminationTime
	}
	if o.VirtualMachineTags != nil {
		toSerialize["VirtualMachineTags"] = o.VirtualMachineTags
	}
	if o.VolumeAttachments != nil {
		toSerialize["VolumeAttachments"] = o.VolumeAttachments
	}
	if o.Zone.IsSet() {
		toSerialize["Zone"] = o.Zone.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudBaseVirtualMachine) UnmarshalJSON(bytes []byte) (err error) {
	type CloudBaseVirtualMachineWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType                  string                            `json:"ObjectType"`
		BillingUnit                 NullableCloudBillingUnit          `json:"BillingUnit,omitempty"`
		ImageInfo                   NullableCloudImageReference       `json:"ImageInfo,omitempty"`
		InstanceType                NullableCloudInstanceType         `json:"InstanceType,omitempty"`
		NetworkInterfaceAttachments []CloudNetworkInterfaceAttachment `json:"NetworkInterfaceAttachments,omitempty"`
		// The private DNS hostname name assigned to the instance.
		PrivateDns *string `json:"PrivateDns,omitempty"`
		// The public DNS name assigned to the instance.
		PublicDns *string                  `json:"PublicDns,omitempty"`
		Region    NullableCloudCloudRegion `json:"Region,omitempty"`
		// How virtual machines are distributed across physical hardware and affects pricing.
		Tenancy *string `json:"Tenancy,omitempty"`
		// Time when this virtualmachine is terminated.
		TerminationTime    *time.Time                    `json:"TerminationTime,omitempty"`
		VirtualMachineTags []CloudCloudTag               `json:"VirtualMachineTags,omitempty"`
		VolumeAttachments  []CloudVolumeAttachment       `json:"VolumeAttachments,omitempty"`
		Zone               NullableCloudAvailabilityZone `json:"Zone,omitempty"`
	}

	varCloudBaseVirtualMachineWithoutEmbeddedStruct := CloudBaseVirtualMachineWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudBaseVirtualMachineWithoutEmbeddedStruct)
	if err == nil {
		varCloudBaseVirtualMachine := _CloudBaseVirtualMachine{}
		varCloudBaseVirtualMachine.ClassId = varCloudBaseVirtualMachineWithoutEmbeddedStruct.ClassId
		varCloudBaseVirtualMachine.ObjectType = varCloudBaseVirtualMachineWithoutEmbeddedStruct.ObjectType
		varCloudBaseVirtualMachine.BillingUnit = varCloudBaseVirtualMachineWithoutEmbeddedStruct.BillingUnit
		varCloudBaseVirtualMachine.ImageInfo = varCloudBaseVirtualMachineWithoutEmbeddedStruct.ImageInfo
		varCloudBaseVirtualMachine.InstanceType = varCloudBaseVirtualMachineWithoutEmbeddedStruct.InstanceType
		varCloudBaseVirtualMachine.NetworkInterfaceAttachments = varCloudBaseVirtualMachineWithoutEmbeddedStruct.NetworkInterfaceAttachments
		varCloudBaseVirtualMachine.PrivateDns = varCloudBaseVirtualMachineWithoutEmbeddedStruct.PrivateDns
		varCloudBaseVirtualMachine.PublicDns = varCloudBaseVirtualMachineWithoutEmbeddedStruct.PublicDns
		varCloudBaseVirtualMachine.Region = varCloudBaseVirtualMachineWithoutEmbeddedStruct.Region
		varCloudBaseVirtualMachine.Tenancy = varCloudBaseVirtualMachineWithoutEmbeddedStruct.Tenancy
		varCloudBaseVirtualMachine.TerminationTime = varCloudBaseVirtualMachineWithoutEmbeddedStruct.TerminationTime
		varCloudBaseVirtualMachine.VirtualMachineTags = varCloudBaseVirtualMachineWithoutEmbeddedStruct.VirtualMachineTags
		varCloudBaseVirtualMachine.VolumeAttachments = varCloudBaseVirtualMachineWithoutEmbeddedStruct.VolumeAttachments
		varCloudBaseVirtualMachine.Zone = varCloudBaseVirtualMachineWithoutEmbeddedStruct.Zone
		*o = CloudBaseVirtualMachine(varCloudBaseVirtualMachine)
	} else {
		return err
	}

	varCloudBaseVirtualMachine := _CloudBaseVirtualMachine{}

	err = json.Unmarshal(bytes, &varCloudBaseVirtualMachine)
	if err == nil {
		o.VirtualizationBaseVirtualMachine = varCloudBaseVirtualMachine.VirtualizationBaseVirtualMachine
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BillingUnit")
		delete(additionalProperties, "ImageInfo")
		delete(additionalProperties, "InstanceType")
		delete(additionalProperties, "NetworkInterfaceAttachments")
		delete(additionalProperties, "PrivateDns")
		delete(additionalProperties, "PublicDns")
		delete(additionalProperties, "Region")
		delete(additionalProperties, "Tenancy")
		delete(additionalProperties, "TerminationTime")
		delete(additionalProperties, "VirtualMachineTags")
		delete(additionalProperties, "VolumeAttachments")
		delete(additionalProperties, "Zone")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualMachine := reflect.ValueOf(o.VirtualizationBaseVirtualMachine)
		for i := 0; i < reflectVirtualizationBaseVirtualMachine.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualMachine.Type().Field(i)

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

type NullableCloudBaseVirtualMachine struct {
	value *CloudBaseVirtualMachine
	isSet bool
}

func (v NullableCloudBaseVirtualMachine) Get() *CloudBaseVirtualMachine {
	return v.value
}

func (v *NullableCloudBaseVirtualMachine) Set(val *CloudBaseVirtualMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBaseVirtualMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBaseVirtualMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBaseVirtualMachine(val *CloudBaseVirtualMachine) *NullableCloudBaseVirtualMachine {
	return &NullableCloudBaseVirtualMachine{value: val, isSet: true}
}

func (v NullableCloudBaseVirtualMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBaseVirtualMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
