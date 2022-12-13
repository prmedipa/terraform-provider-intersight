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

// ResourcepoolPoolMemberAllOf Definition of the list of properties defined in 'resourcepool.PoolMember', excluding properties defined in parent classes.
type ResourcepoolPoolMemberAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string   `json:"ObjectType"`
	Features   []string `json:"Features,omitempty"`
	// An array of relationships to moBaseMo resources.
	AssignedToEntity     []MoBaseMoRelationship         `json:"AssignedToEntity,omitempty"`
	Peer                 *ResourcepoolLeaseRelationship `json:"Peer,omitempty"`
	Pool                 *ResourcepoolPoolRelationship  `json:"Pool,omitempty"`
	Resource             *MoBaseMoRelationship          `json:"Resource,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourcepoolPoolMemberAllOf ResourcepoolPoolMemberAllOf

// NewResourcepoolPoolMemberAllOf instantiates a new ResourcepoolPoolMemberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcepoolPoolMemberAllOf(classId string, objectType string) *ResourcepoolPoolMemberAllOf {
	this := ResourcepoolPoolMemberAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewResourcepoolPoolMemberAllOfWithDefaults instantiates a new ResourcepoolPoolMemberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcepoolPoolMemberAllOfWithDefaults() *ResourcepoolPoolMemberAllOf {
	this := ResourcepoolPoolMemberAllOf{}
	var classId string = "resourcepool.PoolMember"
	this.ClassId = classId
	var objectType string = "resourcepool.PoolMember"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ResourcepoolPoolMemberAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolPoolMemberAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ResourcepoolPoolMemberAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ResourcepoolPoolMemberAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ResourcepoolPoolMemberAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ResourcepoolPoolMemberAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolPoolMemberAllOf) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolPoolMemberAllOf) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ResourcepoolPoolMemberAllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *ResourcepoolPoolMemberAllOf) SetFeatures(v []string) {
	o.Features = v
}

// GetAssignedToEntity returns the AssignedToEntity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourcepoolPoolMemberAllOf) GetAssignedToEntity() []MoBaseMoRelationship {
	if o == nil {
		var ret []MoBaseMoRelationship
		return ret
	}
	return o.AssignedToEntity
}

// GetAssignedToEntityOk returns a tuple with the AssignedToEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourcepoolPoolMemberAllOf) GetAssignedToEntityOk() ([]MoBaseMoRelationship, bool) {
	if o == nil || o.AssignedToEntity == nil {
		return nil, false
	}
	return o.AssignedToEntity, true
}

// HasAssignedToEntity returns a boolean if a field has been set.
func (o *ResourcepoolPoolMemberAllOf) HasAssignedToEntity() bool {
	if o != nil && o.AssignedToEntity != nil {
		return true
	}

	return false
}

// SetAssignedToEntity gets a reference to the given []MoBaseMoRelationship and assigns it to the AssignedToEntity field.
func (o *ResourcepoolPoolMemberAllOf) SetAssignedToEntity(v []MoBaseMoRelationship) {
	o.AssignedToEntity = v
}

// GetPeer returns the Peer field value if set, zero value otherwise.
func (o *ResourcepoolPoolMemberAllOf) GetPeer() ResourcepoolLeaseRelationship {
	if o == nil || o.Peer == nil {
		var ret ResourcepoolLeaseRelationship
		return ret
	}
	return *o.Peer
}

// GetPeerOk returns a tuple with the Peer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPoolMemberAllOf) GetPeerOk() (*ResourcepoolLeaseRelationship, bool) {
	if o == nil || o.Peer == nil {
		return nil, false
	}
	return o.Peer, true
}

// HasPeer returns a boolean if a field has been set.
func (o *ResourcepoolPoolMemberAllOf) HasPeer() bool {
	if o != nil && o.Peer != nil {
		return true
	}

	return false
}

// SetPeer gets a reference to the given ResourcepoolLeaseRelationship and assigns it to the Peer field.
func (o *ResourcepoolPoolMemberAllOf) SetPeer(v ResourcepoolLeaseRelationship) {
	o.Peer = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *ResourcepoolPoolMemberAllOf) GetPool() ResourcepoolPoolRelationship {
	if o == nil || o.Pool == nil {
		var ret ResourcepoolPoolRelationship
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPoolMemberAllOf) GetPoolOk() (*ResourcepoolPoolRelationship, bool) {
	if o == nil || o.Pool == nil {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *ResourcepoolPoolMemberAllOf) HasPool() bool {
	if o != nil && o.Pool != nil {
		return true
	}

	return false
}

// SetPool gets a reference to the given ResourcepoolPoolRelationship and assigns it to the Pool field.
func (o *ResourcepoolPoolMemberAllOf) SetPool(v ResourcepoolPoolRelationship) {
	o.Pool = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourcepoolPoolMemberAllOf) GetResource() MoBaseMoRelationship {
	if o == nil || o.Resource == nil {
		var ret MoBaseMoRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcepoolPoolMemberAllOf) GetResourceOk() (*MoBaseMoRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourcepoolPoolMemberAllOf) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given MoBaseMoRelationship and assigns it to the Resource field.
func (o *ResourcepoolPoolMemberAllOf) SetResource(v MoBaseMoRelationship) {
	o.Resource = &v
}

func (o ResourcepoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Features != nil {
		toSerialize["Features"] = o.Features
	}
	if o.AssignedToEntity != nil {
		toSerialize["AssignedToEntity"] = o.AssignedToEntity
	}
	if o.Peer != nil {
		toSerialize["Peer"] = o.Peer
	}
	if o.Pool != nil {
		toSerialize["Pool"] = o.Pool
	}
	if o.Resource != nil {
		toSerialize["Resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourcepoolPoolMemberAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourcepoolPoolMemberAllOf := _ResourcepoolPoolMemberAllOf{}

	if err = json.Unmarshal(bytes, &varResourcepoolPoolMemberAllOf); err == nil {
		*o = ResourcepoolPoolMemberAllOf(varResourcepoolPoolMemberAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Features")
		delete(additionalProperties, "AssignedToEntity")
		delete(additionalProperties, "Peer")
		delete(additionalProperties, "Pool")
		delete(additionalProperties, "Resource")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourcepoolPoolMemberAllOf struct {
	value *ResourcepoolPoolMemberAllOf
	isSet bool
}

func (v NullableResourcepoolPoolMemberAllOf) Get() *ResourcepoolPoolMemberAllOf {
	return v.value
}

func (v *NullableResourcepoolPoolMemberAllOf) Set(val *ResourcepoolPoolMemberAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcepoolPoolMemberAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcepoolPoolMemberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcepoolPoolMemberAllOf(val *ResourcepoolPoolMemberAllOf) *NullableResourcepoolPoolMemberAllOf {
	return &NullableResourcepoolPoolMemberAllOf{value: val, isSet: true}
}

func (v NullableResourcepoolPoolMemberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcepoolPoolMemberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
