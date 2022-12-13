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

// AssetOrchestrationHsmVmwareVcenterOptions Captures configuration specific to the VMware Vcenter target for the Hardware Support Manager.
type AssetOrchestrationHsmVmwareVcenterOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later.
	HsmEnabled *bool `json:"HsmEnabled,omitempty"`
	// Indicates whether the value of the 'clientSecret' property has been set.
	IsClientSecretSet    *bool `json:"IsClientSecretSet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetOrchestrationHsmVmwareVcenterOptions AssetOrchestrationHsmVmwareVcenterOptions

// NewAssetOrchestrationHsmVmwareVcenterOptions instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetOrchestrationHsmVmwareVcenterOptions(classId string, objectType string) *AssetOrchestrationHsmVmwareVcenterOptions {
	this := AssetOrchestrationHsmVmwareVcenterOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults instantiates a new AssetOrchestrationHsmVmwareVcenterOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetOrchestrationHsmVmwareVcenterOptionsWithDefaults() *AssetOrchestrationHsmVmwareVcenterOptions {
	this := AssetOrchestrationHsmVmwareVcenterOptions{}
	var classId string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ClassId = classId
	var objectType string = "asset.OrchestrationHsmVmwareVcenterOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHsmEnabled returns the HsmEnabled field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabled() bool {
	if o == nil || o.HsmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HsmEnabled
}

// GetHsmEnabledOk returns a tuple with the HsmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetHsmEnabledOk() (*bool, bool) {
	if o == nil || o.HsmEnabled == nil {
		return nil, false
	}
	return o.HsmEnabled, true
}

// HasHsmEnabled returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasHsmEnabled() bool {
	if o != nil && o.HsmEnabled != nil {
		return true
	}

	return false
}

// SetHsmEnabled gets a reference to the given bool and assigns it to the HsmEnabled field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetHsmEnabled(v bool) {
	o.HsmEnabled = &v
}

// GetIsClientSecretSet returns the IsClientSecretSet field value if set, zero value otherwise.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSet() bool {
	if o == nil || o.IsClientSecretSet == nil {
		var ret bool
		return ret
	}
	return *o.IsClientSecretSet
}

// GetIsClientSecretSetOk returns a tuple with the IsClientSecretSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) GetIsClientSecretSetOk() (*bool, bool) {
	if o == nil || o.IsClientSecretSet == nil {
		return nil, false
	}
	return o.IsClientSecretSet, true
}

// HasIsClientSecretSet returns a boolean if a field has been set.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) HasIsClientSecretSet() bool {
	if o != nil && o.IsClientSecretSet != nil {
		return true
	}

	return false
}

// SetIsClientSecretSet gets a reference to the given bool and assigns it to the IsClientSecretSet field.
func (o *AssetOrchestrationHsmVmwareVcenterOptions) SetIsClientSecretSet(v bool) {
	o.IsClientSecretSet = &v
}

func (o AssetOrchestrationHsmVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HsmEnabled != nil {
		toSerialize["HsmEnabled"] = o.HsmEnabled
	}
	if o.IsClientSecretSet != nil {
		toSerialize["IsClientSecretSet"] = o.IsClientSecretSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetOrchestrationHsmVmwareVcenterOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// HsmEnabled controls whether Hardware Support Manager is enabled or not. vCenter Server version 7.0 or later.
		HsmEnabled *bool `json:"HsmEnabled,omitempty"`
		// Indicates whether the value of the 'clientSecret' property has been set.
		IsClientSecretSet *bool `json:"IsClientSecretSet,omitempty"`
	}

	varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct := AssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetOrchestrationHsmVmwareVcenterOptions := _AssetOrchestrationHsmVmwareVcenterOptions{}
		varAssetOrchestrationHsmVmwareVcenterOptions.ClassId = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.ClassId
		varAssetOrchestrationHsmVmwareVcenterOptions.ObjectType = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.ObjectType
		varAssetOrchestrationHsmVmwareVcenterOptions.HsmEnabled = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.HsmEnabled
		varAssetOrchestrationHsmVmwareVcenterOptions.IsClientSecretSet = varAssetOrchestrationHsmVmwareVcenterOptionsWithoutEmbeddedStruct.IsClientSecretSet
		*o = AssetOrchestrationHsmVmwareVcenterOptions(varAssetOrchestrationHsmVmwareVcenterOptions)
	} else {
		return err
	}

	varAssetOrchestrationHsmVmwareVcenterOptions := _AssetOrchestrationHsmVmwareVcenterOptions{}

	err = json.Unmarshal(bytes, &varAssetOrchestrationHsmVmwareVcenterOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetOrchestrationHsmVmwareVcenterOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "HsmEnabled")
		delete(additionalProperties, "IsClientSecretSet")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetOrchestrationHsmVmwareVcenterOptions struct {
	value *AssetOrchestrationHsmVmwareVcenterOptions
	isSet bool
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) Get() *AssetOrchestrationHsmVmwareVcenterOptions {
	return v.value
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) Set(val *AssetOrchestrationHsmVmwareVcenterOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOrchestrationHsmVmwareVcenterOptions(val *AssetOrchestrationHsmVmwareVcenterOptions) *NullableAssetOrchestrationHsmVmwareVcenterOptions {
	return &NullableAssetOrchestrationHsmVmwareVcenterOptions{value: val, isSet: true}
}

func (v NullableAssetOrchestrationHsmVmwareVcenterOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOrchestrationHsmVmwareVcenterOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
