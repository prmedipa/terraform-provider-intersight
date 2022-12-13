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
	"reflect"
	"strings"
)

// StorageVirtualDriveIdentity Identity object that uniquely represents a VirtualDrive object under a Server Profile.
type StorageVirtualDriveIdentity struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The VirtualDrive Name which belongs to the Storage VirtualDrive.
	Name                 *string                           `json:"Name,omitempty"`
	ServerProfile        *ServerProfileRelationship        `json:"ServerProfile,omitempty"`
	StoragePolicy        *StorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
	VirtualDrive         *StorageVirtualDriveRelationship  `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveIdentity StorageVirtualDriveIdentity

// NewStorageVirtualDriveIdentity instantiates a new StorageVirtualDriveIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveIdentity(classId string, objectType string) *StorageVirtualDriveIdentity {
	this := StorageVirtualDriveIdentity{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveIdentityWithDefaults instantiates a new StorageVirtualDriveIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveIdentityWithDefaults() *StorageVirtualDriveIdentity {
	this := StorageVirtualDriveIdentity{}
	var classId string = "storage.VirtualDriveIdentity"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveIdentity) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveIdentity) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveIdentity) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveIdentity) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentity) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentity) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveIdentity) SetName(v string) {
	o.Name = &v
}

// GetServerProfile returns the ServerProfile field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentity) GetServerProfile() ServerProfileRelationship {
	if o == nil || o.ServerProfile == nil {
		var ret ServerProfileRelationship
		return ret
	}
	return *o.ServerProfile
}

// GetServerProfileOk returns a tuple with the ServerProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetServerProfileOk() (*ServerProfileRelationship, bool) {
	if o == nil || o.ServerProfile == nil {
		return nil, false
	}
	return o.ServerProfile, true
}

// HasServerProfile returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentity) HasServerProfile() bool {
	if o != nil && o.ServerProfile != nil {
		return true
	}

	return false
}

// SetServerProfile gets a reference to the given ServerProfileRelationship and assigns it to the ServerProfile field.
func (o *StorageVirtualDriveIdentity) SetServerProfile(v ServerProfileRelationship) {
	o.ServerProfile = &v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentity) GetStoragePolicy() StorageStoragePolicyRelationship {
	if o == nil || o.StoragePolicy == nil {
		var ret StorageStoragePolicyRelationship
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetStoragePolicyOk() (*StorageStoragePolicyRelationship, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentity) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given StorageStoragePolicyRelationship and assigns it to the StoragePolicy field.
func (o *StorageVirtualDriveIdentity) SetStoragePolicy(v StorageStoragePolicyRelationship) {
	o.StoragePolicy = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveIdentity) GetVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || o.VirtualDrive == nil {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveIdentity) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveIdentity) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveIdentity) SetVirtualDrive(v StorageVirtualDriveRelationship) {
	o.VirtualDrive = &v
}

func (o StorageVirtualDriveIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ServerProfile != nil {
		toSerialize["ServerProfile"] = o.ServerProfile
	}
	if o.StoragePolicy != nil {
		toSerialize["StoragePolicy"] = o.StoragePolicy
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveIdentity) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVirtualDriveIdentityWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The VirtualDrive Name which belongs to the Storage VirtualDrive.
		Name          *string                           `json:"Name,omitempty"`
		ServerProfile *ServerProfileRelationship        `json:"ServerProfile,omitempty"`
		StoragePolicy *StorageStoragePolicyRelationship `json:"StoragePolicy,omitempty"`
		VirtualDrive  *StorageVirtualDriveRelationship  `json:"VirtualDrive,omitempty"`
	}

	varStorageVirtualDriveIdentityWithoutEmbeddedStruct := StorageVirtualDriveIdentityWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveIdentityWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDriveIdentity := _StorageVirtualDriveIdentity{}
		varStorageVirtualDriveIdentity.ClassId = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.ClassId
		varStorageVirtualDriveIdentity.ObjectType = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDriveIdentity.Name = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.Name
		varStorageVirtualDriveIdentity.ServerProfile = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.ServerProfile
		varStorageVirtualDriveIdentity.StoragePolicy = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.StoragePolicy
		varStorageVirtualDriveIdentity.VirtualDrive = varStorageVirtualDriveIdentityWithoutEmbeddedStruct.VirtualDrive
		*o = StorageVirtualDriveIdentity(varStorageVirtualDriveIdentity)
	} else {
		return err
	}

	varStorageVirtualDriveIdentity := _StorageVirtualDriveIdentity{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveIdentity)
	if err == nil {
		o.MoBaseMo = varStorageVirtualDriveIdentity.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ServerProfile")
		delete(additionalProperties, "StoragePolicy")
		delete(additionalProperties, "VirtualDrive")

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

type NullableStorageVirtualDriveIdentity struct {
	value *StorageVirtualDriveIdentity
	isSet bool
}

func (v NullableStorageVirtualDriveIdentity) Get() *StorageVirtualDriveIdentity {
	return v.value
}

func (v *NullableStorageVirtualDriveIdentity) Set(val *StorageVirtualDriveIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveIdentity(val *StorageVirtualDriveIdentity) *NullableStorageVirtualDriveIdentity {
	return &NullableStorageVirtualDriveIdentity{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
