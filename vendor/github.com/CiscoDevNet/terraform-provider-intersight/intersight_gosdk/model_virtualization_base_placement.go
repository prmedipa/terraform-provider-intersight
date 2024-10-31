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

// checks if the VirtualizationBasePlacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationBasePlacement{}

// VirtualizationBasePlacement BasePlacement is an abstract base type for all concrete placement types. The concept of Placement is uniformly used in all hypervisors. With VMware it is the Datacenter.
type VirtualizationBasePlacement struct {
	VirtualizationBaseSourceDevice
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Name of the virtual machine placement. It is datacenter name in case of VMware virtual machine.
	Name *string `json:"Name,omitempty"`
	// The uuid of this placement. The uuid is internally generated and not user specified.
	Uuid                 *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationBasePlacement VirtualizationBasePlacement

// NewVirtualizationBasePlacement instantiates a new VirtualizationBasePlacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationBasePlacement(classId string, objectType string) *VirtualizationBasePlacement {
	this := VirtualizationBasePlacement{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationBasePlacementWithDefaults instantiates a new VirtualizationBasePlacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationBasePlacementWithDefaults() *VirtualizationBasePlacement {
	this := VirtualizationBasePlacement{}
	var classId string = "virtualization.VmwareDatacenter"
	this.ClassId = classId
	var objectType string = "virtualization.VmwareDatacenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationBasePlacement) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacement) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationBasePlacement) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwareDatacenter" of the ClassId field.
func (o *VirtualizationBasePlacement) GetDefaultClassId() interface{} {
	return "virtualization.VmwareDatacenter"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationBasePlacement) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacement) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationBasePlacement) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwareDatacenter" of the ObjectType field.
func (o *VirtualizationBasePlacement) GetDefaultObjectType() interface{} {
	return "virtualization.VmwareDatacenter"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualizationBasePlacement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualizationBasePlacement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualizationBasePlacement) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualizationBasePlacement) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationBasePlacement) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualizationBasePlacement) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualizationBasePlacement) SetUuid(v string) {
	o.Uuid = &v
}

func (o VirtualizationBasePlacement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationBasePlacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseSourceDevice, errVirtualizationBaseSourceDevice := json.Marshal(o.VirtualizationBaseSourceDevice)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
	}
	errVirtualizationBaseSourceDevice = json.Unmarshal([]byte(serializedVirtualizationBaseSourceDevice), &toSerialize)
	if errVirtualizationBaseSourceDevice != nil {
		return map[string]interface{}{}, errVirtualizationBaseSourceDevice
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationBasePlacement) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationBasePlacementWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Name of the virtual machine placement. It is datacenter name in case of VMware virtual machine.
		Name *string `json:"Name,omitempty"`
		// The uuid of this placement. The uuid is internally generated and not user specified.
		Uuid *string `json:"Uuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	}

	varVirtualizationBasePlacementWithoutEmbeddedStruct := VirtualizationBasePlacementWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationBasePlacementWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationBasePlacement := _VirtualizationBasePlacement{}
		varVirtualizationBasePlacement.ClassId = varVirtualizationBasePlacementWithoutEmbeddedStruct.ClassId
		varVirtualizationBasePlacement.ObjectType = varVirtualizationBasePlacementWithoutEmbeddedStruct.ObjectType
		varVirtualizationBasePlacement.Name = varVirtualizationBasePlacementWithoutEmbeddedStruct.Name
		varVirtualizationBasePlacement.Uuid = varVirtualizationBasePlacementWithoutEmbeddedStruct.Uuid
		*o = VirtualizationBasePlacement(varVirtualizationBasePlacement)
	} else {
		return err
	}

	varVirtualizationBasePlacement := _VirtualizationBasePlacement{}

	err = json.Unmarshal(data, &varVirtualizationBasePlacement)
	if err == nil {
		o.VirtualizationBaseSourceDevice = varVirtualizationBasePlacement.VirtualizationBaseSourceDevice
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Uuid")

		// remove fields from embedded structs
		reflectVirtualizationBaseSourceDevice := reflect.ValueOf(o.VirtualizationBaseSourceDevice)
		for i := 0; i < reflectVirtualizationBaseSourceDevice.Type().NumField(); i++ {
			t := reflectVirtualizationBaseSourceDevice.Type().Field(i)

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

type NullableVirtualizationBasePlacement struct {
	value *VirtualizationBasePlacement
	isSet bool
}

func (v NullableVirtualizationBasePlacement) Get() *VirtualizationBasePlacement {
	return v.value
}

func (v *NullableVirtualizationBasePlacement) Set(val *VirtualizationBasePlacement) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationBasePlacement) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationBasePlacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationBasePlacement(val *VirtualizationBasePlacement) *NullableVirtualizationBasePlacement {
	return &NullableVirtualizationBasePlacement{value: val, isSet: true}
}

func (v NullableVirtualizationBasePlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationBasePlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
