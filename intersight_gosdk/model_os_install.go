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

// checks if the OsInstall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OsInstall{}

// OsInstall This MO models the target server, answers and other properties needed for OS installation. The OS installation can be started in the target server by doing a POST on this MO. The requests to this MO starts a OS installation workflow that can be tracked using workflow engine MO workflow.WorkflowInfo.
type OsInstall struct {
	OsBaseInstallConfig
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the OS install configuration.
	Name                 *string                                                             `json:"Name,omitempty"`
	ConfigurationFile    NullableOsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
	Image                NullableSoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
	Organization         NullableOrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
	OsduImage            NullableFirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
	Server               NullableComputePhysicalRelationship                                 `json:"Server,omitempty"`
	WorkflowInfo         NullableWorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OsInstall OsInstall

// NewOsInstall instantiates a new OsInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOsInstall(classId string, objectType string) *OsInstall {
	this := OsInstall{}
	this.ClassId = classId
	this.ObjectType = objectType
	var installMethod string = "vMedia"
	this.InstallMethod = &installMethod
	return &this
}

// NewOsInstallWithDefaults instantiates a new OsInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOsInstallWithDefaults() *OsInstall {
	this := OsInstall{}
	var classId string = "os.Install"
	this.ClassId = classId
	var objectType string = "os.Install"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *OsInstall) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *OsInstall) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *OsInstall) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "os.Install" of the ClassId field.
func (o *OsInstall) GetDefaultClassId() interface{} {
	return "os.Install"
}

// GetObjectType returns the ObjectType field value
func (o *OsInstall) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *OsInstall) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *OsInstall) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "os.Install" of the ObjectType field.
func (o *OsInstall) GetDefaultObjectType() interface{} {
	return "os.Install"
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OsInstall) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OsInstall) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OsInstall) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OsInstall) SetName(v string) {
	o.Name = &v
}

// GetConfigurationFile returns the ConfigurationFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetConfigurationFile() OsConfigurationFileRelationship {
	if o == nil || IsNil(o.ConfigurationFile.Get()) {
		var ret OsConfigurationFileRelationship
		return ret
	}
	return *o.ConfigurationFile.Get()
}

// GetConfigurationFileOk returns a tuple with the ConfigurationFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetConfigurationFileOk() (*OsConfigurationFileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationFile.Get(), o.ConfigurationFile.IsSet()
}

// HasConfigurationFile returns a boolean if a field has been set.
func (o *OsInstall) HasConfigurationFile() bool {
	if o != nil && o.ConfigurationFile.IsSet() {
		return true
	}

	return false
}

// SetConfigurationFile gets a reference to the given NullableOsConfigurationFileRelationship and assigns it to the ConfigurationFile field.
func (o *OsInstall) SetConfigurationFile(v OsConfigurationFileRelationship) {
	o.ConfigurationFile.Set(&v)
}

// SetConfigurationFileNil sets the value for ConfigurationFile to be an explicit nil
func (o *OsInstall) SetConfigurationFileNil() {
	o.ConfigurationFile.Set(nil)
}

// UnsetConfigurationFile ensures that no value is present for ConfigurationFile, not even an explicit nil
func (o *OsInstall) UnsetConfigurationFile() {
	o.ConfigurationFile.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetImage() SoftwarerepositoryOperatingSystemFileRelationship {
	if o == nil || IsNil(o.Image.Get()) {
		var ret SoftwarerepositoryOperatingSystemFileRelationship
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetImageOk() (*SoftwarerepositoryOperatingSystemFileRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *OsInstall) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableSoftwarerepositoryOperatingSystemFileRelationship and assigns it to the Image field.
func (o *OsInstall) SetImage(v SoftwarerepositoryOperatingSystemFileRelationship) {
	o.Image.Set(&v)
}

// SetImageNil sets the value for Image to be an explicit nil
func (o *OsInstall) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *OsInstall) UnsetImage() {
	o.Image.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *OsInstall) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *OsInstall) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *OsInstall) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *OsInstall) UnsetOrganization() {
	o.Organization.Unset()
}

// GetOsduImage returns the OsduImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetOsduImage() FirmwareServerConfigurationUtilityDistributableRelationship {
	if o == nil || IsNil(o.OsduImage.Get()) {
		var ret FirmwareServerConfigurationUtilityDistributableRelationship
		return ret
	}
	return *o.OsduImage.Get()
}

// GetOsduImageOk returns a tuple with the OsduImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetOsduImageOk() (*FirmwareServerConfigurationUtilityDistributableRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsduImage.Get(), o.OsduImage.IsSet()
}

// HasOsduImage returns a boolean if a field has been set.
func (o *OsInstall) HasOsduImage() bool {
	if o != nil && o.OsduImage.IsSet() {
		return true
	}

	return false
}

// SetOsduImage gets a reference to the given NullableFirmwareServerConfigurationUtilityDistributableRelationship and assigns it to the OsduImage field.
func (o *OsInstall) SetOsduImage(v FirmwareServerConfigurationUtilityDistributableRelationship) {
	o.OsduImage.Set(&v)
}

// SetOsduImageNil sets the value for OsduImage to be an explicit nil
func (o *OsInstall) SetOsduImageNil() {
	o.OsduImage.Set(nil)
}

