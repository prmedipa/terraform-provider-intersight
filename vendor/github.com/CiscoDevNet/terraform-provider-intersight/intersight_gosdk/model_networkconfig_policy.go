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

// NetworkconfigPolicy Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
type NetworkconfigPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address of the secondary DNS server.
	AlternateIpv4dnsServer *string `json:"AlternateIpv4dnsServer,omitempty"`
	// IP address of the secondary DNS server.
	AlternateIpv6dnsServer *string `json:"AlternateIpv6dnsServer,omitempty"`
	// The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.
	DynamicDnsDomain *string `json:"DynamicDnsDomain,omitempty"`
	// If enabled, updates the resource records to the DNS from Cisco IMC.
	EnableDynamicDns *bool `json:"EnableDynamicDns,omitempty"`
	// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.
	EnableIpv4dnsFromDhcp *bool `json:"EnableIpv4dnsFromDhcp,omitempty"`
	// If enabled, allows to configure IPv6 properties.
	EnableIpv6 *bool `json:"EnableIpv6,omitempty"`
	// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.
	EnableIpv6dnsFromDhcp *bool `json:"EnableIpv6dnsFromDhcp,omitempty"`
	// IP address of the primary DNS server.
	PreferredIpv4dnsServer *string `json:"PreferredIpv4dnsServer,omitempty"`
	// IP address of the primary DNS server.
	PreferredIpv6dnsServer *string                               `json:"PreferredIpv6dnsServer,omitempty"`
	ApplianceAccount       *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Organization           *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkconfigPolicy NetworkconfigPolicy

// NewNetworkconfigPolicy instantiates a new NetworkconfigPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkconfigPolicy(classId string, objectType string) *NetworkconfigPolicy {
	this := NetworkconfigPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNetworkconfigPolicyWithDefaults instantiates a new NetworkconfigPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkconfigPolicyWithDefaults() *NetworkconfigPolicy {
	this := NetworkconfigPolicy{}
	var classId string = "networkconfig.Policy"
	this.ClassId = classId
	var objectType string = "networkconfig.Policy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NetworkconfigPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NetworkconfigPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NetworkconfigPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NetworkconfigPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAlternateIpv4dnsServer returns the AlternateIpv4dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetAlternateIpv4dnsServer() string {
	if o == nil || o.AlternateIpv4dnsServer == nil {
		var ret string
		return ret
	}
	return *o.AlternateIpv4dnsServer
}

// GetAlternateIpv4dnsServerOk returns a tuple with the AlternateIpv4dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetAlternateIpv4dnsServerOk() (*string, bool) {
	if o == nil || o.AlternateIpv4dnsServer == nil {
		return nil, false
	}
	return o.AlternateIpv4dnsServer, true
}

// HasAlternateIpv4dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasAlternateIpv4dnsServer() bool {
	if o != nil && o.AlternateIpv4dnsServer != nil {
		return true
	}

	return false
}

// SetAlternateIpv4dnsServer gets a reference to the given string and assigns it to the AlternateIpv4dnsServer field.
func (o *NetworkconfigPolicy) SetAlternateIpv4dnsServer(v string) {
	o.AlternateIpv4dnsServer = &v
}

// GetAlternateIpv6dnsServer returns the AlternateIpv6dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetAlternateIpv6dnsServer() string {
	if o == nil || o.AlternateIpv6dnsServer == nil {
		var ret string
		return ret
	}
	return *o.AlternateIpv6dnsServer
}

// GetAlternateIpv6dnsServerOk returns a tuple with the AlternateIpv6dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetAlternateIpv6dnsServerOk() (*string, bool) {
	if o == nil || o.AlternateIpv6dnsServer == nil {
		return nil, false
	}
	return o.AlternateIpv6dnsServer, true
}

// HasAlternateIpv6dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasAlternateIpv6dnsServer() bool {
	if o != nil && o.AlternateIpv6dnsServer != nil {
		return true
	}

	return false
}

// SetAlternateIpv6dnsServer gets a reference to the given string and assigns it to the AlternateIpv6dnsServer field.
func (o *NetworkconfigPolicy) SetAlternateIpv6dnsServer(v string) {
	o.AlternateIpv6dnsServer = &v
}

// GetDynamicDnsDomain returns the DynamicDnsDomain field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetDynamicDnsDomain() string {
	if o == nil || o.DynamicDnsDomain == nil {
		var ret string
		return ret
	}
	return *o.DynamicDnsDomain
}

