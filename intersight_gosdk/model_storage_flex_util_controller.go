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

// checks if the StorageFlexUtilController type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageFlexUtilController{}

// StorageFlexUtilController Storage Flex Util Adapter.
type StorageFlexUtilController struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Flex Util Controller.
	ControllerName *string `json:"ControllerName,omitempty"`
	// The current status of the controller.
	ControllerStatus *string `json:"ControllerStatus,omitempty"`
	// Identifier for the Storage Flex Util Controller.
	FfControllerId *string `json:"FfControllerId,omitempty"`
	// The internal state of the controller.
	InternalState *string                          `json:"InternalState,omitempty"`
	ComputeBoard  NullableComputeBoardRelationship `json:"ComputeBoard,omitempty"`
	// An array of relationships to storageFlexUtilPhysicalDrive resources.
	FlexUtilPhysicalDrives []StorageFlexUtilPhysicalDriveRelationship `json:"FlexUtilPhysicalDrives,omitempty"`
	// An array of relationships to storageFlexUtilVirtualDrive resources.
	FlexUtilVirtualDrives []StorageFlexUtilVirtualDriveRelationship   `json:"FlexUtilVirtualDrives,omitempty"`
	InventoryDeviceInfo   NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice      NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _StorageFlexUtilController StorageFlexUtilController

// NewStorageFlexUtilController instantiates a new StorageFlexUtilController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexUtilController(classId string, objectType string) *StorageFlexUtilController {
	this := StorageFlexUtilController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexUtilControllerWithDefaults instantiates a new StorageFlexUtilController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexUtilControllerWithDefaults() *StorageFlexUtilController {
	this := StorageFlexUtilController{}
	var classId string = "storage.FlexUtilController"
	this.ClassId = classId
	var objectType string = "storage.FlexUtilController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexUtilController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexUtilController) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.FlexUtilController" of the ClassId field.
func (o *StorageFlexUtilController) GetDefaultClassId() interface{} {
	return "storage.FlexUtilController"
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexUtilController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexUtilController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.FlexUtilController" of the ObjectType field.
func (o *StorageFlexUtilController) GetDefaultObjectType() interface{} {
	return "storage.FlexUtilController"
}

// GetControllerName returns the ControllerName field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetControllerName() string {
	if o == nil || IsNil(o.ControllerName) {
		var ret string
		return ret
	}
	return *o.ControllerName
}

// GetControllerNameOk returns a tuple with the ControllerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetControllerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerName) {
		return nil, false
	}
	return o.ControllerName, true
}

// HasControllerName returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasControllerName() bool {
	if o != nil && !IsNil(o.ControllerName) {
		return true
	}

	return false
}

// SetControllerName gets a reference to the given string and assigns it to the ControllerName field.
func (o *StorageFlexUtilController) SetControllerName(v string) {
	o.ControllerName = &v
}

// GetControllerStatus returns the ControllerStatus field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetControllerStatus() string {
	if o == nil || IsNil(o.ControllerStatus) {
		var ret string
		return ret
	}
	return *o.ControllerStatus
}

// GetControllerStatusOk returns a tuple with the ControllerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetControllerStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerStatus) {
		return nil, false
	}
	return o.ControllerStatus, true
}

// HasControllerStatus returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasControllerStatus() bool {
	if o != nil && !IsNil(o.ControllerStatus) {
		return true
	}

	return false
}

// SetControllerStatus gets a reference to the given string and assigns it to the ControllerStatus field.
func (o *StorageFlexUtilController) SetControllerStatus(v string) {
	o.ControllerStatus = &v
}

// GetFfControllerId returns the FfControllerId field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetFfControllerId() string {
	if o == nil || IsNil(o.FfControllerId) {
		var ret string
		return ret
	}
	return *o.FfControllerId
}

// GetFfControllerIdOk returns a tuple with the FfControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetFfControllerIdOk() (*string, bool) {
	if o == nil || IsNil(o.FfControllerId) {
		return nil, false
	}
	return o.FfControllerId, true
}

// HasFfControllerId returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFfControllerId() bool {
	if o != nil && !IsNil(o.FfControllerId) {
		return true
	}

	return false
}

// SetFfControllerId gets a reference to the given string and assigns it to the FfControllerId field.
func (o *StorageFlexUtilController) SetFfControllerId(v string) {
	o.FfControllerId = &v
}

// GetInternalState returns the InternalState field value if set, zero value otherwise.
func (o *StorageFlexUtilController) GetInternalState() string {
	if o == nil || IsNil(o.InternalState) {
		var ret string
		return ret
	}
	return *o.InternalState
}

