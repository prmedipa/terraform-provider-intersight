/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-7078
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// VirtualizationIweDatacenter A datacenter object in Intersight Workload Engine. It is a pre-defined object created internally by the system which acts as a container (logically) for all other objects (Host, VirtualMachine, Volume etc).
type VirtualizationIweDatacenter struct {
	VirtualizationBaseDatacenter
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                            `json:"ObjectType"`
	Account              *IamAccountRelationship                           `json:"Account,omitempty"`
	HypervisorManager    *VirtualizationCiscoHypervisorManagerRelationship `json:"HypervisorManager,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationIweDatacenter VirtualizationIweDatacenter

// NewVirtualizationIweDatacenter instantiates a new VirtualizationIweDatacenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationIweDatacenter(classId string, objectType string) *VirtualizationIweDatacenter {
	this := VirtualizationIweDatacenter{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationIweDatacenterWithDefaults instantiates a new VirtualizationIweDatacenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationIweDatacenterWithDefaults() *VirtualizationIweDatacenter {
	this := VirtualizationIweDatacenter{}
	var classId string = "virtualization.IweDatacenter"
	this.ClassId = classId
	var objectType string = "virtualization.IweDatacenter"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationIweDatacenter) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDatacenter) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationIweDatacenter) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationIweDatacenter) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDatacenter) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationIweDatacenter) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *VirtualizationIweDatacenter) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDatacenter) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *VirtualizationIweDatacenter) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *VirtualizationIweDatacenter) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetHypervisorManager returns the HypervisorManager field value if set, zero value otherwise.
func (o *VirtualizationIweDatacenter) GetHypervisorManager() VirtualizationCiscoHypervisorManagerRelationship {
	if o == nil || o.HypervisorManager == nil {
		var ret VirtualizationCiscoHypervisorManagerRelationship
		return ret
	}
	return *o.HypervisorManager
}

// GetHypervisorManagerOk returns a tuple with the HypervisorManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationIweDatacenter) GetHypervisorManagerOk() (*VirtualizationCiscoHypervisorManagerRelationship, bool) {
	if o == nil || o.HypervisorManager == nil {
		return nil, false
	}
	return o.HypervisorManager, true
}

// HasHypervisorManager returns a boolean if a field has been set.
func (o *VirtualizationIweDatacenter) HasHypervisorManager() bool {
	if o != nil && o.HypervisorManager != nil {
		return true
	}

	return false
}

// SetHypervisorManager gets a reference to the given VirtualizationCiscoHypervisorManagerRelationship and assigns it to the HypervisorManager field.
func (o *VirtualizationIweDatacenter) SetHypervisorManager(v VirtualizationCiscoHypervisorManagerRelationship) {
	o.HypervisorManager = &v
}

func (o VirtualizationIweDatacenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBaseDatacenter, errVirtualizationBaseDatacenter := json.Marshal(o.VirtualizationBaseDatacenter)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	errVirtualizationBaseDatacenter = json.Unmarshal([]byte(serializedVirtualizationBaseDatacenter), &toSerialize)
	if errVirtualizationBaseDatacenter != nil {
		return []byte{}, errVirtualizationBaseDatacenter
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.HypervisorManager != nil {
		toSerialize["HypervisorManager"] = o.HypervisorManager
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VirtualizationIweDatacenter) UnmarshalJSON(bytes []byte) (err error) {
	type VirtualizationIweDatacenterWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType        string                                            `json:"ObjectType"`
		Account           *IamAccountRelationship                           `json:"Account,omitempty"`
		HypervisorManager *VirtualizationCiscoHypervisorManagerRelationship `json:"HypervisorManager,omitempty"`
	}

	varVirtualizationIweDatacenterWithoutEmbeddedStruct := VirtualizationIweDatacenterWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDatacenterWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationIweDatacenter := _VirtualizationIweDatacenter{}
		varVirtualizationIweDatacenter.ClassId = varVirtualizationIweDatacenterWithoutEmbeddedStruct.ClassId
		varVirtualizationIweDatacenter.ObjectType = varVirtualizationIweDatacenterWithoutEmbeddedStruct.ObjectType
		varVirtualizationIweDatacenter.Account = varVirtualizationIweDatacenterWithoutEmbeddedStruct.Account
		varVirtualizationIweDatacenter.HypervisorManager = varVirtualizationIweDatacenterWithoutEmbeddedStruct.HypervisorManager
		*o = VirtualizationIweDatacenter(varVirtualizationIweDatacenter)
	} else {
		return err
	}

	varVirtualizationIweDatacenter := _VirtualizationIweDatacenter{}

	err = json.Unmarshal(bytes, &varVirtualizationIweDatacenter)
	if err == nil {
		o.VirtualizationBaseDatacenter = varVirtualizationIweDatacenter.VirtualizationBaseDatacenter
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Account")
		delete(additionalProperties, "HypervisorManager")

		// remove fields from embedded structs
		reflectVirtualizationBaseDatacenter := reflect.ValueOf(o.VirtualizationBaseDatacenter)
		for i := 0; i < reflectVirtualizationBaseDatacenter.Type().NumField(); i++ {
			t := reflectVirtualizationBaseDatacenter.Type().Field(i)

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

type NullableVirtualizationIweDatacenter struct {
	value *VirtualizationIweDatacenter
	isSet bool
}

func (v NullableVirtualizationIweDatacenter) Get() *VirtualizationIweDatacenter {
	return v.value
}

func (v *NullableVirtualizationIweDatacenter) Set(val *VirtualizationIweDatacenter) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationIweDatacenter) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationIweDatacenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationIweDatacenter(val *VirtualizationIweDatacenter) *NullableVirtualizationIweDatacenter {
	return &NullableVirtualizationIweDatacenter{value: val, isSet: true}
}

func (v NullableVirtualizationIweDatacenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationIweDatacenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}