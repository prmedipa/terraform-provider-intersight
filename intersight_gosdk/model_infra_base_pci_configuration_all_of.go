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
)

// InfraBasePciConfigurationAllOf Definition of the list of properties defined in 'infra.BasePciConfiguration', excluding properties defined in parent classes.
type InfraBasePciConfigurationAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The device Id of the GPU device.
	DeviceId *int64 `json:"DeviceId,omitempty"`
	// The vendor Id of the GPU device.
	VendorId             *int64 `json:"VendorId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraBasePciConfigurationAllOf InfraBasePciConfigurationAllOf

// NewInfraBasePciConfigurationAllOf instantiates a new InfraBasePciConfigurationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraBasePciConfigurationAllOf(classId string, objectType string) *InfraBasePciConfigurationAllOf {
	this := InfraBasePciConfigurationAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewInfraBasePciConfigurationAllOfWithDefaults instantiates a new InfraBasePciConfigurationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraBasePciConfigurationAllOfWithDefaults() *InfraBasePciConfigurationAllOf {
	this := InfraBasePciConfigurationAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *InfraBasePciConfigurationAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InfraBasePciConfigurationAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InfraBasePciConfigurationAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *InfraBasePciConfigurationAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *InfraBasePciConfigurationAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *InfraBasePciConfigurationAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InfraBasePciConfigurationAllOf) GetDeviceId() int64 {
	if o == nil || o.DeviceId == nil {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraBasePciConfigurationAllOf) GetDeviceIdOk() (*int64, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InfraBasePciConfigurationAllOf) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *InfraBasePciConfigurationAllOf) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *InfraBasePciConfigurationAllOf) GetVendorId() int64 {
	if o == nil || o.VendorId == nil {
		var ret int64
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraBasePciConfigurationAllOf) GetVendorIdOk() (*int64, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *InfraBasePciConfigurationAllOf) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given int64 and assigns it to the VendorId field.
func (o *InfraBasePciConfigurationAllOf) SetVendorId(v int64) {
	o.VendorId = &v
}

func (o InfraBasePciConfigurationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InfraBasePciConfigurationAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInfraBasePciConfigurationAllOf := _InfraBasePciConfigurationAllOf{}

	if err = json.Unmarshal(bytes, &varInfraBasePciConfigurationAllOf); err == nil {
		*o = InfraBasePciConfigurationAllOf(varInfraBasePciConfigurationAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "VendorId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfraBasePciConfigurationAllOf struct {
	value *InfraBasePciConfigurationAllOf
	isSet bool
}

func (v NullableInfraBasePciConfigurationAllOf) Get() *InfraBasePciConfigurationAllOf {
	return v.value
}

func (v *NullableInfraBasePciConfigurationAllOf) Set(val *InfraBasePciConfigurationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraBasePciConfigurationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraBasePciConfigurationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraBasePciConfigurationAllOf(val *InfraBasePciConfigurationAllOf) *NullableInfraBasePciConfigurationAllOf {
	return &NullableInfraBasePciConfigurationAllOf{value: val, isSet: true}
}

func (v NullableInfraBasePciConfigurationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraBasePciConfigurationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
