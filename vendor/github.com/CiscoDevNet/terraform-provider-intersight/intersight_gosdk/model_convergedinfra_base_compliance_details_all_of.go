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

// ConvergedinfraBaseComplianceDetailsAllOf Definition of the list of properties defined in 'convergedinfra.BaseComplianceDetails', excluding properties defined in parent classes.
type ConvergedinfraBaseComplianceDetailsAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The name of the component for which the compliance is evaluated.
	Name *string `json:"Name,omitempty"`
	// Reason for the status, when the status is Incomplete or NotApproved. Reason should help to identify the component that is not compliant. * `NotEvaluated` - The validation for the HCL or Interop status is yet to be performed. * `Missing-Os-Info` - This means the Interop status for the sever cannot be computed because we have missing OS information. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Incompatible-Storage-Os` - The validation failed because the given storage name and version combination is not found in Interop matrix. * `Incompatible-Os-Info` - The validation failed because the given OS name and version combination is not found in HCL. * `Incompatible-Processor` - The validation failed because the given processor is not found for the given server PID in HCL. * `Incompatible-Server-Platform` - The validation failed because the given server platform is not found in the Interop matrix. * `Incompatible-Server-Model` - The validation failed because the given server model is not found in the Interop matrix. * `Incompatible-Adapter-Model` - The validation failed because the given adapter model is not found in the Interop matrix. * `Incompatible-Switch-Model` - The validation failed because the given switch model is not found in the Interop matrix. * `Incompatible-Server-Firmware` - The validation failed because the given server firmware version is not found in HCL. * `Incompatible-Switch-Firmware` - The validation failed because the given switch firmware version is not found in Interop matrix. * `Incompatible-Firmware` - The validation failed because the given adapter firmware version is not found in either HCL or Interop matrix. * `Incompatible-Driver` - The validation failed because the given driver version is not found in either HCL or Interop matrix. * `Incompatible-Firmware-Driver` - The validation failed because the given adapter firmware and driver version is not found in either HCL or Interop matrix. * `Missing-Os-Driver-Info` - The validation failed because the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information. * `Missing-Os-Or-Driver-Info` - This means the Interop status for the switch or storage array cannot be computed because we have either missing Host OS information or missing required driver information. Either install ucstools vib or use power shell scripts to tag proper OS information or install required drivers. * `Incompatible-Components` - The validation failed for the given server because one or more of its components failed validation. * `Compatible` - This means the HCL status and Interop status for the component have passed all validations for all of its related components.
	Reason *string `json:"Reason,omitempty"`
	// Compliance status for the component. * `NotEvaluated` - The interoperability compliance for the component has not be checked. * `Approved` - The component is valid as per the interoperability compliance check. * `NotApproved` - The component is not valid as per the interoperability compliance check. * `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data.
	Status               *string `json:"Status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConvergedinfraBaseComplianceDetailsAllOf ConvergedinfraBaseComplianceDetailsAllOf

// NewConvergedinfraBaseComplianceDetailsAllOf instantiates a new ConvergedinfraBaseComplianceDetailsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvergedinfraBaseComplianceDetailsAllOf(classId string, objectType string) *ConvergedinfraBaseComplianceDetailsAllOf {
	this := ConvergedinfraBaseComplianceDetailsAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConvergedinfraBaseComplianceDetailsAllOfWithDefaults instantiates a new ConvergedinfraBaseComplianceDetailsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvergedinfraBaseComplianceDetailsAllOfWithDefaults() *ConvergedinfraBaseComplianceDetailsAllOf {
	this := ConvergedinfraBaseComplianceDetailsAllOf{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConvergedinfraBaseComplianceDetailsAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConvergedinfraBaseComplianceDetailsAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) SetName(v string) {
	o.Name = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConvergedinfraBaseComplianceDetailsAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o ConvergedinfraBaseComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConvergedinfraBaseComplianceDetailsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConvergedinfraBaseComplianceDetailsAllOf := _ConvergedinfraBaseComplianceDetailsAllOf{}

	if err = json.Unmarshal(bytes, &varConvergedinfraBaseComplianceDetailsAllOf); err == nil {
		*o = ConvergedinfraBaseComplianceDetailsAllOf(varConvergedinfraBaseComplianceDetailsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConvergedinfraBaseComplianceDetailsAllOf struct {
	value *ConvergedinfraBaseComplianceDetailsAllOf
	isSet bool
}

func (v NullableConvergedinfraBaseComplianceDetailsAllOf) Get() *ConvergedinfraBaseComplianceDetailsAllOf {
	return v.value
}

func (v *NullableConvergedinfraBaseComplianceDetailsAllOf) Set(val *ConvergedinfraBaseComplianceDetailsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConvergedinfraBaseComplianceDetailsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConvergedinfraBaseComplianceDetailsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvergedinfraBaseComplianceDetailsAllOf(val *ConvergedinfraBaseComplianceDetailsAllOf) *NullableConvergedinfraBaseComplianceDetailsAllOf {
	return &NullableConvergedinfraBaseComplianceDetailsAllOf{value: val, isSet: true}
}

func (v NullableConvergedinfraBaseComplianceDetailsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvergedinfraBaseComplianceDetailsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
