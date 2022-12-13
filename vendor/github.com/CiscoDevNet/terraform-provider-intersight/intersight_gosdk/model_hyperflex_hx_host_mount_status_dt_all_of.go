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

// HyperflexHxHostMountStatusDtAllOf Definition of the list of properties defined in 'hyperflex.HxHostMountStatusDt', excluding properties defined in parent classes.
type HyperflexHxHostMountStatusDtAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Accessibility of datastore. * `ACCESSIBLE` - The HyperFlex datastore accessibility state is accessible. * `NOT_ACCESSIBLE` - The HyperFlex datastore accessibility state is not accessible. * `PARTIALLY_ACCESSIBLE` - The HyperFlex datastore accessibility state is partially accessible. * `READONLY` - The HyperFlex datastore accessibility state is read-only. * `HOST_NOT_REACHABLE` - The HyperFlex datastore accessibility state is host not reachable. * `UNKNOWN` - The HyperFlex datastore accessibility state is unknown.
	Accessibility *string `json:"Accessibility,omitempty"`
	// Name of host for the HyperFlex datastore.
	HostName *string `json:"HostName,omitempty"`
	// Is the datastore mounted or not.
	Mounted *bool `json:"Mounted,omitempty"`
	// Reason for inaccessibility.
	Reason               *string `json:"Reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexHxHostMountStatusDtAllOf HyperflexHxHostMountStatusDtAllOf

// NewHyperflexHxHostMountStatusDtAllOf instantiates a new HyperflexHxHostMountStatusDtAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxHostMountStatusDtAllOf(classId string, objectType string) *HyperflexHxHostMountStatusDtAllOf {
	this := HyperflexHxHostMountStatusDtAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexHxHostMountStatusDtAllOfWithDefaults instantiates a new HyperflexHxHostMountStatusDtAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxHostMountStatusDtAllOfWithDefaults() *HyperflexHxHostMountStatusDtAllOf {
	this := HyperflexHxHostMountStatusDtAllOf{}
	var classId string = "hyperflex.HxHostMountStatusDt"
	this.ClassId = classId
	var objectType string = "hyperflex.HxHostMountStatusDt"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexHxHostMountStatusDtAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexHxHostMountStatusDtAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexHxHostMountStatusDtAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexHxHostMountStatusDtAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *HyperflexHxHostMountStatusDtAllOf) GetAccessibility() string {
	if o == nil || o.Accessibility == nil {
		var ret string
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetAccessibilityOk() (*string, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given string and assigns it to the Accessibility field.
func (o *HyperflexHxHostMountStatusDtAllOf) SetAccessibility(v string) {
	o.Accessibility = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexHxHostMountStatusDtAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexHxHostMountStatusDtAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetMounted returns the Mounted field value if set, zero value otherwise.
func (o *HyperflexHxHostMountStatusDtAllOf) GetMounted() bool {
	if o == nil || o.Mounted == nil {
		var ret bool
		return ret
	}
	return *o.Mounted
}

// GetMountedOk returns a tuple with the Mounted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetMountedOk() (*bool, bool) {
	if o == nil || o.Mounted == nil {
		return nil, false
	}
	return o.Mounted, true
}

// HasMounted returns a boolean if a field has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) HasMounted() bool {
	if o != nil && o.Mounted != nil {
		return true
	}

	return false
}

// SetMounted gets a reference to the given bool and assigns it to the Mounted field.
func (o *HyperflexHxHostMountStatusDtAllOf) SetMounted(v bool) {
	o.Mounted = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *HyperflexHxHostMountStatusDtAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *HyperflexHxHostMountStatusDtAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *HyperflexHxHostMountStatusDtAllOf) SetReason(v string) {
	o.Reason = &v
}

func (o HyperflexHxHostMountStatusDtAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Accessibility != nil {
		toSerialize["Accessibility"] = o.Accessibility
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.Mounted != nil {
		toSerialize["Mounted"] = o.Mounted
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexHxHostMountStatusDtAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexHxHostMountStatusDtAllOf := _HyperflexHxHostMountStatusDtAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexHxHostMountStatusDtAllOf); err == nil {
		*o = HyperflexHxHostMountStatusDtAllOf(varHyperflexHxHostMountStatusDtAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Accessibility")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "Mounted")
		delete(additionalProperties, "Reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexHxHostMountStatusDtAllOf struct {
	value *HyperflexHxHostMountStatusDtAllOf
	isSet bool
}

func (v NullableHyperflexHxHostMountStatusDtAllOf) Get() *HyperflexHxHostMountStatusDtAllOf {
	return v.value
}

func (v *NullableHyperflexHxHostMountStatusDtAllOf) Set(val *HyperflexHxHostMountStatusDtAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxHostMountStatusDtAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxHostMountStatusDtAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxHostMountStatusDtAllOf(val *HyperflexHxHostMountStatusDtAllOf) *NullableHyperflexHxHostMountStatusDtAllOf {
	return &NullableHyperflexHxHostMountStatusDtAllOf{value: val, isSet: true}
}

func (v NullableHyperflexHxHostMountStatusDtAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxHostMountStatusDtAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
