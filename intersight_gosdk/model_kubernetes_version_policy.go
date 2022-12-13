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

// KubernetesVersionPolicy Policy that defines which kubernetes version to use.
type KubernetesVersionPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                                `json:"ObjectType"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to kubernetesNodeGroupProfile resources.
	Profiles             []KubernetesNodeGroupProfileRelationship `json:"Profiles,omitempty"`
	Version              *KubernetesVersionRelationship           `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVersionPolicy KubernetesVersionPolicy

// NewKubernetesVersionPolicy instantiates a new KubernetesVersionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVersionPolicy(classId string, objectType string) *KubernetesVersionPolicy {
	this := KubernetesVersionPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesVersionPolicyWithDefaults instantiates a new KubernetesVersionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVersionPolicyWithDefaults() *KubernetesVersionPolicy {
	this := KubernetesVersionPolicy{}
	var classId string = "kubernetes.VersionPolicy"
	this.ClassId = classId
	var objectType string = "kubernetes.VersionPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVersionPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVersionPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVersionPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVersionPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVersionPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVersionPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVersionPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVersionPolicy) GetProfiles() []KubernetesNodeGroupProfileRelationship {
	if o == nil {
		var ret []KubernetesNodeGroupProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVersionPolicy) GetProfilesOk() ([]KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVersionPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesNodeGroupProfileRelationship and assigns it to the Profiles field.
func (o *KubernetesVersionPolicy) SetProfiles(v []KubernetesNodeGroupProfileRelationship) {
	o.Profiles = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KubernetesVersionPolicy) GetVersion() KubernetesVersionRelationship {
	if o == nil || o.Version == nil {
		var ret KubernetesVersionRelationship
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVersionPolicy) GetVersionOk() (*KubernetesVersionRelationship, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KubernetesVersionPolicy) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given KubernetesVersionRelationship and assigns it to the Version field.
func (o *KubernetesVersionPolicy) SetVersion(v KubernetesVersionRelationship) {
	o.Version = &v
}

func (o KubernetesVersionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVersionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesVersionPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                `json:"ObjectType"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to kubernetesNodeGroupProfile resources.
		Profiles []KubernetesNodeGroupProfileRelationship `json:"Profiles,omitempty"`
		Version  *KubernetesVersionRelationship           `json:"Version,omitempty"`
	}

	varKubernetesVersionPolicyWithoutEmbeddedStruct := KubernetesVersionPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesVersionPolicyWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesVersionPolicy := _KubernetesVersionPolicy{}
		varKubernetesVersionPolicy.ClassId = varKubernetesVersionPolicyWithoutEmbeddedStruct.ClassId
		varKubernetesVersionPolicy.ObjectType = varKubernetesVersionPolicyWithoutEmbeddedStruct.ObjectType
		varKubernetesVersionPolicy.Organization = varKubernetesVersionPolicyWithoutEmbeddedStruct.Organization
		varKubernetesVersionPolicy.Profiles = varKubernetesVersionPolicyWithoutEmbeddedStruct.Profiles
		varKubernetesVersionPolicy.Version = varKubernetesVersionPolicyWithoutEmbeddedStruct.Version
		*o = KubernetesVersionPolicy(varKubernetesVersionPolicy)
	} else {
		return err
	}

	varKubernetesVersionPolicy := _KubernetesVersionPolicy{}

	err = json.Unmarshal(bytes, &varKubernetesVersionPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesVersionPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")
		delete(additionalProperties, "Version")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableKubernetesVersionPolicy struct {
	value *KubernetesVersionPolicy
	isSet bool
}

func (v NullableKubernetesVersionPolicy) Get() *KubernetesVersionPolicy {
	return v.value
}

func (v *NullableKubernetesVersionPolicy) Set(val *KubernetesVersionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVersionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVersionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVersionPolicy(val *KubernetesVersionPolicy) *NullableKubernetesVersionPolicy {
	return &NullableKubernetesVersionPolicy{value: val, isSet: true}
}

func (v NullableKubernetesVersionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVersionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
