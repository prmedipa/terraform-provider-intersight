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

// checks if the IamBannerMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamBannerMessage{}

// IamBannerMessage Configuration for the banner message, including title, contents, and display toggle.
type IamBannerMessage struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Contents of the banner message.
	BannerContents *string `json:"BannerContents,omitempty"`
	// Whether or not to display the banner message.
	BannerDisplay *bool `json:"BannerDisplay,omitempty"`
	// Title of the banner message.
	BannerTitle          *string                        `json:"BannerTitle,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamBannerMessage IamBannerMessage

// NewIamBannerMessage instantiates a new IamBannerMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamBannerMessage(classId string, objectType string) *IamBannerMessage {
	this := IamBannerMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamBannerMessageWithDefaults instantiates a new IamBannerMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamBannerMessageWithDefaults() *IamBannerMessage {
	this := IamBannerMessage{}
	var classId string = "iam.BannerMessage"
	this.ClassId = classId
	var objectType string = "iam.BannerMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamBannerMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamBannerMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.BannerMessage" of the ClassId field.
func (o *IamBannerMessage) GetDefaultClassId() interface{} {
	return "iam.BannerMessage"
}

// GetObjectType returns the ObjectType field value
func (o *IamBannerMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamBannerMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.BannerMessage" of the ObjectType field.
func (o *IamBannerMessage) GetDefaultObjectType() interface{} {
	return "iam.BannerMessage"
}

// GetBannerContents returns the BannerContents field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerContents() string {
	if o == nil || IsNil(o.BannerContents) {
		var ret string
		return ret
	}
	return *o.BannerContents
}

// GetBannerContentsOk returns a tuple with the BannerContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerContentsOk() (*string, bool) {
	if o == nil || IsNil(o.BannerContents) {
		return nil, false
	}
	return o.BannerContents, true
}

// HasBannerContents returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerContents() bool {
	if o != nil && !IsNil(o.BannerContents) {
		return true
	}

	return false
}

// SetBannerContents gets a reference to the given string and assigns it to the BannerContents field.
func (o *IamBannerMessage) SetBannerContents(v string) {
	o.BannerContents = &v
}

// GetBannerDisplay returns the BannerDisplay field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerDisplay() bool {
	if o == nil || IsNil(o.BannerDisplay) {
		var ret bool
		return ret
	}
	return *o.BannerDisplay
}

// GetBannerDisplayOk returns a tuple with the BannerDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerDisplayOk() (*bool, bool) {
	if o == nil || IsNil(o.BannerDisplay) {
		return nil, false
	}
	return o.BannerDisplay, true
}

// HasBannerDisplay returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerDisplay() bool {
	if o != nil && !IsNil(o.BannerDisplay) {
		return true
	}

	return false
}

// SetBannerDisplay gets a reference to the given bool and assigns it to the BannerDisplay field.
func (o *IamBannerMessage) SetBannerDisplay(v bool) {
	o.BannerDisplay = &v
}

// GetBannerTitle returns the BannerTitle field value if set, zero value otherwise.
func (o *IamBannerMessage) GetBannerTitle() string {
	if o == nil || IsNil(o.BannerTitle) {
		var ret string
		return ret
	}
	return *o.BannerTitle
}

// GetBannerTitleOk returns a tuple with the BannerTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamBannerMessage) GetBannerTitleOk() (*string, bool) {
	if o == nil || IsNil(o.BannerTitle) {
		return nil, false
	}
	return o.BannerTitle, true
}

// HasBannerTitle returns a boolean if a field has been set.
func (o *IamBannerMessage) HasBannerTitle() bool {
	if o != nil && !IsNil(o.BannerTitle) {
		return true
	}

	return false
}

// SetBannerTitle gets a reference to the given string and assigns it to the BannerTitle field.
func (o *IamBannerMessage) SetBannerTitle(v string) {
	o.BannerTitle = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamBannerMessage) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamBannerMessage) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamBannerMessage) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamBannerMessage) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamBannerMessage) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamBannerMessage) UnsetAccount() {
	o.Account.Unset()
}

func (o IamBannerMessage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamBannerMessage) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BannerContents) {
		toSerialize["BannerContents"] = o.BannerContents
	}
	if !IsNil(o.BannerDisplay) {
		toSerialize["BannerDisplay"] = o.BannerDisplay
	}
	if !IsNil(o.BannerTitle) {
		toSerialize["BannerTitle"] = o.BannerTitle
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamBannerMessage) UnmarshalJSON(data []byte) (err error) {
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
	type IamBannerMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Contents of the banner message.
		BannerContents *string `json:"BannerContents,omitempty"`
		// Whether or not to display the banner message.
		BannerDisplay *bool `json:"BannerDisplay,omitempty"`
		// Title of the banner message.
		BannerTitle *string                        `json:"BannerTitle,omitempty"`
		Account     NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varIamBannerMessageWithoutEmbeddedStruct := IamBannerMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamBannerMessageWithoutEmbeddedStruct)
	if err == nil {
		varIamBannerMessage := _IamBannerMessage{}
		varIamBannerMessage.ClassId = varIamBannerMessageWithoutEmbeddedStruct.ClassId
		varIamBannerMessage.ObjectType = varIamBannerMessageWithoutEmbeddedStruct.ObjectType
		varIamBannerMessage.BannerContents = varIamBannerMessageWithoutEmbeddedStruct.BannerContents
		varIamBannerMessage.BannerDisplay = varIamBannerMessageWithoutEmbeddedStruct.BannerDisplay
		varIamBannerMessage.BannerTitle = varIamBannerMessageWithoutEmbeddedStruct.BannerTitle
		varIamBannerMessage.Account = varIamBannerMessageWithoutEmbeddedStruct.Account
		*o = IamBannerMessage(varIamBannerMessage)
	} else {
		return err
	}

	varIamBannerMessage := _IamBannerMessage{}

	err = json.Unmarshal(data, &varIamBannerMessage)
	if err == nil {
		o.MoBaseMo = varIamBannerMessage.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BannerContents")
		delete(additionalProperties, "BannerDisplay")
		delete(additionalProperties, "BannerTitle")
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

type NullableIamBannerMessage struct {
	value *IamBannerMessage
	isSet bool
}

func (v NullableIamBannerMessage) Get() *IamBannerMessage {
	return v.value
}

func (v *NullableIamBannerMessage) Set(val *IamBannerMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableIamBannerMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableIamBannerMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamBannerMessage(val *IamBannerMessage) *NullableIamBannerMessage {
	return &NullableIamBannerMessage{value: val, isSet: true}
}

func (v NullableIamBannerMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamBannerMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
