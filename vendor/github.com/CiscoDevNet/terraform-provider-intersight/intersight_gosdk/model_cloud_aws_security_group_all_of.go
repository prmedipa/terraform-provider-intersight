/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4929
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// CloudAwsSecurityGroupAllOf Definition of the list of properties defined in 'cloud.AwsSecurityGroup', excluding properties defined in parent classes.
type CloudAwsSecurityGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	EgressRules          []CloudSecurityGroupRule `json:"EgressRules,omitempty"`
	IngressRules         []CloudSecurityGroupRule `json:"IngressRules,omitempty"`
	SecurityGroupTags    []CloudCloudTag          `json:"SecurityGroupTags,omitempty"`
	Location             *CloudAwsVpcRelationship `json:"Location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudAwsSecurityGroupAllOf CloudAwsSecurityGroupAllOf

// NewCloudAwsSecurityGroupAllOf instantiates a new CloudAwsSecurityGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAwsSecurityGroupAllOf(classId string, objectType string) *CloudAwsSecurityGroupAllOf {
	this := CloudAwsSecurityGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudAwsSecurityGroupAllOfWithDefaults instantiates a new CloudAwsSecurityGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAwsSecurityGroupAllOfWithDefaults() *CloudAwsSecurityGroupAllOf {
	this := CloudAwsSecurityGroupAllOf{}
	var classId string = "cloud.AwsSecurityGroup"
	this.ClassId = classId
	var objectType string = "cloud.AwsSecurityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudAwsSecurityGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSecurityGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudAwsSecurityGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudAwsSecurityGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudAwsSecurityGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudAwsSecurityGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEgressRules returns the EgressRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSecurityGroupAllOf) GetEgressRules() []CloudSecurityGroupRule {
	if o == nil {
		var ret []CloudSecurityGroupRule
		return ret
	}
	return o.EgressRules
}

// GetEgressRulesOk returns a tuple with the EgressRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSecurityGroupAllOf) GetEgressRulesOk() (*[]CloudSecurityGroupRule, bool) {
	if o == nil || o.EgressRules == nil {
		return nil, false
	}
	return &o.EgressRules, true
}

// HasEgressRules returns a boolean if a field has been set.
func (o *CloudAwsSecurityGroupAllOf) HasEgressRules() bool {
	if o != nil && o.EgressRules != nil {
		return true
	}

	return false
}

// SetEgressRules gets a reference to the given []CloudSecurityGroupRule and assigns it to the EgressRules field.
func (o *CloudAwsSecurityGroupAllOf) SetEgressRules(v []CloudSecurityGroupRule) {
	o.EgressRules = v
}

// GetIngressRules returns the IngressRules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSecurityGroupAllOf) GetIngressRules() []CloudSecurityGroupRule {
	if o == nil {
		var ret []CloudSecurityGroupRule
		return ret
	}
	return o.IngressRules
}

// GetIngressRulesOk returns a tuple with the IngressRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSecurityGroupAllOf) GetIngressRulesOk() (*[]CloudSecurityGroupRule, bool) {
	if o == nil || o.IngressRules == nil {
		return nil, false
	}
	return &o.IngressRules, true
}

// HasIngressRules returns a boolean if a field has been set.
func (o *CloudAwsSecurityGroupAllOf) HasIngressRules() bool {
	if o != nil && o.IngressRules != nil {
		return true
	}

	return false
}

// SetIngressRules gets a reference to the given []CloudSecurityGroupRule and assigns it to the IngressRules field.
func (o *CloudAwsSecurityGroupAllOf) SetIngressRules(v []CloudSecurityGroupRule) {
	o.IngressRules = v
}

// GetSecurityGroupTags returns the SecurityGroupTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAwsSecurityGroupAllOf) GetSecurityGroupTags() []CloudCloudTag {
	if o == nil {
		var ret []CloudCloudTag
		return ret
	}
	return o.SecurityGroupTags
}

// GetSecurityGroupTagsOk returns a tuple with the SecurityGroupTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAwsSecurityGroupAllOf) GetSecurityGroupTagsOk() (*[]CloudCloudTag, bool) {
	if o == nil || o.SecurityGroupTags == nil {
		return nil, false
	}
	return &o.SecurityGroupTags, true
}

// HasSecurityGroupTags returns a boolean if a field has been set.
func (o *CloudAwsSecurityGroupAllOf) HasSecurityGroupTags() bool {
	if o != nil && o.SecurityGroupTags != nil {
		return true
	}

	return false
}

// SetSecurityGroupTags gets a reference to the given []CloudCloudTag and assigns it to the SecurityGroupTags field.
func (o *CloudAwsSecurityGroupAllOf) SetSecurityGroupTags(v []CloudCloudTag) {
	o.SecurityGroupTags = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudAwsSecurityGroupAllOf) GetLocation() CloudAwsVpcRelationship {
	if o == nil || o.Location == nil {
		var ret CloudAwsVpcRelationship
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudAwsSecurityGroupAllOf) GetLocationOk() (*CloudAwsVpcRelationship, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudAwsSecurityGroupAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given CloudAwsVpcRelationship and assigns it to the Location field.
func (o *CloudAwsSecurityGroupAllOf) SetLocation(v CloudAwsVpcRelationship) {
	o.Location = &v
}

func (o CloudAwsSecurityGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EgressRules != nil {
		toSerialize["EgressRules"] = o.EgressRules
	}
	if o.IngressRules != nil {
		toSerialize["IngressRules"] = o.IngressRules
	}
	if o.SecurityGroupTags != nil {
		toSerialize["SecurityGroupTags"] = o.SecurityGroupTags
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudAwsSecurityGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudAwsSecurityGroupAllOf := _CloudAwsSecurityGroupAllOf{}

	if err = json.Unmarshal(bytes, &varCloudAwsSecurityGroupAllOf); err == nil {
		*o = CloudAwsSecurityGroupAllOf(varCloudAwsSecurityGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EgressRules")
		delete(additionalProperties, "IngressRules")
		delete(additionalProperties, "SecurityGroupTags")
		delete(additionalProperties, "Location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudAwsSecurityGroupAllOf struct {
	value *CloudAwsSecurityGroupAllOf
	isSet bool
}

func (v NullableCloudAwsSecurityGroupAllOf) Get() *CloudAwsSecurityGroupAllOf {
	return v.value
}

func (v *NullableCloudAwsSecurityGroupAllOf) Set(val *CloudAwsSecurityGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAwsSecurityGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAwsSecurityGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAwsSecurityGroupAllOf(val *CloudAwsSecurityGroupAllOf) *NullableCloudAwsSecurityGroupAllOf {
	return &NullableCloudAwsSecurityGroupAllOf{value: val, isSet: true}
}

func (v NullableCloudAwsSecurityGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAwsSecurityGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
