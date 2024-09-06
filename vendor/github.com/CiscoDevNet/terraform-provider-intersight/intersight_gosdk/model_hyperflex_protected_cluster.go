/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18369
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

// checks if the HyperflexProtectedCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexProtectedCluster{}

// HyperflexProtectedCluster Object for the protected clusters view.
type HyperflexProtectedCluster struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Version of the Hyperflex cluster.
	HxVersion *string `json:"HxVersion,omitempty"`
	// The version of hypervisor running on this cluster.
	HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
	// Name of the protected datastore.
	ProtectedDatastoreName *string `json:"ProtectedDatastoreName,omitempty"`
	// Number of VMs protected on this cluster.
	ProtectedVmsCount *int64 `json:"ProtectedVmsCount,omitempty"`
	// Name of the source cluster.
	SourceClusterName *string `json:"SourceClusterName,omitempty"`
	// Name of the target cluster.
	TargetClusterName *string `json:"TargetClusterName,omitempty"`
	// Name of the target datastore.
	TargetDatastoreName *string `json:"TargetDatastoreName,omitempty"`
	// Percent usage of the datastore.
	TargetDatastoreUtilization *float32                                                  `json:"TargetDatastoreUtilization,omitempty"`
	BackupPolicy               NullableHyperflexClusterBackupPolicyInventoryRelationship `json:"BackupPolicy,omitempty"`
	DatastoreStatistic         NullableHyperflexDatastoreStatisticRelationship           `json:"DatastoreStatistic,omitempty"`
	SrcCluster                 NullableHyperflexClusterRelationship                      `json:"SrcCluster,omitempty"`
	TgtCluster                 NullableHyperflexClusterRelationship                      `json:"TgtCluster,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _HyperflexProtectedCluster HyperflexProtectedCluster

// NewHyperflexProtectedCluster instantiates a new HyperflexProtectedCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexProtectedCluster(classId string, objectType string) *HyperflexProtectedCluster {
	this := HyperflexProtectedCluster{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexProtectedClusterWithDefaults instantiates a new HyperflexProtectedCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexProtectedClusterWithDefaults() *HyperflexProtectedCluster {
	this := HyperflexProtectedCluster{}
	var classId string = "hyperflex.ProtectedCluster"
	this.ClassId = classId
	var objectType string = "hyperflex.ProtectedCluster"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexProtectedCluster) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexProtectedCluster) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ProtectedCluster" of the ClassId field.
func (o *HyperflexProtectedCluster) GetDefaultClassId() interface{} {
	return "hyperflex.ProtectedCluster"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexProtectedCluster) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexProtectedCluster) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ProtectedCluster" of the ObjectType field.
func (o *HyperflexProtectedCluster) GetDefaultObjectType() interface{} {
	return "hyperflex.ProtectedCluster"
}

// GetHxVersion returns the HxVersion field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetHxVersion() string {
	if o == nil || IsNil(o.HxVersion) {
		var ret string
		return ret
	}
	return *o.HxVersion
}

// GetHxVersionOk returns a tuple with the HxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetHxVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HxVersion) {
		return nil, false
	}
	return o.HxVersion, true
}

// HasHxVersion returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasHxVersion() bool {
	if o != nil && !IsNil(o.HxVersion) {
		return true
	}

	return false
}

// SetHxVersion gets a reference to the given string and assigns it to the HxVersion field.
func (o *HyperflexProtectedCluster) SetHxVersion(v string) {
	o.HxVersion = &v
}

// GetHypervisorVersion returns the HypervisorVersion field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetHypervisorVersion() string {
	if o == nil || IsNil(o.HypervisorVersion) {
		var ret string
		return ret
	}
	return *o.HypervisorVersion
}

// GetHypervisorVersionOk returns a tuple with the HypervisorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetHypervisorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.HypervisorVersion) {
		return nil, false
	}
	return o.HypervisorVersion, true
}

// HasHypervisorVersion returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasHypervisorVersion() bool {
	if o != nil && !IsNil(o.HypervisorVersion) {
		return true
	}

	return false
}

// SetHypervisorVersion gets a reference to the given string and assigns it to the HypervisorVersion field.
func (o *HyperflexProtectedCluster) SetHypervisorVersion(v string) {
	o.HypervisorVersion = &v
}

// GetProtectedDatastoreName returns the ProtectedDatastoreName field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetProtectedDatastoreName() string {
	if o == nil || IsNil(o.ProtectedDatastoreName) {
		var ret string
		return ret
	}
	return *o.ProtectedDatastoreName
}

// GetProtectedDatastoreNameOk returns a tuple with the ProtectedDatastoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetProtectedDatastoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectedDatastoreName) {
		return nil, false
	}
	return o.ProtectedDatastoreName, true
}

// HasProtectedDatastoreName returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasProtectedDatastoreName() bool {
	if o != nil && !IsNil(o.ProtectedDatastoreName) {
		return true
	}

	return false
}

// SetProtectedDatastoreName gets a reference to the given string and assigns it to the ProtectedDatastoreName field.
func (o *HyperflexProtectedCluster) SetProtectedDatastoreName(v string) {
	o.ProtectedDatastoreName = &v
}

// GetProtectedVmsCount returns the ProtectedVmsCount field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetProtectedVmsCount() int64 {
	if o == nil || IsNil(o.ProtectedVmsCount) {
		var ret int64
		return ret
	}
	return *o.ProtectedVmsCount
}

// GetProtectedVmsCountOk returns a tuple with the ProtectedVmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetProtectedVmsCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ProtectedVmsCount) {
		return nil, false
	}
	return o.ProtectedVmsCount, true
}

// HasProtectedVmsCount returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasProtectedVmsCount() bool {
	if o != nil && !IsNil(o.ProtectedVmsCount) {
		return true
	}

	return false
}

// SetProtectedVmsCount gets a reference to the given int64 and assigns it to the ProtectedVmsCount field.
func (o *HyperflexProtectedCluster) SetProtectedVmsCount(v int64) {
	o.ProtectedVmsCount = &v
}

// GetSourceClusterName returns the SourceClusterName field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetSourceClusterName() string {
	if o == nil || IsNil(o.SourceClusterName) {
		var ret string
		return ret
	}
	return *o.SourceClusterName
}

// GetSourceClusterNameOk returns a tuple with the SourceClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetSourceClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceClusterName) {
		return nil, false
	}
	return o.SourceClusterName, true
}

// HasSourceClusterName returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasSourceClusterName() bool {
	if o != nil && !IsNil(o.SourceClusterName) {
		return true
	}

	return false
}

// SetSourceClusterName gets a reference to the given string and assigns it to the SourceClusterName field.
func (o *HyperflexProtectedCluster) SetSourceClusterName(v string) {
	o.SourceClusterName = &v
}

// GetTargetClusterName returns the TargetClusterName field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetTargetClusterName() string {
	if o == nil || IsNil(o.TargetClusterName) {
		var ret string
		return ret
	}
	return *o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetTargetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterName) {
		return nil, false
	}
	return o.TargetClusterName, true
}

// HasTargetClusterName returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasTargetClusterName() bool {
	if o != nil && !IsNil(o.TargetClusterName) {
		return true
	}

	return false
}

// SetTargetClusterName gets a reference to the given string and assigns it to the TargetClusterName field.
func (o *HyperflexProtectedCluster) SetTargetClusterName(v string) {
	o.TargetClusterName = &v
}

// GetTargetDatastoreName returns the TargetDatastoreName field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetTargetDatastoreName() string {
	if o == nil || IsNil(o.TargetDatastoreName) {
		var ret string
		return ret
	}
	return *o.TargetDatastoreName
}

// GetTargetDatastoreNameOk returns a tuple with the TargetDatastoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetTargetDatastoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetDatastoreName) {
		return nil, false
	}
	return o.TargetDatastoreName, true
}

// HasTargetDatastoreName returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasTargetDatastoreName() bool {
	if o != nil && !IsNil(o.TargetDatastoreName) {
		return true
	}

	return false
}

// SetTargetDatastoreName gets a reference to the given string and assigns it to the TargetDatastoreName field.
func (o *HyperflexProtectedCluster) SetTargetDatastoreName(v string) {
	o.TargetDatastoreName = &v
}

// GetTargetDatastoreUtilization returns the TargetDatastoreUtilization field value if set, zero value otherwise.
func (o *HyperflexProtectedCluster) GetTargetDatastoreUtilization() float32 {
	if o == nil || IsNil(o.TargetDatastoreUtilization) {
		var ret float32
		return ret
	}
	return *o.TargetDatastoreUtilization
}

// GetTargetDatastoreUtilizationOk returns a tuple with the TargetDatastoreUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexProtectedCluster) GetTargetDatastoreUtilizationOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetDatastoreUtilization) {
		return nil, false
	}
	return o.TargetDatastoreUtilization, true
}

// HasTargetDatastoreUtilization returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasTargetDatastoreUtilization() bool {
	if o != nil && !IsNil(o.TargetDatastoreUtilization) {
		return true
	}

	return false
}

// SetTargetDatastoreUtilization gets a reference to the given float32 and assigns it to the TargetDatastoreUtilization field.
func (o *HyperflexProtectedCluster) SetTargetDatastoreUtilization(v float32) {
	o.TargetDatastoreUtilization = &v
}

// GetBackupPolicy returns the BackupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectedCluster) GetBackupPolicy() HyperflexClusterBackupPolicyInventoryRelationship {
	if o == nil || IsNil(o.BackupPolicy.Get()) {
		var ret HyperflexClusterBackupPolicyInventoryRelationship
		return ret
	}
	return *o.BackupPolicy.Get()
}

// GetBackupPolicyOk returns a tuple with the BackupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectedCluster) GetBackupPolicyOk() (*HyperflexClusterBackupPolicyInventoryRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupPolicy.Get(), o.BackupPolicy.IsSet()
}

// HasBackupPolicy returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasBackupPolicy() bool {
	if o != nil && o.BackupPolicy.IsSet() {
		return true
	}

	return false
}

// SetBackupPolicy gets a reference to the given NullableHyperflexClusterBackupPolicyInventoryRelationship and assigns it to the BackupPolicy field.
func (o *HyperflexProtectedCluster) SetBackupPolicy(v HyperflexClusterBackupPolicyInventoryRelationship) {
	o.BackupPolicy.Set(&v)
}

// SetBackupPolicyNil sets the value for BackupPolicy to be an explicit nil
func (o *HyperflexProtectedCluster) SetBackupPolicyNil() {
	o.BackupPolicy.Set(nil)
}

// UnsetBackupPolicy ensures that no value is present for BackupPolicy, not even an explicit nil
func (o *HyperflexProtectedCluster) UnsetBackupPolicy() {
	o.BackupPolicy.Unset()
}

// GetDatastoreStatistic returns the DatastoreStatistic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectedCluster) GetDatastoreStatistic() HyperflexDatastoreStatisticRelationship {
	if o == nil || IsNil(o.DatastoreStatistic.Get()) {
		var ret HyperflexDatastoreStatisticRelationship
		return ret
	}
	return *o.DatastoreStatistic.Get()
}

// GetDatastoreStatisticOk returns a tuple with the DatastoreStatistic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectedCluster) GetDatastoreStatisticOk() (*HyperflexDatastoreStatisticRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatastoreStatistic.Get(), o.DatastoreStatistic.IsSet()
}

// HasDatastoreStatistic returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasDatastoreStatistic() bool {
	if o != nil && o.DatastoreStatistic.IsSet() {
		return true
	}

	return false
}

// SetDatastoreStatistic gets a reference to the given NullableHyperflexDatastoreStatisticRelationship and assigns it to the DatastoreStatistic field.
func (o *HyperflexProtectedCluster) SetDatastoreStatistic(v HyperflexDatastoreStatisticRelationship) {
	o.DatastoreStatistic.Set(&v)
}

// SetDatastoreStatisticNil sets the value for DatastoreStatistic to be an explicit nil
func (o *HyperflexProtectedCluster) SetDatastoreStatisticNil() {
	o.DatastoreStatistic.Set(nil)
}

// UnsetDatastoreStatistic ensures that no value is present for DatastoreStatistic, not even an explicit nil
func (o *HyperflexProtectedCluster) UnsetDatastoreStatistic() {
	o.DatastoreStatistic.Unset()
}

// GetSrcCluster returns the SrcCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectedCluster) GetSrcCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.SrcCluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.SrcCluster.Get()
}

// GetSrcClusterOk returns a tuple with the SrcCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectedCluster) GetSrcClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcCluster.Get(), o.SrcCluster.IsSet()
}

// HasSrcCluster returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasSrcCluster() bool {
	if o != nil && o.SrcCluster.IsSet() {
		return true
	}

	return false
}

// SetSrcCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the SrcCluster field.
func (o *HyperflexProtectedCluster) SetSrcCluster(v HyperflexClusterRelationship) {
	o.SrcCluster.Set(&v)
}

// SetSrcClusterNil sets the value for SrcCluster to be an explicit nil
func (o *HyperflexProtectedCluster) SetSrcClusterNil() {
	o.SrcCluster.Set(nil)
}

// UnsetSrcCluster ensures that no value is present for SrcCluster, not even an explicit nil
func (o *HyperflexProtectedCluster) UnsetSrcCluster() {
	o.SrcCluster.Unset()
}

// GetTgtCluster returns the TgtCluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexProtectedCluster) GetTgtCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.TgtCluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.TgtCluster.Get()
}

// GetTgtClusterOk returns a tuple with the TgtCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexProtectedCluster) GetTgtClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TgtCluster.Get(), o.TgtCluster.IsSet()
}

// HasTgtCluster returns a boolean if a field has been set.
func (o *HyperflexProtectedCluster) HasTgtCluster() bool {
	if o != nil && o.TgtCluster.IsSet() {
		return true
	}

	return false
}

// SetTgtCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the TgtCluster field.
func (o *HyperflexProtectedCluster) SetTgtCluster(v HyperflexClusterRelationship) {
	o.TgtCluster.Set(&v)
}

// SetTgtClusterNil sets the value for TgtCluster to be an explicit nil
func (o *HyperflexProtectedCluster) SetTgtClusterNil() {
	o.TgtCluster.Set(nil)
}

// UnsetTgtCluster ensures that no value is present for TgtCluster, not even an explicit nil
func (o *HyperflexProtectedCluster) UnsetTgtCluster() {
	o.TgtCluster.Unset()
}

func (o HyperflexProtectedCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexProtectedCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.HxVersion) {
		toSerialize["HxVersion"] = o.HxVersion
	}
	if !IsNil(o.HypervisorVersion) {
		toSerialize["HypervisorVersion"] = o.HypervisorVersion
	}
	if !IsNil(o.ProtectedDatastoreName) {
		toSerialize["ProtectedDatastoreName"] = o.ProtectedDatastoreName
	}
	if !IsNil(o.ProtectedVmsCount) {
		toSerialize["ProtectedVmsCount"] = o.ProtectedVmsCount
	}
	if !IsNil(o.SourceClusterName) {
		toSerialize["SourceClusterName"] = o.SourceClusterName
	}
	if !IsNil(o.TargetClusterName) {
		toSerialize["TargetClusterName"] = o.TargetClusterName
	}
	if !IsNil(o.TargetDatastoreName) {
		toSerialize["TargetDatastoreName"] = o.TargetDatastoreName
	}
	if !IsNil(o.TargetDatastoreUtilization) {
		toSerialize["TargetDatastoreUtilization"] = o.TargetDatastoreUtilization
	}
	if o.BackupPolicy.IsSet() {
		toSerialize["BackupPolicy"] = o.BackupPolicy.Get()
	}
	if o.DatastoreStatistic.IsSet() {
		toSerialize["DatastoreStatistic"] = o.DatastoreStatistic.Get()
	}
	if o.SrcCluster.IsSet() {
		toSerialize["SrcCluster"] = o.SrcCluster.Get()
	}
	if o.TgtCluster.IsSet() {
		toSerialize["TgtCluster"] = o.TgtCluster.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexProtectedCluster) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexProtectedClusterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Version of the Hyperflex cluster.
		HxVersion *string `json:"HxVersion,omitempty"`
		// The version of hypervisor running on this cluster.
		HypervisorVersion *string `json:"HypervisorVersion,omitempty"`
		// Name of the protected datastore.
		ProtectedDatastoreName *string `json:"ProtectedDatastoreName,omitempty"`
		// Number of VMs protected on this cluster.
		ProtectedVmsCount *int64 `json:"ProtectedVmsCount,omitempty"`
		// Name of the source cluster.
		SourceClusterName *string `json:"SourceClusterName,omitempty"`
		// Name of the target cluster.
		TargetClusterName *string `json:"TargetClusterName,omitempty"`
		// Name of the target datastore.
		TargetDatastoreName *string `json:"TargetDatastoreName,omitempty"`
		// Percent usage of the datastore.
		TargetDatastoreUtilization *float32                                                  `json:"TargetDatastoreUtilization,omitempty"`
		BackupPolicy               NullableHyperflexClusterBackupPolicyInventoryRelationship `json:"BackupPolicy,omitempty"`
		DatastoreStatistic         NullableHyperflexDatastoreStatisticRelationship           `json:"DatastoreStatistic,omitempty"`
		SrcCluster                 NullableHyperflexClusterRelationship                      `json:"SrcCluster,omitempty"`
		TgtCluster                 NullableHyperflexClusterRelationship                      `json:"TgtCluster,omitempty"`
	}

	varHyperflexProtectedClusterWithoutEmbeddedStruct := HyperflexProtectedClusterWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexProtectedClusterWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexProtectedCluster := _HyperflexProtectedCluster{}
		varHyperflexProtectedCluster.ClassId = varHyperflexProtectedClusterWithoutEmbeddedStruct.ClassId
		varHyperflexProtectedCluster.ObjectType = varHyperflexProtectedClusterWithoutEmbeddedStruct.ObjectType
		varHyperflexProtectedCluster.HxVersion = varHyperflexProtectedClusterWithoutEmbeddedStruct.HxVersion
		varHyperflexProtectedCluster.HypervisorVersion = varHyperflexProtectedClusterWithoutEmbeddedStruct.HypervisorVersion
		varHyperflexProtectedCluster.ProtectedDatastoreName = varHyperflexProtectedClusterWithoutEmbeddedStruct.ProtectedDatastoreName
		varHyperflexProtectedCluster.ProtectedVmsCount = varHyperflexProtectedClusterWithoutEmbeddedStruct.ProtectedVmsCount
		varHyperflexProtectedCluster.SourceClusterName = varHyperflexProtectedClusterWithoutEmbeddedStruct.SourceClusterName
		varHyperflexProtectedCluster.TargetClusterName = varHyperflexProtectedClusterWithoutEmbeddedStruct.TargetClusterName
		varHyperflexProtectedCluster.TargetDatastoreName = varHyperflexProtectedClusterWithoutEmbeddedStruct.TargetDatastoreName
		varHyperflexProtectedCluster.TargetDatastoreUtilization = varHyperflexProtectedClusterWithoutEmbeddedStruct.TargetDatastoreUtilization
		varHyperflexProtectedCluster.BackupPolicy = varHyperflexProtectedClusterWithoutEmbeddedStruct.BackupPolicy
		varHyperflexProtectedCluster.DatastoreStatistic = varHyperflexProtectedClusterWithoutEmbeddedStruct.DatastoreStatistic
		varHyperflexProtectedCluster.SrcCluster = varHyperflexProtectedClusterWithoutEmbeddedStruct.SrcCluster
		varHyperflexProtectedCluster.TgtCluster = varHyperflexProtectedClusterWithoutEmbeddedStruct.TgtCluster
		*o = HyperflexProtectedCluster(varHyperflexProtectedCluster)
	} else {
		return err
	}

	varHyperflexProtectedCluster := _HyperflexProtectedCluster{}

	err = json.Unmarshal(data, &varHyperflexProtectedCluster)
	if err == nil {
		o.MoBaseMo = varHyperflexProtectedCluster.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HxVersion")
		delete(additionalProperties, "HypervisorVersion")
		delete(additionalProperties, "ProtectedDatastoreName")
		delete(additionalProperties, "ProtectedVmsCount")
		delete(additionalProperties, "SourceClusterName")
		delete(additionalProperties, "TargetClusterName")
		delete(additionalProperties, "TargetDatastoreName")
		delete(additionalProperties, "TargetDatastoreUtilization")
		delete(additionalProperties, "BackupPolicy")
		delete(additionalProperties, "DatastoreStatistic")
		delete(additionalProperties, "SrcCluster")
		delete(additionalProperties, "TgtCluster")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableHyperflexProtectedCluster struct {
	value *HyperflexProtectedCluster
	isSet bool
}

func (v NullableHyperflexProtectedCluster) Get() *HyperflexProtectedCluster {
	return v.value
}

func (v *NullableHyperflexProtectedCluster) Set(val *HyperflexProtectedCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexProtectedCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexProtectedCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexProtectedCluster(val *HyperflexProtectedCluster) *NullableHyperflexProtectedCluster {
	return &NullableHyperflexProtectedCluster{value: val, isSet: true}
}

func (v NullableHyperflexProtectedCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexProtectedCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
