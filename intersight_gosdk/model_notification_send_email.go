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

// NotificationSendEmail SendEmail complexType holds the recipient's email address.
type NotificationSendEmail struct {
	NotificationAction
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Email of the recipient, who expects to be notified about the event that happens in Intersight.
	Email                *string `json:"Email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationSendEmail NotificationSendEmail

// NewNotificationSendEmail instantiates a new NotificationSendEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationSendEmail(classId string, objectType string) *NotificationSendEmail {
	this := NotificationSendEmail{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNotificationSendEmailWithDefaults instantiates a new NotificationSendEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationSendEmailWithDefaults() *NotificationSendEmail {
	this := NotificationSendEmail{}
	var classId string = "notification.SendEmail"
	this.ClassId = classId
	var objectType string = "notification.SendEmail"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NotificationSendEmail) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NotificationSendEmail) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NotificationSendEmail) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NotificationSendEmail) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NotificationSendEmail) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NotificationSendEmail) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NotificationSendEmail) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationSendEmail) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NotificationSendEmail) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NotificationSendEmail) SetEmail(v string) {
	o.Email = &v
}

func (o NotificationSendEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedNotificationAction, errNotificationAction := json.Marshal(o.NotificationAction)
	if errNotificationAction != nil {
		return []byte{}, errNotificationAction
	}
	errNotificationAction = json.Unmarshal([]byte(serializedNotificationAction), &toSerialize)
	if errNotificationAction != nil {
		return []byte{}, errNotificationAction
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NotificationSendEmail) UnmarshalJSON(bytes []byte) (err error) {
	type NotificationSendEmailWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Email of the recipient, who expects to be notified about the event that happens in Intersight.
		Email *string `json:"Email,omitempty"`
	}

	varNotificationSendEmailWithoutEmbeddedStruct := NotificationSendEmailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNotificationSendEmailWithoutEmbeddedStruct)
	if err == nil {
		varNotificationSendEmail := _NotificationSendEmail{}
		varNotificationSendEmail.ClassId = varNotificationSendEmailWithoutEmbeddedStruct.ClassId
		varNotificationSendEmail.ObjectType = varNotificationSendEmailWithoutEmbeddedStruct.ObjectType
		varNotificationSendEmail.Email = varNotificationSendEmailWithoutEmbeddedStruct.Email
		*o = NotificationSendEmail(varNotificationSendEmail)
	} else {
		return err
	}

	varNotificationSendEmail := _NotificationSendEmail{}

	err = json.Unmarshal(bytes, &varNotificationSendEmail)
	if err == nil {
		o.NotificationAction = varNotificationSendEmail.NotificationAction
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Email")

		// remove fields from embedded structs
		reflectNotificationAction := reflect.ValueOf(o.NotificationAction)
		for i := 0; i < reflectNotificationAction.Type().NumField(); i++ {
			t := reflectNotificationAction.Type().Field(i)

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

type NullableNotificationSendEmail struct {
	value *NotificationSendEmail
	isSet bool
}

func (v NullableNotificationSendEmail) Get() *NotificationSendEmail {
	return v.value
}

func (v *NullableNotificationSendEmail) Set(val *NotificationSendEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationSendEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationSendEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationSendEmail(val *NotificationSendEmail) *NullableNotificationSendEmail {
	return &NullableNotificationSendEmail{value: val, isSet: true}
}

func (v NullableNotificationSendEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationSendEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
