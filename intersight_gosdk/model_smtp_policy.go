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
)

// checks if the SmtpPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpPolicy{}

// SmtpPolicy Name that identifies the SMTP Policy.
type SmtpPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Authorization password for the process.
	AuthPassword *string `json:"AuthPassword,omitempty" validate:"regexp=^[\\\\S+]{0,254}$"`
	// If enabled, lets user input username and password.
	EnableAuth *bool `json:"EnableAuth,omitempty"`
	// If enabled, lets user input valid CA certificates for authorization.
	EnableTls *bool `json:"EnableTls,omitempty"`
	// If enabled, controls the state of the SMTP client service on the managed device.
	Enabled *bool `json:"Enabled,omitempty"`
	// Indicates whether the value of the 'authPassword' property has been set.
	IsAuthPasswordSet *bool `json:"IsAuthPasswordSet,omitempty"`
	// Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * `critical` - Minimum severity to report is critical. * `condition` - Minimum severity to report is informational. * `warning` - Minimum severity to report is warning. * `minor` - Minimum severity to report is minor. * `major` - Minimum severity to report is major.
	MinSeverity *string `json:"MinSeverity,omitempty"`
	// The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.
	SenderEmail *string "json:\"SenderEmail,omitempty\" validate:\"regexp=^$|^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$\""
	// Port number used by the SMTP server for outgoing SMTP communication.
	SmtpPort       *int64   `json:"SmtpPort,omitempty"`
	SmtpRecipients []string `json:"SmtpRecipients,omitempty"`
	// IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.
	SmtpServer *string `json:"SmtpServer,omitempty"`
	// SMTP username from which email notification is sent.
	UserName         *string                                      `json:"UserName,omitempty"`
	ApplianceAccount NullableIamAccountRelationship               `json:"ApplianceAccount,omitempty"`
	Certificate      NullableIamTrustPointRelationship            `json:"Certificate,omitempty"`
	Organization     NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to policyAbstractConfigProfile resources.
	Profiles             []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SmtpPolicy SmtpPolicy

// NewSmtpPolicy instantiates a new SmtpPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpPolicy(classId string, objectType string) *SmtpPolicy {
	this := SmtpPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enableAuth bool = false
	this.EnableAuth = &enableAuth
	var enableTls bool = false
	this.EnableTls = &enableTls
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// NewSmtpPolicyWithDefaults instantiates a new SmtpPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpPolicyWithDefaults() *SmtpPolicy {
	this := SmtpPolicy{}
	var classId string = "smtp.Policy"
	this.ClassId = classId
	var objectType string = "smtp.Policy"
	this.ObjectType = objectType
	var enableAuth bool = false
	this.EnableAuth = &enableAuth
	var enableTls bool = false
	this.EnableTls = &enableTls
	var enabled bool = true
	this.Enabled = &enabled
	var minSeverity string = "critical"
	this.MinSeverity = &minSeverity
	var smtpPort int64 = 25
	this.SmtpPort = &smtpPort
	return &this
}

// GetClassId returns the ClassId field value
func (o *SmtpPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SmtpPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "smtp.Policy" of the ClassId field.
func (o *SmtpPolicy) GetDefaultClassId() interface{} {
	return "smtp.Policy"
}

// GetObjectType returns the ObjectType field value
func (o *SmtpPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SmtpPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "smtp.Policy" of the ObjectType field.
func (o *SmtpPolicy) GetDefaultObjectType() interface{} {
	return "smtp.Policy"
}

// GetAuthPassword returns the AuthPassword field value if set, zero value otherwise.
func (o *SmtpPolicy) GetAuthPassword() string {
	if o == nil || IsNil(o.AuthPassword) {
		var ret string
		return ret
	}
	return *o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPassword) {
		return nil, false
	}
	return o.AuthPassword, true
}

// HasAuthPassword returns a boolean if a field has been set.
func (o *SmtpPolicy) HasAuthPassword() bool {
	if o != nil && !IsNil(o.AuthPassword) {
		return true
	}

	return false
}

// SetAuthPassword gets a reference to the given string and assigns it to the AuthPassword field.
func (o *SmtpPolicy) SetAuthPassword(v string) {
	o.AuthPassword = &v
}

// GetEnableAuth returns the EnableAuth field value if set, zero value otherwise.
func (o *SmtpPolicy) GetEnableAuth() bool {
	if o == nil || IsNil(o.EnableAuth) {
		var ret bool
		return ret
	}
	return *o.EnableAuth
}

// GetEnableAuthOk returns a tuple with the EnableAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetEnableAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAuth) {
		return nil, false
	}
	return o.EnableAuth, true
}

