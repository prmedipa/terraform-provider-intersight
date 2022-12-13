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

// MemoryPersistentMemoryPolicy The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.
type MemoryPersistentMemoryPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string                                      `json:"ObjectType"`
	Goals             []MemoryPersistentMemoryGoal                `json:"Goals,omitempty"`
	LocalSecurity     NullableMemoryPersistentMemoryLocalSecurity `json:"LocalSecurity,omitempty"`
	LogicalNamespaces []MemoryPersistentMemoryLogicalNamespace    `json:"LogicalNamespaces,omitempty"`
	// Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System. * `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy. * `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools.
	ManagementMode *string `json:"ManagementMode,omitempty"`
	// Persistent Memory Namespaces to be retained or not.
	RetainNamespaces *bool                                 `json:"RetainNamespaces,omitempty"`
	Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryPersistentMemoryPolicy MemoryPersistentMemoryPolicy

// NewMemoryPersistentMemoryPolicy instantiates a new MemoryPersistentMemoryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryPersistentMemoryPolicy(classId string, objectType string) *MemoryPersistentMemoryPolicy {
	this := MemoryPersistentMemoryPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var managementMode string = "configured-from-intersight"
	this.ManagementMode = &managementMode
	var retainNamespaces bool = true
	this.RetainNamespaces = &retainNamespaces
	return &this
}

// NewMemoryPersistentMemoryPolicyWithDefaults instantiates a new MemoryPersistentMemoryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryPersistentMemoryPolicyWithDefaults() *MemoryPersistentMemoryPolicy {
	this := MemoryPersistentMemoryPolicy{}
	var classId string = "memory.PersistentMemoryPolicy"
	this.ClassId = classId
	var objectType string = "memory.PersistentMemoryPolicy"
	this.ObjectType = objectType
	var managementMode string = "configured-from-intersight"
	this.ManagementMode = &managementMode
	var retainNamespaces bool = true
	this.RetainNamespaces = &retainNamespaces
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryPersistentMemoryPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryPersistentMemoryPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryPersistentMemoryPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryPersistentMemoryPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGoals returns the Goals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryPolicy) GetGoals() []MemoryPersistentMemoryGoal {
	if o == nil {
		var ret []MemoryPersistentMemoryGoal
		return ret
	}
	return o.Goals
}

// GetGoalsOk returns a tuple with the Goals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryPolicy) GetGoalsOk() ([]MemoryPersistentMemoryGoal, bool) {
	if o == nil || o.Goals == nil {
		return nil, false
	}
	return o.Goals, true
}

// HasGoals returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasGoals() bool {
	if o != nil && o.Goals != nil {
		return true
	}

	return false
}

// SetGoals gets a reference to the given []MemoryPersistentMemoryGoal and assigns it to the Goals field.
func (o *MemoryPersistentMemoryPolicy) SetGoals(v []MemoryPersistentMemoryGoal) {
	o.Goals = v
}

// GetLocalSecurity returns the LocalSecurity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryPolicy) GetLocalSecurity() MemoryPersistentMemoryLocalSecurity {
	if o == nil || o.LocalSecurity.Get() == nil {
		var ret MemoryPersistentMemoryLocalSecurity
		return ret
	}
	return *o.LocalSecurity.Get()
}

// GetLocalSecurityOk returns a tuple with the LocalSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryPolicy) GetLocalSecurityOk() (*MemoryPersistentMemoryLocalSecurity, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalSecurity.Get(), o.LocalSecurity.IsSet()
}

// HasLocalSecurity returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasLocalSecurity() bool {
	if o != nil && o.LocalSecurity.IsSet() {
		return true
	}

	return false
}

// SetLocalSecurity gets a reference to the given NullableMemoryPersistentMemoryLocalSecurity and assigns it to the LocalSecurity field.
func (o *MemoryPersistentMemoryPolicy) SetLocalSecurity(v MemoryPersistentMemoryLocalSecurity) {
	o.LocalSecurity.Set(&v)
}

// SetLocalSecurityNil sets the value for LocalSecurity to be an explicit nil
func (o *MemoryPersistentMemoryPolicy) SetLocalSecurityNil() {
	o.LocalSecurity.Set(nil)
}

// UnsetLocalSecurity ensures that no value is present for LocalSecurity, not even an explicit nil
func (o *MemoryPersistentMemoryPolicy) UnsetLocalSecurity() {
	o.LocalSecurity.Unset()
}

// GetLogicalNamespaces returns the LogicalNamespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryPolicy) GetLogicalNamespaces() []MemoryPersistentMemoryLogicalNamespace {
	if o == nil {
		var ret []MemoryPersistentMemoryLogicalNamespace
		return ret
	}
	return o.LogicalNamespaces
}

// GetLogicalNamespacesOk returns a tuple with the LogicalNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryPolicy) GetLogicalNamespacesOk() ([]MemoryPersistentMemoryLogicalNamespace, bool) {
	if o == nil || o.LogicalNamespaces == nil {
		return nil, false
	}
	return o.LogicalNamespaces, true
}

// HasLogicalNamespaces returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasLogicalNamespaces() bool {
	if o != nil && o.LogicalNamespaces != nil {
		return true
	}

	return false
}

