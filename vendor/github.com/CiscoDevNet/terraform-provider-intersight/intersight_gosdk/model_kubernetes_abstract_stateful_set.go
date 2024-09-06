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

// checks if the KubernetesAbstractStatefulSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesAbstractStatefulSet{}

// KubernetesAbstractStatefulSet Abstract StatefulSet represents Kubernetes StatefulSet.
type KubernetesAbstractStatefulSet struct {
	KubernetesKubernetesResource
	AdditionalProperties map[string]interface{}
}

type _KubernetesAbstractStatefulSet KubernetesAbstractStatefulSet

// NewKubernetesAbstractStatefulSet instantiates a new KubernetesAbstractStatefulSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesAbstractStatefulSet(classId string, objectType string) *KubernetesAbstractStatefulSet {
	this := KubernetesAbstractStatefulSet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesAbstractStatefulSetWithDefaults instantiates a new KubernetesAbstractStatefulSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesAbstractStatefulSetWithDefaults() *KubernetesAbstractStatefulSet {
	this := KubernetesAbstractStatefulSet{}
	return &this
}

func (o KubernetesAbstractStatefulSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesAbstractStatefulSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedKubernetesKubernetesResource, errKubernetesKubernetesResource := json.Marshal(o.KubernetesKubernetesResource)
	if errKubernetesKubernetesResource != nil {
		return map[string]interface{}{}, errKubernetesKubernetesResource
	}
	errKubernetesKubernetesResource = json.Unmarshal([]byte(serializedKubernetesKubernetesResource), &toSerialize)
	if errKubernetesKubernetesResource != nil {
		return map[string]interface{}{}, errKubernetesKubernetesResource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KubernetesAbstractStatefulSet) UnmarshalJSON(data []byte) (err error) {
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
	type KubernetesAbstractStatefulSetWithoutEmbeddedStruct struct {
	}

	varKubernetesAbstractStatefulSetWithoutEmbeddedStruct := KubernetesAbstractStatefulSetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varKubernetesAbstractStatefulSetWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesAbstractStatefulSet := _KubernetesAbstractStatefulSet{}
		*o = KubernetesAbstractStatefulSet(varKubernetesAbstractStatefulSet)
	} else {
		return err
	}

	varKubernetesAbstractStatefulSet := _KubernetesAbstractStatefulSet{}

	err = json.Unmarshal(data, &varKubernetesAbstractStatefulSet)
	if err == nil {
		o.KubernetesKubernetesResource = varKubernetesAbstractStatefulSet.KubernetesKubernetesResource
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectKubernetesKubernetesResource := reflect.ValueOf(o.KubernetesKubernetesResource)
		for i := 0; i < reflectKubernetesKubernetesResource.Type().NumField(); i++ {
			t := reflectKubernetesKubernetesResource.Type().Field(i)

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

type NullableKubernetesAbstractStatefulSet struct {
	value *KubernetesAbstractStatefulSet
	isSet bool
}

func (v NullableKubernetesAbstractStatefulSet) Get() *KubernetesAbstractStatefulSet {
	return v.value
}

func (v *NullableKubernetesAbstractStatefulSet) Set(val *KubernetesAbstractStatefulSet) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesAbstractStatefulSet) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesAbstractStatefulSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesAbstractStatefulSet(val *KubernetesAbstractStatefulSet) *NullableKubernetesAbstractStatefulSet {
	return &NullableKubernetesAbstractStatefulSet{value: val, isSet: true}
}

func (v NullableKubernetesAbstractStatefulSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesAbstractStatefulSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
