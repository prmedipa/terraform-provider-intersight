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

// checks if the ApplianceNetworkLinkStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceNetworkLinkStatus{}

// ApplianceNetworkLinkStatus Link between two nodes of an Intersight Appliance cluster.
type ApplianceNetworkLinkStatus struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Hostname of the destination endpoint.
	DestinationHostname *string `json:"DestinationHostname,omitempty"`
	// Time to reach the destination endpoint in milliseconds from the source endpoint.
	PingTime *float32 `json:"PingTime,omitempty"`
	// Hostname of the source endpoint.
	SourceHostname       *string                                     `json:"SourceHostname,omitempty"`
	NodeOpStatus         NullableApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceNetworkLinkStatus ApplianceNetworkLinkStatus

// NewApplianceNetworkLinkStatus instantiates a new ApplianceNetworkLinkStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceNetworkLinkStatus(classId string, objectType string) *ApplianceNetworkLinkStatus {
	this := ApplianceNetworkLinkStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceNetworkLinkStatusWithDefaults instantiates a new ApplianceNetworkLinkStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceNetworkLinkStatusWithDefaults() *ApplianceNetworkLinkStatus {
	this := ApplianceNetworkLinkStatus{}
	var classId string = "appliance.NetworkLinkStatus"
	this.ClassId = classId
	var objectType string = "appliance.NetworkLinkStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceNetworkLinkStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceNetworkLinkStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.NetworkLinkStatus" of the ClassId field.
func (o *ApplianceNetworkLinkStatus) GetDefaultClassId() interface{} {
	return "appliance.NetworkLinkStatus"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceNetworkLinkStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceNetworkLinkStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.NetworkLinkStatus" of the ObjectType field.
func (o *ApplianceNetworkLinkStatus) GetDefaultObjectType() interface{} {
	return "appliance.NetworkLinkStatus"
}

// GetDestinationHostname returns the DestinationHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetDestinationHostname() string {
	if o == nil || IsNil(o.DestinationHostname) {
		var ret string
		return ret
	}
	return *o.DestinationHostname
}

// GetDestinationHostnameOk returns a tuple with the DestinationHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetDestinationHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationHostname) {
		return nil, false
	}
	return o.DestinationHostname, true
}

// HasDestinationHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasDestinationHostname() bool {
	if o != nil && !IsNil(o.DestinationHostname) {
		return true
	}

	return false
}

// SetDestinationHostname gets a reference to the given string and assigns it to the DestinationHostname field.
func (o *ApplianceNetworkLinkStatus) SetDestinationHostname(v string) {
	o.DestinationHostname = &v
}

// GetPingTime returns the PingTime field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetPingTime() float32 {
	if o == nil || IsNil(o.PingTime) {
		var ret float32
		return ret
	}
	return *o.PingTime
}

// GetPingTimeOk returns a tuple with the PingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetPingTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.PingTime) {
		return nil, false
	}
	return o.PingTime, true
}

// HasPingTime returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasPingTime() bool {
	if o != nil && !IsNil(o.PingTime) {
		return true
	}

	return false
}

// SetPingTime gets a reference to the given float32 and assigns it to the PingTime field.
func (o *ApplianceNetworkLinkStatus) SetPingTime(v float32) {
	o.PingTime = &v
}

// GetSourceHostname returns the SourceHostname field value if set, zero value otherwise.
func (o *ApplianceNetworkLinkStatus) GetSourceHostname() string {
	if o == nil || IsNil(o.SourceHostname) {
		var ret string
		return ret
	}
	return *o.SourceHostname
}

// GetSourceHostnameOk returns a tuple with the SourceHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceNetworkLinkStatus) GetSourceHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceHostname) {
		return nil, false
	}
	return o.SourceHostname, true
}

// HasSourceHostname returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasSourceHostname() bool {
	if o != nil && !IsNil(o.SourceHostname) {
		return true
	}

	return false
}

// SetSourceHostname gets a reference to the given string and assigns it to the SourceHostname field.
func (o *ApplianceNetworkLinkStatus) SetSourceHostname(v string) {
	o.SourceHostname = &v
}

// GetNodeOpStatus returns the NodeOpStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNetworkLinkStatus) GetNodeOpStatus() ApplianceNodeOpStatusRelationship {
	if o == nil || IsNil(o.NodeOpStatus.Get()) {
		var ret ApplianceNodeOpStatusRelationship
		return ret
	}
	return *o.NodeOpStatus.Get()
}