// GetDynamicDnsDomainOk returns a tuple with the DynamicDnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetDynamicDnsDomainOk() (*string, bool) {
	if o == nil || o.DynamicDnsDomain == nil {
		return nil, false
	}
	return o.DynamicDnsDomain, true
}

// HasDynamicDnsDomain returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasDynamicDnsDomain() bool {
	if o != nil && o.DynamicDnsDomain != nil {
		return true
	}

	return false
}

// SetDynamicDnsDomain gets a reference to the given string and assigns it to the DynamicDnsDomain field.
func (o *NetworkconfigPolicy) SetDynamicDnsDomain(v string) {
	o.DynamicDnsDomain = &v
}

// GetEnableDynamicDns returns the EnableDynamicDns field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetEnableDynamicDns() bool {
	if o == nil || o.EnableDynamicDns == nil {
		var ret bool
		return ret
	}
	return *o.EnableDynamicDns
}

// GetEnableDynamicDnsOk returns a tuple with the EnableDynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetEnableDynamicDnsOk() (*bool, bool) {
	if o == nil || o.EnableDynamicDns == nil {
		return nil, false
	}
	return o.EnableDynamicDns, true
}

// HasEnableDynamicDns returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasEnableDynamicDns() bool {
	if o != nil && o.EnableDynamicDns != nil {
		return true
	}

	return false
}

// SetEnableDynamicDns gets a reference to the given bool and assigns it to the EnableDynamicDns field.
func (o *NetworkconfigPolicy) SetEnableDynamicDns(v bool) {
	o.EnableDynamicDns = &v
}

// GetEnableIpv4dnsFromDhcp returns the EnableIpv4dnsFromDhcp field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetEnableIpv4dnsFromDhcp() bool {
	if o == nil || o.EnableIpv4dnsFromDhcp == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv4dnsFromDhcp
}

// GetEnableIpv4dnsFromDhcpOk returns a tuple with the EnableIpv4dnsFromDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetEnableIpv4dnsFromDhcpOk() (*bool, bool) {
	if o == nil || o.EnableIpv4dnsFromDhcp == nil {
		return nil, false
	}
	return o.EnableIpv4dnsFromDhcp, true
}

// HasEnableIpv4dnsFromDhcp returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasEnableIpv4dnsFromDhcp() bool {
	if o != nil && o.EnableIpv4dnsFromDhcp != nil {
		return true
	}

	return false
}

// SetEnableIpv4dnsFromDhcp gets a reference to the given bool and assigns it to the EnableIpv4dnsFromDhcp field.
func (o *NetworkconfigPolicy) SetEnableIpv4dnsFromDhcp(v bool) {
	o.EnableIpv4dnsFromDhcp = &v
}

// GetEnableIpv6 returns the EnableIpv6 field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetEnableIpv6() bool {
	if o == nil || o.EnableIpv6 == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv6
}

// GetEnableIpv6Ok returns a tuple with the EnableIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetEnableIpv6Ok() (*bool, bool) {
	if o == nil || o.EnableIpv6 == nil {
		return nil, false
	}
	return o.EnableIpv6, true
}

// HasEnableIpv6 returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasEnableIpv6() bool {
	if o != nil && o.EnableIpv6 != nil {
		return true
	}

	return false
}

// SetEnableIpv6 gets a reference to the given bool and assigns it to the EnableIpv6 field.
func (o *NetworkconfigPolicy) SetEnableIpv6(v bool) {
	o.EnableIpv6 = &v
}

// GetEnableIpv6dnsFromDhcp returns the EnableIpv6dnsFromDhcp field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetEnableIpv6dnsFromDhcp() bool {
	if o == nil || o.EnableIpv6dnsFromDhcp == nil {
		var ret bool
		return ret
	}
	return *o.EnableIpv6dnsFromDhcp
}

// GetEnableIpv6dnsFromDhcpOk returns a tuple with the EnableIpv6dnsFromDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetEnableIpv6dnsFromDhcpOk() (*bool, bool) {
	if o == nil || o.EnableIpv6dnsFromDhcp == nil {
		return nil, false
	}
	return o.EnableIpv6dnsFromDhcp, true
}

// HasEnableIpv6dnsFromDhcp returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasEnableIpv6dnsFromDhcp() bool {
	if o != nil && o.EnableIpv6dnsFromDhcp != nil {
		return true
	}

	return false
}

// SetEnableIpv6dnsFromDhcp gets a reference to the given bool and assigns it to the EnableIpv6dnsFromDhcp field.
func (o *NetworkconfigPolicy) SetEnableIpv6dnsFromDhcp(v bool) {
	o.EnableIpv6dnsFromDhcp = &v
}

