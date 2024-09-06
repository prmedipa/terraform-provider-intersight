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

// checks if the AssetDeviceClaim type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetDeviceClaim{}

// AssetDeviceClaim DeviceClaim captures the intent to claim a device to an Intersight account. A device can be unclaimed by performing a DELETE on a DeviceClaim instance. When performing a claim, a secret passphrase must be obtained from the device connector UI/API by a sufficiently privileged user. The passphrase is timebound and proves that the user currently has privileged administrative access to the device being claimed.
type AssetDeviceClaim struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/SecurityTokens').
	SecurityToken *string `json:"SecurityToken,omitempty"`
	// Obtained from the device connector management UI or API (REST endpoint '/connector/DeviceIdentifiers').
	SerialNumber *string                        `json:"SerialNumber,omitempty"`
	Account      NullableIamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to moBaseMo resources.
	CustomPermissionResources []MoBaseMoRelationship                      `json:"CustomPermissionResources,omitempty"`
	Device                    NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
	Reservation               NullableResourceReservationRelationship     `json:"Reservation,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _AssetDeviceClaim AssetDeviceClaim

// NewAssetDeviceClaim instantiates a new AssetDeviceClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceClaim(classId string, objectType string) *AssetDeviceClaim {
	this := AssetDeviceClaim{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetDeviceClaimWithDefaults instantiates a new AssetDeviceClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceClaimWithDefaults() *AssetDeviceClaim {
	this := AssetDeviceClaim{}
	var classId string = "asset.DeviceClaim"
	this.ClassId = classId
	var objectType string = "asset.DeviceClaim"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetDeviceClaim) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetDeviceClaim) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "asset.DeviceClaim" of the ClassId field.
func (o *AssetDeviceClaim) GetDefaultClassId() interface{} {
	return "asset.DeviceClaim"
}

// GetObjectType returns the ObjectType field value
func (o *AssetDeviceClaim) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetDeviceClaim) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "asset.DeviceClaim" of the ObjectType field.
func (o *AssetDeviceClaim) GetDefaultObjectType() interface{} {
	return "asset.DeviceClaim"
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetSecurityToken() string {
	if o == nil || IsNil(o.SecurityToken) {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetSecurityTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityToken) {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasSecurityToken() bool {
	if o != nil && !IsNil(o.SecurityToken) {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *AssetDeviceClaim) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetDeviceClaim) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceClaim) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetDeviceClaim) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceClaim) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceClaim) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *AssetDeviceClaim) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *AssetDeviceClaim) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *AssetDeviceClaim) UnsetAccount() {
	o.Account.Unset()
}

// GetCustomPermissionResources returns the CustomPermissionResources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceClaim) GetCustomPermissionResources() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.CustomPermissionResources
}

// GetCustomPermissionResourcesOk returns a tuple with the CustomPermissionResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceClaim) GetCustomPermissionResourcesOk() ([]MoBaseMoRelationship, bool) {
	if o == nil || IsNil(o.CustomPermissionResources) {
		return nil, false
	}
	return o.CustomPermissionResources, true
}

// HasCustomPermissionResources returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasCustomPermissionResources() bool {
	if o != nil && !IsNil(o.CustomPermissionResources) {
		return true
	}

	return false
}

// SetCustomPermissionResources gets a reference to the given []MoBaseMoRelationship and assigns it to the CustomPermissionResources field.
func (o *AssetDeviceClaim) SetCustomPermissionResources(v []MoBaseMoRelationship) {
	o.CustomPermissionResources = v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceClaim) GetDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.Device.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceClaim) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the Device field.
func (o *AssetDeviceClaim) SetDevice(v AssetDeviceRegistrationRelationship) {
	o.Device.Set(&v)
}

// SetDeviceNil sets the value for Device to be an explicit nil
func (o *AssetDeviceClaim) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *AssetDeviceClaim) UnsetDevice() {
	o.Device.Unset()
}

// GetReservation returns the Reservation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetDeviceClaim) GetReservation() ResourceReservationRelationship {
	if o == nil || IsNil(o.Reservation.Get()) {
		var ret ResourceReservationRelationship
		return ret
	}
	return *o.Reservation.Get()
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetDeviceClaim) GetReservationOk() (*ResourceReservationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reservation.Get(), o.Reservation.IsSet()
}

// HasReservation returns a boolean if a field has been set.
func (o *AssetDeviceClaim) HasReservation() bool {
	if o != nil && o.Reservation.IsSet() {
		return true
	}

	return false
}

// SetReservation gets a reference to the given NullableResourceReservationRelationship and assigns it to the Reservation field.
func (o *AssetDeviceClaim) SetReservation(v ResourceReservationRelationship) {
	o.Reservation.Set(&v)
}

// SetReservationNil sets the value for Reservation to be an explicit nil
func (o *AssetDeviceClaim) SetReservationNil() {
	o.Reservation.Set(nil)
}

// UnsetReservation ensures that no value is present for Reservation, not even an explicit nil
func (o *AssetDeviceClaim) UnsetReservation() {
	o.Reservation.Unset()
}

func (o AssetDeviceClaim) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetDeviceClaim) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SecurityToken) {
		toSerialize["SecurityToken"] = o.SecurityToken
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.CustomPermissionResources != nil {
		toSerialize["CustomPermissionResources"] = o.CustomPermissionResources
	}
	if o.Device.IsSet() {
		toSerialize["Device"] = o.Device.Get()
	}
	if o.Reservation.IsSet() {
		toSerialize["Reservation"] = o.Reservation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetDeviceClaim) UnmarshalJSON(data []byte) (err error) {
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
	type AssetDeviceClaimWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Obtained from the device connector management UI or API (REST endpoint '/connector/SecurityTokens').
		SecurityToken *string `json:"SecurityToken,omitempty"`
		// Obtained from the device connector management UI or API (REST endpoint '/connector/DeviceIdentifiers').
		SerialNumber *string                        `json:"SerialNumber,omitempty"`
		Account      NullableIamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to moBaseMo resources.
		CustomPermissionResources []MoBaseMoRelationship                      `json:"CustomPermissionResources,omitempty"`
		Device                    NullableAssetDeviceRegistrationRelationship `json:"Device,omitempty"`
		Reservation               NullableResourceReservationRelationship     `json:"Reservation,omitempty"`
	}

	varAssetDeviceClaimWithoutEmbeddedStruct := AssetDeviceClaimWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssetDeviceClaimWithoutEmbeddedStruct)
	if err == nil {
		varAssetDeviceClaim := _AssetDeviceClaim{}
		varAssetDeviceClaim.ClassId = varAssetDeviceClaimWithoutEmbeddedStruct.ClassId
		varAssetDeviceClaim.ObjectType = varAssetDeviceClaimWithoutEmbeddedStruct.ObjectType
		varAssetDeviceClaim.SecurityToken = varAssetDeviceClaimWithoutEmbeddedStruct.SecurityToken
		varAssetDeviceClaim.SerialNumber = varAssetDeviceClaimWithoutEmbeddedStruct.SerialNumber
		varAssetDeviceClaim.Account = varAssetDeviceClaimWithoutEmbeddedStruct.Account
		varAssetDeviceClaim.CustomPermissionResources = varAssetDeviceClaimWithoutEmbeddedStruct.CustomPermissionResources
		varAssetDeviceClaim.Device = varAssetDeviceClaimWithoutEmbeddedStruct.Device
		varAssetDeviceClaim.Reservation = varAssetDeviceClaimWithoutEmbeddedStruct.Reservation
		*o = AssetDeviceClaim(varAssetDeviceClaim)
	} else {
		return err
	}

	varAssetDeviceClaim := _AssetDeviceClaim{}

	err = json.Unmarshal(data, &varAssetDeviceClaim)
	if err == nil {
		o.MoBaseMo = varAssetDeviceClaim.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "SecurityToken")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "CustomPermissionResources")
		delete(additionalProperties, "Device")
		delete(additionalProperties, "Reservation")

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

type NullableAssetDeviceClaim struct {
	value *AssetDeviceClaim
	isSet bool
}

func (v NullableAssetDeviceClaim) Get() *AssetDeviceClaim {
	return v.value
}

func (v *NullableAssetDeviceClaim) Set(val *AssetDeviceClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceClaim(val *AssetDeviceClaim) *NullableAssetDeviceClaim {
	return &NullableAssetDeviceClaim{value: val, isSet: true}
}

func (v NullableAssetDeviceClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
