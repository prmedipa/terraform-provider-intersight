/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the MetricsResourceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsResourceConfiguration{}

// MetricsResourceConfiguration The configuration of metric collection for a given resource within an account. Each resource that is capable of collecting metrics into the Intersight system may be configured with options to control the metric collection behavior.
type MetricsResourceConfiguration struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Metric collection is enabled for this resource, when enabled all available metrics are collected from the resource into Intersight.
	Enabled              *bool                           `json:"Enabled,omitempty"`
	Domain               NullableAssetTargetRelationship `json:"Domain,omitempty"`
	Resource             NullableMoBaseMoRelationship    `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetricsResourceConfiguration MetricsResourceConfiguration

// NewMetricsResourceConfiguration instantiates a new MetricsResourceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsResourceConfiguration(classId string, objectType string) *MetricsResourceConfiguration {
	this := MetricsResourceConfiguration{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMetricsResourceConfigurationWithDefaults instantiates a new MetricsResourceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsResourceConfigurationWithDefaults() *MetricsResourceConfiguration {
	this := MetricsResourceConfiguration{}
	var classId string = "metrics.ResourceConfiguration"
	this.ClassId = classId
	var objectType string = "metrics.ResourceConfiguration"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MetricsResourceConfiguration) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MetricsResourceConfiguration) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MetricsResourceConfiguration) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "metrics.ResourceConfiguration" of the ClassId field.
func (o *MetricsResourceConfiguration) GetDefaultClassId() interface{} {
	return "metrics.ResourceConfiguration"
}

// GetObjectType returns the ObjectType field value
func (o *MetricsResourceConfiguration) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MetricsResourceConfiguration) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MetricsResourceConfiguration) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "metrics.ResourceConfiguration" of the ObjectType field.
func (o *MetricsResourceConfiguration) GetDefaultObjectType() interface{} {
	return "metrics.ResourceConfiguration"
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MetricsResourceConfiguration) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsResourceConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MetricsResourceConfiguration) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MetricsResourceConfiguration) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsResourceConfiguration) GetDomain() AssetTargetRelationship {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret AssetTargetRelationship
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsResourceConfiguration) GetDomainOk() (*AssetTargetRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *MetricsResourceConfiguration) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableAssetTargetRelationship and assigns it to the Domain field.
func (o *MetricsResourceConfiguration) SetDomain(v AssetTargetRelationship) {
	o.Domain.Set(&v)
}

// SetDomainNil sets the value for Domain to be an explicit nil
func (o *MetricsResourceConfiguration) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *MetricsResourceConfiguration) UnsetDomain() {
	o.Domain.Unset()
}

// GetResource returns the Resource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricsResourceConfiguration) GetResource() MoBaseMoRelationship {
	if o == nil || IsNil(o.Resource.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource.Get()
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricsResourceConfiguration) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resource.Get(), o.Resource.IsSet()
}

// HasResource returns a boolean if a field has been set.
func (o *MetricsResourceConfiguration) HasResource() bool {
	if o != nil && o.Resource.IsSet() {
		return true
	}

	return false
}

// SetResource gets a reference to the given NullableMoBaseMoRelationship and assigns it to the Resource field.
func (o *MetricsResourceConfiguration) SetResource(v MoBaseMoRelationship) {
	o.Resource.Set(&v)
}

// SetResourceNil sets the value for Resource to be an explicit nil
func (o *MetricsResourceConfiguration) SetResourceNil() {
	o.Resource.Set(nil)
}

// UnsetResource ensures that no value is present for Resource, not even an explicit nil
func (o *MetricsResourceConfiguration) UnsetResource() {
	o.Resource.Unset()
}

func (o MetricsResourceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsResourceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Domain.IsSet() {
		toSerialize["Domain"] = o.Domain.Get()
	}
	if o.Resource.IsSet() {
		toSerialize["Resource"] = o.Resource.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetricsResourceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type MetricsResourceConfigurationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Metric collection is enabled for this resource, when enabled all available metrics are collected from the resource into Intersight.
		Enabled  *bool                           `json:"Enabled,omitempty"`
		Domain   NullableAssetTargetRelationship `json:"Domain,omitempty"`
		Resource NullableMoBaseMoRelationship    `json:"Resource,omitempty"`
	}

	varMetricsResourceConfigurationWithoutEmbeddedStruct := MetricsResourceConfigurationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varMetricsResourceConfigurationWithoutEmbeddedStruct)
	if err == nil {
		varMetricsResourceConfiguration := _MetricsResourceConfiguration{}
		varMetricsResourceConfiguration.ClassId = varMetricsResourceConfigurationWithoutEmbeddedStruct.ClassId
		varMetricsResourceConfiguration.ObjectType = varMetricsResourceConfigurationWithoutEmbeddedStruct.ObjectType
		varMetricsResourceConfiguration.Enabled = varMetricsResourceConfigurationWithoutEmbeddedStruct.Enabled
		varMetricsResourceConfiguration.Domain = varMetricsResourceConfigurationWithoutEmbeddedStruct.Domain
		varMetricsResourceConfiguration.Resource = varMetricsResourceConfigurationWithoutEmbeddedStruct.Resource
		*o = MetricsResourceConfiguration(varMetricsResourceConfiguration)
	} else {
		return err
	}

	varMetricsResourceConfiguration := _MetricsResourceConfiguration{}

	err = json.Unmarshal(data, &varMetricsResourceConfiguration)
	if err == nil {
		o.MoBaseMo = varMetricsResourceConfiguration.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Domain")
		delete(additionalProperties, "Resource")

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

type NullableMetricsResourceConfiguration struct {
	value *MetricsResourceConfiguration
	isSet bool
}

func (v NullableMetricsResourceConfiguration) Get() *MetricsResourceConfiguration {
	return v.value
}

func (v *NullableMetricsResourceConfiguration) Set(val *MetricsResourceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsResourceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsResourceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsResourceConfiguration(val *MetricsResourceConfiguration) *NullableMetricsResourceConfiguration {
	return &NullableMetricsResourceConfiguration{value: val, isSet: true}
}

func (v NullableMetricsResourceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsResourceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
