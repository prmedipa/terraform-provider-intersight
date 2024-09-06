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

// checks if the SolPolicyInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolPolicyInventory{}

// SolPolicyInventory Policy for configuring Serial Over LAN settings on endpoint.
type SolPolicyInventory struct {
	PolicyAbstractPolicyInventory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Baud Rate used for Serial Over LAN communication. * `9600` - Use baud rate 9600 for communication. * `19200` - Use baud rate 19200 for communication. * `38400` - Use baud rate 38400 for communication. * `57600` - Use baud rate 57600 for communication. * `115200` - Use baud rate 115200 for communication.
	BaudRate *int32 `json:"BaudRate,omitempty"`
	// Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * `com0` - Use serial port com0 for communication. * `com1` - Use serial port com1 for communication.
	ComPort *string `json:"ComPort,omitempty"`
	// State of Serial Over LAN service on the endpoint.
	Enabled *bool `json:"Enabled,omitempty"`
	// SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
	SshPort              *int64                       `json:"SshPort,omitempty"`
	TargetMo             NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SolPolicyInventory SolPolicyInventory

// NewSolPolicyInventory instantiates a new SolPolicyInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolPolicyInventory(classId string, objectType string) *SolPolicyInventory {
	this := SolPolicyInventory{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSolPolicyInventoryWithDefaults instantiates a new SolPolicyInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolPolicyInventoryWithDefaults() *SolPolicyInventory {
	this := SolPolicyInventory{}
	var classId string = "sol.PolicyInventory"
	this.ClassId = classId
	var objectType string = "sol.PolicyInventory"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SolPolicyInventory) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SolPolicyInventory) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "sol.PolicyInventory" of the ClassId field.
func (o *SolPolicyInventory) GetDefaultClassId() interface{} {
	return "sol.PolicyInventory"
}

// GetObjectType returns the ObjectType field value
func (o *SolPolicyInventory) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SolPolicyInventory) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "sol.PolicyInventory" of the ObjectType field.
func (o *SolPolicyInventory) GetDefaultObjectType() interface{} {
	return "sol.PolicyInventory"
}

// GetBaudRate returns the BaudRate field value if set, zero value otherwise.
func (o *SolPolicyInventory) GetBaudRate() int32 {
	if o == nil || IsNil(o.BaudRate) {
		var ret int32
		return ret
	}
	return *o.BaudRate
}

// GetBaudRateOk returns a tuple with the BaudRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetBaudRateOk() (*int32, bool) {
	if o == nil || IsNil(o.BaudRate) {
		return nil, false
	}
	return o.BaudRate, true
}

// HasBaudRate returns a boolean if a field has been set.
func (o *SolPolicyInventory) HasBaudRate() bool {
	if o != nil && !IsNil(o.BaudRate) {
		return true
	}

	return false
}

// SetBaudRate gets a reference to the given int32 and assigns it to the BaudRate field.
func (o *SolPolicyInventory) SetBaudRate(v int32) {
	o.BaudRate = &v
}

// GetComPort returns the ComPort field value if set, zero value otherwise.
func (o *SolPolicyInventory) GetComPort() string {
	if o == nil || IsNil(o.ComPort) {
		var ret string
		return ret
	}
	return *o.ComPort
}

// GetComPortOk returns a tuple with the ComPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetComPortOk() (*string, bool) {
	if o == nil || IsNil(o.ComPort) {
		return nil, false
	}
	return o.ComPort, true
}

// HasComPort returns a boolean if a field has been set.
func (o *SolPolicyInventory) HasComPort() bool {
	if o != nil && !IsNil(o.ComPort) {
		return true
	}

	return false
}

// SetComPort gets a reference to the given string and assigns it to the ComPort field.
func (o *SolPolicyInventory) SetComPort(v string) {
	o.ComPort = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SolPolicyInventory) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SolPolicyInventory) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SolPolicyInventory) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *SolPolicyInventory) GetSshPort() int64 {
	if o == nil || IsNil(o.SshPort) {
		var ret int64
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolPolicyInventory) GetSshPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *SolPolicyInventory) HasSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int64 and assigns it to the SshPort field.
func (o *SolPolicyInventory) SetSshPort(v int64) {
	o.SshPort = &v
}

// GetTargetMo returns the TargetMo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SolPolicyInventory) GetTargetMo() MoBaseMoRelationship {
	if o == nil || IsNil(o.TargetMo.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.TargetMo.Get()
}

// GetTargetMoOk returns a tuple with the TargetMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SolPolicyInventory) GetTargetMoOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetMo.Get(), o.TargetMo.IsSet()
}

// HasTargetMo returns a boolean if a field has been set.
func (o *SolPolicyInventory) HasTargetMo() bool {
	if o != nil && o.TargetMo.IsSet() {
		return true
	}

	return false
}

