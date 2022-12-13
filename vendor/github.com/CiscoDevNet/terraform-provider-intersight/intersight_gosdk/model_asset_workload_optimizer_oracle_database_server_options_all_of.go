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

// AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerOracleDatabaseServerOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A unique name for an Oracle database instance on a specific host.
	DatabaseId           *string `json:"DatabaseId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf

// NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf instantiates a new AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf {
	this := AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOfWithDefaults() *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf {
	this := AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerOracleDatabaseServerOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerOracleDatabaseServerOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDatabaseId returns the DatabaseId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetDatabaseId() string {
	if o == nil || o.DatabaseId == nil {
		var ret string
		return ret
	}
	return *o.DatabaseId
}

// GetDatabaseIdOk returns a tuple with the DatabaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) GetDatabaseIdOk() (*string, bool) {
	if o == nil || o.DatabaseId == nil {
		return nil, false
	}
	return o.DatabaseId, true
}

// HasDatabaseId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) HasDatabaseId() bool {
	if o != nil && o.DatabaseId != nil {
		return true
	}

	return false
}

// SetDatabaseId gets a reference to the given string and assigns it to the DatabaseId field.
func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) SetDatabaseId(v string) {
	o.DatabaseId = &v
}

func (o AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DatabaseId != nil {
		toSerialize["DatabaseId"] = o.DatabaseId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf := _AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf(varAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DatabaseId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf struct {
	value *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) Get() *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) Set(val *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf(val *AssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) *NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf {
	return &NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerOracleDatabaseServerOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