// GetInternalStateOk returns a tuple with the InternalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexUtilController) GetInternalStateOk() (*string, bool) {
	if o == nil || IsNil(o.InternalState) {
		return nil, false
	}
	return o.InternalState, true
}

// HasInternalState returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasInternalState() bool {
	if o != nil && !IsNil(o.InternalState) {
		return true
	}

	return false
}

// SetInternalState gets a reference to the given string and assigns it to the InternalState field.
func (o *StorageFlexUtilController) SetInternalState(v string) {
	o.InternalState = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || IsNil(o.ComputeBoard.Get()) {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard.Get()
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputeBoard.Get(), o.ComputeBoard.IsSet()
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard.IsSet() {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given NullableComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageFlexUtilController) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard.Set(&v)
}

// SetComputeBoardNil sets the value for ComputeBoard to be an explicit nil
func (o *StorageFlexUtilController) SetComputeBoardNil() {
	o.ComputeBoard.Set(nil)
}

// UnsetComputeBoard ensures that no value is present for ComputeBoard, not even an explicit nil
func (o *StorageFlexUtilController) UnsetComputeBoard() {
	o.ComputeBoard.Unset()
}

// GetFlexUtilPhysicalDrives returns the FlexUtilPhysicalDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrives() []StorageFlexUtilPhysicalDriveRelationship {
	if o == nil {
		var ret []StorageFlexUtilPhysicalDriveRelationship
		return ret
	}
	return o.FlexUtilPhysicalDrives
}

// GetFlexUtilPhysicalDrivesOk returns a tuple with the FlexUtilPhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetFlexUtilPhysicalDrivesOk() ([]StorageFlexUtilPhysicalDriveRelationship, bool) {
	if o == nil || IsNil(o.FlexUtilPhysicalDrives) {
		return nil, false
	}
	return o.FlexUtilPhysicalDrives, true
}

// HasFlexUtilPhysicalDrives returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFlexUtilPhysicalDrives() bool {
	if o != nil && !IsNil(o.FlexUtilPhysicalDrives) {
		return true
	}

	return false
}

// SetFlexUtilPhysicalDrives gets a reference to the given []StorageFlexUtilPhysicalDriveRelationship and assigns it to the FlexUtilPhysicalDrives field.
func (o *StorageFlexUtilController) SetFlexUtilPhysicalDrives(v []StorageFlexUtilPhysicalDriveRelationship) {
	o.FlexUtilPhysicalDrives = v
}

// GetFlexUtilVirtualDrives returns the FlexUtilVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetFlexUtilVirtualDrives() []StorageFlexUtilVirtualDriveRelationship {
	if o == nil {
		var ret []StorageFlexUtilVirtualDriveRelationship
		return ret
	}
	return o.FlexUtilVirtualDrives
}

// GetFlexUtilVirtualDrivesOk returns a tuple with the FlexUtilVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetFlexUtilVirtualDrivesOk() ([]StorageFlexUtilVirtualDriveRelationship, bool) {
	if o == nil || IsNil(o.FlexUtilVirtualDrives) {
		return nil, false
	}
	return o.FlexUtilVirtualDrives, true
}

// HasFlexUtilVirtualDrives returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasFlexUtilVirtualDrives() bool {
	if o != nil && !IsNil(o.FlexUtilVirtualDrives) {
		return true
	}

	return false
}

