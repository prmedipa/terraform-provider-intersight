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

// checks if the IamIpAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamIpAddress{}

// IamIpAddress Add an IP address to enable IP address based access management.
type IamIpAddress struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example '12.13.14.15', '12.13.14.0/24', and '12.13.14.15-12.13.14.200'. Reserved IP ranges '127.0.0.1', '10.0.0.0/8', '172.16.0.0/12', and '192.168.0.0/16' are not allowed.
	Address *string `json:"Address,omitempty"`
	// Description of Trusted IP address range.
	Description          *string                                   `json:"Description,omitempty"`
	IpAccessManagement   NullableIamIpAccessManagementRelationship `json:"IpAccessManagement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamIpAddress IamIpAddress

// NewIamIpAddress instantiates a new IamIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamIpAddress(classId string, objectType string) *IamIpAddress {
	this := IamIpAddress{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamIpAddressWithDefaults instantiates a new IamIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamIpAddressWithDefaults() *IamIpAddress {
	this := IamIpAddress{}
	var classId string = "iam.IpAddress"
	this.ClassId = classId
	var objectType string = "iam.IpAddress"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamIpAddress) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamIpAddress) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamIpAddress) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.IpAddress" of the ClassId field.
func (o *IamIpAddress) GetDefaultClassId() interface{} {
	return "iam.IpAddress"
}

// GetObjectType returns the ObjectType field value
func (o *IamIpAddress) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamIpAddress) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamIpAddress) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.IpAddress" of the ObjectType field.
func (o *IamIpAddress) GetDefaultObjectType() interface{} {
	return "iam.IpAddress"
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IamIpAddress) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAddress) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IamIpAddress) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *IamIpAddress) SetAddress(v string) {
	o.Address = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamIpAddress) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamIpAddress) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamIpAddress) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamIpAddress) SetDescription(v string) {
	o.Description = &v
}

// GetIpAccessManagement returns the IpAccessManagement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamIpAddress) GetIpAccessManagement() IamIpAccessManagementRelationship {
	if o == nil || IsNil(o.IpAccessManagement.Get()) {
		var ret IamIpAccessManagementRelationship
		return ret
	}
	return *o.IpAccessManagement.Get()
}

// GetIpAccessManagementOk returns a tuple with the IpAccessManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamIpAddress) GetIpAccessManagementOk() (*IamIpAccessManagementRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAccessManagement.Get(), o.IpAccessManagement.IsSet()
}

// HasIpAccessManagement returns a boolean if a field has been set.
func (o *IamIpAddress) HasIpAccessManagement() bool {
	if o != nil && o.IpAccessManagement.IsSet() {
		return true
	}

	return false
}

// SetIpAccessManagement gets a reference to the given NullableIamIpAccessManagementRelationship and assigns it to the IpAccessManagement field.
func (o *IamIpAddress) SetIpAccessManagement(v IamIpAccessManagementRelationship) {
	o.IpAccessManagement.Set(&v)
}

// SetIpAccessManagementNil sets the value for IpAccessManagement to be an explicit nil
func (o *IamIpAddress) SetIpAccessManagementNil() {
	o.IpAccessManagement.Set(nil)
}

// UnsetIpAccessManagement ensures that no value is present for IpAccessManagement, not even an explicit nil
func (o *IamIpAddress) UnsetIpAccessManagement() {
	o.IpAccessManagement.Unset()
}

func (o IamIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamIpAddress) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Address) {
		toSerialize["Address"] = o.Address
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.IpAccessManagement.IsSet() {
		toSerialize["IpAccessManagement"] = o.IpAccessManagement.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamIpAddress) UnmarshalJSON(data []byte) (err error) {
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
	type IamIpAddressWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example '12.13.14.15', '12.13.14.0/24', and '12.13.14.15-12.13.14.200'. Reserved IP ranges '127.0.0.1', '10.0.0.0/8', '172.16.0.0/12', and '192.168.0.0/16' are not allowed.
		Address *string `json:"Address,omitempty"`
		// Description of Trusted IP address range.
		Description        *string                                   `json:"Description,omitempty"`
		IpAccessManagement NullableIamIpAccessManagementRelationship `json:"IpAccessManagement,omitempty"`
	}

	varIamIpAddressWithoutEmbeddedStruct := IamIpAddressWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamIpAddressWithoutEmbeddedStruct)
	if err == nil {
		varIamIpAddress := _IamIpAddress{}
		varIamIpAddress.ClassId = varIamIpAddressWithoutEmbeddedStruct.ClassId
		varIamIpAddress.ObjectType = varIamIpAddressWithoutEmbeddedStruct.ObjectType
		varIamIpAddress.Address = varIamIpAddressWithoutEmbeddedStruct.Address
		varIamIpAddress.Description = varIamIpAddressWithoutEmbeddedStruct.Description
		varIamIpAddress.IpAccessManagement = varIamIpAddressWithoutEmbeddedStruct.IpAccessManagement
		*o = IamIpAddress(varIamIpAddress)
	} else {
		return err
	}

	varIamIpAddress := _IamIpAddress{}

	err = json.Unmarshal(data, &varIamIpAddress)
	if err == nil {
		o.MoBaseMo = varIamIpAddress.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Address")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "IpAccessManagement")

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

type NullableIamIpAddress struct {
	value *IamIpAddress
	isSet bool
}

func (v NullableIamIpAddress) Get() *IamIpAddress {
	return v.value
}

func (v *NullableIamIpAddress) Set(val *IamIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIamIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIamIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamIpAddress(val *IamIpAddress) *NullableIamIpAddress {
	return &NullableIamIpAddress{value: val, isSet: true}
}

func (v NullableIamIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