// SetLogicalNamespaces gets a reference to the given []MemoryPersistentMemoryLogicalNamespace and assigns it to the LogicalNamespaces field.
func (o *MemoryPersistentMemoryPolicy) SetLogicalNamespaces(v []MemoryPersistentMemoryLogicalNamespace) {
	o.LogicalNamespaces = v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryPolicy) GetManagementMode() string {
	if o == nil || o.ManagementMode == nil {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryPolicy) GetManagementModeOk() (*string, bool) {
	if o == nil || o.ManagementMode == nil {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasManagementMode() bool {
	if o != nil && o.ManagementMode != nil {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *MemoryPersistentMemoryPolicy) SetManagementMode(v string) {
	o.ManagementMode = &v
}

// GetRetainNamespaces returns the RetainNamespaces field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryPolicy) GetRetainNamespaces() bool {
	if o == nil || o.RetainNamespaces == nil {
		var ret bool
		return ret
	}
	return *o.RetainNamespaces
}

// GetRetainNamespacesOk returns a tuple with the RetainNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryPolicy) GetRetainNamespacesOk() (*bool, bool) {
	if o == nil || o.RetainNamespaces == nil {
		return nil, false
	}
	return o.RetainNamespaces, true
}

// HasRetainNamespaces returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasRetainNamespaces() bool {
	if o != nil && o.RetainNamespaces != nil {
		return true
	}

	return false
}

// SetRetainNamespaces gets a reference to the given bool and assigns it to the RetainNamespaces field.
func (o *MemoryPersistentMemoryPolicy) SetRetainNamespaces(v bool) {
	o.RetainNamespaces = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *MemoryPersistentMemoryPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryPersistentMemoryPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *MemoryPersistentMemoryPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryPersistentMemoryPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryPersistentMemoryPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *MemoryPersistentMemoryPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *MemoryPersistentMemoryPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o MemoryPersistentMemoryPolicy) MarshalJSON() ([]byte, error) {
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
	if o.Goals != nil {
		toSerialize["Goals"] = o.Goals
	}
	if o.LocalSecurity.IsSet() {
		toSerialize["LocalSecurity"] = o.LocalSecurity.Get()
	}
	if o.LogicalNamespaces != nil {
		toSerialize["LogicalNamespaces"] = o.LogicalNamespaces
	}
	if o.ManagementMode != nil {
		toSerialize["ManagementMode"] = o.ManagementMode
	}
	if o.RetainNamespaces != nil {
		toSerialize["RetainNamespaces"] = o.RetainNamespaces
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

func (o *MemoryPersistentMemoryPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryPersistentMemoryPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                                      `json:"ObjectType"`
		Goals             []MemoryPersistentMemoryGoal                `json:"Goals,omitempty"`
		LocalSecurity     NullableMemoryPersistentMemoryLocalSecurity `json:"LocalSecurity,omitempty"`
		LogicalNamespaces []MemoryPersistentMemoryLogicalNamespace    `json:"LogicalNamespaces,omitempty"`
		// Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System. * `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy. * `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools.
		ManagementMode *string `json:"ManagementMode,omitempty"`
		// Persistent Memory Namespaces to be retained or not.
		RetainNamespaces *bool                                 `json:"RetainNamespaces,omitempty"`
		Organization     *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct := MemoryPersistentMemoryPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct)
	if err == nil {
		varMemoryPersistentMemoryPolicy := _MemoryPersistentMemoryPolicy{}
		varMemoryPersistentMemoryPolicy.ClassId = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.ClassId
		varMemoryPersistentMemoryPolicy.ObjectType = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.ObjectType
		varMemoryPersistentMemoryPolicy.Goals = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.Goals
		varMemoryPersistentMemoryPolicy.LocalSecurity = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.LocalSecurity
		varMemoryPersistentMemoryPolicy.LogicalNamespaces = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.LogicalNamespaces
		varMemoryPersistentMemoryPolicy.ManagementMode = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.ManagementMode
		varMemoryPersistentMemoryPolicy.RetainNamespaces = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.RetainNamespaces
		varMemoryPersistentMemoryPolicy.Organization = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.Organization
		varMemoryPersistentMemoryPolicy.Profiles = varMemoryPersistentMemoryPolicyWithoutEmbeddedStruct.Profiles
		*o = MemoryPersistentMemoryPolicy(varMemoryPersistentMemoryPolicy)
	} else {
		return err
	}

	varMemoryPersistentMemoryPolicy := _MemoryPersistentMemoryPolicy{}

	err = json.Unmarshal(bytes, &varMemoryPersistentMemoryPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varMemoryPersistentMemoryPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Goals")
		delete(additionalProperties, "LocalSecurity")
		delete(additionalProperties, "LogicalNamespaces")
		delete(additionalProperties, "ManagementMode")
		delete(additionalProperties, "RetainNamespaces")
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

type NullableMemoryPersistentMemoryPolicy struct {
	value *MemoryPersistentMemoryPolicy
	isSet bool
}

func (v NullableMemoryPersistentMemoryPolicy) Get() *MemoryPersistentMemoryPolicy {
	return v.value
}

func (v *NullableMemoryPersistentMemoryPolicy) Set(val *MemoryPersistentMemoryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryPersistentMemoryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryPersistentMemoryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryPersistentMemoryPolicy(val *MemoryPersistentMemoryPolicy) *NullableMemoryPersistentMemoryPolicy {
	return &NullableMemoryPersistentMemoryPolicy{value: val, isSet: true}
}

func (v NullableMemoryPersistentMemoryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryPersistentMemoryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
