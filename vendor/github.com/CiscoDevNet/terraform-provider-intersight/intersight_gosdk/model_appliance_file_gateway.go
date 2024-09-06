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
	"time"
)

// checks if the ApplianceFileGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceFileGateway{}

// ApplianceFileGateway FileGateway is a non-persistent model for accessing files from the Intersight. Intersight Appliances queries the FileGateway API, supplying a filename and version, to get the signed URL from the Intersight. Intersight Appliance services uses the signed URLs to download files and store them in the local image cache. Intersight will not store any record of the file access in the initial revision. This model is a pure pass through proxy for the cloud storage service.
type ApplianceFileGateway struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Bucket name in the cloud storage service where the file is stored.
	BucketName *string `json:"BucketName,omitempty"`
	// Size of the file in bytes. FileSize maybe zero if the storage service does not report file size.
	FileSize *int64 `json:"FileSize,omitempty"`
	// File timestamp as reported by the cloud storage service.
	FileTime *time.Time `json:"FileTime,omitempty"`
	// User defined file type supplied by the caller.
	FileType *string `json:"FileType,omitempty"`
	// Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path.
	Filename *string `json:"Filename,omitempty"`
	// Pre-signed URL obtained from the cloud storage service.
	PresignedUrl *string `json:"PresignedUrl,omitempty"`
	// SSL certificate of the cloud storage service.
	ServerCert *string `json:"ServerCert,omitempty"`
	// Signed URL's validity period in minutes.
	ValidityPeriod *int64 `json:"ValidityPeriod,omitempty"`
	// File version as reported by the cloud storage service.
	Version              *string                        `json:"Version,omitempty"`
	Account              NullableIamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceFileGateway ApplianceFileGateway

// NewApplianceFileGateway instantiates a new ApplianceFileGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceFileGateway(classId string, objectType string) *ApplianceFileGateway {
	this := ApplianceFileGateway{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceFileGatewayWithDefaults instantiates a new ApplianceFileGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceFileGatewayWithDefaults() *ApplianceFileGateway {
	this := ApplianceFileGateway{}
	var classId string = "appliance.FileGateway"
	this.ClassId = classId
	var objectType string = "appliance.FileGateway"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceFileGateway) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceFileGateway) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.FileGateway" of the ClassId field.
func (o *ApplianceFileGateway) GetDefaultClassId() interface{} {
	return "appliance.FileGateway"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceFileGateway) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceFileGateway) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.FileGateway" of the ObjectType field.
func (o *ApplianceFileGateway) GetDefaultObjectType() interface{} {
	return "appliance.FileGateway"
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetBucketName() string {
	if o == nil || IsNil(o.BucketName) {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetBucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.BucketName) {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasBucketName() bool {
	if o != nil && !IsNil(o.BucketName) {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *ApplianceFileGateway) SetBucketName(v string) {
	o.BucketName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize) {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *ApplianceFileGateway) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetFileTime returns the FileTime field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetFileTime() time.Time {
	if o == nil || IsNil(o.FileTime) {
		var ret time.Time
		return ret
	}
	return *o.FileTime
}

// GetFileTimeOk returns a tuple with the FileTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetFileTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FileTime) {
		return nil, false
	}
	return o.FileTime, true
}

// HasFileTime returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasFileTime() bool {
	if o != nil && !IsNil(o.FileTime) {
		return true
	}

	return false
}

// SetFileTime gets a reference to the given time.Time and assigns it to the FileTime field.
func (o *ApplianceFileGateway) SetFileTime(v time.Time) {
	o.FileTime = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *ApplianceFileGateway) SetFileType(v string) {
	o.FileType = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceFileGateway) SetFilename(v string) {
	o.Filename = &v
}

// GetPresignedUrl returns the PresignedUrl field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetPresignedUrl() string {
	if o == nil || IsNil(o.PresignedUrl) {
		var ret string
		return ret
	}
	return *o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetPresignedUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PresignedUrl) {
		return nil, false
	}
	return o.PresignedUrl, true
}

// HasPresignedUrl returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasPresignedUrl() bool {
	if o != nil && !IsNil(o.PresignedUrl) {
		return true
	}

	return false
}

// SetPresignedUrl gets a reference to the given string and assigns it to the PresignedUrl field.
func (o *ApplianceFileGateway) SetPresignedUrl(v string) {
	o.PresignedUrl = &v
}

// GetServerCert returns the ServerCert field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetServerCert() string {
	if o == nil || IsNil(o.ServerCert) {
		var ret string
		return ret
	}
	return *o.ServerCert
}

// GetServerCertOk returns a tuple with the ServerCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetServerCertOk() (*string, bool) {
	if o == nil || IsNil(o.ServerCert) {
		return nil, false
	}
	return o.ServerCert, true
}

// HasServerCert returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasServerCert() bool {
	if o != nil && !IsNil(o.ServerCert) {
		return true
	}

	return false
}

