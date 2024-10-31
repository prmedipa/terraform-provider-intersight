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

// checks if the StorageNetAppClusterSnapshotPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppClusterSnapshotPolicy{}

// StorageNetAppClusterSnapshotPolicy NetApp Snapshot policy that is scoped to a cluster. The policy controls the behavior and schedule of snapshots when applied to a volume.
type StorageNetAppClusterSnapshotPolicy struct {
	StorageNetAppBaseSnapshotPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                   `json:"ObjectType"`
	Array                NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppClusterSnapshotPolicy StorageNetAppClusterSnapshotPolicy

// NewStorageNetAppClusterSnapshotPolicy instantiates a new StorageNetAppClusterSnapshotPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppClusterSnapshotPolicy(classId string, objectType string) *StorageNetAppClusterSnapshotPolicy {
	this := StorageNetAppClusterSnapshotPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppClusterSnapshotPolicyWithDefaults instantiates a new StorageNetAppClusterSnapshotPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppClusterSnapshotPolicyWithDefaults() *StorageNetAppClusterSnapshotPolicy {
	this := StorageNetAppClusterSnapshotPolicy{}
	var classId string = "storage.NetAppClusterSnapshotPolicy"
	this.ClassId = classId
	var objectType string = "storage.NetAppClusterSnapshotPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppClusterSnapshotPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterSnapshotPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppClusterSnapshotPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppClusterSnapshotPolicy" of the ClassId field.
func (o *StorageNetAppClusterSnapshotPolicy) GetDefaultClassId() interface{} {
	return "storage.NetAppClusterSnapshotPolicy"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppClusterSnapshotPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppClusterSnapshotPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppClusterSnapshotPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppClusterSnapshotPolicy" of the ObjectType field.
func (o *StorageNetAppClusterSnapshotPolicy) GetDefaultObjectType() interface{} {
	return "storage.NetAppClusterSnapshotPolicy"
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppClusterSnapshotPolicy) GetArray() StorageNetAppClusterRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppClusterSnapshotPolicy) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppClusterSnapshotPolicy) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppClusterSnapshotPolicy) SetArray(v StorageNetAppClusterRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageNetAppClusterSnapshotPolicy) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageNetAppClusterSnapshotPolicy) UnsetArray() {
	o.Array.Unset()
}

func (o StorageNetAppClusterSnapshotPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppClusterSnapshotPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseSnapshotPolicy, errStorageNetAppBaseSnapshotPolicy := json.Marshal(o.StorageNetAppBaseSnapshotPolicy)
	if errStorageNetAppBaseSnapshotPolicy != nil {
		return map[string]interface{}{}, errStorageNetAppBaseSnapshotPolicy
	}
	errStorageNetAppBaseSnapshotPolicy = json.Unmarshal([]byte(serializedStorageNetAppBaseSnapshotPolicy), &toSerialize)
	if errStorageNetAppBaseSnapshotPolicy != nil {
		return map[string]interface{}{}, errStorageNetAppBaseSnapshotPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppClusterSnapshotPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                                   `json:"ObjectType"`
		Array      NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct := StorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppClusterSnapshotPolicy := _StorageNetAppClusterSnapshotPolicy{}
		varStorageNetAppClusterSnapshotPolicy.ClassId = varStorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct.ClassId
		varStorageNetAppClusterSnapshotPolicy.ObjectType = varStorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct.ObjectType
		varStorageNetAppClusterSnapshotPolicy.Array = varStorageNetAppClusterSnapshotPolicyWithoutEmbeddedStruct.Array
		*o = StorageNetAppClusterSnapshotPolicy(varStorageNetAppClusterSnapshotPolicy)
	} else {
		return err
	}

	varStorageNetAppClusterSnapshotPolicy := _StorageNetAppClusterSnapshotPolicy{}

	err = json.Unmarshal(data, &varStorageNetAppClusterSnapshotPolicy)
	if err == nil {
		o.StorageNetAppBaseSnapshotPolicy = varStorageNetAppClusterSnapshotPolicy.StorageNetAppBaseSnapshotPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Array")

		// remove fields from embedded structs
		reflectStorageNetAppBaseSnapshotPolicy := reflect.ValueOf(o.StorageNetAppBaseSnapshotPolicy)
		for i := 0; i < reflectStorageNetAppBaseSnapshotPolicy.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseSnapshotPolicy.Type().Field(i)

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

type NullableStorageNetAppClusterSnapshotPolicy struct {
	value *StorageNetAppClusterSnapshotPolicy
	isSet bool
}

func (v NullableStorageNetAppClusterSnapshotPolicy) Get() *StorageNetAppClusterSnapshotPolicy {
	return v.value
}

func (v *NullableStorageNetAppClusterSnapshotPolicy) Set(val *StorageNetAppClusterSnapshotPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppClusterSnapshotPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppClusterSnapshotPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppClusterSnapshotPolicy(val *StorageNetAppClusterSnapshotPolicy) *NullableStorageNetAppClusterSnapshotPolicy {
	return &NullableStorageNetAppClusterSnapshotPolicy{value: val, isSet: true}
}

func (v NullableStorageNetAppClusterSnapshotPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppClusterSnapshotPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
