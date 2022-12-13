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
)

// VnicLanConnectivityPolicyInventoryAllOf Definition of the list of properties defined in 'vnic.LanConnectivityPolicyInventory', excluding properties defined in parent classes.
type VnicLanConnectivityPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it.
	AzureQosEnabled *bool `json:"AzureQosEnabled,omitempty"`
	// Allocation Type of iSCSI Qualified Name - Static/Pool/None. * `None` - Type indicates that there is no IQN associated to an interface. * `Static` - Type represents that static IQN is associated to an interface. * `Pool` - Type indicates that IQN value is sourced from an associated pool.
	IqnAllocationType *string `json:"IqnAllocationType,omitempty"`
	// The mode used for placement of vNICs on network adapters. It can either be Auto or Custom. * `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user. * `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.
	PlacementMode *string `json:"PlacementMode,omitempty"`
	// User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain.
	StaticIqnName *string `json:"StaticIqnName,omitempty"`
	// The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight. * `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected. * `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.
	// Deprecated
	TargetPlatform *string `json:"TargetPlatform,omitempty"`
	// An array of relationships to vnicEthIfInventory resources.
	EthIfs               []VnicEthIfInventoryRelationship `json:"EthIfs,omitempty"`
	IqnPool              *IqnpoolPoolRelationship         `json:"IqnPool,omitempty"`
	TargetMo             *MoBaseMoRelationship            `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLanConnectivityPolicyInventoryAllOf VnicLanConnectivityPolicyInventoryAllOf

// NewVnicLanConnectivityPolicyInventoryAllOf instantiates a new VnicLanConnectivityPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLanConnectivityPolicyInventoryAllOf(classId string, objectType string) *VnicLanConnectivityPolicyInventoryAllOf {
	this := VnicLanConnectivityPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicLanConnectivityPolicyInventoryAllOfWithDefaults instantiates a new VnicLanConnectivityPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLanConnectivityPolicyInventoryAllOfWithDefaults() *VnicLanConnectivityPolicyInventoryAllOf {
	this := VnicLanConnectivityPolicyInventoryAllOf{}
	var classId string = "vnic.LanConnectivityPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.LanConnectivityPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAzureQosEnabled returns the AzureQosEnabled field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetAzureQosEnabled() bool {
	if o == nil || o.AzureQosEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AzureQosEnabled
}

// GetAzureQosEnabledOk returns a tuple with the AzureQosEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetAzureQosEnabledOk() (*bool, bool) {
	if o == nil || o.AzureQosEnabled == nil {
		return nil, false
	}
	return o.AzureQosEnabled, true
}

// HasAzureQosEnabled returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasAzureQosEnabled() bool {
	if o != nil && o.AzureQosEnabled != nil {
		return true
	}

	return false
}

// SetAzureQosEnabled gets a reference to the given bool and assigns it to the AzureQosEnabled field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetAzureQosEnabled(v bool) {
	o.AzureQosEnabled = &v
}

// GetIqnAllocationType returns the IqnAllocationType field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetIqnAllocationType() string {
	if o == nil || o.IqnAllocationType == nil {
		var ret string
		return ret
	}
	return *o.IqnAllocationType
}

// GetIqnAllocationTypeOk returns a tuple with the IqnAllocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetIqnAllocationTypeOk() (*string, bool) {
	if o == nil || o.IqnAllocationType == nil {
		return nil, false
	}
	return o.IqnAllocationType, true
}

// HasIqnAllocationType returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasIqnAllocationType() bool {
	if o != nil && o.IqnAllocationType != nil {
		return true
	}

	return false
}

// SetIqnAllocationType gets a reference to the given string and assigns it to the IqnAllocationType field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetIqnAllocationType(v string) {
	o.IqnAllocationType = &v
}

// GetPlacementMode returns the PlacementMode field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetPlacementMode() string {
	if o == nil || o.PlacementMode == nil {
		var ret string
		return ret
	}
	return *o.PlacementMode
}

// GetPlacementModeOk returns a tuple with the PlacementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetPlacementModeOk() (*string, bool) {
	if o == nil || o.PlacementMode == nil {
		return nil, false
	}
	return o.PlacementMode, true
}

// HasPlacementMode returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasPlacementMode() bool {
	if o != nil && o.PlacementMode != nil {
		return true
	}

	return false
}

// SetPlacementMode gets a reference to the given string and assigns it to the PlacementMode field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetPlacementMode(v string) {
	o.PlacementMode = &v
}

// GetStaticIqnName returns the StaticIqnName field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetStaticIqnName() string {
	if o == nil || o.StaticIqnName == nil {
		var ret string
		return ret
	}
	return *o.StaticIqnName
}

// GetStaticIqnNameOk returns a tuple with the StaticIqnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetStaticIqnNameOk() (*string, bool) {
	if o == nil || o.StaticIqnName == nil {
		return nil, false
	}
	return o.StaticIqnName, true
}

// HasStaticIqnName returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasStaticIqnName() bool {
	if o != nil && o.StaticIqnName != nil {
		return true
	}

	return false
}

// SetStaticIqnName gets a reference to the given string and assigns it to the StaticIqnName field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetStaticIqnName(v string) {
	o.StaticIqnName = &v
}

// GetTargetPlatform returns the TargetPlatform field value if set, zero value otherwise.
// Deprecated
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetTargetPlatform() string {
	if o == nil || o.TargetPlatform == nil {
		var ret string
		return ret
	}
	return *o.TargetPlatform
}

// GetTargetPlatformOk returns a tuple with the TargetPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetTargetPlatformOk() (*string, bool) {
	if o == nil || o.TargetPlatform == nil {
		return nil, false
	}
	return o.TargetPlatform, true
}

// HasTargetPlatform returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasTargetPlatform() bool {
	if o != nil && o.TargetPlatform != nil {
		return true
	}

	return false
}

// SetTargetPlatform gets a reference to the given string and assigns it to the TargetPlatform field.
// Deprecated
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetTargetPlatform(v string) {
	o.TargetPlatform = &v
}

// GetEthIfs returns the EthIfs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetEthIfs() []VnicEthIfInventoryRelationship {
	if o == nil {
		var ret []VnicEthIfInventoryRelationship
		return ret
	}
	return o.EthIfs
}

// GetEthIfsOk returns a tuple with the EthIfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetEthIfsOk() ([]VnicEthIfInventoryRelationship, bool) {
	if o == nil || o.EthIfs == nil {
		return nil, false
	}
	return o.EthIfs, true
}

// HasEthIfs returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasEthIfs() bool {
	if o != nil && o.EthIfs != nil {
		return true
	}

	return false
}

// SetEthIfs gets a reference to the given []VnicEthIfInventoryRelationship and assigns it to the EthIfs field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetEthIfs(v []VnicEthIfInventoryRelationship) {
	o.EthIfs = v
}

// GetIqnPool returns the IqnPool field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetIqnPool() IqnpoolPoolRelationship {
	if o == nil || o.IqnPool == nil {
		var ret IqnpoolPoolRelationship
		return ret
	}
	return *o.IqnPool
}

// GetIqnPoolOk returns a tuple with the IqnPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetIqnPoolOk() (*IqnpoolPoolRelationship, bool) {
	if o == nil || o.IqnPool == nil {
		return nil, false
	}
	return o.IqnPool, true
}

// HasIqnPool returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasIqnPool() bool {
	if o != nil && o.IqnPool != nil {
		return true
	}

	return false
}

// SetIqnPool gets a reference to the given IqnpoolPoolRelationship and assigns it to the IqnPool field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetIqnPool(v IqnpoolPoolRelationship) {
	o.IqnPool = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicLanConnectivityPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicLanConnectivityPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicLanConnectivityPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AzureQosEnabled != nil {
		toSerialize["AzureQosEnabled"] = o.AzureQosEnabled
	}
	if o.IqnAllocationType != nil {
		toSerialize["IqnAllocationType"] = o.IqnAllocationType
	}
	if o.PlacementMode != nil {
		toSerialize["PlacementMode"] = o.PlacementMode
	}
	if o.StaticIqnName != nil {
		toSerialize["StaticIqnName"] = o.StaticIqnName
	}
	if o.TargetPlatform != nil {
		toSerialize["TargetPlatform"] = o.TargetPlatform
	}
	if o.EthIfs != nil {
		toSerialize["EthIfs"] = o.EthIfs
	}
	if o.IqnPool != nil {
		toSerialize["IqnPool"] = o.IqnPool
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLanConnectivityPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicLanConnectivityPolicyInventoryAllOf := _VnicLanConnectivityPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varVnicLanConnectivityPolicyInventoryAllOf); err == nil {
		*o = VnicLanConnectivityPolicyInventoryAllOf(varVnicLanConnectivityPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AzureQosEnabled")
		delete(additionalProperties, "IqnAllocationType")
		delete(additionalProperties, "PlacementMode")
		delete(additionalProperties, "StaticIqnName")
		delete(additionalProperties, "TargetPlatform")
		delete(additionalProperties, "EthIfs")
		delete(additionalProperties, "IqnPool")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicLanConnectivityPolicyInventoryAllOf struct {
	value *VnicLanConnectivityPolicyInventoryAllOf
	isSet bool
}

func (v NullableVnicLanConnectivityPolicyInventoryAllOf) Get() *VnicLanConnectivityPolicyInventoryAllOf {
	return v.value
}

func (v *NullableVnicLanConnectivityPolicyInventoryAllOf) Set(val *VnicLanConnectivityPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLanConnectivityPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLanConnectivityPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLanConnectivityPolicyInventoryAllOf(val *VnicLanConnectivityPolicyInventoryAllOf) *NullableVnicLanConnectivityPolicyInventoryAllOf {
	return &NullableVnicLanConnectivityPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableVnicLanConnectivityPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLanConnectivityPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
