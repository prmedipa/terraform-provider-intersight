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
	"time"
)

// checks if the X509Certificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &X509Certificate{}

// X509Certificate The representation of an X.509 certificate.
type X509Certificate struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                        `json:"ObjectType"`
	Issuer     NullablePkixDistinguishedName `json:"Issuer,omitempty"`
	// The date on which the certificate's validity period ends.
	NotAfter *time.Time `json:"NotAfter,omitempty"`
	// The date on which the certificate's validity period begins.
	NotBefore *time.Time `json:"NotBefore,omitempty"`
	// The base64 encoded certificate in PEM format.
	PemCertificate *string `json:"PemCertificate,omitempty"`
	// The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.
	Sha256Fingerprint *string `json:"Sha256Fingerprint,omitempty"`
	// Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).
	SignatureAlgorithm   *string                       `json:"SignatureAlgorithm,omitempty"`
	Subject              NullablePkixDistinguishedName `json:"Subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _X509Certificate X509Certificate

// NewX509Certificate instantiates a new X509Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate(classId string, objectType string) *X509Certificate {
	this := X509Certificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewX509CertificateWithDefaults instantiates a new X509Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateWithDefaults() *X509Certificate {
	this := X509Certificate{}
	var classId string = "x509.Certificate"
	this.ClassId = classId
	var objectType string = "x509.Certificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *X509Certificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *X509Certificate) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "x509.Certificate" of the ClassId field.
func (o *X509Certificate) GetDefaultClassId() interface{} {
	return "x509.Certificate"
}

// GetObjectType returns the ObjectType field value
func (o *X509Certificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *X509Certificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "x509.Certificate" of the ObjectType field.
func (o *X509Certificate) GetDefaultObjectType() interface{} {
	return "x509.Certificate"
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate) GetIssuer() PkixDistinguishedName {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate) GetIssuerOk() (*PkixDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullablePkixDistinguishedName and assigns it to the Issuer field.
func (o *X509Certificate) SetIssuer(v PkixDistinguishedName) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *X509Certificate) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *X509Certificate) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509Certificate) GetNotAfter() time.Time {
	if o == nil || IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509Certificate) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509Certificate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509Certificate) GetNotBefore() time.Time {
	if o == nil || IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509Certificate) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509Certificate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetPemCertificate returns the PemCertificate field value if set, zero value otherwise.
func (o *X509Certificate) GetPemCertificate() string {
	if o == nil || IsNil(o.PemCertificate) {
		var ret string
		return ret
	}
	return *o.PemCertificate
}

// GetPemCertificateOk returns a tuple with the PemCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetPemCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.PemCertificate) {
		return nil, false
	}
	return o.PemCertificate, true
}

// HasPemCertificate returns a boolean if a field has been set.
func (o *X509Certificate) HasPemCertificate() bool {
	if o != nil && !IsNil(o.PemCertificate) {
		return true
	}

	return false
}

// SetPemCertificate gets a reference to the given string and assigns it to the PemCertificate field.
func (o *X509Certificate) SetPemCertificate(v string) {
	o.PemCertificate = &v
}

// GetSha256Fingerprint returns the Sha256Fingerprint field value if set, zero value otherwise.
func (o *X509Certificate) GetSha256Fingerprint() string {
	if o == nil || IsNil(o.Sha256Fingerprint) {
		var ret string
		return ret
	}
	return *o.Sha256Fingerprint
}

// GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSha256FingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Sha256Fingerprint) {
		return nil, false
	}
	return o.Sha256Fingerprint, true
}

// HasSha256Fingerprint returns a boolean if a field has been set.
func (o *X509Certificate) HasSha256Fingerprint() bool {
	if o != nil && !IsNil(o.Sha256Fingerprint) {
		return true
	}

	return false
}

// SetSha256Fingerprint gets a reference to the given string and assigns it to the Sha256Fingerprint field.
func (o *X509Certificate) SetSha256Fingerprint(v string) {
	o.Sha256Fingerprint = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *X509Certificate) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *X509Certificate) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *X509Certificate) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate) GetSubject() PkixDistinguishedName {
	if o == nil || IsNil(o.Subject.Get()) {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate) GetSubjectOk() (*PkixDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *X509Certificate) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullablePkixDistinguishedName and assigns it to the Subject field.
func (o *X509Certificate) SetSubject(v PkixDistinguishedName) {
	o.Subject.Set(&v)
}

// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *X509Certificate) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *X509Certificate) UnsetSubject() {
	o.Subject.Unset()
}

func (o X509Certificate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o X509Certificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return map[string]interface{}{}, errMoBaseComplexType
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if o.Issuer.IsSet() {
		toSerialize["Issuer"] = o.Issuer.Get()
	}
	if !IsNil(o.NotAfter) {
		toSerialize["NotAfter"] = o.NotAfter
	}
	if !IsNil(o.NotBefore) {
		toSerialize["NotBefore"] = o.NotBefore
	}
	if !IsNil(o.PemCertificate) {
		toSerialize["PemCertificate"] = o.PemCertificate
	}
	if !IsNil(o.Sha256Fingerprint) {
		toSerialize["Sha256Fingerprint"] = o.Sha256Fingerprint
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["SignatureAlgorithm"] = o.SignatureAlgorithm
	}
	if o.Subject.IsSet() {
		toSerialize["Subject"] = o.Subject.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *X509Certificate) UnmarshalJSON(data []byte) (err error) {
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
	type X509CertificateWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string                        `json:"ObjectType"`
		Issuer     NullablePkixDistinguishedName `json:"Issuer,omitempty"`
		// The date on which the certificate's validity period ends.
		NotAfter *time.Time `json:"NotAfter,omitempty"`
		// The date on which the certificate's validity period begins.
		NotBefore *time.Time `json:"NotBefore,omitempty"`
		// The base64 encoded certificate in PEM format.
		PemCertificate *string `json:"PemCertificate,omitempty"`
		// The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.
		Sha256Fingerprint *string `json:"Sha256Fingerprint,omitempty"`
		// Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).
		SignatureAlgorithm *string                       `json:"SignatureAlgorithm,omitempty"`
		Subject            NullablePkixDistinguishedName `json:"Subject,omitempty"`
	}

	varX509CertificateWithoutEmbeddedStruct := X509CertificateWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varX509CertificateWithoutEmbeddedStruct)
	if err == nil {
		varX509Certificate := _X509Certificate{}
		varX509Certificate.ClassId = varX509CertificateWithoutEmbeddedStruct.ClassId
		varX509Certificate.ObjectType = varX509CertificateWithoutEmbeddedStruct.ObjectType
		varX509Certificate.Issuer = varX509CertificateWithoutEmbeddedStruct.Issuer
		varX509Certificate.NotAfter = varX509CertificateWithoutEmbeddedStruct.NotAfter
		varX509Certificate.NotBefore = varX509CertificateWithoutEmbeddedStruct.NotBefore
		varX509Certificate.PemCertificate = varX509CertificateWithoutEmbeddedStruct.PemCertificate
		varX509Certificate.Sha256Fingerprint = varX509CertificateWithoutEmbeddedStruct.Sha256Fingerprint
		varX509Certificate.SignatureAlgorithm = varX509CertificateWithoutEmbeddedStruct.SignatureAlgorithm
		varX509Certificate.Subject = varX509CertificateWithoutEmbeddedStruct.Subject
		*o = X509Certificate(varX509Certificate)
	} else {
		return err
	}

	varX509Certificate := _X509Certificate{}

	err = json.Unmarshal(data, &varX509Certificate)
	if err == nil {
		o.MoBaseComplexType = varX509Certificate.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Issuer")
		delete(additionalProperties, "NotAfter")
		delete(additionalProperties, "NotBefore")
		delete(additionalProperties, "PemCertificate")
		delete(additionalProperties, "Sha256Fingerprint")
		delete(additionalProperties, "SignatureAlgorithm")
		delete(additionalProperties, "Subject")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableX509Certificate struct {
	value *X509Certificate
	isSet bool
}

func (v NullableX509Certificate) Get() *X509Certificate {
	return v.value
}

func (v *NullableX509Certificate) Set(val *X509Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableX509Certificate) IsSet() bool {
	return v.isSet
}

func (v *NullableX509Certificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509Certificate(val *X509Certificate) *NullableX509Certificate {
	return &NullableX509Certificate{value: val, isSet: true}
}

func (v NullableX509Certificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509Certificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
