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

// checks if the FabricAbstractSpanSourcePortChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricAbstractSpanSourcePortChannel{}

// FabricAbstractSpanSourcePortChannel Abstract class for all SPAN Source Port Channels including Ethernet and FC.
type FabricAbstractSpanSourcePortChannel struct {
	FabricAbstractSpanSource
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Port-channel ID of SPAN source.
	PcId *int64 `json:"PcId,omitempty"`
	// Role of the port-channel configured as SPAN source. * `Uplink` - Uplink Role corresponding to PortRole in PortPolicy. * `FcoeUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `FcUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `Appliance` - FcoeUplink Role corresponding to PortRole in PortPolicy.
	SourceRole           *string                                `json:"SourceRole,omitempty"`
	PhysicalPortChannel  NullableInventoryInterfaceRelationship `json:"PhysicalPortChannel,omitempty"`
	SpanSession          NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricAbstractSpanSourcePortChannel FabricAbstractSpanSourcePortChannel

// NewFabricAbstractSpanSourcePortChannel instantiates a new FabricAbstractSpanSourcePortChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAbstractSpanSourcePortChannel(classId string, objectType string) *FabricAbstractSpanSourcePortChannel {
	this := FabricAbstractSpanSourcePortChannel{}
	this.ClassId = classId
	this.ObjectType = objectType
	var direction string = "Receive"
	this.Direction = &direction
	return &this
}

// NewFabricAbstractSpanSourcePortChannelWithDefaults instantiates a new FabricAbstractSpanSourcePortChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAbstractSpanSourcePortChannelWithDefaults() *FabricAbstractSpanSourcePortChannel {
	this := FabricAbstractSpanSourcePortChannel{}
	var classId string = "fabric.SpanSourceEthPortChannel"
	this.ClassId = classId
	var objectType string = "fabric.SpanSourceEthPortChannel"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAbstractSpanSourcePortChannel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePortChannel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAbstractSpanSourcePortChannel) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.SpanSourceEthPortChannel" of the ClassId field.
func (o *FabricAbstractSpanSourcePortChannel) GetDefaultClassId() interface{} {
	return "fabric.SpanSourceEthPortChannel"
}

// GetObjectType returns the ObjectType field value
func (o *FabricAbstractSpanSourcePortChannel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePortChannel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAbstractSpanSourcePortChannel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.SpanSourceEthPortChannel" of the ObjectType field.
func (o *FabricAbstractSpanSourcePortChannel) GetDefaultObjectType() interface{} {
	return "fabric.SpanSourceEthPortChannel"
}

// GetPcId returns the PcId field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePortChannel) GetPcId() int64 {
	if o == nil || IsNil(o.PcId) {
		var ret int64
		return ret
	}
	return *o.PcId
}

// GetPcIdOk returns a tuple with the PcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePortChannel) GetPcIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PcId) {
		return nil, false
	}
	return o.PcId, true
}

// HasPcId returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePortChannel) HasPcId() bool {
	if o != nil && !IsNil(o.PcId) {
		return true
	}

	return false
}

// SetPcId gets a reference to the given int64 and assigns it to the PcId field.
func (o *FabricAbstractSpanSourcePortChannel) SetPcId(v int64) {
	o.PcId = &v
}

// GetSourceRole returns the SourceRole field value if set, zero value otherwise.
func (o *FabricAbstractSpanSourcePortChannel) GetSourceRole() string {
	if o == nil || IsNil(o.SourceRole) {
		var ret string
		return ret
	}
	return *o.SourceRole
}

// GetSourceRoleOk returns a tuple with the SourceRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAbstractSpanSourcePortChannel) GetSourceRoleOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRole) {
		return nil, false
	}
	return o.SourceRole, true
}

// HasSourceRole returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePortChannel) HasSourceRole() bool {
	if o != nil && !IsNil(o.SourceRole) {
		return true
	}

	return false
}

// SetSourceRole gets a reference to the given string and assigns it to the SourceRole field.
func (o *FabricAbstractSpanSourcePortChannel) SetSourceRole(v string) {
	o.SourceRole = &v
}

// GetPhysicalPortChannel returns the PhysicalPortChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSourcePortChannel) GetPhysicalPortChannel() InventoryInterfaceRelationship {
	if o == nil || IsNil(o.PhysicalPortChannel.Get()) {
		var ret InventoryInterfaceRelationship
		return ret
	}
	return *o.PhysicalPortChannel.Get()
}

// GetPhysicalPortChannelOk returns a tuple with the PhysicalPortChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSourcePortChannel) GetPhysicalPortChannelOk() (*InventoryInterfaceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhysicalPortChannel.Get(), o.PhysicalPortChannel.IsSet()
}

// HasPhysicalPortChannel returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePortChannel) HasPhysicalPortChannel() bool {
	if o != nil && o.PhysicalPortChannel.IsSet() {
		return true
	}

	return false
}

// SetPhysicalPortChannel gets a reference to the given NullableInventoryInterfaceRelationship and assigns it to the PhysicalPortChannel field.
func (o *FabricAbstractSpanSourcePortChannel) SetPhysicalPortChannel(v InventoryInterfaceRelationship) {
	o.PhysicalPortChannel.Set(&v)
}

// SetPhysicalPortChannelNil sets the value for PhysicalPortChannel to be an explicit nil
func (o *FabricAbstractSpanSourcePortChannel) SetPhysicalPortChannelNil() {
	o.PhysicalPortChannel.Set(nil)
}

// UnsetPhysicalPortChannel ensures that no value is present for PhysicalPortChannel, not even an explicit nil
func (o *FabricAbstractSpanSourcePortChannel) UnsetPhysicalPortChannel() {
	o.PhysicalPortChannel.Unset()
}

