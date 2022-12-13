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

// CloudVolumeIopsInfo Volume input/output operations per second.
type CloudVolumeIopsInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of disk read commands that can be performed per second.
	IopsReadLimit *int64 `json:"IopsReadLimit,omitempty"`
	// Number of disk write commands that can be performed per second.
	IopsWriteLimit *int64 `json:"IopsWriteLimit,omitempty"`
	// Data transfer rate limit from the disk, specified in mebibytes (MiB) per second.
	ThroughputReadLimit *int64 `json:"ThroughputReadLimit,omitempty"`
	// Data transfer rate limit to the disk, specified in mebibytes (MiB) per second.
	ThroughputWriteLimit *int64 `json:"ThroughputWriteLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudVolumeIopsInfo CloudVolumeIopsInfo

// NewCloudVolumeIopsInfo instantiates a new CloudVolumeIopsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudVolumeIopsInfo(classId string, objectType string) *CloudVolumeIopsInfo {
	this := CloudVolumeIopsInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudVolumeIopsInfoWithDefaults instantiates a new CloudVolumeIopsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudVolumeIopsInfoWithDefaults() *CloudVolumeIopsInfo {
	this := CloudVolumeIopsInfo{}
	var classId string = "cloud.VolumeIopsInfo"
	this.ClassId = classId
	var objectType string = "cloud.VolumeIopsInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudVolumeIopsInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudVolumeIopsInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudVolumeIopsInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudVolumeIopsInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIopsReadLimit returns the IopsReadLimit field value if set, zero value otherwise.
func (o *CloudVolumeIopsInfo) GetIopsReadLimit() int64 {
	if o == nil || o.IopsReadLimit == nil {
		var ret int64
		return ret
	}
	return *o.IopsReadLimit
}

// GetIopsReadLimitOk returns a tuple with the IopsReadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetIopsReadLimitOk() (*int64, bool) {
	if o == nil || o.IopsReadLimit == nil {
		return nil, false
	}
	return o.IopsReadLimit, true
}

// HasIopsReadLimit returns a boolean if a field has been set.
func (o *CloudVolumeIopsInfo) HasIopsReadLimit() bool {
	if o != nil && o.IopsReadLimit != nil {
		return true
	}

	return false
}

// SetIopsReadLimit gets a reference to the given int64 and assigns it to the IopsReadLimit field.
func (o *CloudVolumeIopsInfo) SetIopsReadLimit(v int64) {
	o.IopsReadLimit = &v
}

// GetIopsWriteLimit returns the IopsWriteLimit field value if set, zero value otherwise.
func (o *CloudVolumeIopsInfo) GetIopsWriteLimit() int64 {
	if o == nil || o.IopsWriteLimit == nil {
		var ret int64
		return ret
	}
	return *o.IopsWriteLimit
}

// GetIopsWriteLimitOk returns a tuple with the IopsWriteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetIopsWriteLimitOk() (*int64, bool) {
	if o == nil || o.IopsWriteLimit == nil {
		return nil, false
	}
	return o.IopsWriteLimit, true
}

// HasIopsWriteLimit returns a boolean if a field has been set.
func (o *CloudVolumeIopsInfo) HasIopsWriteLimit() bool {
	if o != nil && o.IopsWriteLimit != nil {
		return true
	}

	return false
}

// SetIopsWriteLimit gets a reference to the given int64 and assigns it to the IopsWriteLimit field.
func (o *CloudVolumeIopsInfo) SetIopsWriteLimit(v int64) {
	o.IopsWriteLimit = &v
}

// GetThroughputReadLimit returns the ThroughputReadLimit field value if set, zero value otherwise.
func (o *CloudVolumeIopsInfo) GetThroughputReadLimit() int64 {
	if o == nil || o.ThroughputReadLimit == nil {
		var ret int64
		return ret
	}
	return *o.ThroughputReadLimit
}

// GetThroughputReadLimitOk returns a tuple with the ThroughputReadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetThroughputReadLimitOk() (*int64, bool) {
	if o == nil || o.ThroughputReadLimit == nil {
		return nil, false
	}
	return o.ThroughputReadLimit, true
}

// HasThroughputReadLimit returns a boolean if a field has been set.
func (o *CloudVolumeIopsInfo) HasThroughputReadLimit() bool {
	if o != nil && o.ThroughputReadLimit != nil {
		return true
	}

	return false
}

// SetThroughputReadLimit gets a reference to the given int64 and assigns it to the ThroughputReadLimit field.
func (o *CloudVolumeIopsInfo) SetThroughputReadLimit(v int64) {
	o.ThroughputReadLimit = &v
}

// GetThroughputWriteLimit returns the ThroughputWriteLimit field value if set, zero value otherwise.
func (o *CloudVolumeIopsInfo) GetThroughputWriteLimit() int64 {
	if o == nil || o.ThroughputWriteLimit == nil {
		var ret int64
		return ret
	}
	return *o.ThroughputWriteLimit
}

// GetThroughputWriteLimitOk returns a tuple with the ThroughputWriteLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudVolumeIopsInfo) GetThroughputWriteLimitOk() (*int64, bool) {
	if o == nil || o.ThroughputWriteLimit == nil {
		return nil, false
	}
	return o.ThroughputWriteLimit, true
}

// HasThroughputWriteLimit returns a boolean if a field has been set.
func (o *CloudVolumeIopsInfo) HasThroughputWriteLimit() bool {
	if o != nil && o.ThroughputWriteLimit != nil {
		return true
	}

	return false
}

// SetThroughputWriteLimit gets a reference to the given int64 and assigns it to the ThroughputWriteLimit field.
func (o *CloudVolumeIopsInfo) SetThroughputWriteLimit(v int64) {
	o.ThroughputWriteLimit = &v
}

func (o CloudVolumeIopsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IopsReadLimit != nil {
		toSerialize["IopsReadLimit"] = o.IopsReadLimit
	}
	if o.IopsWriteLimit != nil {
		toSerialize["IopsWriteLimit"] = o.IopsWriteLimit
	}
	if o.ThroughputReadLimit != nil {
		toSerialize["ThroughputReadLimit"] = o.ThroughputReadLimit
	}
	if o.ThroughputWriteLimit != nil {
		toSerialize["ThroughputWriteLimit"] = o.ThroughputWriteLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudVolumeIopsInfo) UnmarshalJSON(bytes []byte) (err error) {
	type CloudVolumeIopsInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of disk read commands that can be performed per second.
		IopsReadLimit *int64 `json:"IopsReadLimit,omitempty"`
		// Number of disk write commands that can be performed per second.
		IopsWriteLimit *int64 `json:"IopsWriteLimit,omitempty"`
		// Data transfer rate limit from the disk, specified in mebibytes (MiB) per second.
		ThroughputReadLimit *int64 `json:"ThroughputReadLimit,omitempty"`
		// Data transfer rate limit to the disk, specified in mebibytes (MiB) per second.
		ThroughputWriteLimit *int64 `json:"ThroughputWriteLimit,omitempty"`
	}

	varCloudVolumeIopsInfoWithoutEmbeddedStruct := CloudVolumeIopsInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCloudVolumeIopsInfoWithoutEmbeddedStruct)
	if err == nil {
		varCloudVolumeIopsInfo := _CloudVolumeIopsInfo{}
		varCloudVolumeIopsInfo.ClassId = varCloudVolumeIopsInfoWithoutEmbeddedStruct.ClassId
		varCloudVolumeIopsInfo.ObjectType = varCloudVolumeIopsInfoWithoutEmbeddedStruct.ObjectType
		varCloudVolumeIopsInfo.IopsReadLimit = varCloudVolumeIopsInfoWithoutEmbeddedStruct.IopsReadLimit
		varCloudVolumeIopsInfo.IopsWriteLimit = varCloudVolumeIopsInfoWithoutEmbeddedStruct.IopsWriteLimit
		varCloudVolumeIopsInfo.ThroughputReadLimit = varCloudVolumeIopsInfoWithoutEmbeddedStruct.ThroughputReadLimit
		varCloudVolumeIopsInfo.ThroughputWriteLimit = varCloudVolumeIopsInfoWithoutEmbeddedStruct.ThroughputWriteLimit
		*o = CloudVolumeIopsInfo(varCloudVolumeIopsInfo)
	} else {
		return err
	}

	varCloudVolumeIopsInfo := _CloudVolumeIopsInfo{}

	err = json.Unmarshal(bytes, &varCloudVolumeIopsInfo)
	if err == nil {
		o.MoBaseComplexType = varCloudVolumeIopsInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IopsReadLimit")
		delete(additionalProperties, "IopsWriteLimit")
		delete(additionalProperties, "ThroughputReadLimit")
		delete(additionalProperties, "ThroughputWriteLimit")

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

type NullableCloudVolumeIopsInfo struct {
	value *CloudVolumeIopsInfo
	isSet bool
}

func (v NullableCloudVolumeIopsInfo) Get() *CloudVolumeIopsInfo {
	return v.value
}

func (v *NullableCloudVolumeIopsInfo) Set(val *CloudVolumeIopsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudVolumeIopsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudVolumeIopsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudVolumeIopsInfo(val *CloudVolumeIopsInfo) *NullableCloudVolumeIopsInfo {
	return &NullableCloudVolumeIopsInfo{value: val, isSet: true}
}

func (v NullableCloudVolumeIopsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudVolumeIopsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
