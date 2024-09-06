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

// checks if the NiaapiSnValidatorMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiaapiSnValidatorMetadata{}

// NiaapiSnValidatorMetadata Contains the Serial Metadata Version.
type NiaapiSnValidatorMetadata struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Checksum of SnValidatorData file.
	Checksum *string `json:"Checksum,omitempty"`
	// The filename of sn metadata file.
	FileName *string `json:"FileName,omitempty"`
	// The version number of the SerialNumber Metadata.
	Version              *int64 `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiSnValidatorMetadata NiaapiSnValidatorMetadata

// NewNiaapiSnValidatorMetadata instantiates a new NiaapiSnValidatorMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiSnValidatorMetadata(classId string, objectType string) *NiaapiSnValidatorMetadata {
	this := NiaapiSnValidatorMetadata{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiSnValidatorMetadataWithDefaults instantiates a new NiaapiSnValidatorMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiSnValidatorMetadataWithDefaults() *NiaapiSnValidatorMetadata {
	this := NiaapiSnValidatorMetadata{}
	var classId string = "niaapi.SnValidatorMetadata"
	this.ClassId = classId
	var objectType string = "niaapi.SnValidatorMetadata"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiSnValidatorMetadata) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiSnValidatorMetadata) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiSnValidatorMetadata) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niaapi.SnValidatorMetadata" of the ClassId field.
func (o *NiaapiSnValidatorMetadata) GetDefaultClassId() interface{} {
	return "niaapi.SnValidatorMetadata"
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiSnValidatorMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiSnValidatorMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiSnValidatorMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niaapi.SnValidatorMetadata" of the ObjectType field.
func (o *NiaapiSnValidatorMetadata) GetDefaultObjectType() interface{} {
	return "niaapi.SnValidatorMetadata"
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *NiaapiSnValidatorMetadata) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSnValidatorMetadata) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *NiaapiSnValidatorMetadata) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *NiaapiSnValidatorMetadata) SetChecksum(v string) {
	o.Checksum = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *NiaapiSnValidatorMetadata) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSnValidatorMetadata) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *NiaapiSnValidatorMetadata) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *NiaapiSnValidatorMetadata) SetFileName(v string) {
	o.FileName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiaapiSnValidatorMetadata) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSnValidatorMetadata) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiaapiSnValidatorMetadata) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *NiaapiSnValidatorMetadata) SetVersion(v int64) {
	o.Version = &v
}

func (o NiaapiSnValidatorMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiaapiSnValidatorMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Checksum) {
		toSerialize["Checksum"] = o.Checksum
	}
	if !IsNil(o.FileName) {
		toSerialize["FileName"] = o.FileName
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiaapiSnValidatorMetadata) UnmarshalJSON(data []byte) (err error) {
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
	type NiaapiSnValidatorMetadataWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Checksum of SnValidatorData file.
		Checksum *string `json:"Checksum,omitempty"`
		// The filename of sn metadata file.
		FileName *string `json:"FileName,omitempty"`
		// The version number of the SerialNumber Metadata.
		Version *int64 `json:"Version,omitempty"`
	}

	varNiaapiSnValidatorMetadataWithoutEmbeddedStruct := NiaapiSnValidatorMetadataWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiaapiSnValidatorMetadataWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiSnValidatorMetadata := _NiaapiSnValidatorMetadata{}
		varNiaapiSnValidatorMetadata.ClassId = varNiaapiSnValidatorMetadataWithoutEmbeddedStruct.ClassId
		varNiaapiSnValidatorMetadata.ObjectType = varNiaapiSnValidatorMetadataWithoutEmbeddedStruct.ObjectType
		varNiaapiSnValidatorMetadata.Checksum = varNiaapiSnValidatorMetadataWithoutEmbeddedStruct.Checksum
		varNiaapiSnValidatorMetadata.FileName = varNiaapiSnValidatorMetadataWithoutEmbeddedStruct.FileName
		varNiaapiSnValidatorMetadata.Version = varNiaapiSnValidatorMetadataWithoutEmbeddedStruct.Version
		*o = NiaapiSnValidatorMetadata(varNiaapiSnValidatorMetadata)
	} else {
		return err
	}

	varNiaapiSnValidatorMetadata := _NiaapiSnValidatorMetadata{}

	err = json.Unmarshal(data, &varNiaapiSnValidatorMetadata)
	if err == nil {
		o.MoBaseMo = varNiaapiSnValidatorMetadata.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Checksum")
		delete(additionalProperties, "FileName")
		delete(additionalProperties, "Version")

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

type NullableNiaapiSnValidatorMetadata struct {
	value *NiaapiSnValidatorMetadata
	isSet bool
}

func (v NullableNiaapiSnValidatorMetadata) Get() *NiaapiSnValidatorMetadata {
	return v.value
}

func (v *NullableNiaapiSnValidatorMetadata) Set(val *NiaapiSnValidatorMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiSnValidatorMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiSnValidatorMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiSnValidatorMetadata(val *NiaapiSnValidatorMetadata) *NullableNiaapiSnValidatorMetadata {
	return &NullableNiaapiSnValidatorMetadata{value: val, isSet: true}
}

func (v NullableNiaapiSnValidatorMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiSnValidatorMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
