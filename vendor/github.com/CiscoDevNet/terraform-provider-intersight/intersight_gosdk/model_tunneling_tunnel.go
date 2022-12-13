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

// TunnelingTunnel The tunnel that exists between the client and the DC.
type TunnelingTunnel struct {
	SessionAbstractSubSession
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The multiplexer URL for the client to connect on.
	ClientUrl            *string `json:"ClientUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TunnelingTunnel TunnelingTunnel

// NewTunnelingTunnel instantiates a new TunnelingTunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnelingTunnel(classId string, objectType string) *TunnelingTunnel {
	this := TunnelingTunnel{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "Active"
	this.Status = &status
	return &this
}

// NewTunnelingTunnelWithDefaults instantiates a new TunnelingTunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelingTunnelWithDefaults() *TunnelingTunnel {
	this := TunnelingTunnel{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *TunnelingTunnel) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TunnelingTunnel) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TunnelingTunnel) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TunnelingTunnel) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TunnelingTunnel) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TunnelingTunnel) SetObjectType(v string) {
	o.ObjectType = v
}

// GetClientUrl returns the ClientUrl field value if set, zero value otherwise.
func (o *TunnelingTunnel) GetClientUrl() string {
	if o == nil || o.ClientUrl == nil {
		var ret string
		return ret
	}
	return *o.ClientUrl
}

// GetClientUrlOk returns a tuple with the ClientUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunnelingTunnel) GetClientUrlOk() (*string, bool) {
	if o == nil || o.ClientUrl == nil {
		return nil, false
	}
	return o.ClientUrl, true
}

// HasClientUrl returns a boolean if a field has been set.
func (o *TunnelingTunnel) HasClientUrl() bool {
	if o != nil && o.ClientUrl != nil {
		return true
	}

	return false
}

// SetClientUrl gets a reference to the given string and assigns it to the ClientUrl field.
func (o *TunnelingTunnel) SetClientUrl(v string) {
	o.ClientUrl = &v
}

func (o TunnelingTunnel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSessionAbstractSubSession, errSessionAbstractSubSession := json.Marshal(o.SessionAbstractSubSession)
	if errSessionAbstractSubSession != nil {
		return []byte{}, errSessionAbstractSubSession
	}
	errSessionAbstractSubSession = json.Unmarshal([]byte(serializedSessionAbstractSubSession), &toSerialize)
	if errSessionAbstractSubSession != nil {
		return []byte{}, errSessionAbstractSubSession
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ClientUrl != nil {
		toSerialize["ClientUrl"] = o.ClientUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TunnelingTunnel) UnmarshalJSON(bytes []byte) (err error) {
	type TunnelingTunnelWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The multiplexer URL for the client to connect on.
		ClientUrl *string `json:"ClientUrl,omitempty"`
	}

	varTunnelingTunnelWithoutEmbeddedStruct := TunnelingTunnelWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTunnelingTunnelWithoutEmbeddedStruct)
	if err == nil {
		varTunnelingTunnel := _TunnelingTunnel{}
		varTunnelingTunnel.ClassId = varTunnelingTunnelWithoutEmbeddedStruct.ClassId
		varTunnelingTunnel.ObjectType = varTunnelingTunnelWithoutEmbeddedStruct.ObjectType
		varTunnelingTunnel.ClientUrl = varTunnelingTunnelWithoutEmbeddedStruct.ClientUrl
		*o = TunnelingTunnel(varTunnelingTunnel)
	} else {
		return err
	}

	varTunnelingTunnel := _TunnelingTunnel{}

	err = json.Unmarshal(bytes, &varTunnelingTunnel)
	if err == nil {
		o.SessionAbstractSubSession = varTunnelingTunnel.SessionAbstractSubSession
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientUrl")

		// remove fields from embedded structs
		reflectSessionAbstractSubSession := reflect.ValueOf(o.SessionAbstractSubSession)
		for i := 0; i < reflectSessionAbstractSubSession.Type().NumField(); i++ {
			t := reflectSessionAbstractSubSession.Type().Field(i)

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

type NullableTunnelingTunnel struct {
	value *TunnelingTunnel
	isSet bool
}

func (v NullableTunnelingTunnel) Get() *TunnelingTunnel {
	return v.value
}

func (v *NullableTunnelingTunnel) Set(val *TunnelingTunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelingTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelingTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelingTunnel(val *TunnelingTunnel) *NullableTunnelingTunnel {
	return &NullableTunnelingTunnel{value: val, isSet: true}
}

func (v NullableTunnelingTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelingTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
