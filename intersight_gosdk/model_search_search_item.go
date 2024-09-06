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

// checks if the SearchSearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchItem{}

// SearchSearchItem The Search service entry point to search Intersight REST resources using OData query syntax. See [Search API query syntax](/apidocs/introduction/query/#search-api) for details about the query syntax.
type SearchSearchItem struct {
	MoBaseMo
	AdditionalProperties map[string]interface{}
}

type _SearchSearchItem SearchSearchItem

// NewSearchSearchItem instantiates a new SearchSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchItem(classId string, objectType string) *SearchSearchItem {
	this := SearchSearchItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSearchSearchItemWithDefaults instantiates a new SearchSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchItemWithDefaults() *SearchSearchItem {
	this := SearchSearchItem{}
	return &this
}

func (o SearchSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchSearchItem) UnmarshalJSON(data []byte) (err error) {
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
	type SearchSearchItemWithoutEmbeddedStruct struct {
	}

	varSearchSearchItemWithoutEmbeddedStruct := SearchSearchItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSearchSearchItemWithoutEmbeddedStruct)
	if err == nil {
		varSearchSearchItem := _SearchSearchItem{}
		*o = SearchSearchItem(varSearchSearchItem)
	} else {
		return err
	}

	varSearchSearchItem := _SearchSearchItem{}

	err = json.Unmarshal(data, &varSearchSearchItem)
	if err == nil {
		o.MoBaseMo = varSearchSearchItem.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableSearchSearchItem struct {
	value *SearchSearchItem
	isSet bool
}

func (v NullableSearchSearchItem) Get() *SearchSearchItem {
	return v.value
}

func (v *NullableSearchSearchItem) Set(val *SearchSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchItem(val *SearchSearchItem) *NullableSearchSearchItem {
	return &NullableSearchSearchItem{value: val, isSet: true}
}

func (v NullableSearchSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