// HasEnableAuth returns a boolean if a field has been set.
func (o *SmtpPolicy) HasEnableAuth() bool {
	if o != nil && !IsNil(o.EnableAuth) {
		return true
	}

	return false
}

// SetEnableAuth gets a reference to the given bool and assigns it to the EnableAuth field.
func (o *SmtpPolicy) SetEnableAuth(v bool) {
	o.EnableAuth = &v
}

// GetEnableTls returns the EnableTls field value if set, zero value otherwise.
func (o *SmtpPolicy) GetEnableTls() bool {
	if o == nil || IsNil(o.EnableTls) {
		var ret bool
		return ret
	}
	return *o.EnableTls
}

// GetEnableTlsOk returns a tuple with the EnableTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetEnableTlsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTls) {
		return nil, false
	}
	return o.EnableTls, true
}

// HasEnableTls returns a boolean if a field has been set.
func (o *SmtpPolicy) HasEnableTls() bool {
	if o != nil && !IsNil(o.EnableTls) {
		return true
	}

	return false
}

// SetEnableTls gets a reference to the given bool and assigns it to the EnableTls field.
func (o *SmtpPolicy) SetEnableTls(v bool) {
	o.EnableTls = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SmtpPolicy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SmtpPolicy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SmtpPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIsAuthPasswordSet returns the IsAuthPasswordSet field value if set, zero value otherwise.
func (o *SmtpPolicy) GetIsAuthPasswordSet() bool {
	if o == nil || IsNil(o.IsAuthPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsAuthPasswordSet
}

// GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetIsAuthPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAuthPasswordSet) {
		return nil, false
	}
	return o.IsAuthPasswordSet, true
}

// HasIsAuthPasswordSet returns a boolean if a field has been set.
func (o *SmtpPolicy) HasIsAuthPasswordSet() bool {
	if o != nil && !IsNil(o.IsAuthPasswordSet) {
		return true
	}

	return false
}

// SetIsAuthPasswordSet gets a reference to the given bool and assigns it to the IsAuthPasswordSet field.
func (o *SmtpPolicy) SetIsAuthPasswordSet(v bool) {
	o.IsAuthPasswordSet = &v
}

// GetMinSeverity returns the MinSeverity field value if set, zero value otherwise.
func (o *SmtpPolicy) GetMinSeverity() string {
	if o == nil || IsNil(o.MinSeverity) {
		var ret string
		return ret
	}
	return *o.MinSeverity
}

// GetMinSeverityOk returns a tuple with the MinSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetMinSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.MinSeverity) {
		return nil, false
	}
	return o.MinSeverity, true
}

// HasMinSeverity returns a boolean if a field has been set.
func (o *SmtpPolicy) HasMinSeverity() bool {
	if o != nil && !IsNil(o.MinSeverity) {
		return true
	}

	return false
}

// SetMinSeverity gets a reference to the given string and assigns it to the MinSeverity field.
func (o *SmtpPolicy) SetMinSeverity(v string) {
	o.MinSeverity = &v
}

// GetSenderEmail returns the SenderEmail field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSenderEmail() string {
	if o == nil || IsNil(o.SenderEmail) {
		var ret string
		return ret
	}
	return *o.SenderEmail
}

// GetSenderEmailOk returns a tuple with the SenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSenderEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SenderEmail) {
		return nil, false
	}
	return o.SenderEmail, true
}

// HasSenderEmail returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSenderEmail() bool {
	if o != nil && !IsNil(o.SenderEmail) {
		return true
	}

	return false
}

