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

// VnicFcVhbaPolicyInventory A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
type VnicFcVhbaPolicyInventory struct {
	PolicyAbstractInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
	MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
	// Name of the virtual HBA interface.
	Name *string `json:"Name,omitempty"`
	// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority             *string               `json:"Priority,omitempty"`
	TargetMo             *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcVhbaPolicyInventory VnicFcVhbaPolicyInventory

// NewVnicFcVhbaPolicyInventory instantiates a new VnicFcVhbaPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcVhbaPolicyInventory(classId string, objectType string) *VnicFcVhbaPolicyInventory {
	this := VnicFcVhbaPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	return &this
}

// NewVnicFcVhbaPolicyInventoryWithDefaults instantiates a new VnicFcVhbaPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcVhbaPolicyInventoryWithDefaults() *VnicFcVhbaPolicyInventory {
	this := VnicFcVhbaPolicyInventory{}
	var classId string = "vnic.FcVhbaPolicyInventory"
	this.ClassId = classId
	var objectType string = "vnic.FcVhbaPolicyInventory"
	this.ObjectType = objectType
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcVhbaPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcVhbaPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcVhbaPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcVhbaPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicFcVhbaPolicyInventory) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicFcVhbaPolicyInventory) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicFcVhbaPolicyInventory) SetCos(v int64) {
	o.Cos = &v
}

// GetMaxDataFieldSize returns the MaxDataFieldSize field value if set, zero value otherwise.
func (o *VnicFcVhbaPolicyInventory) GetMaxDataFieldSize() int64 {
	if o == nil || o.MaxDataFieldSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxDataFieldSize
}

// GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetMaxDataFieldSizeOk() (*int64, bool) {
	if o == nil || o.MaxDataFieldSize == nil {
		return nil, false
	}
	return o.MaxDataFieldSize, true
}

// HasMaxDataFieldSize returns a boolean if a field has been set.
func (o *VnicFcVhbaPolicyInventory) HasMaxDataFieldSize() bool {
	if o != nil && o.MaxDataFieldSize != nil {
		return true
	}

	return false
}

// SetMaxDataFieldSize gets a reference to the given int64 and assigns it to the MaxDataFieldSize field.
func (o *VnicFcVhbaPolicyInventory) SetMaxDataFieldSize(v int64) {
	o.MaxDataFieldSize = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicFcVhbaPolicyInventory) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicFcVhbaPolicyInventory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicFcVhbaPolicyInventory) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicFcVhbaPolicyInventory) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicFcVhbaPolicyInventory) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicFcVhbaPolicyInventory) SetPriority(v string) {
	o.Priority = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise.
func (o *VnicFcVhbaPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || o.TargetMo == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcVhbaPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.TargetMo == nil {
		return nil, false
	}
	return o.TargetMo, true
}

// HasTargetMo returns a boolean if a field has been set.
func (o *VnicFcVhbaPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo != nil {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given MoBaseMoRelationship and assigns it to the TargetMo field.
func (o *VnicFcVhbaPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo = &v
}

func (o VnicFcVhbaPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractInventory, errPolicyAbstractInventory := json.Marshal(o.PolicyAbstractInventory)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	errPolicyAbstractInventory = json.Unmarshal([]byte(serializedPolicyAbstractInventory), &toSerialize)
	if errPolicyAbstractInventory != nil {
		return []byte{}, errPolicyAbstractInventory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.MaxDataFieldSize != nil {
		toSerialize["MaxDataFieldSize"] = o.MaxDataFieldSize
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.TargetMo != nil {
		toSerialize["TargetMo"] = o.TargetMo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcVhbaPolicyInventory) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcVhbaPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Class of Service to be associated to the traffic on the virtual interface.
		Cos *int64 `json:"Cos,omitempty"`
		// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
		MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
		// Name of the virtual HBA interface.
		Name *string `json:"Name,omitempty"`
		// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority *string               `json:"Priority,omitempty"`
		TargetMo *MoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct := VnicFcVhbaPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcVhbaPolicyInventory := _VnicFcVhbaPolicyInventory{}
		varVnicFcVhbaPolicyInventory.ClassId = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.ClassId
		varVnicFcVhbaPolicyInventory.ObjectType = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varVnicFcVhbaPolicyInventory.Cos = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.Cos
		varVnicFcVhbaPolicyInventory.MaxDataFieldSize = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.MaxDataFieldSize
		varVnicFcVhbaPolicyInventory.Name = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.Name
		varVnicFcVhbaPolicyInventory.Priority = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.Priority
		varVnicFcVhbaPolicyInventory.TargetMo = varVnicFcVhbaPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = VnicFcVhbaPolicyInventory(varVnicFcVhbaPolicyInventory)
	} else {
		return err
	}

	varVnicFcVhbaPolicyInventory := _VnicFcVhbaPolicyInventory{}

	err = json.Unmarshal(bytes, &varVnicFcVhbaPolicyInventory)
	if err == nil {
		o.PolicyAbstractInventory = varVnicFcVhbaPolicyInventory.PolicyAbstractInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "MaxDataFieldSize")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractInventory := reflect.ValueOf(o.PolicyAbstractInventory)
		for i := 0; i < reflectPolicyAbstractInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractInventory.Type().Field(i)

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

type NullableVnicFcVhbaPolicyInventory struct {
	value *VnicFcVhbaPolicyInventory
	isSet bool
}

func (v NullableVnicFcVhbaPolicyInventory) Get() *VnicFcVhbaPolicyInventory {
	return v.value
}

func (v *NullableVnicFcVhbaPolicyInventory) Set(val *VnicFcVhbaPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcVhbaPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcVhbaPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcVhbaPolicyInventory(val *VnicFcVhbaPolicyInventory) *NullableVnicFcVhbaPolicyInventory {
	return &NullableVnicFcVhbaPolicyInventory{value: val, isSet: true}
}

func (v NullableVnicFcVhbaPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcVhbaPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}