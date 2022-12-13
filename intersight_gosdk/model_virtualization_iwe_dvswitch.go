/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9661
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationIweDvswitch A Intersight Workload Engine cluster wise distributed vSwitch entity.
type VirtualizationIweDvswitch struct {
	VirtualizationBaseDvswitch
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the dvUplink referenced by this dvswitch.
	DvUplink *string `json:"DvUplink,omitempty"`
	// The last host that update this object.
	LastHostname *string                               `json:"LastHostname,omitempty"`
	Cluster      *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
	// An array of relationships to virtualizationIweHost resources.
	MemberHosts []VirtualizationIweHostRelationship `json:"MemberHosts,omitempty"`
	// An array of relationships to virtualizationIweHostVswitch resources.
	MemberVswitches      []VirtualizationIweHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweDvswitch VirtualizationIweDvswitch

// NewVirtualizationIweDvswitch instantiates a new VirtualizationIweDvswitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweDvswitch(classId string, objectType string) *VirtualizationIweDvswitch {
	this := VirtualizationIweDvswitch{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIweDvswitchWithDefaults instantiates a new VirtualizationIweDvswitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweDvswitchWithDefaults() *VirtualizationIweDvswitch {
	this := VirtualizationIweDvswitch{}
	var classId string = "virtualization.IweDvswitch"
	this.ClassId = classId
	var objectType string = "virtualization.IweDvswitch"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweDvswitch) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitch) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweDvswitch) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweDvswitch) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitch) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweDvswitch) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDvUplink returns the DvUplink field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitch) GetDvUplink() string {
	if o == nil || o.DvUplink == nil {
		var ret string
		return ret
	}
	return *o.DvUplink
}

// GetDvUplinkOk returns a tuple with the DvUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitch) GetDvUplinkOk() (*string, bool) {
	if o == nil || o.DvUplink == nil {
		return nil, false
	}
	return o.DvUplink, true
}

// HasDvUplink returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitch) HasDvUplink() bool {
	if o != nil && o.DvUplink != nil {
		return true
	}

	return false
}

// SetDvUplink gets a reference to the given string and assigns it to the DvUplink field.
func (o *VirtualizationIweDvswitch) SetDvUplink(v string) {
	o.DvUplink = &v
}

// GetLastHostname returns the LastHostname field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitch) GetLastHostname() string {
	if o == nil || o.LastHostname == nil {
		var ret string
		return ret
	}
	return *o.LastHostname
}

// GetLastHostnameOk returns a tuple with the LastHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitch) GetLastHostnameOk() (*string, bool) {
	if o == nil || o.LastHostname == nil {
		return nil, false
	}
	return o.LastHostname, true
}

// HasLastHostname returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitch) HasLastHostname() bool {
	if o != nil && o.LastHostname != nil {
		return true
	}

	return false
}

// SetLastHostname gets a reference to the given string and assigns it to the LastHostname field.
func (o *VirtualizationIweDvswitch) SetLastHostname(v string) {
	o.LastHostname = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VirtualizationIweDvswitch) GetCluster() VirtualizationIweClusterRelationship {
	if o == nil || o.Cluster == nil {
		var ret VirtualizationIweClusterRelationship
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDvswitch) GetClusterOk() (*VirtualizationIweClusterRelationship, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitch) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given VirtualizationIweClusterRelationship and assigns it to the Cluster field.
func (o *VirtualizationIweDvswitch) SetCluster(v VirtualizationIweClusterRelationship) {
	o.Cluster = &v
}

// GetMemberHosts returns the MemberHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvswitch) GetMemberHosts() []VirtualizationIweHostRelationship {
	if o == nil {
		var ret []VirtualizationIweHostRelationship
		return ret
	}
	return o.MemberHosts
}

// GetMemberHostsOk returns a tuple with the MemberHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvswitch) GetMemberHostsOk() ([]VirtualizationIweHostRelationship, bool) {
	if o == nil || o.MemberHosts == nil {
		return nil, false
	}
	return o.MemberHosts, true
}

// HasMemberHosts returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitch) HasMemberHosts() bool {
	if o != nil && o.MemberHosts != nil {
		return true
	}

	return false
}

// SetMemberHosts gets a reference to the given []VirtualizationIweHostRelationship and assigns it to the MemberHosts field.
func (o *VirtualizationIweDvswitch) SetMemberHosts(v []VirtualizationIweHostRelationship) {
	o.MemberHosts = v
}

// GetMemberVswitches returns the MemberVswitches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationIweDvswitch) GetMemberVswitches() []VirtualizationIweHostVswitchRelationship {
	if o == nil {
		var ret []VirtualizationIweHostVswitchRelationship
		return ret
	}
	return o.MemberVswitches
}

// GetMemberVswitchesOk returns a tuple with the MemberVswitches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationIweDvswitch) GetMemberVswitchesOk() ([]VirtualizationIweHostVswitchRelationship, bool) {
	if o == nil || o.MemberVswitches == nil {
		return nil, false
	}
	return o.MemberVswitches, true
}

