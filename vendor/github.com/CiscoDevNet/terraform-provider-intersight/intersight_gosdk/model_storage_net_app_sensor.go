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
	"reflect"
	"strings"
)

// StorageNetAppSensor Information for a particular sensor on a NetApp storage array controller.
type StorageNetAppSensor struct {
	EquipmentBaseSensor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                         `json:"ObjectType"`
	ArrayController      *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppSensor StorageNetAppSensor

// NewStorageNetAppSensor instantiates a new StorageNetAppSensor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppSensor(classId string, objectType string) *StorageNetAppSensor {
	this := StorageNetAppSensor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppSensorWithDefaults instantiates a new StorageNetAppSensor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppSensorWithDefaults() *StorageNetAppSensor {
	this := StorageNetAppSensor{}
	var classId string = "storage.NetAppSensor"
	this.ClassId = classId
	var objectType string = "storage.NetAppSensor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppSensor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSensor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppSensor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppSensor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppSensor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppSensor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppSensor) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppSensor) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppSensor) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppSensor) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

func (o StorageNetAppSensor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBaseSensor, errEquipmentBaseSensor := json.Marshal(o.EquipmentBaseSensor)
	if errEquipmentBaseSensor != nil {
		return []byte{}, errEquipmentBaseSensor
	}
	errEquipmentBaseSensor = json.Unmarshal([]byte(serializedEquipmentBaseSensor), &toSerialize)
	if errEquipmentBaseSensor != nil {
		return []byte{}, errEquipmentBaseSensor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppSensor) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppSensorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                         `json:"ObjectType"`
		ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	}

	varStorageNetAppSensorWithoutEmbeddedStruct := StorageNetAppSensorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppSensorWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppSensor := _StorageNetAppSensor{}
		varStorageNetAppSensor.ClassId = varStorageNetAppSensorWithoutEmbeddedStruct.ClassId
		varStorageNetAppSensor.ObjectType = varStorageNetAppSensorWithoutEmbeddedStruct.ObjectType
		varStorageNetAppSensor.ArrayController = varStorageNetAppSensorWithoutEmbeddedStruct.ArrayController
		*o = StorageNetAppSensor(varStorageNetAppSensor)
	} else {
		return err
	}

	varStorageNetAppSensor := _StorageNetAppSensor{}

	err = json.Unmarshal(bytes, &varStorageNetAppSensor)
	if err == nil {
		o.EquipmentBaseSensor = varStorageNetAppSensor.EquipmentBaseSensor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayController")

		// remove fields from embedded structs
		reflectEquipmentBaseSensor := reflect.ValueOf(o.EquipmentBaseSensor)
		for i := 0; i < reflectEquipmentBaseSensor.Type().NumField(); i++ {
			t := reflectEquipmentBaseSensor.Type().Field(i)

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

type NullableStorageNetAppSensor struct {
	value *StorageNetAppSensor
	isSet bool
}

func (v NullableStorageNetAppSensor) Get() *StorageNetAppSensor {
	return v.value
}

func (v *NullableStorageNetAppSensor) Set(val *StorageNetAppSensor) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppSensor) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppSensor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppSensor(val *StorageNetAppSensor) *NullableStorageNetAppSensor {
	return &NullableStorageNetAppSensor{value: val, isSet: true}
}

func (v NullableStorageNetAppSensor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppSensor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
