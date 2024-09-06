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

// checks if the VirtualizationVmwarePhysicalNetworkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualizationVmwarePhysicalNetworkInterface{}

// VirtualizationVmwarePhysicalNetworkInterface Details of VMware physical network interface.
type VirtualizationVmwarePhysicalNetworkInterface struct {
	VirtualizationBasePhysicalNetworkInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Driver type associated with physical network interface.
	Driver *string `json:"Driver,omitempty"`
	// Link speed of the physical network interface.
	LinkSpeed *int32 `json:"LinkSpeed,omitempty"`
	// Standard MAC address assigned to physical network interface.
	MacAddress *string `json:"MacAddress,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`
	// PCI info for physical network interface.
	Pci *string `json:"Pci,omitempty"`
	// Switch associated with the physical network interface.
	SwitchName           *string                                      `json:"SwitchName,omitempty"`
	Host                 NullableVirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualizationVmwarePhysicalNetworkInterface VirtualizationVmwarePhysicalNetworkInterface

// NewVirtualizationVmwarePhysicalNetworkInterface instantiates a new VirtualizationVmwarePhysicalNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualizationVmwarePhysicalNetworkInterface(classId string, objectType string) *VirtualizationVmwarePhysicalNetworkInterface {
	this := VirtualizationVmwarePhysicalNetworkInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVirtualizationVmwarePhysicalNetworkInterfaceWithDefaults instantiates a new VirtualizationVmwarePhysicalNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualizationVmwarePhysicalNetworkInterfaceWithDefaults() *VirtualizationVmwarePhysicalNetworkInterface {
	this := VirtualizationVmwarePhysicalNetworkInterface{}
	var classId string = "virtualization.VmwarePhysicalNetworkInterface"
	this.ClassId = classId
	var objectType string = "virtualization.VmwarePhysicalNetworkInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "virtualization.VmwarePhysicalNetworkInterface" of the ClassId field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDefaultClassId() interface{} {
	return "virtualization.VmwarePhysicalNetworkInterface"
}

// GetObjectType returns the ObjectType field value
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "virtualization.VmwarePhysicalNetworkInterface" of the ObjectType field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDefaultObjectType() interface{} {
	return "virtualization.VmwarePhysicalNetworkInterface"
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDriver() string {
	if o == nil || IsNil(o.Driver) {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDriverOk() (*string, bool) {
	if o == nil || IsNil(o.Driver) {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasDriver() bool {
	if o != nil && !IsNil(o.Driver) {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetDriver(v string) {
	o.Driver = &v
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetLinkSpeed() int32 {
	if o == nil || IsNil(o.LinkSpeed) {
		var ret int32
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetLinkSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasLinkSpeed() bool {
	if o != nil && !IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given int32 and assigns it to the LinkSpeed field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetLinkSpeed(v int32) {
	o.LinkSpeed = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetPci returns the Pci field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetPci() string {
	if o == nil || IsNil(o.Pci) {
		var ret string
		return ret
	}
	return *o.Pci
}

// GetPciOk returns a tuple with the Pci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetPciOk() (*string, bool) {
	if o == nil || IsNil(o.Pci) {
		return nil, false
	}
	return o.Pci, true
}

// HasPci returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasPci() bool {
	if o != nil && !IsNil(o.Pci) {
		return true
	}

	return false
}

// SetPci gets a reference to the given string and assigns it to the Pci field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetPci(v string) {
	o.Pci = &v
}

// GetSwitchName returns the SwitchName field value if set, zero value otherwise.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetSwitchName() string {
	if o == nil || IsNil(o.SwitchName) {
		var ret string
		return ret
	}
	return *o.SwitchName
}

// GetSwitchNameOk returns a tuple with the SwitchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetSwitchNameOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchName) {
		return nil, false
	}
	return o.SwitchName, true
}

// HasSwitchName returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasSwitchName() bool {
	if o != nil && !IsNil(o.SwitchName) {
		return true
	}

	return false
}

// SetSwitchName gets a reference to the given string and assigns it to the SwitchName field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetSwitchName(v string) {
	o.SwitchName = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetHost() VirtualizationVmwareHostRelationship {
	if o == nil || IsNil(o.Host.Get()) {
		var ret VirtualizationVmwareHostRelationship
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualizationVmwarePhysicalNetworkInterface) GetHostOk() (*VirtualizationVmwareHostRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *VirtualizationVmwarePhysicalNetworkInterface) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableVirtualizationVmwareHostRelationship and assigns it to the Host field.
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetHost(v VirtualizationVmwareHostRelationship) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *VirtualizationVmwarePhysicalNetworkInterface) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *VirtualizationVmwarePhysicalNetworkInterface) UnsetHost() {
	o.Host.Unset()
}

func (o VirtualizationVmwarePhysicalNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualizationVmwarePhysicalNetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVirtualizationBasePhysicalNetworkInterface, errVirtualizationBasePhysicalNetworkInterface := json.Marshal(o.VirtualizationBasePhysicalNetworkInterface)
	if errVirtualizationBasePhysicalNetworkInterface != nil {
		return map[string]interface{}{}, errVirtualizationBasePhysicalNetworkInterface
	}
	errVirtualizationBasePhysicalNetworkInterface = json.Unmarshal([]byte(serializedVirtualizationBasePhysicalNetworkInterface), &toSerialize)
	if errVirtualizationBasePhysicalNetworkInterface != nil {
		return map[string]interface{}{}, errVirtualizationBasePhysicalNetworkInterface
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Driver) {
		toSerialize["Driver"] = o.Driver
	}
	if !IsNil(o.LinkSpeed) {
		toSerialize["LinkSpeed"] = o.LinkSpeed
	}
	if !IsNil(o.MacAddress) {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if !IsNil(o.Pci) {
		toSerialize["Pci"] = o.Pci
	}
	if !IsNil(o.SwitchName) {
		toSerialize["SwitchName"] = o.SwitchName
	}
	if o.Host.IsSet() {
		toSerialize["Host"] = o.Host.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualizationVmwarePhysicalNetworkInterface) UnmarshalJSON(data []byte) (err error) {
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
	type VirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Driver type associated with physical network interface.
		Driver *string `json:"Driver,omitempty"`
		// Link speed of the physical network interface.
		LinkSpeed *int32 `json:"LinkSpeed,omitempty"`
		// Standard MAC address assigned to physical network interface.
		MacAddress *string `json:"MacAddress,omitempty" validate:"regexp=^$|^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"`
		// PCI info for physical network interface.
		Pci *string `json:"Pci,omitempty"`
		// Switch associated with the physical network interface.
		SwitchName *string                                      `json:"SwitchName,omitempty"`
		Host       NullableVirtualizationVmwareHostRelationship `json:"Host,omitempty"`
	}

	varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct := VirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varVirtualizationVmwarePhysicalNetworkInterface := _VirtualizationVmwarePhysicalNetworkInterface{}
		varVirtualizationVmwarePhysicalNetworkInterface.ClassId = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.ClassId
		varVirtualizationVmwarePhysicalNetworkInterface.ObjectType = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.ObjectType
		varVirtualizationVmwarePhysicalNetworkInterface.Driver = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.Driver
		varVirtualizationVmwarePhysicalNetworkInterface.LinkSpeed = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.LinkSpeed
		varVirtualizationVmwarePhysicalNetworkInterface.MacAddress = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.MacAddress
		varVirtualizationVmwarePhysicalNetworkInterface.Pci = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.Pci
		varVirtualizationVmwarePhysicalNetworkInterface.SwitchName = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.SwitchName
		varVirtualizationVmwarePhysicalNetworkInterface.Host = varVirtualizationVmwarePhysicalNetworkInterfaceWithoutEmbeddedStruct.Host
		*o = VirtualizationVmwarePhysicalNetworkInterface(varVirtualizationVmwarePhysicalNetworkInterface)
	} else {
		return err
	}

	varVirtualizationVmwarePhysicalNetworkInterface := _VirtualizationVmwarePhysicalNetworkInterface{}

	err = json.Unmarshal(data, &varVirtualizationVmwarePhysicalNetworkInterface)
	if err == nil {
		o.VirtualizationBasePhysicalNetworkInterface = varVirtualizationVmwarePhysicalNetworkInterface.VirtualizationBasePhysicalNetworkInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Driver")
		delete(additionalProperties, "LinkSpeed")
		delete(additionalProperties, "MacAddress")
		delete(additionalProperties, "Pci")
		delete(additionalProperties, "SwitchName")
		delete(additionalProperties, "Host")

		// remove fields from embedded structs
		reflectVirtualizationBasePhysicalNetworkInterface := reflect.ValueOf(o.VirtualizationBasePhysicalNetworkInterface)
		for i := 0; i < reflectVirtualizationBasePhysicalNetworkInterface.Type().NumField(); i++ {
			t := reflectVirtualizationBasePhysicalNetworkInterface.Type().Field(i)

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

type NullableVirtualizationVmwarePhysicalNetworkInterface struct {
	value *VirtualizationVmwarePhysicalNetworkInterface
	isSet bool
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterface) Get() *VirtualizationVmwarePhysicalNetworkInterface {
	return v.value
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterface) Set(val *VirtualizationVmwarePhysicalNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualizationVmwarePhysicalNetworkInterface(val *VirtualizationVmwarePhysicalNetworkInterface) *NullableVirtualizationVmwarePhysicalNetworkInterface {
	return &NullableVirtualizationVmwarePhysicalNetworkInterface{value: val, isSet: true}
}

func (v NullableVirtualizationVmwarePhysicalNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualizationVmwarePhysicalNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
