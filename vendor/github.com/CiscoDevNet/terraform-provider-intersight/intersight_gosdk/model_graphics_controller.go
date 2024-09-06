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

// checks if the GraphicsController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphicsController{}

// GraphicsController Controller for a Graphics Card.
type GraphicsController struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The id of the graphics controller.
	ControllerId *int64 `json:"ControllerId,omitempty"`
	// The PCI address of the graphics controller.
	PciAddr *string `json:"PciAddr,omitempty"`
	// The PCI slot information of the graphics controller.
	PciSlot              *string                                     `json:"PciSlot,omitempty"`
	GraphicsCard         NullableGraphicsCardRelationship            `json:"GraphicsCard,omitempty"`
	InventoryDeviceInfo  NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GraphicsController GraphicsController

// NewGraphicsController instantiates a new GraphicsController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphicsController(classId string, objectType string) *GraphicsController {
	this := GraphicsController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewGraphicsControllerWithDefaults instantiates a new GraphicsController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphicsControllerWithDefaults() *GraphicsController {
	this := GraphicsController{}
	var classId string = "graphics.Controller"
	this.ClassId = classId
	var objectType string = "graphics.Controller"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *GraphicsController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *GraphicsController) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "graphics.Controller" of the ClassId field.
func (o *GraphicsController) GetDefaultClassId() interface{} {
	return "graphics.Controller"
}

// GetObjectType returns the ObjectType field value
func (o *GraphicsController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *GraphicsController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "graphics.Controller" of the ObjectType field.
func (o *GraphicsController) GetDefaultObjectType() interface{} {
	return "graphics.Controller"
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *GraphicsController) GetControllerId() int64 {
	if o == nil || IsNil(o.ControllerId) {
		var ret int64
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetControllerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ControllerId) {
		return nil, false
	}
	return o.ControllerId, true
}

// HasControllerId returns a boolean if a field has been set.
func (o *GraphicsController) HasControllerId() bool {
	if o != nil && !IsNil(o.ControllerId) {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given int64 and assigns it to the ControllerId field.
func (o *GraphicsController) SetControllerId(v int64) {
	o.ControllerId = &v
}

// GetPciAddr returns the PciAddr field value if set, zero value otherwise.
func (o *GraphicsController) GetPciAddr() string {
	if o == nil || IsNil(o.PciAddr) {
		var ret string
		return ret
	}
	return *o.PciAddr
}

// GetPciAddrOk returns a tuple with the PciAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetPciAddrOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddr) {
		return nil, false
	}
	return o.PciAddr, true
}

// HasPciAddr returns a boolean if a field has been set.
func (o *GraphicsController) HasPciAddr() bool {
	if o != nil && !IsNil(o.PciAddr) {
		return true
	}

	return false
}

// SetPciAddr gets a reference to the given string and assigns it to the PciAddr field.
func (o *GraphicsController) SetPciAddr(v string) {
	o.PciAddr = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *GraphicsController) GetPciSlot() string {
	if o == nil || IsNil(o.PciSlot) {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsController) GetPciSlotOk() (*string, bool) {
	if o == nil || IsNil(o.PciSlot) {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *GraphicsController) HasPciSlot() bool {
	if o != nil && !IsNil(o.PciSlot) {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *GraphicsController) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetGraphicsCard returns the GraphicsCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsController) GetGraphicsCard() GraphicsCardRelationship {
	if o == nil || IsNil(o.GraphicsCard.Get()) {
		var ret GraphicsCardRelationship
		return ret
	}
	return *o.GraphicsCard.Get()
}

// GetGraphicsCardOk returns a tuple with the GraphicsCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsController) GetGraphicsCardOk() (*GraphicsCardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.GraphicsCard.Get(), o.GraphicsCard.IsSet()
}

// HasGraphicsCard returns a boolean if a field has been set.
func (o *GraphicsController) HasGraphicsCard() bool {
	if o != nil && o.GraphicsCard.IsSet() {
		return true
	}

	return false
}

// SetGraphicsCard gets a reference to the given NullableGraphicsCardRelationship and assigns it to the GraphicsCard field.
func (o *GraphicsController) SetGraphicsCard(v GraphicsCardRelationship) {
	o.GraphicsCard.Set(&v)
}

// SetGraphicsCardNil sets the value for GraphicsCard to be an explicit nil
func (o *GraphicsController) SetGraphicsCardNil() {
	o.GraphicsCard.Set(nil)
}

// UnsetGraphicsCard ensures that no value is present for GraphicsCard, not even an explicit nil
func (o *GraphicsController) UnsetGraphicsCard() {
	o.GraphicsCard.Unset()
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *GraphicsController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *GraphicsController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *GraphicsController) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *GraphicsController) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *GraphicsController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *GraphicsController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *GraphicsController) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *GraphicsController) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o GraphicsController) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphicsController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return map[string]interface{}{}, errEquipmentBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ControllerId) {
		toSerialize["ControllerId"] = o.ControllerId
	}
	if !IsNil(o.PciAddr) {
		toSerialize["PciAddr"] = o.PciAddr
	}
	if !IsNil(o.PciSlot) {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.GraphicsCard.IsSet() {
		toSerialize["GraphicsCard"] = o.GraphicsCard.Get()
	}
	if o.InventoryDeviceInfo.IsSet() {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GraphicsController) UnmarshalJSON(data []byte) (err error) {
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
	type GraphicsControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The id of the graphics controller.
		ControllerId *int64 `json:"ControllerId,omitempty"`
		// The PCI address of the graphics controller.
		PciAddr *string `json:"PciAddr,omitempty"`
		// The PCI slot information of the graphics controller.
		PciSlot             *string                                     `json:"PciSlot,omitempty"`
		GraphicsCard        NullableGraphicsCardRelationship            `json:"GraphicsCard,omitempty"`
		InventoryDeviceInfo NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varGraphicsControllerWithoutEmbeddedStruct := GraphicsControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varGraphicsControllerWithoutEmbeddedStruct)
	if err == nil {
		varGraphicsController := _GraphicsController{}
		varGraphicsController.ClassId = varGraphicsControllerWithoutEmbeddedStruct.ClassId
		varGraphicsController.ObjectType = varGraphicsControllerWithoutEmbeddedStruct.ObjectType
		varGraphicsController.ControllerId = varGraphicsControllerWithoutEmbeddedStruct.ControllerId
		varGraphicsController.PciAddr = varGraphicsControllerWithoutEmbeddedStruct.PciAddr
		varGraphicsController.PciSlot = varGraphicsControllerWithoutEmbeddedStruct.PciSlot
		varGraphicsController.GraphicsCard = varGraphicsControllerWithoutEmbeddedStruct.GraphicsCard
		varGraphicsController.InventoryDeviceInfo = varGraphicsControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varGraphicsController.RegisteredDevice = varGraphicsControllerWithoutEmbeddedStruct.RegisteredDevice
		*o = GraphicsController(varGraphicsController)
	} else {
		return err
	}

	varGraphicsController := _GraphicsController{}

	err = json.Unmarshal(data, &varGraphicsController)
	if err == nil {
		o.EquipmentBase = varGraphicsController.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerId")
		delete(additionalProperties, "PciAddr")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "GraphicsCard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableGraphicsController struct {
	value *GraphicsController
	isSet bool
}

func (v NullableGraphicsController) Get() *GraphicsController {
	return v.value
}

func (v *NullableGraphicsController) Set(val *GraphicsController) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphicsController) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphicsController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphicsController(val *GraphicsController) *NullableGraphicsController {
	return &NullableGraphicsController{value: val, isSet: true}
}

func (v NullableGraphicsController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphicsController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
