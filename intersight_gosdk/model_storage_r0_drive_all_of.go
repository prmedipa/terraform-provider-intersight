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

// StorageR0DriveAllOf Definition of the list of properties defined in 'storage.R0Drive', excluding properties defined in parent classes.
type StorageR0DriveAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The set of drive slots where RAID0 virtual drives must be created.
	DriveSlots *string `json:"DriveSlots,omitempty"`
	// The list of drive slots where RAID0 virtual drives must be created (comma seperated).
	DriveSlotsList *string `json:"DriveSlotsList,omitempty"`
	// If enabled, this will create a RAID0 virtual drive per disk and encompassing the whole disk.
	Enable               *bool                             `json:"Enable,omitempty"`
	VirtualDrivePolicy   NullableStorageVirtualDrivePolicy `json:"VirtualDrivePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageR0DriveAllOf StorageR0DriveAllOf

// NewStorageR0DriveAllOf instantiates a new StorageR0DriveAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageR0DriveAllOf(classId string, objectType string) *StorageR0DriveAllOf {
	this := StorageR0DriveAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewStorageR0DriveAllOfWithDefaults instantiates a new StorageR0DriveAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageR0DriveAllOfWithDefaults() *StorageR0DriveAllOf {
	this := StorageR0DriveAllOf{}
	var classId string = "storage.R0Drive"
	this.ClassId = classId
	var objectType string = "storage.R0Drive"
	this.ObjectType = objectType
	var enable bool = false
	this.Enable = &enable
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageR0DriveAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageR0DriveAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageR0DriveAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageR0DriveAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageR0DriveAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageR0DriveAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDriveSlots returns the DriveSlots field value if set, zero value otherwise.
func (o *StorageR0DriveAllOf) GetDriveSlots() string {
	if o == nil || o.DriveSlots == nil {
		var ret string
		return ret
	}
	return *o.DriveSlots
}

// GetDriveSlotsOk returns a tuple with the DriveSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0DriveAllOf) GetDriveSlotsOk() (*string, bool) {
	if o == nil || o.DriveSlots == nil {
		return nil, false
	}
	return o.DriveSlots, true
}

// HasDriveSlots returns a boolean if a field has been set.
func (o *StorageR0DriveAllOf) HasDriveSlots() bool {
	if o != nil && o.DriveSlots != nil {
		return true
	}

	return false
}

// SetDriveSlots gets a reference to the given string and assigns it to the DriveSlots field.
func (o *StorageR0DriveAllOf) SetDriveSlots(v string) {
	o.DriveSlots = &v
}

// GetDriveSlotsList returns the DriveSlotsList field value if set, zero value otherwise.
func (o *StorageR0DriveAllOf) GetDriveSlotsList() string {
	if o == nil || o.DriveSlotsList == nil {
		var ret string
		return ret
	}
	return *o.DriveSlotsList
}

// GetDriveSlotsListOk returns a tuple with the DriveSlotsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0DriveAllOf) GetDriveSlotsListOk() (*string, bool) {
	if o == nil || o.DriveSlotsList == nil {
		return nil, false
	}
	return o.DriveSlotsList, true
}

// HasDriveSlotsList returns a boolean if a field has been set.
func (o *StorageR0DriveAllOf) HasDriveSlotsList() bool {
	if o != nil && o.DriveSlotsList != nil {
		return true
	}

	return false
}

// SetDriveSlotsList gets a reference to the given string and assigns it to the DriveSlotsList field.
func (o *StorageR0DriveAllOf) SetDriveSlotsList(v string) {
	o.DriveSlotsList = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *StorageR0DriveAllOf) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageR0DriveAllOf) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *StorageR0DriveAllOf) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *StorageR0DriveAllOf) SetEnable(v bool) {
	o.Enable = &v
}

// GetVirtualDrivePolicy returns the VirtualDrivePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageR0DriveAllOf) GetVirtualDrivePolicy() StorageVirtualDrivePolicy {
	if o == nil || o.VirtualDrivePolicy.Get() == nil {
		var ret StorageVirtualDrivePolicy
		return ret
	}
	return *o.VirtualDrivePolicy.Get()
}

// GetVirtualDrivePolicyOk returns a tuple with the VirtualDrivePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageR0DriveAllOf) GetVirtualDrivePolicyOk() (*StorageVirtualDrivePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualDrivePolicy.Get(), o.VirtualDrivePolicy.IsSet()
}

// HasVirtualDrivePolicy returns a boolean if a field has been set.
func (o *StorageR0DriveAllOf) HasVirtualDrivePolicy() bool {
	if o != nil && o.VirtualDrivePolicy.IsSet() {
		return true
	}

	return false
}

// SetVirtualDrivePolicy gets a reference to the given NullableStorageVirtualDrivePolicy and assigns it to the VirtualDrivePolicy field.
func (o *StorageR0DriveAllOf) SetVirtualDrivePolicy(v StorageVirtualDrivePolicy) {
	o.VirtualDrivePolicy.Set(&v)
}

// SetVirtualDrivePolicyNil sets the value for VirtualDrivePolicy to be an explicit nil
func (o *StorageR0DriveAllOf) SetVirtualDrivePolicyNil() {
	o.VirtualDrivePolicy.Set(nil)
}

// UnsetVirtualDrivePolicy ensures that no value is present for VirtualDrivePolicy, not even an explicit nil
func (o *StorageR0DriveAllOf) UnsetVirtualDrivePolicy() {
	o.VirtualDrivePolicy.Unset()
}

func (o StorageR0DriveAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DriveSlots != nil {
		toSerialize["DriveSlots"] = o.DriveSlots
	}
	if o.DriveSlotsList != nil {
		toSerialize["DriveSlotsList"] = o.DriveSlotsList
	}
	if o.Enable != nil {
		toSerialize["Enable"] = o.Enable
	}
	if o.VirtualDrivePolicy.IsSet() {
		toSerialize["VirtualDrivePolicy"] = o.VirtualDrivePolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageR0DriveAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageR0DriveAllOf := _StorageR0DriveAllOf{}

	if err = json.Unmarshal(bytes, &varStorageR0DriveAllOf); err == nil {
		*o = StorageR0DriveAllOf(varStorageR0DriveAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DriveSlots")
		delete(additionalProperties, "DriveSlotsList")
		delete(additionalProperties, "Enable")
		delete(additionalProperties, "VirtualDrivePolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageR0DriveAllOf struct {
	value *StorageR0DriveAllOf
	isSet bool
}

func (v NullableStorageR0DriveAllOf) Get() *StorageR0DriveAllOf {
	return v.value
}

func (v *NullableStorageR0DriveAllOf) Set(val *StorageR0DriveAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageR0DriveAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageR0DriveAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageR0DriveAllOf(val *StorageR0DriveAllOf) *NullableStorageR0DriveAllOf {
	return &NullableStorageR0DriveAllOf{value: val, isSet: true}
}

func (v NullableStorageR0DriveAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageR0DriveAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