// HasMemberVswitches returns a boolean if a field has been set.
func (o *VirtualizationIweDvswitch) HasMemberVswitches() bool {
	if o != nil && o.MemberVswitches != nil {
		return true
	}

	return false
}

// SetMemberVswitches gets a reference to the given []VirtualizationIweHostVswitchRelationship and assigns it to the MemberVswitches field.
func (o *VirtualizationIweDvswitch) SetMemberVswitches(v []VirtualizationIweHostVswitchRelationship) {
	o.MemberVswitches = v
}

func (o VirtualizationIweDvswitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDvswitch, errVirtualizationBaseDvswitch := json.Marshal(o.VirtualizationBaseDvswitch)
	if errVirtualizationBaseDvswitch != nil {
		return []byte{}, errVirtualizationBaseDvswitch
	}
	errVirtualizationBaseDvswitch = json.Unmarshal([]byte(serializedVirtualizationBaseDvswitch), &toSerialize)
	if errVirtualizationBaseDvswitch != nil {
		return []byte{}, errVirtualizationBaseDvswitch
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DvUplink != nil {
		toSerialize["DvUplink"] = o.DvUplink
	}
	if o.LastHostname != nil {
		toSerialize["LastHostname"] = o.LastHostname
	}
	if o.Cluster != nil {
		toSerialize["Cluster"] = o.Cluster
	}
	if o.MemberHosts != nil {
		toSerialize["MemberHosts"] = o.MemberHosts
	}
	if o.MemberVswitches != nil {
		toSerialize["MemberVswitches"] = o.MemberVswitches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweDvswitch) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIweDvswitchWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the dvUplink referenced by this dvswitch.
		DvUplink *string `json:"DvUplink,omitempty"`
		// The last host that update this object.
		LastHostname *string                               `json:"LastHostname,omitempty"`
		Cluster      *VirtualizationIweClusterRelationship `json:"Cluster,omitempty"`
		// An array of relationships to virtualizationIweHost resources.
		MemberHosts []VirtualizationIweHostRelationship `json:"MemberHosts,omitempty"`
		// An array of relationships to virtualizationIweHostVswitch resources.
		MemberVswitches []VirtualizationIweHostVswitchRelationship `json:"MemberVswitches,omitempty"`
	}

	varVirtualizationIweDvswitchWithoutEmbeddedStruct := VirtualizationIweDvswitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDvswitchWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIweDvswitch := _VirtualizationIweDvswitch{}
		varVirtualizationIweDvswitch.ClassId = varVirtualizationIweDvswitchWithoutEmbeddedStruct.ClassId
		varVirtualizationIweDvswitch.ObjectType = varVirtualizationIweDvswitchWithoutEmbeddedStruct.ObjectType
		varVirtualizationIweDvswitch.DvUplink = varVirtualizationIweDvswitchWithoutEmbeddedStruct.DvUplink
		varVirtualizationIweDvswitch.LastHostname = varVirtualizationIweDvswitchWithoutEmbeddedStruct.LastHostname
		varVirtualizationIweDvswitch.Cluster = varVirtualizationIweDvswitchWithoutEmbeddedStruct.Cluster
		varVirtualizationIweDvswitch.MemberHosts = varVirtualizationIweDvswitchWithoutEmbeddedStruct.MemberHosts
		varVirtualizationIweDvswitch.MemberVswitches = varVirtualizationIweDvswitchWithoutEmbeddedStruct.MemberVswitches
		*o = VirtualizationIweDvswitch(varVirtualizationIweDvswitch)
	} else {
		return err
	}

	varVirtualizationIweDvswitch := _VirtualizationIweDvswitch{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDvswitch)
	if err == nil {
		o.VirtualizationBaseDvswitch = varVirtualizationIweDvswitch.VirtualizationBaseDvswitch
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DvUplink")
		delete(additionalProperties, "LastHostname")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "MemberHosts")
		delete(additionalProperties, "MemberVswitches")

		// remove fields from embedded structs
		reflectVirtualizationBaseDvswitch := reflect.ValueOf(o.VirtualizationBaseDvswitch)
		for i := 0; i < reflectVirtualizationBaseDvswitch.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDvswitch.Type().Field(i)

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

type NullableVirtualizationIweDvswitch struct {
	value *VirtualizationIweDvswitch
	isSet bool
}

func (v NullableVirtualizationIweDvswitch) Get() *VirtualizationIweDvswitch {
	return v.value
}

func (v *NullableVirtualizationIweDvswitch) Set(val *VirtualizationIweDvswitch) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweDvswitch) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweDvswitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweDvswitch(val *VirtualizationIweDvswitch) *NullableVirtualizationIweDvswitch {
	return &NullableVirtualizationIweDvswitch{value: val, isSet: true}
}

func (v NullableVirtualizationIweDvswitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweDvswitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
