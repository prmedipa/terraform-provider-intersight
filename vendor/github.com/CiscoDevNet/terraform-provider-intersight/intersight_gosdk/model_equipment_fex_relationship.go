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
)

// EquipmentFexRelationship - A relationship to the 'equipment.Fex' resource, or the expanded 'equipment.Fex' resource, or the 'null' value.
type EquipmentFexRelationship struct {
	EquipmentFex *EquipmentFex
	MoMoRef      *MoMoRef
}

// EquipmentFexAsEquipmentFexRelationship is a convenience function that returns EquipmentFex wrapped in EquipmentFexRelationship
func EquipmentFexAsEquipmentFexRelationship(v *EquipmentFex) EquipmentFexRelationship {
	return EquipmentFexRelationship{
		EquipmentFex: v,
	}
}

// MoMoRefAsEquipmentFexRelationship is a convenience function that returns MoMoRef wrapped in EquipmentFexRelationship
func MoMoRefAsEquipmentFexRelationship(v *MoMoRef) EquipmentFexRelationship {
	return EquipmentFexRelationship{
		MoMoRef: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EquipmentFexRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'equipment.Fex'
	if jsonDict["ClassId"] == "equipment.Fex" {
		// try to unmarshal JSON data into EquipmentFex
		err = json.Unmarshal(data, &dst.EquipmentFex)
		if err == nil {
			return nil // data stored in dst.EquipmentFex, return on the first match
		} else {
			dst.EquipmentFex = nil
			return fmt.Errorf("failed to unmarshal EquipmentFexRelationship as EquipmentFex: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("failed to unmarshal EquipmentFexRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EquipmentFexRelationship) MarshalJSON() ([]byte, error) {
	if src.EquipmentFex != nil {
		return json.Marshal(&src.EquipmentFex)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EquipmentFexRelationship) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.EquipmentFex != nil {
		return obj.EquipmentFex
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableEquipmentFexRelationship struct {
	value *EquipmentFexRelationship
	isSet bool
}

func (v NullableEquipmentFexRelationship) Get() *EquipmentFexRelationship {
	return v.value
}

func (v *NullableEquipmentFexRelationship) Set(val *EquipmentFexRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFexRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFexRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFexRelationship(val *EquipmentFexRelationship) *NullableEquipmentFexRelationship {
	return &NullableEquipmentFexRelationship{value: val, isSet: true}
}

func (v NullableEquipmentFexRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFexRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
