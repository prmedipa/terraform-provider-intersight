/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// EquipmentSystemIoControllerAllOf Definition of the list of properties defined in 'equipment.SystemIoController', excluding properties defined in parent classes.
type EquipmentSystemIoControllerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The assigned identifier for a chassis.
	ChassisId *string `json:"ChassisId,omitempty"`
	// Connection Path identifies the data path available between IOModule and FI.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// Connection status identifies the status of data path.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field gives a brief information on systemIOController.
	Description *string `json:"Description,omitempty"`
	// This field identifies the CIMC that manages the controller.
	ManagingInstance *string `json:"ManagingInstance,omitempty"`
	// This field identifies the SIOC operational state.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the IO module.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for systemIOController.
	Pid *string `json:"Pid,omitempty"`
	// This represents system I/O Controller identifier.
	SystemIoControllerId *int64                               `json:"SystemIoControllerId,omitempty"`
	Cmc                  *ManagementControllerRelationship    `json:"Cmc,omitempty"`
	EquipmentChassis     *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	SharedIoModule       *EquipmentSharedIoModuleRelationship `json:"SharedIoModule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSystemIoControllerAllOf EquipmentSystemIoControllerAllOf

// NewEquipmentSystemIoControllerAllOf instantiates a new EquipmentSystemIoControllerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSystemIoControllerAllOf(classId string, objectType string) *EquipmentSystemIoControllerAllOf {
	this := EquipmentSystemIoControllerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentSystemIoControllerAllOfWithDefaults instantiates a new EquipmentSystemIoControllerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSystemIoControllerAllOfWithDefaults() *EquipmentSystemIoControllerAllOf {
	this := EquipmentSystemIoControllerAllOf{}
	var classId string = "equipment.SystemIoController"
	this.ClassId = classId
	var objectType string = "equipment.SystemIoController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentSystemIoControllerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentSystemIoControllerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentSystemIoControllerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentSystemIoControllerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetChassisId() string {
	if o == nil || o.ChassisId == nil {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetChassisIdOk() (*string, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *EquipmentSystemIoControllerAllOf) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetConnectionPath() string {
	if o == nil || o.ConnectionPath == nil {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetConnectionPathOk() (*string, bool) {
	if o == nil || o.ConnectionPath == nil {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasConnectionPath() bool {
	if o != nil && o.ConnectionPath != nil {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentSystemIoControllerAllOf) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentSystemIoControllerAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentSystemIoControllerAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetManagingInstance returns the ManagingInstance field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetManagingInstance() string {
	if o == nil || o.ManagingInstance == nil {
		var ret string
		return ret
	}
	return *o.ManagingInstance
}

// GetManagingInstanceOk returns a tuple with the ManagingInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetManagingInstanceOk() (*string, bool) {
	if o == nil || o.ManagingInstance == nil {
		return nil, false
	}
	return o.ManagingInstance, true
}

// HasManagingInstance returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasManagingInstance() bool {
	if o != nil && o.ManagingInstance != nil {
		return true
	}

	return false
}

// SetManagingInstance gets a reference to the given string and assigns it to the ManagingInstance field.
func (o *EquipmentSystemIoControllerAllOf) SetManagingInstance(v string) {
	o.ManagingInstance = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSystemIoControllerAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSystemIoControllerAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentSystemIoControllerAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetSystemIoControllerId returns the SystemIoControllerId field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetSystemIoControllerId() int64 {
	if o == nil || o.SystemIoControllerId == nil {
		var ret int64
		return ret
	}
	return *o.SystemIoControllerId
}

// GetSystemIoControllerIdOk returns a tuple with the SystemIoControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetSystemIoControllerIdOk() (*int64, bool) {
	if o == nil || o.SystemIoControllerId == nil {
		return nil, false
	}
	return o.SystemIoControllerId, true
}

// HasSystemIoControllerId returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasSystemIoControllerId() bool {
	if o != nil && o.SystemIoControllerId != nil {
		return true
	}

	return false
}

// SetSystemIoControllerId gets a reference to the given int64 and assigns it to the SystemIoControllerId field.
func (o *EquipmentSystemIoControllerAllOf) SetSystemIoControllerId(v int64) {
	o.SystemIoControllerId = &v
}

// GetCmc returns the Cmc field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetCmc() ManagementControllerRelationship {
	if o == nil || o.Cmc == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Cmc
}

// GetCmcOk returns a tuple with the Cmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetCmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Cmc == nil {
		return nil, false
	}
	return o.Cmc, true
}

// HasCmc returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasCmc() bool {
	if o != nil && o.Cmc != nil {
		return true
	}

	return false
}

// SetCmc gets a reference to the given ManagementControllerRelationship and assigns it to the Cmc field.
func (o *EquipmentSystemIoControllerAllOf) SetCmc(v ManagementControllerRelationship) {
	o.Cmc = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentSystemIoControllerAllOf) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSystemIoControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSystemIoControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSharedIoModule returns the SharedIoModule field value if set, zero value otherwise.
func (o *EquipmentSystemIoControllerAllOf) GetSharedIoModule() EquipmentSharedIoModuleRelationship {
	if o == nil || o.SharedIoModule == nil {
		var ret EquipmentSharedIoModuleRelationship
		return ret
	}
	return *o.SharedIoModule
}

// GetSharedIoModuleOk returns a tuple with the SharedIoModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSystemIoControllerAllOf) GetSharedIoModuleOk() (*EquipmentSharedIoModuleRelationship, bool) {
	if o == nil || o.SharedIoModule == nil {
		return nil, false
	}
	return o.SharedIoModule, true
}

// HasSharedIoModule returns a boolean if a field has been set.
func (o *EquipmentSystemIoControllerAllOf) HasSharedIoModule() bool {
	if o != nil && o.SharedIoModule != nil {
		return true
	}

	return false
}

// SetSharedIoModule gets a reference to the given EquipmentSharedIoModuleRelationship and assigns it to the SharedIoModule field.
func (o *EquipmentSystemIoControllerAllOf) SetSharedIoModule(v EquipmentSharedIoModuleRelationship) {
	o.SharedIoModule = &v
}

func (o EquipmentSystemIoControllerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.ConnectionPath != nil {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ManagingInstance != nil {
		toSerialize["ManagingInstance"] = o.ManagingInstance
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SystemIoControllerId != nil {
		toSerialize["SystemIoControllerId"] = o.SystemIoControllerId
	}
	if o.Cmc != nil {
		toSerialize["Cmc"] = o.Cmc
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SharedIoModule != nil {
		toSerialize["SharedIoModule"] = o.SharedIoModule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSystemIoControllerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentSystemIoControllerAllOf := _EquipmentSystemIoControllerAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentSystemIoControllerAllOf); err == nil {
		*o = EquipmentSystemIoControllerAllOf(varEquipmentSystemIoControllerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ManagingInstance")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SystemIoControllerId")
		delete(additionalProperties, "Cmc")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SharedIoModule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentSystemIoControllerAllOf struct {
	value *EquipmentSystemIoControllerAllOf
	isSet bool
}

func (v NullableEquipmentSystemIoControllerAllOf) Get() *EquipmentSystemIoControllerAllOf {
	return v.value
}

func (v *NullableEquipmentSystemIoControllerAllOf) Set(val *EquipmentSystemIoControllerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSystemIoControllerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSystemIoControllerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSystemIoControllerAllOf(val *EquipmentSystemIoControllerAllOf) *NullableEquipmentSystemIoControllerAllOf {
	return &NullableEquipmentSystemIoControllerAllOf{value: val, isSet: true}
}

func (v NullableEquipmentSystemIoControllerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSystemIoControllerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
