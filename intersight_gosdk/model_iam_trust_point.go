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

// checks if the IamTrustPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamTrustPoint{}

// IamTrustPoint To affirm the identity of trusted source. Allows import of third-party CA certificates in X.509 (CER) format. It can be a root CA or an trust chain that leads to a root CA.
type IamTrustPoint struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string            `json:"ObjectType"`
	Certificates []X509Certificate `json:"Certificates,omitempty"`
	// The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format.
	Chain   *string                        `json:"Chain,omitempty"`
	Account NullableIamAccountRelationship `json:"Account,omitempty"`
	// An array of relationships to moBaseMo resources.
	AssignedToEntity     []MoBaseMoRelationship `json:"AssignedToEntity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamTrustPoint IamTrustPoint

// NewIamTrustPoint instantiates a new IamTrustPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamTrustPoint(classId string, objectType string) *IamTrustPoint {
	this := IamTrustPoint{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamTrustPointWithDefaults instantiates a new IamTrustPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamTrustPointWithDefaults() *IamTrustPoint {
	this := IamTrustPoint{}
	var classId string = "iam.TrustPoint"
	this.ClassId = classId
	var objectType string = "iam.TrustPoint"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamTrustPoint) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamTrustPoint) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamTrustPoint) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.TrustPoint" of the ClassId field.
func (o *IamTrustPoint) GetDefaultClassId() interface{} {
	return "iam.TrustPoint"
}

// GetObjectType returns the ObjectType field value
func (o *IamTrustPoint) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamTrustPoint) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamTrustPoint) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.TrustPoint" of the ObjectType field.
func (o *IamTrustPoint) GetDefaultObjectType() interface{} {
	return "iam.TrustPoint"
}

// GetCertificates returns the Certificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamTrustPoint) GetCertificates() []X509Certificate {
	if o == nil {
		var ret []X509Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamTrustPoint) GetCertificatesOk() ([]X509Certificate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *IamTrustPoint) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []X509Certificate and assigns it to the Certificates field.
func (o *IamTrustPoint) SetCertificates(v []X509Certificate) {
	o.Certificates = v
}

// GetChain returns the Chain field value if set, zero value otherwise.
func (o *IamTrustPoint) GetChain() string {
	if o == nil || IsNil(o.Chain) {
		var ret string
		return ret
	}
	return *o.Chain
}

// GetChainOk returns a tuple with the Chain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamTrustPoint) GetChainOk() (*string, bool) {
	if o == nil || IsNil(o.Chain) {
		return nil, false
	}
	return o.Chain, true
}

// HasChain returns a boolean if a field has been set.
func (o *IamTrustPoint) HasChain() bool {
	if o != nil && !IsNil(o.Chain) {
		return true
	}

	return false
}

// SetChain gets a reference to the given string and assigns it to the Chain field.
func (o *IamTrustPoint) SetChain(v string) {
	o.Chain = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamTrustPoint) GetAccount() IamAccountRelationship {
	if o == nil || IsNil(o.Account.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamTrustPoint) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *IamTrustPoint) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the Account field.
func (o *IamTrustPoint) SetAccount(v IamAccountRelationship) {
	o.Account.Set(&v)
}

// SetAccountNil sets the value for Account to be an explicit nil
func (o *IamTrustPoint) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *IamTrustPoint) UnsetAccount() {
	o.Account.Unset()
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamTrustPoint) GetAssignedToEntity() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamTrustPoint) GetAssignedToEntityOk() ([]MoBaseMoRelationship, bool) {
	if o == nil || IsNil(o.AssignedToEntity) {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *IamTrustPoint) HasAssignedToEntity() bool {
	if o != nil && !IsNil(o.AssignedToEntity) {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given []MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *IamTrustPoint) SetAssignedToEntity(v []MoBaseMoRelationship) {
	o.AssignedToEntity = v
}

func (o IamTrustPoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamTrustPoint) ToMap() (map[string]interface{}, error) {
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
	if o.Certificates != nil {
		toSerialize["Certificates"] = o.Certificates
	}
	if !IsNil(o.Chain) {
		toSerialize["Chain"] = o.Chain
	}
	if o.Account.IsSet() {
		toSerialize["Account"] = o.Account.Get()
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamTrustPoint) UnmarshalJSON(data []byte) (err error) {
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
	type IamTrustPointWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string            `json:"ObjectType"`
		Certificates []X509Certificate `json:"Certificates,omitempty"`
		// The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format.
		Chain   *string                        `json:"Chain,omitempty"`
		Account NullableIamAccountRelationship `json:"Account,omitempty"`
		// An array of relationships to moBaseMo resources.
		AssignedToEntity []MoBaseMoRelationship `json:"AssignedToEntity,omitempty"`
	}

	varIamTrustPointWithoutEmbeddedStruct := IamTrustPointWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamTrustPointWithoutEmbeddedStruct)
	if err == nil {
		varIamTrustPoint := _IamTrustPoint{}
		varIamTrustPoint.ClassId = varIamTrustPointWithoutEmbeddedStruct.ClassId
		varIamTrustPoint.ObjectType = varIamTrustPointWithoutEmbeddedStruct.ObjectType
		varIamTrustPoint.Certificates = varIamTrustPointWithoutEmbeddedStruct.Certificates
		varIamTrustPoint.Chain = varIamTrustPointWithoutEmbeddedStruct.Chain
		varIamTrustPoint.Account = varIamTrustPointWithoutEmbeddedStruct.Account
		varIamTrustPoint.AssignedToEntity = varIamTrustPointWithoutEmbeddedStruct.AssignedToEntity
		*o = IamTrustPoint(varIamTrustPoint)
	} else {
		return err
	}

	varIamTrustPoint := _IamTrustPoint{}

	err = json.Unmarshal(data, &varIamTrustPoint)
	if err == nil {
		o.MoBaseMo = varIamTrustPoint.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Certificates")
		delete(additionalProperties, "Chain")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "AssignedToEntity")

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

type NullableIamTrustPoint struct {
	value *IamTrustPoint
	isSet bool
}

func (v NullableIamTrustPoint) Get() *IamTrustPoint {
	return v.value
}

func (v *NullableIamTrustPoint) Set(val *IamTrustPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIamTrustPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIamTrustPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamTrustPoint(val *IamTrustPoint) *NullableIamTrustPoint {
	return &NullableIamTrustPoint{value: val, isSet: true}
}

func (v NullableIamTrustPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamTrustPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
