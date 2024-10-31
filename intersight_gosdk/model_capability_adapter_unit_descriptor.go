/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the CapabilityAdapterUnitDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CapabilityAdapterUnitDescriptor{}

// CapabilityAdapterUnitDescriptor Descriptor that uniquely identifies an adapter.
type CapabilityAdapterUnitDescriptor struct {
	CapabilityHardwareDescriptor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Generation of the adapter. * `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.
	AdapterGeneration *int32 `json:"AdapterGeneration,omitempty"`
	// Order in which the ports are connected; sequential or cyclic.
	ConnectivityOrder *string `json:"ConnectivityOrder,omitempty"`
	// The port speed for ethernet ports in Mbps.
	EthernetPortSpeed *int64                    `json:"EthernetPortSpeed,omitempty"`
	Features          []CapabilityFeatureConfig `json:"Features,omitempty"`
	// The port speed for fibre channel ports in Mbps.
	FibreChannelPortSpeed *int64 `json:"FibreChannelPortSpeed,omitempty"`
	// The number of SCSI I/O Queue resources to allocate.
	FibreChannelScsiIoqLimit *int64 `json:"FibreChannelScsiIoqLimit,omitempty"`
	// Indicates that the Azure Stack Host QoS feature is supported by this adapter.
	IsAzureQosSupported *bool `json:"IsAzureQosSupported,omitempty"`
	// Indicates that the GENEVE offload feature is supported by this adapter.
	IsGeneveSupported *bool `json:"IsGeneveSupported,omitempty"`
	// Indicates support for secure boot.
	IsSecureBootSupported *bool `json:"IsSecureBootSupported,omitempty"`
	// Maximum Ring Size value for vNIC Receive Queue.
	MaxEthRxRingSize *int64 `json:"MaxEthRxRingSize,omitempty"`
	// Maximum Ring Size value for vNIC Transmit Queue.
	MaxEthTxRingSize *int64 `json:"MaxEthTxRingSize,omitempty"`
	// Maximum number of vNIC interfaces that can be RoCEv2 enabled.
	MaxRocev2Interfaces *int64 `json:"MaxRocev2Interfaces,omitempty"`
	// Number of Dce Ports for the adapter.
	NumDcePorts *int64 `json:"NumDcePorts,omitempty"`
	// Indicates number of PCI Links of the adapter.
	NumberOfPciLinks *int64 `json:"NumberOfPciLinks,omitempty"`
	// Indicates PCI Link status of adapter.
	PciLink *int64 `json:"PciLink,omitempty"`
	// Prom card type for the adapter.
	PromCardType *string `json:"PromCardType,omitempty"`
	// Vic Id assigned for the adapter.
	VicId                *string `json:"VicId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityAdapterUnitDescriptor CapabilityAdapterUnitDescriptor

// NewCapabilityAdapterUnitDescriptor instantiates a new CapabilityAdapterUnitDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityAdapterUnitDescriptor(classId string, objectType string) *CapabilityAdapterUnitDescriptor {
	this := CapabilityAdapterUnitDescriptor{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adapterGeneration int32 = 4
	this.AdapterGeneration = &adapterGeneration
	var isAzureQosSupported bool = true
	this.IsAzureQosSupported = &isAzureQosSupported
	var isGeneveSupported bool = true
	this.IsGeneveSupported = &isGeneveSupported
	var isSecureBootSupported bool = false
	this.IsSecureBootSupported = &isSecureBootSupported
	var maxEthRxRingSize int64 = 4096
	this.MaxEthRxRingSize = &maxEthRxRingSize
	var maxEthTxRingSize int64 = 4096
	this.MaxEthTxRingSize = &maxEthTxRingSize
	var maxRocev2Interfaces int64 = 2
	this.MaxRocev2Interfaces = &maxRocev2Interfaces
	var numberOfPciLinks int64 = 1
	this.NumberOfPciLinks = &numberOfPciLinks
	var pciLink int64 = 0
	this.PciLink = &pciLink
	return &this
}

// NewCapabilityAdapterUnitDescriptorWithDefaults instantiates a new CapabilityAdapterUnitDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityAdapterUnitDescriptorWithDefaults() *CapabilityAdapterUnitDescriptor {
	this := CapabilityAdapterUnitDescriptor{}
	var classId string = "capability.AdapterUnitDescriptor"
	this.ClassId = classId
	var objectType string = "capability.AdapterUnitDescriptor"
	this.ObjectType = objectType
	var adapterGeneration int32 = 4
	this.AdapterGeneration = &adapterGeneration
	var isAzureQosSupported bool = true
	this.IsAzureQosSupported = &isAzureQosSupported
	var isGeneveSupported bool = true
	this.IsGeneveSupported = &isGeneveSupported
	var isSecureBootSupported bool = false
	this.IsSecureBootSupported = &isSecureBootSupported
	var maxEthRxRingSize int64 = 4096
	this.MaxEthRxRingSize = &maxEthRxRingSize
	var maxEthTxRingSize int64 = 4096
	this.MaxEthTxRingSize = &maxEthTxRingSize
	var maxRocev2Interfaces int64 = 2
	this.MaxRocev2Interfaces = &maxRocev2Interfaces
	var numberOfPciLinks int64 = 1
	this.NumberOfPciLinks = &numberOfPciLinks
	var pciLink int64 = 0
	this.PciLink = &pciLink
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityAdapterUnitDescriptor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityAdapterUnitDescriptor) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "capability.AdapterUnitDescriptor" of the ClassId field.
func (o *CapabilityAdapterUnitDescriptor) GetDefaultClassId() interface{} {
	return "capability.AdapterUnitDescriptor"
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityAdapterUnitDescriptor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityAdapterUnitDescriptor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "capability.AdapterUnitDescriptor" of the ObjectType field.
func (o *CapabilityAdapterUnitDescriptor) GetDefaultObjectType() interface{} {
	return "capability.AdapterUnitDescriptor"
}

// GetAdapterGeneration returns the AdapterGeneration field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetAdapterGeneration() int32 {
	if o == nil || IsNil(o.AdapterGeneration) {
		var ret int32
		return ret
	}
	return *o.AdapterGeneration
}

