/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
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

// VirtualizationVmwareVirtualMachineSnapshot The virtual machine snapshot is represented here.
type VirtualizationVmwareVirtualMachineSnapshot struct {
	VirtualizationBaseVirtualMachineSnapshot
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Snapshot creation time. Time at which snapshot gets created.
	CreationTime *time.Time `json:"CreationTime,omitempty"`
	// If yes, it determines it is the latest snapshot of the virtual machine.
	CurrentSnapshot *bool `json:"CurrentSnapshot,omitempty"`
	// User provided description of the virtual machine snapshot.
	Description *string `json:"Description,omitempty"`
	// If yes, the virtual machine snapshot cannot be deleted.
	Golden *bool `json:"Golden,omitempty"`
	// The internally assigned id/key of virtual machine snapshot.
	Key *int64 `json:"Key,omitempty"`
	// Predecessor id is the id of the parent snapshot.
	PredecessorId *int64 `json:"PredecessorId,omitempty"`
	// Quiesce pauses all the I/O operations on virtual machine till the snapshot is taken.
	Quiesced *bool `json:"Quiesced,omitempty"`
	// Internally assigned MOR reference value.
	RefValue *string `json:"RefValue,omitempty"`
	// Size of the snapshot file created of the virtual machine, stored in bytes.
	SnapshotSize         *int64                                          `json:"SnapshotSize,omitempty"`
	VirtualMachine       *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwareVirtualMachineSnapshot VirtualizationVmwareVirtualMachineSnapshot

// NewVirtualizationVmwareVirtualMachineSnapshot instantiates a new VirtualizationVmwareVirtualMachineSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwareVirtualMachineSnapshot(classId string, objectType string) *VirtualizationVmwareVirtualMachineSnapshot {
	this := VirtualizationVmwareVirtualMachineSnapshot{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwareVirtualMachineSnapshotWithDefaults instantiates a new VirtualizationVmwareVirtualMachineSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwareVirtualMachineSnapshotWithDefaults() *VirtualizationVmwareVirtualMachineSnapshot {
	this := VirtualizationVmwareVirtualMachineSnapshot{}
	var classId string = "virtualization.VmwareVirtualMachineSnapshot"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareVirtualMachineSnapshot"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCreationTime() time.Time {
	if o == nil || o.CreationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetCurrentSnapshot returns the CurrentSnapshot field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCurrentSnapshot() bool {
	if o == nil || o.CurrentSnapshot == nil {
		var ret bool
		return ret
	}
	return *o.CurrentSnapshot
}

// GetCurrentSnapshotOk returns a tuple with the CurrentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetCurrentSnapshotOk() (*bool, bool) {
	if o == nil || o.CurrentSnapshot == nil {
		return nil, false
	}
	return o.CurrentSnapshot, true
}

// HasCurrentSnapshot returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasCurrentSnapshot() bool {
	if o != nil && o.CurrentSnapshot != nil {
		return true
	}

	return false
}

// SetCurrentSnapshot gets a reference to the given bool and assigns it to the CurrentSnapshot field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetCurrentSnapshot(v bool) {
	o.CurrentSnapshot = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetDescription(v string) {
	o.Description = &v
}

// GetGolden returns the Golden field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetGolden() bool {
	if o == nil || o.Golden == nil {
		var ret bool
		return ret
	}
	return *o.Golden
}

// GetGoldenOk returns a tuple with the Golden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetGoldenOk() (*bool, bool) {
	if o == nil || o.Golden == nil {
		return nil, false
	}
	return o.Golden, true
}

// HasGolden returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasGolden() bool {
	if o != nil && o.Golden != nil {
		return true
	}

	return false
}

// SetGolden gets a reference to the given bool and assigns it to the Golden field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetGolden(v bool) {
	o.Golden = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetKey() int64 {
	if o == nil || o.Key == nil {
		var ret int64
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetKeyOk() (*int64, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given int64 and assigns it to the Key field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetKey(v int64) {
	o.Key = &v
}

// GetPredecessorId returns the PredecessorId field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetPredecessorId() int64 {
	if o == nil || o.PredecessorId == nil {
		var ret int64
		return ret
	}
	return *o.PredecessorId
}

// GetPredecessorIdOk returns a tuple with the PredecessorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetPredecessorIdOk() (*int64, bool) {
	if o == nil || o.PredecessorId == nil {
		return nil, false
	}
	return o.PredecessorId, true
}

// HasPredecessorId returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasPredecessorId() bool {
	if o != nil && o.PredecessorId != nil {
		return true
	}

	return false
}

