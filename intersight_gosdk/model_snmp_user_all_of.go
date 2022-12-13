/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// SnmpUserAllOf Definition of the list of properties defined in 'snmp.User', excluding properties defined in parent classes.
type SnmpUserAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authorization password for the user.
	AuthPassword *string `json:"AuthPassword,omitempty"`
	// Authorization protocol for authenticating the user. * `NA` - Authentication protocol is not applicable. * `MD5` - MD5 protocol is used to authenticate SNMP user. * `SHA` - SHA protocol is used to authenticate SNMP user. * `SHA-224` - SHA-224 protocol is used to authenticate SNMP user. * `SHA-256` - SHA-256 protocol is used to authenticate SNMP user. * `SHA-384` - SHA-384 protocol is used to authenticate SNMP user. * `SHA-512` - SHA-512 protocol is used to authenticate SNMP user.
	AuthType *string `json:"AuthType,omitempty"`
	// Indicates whether the value of the 'authPassword' property has been set.
	IsAuthPasswordSet *bool `json:"IsAuthPasswordSet,omitempty"`
	// Indicates whether the value of the 'privacyPassword' property has been set.
	IsPrivacyPasswordSet *bool `json:"IsPrivacyPasswordSet,omitempty"`
	// SNMP username. Must have a minimum of 1 and and a maximum of 31 characters.
	Name *string `json:"Name,omitempty"`
	// Privacy password for the user.
	PrivacyPassword *string `json:"PrivacyPassword,omitempty"`
	// Privacy protocol for the user. * `NA` - Privacy protocol is not applicable. * `DES` - DES privacy protocol is used for SNMP user. * `AES` - AES privacy protocol is used for SNMP user.
	PrivacyType *string `json:"PrivacyType,omitempty"`
	// Security mechanism used for communication between agent and manager. * `AuthPriv` - The user requires both an authorization password and a privacy password. * `NoAuthNoPriv` - The user does not require an authorization or privacy password. * `AuthNoPriv` - The user requires an authorization password but not a privacy password.
	SecurityLevel        *string `json:"SecurityLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SnmpUserAllOf SnmpUserAllOf

// NewSnmpUserAllOf instantiates a new SnmpUserAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnmpUserAllOf(classId string, objectType string) *SnmpUserAllOf {
	this := SnmpUserAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var authType string = "NA"
	this.AuthType = &authType
	var privacyType string = "NA"
	this.PrivacyType = &privacyType
	var securityLevel string = "AuthPriv"
	this.SecurityLevel = &securityLevel
	return &this
}

// NewSnmpUserAllOfWithDefaults instantiates a new SnmpUserAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnmpUserAllOfWithDefaults() *SnmpUserAllOf {
	this := SnmpUserAllOf{}
	var classId string = "snmp.User"
	this.ClassId = classId
	var objectType string = "snmp.User"
	this.ObjectType = objectType
	var authType string = "NA"
	this.AuthType = &authType
	var privacyType string = "NA"
	this.PrivacyType = &privacyType
	var securityLevel string = "AuthPriv"
	this.SecurityLevel = &securityLevel
	return &this
}

// GetClassId returns the ClassId field value
func (o *SnmpUserAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SnmpUserAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SnmpUserAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SnmpUserAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetAuthPassword() string {
	if o == nil || o.AuthPassword == nil {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetAuthPasswordOk() (*string, bool) {
	if o == nil || o.AuthPassword == nil {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasAuthPassword() bool {
	if o != nil && o.AuthPassword != nil {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *SnmpUserAllOf) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *SnmpUserAllOf) SetAuthType(v string) {
	o.AuthType = &v
}

// GetIsAuthPasswordSet returns the IsAuthPasswordSet field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetIsAuthPasswordSet() bool {
	if o == nil || o.IsAuthPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsAuthPasswordSet
}

// GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetIsAuthPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsAuthPasswordSet == nil {
		return nil, false
	}
	return o.IsAuthPasswordSet, true
}

// HasIsAuthPasswordSet returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasIsAuthPasswordSet() bool {
	if o != nil && o.IsAuthPasswordSet != nil {
		return true
	}

	return false
}

// SetIsAuthPasswordSet gets a reference to the given bool and assigns it to the IsAuthPasswordSet field.
func (o *SnmpUserAllOf) SetIsAuthPasswordSet(v bool) {
	o.IsAuthPasswordSet = &v
}

// GetIsPrivacyPasswordSet returns the IsPrivacyPasswordSet field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetIsPrivacyPasswordSet() bool {
	if o == nil || o.IsPrivacyPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPrivacyPasswordSet
}

// GetIsPrivacyPasswordSetOk returns a tuple with the IsPrivacyPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetIsPrivacyPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPrivacyPasswordSet == nil {
		return nil, false
	}
	return o.IsPrivacyPasswordSet, true
}

// HasIsPrivacyPasswordSet returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasIsPrivacyPasswordSet() bool {
	if o != nil && o.IsPrivacyPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPrivacyPasswordSet gets a reference to the given bool and assigns it to the IsPrivacyPasswordSet field.
func (o *SnmpUserAllOf) SetIsPrivacyPasswordSet(v bool) {
	o.IsPrivacyPasswordSet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnmpUserAllOf) SetName(v string) {
	o.Name = &v
}

// GetPrivacyPassword returns the PrivacyPassword field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetPrivacyPassword() string {
	if o == nil || o.PrivacyPassword == nil {
		var ret string
		return ret
	}
	return *o.PrivacyPassword
}

// GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetPrivacyPasswordOk() (*string, bool) {
	if o == nil || o.PrivacyPassword == nil {
		return nil, false
	}
	return o.PrivacyPassword, true
}

// HasPrivacyPassword returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasPrivacyPassword() bool {
	if o != nil && o.PrivacyPassword != nil {
		return true
	}

	return false
}

// SetPrivacyPassword gets a reference to the given string and assigns it to the PrivacyPassword field.
func (o *SnmpUserAllOf) SetPrivacyPassword(v string) {
	o.PrivacyPassword = &v
}

// GetPrivacyType returns the PrivacyType field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetPrivacyType() string {
	if o == nil || o.PrivacyType == nil {
		var ret string
		return ret
	}
	return *o.PrivacyType
}

// GetPrivacyTypeOk returns a tuple with the PrivacyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetPrivacyTypeOk() (*string, bool) {
	if o == nil || o.PrivacyType == nil {
		return nil, false
	}
	return o.PrivacyType, true
}

// HasPrivacyType returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasPrivacyType() bool {
	if o != nil && o.PrivacyType != nil {
		return true
	}

	return false
}

// SetPrivacyType gets a reference to the given string and assigns it to the PrivacyType field.
func (o *SnmpUserAllOf) SetPrivacyType(v string) {
	o.PrivacyType = &v
}

// GetSecurityLevel returns the SecurityLevel field value if set, zero value otherwise.
func (o *SnmpUserAllOf) GetSecurityLevel() string {
	if o == nil || o.SecurityLevel == nil {
		var ret string
		return ret
	}
	return *o.SecurityLevel
}

// GetSecurityLevelOk returns a tuple with the SecurityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnmpUserAllOf) GetSecurityLevelOk() (*string, bool) {
	if o == nil || o.SecurityLevel == nil {
		return nil, false
	}
	return o.SecurityLevel, true
}

// HasSecurityLevel returns a boolean if a field has been set.
func (o *SnmpUserAllOf) HasSecurityLevel() bool {
	if o != nil && o.SecurityLevel != nil {
		return true
	}

	return false
}

// SetSecurityLevel gets a reference to the given string and assigns it to the SecurityLevel field.
func (o *SnmpUserAllOf) SetSecurityLevel(v string) {
	o.SecurityLevel = &v
}

func (o SnmpUserAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthPassword != nil {
		toSerialize["AuthPassword"] = o.AuthPassword
	}
	if o.AuthType != nil {
		toSerialize["AuthType"] = o.AuthType
	}
	if o.IsAuthPasswordSet != nil {
		toSerialize["IsAuthPasswordSet"] = o.IsAuthPasswordSet
	}
	if o.IsPrivacyPasswordSet != nil {
		toSerialize["IsPrivacyPasswordSet"] = o.IsPrivacyPasswordSet
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.PrivacyPassword != nil {
		toSerialize["PrivacyPassword"] = o.PrivacyPassword
	}
	if o.PrivacyType != nil {
		toSerialize["PrivacyType"] = o.PrivacyType
	}
	if o.SecurityLevel != nil {
		toSerialize["SecurityLevel"] = o.SecurityLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SnmpUserAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnmpUserAllOf := _SnmpUserAllOf{}

	if err = json.Unmarshal(bytes, &varSnmpUserAllOf); err == nil {
		*o = SnmpUserAllOf(varSnmpUserAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthPassword")
		delete(additionalProperties, "AuthType")
		delete(additionalProperties, "IsAuthPasswordSet")
		delete(additionalProperties, "IsPrivacyPasswordSet")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "PrivacyPassword")
		delete(additionalProperties, "PrivacyType")
		delete(additionalProperties, "SecurityLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSnmpUserAllOf struct {
	value *SnmpUserAllOf
	isSet bool
}

func (v NullableSnmpUserAllOf) Get() *SnmpUserAllOf {
	return v.value
}

func (v *NullableSnmpUserAllOf) Set(val *SnmpUserAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSnmpUserAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSnmpUserAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnmpUserAllOf(val *SnmpUserAllOf) *NullableSnmpUserAllOf {
	return &NullableSnmpUserAllOf{value: val, isSet: true}
}

func (v NullableSnmpUserAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnmpUserAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
