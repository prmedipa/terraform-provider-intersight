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
	"reflect"
	"strings"
)

// RecommendationPhysicalItem Entity representing the recommended physical device.
type RecommendationPhysicalItem struct {
	RecommendationAbstractItem
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Capacity of the physical entity added.
	Capacity *int64 `json:"Capacity,omitempty"`
	// Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation.
	Count *int64 `json:"Count,omitempty"`
	// If the PhysicalItem is new, this is set to true, else false.
	IsNew *bool `json:"IsNew,omitempty"`
	// Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation.
	MaxCount *int64 `json:"MaxCount,omitempty"`
	// Model of the recommended physical device which is externally identifiable.
	Model *string `json:"Model,omitempty"`
	// Moid of the managed object which represents the parent physical entity.
	ParentMoid *string `json:"ParentMoid,omitempty"`
	// Moid of the managed object which represents the existing physical entity.
	SourceMoid *string `json:"SourceMoid,omitempty"`
	// Unit of the new capacity. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
	Unit *string `json:"Unit,omitempty"`
	// Uuid of the recommended physical device.
	Uuid             *string                                     `json:"Uuid,omitempty"`
	CapacityRunway   *RecommendationCapacityRunwayRelationship   `json:"CapacityRunway,omitempty"`
	ClusterExpansion *RecommendationClusterExpansionRelationship `json:"ClusterExpansion,omitempty"`
	// An array of relationships to recommendationPhysicalItem resources.
	PhysicalItem         []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecommendationPhysicalItem RecommendationPhysicalItem

// NewRecommendationPhysicalItem instantiates a new RecommendationPhysicalItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationPhysicalItem(classId string, objectType string) *RecommendationPhysicalItem {
	this := RecommendationPhysicalItem{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewRecommendationPhysicalItemWithDefaults instantiates a new RecommendationPhysicalItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationPhysicalItemWithDefaults() *RecommendationPhysicalItem {
	this := RecommendationPhysicalItem{}
	var classId string = "recommendation.PhysicalItem"
	this.ClassId = classId
	var objectType string = "recommendation.PhysicalItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecommendationPhysicalItem) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecommendationPhysicalItem) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecommendationPhysicalItem) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecommendationPhysicalItem) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *RecommendationPhysicalItem) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RecommendationPhysicalItem) SetCount(v int64) {
	o.Count = &v
}

// GetIsNew returns the IsNew field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetIsNew() bool {
	if o == nil || o.IsNew == nil {
		var ret bool
		return ret
	}
	return *o.IsNew
}

// GetIsNewOk returns a tuple with the IsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetIsNewOk() (*bool, bool) {
	if o == nil || o.IsNew == nil {
		return nil, false
	}
	return o.IsNew, true
}

// HasIsNew returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasIsNew() bool {
	if o != nil && o.IsNew != nil {
		return true
	}

	return false
}

// SetIsNew gets a reference to the given bool and assigns it to the IsNew field.
func (o *RecommendationPhysicalItem) SetIsNew(v bool) {
	o.IsNew = &v
}

// GetMaxCount returns the MaxCount field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetMaxCount() int64 {
	if o == nil || o.MaxCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxCount
}

// GetMaxCountOk returns a tuple with the MaxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetMaxCountOk() (*int64, bool) {
	if o == nil || o.MaxCount == nil {
		return nil, false
	}
	return o.MaxCount, true
}

// HasMaxCount returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasMaxCount() bool {
	if o != nil && o.MaxCount != nil {
		return true
	}

	return false
}

// SetMaxCount gets a reference to the given int64 and assigns it to the MaxCount field.
func (o *RecommendationPhysicalItem) SetMaxCount(v int64) {
	o.MaxCount = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *RecommendationPhysicalItem) SetModel(v string) {
	o.Model = &v
}

// GetParentMoid returns the ParentMoid field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetParentMoid() string {
	if o == nil || o.ParentMoid == nil {
		var ret string
		return ret
	}
	return *o.ParentMoid
}

