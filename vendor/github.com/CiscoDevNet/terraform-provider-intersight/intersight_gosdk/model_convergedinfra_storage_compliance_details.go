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

// checks if the ConvergedinfraStorageComplianceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvergedinfraStorageComplianceDetails{}

// ConvergedinfraStorageComplianceDetails The compliance details of a storage in a converged infrastructure pod.
type ConvergedinfraStorageComplianceDetails struct {
	ConvergedinfraBaseComplianceDetails
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The operating system name and version (e.g. NetApp ONTAP 9.10) running on the storage array for which the compliance is getting evaluated.
	Os *string `json:"Os,omitempty"`
	// The protocol configured for the communication between the switch and the storage array.
	Protocol *string `json:"Protocol,omitempty"`
	// The reference device (e.g. adapter, fabric interconnect) against which the storage compliance is getting evaluated. * `Server` - The component type for a server in a converged infrastructure pod. * `Adapter` - The component type for an adapter on a server in a converged infrastructure pod. * `FabricInterconnect` - The component type for a fabric interconnect in a converged infrastructure pod. * `Nexus` - The component type for a nexus switch in a converged infrastructure pod. * `Storage` - The component type for a storage array in a converged infrastructure pod.
	RefDevice            *string                                                    `json:"RefDevice,omitempty"`
	AdapterCompliance    NullableConvergedinfraAdapterComplianceDetailsRelationship `json:"AdapterCompliance,omitempty"`
	PodCompliance        NullableConvergedinfraPodComplianceInfoRelationship        `json:"PodCompliance,omitempty"`
	StorageArray         NullableStorageBaseArrayRelationship                       `json:"StorageArray,omitempty"`
	SwitchCompliance     NullableConvergedinfraSwitchComplianceDetailsRelationship  `json:"SwitchCompliance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraStorageComplianceDetails ConvergedinfraStorageComplianceDetails

// NewConvergedinfraStorageComplianceDetails instantiates a new ConvergedinfraStorageComplianceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraStorageComplianceDetails(classId string, objectType string) *ConvergedinfraStorageComplianceDetails {
	this := ConvergedinfraStorageComplianceDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraStorageComplianceDetailsWithDefaults instantiates a new ConvergedinfraStorageComplianceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraStorageComplianceDetailsWithDefaults() *ConvergedinfraStorageComplianceDetails {
	this := ConvergedinfraStorageComplianceDetails{}
	var classId string = "convergedinfra.StorageComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.StorageComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraStorageComplianceDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraStorageComplianceDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "convergedinfra.StorageComplianceDetails" of the ClassId field.
func (o *ConvergedinfraStorageComplianceDetails) GetDefaultClassId() interface{} {
	return "convergedinfra.StorageComplianceDetails"
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraStorageComplianceDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraStorageComplianceDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "convergedinfra.StorageComplianceDetails" of the ObjectType field.
func (o *ConvergedinfraStorageComplianceDetails) GetDefaultObjectType() interface{} {
	return "convergedinfra.StorageComplianceDetails"
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetails) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetails) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ConvergedinfraStorageComplianceDetails) SetOs(v string) {
	o.Os = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetails) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetails) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ConvergedinfraStorageComplianceDetails) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRefDevice returns the RefDevice field value if set, zero value otherwise.
func (o *ConvergedinfraStorageComplianceDetails) GetRefDevice() string {
	if o == nil || IsNil(o.RefDevice) {
		var ret string
		return ret
	}
	return *o.RefDevice
}

// GetRefDeviceOk returns a tuple with the RefDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraStorageComplianceDetails) GetRefDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.RefDevice) {
		return nil, false
	}
	return o.RefDevice, true
}

// HasRefDevice returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasRefDevice() bool {
	if o != nil && !IsNil(o.RefDevice) {
		return true
	}

	return false
}

// SetRefDevice gets a reference to the given string and assigns it to the RefDevice field.
func (o *ConvergedinfraStorageComplianceDetails) SetRefDevice(v string) {
	o.RefDevice = &v
}

// GetAdapterCompliance returns the AdapterCompliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraStorageComplianceDetails) GetAdapterCompliance() ConvergedinfraAdapterComplianceDetailsRelationship {
	if o == nil || IsNil(o.AdapterCompliance.Get()) {
		var ret ConvergedinfraAdapterComplianceDetailsRelationship
		return ret
	}
	return *o.AdapterCompliance.Get()
}

// GetAdapterComplianceOk returns a tuple with the AdapterCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraStorageComplianceDetails) GetAdapterComplianceOk() (*ConvergedinfraAdapterComplianceDetailsRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdapterCompliance.Get(), o.AdapterCompliance.IsSet()
}

// HasAdapterCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasAdapterCompliance() bool {
	if o != nil && o.AdapterCompliance.IsSet() {
		return true
	}

	return false
}

// SetAdapterCompliance gets a reference to the given NullableConvergedinfraAdapterComplianceDetailsRelationship and assigns it to the AdapterCompliance field.
func (o *ConvergedinfraStorageComplianceDetails) SetAdapterCompliance(v ConvergedinfraAdapterComplianceDetailsRelationship) {
	o.AdapterCompliance.Set(&v)
}

// SetAdapterComplianceNil sets the value for AdapterCompliance to be an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) SetAdapterComplianceNil() {
	o.AdapterCompliance.Set(nil)
}

// UnsetAdapterCompliance ensures that no value is present for AdapterCompliance, not even an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) UnsetAdapterCompliance() {
	o.AdapterCompliance.Unset()
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraStorageComplianceDetails) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || IsNil(o.PodCompliance.Get()) {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance.Get()
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraStorageComplianceDetails) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PodCompliance.Get(), o.PodCompliance.IsSet()
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasPodCompliance() bool {
	if o != nil && o.PodCompliance.IsSet() {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given NullableConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraStorageComplianceDetails) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance.Set(&v)
}

// SetPodComplianceNil sets the value for PodCompliance to be an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) SetPodComplianceNil() {
	o.PodCompliance.Set(nil)
}

// UnsetPodCompliance ensures that no value is present for PodCompliance, not even an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) UnsetPodCompliance() {
	o.PodCompliance.Unset()
}

// GetStorageArray returns the StorageArray field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraStorageComplianceDetails) GetStorageArray() StorageBaseArrayRelationship {
	if o == nil || IsNil(o.StorageArray.Get()) {
		var ret StorageBaseArrayRelationship
		return ret
	}
	return *o.StorageArray.Get()
}

// GetStorageArrayOk returns a tuple with the StorageArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraStorageComplianceDetails) GetStorageArrayOk() (*StorageBaseArrayRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageArray.Get(), o.StorageArray.IsSet()
}

// HasStorageArray returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasStorageArray() bool {
	if o != nil && o.StorageArray.IsSet() {
		return true
	}

	return false
}

// SetStorageArray gets a reference to the given NullableStorageBaseArrayRelationship and assigns it to the StorageArray field.
func (o *ConvergedinfraStorageComplianceDetails) SetStorageArray(v StorageBaseArrayRelationship) {
	o.StorageArray.Set(&v)
}

// SetStorageArrayNil sets the value for StorageArray to be an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) SetStorageArrayNil() {
	o.StorageArray.Set(nil)
}

// UnsetStorageArray ensures that no value is present for StorageArray, not even an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) UnsetStorageArray() {
	o.StorageArray.Unset()
}

// GetSwitchCompliance returns the SwitchCompliance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraStorageComplianceDetails) GetSwitchCompliance() ConvergedinfraSwitchComplianceDetailsRelationship {
	if o == nil || IsNil(o.SwitchCompliance.Get()) {
		var ret ConvergedinfraSwitchComplianceDetailsRelationship
		return ret
	}
	return *o.SwitchCompliance.Get()
}

// GetSwitchComplianceOk returns a tuple with the SwitchCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraStorageComplianceDetails) GetSwitchComplianceOk() (*ConvergedinfraSwitchComplianceDetailsRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.SwitchCompliance.Get(), o.SwitchCompliance.IsSet()
}

// HasSwitchCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraStorageComplianceDetails) HasSwitchCompliance() bool {
	if o != nil && o.SwitchCompliance.IsSet() {
		return true
	}

	return false
}

// SetSwitchCompliance gets a reference to the given NullableConvergedinfraSwitchComplianceDetailsRelationship and assigns it to the SwitchCompliance field.
func (o *ConvergedinfraStorageComplianceDetails) SetSwitchCompliance(v ConvergedinfraSwitchComplianceDetailsRelationship) {
	o.SwitchCompliance.Set(&v)
}

// SetSwitchComplianceNil sets the value for SwitchCompliance to be an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) SetSwitchComplianceNil() {
	o.SwitchCompliance.Set(nil)
}

// UnsetSwitchCompliance ensures that no value is present for SwitchCompliance, not even an explicit nil
func (o *ConvergedinfraStorageComplianceDetails) UnsetSwitchCompliance() {
	o.SwitchCompliance.Unset()
}

func (o ConvergedinfraStorageComplianceDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvergedinfraStorageComplianceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedConvergedinfraBaseComplianceDetails, errConvergedinfraBaseComplianceDetails := json.Marshal(o.ConvergedinfraBaseComplianceDetails)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	errConvergedinfraBaseComplianceDetails = json.Unmarshal([]byte(serializedConvergedinfraBaseComplianceDetails), &toSerialize)
	if errConvergedinfraBaseComplianceDetails != nil {
		return map[string]interface{}{}, errConvergedinfraBaseComplianceDetails
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Os) {
		toSerialize["Os"] = o.Os
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.RefDevice) {
		toSerialize["RefDevice"] = o.RefDevice
	}
	if o.AdapterCompliance.IsSet() {
		toSerialize["AdapterCompliance"] = o.AdapterCompliance.Get()
	}
	if o.PodCompliance.IsSet() {
		toSerialize["PodCompliance"] = o.PodCompliance.Get()
	}
	if o.StorageArray.IsSet() {
		toSerialize["StorageArray"] = o.StorageArray.Get()
	}
	if o.SwitchCompliance.IsSet() {
		toSerialize["SwitchCompliance"] = o.SwitchCompliance.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConvergedinfraStorageComplianceDetails) UnmarshalJSON(data []byte) (err error) {
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
	type ConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The operating system name and version (e.g. NetApp ONTAP 9.10) running on the storage array for which the compliance is getting evaluated.
		Os *string `json:"Os,omitempty"`
		// The protocol configured for the communication between the switch and the storage array.
		Protocol *string `json:"Protocol,omitempty"`
		// The reference device (e.g. adapter, fabric interconnect) against which the storage compliance is getting evaluated. * `Server` - The component type for a server in a converged infrastructure pod. * `Adapter` - The component type for an adapter on a server in a converged infrastructure pod. * `FabricInterconnect` - The component type for a fabric interconnect in a converged infrastructure pod. * `Nexus` - The component type for a nexus switch in a converged infrastructure pod. * `Storage` - The component type for a storage array in a converged infrastructure pod.
		RefDevice         *string                                                    `json:"RefDevice,omitempty"`
		AdapterCompliance NullableConvergedinfraAdapterComplianceDetailsRelationship `json:"AdapterCompliance,omitempty"`
		PodCompliance     NullableConvergedinfraPodComplianceInfoRelationship        `json:"PodCompliance,omitempty"`
		StorageArray      NullableStorageBaseArrayRelationship                       `json:"StorageArray,omitempty"`
		SwitchCompliance  NullableConvergedinfraSwitchComplianceDetailsRelationship  `json:"SwitchCompliance,omitempty"`
	}

	varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct := ConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct)
	if err == nil {
		varConvergedinfraStorageComplianceDetails := _ConvergedinfraStorageComplianceDetails{}
		varConvergedinfraStorageComplianceDetails.ClassId = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.ClassId
		varConvergedinfraStorageComplianceDetails.ObjectType = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.ObjectType
		varConvergedinfraStorageComplianceDetails.Os = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.Os
		varConvergedinfraStorageComplianceDetails.Protocol = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.Protocol
		varConvergedinfraStorageComplianceDetails.RefDevice = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.RefDevice
		varConvergedinfraStorageComplianceDetails.AdapterCompliance = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.AdapterCompliance
		varConvergedinfraStorageComplianceDetails.PodCompliance = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.PodCompliance
		varConvergedinfraStorageComplianceDetails.StorageArray = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.StorageArray
		varConvergedinfraStorageComplianceDetails.SwitchCompliance = varConvergedinfraStorageComplianceDetailsWithoutEmbeddedStruct.SwitchCompliance
		*o = ConvergedinfraStorageComplianceDetails(varConvergedinfraStorageComplianceDetails)
	} else {
		return err
	}

	varConvergedinfraStorageComplianceDetails := _ConvergedinfraStorageComplianceDetails{}

	err = json.Unmarshal(data, &varConvergedinfraStorageComplianceDetails)
	if err == nil {
		o.ConvergedinfraBaseComplianceDetails = varConvergedinfraStorageComplianceDetails.ConvergedinfraBaseComplianceDetails
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Os")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RefDevice")
		delete(additionalProperties, "AdapterCompliance")
		delete(additionalProperties, "PodCompliance")
		delete(additionalProperties, "StorageArray")
		delete(additionalProperties, "SwitchCompliance")

		// remove fields from embedded structs
		reflectConvergedinfraBaseComplianceDetails := reflect.ValueOf(o.ConvergedinfraBaseComplianceDetails)
		for i := 0; i < reflectConvergedinfraBaseComplianceDetails.Type().NumField(); i++ {
			t := reflectConvergedinfraBaseComplianceDetails.Type().Field(i)

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

type NullableConvergedinfraStorageComplianceDetails struct {
	value *ConvergedinfraStorageComplianceDetails
	isSet bool
}

func (v NullableConvergedinfraStorageComplianceDetails) Get() *ConvergedinfraStorageComplianceDetails {
	return v.value
}

func (v *NullableConvergedinfraStorageComplianceDetails) Set(val *ConvergedinfraStorageComplianceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraStorageComplianceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraStorageComplianceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraStorageComplianceDetails(val *ConvergedinfraStorageComplianceDetails) *NullableConvergedinfraStorageComplianceDetails {
	return &NullableConvergedinfraStorageComplianceDetails{value: val, isSet: true}
}

func (v NullableConvergedinfraStorageComplianceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraStorageComplianceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
