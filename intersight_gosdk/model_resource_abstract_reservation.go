/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2024100405
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the ResourceAbstractReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceAbstractReservation{}

// ResourceAbstractReservation You can reserve resources for use cases related to Pools or Groups. For example, resources that are  decommissioned or reserved for future allocation are maintained in the Reservation model, where  AbstractReservation serves as the base class for reservations.
type ResourceAbstractReservation struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Details of the use case for which the reservation was created, such as decommissioning.
	Description *string `json:"Description,omitempty"`
	// The resource reservation includes an expiration date and a timestamp indicating when this management object will be cleared. The expiration date is set during the decommissioning process and is maintained for a period of 3 months.
	Expiration *time.Time `json:"Expiration,omitempty"`
	// The unique identification of the resource is based on the resource OData string, which is mentioned as part of the ReservationSelector. For example, 'Serial eq 'EM6259AE6B'.
	ReservationSelector *string `json:"ReservationSelector,omitempty"`
	// The type of resource that is placed into resource groups or pools. Resource Type can be either 'compute.Blade' or 'compute.RackUnit' for pools.
	ResourceType *string `json:"ResourceType,omitempty"`
	// The reservation status can be in the 'Created', 'Processing', 'Failed', or 'Finished' state. * `Created` - By default, a reservation is in Created status. * `Processing` - A reservation is changed to Processing status for appliance mode resource claim requests. * `Failed` - A reservation is changed to Failed status if the validations on resources, resource groups fails. * `Finished` - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters.
	Status               *string                      `json:"Status,omitempty"`
	Identity             NullableMoBaseMoRelationship `json:"Identity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceAbstractReservation ResourceAbstractReservation

// NewResourceAbstractReservation instantiates a new ResourceAbstractReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAbstractReservation(classId string, objectType string) *ResourceAbstractReservation {
	this := ResourceAbstractReservation{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourceAbstractReservationWithDefaults instantiates a new ResourceAbstractReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAbstractReservationWithDefaults() *ResourceAbstractReservation {
	this := ResourceAbstractReservation{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourceAbstractReservation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourceAbstractReservation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourceAbstractReservation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourceAbstractReservation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceAbstractReservation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceAbstractReservation) SetDescription(v string) {
	o.Description = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *ResourceAbstractReservation) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *ResourceAbstractReservation) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetReservationSelector returns the ReservationSelector field value if set, zero value otherwise.
func (o *ResourceAbstractReservation) GetReservationSelector() string {
	if o == nil || IsNil(o.ReservationSelector) {
		var ret string
		return ret
	}
	return *o.ReservationSelector
}

// GetReservationSelectorOk returns a tuple with the ReservationSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetReservationSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.ReservationSelector) {
		return nil, false
	}
	return o.ReservationSelector, true
}

// HasReservationSelector returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasReservationSelector() bool {
	if o != nil && !IsNil(o.ReservationSelector) {
		return true
	}

	return false
}

// SetReservationSelector gets a reference to the given string and assigns it to the ReservationSelector field.
func (o *ResourceAbstractReservation) SetReservationSelector(v string) {
	o.ReservationSelector = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceAbstractReservation) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourceAbstractReservation) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResourceAbstractReservation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAbstractReservation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResourceAbstractReservation) SetStatus(v string) {
	o.Status = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceAbstractReservation) GetIdentity() MoBaseMoRelationship {
	if o == nil || IsNil(o.Identity.Get()) {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Identity.Get()
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceAbstractReservation) GetIdentityOk() (*MoBaseMoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identity.Get(), o.Identity.IsSet()
}

// HasIdentity returns a boolean if a field has been set.
func (o *ResourceAbstractReservation) HasIdentity() bool {
	if o != nil && o.Identity.IsSet() {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given NullableMoBaseMoRelationship and assigns it to the Identity field.
func (o *ResourceAbstractReservation) SetIdentity(v MoBaseMoRelationship) {
	o.Identity.Set(&v)
}

// SetIdentityNil sets the value for Identity to be an explicit nil
func (o *ResourceAbstractReservation) SetIdentityNil() {
	o.Identity.Set(nil)
}

// UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
func (o *ResourceAbstractReservation) UnsetIdentity() {
	o.Identity.Unset()
}

func (o ResourceAbstractReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceAbstractReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Expiration) {
		toSerialize["Expiration"] = o.Expiration
	}
	if !IsNil(o.ReservationSelector) {
		toSerialize["ReservationSelector"] = o.ReservationSelector
	}
	if !IsNil(o.ResourceType) {
		toSerialize["ResourceType"] = o.ResourceType
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.Identity.IsSet() {
		toSerialize["Identity"] = o.Identity.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceAbstractReservation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
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
	type ResourceAbstractReservationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Details of the use case for which the reservation was created, such as decommissioning.
		Description *string `json:"Description,omitempty"`
		// The resource reservation includes an expiration date and a timestamp indicating when this management object will be cleared. The expiration date is set during the decommissioning process and is maintained for a period of 3 months.
		Expiration *time.Time `json:"Expiration,omitempty"`
		// The unique identification of the resource is based on the resource OData string, which is mentioned as part of the ReservationSelector. For example, 'Serial eq 'EM6259AE6B'.
		ReservationSelector *string `json:"ReservationSelector,omitempty"`
		// The type of resource that is placed into resource groups or pools. Resource Type can be either 'compute.Blade' or 'compute.RackUnit' for pools.
		ResourceType *string `json:"ResourceType,omitempty"`
		// The reservation status can be in the 'Created', 'Processing', 'Failed', or 'Finished' state. * `Created` - By default, a reservation is in Created status. * `Processing` - A reservation is changed to Processing status for appliance mode resource claim requests. * `Failed` - A reservation is changed to Failed status if the validations on resources, resource groups fails. * `Finished` - A reservation is changed to Finished status if the validations on resources, resource groups are successful. The resource moids in reservation will be added to resource groups using OData filters.
		Status   *string                      `json:"Status,omitempty"`
		Identity NullableMoBaseMoRelationship `json:"Identity,omitempty"`
	}

	varResourceAbstractReservationWithoutEmbeddedStruct := ResourceAbstractReservationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varResourceAbstractReservationWithoutEmbeddedStruct)
	if err == nil {
		varResourceAbstractReservation := _ResourceAbstractReservation{}
		varResourceAbstractReservation.ClassId = varResourceAbstractReservationWithoutEmbeddedStruct.ClassId
		varResourceAbstractReservation.ObjectType = varResourceAbstractReservationWithoutEmbeddedStruct.ObjectType
		varResourceAbstractReservation.Description = varResourceAbstractReservationWithoutEmbeddedStruct.Description
		varResourceAbstractReservation.Expiration = varResourceAbstractReservationWithoutEmbeddedStruct.Expiration
		varResourceAbstractReservation.ReservationSelector = varResourceAbstractReservationWithoutEmbeddedStruct.ReservationSelector
		varResourceAbstractReservation.ResourceType = varResourceAbstractReservationWithoutEmbeddedStruct.ResourceType
		varResourceAbstractReservation.Status = varResourceAbstractReservationWithoutEmbeddedStruct.Status
		varResourceAbstractReservation.Identity = varResourceAbstractReservationWithoutEmbeddedStruct.Identity
		*o = ResourceAbstractReservation(varResourceAbstractReservation)
	} else {
		return err
	}

	varResourceAbstractReservation := _ResourceAbstractReservation{}

	err = json.Unmarshal(data, &varResourceAbstractReservation)
	if err == nil {
		o.MoBaseMo = varResourceAbstractReservation.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Expiration")
		delete(additionalProperties, "ReservationSelector")
		delete(additionalProperties, "ResourceType")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Identity")

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

type NullableResourceAbstractReservation struct {
	value *ResourceAbstractReservation
	isSet bool
}

func (v NullableResourceAbstractReservation) Get() *ResourceAbstractReservation {
	return v.value
}

func (v *NullableResourceAbstractReservation) Set(val *ResourceAbstractReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAbstractReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAbstractReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAbstractReservation(val *ResourceAbstractReservation) *NullableResourceAbstractReservation {
	return &NullableResourceAbstractReservation{value: val, isSet: true}
}

func (v NullableResourceAbstractReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAbstractReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
