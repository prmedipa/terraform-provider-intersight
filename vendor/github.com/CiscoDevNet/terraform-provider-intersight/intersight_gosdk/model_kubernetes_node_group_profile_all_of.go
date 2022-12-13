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

// KubernetesNodeGroupProfileAllOf Definition of the list of properties defined in 'kubernetes.NodeGroupProfile', excluding properties defined in parent classes.
type KubernetesNodeGroupProfileAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Current number of nodes in this node group at any given point in time.
	Currentsize *int64 `json:"Currentsize,omitempty"`
	// Desired number of nodes in this node group, same as minsize initially and is updated by the auto-scaler.
	Desiredsize *int64                      `json:"Desiredsize,omitempty"`
	GpuConfig   []InfraBaseGpuConfiguration `json:"GpuConfig,omitempty"`
	Labels      []KubernetesNodeGroupLabel  `json:"Labels,omitempty"`
	// Maximum number of nodes this node group can scale up to during repair, replacement or upgrade operations.
	Maxsize *int64 `json:"Maxsize,omitempty"`
	// Minimum number of available nodes this node group can scale down to during repair, replacement or upgrade operations.
	Minsize *int64 `json:"Minsize,omitempty"`
	// The node type ControlPlane, Worker or ControlPlaneWorker. * `Worker` - Node will be marked as a worker node. * `ControlPlane` - Node will be marked as a control plane node. * `ControlPlaneWorker` - Node will be both a controle plane and a worker.
	NodeType       *string                                           `json:"NodeType,omitempty"`
	Taints         []KubernetesNodeGroupTaint                        `json:"Taints,omitempty"`
	ClusterProfile *KubernetesClusterProfileRelationship             `json:"ClusterProfile,omitempty"`
	InfraProvider  *KubernetesBaseInfrastructureProviderRelationship `json:"InfraProvider,omitempty"`
	// An array of relationships to ippoolPool resources.
	// Deprecated
	IpPools           []IppoolPoolRelationship             `json:"IpPools,omitempty"`
	KubernetesVersion *KubernetesVersionPolicyRelationship `json:"KubernetesVersion,omitempty"`
	// An array of relationships to kubernetesNodeProfile resources.
	Nodes                []KubernetesNodeProfileRelationship `json:"Nodes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeGroupProfileAllOf KubernetesNodeGroupProfileAllOf

// NewKubernetesNodeGroupProfileAllOf instantiates a new KubernetesNodeGroupProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeGroupProfileAllOf(classId string, objectType string) *KubernetesNodeGroupProfileAllOf {
	this := KubernetesNodeGroupProfileAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var desiredsize int64 = 3
	this.Desiredsize = &desiredsize
	var nodeType string = "Worker"
	this.NodeType = &nodeType
	return &this
}

// NewKubernetesNodeGroupProfileAllOfWithDefaults instantiates a new KubernetesNodeGroupProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeGroupProfileAllOfWithDefaults() *KubernetesNodeGroupProfileAllOf {
	this := KubernetesNodeGroupProfileAllOf{}
	var classId string = "kubernetes.NodeGroupProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.NodeGroupProfile"
	this.ObjectType = objectType
	var desiredsize int64 = 3
	this.Desiredsize = &desiredsize
	var nodeType string = "Worker"
	this.NodeType = &nodeType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeGroupProfileAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeGroupProfileAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeGroupProfileAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeGroupProfileAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCurrentsize returns the Currentsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetCurrentsize() int64 {
	if o == nil || o.Currentsize == nil {
		var ret int64
		return ret
	}
	return *o.Currentsize
}

// GetCurrentsizeOk returns a tuple with the Currentsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetCurrentsizeOk() (*int64, bool) {
	if o == nil || o.Currentsize == nil {
		return nil, false
	}
	return o.Currentsize, true
}

// HasCurrentsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasCurrentsize() bool {
	if o != nil && o.Currentsize != nil {
		return true
	}

	return false
}

// SetCurrentsize gets a reference to the given int64 and assigns it to the Currentsize field.
func (o *KubernetesNodeGroupProfileAllOf) SetCurrentsize(v int64) {
	o.Currentsize = &v
}

// GetDesiredsize returns the Desiredsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetDesiredsize() int64 {
	if o == nil || o.Desiredsize == nil {
		var ret int64
		return ret
	}
	return *o.Desiredsize
}

// GetDesiredsizeOk returns a tuple with the Desiredsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetDesiredsizeOk() (*int64, bool) {
	if o == nil || o.Desiredsize == nil {
		return nil, false
	}
	return o.Desiredsize, true
}

// HasDesiredsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasDesiredsize() bool {
	if o != nil && o.Desiredsize != nil {
		return true
	}

	return false
}

// SetDesiredsize gets a reference to the given int64 and assigns it to the Desiredsize field.
func (o *KubernetesNodeGroupProfileAllOf) SetDesiredsize(v int64) {
	o.Desiredsize = &v
}

// GetGpuConfig returns the GpuConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfileAllOf) GetGpuConfig() []InfraBaseGpuConfiguration {
	if o == nil {
		var ret []InfraBaseGpuConfiguration
		return ret
	}
	return o.GpuConfig
}

// GetGpuConfigOk returns a tuple with the GpuConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfileAllOf) GetGpuConfigOk() ([]InfraBaseGpuConfiguration, bool) {
	if o == nil || o.GpuConfig == nil {
		return nil, false
	}
	return o.GpuConfig, true
}

// HasGpuConfig returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasGpuConfig() bool {
	if o != nil && o.GpuConfig != nil {
		return true
	}

	return false
}

// SetGpuConfig gets a reference to the given []InfraBaseGpuConfiguration and assigns it to the GpuConfig field.
func (o *KubernetesNodeGroupProfileAllOf) SetGpuConfig(v []InfraBaseGpuConfiguration) {
	o.GpuConfig = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfileAllOf) GetLabels() []KubernetesNodeGroupLabel {
	if o == nil {
		var ret []KubernetesNodeGroupLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfileAllOf) GetLabelsOk() ([]KubernetesNodeGroupLabel, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []KubernetesNodeGroupLabel and assigns it to the Labels field.
func (o *KubernetesNodeGroupProfileAllOf) SetLabels(v []KubernetesNodeGroupLabel) {
	o.Labels = v
}

// GetMaxsize returns the Maxsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetMaxsize() int64 {
	if o == nil || o.Maxsize == nil {
		var ret int64
		return ret
	}
	return *o.Maxsize
}

// GetMaxsizeOk returns a tuple with the Maxsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetMaxsizeOk() (*int64, bool) {
	if o == nil || o.Maxsize == nil {
		return nil, false
	}
	return o.Maxsize, true
}

// HasMaxsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasMaxsize() bool {
	if o != nil && o.Maxsize != nil {
		return true
	}

	return false
}

// SetMaxsize gets a reference to the given int64 and assigns it to the Maxsize field.
func (o *KubernetesNodeGroupProfileAllOf) SetMaxsize(v int64) {
	o.Maxsize = &v
}

// GetMinsize returns the Minsize field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetMinsize() int64 {
	if o == nil || o.Minsize == nil {
		var ret int64
		return ret
	}
	return *o.Minsize
}

// GetMinsizeOk returns a tuple with the Minsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetMinsizeOk() (*int64, bool) {
	if o == nil || o.Minsize == nil {
		return nil, false
	}
	return o.Minsize, true
}

// HasMinsize returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasMinsize() bool {
	if o != nil && o.Minsize != nil {
		return true
	}

	return false
}

// SetMinsize gets a reference to the given int64 and assigns it to the Minsize field.
func (o *KubernetesNodeGroupProfileAllOf) SetMinsize(v int64) {
	o.Minsize = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasNodeType() bool {
	if o != nil && o.NodeType != nil {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *KubernetesNodeGroupProfileAllOf) SetNodeType(v string) {
	o.NodeType = &v
}

// GetTaints returns the Taints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfileAllOf) GetTaints() []KubernetesNodeGroupTaint {
	if o == nil {
		var ret []KubernetesNodeGroupTaint
		return ret
	}
	return o.Taints
}

// GetTaintsOk returns a tuple with the Taints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfileAllOf) GetTaintsOk() ([]KubernetesNodeGroupTaint, bool) {
	if o == nil || o.Taints == nil {
		return nil, false
	}
	return o.Taints, true
}

// HasTaints returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasTaints() bool {
	if o != nil && o.Taints != nil {
		return true
	}

	return false
}

// SetTaints gets a reference to the given []KubernetesNodeGroupTaint and assigns it to the Taints field.
func (o *KubernetesNodeGroupProfileAllOf) SetTaints(v []KubernetesNodeGroupTaint) {
	o.Taints = v
}

// GetClusterProfile returns the ClusterProfile field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetClusterProfile() KubernetesClusterProfileRelationship {
	if o == nil || o.ClusterProfile == nil {
		var ret KubernetesClusterProfileRelationship
		return ret
	}
	return *o.ClusterProfile
}

// GetClusterProfileOk returns a tuple with the ClusterProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetClusterProfileOk() (*KubernetesClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfile == nil {
		return nil, false
	}
	return o.ClusterProfile, true
}

// HasClusterProfile returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasClusterProfile() bool {
	if o != nil && o.ClusterProfile != nil {
		return true
	}

	return false
}

// SetClusterProfile gets a reference to the given KubernetesClusterProfileRelationship and assigns it to the ClusterProfile field.
func (o *KubernetesNodeGroupProfileAllOf) SetClusterProfile(v KubernetesClusterProfileRelationship) {
	o.ClusterProfile = &v
}

// GetInfraProvider returns the InfraProvider field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetInfraProvider() KubernetesBaseInfrastructureProviderRelationship {
	if o == nil || o.InfraProvider == nil {
		var ret KubernetesBaseInfrastructureProviderRelationship
		return ret
	}
	return *o.InfraProvider
}

// GetInfraProviderOk returns a tuple with the InfraProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetInfraProviderOk() (*KubernetesBaseInfrastructureProviderRelationship, bool) {
	if o == nil || o.InfraProvider == nil {
		return nil, false
	}
	return o.InfraProvider, true
}

// HasInfraProvider returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasInfraProvider() bool {
	if o != nil && o.InfraProvider != nil {
		return true
	}

	return false
}

// SetInfraProvider gets a reference to the given KubernetesBaseInfrastructureProviderRelationship and assigns it to the InfraProvider field.
func (o *KubernetesNodeGroupProfileAllOf) SetInfraProvider(v KubernetesBaseInfrastructureProviderRelationship) {
	o.InfraProvider = &v
}

// GetIpPools returns the IpPools field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *KubernetesNodeGroupProfileAllOf) GetIpPools() []IppoolPoolRelationship {
	if o == nil {
		var ret []IppoolPoolRelationship
		return ret
	}
	return o.IpPools
}

// GetIpPoolsOk returns a tuple with the IpPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *KubernetesNodeGroupProfileAllOf) GetIpPoolsOk() ([]IppoolPoolRelationship, bool) {
	if o == nil || o.IpPools == nil {
		return nil, false
	}
	return o.IpPools, true
}

// HasIpPools returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasIpPools() bool {
	if o != nil && o.IpPools != nil {
		return true
	}

	return false
}

// SetIpPools gets a reference to the given []IppoolPoolRelationship and assigns it to the IpPools field.
// Deprecated
func (o *KubernetesNodeGroupProfileAllOf) SetIpPools(v []IppoolPoolRelationship) {
	o.IpPools = v
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *KubernetesNodeGroupProfileAllOf) GetKubernetesVersion() KubernetesVersionPolicyRelationship {
	if o == nil || o.KubernetesVersion == nil {
		var ret KubernetesVersionPolicyRelationship
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeGroupProfileAllOf) GetKubernetesVersionOk() (*KubernetesVersionPolicyRelationship, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion != nil {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given KubernetesVersionPolicyRelationship and assigns it to the KubernetesVersion field.
func (o *KubernetesNodeGroupProfileAllOf) SetKubernetesVersion(v KubernetesVersionPolicyRelationship) {
	o.KubernetesVersion = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesNodeGroupProfileAllOf) GetNodes() []KubernetesNodeProfileRelationship {
	if o == nil {
		var ret []KubernetesNodeProfileRelationship
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesNodeGroupProfileAllOf) GetNodesOk() ([]KubernetesNodeProfileRelationship, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *KubernetesNodeGroupProfileAllOf) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []KubernetesNodeProfileRelationship and assigns it to the Nodes field.
func (o *KubernetesNodeGroupProfileAllOf) SetNodes(v []KubernetesNodeProfileRelationship) {
	o.Nodes = v
}

func (o KubernetesNodeGroupProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Currentsize != nil {
		toSerialize["Currentsize"] = o.Currentsize
	}
	if o.Desiredsize != nil {
		toSerialize["Desiredsize"] = o.Desiredsize
	}
	if o.GpuConfig != nil {
		toSerialize["GpuConfig"] = o.GpuConfig
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.Maxsize != nil {
		toSerialize["Maxsize"] = o.Maxsize
	}
	if o.Minsize != nil {
		toSerialize["Minsize"] = o.Minsize
	}
	if o.NodeType != nil {
		toSerialize["NodeType"] = o.NodeType
	}
	if o.Taints != nil {
		toSerialize["Taints"] = o.Taints
	}
	if o.ClusterProfile != nil {
		toSerialize["ClusterProfile"] = o.ClusterProfile
	}
	if o.InfraProvider != nil {
		toSerialize["InfraProvider"] = o.InfraProvider
	}
	if o.IpPools != nil {
		toSerialize["IpPools"] = o.IpPools
	}
	if o.KubernetesVersion != nil {
		toSerialize["KubernetesVersion"] = o.KubernetesVersion
	}
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeGroupProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesNodeGroupProfileAllOf := _KubernetesNodeGroupProfileAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesNodeGroupProfileAllOf); err == nil {
		*o = KubernetesNodeGroupProfileAllOf(varKubernetesNodeGroupProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Currentsize")
		delete(additionalProperties, "Desiredsize")
		delete(additionalProperties, "GpuConfig")
		delete(additionalProperties, "Labels")
		delete(additionalProperties, "Maxsize")
		delete(additionalProperties, "Minsize")
		delete(additionalProperties, "NodeType")
		delete(additionalProperties, "Taints")
		delete(additionalProperties, "ClusterProfile")
		delete(additionalProperties, "InfraProvider")
		delete(additionalProperties, "IpPools")
		delete(additionalProperties, "KubernetesVersion")
		delete(additionalProperties, "Nodes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesNodeGroupProfileAllOf struct {
	value *KubernetesNodeGroupProfileAllOf
	isSet bool
}

func (v NullableKubernetesNodeGroupProfileAllOf) Get() *KubernetesNodeGroupProfileAllOf {
	return v.value
}

func (v *NullableKubernetesNodeGroupProfileAllOf) Set(val *KubernetesNodeGroupProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeGroupProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeGroupProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeGroupProfileAllOf(val *KubernetesNodeGroupProfileAllOf) *NullableKubernetesNodeGroupProfileAllOf {
	return &NullableKubernetesNodeGroupProfileAllOf{value: val, isSet: true}
}

func (v NullableKubernetesNodeGroupProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeGroupProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
