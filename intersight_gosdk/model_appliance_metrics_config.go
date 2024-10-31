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
	"time"
)

// checks if the ApplianceMetricsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceMetricsConfig{}

// ApplianceMetricsConfig MetricsConfig provides system configuration parameters for managing metrics collection on appliance.
type ApplianceMetricsConfig struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of discovered endpoints from where metrics is being collected currently.
	CurrentEndpointCount *int64 `json:"CurrentEndpointCount,omitempty"`
	// Usage percentage of the discovered endpoints.
	EndpointUsagePercent *int64 `json:"EndpointUsagePercent,omitempty"`
	// Disabled date of the metrics collection feature.
	LastDisabledDate *time.Time `json:"LastDisabledDate,omitempty"`
	// Enabled date of the metrics collection feature.
	LastEnabledDate *time.Time `json:"LastEnabledDate,omitempty"`
	// The maximum number of supported endpoints for an appliance deployment type.
	MaxEndpointCount *int64 `json:"MaxEndpointCount,omitempty"`
	// The overall metrics collection Status based on resource constraints.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// Metric collection state defined by the system.
	SystemEnabled *bool `json:"SystemEnabled,omitempty"`
	// Configured metric collection state by the account administrator.
	UserEnabled          *bool                          `json:"UserEnabled,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceMetricsConfig ApplianceMetricsConfig

// NewApplianceMetricsConfig instantiates a new ApplianceMetricsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceMetricsConfig(classId string, objectType string) *ApplianceMetricsConfig {
	this := ApplianceMetricsConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var userEnabled bool = false
	this.UserEnabled = &userEnabled
	return &this
}

// NewApplianceMetricsConfigWithDefaults instantiates a new ApplianceMetricsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceMetricsConfigWithDefaults() *ApplianceMetricsConfig {
	this := ApplianceMetricsConfig{}
	var classId string = "appliance.MetricsConfig"
	this.ClassId = classId
	var objectType string = "appliance.MetricsConfig"
	this.ObjectType = objectType
	var userEnabled bool = false
	this.UserEnabled = &userEnabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceMetricsConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceMetricsConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.MetricsConfig" of the ClassId field.
func (o *ApplianceMetricsConfig) GetDefaultClassId() interface{} {
	return "appliance.MetricsConfig"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceMetricsConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceMetricsConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.MetricsConfig" of the ObjectType field.
func (o *ApplianceMetricsConfig) GetDefaultObjectType() interface{} {
	return "appliance.MetricsConfig"
}

// GetCurrentEndpointCount returns the CurrentEndpointCount field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetCurrentEndpointCount() int64 {
	if o == nil || IsNil(o.CurrentEndpointCount) {
		var ret int64
		return ret
	}
	return *o.CurrentEndpointCount
}

// GetCurrentEndpointCountOk returns a tuple with the CurrentEndpointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetCurrentEndpointCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentEndpointCount) {
		return nil, false
	}
	return o.CurrentEndpointCount, true
}

// HasCurrentEndpointCount returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasCurrentEndpointCount() bool {
	if o != nil && !IsNil(o.CurrentEndpointCount) {
		return true
	}

	return false
}

// SetCurrentEndpointCount gets a reference to the given int64 and assigns it to the CurrentEndpointCount field.
func (o *ApplianceMetricsConfig) SetCurrentEndpointCount(v int64) {
	o.CurrentEndpointCount = &v
}

// GetEndpointUsagePercent returns the EndpointUsagePercent field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetEndpointUsagePercent() int64 {
	if o == nil || IsNil(o.EndpointUsagePercent) {
		var ret int64
		return ret
	}
	return *o.EndpointUsagePercent
}

// GetEndpointUsagePercentOk returns a tuple with the EndpointUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetEndpointUsagePercentOk() (*int64, bool) {
	if o == nil || IsNil(o.EndpointUsagePercent) {
		return nil, false
	}
	return o.EndpointUsagePercent, true
}

// HasEndpointUsagePercent returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasEndpointUsagePercent() bool {
	if o != nil && !IsNil(o.EndpointUsagePercent) {
		return true
	}

	return false
}

// SetEndpointUsagePercent gets a reference to the given int64 and assigns it to the EndpointUsagePercent field.
func (o *ApplianceMetricsConfig) SetEndpointUsagePercent(v int64) {
	o.EndpointUsagePercent = &v
}

// GetLastDisabledDate returns the LastDisabledDate field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetLastDisabledDate() time.Time {
	if o == nil || IsNil(o.LastDisabledDate) {
		var ret time.Time
		return ret
	}
	return *o.LastDisabledDate
}

// GetLastDisabledDateOk returns a tuple with the LastDisabledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetLastDisabledDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDisabledDate) {
		return nil, false
	}
	return o.LastDisabledDate, true
}

// HasLastDisabledDate returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasLastDisabledDate() bool {
	if o != nil && !IsNil(o.LastDisabledDate) {
		return true
	}

	return false
}

// SetLastDisabledDate gets a reference to the given time.Time and assigns it to the LastDisabledDate field.
func (o *ApplianceMetricsConfig) SetLastDisabledDate(v time.Time) {
	o.LastDisabledDate = &v
}

// GetLastEnabledDate returns the LastEnabledDate field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetLastEnabledDate() time.Time {
	if o == nil || IsNil(o.LastEnabledDate) {
		var ret time.Time
		return ret
	}
	return *o.LastEnabledDate
}

// GetLastEnabledDateOk returns a tuple with the LastEnabledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetLastEnabledDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastEnabledDate) {
		return nil, false
	}
	return o.LastEnabledDate, true
}

// HasLastEnabledDate returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasLastEnabledDate() bool {
	if o != nil && !IsNil(o.LastEnabledDate) {
		return true
	}

	return false
}

// SetLastEnabledDate gets a reference to the given time.Time and assigns it to the LastEnabledDate field.
func (o *ApplianceMetricsConfig) SetLastEnabledDate(v time.Time) {
	o.LastEnabledDate = &v
}

// GetMaxEndpointCount returns the MaxEndpointCount field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetMaxEndpointCount() int64 {
	if o == nil || IsNil(o.MaxEndpointCount) {
		var ret int64
		return ret
	}
	return *o.MaxEndpointCount
}

// GetMaxEndpointCountOk returns a tuple with the MaxEndpointCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetMaxEndpointCountOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxEndpointCount) {
		return nil, false
	}
	return o.MaxEndpointCount, true
}

// HasMaxEndpointCount returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasMaxEndpointCount() bool {
	if o != nil && !IsNil(o.MaxEndpointCount) {
		return true
	}

	return false
}

// SetMaxEndpointCount gets a reference to the given int64 and assigns it to the MaxEndpointCount field.
func (o *ApplianceMetricsConfig) SetMaxEndpointCount(v int64) {
	o.MaxEndpointCount = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ApplianceMetricsConfig) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetSystemEnabled returns the SystemEnabled field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetSystemEnabled() bool {
	if o == nil || IsNil(o.SystemEnabled) {
		var ret bool
		return ret
	}
	return *o.SystemEnabled
}

// GetSystemEnabledOk returns a tuple with the SystemEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetSystemEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemEnabled) {
		return nil, false
	}
	return o.SystemEnabled, true
}

// HasSystemEnabled returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasSystemEnabled() bool {
	if o != nil && !IsNil(o.SystemEnabled) {
		return true
	}

	return false
}

// SetSystemEnabled gets a reference to the given bool and assigns it to the SystemEnabled field.
func (o *ApplianceMetricsConfig) SetSystemEnabled(v bool) {
	o.SystemEnabled = &v
}

// GetUserEnabled returns the UserEnabled field value if set, zero value otherwise.
func (o *ApplianceMetricsConfig) GetUserEnabled() bool {
	if o == nil || IsNil(o.UserEnabled) {
		var ret bool
		return ret
	}
	return *o.UserEnabled
}

// GetUserEnabledOk returns a tuple with the UserEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceMetricsConfig) GetUserEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.UserEnabled) {
		return nil, false
	}
	return o.UserEnabled, true
}

// HasUserEnabled returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasUserEnabled() bool {
	if o != nil && !IsNil(o.UserEnabled) {
		return true
	}

	return false
}

// SetUserEnabled gets a reference to the given bool and assigns it to the UserEnabled field.
func (o *ApplianceMetricsConfig) SetUserEnabled(v bool) {
	o.UserEnabled = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceMetricsConfig) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceMetricsConfig) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceMetricsConfig) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *ApplianceMetricsConfig) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *ApplianceMetricsConfig) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ApplianceMetricsConfig) UnsetAccount() {
	o.Account.Unset()
}

func (o ApplianceMetricsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceMetricsConfig) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CurrentEndpointCount) {
		toSerialize["CurrentEndpointCount"] = o.CurrentEndpointCount
	}
	if !IsNil(o.EndpointUsagePercent) {
		toSerialize["EndpointUsagePercent"] = o.EndpointUsagePercent
	}
	if !IsNil(o.LastDisabledDate) {
		toSerialize["LastDisabledDate"] = o.LastDisabledDate
	}
	if !IsNil(o.LastEnabledDate) {
		toSerialize["LastEnabledDate"] = o.LastEnabledDate
	}
	if !IsNil(o.MaxEndpointCount) {
		toSerialize["MaxEndpointCount"] = o.MaxEndpointCount
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if !IsNil(o.SystemEnabled) {
		toSerialize["SystemEnabled"] = o.SystemEnabled
	}
	if !IsNil(o.UserEnabled) {
		toSerialize["UserEnabled"] = o.UserEnabled
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceMetricsConfig) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceMetricsConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of discovered endpoints from where metrics is being collected currently.
		CurrentEndpointCount *int64 `json:"CurrentEndpointCount,omitempty"`
		// Usage percentage of the discovered endpoints.
		EndpointUsagePercent *int64 `json:"EndpointUsagePercent,omitempty"`
		// Disabled date of the metrics collection feature.
		LastDisabledDate *time.Time `json:"LastDisabledDate,omitempty"`
		// Enabled date of the metrics collection feature.
		LastEnabledDate *time.Time `json:"LastEnabledDate,omitempty"`
		// The maximum number of supported endpoints for an appliance deployment type.
		MaxEndpointCount *int64 `json:"MaxEndpointCount,omitempty"`
		// The overall metrics collection Status based on resource constraints.
		StatusMessage *string `json:"StatusMessage,omitempty"`
		// Metric collection state defined by the system.
		SystemEnabled *bool `json:"SystemEnabled,omitempty"`
		// Configured metric collection state by the account administrator.
		UserEnabled *bool                          `json:"UserEnabled,omitempty"`
		Account     NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceMetricsConfigWithoutEmbeddedStruct := ApplianceMetricsConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceMetricsConfigWithoutEmbeddedStruct)
	if err == nil {
		varApplianceMetricsConfig := _ApplianceMetricsConfig{}
		varApplianceMetricsConfig.ClassId = varApplianceMetricsConfigWithoutEmbeddedStruct.ClassId
		varApplianceMetricsConfig.ObjectType = varApplianceMetricsConfigWithoutEmbeddedStruct.ObjectType
		varApplianceMetricsConfig.CurrentEndpointCount = varApplianceMetricsConfigWithoutEmbeddedStruct.CurrentEndpointCount
		varApplianceMetricsConfig.EndpointUsagePercent = varApplianceMetricsConfigWithoutEmbeddedStruct.EndpointUsagePercent
		varApplianceMetricsConfig.LastDisabledDate = varApplianceMetricsConfigWithoutEmbeddedStruct.LastDisabledDate
		varApplianceMetricsConfig.LastEnabledDate = varApplianceMetricsConfigWithoutEmbeddedStruct.LastEnabledDate
		varApplianceMetricsConfig.MaxEndpointCount = varApplianceMetricsConfigWithoutEmbeddedStruct.MaxEndpointCount
		varApplianceMetricsConfig.StatusMessage = varApplianceMetricsConfigWithoutEmbeddedStruct.StatusMessage
		varApplianceMetricsConfig.SystemEnabled = varApplianceMetricsConfigWithoutEmbeddedStruct.SystemEnabled
		varApplianceMetricsConfig.UserEnabled = varApplianceMetricsConfigWithoutEmbeddedStruct.UserEnabled
		varApplianceMetricsConfig.Account = varApplianceMetricsConfigWithoutEmbeddedStruct.Account
		*o = ApplianceMetricsConfig(varApplianceMetricsConfig)
	} else {
		return err
	}

	varApplianceMetricsConfig := _ApplianceMetricsConfig{}

	err = json.Unmarshal(data, &varApplianceMetricsConfig)
	if err == nil {
		o.MoBaseMo = varApplianceMetricsConfig.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CurrentEndpointCount")
		delete(additionalProperties, "EndpointUsagePercent")
		delete(additionalProperties, "LastDisabledDate")
		delete(additionalProperties, "LastEnabledDate")
		delete(additionalProperties, "MaxEndpointCount")
		delete(additionalProperties, "StatusMessage")
		delete(additionalProperties, "SystemEnabled")
		delete(additionalProperties, "UserEnabled")
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

type NullableApplianceMetricsConfig struct {
	value *ApplianceMetricsConfig
	isSet bool
}

func (v NullableApplianceMetricsConfig) Get() *ApplianceMetricsConfig {
	return v.value
}

func (v *NullableApplianceMetricsConfig) Set(val *ApplianceMetricsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceMetricsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceMetricsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceMetricsConfig(val *ApplianceMetricsConfig) *NullableApplianceMetricsConfig {
	return &NullableApplianceMetricsConfig{value: val, isSet: true}
}

func (v NullableApplianceMetricsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceMetricsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
