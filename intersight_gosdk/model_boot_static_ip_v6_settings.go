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

// checks if the BootStaticIpV6Settings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BootStaticIpV6Settings{}

// BootStaticIpV6Settings Lists the properties that can be used for static IPv6 configuration in HTTP boot.
type BootStaticIpV6Settings struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of DNS server.
	DnsIp *string `json:"DnsIp,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	// IP address of default gateway.
	GatewayIp *string `json:"GatewayIp,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	// Ipv6 static Internet Protocol address.
	Ip *string `json:"Ip,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
	// A prefix length which masks the IP address and divides the IP address into network address and host address.
	PrefixLength         *int64 `json:"PrefixLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootStaticIpV6Settings BootStaticIpV6Settings

// NewBootStaticIpV6Settings instantiates a new BootStaticIpV6Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootStaticIpV6Settings(classId string, objectType string) *BootStaticIpV6Settings {
	this := BootStaticIpV6Settings{}
	this.ClassId = classId
	this.ObjectType = objectType
	var prefixLength int64 = 1
	this.PrefixLength = &prefixLength
	return &this
}

// NewBootStaticIpV6SettingsWithDefaults instantiates a new BootStaticIpV6Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootStaticIpV6SettingsWithDefaults() *BootStaticIpV6Settings {
	this := BootStaticIpV6Settings{}
	var classId string = "boot.StaticIpV6Settings"
	this.ClassId = classId
	var objectType string = "boot.StaticIpV6Settings"
	this.ObjectType = objectType
	var prefixLength int64 = 1
	this.PrefixLength = &prefixLength
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootStaticIpV6Settings) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootStaticIpV6Settings) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "boot.StaticIpV6Settings" of the ClassId field.
func (o *BootStaticIpV6Settings) GetDefaultClassId() interface{} {
	return "boot.StaticIpV6Settings"
}

// GetObjectType returns the ObjectType field value
func (o *BootStaticIpV6Settings) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootStaticIpV6Settings) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "boot.StaticIpV6Settings" of the ObjectType field.
func (o *BootStaticIpV6Settings) GetDefaultObjectType() interface{} {
	return "boot.StaticIpV6Settings"
}

// GetDnsIp returns the DnsIp field value if set, zero value otherwise.
func (o *BootStaticIpV6Settings) GetDnsIp() string {
	if o == nil || IsNil(o.DnsIp) {
		var ret string
		return ret
	}
	return *o.DnsIp
}

// GetDnsIpOk returns a tuple with the DnsIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetDnsIpOk() (*string, bool) {
	if o == nil || IsNil(o.DnsIp) {
		return nil, false
	}
	return o.DnsIp, true
}

// HasDnsIp returns a boolean if a field has been set.
func (o *BootStaticIpV6Settings) HasDnsIp() bool {
	if o != nil && !IsNil(o.DnsIp) {
		return true
	}

	return false
}

// SetDnsIp gets a reference to the given string and assigns it to the DnsIp field.
func (o *BootStaticIpV6Settings) SetDnsIp(v string) {
	o.DnsIp = &v
}

// GetGatewayIp returns the GatewayIp field value if set, zero value otherwise.
func (o *BootStaticIpV6Settings) GetGatewayIp() string {
	if o == nil || IsNil(o.GatewayIp) {
		var ret string
		return ret
	}
	return *o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetGatewayIpOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayIp) {
		return nil, false
	}
	return o.GatewayIp, true
}

// HasGatewayIp returns a boolean if a field has been set.
func (o *BootStaticIpV6Settings) HasGatewayIp() bool {
	if o != nil && !IsNil(o.GatewayIp) {
		return true
	}

	return false
}

// SetGatewayIp gets a reference to the given string and assigns it to the GatewayIp field.
func (o *BootStaticIpV6Settings) SetGatewayIp(v string) {
	o.GatewayIp = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *BootStaticIpV6Settings) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *BootStaticIpV6Settings) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *BootStaticIpV6Settings) SetIp(v string) {
	o.Ip = &v
}

// GetPrefixLength returns the PrefixLength field value if set, zero value otherwise.
func (o *BootStaticIpV6Settings) GetPrefixLength() int64 {
	if o == nil || IsNil(o.PrefixLength) {
		var ret int64
		return ret
	}
	return *o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootStaticIpV6Settings) GetPrefixLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixLength) {
		return nil, false
	}
	return o.PrefixLength, true
}

