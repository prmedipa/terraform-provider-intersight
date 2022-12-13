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

// AssetApiKeyCredential API key based authentication.
type AssetApiKeyCredential struct {
	AssetCredential
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This a secret API key which can be used for authentication purposes for different targets like Azure Enterprise Agreement.
	ApiKey *string `json:"ApiKey,omitempty"`
	// Indicates whether the value of the 'apiKey' property has been set.
	IsApiKeySet          *bool `json:"IsApiKeySet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetApiKeyCredential AssetApiKeyCredential

// NewAssetApiKeyCredential instantiates a new AssetApiKeyCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetApiKeyCredential(classId string, objectType string) *AssetApiKeyCredential {
	this := AssetApiKeyCredential{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetApiKeyCredentialWithDefaults instantiates a new AssetApiKeyCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetApiKeyCredentialWithDefaults() *AssetApiKeyCredential {
	this := AssetApiKeyCredential{}
	var classId string = "asset.ApiKeyCredential"
	this.ClassId = classId
	var objectType string = "asset.ApiKeyCredential"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetApiKeyCredential) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetApiKeyCredential) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetApiKeyCredential) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetApiKeyCredential) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetApiKeyCredential) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetApiKeyCredential) SetObjectType(v string) {
	o.ObjectType = v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *AssetApiKeyCredential) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetApiKeyCredential) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *AssetApiKeyCredential) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *AssetApiKeyCredential) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetIsApiKeySet returns the IsApiKeySet field value if set, zero value otherwise.
func (o *AssetApiKeyCredential) GetIsApiKeySet() bool {
	if o == nil || o.IsApiKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsApiKeySet
}

// GetIsApiKeySetOk returns a tuple with the IsApiKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetApiKeyCredential) GetIsApiKeySetOk() (*bool, bool) {
	if o == nil || o.IsApiKeySet == nil {
		return nil, false
	}
	return o.IsApiKeySet, true
}

// HasIsApiKeySet returns a boolean if a field has been set.
func (o *AssetApiKeyCredential) HasIsApiKeySet() bool {
	if o != nil && o.IsApiKeySet != nil {
		return true
	}

	return false
}

// SetIsApiKeySet gets a reference to the given bool and assigns it to the IsApiKeySet field.
func (o *AssetApiKeyCredential) SetIsApiKeySet(v bool) {
	o.IsApiKeySet = &v
}

func (o AssetApiKeyCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetCredential, errAssetCredential := json.Marshal(o.AssetCredential)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	errAssetCredential = json.Unmarshal([]byte(serializedAssetCredential), &toSerialize)
	if errAssetCredential != nil {
		return []byte{}, errAssetCredential
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ApiKey != nil {
		toSerialize["ApiKey"] = o.ApiKey
	}
	if o.IsApiKeySet != nil {
		toSerialize["IsApiKeySet"] = o.IsApiKeySet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetApiKeyCredential) UnmarshalJSON(bytes []byte) (err error) {
	type AssetApiKeyCredentialWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// This a secret API key which can be used for authentication purposes for different targets like Azure Enterprise Agreement.
		ApiKey *string `json:"ApiKey,omitempty"`
		// Indicates whether the value of the 'apiKey' property has been set.
		IsApiKeySet *bool `json:"IsApiKeySet,omitempty"`
	}

	varAssetApiKeyCredentialWithoutEmbeddedStruct := AssetApiKeyCredentialWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetApiKeyCredentialWithoutEmbeddedStruct)
	if err == nil {
		varAssetApiKeyCredential := _AssetApiKeyCredential{}
		varAssetApiKeyCredential.ClassId = varAssetApiKeyCredentialWithoutEmbeddedStruct.ClassId
		varAssetApiKeyCredential.ObjectType = varAssetApiKeyCredentialWithoutEmbeddedStruct.ObjectType
		varAssetApiKeyCredential.ApiKey = varAssetApiKeyCredentialWithoutEmbeddedStruct.ApiKey
		varAssetApiKeyCredential.IsApiKeySet = varAssetApiKeyCredentialWithoutEmbeddedStruct.IsApiKeySet
		*o = AssetApiKeyCredential(varAssetApiKeyCredential)
	} else {
		return err
	}

	varAssetApiKeyCredential := _AssetApiKeyCredential{}

	err = json.Unmarshal(bytes, &varAssetApiKeyCredential)
	if err == nil {
		o.AssetCredential = varAssetApiKeyCredential.AssetCredential
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ApiKey")
		delete(additionalProperties, "IsApiKeySet")

		// remove fields from embedded structs
		reflectAssetCredential := reflect.ValueOf(o.AssetCredential)
		for i := 0; i < reflectAssetCredential.Type().NumField(); i++ {
			t := reflectAssetCredential.Type().Field(i)

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

type NullableAssetApiKeyCredential struct {
	value *AssetApiKeyCredential
	isSet bool
}

func (v NullableAssetApiKeyCredential) Get() *AssetApiKeyCredential {
	return v.value
}

func (v *NullableAssetApiKeyCredential) Set(val *AssetApiKeyCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetApiKeyCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetApiKeyCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetApiKeyCredential(val *AssetApiKeyCredential) *NullableAssetApiKeyCredential {
	return &NullableAssetApiKeyCredential{value: val, isSet: true}
}

func (v NullableAssetApiKeyCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetApiKeyCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
