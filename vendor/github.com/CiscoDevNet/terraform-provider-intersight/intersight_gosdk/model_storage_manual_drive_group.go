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

// StorageManualDriveGroup This models the manual drive selection configuration.
type StorageManualDriveGroup struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A collection of drives to be used as hot spares for this Drive Group.
	DedicatedHotSpares   *string             `json:"DedicatedHotSpares,omitempty"`
	SpanGroups           []StorageSpanDrives `json:"SpanGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageManualDriveGroup StorageManualDriveGroup

// NewStorageManualDriveGroup instantiates a new StorageManualDriveGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageManualDriveGroup(classId string, objectType string) *StorageManualDriveGroup {
	this := StorageManualDriveGroup{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageManualDriveGroupWithDefaults instantiates a new StorageManualDriveGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageManualDriveGroupWithDefaults() *StorageManualDriveGroup {
	this := StorageManualDriveGroup{}
	var classId string = "storage.ManualDriveGroup"
	this.ClassId = classId
	var objectType string = "storage.ManualDriveGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageManualDriveGroup) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageManualDriveGroup) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageManualDriveGroup) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageManualDriveGroup) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageManualDriveGroup) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageManualDriveGroup) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDedicatedHotSpares returns the DedicatedHotSpares field value if set, zero value otherwise.
func (o *StorageManualDriveGroup) GetDedicatedHotSpares() string {
	if o == nil || o.DedicatedHotSpares == nil {
		var ret string
		return ret
	}
	return *o.DedicatedHotSpares
}

// GetDedicatedHotSparesOk returns a tuple with the DedicatedHotSpares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageManualDriveGroup) GetDedicatedHotSparesOk() (*string, bool) {
	if o == nil || o.DedicatedHotSpares == nil {
		return nil, false
	}
	return o.DedicatedHotSpares, true
}

// HasDedicatedHotSpares returns a boolean if a field has been set.
func (o *StorageManualDriveGroup) HasDedicatedHotSpares() bool {
	if o != nil && o.DedicatedHotSpares != nil {
		return true
	}

	return false
}

// SetDedicatedHotSpares gets a reference to the given string and assigns it to the DedicatedHotSpares field.
func (o *StorageManualDriveGroup) SetDedicatedHotSpares(v string) {
	o.DedicatedHotSpares = &v
}

// GetSpanGroups returns the SpanGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageManualDriveGroup) GetSpanGroups() []StorageSpanDrives {
	if o == nil {
		var ret []StorageSpanDrives
		return ret
	}
	return o.SpanGroups
}

// GetSpanGroupsOk returns a tuple with the SpanGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageManualDriveGroup) GetSpanGroupsOk() ([]StorageSpanDrives, bool) {
	if o == nil || o.SpanGroups == nil {
		return nil, false
	}
	return o.SpanGroups, true
}

// HasSpanGroups returns a boolean if a field has been set.
func (o *StorageManualDriveGroup) HasSpanGroups() bool {
	if o != nil && o.SpanGroups != nil {
		return true
	}

	return false
}

// SetSpanGroups gets a reference to the given []StorageSpanDrives and assigns it to the SpanGroups field.
func (o *StorageManualDriveGroup) SetSpanGroups(v []StorageSpanDrives) {
	o.SpanGroups = v
}

func (o StorageManualDriveGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DedicatedHotSpares != nil {
		toSerialize["DedicatedHotSpares"] = o.DedicatedHotSpares
	}
	if o.SpanGroups != nil {
		toSerialize["SpanGroups"] = o.SpanGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageManualDriveGroup) UnmarshalJSON(bytes []byte) (err error) {
	type StorageManualDriveGroupWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A collection of drives to be used as hot spares for this Drive Group.
		DedicatedHotSpares *string             `json:"DedicatedHotSpares,omitempty"`
		SpanGroups         []StorageSpanDrives `json:"SpanGroups,omitempty"`
	}

	varStorageManualDriveGroupWithoutEmbeddedStruct := StorageManualDriveGroupWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageManualDriveGroupWithoutEmbeddedStruct)
	if err == nil {
		varStorageManualDriveGroup := _StorageManualDriveGroup{}
		varStorageManualDriveGroup.ClassId = varStorageManualDriveGroupWithoutEmbeddedStruct.ClassId
		varStorageManualDriveGroup.ObjectType = varStorageManualDriveGroupWithoutEmbeddedStruct.ObjectType
		varStorageManualDriveGroup.DedicatedHotSpares = varStorageManualDriveGroupWithoutEmbeddedStruct.DedicatedHotSpares
		varStorageManualDriveGroup.SpanGroups = varStorageManualDriveGroupWithoutEmbeddedStruct.SpanGroups
		*o = StorageManualDriveGroup(varStorageManualDriveGroup)
	} else {
		return err
	}

	varStorageManualDriveGroup := _StorageManualDriveGroup{}

	err = json.Unmarshal(bytes, &varStorageManualDriveGroup)
	if err == nil {
		o.MoBaseComplexType = varStorageManualDriveGroup.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DedicatedHotSpares")
		delete(additionalProperties, "SpanGroups")

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

type NullableStorageManualDriveGroup struct {
	value *StorageManualDriveGroup
	isSet bool
}

func (v NullableStorageManualDriveGroup) Get() *StorageManualDriveGroup {
	return v.value
}

func (v *NullableStorageManualDriveGroup) Set(val *StorageManualDriveGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageManualDriveGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageManualDriveGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageManualDriveGroup(val *StorageManualDriveGroup) *NullableStorageManualDriveGroup {
	return &NullableStorageManualDriveGroup{value: val, isSet: true}
}

func (v NullableStorageManualDriveGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageManualDriveGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
