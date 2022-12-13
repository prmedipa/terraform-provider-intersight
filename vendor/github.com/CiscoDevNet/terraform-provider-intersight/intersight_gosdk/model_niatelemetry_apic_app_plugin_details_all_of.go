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
)

// NiatelemetryApicAppPluginDetailsAllOf Definition of the list of properties defined in 'niatelemetry.ApicAppPluginDetails', excluding properties defined in parent classes.
type NiatelemetryApicAppPluginDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// App Id of the plugin in APIC.
	AppId *string `json:"AppId,omitempty"`
	// Permissions to the app plugin in APIC.
	Permission *string `json:"Permission,omitempty"`
	// Permission Level of the app plugin in APIC.
	PermissionLevel *string `json:"PermissionLevel,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryApicAppPluginDetailsAllOf NiatelemetryApicAppPluginDetailsAllOf

// NewNiatelemetryApicAppPluginDetailsAllOf instantiates a new NiatelemetryApicAppPluginDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryApicAppPluginDetailsAllOf(classId string, objectType string) *NiatelemetryApicAppPluginDetailsAllOf {
	this := NiatelemetryApicAppPluginDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryApicAppPluginDetailsAllOfWithDefaults instantiates a new NiatelemetryApicAppPluginDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryApicAppPluginDetailsAllOfWithDefaults() *NiatelemetryApicAppPluginDetailsAllOf {
	this := NiatelemetryApicAppPluginDetailsAllOf{}
	var classId string = "niatelemetry.ApicAppPluginDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.ApicAppPluginDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetAppId(v string) {
	o.AppId = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetPermission(v string) {
	o.Permission = &v
}

// GetPermissionLevel returns the PermissionLevel field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionLevel() string {
	if o == nil || o.PermissionLevel == nil {
		var ret string
		return ret
	}
	return *o.PermissionLevel
}

// GetPermissionLevelOk returns a tuple with the PermissionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetPermissionLevelOk() (*string, bool) {
	if o == nil || o.PermissionLevel == nil {
		return nil, false
	}
	return o.PermissionLevel, true
}

// HasPermissionLevel returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasPermissionLevel() bool {
	if o != nil && o.PermissionLevel != nil {
		return true
	}

	return false
}

// SetPermissionLevel gets a reference to the given string and assigns it to the PermissionLevel field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetPermissionLevel(v string) {
	o.PermissionLevel = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryApicAppPluginDetailsAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryApicAppPluginDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryApicAppPluginDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AppId != nil {
		toSerialize["AppId"] = o.AppId
	}
	if o.Permission != nil {
		toSerialize["Permission"] = o.Permission
	}
	if o.PermissionLevel != nil {
		toSerialize["PermissionLevel"] = o.PermissionLevel
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

func (o *NiatelemetryApicAppPluginDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryApicAppPluginDetailsAllOf := _NiatelemetryApicAppPluginDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryApicAppPluginDetailsAllOf); err == nil {
		*o = NiatelemetryApicAppPluginDetailsAllOf(varNiatelemetryApicAppPluginDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AppId")
		delete(additionalProperties, "Permission")
		delete(additionalProperties, "PermissionLevel")
		delete(additionalProperties, "RecordType")
		delete(additionalProperties, "RecordVersion")
		delete(additionalProperties, "SiteName")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryApicAppPluginDetailsAllOf struct {
	value *NiatelemetryApicAppPluginDetailsAllOf
	isSet bool
}

func (v NullableNiatelemetryApicAppPluginDetailsAllOf) Get() *NiatelemetryApicAppPluginDetailsAllOf {
	return v.value
}

func (v *NullableNiatelemetryApicAppPluginDetailsAllOf) Set(val *NiatelemetryApicAppPluginDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryApicAppPluginDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryApicAppPluginDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryApicAppPluginDetailsAllOf(val *NiatelemetryApicAppPluginDetailsAllOf) *NullableNiatelemetryApicAppPluginDetailsAllOf {
	return &NullableNiatelemetryApicAppPluginDetailsAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryApicAppPluginDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryApicAppPluginDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
