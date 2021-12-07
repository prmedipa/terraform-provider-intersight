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
	"reflect"
	"strings"
)

// FabricFcNetworkPolicy A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.
type FabricFcNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or Disable Trunking on all of configured FC uplink ports.
	EnableTrunking *bool                                 `json:"EnableTrunking,omitempty"`
	Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to fabricSwitchProfile resources.
	Profiles             []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricFcNetworkPolicy FabricFcNetworkPolicy

// NewFabricFcNetworkPolicy instantiates a new FabricFcNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricFcNetworkPolicy(classId string, objectType string) *FabricFcNetworkPolicy {
	this := FabricFcNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFabricFcNetworkPolicyWithDefaults instantiates a new FabricFcNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricFcNetworkPolicyWithDefaults() *FabricFcNetworkPolicy {
	this := FabricFcNetworkPolicy{}
	var classId string = "fabric.FcNetworkPolicy"
	this.ClassId = classId
	var objectType string = "fabric.FcNetworkPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FabricFcNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FabricFcNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FabricFcNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FabricFcNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnableTrunking returns the EnableTrunking field value if set, zero value otherwise.
func (o *FabricFcNetworkPolicy) GetEnableTrunking() bool {
	if o == nil || o.EnableTrunking == nil {
		var ret bool
		return ret
	}
	return *o.EnableTrunking
}

// GetEnableTrunkingOk returns a tuple with the EnableTrunking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetEnableTrunkingOk() (*bool, bool) {
	if o == nil || o.EnableTrunking == nil {
		return nil, false
	}
	return o.EnableTrunking, true
}

// HasEnableTrunking returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasEnableTrunking() bool {
	if o != nil && o.EnableTrunking != nil {
		return true
	}

	return false
}

// SetEnableTrunking gets a reference to the given bool and assigns it to the EnableTrunking field.
func (o *FabricFcNetworkPolicy) SetEnableTrunking(v bool) {
	o.EnableTrunking = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FabricFcNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricFcNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *FabricFcNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FabricFcNetworkPolicy) GetProfiles() []FabricSwitchProfileRelationship {
	if o == nil {
		var ret []FabricSwitchProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FabricFcNetworkPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *FabricFcNetworkPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []FabricSwitchProfileRelationship and assigns it to the Profiles field.
func (o *FabricFcNetworkPolicy) SetProfiles(v []FabricSwitchProfileRelationship) {
	o.Profiles = v
}

func (o FabricFcNetworkPolicy) MarshalJSON() ([]byte, error) {
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
	if o.EnableTrunking != nil {
		toSerialize["EnableTrunking"] = o.EnableTrunking
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FabricFcNetworkPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type FabricFcNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enable or Disable Trunking on all of configured FC uplink ports.
		EnableTrunking *bool                                 `json:"EnableTrunking,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to fabricSwitchProfile resources.
		Profiles []FabricSwitchProfileRelationship `json:"Profiles,omitempty"`
	}

	varFabricFcNetworkPolicyWithoutEmbeddedStruct := FabricFcNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFabricFcNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varFabricFcNetworkPolicy := _FabricFcNetworkPolicy{}
		varFabricFcNetworkPolicy.ClassId = varFabricFcNetworkPolicyWithoutEmbeddedStruct.ClassId
		varFabricFcNetworkPolicy.ObjectType = varFabricFcNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varFabricFcNetworkPolicy.EnableTrunking = varFabricFcNetworkPolicyWithoutEmbeddedStruct.EnableTrunking
		varFabricFcNetworkPolicy.Organization = varFabricFcNetworkPolicyWithoutEmbeddedStruct.Organization
		varFabricFcNetworkPolicy.Profiles = varFabricFcNetworkPolicyWithoutEmbeddedStruct.Profiles
		*o = FabricFcNetworkPolicy(varFabricFcNetworkPolicy)
	} else {
		return err
	}

	varFabricFcNetworkPolicy := _FabricFcNetworkPolicy{}

	err = json.Unmarshal(bytes, &varFabricFcNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varFabricFcNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnableTrunking")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableFabricFcNetworkPolicy struct {
	value *FabricFcNetworkPolicy
	isSet bool
}

func (v NullableFabricFcNetworkPolicy) Get() *FabricFcNetworkPolicy {
	return v.value
}

func (v *NullableFabricFcNetworkPolicy) Set(val *FabricFcNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFcNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFcNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFcNetworkPolicy(val *FabricFcNetworkPolicy) *NullableFabricFcNetworkPolicy {
	return &NullableFabricFcNetworkPolicy{value: val, isSet: true}
}

func (v NullableFabricFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFcNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
