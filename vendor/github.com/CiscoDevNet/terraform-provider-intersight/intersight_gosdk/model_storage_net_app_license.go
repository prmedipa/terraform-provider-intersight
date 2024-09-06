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

// checks if the StorageNetAppLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageNetAppLicense{}

// StorageNetAppLicense NetApp licenses for NetApp Ontap.
type StorageNetAppLicense struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Unique identity of the device.
	ClusterUuid *string `json:"ClusterUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
	// The name of the license package.
	Name *string `json:"Name,omitempty"`
	// Summary state of license package based on all installed licenses. * `Unknown` - The summary state of the license package is unknown. * `Compliant` - The summary state of the license package is compliant. * `Noncompliant` - The summary state of the license package is noncompliant. * `Unlicensed` - The summary state of the license package is unlicensed.
	State                *string                                  `json:"State,omitempty"`
	Array                NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppLicense StorageNetAppLicense

// NewStorageNetAppLicense instantiates a new StorageNetAppLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppLicense(classId string, objectType string) *StorageNetAppLicense {
	this := StorageNetAppLicense{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppLicenseWithDefaults instantiates a new StorageNetAppLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppLicenseWithDefaults() *StorageNetAppLicense {
	this := StorageNetAppLicense{}
	var classId string = "storage.NetAppLicense"
	this.ClassId = classId
	var objectType string = "storage.NetAppLicense"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppLicense) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLicense) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppLicense) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "storage.NetAppLicense" of the ClassId field.
func (o *StorageNetAppLicense) GetDefaultClassId() interface{} {
	return "storage.NetAppLicense"
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppLicense) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppLicense) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppLicense) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "storage.NetAppLicense" of the ObjectType field.
func (o *StorageNetAppLicense) GetDefaultObjectType() interface{} {
	return "storage.NetAppLicense"
}

// GetClusterUuid returns the ClusterUuid field value if set, zero value otherwise.
func (o *StorageNetAppLicense) GetClusterUuid() string {
	if o == nil || IsNil(o.ClusterUuid) {
		var ret string
		return ret
	}
	return *o.ClusterUuid
}

// GetClusterUuidOk returns a tuple with the ClusterUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLicense) GetClusterUuidOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterUuid) {
		return nil, false
	}
	return o.ClusterUuid, true
}

// HasClusterUuid returns a boolean if a field has been set.
func (o *StorageNetAppLicense) HasClusterUuid() bool {
	if o != nil && !IsNil(o.ClusterUuid) {
		return true
	}

	return false
}

// SetClusterUuid gets a reference to the given string and assigns it to the ClusterUuid field.
func (o *StorageNetAppLicense) SetClusterUuid(v string) {
	o.ClusterUuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageNetAppLicense) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLicense) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageNetAppLicense) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageNetAppLicense) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StorageNetAppLicense) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppLicense) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StorageNetAppLicense) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StorageNetAppLicense) SetState(v string) {
	o.State = &v
}

// GetArray returns the Array field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppLicense) GetArray() StorageNetAppClusterRelationship {
	if o == nil || IsNil(o.Array.Get()) {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array.Get()
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppLicense) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Array.Get(), o.Array.IsSet()
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppLicense) HasArray() bool {
	if o != nil && o.Array.IsSet() {
		return true
	}

	return false
}

// SetArray gets a reference to the given NullableStorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppLicense) SetArray(v StorageNetAppClusterRelationship) {
	o.Array.Set(&v)
}

// SetArrayNil sets the value for Array to be an explicit nil
func (o *StorageNetAppLicense) SetArrayNil() {
	o.Array.Set(nil)
}

// UnsetArray ensures that no value is present for Array, not even an explicit nil
func (o *StorageNetAppLicense) UnsetArray() {
	o.Array.Unset()
}

func (o StorageNetAppLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageNetAppLicense) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClusterUuid) {
		toSerialize["ClusterUuid"] = o.ClusterUuid
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if o.Array.IsSet() {
		toSerialize["Array"] = o.Array.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorageNetAppLicense) UnmarshalJSON(data []byte) (err error) {
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
	type StorageNetAppLicenseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Unique identity of the device.
		ClusterUuid *string `json:"ClusterUuid,omitempty" validate:"regexp=^$|^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
		// The name of the license package.
		Name *string `json:"Name,omitempty"`
		// Summary state of license package based on all installed licenses. * `Unknown` - The summary state of the license package is unknown. * `Compliant` - The summary state of the license package is compliant. * `Noncompliant` - The summary state of the license package is noncompliant. * `Unlicensed` - The summary state of the license package is unlicensed.
		State *string                                  `json:"State,omitempty"`
		Array NullableStorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppLicenseWithoutEmbeddedStruct := StorageNetAppLicenseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varStorageNetAppLicenseWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppLicense := _StorageNetAppLicense{}
		varStorageNetAppLicense.ClassId = varStorageNetAppLicenseWithoutEmbeddedStruct.ClassId
		varStorageNetAppLicense.ObjectType = varStorageNetAppLicenseWithoutEmbeddedStruct.ObjectType
		varStorageNetAppLicense.ClusterUuid = varStorageNetAppLicenseWithoutEmbeddedStruct.ClusterUuid
		varStorageNetAppLicense.Name = varStorageNetAppLicenseWithoutEmbeddedStruct.Name
		varStorageNetAppLicense.State = varStorageNetAppLicenseWithoutEmbeddedStruct.State
		varStorageNetAppLicense.Array = varStorageNetAppLicenseWithoutEmbeddedStruct.Array
		*o = StorageNetAppLicense(varStorageNetAppLicense)
	} else {
		return err
	}

	varStorageNetAppLicense := _StorageNetAppLicense{}

	err = json.Unmarshal(data, &varStorageNetAppLicense)
	if err == nil {
		o.MoBaseMo = varStorageNetAppLicense.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClusterUuid")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "State")
		delete(additionalProperties, "Array")

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

type NullableStorageNetAppLicense struct {
	value *StorageNetAppLicense
	isSet bool
}

func (v NullableStorageNetAppLicense) Get() *StorageNetAppLicense {
	return v.value
}

func (v *NullableStorageNetAppLicense) Set(val *StorageNetAppLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppLicense(val *StorageNetAppLicense) *NullableStorageNetAppLicense {
	return &NullableStorageNetAppLicense{value: val, isSet: true}
}

func (v NullableStorageNetAppLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
