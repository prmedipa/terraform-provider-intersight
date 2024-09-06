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

// checks if the StorageHitachiExternalPathGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageHitachiExternalPathGroup{}

// StorageHitachiExternalPathGroup A external path group in Hitachi storage array.
type StorageHitachiExternalPathGroup struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                       `json:"ObjectType"`
	ExternalParityGroups []StorageExternalParityGroup `json:"ExternalParityGroups,omitempty"`
	ExternalPaths        []StorageExternalPath        `json:"ExternalPaths,omitempty"`
	// Product ID of the external storage system.
	ExternalProductId *string `json:"ExternalProductId,omitempty"`
	// Serial number of the external storage system.
	ExternalSerialNumber *string `json:"ExternalSerialNumber,omitempty"`
	// External path group number.
	Name                 *string                                     `json:"Name,omitempty"`
	Array                NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiExternalPathGroup StorageHitachiExternalPathGroup

// NewStorageHitachiExternalPathGroup instantiates a new StorageHitachiExternalPathGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiExternalPathGroup(classId string, objectType string) *StorageHitachiExternalPathGroup {
	this := StorageHitachiExternalPathGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiExternalPathGroupWithDefaults instantiates a new StorageHitachiExternalPathGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiExternalPathGroupWithDefaults() *StorageHitachiExternalPathGroup {
	this := StorageHitachiExternalPathGroup{}
	var classId string = "storage.HitachiExternalPathGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiExternalPathGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiExternalPathGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiExternalPathGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.HitachiExternalPathGroup" of the ClassId field.
func (o *StorageHitachiExternalPathGroup) GetDefaultClassId() interface{} {
	return "storage.HitachiExternalPathGroup"
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiExternalPathGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiExternalPathGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.HitachiExternalPathGroup" of the ObjectType field.
func (o *StorageHitachiExternalPathGroup) GetDefaultObjectType() interface{} {
	return "storage.HitachiExternalPathGroup"
}

// GetExternalParityGroups returns the ExternalParityGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroup) GetExternalParityGroups() []StorageExternalParityGroup {
	if o == nil {
		var ret []StorageExternalParityGroup
		return ret
	}
	return o.ExternalParityGroups
}

// GetExternalParityGroupsOk returns a tuple with the ExternalParityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroup) GetExternalParityGroupsOk() ([]StorageExternalParityGroup, bool) {
	if o == nil || IsNil(o.ExternalParityGroups) {
		return nil, false
	}
	return o.ExternalParityGroups, true
}

// HasExternalParityGroups returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasExternalParityGroups() bool {
	if o != nil && !IsNil(o.ExternalParityGroups) {
		return true
	}

	return false
}

// SetExternalParityGroups gets a reference to the given []StorageExternalParityGroup and assigns it to the ExternalParityGroups field.
func (o *StorageHitachiExternalPathGroup) SetExternalParityGroups(v []StorageExternalParityGroup) {
	o.ExternalParityGroups = v
}

// GetExternalPaths returns the ExternalPaths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroup) GetExternalPaths() []StorageExternalPath {
	if o == nil {
		var ret []StorageExternalPath
		return ret
	}
	return o.ExternalPaths
}

// GetExternalPathsOk returns a tuple with the ExternalPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroup) GetExternalPathsOk() ([]StorageExternalPath, bool) {
	if o == nil || IsNil(o.ExternalPaths) {
		return nil, false
	}
	return o.ExternalPaths, true
}

// HasExternalPaths returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasExternalPaths() bool {
	if o != nil && !IsNil(o.ExternalPaths) {
		return true
	}

	return false
}

// SetExternalPaths gets a reference to the given []StorageExternalPath and assigns it to the ExternalPaths field.
func (o *StorageHitachiExternalPathGroup) SetExternalPaths(v []StorageExternalPath) {
	o.ExternalPaths = v
}

// GetExternalProductId returns the ExternalProductId field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroup) GetExternalProductId() string {
	if o == nil || IsNil(o.ExternalProductId) {
		var ret string
		return ret
	}
	return *o.ExternalProductId
}

// GetExternalProductIdOk returns a tuple with the ExternalProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroup) GetExternalProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalProductId) {
		return nil, false
	}
	return o.ExternalProductId, true
}

// HasExternalProductId returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasExternalProductId() bool {
	if o != nil && !IsNil(o.ExternalProductId) {
		return true
	}

	return false
}

// SetExternalProductId gets a reference to the given string and assigns it to the ExternalProductId field.
func (o *StorageHitachiExternalPathGroup) SetExternalProductId(v string) {
	o.ExternalProductId = &v
}

// GetExternalSerialNumber returns the ExternalSerialNumber field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroup) GetExternalSerialNumber() string {
	if o == nil || IsNil(o.ExternalSerialNumber) {
		var ret string
		return ret
	}
	return *o.ExternalSerialNumber
}