// GetParentMoidOk returns a tuple with the ParentMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetParentMoidOk() (*string, bool) {
	if o == nil || o.ParentMoid == nil {
		return nil, false
	}
	return o.ParentMoid, true
}

// HasParentMoid returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasParentMoid() bool {
	if o != nil && o.ParentMoid != nil {
		return true
	}

	return false
}

// SetParentMoid gets a reference to the given string and assigns it to the ParentMoid field.
func (o *RecommendationPhysicalItem) SetParentMoid(v string) {
	o.ParentMoid = &v
}

// GetSourceMoid returns the SourceMoid field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetSourceMoid() string {
	if o == nil || o.SourceMoid == nil {
		var ret string
		return ret
	}
	return *o.SourceMoid
}

// GetSourceMoidOk returns a tuple with the SourceMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetSourceMoidOk() (*string, bool) {
	if o == nil || o.SourceMoid == nil {
		return nil, false
	}
	return o.SourceMoid, true
}

// HasSourceMoid returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasSourceMoid() bool {
	if o != nil && o.SourceMoid != nil {
		return true
	}

	return false
}

// SetSourceMoid gets a reference to the given string and assigns it to the SourceMoid field.
func (o *RecommendationPhysicalItem) SetSourceMoid(v string) {
	o.SourceMoid = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *RecommendationPhysicalItem) SetUnit(v string) {
	o.Unit = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RecommendationPhysicalItem) SetUuid(v string) {
	o.Uuid = &v
}

// GetCapacityRunway returns the CapacityRunway field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetCapacityRunway() RecommendationCapacityRunwayRelationship {
	if o == nil || o.CapacityRunway == nil {
		var ret RecommendationCapacityRunwayRelationship
		return ret
	}
	return *o.CapacityRunway
}

// GetCapacityRunwayOk returns a tuple with the CapacityRunway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetCapacityRunwayOk() (*RecommendationCapacityRunwayRelationship, bool) {
	if o == nil || o.CapacityRunway == nil {
		return nil, false
	}
	return o.CapacityRunway, true
}

// HasCapacityRunway returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasCapacityRunway() bool {
	if o != nil && o.CapacityRunway != nil {
		return true
	}

	return false
}

// SetCapacityRunway gets a reference to the given RecommendationCapacityRunwayRelationship and assigns it to the CapacityRunway field.
func (o *RecommendationPhysicalItem) SetCapacityRunway(v RecommendationCapacityRunwayRelationship) {
	o.CapacityRunway = &v
}

// GetClusterExpansion returns the ClusterExpansion field value if set, zero value otherwise.
func (o *RecommendationPhysicalItem) GetClusterExpansion() RecommendationClusterExpansionRelationship {
	if o == nil || o.ClusterExpansion == nil {
		var ret RecommendationClusterExpansionRelationship
		return ret
	}
	return *o.ClusterExpansion
}

// GetClusterExpansionOk returns a tuple with the ClusterExpansion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationPhysicalItem) GetClusterExpansionOk() (*RecommendationClusterExpansionRelationship, bool) {
	if o == nil || o.ClusterExpansion == nil {
		return nil, false
	}
	return o.ClusterExpansion, true
}

// HasClusterExpansion returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasClusterExpansion() bool {
	if o != nil && o.ClusterExpansion != nil {
		return true
	}

	return false
}

// SetClusterExpansion gets a reference to the given RecommendationClusterExpansionRelationship and assigns it to the ClusterExpansion field.
func (o *RecommendationPhysicalItem) SetClusterExpansion(v RecommendationClusterExpansionRelationship) {
	o.ClusterExpansion = &v
}

// GetPhysicalItem returns the PhysicalItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecommendationPhysicalItem) GetPhysicalItem() []RecommendationPhysicalItemRelationship {
	if o == nil {
		var ret []RecommendationPhysicalItemRelationship
		return ret
	}
	return o.PhysicalItem
}

