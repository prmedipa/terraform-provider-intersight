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

// checks if the BootConfiguredDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootConfiguredDevice{}

// BootConfiguredDevice Abstract class for all configured boot devices.
type BootConfiguredDevice struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the boot device configured in the boot policy.
	Name *string `json:"Name,omitempty"`
	// The order of the boot device configured in the boot policy.
	Order *int64 `json:"Order,omitempty"`
	// The state of the boot device configured in the boot policy.
	State *string `json:"State,omitempty"`
	// The type of the boot device configured in the boot policy.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootConfiguredDevice BootConfiguredDevice

// NewBootConfiguredDevice instantiates a new BootConfiguredDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootConfiguredDevice(classId string, objectType string) *BootConfiguredDevice {
	this := BootConfiguredDevice{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewBootConfiguredDeviceWithDefaults instantiates a new BootConfiguredDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootConfiguredDeviceWithDefaults() *BootConfiguredDevice {
	this := BootConfiguredDevice{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootConfiguredDevice) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootConfiguredDevice) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootConfiguredDevice) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootConfiguredDevice) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BootConfiguredDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BootConfiguredDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BootConfiguredDevice) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *BootConfiguredDevice) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *BootConfiguredDevice) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *BootConfiguredDevice) SetOrder(v int64) {
	o.Order = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BootConfiguredDevice) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BootConfiguredDevice) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BootConfiguredDevice) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BootConfiguredDevice) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootConfiguredDevice) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BootConfiguredDevice) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BootConfiguredDevice) SetType(v string) {
	o.Type = &v
}

func (o BootConfiguredDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootConfiguredDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["Order"] = o.Order
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BootConfiguredDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type BootConfiguredDeviceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The name of the boot device configured in the boot policy.
		Name *string `json:"Name,omitempty"`
		// The order of the boot device configured in the boot policy.
		Order *int64 `json:"Order,omitempty"`
		// The state of the boot device configured in the boot policy.
		State *string `json:"State,omitempty"`
		// The type of the boot device configured in the boot policy.
		Type *string `json:"Type,omitempty"`
	}

	varBootConfiguredDeviceWithoutEmbeddedStruct := BootConfiguredDeviceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootConfiguredDeviceWithoutEmbeddedStruct)
	if err == nil {
		varBootConfiguredDevice := _BootConfiguredDevice{}
		varBootConfiguredDevice.ClassId = varBootConfiguredDeviceWithoutEmbeddedStruct.ClassId
		varBootConfiguredDevice.ObjectType = varBootConfiguredDeviceWithoutEmbeddedStruct.ObjectType
		varBootConfiguredDevice.Name = varBootConfiguredDeviceWithoutEmbeddedStruct.Name
		varBootConfiguredDevice.Order = varBootConfiguredDeviceWithoutEmbeddedStruct.Order
		varBootConfiguredDevice.State = varBootConfiguredDeviceWithoutEmbeddedStruct.State
		varBootConfiguredDevice.Type = varBootConfiguredDeviceWithoutEmbeddedStruct.Type
		*o = BootConfiguredDevice(varBootConfiguredDevice)
	} else {
		return err
	}

	varBootConfiguredDevice := _BootConfiguredDevice{}

	err = json.Unmarshal(data, &varBootConfiguredDevice)
	if err == nil {
		o.EquipmentBase = varBootConfiguredDevice.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Order")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Type")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableBootConfiguredDevice struct {
	value *BootConfiguredDevice
	isSet bool
}

func (v NullableBootConfiguredDevice) Get() *BootConfiguredDevice {
	return v.value
}

func (v *NullableBootConfiguredDevice) Set(val *BootConfiguredDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableBootConfiguredDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableBootConfiguredDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootConfiguredDevice(val *BootConfiguredDevice) *NullableBootConfiguredDevice {
	return &NullableBootConfiguredDevice{value: val, isSet: true}
}

func (v NullableBootConfiguredDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootConfiguredDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