// GetExternalSerialNumberOk returns a tuple with the ExternalSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroup) GetExternalSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSerialNumber) {
		return nil, false
	}
	return o.ExternalSerialNumber, true
}

// HasExternalSerialNumber returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasExternalSerialNumber() bool {
	if o != nil && !IsNil(o.ExternalSerialNumber) {
		return true
	}

	return false
}

// SetExternalSerialNumber gets a reference to the given string and assigns it to the ExternalSerialNumber field.
func (o *StorageHitachiExternalPathGroup) SetExternalSerialNumber(v string) {
	o.ExternalSerialNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageHitachiExternalPathGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiExternalPathGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageHitachiExternalPathGroup) SetName(v string) {
	o.Name = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroup) GetArray() StorageHitachiArrayRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroup) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiExternalPathGroup) SetArray(v StorageHitachiArrayRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageHitachiExternalPathGroup) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageHitachiExternalPathGroup) UnsetArray() {
	o.Array.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageHitachiExternalPathGroup) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageHitachiExternalPathGroup) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiExternalPathGroup) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiExternalPathGroup) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageHitachiExternalPathGroup) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageHitachiExternalPathGroup) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageHitachiExternalPathGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageHitachiExternalPathGroup) ToMap() (map[string]interface{}, error) {
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
	if o.ExternalParityGroups != nil {
		toSerialize["ExternalParityGroups"] = o.ExternalParityGroups
	}
	if o.ExternalPaths != nil {
		toSerialize["ExternalPaths"] = o.ExternalPaths
	}
	if !IsNil(o.ExternalProductId) {
		toSerialize["ExternalProductId"] = o.ExternalProductId
	}
	if !IsNil(o.ExternalSerialNumber) {
		toSerialize["ExternalSerialNumber"] = o.ExternalSerialNumber
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageHitachiExternalPathGroup) UnmarshalJSON(data []byte) (err error) {
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
	type StorageHitachiExternalPathGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                       `json:"ObjectType"`
		ExternalParityGroups []StorageExternalParityGroup `json:"ExternalParityGroups,omitempty"`
		ExternalPaths        []StorageExternalPath        `json:"ExternalPaths,omitempty"`
		// Product ID of the external storage system.
		ExternalProductId *string `json:"ExternalProductId,omitempty"`
		// Serial number of the external storage system.
		ExternalSerialNumber *string `json:"ExternalSerialNumber,omitempty"`
		// External path group number.
		Name             *string                                     `json:"Name,omitempty"`
		Array            NullableStorageHitachiArrayRelationship     `json:"Array,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageHitachiExternalPathGroupWithoutEmbeddedStruct := StorageHitachiExternalPathGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageHitachiExternalPathGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageHitachiExternalPathGroup := _StorageHitachiExternalPathGroup{}
		varStorageHitachiExternalPathGroup.ClassId = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ClassId
		varStorageHitachiExternalPathGroup.ObjectType = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ObjectType
		varStorageHitachiExternalPathGroup.ExternalParityGroups = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ExternalParityGroups
		varStorageHitachiExternalPathGroup.ExternalPaths = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ExternalPaths
		varStorageHitachiExternalPathGroup.ExternalProductId = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ExternalProductId
		varStorageHitachiExternalPathGroup.ExternalSerialNumber = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.ExternalSerialNumber
		varStorageHitachiExternalPathGroup.Name = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.Name
		varStorageHitachiExternalPathGroup.Array = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.Array
		varStorageHitachiExternalPathGroup.RegisteredDevice = varStorageHitachiExternalPathGroupWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageHitachiExternalPathGroup(varStorageHitachiExternalPathGroup)
	} else {
		return err
	}

	varStorageHitachiExternalPathGroup := _StorageHitachiExternalPathGroup{}

	err = json.Unmarshal(data, &varStorageHitachiExternalPathGroup)
	if err == nil {
		o.MoBaseMo = varStorageHitachiExternalPathGroup.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ExternalParityGroups")
		delete(additionalProperties, "ExternalPaths")
		delete(additionalProperties, "ExternalProductId")
		delete(additionalProperties, "ExternalSerialNumber")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableStorageHitachiExternalPathGroup struct {
	value *StorageHitachiExternalPathGroup
	isSet bool
}

func (v NullableStorageHitachiExternalPathGroup) Get() *StorageHitachiExternalPathGroup {
	return v.value
}

func (v *NullableStorageHitachiExternalPathGroup) Set(val *StorageHitachiExternalPathGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiExternalPathGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiExternalPathGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiExternalPathGroup(val *StorageHitachiExternalPathGroup) *NullableStorageHitachiExternalPathGroup {
	return &NullableStorageHitachiExternalPathGroup{value: val, isSet: true}
}

func (v NullableStorageHitachiExternalPathGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiExternalPathGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
