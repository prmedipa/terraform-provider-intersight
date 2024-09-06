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

// checks if the StorageStoragePolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageStoragePolicy{}

// StorageStoragePolicy The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The security of drives can be enabled through this policy using remote keys from a KMIP server or Manually configured keys.
type StorageStoragePolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// All unconfigured drives will move to the selected state on deployment. Newly inserted drives will move to the selected state. Select Unconfigured Good option to retain the existing configuration. Select JBOD to move the unconfigured drives to JBOD state. Select RAID0 to create a RAID0 virtual drive on each of the unconfigured drives. If JBOD is selected, unconfigured drives will move to JBOD state on host reboot. This setting is applicable only to selected set of controllers on FI attached servers. * `UnconfiguredGood` - Newly inserted drives or on reboot, drives will remain the same state. * `Jbod` - Newly inserted drives or on reboot, drives will automatically move to JBOD state if drive state was UnconfiguredGood. * `RAID0` - Newly inserted drives or on reboot, virtual drives will be created, respective drives will move to Online state.
	DefaultDriveMode *string `json:"DefaultDriveMode,omitempty"`
	// Only U.3 NVMe drives has to be specified, entered slots will be moved to Direct attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range.
	DirectAttachedNvmeSlots *string `json:"DirectAttachedNvmeSlots,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
	// A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen.
	GlobalHotSpares *string                             `json:"GlobalHotSpares,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
	M2VirtualDrive  NullableStorageM2VirtualDriveConfig `json:"M2VirtualDrive,omitempty"`
	Raid0Drive      NullableStorageR0Drive              `json:"Raid0Drive,omitempty"`
	// Only U.3 NVMe drives has to be specified, entered slots will be moved to RAID attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range.
	RaidAttachedNvmeSlots *string `json:"RaidAttachedNvmeSlots,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
	// JBOD drives specified in this slot range will be encrypted. Allowed values are 'ALL', or a comma or hyphen separated number range. Sample format is ALL or 1, 3 or 4-6, 8. Setting the value to 'ALL' will encrypt all the unused UnconfigureGood/JBOD disks.
	SecureJbods *string `json:"SecureJbods,omitempty" validate:"regexp=^$|^((((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*)|(a|A)(l|L)(l|L))$"`
	// State to which drives, not used in this policy, are to be moved. NoChange will not change the drive state. No Change must be selected if Default Drive State is set to JBOD or RAID0. * `NoChange` - Drive state will not be modified by Storage Policy. * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group. * `Jbod` - JBOD state where the disks start showing up to Host OS.
	UnusedDisksState *string `json:"UnusedDisksState,omitempty"`
	// Disks in JBOD State are used to create virtual drives. This setting must be disabled if Default Drive State is set to JBOD.
	UseJbodForVdCreation *bool `json:"UseJbodForVdCreation,omitempty"`
	// An array of relationships to storageDriveGroup resources.
	DriveGroup   []StorageDriveGroupRelationship              `json:"DriveGroup,omitempty"`
	Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageStoragePolicy StorageStoragePolicy

// NewStorageStoragePolicy instantiates a new StorageStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageStoragePolicy(classId string, objectType string) *StorageStoragePolicy {
	this := StorageStoragePolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultDriveMode string = "UnconfiguredGood"
	this.DefaultDriveMode = &defaultDriveMode
	var unusedDisksState string = "NoChange"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// NewStorageStoragePolicyWithDefaults instantiates a new StorageStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageStoragePolicyWithDefaults() *StorageStoragePolicy {
	this := StorageStoragePolicy{}
	var classId string = "storage.StoragePolicy"
	this.ClassId = classId
	var objectType string = "storage.StoragePolicy"
	this.ObjectType = objectType
	var defaultDriveMode string = "UnconfiguredGood"
	this.DefaultDriveMode = &defaultDriveMode
	var unusedDisksState string = "NoChange"
	this.UnusedDisksState = &unusedDisksState
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageStoragePolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageStoragePolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.StoragePolicy" of the ClassId field.
func (o *StorageStoragePolicy) GetDefaultClassId() interface{} {
	return "storage.StoragePolicy"
}

// GetObjectType returns the ObjectType field value
func (o *StorageStoragePolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageStoragePolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.StoragePolicy" of the ObjectType field.
func (o *StorageStoragePolicy) GetDefaultObjectType() interface{} {
	return "storage.StoragePolicy"
}

// GetDefaultDriveMode returns the DefaultDriveMode field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetDefaultDriveMode() string {
	if o == nil || IsNil(o.DefaultDriveMode) {
		var ret string
		return ret
	}
	return *o.DefaultDriveMode
}

// GetDefaultDriveModeOk returns a tuple with the DefaultDriveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetDefaultDriveModeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultDriveMode) {
		return nil, false
	}
	return o.DefaultDriveMode, true
}

// HasDefaultDriveMode returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasDefaultDriveMode() bool {
	if o != nil && !IsNil(o.DefaultDriveMode) {
		return true
	}

	return false
}

// SetDefaultDriveMode gets a reference to the given string and assigns it to the DefaultDriveMode field.
func (o *StorageStoragePolicy) SetDefaultDriveMode(v string) {
	o.DefaultDriveMode = &v
}

// GetDirectAttachedNvmeSlots returns the DirectAttachedNvmeSlots field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetDirectAttachedNvmeSlots() string {
	if o == nil || IsNil(o.DirectAttachedNvmeSlots) {
		var ret string
		return ret
	}
	return *o.DirectAttachedNvmeSlots
}

// GetDirectAttachedNvmeSlotsOk returns a tuple with the DirectAttachedNvmeSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetDirectAttachedNvmeSlotsOk() (*string, bool) {
	if o == nil || IsNil(o.DirectAttachedNvmeSlots) {
		return nil, false
	}
	return o.DirectAttachedNvmeSlots, true
}

// HasDirectAttachedNvmeSlots returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasDirectAttachedNvmeSlots() bool {
	if o != nil && !IsNil(o.DirectAttachedNvmeSlots) {
		return true
	}

	return false
}

// SetDirectAttachedNvmeSlots gets a reference to the given string and assigns it to the DirectAttachedNvmeSlots field.
func (o *StorageStoragePolicy) SetDirectAttachedNvmeSlots(v string) {
	o.DirectAttachedNvmeSlots = &v
}

// GetGlobalHotSpares returns the GlobalHotSpares field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetGlobalHotSpares() string {
	if o == nil || IsNil(o.GlobalHotSpares) {
		var ret string
		return ret
	}
	return *o.GlobalHotSpares
}

// GetGlobalHotSparesOk returns a tuple with the GlobalHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetGlobalHotSparesOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalHotSpares) {
		return nil, false
	}
	return o.GlobalHotSpares, true
}

// HasGlobalHotSpares returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasGlobalHotSpares() bool {
	if o != nil && !IsNil(o.GlobalHotSpares) {
		return true
	}

	return false
}

// SetGlobalHotSpares gets a reference to the given string and assigns it to the GlobalHotSpares field.
func (o *StorageStoragePolicy) SetGlobalHotSpares(v string) {
	o.GlobalHotSpares = &v
}

// GetM2VirtualDrive returns the M2VirtualDrive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetM2VirtualDrive() StorageM2VirtualDriveConfig {
	if o == nil || IsNil(o.M2VirtualDrive.Get()) {
		var ret StorageM2VirtualDriveConfig
		return ret
	}
	return *o.M2VirtualDrive.Get()
}

// GetM2VirtualDriveOk returns a tuple with the M2VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetM2VirtualDriveOk() (*StorageM2VirtualDriveConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.M2VirtualDrive.Get(), o.M2VirtualDrive.IsSet()
}

// HasM2VirtualDrive returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasM2VirtualDrive() bool {
	if o != nil && o.M2VirtualDrive.IsSet() {
		return true
	}

	return false
}

// SetM2VirtualDrive gets a reference to the given NullableStorageM2VirtualDriveConfig and assigns it to the M2VirtualDrive field.
func (o *StorageStoragePolicy) SetM2VirtualDrive(v StorageM2VirtualDriveConfig) {
	o.M2VirtualDrive.Set(&v)
}

// SetM2VirtualDriveNil sets the value for M2VirtualDrive to be an explicit nil
func (o *StorageStoragePolicy) SetM2VirtualDriveNil() {
	o.M2VirtualDrive.Set(nil)
}

// UnsetM2VirtualDrive ensures that no value is present for M2VirtualDrive, not even an explicit nil
func (o *StorageStoragePolicy) UnsetM2VirtualDrive() {
	o.M2VirtualDrive.Unset()
}

// GetRaid0Drive returns the Raid0Drive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetRaid0Drive() StorageR0Drive {
	if o == nil || IsNil(o.Raid0Drive.Get()) {
		var ret StorageR0Drive
		return ret
	}
	return *o.Raid0Drive.Get()
}

// GetRaid0DriveOk returns a tuple with the Raid0Drive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetRaid0DriveOk() (*StorageR0Drive, bool) {
	if o == nil {
		return nil, false
	}
	return o.Raid0Drive.Get(), o.Raid0Drive.IsSet()
}

// HasRaid0Drive returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasRaid0Drive() bool {
	if o != nil && o.Raid0Drive.IsSet() {
		return true
	}

	return false
}

// SetRaid0Drive gets a reference to the given NullableStorageR0Drive and assigns it to the Raid0Drive field.
func (o *StorageStoragePolicy) SetRaid0Drive(v StorageR0Drive) {
	o.Raid0Drive.Set(&v)
}

// SetRaid0DriveNil sets the value for Raid0Drive to be an explicit nil
func (o *StorageStoragePolicy) SetRaid0DriveNil() {
	o.Raid0Drive.Set(nil)
}

// UnsetRaid0Drive ensures that no value is present for Raid0Drive, not even an explicit nil
func (o *StorageStoragePolicy) UnsetRaid0Drive() {
	o.Raid0Drive.Unset()
}

// GetRaidAttachedNvmeSlots returns the RaidAttachedNvmeSlots field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetRaidAttachedNvmeSlots() string {
	if o == nil || IsNil(o.RaidAttachedNvmeSlots) {
		var ret string
		return ret
	}
	return *o.RaidAttachedNvmeSlots
}

// GetRaidAttachedNvmeSlotsOk returns a tuple with the RaidAttachedNvmeSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetRaidAttachedNvmeSlotsOk() (*string, bool) {
	if o == nil || IsNil(o.RaidAttachedNvmeSlots) {
		return nil, false
	}
	return o.RaidAttachedNvmeSlots, true
}

// HasRaidAttachedNvmeSlots returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasRaidAttachedNvmeSlots() bool {
	if o != nil && !IsNil(o.RaidAttachedNvmeSlots) {
		return true
	}

	return false
}

// SetRaidAttachedNvmeSlots gets a reference to the given string and assigns it to the RaidAttachedNvmeSlots field.
func (o *StorageStoragePolicy) SetRaidAttachedNvmeSlots(v string) {
	o.RaidAttachedNvmeSlots = &v
}

// GetSecureJbods returns the SecureJbods field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetSecureJbods() string {
	if o == nil || IsNil(o.SecureJbods) {
		var ret string
		return ret
	}
	return *o.SecureJbods
}

// GetSecureJbodsOk returns a tuple with the SecureJbods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetSecureJbodsOk() (*string, bool) {
	if o == nil || IsNil(o.SecureJbods) {
		return nil, false
	}
	return o.SecureJbods, true
}

// HasSecureJbods returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasSecureJbods() bool {
	if o != nil && !IsNil(o.SecureJbods) {
		return true
	}

	return false
}

// SetSecureJbods gets a reference to the given string and assigns it to the SecureJbods field.
func (o *StorageStoragePolicy) SetSecureJbods(v string) {
	o.SecureJbods = &v
}

// GetUnusedDisksState returns the UnusedDisksState field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetUnusedDisksState() string {
	if o == nil || IsNil(o.UnusedDisksState) {
		var ret string
		return ret
	}
	return *o.UnusedDisksState
}

// GetUnusedDisksStateOk returns a tuple with the UnusedDisksState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetUnusedDisksStateOk() (*string, bool) {
	if o == nil || IsNil(o.UnusedDisksState) {
		return nil, false
	}
	return o.UnusedDisksState, true
}

// HasUnusedDisksState returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasUnusedDisksState() bool {
	if o != nil && !IsNil(o.UnusedDisksState) {
		return true
	}

	return false
}

// SetUnusedDisksState gets a reference to the given string and assigns it to the UnusedDisksState field.
func (o *StorageStoragePolicy) SetUnusedDisksState(v string) {
	o.UnusedDisksState = &v
}

// GetUseJbodForVdCreation returns the UseJbodForVdCreation field value if set, zero value otherwise.
func (o *StorageStoragePolicy) GetUseJbodForVdCreation() bool {
	if o == nil || IsNil(o.UseJbodForVdCreation) {
		var ret bool
		return ret
	}
	return *o.UseJbodForVdCreation
}

// GetUseJbodForVdCreationOk returns a tuple with the UseJbodForVdCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageStoragePolicy) GetUseJbodForVdCreationOk() (*bool, bool) {
	if o == nil || IsNil(o.UseJbodForVdCreation) {
		return nil, false
	}
	return o.UseJbodForVdCreation, true
}

// HasUseJbodForVdCreation returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasUseJbodForVdCreation() bool {
	if o != nil && !IsNil(o.UseJbodForVdCreation) {
		return true
	}

	return false
}

// SetUseJbodForVdCreation gets a reference to the given bool and assigns it to the UseJbodForVdCreation field.
func (o *StorageStoragePolicy) SetUseJbodForVdCreation(v bool) {
	o.UseJbodForVdCreation = &v
}

// GetDriveGroup returns the DriveGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetDriveGroup() []StorageDriveGroupRelationship {
	if o == nil {
		var ret []StorageDriveGroupRelationship
		return ret
	}
	return o.DriveGroup
}

// GetDriveGroupOk returns a tuple with the DriveGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetDriveGroupOk() ([]StorageDriveGroupRelationship, bool) {
	if o == nil || IsNil(o.DriveGroup) {
		return nil, false
	}
	return o.DriveGroup, true
}

// HasDriveGroup returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasDriveGroup() bool {
	if o != nil && !IsNil(o.DriveGroup) {
		return true
	}

	return false
}

// SetDriveGroup gets a reference to the given []StorageDriveGroupRelationship and assigns it to the DriveGroup field.
func (o *StorageStoragePolicy) SetDriveGroup(v []StorageDriveGroupRelationship) {
	o.DriveGroup = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *StorageStoragePolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *StorageStoragePolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *StorageStoragePolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageStoragePolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageStoragePolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *StorageStoragePolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *StorageStoragePolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o StorageStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageStoragePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DefaultDriveMode) {
		toSerialize["DefaultDriveMode"] = o.DefaultDriveMode
	}
	if !IsNil(o.DirectAttachedNvmeSlots) {
		toSerialize["DirectAttachedNvmeSlots"] = o.DirectAttachedNvmeSlots
	}
	if !IsNil(o.GlobalHotSpares) {
		toSerialize["GlobalHotSpares"] = o.GlobalHotSpares
	}
	if o.M2VirtualDrive.IsSet() {
		toSerialize["M2VirtualDrive"] = o.M2VirtualDrive.Get()
	}
	if o.Raid0Drive.IsSet() {
		toSerialize["Raid0Drive"] = o.Raid0Drive.Get()
	}
	if !IsNil(o.RaidAttachedNvmeSlots) {
		toSerialize["RaidAttachedNvmeSlots"] = o.RaidAttachedNvmeSlots
	}
	if !IsNil(o.SecureJbods) {
		toSerialize["SecureJbods"] = o.SecureJbods
	}
	if !IsNil(o.UnusedDisksState) {
		toSerialize["UnusedDisksState"] = o.UnusedDisksState
	}
	if !IsNil(o.UseJbodForVdCreation) {
		toSerialize["UseJbodForVdCreation"] = o.UseJbodForVdCreation
	}
	if o.DriveGroup != nil {
		toSerialize["DriveGroup"] = o.DriveGroup
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageStoragePolicy) UnmarshalJSON(data []byte) (err error) {
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
	type StorageStoragePolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// All unconfigured drives will move to the selected state on deployment. Newly inserted drives will move to the selected state. Select Unconfigured Good option to retain the existing configuration. Select JBOD to move the unconfigured drives to JBOD state. Select RAID0 to create a RAID0 virtual drive on each of the unconfigured drives. If JBOD is selected, unconfigured drives will move to JBOD state on host reboot. This setting is applicable only to selected set of controllers on FI attached servers. * `UnconfiguredGood` - Newly inserted drives or on reboot, drives will remain the same state. * `Jbod` - Newly inserted drives or on reboot, drives will automatically move to JBOD state if drive state was UnconfiguredGood. * `RAID0` - Newly inserted drives or on reboot, virtual drives will be created, respective drives will move to Online state.
		DefaultDriveMode *string `json:"DefaultDriveMode,omitempty"`
		// Only U.3 NVMe drives has to be specified, entered slots will be moved to Direct attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range.
		DirectAttachedNvmeSlots *string `json:"DirectAttachedNvmeSlots,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
		// A collection of disks that is to be used as hot spares, globally, for all the RAID groups. Allowed value is a number range separated by a comma or a hyphen.
		GlobalHotSpares *string                             `json:"GlobalHotSpares,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
		M2VirtualDrive  NullableStorageM2VirtualDriveConfig `json:"M2VirtualDrive,omitempty"`
		Raid0Drive      NullableStorageR0Drive              `json:"Raid0Drive,omitempty"`
		// Only U.3 NVMe drives has to be specified, entered slots will be moved to RAID attached mode. Allowed slots are 1-4, 101-104. Allowed value is a comma or hyphen separated number range.
		RaidAttachedNvmeSlots *string `json:"RaidAttachedNvmeSlots,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
		// JBOD drives specified in this slot range will be encrypted. Allowed values are 'ALL', or a comma or hyphen separated number range. Sample format is ALL or 1, 3 or 4-6, 8. Setting the value to 'ALL' will encrypt all the unused UnconfigureGood/JBOD disks.
		SecureJbods *string `json:"SecureJbods,omitempty" validate:"regexp=^$|^((((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*)|(a|A)(l|L)(l|L))$"`
		// State to which drives, not used in this policy, are to be moved. NoChange will not change the drive state. No Change must be selected if Default Drive State is set to JBOD or RAID0. * `NoChange` - Drive state will not be modified by Storage Policy. * `UnconfiguredGood` - Unconfigured good state -ready to be added in a RAID group. * `Jbod` - JBOD state where the disks start showing up to Host OS.
		UnusedDisksState *string `json:"UnusedDisksState,omitempty"`
		// Disks in JBOD State are used to create virtual drives. This setting must be disabled if Default Drive State is set to JBOD.
		UseJbodForVdCreation *bool `json:"UseJbodForVdCreation,omitempty"`
		// An array of relationships to storageDriveGroup resources.
		DriveGroup   []StorageDriveGroupRelationship              `json:"DriveGroup,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varStorageStoragePolicyWithoutEmbeddedStruct := StorageStoragePolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageStoragePolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageStoragePolicy := _StorageStoragePolicy{}
		varStorageStoragePolicy.ClassId = varStorageStoragePolicyWithoutEmbeddedStruct.ClassId
		varStorageStoragePolicy.ObjectType = varStorageStoragePolicyWithoutEmbeddedStruct.ObjectType
		varStorageStoragePolicy.DefaultDriveMode = varStorageStoragePolicyWithoutEmbeddedStruct.DefaultDriveMode
		varStorageStoragePolicy.DirectAttachedNvmeSlots = varStorageStoragePolicyWithoutEmbeddedStruct.DirectAttachedNvmeSlots
		varStorageStoragePolicy.GlobalHotSpares = varStorageStoragePolicyWithoutEmbeddedStruct.GlobalHotSpares
		varStorageStoragePolicy.M2VirtualDrive = varStorageStoragePolicyWithoutEmbeddedStruct.M2VirtualDrive
		varStorageStoragePolicy.Raid0Drive = varStorageStoragePolicyWithoutEmbeddedStruct.Raid0Drive
		varStorageStoragePolicy.RaidAttachedNvmeSlots = varStorageStoragePolicyWithoutEmbeddedStruct.RaidAttachedNvmeSlots
		varStorageStoragePolicy.SecureJbods = varStorageStoragePolicyWithoutEmbeddedStruct.SecureJbods
		varStorageStoragePolicy.UnusedDisksState = varStorageStoragePolicyWithoutEmbeddedStruct.UnusedDisksState
		varStorageStoragePolicy.UseJbodForVdCreation = varStorageStoragePolicyWithoutEmbeddedStruct.UseJbodForVdCreation
		varStorageStoragePolicy.DriveGroup = varStorageStoragePolicyWithoutEmbeddedStruct.DriveGroup
		varStorageStoragePolicy.Organization = varStorageStoragePolicyWithoutEmbeddedStruct.Organization
		varStorageStoragePolicy.Profiles = varStorageStoragePolicyWithoutEmbeddedStruct.Profiles
		*o = StorageStoragePolicy(varStorageStoragePolicy)
	} else {
		return err
	}

	varStorageStoragePolicy := _StorageStoragePolicy{}

	err = json.Unmarshal(data, &varStorageStoragePolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varStorageStoragePolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DefaultDriveMode")
		delete(additionalProperties, "DirectAttachedNvmeSlots")
		delete(additionalProperties, "GlobalHotSpares")
		delete(additionalProperties, "M2VirtualDrive")
		delete(additionalProperties, "Raid0Drive")
		delete(additionalProperties, "RaidAttachedNvmeSlots")
		delete(additionalProperties, "SecureJbods")
		delete(additionalProperties, "UnusedDisksState")
		delete(additionalProperties, "UseJbodForVdCreation")
		delete(additionalProperties, "DriveGroup")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableStorageStoragePolicy struct {
	value *StorageStoragePolicy
	isSet bool
}

func (v NullableStorageStoragePolicy) Get() *StorageStoragePolicy {
	return v.value
}

func (v *NullableStorageStoragePolicy) Set(val *StorageStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageStoragePolicy(val *StorageStoragePolicy) *NullableStorageStoragePolicy {
	return &NullableStorageStoragePolicy{value: val, isSet: true}
}

func (v NullableStorageStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