// GetPhysicalItemOk returns a tuple with the PhysicalItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecommendationPhysicalItem) GetPhysicalItemOk() ([]RecommendationPhysicalItemRelationship, bool) {
	if o == nil || o.PhysicalItem == nil {
		return nil, false
	}
	return o.PhysicalItem, true
}

// HasPhysicalItem returns a boolean if a field has been set.
func (o *RecommendationPhysicalItem) HasPhysicalItem() bool {
	if o != nil && o.PhysicalItem != nil {
		return true
	}

	return false
}

// SetPhysicalItem gets a reference to the given []RecommendationPhysicalItemRelationship and assigns it to the PhysicalItem field.
func (o *RecommendationPhysicalItem) SetPhysicalItem(v []RecommendationPhysicalItemRelationship) {
	o.PhysicalItem = v
}

func (o RecommendationPhysicalItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecommendationAbstractItem, errRecommendationAbstractItem := json.Marshal(o.RecommendationAbstractItem)
	if errRecommendationAbstractItem != nil {
		return []byte{}, errRecommendationAbstractItem
	}
	errRecommendationAbstractItem = json.Unmarshal([]byte(serializedRecommendationAbstractItem), &toSerialize)
	if errRecommendationAbstractItem != nil {
		return []byte{}, errRecommendationAbstractItem
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.IsNew != nil {
		toSerialize["IsNew"] = o.IsNew
	}
	if o.MaxCount != nil {
		toSerialize["MaxCount"] = o.MaxCount
	}
	if o.Model != nil {
		toSerialize["Model"] = o.Model
	}
	if o.ParentMoid != nil {
		toSerialize["ParentMoid"] = o.ParentMoid
	}
	if o.SourceMoid != nil {
		toSerialize["SourceMoid"] = o.SourceMoid
	}
	if o.Unit != nil {
		toSerialize["Unit"] = o.Unit
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.CapacityRunway != nil {
		toSerialize["CapacityRunway"] = o.CapacityRunway
	}
	if o.ClusterExpansion != nil {
		toSerialize["ClusterExpansion"] = o.ClusterExpansion
	}
	if o.PhysicalItem != nil {
		toSerialize["PhysicalItem"] = o.PhysicalItem
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RecommendationPhysicalItem) UnmarshalJSON(bytes []byte) (err error) {
	type RecommendationPhysicalItemWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Capacity of the physical entity added.
		Capacity *int64 `json:"Capacity,omitempty"`
		// Count of number of items/devices to be added.For example, number of disks to add on a node PhysicalItem in case of HyperFlex Cluster recommendation.
		Count *int64 `json:"Count,omitempty"`
		// If the PhysicalItem is new, this is set to true, else false.
		IsNew *bool `json:"IsNew,omitempty"`
		// Maximum number of items/devices which can be added on this PhysicalItem.For example, maximum number of disks allowed on a node PhysicalItem in case of HyperFlex Cluster recommendation.
		MaxCount *int64 `json:"MaxCount,omitempty"`
		// Model of the recommended physical device which is externally identifiable.
		Model *string `json:"Model,omitempty"`
		// Moid of the managed object which represents the parent physical entity.
		ParentMoid *string `json:"ParentMoid,omitempty"`
		// Moid of the managed object which represents the existing physical entity.
		SourceMoid *string `json:"SourceMoid,omitempty"`
		// Unit of the new capacity. * `TB` - The Enum value TB represents that the measurement unit is in terabytes. * `MB` - The Enum value MB represents that the measurement unit is in megabytes. * `GB` - The Enum value GB represents that the measurement unit is in gigabytes. * `MHz` - The Enum value MHz represents that the measurement unit is in megahertz. * `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz. * `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity.
		Unit *string `json:"Unit,omitempty"`
		// Uuid of the recommended physical device.
		Uuid             *string                                     `json:"Uuid,omitempty"`
		CapacityRunway   *RecommendationCapacityRunwayRelationship   `json:"CapacityRunway,omitempty"`
		ClusterExpansion *RecommendationClusterExpansionRelationship `json:"ClusterExpansion,omitempty"`
		// An array of relationships to recommendationPhysicalItem resources.
		PhysicalItem []RecommendationPhysicalItemRelationship `json:"PhysicalItem,omitempty"`
	}

	varRecommendationPhysicalItemWithoutEmbeddedStruct := RecommendationPhysicalItemWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varRecommendationPhysicalItemWithoutEmbeddedStruct)
	if err == nil {
		varRecommendationPhysicalItem := _RecommendationPhysicalItem{}
		varRecommendationPhysicalItem.ClassId = varRecommendationPhysicalItemWithoutEmbeddedStruct.ClassId
		varRecommendationPhysicalItem.ObjectType = varRecommendationPhysicalItemWithoutEmbeddedStruct.ObjectType
		varRecommendationPhysicalItem.Capacity = varRecommendationPhysicalItemWithoutEmbeddedStruct.Capacity
		varRecommendationPhysicalItem.Count = varRecommendationPhysicalItemWithoutEmbeddedStruct.Count
		varRecommendationPhysicalItem.IsNew = varRecommendationPhysicalItemWithoutEmbeddedStruct.IsNew
		varRecommendationPhysicalItem.MaxCount = varRecommendationPhysicalItemWithoutEmbeddedStruct.MaxCount
		varRecommendationPhysicalItem.Model = varRecommendationPhysicalItemWithoutEmbeddedStruct.Model
		varRecommendationPhysicalItem.ParentMoid = varRecommendationPhysicalItemWithoutEmbeddedStruct.ParentMoid
		varRecommendationPhysicalItem.SourceMoid = varRecommendationPhysicalItemWithoutEmbeddedStruct.SourceMoid
		varRecommendationPhysicalItem.Unit = varRecommendationPhysicalItemWithoutEmbeddedStruct.Unit
		varRecommendationPhysicalItem.Uuid = varRecommendationPhysicalItemWithoutEmbeddedStruct.Uuid
		varRecommendationPhysicalItem.CapacityRunway = varRecommendationPhysicalItemWithoutEmbeddedStruct.CapacityRunway
		varRecommendationPhysicalItem.ClusterExpansion = varRecommendationPhysicalItemWithoutEmbeddedStruct.ClusterExpansion
		varRecommendationPhysicalItem.PhysicalItem = varRecommendationPhysicalItemWithoutEmbeddedStruct.PhysicalItem
		*o = RecommendationPhysicalItem(varRecommendationPhysicalItem)
	} else {
		return err
	}

	varRecommendationPhysicalItem := _RecommendationPhysicalItem{}

	err = json.Unmarshal(bytes, &varRecommendationPhysicalItem)
	if err == nil {
		o.RecommendationAbstractItem = varRecommendationPhysicalItem.RecommendationAbstractItem
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Count")
		delete(additionalProperties, "IsNew")
		delete(additionalProperties, "MaxCount")
		delete(additionalProperties, "Model")
		delete(additionalProperties, "ParentMoid")
		delete(additionalProperties, "SourceMoid")
		delete(additionalProperties, "Unit")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "CapacityRunway")
		delete(additionalProperties, "ClusterExpansion")
		delete(additionalProperties, "PhysicalItem")

		// remove fields from embedded structs
		reflectRecommendationAbstractItem := reflect.ValueOf(o.RecommendationAbstractItem)
		for i := 0; i < reflectRecommendationAbstractItem.Type().NumField(); i++ {
			t := reflectRecommendationAbstractItem.Type().Field(i)

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

type NullableRecommendationPhysicalItem struct {
	value *RecommendationPhysicalItem
	isSet bool
}

func (v NullableRecommendationPhysicalItem) Get() *RecommendationPhysicalItem {
	return v.value
}

func (v *NullableRecommendationPhysicalItem) Set(val *RecommendationPhysicalItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationPhysicalItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationPhysicalItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationPhysicalItem(val *RecommendationPhysicalItem) *NullableRecommendationPhysicalItem {
	return &NullableRecommendationPhysicalItem{value: val, isSet: true}
}

func (v NullableRecommendationPhysicalItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationPhysicalItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
