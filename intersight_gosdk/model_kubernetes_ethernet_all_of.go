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

// KubernetesEthernetAllOf Definition of the list of properties defined in 'kubernetes.Ethernet', excluding properties defined in parent classes.
type KubernetesEthernetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                            `json:"ObjectType"`
	Matcher    NullableKubernetesEthernetMatcher `json:"Matcher,omitempty"`
	// If the infrastructure network is selectable, this indicates which network to attach to the port.
	ProviderName         *string `json:"ProviderName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesEthernetAllOf KubernetesEthernetAllOf

// NewKubernetesEthernetAllOf instantiates a new KubernetesEthernetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEthernetAllOf(classId string, objectType string) *KubernetesEthernetAllOf {
	this := KubernetesEthernetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesEthernetAllOfWithDefaults instantiates a new KubernetesEthernetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEthernetAllOfWithDefaults() *KubernetesEthernetAllOf {
	this := KubernetesEthernetAllOf{}
	var classId string = "kubernetes.Ethernet"
	this.ClassId = classId
	var objectType string = "kubernetes.Ethernet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesEthernetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesEthernetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesEthernetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesEthernetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMatcher returns the Matcher field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesEthernetAllOf) GetMatcher() KubernetesEthernetMatcher {
	if o == nil || o.Matcher.Get() == nil {
		var ret KubernetesEthernetMatcher
		return ret
	}
	return *o.Matcher.Get()
}

// GetMatcherOk returns a tuple with the Matcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesEthernetAllOf) GetMatcherOk() (*KubernetesEthernetMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return o.Matcher.Get(), o.Matcher.IsSet()
}

// HasMatcher returns a boolean if a field has been set.
func (o *KubernetesEthernetAllOf) HasMatcher() bool {
	if o != nil && o.Matcher.IsSet() {
		return true
	}

	return false
}

// SetMatcher gets a reference to the given NullableKubernetesEthernetMatcher and assigns it to the Matcher field.
func (o *KubernetesEthernetAllOf) SetMatcher(v KubernetesEthernetMatcher) {
	o.Matcher.Set(&v)
}

// SetMatcherNil sets the value for Matcher to be an explicit nil
func (o *KubernetesEthernetAllOf) SetMatcherNil() {
	o.Matcher.Set(nil)
}

// UnsetMatcher ensures that no value is present for Matcher, not even an explicit nil
func (o *KubernetesEthernetAllOf) UnsetMatcher() {
	o.Matcher.Unset()
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *KubernetesEthernetAllOf) GetProviderName() string {
	if o == nil || o.ProviderName == nil {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesEthernetAllOf) GetProviderNameOk() (*string, bool) {
	if o == nil || o.ProviderName == nil {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *KubernetesEthernetAllOf) HasProviderName() bool {
	if o != nil && o.ProviderName != nil {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *KubernetesEthernetAllOf) SetProviderName(v string) {
	o.ProviderName = &v
}

func (o KubernetesEthernetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Matcher.IsSet() {
		toSerialize["Matcher"] = o.Matcher.Get()
	}
	if o.ProviderName != nil {
		toSerialize["ProviderName"] = o.ProviderName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesEthernetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesEthernetAllOf := _KubernetesEthernetAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesEthernetAllOf); err == nil {
		*o = KubernetesEthernetAllOf(varKubernetesEthernetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Matcher")
		delete(additionalProperties, "ProviderName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesEthernetAllOf struct {
	value *KubernetesEthernetAllOf
	isSet bool
}

func (v NullableKubernetesEthernetAllOf) Get() *KubernetesEthernetAllOf {
	return v.value
}

func (v *NullableKubernetesEthernetAllOf) Set(val *KubernetesEthernetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEthernetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEthernetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEthernetAllOf(val *KubernetesEthernetAllOf) *NullableKubernetesEthernetAllOf {
	return &NullableKubernetesEthernetAllOf{value: val, isSet: true}
}

func (v NullableKubernetesEthernetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEthernetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
