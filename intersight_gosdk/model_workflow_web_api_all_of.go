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

// WorkflowWebApiAllOf Definition of the list of properties defined in 'workflow.WebApi', excluding properties defined in parent classes.
type WorkflowWebApiAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Collection of key value pairs to set in the request header as Cookie list.
	Cookies interface{} `json:"Cookies,omitempty"`
	// If the target type is Endpoint, this property determines whether the request is to be handled as internal request or external request by the device connector. * `Internal` - The endpoint API executed is an internal request handled by the device connector plugin. * `External` - The endpoint API request is passed through by the device connector.
	EndpointRequestType *string `json:"EndpointRequestType,omitempty"`
	// Collection of key value pairs to set in the request header.
	Headers interface{} `json:"Headers,omitempty"`
	// The HTTP method to be executed in the given URL (GET, POST, PUT, etc). If the value is not specified, GET will be used as default. The supported values are GET, POST, PUT, DELETE, PATCH, HEAD.
	Method *string `json:"Method,omitempty"`
	// The type of the intersight object for which API request is to be made. The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value.
	MoType *string `json:"MoType,omitempty"`
	// The accepted web protocol values are http and https.
	Protocol *string `json:"Protocol,omitempty"`
	// If the web API is to be executed in a remote device connected to the Intersight through device connector, 'Endpoint' is expected as the value whereas if the API is an Intersight API, 'Local' is expected as the value.
	TargetType *string `json:"TargetType,omitempty"`
	// The URL of the resource in the target to which the API request is made.
	Url                  *string `json:"Url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWebApiAllOf WorkflowWebApiAllOf

// NewWorkflowWebApiAllOf instantiates a new WorkflowWebApiAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWebApiAllOf(classId string, objectType string) *WorkflowWebApiAllOf {
	this := WorkflowWebApiAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	return &this
}

// NewWorkflowWebApiAllOfWithDefaults instantiates a new WorkflowWebApiAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWebApiAllOfWithDefaults() *WorkflowWebApiAllOf {
	this := WorkflowWebApiAllOf{}
	var classId string = "workflow.WebApi"
	this.ClassId = classId
	var objectType string = "workflow.WebApi"
	this.ObjectType = objectType
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWebApiAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWebApiAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWebApiAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWebApiAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCookies returns the Cookies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWebApiAllOf) GetCookies() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Cookies
}

// GetCookiesOk returns a tuple with the Cookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWebApiAllOf) GetCookiesOk() (*interface{}, bool) {
	if o == nil || o.Cookies == nil {
		return nil, false
	}
	return &o.Cookies, true
}

// HasCookies returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasCookies() bool {
	if o != nil && o.Cookies != nil {
		return true
	}

	return false
}

// SetCookies gets a reference to the given interface{} and assigns it to the Cookies field.
func (o *WorkflowWebApiAllOf) SetCookies(v interface{}) {
	o.Cookies = v
}

// GetEndpointRequestType returns the EndpointRequestType field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetEndpointRequestType() string {
	if o == nil || o.EndpointRequestType == nil {
		var ret string
		return ret
	}
	return *o.EndpointRequestType
}

// GetEndpointRequestTypeOk returns a tuple with the EndpointRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetEndpointRequestTypeOk() (*string, bool) {
	if o == nil || o.EndpointRequestType == nil {
		return nil, false
	}
	return o.EndpointRequestType, true
}

// HasEndpointRequestType returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasEndpointRequestType() bool {
	if o != nil && o.EndpointRequestType != nil {
		return true
	}

	return false
}

// SetEndpointRequestType gets a reference to the given string and assigns it to the EndpointRequestType field.
func (o *WorkflowWebApiAllOf) SetEndpointRequestType(v string) {
	o.EndpointRequestType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWebApiAllOf) GetHeaders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWebApiAllOf) GetHeadersOk() (*interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given interface{} and assigns it to the Headers field.
func (o *WorkflowWebApiAllOf) SetHeaders(v interface{}) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkflowWebApiAllOf) SetMethod(v string) {
	o.Method = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *WorkflowWebApiAllOf) SetMoType(v string) {
	o.MoType = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *WorkflowWebApiAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowWebApiAllOf) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkflowWebApiAllOf) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApiAllOf) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkflowWebApiAllOf) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkflowWebApiAllOf) SetUrl(v string) {
	o.Url = &v
}

func (o WorkflowWebApiAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cookies != nil {
		toSerialize["Cookies"] = o.Cookies
	}
	if o.EndpointRequestType != nil {
		toSerialize["EndpointRequestType"] = o.EndpointRequestType
	}
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWebApiAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWebApiAllOf := _WorkflowWebApiAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWebApiAllOf); err == nil {
		*o = WorkflowWebApiAllOf(varWorkflowWebApiAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cookies")
		delete(additionalProperties, "EndpointRequestType")
		delete(additionalProperties, "Headers")
		delete(additionalProperties, "Method")
		delete(additionalProperties, "MoType")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "TargetType")
		delete(additionalProperties, "Url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWebApiAllOf struct {
	value *WorkflowWebApiAllOf
	isSet bool
}

func (v NullableWorkflowWebApiAllOf) Get() *WorkflowWebApiAllOf {
	return v.value
}

func (v *NullableWorkflowWebApiAllOf) Set(val *WorkflowWebApiAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWebApiAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWebApiAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWebApiAllOf(val *WorkflowWebApiAllOf) *NullableWorkflowWebApiAllOf {
	return &NullableWorkflowWebApiAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWebApiAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWebApiAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
