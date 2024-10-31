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

// checks if the VirtualizationVmwareDatastore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwareDatastore{}

// VirtualizationVmwareDatastore The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
type VirtualizationVmwareDatastore struct {
	VirtualizationBaseDatastore
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Shows if this datastore is accessible.
	Accessible *bool `json:"Accessible,omitempty"`
	// Inventory path of the Datastore.
	InventoryPath *string `json:"InventoryPath,omitempty"`
	// Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.
	MaintenanceMode *bool `json:"MaintenanceMode,omitempty"`
	// Indicates if this datastore is connected to multiple hosts.
	MultipleHostAccess *bool `json:"MultipleHostAccess,omitempty"`
	// Datastore health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
	Status *string `json:"Status,omitempty"`
	// Indicates if this datastore supports thin provisioning for files.
	ThinProvisioningSupported *bool `json:"ThinProvisioningSupported,omitempty"`
	// Space uncommitted in this datastore in bytes.
	UnCommitted *int64 `json:"UnCommitted,omitempty"`
	// The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').
	Url *string `json:"Url,omitempty"`
	// Number of virtual machine templates relying on (using) this datastore.
	VmTemplateCount *int64                                          `json:"VmTemplateCount,omitempty"`
	Cluster         NullableVirtualizationVmwareClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to virtualizationVmwareCluster resources.
	Clusters         []VirtualizationVmwareClusterRelationship                `json:"Clusters,omitempty"`
	Datacenter       NullableVirtualizationVmwareDatacenterRelationship       `json:"Datacenter,omitempty"`
	DatastoreCluster NullableVirtualizationVmwareDatastoreClusterRelationship `json:"DatastoreCluster,omitempty"`
	// An array of relationships to virtualizationVmwareHost resources.
	Hosts                []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareDatastore VirtualizationVmwareDatastore

// NewVirtualizationVmwareDatastore instantiates a new VirtualizationVmwareDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareDatastore(classId string, objectType string) *VirtualizationVmwareDatastore {
	this := VirtualizationVmwareDatastore{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "Unknown"
	this.Type = &type_
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// NewVirtualizationVmwareDatastoreWithDefaults instantiates a new VirtualizationVmwareDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareDatastoreWithDefaults() *VirtualizationVmwareDatastore {
	this := VirtualizationVmwareDatastore{}
	var classId string = "virtualization.VmwareDatastore"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatastore"
	this.ObjectType = objectType
	var status string = "Unknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareDatastore) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareDatastore) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareDatastore" of the ClassId field.
func (o *VirtualizationVmwareDatastore) GetDefaultClassId() interface{} {
	return "virtualization.VmwareDatastore"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareDatastore) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareDatastore) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareDatastore" of the ObjectType field.
func (o *VirtualizationVmwareDatastore) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareDatastore"
}

// GetAccessible returns the Accessible field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetAccessible() bool {
	if o == nil || IsNil(o.Accessible) {
		var ret bool
		return ret
	}
	return *o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Accessible) {
		return nil, false
	}
	return o.Accessible, true
}

// HasAccessible returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasAccessible() bool {
	if o != nil && !IsNil(o.Accessible) {
		return true
	}

	return false
}

// SetAccessible gets a reference to the given bool and assigns it to the Accessible field.
func (o *VirtualizationVmwareDatastore) SetAccessible(v bool) {
	o.Accessible = &v
}

// GetInventoryPath returns the InventoryPath field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetInventoryPath() string {
	if o == nil || IsNil(o.InventoryPath) {
		var ret string
		return ret
	}
	return *o.InventoryPath
}

// GetInventoryPathOk returns a tuple with the InventoryPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetInventoryPathOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryPath) {
		return nil, false
	}
	return o.InventoryPath, true
}

// HasInventoryPath returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasInventoryPath() bool {
	if o != nil && !IsNil(o.InventoryPath) {
		return true
	}

	return false
}