// GetSpanSession returns the SpanSession field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricAbstractSpanSourcePortChannel) GetSpanSession() FabricSpanSessionRelationship {
	if o == nil || IsNil(o.SpanSession.Get()) {
		var ret FabricSpanSessionRelationship
		return ret
	}
	return *o.SpanSession.Get()
}

// GetSpanSessionOk returns a tuple with the SpanSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricAbstractSpanSourcePortChannel) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpanSession.Get(), o.SpanSession.IsSet()
}

// HasSpanSession returns a boolean if a field has been set.
func (o *FabricAbstractSpanSourcePortChannel) HasSpanSession() bool {
	if o != nil && o.SpanSession.IsSet() {
		return true
	}

	return false
}

// SetSpanSession gets a reference to the given NullableFabricSpanSessionRelationship and assigns it to the SpanSession field.
func (o *FabricAbstractSpanSourcePortChannel) SetSpanSession(v FabricSpanSessionRelationship) {
	o.SpanSession.Set(&v)
}

// SetSpanSessionNil sets the value for SpanSession to be an explicit nil
func (o *FabricAbstractSpanSourcePortChannel) SetSpanSessionNil() {
	o.SpanSession.Set(nil)
}

// UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil
func (o *FabricAbstractSpanSourcePortChannel) UnsetSpanSession() {
	o.SpanSession.Unset()
}

func (o FabricAbstractSpanSourcePortChannel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricAbstractSpanSourcePortChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricAbstractSpanSource, errFabricAbstractSpanSource := json.Marshal(o.FabricAbstractSpanSource)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	errFabricAbstractSpanSource = json.Unmarshal([]byte(serializedFabricAbstractSpanSource), &toSerialize)
	if errFabricAbstractSpanSource != nil {
		return map[string]interface{}{}, errFabricAbstractSpanSource
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.PcId) {
		toSerialize["PcId"] = o.PcId
	}
	if !IsNil(o.SourceRole) {
		toSerialize["SourceRole"] = o.SourceRole
	}
	if o.PhysicalPortChannel.IsSet() {
		toSerialize["PhysicalPortChannel"] = o.PhysicalPortChannel.Get()
	}
	if o.SpanSession.IsSet() {
		toSerialize["SpanSession"] = o.SpanSession.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricAbstractSpanSourcePortChannel) UnmarshalJSON(data []byte) (err error) {
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
	type FabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Port-channel ID of SPAN source.
		PcId *int64 `json:"PcId,omitempty"`
		// Role of the port-channel configured as SPAN source. * `Uplink` - Uplink Role corresponding to PortRole in PortPolicy. * `FcoeUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `FcUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy. * `Appliance` - FcoeUplink Role corresponding to PortRole in PortPolicy.
		SourceRole          *string                                `json:"SourceRole,omitempty"`
		PhysicalPortChannel NullableInventoryInterfaceRelationship `json:"PhysicalPortChannel,omitempty"`
		SpanSession         NullableFabricSpanSessionRelationship  `json:"SpanSession,omitempty"`
	}

	varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct := FabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct)
	if err == nil {
		varFabricAbstractSpanSourcePortChannel := _FabricAbstractSpanSourcePortChannel{}
		varFabricAbstractSpanSourcePortChannel.ClassId = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.ClassId
		varFabricAbstractSpanSourcePortChannel.ObjectType = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.ObjectType
		varFabricAbstractSpanSourcePortChannel.PcId = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.PcId
		varFabricAbstractSpanSourcePortChannel.SourceRole = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.SourceRole
		varFabricAbstractSpanSourcePortChannel.PhysicalPortChannel = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.PhysicalPortChannel
		varFabricAbstractSpanSourcePortChannel.SpanSession = varFabricAbstractSpanSourcePortChannelWithoutEmbeddedStruct.SpanSession
		*o = FabricAbstractSpanSourcePortChannel(varFabricAbstractSpanSourcePortChannel)
	} else {
		return err
	}

	varFabricAbstractSpanSourcePortChannel := _FabricAbstractSpanSourcePortChannel{}

	err = json.Unmarshal(data, &varFabricAbstractSpanSourcePortChannel)
	if err == nil {
		o.FabricAbstractSpanSource = varFabricAbstractSpanSourcePortChannel.FabricAbstractSpanSource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PcId")
		delete(additionalProperties, "SourceRole")
		delete(additionalProperties, "PhysicalPortChannel")
		delete(additionalProperties, "SpanSession")

		// remove fields from embedded structs
		reflectFabricAbstractSpanSource := reflect.ValueOf(o.FabricAbstractSpanSource)
		for i := 0; i < reflectFabricAbstractSpanSource.Type().NumField(); i++ {
			t := reflectFabricAbstractSpanSource.Type().Field(i)

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

type NullableFabricAbstractSpanSourcePortChannel struct {
	value *FabricAbstractSpanSourcePortChannel
	isSet bool
}

func (v NullableFabricAbstractSpanSourcePortChannel) Get() *FabricAbstractSpanSourcePortChannel {
	return v.value
}

func (v *NullableFabricAbstractSpanSourcePortChannel) Set(val *FabricAbstractSpanSourcePortChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAbstractSpanSourcePortChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAbstractSpanSourcePortChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAbstractSpanSourcePortChannel(val *FabricAbstractSpanSourcePortChannel) *NullableFabricAbstractSpanSourcePortChannel {
	return &NullableFabricAbstractSpanSourcePortChannel{value: val, isSet: true}
}

func (v NullableFabricAbstractSpanSourcePortChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAbstractSpanSourcePortChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
