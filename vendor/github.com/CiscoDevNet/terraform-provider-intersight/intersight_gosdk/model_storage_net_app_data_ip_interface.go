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

// StorageNetAppDataIpInterface NetApp Data IP interface is a logical interface for data within the svm scope.
type StorageNetAppDataIpInterface struct {
	StorageNetAppBaseIpInterface
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                         `json:"ObjectType"`
	ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
	// An array of relationships to storageNetAppDataIpInterfaceEvent resources.
	Events               []StorageNetAppDataIpInterfaceEventRelationship `json:"Events,omitempty"`
	NetAppEthernetPort   *StorageNetAppEthernetPortRelationship          `json:"NetAppEthernetPort,omitempty"`
	Tenant               *StorageNetAppStorageVmRelationship             `json:"Tenant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppDataIpInterface StorageNetAppDataIpInterface

// NewStorageNetAppDataIpInterface instantiates a new StorageNetAppDataIpInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppDataIpInterface(classId string, objectType string) *StorageNetAppDataIpInterface {
	this := StorageNetAppDataIpInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppDataIpInterfaceWithDefaults instantiates a new StorageNetAppDataIpInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppDataIpInterfaceWithDefaults() *StorageNetAppDataIpInterface {
	this := StorageNetAppDataIpInterface{}
	var classId string = "storage.NetAppDataIpInterface"
	this.ClassId = classId
	var objectType string = "storage.NetAppDataIpInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppDataIpInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppDataIpInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppDataIpInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppDataIpInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayController returns the ArrayController field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterface) GetArrayController() StorageNetAppNodeRelationship {
	if o == nil || o.ArrayController == nil {
		var ret StorageNetAppNodeRelationship
		return ret
	}
	return *o.ArrayController
}

// GetArrayControllerOk returns a tuple with the ArrayController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool) {
	if o == nil || o.ArrayController == nil {
		return nil, false
	}
	return o.ArrayController, true
}

// HasArrayController returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterface) HasArrayController() bool {
	if o != nil && o.ArrayController != nil {
		return true
	}

	return false
}

// SetArrayController gets a reference to the given StorageNetAppNodeRelationship and assigns it to the ArrayController field.
func (o *StorageNetAppDataIpInterface) SetArrayController(v StorageNetAppNodeRelationship) {
	o.ArrayController = &v
}

// GetEvents returns the Events field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageNetAppDataIpInterface) GetEvents() []StorageNetAppDataIpInterfaceEventRelationship {
	if o == nil {
		var ret []StorageNetAppDataIpInterfaceEventRelationship
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageNetAppDataIpInterface) GetEventsOk() ([]StorageNetAppDataIpInterfaceEventRelationship, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterface) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []StorageNetAppDataIpInterfaceEventRelationship and assigns it to the Events field.
func (o *StorageNetAppDataIpInterface) SetEvents(v []StorageNetAppDataIpInterfaceEventRelationship) {
	o.Events = v
}

// GetNetAppEthernetPort returns the NetAppEthernetPort field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship {
	if o == nil || o.NetAppEthernetPort == nil {
		var ret StorageNetAppEthernetPortRelationship
		return ret
	}
	return *o.NetAppEthernetPort
}

// GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool) {
	if o == nil || o.NetAppEthernetPort == nil {
		return nil, false
	}
	return o.NetAppEthernetPort, true
}

// HasNetAppEthernetPort returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterface) HasNetAppEthernetPort() bool {
	if o != nil && o.NetAppEthernetPort != nil {
		return true
	}

	return false
}

// SetNetAppEthernetPort gets a reference to the given StorageNetAppEthernetPortRelationship and assigns it to the NetAppEthernetPort field.
func (o *StorageNetAppDataIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship) {
	o.NetAppEthernetPort = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *StorageNetAppDataIpInterface) GetTenant() StorageNetAppStorageVmRelationship {
	if o == nil || o.Tenant == nil {
		var ret StorageNetAppStorageVmRelationship
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppDataIpInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *StorageNetAppDataIpInterface) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given StorageNetAppStorageVmRelationship and assigns it to the Tenant field.
func (o *StorageNetAppDataIpInterface) SetTenant(v StorageNetAppStorageVmRelationship) {
	o.Tenant = &v
}

func (o StorageNetAppDataIpInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageNetAppBaseIpInterface, errStorageNetAppBaseIpInterface := json.Marshal(o.StorageNetAppBaseIpInterface)
	if errStorageNetAppBaseIpInterface != nil {
		return []byte{}, errStorageNetAppBaseIpInterface
	}
	errStorageNetAppBaseIpInterface = json.Unmarshal([]byte(serializedStorageNetAppBaseIpInterface), &toSerialize)
	if errStorageNetAppBaseIpInterface != nil {
		return []byte{}, errStorageNetAppBaseIpInterface
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayController != nil {
		toSerialize["ArrayController"] = o.ArrayController
	}
	if o.Events != nil {
		toSerialize["Events"] = o.Events
	}
	if o.NetAppEthernetPort != nil {
		toSerialize["NetAppEthernetPort"] = o.NetAppEthernetPort
	}
	if o.Tenant != nil {
		toSerialize["Tenant"] = o.Tenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppDataIpInterface) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppDataIpInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                         `json:"ObjectType"`
		ArrayController *StorageNetAppNodeRelationship `json:"ArrayController,omitempty"`
		// An array of relationships to storageNetAppDataIpInterfaceEvent resources.
		Events             []StorageNetAppDataIpInterfaceEventRelationship `json:"Events,omitempty"`
		NetAppEthernetPort *StorageNetAppEthernetPortRelationship          `json:"NetAppEthernetPort,omitempty"`
		Tenant             *StorageNetAppStorageVmRelationship             `json:"Tenant,omitempty"`
	}

	varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct := StorageNetAppDataIpInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppDataIpInterface := _StorageNetAppDataIpInterface{}
		varStorageNetAppDataIpInterface.ClassId = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.ClassId
		varStorageNetAppDataIpInterface.ObjectType = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.ObjectType
		varStorageNetAppDataIpInterface.ArrayController = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.ArrayController
		varStorageNetAppDataIpInterface.Events = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.Events
		varStorageNetAppDataIpInterface.NetAppEthernetPort = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.NetAppEthernetPort
		varStorageNetAppDataIpInterface.Tenant = varStorageNetAppDataIpInterfaceWithoutEmbeddedStruct.Tenant
		*o = StorageNetAppDataIpInterface(varStorageNetAppDataIpInterface)
	} else {
		return err
	}

	varStorageNetAppDataIpInterface := _StorageNetAppDataIpInterface{}

	err = json.Unmarshal(bytes, &varStorageNetAppDataIpInterface)
	if err == nil {
		o.StorageNetAppBaseIpInterface = varStorageNetAppDataIpInterface.StorageNetAppBaseIpInterface
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayController")
		delete(additionalProperties, "Events")
		delete(additionalProperties, "NetAppEthernetPort")
		delete(additionalProperties, "Tenant")

		// remove fields from embedded structs
		reflectStorageNetAppBaseIpInterface := reflect.ValueOf(o.StorageNetAppBaseIpInterface)
		for i := 0; i < reflectStorageNetAppBaseIpInterface.Type().NumField(); i++ {
			t := reflectStorageNetAppBaseIpInterface.Type().Field(i)

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

type NullableStorageNetAppDataIpInterface struct {
	value *StorageNetAppDataIpInterface
	isSet bool
}

func (v NullableStorageNetAppDataIpInterface) Get() *StorageNetAppDataIpInterface {
	return v.value
}

func (v *NullableStorageNetAppDataIpInterface) Set(val *StorageNetAppDataIpInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppDataIpInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppDataIpInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppDataIpInterface(val *StorageNetAppDataIpInterface) *NullableStorageNetAppDataIpInterface {
	return &NullableStorageNetAppDataIpInterface{value: val, isSet: true}
}

func (v NullableStorageNetAppDataIpInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppDataIpInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}