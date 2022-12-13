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

// NiatelemetryFabricPodSs Object to capture Fabric PodS details.
type NiatelemetryFabricPodSs struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Dn of the Fabric podS for APIC.
	Dn *string `json:"Dn,omitempty"`
	// Parent PodP of the Fabric podS for APIC.
	FabricPodProf *string `json:"FabricPodProf,omitempty"`
	// Pod Block for the above Fabric PodS.
	PodBlk *string `json:"PodBlk,omitempty"`
	// Policy Group for the above Fabric PodS.
	PodPolGrp *string `json:"PodPolGrp,omitempty"`
	// List of Dn of CommPols, SnmpPols and TimePols.
	PolList *string `json:"PolList,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty"`
	// Name of the APIC site from which this data is being collected.
	SiteName             *string                              `json:"SiteName,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryFabricPodSs NiatelemetryFabricPodSs

// NewNiatelemetryFabricPodSs instantiates a new NiatelemetryFabricPodSs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryFabricPodSs(classId string, objectType string) *NiatelemetryFabricPodSs {
	this := NiatelemetryFabricPodSs{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryFabricPodSsWithDefaults instantiates a new NiatelemetryFabricPodSs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryFabricPodSsWithDefaults() *NiatelemetryFabricPodSs {
	this := NiatelemetryFabricPodSs{}
	var classId string = "niatelemetry.FabricPodSs"
	this.ClassId = classId
	var objectType string = "niatelemetry.FabricPodSs"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryFabricPodSs) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryFabricPodSs) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryFabricPodSs) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryFabricPodSs) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetDn() string {
	if o == nil || o.Dn == nil {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetDnOk() (*string, bool) {
	if o == nil || o.Dn == nil {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasDn() bool {
	if o != nil && o.Dn != nil {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *NiatelemetryFabricPodSs) SetDn(v string) {
	o.Dn = &v
}

// GetFabricPodProf returns the FabricPodProf field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetFabricPodProf() string {
	if o == nil || o.FabricPodProf == nil {
		var ret string
		return ret
	}
	return *o.FabricPodProf
}

// GetFabricPodProfOk returns a tuple with the FabricPodProf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetFabricPodProfOk() (*string, bool) {
	if o == nil || o.FabricPodProf == nil {
		return nil, false
	}
	return o.FabricPodProf, true
}

// HasFabricPodProf returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasFabricPodProf() bool {
	if o != nil && o.FabricPodProf != nil {
		return true
	}

	return false
}

// SetFabricPodProf gets a reference to the given string and assigns it to the FabricPodProf field.
func (o *NiatelemetryFabricPodSs) SetFabricPodProf(v string) {
	o.FabricPodProf = &v
}

// GetPodBlk returns the PodBlk field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetPodBlk() string {
	if o == nil || o.PodBlk == nil {
		var ret string
		return ret
	}
	return *o.PodBlk
}

// GetPodBlkOk returns a tuple with the PodBlk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetPodBlkOk() (*string, bool) {
	if o == nil || o.PodBlk == nil {
		return nil, false
	}
	return o.PodBlk, true
}

// HasPodBlk returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasPodBlk() bool {
	if o != nil && o.PodBlk != nil {
		return true
	}

	return false
}

// SetPodBlk gets a reference to the given string and assigns it to the PodBlk field.
func (o *NiatelemetryFabricPodSs) SetPodBlk(v string) {
	o.PodBlk = &v
}

// GetPodPolGrp returns the PodPolGrp field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetPodPolGrp() string {
	if o == nil || o.PodPolGrp == nil {
		var ret string
		return ret
	}
	return *o.PodPolGrp
}

// GetPodPolGrpOk returns a tuple with the PodPolGrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetPodPolGrpOk() (*string, bool) {
	if o == nil || o.PodPolGrp == nil {
		return nil, false
	}
	return o.PodPolGrp, true
}

// HasPodPolGrp returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasPodPolGrp() bool {
	if o != nil && o.PodPolGrp != nil {
		return true
	}

	return false
}

// SetPodPolGrp gets a reference to the given string and assigns it to the PodPolGrp field.
func (o *NiatelemetryFabricPodSs) SetPodPolGrp(v string) {
	o.PodPolGrp = &v
}

// GetPolList returns the PolList field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetPolList() string {
	if o == nil || o.PolList == nil {
		var ret string
		return ret
	}
	return *o.PolList
}

// GetPolListOk returns a tuple with the PolList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetPolListOk() (*string, bool) {
	if o == nil || o.PolList == nil {
		return nil, false
	}
	return o.PolList, true
}

// HasPolList returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasPolList() bool {
	if o != nil && o.PolList != nil {
		return true
	}

	return false
}

// SetPolList gets a reference to the given string and assigns it to the PolList field.
func (o *NiatelemetryFabricPodSs) SetPolList(v string) {
	o.PolList = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryFabricPodSs) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryFabricPodSs) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *NiatelemetryFabricPodSs) SetSiteName(v string) {
	o.SiteName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryFabricPodSs) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryFabricPodSs) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryFabricPodSs) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryFabricPodSs) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryFabricPodSs) MarshalJSON() ([]byte, error) {
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
	if o.FabricPodProf != nil {
		toSerialize["FabricPodProf"] = o.FabricPodProf
	}
	if o.PodBlk != nil {
		toSerialize["PodBlk"] = o.PodBlk
	}
	if o.PodPolGrp != nil {
		toSerialize["PodPolGrp"] = o.PodPolGrp
	}
	if o.PolList != nil {
		toSerialize["PolList"] = o.PolList
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

func (o *NiatelemetryFabricPodSs) UnmarshalJSON(bytes []byte) (err error) {
	type NiatelemetryFabricPodSsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Dn of the Fabric podS for APIC.
		Dn *string `json:"Dn,omitempty"`
		// Parent PodP of the Fabric podS for APIC.
		FabricPodProf *string `json:"FabricPodProf,omitempty"`
		// Pod Block for the above Fabric PodS.
		PodBlk *string `json:"PodBlk,omitempty"`
		// Policy Group for the above Fabric PodS.
		PodPolGrp *string `json:"PodPolGrp,omitempty"`
		// List of Dn of CommPols, SnmpPols and TimePols.
		PolList *string `json:"PolList,omitempty"`
		// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
		RecordType *string `json:"RecordType,omitempty"`
		// Version of record being pushed. This determines what was the API version for data available from the device.
		RecordVersion *string `json:"RecordVersion,omitempty"`
		// Name of the APIC site from which this data is being collected.
		SiteName         *string                              `json:"SiteName,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryFabricPodSsWithoutEmbeddedStruct := NiatelemetryFabricPodSsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiatelemetryFabricPodSsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryFabricPodSs := _NiatelemetryFabricPodSs{}
		varNiatelemetryFabricPodSs.ClassId = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.ClassId
		varNiatelemetryFabricPodSs.ObjectType = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryFabricPodSs.Dn = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.Dn
		varNiatelemetryFabricPodSs.FabricPodProf = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.FabricPodProf
		varNiatelemetryFabricPodSs.PodBlk = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.PodBlk
		varNiatelemetryFabricPodSs.PodPolGrp = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.PodPolGrp
		varNiatelemetryFabricPodSs.PolList = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.PolList
		varNiatelemetryFabricPodSs.RecordType = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.RecordType
		varNiatelemetryFabricPodSs.RecordVersion = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.RecordVersion
		varNiatelemetryFabricPodSs.SiteName = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.SiteName
		varNiatelemetryFabricPodSs.RegisteredDevice = varNiatelemetryFabricPodSsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryFabricPodSs(varNiatelemetryFabricPodSs)
	} else {
		return err
	}

	varNiatelemetryFabricPodSs := _NiatelemetryFabricPodSs{}

	err = json.Unmarshal(bytes, &varNiatelemetryFabricPodSs)
	if err == nil {
		o.MoBaseMo = varNiatelemetryFabricPodSs.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dn")
		delete(additionalProperties, "FabricPodProf")
		delete(additionalProperties, "PodBlk")
		delete(additionalProperties, "PodPolGrp")
		delete(additionalProperties, "PolList")
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

type NullableNiatelemetryFabricPodSs struct {
	value *NiatelemetryFabricPodSs
	isSet bool
}

func (v NullableNiatelemetryFabricPodSs) Get() *NiatelemetryFabricPodSs {
	return v.value
}

func (v *NullableNiatelemetryFabricPodSs) Set(val *NiatelemetryFabricPodSs) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryFabricPodSs) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryFabricPodSs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryFabricPodSs(val *NiatelemetryFabricPodSs) *NullableNiatelemetryFabricPodSs {
	return &NullableNiatelemetryFabricPodSs{value: val, isSet: true}
}

func (v NullableNiatelemetryFabricPodSs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryFabricPodSs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
