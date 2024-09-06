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

// checks if the WorkflowEnumEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowEnumEntry{}

// WorkflowEnumEntry Captures a single enum entry which has a label and value.
type WorkflowEnumEntry struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), round parenthesis ( () ), or an underscore (_) and must have an alphanumeric character.
	Label *string `json:"Label,omitempty" validate:"regexp=^[a-zA-Z0-9]+[+\\\\s\\/a-zA-Z0-9)(_'.:-]{0,92}$"`
	// Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), round parenthesis ( () ), forward slash (/), or an underscore (_).
	Value                *string `json:"Value,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]*[+\\\\s\\/a-zA-Z0-9)(_.:-]{1,64}$"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowEnumEntry WorkflowEnumEntry

// NewWorkflowEnumEntry instantiates a new WorkflowEnumEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEnumEntry(classId string, objectType string) *WorkflowEnumEntry {
	this := WorkflowEnumEntry{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowEnumEntryWithDefaults instantiates a new WorkflowEnumEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEnumEntryWithDefaults() *WorkflowEnumEntry {
	this := WorkflowEnumEntry{}
	var classId string = "workflow.EnumEntry"
	this.ClassId = classId
	var objectType string = "workflow.EnumEntry"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowEnumEntry) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowEnumEntry) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.EnumEntry" of the ClassId field.
func (o *WorkflowEnumEntry) GetDefaultClassId() interface{} {
	return "workflow.EnumEntry"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowEnumEntry) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowEnumEntry) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.EnumEntry" of the ObjectType field.
func (o *WorkflowEnumEntry) GetDefaultObjectType() interface{} {
	return "workflow.EnumEntry"
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowEnumEntry) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowEnumEntry) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowEnumEntry) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *WorkflowEnumEntry) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEnumEntry) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *WorkflowEnumEntry) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *WorkflowEnumEntry) SetValue(v string) {
	o.Value = &v
}

func (o WorkflowEnumEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowEnumEntry) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowEnumEntry) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowEnumEntryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), round parenthesis ( () ), or an underscore (_) and must have an alphanumeric character.
		Label *string `json:"Label,omitempty" validate:"regexp=^[a-zA-Z0-9]+[+\\\\s\\/a-zA-Z0-9)(_'.:-]{0,92}$"`
		// Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), round parenthesis ( () ), forward slash (/), or an underscore (_).
		Value *string `json:"Value,omitempty" validate:"regexp=^[a-zA-Z0-9_.:-]*[+\\\\s\\/a-zA-Z0-9)(_.:-]{1,64}$"`
	}

	varWorkflowEnumEntryWithoutEmbeddedStruct := WorkflowEnumEntryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowEnumEntryWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowEnumEntry := _WorkflowEnumEntry{}
		varWorkflowEnumEntry.ClassId = varWorkflowEnumEntryWithoutEmbeddedStruct.ClassId
		varWorkflowEnumEntry.ObjectType = varWorkflowEnumEntryWithoutEmbeddedStruct.ObjectType
		varWorkflowEnumEntry.Label = varWorkflowEnumEntryWithoutEmbeddedStruct.Label
		varWorkflowEnumEntry.Value = varWorkflowEnumEntryWithoutEmbeddedStruct.Value
		*o = WorkflowEnumEntry(varWorkflowEnumEntry)
	} else {
		return err
	}

	varWorkflowEnumEntry := _WorkflowEnumEntry{}

	err = json.Unmarshal(data, &varWorkflowEnumEntry)
	if err == nil {
		o.MoBaseComplexType = varWorkflowEnumEntry.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Value")

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

type NullableWorkflowEnumEntry struct {
	value *WorkflowEnumEntry
	isSet bool
}

func (v NullableWorkflowEnumEntry) Get() *WorkflowEnumEntry {
	return v.value
}

func (v *NullableWorkflowEnumEntry) Set(val *WorkflowEnumEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEnumEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEnumEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEnumEntry(val *WorkflowEnumEntry) *NullableWorkflowEnumEntry {
	return &NullableWorkflowEnumEntry{value: val, isSet: true}
}

func (v NullableWorkflowEnumEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEnumEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
