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

// checks if the AaaAbstractAuditRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AaaAbstractAuditRecord{}

// AaaAbstractAuditRecord AbstractAuditRecord is an abstract base type that specifies the common properties for all audit log records concrete sub-types.
type AaaAbstractAuditRecord struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The operation that was performed on this Managed Object. The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type.
	Event *string `json:"Event,omitempty"`
	// The user-friendly names of the changed MO.
	MoDisplayNames interface{} `json:"MoDisplayNames,omitempty"`
	// The object type of the REST resource that was created, modified or deleted.
	MoType *string `json:"MoType,omitempty"`
	// The Moid of the REST resource that was created, modified or deleted.
	ObjectMoid *string `json:"ObjectMoid,omitempty"`
	// The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document.
	Request interface{} `json:"Request,omitempty"`
	// The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team.
	TraceId   *string              `json:"TraceId,omitempty"`
	UserAgent NullableAaaUserAgent `json:"UserAgent,omitempty"`
	// The raw, string representation of the user agent of the request from the user-agent http request header.
	UserAgentString      *string `json:"UserAgentString,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AaaAbstractAuditRecord AaaAbstractAuditRecord

// NewAaaAbstractAuditRecord instantiates a new AaaAbstractAuditRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaaAbstractAuditRecord(classId string, objectType string) *AaaAbstractAuditRecord {
	this := AaaAbstractAuditRecord{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAaaAbstractAuditRecordWithDefaults instantiates a new AaaAbstractAuditRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaaAbstractAuditRecordWithDefaults() *AaaAbstractAuditRecord {
	this := AaaAbstractAuditRecord{}
	var classId string = "aaa.AuditRecord"
	this.ClassId = classId
	var objectType string = "aaa.AuditRecord"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AaaAbstractAuditRecord) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AaaAbstractAuditRecord) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "aaa.AuditRecord" of the ClassId field.
func (o *AaaAbstractAuditRecord) GetDefaultClassId() interface{} {
	return "aaa.AuditRecord"
}

// GetObjectType returns the ObjectType field value
func (o *AaaAbstractAuditRecord) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AaaAbstractAuditRecord) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "aaa.AuditRecord" of the ObjectType field.
func (o *AaaAbstractAuditRecord) GetDefaultObjectType() interface{} {
	return "aaa.AuditRecord"
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecord) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *AaaAbstractAuditRecord) SetEvent(v string) {
	o.Event = &v
}

// GetMoDisplayNames returns the MoDisplayNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecord) GetMoDisplayNames() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MoDisplayNames
}

// GetMoDisplayNamesOk returns a tuple with the MoDisplayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecord) GetMoDisplayNamesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MoDisplayNames) {
		return nil, false
	}
	return &o.MoDisplayNames, true
}

// HasMoDisplayNames returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasMoDisplayNames() bool {
	if o != nil && !IsNil(o.MoDisplayNames) {
		return true
	}

	return false
}

// SetMoDisplayNames gets a reference to the given interface{} and assigns it to the MoDisplayNames field.
func (o *AaaAbstractAuditRecord) SetMoDisplayNames(v interface{}) {
	o.MoDisplayNames = v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecord) GetMoType() string {
	if o == nil || IsNil(o.MoType) {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetMoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MoType) {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasMoType() bool {
	if o != nil && !IsNil(o.MoType) {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *AaaAbstractAuditRecord) SetMoType(v string) {
	o.MoType = &v
}

// GetObjectMoid returns the ObjectMoid field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecord) GetObjectMoid() string {
	if o == nil || IsNil(o.ObjectMoid) {
		var ret string
		return ret
	}
	return *o.ObjectMoid
}

// GetObjectMoidOk returns a tuple with the ObjectMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetObjectMoidOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectMoid) {
		return nil, false
	}
	return o.ObjectMoid, true
}

// HasObjectMoid returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasObjectMoid() bool {
	if o != nil && !IsNil(o.ObjectMoid) {
		return true
	}

	return false
}

// SetObjectMoid gets a reference to the given string and assigns it to the ObjectMoid field.
func (o *AaaAbstractAuditRecord) SetObjectMoid(v string) {
	o.ObjectMoid = &v
}

// GetRequest returns the Request field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecord) GetRequest() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecord) GetRequestOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return &o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given interface{} and assigns it to the Request field.
func (o *AaaAbstractAuditRecord) SetRequest(v interface{}) {
	o.Request = v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecord) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *AaaAbstractAuditRecord) SetTraceId(v string) {
	o.TraceId = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AaaAbstractAuditRecord) GetUserAgent() AaaUserAgent {
	if o == nil || IsNil(o.UserAgent.Get()) {
		var ret AaaUserAgent
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AaaAbstractAuditRecord) GetUserAgentOk() (*AaaUserAgent, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableAaaUserAgent and assigns it to the UserAgent field.
func (o *AaaAbstractAuditRecord) SetUserAgent(v AaaUserAgent) {
	o.UserAgent.Set(&v)
}

// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *AaaAbstractAuditRecord) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *AaaAbstractAuditRecord) UnsetUserAgent() {
	o.UserAgent.Unset()
}

// GetUserAgentString returns the UserAgentString field value if set, zero value otherwise.
func (o *AaaAbstractAuditRecord) GetUserAgentString() string {
	if o == nil || IsNil(o.UserAgentString) {
		var ret string
		return ret
	}
	return *o.UserAgentString
}

// GetUserAgentStringOk returns a tuple with the UserAgentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAbstractAuditRecord) GetUserAgentStringOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgentString) {
		return nil, false
	}
	return o.UserAgentString, true
}

// HasUserAgentString returns a boolean if a field has been set.
func (o *AaaAbstractAuditRecord) HasUserAgentString() bool {
	if o != nil && !IsNil(o.UserAgentString) {
		return true
	}

	return false
}

// SetUserAgentString gets a reference to the given string and assigns it to the UserAgentString field.
func (o *AaaAbstractAuditRecord) SetUserAgentString(v string) {
	o.UserAgentString = &v
}

func (o AaaAbstractAuditRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AaaAbstractAuditRecord) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Event) {
		toSerialize["Event"] = o.Event
	}
	if o.MoDisplayNames != nil {
		toSerialize["MoDisplayNames"] = o.MoDisplayNames
	}
	if !IsNil(o.MoType) {
		toSerialize["MoType"] = o.MoType
	}
	if !IsNil(o.ObjectMoid) {
		toSerialize["ObjectMoid"] = o.ObjectMoid
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if !IsNil(o.TraceId) {
		toSerialize["TraceId"] = o.TraceId
	}
	if o.UserAgent.IsSet() {
		toSerialize["UserAgent"] = o.UserAgent.Get()
	}
	if !IsNil(o.UserAgentString) {
		toSerialize["UserAgentString"] = o.UserAgentString
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AaaAbstractAuditRecord) UnmarshalJSON(data []byte) (err error) {
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
	type AaaAbstractAuditRecordWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The operation that was performed on this Managed Object. The event is a compound string that includes the CRUD operation such as Create, Modify, Delete, and a string representing the Managed Object type.
		Event *string `json:"Event,omitempty"`
		// The user-friendly names of the changed MO.
		MoDisplayNames interface{} `json:"MoDisplayNames,omitempty"`
		// The object type of the REST resource that was created, modified or deleted.
		MoType *string `json:"MoType,omitempty"`
		// The Moid of the REST resource that was created, modified or deleted.
		ObjectMoid *string `json:"ObjectMoid,omitempty"`
		// The body of the REST request that was received from a client to create or modify a REST resource, represented as a JSON document.
		Request interface{} `json:"Request,omitempty"`
		// The trace id of the request that was used to create, modify or delete a REST resource. A trace id is a unique identifier for one particular REST request. It may be used for troubleshooting purpose by the Intersight technical support team.
		TraceId   *string              `json:"TraceId,omitempty"`
		UserAgent NullableAaaUserAgent `json:"UserAgent,omitempty"`
		// The raw, string representation of the user agent of the request from the user-agent http request header.
		UserAgentString *string `json:"UserAgentString,omitempty"`
	}

	varAaaAbstractAuditRecordWithoutEmbeddedStruct := AaaAbstractAuditRecordWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAaaAbstractAuditRecordWithoutEmbeddedStruct)
	if err == nil {
		varAaaAbstractAuditRecord := _AaaAbstractAuditRecord{}
		varAaaAbstractAuditRecord.ClassId = varAaaAbstractAuditRecordWithoutEmbeddedStruct.ClassId
		varAaaAbstractAuditRecord.ObjectType = varAaaAbstractAuditRecordWithoutEmbeddedStruct.ObjectType
		varAaaAbstractAuditRecord.Event = varAaaAbstractAuditRecordWithoutEmbeddedStruct.Event
		varAaaAbstractAuditRecord.MoDisplayNames = varAaaAbstractAuditRecordWithoutEmbeddedStruct.MoDisplayNames
		varAaaAbstractAuditRecord.MoType = varAaaAbstractAuditRecordWithoutEmbeddedStruct.MoType
		varAaaAbstractAuditRecord.ObjectMoid = varAaaAbstractAuditRecordWithoutEmbeddedStruct.ObjectMoid
		varAaaAbstractAuditRecord.Request = varAaaAbstractAuditRecordWithoutEmbeddedStruct.Request
		varAaaAbstractAuditRecord.TraceId = varAaaAbstractAuditRecordWithoutEmbeddedStruct.TraceId
		varAaaAbstractAuditRecord.UserAgent = varAaaAbstractAuditRecordWithoutEmbeddedStruct.UserAgent
		varAaaAbstractAuditRecord.UserAgentString = varAaaAbstractAuditRecordWithoutEmbeddedStruct.UserAgentString
		*o = AaaAbstractAuditRecord(varAaaAbstractAuditRecord)
	} else {
		return err
	}

	varAaaAbstractAuditRecord := _AaaAbstractAuditRecord{}

	err = json.Unmarshal(data, &varAaaAbstractAuditRecord)
	if err == nil {
		o.MoBaseMo = varAaaAbstractAuditRecord.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Event")
		delete(additionalProperties, "MoDisplayNames")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "ObjectMoid")
		delete(additionalProperties, "Request")
		delete(additionalProperties, "TraceId")
		delete(additionalProperties, "UserAgent")
		delete(additionalProperties, "UserAgentString")

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

type NullableAaaAbstractAuditRecord struct {
	value *AaaAbstractAuditRecord
	isSet bool
}

func (v NullableAaaAbstractAuditRecord) Get() *AaaAbstractAuditRecord {
	return v.value
}

func (v *NullableAaaAbstractAuditRecord) Set(val *AaaAbstractAuditRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAaaAbstractAuditRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAaaAbstractAuditRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaaAbstractAuditRecord(val *AaaAbstractAuditRecord) *NullableAaaAbstractAuditRecord {
	return &NullableAaaAbstractAuditRecord{value: val, isSet: true}
}

func (v NullableAaaAbstractAuditRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaaAbstractAuditRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
