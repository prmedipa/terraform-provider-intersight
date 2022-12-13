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
	"reflect"
	"strings"
)

// VirtualizationVmwareCluster A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
type VirtualizationVmwareCluster struct {
	VirtualizationBaseCluster
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CPU over commitment associated with this cluster.
	CpuOverCommitment *int64 `json:"CpuOverCommitment,omitempty"`
	// Count of all datastores associated with this cluster.
	DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
	// Inventory path of the cluster.
	InventoryPath        *string                                     `json:"InventoryPath,omitempty"`
	Datacenter           *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareCluster VirtualizationVmwareCluster

// NewVirtualizationVmwareCluster instantiates a new VirtualizationVmwareCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareCluster(classId string, objectType string) *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	var hypervisorType string = "ESXi"
	this.HypervisorType = &hypervisorType
	return &this
}

// NewVirtualizationVmwareClusterWithDefaults instantiates a new VirtualizationVmwareCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareClusterWithDefaults() *VirtualizationVmwareCluster {
	this := VirtualizationVmwareCluster{}
	var classId string = "virtualization.VmwareCluster"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpuOverCommitment returns the CpuOverCommitment field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetCpuOverCommitment() int64 {
	if o == nil || o.CpuOverCommitment == nil {
		var ret int64
		return ret
	}
	return *o.CpuOverCommitment
}

// GetCpuOverCommitmentOk returns a tuple with the CpuOverCommitment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetCpuOverCommitmentOk() (*int64, bool) {
	if o == nil || o.CpuOverCommitment == nil {
		return nil, false
	}
	return o.CpuOverCommitment, true
}

// HasCpuOverCommitment returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasCpuOverCommitment() bool {
	if o != nil && o.CpuOverCommitment != nil {
		return true
	}

	return false
}

// SetCpuOverCommitment gets a reference to the given int64 and assigns it to the CpuOverCommitment field.
func (o *VirtualizationVmwareCluster) SetCpuOverCommitment(v int64) {
	o.CpuOverCommitment = &v
}

// GetDatastoreCount returns the DatastoreCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetDatastoreCount() int64 {
	if o == nil || o.DatastoreCount == nil {
		var ret int64
		return ret
	}
	return *o.DatastoreCount
}

// GetDatastoreCountOk returns a tuple with the DatastoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetDatastoreCountOk() (*int64, bool) {
	if o == nil || o.DatastoreCount == nil {
		return nil, false
	}
	return o.DatastoreCount, true
}

// HasDatastoreCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatastoreCount() bool {
	if o != nil && o.DatastoreCount != nil {
		return true
	}

	return false
}

// SetDatastoreCount gets a reference to the given int64 and assigns it to the DatastoreCount field.
func (o *VirtualizationVmwareCluster) SetDatastoreCount(v int64) {
	o.DatastoreCount = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetInventoryPath() string {
	if o == nil || o.InventoryPath == nil {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetInventoryPathOk() (*string, bool) {
	if o == nil || o.InventoryPath == nil {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasInventoryPath() bool {
	if o != nil && o.InventoryPath != nil {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareCluster) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || o.Datacenter == nil {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given VirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareCluster) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *VirtualizationVmwareCluster) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareCluster) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *VirtualizationVmwareCluster) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *VirtualizationVmwareCluster) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o VirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseCluster, errVirtualizationBaseCluster := json.Marshal(o.VirtualizationBaseCluster)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	errVirtualizationBaseCluster = json.Unmarshal([]byte(serializedVirtualizationBaseCluster), &toSerialize)
	if errVirtualizationBaseCluster != nil {
		return []byte{}, errVirtualizationBaseCluster
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CpuOverCommitment != nil {
		toSerialize["CpuOverCommitment"] = o.CpuOverCommitment
	}
	if o.DatastoreCount != nil {
		toSerialize["DatastoreCount"] = o.DatastoreCount
	}
	if o.InventoryPath != nil {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareCluster) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CPU over commitment associated with this cluster.
		CpuOverCommitment *int64 `json:"CpuOverCommitment,omitempty"`
		// Count of all datastores associated with this cluster.
		DatastoreCount *int64 `json:"DatastoreCount,omitempty"`
		// Inventory path of the cluster.
		InventoryPath    *string                                     `json:"InventoryPath,omitempty"`
		Datacenter       *VirtualizationVmwareDatacenterRelationship `json:"Datacenter,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship        `json:"RegisteredDevice,omitempty"`
	}

	varVirtualizationVmwareClusterWithoutEmbeddedStruct := VirtualizationVmwareClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareClusterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}
		varVirtualizationVmwareCluster.ClassId = varVirtualizationVmwareClusterWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareCluster.ObjectType = varVirtualizationVmwareClusterWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareCluster.CpuOverCommitment = varVirtualizationVmwareClusterWithoutEmbeddedStruct.CpuOverCommitment
		varVirtualizationVmwareCluster.DatastoreCount = varVirtualizationVmwareClusterWithoutEmbeddedStruct.DatastoreCount
		varVirtualizationVmwareCluster.InventoryPath = varVirtualizationVmwareClusterWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareCluster.Datacenter = varVirtualizationVmwareClusterWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareCluster.RegisteredDevice = varVirtualizationVmwareClusterWithoutEmbeddedStruct.RegisteredDevice
		*o = VirtualizationVmwareCluster(varVirtualizationVmwareCluster)
	} else {
		return err
	}

	varVirtualizationVmwareCluster := _VirtualizationVmwareCluster{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareCluster)
	if err == nil {
		o.VirtualizationBaseCluster = varVirtualizationVmwareCluster.VirtualizationBaseCluster
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CpuOverCommitment")
		delete(additionalProperties, "DatastoreCount")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectVirtualizationBaseCluster := reflect.ValueOf(o.VirtualizationBaseCluster)
		for i := 0; i < reflectVirtualizationBaseCluster.Type().NumField(); i++ {
			t := reflectVirtualizationBaseCluster.Type().Field(i)

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

type NullableVirtualizationVmwareCluster struct {
	value *VirtualizationVmwareCluster
	isSet bool
}

func (v NullableVirtualizationVmwareCluster) Get() *VirtualizationVmwareCluster {
	return v.value
}

func (v *NullableVirtualizationVmwareCluster) Set(val *VirtualizationVmwareCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareCluster(val *VirtualizationVmwareCluster) *NullableVirtualizationVmwareCluster {
	return &NullableVirtualizationVmwareCluster{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
