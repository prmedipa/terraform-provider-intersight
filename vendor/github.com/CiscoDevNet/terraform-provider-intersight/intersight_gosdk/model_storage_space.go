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

// checks if the StorageSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageSpace{}

// StorageSpace Some attributes for the free space and the LDEV defined in the specified external parity group.
type StorageSpace struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Starting location of the LBA of the partition in the external parity group (in a multiple of 512 bytes).
	LbaLocation *string `json:"LbaLocation,omitempty"`
	// Size of the partition in the external parity group (in a multiple of 512 bytes).
	LbaSize *string `json:"LbaSize,omitempty"`
	// LDEV number, property is output only if LDEV has been implemented.
	LdevId *string `json:"LdevId,omitempty"`
	// Number of a partition created as a result of partitioning of an external parity group.
	PartitionNumber *int64 `json:"PartitionNumber,omitempty"`
	// Status about LDEV for partition.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageSpace StorageSpace

// NewStorageSpace instantiates a new StorageSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSpace(classId string, objectType string) *StorageSpace {
	this := StorageSpace{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageSpaceWithDefaults instantiates a new StorageSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSpaceWithDefaults() *StorageSpace {
	this := StorageSpace{}
	var classId string = "storage.Space"
	this.ClassId = classId
	var objectType string = "storage.Space"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageSpace) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageSpace) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.Space" of the ClassId field.
func (o *StorageSpace) GetDefaultClassId() interface{} {
	return "storage.Space"
}

// GetObjectType returns the ObjectType field value
func (o *StorageSpace) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageSpace) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.Space" of the ObjectType field.
func (o *StorageSpace) GetDefaultObjectType() interface{} {
	return "storage.Space"
}

// GetLbaLocation returns the LbaLocation field value if set, zero value otherwise.
func (o *StorageSpace) GetLbaLocation() string {
	if o == nil || IsNil(o.LbaLocation) {
		var ret string
		return ret
	}
	return *o.LbaLocation
}

// GetLbaLocationOk returns a tuple with the LbaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetLbaLocationOk() (*string, bool) {
	if o == nil || IsNil(o.LbaLocation) {
		return nil, false
	}
	return o.LbaLocation, true
}

// HasLbaLocation returns a boolean if a field has been set.
func (o *StorageSpace) HasLbaLocation() bool {
	if o != nil && !IsNil(o.LbaLocation) {
		return true
	}

	return false
}

// SetLbaLocation gets a reference to the given string and assigns it to the LbaLocation field.
func (o *StorageSpace) SetLbaLocation(v string) {
	o.LbaLocation = &v
}

// GetLbaSize returns the LbaSize field value if set, zero value otherwise.
func (o *StorageSpace) GetLbaSize() string {
	if o == nil || IsNil(o.LbaSize) {
		var ret string
		return ret
	}
	return *o.LbaSize
}

// GetLbaSizeOk returns a tuple with the LbaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetLbaSizeOk() (*string, bool) {
	if o == nil || IsNil(o.LbaSize) {
		return nil, false
	}
	return o.LbaSize, true
}

// HasLbaSize returns a boolean if a field has been set.
func (o *StorageSpace) HasLbaSize() bool {
	if o != nil && !IsNil(o.LbaSize) {
		return true
	}

	return false
}

// SetLbaSize gets a reference to the given string and assigns it to the LbaSize field.
func (o *StorageSpace) SetLbaSize(v string) {
	o.LbaSize = &v
}

// GetLdevId returns the LdevId field value if set, zero value otherwise.
func (o *StorageSpace) GetLdevId() string {
	if o == nil || IsNil(o.LdevId) {
		var ret string
		return ret
	}
	return *o.LdevId
}

// GetLdevIdOk returns a tuple with the LdevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetLdevIdOk() (*string, bool) {
	if o == nil || IsNil(o.LdevId) {
		return nil, false
	}
	return o.LdevId, true
}

// HasLdevId returns a boolean if a field has been set.
func (o *StorageSpace) HasLdevId() bool {
	if o != nil && !IsNil(o.LdevId) {
		return true
	}

	return false
}

