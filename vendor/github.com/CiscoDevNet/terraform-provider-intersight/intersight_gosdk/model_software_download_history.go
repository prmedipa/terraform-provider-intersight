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

// checks if the SoftwareDownloadHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftwareDownloadHistory{}

// SoftwareDownloadHistory An object to keep track of software downloads from the Private Appliance portal in SaaS.
type SoftwareDownloadHistory struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of software which was downloaded.
	Name *string `json:"Name,omitempty"`
	// The product type of the downloaded software.
	Product *string `json:"Product,omitempty"`
	// The download time of the software image.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The email id of the user who initiated the software download.
	UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
	// The version of software which was downloaded.
	Version              *string                                       `json:"Version,omitempty"`
	Account              NullableIamAccountRelationship                `json:"Account,omitempty"`
	Image                NullableFirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareDownloadHistory SoftwareDownloadHistory

// NewSoftwareDownloadHistory instantiates a new SoftwareDownloadHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareDownloadHistory(classId string, objectType string) *SoftwareDownloadHistory {
	this := SoftwareDownloadHistory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareDownloadHistoryWithDefaults instantiates a new SoftwareDownloadHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareDownloadHistoryWithDefaults() *SoftwareDownloadHistory {
	this := SoftwareDownloadHistory{}
	var classId string = "software.DownloadHistory"
	this.ClassId = classId
	var objectType string = "software.DownloadHistory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareDownloadHistory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareDownloadHistory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "software.DownloadHistory" of the ClassId field.
func (o *SoftwareDownloadHistory) GetDefaultClassId() interface{} {
	return "software.DownloadHistory"
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareDownloadHistory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareDownloadHistory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "software.DownloadHistory" of the ObjectType field.
func (o *SoftwareDownloadHistory) GetDefaultObjectType() interface{} {
	return "software.DownloadHistory"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoftwareDownloadHistory) SetName(v string) {
	o.Name = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *SoftwareDownloadHistory) SetProduct(v string) {
	o.Product = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SoftwareDownloadHistory) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetUserIdOrEmail() string {
	if o == nil || IsNil(o.UserIdOrEmail) {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserIdOrEmail) {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasUserIdOrEmail() bool {
	if o != nil && !IsNil(o.UserIdOrEmail) {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *SoftwareDownloadHistory) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SoftwareDownloadHistory) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareDownloadHistory) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SoftwareDownloadHistory) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareDownloadHistory) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareDownloadHistory) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *SoftwareDownloadHistory) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *SoftwareDownloadHistory) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *SoftwareDownloadHistory) UnsetAccount() {
	o.Account.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareDownloadHistory) GetImage() FirmwareBaseDistributableRelationship {
	if o == nil || IsNil(o.Image.Get()) {
		var ret FirmwareBaseDistributableRelationship
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareDownloadHistory) GetImageOk() (*FirmwareBaseDistributableRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *SoftwareDownloadHistory) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableFirmwareBaseDistributableRelationship and assigns it to the Image field.
func (o *SoftwareDownloadHistory) SetImage(v FirmwareBaseDistributableRelationship) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *SoftwareDownloadHistory) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *SoftwareDownloadHistory) UnsetImage() {
	o.Image.Unset()
}

func (o SoftwareDownloadHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftwareDownloadHistory) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Product) {
		toSerialize["Product"] = o.Product
	}
	if !IsNil(o.Timestamp) {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if !IsNil(o.UserIdOrEmail) {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.Image.IsSet() {
		toSerialize["Image"] = o.Image.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SoftwareDownloadHistory) UnmarshalJSON(data []byte) (err error) {
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
	type SoftwareDownloadHistoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of software which was downloaded.
		Name *string `json:"Name,omitempty"`
		// The product type of the downloaded software.
		Product *string `json:"Product,omitempty"`
		// The download time of the software image.
		Timestamp *time.Time `json:"Timestamp,omitempty"`
		// The email id of the user who initiated the software download.
		UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
		// The version of software which was downloaded.
		Version *string                                       `json:"Version,omitempty"`
		Account NullableIamAccountRelationship                `json:"Account,omitempty"`
		Image   NullableFirmwareBaseDistributableRelationship `json:"Image,omitempty"`
	}

	varSoftwareDownloadHistoryWithoutEmbeddedStruct := SoftwareDownloadHistoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSoftwareDownloadHistoryWithoutEmbeddedStruct)
	if err == nil {
		varSoftwareDownloadHistory := _SoftwareDownloadHistory{}
		varSoftwareDownloadHistory.ClassId = varSoftwareDownloadHistoryWithoutEmbeddedStruct.ClassId
		varSoftwareDownloadHistory.ObjectType = varSoftwareDownloadHistoryWithoutEmbeddedStruct.ObjectType
		varSoftwareDownloadHistory.Name = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Name
		varSoftwareDownloadHistory.Product = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Product
		varSoftwareDownloadHistory.Timestamp = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Timestamp
		varSoftwareDownloadHistory.UserIdOrEmail = varSoftwareDownloadHistoryWithoutEmbeddedStruct.UserIdOrEmail
		varSoftwareDownloadHistory.Version = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Version
		varSoftwareDownloadHistory.Account = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Account
		varSoftwareDownloadHistory.Image = varSoftwareDownloadHistoryWithoutEmbeddedStruct.Image
		*o = SoftwareDownloadHistory(varSoftwareDownloadHistory)
	} else {
		return err
	}

	varSoftwareDownloadHistory := _SoftwareDownloadHistory{}

	err = json.Unmarshal(data, &varSoftwareDownloadHistory)
	if err == nil {
		o.MoBaseMo = varSoftwareDownloadHistory.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Product")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Image")

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

type NullableSoftwareDownloadHistory struct {
	value *SoftwareDownloadHistory
	isSet bool
}

func (v NullableSoftwareDownloadHistory) Get() *SoftwareDownloadHistory {
	return v.value
}

func (v *NullableSoftwareDownloadHistory) Set(val *SoftwareDownloadHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareDownloadHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareDownloadHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareDownloadHistory(val *SoftwareDownloadHistory) *NullableSoftwareDownloadHistory {
	return &NullableSoftwareDownloadHistory{value: val, isSet: true}
}

func (v NullableSoftwareDownloadHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareDownloadHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
