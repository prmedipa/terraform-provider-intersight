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
)

// CloudAwsSubnetAllOf Definition of the list of properties defined in 'cloud.AwsSubnet', excluding properties defined in parent classes.
type CloudAwsSubnetAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// If true, Ipv4 address is assigned automatically.
	AutoAssignPrivateIpV6 *bool `json:"AutoAssignPrivateIpV6,omitempty"`
	// If true, Ipv4 address is assigned automatically.
	AutoAssignPublicIpV4 *bool `json:"AutoAssignPublicIpV4,omitempty"`
	// The Availability Zone name of the subnet.
	AvailabilityZoneName *string `json:"AvailabilityZoneName,omitempty"`
	// The IPv4 CIDR block assigned to the subnet..
	Ipv4Cidr *string `json:"Ipv4Cidr,omitempty"`
	// The IPv6 CIDR block assigned to the subnet..
	Ipv6Cidr *string `json:"Ipv6Cidr,omitempty"`
	// If true, indicates that this is default subnet.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// The state of the subnet (pending | available).
	State                *string                  `json:"State,omitempty"`
	SubnetTags           []CloudCloudTag          `json:"SubnetTags,omitempty"`
	AwsVpc               *CloudAwsVpcRelationship `json:"AwsVpc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsSubnetAllOf CloudAwsSubnetAllOf

// NewCloudAwsSubnetAllOf instantiates a new CloudAwsSubnetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsSubnetAllOf(classId string, objectType string) *CloudAwsSubnetAllOf {
	this := CloudAwsSubnetAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsSubnetAllOfWithDefaults instantiates a new CloudAwsSubnetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsSubnetAllOfWithDefaults() *CloudAwsSubnetAllOf {
	this := CloudAwsSubnetAllOf{}
	var classId string = "cloud.AwsSubnet"
	this.ClassId = classId
	var objectType string = "cloud.AwsSubnet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsSubnetAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsSubnetAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsSubnetAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsSubnetAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAutoAssignPrivateIpV6 returns the AutoAssignPrivateIpV6 field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetAutoAssignPrivateIpV6() bool {
	if o == nil || o.AutoAssignPrivateIpV6 == nil {
		var ret bool
		return ret
	}
	return *o.AutoAssignPrivateIpV6
}

// GetAutoAssignPrivateIpV6Ok returns a tuple with the AutoAssignPrivateIpV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetAutoAssignPrivateIpV6Ok() (*bool, bool) {
	if o == nil || o.AutoAssignPrivateIpV6 == nil {
		return nil, false
	}
	return o.AutoAssignPrivateIpV6, true
}

// HasAutoAssignPrivateIpV6 returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasAutoAssignPrivateIpV6() bool {
	if o != nil && o.AutoAssignPrivateIpV6 != nil {
		return true
	}

	return false
}

// SetAutoAssignPrivateIpV6 gets a reference to the given bool and assigns it to the AutoAssignPrivateIpV6 field.
func (o *CloudAwsSubnetAllOf) SetAutoAssignPrivateIpV6(v bool) {
	o.AutoAssignPrivateIpV6 = &v
}

// GetAutoAssignPublicIpV4 returns the AutoAssignPublicIpV4 field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetAutoAssignPublicIpV4() bool {
	if o == nil || o.AutoAssignPublicIpV4 == nil {
		var ret bool
		return ret
	}
	return *o.AutoAssignPublicIpV4
}

// GetAutoAssignPublicIpV4Ok returns a tuple with the AutoAssignPublicIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetAutoAssignPublicIpV4Ok() (*bool, bool) {
	if o == nil || o.AutoAssignPublicIpV4 == nil {
		return nil, false
	}
	return o.AutoAssignPublicIpV4, true
}

// HasAutoAssignPublicIpV4 returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasAutoAssignPublicIpV4() bool {
	if o != nil && o.AutoAssignPublicIpV4 != nil {
		return true
	}

	return false
}

// SetAutoAssignPublicIpV4 gets a reference to the given bool and assigns it to the AutoAssignPublicIpV4 field.
func (o *CloudAwsSubnetAllOf) SetAutoAssignPublicIpV4(v bool) {
	o.AutoAssignPublicIpV4 = &v
}

// GetAvailabilityZoneName returns the AvailabilityZoneName field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetAvailabilityZoneName() string {
	if o == nil || o.AvailabilityZoneName == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneName
}

// GetAvailabilityZoneNameOk returns a tuple with the AvailabilityZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetAvailabilityZoneNameOk() (*string, bool) {
	if o == nil || o.AvailabilityZoneName == nil {
		return nil, false
	}
	return o.AvailabilityZoneName, true
}

// HasAvailabilityZoneName returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasAvailabilityZoneName() bool {
	if o != nil && o.AvailabilityZoneName != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneName gets a reference to the given string and assigns it to the AvailabilityZoneName field.
func (o *CloudAwsSubnetAllOf) SetAvailabilityZoneName(v string) {
	o.AvailabilityZoneName = &v
}

// GetIpv4Cidr returns the Ipv4Cidr field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetIpv4Cidr() string {
	if o == nil || o.Ipv4Cidr == nil {
		var ret string
		return ret
	}
	return *o.Ipv4Cidr
}

// GetIpv4CidrOk returns a tuple with the Ipv4Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetIpv4CidrOk() (*string, bool) {
	if o == nil || o.Ipv4Cidr == nil {
		return nil, false
	}
	return o.Ipv4Cidr, true
}

// HasIpv4Cidr returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasIpv4Cidr() bool {
	if o != nil && o.Ipv4Cidr != nil {
		return true
	}

	return false
}

// SetIpv4Cidr gets a reference to the given string and assigns it to the Ipv4Cidr field.
func (o *CloudAwsSubnetAllOf) SetIpv4Cidr(v string) {
	o.Ipv4Cidr = &v
}

// GetIpv6Cidr returns the Ipv6Cidr field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetIpv6Cidr() string {
	if o == nil || o.Ipv6Cidr == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Cidr
}

// GetIpv6CidrOk returns a tuple with the Ipv6Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetIpv6CidrOk() (*string, bool) {
	if o == nil || o.Ipv6Cidr == nil {
		return nil, false
	}
	return o.Ipv6Cidr, true
}

// HasIpv6Cidr returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasIpv6Cidr() bool {
	if o != nil && o.Ipv6Cidr != nil {
		return true
	}

	return false
}

// SetIpv6Cidr gets a reference to the given string and assigns it to the Ipv6Cidr field.
func (o *CloudAwsSubnetAllOf) SetIpv6Cidr(v string) {
	o.Ipv6Cidr = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudAwsSubnetAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudAwsSubnetAllOf) SetState(v string) {
	o.State = &v
}

// GetSubnetTags returns the SubnetTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSubnetAllOf) GetSubnetTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.SubnetTags
}

// GetSubnetTagsOk returns a tuple with the SubnetTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSubnetAllOf) GetSubnetTagsOk() ([]CloudCloudTag, bool) {
	if o == nil || o.SubnetTags == nil {
		return nil, false
	}
	return o.SubnetTags, true
}

// HasSubnetTags returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasSubnetTags() bool {
	if o != nil && o.SubnetTags != nil {
		return true
	}

	return false
}

// SetSubnetTags gets a reference to the given []CloudCloudTag and assigns it to the SubnetTags field.
func (o *CloudAwsSubnetAllOf) SetSubnetTags(v []CloudCloudTag) {
	o.SubnetTags = v
}

// GetAwsVpc returns the AwsVpc field value if set, zero value otherwise.
func (o *CloudAwsSubnetAllOf) GetAwsVpc() CloudAwsVpcRelationship {
	if o == nil || o.AwsVpc == nil {
		var ret CloudAwsVpcRelationship
		return ret
	}
	return *o.AwsVpc
}

// GetAwsVpcOk returns a tuple with the AwsVpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnetAllOf) GetAwsVpcOk() (*CloudAwsVpcRelationship, bool) {
	if o == nil || o.AwsVpc == nil {
		return nil, false
	}
	return o.AwsVpc, true
}

// HasAwsVpc returns a boolean if a field has been set.
func (o *CloudAwsSubnetAllOf) HasAwsVpc() bool {
	if o != nil && o.AwsVpc != nil {
		return true
	}

	return false
}

// SetAwsVpc gets a reference to the given CloudAwsVpcRelationship and assigns it to the AwsVpc field.
func (o *CloudAwsSubnetAllOf) SetAwsVpc(v CloudAwsVpcRelationship) {
	o.AwsVpc = &v
}

func (o CloudAwsSubnetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AutoAssignPrivateIpV6 != nil {
		toSerialize["AutoAssignPrivateIpV6"] = o.AutoAssignPrivateIpV6
	}
	if o.AutoAssignPublicIpV4 != nil {
		toSerialize["AutoAssignPublicIpV4"] = o.AutoAssignPublicIpV4
	}
	if o.AvailabilityZoneName != nil {
		toSerialize["AvailabilityZoneName"] = o.AvailabilityZoneName
	}
	if o.Ipv4Cidr != nil {
		toSerialize["Ipv4Cidr"] = o.Ipv4Cidr
	}
	if o.Ipv6Cidr != nil {
		toSerialize["Ipv6Cidr"] = o.Ipv6Cidr
	}
	if o.IsDefault != nil {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubnetTags != nil {
		toSerialize["SubnetTags"] = o.SubnetTags
	}
	if o.AwsVpc != nil {
		toSerialize["AwsVpc"] = o.AwsVpc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsSubnetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsSubnetAllOf := _CloudAwsSubnetAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsSubnetAllOf); err == nil {
		*o = CloudAwsSubnetAllOf(varCloudAwsSubnetAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AutoAssignPrivateIpV6")
		delete(additionalProperties, "AutoAssignPublicIpV4")
		delete(additionalProperties, "AvailabilityZoneName")
		delete(additionalProperties, "Ipv4Cidr")
		delete(additionalProperties, "Ipv6Cidr")
		delete(additionalProperties, "IsDefault")
		delete(additionalProperties, "State")
		delete(additionalProperties, "SubnetTags")
		delete(additionalProperties, "AwsVpc")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsSubnetAllOf struct {
	value *CloudAwsSubnetAllOf
	isSet bool
}

func (v NullableCloudAwsSubnetAllOf) Get() *CloudAwsSubnetAllOf {
	return v.value
}

func (v *NullableCloudAwsSubnetAllOf) Set(val *CloudAwsSubnetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsSubnetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsSubnetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsSubnetAllOf(val *CloudAwsSubnetAllOf) *NullableCloudAwsSubnetAllOf {
	return &NullableCloudAwsSubnetAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsSubnetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsSubnetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
