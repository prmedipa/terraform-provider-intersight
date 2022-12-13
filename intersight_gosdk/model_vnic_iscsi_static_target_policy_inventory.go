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

// VnicIscsiStaticTargetPolicyInventory Configuration parameters that defines the reachability of iSCSI Target portal.
type VnicIscsiStaticTargetPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IPv4 address assigned to the iSCSI target.
	IpAddress *string         `json:"IpAddress,omitempty"`
	Lun       NullableVnicLun `json:"Lun,omitempty"`
	// The port associated with the iSCSI target.
	Port *int64 `json:"Port,omitempty"`
	// Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target.
	TargetName           *string               `json:"TargetName,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicIscsiStaticTargetPolicyInventory VnicIscsiStaticTargetPolicyInventory

// NewVnicIscsiStaticTargetPolicyInventory instantiates a new VnicIscsiStaticTargetPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicIscsiStaticTargetPolicyInventory(classId string, objectType string) *VnicIscsiStaticTargetPolicyInventory {
	this := VnicIscsiStaticTargetPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicIscsiStaticTargetPolicyInventoryWithDefaults instantiates a new VnicIscsiStaticTargetPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicIscsiStaticTargetPolicyInventoryWithDefaults() *VnicIscsiStaticTargetPolicyInventory {
	this := VnicIscsiStaticTargetPolicyInventory{}
	var classId string = "vnic.IscsiStaticTargetPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.IscsiStaticTargetPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicIscsiStaticTargetPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicIscsiStaticTargetPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicIscsiStaticTargetPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicIscsiStaticTargetPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *VnicIscsiStaticTargetPolicyInventory) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *VnicIscsiStaticTargetPolicyInventory) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLun returns the Lun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicIscsiStaticTargetPolicyInventory) GetLun() VnicLun {
	if o == nil || o.Lun.Get() == nil {
		var ret VnicLun
		return ret
	}
	return *o.Lun.Get()
}

// GetLunOk returns a tuple with the Lun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicIscsiStaticTargetPolicyInventory) GetLunOk() (*VnicLun, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lun.Get(), o.Lun.IsSet()
}

// HasLun returns a boolean if a field has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) HasLun() bool {
	if o != nil && o.Lun.IsSet() {
		return true
	}

	return false
}

// SetLun gets a reference to the given NullableVnicLun and assigns it to the Lun field.
func (o *VnicIscsiStaticTargetPolicyInventory) SetLun(v VnicLun) {
	o.Lun.Set(&v)
}

// SetLunNil sets the value for Lun to be an explicit nil
func (o *VnicIscsiStaticTargetPolicyInventory) SetLunNil() {
	o.Lun.Set(nil)
}

// UnsetLun ensures that no value is present for Lun, not even an explicit nil
func (o *VnicIscsiStaticTargetPolicyInventory) UnsetLun() {
	o.Lun.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VnicIscsiStaticTargetPolicyInventory) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *VnicIscsiStaticTargetPolicyInventory) SetPort(v int64) {
	o.Port = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetName() string {
	if o == nil || o.TargetName == nil {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetNameOk() (*string, bool) {
	if o == nil || o.TargetName == nil {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) HasTargetName() bool {
	if o != nil && o.TargetName != nil {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *VnicIscsiStaticTargetPolicyInventory) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicIscsiStaticTargetPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicIscsiStaticTargetPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicIscsiStaticTargetPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return []byte{}, errPolicyAbstractPolicyInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Lun.IsSet() {
		toSerialize["Lun"] = o.Lun.Get()
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.TargetName != nil {
		toSerialize["TargetName"] = o.TargetName
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicIscsiStaticTargetPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The IPv4 address assigned to the iSCSI target.
		IpAddress *string         `json:"IpAddress,omitempty"`
		Lun       NullableVnicLun `json:"Lun,omitempty"`
		// The port associated with the iSCSI target.
		Port *int64 `json:"Port,omitempty"`
		// Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target.
		TargetName *string               `json:"TargetName,omitempty"`
		TargetMo   *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct := VnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicIscsiStaticTargetPolicyInventory := _VnicIscsiStaticTargetPolicyInventory{}
		varVnicIscsiStaticTargetPolicyInventory.ClassId = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicIscsiStaticTargetPolicyInventory.ObjectType = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicIscsiStaticTargetPolicyInventory.IpAddress = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.IpAddress
		varVnicIscsiStaticTargetPolicyInventory.Lun = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.Lun
		varVnicIscsiStaticTargetPolicyInventory.Port = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.Port
		varVnicIscsiStaticTargetPolicyInventory.TargetName = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.TargetName
		varVnicIscsiStaticTargetPolicyInventory.TargetMo = varVnicIscsiStaticTargetPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicIscsiStaticTargetPolicyInventory(varVnicIscsiStaticTargetPolicyInventory)
	} else {
		return err
	}

	varVnicIscsiStaticTargetPolicyInventory := _VnicIscsiStaticTargetPolicyInventory{}

	err = json.Unmarshal(bytes, &varVnicIscsiStaticTargetPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varVnicIscsiStaticTargetPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Lun")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "TargetName")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableVnicIscsiStaticTargetPolicyInventory struct {
	value *VnicIscsiStaticTargetPolicyInventory
	isSet bool
}

func (v NullableVnicIscsiStaticTargetPolicyInventory) Get() *VnicIscsiStaticTargetPolicyInventory {
	return v.value
}

func (v *NullableVnicIscsiStaticTargetPolicyInventory) Set(val *VnicIscsiStaticTargetPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicIscsiStaticTargetPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicIscsiStaticTargetPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicIscsiStaticTargetPolicyInventory(val *VnicIscsiStaticTargetPolicyInventory) *NullableVnicIscsiStaticTargetPolicyInventory {
	return &NullableVnicIscsiStaticTargetPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicIscsiStaticTargetPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicIscsiStaticTargetPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
