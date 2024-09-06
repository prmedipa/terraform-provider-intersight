/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-18369
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the RecommendationBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecommendationBase{}

// RecommendationBase Entity representing recommendation base properties like name and last updated time.
type RecommendationBase struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The time when the recommendation was last updated.
	LastUpdatedTime *time.Time `json:"LastUpdatedTime,omitempty"`
	// The name of the recommendation.
	Name *string `json:"Name,omitempty"`
	// Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false.
	RequirementMet       *bool `json:"RequirementMet,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationBase RecommendationBase

// NewRecommendationBase instantiates a new RecommendationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationBase(classId string, objectType string) *RecommendationBase {
	this := RecommendationBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationBaseWithDefaults instantiates a new RecommendationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationBaseWithDefaults() *RecommendationBase {
	this := RecommendationBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLastUpdatedTime returns the LastUpdatedTime field value if set, zero value otherwise.
func (o *RecommendationBase) GetLastUpdatedTime() time.Time {
	if o == nil || IsNil(o.LastUpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedTime
}

// GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBase) GetLastUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedTime) {
		return nil, false
	}
	return o.LastUpdatedTime, true
}

// HasLastUpdatedTime returns a boolean if a field has been set.
func (o *RecommendationBase) HasLastUpdatedTime() bool {
	if o != nil && !IsNil(o.LastUpdatedTime) {
		return true
	}

	return false
}

// SetLastUpdatedTime gets a reference to the given time.Time and assigns it to the LastUpdatedTime field.
func (o *RecommendationBase) SetLastUpdatedTime(v time.Time) {
	o.LastUpdatedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecommendationBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecommendationBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecommendationBase) SetName(v string) {
	o.Name = &v
}

// GetRequirementMet returns the RequirementMet field value if set, zero value otherwise.
func (o *RecommendationBase) GetRequirementMet() bool {
	if o == nil || IsNil(o.RequirementMet) {
		var ret bool
		return ret
	}
	return *o.RequirementMet
}

// GetRequirementMetOk returns a tuple with the RequirementMet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationBase) GetRequirementMetOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirementMet) {
		return nil, false
	}
	return o.RequirementMet, true
}

// HasRequirementMet returns a boolean if a field has been set.
func (o *RecommendationBase) HasRequirementMet() bool {
	if o != nil && !IsNil(o.RequirementMet) {
		return true
	}

	return false
}

// SetRequirementMet gets a reference to the given bool and assigns it to the RequirementMet field.
func (o *RecommendationBase) SetRequirementMet(v bool) {
	o.RequirementMet = &v
}

func (o RecommendationBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.LastUpdatedTime) {
		toSerialize["LastUpdatedTime"] = o.LastUpdatedTime
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.RequirementMet) {
		toSerialize["RequirementMet"] = o.RequirementMet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecommendationBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type RecommendationBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The time when the recommendation was last updated.
		LastUpdatedTime *time.Time `json:"LastUpdatedTime,omitempty"`
		// The name of the recommendation.
		Name *string `json:"Name,omitempty"`
		// Indicates if the recommendation requirement is met by the existing setup by adding hardware components to it or it needs expansion. For example if the recommendation is to add 16 disks to a HyperFlex cluster but the cluster is already alost full and only 8 disks can be added, then this property is set to false.
		RequirementMet *bool `json:"RequirementMet,omitempty"`
	}

	varRecommendationBaseWithoutEmbeddedStruct := RecommendationBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecommendationBaseWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationBase := _RecommendationBase{}
		varRecommendationBase.ClassId = varRecommendationBaseWithoutEmbeddedStruct.ClassId
		varRecommendationBase.ObjectType = varRecommendationBaseWithoutEmbeddedStruct.ObjectType
		varRecommendationBase.LastUpdatedTime = varRecommendationBaseWithoutEmbeddedStruct.LastUpdatedTime
		varRecommendationBase.Name = varRecommendationBaseWithoutEmbeddedStruct.Name
		varRecommendationBase.RequirementMet = varRecommendationBaseWithoutEmbeddedStruct.RequirementMet
		*o = RecommendationBase(varRecommendationBase)
	} else {
		return err
	}

	varRecommendationBase := _RecommendationBase{}

	err = json.Unmarshal(data, &varRecommendationBase)
	if err == nil {
		o.MoBaseMo = varRecommendationBase.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LastUpdatedTime")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "RequirementMet")

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

type NullableRecommendationBase struct {
	value *RecommendationBase
	isSet bool
}

func (v NullableRecommendationBase) Get() *RecommendationBase {
	return v.value
}

func (v *NullableRecommendationBase) Set(val *RecommendationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationBase(val *RecommendationBase) *NullableRecommendationBase {
	return &NullableRecommendationBase{value: val, isSet: true}
}

func (v NullableRecommendationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
