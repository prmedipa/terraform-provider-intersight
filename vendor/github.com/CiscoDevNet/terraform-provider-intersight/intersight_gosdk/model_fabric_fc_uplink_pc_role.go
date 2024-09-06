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

// checks if the FabricFcUplinkPcRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricFcUplinkPcRole{}

// FabricFcUplinkPcRole Object sent by user to configure a fc uplink port-channel on the collection of ports.
type FabricFcUplinkPcRole struct {
	FabricPortChannelRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `16Gbps` - Admin configurable speed 16Gbps. * `8Gbps` - Admin configurable speed 8Gbps. * `32Gbps` - Admin configurable speed 32Gbps. * `Auto` - Admin configurable speed AUTO ( default ).
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Fill pattern to differentiate the configs in NPIV. * `Idle` - Fc Fill Pattern type Idle. * `Arbff` - Fc Fill Pattern type Arbff.
	FillPattern *string `json:"FillPattern,omitempty"`
	// Virtual San Identifier associated to the FC port.
	VsanId               *int64 `json:"VsanId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcUplinkPcRole FabricFcUplinkPcRole

// NewFabricFcUplinkPcRole instantiates a new FabricFcUplinkPcRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcUplinkPcRole(classId string, objectType string) *FabricFcUplinkPcRole {
	this := FabricFcUplinkPcRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "16Gbps"
	this.AdminSpeed = &adminSpeed
	var fillPattern string = "Idle"
	this.FillPattern = &fillPattern
	return &this
}

// NewFabricFcUplinkPcRoleWithDefaults instantiates a new FabricFcUplinkPcRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcUplinkPcRoleWithDefaults() *FabricFcUplinkPcRole {
	this := FabricFcUplinkPcRole{}
	var classId string = "fabric.FcUplinkPcRole"
	this.ClassId = classId
	var objectType string = "fabric.FcUplinkPcRole"
	this.ObjectType = objectType
	var adminSpeed string = "16Gbps"
	this.AdminSpeed = &adminSpeed
	var fillPattern string = "Idle"
	this.FillPattern = &fillPattern
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcUplinkPcRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcUplinkPcRole) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.FcUplinkPcRole" of the ClassId field.
func (o *FabricFcUplinkPcRole) GetDefaultClassId() interface{} {
	return "fabric.FcUplinkPcRole"
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcUplinkPcRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcUplinkPcRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.FcUplinkPcRole" of the ObjectType field.
func (o *FabricFcUplinkPcRole) GetDefaultObjectType() interface{} {
	return "fabric.FcUplinkPcRole"
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRole) GetAdminSpeed() string {
	if o == nil || IsNil(o.AdminSpeed) {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.AdminSpeed) {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRole) HasAdminSpeed() bool {
	if o != nil && !IsNil(o.AdminSpeed) {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricFcUplinkPcRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFillPattern returns the FillPattern field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRole) GetFillPattern() string {
	if o == nil || IsNil(o.FillPattern) {
		var ret string
		return ret
	}
	return *o.FillPattern
}

// GetFillPatternOk returns a tuple with the FillPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRole) GetFillPatternOk() (*string, bool) {
	if o == nil || IsNil(o.FillPattern) {
		return nil, false
	}
	return o.FillPattern, true
}

// HasFillPattern returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRole) HasFillPattern() bool {
	if o != nil && !IsNil(o.FillPattern) {
		return true
	}

	return false
}

// SetFillPattern gets a reference to the given string and assigns it to the FillPattern field.
func (o *FabricFcUplinkPcRole) SetFillPattern(v string) {
	o.FillPattern = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricFcUplinkPcRole) GetVsanId() int64 {
	if o == nil || IsNil(o.VsanId) {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcUplinkPcRole) GetVsanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.VsanId) {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricFcUplinkPcRole) HasVsanId() bool {
	if o != nil && !IsNil(o.VsanId) {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricFcUplinkPcRole) SetVsanId(v int64) {
	o.VsanId = &v
}

func (o FabricFcUplinkPcRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricFcUplinkPcRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortChannelRole, errFabricPortChannelRole := json.Marshal(o.FabricPortChannelRole)
	if errFabricPortChannelRole != nil {
		return map[string]interface{}{}, errFabricPortChannelRole
	}
	errFabricPortChannelRole = json.Unmarshal([]byte(serializedFabricPortChannelRole), &toSerialize)
	if errFabricPortChannelRole != nil {
		return map[string]interface{}{}, errFabricPortChannelRole
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminSpeed) {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if !IsNil(o.FillPattern) {
		toSerialize["FillPattern"] = o.FillPattern
	}
	if !IsNil(o.VsanId) {
		toSerialize["VsanId"] = o.VsanId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricFcUplinkPcRole) UnmarshalJSON(data []byte) (err error) {
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
	type FabricFcUplinkPcRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port. * `16Gbps` - Admin configurable speed 16Gbps. * `8Gbps` - Admin configurable speed 8Gbps. * `32Gbps` - Admin configurable speed 32Gbps. * `Auto` - Admin configurable speed AUTO ( default ).
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Fill pattern to differentiate the configs in NPIV. * `Idle` - Fc Fill Pattern type Idle. * `Arbff` - Fc Fill Pattern type Arbff.
		FillPattern *string `json:"FillPattern,omitempty"`
		// Virtual San Identifier associated to the FC port.
		VsanId *int64 `json:"VsanId,omitempty"`
	}

	varFabricFcUplinkPcRoleWithoutEmbeddedStruct := FabricFcUplinkPcRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricFcUplinkPcRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcUplinkPcRole := _FabricFcUplinkPcRole{}
		varFabricFcUplinkPcRole.ClassId = varFabricFcUplinkPcRoleWithoutEmbeddedStruct.ClassId
		varFabricFcUplinkPcRole.ObjectType = varFabricFcUplinkPcRoleWithoutEmbeddedStruct.ObjectType
		varFabricFcUplinkPcRole.AdminSpeed = varFabricFcUplinkPcRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricFcUplinkPcRole.FillPattern = varFabricFcUplinkPcRoleWithoutEmbeddedStruct.FillPattern
		varFabricFcUplinkPcRole.VsanId = varFabricFcUplinkPcRoleWithoutEmbeddedStruct.VsanId
		*o = FabricFcUplinkPcRole(varFabricFcUplinkPcRole)
	} else {
		return err
	}

	varFabricFcUplinkPcRole := _FabricFcUplinkPcRole{}

	err = json.Unmarshal(data, &varFabricFcUplinkPcRole)
	if err == nil {
		o.FabricPortChannelRole = varFabricFcUplinkPcRole.FabricPortChannelRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "FillPattern")
		delete(additionalProperties, "VsanId")

		// remove fields from embedded structs
		reflectFabricPortChannelRole := reflect.ValueOf(o.FabricPortChannelRole)
		for i := 0; i < reflectFabricPortChannelRole.Type().NumField(); i++ {
			t := reflectFabricPortChannelRole.Type().Field(i)

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

type NullableFabricFcUplinkPcRole struct {
	value *FabricFcUplinkPcRole
	isSet bool
}

func (v NullableFabricFcUplinkPcRole) Get() *FabricFcUplinkPcRole {
	return v.value
}

func (v *NullableFabricFcUplinkPcRole) Set(val *FabricFcUplinkPcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcUplinkPcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcUplinkPcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcUplinkPcRole(val *FabricFcUplinkPcRole) *NullableFabricFcUplinkPcRole {
	return &NullableFabricFcUplinkPcRole{value: val, isSet: true}
}

func (v NullableFabricFcUplinkPcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcUplinkPcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
