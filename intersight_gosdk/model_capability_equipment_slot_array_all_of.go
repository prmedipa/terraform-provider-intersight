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

// CapabilityEquipmentSlotArrayAllOf Definition of the list of properties defined in 'capability.EquipmentSlotArray', excluding properties defined in parent classes.
type CapabilityEquipmentSlotArrayAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// First Index information for a Switch/Fabric-Interconnect hardware.
	FirstIndex *float32 `json:"FirstIndex,omitempty"`
	// Height information for a Switch/Fabric-Interconnect hardware.
	Height *float32 `json:"Height,omitempty"`
	// Horizontal Start Offset information for a Switch/Fabric-Interconnect hardware.
	HorizontalStartOffset *float32 `json:"HorizontalStartOffset,omitempty"`
	// Inline Group Separation information for a Switch/Fabric-Interconnect hardware.
	InlineGroupSeparation *float32 `json:"InlineGroupSeparation,omitempty"`
	// Inline Group Size information for a Switch/Fabric-Interconnect hardware.
	InlineGroupSize *float32 `json:"InlineGroupSize,omitempty"`
	// Inline Offset information for a Switch/Fabric-Interconnect hardware.
	InlineOffset *float32 `json:"InlineOffset,omitempty"`
	// Location information for a Switch/Fabric-Interconnect hardware.
	Location *string `json:"Location,omitempty"`
	// Number of Slots information for a Switch/Fabric-Interconnect hardware.
	NumberOfSlots *int64 `json:"NumberOfSlots,omitempty"`
	// Orientation information for a Switch/Fabric-Interconnect hardware.
	Orientation *string `json:"Orientation,omitempty"`
	// Selector information for a Switch/Fabric-Interconnect hardware.
	Selector *string `json:"Selector,omitempty"`
	// Slots per line information for a Switch/Fabric-Interconnect hardware.
	SlotsPerLine *int64 `json:"SlotsPerLine,omitempty"`
	// Transverse Group Separation information for a Switch/Fabric-Interconnect hardware.
	TransverseGroupSeparation *float32 `json:"TransverseGroupSeparation,omitempty"`
	// Transverse Group Size information for a Switch/Fabric-Interconnect hardware.
	TransverseGroupSize *float32 `json:"TransverseGroupSize,omitempty"`
	// Transverse Offset information for a Switch/Fabric-Interconnect hardware.
	TransverseOffset *float32 `json:"TransverseOffset,omitempty"`
	// Vertical Start Offset information for a Switch/Fabric-Interconnect hardware.
	VerticalStartOffset *float32 `json:"VerticalStartOffset,omitempty"`
	// Width information for a Switch/Fabric-Interconnect hardware.
	Width                *float32 `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CapabilityEquipmentSlotArrayAllOf CapabilityEquipmentSlotArrayAllOf

// NewCapabilityEquipmentSlotArrayAllOf instantiates a new CapabilityEquipmentSlotArrayAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityEquipmentSlotArrayAllOf(classId string, objectType string) *CapabilityEquipmentSlotArrayAllOf {
	this := CapabilityEquipmentSlotArrayAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCapabilityEquipmentSlotArrayAllOfWithDefaults instantiates a new CapabilityEquipmentSlotArrayAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityEquipmentSlotArrayAllOfWithDefaults() *CapabilityEquipmentSlotArrayAllOf {
	this := CapabilityEquipmentSlotArrayAllOf{}
	var classId string = "capability.EquipmentSlotArray"
	this.ClassId = classId
	var objectType string = "capability.EquipmentSlotArray"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CapabilityEquipmentSlotArrayAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CapabilityEquipmentSlotArrayAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CapabilityEquipmentSlotArrayAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CapabilityEquipmentSlotArrayAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFirstIndex returns the FirstIndex field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetFirstIndex() float32 {
	if o == nil || o.FirstIndex == nil {
		var ret float32
		return ret
	}
	return *o.FirstIndex
}

// GetFirstIndexOk returns a tuple with the FirstIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetFirstIndexOk() (*float32, bool) {
	if o == nil || o.FirstIndex == nil {
		return nil, false
	}
	return o.FirstIndex, true
}

// HasFirstIndex returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasFirstIndex() bool {
	if o != nil && o.FirstIndex != nil {
		return true
	}

	return false
}

// SetFirstIndex gets a reference to the given float32 and assigns it to the FirstIndex field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetFirstIndex(v float32) {
	o.FirstIndex = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetHeight() float32 {
	if o == nil || o.Height == nil {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetHeightOk() (*float32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetHeight(v float32) {
	o.Height = &v
}

// GetHorizontalStartOffset returns the HorizontalStartOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetHorizontalStartOffset() float32 {
	if o == nil || o.HorizontalStartOffset == nil {
		var ret float32
		return ret
	}
	return *o.HorizontalStartOffset
}

// GetHorizontalStartOffsetOk returns a tuple with the HorizontalStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetHorizontalStartOffsetOk() (*float32, bool) {
	if o == nil || o.HorizontalStartOffset == nil {
		return nil, false
	}
	return o.HorizontalStartOffset, true
}

// HasHorizontalStartOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasHorizontalStartOffset() bool {
	if o != nil && o.HorizontalStartOffset != nil {
		return true
	}

	return false
}

// SetHorizontalStartOffset gets a reference to the given float32 and assigns it to the HorizontalStartOffset field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetHorizontalStartOffset(v float32) {
	o.HorizontalStartOffset = &v
}

// GetInlineGroupSeparation returns the InlineGroupSeparation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineGroupSeparation() float32 {
	if o == nil || o.InlineGroupSeparation == nil {
		var ret float32
		return ret
	}
	return *o.InlineGroupSeparation
}

// GetInlineGroupSeparationOk returns a tuple with the InlineGroupSeparation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineGroupSeparationOk() (*float32, bool) {
	if o == nil || o.InlineGroupSeparation == nil {
		return nil, false
	}
	return o.InlineGroupSeparation, true
}

// HasInlineGroupSeparation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasInlineGroupSeparation() bool {
	if o != nil && o.InlineGroupSeparation != nil {
		return true
	}

	return false
}

// SetInlineGroupSeparation gets a reference to the given float32 and assigns it to the InlineGroupSeparation field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetInlineGroupSeparation(v float32) {
	o.InlineGroupSeparation = &v
}

// GetInlineGroupSize returns the InlineGroupSize field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineGroupSize() float32 {
	if o == nil || o.InlineGroupSize == nil {
		var ret float32
		return ret
	}
	return *o.InlineGroupSize
}

// GetInlineGroupSizeOk returns a tuple with the InlineGroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineGroupSizeOk() (*float32, bool) {
	if o == nil || o.InlineGroupSize == nil {
		return nil, false
	}
	return o.InlineGroupSize, true
}

// HasInlineGroupSize returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasInlineGroupSize() bool {
	if o != nil && o.InlineGroupSize != nil {
		return true
	}

	return false
}

// SetInlineGroupSize gets a reference to the given float32 and assigns it to the InlineGroupSize field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetInlineGroupSize(v float32) {
	o.InlineGroupSize = &v
}

// GetInlineOffset returns the InlineOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineOffset() float32 {
	if o == nil || o.InlineOffset == nil {
		var ret float32
		return ret
	}
	return *o.InlineOffset
}

// GetInlineOffsetOk returns a tuple with the InlineOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetInlineOffsetOk() (*float32, bool) {
	if o == nil || o.InlineOffset == nil {
		return nil, false
	}
	return o.InlineOffset, true
}

// HasInlineOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasInlineOffset() bool {
	if o != nil && o.InlineOffset != nil {
		return true
	}

	return false
}

// SetInlineOffset gets a reference to the given float32 and assigns it to the InlineOffset field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetInlineOffset(v float32) {
	o.InlineOffset = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetLocation(v string) {
	o.Location = &v
}

// GetNumberOfSlots returns the NumberOfSlots field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetNumberOfSlots() int64 {
	if o == nil || o.NumberOfSlots == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSlots
}

// GetNumberOfSlotsOk returns a tuple with the NumberOfSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetNumberOfSlotsOk() (*int64, bool) {
	if o == nil || o.NumberOfSlots == nil {
		return nil, false
	}
	return o.NumberOfSlots, true
}

// HasNumberOfSlots returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasNumberOfSlots() bool {
	if o != nil && o.NumberOfSlots != nil {
		return true
	}

	return false
}

// SetNumberOfSlots gets a reference to the given int64 and assigns it to the NumberOfSlots field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetNumberOfSlots(v int64) {
	o.NumberOfSlots = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetOrientation() string {
	if o == nil || o.Orientation == nil {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetOrientationOk() (*string, bool) {
	if o == nil || o.Orientation == nil {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasOrientation() bool {
	if o != nil && o.Orientation != nil {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetOrientation(v string) {
	o.Orientation = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetSelector(v string) {
	o.Selector = &v
}

// GetSlotsPerLine returns the SlotsPerLine field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetSlotsPerLine() int64 {
	if o == nil || o.SlotsPerLine == nil {
		var ret int64
		return ret
	}
	return *o.SlotsPerLine
}

// GetSlotsPerLineOk returns a tuple with the SlotsPerLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetSlotsPerLineOk() (*int64, bool) {
	if o == nil || o.SlotsPerLine == nil {
		return nil, false
	}
	return o.SlotsPerLine, true
}

// HasSlotsPerLine returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasSlotsPerLine() bool {
	if o != nil && o.SlotsPerLine != nil {
		return true
	}

	return false
}

// SetSlotsPerLine gets a reference to the given int64 and assigns it to the SlotsPerLine field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetSlotsPerLine(v int64) {
	o.SlotsPerLine = &v
}

// GetTransverseGroupSeparation returns the TransverseGroupSeparation field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseGroupSeparation() float32 {
	if o == nil || o.TransverseGroupSeparation == nil {
		var ret float32
		return ret
	}
	return *o.TransverseGroupSeparation
}

// GetTransverseGroupSeparationOk returns a tuple with the TransverseGroupSeparation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseGroupSeparationOk() (*float32, bool) {
	if o == nil || o.TransverseGroupSeparation == nil {
		return nil, false
	}
	return o.TransverseGroupSeparation, true
}

// HasTransverseGroupSeparation returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasTransverseGroupSeparation() bool {
	if o != nil && o.TransverseGroupSeparation != nil {
		return true
	}

	return false
}

// SetTransverseGroupSeparation gets a reference to the given float32 and assigns it to the TransverseGroupSeparation field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetTransverseGroupSeparation(v float32) {
	o.TransverseGroupSeparation = &v
}

// GetTransverseGroupSize returns the TransverseGroupSize field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseGroupSize() float32 {
	if o == nil || o.TransverseGroupSize == nil {
		var ret float32
		return ret
	}
	return *o.TransverseGroupSize
}

// GetTransverseGroupSizeOk returns a tuple with the TransverseGroupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseGroupSizeOk() (*float32, bool) {
	if o == nil || o.TransverseGroupSize == nil {
		return nil, false
	}
	return o.TransverseGroupSize, true
}

// HasTransverseGroupSize returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasTransverseGroupSize() bool {
	if o != nil && o.TransverseGroupSize != nil {
		return true
	}

	return false
}

// SetTransverseGroupSize gets a reference to the given float32 and assigns it to the TransverseGroupSize field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetTransverseGroupSize(v float32) {
	o.TransverseGroupSize = &v
}

// GetTransverseOffset returns the TransverseOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseOffset() float32 {
	if o == nil || o.TransverseOffset == nil {
		var ret float32
		return ret
	}
	return *o.TransverseOffset
}

// GetTransverseOffsetOk returns a tuple with the TransverseOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetTransverseOffsetOk() (*float32, bool) {
	if o == nil || o.TransverseOffset == nil {
		return nil, false
	}
	return o.TransverseOffset, true
}

// HasTransverseOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasTransverseOffset() bool {
	if o != nil && o.TransverseOffset != nil {
		return true
	}

	return false
}

// SetTransverseOffset gets a reference to the given float32 and assigns it to the TransverseOffset field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetTransverseOffset(v float32) {
	o.TransverseOffset = &v
}

// GetVerticalStartOffset returns the VerticalStartOffset field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetVerticalStartOffset() float32 {
	if o == nil || o.VerticalStartOffset == nil {
		var ret float32
		return ret
	}
	return *o.VerticalStartOffset
}

// GetVerticalStartOffsetOk returns a tuple with the VerticalStartOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetVerticalStartOffsetOk() (*float32, bool) {
	if o == nil || o.VerticalStartOffset == nil {
		return nil, false
	}
	return o.VerticalStartOffset, true
}

// HasVerticalStartOffset returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasVerticalStartOffset() bool {
	if o != nil && o.VerticalStartOffset != nil {
		return true
	}

	return false
}

// SetVerticalStartOffset gets a reference to the given float32 and assigns it to the VerticalStartOffset field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetVerticalStartOffset(v float32) {
	o.VerticalStartOffset = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CapabilityEquipmentSlotArrayAllOf) GetWidth() float32 {
	if o == nil || o.Width == nil {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) GetWidthOk() (*float32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CapabilityEquipmentSlotArrayAllOf) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *CapabilityEquipmentSlotArrayAllOf) SetWidth(v float32) {
	o.Width = &v
}

func (o CapabilityEquipmentSlotArrayAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FirstIndex != nil {
		toSerialize["FirstIndex"] = o.FirstIndex
	}
	if o.Height != nil {
		toSerialize["Height"] = o.Height
	}
	if o.HorizontalStartOffset != nil {
		toSerialize["HorizontalStartOffset"] = o.HorizontalStartOffset
	}
	if o.InlineGroupSeparation != nil {
		toSerialize["InlineGroupSeparation"] = o.InlineGroupSeparation
	}
	if o.InlineGroupSize != nil {
		toSerialize["InlineGroupSize"] = o.InlineGroupSize
	}
	if o.InlineOffset != nil {
		toSerialize["InlineOffset"] = o.InlineOffset
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.NumberOfSlots != nil {
		toSerialize["NumberOfSlots"] = o.NumberOfSlots
	}
	if o.Orientation != nil {
		toSerialize["Orientation"] = o.Orientation
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}
	if o.SlotsPerLine != nil {
		toSerialize["SlotsPerLine"] = o.SlotsPerLine
	}
	if o.TransverseGroupSeparation != nil {
		toSerialize["TransverseGroupSeparation"] = o.TransverseGroupSeparation
	}
	if o.TransverseGroupSize != nil {
		toSerialize["TransverseGroupSize"] = o.TransverseGroupSize
	}
	if o.TransverseOffset != nil {
		toSerialize["TransverseOffset"] = o.TransverseOffset
	}
	if o.VerticalStartOffset != nil {
		toSerialize["VerticalStartOffset"] = o.VerticalStartOffset
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CapabilityEquipmentSlotArrayAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varCapabilityEquipmentSlotArrayAllOf := _CapabilityEquipmentSlotArrayAllOf{}

	if err = json.Unmarshal(bytes, &varCapabilityEquipmentSlotArrayAllOf); err == nil {
		*o = CapabilityEquipmentSlotArrayAllOf(varCapabilityEquipmentSlotArrayAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FirstIndex")
		delete(additionalProperties, "Height")
		delete(additionalProperties, "HorizontalStartOffset")
		delete(additionalProperties, "InlineGroupSeparation")
		delete(additionalProperties, "InlineGroupSize")
		delete(additionalProperties, "InlineOffset")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "NumberOfSlots")
		delete(additionalProperties, "Orientation")
		delete(additionalProperties, "Selector")
		delete(additionalProperties, "SlotsPerLine")
		delete(additionalProperties, "TransverseGroupSeparation")
		delete(additionalProperties, "TransverseGroupSize")
		delete(additionalProperties, "TransverseOffset")
		delete(additionalProperties, "VerticalStartOffset")
		delete(additionalProperties, "Width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapabilityEquipmentSlotArrayAllOf struct {
	value *CapabilityEquipmentSlotArrayAllOf
	isSet bool
}

func (v NullableCapabilityEquipmentSlotArrayAllOf) Get() *CapabilityEquipmentSlotArrayAllOf {
	return v.value
}

func (v *NullableCapabilityEquipmentSlotArrayAllOf) Set(val *CapabilityEquipmentSlotArrayAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityEquipmentSlotArrayAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityEquipmentSlotArrayAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityEquipmentSlotArrayAllOf(val *CapabilityEquipmentSlotArrayAllOf) *NullableCapabilityEquipmentSlotArrayAllOf {
	return &NullableCapabilityEquipmentSlotArrayAllOf{value: val, isSet: true}
}

func (v NullableCapabilityEquipmentSlotArrayAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityEquipmentSlotArrayAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