// GetPreferredIpv4dnsServer returns the PreferredIpv4dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetPreferredIpv4dnsServer() string {
	if o == nil || o.PreferredIpv4dnsServer == nil {
		var ret string
		return ret
	}
	return *o.PreferredIpv4dnsServer
}

// GetPreferredIpv4dnsServerOk returns a tuple with the PreferredIpv4dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetPreferredIpv4dnsServerOk() (*string, bool) {
	if o == nil || o.PreferredIpv4dnsServer == nil {
		return nil, false
	}
	return o.PreferredIpv4dnsServer, true
}

// HasPreferredIpv4dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasPreferredIpv4dnsServer() bool {
	if o != nil && o.PreferredIpv4dnsServer != nil {
		return true
	}

	return false
}

// SetPreferredIpv4dnsServer gets a reference to the given string and assigns it to the PreferredIpv4dnsServer field.
func (o *NetworkconfigPolicy) SetPreferredIpv4dnsServer(v string) {
	o.PreferredIpv4dnsServer = &v
}

// GetPreferredIpv6dnsServer returns the PreferredIpv6dnsServer field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetPreferredIpv6dnsServer() string {
	if o == nil || o.PreferredIpv6dnsServer == nil {
		var ret string
		return ret
	}
	return *o.PreferredIpv6dnsServer
}

// GetPreferredIpv6dnsServerOk returns a tuple with the PreferredIpv6dnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetPreferredIpv6dnsServerOk() (*string, bool) {
	if o == nil || o.PreferredIpv6dnsServer == nil {
		return nil, false
	}
	return o.PreferredIpv6dnsServer, true
}

// HasPreferredIpv6dnsServer returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasPreferredIpv6dnsServer() bool {
	if o != nil && o.PreferredIpv6dnsServer != nil {
		return true
	}

	return false
}

