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

// checks if the VnicVlanSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnicVlanSettings{}

// VnicVlanSettings VLAN configuration for the virtual interface.
type VnicVlanSettings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Allowed VLAN IDs of the virtual interface.
	AllowedVlans *string `json:"AllowedVlans,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
	// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface.
	DefaultVlan *int64 `json:"DefaultVlan,omitempty"`
	// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. * `ACCESS` - An access port carries traffic only for a single VLAN on the interface. * `TRUNK` - A trunk port can have two or more VLANs configured on the interface. It can carry traffic for several VLANs simultaneously.
	Mode *string `json:"Mode,omitempty"`
	// Enable QinQ (802.1Q-in-802.1Q) Tunneling on the vNIC.
	QinqEnabled *bool `json:"QinqEnabled,omitempty"`
	// When activating VIC QinQ (802.1Q-in-802.1Q) Tunneling, a particular VLAN ID is set. In Access VLAN mode, this QinQ VLAN ID is established as the default VLAN.
	QinqVlan             *int64 `json:"QinqVlan,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicVlanSettings VnicVlanSettings

// NewVnicVlanSettings instantiates a new VnicVlanSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicVlanSettings(classId string, objectType string) *VnicVlanSettings {
	this := VnicVlanSettings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var defaultVlan int64 = 0
	this.DefaultVlan = &defaultVlan
	var mode string = "ACCESS"
	this.Mode = &mode
	var qinqEnabled bool = false
	this.QinqEnabled = &qinqEnabled
	var qinqVlan int64 = 2
	this.QinqVlan = &qinqVlan
	return &this
}

// NewVnicVlanSettingsWithDefaults instantiates a new VnicVlanSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicVlanSettingsWithDefaults() *VnicVlanSettings {
	this := VnicVlanSettings{}
	var classId string = "vnic.VlanSettings"
	this.ClassId = classId
	var objectType string = "vnic.VlanSettings"
	this.ObjectType = objectType
	var defaultVlan int64 = 0
	this.DefaultVlan = &defaultVlan
	var mode string = "ACCESS"
	this.Mode = &mode
	var qinqEnabled bool = false
	this.QinqEnabled = &qinqEnabled
	var qinqVlan int64 = 2
	this.QinqVlan = &qinqVlan
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicVlanSettings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicVlanSettings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "vnic.VlanSettings" of the ClassId field.
func (o *VnicVlanSettings) GetDefaultClassId() interface{} {
	return "vnic.VlanSettings"
}

// GetObjectType returns the ObjectType field value
func (o *VnicVlanSettings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicVlanSettings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "vnic.VlanSettings" of the ObjectType field.
func (o *VnicVlanSettings) GetDefaultObjectType() interface{} {
	return "vnic.VlanSettings"
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *VnicVlanSettings) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *VnicVlanSettings) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *VnicVlanSettings) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetDefaultVlan returns the DefaultVlan field value if set, zero value otherwise.
func (o *VnicVlanSettings) GetDefaultVlan() int64 {
	if o == nil || IsNil(o.DefaultVlan) {
		var ret int64
		return ret
	}
	return *o.DefaultVlan
}

// GetDefaultVlanOk returns a tuple with the DefaultVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetDefaultVlanOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultVlan) {
		return nil, false
	}
	return o.DefaultVlan, true
}

// HasDefaultVlan returns a boolean if a field has been set.
func (o *VnicVlanSettings) HasDefaultVlan() bool {
	if o != nil && !IsNil(o.DefaultVlan) {
		return true
	}

	return false
}

// SetDefaultVlan gets a reference to the given int64 and assigns it to the DefaultVlan field.
func (o *VnicVlanSettings) SetDefaultVlan(v int64) {
	o.DefaultVlan = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *VnicVlanSettings) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *VnicVlanSettings) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *VnicVlanSettings) SetMode(v string) {
	o.Mode = &v
}

// GetQinqEnabled returns the QinqEnabled field value if set, zero value otherwise.
func (o *VnicVlanSettings) GetQinqEnabled() bool {
	if o == nil || IsNil(o.QinqEnabled) {
		var ret bool
		return ret
	}
	return *o.QinqEnabled
}

// GetQinqEnabledOk returns a tuple with the QinqEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetQinqEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.QinqEnabled) {
		return nil, false
	}
	return o.QinqEnabled, true
}

// HasQinqEnabled returns a boolean if a field has been set.
func (o *VnicVlanSettings) HasQinqEnabled() bool {
	if o != nil && !IsNil(o.QinqEnabled) {
		return true
	}

	return false
}

// SetQinqEnabled gets a reference to the given bool and assigns it to the QinqEnabled field.
func (o *VnicVlanSettings) SetQinqEnabled(v bool) {
	o.QinqEnabled = &v
}

