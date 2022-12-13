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

// OsVirtualDrive Virtual Drive target to do Operating System installation on.
type OsVirtualDrive struct {
	OsInstallTarget
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Virtual Drive ID to be used as Install Target.
	Id *string `json:"Id,omitempty"`
	// The Virtual Drive Name to be used as Install Target.
	Name *string `json:"Name,omitempty"`
	// The SlotID of the Storage Controller associated to the virtual drive.
	StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OsVirtualDrive OsVirtualDrive

// NewOsVirtualDrive instantiates a new OsVirtualDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsVirtualDrive(classId string, objectType string) *OsVirtualDrive {
	this := OsVirtualDrive{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOsVirtualDriveWithDefaults instantiates a new OsVirtualDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsVirtualDriveWithDefaults() *OsVirtualDrive {
	this := OsVirtualDrive{}
	var classId string = "os.VirtualDrive"
	this.ClassId = classId
	var objectType string = "os.VirtualDrive"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsVirtualDrive) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsVirtualDrive) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsVirtualDrive) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *OsVirtualDrive) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsVirtualDrive) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsVirtualDrive) SetObjectType(v string) {
	o.ObjectType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OsVirtualDrive) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDrive) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OsVirtualDrive) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OsVirtualDrive) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsVirtualDrive) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDrive) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsVirtualDrive) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsVirtualDrive) SetName(v string) {
	o.Name = &v
}

// GetStorageControllerSlotId returns the StorageControllerSlotId field value if set, zero value otherwise.
func (o *OsVirtualDrive) GetStorageControllerSlotId() string {
	if o == nil || o.StorageControllerSlotId == nil {
		var ret string
		return ret
	}
	return *o.StorageControllerSlotId
}

// GetStorageControllerSlotIdOk returns a tuple with the StorageControllerSlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsVirtualDrive) GetStorageControllerSlotIdOk() (*string, bool) {
	if o == nil || o.StorageControllerSlotId == nil {
		return nil, false
	}
	return o.StorageControllerSlotId, true
}

// HasStorageControllerSlotId returns a boolean if a field has been set.
func (o *OsVirtualDrive) HasStorageControllerSlotId() bool {
	if o != nil && o.StorageControllerSlotId != nil {
		return true
	}

	return false
}

// SetStorageControllerSlotId gets a reference to the given string and assigns it to the StorageControllerSlotId field.
func (o *OsVirtualDrive) SetStorageControllerSlotId(v string) {
	o.StorageControllerSlotId = &v
}

func (o OsVirtualDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedOsInstallTarget, errOsInstallTarget := json.Marshal(o.OsInstallTarget)
	if errOsInstallTarget != nil {
		return []byte{}, errOsInstallTarget
	}
	errOsInstallTarget = json.Unmarshal([]byte(serializedOsInstallTarget), &toSerialize)
	if errOsInstallTarget != nil {
		return []byte{}, errOsInstallTarget
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StorageControllerSlotId != nil {
		toSerialize["StorageControllerSlotId"] = o.StorageControllerSlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OsVirtualDrive) UnmarshalJSON(bytes []byte) (err error) {
	type OsVirtualDriveWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Virtual Drive ID to be used as Install Target.
		Id *string `json:"Id,omitempty"`
		// The Virtual Drive Name to be used as Install Target.
		Name *string `json:"Name,omitempty"`
		// The SlotID of the Storage Controller associated to the virtual drive.
		StorageControllerSlotId *string `json:"StorageControllerSlotId,omitempty"`
	}

	varOsVirtualDriveWithoutEmbeddedStruct := OsVirtualDriveWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOsVirtualDriveWithoutEmbeddedStruct)
	if err == nil {
		varOsVirtualDrive := _OsVirtualDrive{}
		varOsVirtualDrive.ClassId = varOsVirtualDriveWithoutEmbeddedStruct.ClassId
		varOsVirtualDrive.ObjectType = varOsVirtualDriveWithoutEmbeddedStruct.ObjectType
		varOsVirtualDrive.Id = varOsVirtualDriveWithoutEmbeddedStruct.Id
		varOsVirtualDrive.Name = varOsVirtualDriveWithoutEmbeddedStruct.Name
		varOsVirtualDrive.StorageControllerSlotId = varOsVirtualDriveWithoutEmbeddedStruct.StorageControllerSlotId
		*o = OsVirtualDrive(varOsVirtualDrive)
	} else {
		return err
	}

	varOsVirtualDrive := _OsVirtualDrive{}

	err = json.Unmarshal(bytes, &varOsVirtualDrive)
	if err == nil {
		o.OsInstallTarget = varOsVirtualDrive.OsInstallTarget
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "StorageControllerSlotId")

		// remove fields from embedded structs
		reflectOsInstallTarget := reflect.ValueOf(o.OsInstallTarget)
		for i := 0; i < reflectOsInstallTarget.Type().NumField(); i++ {
			t := reflectOsInstallTarget.Type().Field(i)

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

type NullableOsVirtualDrive struct {
	value *OsVirtualDrive
	isSet bool
}

func (v NullableOsVirtualDrive) Get() *OsVirtualDrive {
	return v.value
}

func (v *NullableOsVirtualDrive) Set(val *OsVirtualDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableOsVirtualDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableOsVirtualDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsVirtualDrive(val *OsVirtualDrive) *NullableOsVirtualDrive {
	return &NullableOsVirtualDrive{value: val, isSet: true}
}

func (v NullableOsVirtualDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsVirtualDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
