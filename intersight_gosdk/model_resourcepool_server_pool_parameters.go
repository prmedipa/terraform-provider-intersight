/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ResourcepoolServerPoolParameters This server pool specific parameters are captured as part this parameter.
type ResourcepoolServerPoolParameters struct {
	ResourcepoolResourcePoolParameters
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform for which the servers in resource pool are applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `IntersightStandalone` - Intersight Standalone mode of operation. * `UCSM` - Unified Computing System Manager mode of operation. * `Intersight` - Intersight managed mode of operation.
	ManagementMode       *string `json:"ManagementMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolServerPoolParameters ResourcepoolServerPoolParameters

// NewResourcepoolServerPoolParameters instantiates a new ResourcepoolServerPoolParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolServerPoolParameters(classId string, objectType string) *ResourcepoolServerPoolParameters {
	this := ResourcepoolServerPoolParameters{}
	this.ClassId = classId
	this.ObjectType = objectType
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// NewResourcepoolServerPoolParametersWithDefaults instantiates a new ResourcepoolServerPoolParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolServerPoolParametersWithDefaults() *ResourcepoolServerPoolParameters {
	this := ResourcepoolServerPoolParameters{}
	var classId string = "resourcepool.ServerPoolParameters"
	this.ClassId = classId
	var objectType string = "resourcepool.ServerPoolParameters"
	this.ObjectType = objectType
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolServerPoolParameters) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParameters) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolServerPoolParameters) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolServerPoolParameters) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParameters) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolServerPoolParameters) SetObjectType(v string) {
	o.ObjectType = v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *ResourcepoolServerPoolParameters) GetManagementMode() string {
	if o == nil || o.ManagementMode == nil {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolServerPoolParameters) GetManagementModeOk() (*string, bool) {
	if o == nil || o.ManagementMode == nil {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *ResourcepoolServerPoolParameters) HasManagementMode() bool {
	if o != nil && o.ManagementMode != nil {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *ResourcepoolServerPoolParameters) SetManagementMode(v string) {
	o.ManagementMode = &v
}

func (o ResourcepoolServerPoolParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedResourcepoolResourcePoolParameters, errResourcepoolResourcePoolParameters := json.Marshal(o.ResourcepoolResourcePoolParameters)
	if errResourcepoolResourcePoolParameters != nil {
		return []byte{}, errResourcepoolResourcePoolParameters
	}
	errResourcepoolResourcePoolParameters = json.Unmarshal([]byte(serializedResourcepoolResourcePoolParameters), &toSerialize)
	if errResourcepoolResourcePoolParameters != nil {
		return []byte{}, errResourcepoolResourcePoolParameters
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ManagementMode != nil {
		toSerialize["ManagementMode"] = o.ManagementMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolServerPoolParameters) UnmarshalJSON(bytes []byte) (err error) {
	type ResourcepoolServerPoolParametersWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The platform for which the servers in resource pool are applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `IntersightStandalone` - Intersight Standalone mode of operation. * `UCSM` - Unified Computing System Manager mode of operation. * `Intersight` - Intersight managed mode of operation.
		ManagementMode *string `json:"ManagementMode,omitempty"`
	}

	varResourcepoolServerPoolParametersWithoutEmbeddedStruct := ResourcepoolServerPoolParametersWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varResourcepoolServerPoolParametersWithoutEmbeddedStruct)
	if err == nil {
		varResourcepoolServerPoolParameters := _ResourcepoolServerPoolParameters{}
		varResourcepoolServerPoolParameters.ClassId = varResourcepoolServerPoolParametersWithoutEmbeddedStruct.ClassId
		varResourcepoolServerPoolParameters.ObjectType = varResourcepoolServerPoolParametersWithoutEmbeddedStruct.ObjectType
		varResourcepoolServerPoolParameters.ManagementMode = varResourcepoolServerPoolParametersWithoutEmbeddedStruct.ManagementMode
		*o = ResourcepoolServerPoolParameters(varResourcepoolServerPoolParameters)
	} else {
		return err
	}

	varResourcepoolServerPoolParameters := _ResourcepoolServerPoolParameters{}

	err = json.Unmarshal(bytes, &varResourcepoolServerPoolParameters)
	if err == nil {
		o.ResourcepoolResourcePoolParameters = varResourcepoolServerPoolParameters.ResourcepoolResourcePoolParameters
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ManagementMode")

		// remove fields from embedded structs
		reflectResourcepoolResourcePoolParameters := reflect.ValueOf(o.ResourcepoolResourcePoolParameters)
		for i := 0; i < reflectResourcepoolResourcePoolParameters.Type().NumField(); i++ {
			t := reflectResourcepoolResourcePoolParameters.Type().Field(i)

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

type NullableResourcepoolServerPoolParameters struct {
	value *ResourcepoolServerPoolParameters
	isSet bool
}

func (v NullableResourcepoolServerPoolParameters) Get() *ResourcepoolServerPoolParameters {
	return v.value
}

func (v *NullableResourcepoolServerPoolParameters) Set(val *ResourcepoolServerPoolParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolServerPoolParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolServerPoolParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolServerPoolParameters(val *ResourcepoolServerPoolParameters) *NullableResourcepoolServerPoolParameters {
	return &NullableResourcepoolServerPoolParameters{value: val, isSet: true}
}

func (v NullableResourcepoolServerPoolParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolServerPoolParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