// SetPredecessorId gets a reference to the given int64 and assigns it to the PredecessorId field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetPredecessorId(v int64) {
	o.PredecessorId = &v
}

// GetQuiesced returns the Quiesced field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetQuiesced() bool {
	if o == nil || o.Quiesced == nil {
		var ret bool
		return ret
	}
	return *o.Quiesced
}

// GetQuiescedOk returns a tuple with the Quiesced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetQuiescedOk() (*bool, bool) {
	if o == nil || o.Quiesced == nil {
		return nil, false
	}
	return o.Quiesced, true
}

// HasQuiesced returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasQuiesced() bool {
	if o != nil && o.Quiesced != nil {
		return true
	}

	return false
}

// SetQuiesced gets a reference to the given bool and assigns it to the Quiesced field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetQuiesced(v bool) {
	o.Quiesced = &v
}

// GetRefValue returns the RefValue field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetRefValue() string {
	if o == nil || o.RefValue == nil {
		var ret string
		return ret
	}
	return *o.RefValue
}

// GetRefValueOk returns a tuple with the RefValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetRefValueOk() (*string, bool) {
	if o == nil || o.RefValue == nil {
		return nil, false
	}
	return o.RefValue, true
}

// HasRefValue returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasRefValue() bool {
	if o != nil && o.RefValue != nil {
		return true
	}

	return false
}

// SetRefValue gets a reference to the given string and assigns it to the RefValue field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetRefValue(v string) {
	o.RefValue = &v
}

// GetSnapshotSize returns the SnapshotSize field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetSnapshotSize() int64 {
	if o == nil || o.SnapshotSize == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotSize
}

// GetSnapshotSizeOk returns a tuple with the SnapshotSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetSnapshotSizeOk() (*int64, bool) {
	if o == nil || o.SnapshotSize == nil {
		return nil, false
	}
	return o.SnapshotSize, true
}

// HasSnapshotSize returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasSnapshotSize() bool {
	if o != nil && o.SnapshotSize != nil {
		return true
	}

	return false
}

