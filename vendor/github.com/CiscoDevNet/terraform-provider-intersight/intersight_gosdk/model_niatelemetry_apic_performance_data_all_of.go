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

// NiatelemetryApicPerformanceDataAllOf Definition of the list of properties defined in 'niatelemetry.ApicPerformanceData', excluding properties defined in parent classes.
type NiatelemetryApicPerformanceDataAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType               string                                 `json:"ObjectType"`
	DigitalOpticalMonitoring []NiatelemetryDigitalOpticalMonitoring `json:"DigitalOpticalMonitoring,omitempty"`
	// Dn of the fabric nodes in the apic.
	Dn                             *string                                            `json:"Dn,omitempty"`
	EqptStorageFirmware            NullableNiatelemetryEqptStorageFirmware            `json:"EqptStorageFirmware,omitempty"`
	EqptcapacityPolUsage5min       NullableNiatelemetryEqptcapacityPolUsage5min       `json:"EqptcapacityPolUsage5min,omitempty"`
	EqptcapacityPrefixEntries15min NullableNiatelemetryEqptcapacityPrefixEntries15min `json:"EqptcapacityPrefixEntries15min,omitempty"`
	EqptcapacityPrefixEntries5min  NullableNiatelemetryEqptcapacityPrefixEntries5min  `json:"EqptcapacityPrefixEntries5min,omitempty"`
	// Health of the fabric nodes in the apic.
	NodeHealth      *int64                              `json:"NodeHealth,omitempty"`
	ProcSysCpu15min NullableNiatelemetryProcSysCpu15min `json:"ProcSysCpu15min,omitempty"`
	ProcSysCpu5min  NullableNiatelemetryProcSysCpu5min  `json:"ProcSysCpu5min,omitempty"`
	ProcSysMem15min NullableNiatelemetryProcSysMem15min `json:"ProcSysMem15min,omitempty"`
	ProcSysMem5min  NullableNiatelemetryProcSysMem5min  `json:"ProcSysMem5min,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected..
	SiteName              *string                              `json:"SiteName,omitempty"`
	SwitchDiskUtilization []NiatelemetrySwitchDiskUtilization  `json:"SwitchDiskUtilization,omitempty"`
	RegisteredDevice      *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _NiatelemetryApicPerformanceDataAllOf NiatelemetryApicPerformanceDataAllOf

// NewNiatelemetryApicPerformanceDataAllOf instantiates a new NiatelemetryApicPerformanceDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicPerformanceDataAllOf(classId string, objectType string) *NiatelemetryApicPerformanceDataAllOf {
	this := NiatelemetryApicPerformanceDataAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicPerformanceDataAllOfWithDefaults instantiates a new NiatelemetryApicPerformanceDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicPerformanceDataAllOfWithDefaults() *NiatelemetryApicPerformanceDataAllOf {
	this := NiatelemetryApicPerformanceDataAllOf{}
	var classId string = "niatelemetry.ApicPerformanceData"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicPerformanceData"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicPerformanceDataAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicPerformanceDataAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicPerformanceDataAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicPerformanceDataAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDigitalOpticalMonitoring returns the DigitalOpticalMonitoring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetDigitalOpticalMonitoring() []NiatelemetryDigitalOpticalMonitoring {
	if o == nil {
		var ret []NiatelemetryDigitalOpticalMonitoring
		return ret
	}
	return o.DigitalOpticalMonitoring
}

// GetDigitalOpticalMonitoringOk returns a tuple with the DigitalOpticalMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetDigitalOpticalMonitoringOk() ([]NiatelemetryDigitalOpticalMonitoring, bool) {
	if o == nil || o.DigitalOpticalMonitoring == nil {
		return nil, false
	}
	return o.DigitalOpticalMonitoring, true
}

// HasDigitalOpticalMonitoring returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasDigitalOpticalMonitoring() bool {
	if o != nil && o.DigitalOpticalMonitoring != nil {
		return true
	}

	return false
}

// SetDigitalOpticalMonitoring gets a reference to the given []NiatelemetryDigitalOpticalMonitoring and assigns it to the DigitalOpticalMonitoring field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetDigitalOpticalMonitoring(v []NiatelemetryDigitalOpticalMonitoring) {
	o.DigitalOpticalMonitoring = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetDn(v string) {
	o.Dn = &v
}

// GetEqptStorageFirmware returns the EqptStorageFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptStorageFirmware() NiatelemetryEqptStorageFirmware {
	if o == nil || o.EqptStorageFirmware.Get() == nil {
		var ret NiatelemetryEqptStorageFirmware
		return ret
	}
	return *o.EqptStorageFirmware.Get()
}

// GetEqptStorageFirmwareOk returns a tuple with the EqptStorageFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptStorageFirmwareOk() (*NiatelemetryEqptStorageFirmware, bool) {
	if o == nil {
		return nil, false
	}
	return o.EqptStorageFirmware.Get(), o.EqptStorageFirmware.IsSet()
}

// HasEqptStorageFirmware returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptStorageFirmware() bool {
	if o != nil && o.EqptStorageFirmware.IsSet() {
		return true
	}

	return false
}

// SetEqptStorageFirmware gets a reference to the given NullableNiatelemetryEqptStorageFirmware and assigns it to the EqptStorageFirmware field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptStorageFirmware(v NiatelemetryEqptStorageFirmware) {
	o.EqptStorageFirmware.Set(&v)
}

// SetEqptStorageFirmwareNil sets the value for EqptStorageFirmware to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptStorageFirmwareNil() {
	o.EqptStorageFirmware.Set(nil)
}

// UnsetEqptStorageFirmware ensures that no value is present for EqptStorageFirmware, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptStorageFirmware() {
	o.EqptStorageFirmware.Unset()
}

// GetEqptcapacityPolUsage5min returns the EqptcapacityPolUsage5min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPolUsage5min() NiatelemetryEqptcapacityPolUsage5min {
	if o == nil || o.EqptcapacityPolUsage5min.Get() == nil {
		var ret NiatelemetryEqptcapacityPolUsage5min
		return ret
	}
	return *o.EqptcapacityPolUsage5min.Get()
}

// GetEqptcapacityPolUsage5minOk returns a tuple with the EqptcapacityPolUsage5min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPolUsage5minOk() (*NiatelemetryEqptcapacityPolUsage5min, bool) {
	if o == nil {
		return nil, false
	}
	return o.EqptcapacityPolUsage5min.Get(), o.EqptcapacityPolUsage5min.IsSet()
}

// HasEqptcapacityPolUsage5min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPolUsage5min() bool {
	if o != nil && o.EqptcapacityPolUsage5min.IsSet() {
		return true
	}

	return false
}

// SetEqptcapacityPolUsage5min gets a reference to the given NullableNiatelemetryEqptcapacityPolUsage5min and assigns it to the EqptcapacityPolUsage5min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPolUsage5min(v NiatelemetryEqptcapacityPolUsage5min) {
	o.EqptcapacityPolUsage5min.Set(&v)
}

// SetEqptcapacityPolUsage5minNil sets the value for EqptcapacityPolUsage5min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPolUsage5minNil() {
	o.EqptcapacityPolUsage5min.Set(nil)
}

// UnsetEqptcapacityPolUsage5min ensures that no value is present for EqptcapacityPolUsage5min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPolUsage5min() {
	o.EqptcapacityPolUsage5min.Unset()
}

// GetEqptcapacityPrefixEntries15min returns the EqptcapacityPrefixEntries15min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries15min() NiatelemetryEqptcapacityPrefixEntries15min {
	if o == nil || o.EqptcapacityPrefixEntries15min.Get() == nil {
		var ret NiatelemetryEqptcapacityPrefixEntries15min
		return ret
	}
	return *o.EqptcapacityPrefixEntries15min.Get()
}

// GetEqptcapacityPrefixEntries15minOk returns a tuple with the EqptcapacityPrefixEntries15min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries15minOk() (*NiatelemetryEqptcapacityPrefixEntries15min, bool) {
	if o == nil {
		return nil, false
	}
	return o.EqptcapacityPrefixEntries15min.Get(), o.EqptcapacityPrefixEntries15min.IsSet()
}

// HasEqptcapacityPrefixEntries15min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPrefixEntries15min() bool {
	if o != nil && o.EqptcapacityPrefixEntries15min.IsSet() {
		return true
	}

	return false
}

// SetEqptcapacityPrefixEntries15min gets a reference to the given NullableNiatelemetryEqptcapacityPrefixEntries15min and assigns it to the EqptcapacityPrefixEntries15min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries15min(v NiatelemetryEqptcapacityPrefixEntries15min) {
	o.EqptcapacityPrefixEntries15min.Set(&v)
}

// SetEqptcapacityPrefixEntries15minNil sets the value for EqptcapacityPrefixEntries15min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries15minNil() {
	o.EqptcapacityPrefixEntries15min.Set(nil)
}

// UnsetEqptcapacityPrefixEntries15min ensures that no value is present for EqptcapacityPrefixEntries15min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPrefixEntries15min() {
	o.EqptcapacityPrefixEntries15min.Unset()
}

// GetEqptcapacityPrefixEntries5min returns the EqptcapacityPrefixEntries5min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries5min() NiatelemetryEqptcapacityPrefixEntries5min {
	if o == nil || o.EqptcapacityPrefixEntries5min.Get() == nil {
		var ret NiatelemetryEqptcapacityPrefixEntries5min
		return ret
	}
	return *o.EqptcapacityPrefixEntries5min.Get()
}

// GetEqptcapacityPrefixEntries5minOk returns a tuple with the EqptcapacityPrefixEntries5min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries5minOk() (*NiatelemetryEqptcapacityPrefixEntries5min, bool) {
	if o == nil {
		return nil, false
	}
	return o.EqptcapacityPrefixEntries5min.Get(), o.EqptcapacityPrefixEntries5min.IsSet()
}

// HasEqptcapacityPrefixEntries5min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPrefixEntries5min() bool {
	if o != nil && o.EqptcapacityPrefixEntries5min.IsSet() {
		return true
	}

	return false
}

// SetEqptcapacityPrefixEntries5min gets a reference to the given NullableNiatelemetryEqptcapacityPrefixEntries5min and assigns it to the EqptcapacityPrefixEntries5min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries5min(v NiatelemetryEqptcapacityPrefixEntries5min) {
	o.EqptcapacityPrefixEntries5min.Set(&v)
}

// SetEqptcapacityPrefixEntries5minNil sets the value for EqptcapacityPrefixEntries5min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries5minNil() {
	o.EqptcapacityPrefixEntries5min.Set(nil)
}

// UnsetEqptcapacityPrefixEntries5min ensures that no value is present for EqptcapacityPrefixEntries5min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPrefixEntries5min() {
	o.EqptcapacityPrefixEntries5min.Unset()
}

// GetNodeHealth returns the NodeHealth field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetNodeHealth() int64 {
	if o == nil || o.NodeHealth == nil {
		var ret int64
		return ret
	}
	return *o.NodeHealth
}

// GetNodeHealthOk returns a tuple with the NodeHealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetNodeHealthOk() (*int64, bool) {
	if o == nil || o.NodeHealth == nil {
		return nil, false
	}
	return o.NodeHealth, true
}

// HasNodeHealth returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasNodeHealth() bool {
	if o != nil && o.NodeHealth != nil {
		return true
	}

	return false
}

// SetNodeHealth gets a reference to the given int64 and assigns it to the NodeHealth field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetNodeHealth(v int64) {
	o.NodeHealth = &v
}

// GetProcSysCpu15min returns the ProcSysCpu15min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu15min() NiatelemetryProcSysCpu15min {
	if o == nil || o.ProcSysCpu15min.Get() == nil {
		var ret NiatelemetryProcSysCpu15min
		return ret
	}
	return *o.ProcSysCpu15min.Get()
}

// GetProcSysCpu15minOk returns a tuple with the ProcSysCpu15min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu15minOk() (*NiatelemetryProcSysCpu15min, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcSysCpu15min.Get(), o.ProcSysCpu15min.IsSet()
}

// HasProcSysCpu15min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysCpu15min() bool {
	if o != nil && o.ProcSysCpu15min.IsSet() {
		return true
	}

	return false
}

// SetProcSysCpu15min gets a reference to the given NullableNiatelemetryProcSysCpu15min and assigns it to the ProcSysCpu15min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu15min(v NiatelemetryProcSysCpu15min) {
	o.ProcSysCpu15min.Set(&v)
}

// SetProcSysCpu15minNil sets the value for ProcSysCpu15min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu15minNil() {
	o.ProcSysCpu15min.Set(nil)
}

// UnsetProcSysCpu15min ensures that no value is present for ProcSysCpu15min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysCpu15min() {
	o.ProcSysCpu15min.Unset()
}

// GetProcSysCpu5min returns the ProcSysCpu5min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu5min() NiatelemetryProcSysCpu5min {
	if o == nil || o.ProcSysCpu5min.Get() == nil {
		var ret NiatelemetryProcSysCpu5min
		return ret
	}
	return *o.ProcSysCpu5min.Get()
}

// GetProcSysCpu5minOk returns a tuple with the ProcSysCpu5min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu5minOk() (*NiatelemetryProcSysCpu5min, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcSysCpu5min.Get(), o.ProcSysCpu5min.IsSet()
}

// HasProcSysCpu5min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysCpu5min() bool {
	if o != nil && o.ProcSysCpu5min.IsSet() {
		return true
	}

	return false
}

// SetProcSysCpu5min gets a reference to the given NullableNiatelemetryProcSysCpu5min and assigns it to the ProcSysCpu5min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu5min(v NiatelemetryProcSysCpu5min) {
	o.ProcSysCpu5min.Set(&v)
}

// SetProcSysCpu5minNil sets the value for ProcSysCpu5min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu5minNil() {
	o.ProcSysCpu5min.Set(nil)
}

// UnsetProcSysCpu5min ensures that no value is present for ProcSysCpu5min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysCpu5min() {
	o.ProcSysCpu5min.Unset()
}

// GetProcSysMem15min returns the ProcSysMem15min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem15min() NiatelemetryProcSysMem15min {
	if o == nil || o.ProcSysMem15min.Get() == nil {
		var ret NiatelemetryProcSysMem15min
		return ret
	}
	return *o.ProcSysMem15min.Get()
}

// GetProcSysMem15minOk returns a tuple with the ProcSysMem15min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem15minOk() (*NiatelemetryProcSysMem15min, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcSysMem15min.Get(), o.ProcSysMem15min.IsSet()
}

// HasProcSysMem15min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysMem15min() bool {
	if o != nil && o.ProcSysMem15min.IsSet() {
		return true
	}

	return false
}

// SetProcSysMem15min gets a reference to the given NullableNiatelemetryProcSysMem15min and assigns it to the ProcSysMem15min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem15min(v NiatelemetryProcSysMem15min) {
	o.ProcSysMem15min.Set(&v)
}

// SetProcSysMem15minNil sets the value for ProcSysMem15min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem15minNil() {
	o.ProcSysMem15min.Set(nil)
}

// UnsetProcSysMem15min ensures that no value is present for ProcSysMem15min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysMem15min() {
	o.ProcSysMem15min.Unset()
}

// GetProcSysMem5min returns the ProcSysMem5min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem5min() NiatelemetryProcSysMem5min {
	if o == nil || o.ProcSysMem5min.Get() == nil {
		var ret NiatelemetryProcSysMem5min
		return ret
	}
	return *o.ProcSysMem5min.Get()
}

// GetProcSysMem5minOk returns a tuple with the ProcSysMem5min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem5minOk() (*NiatelemetryProcSysMem5min, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcSysMem5min.Get(), o.ProcSysMem5min.IsSet()
}

// HasProcSysMem5min returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysMem5min() bool {
	if o != nil && o.ProcSysMem5min.IsSet() {
		return true
	}

	return false
}

// SetProcSysMem5min gets a reference to the given NullableNiatelemetryProcSysMem5min and assigns it to the ProcSysMem5min field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem5min(v NiatelemetryProcSysMem5min) {
	o.ProcSysMem5min.Set(&v)
}

// SetProcSysMem5minNil sets the value for ProcSysMem5min to be an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem5minNil() {
	o.ProcSysMem5min.Set(nil)
}

// UnsetProcSysMem5min ensures that no value is present for ProcSysMem5min, not even an explicit nil
func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysMem5min() {
	o.ProcSysMem5min.Unset()
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSwitchDiskUtilization returns the SwitchDiskUtilization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryApicPerformanceDataAllOf) GetSwitchDiskUtilization() []NiatelemetrySwitchDiskUtilization {
	if o == nil {
		var ret []NiatelemetrySwitchDiskUtilization
		return ret
	}
	return o.SwitchDiskUtilization
}

// GetSwitchDiskUtilizationOk returns a tuple with the SwitchDiskUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryApicPerformanceDataAllOf) GetSwitchDiskUtilizationOk() ([]NiatelemetrySwitchDiskUtilization, bool) {
	if o == nil || o.SwitchDiskUtilization == nil {
		return nil, false
	}
	return o.SwitchDiskUtilization, true
}

// HasSwitchDiskUtilization returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasSwitchDiskUtilization() bool {
	if o != nil && o.SwitchDiskUtilization != nil {
		return true
	}

	return false
}

// SetSwitchDiskUtilization gets a reference to the given []NiatelemetrySwitchDiskUtilization and assigns it to the SwitchDiskUtilization field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetSwitchDiskUtilization(v []NiatelemetrySwitchDiskUtilization) {
	o.SwitchDiskUtilization = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicPerformanceDataAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicPerformanceDataAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicPerformanceDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DigitalOpticalMonitoring != nil {
		toSerialize["DigitalOpticalMonitoring"] = o.DigitalOpticalMonitoring
	}
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.EqptStorageFirmware.IsSet() {
		toSerialize["EqptStorageFirmware"] = o.EqptStorageFirmware.Get()
	}
	if o.EqptcapacityPolUsage5min.IsSet() {
		toSerialize["EqptcapacityPolUsage5min"] = o.EqptcapacityPolUsage5min.Get()
	}
	if o.EqptcapacityPrefixEntries15min.IsSet() {
		toSerialize["EqptcapacityPrefixEntries15min"] = o.EqptcapacityPrefixEntries15min.Get()
	}
	if o.EqptcapacityPrefixEntries5min.IsSet() {
		toSerialize["EqptcapacityPrefixEntries5min"] = o.EqptcapacityPrefixEntries5min.Get()
	}
	if o.NodeHealth != nil {
		toSerialize["NodeHealth"] = o.NodeHealth
	}
	if o.ProcSysCpu15min.IsSet() {
		toSerialize["ProcSysCpu15min"] = o.ProcSysCpu15min.Get()
	}
	if o.ProcSysCpu5min.IsSet() {
		toSerialize["ProcSysCpu5min"] = o.ProcSysCpu5min.Get()
	}
	if o.ProcSysMem15min.IsSet() {
		toSerialize["ProcSysMem15min"] = o.ProcSysMem15min.Get()
	}
	if o.ProcSysMem5min.IsSet() {
		toSerialize["ProcSysMem5min"] = o.ProcSysMem5min.Get()
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.SwitchDiskUtilization != nil {
		toSerialize["SwitchDiskUtilization"] = o.SwitchDiskUtilization
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryApicPerformanceDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicPerformanceDataAllOf := _NiatelemetryApicPerformanceDataAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicPerformanceDataAllOf); err == nil {
		*o = NiatelemetryApicPerformanceDataAllOf(varNiatelemetryApicPerformanceDataAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DigitalOpticalMonitoring")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "EqptStorageFirmware")
		delete(additionalProperties, "EqptcapacityPolUsage5min")
		delete(additionalProperties, "EqptcapacityPrefixEntries15min")
		delete(additionalProperties, "EqptcapacityPrefixEntries5min")
		delete(additionalProperties, "NodeHealth")
		delete(additionalProperties, "ProcSysCpu15min")
		delete(additionalProperties, "ProcSysCpu5min")
		delete(additionalProperties, "ProcSysMem15min")
		delete(additionalProperties, "ProcSysMem5min")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "SwitchDiskUtilization")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicPerformanceDataAllOf struct {
	value *NiatelemetryApicPerformanceDataAllOf
	isSet bool
}

func (v NullableNiatelemetryApicPerformanceDataAllOf) Get() *NiatelemetryApicPerformanceDataAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicPerformanceDataAllOf) Set(val *NiatelemetryApicPerformanceDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicPerformanceDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicPerformanceDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicPerformanceDataAllOf(val *NiatelemetryApicPerformanceDataAllOf) *NullableNiatelemetryApicPerformanceDataAllOf {
	return &NullableNiatelemetryApicPerformanceDataAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicPerformanceDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicPerformanceDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