// GetAdapterGenerationOk returns a tuple with the AdapterGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetAdapterGenerationOk() (*int32, bool) {
	if o == nil || IsNil(o.AdapterGeneration) {
		return nil, false
	}
	return o.AdapterGeneration, true
}

// HasAdapterGeneration returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasAdapterGeneration() bool {
	if o != nil && !IsNil(o.AdapterGeneration) {
		return true
	}

	return false
}

// SetAdapterGeneration gets a reference to the given int32 and assigns it to the AdapterGeneration field.
func (o *CapabilityAdapterUnitDescriptor) SetAdapterGeneration(v int32) {
	o.AdapterGeneration = &v
}

// GetConnectivityOrder returns the ConnectivityOrder field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetConnectivityOrder() string {
	if o == nil || IsNil(o.ConnectivityOrder) {
		var ret string
		return ret
	}
	return *o.ConnectivityOrder
}

// GetConnectivityOrderOk returns a tuple with the ConnectivityOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetConnectivityOrderOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectivityOrder) {
		return nil, false
	}
	return o.ConnectivityOrder, true
}

// HasConnectivityOrder returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasConnectivityOrder() bool {
	if o != nil && !IsNil(o.ConnectivityOrder) {
		return true
	}

	return false
}

// SetConnectivityOrder gets a reference to the given string and assigns it to the ConnectivityOrder field.
func (o *CapabilityAdapterUnitDescriptor) SetConnectivityOrder(v string) {
	o.ConnectivityOrder = &v
}

