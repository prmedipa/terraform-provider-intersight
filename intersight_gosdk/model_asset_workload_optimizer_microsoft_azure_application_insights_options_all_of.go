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

// AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf Definition of the list of properties defined in 'asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions', excluding properties defined in parent classes.
type AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enrollment number of this Azure account.
	EnrollmentNumber *string `json:"EnrollmentNumber,omitempty"`
	// The offer ID of this account. Default value is \"MS-AZR-0003P\", a pay-as-you-go subscription lets you pay for the services and resources that you use on a monthly basis.
	OfferId *string `json:"OfferId,omitempty"`
	// The Azure Subscription ID.
	SubscriptionId *string `json:"SubscriptionId,omitempty"`
	// Tenant ID associated with Azure Account.
	TenantId             *string `json:"TenantId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf

// NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOfWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOfWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf {
	this := AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureApplicationInsightsOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnrollmentNumber returns the EnrollmentNumber field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetEnrollmentNumber() string {
	if o == nil || o.EnrollmentNumber == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentNumber
}

// GetEnrollmentNumberOk returns a tuple with the EnrollmentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetEnrollmentNumberOk() (*string, bool) {
	if o == nil || o.EnrollmentNumber == nil {
		return nil, false
	}
	return o.EnrollmentNumber, true
}

// HasEnrollmentNumber returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) HasEnrollmentNumber() bool {
	if o != nil && o.EnrollmentNumber != nil {
		return true
	}

	return false
}

// SetEnrollmentNumber gets a reference to the given string and assigns it to the EnrollmentNumber field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetEnrollmentNumber(v string) {
	o.EnrollmentNumber = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetOfferId() string {
	if o == nil || o.OfferId == nil {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetOfferIdOk() (*string, bool) {
	if o == nil || o.OfferId == nil {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) HasOfferId() bool {
	if o != nil && o.OfferId != nil {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetOfferId(v string) {
	o.OfferId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnrollmentNumber != nil {
		toSerialize["EnrollmentNumber"] = o.EnrollmentNumber
	}
	if o.OfferId != nil {
		toSerialize["OfferId"] = o.OfferId
	}
	if o.SubscriptionId != nil {
		toSerialize["SubscriptionId"] = o.SubscriptionId
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf := _AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf{}

	if err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf); err == nil {
		*o = AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf(varAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnrollmentNumber")
		delete(additionalProperties, "OfferId")
		delete(additionalProperties, "SubscriptionId")
		delete(additionalProperties, "TenantId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf struct {
	value *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) Get() *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) Set(val *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf(val *AssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureApplicationInsightsOptionsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