// GetNodeOpStatusOk returns a tuple with the NodeOpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNetworkLinkStatus) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeOpStatus.Get(), o.NodeOpStatus.IsSet()
}

// HasNodeOpStatus returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasNodeOpStatus() bool {
	if o != nil && o.NodeOpStatus.IsSet() {
		return true
	}

	return false
}

// SetNodeOpStatus gets a reference to the given NullableApplianceNodeOpStatusRelationship and assigns it to the NodeOpStatus field.
func (o *ApplianceNetworkLinkStatus) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship) {
	o.NodeOpStatus.Set(&v)
}

// SetNodeOpStatusNil sets the value for NodeOpStatus to be an explicit nil
func (o *ApplianceNetworkLinkStatus) SetNodeOpStatusNil() {
	o.NodeOpStatus.Set(nil)
}

// UnsetNodeOpStatus ensures that no value is present for NodeOpStatus, not even an explicit nil
func (o *ApplianceNetworkLinkStatus) UnsetNodeOpStatus() {
	o.NodeOpStatus.Unset()
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceNetworkLinkStatus) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceNetworkLinkStatus) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ApplianceNetworkLinkStatus) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ApplianceNetworkLinkStatus) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *ApplianceNetworkLinkStatus) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *ApplianceNetworkLinkStatus) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o ApplianceNetworkLinkStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceNetworkLinkStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.DestinationHostname) {
		toSerialize["DestinationHostname"] = o.DestinationHostname
	}
	if !IsNil(o.PingTime) {
		toSerialize["PingTime"] = o.PingTime
	}
	if !IsNil(o.SourceHostname) {
		toSerialize["SourceHostname"] = o.SourceHostname
	}
	if o.NodeOpStatus.IsSet() {
		toSerialize["NodeOpStatus"] = o.NodeOpStatus.Get()
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceNetworkLinkStatus) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceNetworkLinkStatusWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Hostname of the destination endpoint.
		DestinationHostname *string `json:"DestinationHostname,omitempty"`
		// Time to reach the destination endpoint in milliseconds from the source endpoint.
		PingTime *float32 `json:"PingTime,omitempty"`
		// Hostname of the source endpoint.
		SourceHostname   *string                                     `json:"SourceHostname,omitempty"`
		NodeOpStatus     NullableApplianceNodeOpStatusRelationship   `json:"NodeOpStatus,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varApplianceNetworkLinkStatusWithoutEmbeddedStruct := ApplianceNetworkLinkStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceNetworkLinkStatusWithoutEmbeddedStruct)
	if err == nil {
		varApplianceNetworkLinkStatus := _ApplianceNetworkLinkStatus{}
		varApplianceNetworkLinkStatus.ClassId = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.ClassId
		varApplianceNetworkLinkStatus.ObjectType = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.ObjectType
		varApplianceNetworkLinkStatus.DestinationHostname = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.DestinationHostname
		varApplianceNetworkLinkStatus.PingTime = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.PingTime
		varApplianceNetworkLinkStatus.SourceHostname = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.SourceHostname
		varApplianceNetworkLinkStatus.NodeOpStatus = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.NodeOpStatus
		varApplianceNetworkLinkStatus.RegisteredDevice = varApplianceNetworkLinkStatusWithoutEmbeddedStruct.RegisteredDevice
		*o = ApplianceNetworkLinkStatus(varApplianceNetworkLinkStatus)
	} else {
		return err
	}

	varApplianceNetworkLinkStatus := _ApplianceNetworkLinkStatus{}

	err = json.Unmarshal(data, &varApplianceNetworkLinkStatus)
	if err == nil {
		o.MoBaseMo = varApplianceNetworkLinkStatus.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DestinationHostname")
		delete(additionalProperties, "PingTime")
		delete(additionalProperties, "SourceHostname")
		delete(additionalProperties, "NodeOpStatus")
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

type NullableApplianceNetworkLinkStatus struct {
	value *ApplianceNetworkLinkStatus
	isSet bool
}

func (v NullableApplianceNetworkLinkStatus) Get() *ApplianceNetworkLinkStatus {
	return v.value
}

func (v *NullableApplianceNetworkLinkStatus) Set(val *ApplianceNetworkLinkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceNetworkLinkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceNetworkLinkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceNetworkLinkStatus(val *ApplianceNetworkLinkStatus) *NullableApplianceNetworkLinkStatus {
	return &NullableApplianceNetworkLinkStatus{value: val, isSet: true}
}

func (v NullableApplianceNetworkLinkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceNetworkLinkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