// GetEthernetPortSpeed returns the EthernetPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetEthernetPortSpeed() int64 {
	if o == nil || IsNil(o.EthernetPortSpeed) {
		var ret int64
		return ret
	}
	return *o.EthernetPortSpeed
}

// GetEthernetPortSpeedOk returns a tuple with the EthernetPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetEthernetPortSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.EthernetPortSpeed) {
		return nil, false
	}
	return o.EthernetPortSpeed, true
}

// HasEthernetPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasEthernetPortSpeed() bool {
	if o != nil && !IsNil(o.EthernetPortSpeed) {
		return true
	}

	return false
}

// SetEthernetPortSpeed gets a reference to the given int64 and assigns it to the EthernetPortSpeed field.
func (o *CapabilityAdapterUnitDescriptor) SetEthernetPortSpeed(v int64) {
	o.EthernetPortSpeed = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityAdapterUnitDescriptor) GetFeatures() []CapabilityFeatureConfig {
	if o == nil {
		var ret []CapabilityFeatureConfig
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityAdapterUnitDescriptor) GetFeaturesOk() ([]CapabilityFeatureConfig, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []CapabilityFeatureConfig and assigns it to the Features field.
func (o *CapabilityAdapterUnitDescriptor) SetFeatures(v []CapabilityFeatureConfig) {
	o.Features = v
}

// GetFibreChannelPortSpeed returns the FibreChannelPortSpeed field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelPortSpeed() int64 {
	if o == nil || IsNil(o.FibreChannelPortSpeed) {
		var ret int64
		return ret
	}
	return *o.FibreChannelPortSpeed
}

// GetFibreChannelPortSpeedOk returns a tuple with the FibreChannelPortSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelPortSpeedOk() (*int64, bool) {
	if o == nil || IsNil(o.FibreChannelPortSpeed) {
		return nil, false
	}
	return o.FibreChannelPortSpeed, true
}

// HasFibreChannelPortSpeed returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasFibreChannelPortSpeed() bool {
	if o != nil && !IsNil(o.FibreChannelPortSpeed) {
		return true
	}

	return false
}

// SetFibreChannelPortSpeed gets a reference to the given int64 and assigns it to the FibreChannelPortSpeed field.
func (o *CapabilityAdapterUnitDescriptor) SetFibreChannelPortSpeed(v int64) {
	o.FibreChannelPortSpeed = &v
}

// GetFibreChannelScsiIoqLimit returns the FibreChannelScsiIoqLimit field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelScsiIoqLimit() int64 {
	if o == nil || IsNil(o.FibreChannelScsiIoqLimit) {
		var ret int64
		return ret
	}
	return *o.FibreChannelScsiIoqLimit
}

// GetFibreChannelScsiIoqLimitOk returns a tuple with the FibreChannelScsiIoqLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetFibreChannelScsiIoqLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.FibreChannelScsiIoqLimit) {
		return nil, false
	}
	return o.FibreChannelScsiIoqLimit, true
}

// HasFibreChannelScsiIoqLimit returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasFibreChannelScsiIoqLimit() bool {
	if o != nil && !IsNil(o.FibreChannelScsiIoqLimit) {
		return true
	}

	return false
}

// SetFibreChannelScsiIoqLimit gets a reference to the given int64 and assigns it to the FibreChannelScsiIoqLimit field.
func (o *CapabilityAdapterUnitDescriptor) SetFibreChannelScsiIoqLimit(v int64) {
	o.FibreChannelScsiIoqLimit = &v
}

// GetIsAzureQosSupported returns the IsAzureQosSupported field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetIsAzureQosSupported() bool {
	if o == nil || IsNil(o.IsAzureQosSupported) {
		var ret bool
		return ret
	}
	return *o.IsAzureQosSupported
}

// GetIsAzureQosSupportedOk returns a tuple with the IsAzureQosSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetIsAzureQosSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAzureQosSupported) {
		return nil, false
	}
	return o.IsAzureQosSupported, true
}