// SetSenderEmail gets a reference to the given string and assigns it to the SenderEmail field.
func (o *SmtpPolicy) SetSenderEmail(v string) {
	o.SenderEmail = &v
}

// GetSmtpPort returns the SmtpPort field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSmtpPort() int64 {
	if o == nil || IsNil(o.SmtpPort) {
		var ret int64
		return ret
	}
	return *o.SmtpPort
}

// GetSmtpPortOk returns a tuple with the SmtpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSmtpPortOk() (*int64, bool) {
	if o == nil || IsNil(o.SmtpPort) {
		return nil, false
	}
	return o.SmtpPort, true
}

// HasSmtpPort returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpPort() bool {
	if o != nil && !IsNil(o.SmtpPort) {
		return true
	}

	return false
}

// SetSmtpPort gets a reference to the given int64 and assigns it to the SmtpPort field.
func (o *SmtpPolicy) SetSmtpPort(v int64) {
	o.SmtpPort = &v
}

// GetSmtpRecipients returns the SmtpRecipients field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetSmtpRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SmtpRecipients
}

// GetSmtpRecipientsOk returns a tuple with the SmtpRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetSmtpRecipientsOk() ([]string, bool) {
	if o == nil || IsNil(o.SmtpRecipients) {
		return nil, false
	}
	return o.SmtpRecipients, true
}

// HasSmtpRecipients returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpRecipients() bool {
	if o != nil && !IsNil(o.SmtpRecipients) {
		return true
	}

	return false
}

// SetSmtpRecipients gets a reference to the given []string and assigns it to the SmtpRecipients field.
func (o *SmtpPolicy) SetSmtpRecipients(v []string) {
	o.SmtpRecipients = v
}

// GetSmtpServer returns the SmtpServer field value if set, zero value otherwise.
func (o *SmtpPolicy) GetSmtpServer() string {
	if o == nil || IsNil(o.SmtpServer) {
		var ret string
		return ret
	}
	return *o.SmtpServer
}

// GetSmtpServerOk returns a tuple with the SmtpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetSmtpServerOk() (*string, bool) {
	if o == nil || IsNil(o.SmtpServer) {
		return nil, false
	}
	return o.SmtpServer, true
}

// HasSmtpServer returns a boolean if a field has been set.
func (o *SmtpPolicy) HasSmtpServer() bool {
	if o != nil && !IsNil(o.SmtpServer) {
		return true
	}

	return false
}

// SetSmtpServer gets a reference to the given string and assigns it to the SmtpServer field.
func (o *SmtpPolicy) SetSmtpServer(v string) {
	o.SmtpServer = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SmtpPolicy) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpPolicy) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SmtpPolicy) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SmtpPolicy) SetUserName(v string) {
	o.UserName = &v
}

// GetApplianceAccount returns the ApplianceAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetApplianceAccount() IamAccountRelationship {
	if o == nil || IsNil(o.ApplianceAccount.Get()) {
		var ret IamAccountRelationship
		return ret
	}
	return *o.ApplianceAccount.Get()
}

// GetApplianceAccountOk returns a tuple with the ApplianceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetApplianceAccountOk() (*IamAccountRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplianceAccount.Get(), o.ApplianceAccount.IsSet()
}

// HasApplianceAccount returns a boolean if a field has been set.
func (o *SmtpPolicy) HasApplianceAccount() bool {
	if o != nil && o.ApplianceAccount.IsSet() {
		return true
	}

	return false
}

// SetApplianceAccount gets a reference to the given NullableIamAccountRelationship and assigns it to the ApplianceAccount field.
func (o *SmtpPolicy) SetApplianceAccount(v IamAccountRelationship) {
	o.ApplianceAccount.Set(&v)
}

// SetApplianceAccountNil sets the value for ApplianceAccount to be an explicit nil
func (o *SmtpPolicy) SetApplianceAccountNil() {
	o.ApplianceAccount.Set(nil)
}

