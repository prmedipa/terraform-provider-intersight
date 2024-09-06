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

// checks if the NiatelemetryAaaLdapProviderDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryAaaLdapProviderDetails{}

// NiatelemetryAaaLdapProviderDetails Object to capture AAA Ldap provider details.
type NiatelemetryAaaLdapProviderDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Base dn of the AAA ldap provider in APIC.
	BaseDn *string `json:"BaseDn,omitempty"`
	// Dn of the AAA ldap provider in APIC.
	Dn *string `json:"Dn,omitempty"`
	// Port of the AAA ldap provider in APIC.
	Port *string `json:"Port,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Root dn of the AAA ldap provider in APIC.
	RootDn *string `json:"RootDn,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                                     `json:"SiteName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryAaaLdapProviderDetails NiatelemetryAaaLdapProviderDetails

// NewNiatelemetryAaaLdapProviderDetails instantiates a new NiatelemetryAaaLdapProviderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryAaaLdapProviderDetails(classId string, objectType string) *NiatelemetryAaaLdapProviderDetails {
	this := NiatelemetryAaaLdapProviderDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryAaaLdapProviderDetailsWithDefaults instantiates a new NiatelemetryAaaLdapProviderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryAaaLdapProviderDetailsWithDefaults() *NiatelemetryAaaLdapProviderDetails {
	this := NiatelemetryAaaLdapProviderDetails{}
	var classId string = "niatelemetry.AaaLdapProviderDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.AaaLdapProviderDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryAaaLdapProviderDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryAaaLdapProviderDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.AaaLdapProviderDetails" of the ClassId field.
func (o *NiatelemetryAaaLdapProviderDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.AaaLdapProviderDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryAaaLdapProviderDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryAaaLdapProviderDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.AaaLdapProviderDetails" of the ObjectType field.
func (o *NiatelemetryAaaLdapProviderDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.AaaLdapProviderDetails"
}

// GetBaseDn returns the BaseDn field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetBaseDn() string {
	if o == nil || IsNil(o.BaseDn) {
		var ret string
		return ret
	}
	return *o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.BaseDn) {
		return nil, false
	}
	return o.BaseDn, true
}

// HasBaseDn returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasBaseDn() bool {
	if o != nil && !IsNil(o.BaseDn) {
		return true
	}

	return false
}

// SetBaseDn gets a reference to the given string and assigns it to the BaseDn field.
func (o *NiatelemetryAaaLdapProviderDetails) SetBaseDn(v string) {
	o.BaseDn = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryAaaLdapProviderDetails) SetDn(v string) {
	o.Dn = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *NiatelemetryAaaLdapProviderDetails) SetPort(v string) {
	o.Port = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryAaaLdapProviderDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetRecordVersion() string {
	if o == nil || IsNil(o.RecordVersion) {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetRecordVersionOk() (*string, bool) {
	if o == nil || IsNil(o.RecordVersion) {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasRecordVersion() bool {
	if o != nil && !IsNil(o.RecordVersion) {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryAaaLdapProviderDetails) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetRootDn returns the RootDn field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetRootDn() string {
	if o == nil || IsNil(o.RootDn) {
		var ret string
		return ret
	}
	return *o.RootDn
}

// GetRootDnOk returns a tuple with the RootDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetRootDnOk() (*string, bool) {
	if o == nil || IsNil(o.RootDn) {
		return nil, false
	}
	return o.RootDn, true
}

// HasRootDn returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasRootDn() bool {
	if o != nil && !IsNil(o.RootDn) {
		return true
	}

	return false
}

// SetRootDn gets a reference to the given string and assigns it to the RootDn field.
func (o *NiatelemetryAaaLdapProviderDetails) SetRootDn(v string) {
	o.RootDn = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryAaaLdapProviderDetails) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryAaaLdapProviderDetails) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryAaaLdapProviderDetails) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryAaaLdapProviderDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryAaaLdapProviderDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryAaaLdapProviderDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryAaaLdapProviderDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryAaaLdapProviderDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryAaaLdapProviderDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryAaaLdapProviderDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryAaaLdapProviderDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BaseDn) {
		toSerialize["BaseDn"] = o.BaseDn
	}
	if !IsNil(o.Dn) {
		toSerialize["Dn"] = o.Dn
	}
	if !IsNil(o.Port) {
		toSerialize["Port"] = o.Port
	}
	if !IsNil(o.RecordType) {
		toSerialize["RecordType"] = o.RecordType
	}
	if !IsNil(o.RecordVersion) {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if !IsNil(o.RootDn) {
		toSerialize["RootDn"] = o.RootDn
	}
	if !IsNil(o.SiteName) {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryAaaLdapProviderDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Base dn of the AAA ldap provider in APIC.
		BaseDn *string `json:"BaseDn,omitempty"`
		// Dn of the AAA ldap provider in APIC.
		Dn *string `json:"Dn,omitempty"`
		// Port of the AAA ldap provider in APIC.
		Port *string `json:"Port,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Root dn of the AAA ldap provider in APIC.
		RootDn *string `json:"RootDn,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                                     `json:"SiteName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct := NiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryAaaLdapProviderDetails := _NiatelemetryAaaLdapProviderDetails{}
		varNiatelemetryAaaLdapProviderDetails.ClassId = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryAaaLdapProviderDetails.ObjectType = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryAaaLdapProviderDetails.BaseDn = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.BaseDn
		varNiatelemetryAaaLdapProviderDetails.Dn = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.Dn
		varNiatelemetryAaaLdapProviderDetails.Port = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.Port
		varNiatelemetryAaaLdapProviderDetails.RecordType = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.RecordType
		varNiatelemetryAaaLdapProviderDetails.RecordVersion = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryAaaLdapProviderDetails.RootDn = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.RootDn
		varNiatelemetryAaaLdapProviderDetails.SiteName = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.SiteName
		varNiatelemetryAaaLdapProviderDetails.RegisteredDevice = varNiatelemetryAaaLdapProviderDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryAaaLdapProviderDetails(varNiatelemetryAaaLdapProviderDetails)
	} else {
		return err
	}

	varNiatelemetryAaaLdapProviderDetails := _NiatelemetryAaaLdapProviderDetails{}

	err = json.Unmarshal(data, &varNiatelemetryAaaLdapProviderDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryAaaLdapProviderDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaseDn")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "Port")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "RootDn")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableNiatelemetryAaaLdapProviderDetails struct {
	value *NiatelemetryAaaLdapProviderDetails
	isSet bool
}

func (v NullableNiatelemetryAaaLdapProviderDetails) Get() *NiatelemetryAaaLdapProviderDetails {
	return v.value
}

func (v *NullableNiatelemetryAaaLdapProviderDetails) Set(val *NiatelemetryAaaLdapProviderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryAaaLdapProviderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryAaaLdapProviderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryAaaLdapProviderDetails(val *NiatelemetryAaaLdapProviderDetails) *NullableNiatelemetryAaaLdapProviderDetails {
	return &NullableNiatelemetryAaaLdapProviderDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryAaaLdapProviderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryAaaLdapProviderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
