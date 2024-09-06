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

// checks if the ApplianceClusterReplaceNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplianceClusterReplaceNode{}

// ApplianceClusterReplaceNode ClusterReplaceNode is a singleton that tracks the Intersight Appliance's process for replacing a cluster node.
type ApplianceClusterReplaceNode struct {
	ApplianceClusterInstallBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Hostname of the node being replaced.
	Hostname *string `json:"Hostname,omitempty"`
	// Node id of the node being replaced.
	NodeId *int64 `json:"NodeId,omitempty"`
	// If the node being replaced has a different IP.
	NodeIpChanged        *bool `json:"NodeIpChanged,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceClusterReplaceNode ApplianceClusterReplaceNode

// NewApplianceClusterReplaceNode instantiates a new ApplianceClusterReplaceNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceClusterReplaceNode(classId string, objectType string) *ApplianceClusterReplaceNode {
	this := ApplianceClusterReplaceNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewApplianceClusterReplaceNodeWithDefaults instantiates a new ApplianceClusterReplaceNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceClusterReplaceNodeWithDefaults() *ApplianceClusterReplaceNode {
	this := ApplianceClusterReplaceNode{}
	var classId string = "appliance.ClusterReplaceNode"
	this.ClassId = classId
	var objectType string = "appliance.ClusterReplaceNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceClusterReplaceNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterReplaceNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceClusterReplaceNode) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "appliance.ClusterReplaceNode" of the ClassId field.
func (o *ApplianceClusterReplaceNode) GetDefaultClassId() interface{} {
	return "appliance.ClusterReplaceNode"
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceClusterReplaceNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceClusterReplaceNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceClusterReplaceNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "appliance.ClusterReplaceNode" of the ObjectType field.
func (o *ApplianceClusterReplaceNode) GetDefaultObjectType() interface{} {
	return "appliance.ClusterReplaceNode"
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceClusterReplaceNode) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterReplaceNode) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceClusterReplaceNode) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceClusterReplaceNode) SetHostname(v string) {
	o.Hostname = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ApplianceClusterReplaceNode) GetNodeId() int64 {
	if o == nil || IsNil(o.NodeId) {
		var ret int64
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterReplaceNode) GetNodeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ApplianceClusterReplaceNode) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given int64 and assigns it to the NodeId field.
func (o *ApplianceClusterReplaceNode) SetNodeId(v int64) {
	o.NodeId = &v
}

// GetNodeIpChanged returns the NodeIpChanged field value if set, zero value otherwise.
func (o *ApplianceClusterReplaceNode) GetNodeIpChanged() bool {
	if o == nil || IsNil(o.NodeIpChanged) {
		var ret bool
		return ret
	}
	return *o.NodeIpChanged
}

// GetNodeIpChangedOk returns a tuple with the NodeIpChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceClusterReplaceNode) GetNodeIpChangedOk() (*bool, bool) {
	if o == nil || IsNil(o.NodeIpChanged) {
		return nil, false
	}
	return o.NodeIpChanged, true
}

// HasNodeIpChanged returns a boolean if a field has been set.
func (o *ApplianceClusterReplaceNode) HasNodeIpChanged() bool {
	if o != nil && !IsNil(o.NodeIpChanged) {
		return true
	}

	return false
}

// SetNodeIpChanged gets a reference to the given bool and assigns it to the NodeIpChanged field.
func (o *ApplianceClusterReplaceNode) SetNodeIpChanged(v bool) {
	o.NodeIpChanged = &v
}

func (o ApplianceClusterReplaceNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplianceClusterReplaceNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApplianceClusterInstallBase, errApplianceClusterInstallBase := json.Marshal(o.ApplianceClusterInstallBase)
	if errApplianceClusterInstallBase != nil {
		return map[string]interface{}{}, errApplianceClusterInstallBase
	}
	errApplianceClusterInstallBase = json.Unmarshal([]byte(serializedApplianceClusterInstallBase), &toSerialize)
	if errApplianceClusterInstallBase != nil {
		return map[string]interface{}{}, errApplianceClusterInstallBase
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Hostname) {
		toSerialize["Hostname"] = o.Hostname
	}
	if !IsNil(o.NodeId) {
		toSerialize["NodeId"] = o.NodeId
	}
	if !IsNil(o.NodeIpChanged) {
		toSerialize["NodeIpChanged"] = o.NodeIpChanged
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplianceClusterReplaceNode) UnmarshalJSON(data []byte) (err error) {
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
	type ApplianceClusterReplaceNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Hostname of the node being replaced.
		Hostname *string `json:"Hostname,omitempty"`
		// Node id of the node being replaced.
		NodeId *int64 `json:"NodeId,omitempty"`
		// If the node being replaced has a different IP.
		NodeIpChanged *bool `json:"NodeIpChanged,omitempty"`
	}

	varApplianceClusterReplaceNodeWithoutEmbeddedStruct := ApplianceClusterReplaceNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varApplianceClusterReplaceNodeWithoutEmbeddedStruct)
	if err == nil {
		varApplianceClusterReplaceNode := _ApplianceClusterReplaceNode{}
		varApplianceClusterReplaceNode.ClassId = varApplianceClusterReplaceNodeWithoutEmbeddedStruct.ClassId
		varApplianceClusterReplaceNode.ObjectType = varApplianceClusterReplaceNodeWithoutEmbeddedStruct.ObjectType
		varApplianceClusterReplaceNode.Hostname = varApplianceClusterReplaceNodeWithoutEmbeddedStruct.Hostname
		varApplianceClusterReplaceNode.NodeId = varApplianceClusterReplaceNodeWithoutEmbeddedStruct.NodeId
		varApplianceClusterReplaceNode.NodeIpChanged = varApplianceClusterReplaceNodeWithoutEmbeddedStruct.NodeIpChanged
		*o = ApplianceClusterReplaceNode(varApplianceClusterReplaceNode)
	} else {
		return err
	}

	varApplianceClusterReplaceNode := _ApplianceClusterReplaceNode{}

	err = json.Unmarshal(data, &varApplianceClusterReplaceNode)
	if err == nil {
		o.ApplianceClusterInstallBase = varApplianceClusterReplaceNode.ApplianceClusterInstallBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "NodeId")
		delete(additionalProperties, "NodeIpChanged")

		// remove fields from embedded structs
		reflectApplianceClusterInstallBase := reflect.ValueOf(o.ApplianceClusterInstallBase)
		for i := 0; i < reflectApplianceClusterInstallBase.Type().NumField(); i++ {
			t := reflectApplianceClusterInstallBase.Type().Field(i)

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

type NullableApplianceClusterReplaceNode struct {
	value *ApplianceClusterReplaceNode
	isSet bool
}

func (v NullableApplianceClusterReplaceNode) Get() *ApplianceClusterReplaceNode {
	return v.value
}

func (v *NullableApplianceClusterReplaceNode) Set(val *ApplianceClusterReplaceNode) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceClusterReplaceNode) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceClusterReplaceNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceClusterReplaceNode(val *ApplianceClusterReplaceNode) *NullableApplianceClusterReplaceNode {
	return &NullableApplianceClusterReplaceNode{value: val, isSet: true}
}

func (v NullableApplianceClusterReplaceNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceClusterReplaceNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