// HasIsAzureQosSupported returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasIsAzureQosSupported() bool {
	if o != nil && !IsNil(o.IsAzureQosSupported) {
		return true
	}

	return false
}

// SetIsAzureQosSupported gets a reference to the given bool and assigns it to the IsAzureQosSupported field.
func (o *CapabilityAdapterUnitDescriptor) SetIsAzureQosSupported(v bool) {
	o.IsAzureQosSupported = &v
}

// GetIsGeneveSupported returns the IsGeneveSupported field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetIsGeneveSupported() bool {
	if o == nil || IsNil(o.IsGeneveSupported) {
		var ret bool
		return ret
	}
	return *o.IsGeneveSupported
}

// GetIsGeneveSupportedOk returns a tuple with the IsGeneveSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetIsGeneveSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGeneveSupported) {
		return nil, false
	}
	return o.IsGeneveSupported, true
}

// HasIsGeneveSupported returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasIsGeneveSupported() bool {
	if o != nil && !IsNil(o.IsGeneveSupported) {
		return true
	}

	return false
}

// SetIsGeneveSupported gets a reference to the given bool and assigns it to the IsGeneveSupported field.
func (o *CapabilityAdapterUnitDescriptor) SetIsGeneveSupported(v bool) {
	o.IsGeneveSupported = &v
}

// GetIsSecureBootSupported returns the IsSecureBootSupported field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetIsSecureBootSupported() bool {
	if o == nil || IsNil(o.IsSecureBootSupported) {
		var ret bool
		return ret
	}
	return *o.IsSecureBootSupported
}

// GetIsSecureBootSupportedOk returns a tuple with the IsSecureBootSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetIsSecureBootSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSecureBootSupported) {
		return nil, false
	}
	return o.IsSecureBootSupported, true
}

// HasIsSecureBootSupported returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasIsSecureBootSupported() bool {
	if o != nil && !IsNil(o.IsSecureBootSupported) {
		return true
	}

	return false
}

// SetIsSecureBootSupported gets a reference to the given bool and assigns it to the IsSecureBootSupported field.
func (o *CapabilityAdapterUnitDescriptor) SetIsSecureBootSupported(v bool) {
	o.IsSecureBootSupported = &v
}

// GetMaxEthRxRingSize returns the MaxEthRxRingSize field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetMaxEthRxRingSize() int64 {
	if o == nil || IsNil(o.MaxEthRxRingSize) {
		var ret int64
		return ret
	}
	return *o.MaxEthRxRingSize
}

// GetMaxEthRxRingSizeOk returns a tuple with the MaxEthRxRingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetMaxEthRxRingSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxEthRxRingSize) {
		return nil, false
	}
	return o.MaxEthRxRingSize, true
}

// HasMaxEthRxRingSize returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasMaxEthRxRingSize() bool {
	if o != nil && !IsNil(o.MaxEthRxRingSize) {
		return true
	}

	return false
}

// SetMaxEthRxRingSize gets a reference to the given int64 and assigns it to the MaxEthRxRingSize field.
func (o *CapabilityAdapterUnitDescriptor) SetMaxEthRxRingSize(v int64) {
	o.MaxEthRxRingSize = &v
}

// GetMaxEthTxRingSize returns the MaxEthTxRingSize field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetMaxEthTxRingSize() int64 {
	if o == nil || IsNil(o.MaxEthTxRingSize) {
		var ret int64
		return ret
	}
	return *o.MaxEthTxRingSize
}

// GetMaxEthTxRingSizeOk returns a tuple with the MaxEthTxRingSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetMaxEthTxRingSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxEthTxRingSize) {
		return nil, false
	}
	return o.MaxEthTxRingSize, true
}

// HasMaxEthTxRingSize returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasMaxEthTxRingSize() bool {
	if o != nil && !IsNil(o.MaxEthTxRingSize) {
		return true
	}

	return false
}

