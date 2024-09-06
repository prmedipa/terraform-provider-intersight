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

// checks if the FabricUplinkPcRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricUplinkPcRole{}

// FabricUplinkPcRole Object sent by user to configure a ethernet uplink port-channel on the collection of ports.
type FabricUplinkPcRole struct {
	FabricPortChannelRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps. * `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Forward error correction configuration for Uplink Port Channel member ports. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
	Fec *string `json:"Fec,omitempty"`
	// An array of relationships to fabricEthNetworkGroupPolicy resources.
	EthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship       `json:"EthNetworkGroupPolicy,omitempty"`
	FlowControlPolicy     NullableFabricFlowControlPolicyRelationship     `json:"FlowControlPolicy,omitempty"`
	LinkAggregationPolicy NullableFabricLinkAggregationPolicyRelationship `json:"LinkAggregationPolicy,omitempty"`
	LinkControlPolicy     NullableFabricLinkControlPolicyRelationship     `json:"LinkControlPolicy,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _FabricUplinkPcRole FabricUplinkPcRole

// NewFabricUplinkPcRole instantiates a new FabricUplinkPcRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricUplinkPcRole(classId string, objectType string) *FabricUplinkPcRole {
	this := FabricUplinkPcRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// NewFabricUplinkPcRoleWithDefaults instantiates a new FabricUplinkPcRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricUplinkPcRoleWithDefaults() *FabricUplinkPcRole {
	this := FabricUplinkPcRole{}
	var classId string = "fabric.UplinkPcRole"
	this.ClassId = classId
	var objectType string = "fabric.UplinkPcRole"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var fec string = "Auto"
	this.Fec = &fec
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricUplinkPcRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricUplinkPcRole) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "fabric.UplinkPcRole" of the ClassId field.
func (o *FabricUplinkPcRole) GetDefaultClassId() interface{} {
	return "fabric.UplinkPcRole"
}

// GetObjectType returns the ObjectType field value
func (o *FabricUplinkPcRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricUplinkPcRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "fabric.UplinkPcRole" of the ObjectType field.
func (o *FabricUplinkPcRole) GetDefaultObjectType() interface{} {
	return "fabric.UplinkPcRole"
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricUplinkPcRole) GetAdminSpeed() string {
	if o == nil || IsNil(o.AdminSpeed) {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.AdminSpeed) {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasAdminSpeed() bool {
	if o != nil && !IsNil(o.AdminSpeed) {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricUplinkPcRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetFec returns the Fec field value if set, zero value otherwise.
func (o *FabricUplinkPcRole) GetFec() string {
	if o == nil || IsNil(o.Fec) {
		var ret string
		return ret
	}
	return *o.Fec
}

// GetFecOk returns a tuple with the Fec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricUplinkPcRole) GetFecOk() (*string, bool) {
	if o == nil || IsNil(o.Fec) {
		return nil, false
	}
	return o.Fec, true
}

// HasFec returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasFec() bool {
	if o != nil && !IsNil(o.Fec) {
		return true
	}

	return false
}

// SetFec gets a reference to the given string and assigns it to the Fec field.
func (o *FabricUplinkPcRole) SetFec(v string) {
	o.Fec = &v
}

// GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkPcRole) GetEthNetworkGroupPolicy() []FabricEthNetworkGroupPolicyRelationship {
	if o == nil {
		var ret []FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return o.EthNetworkGroupPolicy
}

// GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkPcRole) GetEthNetworkGroupPolicyOk() ([]FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || IsNil(o.EthNetworkGroupPolicy) {
		return nil, false
	}
	return o.EthNetworkGroupPolicy, true
}

// HasEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasEthNetworkGroupPolicy() bool {
	if o != nil && !IsNil(o.EthNetworkGroupPolicy) {
		return true
	}

	return false
}

// SetEthNetworkGroupPolicy gets a reference to the given []FabricEthNetworkGroupPolicyRelationship and assigns it to the EthNetworkGroupPolicy field.
func (o *FabricUplinkPcRole) SetEthNetworkGroupPolicy(v []FabricEthNetworkGroupPolicyRelationship) {
	o.EthNetworkGroupPolicy = v
}

// GetFlowControlPolicy returns the FlowControlPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkPcRole) GetFlowControlPolicy() FabricFlowControlPolicyRelationship {
	if o == nil || IsNil(o.FlowControlPolicy.Get()) {
		var ret FabricFlowControlPolicyRelationship
		return ret
	}
	return *o.FlowControlPolicy.Get()
}