// UnsetApplianceAccount ensures that no value is present for ApplianceAccount, not even an explicit nil
func (o *SmtpPolicy) UnsetApplianceAccount() {
	o.ApplianceAccount.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetCertificate() IamTrustPointRelationship {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret IamTrustPointRelationship
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetCertificateOk() (*IamTrustPointRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *SmtpPolicy) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableIamTrustPointRelationship and assigns it to the Certificate field.
func (o *SmtpPolicy) SetCertificate(v IamTrustPointRelationship) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *SmtpPolicy) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *SmtpPolicy) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *SmtpPolicy) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableOrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SmtpPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization.Set(&v)
}

// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *SmtpPolicy) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *SmtpPolicy) UnsetOrganization() {
	o.Organization.Unset()
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpPolicy) GetProfiles() []PolicyAbstractConfigProfileRelationship {
	if o == nil {
		var ret []PolicyAbstractConfigProfileRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpPolicy) GetProfilesOk() ([]PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *SmtpPolicy) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []PolicyAbstractConfigProfileRelationship and assigns it to the Profiles field.
func (o *SmtpPolicy) SetProfiles(v []PolicyAbstractConfigProfileRelationship) {
	o.Profiles = v
}

func (o SmtpPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.AuthPassword) {
		toSerialize["AuthPassword"] = o.AuthPassword
	}
	if !IsNil(o.EnableAuth) {
		toSerialize["EnableAuth"] = o.EnableAuth
	}
	if !IsNil(o.EnableTls) {
		toSerialize["EnableTls"] = o.EnableTls
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.IsAuthPasswordSet) {
		toSerialize["IsAuthPasswordSet"] = o.IsAuthPasswordSet
	}
	if !IsNil(o.MinSeverity) {
		toSerialize["MinSeverity"] = o.MinSeverity
	}
	if !IsNil(o.SenderEmail) {
		toSerialize["SenderEmail"] = o.SenderEmail
	}
	if !IsNil(o.SmtpPort) {
		toSerialize["SmtpPort"] = o.SmtpPort
	}
	if o.SmtpRecipients != nil {
		toSerialize["SmtpRecipients"] = o.SmtpRecipients
	}
	if !IsNil(o.SmtpServer) {
		toSerialize["SmtpServer"] = o.SmtpServer
	}
	if !IsNil(o.UserName) {
		toSerialize["UserName"] = o.UserName
	}
	if o.ApplianceAccount.IsSet() {
		toSerialize["ApplianceAccount"] = o.ApplianceAccount.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["Certificate"] = o.Certificate.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["Organization"] = o.Organization.Get()
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SmtpPolicy) UnmarshalJSON(data []byte) (err error) {
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
	type SmtpPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Authorization password for the process.
		AuthPassword *string `json:"AuthPassword,omitempty" validate:"regexp=^[\\\\S+]{0,254}$"`
		// If enabled, lets user input username and password.
		EnableAuth *bool `json:"EnableAuth,omitempty"`
		// If enabled, lets user input valid CA certificates for authorization.
		EnableTls *bool `json:"EnableTls,omitempty"`
		// If enabled, controls the state of the SMTP client service on the managed device.
		Enabled *bool `json:"Enabled,omitempty"`
		// Indicates whether the value of the 'authPassword' property has been set.
		IsAuthPasswordSet *bool `json:"IsAuthPasswordSet,omitempty"`
		// Minimum fault severity level to receive email notifications. Email notifications are sent for all faults whose severity is equal to or greater than the chosen level. * `critical` - Minimum severity to report is critical. * `condition` - Minimum severity to report is informational. * `warning` - Minimum severity to report is warning. * `minor` - Minimum severity to report is minor. * `major` - Minimum severity to report is major.
		MinSeverity *string `json:"MinSeverity,omitempty"`
		// The email address entered here will be displayed as the from address (mail received from address) of all the SMTP mail alerts that are received. If not configured, the hostname of the server is used in the from address field.
		SenderEmail *string "json:\"SenderEmail,omitempty\" validate:\"regexp=^$|^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$\""
		// Port number used by the SMTP server for outgoing SMTP communication.
		SmtpPort       *int64   `json:"SmtpPort,omitempty"`
		SmtpRecipients []string `json:"SmtpRecipients,omitempty"`
		// IP address or hostname of the SMTP server. The SMTP server is used by the managed device to send email notifications.
		SmtpServer *string `json:"SmtpServer,omitempty"`
		// SMTP username from which email notification is sent.
		UserName         *string                                      `json:"UserName,omitempty"`
		ApplianceAccount NullableIamAccountRelationship               `json:"ApplianceAccount,omitempty"`
		Certificate      NullableIamTrustPointRelationship            `json:"Certificate,omitempty"`
		Organization     NullableOrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to policyAbstractConfigProfile resources.
		Profiles []PolicyAbstractConfigProfileRelationship `json:"Profiles,omitempty"`
	}

	varSmtpPolicyWithoutEmbeddedStruct := SmtpPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSmtpPolicyWithoutEmbeddedStruct)
	if err == nil {
		varSmtpPolicy := _SmtpPolicy{}
		varSmtpPolicy.ClassId = varSmtpPolicyWithoutEmbeddedStruct.ClassId
		varSmtpPolicy.ObjectType = varSmtpPolicyWithoutEmbeddedStruct.ObjectType
		varSmtpPolicy.AuthPassword = varSmtpPolicyWithoutEmbeddedStruct.AuthPassword
		varSmtpPolicy.EnableAuth = varSmtpPolicyWithoutEmbeddedStruct.EnableAuth
		varSmtpPolicy.EnableTls = varSmtpPolicyWithoutEmbeddedStruct.EnableTls
		varSmtpPolicy.Enabled = varSmtpPolicyWithoutEmbeddedStruct.Enabled
		varSmtpPolicy.IsAuthPasswordSet = varSmtpPolicyWithoutEmbeddedStruct.IsAuthPasswordSet
		varSmtpPolicy.MinSeverity = varSmtpPolicyWithoutEmbeddedStruct.MinSeverity
		varSmtpPolicy.SenderEmail = varSmtpPolicyWithoutEmbeddedStruct.SenderEmail
		varSmtpPolicy.SmtpPort = varSmtpPolicyWithoutEmbeddedStruct.SmtpPort
		varSmtpPolicy.SmtpRecipients = varSmtpPolicyWithoutEmbeddedStruct.SmtpRecipients
		varSmtpPolicy.SmtpServer = varSmtpPolicyWithoutEmbeddedStruct.SmtpServer
		varSmtpPolicy.UserName = varSmtpPolicyWithoutEmbeddedStruct.UserName
		varSmtpPolicy.ApplianceAccount = varSmtpPolicyWithoutEmbeddedStruct.ApplianceAccount
		varSmtpPolicy.Certificate = varSmtpPolicyWithoutEmbeddedStruct.Certificate
		varSmtpPolicy.Organization = varSmtpPolicyWithoutEmbeddedStruct.Organization
		varSmtpPolicy.Profiles = varSmtpPolicyWithoutEmbeddedStruct.Profiles
		*o = SmtpPolicy(varSmtpPolicy)
	} else {
		return err
	}

	varSmtpPolicy := _SmtpPolicy{}

	err = json.Unmarshal(data, &varSmtpPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varSmtpPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AuthPassword")
		delete(additionalProperties, "EnableAuth")
		delete(additionalProperties, "EnableTls")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "IsAuthPasswordSet")
		delete(additionalProperties, "MinSeverity")
		delete(additionalProperties, "SenderEmail")
		delete(additionalProperties, "SmtpPort")
		delete(additionalProperties, "SmtpRecipients")
		delete(additionalProperties, "SmtpServer")
		delete(additionalProperties, "UserName")
		delete(additionalProperties, "ApplianceAccount")
		delete(additionalProperties, "Certificate")
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

type NullableSmtpPolicy struct {
	value *SmtpPolicy
	isSet bool
}

func (v NullableSmtpPolicy) Get() *SmtpPolicy {
	return v.value
}

func (v *NullableSmtpPolicy) Set(val *SmtpPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpPolicy(val *SmtpPolicy) *NullableSmtpPolicy {
	return &NullableSmtpPolicy{value: val, isSet: true}
}

func (v NullableSmtpPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
