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

// checks if the RecoveryConfigResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryConfigResult{}

// RecoveryConfigResult Profile configuration (deploy, validation) results with the overall state and detailed result messages.
type RecoveryConfigResult struct {
	PolicyAbstractConfigResult
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string                                    `json:"ObjectType"`
	BackupProfile NullableRecoveryBackupProfileRelationship `json:"BackupProfile,omitempty"`
	// An array of relationships to recoveryConfigResultEntry resources.
	ResultEntries        []RecoveryConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryConfigResult RecoveryConfigResult

// NewRecoveryConfigResult instantiates a new RecoveryConfigResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryConfigResult(classId string, objectType string) *RecoveryConfigResult {
	this := RecoveryConfigResult{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecoveryConfigResultWithDefaults instantiates a new RecoveryConfigResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryConfigResultWithDefaults() *RecoveryConfigResult {
	this := RecoveryConfigResult{}
	var classId string = "recovery.ConfigResult"
	this.ClassId = classId
	var objectType string = "recovery.ConfigResult"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryConfigResult) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryConfigResult) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryConfigResult) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "recovery.ConfigResult" of the ClassId field.
func (o *RecoveryConfigResult) GetDefaultClassId() interface{} {
	return "recovery.ConfigResult"
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryConfigResult) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryConfigResult) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryConfigResult) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "recovery.ConfigResult" of the ObjectType field.
func (o *RecoveryConfigResult) GetDefaultObjectType() interface{} {
	return "recovery.ConfigResult"
}

// GetBackupProfile returns the BackupProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryConfigResult) GetBackupProfile() RecoveryBackupProfileRelationship {
	if o == nil || IsNil(o.BackupProfile.Get()) {
		var ret RecoveryBackupProfileRelationship
		return ret
	}
	return *o.BackupProfile.Get()
}

// GetBackupProfileOk returns a tuple with the BackupProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryConfigResult) GetBackupProfileOk() (*RecoveryBackupProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupProfile.Get(), o.BackupProfile.IsSet()
}

// HasBackupProfile returns a boolean if a field has been set.
func (o *RecoveryConfigResult) HasBackupProfile() bool {
	if o != nil && o.BackupProfile.IsSet() {
		return true
	}

	return false
}

// SetBackupProfile gets a reference to the given NullableRecoveryBackupProfileRelationship and assigns it to the BackupProfile field.
func (o *RecoveryConfigResult) SetBackupProfile(v RecoveryBackupProfileRelationship) {
	o.BackupProfile.Set(&v)
}

// SetBackupProfileNil sets the value for BackupProfile to be an explicit nil
func (o *RecoveryConfigResult) SetBackupProfileNil() {
	o.BackupProfile.Set(nil)
}

// UnsetBackupProfile ensures that no value is present for BackupProfile, not even an explicit nil
func (o *RecoveryConfigResult) UnsetBackupProfile() {
	o.BackupProfile.Unset()
}

// GetResultEntries returns the ResultEntries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecoveryConfigResult) GetResultEntries() []RecoveryConfigResultEntryRelationship {
	if o == nil {
		var ret []RecoveryConfigResultEntryRelationship
		return ret
	}
	return o.ResultEntries
}

// GetResultEntriesOk returns a tuple with the ResultEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecoveryConfigResult) GetResultEntriesOk() ([]RecoveryConfigResultEntryRelationship, bool) {
	if o == nil || IsNil(o.ResultEntries) {
		return nil, false
	}
	return o.ResultEntries, true
}

// HasResultEntries returns a boolean if a field has been set.
func (o *RecoveryConfigResult) HasResultEntries() bool {
	if o != nil && !IsNil(o.ResultEntries) {
		return true
	}

	return false
}

// SetResultEntries gets a reference to the given []RecoveryConfigResultEntryRelationship and assigns it to the ResultEntries field.
func (o *RecoveryConfigResult) SetResultEntries(v []RecoveryConfigResultEntryRelationship) {
	o.ResultEntries = v
}

func (o RecoveryConfigResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryConfigResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigResult, errPolicyAbstractConfigResult := json.Marshal(o.PolicyAbstractConfigResult)
	if errPolicyAbstractConfigResult != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResult
	}
	errPolicyAbstractConfigResult = json.Unmarshal([]byte(serializedPolicyAbstractConfigResult), &toSerialize)
	if errPolicyAbstractConfigResult != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigResult
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.BackupProfile.IsSet() {
		toSerialize["BackupProfile"] = o.BackupProfile.Get()
	}
	if o.ResultEntries != nil {
		toSerialize["ResultEntries"] = o.ResultEntries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryConfigResult) UnmarshalJSON(data []byte) (err error) {
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
	type RecoveryConfigResultWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType    string                                    `json:"ObjectType"`
		BackupProfile NullableRecoveryBackupProfileRelationship `json:"BackupProfile,omitempty"`
		// An array of relationships to recoveryConfigResultEntry resources.
		ResultEntries []RecoveryConfigResultEntryRelationship `json:"ResultEntries,omitempty"`
	}

	varRecoveryConfigResultWithoutEmbeddedStruct := RecoveryConfigResultWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecoveryConfigResultWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryConfigResult := _RecoveryConfigResult{}
		varRecoveryConfigResult.ClassId = varRecoveryConfigResultWithoutEmbeddedStruct.ClassId
		varRecoveryConfigResult.ObjectType = varRecoveryConfigResultWithoutEmbeddedStruct.ObjectType
		varRecoveryConfigResult.BackupProfile = varRecoveryConfigResultWithoutEmbeddedStruct.BackupProfile
		varRecoveryConfigResult.ResultEntries = varRecoveryConfigResultWithoutEmbeddedStruct.ResultEntries
		*o = RecoveryConfigResult(varRecoveryConfigResult)
	} else {
		return err
	}

	varRecoveryConfigResult := _RecoveryConfigResult{}

	err = json.Unmarshal(data, &varRecoveryConfigResult)
	if err == nil {
		o.PolicyAbstractConfigResult = varRecoveryConfigResult.PolicyAbstractConfigResult
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BackupProfile")
		delete(additionalProperties, "ResultEntries")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigResult := reflect.ValueOf(o.PolicyAbstractConfigResult)
		for i := 0; i < reflectPolicyAbstractConfigResult.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigResult.Type().Field(i)

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

type NullableRecoveryConfigResult struct {
	value *RecoveryConfigResult
	isSet bool
}

func (v NullableRecoveryConfigResult) Get() *RecoveryConfigResult {
	return v.value
}

func (v *NullableRecoveryConfigResult) Set(val *RecoveryConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryConfigResult(val *RecoveryConfigResult) *NullableRecoveryConfigResult {
	return &NullableRecoveryConfigResult{value: val, isSet: true}
}

func (v NullableRecoveryConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
