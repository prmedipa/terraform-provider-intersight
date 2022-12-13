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

// FabricVsanInventory Inventory object for VSAN.
type FabricVsanInventory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin state of the VSAN. It can be active or inactive. * `` - Default value to represent the administrative state of isolated vsan. * `Up` - VSAN administrative state is up. * `Down` - VSAN administrative state is down.
	AdminState *string `json:"AdminState,omitempty"`
	// Interoperability mode of the VSAN. It enables products of multiple vendors to communicate with each other. * `` - Default value to represent the interoperability mode of isolated vsan when it is not configured. * `Default` - Default mode for a VSAN that is communicating between a SAN composed entirely of MDS 9000 switches. * `1` - Allows integration with Brocade and McData switches that have been configured for their own interoperability modes. Brocade and McData switches must be running in interop mode to work with this VSAN mode. * `2` - Allows seamless integration with specific Brocade switches running in their own native mode of operation. * `3` - Allows seamless integration with Brocade switches that contains more than 16 ports. * `4` - Allows seamless integration between MDS VSANs and McData switches running in McData Fabric 1.0 interop mode.
	InteropMode *string `json:"InteropMode,omitempty"`
	// These attributes indicate the use of the source-destination ID or the originator exchange OX ID for load balancing path selection. * `` - Default value to represent the load balancing scheme of isolated vsan. * `src-id/dst-id` - Directs the switch to use the source and destination ID for its path selection process. * `src-id/dst-id/oxid` - Directs the switch to use the source ID, the destination ID, and the OX ID for its path selection process.
	LoadBalancing *string `json:"LoadBalancing,omitempty"`
	// User-specified name of the VSAN.
	Name *string `json:"Name,omitempty"`
	// Operational state of the VSAN. * `` - Default value to represent the operational state of isolated vsan. * `Up` - VSAN operational state is up. * `Down` - VSAN operational state is down.
	OperState *string `json:"OperState,omitempty"`
	// Smart zoning status on the VSAN. It can be enabled or disabled. * `` - Default value to represent the smart zoning status of isolated vsan. * `Enabled` - VSAN smart zoning state is enabled. * `Disabled` - VSAN smart zoning state is disabled.
	SmartZoning *string `json:"SmartZoning,omitempty"`
	// Identifier for the VSAN. It is an integer from 1 to 4094.
	VsanId               *int64                               `json:"VsanId,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricVsanInventory FabricVsanInventory

// NewFabricVsanInventory instantiates a new FabricVsanInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricVsanInventory(classId string, objectType string) *FabricVsanInventory {
	this := FabricVsanInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricVsanInventoryWithDefaults instantiates a new FabricVsanInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricVsanInventoryWithDefaults() *FabricVsanInventory {
	this := FabricVsanInventory{}
	var classId string = "fabric.VsanInventory"
	this.ClassId = classId
	var objectType string = "fabric.VsanInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricVsanInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricVsanInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricVsanInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricVsanInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *FabricVsanInventory) SetAdminState(v string) {
	o.AdminState = &v
}

// GetInteropMode returns the InteropMode field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetInteropMode() string {
	if o == nil || o.InteropMode == nil {
		var ret string
		return ret
	}
	return *o.InteropMode
}

// GetInteropModeOk returns a tuple with the InteropMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetInteropModeOk() (*string, bool) {
	if o == nil || o.InteropMode == nil {
		return nil, false
	}
	return o.InteropMode, true
}

// HasInteropMode returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasInteropMode() bool {
	if o != nil && o.InteropMode != nil {
		return true
	}

	return false
}

// SetInteropMode gets a reference to the given string and assigns it to the InteropMode field.
func (o *FabricVsanInventory) SetInteropMode(v string) {
	o.InteropMode = &v
}

// GetLoadBalancing returns the LoadBalancing field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetLoadBalancing() string {
	if o == nil || o.LoadBalancing == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancing
}

// GetLoadBalancingOk returns a tuple with the LoadBalancing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetLoadBalancingOk() (*string, bool) {
	if o == nil || o.LoadBalancing == nil {
		return nil, false
	}
	return o.LoadBalancing, true
}

// HasLoadBalancing returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasLoadBalancing() bool {
	if o != nil && o.LoadBalancing != nil {
		return true
	}

	return false
}

// SetLoadBalancing gets a reference to the given string and assigns it to the LoadBalancing field.
func (o *FabricVsanInventory) SetLoadBalancing(v string) {
	o.LoadBalancing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricVsanInventory) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *FabricVsanInventory) SetOperState(v string) {
	o.OperState = &v
}

// GetSmartZoning returns the SmartZoning field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetSmartZoning() string {
	if o == nil || o.SmartZoning == nil {
		var ret string
		return ret
	}
	return *o.SmartZoning
}

// GetSmartZoningOk returns a tuple with the SmartZoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetSmartZoningOk() (*string, bool) {
	if o == nil || o.SmartZoning == nil {
		return nil, false
	}
	return o.SmartZoning, true
}

// HasSmartZoning returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasSmartZoning() bool {
	if o != nil && o.SmartZoning != nil {
		return true
	}

	return false
}

// SetSmartZoning gets a reference to the given string and assigns it to the SmartZoning field.
func (o *FabricVsanInventory) SetSmartZoning(v string) {
	o.SmartZoning = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricVsanInventory) SetVsanId(v int64) {
	o.VsanId = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *FabricVsanInventory) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *FabricVsanInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricVsanInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *FabricVsanInventory) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *FabricVsanInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o FabricVsanInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.InteropMode != nil {
		toSerialize["InteropMode"] = o.InteropMode
	}
	if o.LoadBalancing != nil {
		toSerialize["LoadBalancing"] = o.LoadBalancing
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.SmartZoning != nil {
		toSerialize["SmartZoning"] = o.SmartZoning
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricVsanInventory) UnmarshalJSON(bytes []byte) (err error) {
	type FabricVsanInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin state of the VSAN. It can be active or inactive. * `` - Default value to represent the administrative state of isolated vsan. * `Up` - VSAN administrative state is up. * `Down` - VSAN administrative state is down.
		AdminState *string `json:"AdminState,omitempty"`
		// Interoperability mode of the VSAN. It enables products of multiple vendors to communicate with each other. * `` - Default value to represent the interoperability mode of isolated vsan when it is not configured. * `Default` - Default mode for a VSAN that is communicating between a SAN composed entirely of MDS 9000 switches. * `1` - Allows integration with Brocade and McData switches that have been configured for their own interoperability modes. Brocade and McData switches must be running in interop mode to work with this VSAN mode. * `2` - Allows seamless integration with specific Brocade switches running in their own native mode of operation. * `3` - Allows seamless integration with Brocade switches that contains more than 16 ports. * `4` - Allows seamless integration between MDS VSANs and McData switches running in McData Fabric 1.0 interop mode.
		InteropMode *string `json:"InteropMode,omitempty"`
		// These attributes indicate the use of the source-destination ID or the originator exchange OX ID for load balancing path selection. * `` - Default value to represent the load balancing scheme of isolated vsan. * `src-id/dst-id` - Directs the switch to use the source and destination ID for its path selection process. * `src-id/dst-id/oxid` - Directs the switch to use the source ID, the destination ID, and the OX ID for its path selection process.
		LoadBalancing *string `json:"LoadBalancing,omitempty"`
		// User-specified name of the VSAN.
		Name *string `json:"Name,omitempty"`
		// Operational state of the VSAN. * `` - Default value to represent the operational state of isolated vsan. * `Up` - VSAN operational state is up. * `Down` - VSAN operational state is down.
		OperState *string `json:"OperState,omitempty"`
		// Smart zoning status on the VSAN. It can be enabled or disabled. * `` - Default value to represent the smart zoning status of isolated vsan. * `Enabled` - VSAN smart zoning state is enabled. * `Disabled` - VSAN smart zoning state is disabled.
		SmartZoning *string `json:"SmartZoning,omitempty"`
		// Identifier for the VSAN. It is an integer from 1 to 4094.
		VsanId           *int64                               `json:"VsanId,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varFabricVsanInventoryWithoutEmbeddedStruct := FabricVsanInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricVsanInventoryWithoutEmbeddedStruct)
	if err == nil {
		varFabricVsanInventory := _FabricVsanInventory{}
		varFabricVsanInventory.ClassId = varFabricVsanInventoryWithoutEmbeddedStruct.ClassId
		varFabricVsanInventory.ObjectType = varFabricVsanInventoryWithoutEmbeddedStruct.ObjectType
		varFabricVsanInventory.AdminState = varFabricVsanInventoryWithoutEmbeddedStruct.AdminState
		varFabricVsanInventory.InteropMode = varFabricVsanInventoryWithoutEmbeddedStruct.InteropMode
		varFabricVsanInventory.LoadBalancing = varFabricVsanInventoryWithoutEmbeddedStruct.LoadBalancing
		varFabricVsanInventory.Name = varFabricVsanInventoryWithoutEmbeddedStruct.Name
		varFabricVsanInventory.OperState = varFabricVsanInventoryWithoutEmbeddedStruct.OperState
		varFabricVsanInventory.SmartZoning = varFabricVsanInventoryWithoutEmbeddedStruct.SmartZoning
		varFabricVsanInventory.VsanId = varFabricVsanInventoryWithoutEmbeddedStruct.VsanId
		varFabricVsanInventory.NetworkElement = varFabricVsanInventoryWithoutEmbeddedStruct.NetworkElement
		varFabricVsanInventory.RegisteredDevice = varFabricVsanInventoryWithoutEmbeddedStruct.RegisteredDevice
		*o = FabricVsanInventory(varFabricVsanInventory)
	} else {
		return err
	}

	varFabricVsanInventory := _FabricVsanInventory{}

	err = json.Unmarshal(bytes, &varFabricVsanInventory)
	if err == nil {
		o.MoBaseMo = varFabricVsanInventory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "InteropMode")
		delete(additionalProperties, "LoadBalancing")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "SmartZoning")
		delete(additionalProperties, "VsanId")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableFabricVsanInventory struct {
	value *FabricVsanInventory
	isSet bool
}

func (v NullableFabricVsanInventory) Get() *FabricVsanInventory {
	return v.value
}

func (v *NullableFabricVsanInventory) Set(val *FabricVsanInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricVsanInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricVsanInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricVsanInventory(val *FabricVsanInventory) *NullableFabricVsanInventory {
	return &NullableFabricVsanInventory{value: val, isSet: true}
}

func (v NullableFabricVsanInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricVsanInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