// HasPrefixLength returns a boolean if a field has been set.
func (o *BootStaticIpV6Settings) HasPrefixLength() bool {
	if o != nil && !IsNil(o.PrefixLength) {
		return true
	}

	return false
}

// SetPrefixLength gets a reference to the given int64 and assigns it to the PrefixLength field.
func (o *BootStaticIpV6Settings) SetPrefixLength(v int64) {
	o.PrefixLength = &v
}

func (o BootStaticIpV6Settings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BootStaticIpV6Settings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DnsIp) {
		toSerialize["DnsIp"] = o.DnsIp
	}
	if !IsNil(o.GatewayIp) {
		toSerialize["GatewayIp"] = o.GatewayIp
	}
	if !IsNil(o.Ip) {
		toSerialize["Ip"] = o.Ip
	}
	if !IsNil(o.PrefixLength) {
		toSerialize["PrefixLength"] = o.PrefixLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BootStaticIpV6Settings) UnmarshalJSON(data []byte) (err error) {
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
	type BootStaticIpV6SettingsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of DNS server.
		DnsIp *string `json:"DnsIp,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		// IP address of default gateway.
		GatewayIp *string `json:"GatewayIp,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		// Ipv6 static Internet Protocol address.
		Ip *string `json:"Ip,omitempty" validate:"regexp=^$|^(([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:([0-9A-Fa-f]{1,4}:[0-9A-Fa-f]{0,4}|:[0-9A-Fa-f]{1,4})?|(:[0-9A-Fa-f]{1,4}){0,2})|(:[0-9A-Fa-f]{1,4}){0,3})|(:[0-9A-Fa-f]{1,4}){0,4})|:(:[0-9A-Fa-f]{1,4}){0,5})((:[0-9A-Fa-f]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\\\\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9A-Fa-f]{1,4}:){1,6}|:):[0-9A-Fa-f]{0,4}|([0-9A-Fa-f]{1,4}:){7}:)$"`
		// A prefix length which masks the IP address and divides the IP address into network address and host address.
		PrefixLength *int64 `json:"PrefixLength,omitempty"`
	}

	varBootStaticIpV6SettingsWithoutEmbeddedStruct := BootStaticIpV6SettingsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBootStaticIpV6SettingsWithoutEmbeddedStruct)
	if err == nil {
		varBootStaticIpV6Settings := _BootStaticIpV6Settings{}
		varBootStaticIpV6Settings.ClassId = varBootStaticIpV6SettingsWithoutEmbeddedStruct.ClassId
		varBootStaticIpV6Settings.ObjectType = varBootStaticIpV6SettingsWithoutEmbeddedStruct.ObjectType
		varBootStaticIpV6Settings.DnsIp = varBootStaticIpV6SettingsWithoutEmbeddedStruct.DnsIp
		varBootStaticIpV6Settings.GatewayIp = varBootStaticIpV6SettingsWithoutEmbeddedStruct.GatewayIp
		varBootStaticIpV6Settings.Ip = varBootStaticIpV6SettingsWithoutEmbeddedStruct.Ip
		varBootStaticIpV6Settings.PrefixLength = varBootStaticIpV6SettingsWithoutEmbeddedStruct.PrefixLength
		*o = BootStaticIpV6Settings(varBootStaticIpV6Settings)
	} else {
		return err
	}

	varBootStaticIpV6Settings := _BootStaticIpV6Settings{}

	err = json.Unmarshal(data, &varBootStaticIpV6Settings)
	if err == nil {
		o.MoBaseComplexType = varBootStaticIpV6Settings.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DnsIp")
		delete(additionalProperties, "GatewayIp")
		delete(additionalProperties, "Ip")
		delete(additionalProperties, "PrefixLength")

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

type NullableBootStaticIpV6Settings struct {
	value *BootStaticIpV6Settings
	isSet bool
}

func (v NullableBootStaticIpV6Settings) Get() *BootStaticIpV6Settings {
	return v.value
}

func (v *NullableBootStaticIpV6Settings) Set(val *BootStaticIpV6Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableBootStaticIpV6Settings) IsSet() bool {
	return v.isSet
}

func (v *NullableBootStaticIpV6Settings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootStaticIpV6Settings(val *BootStaticIpV6Settings) *NullableBootStaticIpV6Settings {
	return &NullableBootStaticIpV6Settings{value: val, isSet: true}
}

func (v NullableBootStaticIpV6Settings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootStaticIpV6Settings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
