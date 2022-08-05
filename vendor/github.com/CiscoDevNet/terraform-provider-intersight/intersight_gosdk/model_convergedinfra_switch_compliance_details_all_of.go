/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7658
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// ConvergedinfraSwitchComplianceDetailsAllOf Definition of the list of properties defined in 'convergedinfra.SwitchComplianceDetails', excluding properties defined in parent classes.
type ConvergedinfraSwitchComplianceDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The firmware of the switch as received from inventory.
	Firmware *string `json:"Firmware,omitempty"`
	// The model information of the switch.
	Model *string `json:"Model,omitempty"`
	// The type of switch component. It must be set to either Fabric Interconnect, Nexus or MDS. * `FabricInterconnect` - The default Switch type of UCSM and IMM mode devices. * `NexusDevice` - Switch type of Nexus devices. * `MDSDevice` - Switch type of Nexus MDS devices.
	Type          *string                                      `json:"Type,omitempty"`
	PodCompliance *ConvergedinfraPodComplianceInfoRelationship `json:"PodCompliance,omitempty"`
	// An array of relationships to convergedinfraStorageComplianceDetails resources.
	StorageCompliances   []ConvergedinfraStorageComplianceDetailsRelationship `json:"StorageCompliances,omitempty"`
	Switch               *NetworkElementSummaryRelationship                   `json:"Switch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraSwitchComplianceDetailsAllOf ConvergedinfraSwitchComplianceDetailsAllOf

// NewConvergedinfraSwitchComplianceDetailsAllOf instantiates a new ConvergedinfraSwitchComplianceDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraSwitchComplianceDetailsAllOf(classId string, objectType string) *ConvergedinfraSwitchComplianceDetailsAllOf {
	this := ConvergedinfraSwitchComplianceDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraSwitchComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraSwitchComplianceDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraSwitchComplianceDetailsAllOfWithDefaults() *ConvergedinfraSwitchComplianceDetailsAllOf {
	this := ConvergedinfraSwitchComplianceDetailsAllOf{}
	var classId string = "convergedinfra.SwitchComplianceDetails"
	this.ClassId = classId
	var objectType string = "convergedinfra.SwitchComplianceDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetFirmware() string {
	if o == nil || o.Firmware == nil {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetFirmwareOk() (*string, bool) {
	if o == nil || o.Firmware == nil {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasFirmware() bool {
	if o != nil && o.Firmware != nil {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetFirmware(v string) {
	o.Firmware = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetModel(v string) {
	o.Model = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetType(v string) {
	o.Type = &v
}

// GetPodCompliance returns the PodCompliance field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetPodCompliance() ConvergedinfraPodComplianceInfoRelationship {
	if o == nil || o.PodCompliance == nil {
		var ret ConvergedinfraPodComplianceInfoRelationship
		return ret
	}
	return *o.PodCompliance
}

// GetPodComplianceOk returns a tuple with the PodCompliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetPodComplianceOk() (*ConvergedinfraPodComplianceInfoRelationship, bool) {
	if o == nil || o.PodCompliance == nil {
		return nil, false
	}
	return o.PodCompliance, true
}

// HasPodCompliance returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasPodCompliance() bool {
	if o != nil && o.PodCompliance != nil {
		return true
	}

	return false
}

// SetPodCompliance gets a reference to the given ConvergedinfraPodComplianceInfoRelationship and assigns it to the PodCompliance field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetPodCompliance(v ConvergedinfraPodComplianceInfoRelationship) {
	o.PodCompliance = &v
}

// GetStorageCompliances returns the StorageCompliances field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetStorageCompliances() []ConvergedinfraStorageComplianceDetailsRelationship {
	if o == nil {
		var ret []ConvergedinfraStorageComplianceDetailsRelationship
		return ret
	}
	return o.StorageCompliances
}

// GetStorageCompliancesOk returns a tuple with the StorageCompliances field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetStorageCompliancesOk() ([]ConvergedinfraStorageComplianceDetailsRelationship, bool) {
	if o == nil || o.StorageCompliances == nil {
		return nil, false
	}
	return o.StorageCompliances, true
}

// HasStorageCompliances returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasStorageCompliances() bool {
	if o != nil && o.StorageCompliances != nil {
		return true
	}

	return false
}

// SetStorageCompliances gets a reference to the given []ConvergedinfraStorageComplianceDetailsRelationship and assigns it to the StorageCompliances field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetStorageCompliances(v []ConvergedinfraStorageComplianceDetailsRelationship) {
	o.StorageCompliances = v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetSwitch() NetworkElementSummaryRelationship {
	if o == nil || o.Switch == nil {
		var ret NetworkElementSummaryRelationship
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) GetSwitchOk() (*NetworkElementSummaryRelationship, bool) {
	if o == nil || o.Switch == nil {
		return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) HasSwitch() bool {
	if o != nil && o.Switch != nil {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given NetworkElementSummaryRelationship and assigns it to the Switch field.
func (o *ConvergedinfraSwitchComplianceDetailsAllOf) SetSwitch(v NetworkElementSummaryRelationship) {
	o.Switch = &v
}

func (o ConvergedinfraSwitchComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Firmware != nil {
		toSerialize["Firmware"] = o.Firmware
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.PodCompliance != nil {
		toSerialize["PodCompliance"] = o.PodCompliance
	}
	if o.StorageCompliances != nil {
		toSerialize["StorageCompliances"] = o.StorageCompliances
	}
	if o.Switch != nil {
		toSerialize["Switch"] = o.Switch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraSwitchComplianceDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraSwitchComplianceDetailsAllOf := _ConvergedinfraSwitchComplianceDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraSwitchComplianceDetailsAllOf); err == nil {
		*o = ConvergedinfraSwitchComplianceDetailsAllOf(varConvergedinfraSwitchComplianceDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Firmware")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "PodCompliance")
		delete(additionalProperties, "StorageCompliances")
		delete(additionalProperties, "Switch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraSwitchComplianceDetailsAllOf struct {
	value *ConvergedinfraSwitchComplianceDetailsAllOf
	isSet bool
}

func (v NullableConvergedinfraSwitchComplianceDetailsAllOf) Get() *ConvergedinfraSwitchComplianceDetailsAllOf {
	return v.value
}

func (v *NullableConvergedinfraSwitchComplianceDetailsAllOf) Set(val *ConvergedinfraSwitchComplianceDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraSwitchComplianceDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraSwitchComplianceDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraSwitchComplianceDetailsAllOf(val *ConvergedinfraSwitchComplianceDetailsAllOf) *NullableConvergedinfraSwitchComplianceDetailsAllOf {
	return &NullableConvergedinfraSwitchComplianceDetailsAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraSwitchComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraSwitchComplianceDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}