// SetInventoryPath gets a reference to the given string and assigns it to the InventoryPath field.
func (o *VirtualizationVmwareDatastore) SetInventoryPath(v string) {
	o.InventoryPath = &v
}

// GetMaintenanceMode returns the MaintenanceMode field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetMaintenanceMode() bool {
	if o == nil || IsNil(o.MaintenanceMode) {
		var ret bool
		return ret
	}
	return *o.MaintenanceMode
}

// GetMaintenanceModeOk returns a tuple with the MaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetMaintenanceModeOk() (*bool, bool) {
	if o == nil || IsNil(o.MaintenanceMode) {
		return nil, false
	}
	return o.MaintenanceMode, true
}

// HasMaintenanceMode returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasMaintenanceMode() bool {
	if o != nil && !IsNil(o.MaintenanceMode) {
		return true
	}

	return false
}

// SetMaintenanceMode gets a reference to the given bool and assigns it to the MaintenanceMode field.
func (o *VirtualizationVmwareDatastore) SetMaintenanceMode(v bool) {
	o.MaintenanceMode = &v
}

// GetMultipleHostAccess returns the MultipleHostAccess field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetMultipleHostAccess() bool {
	if o == nil || IsNil(o.MultipleHostAccess) {
		var ret bool
		return ret
	}
	return *o.MultipleHostAccess
}

// GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetMultipleHostAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.MultipleHostAccess) {
		return nil, false
	}
	return o.MultipleHostAccess, true
}

// HasMultipleHostAccess returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasMultipleHostAccess() bool {
	if o != nil && !IsNil(o.MultipleHostAccess) {
		return true
	}

	return false
}

// SetMultipleHostAccess gets a reference to the given bool and assigns it to the MultipleHostAccess field.
func (o *VirtualizationVmwareDatastore) SetMultipleHostAccess(v bool) {
	o.MultipleHostAccess = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VirtualizationVmwareDatastore) SetStatus(v string) {
	o.Status = &v
}

// GetThinProvisioningSupported returns the ThinProvisioningSupported field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetThinProvisioningSupported() bool {
	if o == nil || IsNil(o.ThinProvisioningSupported) {
		var ret bool
		return ret
	}
	return *o.ThinProvisioningSupported
}

// GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetThinProvisioningSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.ThinProvisioningSupported) {
		return nil, false
	}
	return o.ThinProvisioningSupported, true
}

// HasThinProvisioningSupported returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasThinProvisioningSupported() bool {
	if o != nil && !IsNil(o.ThinProvisioningSupported) {
		return true
	}

	return false
}

// SetThinProvisioningSupported gets a reference to the given bool and assigns it to the ThinProvisioningSupported field.
func (o *VirtualizationVmwareDatastore) SetThinProvisioningSupported(v bool) {
	o.ThinProvisioningSupported = &v
}

// GetUnCommitted returns the UnCommitted field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetUnCommitted() int64 {
	if o == nil || IsNil(o.UnCommitted) {
		var ret int64
		return ret
	}
	return *o.UnCommitted
}

// GetUnCommittedOk returns a tuple with the UnCommitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetUnCommittedOk() (*int64, bool) {
	if o == nil || IsNil(o.UnCommitted) {
		return nil, false
	}
	return o.UnCommitted, true
}

// HasUnCommitted returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasUnCommitted() bool {
	if o != nil && !IsNil(o.UnCommitted) {
		return true
	}

	return false
}

// SetUnCommitted gets a reference to the given int64 and assigns it to the UnCommitted field.
func (o *VirtualizationVmwareDatastore) SetUnCommitted(v int64) {
	o.UnCommitted = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VirtualizationVmwareDatastore) SetUrl(v string) {
	o.Url = &v
}

// GetVmTemplateCount returns the VmTemplateCount field value if set, zero value otherwise.
func (o *VirtualizationVmwareDatastore) GetVmTemplateCount() int64 {
	if o == nil || IsNil(o.VmTemplateCount) {
		var ret int64
		return ret
	}
	return *o.VmTemplateCount
}

