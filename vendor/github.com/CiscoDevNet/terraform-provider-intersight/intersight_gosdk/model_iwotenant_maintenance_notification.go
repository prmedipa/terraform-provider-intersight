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

// checks if the IwotenantMaintenanceNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IwotenantMaintenanceNotification{}

// IwotenantMaintenanceNotification Maintenance related notification to be displayed as UI banner when customer logs in the Intersight UI.
type IwotenantMaintenanceNotification struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Any i18n (internationalization) key defined the message content. If the key already exists then the  message content will be picked based on the key, otherwise provided message value will be used.
	I18nKey *string `json:"I18nKey,omitempty"`
	// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
	IwoId *string `json:"IwoId,omitempty"`
	// The date/time from which the actual maintenance operations will be performed for a Customer's account.
	MaintenanceStartTime *time.Time `json:"MaintenanceStartTime,omitempty"`
	// The notification message content is to display in the UI banner after the Customer's login to inform about planned maintenance operations on IWO.
	NtfnMessage *string `json:"NtfnMessage,omitempty"`
	// The date/time from which the maintenance banner message will be shown to the Customer after login in to  Intersight UI.
	ShowFromTime *time.Time `json:"ShowFromTime,omitempty"`
	// The date/time until which the maintenance banner message will be shown to the Customer after login into  Intersight UI. This will also be the time actual maintenance operation is planned for the finish of a  Customer's account.
	ShowUntilTime        *time.Time                     `json:"ShowUntilTime,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IwotenantMaintenanceNotification IwotenantMaintenanceNotification

// NewIwotenantMaintenanceNotification instantiates a new IwotenantMaintenanceNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIwotenantMaintenanceNotification(classId string, objectType string) *IwotenantMaintenanceNotification {
	this := IwotenantMaintenanceNotification{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIwotenantMaintenanceNotificationWithDefaults instantiates a new IwotenantMaintenanceNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIwotenantMaintenanceNotificationWithDefaults() *IwotenantMaintenanceNotification {
	this := IwotenantMaintenanceNotification{}
	var classId string = "iwotenant.MaintenanceNotification"
	this.ClassId = classId
	var objectType string = "iwotenant.MaintenanceNotification"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IwotenantMaintenanceNotification) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IwotenantMaintenanceNotification) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iwotenant.MaintenanceNotification" of the ClassId field.
func (o *IwotenantMaintenanceNotification) GetDefaultClassId() interface{} {
	return "iwotenant.MaintenanceNotification"
}

// GetObjectType returns the ObjectType field value
func (o *IwotenantMaintenanceNotification) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IwotenantMaintenanceNotification) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iwotenant.MaintenanceNotification" of the ObjectType field.
func (o *IwotenantMaintenanceNotification) GetDefaultObjectType() interface{} {
	return "iwotenant.MaintenanceNotification"
}

// GetI18nKey returns the I18nKey field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetI18nKey() string {
	if o == nil || IsNil(o.I18nKey) {
		var ret string
		return ret
	}
	return *o.I18nKey
}

// GetI18nKeyOk returns a tuple with the I18nKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetI18nKeyOk() (*string, bool) {
	if o == nil || IsNil(o.I18nKey) {
		return nil, false
	}
	return o.I18nKey, true
}

// HasI18nKey returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasI18nKey() bool {
	if o != nil && !IsNil(o.I18nKey) {
		return true
	}

	return false
}

// SetI18nKey gets a reference to the given string and assigns it to the I18nKey field.
func (o *IwotenantMaintenanceNotification) SetI18nKey(v string) {
	o.I18nKey = &v
}

// GetIwoId returns the IwoId field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetIwoId() string {
	if o == nil || IsNil(o.IwoId) {
		var ret string
		return ret
	}
	return *o.IwoId
}

// GetIwoIdOk returns a tuple with the IwoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetIwoIdOk() (*string, bool) {
	if o == nil || IsNil(o.IwoId) {
		return nil, false
	}
	return o.IwoId, true
}

// HasIwoId returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasIwoId() bool {
	if o != nil && !IsNil(o.IwoId) {
		return true
	}

	return false
}

// SetIwoId gets a reference to the given string and assigns it to the IwoId field.
func (o *IwotenantMaintenanceNotification) SetIwoId(v string) {
	o.IwoId = &v
}

// GetMaintenanceStartTime returns the MaintenanceStartTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetMaintenanceStartTime() time.Time {
	if o == nil || IsNil(o.MaintenanceStartTime) {
		var ret time.Time
		return ret
	}
	return *o.MaintenanceStartTime
}

// GetMaintenanceStartTimeOk returns a tuple with the MaintenanceStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetMaintenanceStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MaintenanceStartTime) {
		return nil, false
	}
	return o.MaintenanceStartTime, true
}

// HasMaintenanceStartTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasMaintenanceStartTime() bool {
	if o != nil && !IsNil(o.MaintenanceStartTime) {
		return true
	}

	return false
}

// SetMaintenanceStartTime gets a reference to the given time.Time and assigns it to the MaintenanceStartTime field.
func (o *IwotenantMaintenanceNotification) SetMaintenanceStartTime(v time.Time) {
	o.MaintenanceStartTime = &v
}

// GetNtfnMessage returns the NtfnMessage field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetNtfnMessage() string {
	if o == nil || IsNil(o.NtfnMessage) {
		var ret string
		return ret
	}
	return *o.NtfnMessage
}

// GetNtfnMessageOk returns a tuple with the NtfnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetNtfnMessageOk() (*string, bool) {
	if o == nil || IsNil(o.NtfnMessage) {
		return nil, false
	}
	return o.NtfnMessage, true
}

// HasNtfnMessage returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasNtfnMessage() bool {
	if o != nil && !IsNil(o.NtfnMessage) {
		return true
	}

	return false
}

// SetNtfnMessage gets a reference to the given string and assigns it to the NtfnMessage field.
func (o *IwotenantMaintenanceNotification) SetNtfnMessage(v string) {
	o.NtfnMessage = &v
}

// GetShowFromTime returns the ShowFromTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetShowFromTime() time.Time {
	if o == nil || IsNil(o.ShowFromTime) {
		var ret time.Time
		return ret
	}
	return *o.ShowFromTime
}

// GetShowFromTimeOk returns a tuple with the ShowFromTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetShowFromTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShowFromTime) {
		return nil, false
	}
	return o.ShowFromTime, true
}

// HasShowFromTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasShowFromTime() bool {
	if o != nil && !IsNil(o.ShowFromTime) {
		return true
	}

	return false
}

// SetShowFromTime gets a reference to the given time.Time and assigns it to the ShowFromTime field.
func (o *IwotenantMaintenanceNotification) SetShowFromTime(v time.Time) {
	o.ShowFromTime = &v
}

// GetShowUntilTime returns the ShowUntilTime field value if set, zero value otherwise.
func (o *IwotenantMaintenanceNotification) GetShowUntilTime() time.Time {
	if o == nil || IsNil(o.ShowUntilTime) {
		var ret time.Time
		return ret
	}
	return *o.ShowUntilTime
}

// GetShowUntilTimeOk returns a tuple with the ShowUntilTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IwotenantMaintenanceNotification) GetShowUntilTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShowUntilTime) {
		return nil, false
	}
	return o.ShowUntilTime, true
}

// HasShowUntilTime returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasShowUntilTime() bool {
	if o != nil && !IsNil(o.ShowUntilTime) {
		return true
	}

	return false
}

// SetShowUntilTime gets a reference to the given time.Time and assigns it to the ShowUntilTime field.
func (o *IwotenantMaintenanceNotification) SetShowUntilTime(v time.Time) {
	o.ShowUntilTime = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IwotenantMaintenanceNotification) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IwotenantMaintenanceNotification) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IwotenantMaintenanceNotification) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IwotenantMaintenanceNotification) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IwotenantMaintenanceNotification) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IwotenantMaintenanceNotification) UnsetAccount() {
	o.Account.Unset()
}

func (o IwotenantMaintenanceNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IwotenantMaintenanceNotification) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.I18nKey) {
		toSerialize["I18nKey"] = o.I18nKey
	}
	if !IsNil(o.IwoId) {
		toSerialize["IwoId"] = o.IwoId
	}
	if !IsNil(o.MaintenanceStartTime) {
		toSerialize["MaintenanceStartTime"] = o.MaintenanceStartTime
	}
	if !IsNil(o.NtfnMessage) {
		toSerialize["NtfnMessage"] = o.NtfnMessage
	}
	if !IsNil(o.ShowFromTime) {
		toSerialize["ShowFromTime"] = o.ShowFromTime
	}
	if !IsNil(o.ShowUntilTime) {
		toSerialize["ShowUntilTime"] = o.ShowUntilTime
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IwotenantMaintenanceNotification) UnmarshalJSON(data []byte) (err error) {
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
	type IwotenantMaintenanceNotificationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Any i18n (internationalization) key defined the message content. If the key already exists then the  message content will be picked based on the key, otherwise provided message value will be used.
		I18nKey *string `json:"I18nKey,omitempty"`
		// The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId.
		IwoId *string `json:"IwoId,omitempty"`
		// The date/time from which the actual maintenance operations will be performed for a Customer's account.
		MaintenanceStartTime *time.Time `json:"MaintenanceStartTime,omitempty"`
		// The notification message content is to display in the UI banner after the Customer's login to inform about planned maintenance operations on IWO.
		NtfnMessage *string `json:"NtfnMessage,omitempty"`
		// The date/time from which the maintenance banner message will be shown to the Customer after login in to  Intersight UI.
		ShowFromTime *time.Time `json:"ShowFromTime,omitempty"`
		// The date/time until which the maintenance banner message will be shown to the Customer after login into  Intersight UI. This will also be the time actual maintenance operation is planned for the finish of a  Customer's account.
		ShowUntilTime *time.Time                     `json:"ShowUntilTime,omitempty"`
		Account       NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIwotenantMaintenanceNotificationWithoutEmbeddedStruct := IwotenantMaintenanceNotificationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIwotenantMaintenanceNotificationWithoutEmbeddedStruct)
	if err == nil {
		varIwotenantMaintenanceNotification := _IwotenantMaintenanceNotification{}
		varIwotenantMaintenanceNotification.ClassId = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.ClassId
		varIwotenantMaintenanceNotification.ObjectType = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.ObjectType
		varIwotenantMaintenanceNotification.I18nKey = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.I18nKey
		varIwotenantMaintenanceNotification.IwoId = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.IwoId
		varIwotenantMaintenanceNotification.MaintenanceStartTime = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.MaintenanceStartTime
		varIwotenantMaintenanceNotification.NtfnMessage = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.NtfnMessage
		varIwotenantMaintenanceNotification.ShowFromTime = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.ShowFromTime
		varIwotenantMaintenanceNotification.ShowUntilTime = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.ShowUntilTime
		varIwotenantMaintenanceNotification.Account = varIwotenantMaintenanceNotificationWithoutEmbeddedStruct.Account
		*o = IwotenantMaintenanceNotification(varIwotenantMaintenanceNotification)
	} else {
		return err
	}

	varIwotenantMaintenanceNotification := _IwotenantMaintenanceNotification{}

	err = json.Unmarshal(data, &varIwotenantMaintenanceNotification)
	if err == nil {
		o.MoBaseMo = varIwotenantMaintenanceNotification.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "I18nKey")
		delete(additionalProperties, "IwoId")
		delete(additionalProperties, "MaintenanceStartTime")
		delete(additionalProperties, "NtfnMessage")
		delete(additionalProperties, "ShowFromTime")
		delete(additionalProperties, "ShowUntilTime")
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

type NullableIwotenantMaintenanceNotification struct {
	value *IwotenantMaintenanceNotification
	isSet bool
}

func (v NullableIwotenantMaintenanceNotification) Get() *IwotenantMaintenanceNotification {
	return v.value
}

func (v *NullableIwotenantMaintenanceNotification) Set(val *IwotenantMaintenanceNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableIwotenantMaintenanceNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableIwotenantMaintenanceNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIwotenantMaintenanceNotification(val *IwotenantMaintenanceNotification) *NullableIwotenantMaintenanceNotification {
	return &NullableIwotenantMaintenanceNotification{value: val, isSet: true}
}

func (v NullableIwotenantMaintenanceNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIwotenantMaintenanceNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
