/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-9783
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// IssueOdataConditionAllOf Definition of the list of properties defined in 'issue.OdataCondition', excluding properties defined in parent classes.
type IssueOdataConditionAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string           `json:"ObjectType"`
	DeviceTags []IssueDeviceTag `json:"DeviceTags,omitempty"`
	// The Intersight object type on which the condition is to be applied. The object type may be either a concrete type such as compute.RackUnit or a shared parent type such as compute.Physical.
	Motype *string `json:"Motype,omitempty"`
	// An Odata query string interpreted to be the value portion of a $filter expression. For example, for the function query $filter=Serial eq 'ABC', the 'query' field should be assigned the string Serial eq 'ABC.
	Query                *string `json:"Query,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IssueOdataConditionAllOf IssueOdataConditionAllOf

// NewIssueOdataConditionAllOf instantiates a new IssueOdataConditionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueOdataConditionAllOf(classId string, objectType string) *IssueOdataConditionAllOf {
	this := IssueOdataConditionAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIssueOdataConditionAllOfWithDefaults instantiates a new IssueOdataConditionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueOdataConditionAllOfWithDefaults() *IssueOdataConditionAllOf {
	this := IssueOdataConditionAllOf{}
	var classId string = "issue.OdataCondition"
	this.ClassId = classId
	var objectType string = "issue.OdataCondition"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IssueOdataConditionAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IssueOdataConditionAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IssueOdataConditionAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IssueOdataConditionAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IssueOdataConditionAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IssueOdataConditionAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceTags returns the DeviceTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueOdataConditionAllOf) GetDeviceTags() []IssueDeviceTag {
	if o == nil {
		var ret []IssueDeviceTag
		return ret
	}
	return o.DeviceTags
}

// GetDeviceTagsOk returns a tuple with the DeviceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueOdataConditionAllOf) GetDeviceTagsOk() ([]IssueDeviceTag, bool) {
	if o == nil || o.DeviceTags == nil {
		return nil, false
	}
	return o.DeviceTags, true
}

// HasDeviceTags returns a boolean if a field has been set.
func (o *IssueOdataConditionAllOf) HasDeviceTags() bool {
	if o != nil && o.DeviceTags != nil {
		return true
	}

	return false
}

// SetDeviceTags gets a reference to the given []IssueDeviceTag and assigns it to the DeviceTags field.
func (o *IssueOdataConditionAllOf) SetDeviceTags(v []IssueDeviceTag) {
	o.DeviceTags = v
}

// GetMotype returns the Motype field value if set, zero value otherwise.
func (o *IssueOdataConditionAllOf) GetMotype() string {
	if o == nil || o.Motype == nil {
		var ret string
		return ret
	}
	return *o.Motype
}

// GetMotypeOk returns a tuple with the Motype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueOdataConditionAllOf) GetMotypeOk() (*string, bool) {
	if o == nil || o.Motype == nil {
		return nil, false
	}
	return o.Motype, true
}

// HasMotype returns a boolean if a field has been set.
func (o *IssueOdataConditionAllOf) HasMotype() bool {
	if o != nil && o.Motype != nil {
		return true
	}

	return false
}

// SetMotype gets a reference to the given string and assigns it to the Motype field.
func (o *IssueOdataConditionAllOf) SetMotype(v string) {
	o.Motype = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *IssueOdataConditionAllOf) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueOdataConditionAllOf) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *IssueOdataConditionAllOf) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *IssueOdataConditionAllOf) SetQuery(v string) {
	o.Query = &v
}

func (o IssueOdataConditionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceTags != nil {
		toSerialize["DeviceTags"] = o.DeviceTags
	}
	if o.Motype != nil {
		toSerialize["Motype"] = o.Motype
	}
	if o.Query != nil {
		toSerialize["Query"] = o.Query
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IssueOdataConditionAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIssueOdataConditionAllOf := _IssueOdataConditionAllOf{}

	if err = json.Unmarshal(bytes, &varIssueOdataConditionAllOf); err == nil {
		*o = IssueOdataConditionAllOf(varIssueOdataConditionAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceTags")
		delete(additionalProperties, "Motype")
		delete(additionalProperties, "Query")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIssueOdataConditionAllOf struct {
	value *IssueOdataConditionAllOf
	isSet bool
}

func (v NullableIssueOdataConditionAllOf) Get() *IssueOdataConditionAllOf {
	return v.value
}

func (v *NullableIssueOdataConditionAllOf) Set(val *IssueOdataConditionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueOdataConditionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueOdataConditionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueOdataConditionAllOf(val *IssueOdataConditionAllOf) *NullableIssueOdataConditionAllOf {
	return &NullableIssueOdataConditionAllOf{value: val, isSet: true}
}

func (v NullableIssueOdataConditionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueOdataConditionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
