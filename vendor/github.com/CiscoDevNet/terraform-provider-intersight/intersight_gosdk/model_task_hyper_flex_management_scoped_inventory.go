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

// TaskHyperFlexManagementScopedInventory API to trigger on-demand HyperFlex inventory to update modified objects in Intersight report.
type TaskHyperFlexManagementScopedInventory struct {
	ConnectorScopedInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                               `json:"ObjectType"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TaskHyperFlexManagementScopedInventory TaskHyperFlexManagementScopedInventory

// NewTaskHyperFlexManagementScopedInventory instantiates a new TaskHyperFlexManagementScopedInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskHyperFlexManagementScopedInventory(classId string, objectType string) *TaskHyperFlexManagementScopedInventory {
	this := TaskHyperFlexManagementScopedInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewTaskHyperFlexManagementScopedInventoryWithDefaults instantiates a new TaskHyperFlexManagementScopedInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskHyperFlexManagementScopedInventoryWithDefaults() *TaskHyperFlexManagementScopedInventory {
	this := TaskHyperFlexManagementScopedInventory{}
	var classId string = "task.HyperFlexManagementScopedInventory"
	this.ClassId = classId
	var objectType string = "task.HyperFlexManagementScopedInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *TaskHyperFlexManagementScopedInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TaskHyperFlexManagementScopedInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TaskHyperFlexManagementScopedInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TaskHyperFlexManagementScopedInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TaskHyperFlexManagementScopedInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TaskHyperFlexManagementScopedInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *TaskHyperFlexManagementScopedInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskHyperFlexManagementScopedInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *TaskHyperFlexManagementScopedInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *TaskHyperFlexManagementScopedInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o TaskHyperFlexManagementScopedInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorScopedInventory, errConnectorScopedInventory := json.Marshal(o.ConnectorScopedInventory)
	if errConnectorScopedInventory != nil {
		return []byte{}, errConnectorScopedInventory
	}
	errConnectorScopedInventory = json.Unmarshal([]byte(serializedConnectorScopedInventory), &toSerialize)
	if errConnectorScopedInventory != nil {
		return []byte{}, errConnectorScopedInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TaskHyperFlexManagementScopedInventory) UnmarshalJSON(bytes []byte) (err error) {
	type TaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType       string                               `json:"ObjectType"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varTaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct := TaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct)
	if err == nil {
		varTaskHyperFlexManagementScopedInventory := _TaskHyperFlexManagementScopedInventory{}
		varTaskHyperFlexManagementScopedInventory.ClassId = varTaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct.ClassId
		varTaskHyperFlexManagementScopedInventory.ObjectType = varTaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct.ObjectType
		varTaskHyperFlexManagementScopedInventory.RegisteredDevice = varTaskHyperFlexManagementScopedInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = TaskHyperFlexManagementScopedInventory(varTaskHyperFlexManagementScopedInventory)
	} else {
		return err
	}

	varTaskHyperFlexManagementScopedInventory := _TaskHyperFlexManagementScopedInventory{}

	err = json.Unmarshal(bytes, &varTaskHyperFlexManagementScopedInventory)
	if err == nil {
		o.ConnectorScopedInventory = varTaskHyperFlexManagementScopedInventory.ConnectorScopedInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectConnectorScopedInventory := reflect.ValueOf(o.ConnectorScopedInventory)
		for i := 0; i < reflectConnectorScopedInventory.Type().NumField(); i++ {
			t := reflectConnectorScopedInventory.Type().Field(i)

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

type NullableTaskHyperFlexManagementScopedInventory struct {
	value *TaskHyperFlexManagementScopedInventory
	isSet bool
}

func (v NullableTaskHyperFlexManagementScopedInventory) Get() *TaskHyperFlexManagementScopedInventory {
	return v.value
}

func (v *NullableTaskHyperFlexManagementScopedInventory) Set(val *TaskHyperFlexManagementScopedInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskHyperFlexManagementScopedInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskHyperFlexManagementScopedInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskHyperFlexManagementScopedInventory(val *TaskHyperFlexManagementScopedInventory) *NullableTaskHyperFlexManagementScopedInventory {
	return &NullableTaskHyperFlexManagementScopedInventory{value: val, isSet: true}
}

func (v NullableTaskHyperFlexManagementScopedInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskHyperFlexManagementScopedInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
