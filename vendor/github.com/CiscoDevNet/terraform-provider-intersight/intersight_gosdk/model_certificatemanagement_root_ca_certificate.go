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

// checks if the CertificatemanagementRootCaCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificatemanagementRootCaCertificate{}

// CertificatemanagementRootCaCertificate Root CA Certificate used for HTTPS server authentication.
type CertificatemanagementRootCaCertificate struct {
	CertificatemanagementCertificateBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// A name that helps identify a certificate. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
	CertificateName      *string `json:"CertificateName,omitempty" validate:"regexp=((^[a-zA-Z0-9]$){1,30}|(^(([a-zA-Z0-9])([a-zA-Z0-9_\\\\-]{0,28})([a-zA-Z0-9]))$))"`
	AdditionalProperties map[string]interface{}
}

type _CertificatemanagementRootCaCertificate CertificatemanagementRootCaCertificate

// NewCertificatemanagementRootCaCertificate instantiates a new CertificatemanagementRootCaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatemanagementRootCaCertificate(classId string, objectType string) *CertificatemanagementRootCaCertificate {
	this := CertificatemanagementRootCaCertificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = true
	this.Enabled = &enabled
	return &this
}

// NewCertificatemanagementRootCaCertificateWithDefaults instantiates a new CertificatemanagementRootCaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatemanagementRootCaCertificateWithDefaults() *CertificatemanagementRootCaCertificate {
	this := CertificatemanagementRootCaCertificate{}
	var classId string = "certificatemanagement.RootCaCertificate"
	this.ClassId = classId
	var objectType string = "certificatemanagement.RootCaCertificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CertificatemanagementRootCaCertificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementRootCaCertificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CertificatemanagementRootCaCertificate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "certificatemanagement.RootCaCertificate" of the ClassId field.
func (o *CertificatemanagementRootCaCertificate) GetDefaultClassId() interface{} {
	return "certificatemanagement.RootCaCertificate"
}

// GetObjectType returns the ObjectType field value
func (o *CertificatemanagementRootCaCertificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CertificatemanagementRootCaCertificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CertificatemanagementRootCaCertificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "certificatemanagement.RootCaCertificate" of the ObjectType field.
func (o *CertificatemanagementRootCaCertificate) GetDefaultObjectType() interface{} {
	return "certificatemanagement.RootCaCertificate"
}

// GetCertificateName returns the CertificateName field value if set, zero value otherwise.
func (o *CertificatemanagementRootCaCertificate) GetCertificateName() string {
	if o == nil || IsNil(o.CertificateName) {
		var ret string
		return ret
	}
	return *o.CertificateName
}

// GetCertificateNameOk returns a tuple with the CertificateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificatemanagementRootCaCertificate) GetCertificateNameOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateName) {
		return nil, false
	}
	return o.CertificateName, true
}

// HasCertificateName returns a boolean if a field has been set.
func (o *CertificatemanagementRootCaCertificate) HasCertificateName() bool {
	if o != nil && !IsNil(o.CertificateName) {
		return true
	}

	return false
}

// SetCertificateName gets a reference to the given string and assigns it to the CertificateName field.
func (o *CertificatemanagementRootCaCertificate) SetCertificateName(v string) {
	o.CertificateName = &v
}

func (o CertificatemanagementRootCaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificatemanagementRootCaCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCertificatemanagementCertificateBase, errCertificatemanagementCertificateBase := json.Marshal(o.CertificatemanagementCertificateBase)
	if errCertificatemanagementCertificateBase != nil {
		return map[string]interface{}{}, errCertificatemanagementCertificateBase
	}
	errCertificatemanagementCertificateBase = json.Unmarshal([]byte(serializedCertificatemanagementCertificateBase), &toSerialize)
	if errCertificatemanagementCertificateBase != nil {
		return map[string]interface{}{}, errCertificatemanagementCertificateBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.CertificateName) {
		toSerialize["CertificateName"] = o.CertificateName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificatemanagementRootCaCertificate) UnmarshalJSON(data []byte) (err error) {
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
	type CertificatemanagementRootCaCertificateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// A name that helps identify a certificate. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
		CertificateName *string `json:"CertificateName,omitempty" validate:"regexp=((^[a-zA-Z0-9]$){1,30}|(^(([a-zA-Z0-9])([a-zA-Z0-9_\\\\-]{0,28})([a-zA-Z0-9]))$))"`
	}

	varCertificatemanagementRootCaCertificateWithoutEmbeddedStruct := CertificatemanagementRootCaCertificateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCertificatemanagementRootCaCertificateWithoutEmbeddedStruct)
	if err == nil {
		varCertificatemanagementRootCaCertificate := _CertificatemanagementRootCaCertificate{}
		varCertificatemanagementRootCaCertificate.ClassId = varCertificatemanagementRootCaCertificateWithoutEmbeddedStruct.ClassId
		varCertificatemanagementRootCaCertificate.ObjectType = varCertificatemanagementRootCaCertificateWithoutEmbeddedStruct.ObjectType
		varCertificatemanagementRootCaCertificate.CertificateName = varCertificatemanagementRootCaCertificateWithoutEmbeddedStruct.CertificateName
		*o = CertificatemanagementRootCaCertificate(varCertificatemanagementRootCaCertificate)
	} else {
		return err
	}

	varCertificatemanagementRootCaCertificate := _CertificatemanagementRootCaCertificate{}

	err = json.Unmarshal(data, &varCertificatemanagementRootCaCertificate)
	if err == nil {
		o.CertificatemanagementCertificateBase = varCertificatemanagementRootCaCertificate.CertificatemanagementCertificateBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CertificateName")

		// remove fields from embedded structs
		reflectCertificatemanagementCertificateBase := reflect.ValueOf(o.CertificatemanagementCertificateBase)
		for i := 0; i < reflectCertificatemanagementCertificateBase.Type().NumField(); i++ {
			t := reflectCertificatemanagementCertificateBase.Type().Field(i)

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

type NullableCertificatemanagementRootCaCertificate struct {
	value *CertificatemanagementRootCaCertificate
	isSet bool
}

func (v NullableCertificatemanagementRootCaCertificate) Get() *CertificatemanagementRootCaCertificate {
	return v.value
}

func (v *NullableCertificatemanagementRootCaCertificate) Set(val *CertificatemanagementRootCaCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatemanagementRootCaCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatemanagementRootCaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatemanagementRootCaCertificate(val *CertificatemanagementRootCaCertificate) *NullableCertificatemanagementRootCaCertificate {
	return &NullableCertificatemanagementRootCaCertificate{value: val, isSet: true}
}

func (v NullableCertificatemanagementRootCaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatemanagementRootCaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
