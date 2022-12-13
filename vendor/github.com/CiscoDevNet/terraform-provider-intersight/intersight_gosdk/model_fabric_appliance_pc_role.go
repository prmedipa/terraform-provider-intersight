/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// FabricAppliancePcRole Configuration object sent by user to create an appliance port channel.
type FabricAppliancePcRole struct {
	FabricPortChannelRole
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin configured speed for the port channel. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
	AdminSpeed *string `json:"AdminSpeed,omitempty"`
	// Port mode to be set on the appliance port-channel. * `trunk` - Trunk Mode Switch Port Type. * `access` - Access Mode Switch Port Type.
	Mode *string `json:"Mode,omitempty"`
	// The 'name' of the System QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority                *string                                    `json:"Priority,omitempty"`
	EthNetworkControlPolicy *FabricEthNetworkControlPolicyRelationship `json:"EthNetworkControlPolicy,omitempty"`
	EthNetworkGroupPolicy   *FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
	LinkAggregationPolicy   *FabricLinkAggregationPolicyRelationship   `json:"LinkAggregationPolicy,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _FabricAppliancePcRole FabricAppliancePcRole

// NewFabricAppliancePcRole instantiates a new FabricAppliancePcRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricAppliancePcRole(classId string, objectType string) *FabricAppliancePcRole {
	this := FabricAppliancePcRole{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var mode string = "trunk"
	this.Mode = &mode
	var priority string = "Best Effort"
	this.Priority = &priority
	return &this
}

// NewFabricAppliancePcRoleWithDefaults instantiates a new FabricAppliancePcRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricAppliancePcRoleWithDefaults() *FabricAppliancePcRole {
	this := FabricAppliancePcRole{}
	var classId string = "fabric.AppliancePcRole"
	this.ClassId = classId
	var objectType string = "fabric.AppliancePcRole"
	this.ObjectType = objectType
	var adminSpeed string = "Auto"
	this.AdminSpeed = &adminSpeed
	var mode string = "trunk"
	this.Mode = &mode
	var priority string = "Best Effort"
	this.Priority = &priority
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricAppliancePcRole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricAppliancePcRole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricAppliancePcRole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricAppliancePcRole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminSpeed returns the AdminSpeed field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetAdminSpeed() string {
	if o == nil || o.AdminSpeed == nil {
		var ret string
		return ret
	}
	return *o.AdminSpeed
}

// GetAdminSpeedOk returns a tuple with the AdminSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetAdminSpeedOk() (*string, bool) {
	if o == nil || o.AdminSpeed == nil {
		return nil, false
	}
	return o.AdminSpeed, true
}

// HasAdminSpeed returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasAdminSpeed() bool {
	if o != nil && o.AdminSpeed != nil {
		return true
	}

	return false
}

// SetAdminSpeed gets a reference to the given string and assigns it to the AdminSpeed field.
func (o *FabricAppliancePcRole) SetAdminSpeed(v string) {
	o.AdminSpeed = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *FabricAppliancePcRole) SetMode(v string) {
	o.Mode = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *FabricAppliancePcRole) SetPriority(v string) {
	o.Priority = &v
}

// GetEthNetworkControlPolicy returns the EthNetworkControlPolicy field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetEthNetworkControlPolicy() FabricEthNetworkControlPolicyRelationship {
	if o == nil || o.EthNetworkControlPolicy == nil {
		var ret FabricEthNetworkControlPolicyRelationship
		return ret
	}
	return *o.EthNetworkControlPolicy
}

// GetEthNetworkControlPolicyOk returns a tuple with the EthNetworkControlPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetEthNetworkControlPolicyOk() (*FabricEthNetworkControlPolicyRelationship, bool) {
	if o == nil || o.EthNetworkControlPolicy == nil {
		return nil, false
	}
	return o.EthNetworkControlPolicy, true
}

// HasEthNetworkControlPolicy returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasEthNetworkControlPolicy() bool {
	if o != nil && o.EthNetworkControlPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkControlPolicy gets a reference to the given FabricEthNetworkControlPolicyRelationship and assigns it to the EthNetworkControlPolicy field.
func (o *FabricAppliancePcRole) SetEthNetworkControlPolicy(v FabricEthNetworkControlPolicyRelationship) {
	o.EthNetworkControlPolicy = &v
}

// GetEthNetworkGroupPolicy returns the EthNetworkGroupPolicy field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetEthNetworkGroupPolicy() FabricEthNetworkGroupPolicyRelationship {
	if o == nil || o.EthNetworkGroupPolicy == nil {
		var ret FabricEthNetworkGroupPolicyRelationship
		return ret
	}
	return *o.EthNetworkGroupPolicy
}

// GetEthNetworkGroupPolicyOk returns a tuple with the EthNetworkGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetEthNetworkGroupPolicyOk() (*FabricEthNetworkGroupPolicyRelationship, bool) {
	if o == nil || o.EthNetworkGroupPolicy == nil {
		return nil, false
	}
	return o.EthNetworkGroupPolicy, true
}

// HasEthNetworkGroupPolicy returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasEthNetworkGroupPolicy() bool {
	if o != nil && o.EthNetworkGroupPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkGroupPolicy gets a reference to the given FabricEthNetworkGroupPolicyRelationship and assigns it to the EthNetworkGroupPolicy field.
func (o *FabricAppliancePcRole) SetEthNetworkGroupPolicy(v FabricEthNetworkGroupPolicyRelationship) {
	o.EthNetworkGroupPolicy = &v
}

// GetLinkAggregationPolicy returns the LinkAggregationPolicy field value if set, zero value otherwise.
func (o *FabricAppliancePcRole) GetLinkAggregationPolicy() FabricLinkAggregationPolicyRelationship {
	if o == nil || o.LinkAggregationPolicy == nil {
		var ret FabricLinkAggregationPolicyRelationship
		return ret
	}
	return *o.LinkAggregationPolicy
}

// GetLinkAggregationPolicyOk returns a tuple with the LinkAggregationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricAppliancePcRole) GetLinkAggregationPolicyOk() (*FabricLinkAggregationPolicyRelationship, bool) {
	if o == nil || o.LinkAggregationPolicy == nil {
		return nil, false
	}
	return o.LinkAggregationPolicy, true
}

// HasLinkAggregationPolicy returns a boolean if a field has been set.
func (o *FabricAppliancePcRole) HasLinkAggregationPolicy() bool {
	if o != nil && o.LinkAggregationPolicy != nil {
		return true
	}

	return false
}

// SetLinkAggregationPolicy gets a reference to the given FabricLinkAggregationPolicyRelationship and assigns it to the LinkAggregationPolicy field.
func (o *FabricAppliancePcRole) SetLinkAggregationPolicy(v FabricLinkAggregationPolicyRelationship) {
	o.LinkAggregationPolicy = &v
}

func (o FabricAppliancePcRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFabricPortChannelRole, errFabricPortChannelRole := json.Marshal(o.FabricPortChannelRole)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	errFabricPortChannelRole = json.Unmarshal([]byte(serializedFabricPortChannelRole), &toSerialize)
	if errFabricPortChannelRole != nil {
		return []byte{}, errFabricPortChannelRole
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminSpeed != nil {
		toSerialize["AdminSpeed"] = o.AdminSpeed
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.EthNetworkControlPolicy != nil {
		toSerialize["EthNetworkControlPolicy"] = o.EthNetworkControlPolicy
	}
	if o.EthNetworkGroupPolicy != nil {
		toSerialize["EthNetworkGroupPolicy"] = o.EthNetworkGroupPolicy
	}
	if o.LinkAggregationPolicy != nil {
		toSerialize["LinkAggregationPolicy"] = o.LinkAggregationPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricAppliancePcRole) UnmarshalJSON(bytes []byte) (err error) {
	type FabricAppliancePcRoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin configured speed for the port channel. * `Auto` - Admin configurable speed AUTO ( default ). * `1Gbps` - Admin configurable speed 1Gbps. * `10Gbps` - Admin configurable speed 10Gbps. * `25Gbps` - Admin configurable speed 25Gbps. * `40Gbps` - Admin configurable speed 40Gbps. * `100Gbps` - Admin configurable speed 100Gbps.
		AdminSpeed *string `json:"AdminSpeed,omitempty"`
		// Port mode to be set on the appliance port-channel. * `trunk` - Trunk Mode Switch Port Type. * `access` - Access Mode Switch Port Type.
		Mode *string `json:"Mode,omitempty"`
		// The 'name' of the System QoS Class. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority                *string                                    `json:"Priority,omitempty"`
		EthNetworkControlPolicy *FabricEthNetworkControlPolicyRelationship `json:"EthNetworkControlPolicy,omitempty"`
		EthNetworkGroupPolicy   *FabricEthNetworkGroupPolicyRelationship   `json:"EthNetworkGroupPolicy,omitempty"`
		LinkAggregationPolicy   *FabricLinkAggregationPolicyRelationship   `json:"LinkAggregationPolicy,omitempty"`
	}

	varFabricAppliancePcRoleWithoutEmbeddedStruct := FabricAppliancePcRoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricAppliancePcRoleWithoutEmbeddedStruct)
	if err == nil {
		varFabricAppliancePcRole := _FabricAppliancePcRole{}
		varFabricAppliancePcRole.ClassId = varFabricAppliancePcRoleWithoutEmbeddedStruct.ClassId
		varFabricAppliancePcRole.ObjectType = varFabricAppliancePcRoleWithoutEmbeddedStruct.ObjectType
		varFabricAppliancePcRole.AdminSpeed = varFabricAppliancePcRoleWithoutEmbeddedStruct.AdminSpeed
		varFabricAppliancePcRole.Mode = varFabricAppliancePcRoleWithoutEmbeddedStruct.Mode
		varFabricAppliancePcRole.Priority = varFabricAppliancePcRoleWithoutEmbeddedStruct.Priority
		varFabricAppliancePcRole.EthNetworkControlPolicy = varFabricAppliancePcRoleWithoutEmbeddedStruct.EthNetworkControlPolicy
		varFabricAppliancePcRole.EthNetworkGroupPolicy = varFabricAppliancePcRoleWithoutEmbeddedStruct.EthNetworkGroupPolicy
		varFabricAppliancePcRole.LinkAggregationPolicy = varFabricAppliancePcRoleWithoutEmbeddedStruct.LinkAggregationPolicy
		*o = FabricAppliancePcRole(varFabricAppliancePcRole)
	} else {
		return err
	}

	varFabricAppliancePcRole := _FabricAppliancePcRole{}

	err = json.Unmarshal(bytes, &varFabricAppliancePcRole)
	if err == nil {
		o.FabricPortChannelRole = varFabricAppliancePcRole.FabricPortChannelRole
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminSpeed")
		delete(additionalProperties, "Mode")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "EthNetworkControlPolicy")
		delete(additionalProperties, "EthNetworkGroupPolicy")
		delete(additionalProperties, "LinkAggregationPolicy")

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

type NullableFabricAppliancePcRole struct {
	value *FabricAppliancePcRole
	isSet bool
}

func (v NullableFabricAppliancePcRole) Get() *FabricAppliancePcRole {
	return v.value
}

func (v *NullableFabricAppliancePcRole) Set(val *FabricAppliancePcRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricAppliancePcRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricAppliancePcRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricAppliancePcRole(val *FabricAppliancePcRole) *NullableFabricAppliancePcRole {
	return &NullableFabricAppliancePcRole{value: val, isSet: true}
}

func (v NullableFabricAppliancePcRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricAppliancePcRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