// GetVmTemplateCountOk returns a tuple with the VmTemplateCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareDatastore) GetVmTemplateCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VmTemplateCount) {
		return nil, false
	}
	return o.VmTemplateCount, true
}

// HasVmTemplateCount returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasVmTemplateCount() bool {
	if o != nil && !IsNil(o.VmTemplateCount) {
		return true
	}

	return false
}

// SetVmTemplateCount gets a reference to the given int64 and assigns it to the VmTemplateCount field.
func (o *VirtualizationVmwareDatastore) SetVmTemplateCount(v int64) {
	o.VmTemplateCount = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastore) GetCluster() VirtualizationVmwareClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret VirtualizationVmwareClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastore) GetClusterOk() (*VirtualizationVmwareClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableVirtualizationVmwareClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationVmwareDatastore) SetCluster(v VirtualizationVmwareClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *VirtualizationVmwareDatastore) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *VirtualizationVmwareDatastore) UnsetCluster() {
	o.Cluster.Unset()
}

// GetClusters returns the Clusters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastore) GetClusters() []VirtualizationVmwareClusterRelationship {
	if o == nil {
		var ret []VirtualizationVmwareClusterRelationship
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastore) GetClustersOk() ([]VirtualizationVmwareClusterRelationship, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []VirtualizationVmwareClusterRelationship and assigns it to the Clusters field.
func (o *VirtualizationVmwareDatastore) SetClusters(v []VirtualizationVmwareClusterRelationship) {
	o.Clusters = v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastore) GetDatacenter() VirtualizationVmwareDatacenterRelationship {
	if o == nil || IsNil(o.Datacenter.Get()) {
		var ret VirtualizationVmwareDatacenterRelationship
		return ret
	}
	return *o.Datacenter.Get()
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastore) GetDatacenterOk() (*VirtualizationVmwareDatacenterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Datacenter.Get(), o.Datacenter.IsSet()
}

// HasDatacenter returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasDatacenter() bool {
	if o != nil && o.Datacenter.IsSet() {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given NullableVirtualizationVmwareDatacenterRelationship and assigns it to the Datacenter field.
func (o *VirtualizationVmwareDatastore) SetDatacenter(v VirtualizationVmwareDatacenterRelationship) {
	o.Datacenter.Set(&v)
}

// SetDatacenterNil sets the value for Datacenter to be an explicit nil
func (o *VirtualizationVmwareDatastore) SetDatacenterNil() {
	o.Datacenter.Set(nil)
}

// UnsetDatacenter ensures that no value is present for Datacenter, not even an explicit nil
func (o *VirtualizationVmwareDatastore) UnsetDatacenter() {
	o.Datacenter.Unset()
}

// GetDatastoreCluster returns the DatastoreCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastore) GetDatastoreCluster() VirtualizationVmwareDatastoreClusterRelationship {
	if o == nil || IsNil(o.DatastoreCluster.Get()) {
		var ret VirtualizationVmwareDatastoreClusterRelationship
		return ret
	}
	return *o.DatastoreCluster.Get()
}

// GetDatastoreClusterOk returns a tuple with the DatastoreCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastore) GetDatastoreClusterOk() (*VirtualizationVmwareDatastoreClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatastoreCluster.Get(), o.DatastoreCluster.IsSet()
}

// HasDatastoreCluster returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasDatastoreCluster() bool {
	if o != nil && o.DatastoreCluster.IsSet() {
		return true
	}

	return false
}

// SetDatastoreCluster gets a reference to the given NullableVirtualizationVmwareDatastoreClusterRelationship and assigns it to the DatastoreCluster field.
func (o *VirtualizationVmwareDatastore) SetDatastoreCluster(v VirtualizationVmwareDatastoreClusterRelationship) {
	o.DatastoreCluster.Set(&v)
}

// SetDatastoreClusterNil sets the value for DatastoreCluster to be an explicit nil
func (o *VirtualizationVmwareDatastore) SetDatastoreClusterNil() {
	o.DatastoreCluster.Set(nil)
}

