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

// checks if the IaasServiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IaasServiceRequest{}

// IaasServiceRequest Gets last six months Service Requests from UCSD.
type IaasServiceRequest struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Service request duration.
	Duration *string `json:"Duration,omitempty"`
	// Service Request initiating user.
	InitiatingUser *string `json:"InitiatingUser,omitempty"`
	// Service request end time.
	RequestEndTime *string `json:"RequestEndTime,omitempty"`
	// Service request id of an SR.
	RequestId *string `json:"RequestId,omitempty"`
	// Service request start time.
	RequestStartTime *string `json:"RequestStartTime,omitempty"`
	// Service request type of an SR.
	RequestType *string `json:"RequestType,omitempty"`
	// UCSD service request status.
	Status *string `json:"Status,omitempty"`
	// Executed workflow name for an SR.
	WorkflowName         *string                                     `json:"WorkflowName,omitempty"`
	WorkflowSteps        []IaasWorkflowSteps                         `json:"WorkflowSteps,omitempty"`
	RegisteredDevice     NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IaasServiceRequest IaasServiceRequest

// NewIaasServiceRequest instantiates a new IaasServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasServiceRequest(classId string, objectType string) *IaasServiceRequest {
	this := IaasServiceRequest{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIaasServiceRequestWithDefaults instantiates a new IaasServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasServiceRequestWithDefaults() *IaasServiceRequest {
	this := IaasServiceRequest{}
	var classId string = "iaas.ServiceRequest"
	this.ClassId = classId
	var objectType string = "iaas.ServiceRequest"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IaasServiceRequest) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IaasServiceRequest) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iaas.ServiceRequest" of the ClassId field.
func (o *IaasServiceRequest) GetDefaultClassId() interface{} {
	return "iaas.ServiceRequest"
}

// GetObjectType returns the ObjectType field value
func (o *IaasServiceRequest) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IaasServiceRequest) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iaas.ServiceRequest" of the ObjectType field.
func (o *IaasServiceRequest) GetDefaultObjectType() interface{} {
	return "iaas.ServiceRequest"
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *IaasServiceRequest) SetDuration(v string) {
	o.Duration = &v
}

// GetInitiatingUser returns the InitiatingUser field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetInitiatingUser() string {
	if o == nil || IsNil(o.InitiatingUser) {
		var ret string
		return ret
	}
	return *o.InitiatingUser
}

// GetInitiatingUserOk returns a tuple with the InitiatingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetInitiatingUserOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatingUser) {
		return nil, false
	}
	return o.InitiatingUser, true
}

// HasInitiatingUser returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasInitiatingUser() bool {
	if o != nil && !IsNil(o.InitiatingUser) {
		return true
	}

	return false
}

// SetInitiatingUser gets a reference to the given string and assigns it to the InitiatingUser field.
func (o *IaasServiceRequest) SetInitiatingUser(v string) {
	o.InitiatingUser = &v
}

// GetRequestEndTime returns the RequestEndTime field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestEndTime() string {
	if o == nil || IsNil(o.RequestEndTime) {
		var ret string
		return ret
	}
	return *o.RequestEndTime
}

// GetRequestEndTimeOk returns a tuple with the RequestEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestEndTime) {
		return nil, false
	}
	return o.RequestEndTime, true
}

// HasRequestEndTime returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestEndTime() bool {
	if o != nil && !IsNil(o.RequestEndTime) {
		return true
	}

	return false
}

// SetRequestEndTime gets a reference to the given string and assigns it to the RequestEndTime field.
func (o *IaasServiceRequest) SetRequestEndTime(v string) {
	o.RequestEndTime = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *IaasServiceRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestStartTime returns the RequestStartTime field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestStartTime() string {
	if o == nil || IsNil(o.RequestStartTime) {
		var ret string
		return ret
	}
	return *o.RequestStartTime
}

// GetRequestStartTimeOk returns a tuple with the RequestStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestStartTime) {
		return nil, false
	}
	return o.RequestStartTime, true
}

// HasRequestStartTime returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestStartTime() bool {
	if o != nil && !IsNil(o.RequestStartTime) {
		return true
	}

	return false
}

// SetRequestStartTime gets a reference to the given string and assigns it to the RequestStartTime field.
func (o *IaasServiceRequest) SetRequestStartTime(v string) {
	o.RequestStartTime = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetRequestType() string {
	if o == nil || IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRequestType() bool {
	if o != nil && !IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *IaasServiceRequest) SetRequestType(v string) {
	o.RequestType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IaasServiceRequest) SetStatus(v string) {
	o.Status = &v
}

// GetWorkflowName returns the WorkflowName field value if set, zero value otherwise.
func (o *IaasServiceRequest) GetWorkflowName() string {
	if o == nil || IsNil(o.WorkflowName) {
		var ret string
		return ret
	}
	return *o.WorkflowName
}

// GetWorkflowNameOk returns a tuple with the WorkflowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasServiceRequest) GetWorkflowNameOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowName) {
		return nil, false
	}
	return o.WorkflowName, true
}

// HasWorkflowName returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasWorkflowName() bool {
	if o != nil && !IsNil(o.WorkflowName) {
		return true
	}

	return false
}

// SetWorkflowName gets a reference to the given string and assigns it to the WorkflowName field.
func (o *IaasServiceRequest) SetWorkflowName(v string) {
	o.WorkflowName = &v
}

// GetWorkflowSteps returns the WorkflowSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasServiceRequest) GetWorkflowSteps() []IaasWorkflowSteps {
	if o == nil {
		var ret []IaasWorkflowSteps
		return ret
	}
	return o.WorkflowSteps
}

