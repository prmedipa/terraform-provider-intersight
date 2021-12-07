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
	"time"
)

// AaaAuditRecord AuditRecord presents the configuration changes made by the user per transaction.
type AaaAuditRecord struct {
	AaaAbstractAuditRecord
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information.
	Email *string `json:"Email,omitempty"`
	// The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before.
	InstId *string `json:"InstId,omitempty"`
	// The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information.
	SessionId *string `json:"SessionId,omitempty"`
	// The source IP of the client.
	SourceIp *string `json:"SourceIp,omitempty"`
	// The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted.
	Timestamp *time.Time `json:"Timestamp,omitempty"`
	// The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information.
	UserIdOrEmail        *string                 `json:"UserIdOrEmail,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	Sessions             *IamSessionRelationship `json:"Sessions,omitempty"`
	User                 *IamUserRelationship    `json:"User,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AaaAuditRecord AaaAuditRecord

// NewAaaAuditRecord instantiates a new AaaAuditRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAaaAuditRecord(classId string, objectType string) *AaaAuditRecord {
	this := AaaAuditRecord{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAaaAuditRecordWithDefaults instantiates a new AaaAuditRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAaaAuditRecordWithDefaults() *AaaAuditRecord {
	this := AaaAuditRecord{}
	var classId string = "aaa.AuditRecord"
	this.ClassId = classId
	var objectType string = "aaa.AuditRecord"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AaaAuditRecord) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AaaAuditRecord) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AaaAuditRecord) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AaaAuditRecord) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AaaAuditRecord) SetEmail(v string) {
	o.Email = &v
}

// GetInstId returns the InstId field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetInstId() string {
	if o == nil || o.InstId == nil {
		var ret string
		return ret
	}
	return *o.InstId
}

// GetInstIdOk returns a tuple with the InstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetInstIdOk() (*string, bool) {
	if o == nil || o.InstId == nil {
		return nil, false
	}
	return o.InstId, true
}

// HasInstId returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasInstId() bool {
	if o != nil && o.InstId != nil {
		return true
	}

	return false
}

// SetInstId gets a reference to the given string and assigns it to the InstId field.
func (o *AaaAuditRecord) SetInstId(v string) {
	o.InstId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetSessionId() string {
	if o == nil || o.SessionId == nil {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetSessionIdOk() (*string, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *AaaAuditRecord) SetSessionId(v string) {
	o.SessionId = &v
}

// GetSourceIp returns the SourceIp field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetSourceIp() string {
	if o == nil || o.SourceIp == nil {
		var ret string
		return ret
	}
	return *o.SourceIp
}

// GetSourceIpOk returns a tuple with the SourceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetSourceIpOk() (*string, bool) {
	if o == nil || o.SourceIp == nil {
		return nil, false
	}
	return o.SourceIp, true
}

// HasSourceIp returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasSourceIp() bool {
	if o != nil && o.SourceIp != nil {
		return true
	}

	return false
}