// GetQinqVlan returns the QinqVlan field value if set, zero value otherwise.
func (o *VnicVlanSettings) GetQinqVlan() int64 {
	if o == nil || IsNil(o.QinqVlan) {
		var ret int64
		return ret
	}
	return *o.QinqVlan
}

// GetQinqVlanOk returns a tuple with the QinqVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicVlanSettings) GetQinqVlanOk() (*int64, bool) {
	if o == nil || IsNil(o.QinqVlan) {
		return nil, false
	}
	return o.QinqVlan, true
}

// HasQinqVlan returns a boolean if a field has been set.
func (o *VnicVlanSettings) HasQinqVlan() bool {
	if o != nil && !IsNil(o.QinqVlan) {
		return true
	}

	return false
}

// SetQinqVlan gets a reference to the given int64 and assigns it to the QinqVlan field.
func (o *VnicVlanSettings) SetQinqVlan(v int64) {
	o.QinqVlan = &v
}

func (o VnicVlanSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnicVlanSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AllowedVlans) {
		toSerialize["AllowedVlans"] = o.AllowedVlans
	}
	if !IsNil(o.DefaultVlan) {
		toSerialize["DefaultVlan"] = o.DefaultVlan
	}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	if !IsNil(o.QinqEnabled) {
		toSerialize["QinqEnabled"] = o.QinqEnabled
	}
	if !IsNil(o.QinqVlan) {
		toSerialize["QinqVlan"] = o.QinqVlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VnicVlanSettings) UnmarshalJSON(data []byte) (err error) {
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
	type VnicVlanSettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Allowed VLAN IDs of the virtual interface.
		AllowedVlans *string `json:"AllowedVlans,omitempty" validate:"regexp=^$|^((\\\\d+\\\\-\\\\d+)|(\\\\d+))(,((\\\\d+\\\\-\\\\d+)|(\\\\d+)))*$"`
		// Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface.
		DefaultVlan *int64 `json:"DefaultVlan,omitempty"`
		// Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. * `ACCESS` - An access port carries traffic only for a single VLAN on the interface. * `TRUNK` - A trunk port can have two or more VLANs configured on the interface. It can carry traffic for several VLANs simultaneously.
		Mode *string `json:"Mode,omitempty"`
		// Enable QinQ (802.1Q-in-802.1Q) Tunneling on the vNIC.
		QinqEnabled *bool `json:"QinqEnabled,omitempty"`
		// When activating VIC QinQ (802.1Q-in-802.1Q) Tunneling, a particular VLAN ID is set. In Access VLAN mode, this QinQ VLAN ID is established as the default VLAN.
		QinqVlan *int64 `json:"QinqVlan,omitempty"`
	}

	varVnicVlanSettingsWithoutEmbeddedStruct := VnicVlanSettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVnicVlanSettingsWithoutEmbeddedStruct)
	if err == nil {
		varVnicVlanSettings := _VnicVlanSettings{}
		varVnicVlanSettings.ClassId = varVnicVlanSettingsWithoutEmbeddedStruct.ClassId
		varVnicVlanSettings.ObjectType = varVnicVlanSettingsWithoutEmbeddedStruct.ObjectType
		varVnicVlanSettings.AllowedVlans = varVnicVlanSettingsWithoutEmbeddedStruct.AllowedVlans
		varVnicVlanSettings.DefaultVlan = varVnicVlanSettingsWithoutEmbeddedStruct.DefaultVlan
		varVnicVlanSettings.Mode = varVnicVlanSettingsWithoutEmbeddedStruct.Mode
		varVnicVlanSettings.QinqEnabled = varVnicVlanSettingsWithoutEmbeddedStruct.QinqEnabled
		varVnicVlanSettings.QinqVlan = varVnicVlanSettingsWithoutEmbeddedStruct.QinqVlan
		*o = VnicVlanSettings(varVnicVlanSettings)
	} else {
		return err
	}

	varVnicVlanSettings := _VnicVlanSettings{}

	err = json.Unmarshal(data, &varVnicVlanSettings)
	if err == nil {
		o.MoBaseComplexType = varVnicVlanSettings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowedVlans")
		delete(additionalProperties, "DefaultVlan")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "QinqEnabled")
		delete(additionalProperties, "QinqVlan")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicVlanSettings struct {
	value *VnicVlanSettings
	isSet bool
}

func (v NullableVnicVlanSettings) Get() *VnicVlanSettings {
	return v.value
}

func (v *NullableVnicVlanSettings) Set(val *VnicVlanSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicVlanSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicVlanSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicVlanSettings(val *VnicVlanSettings) *NullableVnicVlanSettings {
	return &NullableVnicVlanSettings{value: val, isSet: true}
}

func (v NullableVnicVlanSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicVlanSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