// GetWorkflowStepsOk returns a tuple with the WorkflowSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasServiceRequest) GetWorkflowStepsOk() ([]IaasWorkflowSteps, bool) {
	if o == nil || IsNil(o.WorkflowSteps) {
		return nil, false
	}
	return o.WorkflowSteps, true
}

// HasWorkflowSteps returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasWorkflowSteps() bool {
	if o != nil && !IsNil(o.WorkflowSteps) {
		return true
	}

	return false
}

// SetWorkflowSteps gets a reference to the given []IaasWorkflowSteps and assigns it to the WorkflowSteps field.
func (o *IaasServiceRequest) SetWorkflowSteps(v []IaasWorkflowSteps) {
	o.WorkflowSteps = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IaasServiceRequest) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || IsNil(o.RegisteredDevice.Get()) {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice.Get()
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IaasServiceRequest) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredDevice.Get(), o.RegisteredDevice.IsSet()
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *IaasServiceRequest) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice.IsSet() {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given NullableAssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *IaasServiceRequest) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice.Set(&v)
}

// SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil
func (o *IaasServiceRequest) SetRegisteredDeviceNil() {
	o.RegisteredDevice.Set(nil)
}

// UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
func (o *IaasServiceRequest) UnsetRegisteredDevice() {
	o.RegisteredDevice.Unset()
}

func (o IaasServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IaasServiceRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Duration) {
		toSerialize["Duration"] = o.Duration
	}
	if !IsNil(o.InitiatingUser) {
		toSerialize["InitiatingUser"] = o.InitiatingUser
	}
	if !IsNil(o.RequestEndTime) {
		toSerialize["RequestEndTime"] = o.RequestEndTime
	}
	if !IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	if !IsNil(o.RequestStartTime) {
		toSerialize["RequestStartTime"] = o.RequestStartTime
	}
	if !IsNil(o.RequestType) {
		toSerialize["RequestType"] = o.RequestType
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.WorkflowName) {
		toSerialize["WorkflowName"] = o.WorkflowName
	}
	if o.WorkflowSteps != nil {
		toSerialize["WorkflowSteps"] = o.WorkflowSteps
	}
	if o.RegisteredDevice.IsSet() {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IaasServiceRequest) UnmarshalJSON(data []byte) (err error) {
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
	type IaasServiceRequestWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Service request duration.
		Duration *string `json:"Duration,omitempty"`
		// Service Request initiating user.
		InitiatingUser *string `json:"InitiatingUser,omitempty"`
		// Service request end time.
		RequestEndTime *string `json:"RequestEndTime,omitempty"`
		// Service request id of an SR.
		RequestId *string `json:"RequestId,omitempty"`
		// Service request start time.
		RequestStartTime *string `json:"RequestStartTime,omitempty"`
		// Service request type of an SR.
		RequestType *string `json:"RequestType,omitempty"`
		// UCSD service request status.
		Status *string `json:"Status,omitempty"`
		// Executed workflow name for an SR.
		WorkflowName     *string                                     `json:"WorkflowName,omitempty"`
		WorkflowSteps    []IaasWorkflowSteps                         `json:"WorkflowSteps,omitempty"`
		RegisteredDevice NullableAssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varIaasServiceRequestWithoutEmbeddedStruct := IaasServiceRequestWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIaasServiceRequestWithoutEmbeddedStruct)
	if err == nil {
		varIaasServiceRequest := _IaasServiceRequest{}
		varIaasServiceRequest.ClassId = varIaasServiceRequestWithoutEmbeddedStruct.ClassId
		varIaasServiceRequest.ObjectType = varIaasServiceRequestWithoutEmbeddedStruct.ObjectType
		varIaasServiceRequest.Duration = varIaasServiceRequestWithoutEmbeddedStruct.Duration
		varIaasServiceRequest.InitiatingUser = varIaasServiceRequestWithoutEmbeddedStruct.InitiatingUser
		varIaasServiceRequest.RequestEndTime = varIaasServiceRequestWithoutEmbeddedStruct.RequestEndTime
		varIaasServiceRequest.RequestId = varIaasServiceRequestWithoutEmbeddedStruct.RequestId
		varIaasServiceRequest.RequestStartTime = varIaasServiceRequestWithoutEmbeddedStruct.RequestStartTime
		varIaasServiceRequest.RequestType = varIaasServiceRequestWithoutEmbeddedStruct.RequestType
		varIaasServiceRequest.Status = varIaasServiceRequestWithoutEmbeddedStruct.Status
		varIaasServiceRequest.WorkflowName = varIaasServiceRequestWithoutEmbeddedStruct.WorkflowName
		varIaasServiceRequest.WorkflowSteps = varIaasServiceRequestWithoutEmbeddedStruct.WorkflowSteps
		varIaasServiceRequest.RegisteredDevice = varIaasServiceRequestWithoutEmbeddedStruct.RegisteredDevice
		*o = IaasServiceRequest(varIaasServiceRequest)
	} else {
		return err
	}

	varIaasServiceRequest := _IaasServiceRequest{}

	err = json.Unmarshal(data, &varIaasServiceRequest)
	if err == nil {
		o.MoBaseMo = varIaasServiceRequest.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Duration")
		delete(additionalProperties, "InitiatingUser")
		delete(additionalProperties, "RequestEndTime")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "RequestStartTime")
		delete(additionalProperties, "RequestType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "WorkflowName")
		delete(additionalProperties, "WorkflowSteps")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableIaasServiceRequest struct {
	value *IaasServiceRequest
	isSet bool
}

func (v NullableIaasServiceRequest) Get() *IaasServiceRequest {
	return v.value
}

func (v *NullableIaasServiceRequest) Set(val *IaasServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasServiceRequest(val *IaasServiceRequest) *NullableIaasServiceRequest {
	return &NullableIaasServiceRequest{value: val, isSet: true}
}

func (v NullableIaasServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
