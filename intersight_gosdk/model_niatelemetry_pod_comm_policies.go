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

// NiatelemetryPodCommPolicies Object to capture Pod Communication Policy details.
type NiatelemetryPodCommPolicies struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Comm Http Admin State of the Comm Pol in APIC.
	CommHttpAdminSt *string `json:"CommHttpAdminSt,omitempty"`
	// Comm Https Admin State of the Comm Pol in APIC.
	CommHttpsAdminSt *string `json:"CommHttpsAdminSt,omitempty"`
	// Comm Ssh Admin State of the Comm Pol in APIC.
	CommSshAdminSt *string `json:"CommSshAdminSt,omitempty"`
	// Comm Telnet Admin State of the Comm Pol in APIC.
	CommTelnetAdminSt *string `json:"CommTelnetAdminSt,omitempty"`
	// Dn of the Comm Pol in APIC.
	PolDn *string `json:"PolDn,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryPodCommPolicies NiatelemetryPodCommPolicies

// NewNiatelemetryPodCommPolicies instantiates a new NiatelemetryPodCommPolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryPodCommPolicies(classId string, objectType string) *NiatelemetryPodCommPolicies {
	this := NiatelemetryPodCommPolicies{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryPodCommPoliciesWithDefaults instantiates a new NiatelemetryPodCommPolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryPodCommPoliciesWithDefaults() *NiatelemetryPodCommPolicies {
	this := NiatelemetryPodCommPolicies{}
	var classId string = "niatelemetry.PodCommPolicies"
	this.ClassId = classId
	var objectType string = "niatelemetry.PodCommPolicies"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryPodCommPolicies) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryPodCommPolicies) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryPodCommPolicies) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryPodCommPolicies) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCommHttpAdminSt returns the CommHttpAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetCommHttpAdminSt() string {
	if o == nil || o.CommHttpAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommHttpAdminSt
}

// GetCommHttpAdminStOk returns a tuple with the CommHttpAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetCommHttpAdminStOk() (*string, bool) {
	if o == nil || o.CommHttpAdminSt == nil {
		return nil, false
	}
	return o.CommHttpAdminSt, true
}

// HasCommHttpAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasCommHttpAdminSt() bool {
	if o != nil && o.CommHttpAdminSt != nil {
		return true
	}

	return false
}

// SetCommHttpAdminSt gets a reference to the given string and assigns it to the CommHttpAdminSt field.
func (o *NiatelemetryPodCommPolicies) SetCommHttpAdminSt(v string) {
	o.CommHttpAdminSt = &v
}

// GetCommHttpsAdminSt returns the CommHttpsAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetCommHttpsAdminSt() string {
	if o == nil || o.CommHttpsAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommHttpsAdminSt
}

// GetCommHttpsAdminStOk returns a tuple with the CommHttpsAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetCommHttpsAdminStOk() (*string, bool) {
	if o == nil || o.CommHttpsAdminSt == nil {
		return nil, false
	}
	return o.CommHttpsAdminSt, true
}

// HasCommHttpsAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasCommHttpsAdminSt() bool {
	if o != nil && o.CommHttpsAdminSt != nil {
		return true
	}

	return false
}

// SetCommHttpsAdminSt gets a reference to the given string and assigns it to the CommHttpsAdminSt field.
func (o *NiatelemetryPodCommPolicies) SetCommHttpsAdminSt(v string) {
	o.CommHttpsAdminSt = &v
}

// GetCommSshAdminSt returns the CommSshAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetCommSshAdminSt() string {
	if o == nil || o.CommSshAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommSshAdminSt
}

// GetCommSshAdminStOk returns a tuple with the CommSshAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetCommSshAdminStOk() (*string, bool) {
	if o == nil || o.CommSshAdminSt == nil {
		return nil, false
	}
	return o.CommSshAdminSt, true
}

// HasCommSshAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasCommSshAdminSt() bool {
	if o != nil && o.CommSshAdminSt != nil {
		return true
	}

	return false
}

// SetCommSshAdminSt gets a reference to the given string and assigns it to the CommSshAdminSt field.
func (o *NiatelemetryPodCommPolicies) SetCommSshAdminSt(v string) {
	o.CommSshAdminSt = &v
}

// GetCommTelnetAdminSt returns the CommTelnetAdminSt field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetCommTelnetAdminSt() string {
	if o == nil || o.CommTelnetAdminSt == nil {
		var ret string
		return ret
	}
	return *o.CommTelnetAdminSt
}

// GetCommTelnetAdminStOk returns a tuple with the CommTelnetAdminSt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetCommTelnetAdminStOk() (*string, bool) {
	if o == nil || o.CommTelnetAdminSt == nil {
		return nil, false
	}
	return o.CommTelnetAdminSt, true
}

// HasCommTelnetAdminSt returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasCommTelnetAdminSt() bool {
	if o != nil && o.CommTelnetAdminSt != nil {
		return true
	}

	return false
}

// SetCommTelnetAdminSt gets a reference to the given string and assigns it to the CommTelnetAdminSt field.
func (o *NiatelemetryPodCommPolicies) SetCommTelnetAdminSt(v string) {
	o.CommTelnetAdminSt = &v
}

// GetPolDn returns the PolDn field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetPolDn() string {
	if o == nil || o.PolDn == nil {
		var ret string
		return ret
	}
	return *o.PolDn
}

// GetPolDnOk returns a tuple with the PolDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetPolDnOk() (*string, bool) {
	if o == nil || o.PolDn == nil {
		return nil, false
	}
	return o.PolDn, true
}

// HasPolDn returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasPolDn() bool {
	if o != nil && o.PolDn != nil {
		return true
	}

	return false
}

// SetPolDn gets a reference to the given string and assigns it to the PolDn field.
func (o *NiatelemetryPodCommPolicies) SetPolDn(v string) {
	o.PolDn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryPodCommPolicies) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryPodCommPolicies) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryPodCommPolicies) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryPodCommPolicies) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryPodCommPolicies) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryPodCommPolicies) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryPodCommPolicies) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryPodCommPolicies) MarshalJSON() ([]byte, error) {
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
	if o.CommHttpAdminSt != nil {
		toSerialize["CommHttpAdminSt"] = o.CommHttpAdminSt
	}
	if o.CommHttpsAdminSt != nil {
		toSerialize["CommHttpsAdminSt"] = o.CommHttpsAdminSt
	}
	if o.CommSshAdminSt != nil {
		toSerialize["CommSshAdminSt"] = o.CommSshAdminSt
	}
	if o.CommTelnetAdminSt != nil {
		toSerialize["CommTelnetAdminSt"] = o.CommTelnetAdminSt
	}
	if o.PolDn != nil {
		toSerialize["PolDn"] = o.PolDn
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.SiteName != nil {
		toSerialize["SiteName"] = o.SiteName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryPodCommPolicies) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryPodCommPoliciesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Comm Http Admin State of the Comm Pol in APIC.
		CommHttpAdminSt *string `json:"CommHttpAdminSt,omitempty"`
		// Comm Https Admin State of the Comm Pol in APIC.
		CommHttpsAdminSt *string `json:"CommHttpsAdminSt,omitempty"`
		// Comm Ssh Admin State of the Comm Pol in APIC.
		CommSshAdminSt *string `json:"CommSshAdminSt,omitempty"`
		// Comm Telnet Admin State of the Comm Pol in APIC.
		CommTelnetAdminSt *string `json:"CommTelnetAdminSt,omitempty"`
		// Dn of the Comm Pol in APIC.
		PolDn *string `json:"PolDn,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct := NiatelemetryPodCommPoliciesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryPodCommPolicies := _NiatelemetryPodCommPolicies{}
		varNiatelemetryPodCommPolicies.ClassId = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.ClassId
		varNiatelemetryPodCommPolicies.ObjectType = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.ObjectType
		varNiatelemetryPodCommPolicies.CommHttpAdminSt = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.CommHttpAdminSt
		varNiatelemetryPodCommPolicies.CommHttpsAdminSt = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.CommHttpsAdminSt
		varNiatelemetryPodCommPolicies.CommSshAdminSt = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.CommSshAdminSt
		varNiatelemetryPodCommPolicies.CommTelnetAdminSt = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.CommTelnetAdminSt
		varNiatelemetryPodCommPolicies.PolDn = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.PolDn
		varNiatelemetryPodCommPolicies.RecordType = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.RecordType
		varNiatelemetryPodCommPolicies.RecordVersion = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryPodCommPolicies.SiteName = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.SiteName
		varNiatelemetryPodCommPolicies.RegisteredDevice = varNiatelemetryPodCommPoliciesWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryPodCommPolicies(varNiatelemetryPodCommPolicies)
	} else {
		return err
	}

	varNiatelemetryPodCommPolicies := _NiatelemetryPodCommPolicies{}

	err = json.Unmarshal(bytes, &varNiatelemetryPodCommPolicies)
	if err == nil {
		o.MoBaseMo = varNiatelemetryPodCommPolicies.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CommHttpAdminSt")
		delete(additionalProperties, "CommHttpsAdminSt")
		delete(additionalProperties, "CommSshAdminSt")
		delete(additionalProperties, "CommTelnetAdminSt")
		delete(additionalProperties, "PolDn")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
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

type NullableNiatelemetryPodCommPolicies struct {
	value *NiatelemetryPodCommPolicies
	isSet bool
}

func (v NullableNiatelemetryPodCommPolicies) Get() *NiatelemetryPodCommPolicies {
	return v.value
}

func (v *NullableNiatelemetryPodCommPolicies) Set(val *NiatelemetryPodCommPolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryPodCommPolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryPodCommPolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryPodCommPolicies(val *NiatelemetryPodCommPolicies) *NullableNiatelemetryPodCommPolicies {
	return &NullableNiatelemetryPodCommPolicies{value: val, isSet: true}
}

func (v NullableNiatelemetryPodCommPolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryPodCommPolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
