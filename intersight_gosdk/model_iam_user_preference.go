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

// IamUserPreference Holder for UI preferences such as theme, columns.
type IamUserPreference struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// UI preferences of the user.
	Preference interface{} `json:"Preference,omitempty"`
	// Unique id of the user used by the identity provider to store the user.
	UserUniqueIdentifier *string                      `json:"UserUniqueIdentifier,omitempty"`
	Idp                  *IamIdpRelationship          `json:"Idp,omitempty"`
	IdpReference         *IamIdpReferenceRelationship `json:"IdpReference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUserPreference IamUserPreference

// NewIamUserPreference instantiates a new IamUserPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserPreference(classId string, objectType string) *IamUserPreference {
	this := IamUserPreference{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamUserPreferenceWithDefaults instantiates a new IamUserPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserPreferenceWithDefaults() *IamUserPreference {
	this := IamUserPreference{}
	var classId string = "iam.UserPreference"
	this.ClassId = classId
	var objectType string = "iam.UserPreference"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUserPreference) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUserPreference) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUserPreference) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamUserPreference) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUserPreference) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUserPreference) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPreference returns the Preference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserPreference) GetPreference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserPreference) GetPreferenceOk() (*interface{}, bool) {
	if o == nil || o.Preference == nil {
		return nil, false
	}
	return &o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *IamUserPreference) HasPreference() bool {
	if o != nil && o.Preference != nil {
		return true
	}

	return false
}

// SetPreference gets a reference to the given interface{} and assigns it to the Preference field.
func (o *IamUserPreference) SetPreference(v interface{}) {
	o.Preference = v
}

// GetUserUniqueIdentifier returns the UserUniqueIdentifier field value if set, zero value otherwise.
func (o *IamUserPreference) GetUserUniqueIdentifier() string {
	if o == nil || o.UserUniqueIdentifier == nil {
		var ret string
		return ret
	}
	return *o.UserUniqueIdentifier
}

// GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreference) GetUserUniqueIdentifierOk() (*string, bool) {
	if o == nil || o.UserUniqueIdentifier == nil {
		return nil, false
	}
	return o.UserUniqueIdentifier, true
}

// HasUserUniqueIdentifier returns a boolean if a field has been set.
func (o *IamUserPreference) HasUserUniqueIdentifier() bool {
	if o != nil && o.UserUniqueIdentifier != nil {
		return true
	}

	return false
}

// SetUserUniqueIdentifier gets a reference to the given string and assigns it to the UserUniqueIdentifier field.
func (o *IamUserPreference) SetUserUniqueIdentifier(v string) {
	o.UserUniqueIdentifier = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IamUserPreference) GetIdp() IamIdpRelationship {
	if o == nil || o.Idp == nil {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreference) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUserPreference) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given IamIdpRelationship and assigns it to the Idp field.
func (o *IamUserPreference) SetIdp(v IamIdpRelationship) {
	o.Idp = &v
}

// GetIdpReference returns the IdpReference field value if set, zero value otherwise.
func (o *IamUserPreference) GetIdpReference() IamIdpReferenceRelationship {
	if o == nil || o.IdpReference == nil {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.IdpReference
}

// GetIdpReferenceOk returns a tuple with the IdpReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserPreference) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil || o.IdpReference == nil {
		return nil, false
	}
	return o.IdpReference, true
}

// HasIdpReference returns a boolean if a field has been set.
func (o *IamUserPreference) HasIdpReference() bool {
	if o != nil && o.IdpReference != nil {
		return true
	}

	return false
}

// SetIdpReference gets a reference to the given IamIdpReferenceRelationship and assigns it to the IdpReference field.
func (o *IamUserPreference) SetIdpReference(v IamIdpReferenceRelationship) {
	o.IdpReference = &v
}

func (o IamUserPreference) MarshalJSON() ([]byte, error) {
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
	if o.Preference != nil {
		toSerialize["Preference"] = o.Preference
	}
	if o.UserUniqueIdentifier != nil {
		toSerialize["UserUniqueIdentifier"] = o.UserUniqueIdentifier
	}
	if o.Idp != nil {
		toSerialize["Idp"] = o.Idp
	}
	if o.IdpReference != nil {
		toSerialize["IdpReference"] = o.IdpReference
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamUserPreference) UnmarshalJSON(bytes []byte) (err error) {
	type IamUserPreferenceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// UI preferences of the user.
		Preference interface{} `json:"Preference,omitempty"`
		// Unique id of the user used by the identity provider to store the user.
		UserUniqueIdentifier *string                      `json:"UserUniqueIdentifier,omitempty"`
		Idp                  *IamIdpRelationship          `json:"Idp,omitempty"`
		IdpReference         *IamIdpReferenceRelationship `json:"IdpReference,omitempty"`
	}

	varIamUserPreferenceWithoutEmbeddedStruct := IamUserPreferenceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamUserPreferenceWithoutEmbeddedStruct)
	if err == nil {
		varIamUserPreference := _IamUserPreference{}
		varIamUserPreference.ClassId = varIamUserPreferenceWithoutEmbeddedStruct.ClassId
		varIamUserPreference.ObjectType = varIamUserPreferenceWithoutEmbeddedStruct.ObjectType
		varIamUserPreference.Preference = varIamUserPreferenceWithoutEmbeddedStruct.Preference
		varIamUserPreference.UserUniqueIdentifier = varIamUserPreferenceWithoutEmbeddedStruct.UserUniqueIdentifier
		varIamUserPreference.Idp = varIamUserPreferenceWithoutEmbeddedStruct.Idp
		varIamUserPreference.IdpReference = varIamUserPreferenceWithoutEmbeddedStruct.IdpReference
		*o = IamUserPreference(varIamUserPreference)
	} else {
		return err
	}

	varIamUserPreference := _IamUserPreference{}

	err = json.Unmarshal(bytes, &varIamUserPreference)
	if err == nil {
		o.MoBaseMo = varIamUserPreference.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Preference")
		delete(additionalProperties, "UserUniqueIdentifier")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "IdpReference")

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

type NullableIamUserPreference struct {
	value *IamUserPreference
	isSet bool
}

func (v NullableIamUserPreference) Get() *IamUserPreference {
	return v.value
}

func (v *NullableIamUserPreference) Set(val *IamUserPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUserPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUserPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUserPreference(val *IamUserPreference) *NullableIamUserPreference {
	return &NullableIamUserPreference{value: val, isSet: true}
}

func (v NullableIamUserPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUserPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