// SetMaxEthTxRingSize gets a reference to the given int64 and assigns it to the MaxEthTxRingSize field.
func (o *CapabilityAdapterUnitDescriptor) SetMaxEthTxRingSize(v int64) {
	o.MaxEthTxRingSize = &v
}

// GetMaxRocev2Interfaces returns the MaxRocev2Interfaces field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetMaxRocev2Interfaces() int64 {
	if o == nil || IsNil(o.MaxRocev2Interfaces) {
		var ret int64
		return ret
	}
	return *o.MaxRocev2Interfaces
}

// GetMaxRocev2InterfacesOk returns a tuple with the MaxRocev2Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetMaxRocev2InterfacesOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxRocev2Interfaces) {
		return nil, false
	}
	return o.MaxRocev2Interfaces, true
}

// HasMaxRocev2Interfaces returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasMaxRocev2Interfaces() bool {
	if o != nil && !IsNil(o.MaxRocev2Interfaces) {
		return true
	}

	return false
}

// SetMaxRocev2Interfaces gets a reference to the given int64 and assigns it to the MaxRocev2Interfaces field.
func (o *CapabilityAdapterUnitDescriptor) SetMaxRocev2Interfaces(v int64) {
	o.MaxRocev2Interfaces = &v
}

// GetNumDcePorts returns the NumDcePorts field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetNumDcePorts() int64 {
	if o == nil || IsNil(o.NumDcePorts) {
		var ret int64
		return ret
	}
	return *o.NumDcePorts
}

// GetNumDcePortsOk returns a tuple with the NumDcePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetNumDcePortsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumDcePorts) {
		return nil, false
	}
	return o.NumDcePorts, true
}

// HasNumDcePorts returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasNumDcePorts() bool {
	if o != nil && !IsNil(o.NumDcePorts) {
		return true
	}

	return false
}

// SetNumDcePorts gets a reference to the given int64 and assigns it to the NumDcePorts field.
func (o *CapabilityAdapterUnitDescriptor) SetNumDcePorts(v int64) {
	o.NumDcePorts = &v
}

// GetNumberOfPciLinks returns the NumberOfPciLinks field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetNumberOfPciLinks() int64 {
	if o == nil || IsNil(o.NumberOfPciLinks) {
		var ret int64
		return ret
	}
	return *o.NumberOfPciLinks
}

// GetNumberOfPciLinksOk returns a tuple with the NumberOfPciLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetNumberOfPciLinksOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfPciLinks) {
		return nil, false
	}
	return o.NumberOfPciLinks, true
}

// HasNumberOfPciLinks returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasNumberOfPciLinks() bool {
	if o != nil && !IsNil(o.NumberOfPciLinks) {
		return true
	}

	return false
}

// SetNumberOfPciLinks gets a reference to the given int64 and assigns it to the NumberOfPciLinks field.
func (o *CapabilityAdapterUnitDescriptor) SetNumberOfPciLinks(v int64) {
	o.NumberOfPciLinks = &v
}

// GetPciLink returns the PciLink field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetPciLink() int64 {
	if o == nil || IsNil(o.PciLink) {
		var ret int64
		return ret
	}
	return *o.PciLink
}

// GetPciLinkOk returns a tuple with the PciLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetPciLinkOk() (*int64, bool) {
	if o == nil || IsNil(o.PciLink) {
		return nil, false
	}
	return o.PciLink, true
}

// HasPciLink returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasPciLink() bool {
	if o != nil && !IsNil(o.PciLink) {
		return true
	}

	return false
}

// SetPciLink gets a reference to the given int64 and assigns it to the PciLink field.
func (o *CapabilityAdapterUnitDescriptor) SetPciLink(v int64) {
	o.PciLink = &v
}

// GetPromCardType returns the PromCardType field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetPromCardType() string {
	if o == nil || IsNil(o.PromCardType) {
		var ret string
		return ret
	}
	return *o.PromCardType
}

