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

// EquipmentChassisIdPool ChassisIdPool object contains pool of chassisId's that can be allocated for newly discovered chassis.
type EquipmentChassisIdPool struct {
	SwIdPoolBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	PreferredIds         []int64                              `json:"PreferredIds,omitempty"`
	DeviceRegistration   *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentChassisIdPool EquipmentChassisIdPool

// NewEquipmentChassisIdPool instantiates a new EquipmentChassisIdPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentChassisIdPool(classId string, objectType string) *EquipmentChassisIdPool {
	this := EquipmentChassisIdPool{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentChassisIdPoolWithDefaults instantiates a new EquipmentChassisIdPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentChassisIdPoolWithDefaults() *EquipmentChassisIdPool {
	this := EquipmentChassisIdPool{}
	var classId string = "equipment.ChassisIdPool"
	this.ClassId = classId
	var objectType string = "equipment.ChassisIdPool"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentChassisIdPool) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentChassisIdPool) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentChassisIdPool) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentChassisIdPool) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentChassisIdPool) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentChassisIdPool) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPreferredIds returns the PreferredIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisIdPool) GetPreferredIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PreferredIds
}

// GetPreferredIdsOk returns a tuple with the PreferredIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisIdPool) GetPreferredIdsOk() ([]int64, bool) {
	if o == nil || o.PreferredIds == nil {
		return nil, false
	}
	return o.PreferredIds, true
}

// HasPreferredIds returns a boolean if a field has been set.
func (o *EquipmentChassisIdPool) HasPreferredIds() bool {
	if o != nil && o.PreferredIds != nil {
		return true
	}

	return false
}

// SetPreferredIds gets a reference to the given []int64 and assigns it to the PreferredIds field.
func (o *EquipmentChassisIdPool) SetPreferredIds(v []int64) {
	o.PreferredIds = v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentChassisIdPool) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisIdPool) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentChassisIdPool) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentChassisIdPool) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

func (o EquipmentChassisIdPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSwIdPoolBase, errSwIdPoolBase := json.Marshal(o.SwIdPoolBase)
	if errSwIdPoolBase != nil {
		return []byte{}, errSwIdPoolBase
	}
	errSwIdPoolBase = json.Unmarshal([]byte(serializedSwIdPoolBase), &toSerialize)
	if errSwIdPoolBase != nil {
		return []byte{}, errSwIdPoolBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PreferredIds != nil {
		toSerialize["PreferredIds"] = o.PreferredIds
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentChassisIdPool) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentChassisIdPoolWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                               `json:"ObjectType"`
		PreferredIds       []int64                              `json:"PreferredIds,omitempty"`
		DeviceRegistration *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	}

	varEquipmentChassisIdPoolWithoutEmbeddedStruct := EquipmentChassisIdPoolWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentChassisIdPoolWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentChassisIdPool := _EquipmentChassisIdPool{}
		varEquipmentChassisIdPool.ClassId = varEquipmentChassisIdPoolWithoutEmbeddedStruct.ClassId
		varEquipmentChassisIdPool.ObjectType = varEquipmentChassisIdPoolWithoutEmbeddedStruct.ObjectType
		varEquipmentChassisIdPool.PreferredIds = varEquipmentChassisIdPoolWithoutEmbeddedStruct.PreferredIds
		varEquipmentChassisIdPool.DeviceRegistration = varEquipmentChassisIdPoolWithoutEmbeddedStruct.DeviceRegistration
		*o = EquipmentChassisIdPool(varEquipmentChassisIdPool)
	} else {
		return err
	}

	varEquipmentChassisIdPool := _EquipmentChassisIdPool{}

	err = json.Unmarshal(bytes, &varEquipmentChassisIdPool)
	if err == nil {
		o.SwIdPoolBase = varEquipmentChassisIdPool.SwIdPoolBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PreferredIds")
		delete(additionalProperties, "DeviceRegistration")

		// remove fields from embedded structs
		reflectSwIdPoolBase := reflect.ValueOf(o.SwIdPoolBase)
		for i := 0; i < reflectSwIdPoolBase.Type().NumField(); i++ {
			t := reflectSwIdPoolBase.Type().Field(i)

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

type NullableEquipmentChassisIdPool struct {
	value *EquipmentChassisIdPool
	isSet bool
}

func (v NullableEquipmentChassisIdPool) Get() *EquipmentChassisIdPool {
	return v.value
}

func (v *NullableEquipmentChassisIdPool) Set(val *EquipmentChassisIdPool) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentChassisIdPool) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentChassisIdPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentChassisIdPool(val *EquipmentChassisIdPool) *NullableEquipmentChassisIdPool {
	return &NullableEquipmentChassisIdPool{value: val, isSet: true}
}

func (v NullableEquipmentChassisIdPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentChassisIdPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
