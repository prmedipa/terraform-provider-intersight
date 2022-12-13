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
	"reflect"
	"strings"
)

// IamSessionLimits The session related configuration limits.
type IamSessionLimits struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds.
	IdleTimeOut *int64 `json:"IdleTimeOut,omitempty"`
	// The maximum number of sessions allowed in an account or permission. The default value is 128.
	MaximumLimit *int64 `json:"MaximumLimit,omitempty"`
	// The maximum number of sessions allowed per user. Default value is 32.
	PerUserLimit *int64 `json:"PerUserLimit,omitempty"`
	// The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds.
	SessionTimeOut       *int64                     `json:"SessionTimeOut,omitempty"`
	Account              *IamAccountRelationship    `json:"Account,omitempty"`
	Permission           *IamPermissionRelationship `json:"Permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSessionLimits IamSessionLimits

// NewIamSessionLimits instantiates a new IamSessionLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSessionLimits(classId string, objectType string) *IamSessionLimits {
	this := IamSessionLimits{}
	this.ClassId = classId
	this.ObjectType = objectType
	var idleTimeOut int64 = 1800
	this.IdleTimeOut = &idleTimeOut
	var maximumLimit int64 = 128
	this.MaximumLimit = &maximumLimit
	var perUserLimit int64 = 32
	this.PerUserLimit = &perUserLimit
	var sessionTimeOut int64 = 57600
	this.SessionTimeOut = &sessionTimeOut
	return &this
}

// NewIamSessionLimitsWithDefaults instantiates a new IamSessionLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSessionLimitsWithDefaults() *IamSessionLimits {
	this := IamSessionLimits{}
	var classId string = "iam.SessionLimits"
	this.ClassId = classId
	var objectType string = "iam.SessionLimits"
	this.ObjectType = objectType
	var idleTimeOut int64 = 1800
	this.IdleTimeOut = &idleTimeOut
	var maximumLimit int64 = 128
	this.MaximumLimit = &maximumLimit
	var perUserLimit int64 = 32
	this.PerUserLimit = &perUserLimit
	var sessionTimeOut int64 = 57600
	this.SessionTimeOut = &sessionTimeOut
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSessionLimits) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSessionLimits) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSessionLimits) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSessionLimits) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIdleTimeOut returns the IdleTimeOut field value if set, zero value otherwise.
func (o *IamSessionLimits) GetIdleTimeOut() int64 {
	if o == nil || o.IdleTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.IdleTimeOut
}

// GetIdleTimeOutOk returns a tuple with the IdleTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetIdleTimeOutOk() (*int64, bool) {
	if o == nil || o.IdleTimeOut == nil {
		return nil, false
	}
	return o.IdleTimeOut, true
}

// HasIdleTimeOut returns a boolean if a field has been set.
func (o *IamSessionLimits) HasIdleTimeOut() bool {
	if o != nil && o.IdleTimeOut != nil {
		return true
	}

	return false
}

// SetIdleTimeOut gets a reference to the given int64 and assigns it to the IdleTimeOut field.
func (o *IamSessionLimits) SetIdleTimeOut(v int64) {
	o.IdleTimeOut = &v
}

// GetMaximumLimit returns the MaximumLimit field value if set, zero value otherwise.
func (o *IamSessionLimits) GetMaximumLimit() int64 {
	if o == nil || o.MaximumLimit == nil {
		var ret int64
		return ret
	}
	return *o.MaximumLimit
}

// GetMaximumLimitOk returns a tuple with the MaximumLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetMaximumLimitOk() (*int64, bool) {
	if o == nil || o.MaximumLimit == nil {
		return nil, false
	}
	return o.MaximumLimit, true
}

// HasMaximumLimit returns a boolean if a field has been set.
func (o *IamSessionLimits) HasMaximumLimit() bool {
	if o != nil && o.MaximumLimit != nil {
		return true
	}

	return false
}

// SetMaximumLimit gets a reference to the given int64 and assigns it to the MaximumLimit field.
func (o *IamSessionLimits) SetMaximumLimit(v int64) {
	o.MaximumLimit = &v
}

// GetPerUserLimit returns the PerUserLimit field value if set, zero value otherwise.
func (o *IamSessionLimits) GetPerUserLimit() int64 {
	if o == nil || o.PerUserLimit == nil {
		var ret int64
		return ret
	}
	return *o.PerUserLimit
}

// GetPerUserLimitOk returns a tuple with the PerUserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetPerUserLimitOk() (*int64, bool) {
	if o == nil || o.PerUserLimit == nil {
		return nil, false
	}
	return o.PerUserLimit, true
}

// HasPerUserLimit returns a boolean if a field has been set.
func (o *IamSessionLimits) HasPerUserLimit() bool {
	if o != nil && o.PerUserLimit != nil {
		return true
	}

	return false
}

// SetPerUserLimit gets a reference to the given int64 and assigns it to the PerUserLimit field.
func (o *IamSessionLimits) SetPerUserLimit(v int64) {
	o.PerUserLimit = &v
}

// GetSessionTimeOut returns the SessionTimeOut field value if set, zero value otherwise.
func (o *IamSessionLimits) GetSessionTimeOut() int64 {
	if o == nil || o.SessionTimeOut == nil {
		var ret int64
		return ret
	}
	return *o.SessionTimeOut
}

// GetSessionTimeOutOk returns a tuple with the SessionTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetSessionTimeOutOk() (*int64, bool) {
	if o == nil || o.SessionTimeOut == nil {
		return nil, false
	}
	return o.SessionTimeOut, true
}

// HasSessionTimeOut returns a boolean if a field has been set.
func (o *IamSessionLimits) HasSessionTimeOut() bool {
	if o != nil && o.SessionTimeOut != nil {
		return true
	}

	return false
}

// SetSessionTimeOut gets a reference to the given int64 and assigns it to the SessionTimeOut field.
func (o *IamSessionLimits) SetSessionTimeOut(v int64) {
	o.SessionTimeOut = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamSessionLimits) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamSessionLimits) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamSessionLimits) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *IamSessionLimits) GetPermission() IamPermissionRelationship {
	if o == nil || o.Permission == nil {
		var ret IamPermissionRelationship
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSessionLimits) GetPermissionOk() (*IamPermissionRelationship, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *IamSessionLimits) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given IamPermissionRelationship and assigns it to the Permission field.
func (o *IamSessionLimits) SetPermission(v IamPermissionRelationship) {
	o.Permission = &v
}

func (o IamSessionLimits) MarshalJSON() ([]byte, error) {
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
	if o.IdleTimeOut != nil {
		toSerialize["IdleTimeOut"] = o.IdleTimeOut
	}
	if o.MaximumLimit != nil {
		toSerialize["MaximumLimit"] = o.MaximumLimit
	}
	if o.PerUserLimit != nil {
		toSerialize["PerUserLimit"] = o.PerUserLimit
	}
	if o.SessionTimeOut != nil {
		toSerialize["SessionTimeOut"] = o.SessionTimeOut
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSessionLimits) UnmarshalJSON(bytes []byte) (err error) {
	type IamSessionLimitsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds.
		IdleTimeOut *int64 `json:"IdleTimeOut,omitempty"`
		// The maximum number of sessions allowed in an account or permission. The default value is 128.
		MaximumLimit *int64 `json:"MaximumLimit,omitempty"`
		// The maximum number of sessions allowed per user. Default value is 32.
		PerUserLimit *int64 `json:"PerUserLimit,omitempty"`
		// The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds.
		SessionTimeOut *int64                     `json:"SessionTimeOut,omitempty"`
		Account        *IamAccountRelationship    `json:"Account,omitempty"`
		Permission     *IamPermissionRelationship `json:"Permission,omitempty"`
	}

	varIamSessionLimitsWithoutEmbeddedStruct := IamSessionLimitsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamSessionLimitsWithoutEmbeddedStruct)
	if err == nil {
		varIamSessionLimits := _IamSessionLimits{}
		varIamSessionLimits.ClassId = varIamSessionLimitsWithoutEmbeddedStruct.ClassId
		varIamSessionLimits.ObjectType = varIamSessionLimitsWithoutEmbeddedStruct.ObjectType
		varIamSessionLimits.IdleTimeOut = varIamSessionLimitsWithoutEmbeddedStruct.IdleTimeOut
		varIamSessionLimits.MaximumLimit = varIamSessionLimitsWithoutEmbeddedStruct.MaximumLimit
		varIamSessionLimits.PerUserLimit = varIamSessionLimitsWithoutEmbeddedStruct.PerUserLimit
		varIamSessionLimits.SessionTimeOut = varIamSessionLimitsWithoutEmbeddedStruct.SessionTimeOut
		varIamSessionLimits.Account = varIamSessionLimitsWithoutEmbeddedStruct.Account
		varIamSessionLimits.Permission = varIamSessionLimitsWithoutEmbeddedStruct.Permission
		*o = IamSessionLimits(varIamSessionLimits)
	} else {
		return err
	}

	varIamSessionLimits := _IamSessionLimits{}

	err = json.Unmarshal(bytes, &varIamSessionLimits)
	if err == nil {
		o.MoBaseMo = varIamSessionLimits.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IdleTimeOut")
		delete(additionalProperties, "MaximumLimit")
		delete(additionalProperties, "PerUserLimit")
		delete(additionalProperties, "SessionTimeOut")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Permission")

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

type NullableIamSessionLimits struct {
	value *IamSessionLimits
	isSet bool
}

func (v NullableIamSessionLimits) Get() *IamSessionLimits {
	return v.value
}

func (v *NullableIamSessionLimits) Set(val *IamSessionLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSessionLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSessionLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSessionLimits(val *IamSessionLimits) *NullableIamSessionLimits {
	return &NullableIamSessionLimits{value: val, isSet: true}
}

func (v NullableIamSessionLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSessionLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
