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
)

// CapabilityEndpointDescriptorAllOf Definition of the list of properties defined in 'capability.EndpointDescriptor', excluding properties defined in parent classes.
type CapabilityEndpointDescriptorAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Detailed information about the endpoint.
	Description *string `json:"Description,omitempty"`
	// The model of the endpoint, for which this capability information is applicable.
	Model *string `json:"Model,omitempty"`
	// The vendor of the endpoint, for which this capability information is applicable.
	Vendor *string `json:"Vendor,omitempty"`
	// The firmware or software version of the endpoint, for which this capability information is applicable.
	Version *string `json:"Version,omitempty"`
	// An array of relationships to capabilityCapability resources.
	Capabilities         []CapabilityCapabilityRelationship `json:"Capabilities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEndpointDescriptorAllOf CapabilityEndpointDescriptorAllOf

// NewCapabilityEndpointDescriptorAllOf instantiates a new CapabilityEndpointDescriptorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEndpointDescriptorAllOf(classId string, objectType string) *CapabilityEndpointDescriptorAllOf {
	this := CapabilityEndpointDescriptorAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityEndpointDescriptorAllOfWithDefaults instantiates a new CapabilityEndpointDescriptorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEndpointDescriptorAllOfWithDefaults() *CapabilityEndpointDescriptorAllOf {
	this := CapabilityEndpointDescriptorAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityEndpointDescriptorAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityEndpointDescriptorAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityEndpointDescriptorAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityEndpointDescriptorAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CapabilityEndpointDescriptorAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CapabilityEndpointDescriptorAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CapabilityEndpointDescriptorAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CapabilityEndpointDescriptorAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CapabilityEndpointDescriptorAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CapabilityEndpointDescriptorAllOf) SetModel(v string) {
	o.Model = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CapabilityEndpointDescriptorAllOf) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CapabilityEndpointDescriptorAllOf) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CapabilityEndpointDescriptorAllOf) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CapabilityEndpointDescriptorAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEndpointDescriptorAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CapabilityEndpointDescriptorAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CapabilityEndpointDescriptorAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CapabilityEndpointDescriptorAllOf) GetCapabilities() []CapabilityCapabilityRelationship {
	if o == nil {
		var ret []CapabilityCapabilityRelationship
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CapabilityEndpointDescriptorAllOf) GetCapabilitiesOk() ([]CapabilityCapabilityRelationship, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *CapabilityEndpointDescriptorAllOf) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []CapabilityCapabilityRelationship and assigns it to the Capabilities field.
func (o *CapabilityEndpointDescriptorAllOf) SetCapabilities(v []CapabilityCapabilityRelationship) {
	o.Capabilities = v
}

func (o CapabilityEndpointDescriptorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Capabilities != nil {
		toSerialize["Capabilities"] = o.Capabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEndpointDescriptorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityEndpointDescriptorAllOf := _CapabilityEndpointDescriptorAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityEndpointDescriptorAllOf); err == nil {
		*o = CapabilityEndpointDescriptorAllOf(varCapabilityEndpointDescriptorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Capabilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityEndpointDescriptorAllOf struct {
	value *CapabilityEndpointDescriptorAllOf
	isSet bool
}

func (v NullableCapabilityEndpointDescriptorAllOf) Get() *CapabilityEndpointDescriptorAllOf {
	return v.value
}

func (v *NullableCapabilityEndpointDescriptorAllOf) Set(val *CapabilityEndpointDescriptorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEndpointDescriptorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEndpointDescriptorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEndpointDescriptorAllOf(val *CapabilityEndpointDescriptorAllOf) *NullableCapabilityEndpointDescriptorAllOf {
	return &NullableCapabilityEndpointDescriptorAllOf{value: val, isSet: true}
}

func (v NullableCapabilityEndpointDescriptorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEndpointDescriptorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