// SetSnapshotSize gets a reference to the given int64 and assigns it to the SnapshotSize field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetSnapshotSize(v int64) {
	o.SnapshotSize = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetVirtualMachine() VirtualizationVmwareVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret VirtualizationVmwareVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) GetVirtualMachineOk() (*VirtualizationVmwareVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *VirtualizationVmwareVirtualMachineSnapshot) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given VirtualizationVmwareVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *VirtualizationVmwareVirtualMachineSnapshot) SetVirtualMachine(v VirtualizationVmwareVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o VirtualizationVmwareVirtualMachineSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseVirtualMachineSnapshot, errVirtualizationBaseVirtualMachineSnapshot := json.Marshal(o.VirtualizationBaseVirtualMachineSnapshot)
	if errVirtualizationBaseVirtualMachineSnapshot != nil {
		return []byte{}, errVirtualizationBaseVirtualMachineSnapshot
	}
	errVirtualizationBaseVirtualMachineSnapshot = json.Unmarshal([]byte(serializedVirtualizationBaseVirtualMachineSnapshot), &toSerialize)
	if errVirtualizationBaseVirtualMachineSnapshot != nil {
		return []byte{}, errVirtualizationBaseVirtualMachineSnapshot
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CreationTime != nil {
		toSerialize["CreationTime"] = o.CreationTime
	}
	if o.CurrentSnapshot != nil {
		toSerialize["CurrentSnapshot"] = o.CurrentSnapshot
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Golden != nil {
		toSerialize["Golden"] = o.Golden
	}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.PredecessorId != nil {
		toSerialize["PredecessorId"] = o.PredecessorId
	}
	if o.Quiesced != nil {
		toSerialize["Quiesced"] = o.Quiesced
	}
	if o.RefValue != nil {
		toSerialize["RefValue"] = o.RefValue
	}
	if o.SnapshotSize != nil {
		toSerialize["SnapshotSize"] = o.SnapshotSize
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationVmwareVirtualMachineSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Snapshot creation time. Time at which snapshot gets created.
		CreationTime *time.Time `json:"CreationTime,omitempty"`
		// If yes, it determines it is the latest snapshot of the virtual machine.
		CurrentSnapshot *bool `json:"CurrentSnapshot,omitempty"`
		// User provided description of the virtual machine snapshot.
		Description *string `json:"Description,omitempty"`
		// If yes, the virtual machine snapshot cannot be deleted.
		Golden *bool `json:"Golden,omitempty"`
		// The internally assigned id/key of virtual machine snapshot.
		Key *int64 `json:"Key,omitempty"`
		// Predecessor id is the id of the parent snapshot.
		PredecessorId *int64 `json:"PredecessorId,omitempty"`
		// Quiesce pauses all the I/O operations on virtual machine till the snapshot is taken.
		Quiesced *bool `json:"Quiesced,omitempty"`
		// Internally assigned MOR reference value.
		RefValue *string `json:"RefValue,omitempty"`
		// Size of the snapshot file created of the virtual machine, stored in bytes.
		SnapshotSize   *int64                                          `json:"SnapshotSize,omitempty"`
		VirtualMachine *VirtualizationVmwareVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct := VirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwareVirtualMachineSnapshot := _VirtualizationVmwareVirtualMachineSnapshot{}
		varVirtualizationVmwareVirtualMachineSnapshot.ClassId = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwareVirtualMachineSnapshot.ObjectType = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwareVirtualMachineSnapshot.CreationTime = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.CreationTime
		varVirtualizationVmwareVirtualMachineSnapshot.CurrentSnapshot = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.CurrentSnapshot
		varVirtualizationVmwareVirtualMachineSnapshot.Description = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.Description
		varVirtualizationVmwareVirtualMachineSnapshot.Golden = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.Golden
		varVirtualizationVmwareVirtualMachineSnapshot.Key = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.Key
		varVirtualizationVmwareVirtualMachineSnapshot.PredecessorId = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.PredecessorId
		varVirtualizationVmwareVirtualMachineSnapshot.Quiesced = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.Quiesced
		varVirtualizationVmwareVirtualMachineSnapshot.RefValue = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.RefValue
		varVirtualizationVmwareVirtualMachineSnapshot.SnapshotSize = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.SnapshotSize
		varVirtualizationVmwareVirtualMachineSnapshot.VirtualMachine = varVirtualizationVmwareVirtualMachineSnapshotWithoutEmbeddedStruct.VirtualMachine
		*o = VirtualizationVmwareVirtualMachineSnapshot(varVirtualizationVmwareVirtualMachineSnapshot)
	} else {
		return err
	}

	varVirtualizationVmwareVirtualMachineSnapshot := _VirtualizationVmwareVirtualMachineSnapshot{}

	err = json.Unmarshal(bytes, &varVirtualizationVmwareVirtualMachineSnapshot)
	if err == nil {
		o.VirtualizationBaseVirtualMachineSnapshot = varVirtualizationVmwareVirtualMachineSnapshot.VirtualizationBaseVirtualMachineSnapshot
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreationTime")
		delete(additionalProperties, "CurrentSnapshot")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Golden")
		delete(additionalProperties, "Key")
		delete(additionalProperties, "PredecessorId")
		delete(additionalProperties, "Quiesced")
		delete(additionalProperties, "RefValue")
		delete(additionalProperties, "SnapshotSize")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectVirtualizationBaseVirtualMachineSnapshot := reflect.ValueOf(o.VirtualizationBaseVirtualMachineSnapshot)
		for i := 0; i < reflectVirtualizationBaseVirtualMachineSnapshot.Type().NumField(); i++ {
			t := reflectVirtualizationBaseVirtualMachineSnapshot.Type().Field(i)

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

type NullableVirtualizationVmwareVirtualMachineSnapshot struct {
	value *VirtualizationVmwareVirtualMachineSnapshot
	isSet bool
}

func (v NullableVirtualizationVmwareVirtualMachineSnapshot) Get() *VirtualizationVmwareVirtualMachineSnapshot {
	return v.value
}

func (v *NullableVirtualizationVmwareVirtualMachineSnapshot) Set(val *VirtualizationVmwareVirtualMachineSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwareVirtualMachineSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwareVirtualMachineSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwareVirtualMachineSnapshot(val *VirtualizationVmwareVirtualMachineSnapshot) *NullableVirtualizationVmwareVirtualMachineSnapshot {
	return &NullableVirtualizationVmwareVirtualMachineSnapshot{value: val, isSet: true}
}

func (v NullableVirtualizationVmwareVirtualMachineSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwareVirtualMachineSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
