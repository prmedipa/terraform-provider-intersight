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

// IamSecurityHolder Holder for organization aggregated permissions and global account permissions. User configures permissions for entire account or subset of organizations and specifies associated roles with each organization. Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
type IamSecurityHolder struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	Account              *IamAccountRelationship            `json:"Account,omitempty"`
	IpRulesConfiguration *IamIpAccessManagementRelationship `json:"IpRulesConfiguration,omitempty"`
	// An array of relationships to iamResourcePermission resources.
	ResourcePermissions  []IamResourcePermissionRelationship `json:"ResourcePermissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamSecurityHolder IamSecurityHolder

// NewIamSecurityHolder instantiates a new IamSecurityHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamSecurityHolder(classId string, objectType string) *IamSecurityHolder {
	this := IamSecurityHolder{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamSecurityHolderWithDefaults instantiates a new IamSecurityHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamSecurityHolderWithDefaults() *IamSecurityHolder {
	this := IamSecurityHolder{}
	var classId string = "iam.SecurityHolder"
	this.ClassId = classId
	var objectType string = "iam.SecurityHolder"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamSecurityHolder) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamSecurityHolder) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamSecurityHolder) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamSecurityHolder) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamSecurityHolder) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamSecurityHolder) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamSecurityHolder) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSecurityHolder) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamSecurityHolder) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamSecurityHolder) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetIpRulesConfiguration returns the IpRulesConfiguration field value if set, zero value otherwise.
func (o *IamSecurityHolder) GetIpRulesConfiguration() IamIpAccessManagementRelationship {
	if o == nil || o.IpRulesConfiguration == nil {
		var ret IamIpAccessManagementRelationship
		return ret
	}
	return *o.IpRulesConfiguration
}

// GetIpRulesConfigurationOk returns a tuple with the IpRulesConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamSecurityHolder) GetIpRulesConfigurationOk() (*IamIpAccessManagementRelationship, bool) {
	if o == nil || o.IpRulesConfiguration == nil {
		return nil, false
	}
	return o.IpRulesConfiguration, true
}

// HasIpRulesConfiguration returns a boolean if a field has been set.
func (o *IamSecurityHolder) HasIpRulesConfiguration() bool {
	if o != nil && o.IpRulesConfiguration != nil {
		return true
	}

	return false
}

// SetIpRulesConfiguration gets a reference to the given IamIpAccessManagementRelationship and assigns it to the IpRulesConfiguration field.
func (o *IamSecurityHolder) SetIpRulesConfiguration(v IamIpAccessManagementRelationship) {
	o.IpRulesConfiguration = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamSecurityHolder) GetResourcePermissions() []IamResourcePermissionRelationship {
	if o == nil {
		var ret []IamResourcePermissionRelationship
		return ret
	}
	return o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamSecurityHolder) GetResourcePermissionsOk() ([]IamResourcePermissionRelationship, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *IamSecurityHolder) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given []IamResourcePermissionRelationship and assigns it to the ResourcePermissions field.
func (o *IamSecurityHolder) SetResourcePermissions(v []IamResourcePermissionRelationship) {
	o.ResourcePermissions = v
}

func (o IamSecurityHolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.IpRulesConfiguration != nil {
		toSerialize["IpRulesConfiguration"] = o.IpRulesConfiguration
	}
	if o.ResourcePermissions != nil {
		toSerialize["ResourcePermissions"] = o.ResourcePermissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamSecurityHolder) UnmarshalJSON(bytes []byte) (err error) {
	type IamSecurityHolderWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                             `json:"ObjectType"`
		Account              *IamAccountRelationship            `json:"Account,omitempty"`
		IpRulesConfiguration *IamIpAccessManagementRelationship `json:"IpRulesConfiguration,omitempty"`
		// An array of relationships to iamResourcePermission resources.
		ResourcePermissions []IamResourcePermissionRelationship `json:"ResourcePermissions,omitempty"`
	}

	varIamSecurityHolderWithoutEmbeddedStruct := IamSecurityHolderWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamSecurityHolderWithoutEmbeddedStruct)
	if err == nil {
		varIamSecurityHolder := _IamSecurityHolder{}
		varIamSecurityHolder.ClassId = varIamSecurityHolderWithoutEmbeddedStruct.ClassId
		varIamSecurityHolder.ObjectType = varIamSecurityHolderWithoutEmbeddedStruct.ObjectType
		varIamSecurityHolder.Account = varIamSecurityHolderWithoutEmbeddedStruct.Account
		varIamSecurityHolder.IpRulesConfiguration = varIamSecurityHolderWithoutEmbeddedStruct.IpRulesConfiguration
		varIamSecurityHolder.ResourcePermissions = varIamSecurityHolderWithoutEmbeddedStruct.ResourcePermissions
		*o = IamSecurityHolder(varIamSecurityHolder)
	} else {
		return err
	}

	varIamSecurityHolder := _IamSecurityHolder{}

	err = json.Unmarshal(bytes, &varIamSecurityHolder)
	if err == nil {
		o.MoBaseMo = varIamSecurityHolder.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "IpRulesConfiguration")
		delete(additionalProperties, "ResourcePermissions")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableIamSecurityHolder struct {
	value *IamSecurityHolder
	isSet bool
}

func (v NullableIamSecurityHolder) Get() *IamSecurityHolder {
	return v.value
}

func (v *NullableIamSecurityHolder) Set(val *IamSecurityHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableIamSecurityHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableIamSecurityHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamSecurityHolder(val *IamSecurityHolder) *NullableIamSecurityHolder {
	return &NullableIamSecurityHolder{value: val, isSet: true}
}

func (v NullableIamSecurityHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamSecurityHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
