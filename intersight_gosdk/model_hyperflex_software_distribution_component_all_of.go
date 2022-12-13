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

// HyperflexSoftwareDistributionComponentAllOf Definition of the list of properties defined in 'hyperflex.SoftwareDistributionComponent', excluding properties defined in parent classes.
type HyperflexSoftwareDistributionComponentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The bucket name where the files are present, if source is external cloud store.
	BucketName *string `json:"BucketName,omitempty"`
	// The HyperFlex Software Distribution Component Identifier.
	ComponentId *string `json:"ComponentId,omitempty"`
	// The HyperFlex Software Distribution Component Name.
	ComponentName *string `json:"ComponentName,omitempty"`
	// File location on the cloud storage.
	FilePath        *string  `json:"FilePath,omitempty"`
	FilesToDownload []string `json:"FilesToDownload,omitempty"`
	// The HyperFlex Software Distribution Component Version.
	Version                     *string                                           `json:"Version,omitempty"`
	SoftwareDistributionVersion *HyperflexSoftwareDistributionVersionRelationship `json:"SoftwareDistributionVersion,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _HyperflexSoftwareDistributionComponentAllOf HyperflexSoftwareDistributionComponentAllOf

// NewHyperflexSoftwareDistributionComponentAllOf instantiates a new HyperflexSoftwareDistributionComponentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSoftwareDistributionComponentAllOf(classId string, objectType string) *HyperflexSoftwareDistributionComponentAllOf {
	this := HyperflexSoftwareDistributionComponentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSoftwareDistributionComponentAllOfWithDefaults instantiates a new HyperflexSoftwareDistributionComponentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSoftwareDistributionComponentAllOfWithDefaults() *HyperflexSoftwareDistributionComponentAllOf {
	this := HyperflexSoftwareDistributionComponentAllOf{}
	var classId string = "hyperflex.SoftwareDistributionComponent"
	this.ClassId = classId
	var objectType string = "hyperflex.SoftwareDistributionComponent"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSoftwareDistributionComponentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSoftwareDistributionComponentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSoftwareDistributionComponentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSoftwareDistributionComponentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasBucketName() bool {
	if o != nil && o.BucketName != nil {
		return true
	}

	return false
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetBucketName(v string) {
	o.BucketName = &v
}

// GetComponentId returns the ComponentId field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentId() string {
	if o == nil || o.ComponentId == nil {
		var ret string
		return ret
	}
	return *o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentIdOk() (*string, bool) {
	if o == nil || o.ComponentId == nil {
		return nil, false
	}
	return o.ComponentId, true
}

// HasComponentId returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasComponentId() bool {
	if o != nil && o.ComponentId != nil {
		return true
	}

	return false
}

// SetComponentId gets a reference to the given string and assigns it to the ComponentId field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetComponentId(v string) {
	o.ComponentId = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasComponentName() bool {
	if o != nil && o.ComponentName != nil {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFilesToDownload returns the FilesToDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilesToDownload() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.FilesToDownload
}

// GetFilesToDownloadOk returns a tuple with the FilesToDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSoftwareDistributionComponentAllOf) GetFilesToDownloadOk() ([]string, bool) {
	if o == nil || o.FilesToDownload == nil {
		return nil, false
	}
	return o.FilesToDownload, true
}

// HasFilesToDownload returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasFilesToDownload() bool {
	if o != nil && o.FilesToDownload != nil {
		return true
	}

	return false
}

// SetFilesToDownload gets a reference to the given []string and assigns it to the FilesToDownload field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetFilesToDownload(v []string) {
	o.FilesToDownload = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetSoftwareDistributionVersion returns the SoftwareDistributionVersion field value if set, zero value otherwise.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetSoftwareDistributionVersion() HyperflexSoftwareDistributionVersionRelationship {
	if o == nil || o.SoftwareDistributionVersion == nil {
		var ret HyperflexSoftwareDistributionVersionRelationship
		return ret
	}
	return *o.SoftwareDistributionVersion
}

// GetSoftwareDistributionVersionOk returns a tuple with the SoftwareDistributionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) GetSoftwareDistributionVersionOk() (*HyperflexSoftwareDistributionVersionRelationship, bool) {
	if o == nil || o.SoftwareDistributionVersion == nil {
		return nil, false
	}
	return o.SoftwareDistributionVersion, true
}

// HasSoftwareDistributionVersion returns a boolean if a field has been set.
func (o *HyperflexSoftwareDistributionComponentAllOf) HasSoftwareDistributionVersion() bool {
	if o != nil && o.SoftwareDistributionVersion != nil {
		return true
	}

	return false
}

// SetSoftwareDistributionVersion gets a reference to the given HyperflexSoftwareDistributionVersionRelationship and assigns it to the SoftwareDistributionVersion field.
func (o *HyperflexSoftwareDistributionComponentAllOf) SetSoftwareDistributionVersion(v HyperflexSoftwareDistributionVersionRelationship) {
	o.SoftwareDistributionVersion = &v
}

func (o HyperflexSoftwareDistributionComponentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BucketName != nil {
		toSerialize["BucketName"] = o.BucketName
	}
	if o.ComponentId != nil {
		toSerialize["ComponentId"] = o.ComponentId
	}
	if o.ComponentName != nil {
		toSerialize["ComponentName"] = o.ComponentName
	}
	if o.FilePath != nil {
		toSerialize["FilePath"] = o.FilePath
	}
	if o.FilesToDownload != nil {
		toSerialize["FilesToDownload"] = o.FilesToDownload
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.SoftwareDistributionVersion != nil {
		toSerialize["SoftwareDistributionVersion"] = o.SoftwareDistributionVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSoftwareDistributionComponentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexSoftwareDistributionComponentAllOf := _HyperflexSoftwareDistributionComponentAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexSoftwareDistributionComponentAllOf); err == nil {
		*o = HyperflexSoftwareDistributionComponentAllOf(varHyperflexSoftwareDistributionComponentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BucketName")
		delete(additionalProperties, "ComponentId")
		delete(additionalProperties, "ComponentName")
		delete(additionalProperties, "FilePath")
		delete(additionalProperties, "FilesToDownload")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "SoftwareDistributionVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexSoftwareDistributionComponentAllOf struct {
	value *HyperflexSoftwareDistributionComponentAllOf
	isSet bool
}

func (v NullableHyperflexSoftwareDistributionComponentAllOf) Get() *HyperflexSoftwareDistributionComponentAllOf {
	return v.value
}

func (v *NullableHyperflexSoftwareDistributionComponentAllOf) Set(val *HyperflexSoftwareDistributionComponentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSoftwareDistributionComponentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSoftwareDistributionComponentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSoftwareDistributionComponentAllOf(val *HyperflexSoftwareDistributionComponentAllOf) *NullableHyperflexSoftwareDistributionComponentAllOf {
	return &NullableHyperflexSoftwareDistributionComponentAllOf{value: val, isSet: true}
}

func (v NullableHyperflexSoftwareDistributionComponentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSoftwareDistributionComponentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
