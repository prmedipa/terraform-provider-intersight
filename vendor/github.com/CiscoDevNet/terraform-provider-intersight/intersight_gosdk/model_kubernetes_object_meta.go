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

// checks if the KubernetesObjectMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesObjectMeta{}

// KubernetesObjectMeta Kubernetes Object Meta Data.
type KubernetesObjectMeta struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists.
	CreationTimestamp *string `json:"CreationTimestamp,omitempty"`
	// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Read-only.
	Name *string `json:"Name,omitempty"`
	// Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Read-only.
	Namespace *string `json:"Namespace,omitempty"`
	// Specific resourceVersion to which this reference is made, if any.
	ResourceVersion *string `json:"ResourceVersion,omitempty"`
	// UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only.
	Uuid                 *string `json:"Uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesObjectMeta KubernetesObjectMeta

// NewKubernetesObjectMeta instantiates a new KubernetesObjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesObjectMeta(classId string, objectType string) *KubernetesObjectMeta {
	this := KubernetesObjectMeta{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesObjectMetaWithDefaults instantiates a new KubernetesObjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesObjectMetaWithDefaults() *KubernetesObjectMeta {
	this := KubernetesObjectMeta{}
	var classId string = "kubernetes.ObjectMeta"
	this.ClassId = classId
	var objectType string = "kubernetes.ObjectMeta"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesObjectMeta) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesObjectMeta) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "kubernetes.ObjectMeta" of the ClassId field.
func (o *KubernetesObjectMeta) GetDefaultClassId() interface{} {
	return "kubernetes.ObjectMeta"
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesObjectMeta) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesObjectMeta) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "kubernetes.ObjectMeta" of the ObjectType field.
func (o *KubernetesObjectMeta) GetDefaultObjectType() interface{} {
	return "kubernetes.ObjectMeta"
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *KubernetesObjectMeta) GetCreationTimestamp() string {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetCreationTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *KubernetesObjectMeta) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *KubernetesObjectMeta) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubernetesObjectMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubernetesObjectMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubernetesObjectMeta) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *KubernetesObjectMeta) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubernetesObjectMeta) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *KubernetesObjectMeta) SetNamespace(v string) {
	o.Namespace = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *KubernetesObjectMeta) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *KubernetesObjectMeta) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *KubernetesObjectMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *KubernetesObjectMeta) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesObjectMeta) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *KubernetesObjectMeta) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *KubernetesObjectMeta) SetUuid(v string) {
	o.Uuid = &v
}

func (o KubernetesObjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesObjectMeta) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CreationTimestamp) {
		toSerialize["CreationTimestamp"] = o.CreationTimestamp
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["Namespace"] = o.Namespace
	}
	if !IsNil(o.ResourceVersion) {
		toSerialize["ResourceVersion"] = o.ResourceVersion
	}
	if !IsNil(o.Uuid) {
		toSerialize["Uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesObjectMeta) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesObjectMetaWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists.
		CreationTimestamp *string `json:"CreationTimestamp,omitempty"`
		// Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Read-only.
		Name *string `json:"Name,omitempty"`
		// Namespace defines the space within each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Read-only.
		Namespace *string `json:"Namespace,omitempty"`
		// Specific resourceVersion to which this reference is made, if any.
		ResourceVersion *string `json:"ResourceVersion,omitempty"`
		// UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations. Populated by the system. Read-only.
		Uuid *string `json:"Uuid,omitempty"`
	}

	varKubernetesObjectMetaWithoutEmbeddedStruct := KubernetesObjectMetaWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesObjectMetaWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesObjectMeta := _KubernetesObjectMeta{}
		varKubernetesObjectMeta.ClassId = varKubernetesObjectMetaWithoutEmbeddedStruct.ClassId
		varKubernetesObjectMeta.ObjectType = varKubernetesObjectMetaWithoutEmbeddedStruct.ObjectType
		varKubernetesObjectMeta.CreationTimestamp = varKubernetesObjectMetaWithoutEmbeddedStruct.CreationTimestamp
		varKubernetesObjectMeta.Name = varKubernetesObjectMetaWithoutEmbeddedStruct.Name
		varKubernetesObjectMeta.Namespace = varKubernetesObjectMetaWithoutEmbeddedStruct.Namespace
		varKubernetesObjectMeta.ResourceVersion = varKubernetesObjectMetaWithoutEmbeddedStruct.ResourceVersion
		varKubernetesObjectMeta.Uuid = varKubernetesObjectMetaWithoutEmbeddedStruct.Uuid
		*o = KubernetesObjectMeta(varKubernetesObjectMeta)
	} else {
		return err
	}

	varKubernetesObjectMeta := _KubernetesObjectMeta{}

	err = json.Unmarshal(data, &varKubernetesObjectMeta)
	if err == nil {
		o.MoBaseComplexType = varKubernetesObjectMeta.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CreationTimestamp")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Namespace")
		delete(additionalProperties, "ResourceVersion")
		delete(additionalProperties, "Uuid")

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

type NullableKubernetesObjectMeta struct {
	value *KubernetesObjectMeta
	isSet bool
}

func (v NullableKubernetesObjectMeta) Get() *KubernetesObjectMeta {
	return v.value
}

func (v *NullableKubernetesObjectMeta) Set(val *KubernetesObjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesObjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesObjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesObjectMeta(val *KubernetesObjectMeta) *NullableKubernetesObjectMeta {
	return &NullableKubernetesObjectMeta{value: val, isSet: true}
}

func (v NullableKubernetesObjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesObjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
