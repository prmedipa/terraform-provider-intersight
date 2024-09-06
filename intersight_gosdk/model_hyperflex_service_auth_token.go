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

// checks if the HyperflexServiceAuthToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperflexServiceAuthToken{}

// HyperflexServiceAuthToken A Service auth token entity that represents HyperFlex Data Platform service AAA token.
type HyperflexServiceAuthToken struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Client Id or tenant Id of the entity that uses the service auth token.
	ClientId *string `json:"ClientId,omitempty"`
	// Version of Container Storage Interface (CSI) that the tokenOwner is associated with.
	CsiVersion *string `json:"CsiVersion,omitempty"`
	// Service auth token that has been created by HyperFlex cluster.
	ServiceAuthToken *string `json:"ServiceAuthToken,omitempty"`
	// Represents status of ervice auth claim or revocation. * `Unknown` - Unknown claim state of the service auth token. * `Claiming` - The service auth token claim is in progress. * `Claimed` - The service auth token has been successfully claimed. * `FailedToClaim` - Cannot claim the service auth token on the underlying HyperFlex cluster. * `Revoking` - The service auth token revocation is in progress. * `Revoked` - The service auth token revocation has been successfully revoked. * `FailedToRevoke` - Cannot revoke the service auth token on the underlying HyperFlex cluster.
	Status               *string                                      `json:"Status,omitempty"`
	Cluster              NullableHyperflexClusterRelationship         `json:"Cluster,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	TokenOwner           NullableMoBaseMoRelationship                 `json:"TokenOwner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServiceAuthToken HyperflexServiceAuthToken

// NewHyperflexServiceAuthToken instantiates a new HyperflexServiceAuthToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServiceAuthToken(classId string, objectType string) *HyperflexServiceAuthToken {
	this := HyperflexServiceAuthToken{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexServiceAuthTokenWithDefaults instantiates a new HyperflexServiceAuthToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServiceAuthTokenWithDefaults() *HyperflexServiceAuthToken {
	this := HyperflexServiceAuthToken{}
	var classId string = "hyperflex.ServiceAuthToken"
	this.ClassId = classId
	var objectType string = "hyperflex.ServiceAuthToken"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServiceAuthToken) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServiceAuthToken) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "hyperflex.ServiceAuthToken" of the ClassId field.
func (o *HyperflexServiceAuthToken) GetDefaultClassId() interface{} {
	return "hyperflex.ServiceAuthToken"
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServiceAuthToken) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServiceAuthToken) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "hyperflex.ServiceAuthToken" of the ObjectType field.
func (o *HyperflexServiceAuthToken) GetDefaultObjectType() interface{} {
	return "hyperflex.ServiceAuthToken"
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *HyperflexServiceAuthToken) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *HyperflexServiceAuthToken) SetClientId(v string) {
	o.ClientId = &v
}

// GetCsiVersion returns the CsiVersion field value if set, zero value otherwise.
func (o *HyperflexServiceAuthToken) GetCsiVersion() string {
	if o == nil || IsNil(o.CsiVersion) {
		var ret string
		return ret
	}
	return *o.CsiVersion
}

// GetCsiVersionOk returns a tuple with the CsiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetCsiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.CsiVersion) {
		return nil, false
	}
	return o.CsiVersion, true
}

// HasCsiVersion returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasCsiVersion() bool {
	if o != nil && !IsNil(o.CsiVersion) {
		return true
	}

	return false
}

// SetCsiVersion gets a reference to the given string and assigns it to the CsiVersion field.
func (o *HyperflexServiceAuthToken) SetCsiVersion(v string) {
	o.CsiVersion = &v
}

// GetServiceAuthToken returns the ServiceAuthToken field value if set, zero value otherwise.
func (o *HyperflexServiceAuthToken) GetServiceAuthToken() string {
	if o == nil || IsNil(o.ServiceAuthToken) {
		var ret string
		return ret
	}
	return *o.ServiceAuthToken
}

// GetServiceAuthTokenOk returns a tuple with the ServiceAuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetServiceAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAuthToken) {
		return nil, false
	}
	return o.ServiceAuthToken, true
}

// HasServiceAuthToken returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasServiceAuthToken() bool {
	if o != nil && !IsNil(o.ServiceAuthToken) {
		return true
	}

	return false
}

// SetServiceAuthToken gets a reference to the given string and assigns it to the ServiceAuthToken field.
func (o *HyperflexServiceAuthToken) SetServiceAuthToken(v string) {
	o.ServiceAuthToken = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HyperflexServiceAuthToken) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServiceAuthToken) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HyperflexServiceAuthToken) SetStatus(v string) {
	o.Status = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServiceAuthToken) GetCluster() HyperflexClusterRelationship {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret HyperflexClusterRelationship
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServiceAuthToken) GetClusterOk() (*HyperflexClusterRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableHyperflexClusterRelationship and assigns it to the Cluster field.
func (o *HyperflexServiceAuthToken) SetCluster(v HyperflexClusterRelationship) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *HyperflexServiceAuthToken) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *HyperflexServiceAuthToken) UnsetCluster() {
	o.Cluster.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServiceAuthToken) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServiceAuthToken) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexServiceAuthToken) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *HyperflexServiceAuthToken) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *HyperflexServiceAuthToken) UnsetOrganization() {
	o.Organization.Unset()
}