// SetSourceIp gets a reference to the given string and assigns it to the SourceIp field.
func (o *AaaAuditRecord) SetSourceIp(v string) {
	o.SourceIp = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AaaAuditRecord) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetUserIdOrEmail() string {
	if o == nil || o.UserIdOrEmail == nil {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || o.UserIdOrEmail == nil {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasUserIdOrEmail() bool {
	if o != nil && o.UserIdOrEmail != nil {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *AaaAuditRecord) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *AaaAuditRecord) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetSessions() IamSessionRelationship {
	if o == nil || o.Sessions == nil {
		var ret IamSessionRelationship
		return ret
	}
	return *o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetSessionsOk() (*IamSessionRelationship, bool) {
	if o == nil || o.Sessions == nil {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasSessions() bool {
	if o != nil && o.Sessions != nil {
		return true
	}

	return false
}

// SetSessions gets a reference to the given IamSessionRelationship and assigns it to the Sessions field.
func (o *AaaAuditRecord) SetSessions(v IamSessionRelationship) {
	o.Sessions = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AaaAuditRecord) GetUser() IamUserRelationship {
	if o == nil || o.User == nil {
		var ret IamUserRelationship
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AaaAuditRecord) GetUserOk() (*IamUserRelationship, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AaaAuditRecord) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given IamUserRelationship and assigns it to the User field.
func (o *AaaAuditRecord) SetUser(v IamUserRelationship) {
	o.User = &v
}

func (o AaaAuditRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAaaAbstractAuditRecord, errAaaAbstractAuditRecord := json.Marshal(o.AaaAbstractAuditRecord)
	if errAaaAbstractAuditRecord != nil {
		return []byte{}, errAaaAbstractAuditRecord
	}
	errAaaAbstractAuditRecord = json.Unmarshal([]byte(serializedAaaAbstractAuditRecord), &toSerialize)
	if errAaaAbstractAuditRecord != nil {
		return []byte{}, errAaaAbstractAuditRecord
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Email != nil {
		toSerialize["Email"] = o.Email
	}
	if o.InstId != nil {
		toSerialize["InstId"] = o.InstId
	}
	if o.SessionId != nil {
		toSerialize["SessionId"] = o.SessionId
	}
	if o.SourceIp != nil {
		toSerialize["SourceIp"] = o.SourceIp
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.UserIdOrEmail != nil {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Sessions != nil {
		toSerialize["Sessions"] = o.Sessions
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AaaAuditRecord) UnmarshalJSON(bytes []byte) (err error) {
	type AaaAuditRecordWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The email of the associated user that made the change.  In case the user is later deleted, we still have some reference to the information.
		Email *string `json:"Email,omitempty"`
		// The instance id of AuditRecordLocal, which is used to identify if the comming AuditRecordLocal was already processed before.
		InstId *string `json:"InstId,omitempty"`
		// The sessionId in which the user made the change. In case that the session is later deleted, we still have some reference to the information.
		SessionId *string `json:"SessionId,omitempty"`
		// The source IP of the client.
		SourceIp *string `json:"SourceIp,omitempty"`
		// The creation time of AuditRecordLocal, which is the time when the affected MO was created/modified/deleted.
		Timestamp *time.Time `json:"Timestamp,omitempty"`
		// The userId or the email of the associated user that made the change. In case that user is later deleted, we still have some reference to the information.
		UserIdOrEmail *string                 `json:"UserIdOrEmail,omitempty"`
		Account       *IamAccountRelationship `json:"Account,omitempty"`
		Sessions      *IamSessionRelationship `json:"Sessions,omitempty"`
		User          *IamUserRelationship    `json:"User,omitempty"`
	}

	varAaaAuditRecordWithoutEmbeddedStruct := AaaAuditRecordWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAaaAuditRecordWithoutEmbeddedStruct)
	if err == nil {
		varAaaAuditRecord := _AaaAuditRecord{}
		varAaaAuditRecord.ClassId = varAaaAuditRecordWithoutEmbeddedStruct.ClassId
		varAaaAuditRecord.ObjectType = varAaaAuditRecordWithoutEmbeddedStruct.ObjectType
		varAaaAuditRecord.Email = varAaaAuditRecordWithoutEmbeddedStruct.Email
		varAaaAuditRecord.InstId = varAaaAuditRecordWithoutEmbeddedStruct.InstId
		varAaaAuditRecord.SessionId = varAaaAuditRecordWithoutEmbeddedStruct.SessionId
		varAaaAuditRecord.SourceIp = varAaaAuditRecordWithoutEmbeddedStruct.SourceIp
		varAaaAuditRecord.Timestamp = varAaaAuditRecordWithoutEmbeddedStruct.Timestamp
		varAaaAuditRecord.UserIdOrEmail = varAaaAuditRecordWithoutEmbeddedStruct.UserIdOrEmail
		varAaaAuditRecord.Account = varAaaAuditRecordWithoutEmbeddedStruct.Account
		varAaaAuditRecord.Sessions = varAaaAuditRecordWithoutEmbeddedStruct.Sessions
		varAaaAuditRecord.User = varAaaAuditRecordWithoutEmbeddedStruct.User
		*o = AaaAuditRecord(varAaaAuditRecord)
	} else {
		return err
	}

	varAaaAuditRecord := _AaaAuditRecord{}

	err = json.Unmarshal(bytes, &varAaaAuditRecord)
	if err == nil {
		o.AaaAbstractAuditRecord = varAaaAuditRecord.AaaAbstractAuditRecord
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "InstId")
		delete(additionalProperties, "SessionId")
		delete(additionalProperties, "SourceIp")
		delete(additionalProperties, "Timestamp")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "Sessions")
		delete(additionalProperties, "User")

		// remove fields from embedded structs
		reflectAaaAbstractAuditRecord := reflect.ValueOf(o.AaaAbstractAuditRecord)
		for i := 0; i < reflectAaaAbstractAuditRecord.Type().NumField(); i++ {
			t := reflectAaaAbstractAuditRecord.Type().Field(i)

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

type NullableAaaAuditRecord struct {
	value *AaaAuditRecord
	isSet bool
}

func (v NullableAaaAuditRecord) Get() *AaaAuditRecord {
	return v.value
}

func (v *NullableAaaAuditRecord) Set(val *AaaAuditRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableAaaAuditRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableAaaAuditRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAaaAuditRecord(val *AaaAuditRecord) *NullableAaaAuditRecord {
	return &NullableAaaAuditRecord{value: val, isSet: true}
}

func (v NullableAaaAuditRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAaaAuditRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
