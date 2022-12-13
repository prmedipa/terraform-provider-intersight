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

// NiaapiDetail The content inside the metadata package with filename and check sum for this file.
type NiaapiDetail struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Checksum of this part of Content.
	Chksum *string `json:"Chksum,omitempty"`
	// The file name within this Metadata file.
	Filename *string `json:"Filename,omitempty"`
	// The name of this Content.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiDetail NiaapiDetail

// NewNiaapiDetail instantiates a new NiaapiDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiDetail(classId string, objectType string) *NiaapiDetail {
	this := NiaapiDetail{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiDetailWithDefaults instantiates a new NiaapiDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiDetailWithDefaults() *NiaapiDetail {
	this := NiaapiDetail{}
	var classId string = "niaapi.Detail"
	this.ClassId = classId
	var objectType string = "niaapi.Detail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiDetail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiDetail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiDetail) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiDetail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiDetail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiDetail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetChksum returns the Chksum field value if set, zero value otherwise.
func (o *NiaapiDetail) GetChksum() string {
	if o == nil || o.Chksum == nil {
		var ret string
		return ret
	}
	return *o.Chksum
}

// GetChksumOk returns a tuple with the Chksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetail) GetChksumOk() (*string, bool) {
	if o == nil || o.Chksum == nil {
		return nil, false
	}
	return o.Chksum, true
}

// HasChksum returns a boolean if a field has been set.
func (o *NiaapiDetail) HasChksum() bool {
	if o != nil && o.Chksum != nil {
		return true
	}

	return false
}

// SetChksum gets a reference to the given string and assigns it to the Chksum field.
func (o *NiaapiDetail) SetChksum(v string) {
	o.Chksum = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *NiaapiDetail) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetail) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *NiaapiDetail) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *NiaapiDetail) SetFilename(v string) {
	o.Filename = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NiaapiDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NiaapiDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NiaapiDetail) SetName(v string) {
	o.Name = &v
}

func (o NiaapiDetail) MarshalJSON() ([]byte, error) {
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
	if o.Chksum != nil {
		toSerialize["Chksum"] = o.Chksum
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiDetail) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiDetailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Checksum of this part of Content.
		Chksum *string `json:"Chksum,omitempty"`
		// The file name within this Metadata file.
		Filename *string `json:"Filename,omitempty"`
		// The name of this Content.
		Name *string `json:"Name,omitempty"`
	}

	varNiaapiDetailWithoutEmbeddedStruct := NiaapiDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiDetailWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiDetail := _NiaapiDetail{}
		varNiaapiDetail.ClassId = varNiaapiDetailWithoutEmbeddedStruct.ClassId
		varNiaapiDetail.ObjectType = varNiaapiDetailWithoutEmbeddedStruct.ObjectType
		varNiaapiDetail.Chksum = varNiaapiDetailWithoutEmbeddedStruct.Chksum
		varNiaapiDetail.Filename = varNiaapiDetailWithoutEmbeddedStruct.Filename
		varNiaapiDetail.Name = varNiaapiDetailWithoutEmbeddedStruct.Name
		*o = NiaapiDetail(varNiaapiDetail)
	} else {
		return err
	}

	varNiaapiDetail := _NiaapiDetail{}

	err = json.Unmarshal(bytes, &varNiaapiDetail)
	if err == nil {
		o.MoBaseComplexType = varNiaapiDetail.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Chksum")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Name")

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

type NullableNiaapiDetail struct {
	value *NiaapiDetail
	isSet bool
}

func (v NullableNiaapiDetail) Get() *NiaapiDetail {
	return v.value
}

func (v *NullableNiaapiDetail) Set(val *NiaapiDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiDetail(val *NiaapiDetail) *NullableNiaapiDetail {
	return &NullableNiaapiDetail{value: val, isSet: true}
}

func (v NullableNiaapiDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
