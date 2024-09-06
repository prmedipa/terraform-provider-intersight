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

// checks if the NiatelemetryJobDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryJobDetail{}

// NiatelemetryJobDetail Stores information related to job details.
type NiatelemetryJobDetail struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the id of the job.
	JobId *int64 `json:"JobId,omitempty"`
	// Returns the status of the job.
	UpgStatus            *string `json:"UpgStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryJobDetail NiatelemetryJobDetail

// NewNiatelemetryJobDetail instantiates a new NiatelemetryJobDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryJobDetail(classId string, objectType string) *NiatelemetryJobDetail {
	this := NiatelemetryJobDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryJobDetailWithDefaults instantiates a new NiatelemetryJobDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryJobDetailWithDefaults() *NiatelemetryJobDetail {
	this := NiatelemetryJobDetail{}
	var classId string = "niatelemetry.JobDetail"
	this.ClassId = classId
	var objectType string = "niatelemetry.JobDetail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryJobDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryJobDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.JobDetail" of the ClassId field.
func (o *NiatelemetryJobDetail) GetDefaultClassId() interface{} {
	return "niatelemetry.JobDetail"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryJobDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryJobDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.JobDetail" of the ObjectType field.
func (o *NiatelemetryJobDetail) GetDefaultObjectType() interface{} {
	return "niatelemetry.JobDetail"
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *NiatelemetryJobDetail) GetJobId() int64 {
	if o == nil || IsNil(o.JobId) {
		var ret int64
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetail) GetJobIdOk() (*int64, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *NiatelemetryJobDetail) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given int64 and assigns it to the JobId field.
func (o *NiatelemetryJobDetail) SetJobId(v int64) {
	o.JobId = &v
}

// GetUpgStatus returns the UpgStatus field value if set, zero value otherwise.
func (o *NiatelemetryJobDetail) GetUpgStatus() string {
	if o == nil || IsNil(o.UpgStatus) {
		var ret string
		return ret
	}
	return *o.UpgStatus
}

// GetUpgStatusOk returns a tuple with the UpgStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryJobDetail) GetUpgStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UpgStatus) {
		return nil, false
	}
	return o.UpgStatus, true
}

// HasUpgStatus returns a boolean if a field has been set.
func (o *NiatelemetryJobDetail) HasUpgStatus() bool {
	if o != nil && !IsNil(o.UpgStatus) {
		return true
	}

	return false
}

// SetUpgStatus gets a reference to the given string and assigns it to the UpgStatus field.
func (o *NiatelemetryJobDetail) SetUpgStatus(v string) {
	o.UpgStatus = &v
}

func (o NiatelemetryJobDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryJobDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.JobId) {
		toSerialize["JobId"] = o.JobId
	}
	if !IsNil(o.UpgStatus) {
		toSerialize["UpgStatus"] = o.UpgStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryJobDetail) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryJobDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Returns the id of the job.
		JobId *int64 `json:"JobId,omitempty"`
		// Returns the status of the job.
		UpgStatus *string `json:"UpgStatus,omitempty"`
	}

	varNiatelemetryJobDetailWithoutEmbeddedStruct := NiatelemetryJobDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryJobDetailWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryJobDetail := _NiatelemetryJobDetail{}
		varNiatelemetryJobDetail.ClassId = varNiatelemetryJobDetailWithoutEmbeddedStruct.ClassId
		varNiatelemetryJobDetail.ObjectType = varNiatelemetryJobDetailWithoutEmbeddedStruct.ObjectType
		varNiatelemetryJobDetail.JobId = varNiatelemetryJobDetailWithoutEmbeddedStruct.JobId
		varNiatelemetryJobDetail.UpgStatus = varNiatelemetryJobDetailWithoutEmbeddedStruct.UpgStatus
		*o = NiatelemetryJobDetail(varNiatelemetryJobDetail)
	} else {
		return err
	}

	varNiatelemetryJobDetail := _NiatelemetryJobDetail{}

	err = json.Unmarshal(data, &varNiatelemetryJobDetail)
	if err == nil {
		o.MoBaseComplexType = varNiatelemetryJobDetail.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "JobId")
		delete(additionalProperties, "UpgStatus")

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

type NullableNiatelemetryJobDetail struct {
	value *NiatelemetryJobDetail
	isSet bool
}

func (v NullableNiatelemetryJobDetail) Get() *NiatelemetryJobDetail {
	return v.value
}

func (v *NullableNiatelemetryJobDetail) Set(val *NiatelemetryJobDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryJobDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryJobDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryJobDetail(val *NiatelemetryJobDetail) *NullableNiatelemetryJobDetail {
	return &NullableNiatelemetryJobDetail{value: val, isSet: true}
}

func (v NullableNiatelemetryJobDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryJobDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
