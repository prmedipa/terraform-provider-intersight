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

// VirtualizationIweVirtualMachineAllOf Definition of the list of properties defined in 'virtualization.IweVirtualMachine', excluding properties defined in parent classes.
type VirtualizationIweVirtualMachineAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string          `json:"ObjectType"`
	AffinitySelectors []InfraMetaData `json:"AffinitySelectors,omitempty"`
	// Denotes age or life time of the VM in nano seconds.
	Age                   *string                `json:"Age,omitempty"`
	AntiAffinitySelectors []InfraMetaData        `json:"AntiAffinitySelectors,omitempty"`
	Disks                 []VirtualizationVmDisk `json:"Disks,omitempty"`
	// Reason of the failure when VM is in failed state.
	FailureReason *string                     `json:"FailureReason,omitempty"`
	Interfaces    []VirtualizationVmInterface `json:"Interfaces,omitempty"`
	Labels        []InfraMetaData             `json:"Labels,omitempty"`
	// Number network interfaces associated with the virtual machine.
	NetworkCount *int64 `json:"NetworkCount,omitempty"`
	// Denotes the VM start timestamp.
	StartTime *string `json:"StartTime,omitempty"`
	// Status of virtual machine. * `Unknown` - Virtual machine state is not available. * `Running` - Virtual machine is running normally. * `Stopped` - Virtual machine has been stopped. * `WaitForLaunch` - Virtual machine is wating to be launched. * `Paused` - Virtual machine is currently paused. * `ImportInProgress` - Virtual machine image is being imported into the platform. * `ImportFailed` - Virtual machine image import operation failed. * `DiskCloneInProgress` - Disk clone operation for the virtual machine is in progress. * `DiskCloneFailed` - Disk clone operation for the virtual machine failed. * `DiskCreateInProgress` - Disk create operation for the virtual machine is in progress. * `DiskCreateFailed` - Disk create operation for the virtual machine failed. * `Processing` - Virtual machine is being created. * `UnSchedulable` - Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications. * `Failed` - Some virtual machine operation has failed. More information is available as part of the results of the operation. * `Warning` - CPU/Memory utilisation for the virtual machine has crossed the threshold value. * `` - Virtual machine status is not available.
	Status               *string                               `json:"Status,omitempty"`
	Cluster              *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	Host                 *VirtualizationIweHostRelationship    `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweVirtualMachineAllOf VirtualizationIweVirtualMachineAllOf

// NewVirtualizationIweVirtualMachineAllOf instantiates a new VirtualizationIweVirtualMachineAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweVirtualMachineAllOf(classId string, objectType string) *VirtualizationIweVirtualMachineAllOf {
	this := VirtualizationIweVirtualMachineAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationIweVirtualMachineAllOfWithDefaults instantiates a new VirtualizationIweVirtualMachineAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweVirtualMachineAllOfWithDefaults() *VirtualizationIweVirtualMachineAllOf {
	this := VirtualizationIweVirtualMachineAllOf{}
	var classId string = "virtualization.IweVirtualMachine"
	this.ClassId = classId
	var objectType string = "virtualization.IweVirtualMachine"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweVirtualMachineAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweVirtualMachineAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweVirtualMachineAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweVirtualMachineAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAffinitySelectors returns the AffinitySelectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachineAllOf) GetAffinitySelectors() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.AffinitySelectors
}

// GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachineAllOf) GetAffinitySelectorsOk() ([]InfraMetaData, bool) {
	if o == nil || o.AffinitySelectors == nil {
		return nil, false
	}
	return o.AffinitySelectors, true
}

// HasAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasAffinitySelectors() bool {
	if o != nil && o.AffinitySelectors != nil {
		return true
	}

	return false
}

// SetAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AffinitySelectors field.
func (o *VirtualizationIweVirtualMachineAllOf) SetAffinitySelectors(v []InfraMetaData) {
	o.AffinitySelectors = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetAge() string {
	if o == nil || o.Age == nil {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetAgeOk() (*string, bool) {
	if o == nil || o.Age == nil {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasAge() bool {
	if o != nil && o.Age != nil {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *VirtualizationIweVirtualMachineAllOf) SetAge(v string) {
	o.Age = &v
}

// GetAntiAffinitySelectors returns the AntiAffinitySelectors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachineAllOf) GetAntiAffinitySelectors() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.AntiAffinitySelectors
}

// GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachineAllOf) GetAntiAffinitySelectorsOk() ([]InfraMetaData, bool) {
	if o == nil || o.AntiAffinitySelectors == nil {
		return nil, false
	}
	return o.AntiAffinitySelectors, true
}

// HasAntiAffinitySelectors returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasAntiAffinitySelectors() bool {
	if o != nil && o.AntiAffinitySelectors != nil {
		return true
	}

	return false
}

// SetAntiAffinitySelectors gets a reference to the given []InfraMetaData and assigns it to the AntiAffinitySelectors field.
func (o *VirtualizationIweVirtualMachineAllOf) SetAntiAffinitySelectors(v []InfraMetaData) {
	o.AntiAffinitySelectors = v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachineAllOf) GetDisks() []VirtualizationVmDisk {
	if o == nil {
		var ret []VirtualizationVmDisk
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachineAllOf) GetDisksOk() ([]VirtualizationVmDisk, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []VirtualizationVmDisk and assigns it to the Disks field.
func (o *VirtualizationIweVirtualMachineAllOf) SetDisks(v []VirtualizationVmDisk) {
	o.Disks = v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *VirtualizationIweVirtualMachineAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachineAllOf) GetInterfaces() []VirtualizationVmInterface {
	if o == nil {
		var ret []VirtualizationVmInterface
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachineAllOf) GetInterfacesOk() ([]VirtualizationVmInterface, bool) {
	if o == nil || o.Interfaces == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasInterfaces() bool {
	if o != nil && o.Interfaces != nil {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []VirtualizationVmInterface and assigns it to the Interfaces field.
func (o *VirtualizationIweVirtualMachineAllOf) SetInterfaces(v []VirtualizationVmInterface) {
	o.Interfaces = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweVirtualMachineAllOf) GetLabels() []InfraMetaData {
	if o == nil {
		var ret []InfraMetaData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweVirtualMachineAllOf) GetLabelsOk() ([]InfraMetaData, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []InfraMetaData and assigns it to the Labels field.
func (o *VirtualizationIweVirtualMachineAllOf) SetLabels(v []InfraMetaData) {
	o.Labels = v
}

// GetNetworkCount returns the NetworkCount field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetNetworkCount() int64 {
	if o == nil || o.NetworkCount == nil {
		var ret int64
		return ret
	}
	return *o.NetworkCount
}

// GetNetworkCountOk returns a tuple with the NetworkCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetNetworkCountOk() (*int64, bool) {
	if o == nil || o.NetworkCount == nil {
		return nil, false
	}
	return o.NetworkCount, true
}

// HasNetworkCount returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasNetworkCount() bool {
	if o != nil && o.NetworkCount != nil {
		return true
	}

	return false
}

// SetNetworkCount gets a reference to the given int64 and assigns it to the NetworkCount field.
func (o *VirtualizationIweVirtualMachineAllOf) SetNetworkCount(v int64) {
	o.NetworkCount = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *VirtualizationIweVirtualMachineAllOf) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationIweVirtualMachineAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweVirtualMachineAllOf) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VirtualizationIweVirtualMachineAllOf) GetHost() VirtualizationIweHostRelationship {
	if o == nil || o.Host == nil {
		var ret VirtualizationIweHostRelationship
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweVirtualMachineAllOf) GetHostOk() (*VirtualizationIweHostRelationship, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationIweVirtualMachineAllOf) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given VirtualizationIweHostRelationship and assigns it to the Host field.
func (o *VirtualizationIweVirtualMachineAllOf) SetHost(v VirtualizationIweHostRelationship) {
	o.Host = &v
}

func (o VirtualizationIweVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AffinitySelectors != nil {
		toSerialize["AffinitySelectors"] = o.AffinitySelectors
	}
	if o.Age != nil {
		toSerialize["Age"] = o.Age
	}
	if o.AntiAffinitySelectors != nil {
		toSerialize["AntiAffinitySelectors"] = o.AntiAffinitySelectors
	}
	if o.Disks != nil {
		toSerialize["Disks"] = o.Disks
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.Interfaces != nil {
		toSerialize["Interfaces"] = o.Interfaces
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.NetworkCount != nil {
		toSerialize["NetworkCount"] = o.NetworkCount
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.Host != nil {
		toSerialize["Host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweVirtualMachineAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVirtualizationIweVirtualMachineAllOf := _VirtualizationIweVirtualMachineAllOf{}

	if err = json.Unmarshal(bytes, &varVirtualizationIweVirtualMachineAllOf); err == nil {
		*o = VirtualizationIweVirtualMachineAllOf(varVirtualizationIweVirtualMachineAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AffinitySelectors")
		delete(additionalProperties, "Age")
		delete(additionalProperties, "AntiAffinitySelectors")
		delete(additionalProperties, "Disks")
		delete(additionalProperties, "FailureReason")
		delete(additionalProperties, "Interfaces")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "NetworkCount")
		delete(additionalProperties, "StartTime")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Host")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualizationIweVirtualMachineAllOf struct {
	value *VirtualizationIweVirtualMachineAllOf
	isSet bool
}

func (v NullableVirtualizationIweVirtualMachineAllOf) Get() *VirtualizationIweVirtualMachineAllOf {
	return v.value
}

func (v *NullableVirtualizationIweVirtualMachineAllOf) Set(val *VirtualizationIweVirtualMachineAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweVirtualMachineAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweVirtualMachineAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweVirtualMachineAllOf(val *VirtualizationIweVirtualMachineAllOf) *NullableVirtualizationIweVirtualMachineAllOf {
	return &NullableVirtualizationIweVirtualMachineAllOf{value: val, isSet: true}
}

func (v NullableVirtualizationIweVirtualMachineAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweVirtualMachineAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
