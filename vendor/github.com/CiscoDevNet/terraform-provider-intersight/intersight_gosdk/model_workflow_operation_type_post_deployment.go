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

// checks if the WorkflowOperationTypePostDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowOperationTypePostDeployment{}

// WorkflowOperationTypePostDeployment Operation details for post-deployment actions.
type WorkflowOperationTypePostDeployment struct {
	WorkflowBaseOperation
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType                  string   `json:"ObjectType"`
	ServiceItemActionDefinition *MoMoRef `json:"ServiceItemActionDefinition,omitempty"`
	ServiceItemInstance         *MoMoRef `json:"ServiceItemInstance,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _WorkflowOperationTypePostDeployment WorkflowOperationTypePostDeployment

// NewWorkflowOperationTypePostDeployment instantiates a new WorkflowOperationTypePostDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowOperationTypePostDeployment(classId string, objectType string) *WorkflowOperationTypePostDeployment {
	this := WorkflowOperationTypePostDeployment{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	return &this
}

// NewWorkflowOperationTypePostDeploymentWithDefaults instantiates a new WorkflowOperationTypePostDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowOperationTypePostDeploymentWithDefaults() *WorkflowOperationTypePostDeployment {
	this := WorkflowOperationTypePostDeployment{}
	var classId string = "workflow.OperationTypePostDeployment"
	this.ClassId = classId
	var objectType string = "workflow.OperationTypePostDeployment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowOperationTypePostDeployment) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypePostDeployment) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowOperationTypePostDeployment) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "workflow.OperationTypePostDeployment" of the ClassId field.
func (o *WorkflowOperationTypePostDeployment) GetDefaultClassId() interface{} {
	return "workflow.OperationTypePostDeployment"
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowOperationTypePostDeployment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypePostDeployment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowOperationTypePostDeployment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "workflow.OperationTypePostDeployment" of the ObjectType field.
func (o *WorkflowOperationTypePostDeployment) GetDefaultObjectType() interface{} {
	return "workflow.OperationTypePostDeployment"
}

// GetServiceItemActionDefinition returns the ServiceItemActionDefinition field value if set, zero value otherwise.
func (o *WorkflowOperationTypePostDeployment) GetServiceItemActionDefinition() MoMoRef {
	if o == nil || IsNil(o.ServiceItemActionDefinition) {
		var ret MoMoRef
		return ret
	}
	return *o.ServiceItemActionDefinition
}

// GetServiceItemActionDefinitionOk returns a tuple with the ServiceItemActionDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypePostDeployment) GetServiceItemActionDefinitionOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.ServiceItemActionDefinition) {
		return nil, false
	}
	return o.ServiceItemActionDefinition, true
}

// HasServiceItemActionDefinition returns a boolean if a field has been set.
func (o *WorkflowOperationTypePostDeployment) HasServiceItemActionDefinition() bool {
	if o != nil && !IsNil(o.ServiceItemActionDefinition) {
		return true
	}

	return false
}

// SetServiceItemActionDefinition gets a reference to the given MoMoRef and assigns it to the ServiceItemActionDefinition field.
func (o *WorkflowOperationTypePostDeployment) SetServiceItemActionDefinition(v MoMoRef) {
	o.ServiceItemActionDefinition = &v
}

// GetServiceItemInstance returns the ServiceItemInstance field value if set, zero value otherwise.
func (o *WorkflowOperationTypePostDeployment) GetServiceItemInstance() MoMoRef {
	if o == nil || IsNil(o.ServiceItemInstance) {
		var ret MoMoRef
		return ret
	}
	return *o.ServiceItemInstance
}

// GetServiceItemInstanceOk returns a tuple with the ServiceItemInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowOperationTypePostDeployment) GetServiceItemInstanceOk() (*MoMoRef, bool) {
	if o == nil || IsNil(o.ServiceItemInstance) {
		return nil, false
	}
	return o.ServiceItemInstance, true
}

// HasServiceItemInstance returns a boolean if a field has been set.
func (o *WorkflowOperationTypePostDeployment) HasServiceItemInstance() bool {
	if o != nil && !IsNil(o.ServiceItemInstance) {
		return true
	}

	return false
}

// SetServiceItemInstance gets a reference to the given MoMoRef and assigns it to the ServiceItemInstance field.
func (o *WorkflowOperationTypePostDeployment) SetServiceItemInstance(v MoMoRef) {
	o.ServiceItemInstance = &v
}

func (o WorkflowOperationTypePostDeployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowOperationTypePostDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBaseOperation, errWorkflowBaseOperation := json.Marshal(o.WorkflowBaseOperation)
	if errWorkflowBaseOperation != nil {
		return map[string]interface{}{}, errWorkflowBaseOperation
	}
	errWorkflowBaseOperation = json.Unmarshal([]byte(serializedWorkflowBaseOperation), &toSerialize)
	if errWorkflowBaseOperation != nil {
		return map[string]interface{}{}, errWorkflowBaseOperation
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ServiceItemActionDefinition) {
		toSerialize["ServiceItemActionDefinition"] = o.ServiceItemActionDefinition
	}
	if !IsNil(o.ServiceItemInstance) {
		toSerialize["ServiceItemInstance"] = o.ServiceItemInstance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowOperationTypePostDeployment) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowOperationTypePostDeploymentWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType                  string   `json:"ObjectType"`
		ServiceItemActionDefinition *MoMoRef `json:"ServiceItemActionDefinition,omitempty"`
		ServiceItemInstance         *MoMoRef `json:"ServiceItemInstance,omitempty"`
	}

	varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct := WorkflowOperationTypePostDeploymentWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowOperationTypePostDeployment := _WorkflowOperationTypePostDeployment{}
		varWorkflowOperationTypePostDeployment.ClassId = varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct.ClassId
		varWorkflowOperationTypePostDeployment.ObjectType = varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct.ObjectType
		varWorkflowOperationTypePostDeployment.ServiceItemActionDefinition = varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct.ServiceItemActionDefinition
		varWorkflowOperationTypePostDeployment.ServiceItemInstance = varWorkflowOperationTypePostDeploymentWithoutEmbeddedStruct.ServiceItemInstance
		*o = WorkflowOperationTypePostDeployment(varWorkflowOperationTypePostDeployment)
	} else {
		return err
	}

	varWorkflowOperationTypePostDeployment := _WorkflowOperationTypePostDeployment{}

	err = json.Unmarshal(data, &varWorkflowOperationTypePostDeployment)
	if err == nil {
		o.WorkflowBaseOperation = varWorkflowOperationTypePostDeployment.WorkflowBaseOperation
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServiceItemActionDefinition")
		delete(additionalProperties, "ServiceItemInstance")

		// remove fields from embedded structs
		reflectWorkflowBaseOperation := reflect.ValueOf(o.WorkflowBaseOperation)
		for i := 0; i < reflectWorkflowBaseOperation.Type().NumField(); i++ {
			t := reflectWorkflowBaseOperation.Type().Field(i)

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

type NullableWorkflowOperationTypePostDeployment struct {
	value *WorkflowOperationTypePostDeployment
	isSet bool
}

func (v NullableWorkflowOperationTypePostDeployment) Get() *WorkflowOperationTypePostDeployment {
	return v.value
}

func (v *NullableWorkflowOperationTypePostDeployment) Set(val *WorkflowOperationTypePostDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowOperationTypePostDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowOperationTypePostDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowOperationTypePostDeployment(val *WorkflowOperationTypePostDeployment) *NullableWorkflowOperationTypePostDeployment {
	return &NullableWorkflowOperationTypePostDeployment{value: val, isSet: true}
}

func (v NullableWorkflowOperationTypePostDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowOperationTypePostDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