// SetFlexUtilVirtualDrives gets a reference to the given []StorageFlexUtilVirtualDriveRelationship and assigns it to the FlexUtilVirtualDrives field.
func (o *StorageFlexUtilController) SetFlexUtilVirtualDrives(v []StorageFlexUtilVirtualDriveRelationship) {
	o.FlexUtilVirtualDrives = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || IsNil(o.InventoryDeviceInfo.Get()) {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo.Get()
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo.Get(), o.InventoryDeviceInfo.IsSet()
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo.IsSet() {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given NullableInventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexUtilController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo.Set(&v)
}

// SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil
func (o *StorageFlexUtilController) SetInventoryDeviceInfoNil() {
	o.InventoryDeviceInfo.Set(nil)
}

// UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
func (o *StorageFlexUtilController) UnsetInventoryDeviceInfo() {
	o.InventoryDeviceInfo.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexUtilController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexUtilController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexUtilController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexUtilController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *StorageFlexUtilController) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *StorageFlexUtilController) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o StorageFlexUtilController) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageFlexUtilController) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ControllerName) {
		toSerialize["ControllerName"] = o.ControllerName
	}
	if !IsNil(o.ControllerStatus) {
		toSerialize["ControllerStatus"] = o.ControllerStatus
	}
	if !IsNil(o.FfControllerId) {
		toSerialize["FfControllerId"] = o.FfControllerId
	}
	if !IsNil(o.InternalState) {
		toSerialize["InternalState"] = o.InternalState
	}
	if o.ComputeBoard.IsSet() {
		toSerialize["ComputeBoard"] = o.ComputeBoard.Get()
	}
	if o.FlexUtilPhysicalDrives != nil {
		toSerialize["FlexUtilPhysicalDrives"] = o.FlexUtilPhysicalDrives
	}
	if o.FlexUtilVirtualDrives != nil {
		toSerialize["FlexUtilVirtualDrives"] = o.FlexUtilVirtualDrives
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

func (o *StorageFlexUtilController) UnmarshalJSON(data []byte) (err error) {
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
	type StorageFlexUtilControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Flex Util Controller.
		ControllerName *string `json:"ControllerName,omitempty"`
		// The current status of the controller.
		ControllerStatus *string `json:"ControllerStatus,omitempty"`
		// Identifier for the Storage Flex Util Controller.
		FfControllerId *string `json:"FfControllerId,omitempty"`
		// The internal state of the controller.
		InternalState *string                          `json:"InternalState,omitempty"`
		ComputeBoard  NullableComputeBoardRelationship `json:"ComputeBoard,omitempty"`
		// An array of relationships to storageFlexUtilPhysicalDrive resources.
		FlexUtilPhysicalDrives []StorageFlexUtilPhysicalDriveRelationship `json:"FlexUtilPhysicalDrives,omitempty"`
		// An array of relationships to storageFlexUtilVirtualDrive resources.
		FlexUtilVirtualDrives []StorageFlexUtilVirtualDriveRelationship   `json:"FlexUtilVirtualDrives,omitempty"`
		InventoryDeviceInfo   NullableInventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice      NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varStorageFlexUtilControllerWithoutEmbeddedStruct := StorageFlexUtilControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageFlexUtilControllerWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexUtilController := _StorageFlexUtilController{}
		varStorageFlexUtilController.ClassId = varStorageFlexUtilControllerWithoutEmbeddedStruct.ClassId
		varStorageFlexUtilController.ObjectType = varStorageFlexUtilControllerWithoutEmbeddedStruct.ObjectType
		varStorageFlexUtilController.ControllerName = varStorageFlexUtilControllerWithoutEmbeddedStruct.ControllerName
		varStorageFlexUtilController.ControllerStatus = varStorageFlexUtilControllerWithoutEmbeddedStruct.ControllerStatus
		varStorageFlexUtilController.FfControllerId = varStorageFlexUtilControllerWithoutEmbeddedStruct.FfControllerId
		varStorageFlexUtilController.InternalState = varStorageFlexUtilControllerWithoutEmbeddedStruct.InternalState
		varStorageFlexUtilController.ComputeBoard = varStorageFlexUtilControllerWithoutEmbeddedStruct.ComputeBoard
		varStorageFlexUtilController.FlexUtilPhysicalDrives = varStorageFlexUtilControllerWithoutEmbeddedStruct.FlexUtilPhysicalDrives
		varStorageFlexUtilController.FlexUtilVirtualDrives = varStorageFlexUtilControllerWithoutEmbeddedStruct.FlexUtilVirtualDrives
		varStorageFlexUtilController.InventoryDeviceInfo = varStorageFlexUtilControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexUtilController.RegisteredDevice = varStorageFlexUtilControllerWithoutEmbeddedStruct.RegisteredDevice
		*o = StorageFlexUtilController(varStorageFlexUtilController)
	} else {
		return err
	}

	varStorageFlexUtilController := _StorageFlexUtilController{}

	err = json.Unmarshal(data, &varStorageFlexUtilController)
	if err == nil {
		o.InventoryBase = varStorageFlexUtilController.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerName")
		delete(additionalProperties, "ControllerStatus")
		delete(additionalProperties, "FfControllerId")
		delete(additionalProperties, "InternalState")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "FlexUtilPhysicalDrives")
		delete(additionalProperties, "FlexUtilVirtualDrives")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageFlexUtilController struct {
	value *StorageFlexUtilController
	isSet bool
}

func (v NullableStorageFlexUtilController) Get() *StorageFlexUtilController {
	return v.value
}

func (v *NullableStorageFlexUtilController) Set(val *StorageFlexUtilController) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexUtilController) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexUtilController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexUtilController(val *StorageFlexUtilController) *NullableStorageFlexUtilController {
	return &NullableStorageFlexUtilController{value: val, isSet: true}
}

func (v NullableStorageFlexUtilController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexUtilController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
