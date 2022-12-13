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

// AssetSudiInfo The SUDI is an X.509v3 certificate, which maintains the product identifier and serial number. The identity is implemented at manufacturing and chained to a publicly identifiable root certificate authority. It can be used as an unchangeable identity for configuration, security, auditing, and management. This structure contains the SUDI information read from the device's Trust Anchor Module (TAM).
type AssetSudiInfo struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
	Pid *string `json:"Pid,omitempty"`
	// The device SerialNumber extracted from the X.509 SUDI Leaf Certificate.
	SerialNumber *string `json:"SerialNumber,omitempty"`
	// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certificate.
	Signature *string `json:"Signature,omitempty"`
	// The validation status of the device. * `DeviceStatusUnknown` - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value. * `Verified` - The device returned a valid PID, Serial Number, Status and X.509 Leaf Certificate. The certificate signing chain was validated. * `CertificateValidationFailed` - Validation of the certificate signing chain failed. * `UnsupportedFirmware` - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation. * `UnsupportedHardware` - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated. * `DeviceNotResponding` - A request was sent to the device, but no response was received.
	Status               *string                 `json:"Status,omitempty"`
	SudiCertificate      NullableX509Certificate `json:"SudiCertificate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetSudiInfo AssetSudiInfo

// NewAssetSudiInfo instantiates a new AssetSudiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetSudiInfo(classId string, objectType string) *AssetSudiInfo {
	this := AssetSudiInfo{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "DeviceStatusUnknown"
	this.Status = &status
	return &this
}

// NewAssetSudiInfoWithDefaults instantiates a new AssetSudiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetSudiInfoWithDefaults() *AssetSudiInfo {
	this := AssetSudiInfo{}
	var classId string = "asset.SudiInfo"
	this.ClassId = classId
	var objectType string = "asset.SudiInfo"
	this.ObjectType = objectType
	var status string = "DeviceStatusUnknown"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetSudiInfo) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetSudiInfo) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetSudiInfo) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetSudiInfo) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *AssetSudiInfo) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *AssetSudiInfo) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *AssetSudiInfo) SetPid(v string) {
	o.Pid = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *AssetSudiInfo) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *AssetSudiInfo) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *AssetSudiInfo) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *AssetSudiInfo) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *AssetSudiInfo) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *AssetSudiInfo) SetSignature(v string) {
	o.Signature = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AssetSudiInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetSudiInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AssetSudiInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AssetSudiInfo) SetStatus(v string) {
	o.Status = &v
}

// GetSudiCertificate returns the SudiCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetSudiInfo) GetSudiCertificate() X509Certificate {
	if o == nil || o.SudiCertificate.Get() == nil {
		var ret X509Certificate
		return ret
	}
	return *o.SudiCertificate.Get()
}

// GetSudiCertificateOk returns a tuple with the SudiCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetSudiInfo) GetSudiCertificateOk() (*X509Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.SudiCertificate.Get(), o.SudiCertificate.IsSet()
}

// HasSudiCertificate returns a boolean if a field has been set.
func (o *AssetSudiInfo) HasSudiCertificate() bool {
	if o != nil && o.SudiCertificate.IsSet() {
		return true
	}

	return false
}

// SetSudiCertificate gets a reference to the given NullableX509Certificate and assigns it to the SudiCertificate field.
func (o *AssetSudiInfo) SetSudiCertificate(v X509Certificate) {
	o.SudiCertificate.Set(&v)
}

// SetSudiCertificateNil sets the value for SudiCertificate to be an explicit nil
func (o *AssetSudiInfo) SetSudiCertificateNil() {
	o.SudiCertificate.Set(nil)
}

// UnsetSudiCertificate ensures that no value is present for SudiCertificate, not even an explicit nil
func (o *AssetSudiInfo) UnsetSudiCertificate() {
	o.SudiCertificate.Unset()
}

func (o AssetSudiInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.SerialNumber != nil {
		toSerialize["SerialNumber"] = o.SerialNumber
	}
	if o.Signature != nil {
		toSerialize["Signature"] = o.Signature
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.SudiCertificate.IsSet() {
		toSerialize["SudiCertificate"] = o.SudiCertificate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetSudiInfo) UnmarshalJSON(bytes []byte) (err error) {
	type AssetSudiInfoWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The device model (PID) extracted from the X.509 SUDI Leaf Certificate.
		Pid *string `json:"Pid,omitempty"`
		// The device SerialNumber extracted from the X.509 SUDI Leaf Certificate.
		SerialNumber *string `json:"SerialNumber,omitempty"`
		// The signature is obtained by taking the base64 encoding of the Serial Number + PID + Status, taking the SHA256 hash and then signing with the SUDI X.509 Leaf Certificate.
		Signature *string `json:"Signature,omitempty"`
		// The validation status of the device. * `DeviceStatusUnknown` - SUDI validation is done on the establishment of a connection. Before a device connects or after it disconnects, the SUDI validation status is set to this value. * `Verified` - The device returned a valid PID, Serial Number, Status and X.509 Leaf Certificate. The certificate signing chain was validated. * `CertificateValidationFailed` - Validation of the certificate signing chain failed. * `UnsupportedFirmware` - The firmware version of the Cisco IMC that is installed does not contain the SUDI APIs needed to perform validation. * `UnsupportedHardware` - The device is a model that does not contain a Trust Anchor Module (TAM) and thus cannot be validated. * `DeviceNotResponding` - A request was sent to the device, but no response was received.
		Status          *string                 `json:"Status,omitempty"`
		SudiCertificate NullableX509Certificate `json:"SudiCertificate,omitempty"`
	}

	varAssetSudiInfoWithoutEmbeddedStruct := AssetSudiInfoWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetSudiInfoWithoutEmbeddedStruct)
	if err == nil {
		varAssetSudiInfo := _AssetSudiInfo{}
		varAssetSudiInfo.ClassId = varAssetSudiInfoWithoutEmbeddedStruct.ClassId
		varAssetSudiInfo.ObjectType = varAssetSudiInfoWithoutEmbeddedStruct.ObjectType
		varAssetSudiInfo.Pid = varAssetSudiInfoWithoutEmbeddedStruct.Pid
		varAssetSudiInfo.SerialNumber = varAssetSudiInfoWithoutEmbeddedStruct.SerialNumber
		varAssetSudiInfo.Signature = varAssetSudiInfoWithoutEmbeddedStruct.Signature
		varAssetSudiInfo.Status = varAssetSudiInfoWithoutEmbeddedStruct.Status
		varAssetSudiInfo.SudiCertificate = varAssetSudiInfoWithoutEmbeddedStruct.SudiCertificate
		*o = AssetSudiInfo(varAssetSudiInfo)
	} else {
		return err
	}

	varAssetSudiInfo := _AssetSudiInfo{}

	err = json.Unmarshal(bytes, &varAssetSudiInfo)
	if err == nil {
		o.MoBaseComplexType = varAssetSudiInfo.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "SerialNumber")
		delete(additionalProperties, "Signature")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "SudiCertificate")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableAssetSudiInfo struct {
	value *AssetSudiInfo
	isSet bool
}

func (v NullableAssetSudiInfo) Get() *AssetSudiInfo {
	return v.value
}

func (v *NullableAssetSudiInfo) Set(val *AssetSudiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetSudiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetSudiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetSudiInfo(val *AssetSudiInfo) *NullableAssetSudiInfo {
	return &NullableAssetSudiInfo{value: val, isSet: true}
}

func (v NullableAssetSudiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetSudiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
