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
)

// checks if the HyperflexReplicationClusterReferenceToSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexReplicationClusterReferenceToSchedule{}

// HyperflexReplicationClusterReferenceToSchedule A map from Replication Cluster EntityReference to the schedule.
type HyperflexReplicationClusterReferenceToSchedule struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                   string                               `json:"ObjectType"`
	Schedule                     NullableHyperflexReplicationSchedule `json:"Schedule,omitempty"`
	TargetClusterEntityReference NullableHyperflexEntityReference     `json:"TargetClusterEntityReference,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _HyperflexReplicationClusterReferenceToSchedule HyperflexReplicationClusterReferenceToSchedule

// NewHyperflexReplicationClusterReferenceToSchedule instantiates a new HyperflexReplicationClusterReferenceToSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexReplicationClusterReferenceToSchedule(classId string, objectType string) *HyperflexReplicationClusterReferenceToSchedule {
	this := HyperflexReplicationClusterReferenceToSchedule{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexReplicationClusterReferenceToScheduleWithDefaults instantiates a new HyperflexReplicationClusterReferenceToSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexReplicationClusterReferenceToScheduleWithDefaults() *HyperflexReplicationClusterReferenceToSchedule {
	this := HyperflexReplicationClusterReferenceToSchedule{}
	var classId string = "hyperflex.ReplicationClusterReferenceToSchedule"
	this.ClassId = classId
	var objectType string = "hyperflex.ReplicationClusterReferenceToSchedule"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexReplicationClusterReferenceToSchedule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationClusterReferenceToSchedule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexReplicationClusterReferenceToSchedule) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ReplicationClusterReferenceToSchedule" of the ClassId field.
func (o *HyperflexReplicationClusterReferenceToSchedule) GetDefaultClassId() interface{} {
	return "hyperflex.ReplicationClusterReferenceToSchedule"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexReplicationClusterReferenceToSchedule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexReplicationClusterReferenceToSchedule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexReplicationClusterReferenceToSchedule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ReplicationClusterReferenceToSchedule" of the ObjectType field.
func (o *HyperflexReplicationClusterReferenceToSchedule) GetDefaultObjectType() interface{} {
	return "hyperflex.ReplicationClusterReferenceToSchedule"
}

// GetSchedule returns the Schedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationClusterReferenceToSchedule) GetSchedule() HyperflexReplicationSchedule {
	if o == nil || IsNil(o.Schedule.Get()) {
		var ret HyperflexReplicationSchedule
		return ret
	}
	return *o.Schedule.Get()
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationClusterReferenceToSchedule) GetScheduleOk() (*HyperflexReplicationSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schedule.Get(), o.Schedule.IsSet()
}

// HasSchedule returns a boolean if a field has been set.
func (o *HyperflexReplicationClusterReferenceToSchedule) HasSchedule() bool {
	if o != nil && o.Schedule.IsSet() {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given NullableHyperflexReplicationSchedule and assigns it to the Schedule field.
func (o *HyperflexReplicationClusterReferenceToSchedule) SetSchedule(v HyperflexReplicationSchedule) {
	o.Schedule.Set(&v)
}

// SetScheduleNil sets the value for Schedule to be an explicit nil
func (o *HyperflexReplicationClusterReferenceToSchedule) SetScheduleNil() {
	o.Schedule.Set(nil)
}

// UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
func (o *HyperflexReplicationClusterReferenceToSchedule) UnsetSchedule() {
	o.Schedule.Unset()
}

// GetTargetClusterEntityReference returns the TargetClusterEntityReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexReplicationClusterReferenceToSchedule) GetTargetClusterEntityReference() HyperflexEntityReference {
	if o == nil || IsNil(o.TargetClusterEntityReference.Get()) {
		var ret HyperflexEntityReference
		return ret
	}
	return *o.TargetClusterEntityReference.Get()
}

// GetTargetClusterEntityReferenceOk returns a tuple with the TargetClusterEntityReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexReplicationClusterReferenceToSchedule) GetTargetClusterEntityReferenceOk() (*HyperflexEntityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetClusterEntityReference.Get(), o.TargetClusterEntityReference.IsSet()
}

// HasTargetClusterEntityReference returns a boolean if a field has been set.
func (o *HyperflexReplicationClusterReferenceToSchedule) HasTargetClusterEntityReference() bool {
	if o != nil && o.TargetClusterEntityReference.IsSet() {
		return true
	}

	return false
}

// SetTargetClusterEntityReference gets a reference to the given NullableHyperflexEntityReference and assigns it to the TargetClusterEntityReference field.
func (o *HyperflexReplicationClusterReferenceToSchedule) SetTargetClusterEntityReference(v HyperflexEntityReference) {
	o.TargetClusterEntityReference.Set(&v)
}

// SetTargetClusterEntityReferenceNil sets the value for TargetClusterEntityReference to be an explicit nil
func (o *HyperflexReplicationClusterReferenceToSchedule) SetTargetClusterEntityReferenceNil() {
	o.TargetClusterEntityReference.Set(nil)
}

// UnsetTargetClusterEntityReference ensures that no value is present for TargetClusterEntityReference, not even an explicit nil
func (o *HyperflexReplicationClusterReferenceToSchedule) UnsetTargetClusterEntityReference() {
	o.TargetClusterEntityReference.Unset()
}

func (o HyperflexReplicationClusterReferenceToSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexReplicationClusterReferenceToSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Schedule.IsSet() {
		toSerialize["Schedule"] = o.Schedule.Get()
	}
	if o.TargetClusterEntityReference.IsSet() {
		toSerialize["TargetClusterEntityReference"] = o.TargetClusterEntityReference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexReplicationClusterReferenceToSchedule) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType                   string                               `json:"ObjectType"`
		Schedule                     NullableHyperflexReplicationSchedule `json:"Schedule,omitempty"`
		TargetClusterEntityReference NullableHyperflexEntityReference     `json:"TargetClusterEntityReference,omitempty"`
	}

	varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct := HyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexReplicationClusterReferenceToSchedule := _HyperflexReplicationClusterReferenceToSchedule{}
		varHyperflexReplicationClusterReferenceToSchedule.ClassId = varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct.ClassId
		varHyperflexReplicationClusterReferenceToSchedule.ObjectType = varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct.ObjectType
		varHyperflexReplicationClusterReferenceToSchedule.Schedule = varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct.Schedule
		varHyperflexReplicationClusterReferenceToSchedule.TargetClusterEntityReference = varHyperflexReplicationClusterReferenceToScheduleWithoutEmbeddedStruct.TargetClusterEntityReference
		*o = HyperflexReplicationClusterReferenceToSchedule(varHyperflexReplicationClusterReferenceToSchedule)
	} else {
		return err
	}

	varHyperflexReplicationClusterReferenceToSchedule := _HyperflexReplicationClusterReferenceToSchedule{}

	err = json.Unmarshal(data, &varHyperflexReplicationClusterReferenceToSchedule)
	if err == nil {
		o.MoBaseComplexType = varHyperflexReplicationClusterReferenceToSchedule.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Schedule")
		delete(additionalProperties, "TargetClusterEntityReference")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableHyperflexReplicationClusterReferenceToSchedule struct {
	value *HyperflexReplicationClusterReferenceToSchedule
	isSet bool
}

func (v NullableHyperflexReplicationClusterReferenceToSchedule) Get() *HyperflexReplicationClusterReferenceToSchedule {
	return v.value
}

func (v *NullableHyperflexReplicationClusterReferenceToSchedule) Set(val *HyperflexReplicationClusterReferenceToSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexReplicationClusterReferenceToSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexReplicationClusterReferenceToSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexReplicationClusterReferenceToSchedule(val *HyperflexReplicationClusterReferenceToSchedule) *NullableHyperflexReplicationClusterReferenceToSchedule {
	return &NullableHyperflexReplicationClusterReferenceToSchedule{value: val, isSet: true}
}

func (v NullableHyperflexReplicationClusterReferenceToSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexReplicationClusterReferenceToSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
