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
	"time"
)

// CloudNetworkInstanceAttachmentAllOf Definition of the list of properties defined in 'cloud.NetworkInstanceAttachment', excluding properties defined in parent classes.
type CloudNetworkInstanceAttachmentAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Time stamp when the attachment was initiated.
	AttachTime *time.Time `json:"AttachTime,omitempty"`
	// Indicates whether the attachment is deleted when an instance is terminated.
	AutoDelete *bool `json:"AutoDelete,omitempty"`
	// The device index to which the network interface is attached.
	DeviceIndex *int64 `json:"DeviceIndex,omitempty"`
	// The ID of the instance to which the network interface is attached.
	InstanceId *string `json:"InstanceId,omitempty"`
	// The status of the attachment. It is one of attaching, attached, detaching, or detached. * `UnAttached` - Network interface is not attached to a virtual machine. * `Attached` - Network interface is attached to a virtual machine. * `Attaching` - Network interface is being attached to a virtual machine. * `Detaching` - Network interface is being attached to a virtual machine. * `Detached` - Network interface is detached from a virtual machine.
	State                *string `json:"State,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudNetworkInstanceAttachmentAllOf CloudNetworkInstanceAttachmentAllOf

// NewCloudNetworkInstanceAttachmentAllOf instantiates a new CloudNetworkInstanceAttachmentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudNetworkInstanceAttachmentAllOf(classId string, objectType string) *CloudNetworkInstanceAttachmentAllOf {
	this := CloudNetworkInstanceAttachmentAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCloudNetworkInstanceAttachmentAllOfWithDefaults instantiates a new CloudNetworkInstanceAttachmentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudNetworkInstanceAttachmentAllOfWithDefaults() *CloudNetworkInstanceAttachmentAllOf {
	this := CloudNetworkInstanceAttachmentAllOf{}
	var classId string = "cloud.NetworkInstanceAttachment"
	this.ClassId = classId
	var objectType string = "cloud.NetworkInstanceAttachment"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CloudNetworkInstanceAttachmentAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CloudNetworkInstanceAttachmentAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CloudNetworkInstanceAttachmentAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CloudNetworkInstanceAttachmentAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAttachTime returns the AttachTime field value if set, zero value otherwise.
func (o *CloudNetworkInstanceAttachmentAllOf) GetAttachTime() time.Time {
	if o == nil || o.AttachTime == nil {
		var ret time.Time
		return ret
	}
	return *o.AttachTime
}

// GetAttachTimeOk returns a tuple with the AttachTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetAttachTimeOk() (*time.Time, bool) {
	if o == nil || o.AttachTime == nil {
		return nil, false
	}
	return o.AttachTime, true
}

// HasAttachTime returns a boolean if a field has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) HasAttachTime() bool {
	if o != nil && o.AttachTime != nil {
		return true
	}

	return false
}

// SetAttachTime gets a reference to the given time.Time and assigns it to the AttachTime field.
func (o *CloudNetworkInstanceAttachmentAllOf) SetAttachTime(v time.Time) {
	o.AttachTime = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *CloudNetworkInstanceAttachmentAllOf) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) HasAutoDelete() bool {
	if o != nil && o.AutoDelete != nil {
		return true
	}

	return false
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *CloudNetworkInstanceAttachmentAllOf) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetDeviceIndex returns the DeviceIndex field value if set, zero value otherwise.
func (o *CloudNetworkInstanceAttachmentAllOf) GetDeviceIndex() int64 {
	if o == nil || o.DeviceIndex == nil {
		var ret int64
		return ret
	}
	return *o.DeviceIndex
}

// GetDeviceIndexOk returns a tuple with the DeviceIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetDeviceIndexOk() (*int64, bool) {
	if o == nil || o.DeviceIndex == nil {
		return nil, false
	}
	return o.DeviceIndex, true
}

// HasDeviceIndex returns a boolean if a field has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) HasDeviceIndex() bool {
	if o != nil && o.DeviceIndex != nil {
		return true
	}

	return false
}

// SetDeviceIndex gets a reference to the given int64 and assigns it to the DeviceIndex field.
func (o *CloudNetworkInstanceAttachmentAllOf) SetDeviceIndex(v int64) {
	o.DeviceIndex = &v
}

// GetInstanceId returns the InstanceId field value if set, zero value otherwise.
func (o *CloudNetworkInstanceAttachmentAllOf) GetInstanceId() string {
	if o == nil || o.InstanceId == nil {
		var ret string
		return ret
	}
	return *o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetInstanceIdOk() (*string, bool) {
	if o == nil || o.InstanceId == nil {
		return nil, false
	}
	return o.InstanceId, true
}

// HasInstanceId returns a boolean if a field has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) HasInstanceId() bool {
	if o != nil && o.InstanceId != nil {
		return true
	}

	return false
}

// SetInstanceId gets a reference to the given string and assigns it to the InstanceId field.
func (o *CloudNetworkInstanceAttachmentAllOf) SetInstanceId(v string) {
	o.InstanceId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudNetworkInstanceAttachmentAllOf) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudNetworkInstanceAttachmentAllOf) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CloudNetworkInstanceAttachmentAllOf) SetState(v string) {
	o.State = &v
}

func (o CloudNetworkInstanceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AttachTime != nil {
		toSerialize["AttachTime"] = o.AttachTime
	}
	if o.AutoDelete != nil {
		toSerialize["AutoDelete"] = o.AutoDelete
	}
	if o.DeviceIndex != nil {
		toSerialize["DeviceIndex"] = o.DeviceIndex
	}
	if o.InstanceId != nil {
		toSerialize["InstanceId"] = o.InstanceId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CloudNetworkInstanceAttachmentAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCloudNetworkInstanceAttachmentAllOf := _CloudNetworkInstanceAttachmentAllOf{}

	if err = json.Unmarshal(bytes, &varCloudNetworkInstanceAttachmentAllOf); err == nil {
		*o = CloudNetworkInstanceAttachmentAllOf(varCloudNetworkInstanceAttachmentAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AttachTime")
		delete(additionalProperties, "AutoDelete")
		delete(additionalProperties, "DeviceIndex")
		delete(additionalProperties, "InstanceId")
		delete(additionalProperties, "State")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudNetworkInstanceAttachmentAllOf struct {
	value *CloudNetworkInstanceAttachmentAllOf
	isSet bool
}

func (v NullableCloudNetworkInstanceAttachmentAllOf) Get() *CloudNetworkInstanceAttachmentAllOf {
	return v.value
}

func (v *NullableCloudNetworkInstanceAttachmentAllOf) Set(val *CloudNetworkInstanceAttachmentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudNetworkInstanceAttachmentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudNetworkInstanceAttachmentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudNetworkInstanceAttachmentAllOf(val *CloudNetworkInstanceAttachmentAllOf) *NullableCloudNetworkInstanceAttachmentAllOf {
	return &NullableCloudNetworkInstanceAttachmentAllOf{value: val, isSet: true}
}

func (v NullableCloudNetworkInstanceAttachmentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudNetworkInstanceAttachmentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
