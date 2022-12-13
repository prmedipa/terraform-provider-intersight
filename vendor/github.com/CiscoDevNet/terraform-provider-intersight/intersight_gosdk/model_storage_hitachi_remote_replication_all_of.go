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

// StorageHitachiRemoteReplicationAllOf Definition of the list of properties defined in 'storage.HitachiRemoteReplication', excluding properties defined in parent classes.
type StorageHitachiRemoteReplicationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Consistency group ID. If no consistency group consists, information is not input.
	ConsistencyGroupId *string `json:"ConsistencyGroupId,omitempty"`
	// Copy speed. Number for the size of tracks to be copied.
	CopyPace *string `json:"CopyPace,omitempty"`
	// Status of the 3DC multi-target configuration that uses delta resync.
	DeltaStatus *string `json:"DeltaStatus,omitempty"`
	// Fence level. Whether the P-VOL can be written to when the pair is split due to error.
	FenceLevel *string `json:"FenceLevel,omitempty"`
	// MU (mirror unit) number of the volume.
	MuNumber *string `json:"MuNumber,omitempty"`
	// Object ID of the remote copy pair.
	Name *string `json:"Name,omitempty"`
	// Path group ID of the RCU.
	PathGroupId *string `json:"PathGroupId,omitempty"`
	// I-O mode of the P-VOL. Information is input only in the case of global-active device.
	PvolIoMode *string `json:"PvolIoMode,omitempty"`
	// Journal ID of the P-VOL. A value is input only in the case of Universal Replicator.
	PvolJournalId *string `json:"PvolJournalId,omitempty"`
	// LDEV number of primary volume.
	PvolLdevId *int64 `json:"PvolLdevId,omitempty"`
	// Serial number of the storage system on the P-VOL.
	PvolStorageSerial *string `json:"PvolStorageSerial,omitempty"`
	// ID of the Quorum disk. A value is input only in the case of global-active device.
	QuorumDiskId *string `json:"QuorumDiskId,omitempty"`
	// Pair type of the remote copy pair.
	ReplicationType *string `json:"ReplicationType,omitempty"`
	// Status of the remote copy pair.
	Status *string `json:"Status,omitempty"`
	// I-O mode of the S-VOL. Information is input only in the case of global-active device.
	SvolIoMode *string `json:"SvolIoMode,omitempty"`
	// Journal ID of the S-VOL. A value is input only in the case of Universal Replicator.
	SvolJournalId *string `json:"SvolJournalId,omitempty"`
	// LDEV number of secondary volume.
	SvolLdevId *int64 `json:"SvolLdevId,omitempty"`
	// Serial number of the storage system on the S-VOL.
	SvolStorageSerial    *string                              `json:"SvolStorageSerial,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiRemoteReplicationAllOf StorageHitachiRemoteReplicationAllOf

// NewStorageHitachiRemoteReplicationAllOf instantiates a new StorageHitachiRemoteReplicationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiRemoteReplicationAllOf(classId string, objectType string) *StorageHitachiRemoteReplicationAllOf {
	this := StorageHitachiRemoteReplicationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiRemoteReplicationAllOfWithDefaults instantiates a new StorageHitachiRemoteReplicationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiRemoteReplicationAllOfWithDefaults() *StorageHitachiRemoteReplicationAllOf {
	this := StorageHitachiRemoteReplicationAllOf{}
	var classId string = "storage.HitachiRemoteReplication"
	this.ClassId = classId
	var objectType string = "storage.HitachiRemoteReplication"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiRemoteReplicationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiRemoteReplicationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiRemoteReplicationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiRemoteReplicationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetConsistencyGroupId returns the ConsistencyGroupId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetConsistencyGroupId() string {
	if o == nil || o.ConsistencyGroupId == nil {
		var ret string
		return ret
	}
	return *o.ConsistencyGroupId
}

// GetConsistencyGroupIdOk returns a tuple with the ConsistencyGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetConsistencyGroupIdOk() (*string, bool) {
	if o == nil || o.ConsistencyGroupId == nil {
		return nil, false
	}
	return o.ConsistencyGroupId, true
}

// HasConsistencyGroupId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasConsistencyGroupId() bool {
	if o != nil && o.ConsistencyGroupId != nil {
		return true
	}

	return false
}

// SetConsistencyGroupId gets a reference to the given string and assigns it to the ConsistencyGroupId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetConsistencyGroupId(v string) {
	o.ConsistencyGroupId = &v
}

// GetCopyPace returns the CopyPace field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetCopyPace() string {
	if o == nil || o.CopyPace == nil {
		var ret string
		return ret
	}
	return *o.CopyPace
}

// GetCopyPaceOk returns a tuple with the CopyPace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetCopyPaceOk() (*string, bool) {
	if o == nil || o.CopyPace == nil {
		return nil, false
	}
	return o.CopyPace, true
}

// HasCopyPace returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasCopyPace() bool {
	if o != nil && o.CopyPace != nil {
		return true
	}

	return false
}

// SetCopyPace gets a reference to the given string and assigns it to the CopyPace field.
func (o *StorageHitachiRemoteReplicationAllOf) SetCopyPace(v string) {
	o.CopyPace = &v
}

// GetDeltaStatus returns the DeltaStatus field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetDeltaStatus() string {
	if o == nil || o.DeltaStatus == nil {
		var ret string
		return ret
	}
	return *o.DeltaStatus
}

// GetDeltaStatusOk returns a tuple with the DeltaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetDeltaStatusOk() (*string, bool) {
	if o == nil || o.DeltaStatus == nil {
		return nil, false
	}
	return o.DeltaStatus, true
}

// HasDeltaStatus returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasDeltaStatus() bool {
	if o != nil && o.DeltaStatus != nil {
		return true
	}

	return false
}

// SetDeltaStatus gets a reference to the given string and assigns it to the DeltaStatus field.
func (o *StorageHitachiRemoteReplicationAllOf) SetDeltaStatus(v string) {
	o.DeltaStatus = &v
}

// GetFenceLevel returns the FenceLevel field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetFenceLevel() string {
	if o == nil || o.FenceLevel == nil {
		var ret string
		return ret
	}
	return *o.FenceLevel
}

// GetFenceLevelOk returns a tuple with the FenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetFenceLevelOk() (*string, bool) {
	if o == nil || o.FenceLevel == nil {
		return nil, false
	}
	return o.FenceLevel, true
}

// HasFenceLevel returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasFenceLevel() bool {
	if o != nil && o.FenceLevel != nil {
		return true
	}

	return false
}

// SetFenceLevel gets a reference to the given string and assigns it to the FenceLevel field.
func (o *StorageHitachiRemoteReplicationAllOf) SetFenceLevel(v string) {
	o.FenceLevel = &v
}

// GetMuNumber returns the MuNumber field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetMuNumber() string {
	if o == nil || o.MuNumber == nil {
		var ret string
		return ret
	}
	return *o.MuNumber
}

// GetMuNumberOk returns a tuple with the MuNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetMuNumberOk() (*string, bool) {
	if o == nil || o.MuNumber == nil {
		return nil, false
	}
	return o.MuNumber, true
}

// HasMuNumber returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasMuNumber() bool {
	if o != nil && o.MuNumber != nil {
		return true
	}

	return false
}

// SetMuNumber gets a reference to the given string and assigns it to the MuNumber field.
func (o *StorageHitachiRemoteReplicationAllOf) SetMuNumber(v string) {
	o.MuNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageHitachiRemoteReplicationAllOf) SetName(v string) {
	o.Name = &v
}

// GetPathGroupId returns the PathGroupId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetPathGroupId() string {
	if o == nil || o.PathGroupId == nil {
		var ret string
		return ret
	}
	return *o.PathGroupId
}

// GetPathGroupIdOk returns a tuple with the PathGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetPathGroupIdOk() (*string, bool) {
	if o == nil || o.PathGroupId == nil {
		return nil, false
	}
	return o.PathGroupId, true
}

// HasPathGroupId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasPathGroupId() bool {
	if o != nil && o.PathGroupId != nil {
		return true
	}

	return false
}

// SetPathGroupId gets a reference to the given string and assigns it to the PathGroupId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetPathGroupId(v string) {
	o.PathGroupId = &v
}

// GetPvolIoMode returns the PvolIoMode field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolIoMode() string {
	if o == nil || o.PvolIoMode == nil {
		var ret string
		return ret
	}
	return *o.PvolIoMode
}

// GetPvolIoModeOk returns a tuple with the PvolIoMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolIoModeOk() (*string, bool) {
	if o == nil || o.PvolIoMode == nil {
		return nil, false
	}
	return o.PvolIoMode, true
}

// HasPvolIoMode returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasPvolIoMode() bool {
	if o != nil && o.PvolIoMode != nil {
		return true
	}

	return false
}

// SetPvolIoMode gets a reference to the given string and assigns it to the PvolIoMode field.
func (o *StorageHitachiRemoteReplicationAllOf) SetPvolIoMode(v string) {
	o.PvolIoMode = &v
}

// GetPvolJournalId returns the PvolJournalId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolJournalId() string {
	if o == nil || o.PvolJournalId == nil {
		var ret string
		return ret
	}
	return *o.PvolJournalId
}

// GetPvolJournalIdOk returns a tuple with the PvolJournalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolJournalIdOk() (*string, bool) {
	if o == nil || o.PvolJournalId == nil {
		return nil, false
	}
	return o.PvolJournalId, true
}

// HasPvolJournalId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasPvolJournalId() bool {
	if o != nil && o.PvolJournalId != nil {
		return true
	}

	return false
}

// SetPvolJournalId gets a reference to the given string and assigns it to the PvolJournalId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetPvolJournalId(v string) {
	o.PvolJournalId = &v
}

// GetPvolLdevId returns the PvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolLdevId() int64 {
	if o == nil || o.PvolLdevId == nil {
		var ret int64
		return ret
	}
	return *o.PvolLdevId
}

// GetPvolLdevIdOk returns a tuple with the PvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolLdevIdOk() (*int64, bool) {
	if o == nil || o.PvolLdevId == nil {
		return nil, false
	}
	return o.PvolLdevId, true
}

// HasPvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasPvolLdevId() bool {
	if o != nil && o.PvolLdevId != nil {
		return true
	}

	return false
}

// SetPvolLdevId gets a reference to the given int64 and assigns it to the PvolLdevId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetPvolLdevId(v int64) {
	o.PvolLdevId = &v
}

// GetPvolStorageSerial returns the PvolStorageSerial field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolStorageSerial() string {
	if o == nil || o.PvolStorageSerial == nil {
		var ret string
		return ret
	}
	return *o.PvolStorageSerial
}

// GetPvolStorageSerialOk returns a tuple with the PvolStorageSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetPvolStorageSerialOk() (*string, bool) {
	if o == nil || o.PvolStorageSerial == nil {
		return nil, false
	}
	return o.PvolStorageSerial, true
}

// HasPvolStorageSerial returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasPvolStorageSerial() bool {
	if o != nil && o.PvolStorageSerial != nil {
		return true
	}

	return false
}

// SetPvolStorageSerial gets a reference to the given string and assigns it to the PvolStorageSerial field.
func (o *StorageHitachiRemoteReplicationAllOf) SetPvolStorageSerial(v string) {
	o.PvolStorageSerial = &v
}

// GetQuorumDiskId returns the QuorumDiskId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetQuorumDiskId() string {
	if o == nil || o.QuorumDiskId == nil {
		var ret string
		return ret
	}
	return *o.QuorumDiskId
}

// GetQuorumDiskIdOk returns a tuple with the QuorumDiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetQuorumDiskIdOk() (*string, bool) {
	if o == nil || o.QuorumDiskId == nil {
		return nil, false
	}
	return o.QuorumDiskId, true
}

// HasQuorumDiskId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasQuorumDiskId() bool {
	if o != nil && o.QuorumDiskId != nil {
		return true
	}

	return false
}

// SetQuorumDiskId gets a reference to the given string and assigns it to the QuorumDiskId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetQuorumDiskId(v string) {
	o.QuorumDiskId = &v
}

// GetReplicationType returns the ReplicationType field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetReplicationType() string {
	if o == nil || o.ReplicationType == nil {
		var ret string
		return ret
	}
	return *o.ReplicationType
}

// GetReplicationTypeOk returns a tuple with the ReplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetReplicationTypeOk() (*string, bool) {
	if o == nil || o.ReplicationType == nil {
		return nil, false
	}
	return o.ReplicationType, true
}

// HasReplicationType returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasReplicationType() bool {
	if o != nil && o.ReplicationType != nil {
		return true
	}

	return false
}

// SetReplicationType gets a reference to the given string and assigns it to the ReplicationType field.
func (o *StorageHitachiRemoteReplicationAllOf) SetReplicationType(v string) {
	o.ReplicationType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageHitachiRemoteReplicationAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetSvolIoMode returns the SvolIoMode field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolIoMode() string {
	if o == nil || o.SvolIoMode == nil {
		var ret string
		return ret
	}
	return *o.SvolIoMode
}

// GetSvolIoModeOk returns a tuple with the SvolIoMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolIoModeOk() (*string, bool) {
	if o == nil || o.SvolIoMode == nil {
		return nil, false
	}
	return o.SvolIoMode, true
}

// HasSvolIoMode returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasSvolIoMode() bool {
	if o != nil && o.SvolIoMode != nil {
		return true
	}

	return false
}

// SetSvolIoMode gets a reference to the given string and assigns it to the SvolIoMode field.
func (o *StorageHitachiRemoteReplicationAllOf) SetSvolIoMode(v string) {
	o.SvolIoMode = &v
}

// GetSvolJournalId returns the SvolJournalId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolJournalId() string {
	if o == nil || o.SvolJournalId == nil {
		var ret string
		return ret
	}
	return *o.SvolJournalId
}

// GetSvolJournalIdOk returns a tuple with the SvolJournalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolJournalIdOk() (*string, bool) {
	if o == nil || o.SvolJournalId == nil {
		return nil, false
	}
	return o.SvolJournalId, true
}

// HasSvolJournalId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasSvolJournalId() bool {
	if o != nil && o.SvolJournalId != nil {
		return true
	}

	return false
}

// SetSvolJournalId gets a reference to the given string and assigns it to the SvolJournalId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetSvolJournalId(v string) {
	o.SvolJournalId = &v
}

// GetSvolLdevId returns the SvolLdevId field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolLdevId() int64 {
	if o == nil || o.SvolLdevId == nil {
		var ret int64
		return ret
	}
	return *o.SvolLdevId
}

// GetSvolLdevIdOk returns a tuple with the SvolLdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolLdevIdOk() (*int64, bool) {
	if o == nil || o.SvolLdevId == nil {
		return nil, false
	}
	return o.SvolLdevId, true
}

// HasSvolLdevId returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasSvolLdevId() bool {
	if o != nil && o.SvolLdevId != nil {
		return true
	}

	return false
}

// SetSvolLdevId gets a reference to the given int64 and assigns it to the SvolLdevId field.
func (o *StorageHitachiRemoteReplicationAllOf) SetSvolLdevId(v int64) {
	o.SvolLdevId = &v
}

// GetSvolStorageSerial returns the SvolStorageSerial field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolStorageSerial() string {
	if o == nil || o.SvolStorageSerial == nil {
		var ret string
		return ret
	}
	return *o.SvolStorageSerial
}

// GetSvolStorageSerialOk returns a tuple with the SvolStorageSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetSvolStorageSerialOk() (*string, bool) {
	if o == nil || o.SvolStorageSerial == nil {
		return nil, false
	}
	return o.SvolStorageSerial, true
}

// HasSvolStorageSerial returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasSvolStorageSerial() bool {
	if o != nil && o.SvolStorageSerial != nil {
		return true
	}

	return false
}

// SetSvolStorageSerial gets a reference to the given string and assigns it to the SvolStorageSerial field.
func (o *StorageHitachiRemoteReplicationAllOf) SetSvolStorageSerial(v string) {
	o.SvolStorageSerial = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiRemoteReplicationAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiRemoteReplicationAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiRemoteReplicationAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiRemoteReplicationAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiRemoteReplicationAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiRemoteReplicationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ConsistencyGroupId != nil {
		toSerialize["ConsistencyGroupId"] = o.ConsistencyGroupId
	}
	if o.CopyPace != nil {
		toSerialize["CopyPace"] = o.CopyPace
	}
	if o.DeltaStatus != nil {
		toSerialize["DeltaStatus"] = o.DeltaStatus
	}
	if o.FenceLevel != nil {
		toSerialize["FenceLevel"] = o.FenceLevel
	}
	if o.MuNumber != nil {
		toSerialize["MuNumber"] = o.MuNumber
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PathGroupId != nil {
		toSerialize["PathGroupId"] = o.PathGroupId
	}
	if o.PvolIoMode != nil {
		toSerialize["PvolIoMode"] = o.PvolIoMode
	}
	if o.PvolJournalId != nil {
		toSerialize["PvolJournalId"] = o.PvolJournalId
	}
	if o.PvolLdevId != nil {
		toSerialize["PvolLdevId"] = o.PvolLdevId
	}
	if o.PvolStorageSerial != nil {
		toSerialize["PvolStorageSerial"] = o.PvolStorageSerial
	}
	if o.QuorumDiskId != nil {
		toSerialize["QuorumDiskId"] = o.QuorumDiskId
	}
	if o.ReplicationType != nil {
		toSerialize["ReplicationType"] = o.ReplicationType
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SvolIoMode != nil {
		toSerialize["SvolIoMode"] = o.SvolIoMode
	}
	if o.SvolJournalId != nil {
		toSerialize["SvolJournalId"] = o.SvolJournalId
	}
	if o.SvolLdevId != nil {
		toSerialize["SvolLdevId"] = o.SvolLdevId
	}
	if o.SvolStorageSerial != nil {
		toSerialize["SvolStorageSerial"] = o.SvolStorageSerial
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiRemoteReplicationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiRemoteReplicationAllOf := _StorageHitachiRemoteReplicationAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiRemoteReplicationAllOf); err == nil {
		*o = StorageHitachiRemoteReplicationAllOf(varStorageHitachiRemoteReplicationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ConsistencyGroupId")
		delete(additionalProperties, "CopyPace")
		delete(additionalProperties, "DeltaStatus")
		delete(additionalProperties, "FenceLevel")
		delete(additionalProperties, "MuNumber")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PathGroupId")
		delete(additionalProperties, "PvolIoMode")
		delete(additionalProperties, "PvolJournalId")
		delete(additionalProperties, "PvolLdevId")
		delete(additionalProperties, "PvolStorageSerial")
		delete(additionalProperties, "QuorumDiskId")
		delete(additionalProperties, "ReplicationType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SvolIoMode")
		delete(additionalProperties, "SvolJournalId")
		delete(additionalProperties, "SvolLdevId")
		delete(additionalProperties, "SvolStorageSerial")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiRemoteReplicationAllOf struct {
	value *StorageHitachiRemoteReplicationAllOf
	isSet bool
}

func (v NullableStorageHitachiRemoteReplicationAllOf) Get() *StorageHitachiRemoteReplicationAllOf {
	return v.value
}

func (v *NullableStorageHitachiRemoteReplicationAllOf) Set(val *StorageHitachiRemoteReplicationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiRemoteReplicationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiRemoteReplicationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiRemoteReplicationAllOf(val *StorageHitachiRemoteReplicationAllOf) *NullableStorageHitachiRemoteReplicationAllOf {
	return &NullableStorageHitachiRemoteReplicationAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiRemoteReplicationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiRemoteReplicationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
