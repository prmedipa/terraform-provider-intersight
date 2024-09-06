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

// checks if the PolicyConfigResultContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyConfigResultContext{}

// PolicyConfigResultContext The context for a validation/config result such as the related entity's name, type, MOID etc.
type PolicyConfigResultContext struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The data of the object present in config result context.
	EntityData interface{} `json:"EntityData,omitempty"`
	// The Moid of the object present in config result context.
	EntityMoid *string `json:"EntityMoid,omitempty"`
	// The name of the object present in config result context.
	EntityName *string `json:"EntityName,omitempty"`
	// The type of the object present in config result context.
	EntityType *string `json:"EntityType,omitempty"`
	// The Moid of the parent object present in config result context.
	ParentMoid *string `json:"ParentMoid,omitempty"`
	// The type of the parent object present in config result context.
	ParentType           *string `json:"ParentType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyConfigResultContext PolicyConfigResultContext

// NewPolicyConfigResultContext instantiates a new PolicyConfigResultContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyConfigResultContext(classId string, objectType string) *PolicyConfigResultContext {
	this := PolicyConfigResultContext{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewPolicyConfigResultContextWithDefaults instantiates a new PolicyConfigResultContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyConfigResultContextWithDefaults() *PolicyConfigResultContext {
	this := PolicyConfigResultContext{}
	var classId string = "policy.ConfigResultContext"
	this.ClassId = classId
	var objectType string = "policy.ConfigResultContext"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyConfigResultContext) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyConfigResultContext) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "policy.ConfigResultContext" of the ClassId field.
func (o *PolicyConfigResultContext) GetDefaultClassId() interface{} {
	return "policy.ConfigResultContext"
}

// GetObjectType returns the ObjectType field value
func (o *PolicyConfigResultContext) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyConfigResultContext) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "policy.ConfigResultContext" of the ObjectType field.
func (o *PolicyConfigResultContext) GetDefaultObjectType() interface{} {
	return "policy.ConfigResultContext"
}

// GetEntityData returns the EntityData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyConfigResultContext) GetEntityData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EntityData
}

// GetEntityDataOk returns a tuple with the EntityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyConfigResultContext) GetEntityDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EntityData) {
		return nil, false
	}
	return &o.EntityData, true
}

// HasEntityData returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityData() bool {
	if o != nil && !IsNil(o.EntityData) {
		return true
	}

	return false
}

// SetEntityData gets a reference to the given interface{} and assigns it to the EntityData field.
func (o *PolicyConfigResultContext) SetEntityData(v interface{}) {
	o.EntityData = v
}

// GetEntityMoid returns the EntityMoid field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityMoid() string {
	if o == nil || IsNil(o.EntityMoid) {
		var ret string
		return ret
	}
	return *o.EntityMoid
}

// GetEntityMoidOk returns a tuple with the EntityMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityMoidOk() (*string, bool) {
	if o == nil || IsNil(o.EntityMoid) {
		return nil, false
	}
	return o.EntityMoid, true
}

// HasEntityMoid returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityMoid() bool {
	if o != nil && !IsNil(o.EntityMoid) {
		return true
	}

	return false
}

// SetEntityMoid gets a reference to the given string and assigns it to the EntityMoid field.
func (o *PolicyConfigResultContext) SetEntityMoid(v string) {
	o.EntityMoid = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityName() string {
	if o == nil || IsNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityName() bool {
	if o != nil && !IsNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *PolicyConfigResultContext) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *PolicyConfigResultContext) SetEntityType(v string) {
	o.EntityType = &v
}

// GetParentMoid returns the ParentMoid field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetParentMoid() string {
	if o == nil || IsNil(o.ParentMoid) {
		var ret string
		return ret
	}
	return *o.ParentMoid
}

// GetParentMoidOk returns a tuple with the ParentMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetParentMoidOk() (*string, bool) {
	if o == nil || IsNil(o.ParentMoid) {
		return nil, false
	}
	return o.ParentMoid, true
}

// HasParentMoid returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasParentMoid() bool {
	if o != nil && !IsNil(o.ParentMoid) {
		return true
	}

	return false
}

// SetParentMoid gets a reference to the given string and assigns it to the ParentMoid field.
func (o *PolicyConfigResultContext) SetParentMoid(v string) {
	o.ParentMoid = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *PolicyConfigResultContext) GetParentType() string {
	if o == nil || IsNil(o.ParentType) {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyConfigResultContext) GetParentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentType) {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *PolicyConfigResultContext) HasParentType() bool {
	if o != nil && !IsNil(o.ParentType) {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *PolicyConfigResultContext) SetParentType(v string) {
	o.ParentType = &v
}

func (o PolicyConfigResultContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyConfigResultContext) ToMap() (map[string]interface{}, error) {
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
	if o.EntityData != nil {
		toSerialize["EntityData"] = o.EntityData
	}
	if !IsNil(o.EntityMoid) {
		toSerialize["EntityMoid"] = o.EntityMoid
	}
	if !IsNil(o.EntityName) {
		toSerialize["EntityName"] = o.EntityName
	}
	if !IsNil(o.EntityType) {
		toSerialize["EntityType"] = o.EntityType
	}
	if !IsNil(o.ParentMoid) {
		toSerialize["ParentMoid"] = o.ParentMoid
	}
	if !IsNil(o.ParentType) {
		toSerialize["ParentType"] = o.ParentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyConfigResultContext) UnmarshalJSON(data []byte) (err error) {
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
	type PolicyConfigResultContextWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The data of the object present in config result context.
		EntityData interface{} `json:"EntityData,omitempty"`
		// The Moid of the object present in config result context.
		EntityMoid *string `json:"EntityMoid,omitempty"`
		// The name of the object present in config result context.
		EntityName *string `json:"EntityName,omitempty"`
		// The type of the object present in config result context.
		EntityType *string `json:"EntityType,omitempty"`
		// The Moid of the parent object present in config result context.
		ParentMoid *string `json:"ParentMoid,omitempty"`
		// The type of the parent object present in config result context.
		ParentType *string `json:"ParentType,omitempty"`
	}

	varPolicyConfigResultContextWithoutEmbeddedStruct := PolicyConfigResultContextWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varPolicyConfigResultContextWithoutEmbeddedStruct)
	if err == nil {
		varPolicyConfigResultContext := _PolicyConfigResultContext{}
		varPolicyConfigResultContext.ClassId = varPolicyConfigResultContextWithoutEmbeddedStruct.ClassId
		varPolicyConfigResultContext.ObjectType = varPolicyConfigResultContextWithoutEmbeddedStruct.ObjectType
		varPolicyConfigResultContext.EntityData = varPolicyConfigResultContextWithoutEmbeddedStruct.EntityData
		varPolicyConfigResultContext.EntityMoid = varPolicyConfigResultContextWithoutEmbeddedStruct.EntityMoid
		varPolicyConfigResultContext.EntityName = varPolicyConfigResultContextWithoutEmbeddedStruct.EntityName
		varPolicyConfigResultContext.EntityType = varPolicyConfigResultContextWithoutEmbeddedStruct.EntityType
		varPolicyConfigResultContext.ParentMoid = varPolicyConfigResultContextWithoutEmbeddedStruct.ParentMoid
		varPolicyConfigResultContext.ParentType = varPolicyConfigResultContextWithoutEmbeddedStruct.ParentType
		*o = PolicyConfigResultContext(varPolicyConfigResultContext)
	} else {
		return err
	}

	varPolicyConfigResultContext := _PolicyConfigResultContext{}

	err = json.Unmarshal(data, &varPolicyConfigResultContext)
	if err == nil {
		o.MoBaseComplexType = varPolicyConfigResultContext.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EntityData")
		delete(additionalProperties, "EntityMoid")
		delete(additionalProperties, "EntityName")
		delete(additionalProperties, "EntityType")
		delete(additionalProperties, "ParentMoid")
		delete(additionalProperties, "ParentType")

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

type NullablePolicyConfigResultContext struct {
	value *PolicyConfigResultContext
	isSet bool
}

func (v NullablePolicyConfigResultContext) Get() *PolicyConfigResultContext {
	return v.value
}

func (v *NullablePolicyConfigResultContext) Set(val *PolicyConfigResultContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyConfigResultContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyConfigResultContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyConfigResultContext(val *PolicyConfigResultContext) *NullablePolicyConfigResultContext {
	return &NullablePolicyConfigResultContext{value: val, isSet: true}
}

func (v NullablePolicyConfigResultContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyConfigResultContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