// SetPreferredIpv6dnsServer gets a reference to the given string and assigns it to the PreferredIpv6dnsServer field.
func (o *NetworkconfigPolicy) SetPreferredIpv6dnsServer(v string) {
	o.PreferredIpv6dnsServer = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetApplianceAccount() IamAccountRelationship {
	if o == nil || o.ApplianceAccount == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.ApplianceAccount == nil {
		return nil, false
	}
	return o.ApplianceAccount, true
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount != nil {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given IamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *NetworkconfigPolicy) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *NetworkconfigPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkconfigPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *NetworkconfigPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkconfigPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkconfigPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *NetworkconfigPolicy) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *NetworkconfigPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o NetworkconfigPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AlternateIpv4dnsServer != nil {
		toSerialize["AlternateIpv4dnsServer"] = o.AlternateIpv4dnsServer
	}
	if o.AlternateIpv6dnsServer != nil {
		toSerialize["AlternateIpv6dnsServer"] = o.AlternateIpv6dnsServer
	}
	if o.DynamicDnsDomain != nil {
		toSerialize["DynamicDnsDomain"] = o.DynamicDnsDomain
	}
	if o.EnableDynamicDns != nil {
		toSerialize["EnableDynamicDns"] = o.EnableDynamicDns
	}
	if o.EnableIpv4dnsFromDhcp != nil {
		toSerialize["EnableIpv4dnsFromDhcp"] = o.EnableIpv4dnsFromDhcp
	}
	if o.EnableIpv6 != nil {
		toSerialize["EnableIpv6"] = o.EnableIpv6
	}
	if o.EnableIpv6dnsFromDhcp != nil {
		toSerialize["EnableIpv6dnsFromDhcp"] = o.EnableIpv6dnsFromDhcp
	}
	if o.PreferredIpv4dnsServer != nil {
		toSerialize["PreferredIpv4dnsServer"] = o.PreferredIpv4dnsServer
	}
	if o.PreferredIpv6dnsServer != nil {
		toSerialize["PreferredIpv6dnsServer"] = o.PreferredIpv6dnsServer
	}
	if o.ApplianceAccount != nil {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NetworkconfigPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type NetworkconfigPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address of the secondary DNS server.
		AlternateIpv4dnsServer *string `json:"AlternateIpv4dnsServer,omitempty"`
		// IP address of the secondary DNS server.
		AlternateIpv6dnsServer *string `json:"AlternateIpv6dnsServer,omitempty"`
		// The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.
		DynamicDnsDomain *string `json:"DynamicDnsDomain,omitempty"`
		// If enabled, updates the resource records to the DNS from Cisco IMC.
		EnableDynamicDns *bool `json:"EnableDynamicDns,omitempty"`
		// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.
		EnableIpv4dnsFromDhcp *bool `json:"EnableIpv4dnsFromDhcp,omitempty"`
		// If enabled, allows to configure IPv6 properties.
		EnableIpv6 *bool `json:"EnableIpv6,omitempty"`
		// If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.
		EnableIpv6dnsFromDhcp *bool `json:"EnableIpv6dnsFromDhcp,omitempty"`
		// IP address of the primary DNS server.
		PreferredIpv4dnsServer *string `json:"PreferredIpv4dnsServer,omitempty"`
		// IP address of the primary DNS server.
		PreferredIpv6dnsServer *string                               `json:"PreferredIpv6dnsServer,omitempty"`
		ApplianceAccount       *IamAccountRelationship               `json:"ApplianceAccount,omitempty"`
		Organization           *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varNetworkconfigPolicyWithoutEmbeddedStruct := NetworkconfigPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNetworkconfigPolicyWithoutEmbeddedStruct)
	if err == nil {
		varNetworkconfigPolicy := _NetworkconfigPolicy{}
		varNetworkconfigPolicy.ClassId = varNetworkconfigPolicyWithoutEmbeddedStruct.ClassId
		varNetworkconfigPolicy.ObjectType = varNetworkconfigPolicyWithoutEmbeddedStruct.ObjectType
		varNetworkconfigPolicy.AlternateIpv4dnsServer = varNetworkconfigPolicyWithoutEmbeddedStruct.AlternateIpv4dnsServer
		varNetworkconfigPolicy.AlternateIpv6dnsServer = varNetworkconfigPolicyWithoutEmbeddedStruct.AlternateIpv6dnsServer
		varNetworkconfigPolicy.DynamicDnsDomain = varNetworkconfigPolicyWithoutEmbeddedStruct.DynamicDnsDomain
		varNetworkconfigPolicy.EnableDynamicDns = varNetworkconfigPolicyWithoutEmbeddedStruct.EnableDynamicDns
		varNetworkconfigPolicy.EnableIpv4dnsFromDhcp = varNetworkconfigPolicyWithoutEmbeddedStruct.EnableIpv4dnsFromDhcp
		varNetworkconfigPolicy.EnableIpv6 = varNetworkconfigPolicyWithoutEmbeddedStruct.EnableIpv6
		varNetworkconfigPolicy.EnableIpv6dnsFromDhcp = varNetworkconfigPolicyWithoutEmbeddedStruct.EnableIpv6dnsFromDhcp
		varNetworkconfigPolicy.PreferredIpv4dnsServer = varNetworkconfigPolicyWithoutEmbeddedStruct.PreferredIpv4dnsServer
		varNetworkconfigPolicy.PreferredIpv6dnsServer = varNetworkconfigPolicyWithoutEmbeddedStruct.PreferredIpv6dnsServer
		varNetworkconfigPolicy.ApplianceAccount = varNetworkconfigPolicyWithoutEmbeddedStruct.ApplianceAccount
		varNetworkconfigPolicy.Organization = varNetworkconfigPolicyWithoutEmbeddedStruct.Organization
		varNetworkconfigPolicy.Profiles = varNetworkconfigPolicyWithoutEmbeddedStruct.Profiles
		*o = NetworkconfigPolicy(varNetworkconfigPolicy)
	} else {
		return err
	}

	varNetworkconfigPolicy := _NetworkconfigPolicy{}

	err = json.Unmarshal(bytes, &varNetworkconfigPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varNetworkconfigPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AlternateIpv4dnsServer")
		delete(additionalProperties, "AlternateIpv6dnsServer")
		delete(additionalProperties, "DynamicDnsDomain")
		delete(additionalProperties, "EnableDynamicDns")
		delete(additionalProperties, "EnableIpv4dnsFromDhcp")
		delete(additionalProperties, "EnableIpv6")
		delete(additionalProperties, "EnableIpv6dnsFromDhcp")
		delete(additionalProperties, "PreferredIpv4dnsServer")
		delete(additionalProperties, "PreferredIpv6dnsServer")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableNetworkconfigPolicy struct {
	value *NetworkconfigPolicy
	isSet bool
}

func (v NullableNetworkconfigPolicy) Get() *NetworkconfigPolicy {
	return v.value
}

func (v *NullableNetworkconfigPolicy) Set(val *NetworkconfigPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkconfigPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkconfigPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkconfigPolicy(val *NetworkconfigPolicy) *NullableNetworkconfigPolicy {
	return &NullableNetworkconfigPolicy{value: val, isSet: true}
}

func (v NullableNetworkconfigPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkconfigPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