// GetFlowControlPolicyOk returns a tuple with the FlowControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkPcRole) GetFlowControlPolicyOk() (*FabricFlowControlPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowControlPolicy.Get(), o.FlowControlPolicy.IsSet()
}

// HasFlowControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasFlowControlPolicy() bool {
	if o != nil && o.FlowControlPolicy.IsSet() {
		return true
	}

	return false
}

// SetFlowControlPolicy gets a reference to the given NullableFabricFlowControlPolicyRelationship and assigns it to the FlowControlPolicy field.
func (o *FabricUplinkPcRole) SetFlowControlPolicy(v FabricFlowControlPolicyRelationship) {
	o.FlowControlPolicy.Set(&v)
}

// SetFlowControlPolicyNil sets the value for FlowControlPolicy to be an explicit nil
func (o *FabricUplinkPcRole) SetFlowControlPolicyNil() {
	o.FlowControlPolicy.Set(nil)
}

// UnsetFlowControlPolicy ensures that no value is present for FlowControlPolicy, not even an explicit nil
func (o *FabricUplinkPcRole) UnsetFlowControlPolicy() {
	o.FlowControlPolicy.Unset()
}

// GetLinkAggregationPolicy returns the LinkAggregationPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkPcRole) GetLinkAggregationPolicy() FabricLinkAggregationPolicyRelationship {
	if o == nil || IsNil(o.LinkAggregationPolicy.Get()) {
		var ret FabricLinkAggregationPolicyRelationship
		return ret
	}
	return *o.LinkAggregationPolicy.Get()
}

// GetLinkAggregationPolicyOk returns a tuple with the LinkAggregationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkPcRole) GetLinkAggregationPolicyOk() (*FabricLinkAggregationPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkAggregationPolicy.Get(), o.LinkAggregationPolicy.IsSet()
}

// HasLinkAggregationPolicy returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasLinkAggregationPolicy() bool {
	if o != nil && o.LinkAggregationPolicy.IsSet() {
		return true
	}

	return false
}

// SetLinkAggregationPolicy gets a reference to the given NullableFabricLinkAggregationPolicyRelationship and assigns it to the LinkAggregationPolicy field.
func (o *FabricUplinkPcRole) SetLinkAggregationPolicy(v FabricLinkAggregationPolicyRelationship) {
	o.LinkAggregationPolicy.Set(&v)
}

// SetLinkAggregationPolicyNil sets the value for LinkAggregationPolicy to be an explicit nil
func (o *FabricUplinkPcRole) SetLinkAggregationPolicyNil() {
	o.LinkAggregationPolicy.Set(nil)
}

// UnsetLinkAggregationPolicy ensures that no value is present for LinkAggregationPolicy, not even an explicit nil
func (o *FabricUplinkPcRole) UnsetLinkAggregationPolicy() {
	o.LinkAggregationPolicy.Unset()
}

// GetLinkControlPolicy returns the LinkControlPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricUplinkPcRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship {
	if o == nil || IsNil(o.LinkControlPolicy.Get()) {
		var ret FabricLinkControlPolicyRelationship
		return ret
	}
	return *o.LinkControlPolicy.Get()
}

// GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricUplinkPcRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkControlPolicy.Get(), o.LinkControlPolicy.IsSet()
}

// HasLinkControlPolicy returns a boolean if a field has been set.
func (o *FabricUplinkPcRole) HasLinkControlPolicy() bool {
	if o != nil && o.LinkControlPolicy.IsSet() {
		return true
	}

	return false
}

// SetLinkControlPolicy gets a reference to the given NullableFabricLinkControlPolicyRelationship and assigns it to the LinkControlPolicy field.
func (o *FabricUplinkPcRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship) {
	o.LinkControlPolicy.Set(&v)
}

// SetLinkControlPolicyNil sets the value for LinkControlPolicy to be an explicit nil
func (o *FabricUplinkPcRole) SetLinkControlPolicyNil() {
	o.LinkControlPolicy.Set(nil)
}

// UnsetLinkControlPolicy ensures that no value is present for LinkControlPolicy, not even an explicit nil
func (o *FabricUplinkPcRole) UnsetLinkControlPolicy() {
	o.LinkControlPolicy.Unset()
}