// UnsetOsduImage ensures that no value is present for OsduImage, not even an explicit nil
func (o *OsInstall) UnsetOsduImage() {
	o.OsduImage.Unset()
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetServer() ComputePhysicalRelationship {
	if o == nil || IsNil(o.Server.Get()) {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *OsInstall) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableComputePhysicalRelationship and assigns it to the Server field.
func (o *OsInstall) SetServer(v ComputePhysicalRelationship) {
	o.Server.Set(&v)
}

// SetServerNil sets the value for Server to be an explicit nil
func (o *OsInstall) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *OsInstall) UnsetServer() {
	o.Server.Unset()
}

// GetWorkflowInfo returns the WorkflowInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OsInstall) GetWorkflowInfo() WorkflowWorkflowInfoRelationship {
	if o == nil || IsNil(o.WorkflowInfo.Get()) {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.WorkflowInfo.Get()
}

// GetWorkflowInfoOk returns a tuple with the WorkflowInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OsInstall) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowInfo.Get(), o.WorkflowInfo.IsSet()
}

// HasWorkflowInfo returns a boolean if a field has been set.
func (o *OsInstall) HasWorkflowInfo() bool {
	if o != nil && o.WorkflowInfo.IsSet() {
		return true
	}

	return false
}

// SetWorkflowInfo gets a reference to the given NullableWorkflowWorkflowInfoRelationship and assigns it to the WorkflowInfo field.
func (o *OsInstall) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship) {
	o.WorkflowInfo.Set(&v)
}

// SetWorkflowInfoNil sets the value for WorkflowInfo to be an explicit nil
func (o *OsInstall) SetWorkflowInfoNil() {
	o.WorkflowInfo.Set(nil)
}

// UnsetWorkflowInfo ensures that no value is present for WorkflowInfo, not even an explicit nil
func (o *OsInstall) UnsetWorkflowInfo() {
	o.WorkflowInfo.Unset()
}

func (o OsInstall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OsInstall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedOsBaseInstallConfig, errOsBaseInstallConfig := json.Marshal(o.OsBaseInstallConfig)
	if errOsBaseInstallConfig != nil {
		return map[string]interface{}{}, errOsBaseInstallConfig
	}
	errOsBaseInstallConfig = json.Unmarshal([]byte(serializedOsBaseInstallConfig), &toSerialize)
	if errOsBaseInstallConfig != nil {
		return map[string]interface{}{}, errOsBaseInstallConfig
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.ConfigurationFile.IsSet() {
		toSerialize["ConfigurationFile"] = o.ConfigurationFile.Get()
	}
	if o.Image.IsSet() {
		toSerialize["Image"] = o.Image.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.OsduImage.IsSet() {
		toSerialize["OsduImage"] = o.OsduImage.Get()
	}
	if o.Server.IsSet() {
		toSerialize["Server"] = o.Server.Get()
	}
	if o.WorkflowInfo.IsSet() {
		toSerialize["WorkflowInfo"] = o.WorkflowInfo.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OsInstall) UnmarshalJSON(data []byte) (err error) {
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
	type OsInstallWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The name of the OS install configuration.
		Name              *string                                                             `json:"Name,omitempty"`
		ConfigurationFile NullableOsConfigurationFileRelationship                             `json:"ConfigurationFile,omitempty"`
		Image             NullableSoftwarerepositoryOperatingSystemFileRelationship           `json:"Image,omitempty"`
		Organization      NullableOrganizationOrganizationRelationship                        `json:"Organization,omitempty"`
		OsduImage         NullableFirmwareServerConfigurationUtilityDistributableRelationship `json:"OsduImage,omitempty"`
		Server            NullableComputePhysicalRelationship                                 `json:"Server,omitempty"`
		WorkflowInfo      NullableWorkflowWorkflowInfoRelationship                            `json:"WorkflowInfo,omitempty"`
	}

	varOsInstallWithoutEmbeddedStruct := OsInstallWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varOsInstallWithoutEmbeddedStruct)
	if err == nil {
		varOsInstall := _OsInstall{}
		varOsInstall.ClassId = varOsInstallWithoutEmbeddedStruct.ClassId
		varOsInstall.ObjectType = varOsInstallWithoutEmbeddedStruct.ObjectType
		varOsInstall.Name = varOsInstallWithoutEmbeddedStruct.Name
		varOsInstall.ConfigurationFile = varOsInstallWithoutEmbeddedStruct.ConfigurationFile
		varOsInstall.Image = varOsInstallWithoutEmbeddedStruct.Image
		varOsInstall.Organization = varOsInstallWithoutEmbeddedStruct.Organization
		varOsInstall.OsduImage = varOsInstallWithoutEmbeddedStruct.OsduImage
		varOsInstall.Server = varOsInstallWithoutEmbeddedStruct.Server
		varOsInstall.WorkflowInfo = varOsInstallWithoutEmbeddedStruct.WorkflowInfo
		*o = OsInstall(varOsInstall)
	} else {
		return err
	}

	varOsInstall := _OsInstall{}

	err = json.Unmarshal(data, &varOsInstall)
	if err == nil {
		o.OsBaseInstallConfig = varOsInstall.OsBaseInstallConfig
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ConfigurationFile")
		delete(additionalProperties, "Image")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "OsduImage")
		delete(additionalProperties, "Server")
		delete(additionalProperties, "WorkflowInfo")

		// remove fields from embedded structs
		reflectOsBaseInstallConfig := reflect.ValueOf(o.OsBaseInstallConfig)
		for i := 0; i < reflectOsBaseInstallConfig.Type().NumField(); i++ {
			t := reflectOsBaseInstallConfig.Type().Field(i)

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

type NullableOsInstall struct {
	value *OsInstall
	isSet bool
}

func (v NullableOsInstall) Get() *OsInstall {
	return v.value
}

func (v *NullableOsInstall) Set(val *OsInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableOsInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableOsInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOsInstall(val *OsInstall) *NullableOsInstall {
	return &NullableOsInstall{value: val, isSet: true}
}

func (v NullableOsInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOsInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
