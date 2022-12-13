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

// LicenseLicenseInfoView Custom LicenseInfo view list to be displayed by the UI.
type LicenseLicenseInfoView struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                 `json:"ObjectType"`
	AccountLicenseData   *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LicenseLicenseInfoView LicenseLicenseInfoView

// NewLicenseLicenseInfoView instantiates a new LicenseLicenseInfoView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseLicenseInfoView(classId string, objectType string) *LicenseLicenseInfoView {
	this := LicenseLicenseInfoView{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewLicenseLicenseInfoViewWithDefaults instantiates a new LicenseLicenseInfoView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseLicenseInfoViewWithDefaults() *LicenseLicenseInfoView {
	this := LicenseLicenseInfoView{}
	var classId string = "license.LicenseInfoView"
	this.ClassId = classId
	var objectType string = "license.LicenseInfoView"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *LicenseLicenseInfoView) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfoView) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *LicenseLicenseInfoView) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *LicenseLicenseInfoView) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfoView) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *LicenseLicenseInfoView) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccountLicenseData returns the AccountLicenseData field value if set, zero value otherwise.
func (o *LicenseLicenseInfoView) GetAccountLicenseData() LicenseAccountLicenseDataRelationship {
	if o == nil || o.AccountLicenseData == nil {
		var ret LicenseAccountLicenseDataRelationship
		return ret
	}
	return *o.AccountLicenseData
}

// GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseLicenseInfoView) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool) {
	if o == nil || o.AccountLicenseData == nil {
		return nil, false
	}
	return o.AccountLicenseData, true
}

// HasAccountLicenseData returns a boolean if a field has been set.
func (o *LicenseLicenseInfoView) HasAccountLicenseData() bool {
	if o != nil && o.AccountLicenseData != nil {
		return true
	}

	return false
}

// SetAccountLicenseData gets a reference to the given LicenseAccountLicenseDataRelationship and assigns it to the AccountLicenseData field.
func (o *LicenseLicenseInfoView) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship) {
	o.AccountLicenseData = &v
}

func (o LicenseLicenseInfoView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AccountLicenseData != nil {
		toSerialize["AccountLicenseData"] = o.AccountLicenseData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LicenseLicenseInfoView) UnmarshalJSON(bytes []byte) (err error) {
	type LicenseLicenseInfoViewWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType         string                                 `json:"ObjectType"`
		AccountLicenseData *LicenseAccountLicenseDataRelationship `json:"AccountLicenseData,omitempty"`
	}

	varLicenseLicenseInfoViewWithoutEmbeddedStruct := LicenseLicenseInfoViewWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varLicenseLicenseInfoViewWithoutEmbeddedStruct)
	if err == nil {
		varLicenseLicenseInfoView := _LicenseLicenseInfoView{}
		varLicenseLicenseInfoView.ClassId = varLicenseLicenseInfoViewWithoutEmbeddedStruct.ClassId
		varLicenseLicenseInfoView.ObjectType = varLicenseLicenseInfoViewWithoutEmbeddedStruct.ObjectType
		varLicenseLicenseInfoView.AccountLicenseData = varLicenseLicenseInfoViewWithoutEmbeddedStruct.AccountLicenseData
		*o = LicenseLicenseInfoView(varLicenseLicenseInfoView)
	} else {
		return err
	}

	varLicenseLicenseInfoView := _LicenseLicenseInfoView{}

	err = json.Unmarshal(bytes, &varLicenseLicenseInfoView)
	if err == nil {
		o.MoBaseMo = varLicenseLicenseInfoView.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccountLicenseData")

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

type NullableLicenseLicenseInfoView struct {
	value *LicenseLicenseInfoView
	isSet bool
}

func (v NullableLicenseLicenseInfoView) Get() *LicenseLicenseInfoView {
	return v.value
}

func (v *NullableLicenseLicenseInfoView) Set(val *LicenseLicenseInfoView) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseLicenseInfoView) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseLicenseInfoView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseLicenseInfoView(val *LicenseLicenseInfoView) *NullableLicenseLicenseInfoView {
	return &NullableLicenseLicenseInfoView{value: val, isSet: true}
}

func (v NullableLicenseLicenseInfoView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseLicenseInfoView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