func (o FabricUplinkPcRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricUplinkPcRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortChannelRole, errFabricPortChannelRole := json.Marshal(o.FabricPortChannelRole)
	if errFabricPortChannelRole != nil {
		return map[string]interface{}{}, errFabricPortChannelRole
	}
	errFabricPortChannelRole = json.Unmarshal([]byte(serializedFabricPortChannelRole), &toSerialize)
	if errFabricPortChannelRole != nil {
		return map[string]interface{}{}, errFabricPortChannelRole
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AdminSpeed) {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if !IsNil(o.Fec) {
		toSerialize["Fec"] = o.Fec
	}
	if o.EthNetworkGroupPolicy != nil {
		toSerialize["EthNetworkGroupPolicy"] = o.EthNetworkGroupPolicy
	}
	if o.FlowControlPolicy.IsSet() {
		toSerialize["FlowControlPolicy"] = o.FlowControlPolicy.Get()
	}
	if o.LinkAggregationPolicy.IsSet() {
		toSerialize["LinkAggregationPolicy"] = o.LinkAggregationPolicy.Get()
	}
	if o.LinkControlPolicy.IsSet() {
		toSerialize["LinkControlPolicy"] = o.LinkControlPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricUplinkPcRole) UnmarshalJSON(data []byte) (err error) {
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
	type FabricUplinkPcRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps. * `NegAuto25Gbps` - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Forward error correction configuration for Uplink Port Channel member ports. * `Auto` - Forward error correction option 'Auto'. * `Cl91` - Forward error correction option 'cl91'. * `Cl74` - Forward error correction option 'cl74'.
		Fec *string `json:"Fec,omitempty"`
		// An array of relationships to fabricEthNetworkGroupPolicy resources.
		EthNetworkGroupPolicy []FabricEthNetworkGroupPolicyRelationship       `json:"EthNetworkGroupPolicy,omitempty"`
		FlowControlPolicy     NullableFabricFlowControlPolicyRelationship     `json:"FlowControlPolicy,omitempty"`
		LinkAggregationPolicy NullableFabricLinkAggregationPolicyRelationship `json:"LinkAggregationPolicy,omitempty"`
		LinkControlPolicy     NullableFabricLinkControlPolicyRelationship     `json:"LinkControlPolicy,omitempty"`
	}

	varFabricUplinkPcRoleWithoutEmbeddedStruct := FabricUplinkPcRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varFabricUplinkPcRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricUplinkPcRole := _FabricUplinkPcRole{}
		varFabricUplinkPcRole.ClassId = varFabricUplinkPcRoleWithoutEmbeddedStruct.ClassId
		varFabricUplinkPcRole.ObjectType = varFabricUplinkPcRoleWithoutEmbeddedStruct.ObjectType
		varFabricUplinkPcRole.AdminSpeed = varFabricUplinkPcRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricUplinkPcRole.Fec = varFabricUplinkPcRoleWithoutEmbeddedStruct.Fec
		varFabricUplinkPcRole.EthNetworkGroupPolicy = varFabricUplinkPcRoleWithoutEmbeddedStruct.EthNetworkGroupPolicy
		varFabricUplinkPcRole.FlowControlPolicy = varFabricUplinkPcRoleWithoutEmbeddedStruct.FlowControlPolicy
		varFabricUplinkPcRole.LinkAggregationPolicy = varFabricUplinkPcRoleWithoutEmbeddedStruct.LinkAggregationPolicy
		varFabricUplinkPcRole.LinkControlPolicy = varFabricUplinkPcRoleWithoutEmbeddedStruct.LinkControlPolicy
		*o = FabricUplinkPcRole(varFabricUplinkPcRole)
	} else {
		return err
	}

	varFabricUplinkPcRole := _FabricUplinkPcRole{}

	err = json.Unmarshal(data, &varFabricUplinkPcRole)
	if err == nil {
		o.FabricPortChannelRole = varFabricUplinkPcRole.FabricPortChannelRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "Fec")
		delete(additionalProperties, "EthNetworkGroupPolicy")
		delete(additionalProperties, "FlowControlPolicy")
		delete(additionalProperties, "LinkAggregationPolicy")
		delete(additionalProperties, "LinkControlPolicy")

		// remove fields from embedded structs
		reflectFabricPortChannelRole := reflect.ValueOf(o.FabricPortChannelRole)
		for i := 0; i < reflectFabricPortChannelRole.Type().NumField(); i++ {
			t := reflectFabricPortChannelRole.Type().Field(i)

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

type NullableFabricUplinkPcRole struct {
	value *FabricUplinkPcRole
	isSet bool
}

func (v NullableFabricUplinkPcRole) Get() *FabricUplinkPcRole {
	return v.value
}

func (v *NullableFabricUplinkPcRole) Set(val *FabricUplinkPcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricUplinkPcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricUplinkPcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricUplinkPcRole(val *FabricUplinkPcRole) *NullableFabricUplinkPcRole {
	return &NullableFabricUplinkPcRole{value: val, isSet: true}
}

func (v NullableFabricUplinkPcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricUplinkPcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
