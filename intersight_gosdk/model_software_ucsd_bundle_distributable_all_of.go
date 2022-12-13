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

// SoftwareUcsdBundleDistributableAllOf Definition of the list of properties defined in 'software.UcsdBundleDistributable', excluding properties defined in parent classes.
type SoftwareUcsdBundleDistributableAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                                 `json:"ObjectType"`
	Catalog    *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	// An array of relationships to softwareUcsdDistributable resources.
	Images               []SoftwareUcsdDistributableRelationship `json:"Images,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwareUcsdBundleDistributableAllOf SoftwareUcsdBundleDistributableAllOf

// NewSoftwareUcsdBundleDistributableAllOf instantiates a new SoftwareUcsdBundleDistributableAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareUcsdBundleDistributableAllOf(classId string, objectType string) *SoftwareUcsdBundleDistributableAllOf {
	this := SoftwareUcsdBundleDistributableAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwareUcsdBundleDistributableAllOfWithDefaults instantiates a new SoftwareUcsdBundleDistributableAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareUcsdBundleDistributableAllOfWithDefaults() *SoftwareUcsdBundleDistributableAllOf {
	this := SoftwareUcsdBundleDistributableAllOf{}
	var classId string = "software.UcsdBundleDistributable"
	this.ClassId = classId
	var objectType string = "software.UcsdBundleDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwareUcsdBundleDistributableAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdBundleDistributableAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwareUcsdBundleDistributableAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwareUcsdBundleDistributableAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdBundleDistributableAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwareUcsdBundleDistributableAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareUcsdBundleDistributableAllOf) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareUcsdBundleDistributableAllOf) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareUcsdBundleDistributableAllOf) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareUcsdBundleDistributableAllOf) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SoftwareUcsdBundleDistributableAllOf) GetImages() []SoftwareUcsdDistributableRelationship {
	if o == nil {
		var ret []SoftwareUcsdDistributableRelationship
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SoftwareUcsdBundleDistributableAllOf) GetImagesOk() ([]SoftwareUcsdDistributableRelationship, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *SoftwareUcsdBundleDistributableAllOf) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []SoftwareUcsdDistributableRelationship and assigns it to the Images field.
func (o *SoftwareUcsdBundleDistributableAllOf) SetImages(v []SoftwareUcsdDistributableRelationship) {
	o.Images = v
}

func (o SoftwareUcsdBundleDistributableAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	if o.Images != nil {
		toSerialize["Images"] = o.Images
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwareUcsdBundleDistributableAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSoftwareUcsdBundleDistributableAllOf := _SoftwareUcsdBundleDistributableAllOf{}

	if err = json.Unmarshal(bytes, &varSoftwareUcsdBundleDistributableAllOf); err == nil {
		*o = SoftwareUcsdBundleDistributableAllOf(varSoftwareUcsdBundleDistributableAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Catalog")
		delete(additionalProperties, "Images")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSoftwareUcsdBundleDistributableAllOf struct {
	value *SoftwareUcsdBundleDistributableAllOf
	isSet bool
}

func (v NullableSoftwareUcsdBundleDistributableAllOf) Get() *SoftwareUcsdBundleDistributableAllOf {
	return v.value
}

func (v *NullableSoftwareUcsdBundleDistributableAllOf) Set(val *SoftwareUcsdBundleDistributableAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareUcsdBundleDistributableAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareUcsdBundleDistributableAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareUcsdBundleDistributableAllOf(val *SoftwareUcsdBundleDistributableAllOf) *NullableSoftwareUcsdBundleDistributableAllOf {
	return &NullableSoftwareUcsdBundleDistributableAllOf{value: val, isSet: true}
}

func (v NullableSoftwareUcsdBundleDistributableAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareUcsdBundleDistributableAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
