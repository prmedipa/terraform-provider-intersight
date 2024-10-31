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

// checks if the AssetVirtualizationService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetVirtualizationService{}

// AssetVirtualizationService The necessary configuration details to enable  Intersight Virtualization features on the selected managed target.
type AssetVirtualizationService struct {
	AssetService
	AdditionalProperties map[string]interface{}
}

type _AssetVirtualizationService AssetVirtualizationService

// NewAssetVirtualizationService instantiates a new AssetVirtualizationService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetVirtualizationService(classId string, objectType string) *AssetVirtualizationService {
	this := AssetVirtualizationService{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetVirtualizationServiceWithDefaults instantiates a new AssetVirtualizationService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetVirtualizationServiceWithDefaults() *AssetVirtualizationService {
	this := AssetVirtualizationService{}
	return &this
}

func (o AssetVirtualizationService) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetVirtualizationService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetService, errAssetService := json.Marshal(o.AssetService)
	if errAssetService != nil {
		return map[string]interface{}{}, errAssetService
	}
	errAssetService = json.Unmarshal([]byte(serializedAssetService), &toSerialize)
	if errAssetService != nil {
		return map[string]interface{}{}, errAssetService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetVirtualizationService) UnmarshalJSON(data []byte) (err error) {
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
	type AssetVirtualizationServiceWithoutEmbeddedStruct struct {
	}

	varAssetVirtualizationServiceWithoutEmbeddedStruct := AssetVirtualizationServiceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetVirtualizationServiceWithoutEmbeddedStruct)
	if err == nil {
		varAssetVirtualizationService := _AssetVirtualizationService{}
		*o = AssetVirtualizationService(varAssetVirtualizationService)
	} else {
		return err
	}

	varAssetVirtualizationService := _AssetVirtualizationService{}

	err = json.Unmarshal(data, &varAssetVirtualizationService)
	if err == nil {
		o.AssetService = varAssetVirtualizationService.AssetService
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectAssetService := reflect.ValueOf(o.AssetService)
		for i := 0; i < reflectAssetService.Type().NumField(); i++ {
			t := reflectAssetService.Type().Field(i)

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

type NullableAssetVirtualizationService struct {
	value *AssetVirtualizationService
	isSet bool
}

func (v NullableAssetVirtualizationService) Get() *AssetVirtualizationService {
	return v.value
}

func (v *NullableAssetVirtualizationService) Set(val *AssetVirtualizationService) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetVirtualizationService) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetVirtualizationService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetVirtualizationService(val *AssetVirtualizationService) *NullableAssetVirtualizationService {
	return &NullableAssetVirtualizationService{value: val, isSet: true}
}

func (v NullableAssetVirtualizationService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetVirtualizationService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
