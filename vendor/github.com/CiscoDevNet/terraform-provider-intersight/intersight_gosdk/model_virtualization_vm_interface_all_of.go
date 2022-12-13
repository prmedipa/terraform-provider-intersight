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

// VirtualizationVmInterfaceAllOf Definition of the list of properties defined in 'virtualization.VmInterface', excluding properties defined in parent classes.
type VirtualizationVmInterfaceAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual machine brige name of interface.
	Bridge    *string  `json:"Bridge,omitempty"`
	IpAddress []string `json:"IpAddress,omitempty"`
	// Virtual machine interface mac address.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Virtual machine interface model. * `Unknown` - The type of the network adaptor type is unknown. * `E1000` - Emulated version of the Intel 82545EM Gigabit Ethernet NIC. * `SRIOV` - Representation of a virtual function (VF) on a physical NIC with SR-IOV support. * `VMXNET2` - VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later. * `VMXNET3` - VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. * `E1000E` - E1000E – emulates a newer real network adapter, the 1 Gbit Intel 82574, and is available for Windows 2012 and later. The E1000E needs virtual machine hardware version 8 or later. * `NE2K_PCI` - The Ne2000 network card uses two ring buffers for packet handling. These are circular buffers made of 256-byte pages that the chip's DMA logic will use to store received packets or to get received packets. * `PCnet` - The PCnet-PCI II is a PCI network adapter. It has built-in support for CRC checks and can automatically pad short packets to the minimum Ethernet length. * `RTL8139` - The RTL8139 is a fast Ethernet card that operates at 10/100 Mbps. It is compliant with PCI version 2.0/2.1 and it is known for reliability and superior performance. * `VirtIO` - VirtIO is a standardized interface which allows virtual machines access to simplified \"virtual\" devices, such as block devices, network adapters and consoles. Accessing devices through VirtIO on a guest VM improves performance over more traditional \"emulated\" devices, as VirtIO devices require only the bare minimum setup and configuration needed to send and receive data, while the host machine handles the majority of the setup and maintenance of the actual physical hardware. * `` - Default network adaptor type supported by the hypervisor.
	Model *string `json:"Model,omitempty"`
	// Name of the virtual machine interface.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmInterfaceAllOf VirtualizationVmInterfaceAllOf

// NewVirtualizationVmInterfaceAllOf instantiates a new VirtualizationVmInterfaceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmInterfaceAllOf(classId string, objectType string) *VirtualizationVmInterfaceAllOf {
	this := VirtualizationVmInterfaceAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmInterfaceAllOfWithDefaults instantiates a new VirtualizationVmInterfaceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmInterfaceAllOfWithDefaults() *VirtualizationVmInterfaceAllOf {
	this := VirtualizationVmInterfaceAllOf{}
	var classId string = "virtualization.VmInterface"
	this.ClassId = classId
	var objectType string = "virtualization.VmInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmInterfaceAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmInterfaceAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmInterfaceAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmInterfaceAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *VirtualizationVmInterfaceAllOf) GetBridge() string {
	if o == nil || o.Bridge == nil {
		var ret string
		return ret
	}
	return *o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetBridgeOk() (*string, bool) {
	if o == nil || o.Bridge == nil {
		return nil, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *VirtualizationVmInterfaceAllOf) HasBridge() bool {
	if o != nil && o.Bridge != nil {
		return true
	}

	return false
}

// SetBridge gets a reference to the given string and assigns it to the Bridge field.
func (o *VirtualizationVmInterfaceAllOf) SetBridge(v string) {
	o.Bridge = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmInterfaceAllOf) GetIpAddress() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmInterfaceAllOf) GetIpAddressOk() ([]string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VirtualizationVmInterfaceAllOf) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given []string and assigns it to the IpAddress field.
func (o *VirtualizationVmInterfaceAllOf) SetIpAddress(v []string) {
	o.IpAddress = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationVmInterfaceAllOf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationVmInterfaceAllOf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationVmInterfaceAllOf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *VirtualizationVmInterfaceAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *VirtualizationVmInterfaceAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *VirtualizationVmInterfaceAllOf) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationVmInterfaceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmInterfaceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationVmInterfaceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationVmInterfaceAllOf) SetName(v string) {
	o.Name = &v
}

func (o VirtualizationVmInterfaceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bridge != nil {
		toSerialize["Bridge"] = o.Bridge
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmInterfaceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationVmInterfaceAllOf := _VirtualizationVmInterfaceAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationVmInterfaceAllOf); err == nil {
		*o = VirtualizationVmInterfaceAllOf(varVirtualizationVmInterfaceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bridge")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationVmInterfaceAllOf struct {
	value *VirtualizationVmInterfaceAllOf
	isSet bool
}

func (v NullableVirtualizationVmInterfaceAllOf) Get() *VirtualizationVmInterfaceAllOf {
	return v.value
}

func (v *NullableVirtualizationVmInterfaceAllOf) Set(val *VirtualizationVmInterfaceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmInterfaceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmInterfaceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmInterfaceAllOf(val *VirtualizationVmInterfaceAllOf) *NullableVirtualizationVmInterfaceAllOf {
	return &NullableVirtualizationVmInterfaceAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationVmInterfaceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmInterfaceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