// GetTokenOwner returns the TokenOwner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexServiceAuthToken) GetTokenOwner() MoBaseMoRelationship {
	if o == nil || IsNil(o.TokenOwner.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TokenOwner.Get()
}

// GetTokenOwnerOk returns a tuple with the TokenOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexServiceAuthToken) GetTokenOwnerOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenOwner.Get(), o.TokenOwner.IsSet()
}

// HasTokenOwner returns a boolean if a field has been set.
func (o *HyperflexServiceAuthToken) HasTokenOwner() bool {
	if o != nil && o.TokenOwner.IsSet() {
		return true
	}

	return false
}

// SetTokenOwner gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TokenOwner field.
func (o *HyperflexServiceAuthToken) SetTokenOwner(v MoBaseMoRelationship) {
	o.TokenOwner.Set(&v)
}

// SetTokenOwnerNil sets the value for TokenOwner to be an explicit nil
func (o *HyperflexServiceAuthToken) SetTokenOwnerNil() {
	o.TokenOwner.Set(nil)
}

// UnsetTokenOwner ensures that no value is present for TokenOwner, not even an explicit nil
func (o *HyperflexServiceAuthToken) UnsetTokenOwner() {
	o.TokenOwner.Unset()
}

func (o HyperflexServiceAuthToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperflexServiceAuthToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ClientId) {
		toSerialize["ClientId"] = o.ClientId
	}
	if !IsNil(o.CsiVersion) {
		toSerialize["CsiVersion"] = o.CsiVersion
	}
	if !IsNil(o.ServiceAuthToken) {
		toSerialize["ServiceAuthToken"] = o.ServiceAuthToken
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.Cluster.IsSet() {
		toSerialize["Cluster"] = o.Cluster.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.TokenOwner.IsSet() {
		toSerialize["TokenOwner"] = o.TokenOwner.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HyperflexServiceAuthToken) UnmarshalJSON(data []byte) (err error) {
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
	type HyperflexServiceAuthTokenWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Client Id or tenant Id of the entity that uses the service auth token.
		ClientId *string `json:"ClientId,omitempty"`
		// Version of Container Storage Interface (CSI) that the tokenOwner is associated with.
		CsiVersion *string `json:"CsiVersion,omitempty"`
		// Service auth token that has been created by HyperFlex cluster.
		ServiceAuthToken *string `json:"ServiceAuthToken,omitempty"`
		// Represents status of ervice auth claim or revocation. * `Unknown` - Unknown claim state of the service auth token. * `Claiming` - The service auth token claim is in progress. * `Claimed` - The service auth token has been successfully claimed. * `FailedToClaim` - Cannot claim the service auth token on the underlying HyperFlex cluster. * `Revoking` - The service auth token revocation is in progress. * `Revoked` - The service auth token revocation has been successfully revoked. * `FailedToRevoke` - Cannot revoke the service auth token on the underlying HyperFlex cluster.
		Status       *string                                      `json:"Status,omitempty"`
		Cluster      NullableHyperflexClusterRelationship         `json:"Cluster,omitempty"`
		Organization NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		TokenOwner   NullableMoBaseMoRelationship                 `json:"TokenOwner,omitempty"`
	}

	varHyperflexServiceAuthTokenWithoutEmbeddedStruct := HyperflexServiceAuthTokenWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varHyperflexServiceAuthTokenWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexServiceAuthToken := _HyperflexServiceAuthToken{}
		varHyperflexServiceAuthToken.ClassId = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.ClassId
		varHyperflexServiceAuthToken.ObjectType = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.ObjectType
		varHyperflexServiceAuthToken.ClientId = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.ClientId
		varHyperflexServiceAuthToken.CsiVersion = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.CsiVersion
		varHyperflexServiceAuthToken.ServiceAuthToken = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.ServiceAuthToken
		varHyperflexServiceAuthToken.Status = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.Status
		varHyperflexServiceAuthToken.Cluster = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.Cluster
		varHyperflexServiceAuthToken.Organization = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.Organization
		varHyperflexServiceAuthToken.TokenOwner = varHyperflexServiceAuthTokenWithoutEmbeddedStruct.TokenOwner
		*o = HyperflexServiceAuthToken(varHyperflexServiceAuthToken)
	} else {
		return err
	}

	varHyperflexServiceAuthToken := _HyperflexServiceAuthToken{}

	err = json.Unmarshal(data, &varHyperflexServiceAuthToken)
	if err == nil {
		o.MoBaseMo = varHyperflexServiceAuthToken.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientId")
		delete(additionalProperties, "CsiVersion")
		delete(additionalProperties, "ServiceAuthToken")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Cluster")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "TokenOwner")

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

type NullableHyperflexServiceAuthToken struct {
	value *HyperflexServiceAuthToken
	isSet bool
}

func (v NullableHyperflexServiceAuthToken) Get() *HyperflexServiceAuthToken {
	return v.value
}

func (v *NullableHyperflexServiceAuthToken) Set(val *HyperflexServiceAuthToken) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServiceAuthToken) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServiceAuthToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServiceAuthToken(val *HyperflexServiceAuthToken) *NullableHyperflexServiceAuthToken {
	return &NullableHyperflexServiceAuthToken{value: val, isSet: true}
}

func (v NullableHyperflexServiceAuthToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServiceAuthToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
