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

// checks if the OpenapiFailedTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenapiFailedTask{}

// OpenapiFailedTask A type that holds information about failed tasks.
type OpenapiFailedTask struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the task.
	Name *string `json:"Name,omitempty"`
	// REST API path of the task.
	Path *string `json:"Path,omitempty"`
	// Indicates the reason for task creation failure.
	Reason               *string `json:"Reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenapiFailedTask OpenapiFailedTask

// NewOpenapiFailedTask instantiates a new OpenapiFailedTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenapiFailedTask(classId string, objectType string) *OpenapiFailedTask {
	this := OpenapiFailedTask{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewOpenapiFailedTaskWithDefaults instantiates a new OpenapiFailedTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenapiFailedTaskWithDefaults() *OpenapiFailedTask {
	this := OpenapiFailedTask{}
	var classId string = "openapi.FailedTask"
	this.ClassId = classId
	var objectType string = "openapi.FailedTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OpenapiFailedTask) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTask) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OpenapiFailedTask) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "openapi.FailedTask" of the ClassId field.
func (o *OpenapiFailedTask) GetDefaultClassId() interface{} {
	return "openapi.FailedTask"
}

// GetObjectType returns the ObjectType field value
func (o *OpenapiFailedTask) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTask) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OpenapiFailedTask) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "openapi.FailedTask" of the ObjectType field.
func (o *OpenapiFailedTask) GetDefaultObjectType() interface{} {
	return "openapi.FailedTask"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OpenapiFailedTask) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTask) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OpenapiFailedTask) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OpenapiFailedTask) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OpenapiFailedTask) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTask) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OpenapiFailedTask) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OpenapiFailedTask) SetPath(v string) {
	o.Path = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *OpenapiFailedTask) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenapiFailedTask) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *OpenapiFailedTask) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *OpenapiFailedTask) SetReason(v string) {
	o.Reason = &v
}

func (o OpenapiFailedTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenapiFailedTask) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenapiFailedTask) UnmarshalJSON(data []byte) (err error) {
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
	type OpenapiFailedTaskWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the task.
		Name *string `json:"Name,omitempty"`
		// REST API path of the task.
		Path *string `json:"Path,omitempty"`
		// Indicates the reason for task creation failure.
		Reason *string `json:"Reason,omitempty"`
	}

	varOpenapiFailedTaskWithoutEmbeddedStruct := OpenapiFailedTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOpenapiFailedTaskWithoutEmbeddedStruct)
	if err == nil {
		varOpenapiFailedTask := _OpenapiFailedTask{}
		varOpenapiFailedTask.ClassId = varOpenapiFailedTaskWithoutEmbeddedStruct.ClassId
		varOpenapiFailedTask.ObjectType = varOpenapiFailedTaskWithoutEmbeddedStruct.ObjectType
		varOpenapiFailedTask.Name = varOpenapiFailedTaskWithoutEmbeddedStruct.Name
		varOpenapiFailedTask.Path = varOpenapiFailedTaskWithoutEmbeddedStruct.Path
		varOpenapiFailedTask.Reason = varOpenapiFailedTaskWithoutEmbeddedStruct.Reason
		*o = OpenapiFailedTask(varOpenapiFailedTask)
	} else {
		return err
	}

	varOpenapiFailedTask := _OpenapiFailedTask{}

	err = json.Unmarshal(data, &varOpenapiFailedTask)
	if err == nil {
		o.MoBaseComplexType = varOpenapiFailedTask.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Reason")

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

type NullableOpenapiFailedTask struct {
	value *OpenapiFailedTask
	isSet bool
}

func (v NullableOpenapiFailedTask) Get() *OpenapiFailedTask {
	return v.value
}

func (v *NullableOpenapiFailedTask) Set(val *OpenapiFailedTask) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenapiFailedTask) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenapiFailedTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenapiFailedTask(val *OpenapiFailedTask) *NullableOpenapiFailedTask {
	return &NullableOpenapiFailedTask{value: val, isSet: true}
}

func (v NullableOpenapiFailedTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenapiFailedTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
