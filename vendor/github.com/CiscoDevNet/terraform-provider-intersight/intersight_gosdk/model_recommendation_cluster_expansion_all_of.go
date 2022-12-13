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

// RecommendationClusterExpansionAllOf Definition of the list of properties defined in 'recommendation.ClusterExpansion', excluding properties defined in parent classes.
type RecommendationClusterExpansionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the cluster for which the expansion recommendation is provided.
	ClusterName              *string                                             `json:"ClusterName,omitempty"`
	HardwareExpansionRequest *RecommendationHardwareExpansionRequestRelationship `json:"HardwareExpansionRequest,omitempty"`
	// An array of relationships to recommendationPhysicalItem resources.
	PhysicalItem     []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty"`
	// An array of relationships to recommendationSoftwareItem resources.
	SoftwareItem         []RecommendationSoftwareItemRelationship `json:"SoftwareItem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationClusterExpansionAllOf RecommendationClusterExpansionAllOf

// NewRecommendationClusterExpansionAllOf instantiates a new RecommendationClusterExpansionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationClusterExpansionAllOf(classId string, objectType string) *RecommendationClusterExpansionAllOf {
	this := RecommendationClusterExpansionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationClusterExpansionAllOfWithDefaults instantiates a new RecommendationClusterExpansionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationClusterExpansionAllOfWithDefaults() *RecommendationClusterExpansionAllOf {
	this := RecommendationClusterExpansionAllOf{}
	var classId string = "recommendation.ClusterExpansion"
	this.ClassId = classId
	var objectType string = "recommendation.ClusterExpansion"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationClusterExpansionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationClusterExpansionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationClusterExpansionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationClusterExpansionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationClusterExpansionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationClusterExpansionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *RecommendationClusterExpansionAllOf) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationClusterExpansionAllOf) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *RecommendationClusterExpansionAllOf) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *RecommendationClusterExpansionAllOf) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetHardwareExpansionRequest returns the HardwareExpansionRequest field value if set, zero value otherwise.
func (o *RecommendationClusterExpansionAllOf) GetHardwareExpansionRequest() RecommendationHardwareExpansionRequestRelationship {
	if o == nil || o.HardwareExpansionRequest == nil {
		var ret RecommendationHardwareExpansionRequestRelationship
		return ret
	}
	return *o.HardwareExpansionRequest
}

// GetHardwareExpansionRequestOk returns a tuple with the HardwareExpansionRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationClusterExpansionAllOf) GetHardwareExpansionRequestOk() (*RecommendationHardwareExpansionRequestRelationship, bool) {
	if o == nil || o.HardwareExpansionRequest == nil {
		return nil, false
	}
	return o.HardwareExpansionRequest, true
}

// HasHardwareExpansionRequest returns a boolean if a field has been set.
func (o *RecommendationClusterExpansionAllOf) HasHardwareExpansionRequest() bool {
	if o != nil && o.HardwareExpansionRequest != nil {
		return true
	}

	return false
}

// SetHardwareExpansionRequest gets a reference to the given RecommendationHardwareExpansionRequestRelationship and assigns it to the HardwareExpansionRequest field.
func (o *RecommendationClusterExpansionAllOf) SetHardwareExpansionRequest(v RecommendationHardwareExpansionRequestRelationship) {
	o.HardwareExpansionRequest = &v
}

// GetPhysicalItem returns the PhysicalItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationClusterExpansionAllOf) GetPhysicalItem() []RecommendationPhysicalItemRelationship {
	if o == nil {
		var ret []RecommendationPhysicalItemRelationship
		return ret
	}
	return o.PhysicalItem
}

// GetPhysicalItemOk returns a tuple with the PhysicalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationClusterExpansionAllOf) GetPhysicalItemOk() ([]RecommendationPhysicalItemRelationship, bool) {
	if o == nil || o.PhysicalItem == nil {
		return nil, false
	}
	return o.PhysicalItem, true
}

// HasPhysicalItem returns a boolean if a field has been set.
func (o *RecommendationClusterExpansionAllOf) HasPhysicalItem() bool {
	if o != nil && o.PhysicalItem != nil {
		return true
	}

	return false
}

// SetPhysicalItem gets a reference to the given []RecommendationPhysicalItemRelationship and assigns it to the PhysicalItem field.
func (o *RecommendationClusterExpansionAllOf) SetPhysicalItem(v []RecommendationPhysicalItemRelationship) {
	o.PhysicalItem = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *RecommendationClusterExpansionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationClusterExpansionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *RecommendationClusterExpansionAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *RecommendationClusterExpansionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSoftwareItem returns the SoftwareItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationClusterExpansionAllOf) GetSoftwareItem() []RecommendationSoftwareItemRelationship {
	if o == nil {
		var ret []RecommendationSoftwareItemRelationship
		return ret
	}
	return o.SoftwareItem
}

// GetSoftwareItemOk returns a tuple with the SoftwareItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationClusterExpansionAllOf) GetSoftwareItemOk() ([]RecommendationSoftwareItemRelationship, bool) {
	if o == nil || o.SoftwareItem == nil {
		return nil, false
	}
	return o.SoftwareItem, true
}

// HasSoftwareItem returns a boolean if a field has been set.
func (o *RecommendationClusterExpansionAllOf) HasSoftwareItem() bool {
	if o != nil && o.SoftwareItem != nil {
		return true
	}

	return false
}

// SetSoftwareItem gets a reference to the given []RecommendationSoftwareItemRelationship and assigns it to the SoftwareItem field.
func (o *RecommendationClusterExpansionAllOf) SetSoftwareItem(v []RecommendationSoftwareItemRelationship) {
	o.SoftwareItem = v
}

func (o RecommendationClusterExpansionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClusterName != nil {
		toSerialize["ClusterName"] = o.ClusterName
	}
	if o.HardwareExpansionRequest != nil {
		toSerialize["HardwareExpansionRequest"] = o.HardwareExpansionRequest
	}
	if o.PhysicalItem != nil {
		toSerialize["PhysicalItem"] = o.PhysicalItem
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.SoftwareItem != nil {
		toSerialize["SoftwareItem"] = o.SoftwareItem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationClusterExpansionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRecommendationClusterExpansionAllOf := _RecommendationClusterExpansionAllOf{}

	if err = json.Unmarshal(bytes, &varRecommendationClusterExpansionAllOf); err == nil {
		*o = RecommendationClusterExpansionAllOf(varRecommendationClusterExpansionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterName")
		delete(additionalProperties, "HardwareExpansionRequest")
		delete(additionalProperties, "PhysicalItem")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "SoftwareItem")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecommendationClusterExpansionAllOf struct {
	value *RecommendationClusterExpansionAllOf
	isSet bool
}

func (v NullableRecommendationClusterExpansionAllOf) Get() *RecommendationClusterExpansionAllOf {
	return v.value
}

func (v *NullableRecommendationClusterExpansionAllOf) Set(val *RecommendationClusterExpansionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationClusterExpansionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationClusterExpansionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationClusterExpansionAllOf(val *RecommendationClusterExpansionAllOf) *NullableRecommendationClusterExpansionAllOf {
	return &NullableRecommendationClusterExpansionAllOf{value: val, isSet: true}
}

func (v NullableRecommendationClusterExpansionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationClusterExpansionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
