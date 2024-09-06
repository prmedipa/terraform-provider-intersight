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

// checks if the PolicyAbstractProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAbstractProfile{}

// PolicyAbstractProfile Abstract base type for all profiles.
type PolicyAbstractProfile struct {
	PolicyAbstractConfigurationObject
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Description of the profile.
	Description *string `json:"Description,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]+[\\\\x00-\\\\xFF]*$"`
	// Name of the profile instance or profile template.
	Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
	// Defines the type of the profile. Accepted values are instance or template. * `instance` - The profile defines the configuration for a specific instance of a target.
	Type                 *string                                   `json:"Type,omitempty"`
	SrcTemplate          NullablePolicyAbstractProfileRelationship `json:"SrcTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAbstractProfile PolicyAbstractProfile

// NewPolicyAbstractProfile instantiates a new PolicyAbstractProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAbstractProfile(classId string, objectType string) *PolicyAbstractProfile {
	this := PolicyAbstractProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// NewPolicyAbstractProfileWithDefaults instantiates a new PolicyAbstractProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAbstractProfileWithDefaults() *PolicyAbstractProfile {
	this := PolicyAbstractProfile{}
	var type_ string = "instance"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyAbstractProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyAbstractProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyAbstractProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyAbstractProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyAbstractProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyAbstractProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyAbstractProfile) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyAbstractProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyAbstractProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyAbstractProfile) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyAbstractProfile) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAbstractProfile) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyAbstractProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyAbstractProfile) SetType(v string) {
	o.Type = &v
}

// GetSrcTemplate returns the SrcTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAbstractProfile) GetSrcTemplate() PolicyAbstractProfileRelationship {
	if o == nil || IsNil(o.SrcTemplate.Get()) {
		var ret PolicyAbstractProfileRelationship
		return ret
	}
	return *o.SrcTemplate.Get()
}

// GetSrcTemplateOk returns a tuple with the SrcTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAbstractProfile) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SrcTemplate.Get(), o.SrcTemplate.IsSet()
}

// HasSrcTemplate returns a boolean if a field has been set.
func (o *PolicyAbstractProfile) HasSrcTemplate() bool {
	if o != nil && o.SrcTemplate.IsSet() {
		return true
	}

	return false
}

// SetSrcTemplate gets a reference to the given NullablePolicyAbstractProfileRelationship and assigns it to the SrcTemplate field.
func (o *PolicyAbstractProfile) SetSrcTemplate(v PolicyAbstractProfileRelationship) {
	o.SrcTemplate.Set(&v)
}

// SetSrcTemplateNil sets the value for SrcTemplate to be an explicit nil
func (o *PolicyAbstractProfile) SetSrcTemplateNil() {
	o.SrcTemplate.Set(nil)
}

// UnsetSrcTemplate ensures that no value is present for SrcTemplate, not even an explicit nil
func (o *PolicyAbstractProfile) UnsetSrcTemplate() {
	o.SrcTemplate.Unset()
}

func (o PolicyAbstractProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAbstractProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractConfigurationObject, errPolicyAbstractConfigurationObject := json.Marshal(o.PolicyAbstractConfigurationObject)
	if errPolicyAbstractConfigurationObject != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigurationObject
	}
	errPolicyAbstractConfigurationObject = json.Unmarshal([]byte(serializedPolicyAbstractConfigurationObject), &toSerialize)
	if errPolicyAbstractConfigurationObject != nil {
		return map[string]interface{}{}, errPolicyAbstractConfigurationObject
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.SrcTemplate.IsSet() {
		toSerialize["SrcTemplate"] = o.SrcTemplate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAbstractProfile) UnmarshalJSON(data []byte) (err error) {
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
	type PolicyAbstractProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Description of the profile.
		Description *string `json:"Description,omitempty" validate:"regexp=^$|^[a-zA-Z0-9]+[\\\\x00-\\\\xFF]*$"`
		// Name of the profile instance or profile template.
		Name *string `json:"Name,omitempty" validate:"regexp=^[a-zA-Z0-9_.-]{1,64}$"`
		// Defines the type of the profile. Accepted values are instance or template. * `instance` - The profile defines the configuration for a specific instance of a target.
		Type        *string                                   `json:"Type,omitempty"`
		SrcTemplate NullablePolicyAbstractProfileRelationship `json:"SrcTemplate,omitempty"`
	}

	varPolicyAbstractProfileWithoutEmbeddedStruct := PolicyAbstractProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPolicyAbstractProfileWithoutEmbeddedStruct)
	if err == nil {
		varPolicyAbstractProfile := _PolicyAbstractProfile{}
		varPolicyAbstractProfile.ClassId = varPolicyAbstractProfileWithoutEmbeddedStruct.ClassId
		varPolicyAbstractProfile.ObjectType = varPolicyAbstractProfileWithoutEmbeddedStruct.ObjectType
		varPolicyAbstractProfile.Description = varPolicyAbstractProfileWithoutEmbeddedStruct.Description
		varPolicyAbstractProfile.Name = varPolicyAbstractProfileWithoutEmbeddedStruct.Name
		varPolicyAbstractProfile.Type = varPolicyAbstractProfileWithoutEmbeddedStruct.Type
		varPolicyAbstractProfile.SrcTemplate = varPolicyAbstractProfileWithoutEmbeddedStruct.SrcTemplate
		*o = PolicyAbstractProfile(varPolicyAbstractProfile)
	} else {
		return err
	}

	varPolicyAbstractProfile := _PolicyAbstractProfile{}

	err = json.Unmarshal(data, &varPolicyAbstractProfile)
	if err == nil {
		o.PolicyAbstractConfigurationObject = varPolicyAbstractProfile.PolicyAbstractConfigurationObject
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "SrcTemplate")

		// remove fields from embedded structs
		reflectPolicyAbstractConfigurationObject := reflect.ValueOf(o.PolicyAbstractConfigurationObject)
		for i := 0; i < reflectPolicyAbstractConfigurationObject.Type().NumField(); i++ {
			t := reflectPolicyAbstractConfigurationObject.Type().Field(i)

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

type NullablePolicyAbstractProfile struct {
	value *PolicyAbstractProfile
	isSet bool
}

func (v NullablePolicyAbstractProfile) Get() *PolicyAbstractProfile {
	return v.value
}

func (v *NullablePolicyAbstractProfile) Set(val *PolicyAbstractProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAbstractProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAbstractProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAbstractProfile(val *PolicyAbstractProfile) *NullablePolicyAbstractProfile {
	return &NullablePolicyAbstractProfile{value: val, isSet: true}
}

func (v NullablePolicyAbstractProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAbstractProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
