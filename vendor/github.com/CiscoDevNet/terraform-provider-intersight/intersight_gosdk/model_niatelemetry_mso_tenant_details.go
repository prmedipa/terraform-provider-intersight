/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
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

// checks if the NiatelemetryMsoTenantDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiatelemetryMsoTenantDetails{}

// NiatelemetryMsoTenantDetails Details of tenant in Multi-Site Orchestrator.
type NiatelemetryMsoTenantDetails struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Site IDs to which this tenant is deployed to.
	DeployedSites *string `json:"DeployedSites,omitempty"`
	// Number of schemas assigned to each tenant in Multi-Site Orchestrator.
	NumberOfSchemasAssignedPerTenantInMso *int64 `json:"NumberOfSchemasAssignedPerTenantInMso,omitempty"`
	// Number of sites each tenant is deployed to.
	SitesEachTenantIsDeployedToInMso *int64 `json:"SitesEachTenantIsDeployedToInMso,omitempty"`
	// ID of tenant in Multi-Site Orchestrator.
	TenantId *string `json:"TenantId,omitempty"`
	// Name of the tenant in Multi-Site Orchestrator.
	TenantName           *string                                     `json:"TenantName,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryMsoTenantDetails NiatelemetryMsoTenantDetails

// NewNiatelemetryMsoTenantDetails instantiates a new NiatelemetryMsoTenantDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoTenantDetails(classId string, objectType string) *NiatelemetryMsoTenantDetails {
	this := NiatelemetryMsoTenantDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoTenantDetailsWithDefaults instantiates a new NiatelemetryMsoTenantDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoTenantDetailsWithDefaults() *NiatelemetryMsoTenantDetails {
	this := NiatelemetryMsoTenantDetails{}
	var classId string = "niatelemetry.MsoTenantDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoTenantDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoTenantDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoTenantDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "niatelemetry.MsoTenantDetails" of the ClassId field.
func (o *NiatelemetryMsoTenantDetails) GetDefaultClassId() interface{} {
	return "niatelemetry.MsoTenantDetails"
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoTenantDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoTenantDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "niatelemetry.MsoTenantDetails" of the ObjectType field.
func (o *NiatelemetryMsoTenantDetails) GetDefaultObjectType() interface{} {
	return "niatelemetry.MsoTenantDetails"
}

// GetDeployedSites returns the DeployedSites field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetDeployedSites() string {
	if o == nil || IsNil(o.DeployedSites) {
		var ret string
		return ret
	}
	return *o.DeployedSites
}

// GetDeployedSitesOk returns a tuple with the DeployedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetDeployedSitesOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedSites) {
		return nil, false
	}
	return o.DeployedSites, true
}

// HasDeployedSites returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasDeployedSites() bool {
	if o != nil && !IsNil(o.DeployedSites) {
		return true
	}

	return false
}

// SetDeployedSites gets a reference to the given string and assigns it to the DeployedSites field.
func (o *NiatelemetryMsoTenantDetails) SetDeployedSites(v string) {
	o.DeployedSites = &v
}

// GetNumberOfSchemasAssignedPerTenantInMso returns the NumberOfSchemasAssignedPerTenantInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMso() int64 {
	if o == nil || IsNil(o.NumberOfSchemasAssignedPerTenantInMso) {
		var ret int64
		return ret
	}
	return *o.NumberOfSchemasAssignedPerTenantInMso
}

// GetNumberOfSchemasAssignedPerTenantInMsoOk returns a tuple with the NumberOfSchemasAssignedPerTenantInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMsoOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfSchemasAssignedPerTenantInMso) {
		return nil, false
	}
	return o.NumberOfSchemasAssignedPerTenantInMso, true
}

// HasNumberOfSchemasAssignedPerTenantInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasNumberOfSchemasAssignedPerTenantInMso() bool {
	if o != nil && !IsNil(o.NumberOfSchemasAssignedPerTenantInMso) {
		return true
	}

	return false
}

// SetNumberOfSchemasAssignedPerTenantInMso gets a reference to the given int64 and assigns it to the NumberOfSchemasAssignedPerTenantInMso field.
func (o *NiatelemetryMsoTenantDetails) SetNumberOfSchemasAssignedPerTenantInMso(v int64) {
	o.NumberOfSchemasAssignedPerTenantInMso = &v
}

// GetSitesEachTenantIsDeployedToInMso returns the SitesEachTenantIsDeployedToInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMso() int64 {
	if o == nil || IsNil(o.SitesEachTenantIsDeployedToInMso) {
		var ret int64
		return ret
	}
	return *o.SitesEachTenantIsDeployedToInMso
}

// GetSitesEachTenantIsDeployedToInMsoOk returns a tuple with the SitesEachTenantIsDeployedToInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMsoOk() (*int64, bool) {
	if o == nil || IsNil(o.SitesEachTenantIsDeployedToInMso) {
		return nil, false
	}
	return o.SitesEachTenantIsDeployedToInMso, true
}

// HasSitesEachTenantIsDeployedToInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasSitesEachTenantIsDeployedToInMso() bool {
	if o != nil && !IsNil(o.SitesEachTenantIsDeployedToInMso) {
		return true
	}

	return false
}

// SetSitesEachTenantIsDeployedToInMso gets a reference to the given int64 and assigns it to the SitesEachTenantIsDeployedToInMso field.
func (o *NiatelemetryMsoTenantDetails) SetSitesEachTenantIsDeployedToInMso(v int64) {
	o.SitesEachTenantIsDeployedToInMso = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *NiatelemetryMsoTenantDetails) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetTenantName() string {
	if o == nil || IsNil(o.TenantName) {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetTenantNameOk() (*string, bool) {
	if o == nil || IsNil(o.TenantName) {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasTenantName() bool {
	if o != nil && !IsNil(o.TenantName) {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *NiatelemetryMsoTenantDetails) SetTenantName(v string) {
	o.TenantName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NiatelemetryMsoTenantDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NiatelemetryMsoTenantDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoTenantDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *NiatelemetryMsoTenantDetails) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *NiatelemetryMsoTenantDetails) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o NiatelemetryMsoTenantDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiatelemetryMsoTenantDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeployedSites) {
		toSerialize["DeployedSites"] = o.DeployedSites
	}
	if !IsNil(o.NumberOfSchemasAssignedPerTenantInMso) {
		toSerialize["NumberOfSchemasAssignedPerTenantInMso"] = o.NumberOfSchemasAssignedPerTenantInMso
	}
	if !IsNil(o.SitesEachTenantIsDeployedToInMso) {
		toSerialize["SitesEachTenantIsDeployedToInMso"] = o.SitesEachTenantIsDeployedToInMso
	}
	if !IsNil(o.TenantId) {
		toSerialize["TenantId"] = o.TenantId
	}
	if !IsNil(o.TenantName) {
		toSerialize["TenantName"] = o.TenantName
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NiatelemetryMsoTenantDetails) UnmarshalJSON(data []byte) (err error) {
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
	type NiatelemetryMsoTenantDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Site IDs to which this tenant is deployed to.
		DeployedSites *string `json:"DeployedSites,omitempty"`
		// Number of schemas assigned to each tenant in Multi-Site Orchestrator.
		NumberOfSchemasAssignedPerTenantInMso *int64 `json:"NumberOfSchemasAssignedPerTenantInMso,omitempty"`
		// Number of sites each tenant is deployed to.
		SitesEachTenantIsDeployedToInMso *int64 `json:"SitesEachTenantIsDeployedToInMso,omitempty"`
		// ID of tenant in Multi-Site Orchestrator.
		TenantId *string `json:"TenantId,omitempty"`
		// Name of the tenant in Multi-Site Orchestrator.
		TenantName       *string                                     `json:"TenantName,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct := NiatelemetryMsoTenantDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct)
	if err == nil {
		varNiatelemetryMsoTenantDetails := _NiatelemetryMsoTenantDetails{}
		varNiatelemetryMsoTenantDetails.ClassId = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.ClassId
		varNiatelemetryMsoTenantDetails.ObjectType = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.ObjectType
		varNiatelemetryMsoTenantDetails.DeployedSites = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.DeployedSites
		varNiatelemetryMsoTenantDetails.NumberOfSchemasAssignedPerTenantInMso = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.NumberOfSchemasAssignedPerTenantInMso
		varNiatelemetryMsoTenantDetails.SitesEachTenantIsDeployedToInMso = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.SitesEachTenantIsDeployedToInMso
		varNiatelemetryMsoTenantDetails.TenantId = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.TenantId
		varNiatelemetryMsoTenantDetails.TenantName = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.TenantName
		varNiatelemetryMsoTenantDetails.RegisteredDevice = varNiatelemetryMsoTenantDetailsWithoutEmbeddedStruct.RegisteredDevice
		*o = NiatelemetryMsoTenantDetails(varNiatelemetryMsoTenantDetails)
	} else {
		return err
	}

	varNiatelemetryMsoTenantDetails := _NiatelemetryMsoTenantDetails{}

	err = json.Unmarshal(data, &varNiatelemetryMsoTenantDetails)
	if err == nil {
		o.MoBaseMo = varNiatelemetryMsoTenantDetails.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeployedSites")
		delete(additionalProperties, "NumberOfSchemasAssignedPerTenantInMso")
		delete(additionalProperties, "SitesEachTenantIsDeployedToInMso")
		delete(additionalProperties, "TenantId")
		delete(additionalProperties, "TenantName")
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

type NullableNiatelemetryMsoTenantDetails struct {
	value *NiatelemetryMsoTenantDetails
	isSet bool
}

func (v NullableNiatelemetryMsoTenantDetails) Get() *NiatelemetryMsoTenantDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoTenantDetails) Set(val *NiatelemetryMsoTenantDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoTenantDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoTenantDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoTenantDetails(val *NiatelemetryMsoTenantDetails) *NullableNiatelemetryMsoTenantDetails {
	return &NullableNiatelemetryMsoTenantDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoTenantDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoTenantDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
