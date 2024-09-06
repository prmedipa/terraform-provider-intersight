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

// checks if the CloudAwsSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudAwsSubnet{}

// CloudAwsSubnet Subnet object in AWS inventory. It is a range of IP addresses in a VPC that can be used to isolate different EC2 resources from each other or from the Internet.
type CloudAwsSubnet struct {
	CloudBaseNetwork
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
	State                *string                         `json:"State,omitempty"`
	SubnetTags           []CloudCloudTag                 `json:"SubnetTags,omitempty"`
	AwsVpc               NullableCloudAwsVpcRelationship `json:"AwsVpc,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsSubnet CloudAwsSubnet

// NewCloudAwsSubnet instantiates a new CloudAwsSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsSubnet(classId string, objectType string) *CloudAwsSubnet {
	this := CloudAwsSubnet{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsSubnetWithDefaults instantiates a new CloudAwsSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsSubnetWithDefaults() *CloudAwsSubnet {
	this := CloudAwsSubnet{}
	var classId string = "cloud.AwsSubnet"
	this.ClassId = classId
	var objectType string = "cloud.AwsSubnet"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsSubnet) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsSubnet) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "cloud.AwsSubnet" of the ClassId field.
func (o *CloudAwsSubnet) GetDefaultClassId() interface{} {
	return "cloud.AwsSubnet"
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsSubnet) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsSubnet) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "cloud.AwsSubnet" of the ObjectType field.
func (o *CloudAwsSubnet) GetDefaultObjectType() interface{} {
	return "cloud.AwsSubnet"
}

// GetAutoAssignPrivateIpV6 returns the AutoAssignPrivateIpV6 field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetAutoAssignPrivateIpV6() bool {
	if o == nil || IsNil(o.AutoAssignPrivateIpV6) {
		var ret bool
		return ret
	}
	return *o.AutoAssignPrivateIpV6
}

// GetAutoAssignPrivateIpV6Ok returns a tuple with the AutoAssignPrivateIpV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetAutoAssignPrivateIpV6Ok() (*bool, bool) {
	if o == nil || IsNil(o.AutoAssignPrivateIpV6) {
		return nil, false
	}
	return o.AutoAssignPrivateIpV6, true
}

// HasAutoAssignPrivateIpV6 returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasAutoAssignPrivateIpV6() bool {
	if o != nil && !IsNil(o.AutoAssignPrivateIpV6) {
		return true
	}

	return false
}

// SetAutoAssignPrivateIpV6 gets a reference to the given bool and assigns it to the AutoAssignPrivateIpV6 field.
func (o *CloudAwsSubnet) SetAutoAssignPrivateIpV6(v bool) {
	o.AutoAssignPrivateIpV6 = &v
}

// GetAutoAssignPublicIpV4 returns the AutoAssignPublicIpV4 field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetAutoAssignPublicIpV4() bool {
	if o == nil || IsNil(o.AutoAssignPublicIpV4) {
		var ret bool
		return ret
	}
	return *o.AutoAssignPublicIpV4
}

// GetAutoAssignPublicIpV4Ok returns a tuple with the AutoAssignPublicIpV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetAutoAssignPublicIpV4Ok() (*bool, bool) {
	if o == nil || IsNil(o.AutoAssignPublicIpV4) {
		return nil, false
	}
	return o.AutoAssignPublicIpV4, true
}

// HasAutoAssignPublicIpV4 returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasAutoAssignPublicIpV4() bool {
	if o != nil && !IsNil(o.AutoAssignPublicIpV4) {
		return true
	}

	return false
}

// SetAutoAssignPublicIpV4 gets a reference to the given bool and assigns it to the AutoAssignPublicIpV4 field.
func (o *CloudAwsSubnet) SetAutoAssignPublicIpV4(v bool) {
	o.AutoAssignPublicIpV4 = &v
}

// GetAvailabilityZoneName returns the AvailabilityZoneName field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetAvailabilityZoneName() string {
	if o == nil || IsNil(o.AvailabilityZoneName) {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneName
}

// GetAvailabilityZoneNameOk returns a tuple with the AvailabilityZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetAvailabilityZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZoneName) {
		return nil, false
	}
	return o.AvailabilityZoneName, true
}

// HasAvailabilityZoneName returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasAvailabilityZoneName() bool {
	if o != nil && !IsNil(o.AvailabilityZoneName) {
		return true
	}

	return false
}

// SetAvailabilityZoneName gets a reference to the given string and assigns it to the AvailabilityZoneName field.
func (o *CloudAwsSubnet) SetAvailabilityZoneName(v string) {
	o.AvailabilityZoneName = &v
}

// GetIpv4Cidr returns the Ipv4Cidr field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetIpv4Cidr() string {
	if o == nil || IsNil(o.Ipv4Cidr) {
		var ret string
		return ret
	}
	return *o.Ipv4Cidr
}

// GetIpv4CidrOk returns a tuple with the Ipv4Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetIpv4CidrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Cidr) {
		return nil, false
	}
	return o.Ipv4Cidr, true
}

// HasIpv4Cidr returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasIpv4Cidr() bool {
	if o != nil && !IsNil(o.Ipv4Cidr) {
		return true
	}

	return false
}

// SetIpv4Cidr gets a reference to the given string and assigns it to the Ipv4Cidr field.
func (o *CloudAwsSubnet) SetIpv4Cidr(v string) {
	o.Ipv4Cidr = &v
}

// GetIpv6Cidr returns the Ipv6Cidr field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetIpv6Cidr() string {
	if o == nil || IsNil(o.Ipv6Cidr) {
		var ret string
		return ret
	}
	return *o.Ipv6Cidr
}

// GetIpv6CidrOk returns a tuple with the Ipv6Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetIpv6CidrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6Cidr) {
		return nil, false
	}
	return o.Ipv6Cidr, true
}

// HasIpv6Cidr returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasIpv6Cidr() bool {
	if o != nil && !IsNil(o.Ipv6Cidr) {
		return true
	}

	return false
}

// SetIpv6Cidr gets a reference to the given string and assigns it to the Ipv6Cidr field.
func (o *CloudAwsSubnet) SetIpv6Cidr(v string) {
	o.Ipv6Cidr = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CloudAwsSubnet) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudAwsSubnet) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSubnet) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudAwsSubnet) SetState(v string) {
	o.State = &v
}

// GetSubnetTags returns the SubnetTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSubnet) GetSubnetTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.SubnetTags
}

// GetSubnetTagsOk returns a tuple with the SubnetTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSubnet) GetSubnetTagsOk() ([]CloudCloudTag, bool) {
	if o == nil || IsNil(o.SubnetTags) {
		return nil, false
	}
	return o.SubnetTags, true
}

// HasSubnetTags returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasSubnetTags() bool {
	if o != nil && !IsNil(o.SubnetTags) {
		return true
	}

	return false
}

// SetSubnetTags gets a reference to the given []CloudCloudTag and assigns it to the SubnetTags field.
func (o *CloudAwsSubnet) SetSubnetTags(v []CloudCloudTag) {
	o.SubnetTags = v
}

// GetAwsVpc returns the AwsVpc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSubnet) GetAwsVpc() CloudAwsVpcRelationship {
	if o == nil || IsNil(o.AwsVpc.Get()) {
		var ret CloudAwsVpcRelationship
		return ret
	}
	return *o.AwsVpc.Get()
}

// GetAwsVpcOk returns a tuple with the AwsVpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSubnet) GetAwsVpcOk() (*CloudAwsVpcRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.AwsVpc.Get(), o.AwsVpc.IsSet()
}

// HasAwsVpc returns a boolean if a field has been set.
func (o *CloudAwsSubnet) HasAwsVpc() bool {
	if o != nil && o.AwsVpc.IsSet() {
		return true
	}

	return false
}

// SetAwsVpc gets a reference to the given NullableCloudAwsVpcRelationship and assigns it to the AwsVpc field.
func (o *CloudAwsSubnet) SetAwsVpc(v CloudAwsVpcRelationship) {
	o.AwsVpc.Set(&v)
}

// SetAwsVpcNil sets the value for AwsVpc to be an explicit nil
func (o *CloudAwsSubnet) SetAwsVpcNil() {
	o.AwsVpc.Set(nil)
}

// UnsetAwsVpc ensures that no value is present for AwsVpc, not even an explicit nil
func (o *CloudAwsSubnet) UnsetAwsVpc() {
	o.AwsVpc.Unset()
}

func (o CloudAwsSubnet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudAwsSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseNetwork, errCloudBaseNetwork := json.Marshal(o.CloudBaseNetwork)
	if errCloudBaseNetwork != nil {
		return map[string]interface{}{}, errCloudBaseNetwork
	}
	errCloudBaseNetwork = json.Unmarshal([]byte(serializedCloudBaseNetwork), &toSerialize)
	if errCloudBaseNetwork != nil {
		return map[string]interface{}{}, errCloudBaseNetwork
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AutoAssignPrivateIpV6) {
		toSerialize["AutoAssignPrivateIpV6"] = o.AutoAssignPrivateIpV6
	}
	if !IsNil(o.AutoAssignPublicIpV4) {
		toSerialize["AutoAssignPublicIpV4"] = o.AutoAssignPublicIpV4
	}
	if !IsNil(o.AvailabilityZoneName) {
		toSerialize["AvailabilityZoneName"] = o.AvailabilityZoneName
	}
	if !IsNil(o.Ipv4Cidr) {
		toSerialize["Ipv4Cidr"] = o.Ipv4Cidr
	}
	if !IsNil(o.Ipv6Cidr) {
		toSerialize["Ipv6Cidr"] = o.Ipv6Cidr
	}
	if !IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if o.SubnetTags != nil {
		toSerialize["SubnetTags"] = o.SubnetTags
	}
	if o.AwsVpc.IsSet() {
		toSerialize["AwsVpc"] = o.AwsVpc.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudAwsSubnet) UnmarshalJSON(data []byte) (err error) {
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
	type CloudAwsSubnetWithoutEmbeddedStruct struct {
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
		State      *string                         `json:"State,omitempty"`
		SubnetTags []CloudCloudTag                 `json:"SubnetTags,omitempty"`
		AwsVpc     NullableCloudAwsVpcRelationship `json:"AwsVpc,omitempty"`
	}

	varCloudAwsSubnetWithoutEmbeddedStruct := CloudAwsSubnetWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCloudAwsSubnetWithoutEmbeddedStruct)
	if err == nil {
		varCloudAwsSubnet := _CloudAwsSubnet{}
		varCloudAwsSubnet.ClassId = varCloudAwsSubnetWithoutEmbeddedStruct.ClassId
		varCloudAwsSubnet.ObjectType = varCloudAwsSubnetWithoutEmbeddedStruct.ObjectType
		varCloudAwsSubnet.AutoAssignPrivateIpV6 = varCloudAwsSubnetWithoutEmbeddedStruct.AutoAssignPrivateIpV6
		varCloudAwsSubnet.AutoAssignPublicIpV4 = varCloudAwsSubnetWithoutEmbeddedStruct.AutoAssignPublicIpV4
		varCloudAwsSubnet.AvailabilityZoneName = varCloudAwsSubnetWithoutEmbeddedStruct.AvailabilityZoneName
		varCloudAwsSubnet.Ipv4Cidr = varCloudAwsSubnetWithoutEmbeddedStruct.Ipv4Cidr
		varCloudAwsSubnet.Ipv6Cidr = varCloudAwsSubnetWithoutEmbeddedStruct.Ipv6Cidr
		varCloudAwsSubnet.IsDefault = varCloudAwsSubnetWithoutEmbeddedStruct.IsDefault
		varCloudAwsSubnet.State = varCloudAwsSubnetWithoutEmbeddedStruct.State
		varCloudAwsSubnet.SubnetTags = varCloudAwsSubnetWithoutEmbeddedStruct.SubnetTags
		varCloudAwsSubnet.AwsVpc = varCloudAwsSubnetWithoutEmbeddedStruct.AwsVpc
		*o = CloudAwsSubnet(varCloudAwsSubnet)
	} else {
		return err
	}

	varCloudAwsSubnet := _CloudAwsSubnet{}

	err = json.Unmarshal(data, &varCloudAwsSubnet)
	if err == nil {
		o.CloudBaseNetwork = varCloudAwsSubnet.CloudBaseNetwork
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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

		// remove fields from embedded structs
		reflectCloudBaseNetwork := reflect.ValueOf(o.CloudBaseNetwork)
		for i := 0; i < reflectCloudBaseNetwork.Type().NumField(); i++ {
			t := reflectCloudBaseNetwork.Type().Field(i)

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

type NullableCloudAwsSubnet struct {
	value *CloudAwsSubnet
	isSet bool
}

func (v NullableCloudAwsSubnet) Get() *CloudAwsSubnet {
	return v.value
}

func (v *NullableCloudAwsSubnet) Set(val *CloudAwsSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsSubnet(val *CloudAwsSubnet) *NullableCloudAwsSubnet {
	return &NullableCloudAwsSubnet{value: val, isSet: true}
}

func (v NullableCloudAwsSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
