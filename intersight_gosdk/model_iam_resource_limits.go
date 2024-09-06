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

// checks if the IamResourceLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamResourceLimits{}

// IamResourceLimits The resource limits used to limit resources such as User accounts.
type IamResourceLimits struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Boolean value used to decide whether API keys that never expire are allowed for the account. This allows creation of API keys which are perpetual which can used for specific applications where rotation of API keys are not feasible.
	AllowApiKeysWithoutExpiry *bool `json:"AllowApiKeysWithoutExpiry,omitempty"`
	// Boolean value used to decide whether App Registration that never expire are allowed for the account.
	AllowAppRegistrationsWithoutExpiry *bool `json:"AllowAppRegistrationsWithoutExpiry,omitempty"`
	// The maximum expiration period (in seconds) allowed for API keys. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability.
	MaxApiKeyExpiry *int64 `json:"MaxApiKeyExpiry,omitempty"`
	// The maximum expiration period (in seconds) allowed for App Registration. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability.
	MaxAppRegistrationExpiry *int64 `json:"MaxAppRegistrationExpiry,omitempty"`
	// The maximum number of users allowed in an account. The default value is 200.
	PerAccountUserLimit  *int64                         `json:"PerAccountUserLimit,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamResourceLimits IamResourceLimits

// NewIamResourceLimits instantiates a new IamResourceLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamResourceLimits(classId string, objectType string) *IamResourceLimits {
	this := IamResourceLimits{}
	this.ClassId = classId
	this.ObjectType = objectType
	var allowApiKeysWithoutExpiry bool = false
	this.AllowApiKeysWithoutExpiry = &allowApiKeysWithoutExpiry
	var allowAppRegistrationsWithoutExpiry bool = false
	this.AllowAppRegistrationsWithoutExpiry = &allowAppRegistrationsWithoutExpiry
	var maxApiKeyExpiry int64 = 15552000
	this.MaxApiKeyExpiry = &maxApiKeyExpiry
	var maxAppRegistrationExpiry int64 = 15552000
	this.MaxAppRegistrationExpiry = &maxAppRegistrationExpiry
	return &this
}

// NewIamResourceLimitsWithDefaults instantiates a new IamResourceLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamResourceLimitsWithDefaults() *IamResourceLimits {
	this := IamResourceLimits{}
	var classId string = "iam.ResourceLimits"
	this.ClassId = classId
	var objectType string = "iam.ResourceLimits"
	this.ObjectType = objectType
	var allowApiKeysWithoutExpiry bool = false
	this.AllowApiKeysWithoutExpiry = &allowApiKeysWithoutExpiry
	var allowAppRegistrationsWithoutExpiry bool = false
	this.AllowAppRegistrationsWithoutExpiry = &allowAppRegistrationsWithoutExpiry
	var maxApiKeyExpiry int64 = 15552000
	this.MaxApiKeyExpiry = &maxApiKeyExpiry
	var maxAppRegistrationExpiry int64 = 15552000
	this.MaxAppRegistrationExpiry = &maxAppRegistrationExpiry
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamResourceLimits) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamResourceLimits) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.ResourceLimits" of the ClassId field.
func (o *IamResourceLimits) GetDefaultClassId() interface{} {
	return "iam.ResourceLimits"
}

// GetObjectType returns the ObjectType field value
func (o *IamResourceLimits) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamResourceLimits) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.ResourceLimits" of the ObjectType field.
func (o *IamResourceLimits) GetDefaultObjectType() interface{} {
	return "iam.ResourceLimits"
}

// GetAllowApiKeysWithoutExpiry returns the AllowApiKeysWithoutExpiry field value if set, zero value otherwise.
func (o *IamResourceLimits) GetAllowApiKeysWithoutExpiry() bool {
	if o == nil || IsNil(o.AllowApiKeysWithoutExpiry) {
		var ret bool
		return ret
	}
	return *o.AllowApiKeysWithoutExpiry
}

// GetAllowApiKeysWithoutExpiryOk returns a tuple with the AllowApiKeysWithoutExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetAllowApiKeysWithoutExpiryOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowApiKeysWithoutExpiry) {
		return nil, false
	}
	return o.AllowApiKeysWithoutExpiry, true
}

// HasAllowApiKeysWithoutExpiry returns a boolean if a field has been set.
func (o *IamResourceLimits) HasAllowApiKeysWithoutExpiry() bool {
	if o != nil && !IsNil(o.AllowApiKeysWithoutExpiry) {
		return true
	}

	return false
}

// SetAllowApiKeysWithoutExpiry gets a reference to the given bool and assigns it to the AllowApiKeysWithoutExpiry field.
func (o *IamResourceLimits) SetAllowApiKeysWithoutExpiry(v bool) {
	o.AllowApiKeysWithoutExpiry = &v
}

// GetAllowAppRegistrationsWithoutExpiry returns the AllowAppRegistrationsWithoutExpiry field value if set, zero value otherwise.
func (o *IamResourceLimits) GetAllowAppRegistrationsWithoutExpiry() bool {
	if o == nil || IsNil(o.AllowAppRegistrationsWithoutExpiry) {
		var ret bool
		return ret
	}
	return *o.AllowAppRegistrationsWithoutExpiry
}

// GetAllowAppRegistrationsWithoutExpiryOk returns a tuple with the AllowAppRegistrationsWithoutExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetAllowAppRegistrationsWithoutExpiryOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAppRegistrationsWithoutExpiry) {
		return nil, false
	}
	return o.AllowAppRegistrationsWithoutExpiry, true
}

// HasAllowAppRegistrationsWithoutExpiry returns a boolean if a field has been set.
func (o *IamResourceLimits) HasAllowAppRegistrationsWithoutExpiry() bool {
	if o != nil && !IsNil(o.AllowAppRegistrationsWithoutExpiry) {
		return true
	}

	return false
}

// SetAllowAppRegistrationsWithoutExpiry gets a reference to the given bool and assigns it to the AllowAppRegistrationsWithoutExpiry field.
func (o *IamResourceLimits) SetAllowAppRegistrationsWithoutExpiry(v bool) {
	o.AllowAppRegistrationsWithoutExpiry = &v
}

// GetMaxApiKeyExpiry returns the MaxApiKeyExpiry field value if set, zero value otherwise.
func (o *IamResourceLimits) GetMaxApiKeyExpiry() int64 {
	if o == nil || IsNil(o.MaxApiKeyExpiry) {
		var ret int64
		return ret
	}
	return *o.MaxApiKeyExpiry
}

// GetMaxApiKeyExpiryOk returns a tuple with the MaxApiKeyExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetMaxApiKeyExpiryOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxApiKeyExpiry) {
		return nil, false
	}
	return o.MaxApiKeyExpiry, true
}

// HasMaxApiKeyExpiry returns a boolean if a field has been set.
func (o *IamResourceLimits) HasMaxApiKeyExpiry() bool {
	if o != nil && !IsNil(o.MaxApiKeyExpiry) {
		return true
	}

	return false
}

// SetMaxApiKeyExpiry gets a reference to the given int64 and assigns it to the MaxApiKeyExpiry field.
func (o *IamResourceLimits) SetMaxApiKeyExpiry(v int64) {
	o.MaxApiKeyExpiry = &v
}

// GetMaxAppRegistrationExpiry returns the MaxAppRegistrationExpiry field value if set, zero value otherwise.
func (o *IamResourceLimits) GetMaxAppRegistrationExpiry() int64 {
	if o == nil || IsNil(o.MaxAppRegistrationExpiry) {
		var ret int64
		return ret
	}
	return *o.MaxAppRegistrationExpiry
}

// GetMaxAppRegistrationExpiryOk returns a tuple with the MaxAppRegistrationExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetMaxAppRegistrationExpiryOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxAppRegistrationExpiry) {
		return nil, false
	}
	return o.MaxAppRegistrationExpiry, true
}

// HasMaxAppRegistrationExpiry returns a boolean if a field has been set.
func (o *IamResourceLimits) HasMaxAppRegistrationExpiry() bool {
	if o != nil && !IsNil(o.MaxAppRegistrationExpiry) {
		return true
	}

	return false
}

// SetMaxAppRegistrationExpiry gets a reference to the given int64 and assigns it to the MaxAppRegistrationExpiry field.
func (o *IamResourceLimits) SetMaxAppRegistrationExpiry(v int64) {
	o.MaxAppRegistrationExpiry = &v
}

// GetPerAccountUserLimit returns the PerAccountUserLimit field value if set, zero value otherwise.
func (o *IamResourceLimits) GetPerAccountUserLimit() int64 {
	if o == nil || IsNil(o.PerAccountUserLimit) {
		var ret int64
		return ret
	}
	return *o.PerAccountUserLimit
}

// GetPerAccountUserLimitOk returns a tuple with the PerAccountUserLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamResourceLimits) GetPerAccountUserLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.PerAccountUserLimit) {
		return nil, false
	}
	return o.PerAccountUserLimit, true
}

// HasPerAccountUserLimit returns a boolean if a field has been set.
func (o *IamResourceLimits) HasPerAccountUserLimit() bool {
	if o != nil && !IsNil(o.PerAccountUserLimit) {
		return true
	}

	return false
}

// SetPerAccountUserLimit gets a reference to the given int64 and assigns it to the PerAccountUserLimit field.
func (o *IamResourceLimits) SetPerAccountUserLimit(v int64) {
	o.PerAccountUserLimit = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamResourceLimits) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamResourceLimits) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamResourceLimits) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamResourceLimits) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamResourceLimits) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamResourceLimits) UnsetAccount() {
	o.Account.Unset()
}

func (o IamResourceLimits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamResourceLimits) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AllowApiKeysWithoutExpiry) {
		toSerialize["AllowApiKeysWithoutExpiry"] = o.AllowApiKeysWithoutExpiry
	}
	if !IsNil(o.AllowAppRegistrationsWithoutExpiry) {
		toSerialize["AllowAppRegistrationsWithoutExpiry"] = o.AllowAppRegistrationsWithoutExpiry
	}
	if !IsNil(o.MaxApiKeyExpiry) {
		toSerialize["MaxApiKeyExpiry"] = o.MaxApiKeyExpiry
	}
	if !IsNil(o.MaxAppRegistrationExpiry) {
		toSerialize["MaxAppRegistrationExpiry"] = o.MaxAppRegistrationExpiry
	}
	if !IsNil(o.PerAccountUserLimit) {
		toSerialize["PerAccountUserLimit"] = o.PerAccountUserLimit
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamResourceLimits) UnmarshalJSON(data []byte) (err error) {
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
	type IamResourceLimitsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Boolean value used to decide whether API keys that never expire are allowed for the account. This allows creation of API keys which are perpetual which can used for specific applications where rotation of API keys are not feasible.
		AllowApiKeysWithoutExpiry *bool `json:"AllowApiKeysWithoutExpiry,omitempty"`
		// Boolean value used to decide whether App Registration that never expire are allowed for the account.
		AllowAppRegistrationsWithoutExpiry *bool `json:"AllowAppRegistrationsWithoutExpiry,omitempty"`
		// The maximum expiration period (in seconds) allowed for API keys. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability.
		MaxApiKeyExpiry *int64 `json:"MaxApiKeyExpiry,omitempty"`
		// The maximum expiration period (in seconds) allowed for App Registration. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability.
		MaxAppRegistrationExpiry *int64 `json:"MaxAppRegistrationExpiry,omitempty"`
		// The maximum number of users allowed in an account. The default value is 200.
		PerAccountUserLimit *int64                         `json:"PerAccountUserLimit,omitempty"`
		Account             NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIamResourceLimitsWithoutEmbeddedStruct := IamResourceLimitsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamResourceLimitsWithoutEmbeddedStruct)
	if err == nil {
		varIamResourceLimits := _IamResourceLimits{}
		varIamResourceLimits.ClassId = varIamResourceLimitsWithoutEmbeddedStruct.ClassId
		varIamResourceLimits.ObjectType = varIamResourceLimitsWithoutEmbeddedStruct.ObjectType
		varIamResourceLimits.AllowApiKeysWithoutExpiry = varIamResourceLimitsWithoutEmbeddedStruct.AllowApiKeysWithoutExpiry
		varIamResourceLimits.AllowAppRegistrationsWithoutExpiry = varIamResourceLimitsWithoutEmbeddedStruct.AllowAppRegistrationsWithoutExpiry
		varIamResourceLimits.MaxApiKeyExpiry = varIamResourceLimitsWithoutEmbeddedStruct.MaxApiKeyExpiry
		varIamResourceLimits.MaxAppRegistrationExpiry = varIamResourceLimitsWithoutEmbeddedStruct.MaxAppRegistrationExpiry
		varIamResourceLimits.PerAccountUserLimit = varIamResourceLimitsWithoutEmbeddedStruct.PerAccountUserLimit
		varIamResourceLimits.Account = varIamResourceLimitsWithoutEmbeddedStruct.Account
		*o = IamResourceLimits(varIamResourceLimits)
	} else {
		return err
	}

	varIamResourceLimits := _IamResourceLimits{}

	err = json.Unmarshal(data, &varIamResourceLimits)
	if err == nil {
		o.MoBaseMo = varIamResourceLimits.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AllowApiKeysWithoutExpiry")
		delete(additionalProperties, "AllowAppRegistrationsWithoutExpiry")
		delete(additionalProperties, "MaxApiKeyExpiry")
		delete(additionalProperties, "MaxAppRegistrationExpiry")
		delete(additionalProperties, "PerAccountUserLimit")
		delete(additionalProperties, "Account")

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

type NullableIamResourceLimits struct {
	value *IamResourceLimits
	isSet bool
}

func (v NullableIamResourceLimits) Get() *IamResourceLimits {
	return v.value
}

func (v *NullableIamResourceLimits) Set(val *IamResourceLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableIamResourceLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableIamResourceLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamResourceLimits(val *IamResourceLimits) *NullableIamResourceLimits {
	return &NullableIamResourceLimits{value: val, isSet: true}
}

func (v NullableIamResourceLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamResourceLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
