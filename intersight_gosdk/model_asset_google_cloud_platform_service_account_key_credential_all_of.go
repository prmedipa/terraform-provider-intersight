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

// AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf Definition of the list of properties defined in 'asset.GoogleCloudPlatformServiceAccountKeyCredential', excluding properties defined in parent classes.
type AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'serviceAccountKey' property has been set.
	IsServiceAccountKeySet *bool `json:"IsServiceAccountKeySet,omitempty"`
	// Google Cloud Platform (GCP) service account's key type. * `JSON` - JavaScript Object Notation (JSON) is a standard text-based format for representing structured data based on JavaScript object syntax. It is commonly used for transmitting data in web applications. * `P12` - PKCS #12 is an archive file format for storing many cryptography objects as a single file. It is commonly used to bundle a private key with its X.509 certificate or to bundle all the members of a chain of trust.
	KeyType *string `json:"KeyType,omitempty"`
	// Google Cloud Platform (GCP) service account's key file content in the format of the keyType specified.
	ServiceAccountKey    *string `json:"ServiceAccountKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf

// NewAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf instantiates a new AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf(classId string, objectType string) *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf {
	this := AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var keyType string = "JSON"
	this.KeyType = &keyType
	return &this
}

// NewAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOfWithDefaults instantiates a new AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOfWithDefaults() *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf {
	this := AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf{}
	var classId string = "asset.GoogleCloudPlatformServiceAccountKeyCredential"
	this.ClassId = classId
	var objectType string = "asset.GoogleCloudPlatformServiceAccountKeyCredential"
	this.ObjectType = objectType
	var keyType string = "JSON"
	this.KeyType = &keyType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsServiceAccountKeySet returns the IsServiceAccountKeySet field value if set, zero value otherwise.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetIsServiceAccountKeySet() bool {
	if o == nil || o.IsServiceAccountKeySet == nil {
		var ret bool
		return ret
	}
	return *o.IsServiceAccountKeySet
}

// GetIsServiceAccountKeySetOk returns a tuple with the IsServiceAccountKeySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetIsServiceAccountKeySetOk() (*bool, bool) {
	if o == nil || o.IsServiceAccountKeySet == nil {
		return nil, false
	}
	return o.IsServiceAccountKeySet, true
}

// HasIsServiceAccountKeySet returns a boolean if a field has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) HasIsServiceAccountKeySet() bool {
	if o != nil && o.IsServiceAccountKeySet != nil {
		return true
	}

	return false
}

// SetIsServiceAccountKeySet gets a reference to the given bool and assigns it to the IsServiceAccountKeySet field.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) SetIsServiceAccountKeySet(v bool) {
	o.IsServiceAccountKeySet = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) SetKeyType(v string) {
	o.KeyType = &v
}

// GetServiceAccountKey returns the ServiceAccountKey field value if set, zero value otherwise.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetServiceAccountKey() string {
	if o == nil || o.ServiceAccountKey == nil {
		var ret string
		return ret
	}
	return *o.ServiceAccountKey
}

// GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) GetServiceAccountKeyOk() (*string, bool) {
	if o == nil || o.ServiceAccountKey == nil {
		return nil, false
	}
	return o.ServiceAccountKey, true
}

// HasServiceAccountKey returns a boolean if a field has been set.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) HasServiceAccountKey() bool {
	if o != nil && o.ServiceAccountKey != nil {
		return true
	}

	return false
}

// SetServiceAccountKey gets a reference to the given string and assigns it to the ServiceAccountKey field.
func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) SetServiceAccountKey(v string) {
	o.ServiceAccountKey = &v
}

func (o AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsServiceAccountKeySet != nil {
		toSerialize["IsServiceAccountKeySet"] = o.IsServiceAccountKeySet
	}
	if o.KeyType != nil {
		toSerialize["KeyType"] = o.KeyType
	}
	if o.ServiceAccountKey != nil {
		toSerialize["ServiceAccountKey"] = o.ServiceAccountKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf := _AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf{}

	if err = json.Unmarshal(bytes, &varAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf); err == nil {
		*o = AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf(varAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsServiceAccountKeySet")
		delete(additionalProperties, "KeyType")
		delete(additionalProperties, "ServiceAccountKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf struct {
	value *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf
	isSet bool
}

func (v NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) Get() *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf {
	return v.value
}

func (v *NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) Set(val *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf(val *AssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) *NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf {
	return &NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf{value: val, isSet: true}
}

func (v NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetGoogleCloudPlatformServiceAccountKeyCredentialAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