// SetTargetMo gets a reference to the given NullableMoBaseMoRelationship and assigns it to the TargetMo field.
func (o *SolPolicyInventory) SetTargetMo(v MoBaseMoRelationship) {
	o.TargetMo.Set(&v)
}

// SetTargetMoNil sets the value for TargetMo to be an explicit nil
func (o *SolPolicyInventory) SetTargetMoNil() {
	o.TargetMo.Set(nil)
}

// UnsetTargetMo ensures that no value is present for TargetMo, not even an explicit nil
func (o *SolPolicyInventory) UnsetTargetMo() {
	o.TargetMo.Unset()
}

func (o SolPolicyInventory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolPolicyInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicyInventory, errPolicyAbstractPolicyInventory := json.Marshal(o.PolicyAbstractPolicyInventory)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	errPolicyAbstractPolicyInventory = json.Unmarshal([]byte(serializedPolicyAbstractPolicyInventory), &toSerialize)
	if errPolicyAbstractPolicyInventory != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicyInventory
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.BaudRate) {
		toSerialize["BaudRate"] = o.BaudRate
	}
	if !IsNil(o.ComPort) {
		toSerialize["ComPort"] = o.ComPort
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.SshPort) {
		toSerialize["SshPort"] = o.SshPort
	}
	if o.TargetMo.IsSet() {
		toSerialize["TargetMo"] = o.TargetMo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SolPolicyInventory) UnmarshalJSON(data []byte) (err error) {
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
	type SolPolicyInventoryWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Baud Rate used for Serial Over LAN communication. * `9600` - Use baud rate 9600 for communication. * `19200` - Use baud rate 19200 for communication. * `38400` - Use baud rate 38400 for communication. * `57600` - Use baud rate 57600 for communication. * `115200` - Use baud rate 115200 for communication.
		BaudRate *int32 `json:"BaudRate,omitempty"`
		// Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default. * `com0` - Use serial port com0 for communication. * `com1` - Use serial port com1 for communication.
		ComPort *string `json:"ComPort,omitempty"`
		// State of Serial Over LAN service on the endpoint.
		Enabled *bool `json:"Enabled,omitempty"`
		// SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
		SshPort  *int64                       `json:"SshPort,omitempty"`
		TargetMo NullableMoBaseMoRelationship `json:"TargetMo,omitempty"`
	}

	varSolPolicyInventoryWithoutEmbeddedStruct := SolPolicyInventoryWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSolPolicyInventoryWithoutEmbeddedStruct)
	if err == nil {
		varSolPolicyInventory := _SolPolicyInventory{}
		varSolPolicyInventory.ClassId = varSolPolicyInventoryWithoutEmbeddedStruct.ClassId
		varSolPolicyInventory.ObjectType = varSolPolicyInventoryWithoutEmbeddedStruct.ObjectType
		varSolPolicyInventory.BaudRate = varSolPolicyInventoryWithoutEmbeddedStruct.BaudRate
		varSolPolicyInventory.ComPort = varSolPolicyInventoryWithoutEmbeddedStruct.ComPort
		varSolPolicyInventory.Enabled = varSolPolicyInventoryWithoutEmbeddedStruct.Enabled
		varSolPolicyInventory.SshPort = varSolPolicyInventoryWithoutEmbeddedStruct.SshPort
		varSolPolicyInventory.TargetMo = varSolPolicyInventoryWithoutEmbeddedStruct.TargetMo
		*o = SolPolicyInventory(varSolPolicyInventory)
	} else {
		return err
	}

	varSolPolicyInventory := _SolPolicyInventory{}

	err = json.Unmarshal(data, &varSolPolicyInventory)
	if err == nil {
		o.PolicyAbstractPolicyInventory = varSolPolicyInventory.PolicyAbstractPolicyInventory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BaudRate")
		delete(additionalProperties, "ComPort")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "SshPort")
		delete(additionalProperties, "TargetMo")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicyInventory := reflect.ValueOf(o.PolicyAbstractPolicyInventory)
		for i := 0; i < reflectPolicyAbstractPolicyInventory.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicyInventory.Type().Field(i)

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

type NullableSolPolicyInventory struct {
	value *SolPolicyInventory
	isSet bool
}

func (v NullableSolPolicyInventory) Get() *SolPolicyInventory {
	return v.value
}

func (v *NullableSolPolicyInventory) Set(val *SolPolicyInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableSolPolicyInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableSolPolicyInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolPolicyInventory(val *SolPolicyInventory) *NullableSolPolicyInventory {
	return &NullableSolPolicyInventory{value: val, isSet: true}
}

func (v NullableSolPolicyInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolPolicyInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