// SetServerCert gets a reference to the given string and assigns it to the ServerCert field.
func (o *ApplianceFileGateway) SetServerCert(v string) {
	o.ServerCert = &v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetValidityPeriod() int64 {
	if o == nil || IsNil(o.ValidityPeriod) {
		var ret int64
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetValidityPeriodOk() (*int64, bool) {
	if o == nil || IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasValidityPeriod() bool {
	if o != nil && !IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given int64 and assigns it to the ValidityPeriod field.
func (o *ApplianceFileGateway) SetValidityPeriod(v int64) {
	o.ValidityPeriod = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplianceFileGateway) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceFileGateway) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplianceFileGateway) SetVersion(v string) {
	o.Version = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplianceFileGateway) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplianceFileGateway) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceFileGateway) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *ApplianceFileGateway) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *ApplianceFileGateway) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ApplianceFileGateway) UnsetAccount() {
	o.Account.Unset()
}

func (o ApplianceFileGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceFileGateway) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.BucketName) {
		toSerialize["BucketName"] = o.BucketName
	}
	if !IsNil(o.FileSize) {
		toSerialize["FileSize"] = o.FileSize
	}
	if !IsNil(o.FileTime) {
		toSerialize["FileTime"] = o.FileTime
	}
	if !IsNil(o.FileType) {
		toSerialize["FileType"] = o.FileType
	}
	if !IsNil(o.Filename) {
		toSerialize["Filename"] = o.Filename
	}
	if !IsNil(o.PresignedUrl) {
		toSerialize["PresignedUrl"] = o.PresignedUrl
	}
	if !IsNil(o.ServerCert) {
		toSerialize["ServerCert"] = o.ServerCert
	}
	if !IsNil(o.ValidityPeriod) {
		toSerialize["ValidityPeriod"] = o.ValidityPeriod
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceFileGateway) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceFileGatewayWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Bucket name in the cloud storage service where the file is stored.
		BucketName *string `json:"BucketName,omitempty"`
		// Size of the file in bytes. FileSize maybe zero if the storage service does not report file size.
		FileSize *int64 `json:"FileSize,omitempty"`
		// File timestamp as reported by the cloud storage service.
		FileTime *time.Time `json:"FileTime,omitempty"`
		// User defined file type supplied by the caller.
		FileType *string `json:"FileType,omitempty"`
		// Full name of the file as it exists in the cloud storage service. If the file is in a subfolder, then the filename must contain the full path.
		Filename *string `json:"Filename,omitempty"`
		// Pre-signed URL obtained from the cloud storage service.
		PresignedUrl *string `json:"PresignedUrl,omitempty"`
		// SSL certificate of the cloud storage service.
		ServerCert *string `json:"ServerCert,omitempty"`
		// Signed URL's validity period in minutes.
		ValidityPeriod *int64 `json:"ValidityPeriod,omitempty"`
		// File version as reported by the cloud storage service.
		Version *string                        `json:"Version,omitempty"`
		Account NullableIamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceFileGatewayWithoutEmbeddedStruct := ApplianceFileGatewayWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceFileGatewayWithoutEmbeddedStruct)
	if err == nil {
		varApplianceFileGateway := _ApplianceFileGateway{}
		varApplianceFileGateway.ClassId = varApplianceFileGatewayWithoutEmbeddedStruct.ClassId
		varApplianceFileGateway.ObjectType = varApplianceFileGatewayWithoutEmbeddedStruct.ObjectType
		varApplianceFileGateway.BucketName = varApplianceFileGatewayWithoutEmbeddedStruct.BucketName
		varApplianceFileGateway.FileSize = varApplianceFileGatewayWithoutEmbeddedStruct.FileSize
		varApplianceFileGateway.FileTime = varApplianceFileGatewayWithoutEmbeddedStruct.FileTime
		varApplianceFileGateway.FileType = varApplianceFileGatewayWithoutEmbeddedStruct.FileType
		varApplianceFileGateway.Filename = varApplianceFileGatewayWithoutEmbeddedStruct.Filename
		varApplianceFileGateway.PresignedUrl = varApplianceFileGatewayWithoutEmbeddedStruct.PresignedUrl
		varApplianceFileGateway.ServerCert = varApplianceFileGatewayWithoutEmbeddedStruct.ServerCert
		varApplianceFileGateway.ValidityPeriod = varApplianceFileGatewayWithoutEmbeddedStruct.ValidityPeriod
		varApplianceFileGateway.Version = varApplianceFileGatewayWithoutEmbeddedStruct.Version
		varApplianceFileGateway.Account = varApplianceFileGatewayWithoutEmbeddedStruct.Account
		*o = ApplianceFileGateway(varApplianceFileGateway)
	} else {
		return err
	}

	varApplianceFileGateway := _ApplianceFileGateway{}

	err = json.Unmarshal(data, &varApplianceFileGateway)
	if err == nil {
		o.MoBaseMo = varApplianceFileGateway.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "FileSize")
		delete(additionalProperties, "FileTime")
		delete(additionalProperties, "FileType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "PresignedUrl")
		delete(additionalProperties, "ServerCert")
		delete(additionalProperties, "ValidityPeriod")
		delete(additionalProperties, "Version")
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

type NullableApplianceFileGateway struct {
	value *ApplianceFileGateway
	isSet bool
}

func (v NullableApplianceFileGateway) Get() *ApplianceFileGateway {
	return v.value
}

func (v *NullableApplianceFileGateway) Set(val *ApplianceFileGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceFileGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceFileGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceFileGateway(val *ApplianceFileGateway) *NullableApplianceFileGateway {
	return &NullableApplianceFileGateway{value: val, isSet: true}
}

func (v NullableApplianceFileGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceFileGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
