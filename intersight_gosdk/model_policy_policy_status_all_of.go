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

// PolicyPolicyStatusAllOf Definition of the list of properties defined in 'policy.PolicyStatus', excluding properties defined in parent classes.
type PolicyPolicyStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The object id of the policy being attached.
	Moid *string `json:"Moid,omitempty"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the policy attach/detach was successful or not. Values  -- ok, errored, validating. * `ok` - The policy attach/detach is successful. * `error` - The policy cannot be attached/detached due to an error. * `validating` - The policy preconfig validation is in progress.
	Status *string `json:"Status,omitempty"`
	// The object type of the policy being attached.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyPolicyStatusAllOf PolicyPolicyStatusAllOf

// NewPolicyPolicyStatusAllOf instantiates a new PolicyPolicyStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyPolicyStatusAllOf(classId string, objectType string) *PolicyPolicyStatusAllOf {
	this := PolicyPolicyStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewPolicyPolicyStatusAllOfWithDefaults instantiates a new PolicyPolicyStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyPolicyStatusAllOfWithDefaults() *PolicyPolicyStatusAllOf {
	this := PolicyPolicyStatusAllOf{}
	var classId string = "policy.PolicyStatus"
	this.ClassId = classId
	var objectType string = "policy.PolicyStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyPolicyStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyPolicyStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyPolicyStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyPolicyStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *PolicyPolicyStatusAllOf) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *PolicyPolicyStatusAllOf) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *PolicyPolicyStatusAllOf) SetMoid(v string) {
	o.Moid = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PolicyPolicyStatusAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PolicyPolicyStatusAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PolicyPolicyStatusAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PolicyPolicyStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PolicyPolicyStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PolicyPolicyStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PolicyPolicyStatusAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyPolicyStatusAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PolicyPolicyStatusAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PolicyPolicyStatusAllOf) SetType(v string) {
	o.Type = &v
}

func (o PolicyPolicyStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyPolicyStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyPolicyStatusAllOf := _PolicyPolicyStatusAllOf{}

	if err = json.Unmarshal(bytes, &varPolicyPolicyStatusAllOf); err == nil {
		*o = PolicyPolicyStatusAllOf(varPolicyPolicyStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyPolicyStatusAllOf struct {
	value *PolicyPolicyStatusAllOf
	isSet bool
}

func (v NullablePolicyPolicyStatusAllOf) Get() *PolicyPolicyStatusAllOf {
	return v.value
}

func (v *NullablePolicyPolicyStatusAllOf) Set(val *PolicyPolicyStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyPolicyStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyPolicyStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyPolicyStatusAllOf(val *PolicyPolicyStatusAllOf) *NullablePolicyPolicyStatusAllOf {
	return &NullablePolicyPolicyStatusAllOf{value: val, isSet: true}
}

func (v NullablePolicyPolicyStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyPolicyStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