// SetLdevId gets a reference to the given string and assigns it to the LdevId field.
func (o *StorageSpace) SetLdevId(v string) {
	o.LdevId = &v
}

// GetPartitionNumber returns the PartitionNumber field value if set, zero value otherwise.
func (o *StorageSpace) GetPartitionNumber() int64 {
	if o == nil || IsNil(o.PartitionNumber) {
		var ret int64
		return ret
	}
	return *o.PartitionNumber
}

// GetPartitionNumberOk returns a tuple with the PartitionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetPartitionNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.PartitionNumber) {
		return nil, false
	}
	return o.PartitionNumber, true
}

// HasPartitionNumber returns a boolean if a field has been set.
func (o *StorageSpace) HasPartitionNumber() bool {
	if o != nil && !IsNil(o.PartitionNumber) {
		return true
	}

	return false
}

// SetPartitionNumber gets a reference to the given int64 and assigns it to the PartitionNumber field.
func (o *StorageSpace) SetPartitionNumber(v int64) {
	o.PartitionNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StorageSpace) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSpace) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StorageSpace) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StorageSpace) SetStatus(v string) {
	o.Status = &v
}

func (o StorageSpace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.LbaLocation) {
		toSerialize["LbaLocation"] = o.LbaLocation
	}
	if !IsNil(o.LbaSize) {
		toSerialize["LbaSize"] = o.LbaSize
	}
	if !IsNil(o.LdevId) {
		toSerialize["LdevId"] = o.LdevId
	}
	if !IsNil(o.PartitionNumber) {
		toSerialize["PartitionNumber"] = o.PartitionNumber
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageSpace) UnmarshalJSON(data []byte) (err error) {
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
	type StorageSpaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Starting location of the LBA of the partition in the external parity group (in a multiple of 512 bytes).
		LbaLocation *string `json:"LbaLocation,omitempty"`
		// Size of the partition in the external parity group (in a multiple of 512 bytes).
		LbaSize *string `json:"LbaSize,omitempty"`
		// LDEV number, property is output only if LDEV has been implemented.
		LdevId *string `json:"LdevId,omitempty"`
		// Number of a partition created as a result of partitioning of an external parity group.
		PartitionNumber *int64 `json:"PartitionNumber,omitempty"`
		// Status about LDEV for partition.
		Status *string `json:"Status,omitempty"`
	}

	varStorageSpaceWithoutEmbeddedStruct := StorageSpaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageSpaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageSpace := _StorageSpace{}
		varStorageSpace.ClassId = varStorageSpaceWithoutEmbeddedStruct.ClassId
		varStorageSpace.ObjectType = varStorageSpaceWithoutEmbeddedStruct.ObjectType
		varStorageSpace.LbaLocation = varStorageSpaceWithoutEmbeddedStruct.LbaLocation
		varStorageSpace.LbaSize = varStorageSpaceWithoutEmbeddedStruct.LbaSize
		varStorageSpace.LdevId = varStorageSpaceWithoutEmbeddedStruct.LdevId
		varStorageSpace.PartitionNumber = varStorageSpaceWithoutEmbeddedStruct.PartitionNumber
		varStorageSpace.Status = varStorageSpaceWithoutEmbeddedStruct.Status
		*o = StorageSpace(varStorageSpace)
	} else {
		return err
	}

	varStorageSpace := _StorageSpace{}

	err = json.Unmarshal(data, &varStorageSpace)
	if err == nil {
		o.MoBaseComplexType = varStorageSpace.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LbaLocation")
		delete(additionalProperties, "LbaSize")
		delete(additionalProperties, "LdevId")
		delete(additionalProperties, "PartitionNumber")
		delete(additionalProperties, "Status")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableStorageSpace struct {
	value *StorageSpace
	isSet bool
}

func (v NullableStorageSpace) Get() *StorageSpace {
	return v.value
}

func (v *NullableStorageSpace) Set(val *StorageSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSpace(val *StorageSpace) *NullableStorageSpace {
	return &NullableStorageSpace{value: val, isSet: true}
}

func (v NullableStorageSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
