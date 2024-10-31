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

// checks if the EquipmentBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquipmentBase{}

// EquipmentBase Abstract base class for all equipments which have a vendor /model / serial.
type EquipmentBase struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// This field indicates the compute status of the catalog values for the associated component or hardware.
	IsUpgraded *bool `json:"IsUpgraded,omitempty"`
	// This field displays the model number of the associated component or hardware.
	Model *string `json:"Model,omitempty"`
	// This field indicates the presence (equipped) or absence (absent) of the associated component or hardware.
	Presence *string `json:"Presence,omitempty"`
	// This field displays the revised version of the associated component or hardware (if any).
	Revision *string `json:"Revision,omitempty"`
	// This field displays the serial number of the associated component or hardware.
	Serial *string `json:"Serial,omitempty"`
	// This field displays the vendor information of the associated component or hardware.
	Vendor               *string                          `json:"Vendor,omitempty"`
	PreviousFru          NullableEquipmentFruRelationship `json:"PreviousFru,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentBase EquipmentBase

// NewEquipmentBase instantiates a new EquipmentBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentBase(classId string, objectType string) *EquipmentBase {
	this := EquipmentBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentBaseWithDefaults instantiates a new EquipmentBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentBaseWithDefaults() *EquipmentBase {
	this := EquipmentBase{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsUpgraded returns the IsUpgraded field value if set, zero value otherwise.
func (o *EquipmentBase) GetIsUpgraded() bool {
	if o == nil || IsNil(o.IsUpgraded) {
		var ret bool
		return ret
	}
	return *o.IsUpgraded
}

// GetIsUpgradedOk returns a tuple with the IsUpgraded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetIsUpgradedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUpgraded) {
		return nil, false
	}
	return o.IsUpgraded, true
}

// HasIsUpgraded returns a boolean if a field has been set.
func (o *EquipmentBase) HasIsUpgraded() bool {
	if o != nil && !IsNil(o.IsUpgraded) {
		return true
	}

	return false
}

// SetIsUpgraded gets a reference to the given bool and assigns it to the IsUpgraded field.
func (o *EquipmentBase) SetIsUpgraded(v bool) {
	o.IsUpgraded = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EquipmentBase) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EquipmentBase) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EquipmentBase) SetModel(v string) {
	o.Model = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentBase) GetPresence() string {
	if o == nil || IsNil(o.Presence) {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetPresenceOk() (*string, bool) {
	if o == nil || IsNil(o.Presence) {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentBase) HasPresence() bool {
	if o != nil && !IsNil(o.Presence) {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentBase) SetPresence(v string) {
	o.Presence = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *EquipmentBase) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *EquipmentBase) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *EquipmentBase) SetRevision(v string) {
	o.Revision = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EquipmentBase) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EquipmentBase) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EquipmentBase) SetSerial(v string) {
	o.Serial = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *EquipmentBase) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentBase) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *EquipmentBase) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *EquipmentBase) SetVendor(v string) {
	o.Vendor = &v
}

// GetPreviousFru returns the PreviousFru field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentBase) GetPreviousFru() EquipmentFruRelationship {
	if o == nil || IsNil(o.PreviousFru.Get()) {
		var ret EquipmentFruRelationship
		return ret
	}
	return *o.PreviousFru.Get()
}

// GetPreviousFruOk returns a tuple with the PreviousFru field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentBase) GetPreviousFruOk() (*EquipmentFruRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousFru.Get(), o.PreviousFru.IsSet()
}

// HasPreviousFru returns a boolean if a field has been set.
func (o *EquipmentBase) HasPreviousFru() bool {
	if o != nil && o.PreviousFru.IsSet() {
		return true
	}

	return false
}

// SetPreviousFru gets a reference to the given NullableEquipmentFruRelationship and assigns it to the PreviousFru field.
func (o *EquipmentBase) SetPreviousFru(v EquipmentFruRelationship) {
	o.PreviousFru.Set(&v)
}

// SetPreviousFruNil sets the value for PreviousFru to be an explicit nil
func (o *EquipmentBase) SetPreviousFruNil() {
	o.PreviousFru.Set(nil)
}

// UnsetPreviousFru ensures that no value is present for PreviousFru, not even an explicit nil
func (o *EquipmentBase) UnsetPreviousFru() {
	o.PreviousFru.Unset()
}

func (o EquipmentBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquipmentBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return map[string]interface{}{}, errInventoryBase
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.IsUpgraded) {
		toSerialize["IsUpgraded"] = o.IsUpgraded
	}
	if !IsNil(o.Model) {
		toSerialize["Model"] = o.Model
	}
	if !IsNil(o.Presence) {
		toSerialize["Presence"] = o.Presence
	}
	if !IsNil(o.Revision) {
		toSerialize["Revision"] = o.Revision
	}
	if !IsNil(o.Serial) {
		toSerialize["Serial"] = o.Serial
	}
	if !IsNil(o.Vendor) {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.PreviousFru.IsSet() {
		toSerialize["PreviousFru"] = o.PreviousFru.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EquipmentBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type EquipmentBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// This field indicates the compute status of the catalog values for the associated component or hardware.
		IsUpgraded *bool `json:"IsUpgraded,omitempty"`
		// This field displays the model number of the associated component or hardware.
		Model *string `json:"Model,omitempty"`
		// This field indicates the presence (equipped) or absence (absent) of the associated component or hardware.
		Presence *string `json:"Presence,omitempty"`
		// This field displays the revised version of the associated component or hardware (if any).
		Revision *string `json:"Revision,omitempty"`
		// This field displays the serial number of the associated component or hardware.
		Serial *string `json:"Serial,omitempty"`
		// This field displays the vendor information of the associated component or hardware.
		Vendor      *string                          `json:"Vendor,omitempty"`
		PreviousFru NullableEquipmentFruRelationship `json:"PreviousFru,omitempty"`
	}

	varEquipmentBaseWithoutEmbeddedStruct := EquipmentBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varEquipmentBaseWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentBase := _EquipmentBase{}
		varEquipmentBase.ClassId = varEquipmentBaseWithoutEmbeddedStruct.ClassId
		varEquipmentBase.ObjectType = varEquipmentBaseWithoutEmbeddedStruct.ObjectType
		varEquipmentBase.IsUpgraded = varEquipmentBaseWithoutEmbeddedStruct.IsUpgraded
		varEquipmentBase.Model = varEquipmentBaseWithoutEmbeddedStruct.Model
		varEquipmentBase.Presence = varEquipmentBaseWithoutEmbeddedStruct.Presence
		varEquipmentBase.Revision = varEquipmentBaseWithoutEmbeddedStruct.Revision
		varEquipmentBase.Serial = varEquipmentBaseWithoutEmbeddedStruct.Serial
		varEquipmentBase.Vendor = varEquipmentBaseWithoutEmbeddedStruct.Vendor
		varEquipmentBase.PreviousFru = varEquipmentBaseWithoutEmbeddedStruct.PreviousFru
		*o = EquipmentBase(varEquipmentBase)
	} else {
		return err
	}

	varEquipmentBase := _EquipmentBase{}

	err = json.Unmarshal(data, &varEquipmentBase)
	if err == nil {
		o.InventoryBase = varEquipmentBase.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsUpgraded")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Revision")
		delete(additionalProperties, "Serial")
		delete(additionalProperties, "Vendor")
		delete(additionalProperties, "PreviousFru")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableEquipmentBase struct {
	value *EquipmentBase
	isSet bool
}

func (v NullableEquipmentBase) Get() *EquipmentBase {
	return v.value
}

func (v *NullableEquipmentBase) Set(val *EquipmentBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentBase(val *EquipmentBase) *NullableEquipmentBase {
	return &NullableEquipmentBase{value: val, isSet: true}
}

func (v NullableEquipmentBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