// UnsetDatastoreCluster ensures that no value is present for DatastoreCluster, not even an explicit nil
func (o *VirtualizationVmwareDatastore) UnsetDatastoreCluster() {
	o.DatastoreCluster.Unset()
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwareDatastore) GetHosts() []VirtualizationVmwareHostRelationship {
	if o == nil {
		var ret []VirtualizationVmwareHostRelationship
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwareDatastore) GetHostsOk() ([]VirtualizationVmwareHostRelationship, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualizationVmwareDatastore) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VirtualizationVmwareHostRelationship and assigns it to the Hosts field.
func (o *VirtualizationVmwareDatastore) SetHosts(v []VirtualizationVmwareHostRelationship) {
	o.Hosts = v
}

func (o VirtualizationVmwareDatastore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwareDatastore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDatastore, errVirtualizationBaseDatastore := json.Marshal(o.VirtualizationBaseDatastore)
	if errVirtualizationBaseDatastore != nil {
		return map[string]interface{}{}, errVirtualizationBaseDatastore
	}
	errVirtualizationBaseDatastore = json.Unmarshal([]byte(serializedVirtualizationBaseDatastore), &toSerialize)
	if errVirtualizationBaseDatastore != nil {
		return map[string]interface{}{}, errVirtualizationBaseDatastore
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Accessible) {
		toSerialize["Accessible"] = o.Accessible
	}
	if !IsNil(o.InventoryPath) {
		toSerialize["InventoryPath"] = o.InventoryPath
	}
	if !IsNil(o.MaintenanceMode) {
		toSerialize["MaintenanceMode"] = o.MaintenanceMode
	}
	if !IsNil(o.MultipleHostAccess) {
		toSerialize["MultipleHostAccess"] = o.MultipleHostAccess
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.ThinProvisioningSupported) {
		toSerialize["ThinProvisioningSupported"] = o.ThinProvisioningSupported
	}
	if !IsNil(o.UnCommitted) {
		toSerialize["UnCommitted"] = o.UnCommitted
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.VmTemplateCount) {
		toSerialize["VmTemplateCount"] = o.VmTemplateCount
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}
	if o.Clusters != nil {
		toSerialize["Clusters"] = o.Clusters
	}
	if o.Datacenter.IsSet() {
		toSerialize["Datacenter"] = o.Datacenter.Get()
	}
	if o.DatastoreCluster.IsSet() {
		toSerialize["DatastoreCluster"] = o.DatastoreCluster.Get()
	}
	if o.Hosts != nil {
		toSerialize["Hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwareDatastore) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwareDatastoreWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Shows if this datastore is accessible.
		Accessible *bool `json:"Accessible,omitempty"`
		// Inventory path of the Datastore.
		InventoryPath *string `json:"InventoryPath,omitempty"`
		// Indicates if the datastore is in maintenance mode. Will be set to True, when in maintenance mode.
		MaintenanceMode *bool `json:"MaintenanceMode,omitempty"`
		// Indicates if this datastore is connected to multiple hosts.
		MultipleHostAccess *bool `json:"MultipleHostAccess,omitempty"`
		// Datastore health status, as reported by the hypervisor platform. * `Unknown` - Entity status is unknown. * `Degraded` - State is degraded, and might impact normal operation of the entity. * `Critical` - Entity is in a critical state, impacting operations. * `Ok` - Entity status is in a stable state, operating normally.
		Status *string `json:"Status,omitempty"`
		// Indicates if this datastore supports thin provisioning for files.
		ThinProvisioningSupported *bool `json:"ThinProvisioningSupported,omitempty"`
		// Space uncommitted in this datastore in bytes.
		UnCommitted *int64 `json:"UnCommitted,omitempty"`
		// The URL to access this datastore (example - 'ds:///vmfs/volumes/562a4e8a-0eeb5372-dd61-78baf9cb9afa/').
		Url *string `json:"Url,omitempty"`
		// Number of virtual machine templates relying on (using) this datastore.
		VmTemplateCount *int64                                          `json:"VmTemplateCount,omitempty"`
		Cluster         NullableVirtualizationVmwareClusterRelationship `json:"Cluster,omitempty"`
		// An array of relationships to virtualizationVmwareCluster resources.
		Clusters         []VirtualizationVmwareClusterRelationship                `json:"Clusters,omitempty"`
		Datacenter       NullableVirtualizationVmwareDatacenterRelationship       `json:"Datacenter,omitempty"`
		DatastoreCluster NullableVirtualizationVmwareDatastoreClusterRelationship `json:"DatastoreCluster,omitempty"`
		// An array of relationships to virtualizationVmwareHost resources.
		Hosts []VirtualizationVmwareHostRelationship `json:"Hosts,omitempty"`
	}

	varVirtualizationVmwareDatastoreWithoutEmbeddedStruct := VirtualizationVmwareDatastoreWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwareDatastoreWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareDatastore := _VirtualizationVmwareDatastore{}
		varVirtualizationVmwareDatastore.ClassId = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareDatastore.ObjectType = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareDatastore.Accessible = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Accessible
		varVirtualizationVmwareDatastore.InventoryPath = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.InventoryPath
		varVirtualizationVmwareDatastore.MaintenanceMode = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.MaintenanceMode
		varVirtualizationVmwareDatastore.MultipleHostAccess = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.MultipleHostAccess
		varVirtualizationVmwareDatastore.Status = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Status
		varVirtualizationVmwareDatastore.ThinProvisioningSupported = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.ThinProvisioningSupported
		varVirtualizationVmwareDatastore.UnCommitted = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.UnCommitted
		varVirtualizationVmwareDatastore.Url = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Url
		varVirtualizationVmwareDatastore.VmTemplateCount = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.VmTemplateCount
		varVirtualizationVmwareDatastore.Cluster = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Cluster
		varVirtualizationVmwareDatastore.Clusters = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Clusters
		varVirtualizationVmwareDatastore.Datacenter = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Datacenter
		varVirtualizationVmwareDatastore.DatastoreCluster = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.DatastoreCluster
		varVirtualizationVmwareDatastore.Hosts = varVirtualizationVmwareDatastoreWithoutEmbeddedStruct.Hosts
		*o = VirtualizationVmwareDatastore(varVirtualizationVmwareDatastore)
	} else {
		return err
	}

	varVirtualizationVmwareDatastore := _VirtualizationVmwareDatastore{}

	err = json.Unmarshal(data, &varVirtualizationVmwareDatastore)
	if err == nil {
		o.VirtualizationBaseDatastore = varVirtualizationVmwareDatastore.VirtualizationBaseDatastore
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accessible")
		delete(additionalProperties, "InventoryPath")
		delete(additionalProperties, "MaintenanceMode")
		delete(additionalProperties, "MultipleHostAccess")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "ThinProvisioningSupported")
		delete(additionalProperties, "UnCommitted")
		delete(additionalProperties, "Url")
		delete(additionalProperties, "VmTemplateCount")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Clusters")
		delete(additionalProperties, "Datacenter")
		delete(additionalProperties, "DatastoreCluster")
		delete(additionalProperties, "Hosts")

		// remove fields from embedded structs
		reflectVirtualizationBaseDatastore := reflect.ValueOf(o.VirtualizationBaseDatastore)
		for i := 0; i < reflectVirtualizationBaseDatastore.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDatastore.Type().Field(i)

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

type NullableVirtualizationVmwareDatastore struct {
	value *VirtualizationVmwareDatastore
	isSet bool
}

func (v NullableVirtualizationVmwareDatastore) Get() *VirtualizationVmwareDatastore {
	return v.value
}

func (v *NullableVirtualizationVmwareDatastore) Set(val *VirtualizationVmwareDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareDatastore(val *VirtualizationVmwareDatastore) *NullableVirtualizationVmwareDatastore {
	return &NullableVirtualizationVmwareDatastore{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