// GetPromCardTypeOk returns a tuple with the PromCardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetPromCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PromCardType) {
		return nil, false
	}
	return o.PromCardType, true
}

// HasPromCardType returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasPromCardType() bool {
	if o != nil && !IsNil(o.PromCardType) {
		return true
	}

	return false
}

// SetPromCardType gets a reference to the given string and assigns it to the PromCardType field.
func (o *CapabilityAdapterUnitDescriptor) SetPromCardType(v string) {
	o.PromCardType = &v
}

// GetVicId returns the VicId field value if set, zero value otherwise.
func (o *CapabilityAdapterUnitDescriptor) GetVicId() string {
	if o == nil || IsNil(o.VicId) {
		var ret string
		return ret
	}
	return *o.VicId
}

// GetVicIdOk returns a tuple with the VicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityAdapterUnitDescriptor) GetVicIdOk() (*string, bool) {
	if o == nil || IsNil(o.VicId) {
		return nil, false
	}
	return o.VicId, true
}

// HasVicId returns a boolean if a field has been set.
func (o *CapabilityAdapterUnitDescriptor) HasVicId() bool {
	if o != nil && !IsNil(o.VicId) {
		return true
	}

	return false
}

// SetVicId gets a reference to the given string and assigns it to the VicId field.
func (o *CapabilityAdapterUnitDescriptor) SetVicId(v string) {
	o.VicId = &v
}

func (o CapabilityAdapterUnitDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityAdapterUnitDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCapabilityHardwareDescriptor, errCapabilityHardwareDescriptor := json.Marshal(o.CapabilityHardwareDescriptor)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	errCapabilityHardwareDescriptor = json.Unmarshal([]byte(serializedCapabilityHardwareDescriptor), &toSerialize)
	if errCapabilityHardwareDescriptor != nil {
		return map[string]interface{}{}, errCapabilityHardwareDescriptor
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdapterGeneration) {
		toSerialize["AdapterGeneration"] = o.AdapterGeneration
	}
	if !IsNil(o.ConnectivityOrder) {
		toSerialize["ConnectivityOrder"] = o.ConnectivityOrder
	}
	if !IsNil(o.EthernetPortSpeed) {
		toSerialize["EthernetPortSpeed"] = o.EthernetPortSpeed
	}
	if o.Features != nil {
		toSerialize["Features"] = o.Features
	}
	if !IsNil(o.FibreChannelPortSpeed) {
		toSerialize["FibreChannelPortSpeed"] = o.FibreChannelPortSpeed
	}
	if !IsNil(o.FibreChannelScsiIoqLimit) {
		toSerialize["FibreChannelScsiIoqLimit"] = o.FibreChannelScsiIoqLimit
	}
	if !IsNil(o.IsAzureQosSupported) {
		toSerialize["IsAzureQosSupported"] = o.IsAzureQosSupported
	}
	if !IsNil(o.IsGeneveSupported) {
		toSerialize["IsGeneveSupported"] = o.IsGeneveSupported
	}
	if !IsNil(o.IsSecureBootSupported) {
		toSerialize["IsSecureBootSupported"] = o.IsSecureBootSupported
	}
	if !IsNil(o.MaxEthRxRingSize) {
		toSerialize["MaxEthRxRingSize"] = o.MaxEthRxRingSize
	}
	if !IsNil(o.MaxEthTxRingSize) {
		toSerialize["MaxEthTxRingSize"] = o.MaxEthTxRingSize
	}
	if !IsNil(o.MaxRocev2Interfaces) {
		toSerialize["MaxRocev2Interfaces"] = o.MaxRocev2Interfaces
	}
	if !IsNil(o.NumDcePorts) {
		toSerialize["NumDcePorts"] = o.NumDcePorts
	}
	if !IsNil(o.NumberOfPciLinks) {
		toSerialize["NumberOfPciLinks"] = o.NumberOfPciLinks
	}
	if !IsNil(o.PciLink) {
		toSerialize["PciLink"] = o.PciLink
	}
	if !IsNil(o.PromCardType) {
		toSerialize["PromCardType"] = o.PromCardType
	}
	if !IsNil(o.VicId) {
		toSerialize["VicId"] = o.VicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CapabilityAdapterUnitDescriptor) UnmarshalJSON(data []byte) (err error) {
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
	type CapabilityAdapterUnitDescriptorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Generation of the adapter. * `4` - Fourth generation adapters (14xx). The PIDs of these adapters end with the string 04. * `2` - Second generation VIC adapters (12xx). The PIDs of these adapters end with the string 02. * `3` - Third generation adapters (13xx). The PIDs of these adapters end with the string 03. * `5` - Fifth generation adapters (15xx). The PIDs of these adapters contain the V5 string.
		AdapterGeneration *int32 `json:"AdapterGeneration,omitempty"`
		// Order in which the ports are connected; sequential or cyclic.
		ConnectivityOrder *string `json:"ConnectivityOrder,omitempty"`
		// The port speed for ethernet ports in Mbps.
		EthernetPortSpeed *int64                    `json:"EthernetPortSpeed,omitempty"`
		Features          []CapabilityFeatureConfig `json:"Features,omitempty"`
		// The port speed for fibre channel ports in Mbps.
		FibreChannelPortSpeed *int64 `json:"FibreChannelPortSpeed,omitempty"`
		// The number of SCSI I/O Queue resources to allocate.
		FibreChannelScsiIoqLimit *int64 `json:"FibreChannelScsiIoqLimit,omitempty"`
		// Indicates that the Azure Stack Host QoS feature is supported by this adapter.
		IsAzureQosSupported *bool `json:"IsAzureQosSupported,omitempty"`
		// Indicates that the GENEVE offload feature is supported by this adapter.
		IsGeneveSupported *bool `json:"IsGeneveSupported,omitempty"`
		// Indicates support for secure boot.
		IsSecureBootSupported *bool `json:"IsSecureBootSupported,omitempty"`
		// Maximum Ring Size value for vNIC Receive Queue.
		MaxEthRxRingSize *int64 `json:"MaxEthRxRingSize,omitempty"`
		// Maximum Ring Size value for vNIC Transmit Queue.
		MaxEthTxRingSize *int64 `json:"MaxEthTxRingSize,omitempty"`
		// Maximum number of vNIC interfaces that can be RoCEv2 enabled.
		MaxRocev2Interfaces *int64 `json:"MaxRocev2Interfaces,omitempty"`
		// Number of Dce Ports for the adapter.
		NumDcePorts *int64 `json:"NumDcePorts,omitempty"`
		// Indicates number of PCI Links of the adapter.
		NumberOfPciLinks *int64 `json:"NumberOfPciLinks,omitempty"`
		// Indicates PCI Link status of adapter.
		PciLink *int64 `json:"PciLink,omitempty"`
		// Prom card type for the adapter.
		PromCardType *string `json:"PromCardType,omitempty"`
		// Vic Id assigned for the adapter.
		VicId *string `json:"VicId,omitempty"`
	}

	varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct := CapabilityAdapterUnitDescriptorWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct)
	if err == nil {
		varCapabilityAdapterUnitDescriptor := _CapabilityAdapterUnitDescriptor{}
		varCapabilityAdapterUnitDescriptor.ClassId = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.ClassId
		varCapabilityAdapterUnitDescriptor.ObjectType = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.ObjectType
		varCapabilityAdapterUnitDescriptor.AdapterGeneration = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.AdapterGeneration
		varCapabilityAdapterUnitDescriptor.ConnectivityOrder = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.ConnectivityOrder
		varCapabilityAdapterUnitDescriptor.EthernetPortSpeed = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.EthernetPortSpeed
		varCapabilityAdapterUnitDescriptor.Features = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.Features
		varCapabilityAdapterUnitDescriptor.FibreChannelPortSpeed = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.FibreChannelPortSpeed
		varCapabilityAdapterUnitDescriptor.FibreChannelScsiIoqLimit = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.FibreChannelScsiIoqLimit
		varCapabilityAdapterUnitDescriptor.IsAzureQosSupported = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.IsAzureQosSupported
		varCapabilityAdapterUnitDescriptor.IsGeneveSupported = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.IsGeneveSupported
		varCapabilityAdapterUnitDescriptor.IsSecureBootSupported = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.IsSecureBootSupported
		varCapabilityAdapterUnitDescriptor.MaxEthRxRingSize = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.MaxEthRxRingSize
		varCapabilityAdapterUnitDescriptor.MaxEthTxRingSize = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.MaxEthTxRingSize
		varCapabilityAdapterUnitDescriptor.MaxRocev2Interfaces = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.MaxRocev2Interfaces
		varCapabilityAdapterUnitDescriptor.NumDcePorts = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.NumDcePorts
		varCapabilityAdapterUnitDescriptor.NumberOfPciLinks = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.NumberOfPciLinks
		varCapabilityAdapterUnitDescriptor.PciLink = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.PciLink
		varCapabilityAdapterUnitDescriptor.PromCardType = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.PromCardType
		varCapabilityAdapterUnitDescriptor.VicId = varCapabilityAdapterUnitDescriptorWithoutEmbeddedStruct.VicId
		*o = CapabilityAdapterUnitDescriptor(varCapabilityAdapterUnitDescriptor)
	} else {
		return err
	}

	varCapabilityAdapterUnitDescriptor := _CapabilityAdapterUnitDescriptor{}

	err = json.Unmarshal(data, &varCapabilityAdapterUnitDescriptor)
	if err == nil {
		o.CapabilityHardwareDescriptor = varCapabilityAdapterUnitDescriptor.CapabilityHardwareDescriptor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdapterGeneration")
		delete(additionalProperties, "ConnectivityOrder")
		delete(additionalProperties, "EthernetPortSpeed")
		delete(additionalProperties, "Features")
		delete(additionalProperties, "FibreChannelPortSpeed")
		delete(additionalProperties, "FibreChannelScsiIoqLimit")
		delete(additionalProperties, "IsAzureQosSupported")
		delete(additionalProperties, "IsGeneveSupported")
		delete(additionalProperties, "IsSecureBootSupported")
		delete(additionalProperties, "MaxEthRxRingSize")
		delete(additionalProperties, "MaxEthTxRingSize")
		delete(additionalProperties, "MaxRocev2Interfaces")
		delete(additionalProperties, "NumDcePorts")
		delete(additionalProperties, "NumberOfPciLinks")
		delete(additionalProperties, "PciLink")
		delete(additionalProperties, "PromCardType")
		delete(additionalProperties, "VicId")

		// remove fields from embedded structs
		reflectCapabilityHardwareDescriptor := reflect.ValueOf(o.CapabilityHardwareDescriptor)
		for i := 0; i < reflectCapabilityHardwareDescriptor.Type().NumField(); i++ {
			t := reflectCapabilityHardwareDescriptor.Type().Field(i)

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

type NullableCapabilityAdapterUnitDescriptor struct {
	value *CapabilityAdapterUnitDescriptor
	isSet bool
}

func (v NullableCapabilityAdapterUnitDescriptor) Get() *CapabilityAdapterUnitDescriptor {
	return v.value
}

func (v *NullableCapabilityAdapterUnitDescriptor) Set(val *CapabilityAdapterUnitDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityAdapterUnitDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityAdapterUnitDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityAdapterUnitDescriptor(val *CapabilityAdapterUnitDescriptor) *NullableCapabilityAdapterUnitDescriptor {
	return &NullableCapabilityAdapterUnitDescriptor{value: val, isSet: true}
}

func (v NullableCapabilityAdapterUnitDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityAdapterUnitDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
