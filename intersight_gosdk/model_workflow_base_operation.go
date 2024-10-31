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

// checks if the WorkflowBaseOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowBaseOperation{}

// WorkflowBaseOperation Specify operation type on a service item like deployment or post-deployment to be performed on catalog.
type WorkflowBaseOperation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Type of action operation to be performed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
	OperationType        *string `json:"OperationType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBaseOperation WorkflowBaseOperation

// NewWorkflowBaseOperation instantiates a new WorkflowBaseOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBaseOperation(classId string, objectType string) *WorkflowBaseOperation {
	this := WorkflowBaseOperation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	return &this
}

// NewWorkflowBaseOperationWithDefaults instantiates a new WorkflowBaseOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBaseOperationWithDefaults() *WorkflowBaseOperation {
	this := WorkflowBaseOperation{}
	var operationType string = "PostDeployment"
	this.OperationType = &operationType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBaseOperation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBaseOperation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBaseOperation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBaseOperation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBaseOperation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBaseOperation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *WorkflowBaseOperation) GetOperationType() string {
	if o == nil || IsNil(o.OperationType) {
		var ret string
		return ret
	}
	return *o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBaseOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *WorkflowBaseOperation) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given string and assigns it to the OperationType field.
func (o *WorkflowBaseOperation) SetOperationType(v string) {
	o.OperationType = &v
}

func (o WorkflowBaseOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowBaseOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.OperationType) {
		toSerialize["OperationType"] = o.OperationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowBaseOperation) UnmarshalJSON(data []byte) (err error) {
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
	type WorkflowBaseOperationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Type of action operation to be performed on the service item. * `PostDeployment` - This represents the post-deployment actions for the resources created or defined through the deployment action. There can be more than one post-deployment operations associated with a service item. * `Deployment` - This represents the deploy action, for the service item action definition. This operation type is used to create or define resources that is managed by the service item. There can only be one Service Item Action Definition that can be marked with the operation type as Deployment and this is a mandatory operation type. All valid Service Items must have one and only one operation type marked as type Deployment. * `Decommission` - This represents the decommission action, used to decommission the created resources. All valid Service Items must have one and only one operation type marked as type Decommission. Once a decommission action is run on a Service Item, no further operations are allowed on that Service Item.
		OperationType *string `json:"OperationType,omitempty"`
	}

	varWorkflowBaseOperationWithoutEmbeddedStruct := WorkflowBaseOperationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowBaseOperationWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowBaseOperation := _WorkflowBaseOperation{}
		varWorkflowBaseOperation.ClassId = varWorkflowBaseOperationWithoutEmbeddedStruct.ClassId
		varWorkflowBaseOperation.ObjectType = varWorkflowBaseOperationWithoutEmbeddedStruct.ObjectType
		varWorkflowBaseOperation.OperationType = varWorkflowBaseOperationWithoutEmbeddedStruct.OperationType
		*o = WorkflowBaseOperation(varWorkflowBaseOperation)
	} else {
		return err
	}

	varWorkflowBaseOperation := _WorkflowBaseOperation{}

	err = json.Unmarshal(data, &varWorkflowBaseOperation)
	if err == nil {
		o.MoBaseComplexType = varWorkflowBaseOperation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperationType")

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

type NullableWorkflowBaseOperation struct {
	value *WorkflowBaseOperation
	isSet bool
}

func (v NullableWorkflowBaseOperation) Get() *WorkflowBaseOperation {
	return v.value
}

func (v *NullableWorkflowBaseOperation) Set(val *WorkflowBaseOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBaseOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBaseOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBaseOperation(val *WorkflowBaseOperation) *NullableWorkflowBaseOperation {
	return &NullableWorkflowBaseOperation{value: val, isSet: true}
}

func (v NullableWorkflowBaseOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBaseOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
