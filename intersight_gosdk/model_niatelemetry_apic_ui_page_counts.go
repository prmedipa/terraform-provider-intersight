/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// NiatelemetryApicUiPageCounts Object to capture the UI page counts in APIC.
type NiatelemetryApicUiPageCounts struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the page opened by the user.
	Dn *string `json:"Dn,omitempty"`
	// Number of times that the user has opened this page.
	PageCount *int64 `json:"PageCount,omitempty"`
	// Name of the page for which page count is recorded.
	PageName *string `json:"PageName,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicUiPageCounts NiatelemetryApicUiPageCounts

// NewNiatelemetryApicUiPageCounts instantiates a new NiatelemetryApicUiPageCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicUiPageCounts(classId string, objectType string) *NiatelemetryApicUiPageCounts {
	this := NiatelemetryApicUiPageCounts{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicUiPageCountsWithDefaults instantiates a new NiatelemetryApicUiPageCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicUiPageCountsWithDefaults() *NiatelemetryApicUiPageCounts {
	this := NiatelemetryApicUiPageCounts{}
	var classId string = "niatelemetry.ApicUiPageCounts"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicUiPageCounts"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicUiPageCounts) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicUiPageCounts) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicUiPageCounts) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicUiPageCounts) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryApicUiPageCounts) SetDn(v string) {
	o.Dn = &v
}

// GetPageCount returns the PageCount field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetPageCount() int64 {
	if o == nil || o.PageCount == nil {
		var ret int64
		return ret
	}
	return *o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetPageCountOk() (*int64, bool) {
	if o == nil || o.PageCount == nil {
		return nil, false
	}
	return o.PageCount, true
}

// HasPageCount returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasPageCount() bool {
	if o != nil && o.PageCount != nil {
		return true
	}

	return false
}

// SetPageCount gets a reference to the given int64 and assigns it to the PageCount field.
func (o *NiatelemetryApicUiPageCounts) SetPageCount(v int64) {
	o.PageCount = &v
}

// GetPageName returns the PageName field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetPageName() string {
	if o == nil || o.PageName == nil {
		var ret string
		return ret
	}
	return *o.PageName
}

// GetPageNameOk returns a tuple with the PageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetPageNameOk() (*string, bool) {
	if o == nil || o.PageName == nil {
		return nil, false
	}
	return o.PageName, true
}

// HasPageName returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasPageName() bool {
	if o != nil && o.PageName != nil {
		return true
	}

	return false
}

// SetPageName gets a reference to the given string and assigns it to the PageName field.
func (o *NiatelemetryApicUiPageCounts) SetPageName(v string) {
	o.PageName = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicUiPageCounts) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicUiPageCounts) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicUiPageCounts) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicUiPageCounts) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicUiPageCounts) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicUiPageCounts) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicUiPageCounts) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicUiPageCounts) MarshalJSON() ([]byte, error) {
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
	if o.Dn != nil {
		toSerialize["Dn"] = o.Dn
	}
	if o.PageCount != nil {
		toSerialize["PageCount"] = o.PageCount
	}
	if o.PageName != nil {
		toSerialize["PageName"] = o.PageName
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

func (o *NiatelemetryApicUiPageCounts) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryApicUiPageCountsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the page opened by the user.
		Dn *string `json:"Dn,omitempty"`
		// Number of times that the user has opened this page.
		PageCount *int64 `json:"PageCount,omitempty"`
		// Name of the page for which page count is recorded.
		PageName *string `json:"PageName,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct := NiatelemetryApicUiPageCountsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryApicUiPageCounts := _NiatelemetryApicUiPageCounts{}
		varNiatelemetryApicUiPageCounts.ClassId = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.ClassId
		varNiatelemetryApicUiPageCounts.ObjectType = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryApicUiPageCounts.Dn = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.Dn
		varNiatelemetryApicUiPageCounts.PageCount = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.PageCount
		varNiatelemetryApicUiPageCounts.PageName = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.PageName
		varNiatelemetryApicUiPageCounts.RecordType = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.RecordType
		varNiatelemetryApicUiPageCounts.RecordVersion = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryApicUiPageCounts.SiteName = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.SiteName
		varNiatelemetryApicUiPageCounts.RegisteredDevice = varNiatelemetryApicUiPageCountsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryApicUiPageCounts(varNiatelemetryApicUiPageCounts)
	} else {
		return err
	}

	varNiatelemetryApicUiPageCounts := _NiatelemetryApicUiPageCounts{}

	err = json.Unmarshal(bytes, &varNiatelemetryApicUiPageCounts)
	if err == nil {
		o.MoBaseMo = varNiatelemetryApicUiPageCounts.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "PageCount")
		delete(additionalProperties, "PageName")
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

type NullableNiatelemetryApicUiPageCounts struct {
	value *NiatelemetryApicUiPageCounts
	isSet bool
}

func (v NullableNiatelemetryApicUiPageCounts) Get() *NiatelemetryApicUiPageCounts {
	return v.value
}

func (v *NullableNiatelemetryApicUiPageCounts) Set(val *NiatelemetryApicUiPageCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicUiPageCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicUiPageCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicUiPageCounts(val *NiatelemetryApicUiPageCounts) *NullableNiatelemetryApicUiPageCounts {
	return &NullableNiatelemetryApicUiPageCounts{value: val, isSet: true}
}

func (v NullableNiatelemetryApicUiPageCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicUiPageCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
