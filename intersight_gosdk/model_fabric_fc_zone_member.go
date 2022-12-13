/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricFcZoneMember SAN target or WWPN that is to be part of the FC zone.
type FabricFcZoneMember struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name given to the target member.
	Name *string `json:"Name,omitempty"`
	// Unique identifier for the Fabric object. * `A` - Switch Identifier of Fabric Interconnect A. * `B` - Switch Identifier of Fabric Interconnect B.
	SwitchId *string `json:"SwitchId,omitempty"`
	// VSAN with scope defined as Storage in the VSAN policy.
	VsanId *int64 `json:"VsanId,omitempty"`
	// WWPN that is a member of the FC zone.
	Wwpn                 *string `json:"Wwpn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcZoneMember FabricFcZoneMember

// NewFabricFcZoneMember instantiates a new FabricFcZoneMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcZoneMember(classId string, objectType string) *FabricFcZoneMember {
	this := FabricFcZoneMember{}
	this.ClassId = classId
	this.ObjectType = objectType
	var switchId string = "A"
	this.SwitchId = &switchId
	return &this
}

// NewFabricFcZoneMemberWithDefaults instantiates a new FabricFcZoneMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcZoneMemberWithDefaults() *FabricFcZoneMember {
	this := FabricFcZoneMember{}
	var classId string = "fabric.FcZoneMember"
	this.ClassId = classId
	var objectType string = "fabric.FcZoneMember"
	this.ObjectType = objectType
	var switchId string = "A"
	this.SwitchId = &switchId
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcZoneMember) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcZoneMember) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcZoneMember) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcZoneMember) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FabricFcZoneMember) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FabricFcZoneMember) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FabricFcZoneMember) SetName(v string) {
	o.Name = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *FabricFcZoneMember) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *FabricFcZoneMember) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *FabricFcZoneMember) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetVsanId returns the VsanId field value if set, zero value otherwise.
func (o *FabricFcZoneMember) GetVsanId() int64 {
	if o == nil || o.VsanId == nil {
		var ret int64
		return ret
	}
	return *o.VsanId
}

// GetVsanIdOk returns a tuple with the VsanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetVsanIdOk() (*int64, bool) {
	if o == nil || o.VsanId == nil {
		return nil, false
	}
	return o.VsanId, true
}

// HasVsanId returns a boolean if a field has been set.
func (o *FabricFcZoneMember) HasVsanId() bool {
	if o != nil && o.VsanId != nil {
		return true
	}

	return false
}

// SetVsanId gets a reference to the given int64 and assigns it to the VsanId field.
func (o *FabricFcZoneMember) SetVsanId(v int64) {
	o.VsanId = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *FabricFcZoneMember) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcZoneMember) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *FabricFcZoneMember) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *FabricFcZoneMember) SetWwpn(v string) {
	o.Wwpn = &v
}

func (o FabricFcZoneMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.VsanId != nil {
		toSerialize["VsanId"] = o.VsanId
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcZoneMember) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcZoneMemberWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name given to the target member.
		Name *string `json:"Name,omitempty"`
		// Unique identifier for the Fabric object. * `A` - Switch Identifier of Fabric Interconnect A. * `B` - Switch Identifier of Fabric Interconnect B.
		SwitchId *string `json:"SwitchId,omitempty"`
		// VSAN with scope defined as Storage in the VSAN policy.
		VsanId *int64 `json:"VsanId,omitempty"`
		// WWPN that is a member of the FC zone.
		Wwpn *string `json:"Wwpn,omitempty"`
	}

	varFabricFcZoneMemberWithoutEmbeddedStruct := FabricFcZoneMemberWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcZoneMemberWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcZoneMember := _FabricFcZoneMember{}
		varFabricFcZoneMember.ClassId = varFabricFcZoneMemberWithoutEmbeddedStruct.ClassId
		varFabricFcZoneMember.ObjectType = varFabricFcZoneMemberWithoutEmbeddedStruct.ObjectType
		varFabricFcZoneMember.Name = varFabricFcZoneMemberWithoutEmbeddedStruct.Name
		varFabricFcZoneMember.SwitchId = varFabricFcZoneMemberWithoutEmbeddedStruct.SwitchId
		varFabricFcZoneMember.VsanId = varFabricFcZoneMemberWithoutEmbeddedStruct.VsanId
		varFabricFcZoneMember.Wwpn = varFabricFcZoneMemberWithoutEmbeddedStruct.Wwpn
		*o = FabricFcZoneMember(varFabricFcZoneMember)
	} else {
		return err
	}

	varFabricFcZoneMember := _FabricFcZoneMember{}

	err = json.Unmarshal(bytes, &varFabricFcZoneMember)
	if err == nil {
		o.MoBaseComplexType = varFabricFcZoneMember.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "SwitchId")
		delete(additionalProperties, "VsanId")
		delete(additionalProperties, "Wwpn")

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

type NullableFabricFcZoneMember struct {
	value *FabricFcZoneMember
	isSet bool
}

func (v NullableFabricFcZoneMember) Get() *FabricFcZoneMember {
	return v.value
}

func (v *NullableFabricFcZoneMember) Set(val *FabricFcZoneMember) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcZoneMember) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcZoneMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcZoneMember(val *FabricFcZoneMember) *NullableFabricFcZoneMember {
	return &NullableFabricFcZoneMember{value: val, isSet: true}
}

func (v NullableFabricFcZoneMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcZoneMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
