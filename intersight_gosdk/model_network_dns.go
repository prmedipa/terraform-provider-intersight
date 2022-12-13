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

// NetworkDns Concrete class for list of DNS servers configured on the nexus end point.
type NetworkDns struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string   `json:"ObjectType"`
	AdditionalDomains []string `json:"AdditionalDomains,omitempty"`
	// Default domain configured for VRF.
	DefaultDomain *string  `json:"DefaultDomain,omitempty"`
	NameServers   []string `json:"NameServers,omitempty"`
	// Name of the VRF configured for the DNS.
	VrfName              *string                              `json:"VrfName,omitempty"`
	NetworkElement       *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDns NetworkDns

// NewNetworkDns instantiates a new NetworkDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDns(classId string, objectType string) *NetworkDns {
	this := NetworkDns{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkDnsWithDefaults instantiates a new NetworkDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDnsWithDefaults() *NetworkDns {
	this := NetworkDns{}
	var classId string = "network.Dns"
	this.ClassId = classId
	var objectType string = "network.Dns"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkDns) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkDns) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkDns) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkDns) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdditionalDomains returns the AdditionalDomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkDns) GetAdditionalDomains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdditionalDomains
}

// GetAdditionalDomainsOk returns a tuple with the AdditionalDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkDns) GetAdditionalDomainsOk() ([]string, bool) {
	if o == nil || o.AdditionalDomains == nil {
		return nil, false
	}
	return o.AdditionalDomains, true
}

// HasAdditionalDomains returns a boolean if a field has been set.
func (o *NetworkDns) HasAdditionalDomains() bool {
	if o != nil && o.AdditionalDomains != nil {
		return true
	}

	return false
}

// SetAdditionalDomains gets a reference to the given []string and assigns it to the AdditionalDomains field.
func (o *NetworkDns) SetAdditionalDomains(v []string) {
	o.AdditionalDomains = v
}

// GetDefaultDomain returns the DefaultDomain field value if set, zero value otherwise.
func (o *NetworkDns) GetDefaultDomain() string {
	if o == nil || o.DefaultDomain == nil {
		var ret string
		return ret
	}
	return *o.DefaultDomain
}

// GetDefaultDomainOk returns a tuple with the DefaultDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetDefaultDomainOk() (*string, bool) {
	if o == nil || o.DefaultDomain == nil {
		return nil, false
	}
	return o.DefaultDomain, true
}

// HasDefaultDomain returns a boolean if a field has been set.
func (o *NetworkDns) HasDefaultDomain() bool {
	if o != nil && o.DefaultDomain != nil {
		return true
	}

	return false
}

// SetDefaultDomain gets a reference to the given string and assigns it to the DefaultDomain field.
func (o *NetworkDns) SetDefaultDomain(v string) {
	o.DefaultDomain = &v
}

// GetNameServers returns the NameServers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkDns) GetNameServers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NameServers
}

// GetNameServersOk returns a tuple with the NameServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkDns) GetNameServersOk() ([]string, bool) {
	if o == nil || o.NameServers == nil {
		return nil, false
	}
	return o.NameServers, true
}

// HasNameServers returns a boolean if a field has been set.
func (o *NetworkDns) HasNameServers() bool {
	if o != nil && o.NameServers != nil {
		return true
	}

	return false
}

// SetNameServers gets a reference to the given []string and assigns it to the NameServers field.
func (o *NetworkDns) SetNameServers(v []string) {
	o.NameServers = v
}

// GetVrfName returns the VrfName field value if set, zero value otherwise.
func (o *NetworkDns) GetVrfName() string {
	if o == nil || o.VrfName == nil {
		var ret string
		return ret
	}
	return *o.VrfName
}

// GetVrfNameOk returns a tuple with the VrfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetVrfNameOk() (*string, bool) {
	if o == nil || o.VrfName == nil {
		return nil, false
	}
	return o.VrfName, true
}

// HasVrfName returns a boolean if a field has been set.
func (o *NetworkDns) HasVrfName() bool {
	if o != nil && o.VrfName != nil {
		return true
	}

	return false
}

// SetVrfName gets a reference to the given string and assigns it to the VrfName field.
func (o *NetworkDns) SetVrfName(v string) {
	o.VrfName = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *NetworkDns) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *NetworkDns) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *NetworkDns) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkDns) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDns) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkDns) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkDns) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NetworkDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdditionalDomains != nil {
		toSerialize["AdditionalDomains"] = o.AdditionalDomains
	}
	if o.DefaultDomain != nil {
		toSerialize["DefaultDomain"] = o.DefaultDomain
	}
	if o.NameServers != nil {
		toSerialize["NameServers"] = o.NameServers
	}
	if o.VrfName != nil {
		toSerialize["VrfName"] = o.VrfName
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

func (o *NetworkDns) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkDnsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string   `json:"ObjectType"`
		AdditionalDomains []string `json:"AdditionalDomains,omitempty"`
		// Default domain configured for VRF.
		DefaultDomain *string  `json:"DefaultDomain,omitempty"`
		NameServers   []string `json:"NameServers,omitempty"`
		// Name of the VRF configured for the DNS.
		VrfName          *string                              `json:"VrfName,omitempty"`
		NetworkElement   *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNetworkDnsWithoutEmbeddedStruct := NetworkDnsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkDnsWithoutEmbeddedStruct)
	if err == nil {
		varNetworkDns := _NetworkDns{}
		varNetworkDns.ClassId = varNetworkDnsWithoutEmbeddedStruct.ClassId
		varNetworkDns.ObjectType = varNetworkDnsWithoutEmbeddedStruct.ObjectType
		varNetworkDns.AdditionalDomains = varNetworkDnsWithoutEmbeddedStruct.AdditionalDomains
		varNetworkDns.DefaultDomain = varNetworkDnsWithoutEmbeddedStruct.DefaultDomain
		varNetworkDns.NameServers = varNetworkDnsWithoutEmbeddedStruct.NameServers
		varNetworkDns.VrfName = varNetworkDnsWithoutEmbeddedStruct.VrfName
		varNetworkDns.NetworkElement = varNetworkDnsWithoutEmbeddedStruct.NetworkElement
		varNetworkDns.RegisteredDevice = varNetworkDnsWithoutEmbeddedStruct.RegisteredDevice
		*o = NetworkDns(varNetworkDns)
	} else {
		return err
	}

	varNetworkDns := _NetworkDns{}

	err = json.Unmarshal(bytes, &varNetworkDns)
	if err == nil {
		o.InventoryBase = varNetworkDns.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdditionalDomains")
		delete(additionalProperties, "DefaultDomain")
		delete(additionalProperties, "NameServers")
		delete(additionalProperties, "VrfName")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableNetworkDns struct {
	value *NetworkDns
	isSet bool
}

func (v NullableNetworkDns) Get() *NetworkDns {
	return v.value
}

func (v *NullableNetworkDns) Set(val *NetworkDns) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDns) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDns(val *NetworkDns) *NullableNetworkDns {
	return &NullableNetworkDns{value: val, isSet: true}
}

func (v NullableNetworkDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
