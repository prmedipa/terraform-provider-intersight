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

// IamEndPointUserPolicyInventoryAllOf Definition of the list of properties defined in 'iam.EndPointUserPolicyInventory', excluding properties defined in parent classes.
type IamEndPointUserPolicyInventoryAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType         string                                `json:"ObjectType"`
	PasswordProperties NullableIamEndPointPasswordProperties `json:"PasswordProperties,omitempty"`
	// An array of relationships to iamEndPointUserRoleInventory resources.
	EndPointUserRoles    []IamEndPointUserRoleInventoryRelationship `json:"EndPointUserRoles,omitempty"`
	TargetMo             *MoBaseMoRelationship                      `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamEndPointUserPolicyInventoryAllOf IamEndPointUserPolicyInventoryAllOf

// NewIamEndPointUserPolicyInventoryAllOf instantiates a new IamEndPointUserPolicyInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamEndPointUserPolicyInventoryAllOf(classId string, objectType string) *IamEndPointUserPolicyInventoryAllOf {
	this := IamEndPointUserPolicyInventoryAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamEndPointUserPolicyInventoryAllOfWithDefaults instantiates a new IamEndPointUserPolicyInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamEndPointUserPolicyInventoryAllOfWithDefaults() *IamEndPointUserPolicyInventoryAllOf {
	this := IamEndPointUserPolicyInventoryAllOf{}
	var classId string = "iam.EndPointUserPolicyInventory"
	this.ClassId = classId
	var objectType string = "iam.EndPointUserPolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamEndPointUserPolicyInventoryAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamEndPointUserPolicyInventoryAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamEndPointUserPolicyInventoryAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamEndPointUserPolicyInventoryAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPasswordProperties returns the PasswordProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyInventoryAllOf) GetPasswordProperties() IamEndPointPasswordProperties {
	if o == nil || o.PasswordProperties.Get() == nil {
		var ret IamEndPointPasswordProperties
		return ret
	}
	return *o.PasswordProperties.Get()
}

// GetPasswordPropertiesOk returns a tuple with the PasswordProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyInventoryAllOf) GetPasswordPropertiesOk() (*IamEndPointPasswordProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordProperties.Get(), o.PasswordProperties.IsSet()
}

// HasPasswordProperties returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) HasPasswordProperties() bool {
	if o != nil && o.PasswordProperties.IsSet() {
		return true
	}

	return false
}

// SetPasswordProperties gets a reference to the given NullableIamEndPointPasswordProperties and assigns it to the PasswordProperties field.
func (o *IamEndPointUserPolicyInventoryAllOf) SetPasswordProperties(v IamEndPointPasswordProperties) {
	o.PasswordProperties.Set(&v)
}

// SetPasswordPropertiesNil sets the value for PasswordProperties to be an explicit nil
func (o *IamEndPointUserPolicyInventoryAllOf) SetPasswordPropertiesNil() {
	o.PasswordProperties.Set(nil)
}

// UnsetPasswordProperties ensures that no value is present for PasswordProperties, not even an explicit nil
func (o *IamEndPointUserPolicyInventoryAllOf) UnsetPasswordProperties() {
	o.PasswordProperties.Unset()
}

// GetEndPointUserRoles returns the EndPointUserRoles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamEndPointUserPolicyInventoryAllOf) GetEndPointUserRoles() []IamEndPointUserRoleInventoryRelationship {
	if o == nil {
		var ret []IamEndPointUserRoleInventoryRelationship
		return ret
	}
	return o.EndPointUserRoles
}

// GetEndPointUserRolesOk returns a tuple with the EndPointUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamEndPointUserPolicyInventoryAllOf) GetEndPointUserRolesOk() ([]IamEndPointUserRoleInventoryRelationship, bool) {
	if o == nil || o.EndPointUserRoles == nil {
		return nil, false
	}
	return o.EndPointUserRoles, true
}

// HasEndPointUserRoles returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) HasEndPointUserRoles() bool {
	if o != nil && o.EndPointUserRoles != nil {
		return true
	}

	return false
}

// SetEndPointUserRoles gets a reference to the given []IamEndPointUserRoleInventoryRelationship and assigns it to the EndPointUserRoles field.
func (o *IamEndPointUserPolicyInventoryAllOf) SetEndPointUserRoles(v []IamEndPointUserRoleInventoryRelationship) {
	o.EndPointUserRoles = v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *IamEndPointUserPolicyInventoryAllOf) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *IamEndPointUserPolicyInventoryAllOf) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *IamEndPointUserPolicyInventoryAllOf) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o IamEndPointUserPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.PasswordProperties.IsSet() {
		toSerialize["PasswordProperties"] = o.PasswordProperties.Get()
	}
	if o.EndPointUserRoles != nil {
		toSerialize["EndPointUserRoles"] = o.EndPointUserRoles
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamEndPointUserPolicyInventoryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamEndPointUserPolicyInventoryAllOf := _IamEndPointUserPolicyInventoryAllOf{}

	if err = json.Unmarshal(bytes, &varIamEndPointUserPolicyInventoryAllOf); err == nil {
		*o = IamEndPointUserPolicyInventoryAllOf(varIamEndPointUserPolicyInventoryAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "PasswordProperties")
		delete(additionalProperties, "EndPointUserRoles")
		delete(additionalProperties, "TargetMo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamEndPointUserPolicyInventoryAllOf struct {
	value *IamEndPointUserPolicyInventoryAllOf
	isSet bool
}

func (v NullableIamEndPointUserPolicyInventoryAllOf) Get() *IamEndPointUserPolicyInventoryAllOf {
	return v.value
}

func (v *NullableIamEndPointUserPolicyInventoryAllOf) Set(val *IamEndPointUserPolicyInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamEndPointUserPolicyInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamEndPointUserPolicyInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamEndPointUserPolicyInventoryAllOf(val *IamEndPointUserPolicyInventoryAllOf) *NullableIamEndPointUserPolicyInventoryAllOf {
	return &NullableIamEndPointUserPolicyInventoryAllOf{value: val, isSet: true}
}

func (v NullableIamEndPointUserPolicyInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamEndPointUserPolicyInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
