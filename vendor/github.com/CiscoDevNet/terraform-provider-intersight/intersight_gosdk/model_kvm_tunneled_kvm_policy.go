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

// KvmTunneledKvmPolicy Policy to control Tunelled vKVM for a specific account.
type KvmTunneledKvmPolicy struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or Disable configuration of tunneled KVM for a specific account.
	TunneledKvmConfiguration *bool `json:"TunneledKvmConfiguration,omitempty"`
	// Enable or Disable launching tunneled KVM for a specific account.
	TunneledKvmLaunch    *bool                   `json:"TunneledKvmLaunch,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmTunneledKvmPolicy KvmTunneledKvmPolicy

// NewKvmTunneledKvmPolicy instantiates a new KvmTunneledKvmPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmTunneledKvmPolicy(classId string, objectType string) *KvmTunneledKvmPolicy {
	this := KvmTunneledKvmPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var tunneledKvmConfiguration bool = false
	this.TunneledKvmConfiguration = &tunneledKvmConfiguration
	var tunneledKvmLaunch bool = false
	this.TunneledKvmLaunch = &tunneledKvmLaunch
	return &this
}

// NewKvmTunneledKvmPolicyWithDefaults instantiates a new KvmTunneledKvmPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmTunneledKvmPolicyWithDefaults() *KvmTunneledKvmPolicy {
	this := KvmTunneledKvmPolicy{}
	var classId string = "kvm.TunneledKvmPolicy"
	this.ClassId = classId
	var objectType string = "kvm.TunneledKvmPolicy"
	this.ObjectType = objectType
	var tunneledKvmConfiguration bool = false
	this.TunneledKvmConfiguration = &tunneledKvmConfiguration
	var tunneledKvmLaunch bool = false
	this.TunneledKvmLaunch = &tunneledKvmLaunch
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmTunneledKvmPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmTunneledKvmPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmTunneledKvmPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmTunneledKvmPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetTunneledKvmConfiguration returns the TunneledKvmConfiguration field value if set, zero value otherwise.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfiguration() bool {
	if o == nil || o.TunneledKvmConfiguration == nil {
		var ret bool
		return ret
	}
	return *o.TunneledKvmConfiguration
}

// GetTunneledKvmConfigurationOk returns a tuple with the TunneledKvmConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmConfigurationOk() (*bool, bool) {
	if o == nil || o.TunneledKvmConfiguration == nil {
		return nil, false
	}
	return o.TunneledKvmConfiguration, true
}

// HasTunneledKvmConfiguration returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasTunneledKvmConfiguration() bool {
	if o != nil && o.TunneledKvmConfiguration != nil {
		return true
	}

	return false
}

// SetTunneledKvmConfiguration gets a reference to the given bool and assigns it to the TunneledKvmConfiguration field.
func (o *KvmTunneledKvmPolicy) SetTunneledKvmConfiguration(v bool) {
	o.TunneledKvmConfiguration = &v
}

// GetTunneledKvmLaunch returns the TunneledKvmLaunch field value if set, zero value otherwise.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunch() bool {
	if o == nil || o.TunneledKvmLaunch == nil {
		var ret bool
		return ret
	}
	return *o.TunneledKvmLaunch
}

// GetTunneledKvmLaunchOk returns a tuple with the TunneledKvmLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetTunneledKvmLaunchOk() (*bool, bool) {
	if o == nil || o.TunneledKvmLaunch == nil {
		return nil, false
	}
	return o.TunneledKvmLaunch, true
}

// HasTunneledKvmLaunch returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasTunneledKvmLaunch() bool {
	if o != nil && o.TunneledKvmLaunch != nil {
		return true
	}

	return false
}

// SetTunneledKvmLaunch gets a reference to the given bool and assigns it to the TunneledKvmLaunch field.
func (o *KvmTunneledKvmPolicy) SetTunneledKvmLaunch(v bool) {
	o.TunneledKvmLaunch = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *KvmTunneledKvmPolicy) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmTunneledKvmPolicy) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *KvmTunneledKvmPolicy) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *KvmTunneledKvmPolicy) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o KvmTunneledKvmPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.TunneledKvmConfiguration != nil {
		toSerialize["TunneledKvmConfiguration"] = o.TunneledKvmConfiguration
	}
	if o.TunneledKvmLaunch != nil {
		toSerialize["TunneledKvmLaunch"] = o.TunneledKvmLaunch
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmTunneledKvmPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KvmTunneledKvmPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or Disable configuration of tunneled KVM for a specific account.
		TunneledKvmConfiguration *bool `json:"TunneledKvmConfiguration,omitempty"`
		// Enable or Disable launching tunneled KVM for a specific account.
		TunneledKvmLaunch *bool                   `json:"TunneledKvmLaunch,omitempty"`
		Account           *IamAccountRelationship `json:"Account,omitempty"`
	}

	varKvmTunneledKvmPolicyWithoutEmbeddedStruct := KvmTunneledKvmPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmTunneledKvmPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKvmTunneledKvmPolicy := _KvmTunneledKvmPolicy{}
		varKvmTunneledKvmPolicy.ClassId = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.ClassId
		varKvmTunneledKvmPolicy.ObjectType = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.ObjectType
		varKvmTunneledKvmPolicy.TunneledKvmConfiguration = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.TunneledKvmConfiguration
		varKvmTunneledKvmPolicy.TunneledKvmLaunch = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.TunneledKvmLaunch
		varKvmTunneledKvmPolicy.Account = varKvmTunneledKvmPolicyWithoutEmbeddedStruct.Account
		*o = KvmTunneledKvmPolicy(varKvmTunneledKvmPolicy)
	} else {
		return err
	}

	varKvmTunneledKvmPolicy := _KvmTunneledKvmPolicy{}

	err = json.Unmarshal(bytes, &varKvmTunneledKvmPolicy)
	if err == nil {
		o.MoBaseMo = varKvmTunneledKvmPolicy.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "TunneledKvmConfiguration")
		delete(additionalProperties, "TunneledKvmLaunch")
		delete(additionalProperties, "Account")

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

type NullableKvmTunneledKvmPolicy struct {
	value *KvmTunneledKvmPolicy
	isSet bool
}

func (v NullableKvmTunneledKvmPolicy) Get() *KvmTunneledKvmPolicy {
	return v.value
}

func (v *NullableKvmTunneledKvmPolicy) Set(val *KvmTunneledKvmPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmTunneledKvmPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmTunneledKvmPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmTunneledKvmPolicy(val *KvmTunneledKvmPolicy) *NullableKvmTunneledKvmPolicy {
	return &NullableKvmTunneledKvmPolicy{value: val, isSet: true}
}

func (v NullableKvmTunneledKvmPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmTunneledKvmPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
