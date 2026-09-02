package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DnsIps() *[]*string
	// Experimental.
	SetDnsIps(val *[]*string)
	// Experimental.
	DnsIpsInput() *[]*string
	// Experimental.
	DomainJoinServiceAccountSecret() *string
	// Experimental.
	SetDomainJoinServiceAccountSecret(val *string)
	// Experimental.
	DomainJoinServiceAccountSecretInput() *string
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	FileSystemAdministratorsGroup() *string
	// Experimental.
	SetFileSystemAdministratorsGroup(val *string)
	// Experimental.
	FileSystemAdministratorsGroupInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty
	// Experimental.
	SetInternalValue(val *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty)
	// Experimental.
	OrganizationalUnitDistinguishedName() *string
	// Experimental.
	SetOrganizationalUnitDistinguishedName(val *string)
	// Experimental.
	OrganizationalUnitDistinguishedNameInput() *string
	// Experimental.
	Password() *string
	// Experimental.
	SetPassword(val *string)
	// Experimental.
	PasswordInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	PasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetPasswordWo(val *string)
	// Experimental.
	PasswordWoInput() *string
	// Experimental.
	PasswordWoVersion() *float64
	// Experimental.
	SetPasswordWoVersion(val *float64)
	// Experimental.
	PasswordWoVersionInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Username() *string
	// Experimental.
	SetUsername(val *string)
	// Experimental.
	UsernameInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	ResetDomainJoinServiceAccountSecret()
	// Experimental.
	ResetFileSystemAdministratorsGroup()
	// Experimental.
	ResetOrganizationalUnitDistinguishedName()
	// Experimental.
	ResetPassword()
	// Experimental.
	ResetPasswordWo()
	// Experimental.
	ResetPasswordWoVersion()
	// Experimental.
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference
type jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DomainJoinServiceAccountSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainJoinServiceAccountSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DomainJoinServiceAccountSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainJoinServiceAccountSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) FileSystemAdministratorsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) FileSystemAdministratorsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) InternalValue() *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty {
	var returns *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) PasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) PasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfWindowsFileSystem.SelfManagedActiveDirectoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference_Override(t TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfWindowsFileSystem.SelfManagedActiveDirectoryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetDnsIps(val *[]*string) {
	if err := j.validateSetDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsIps",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetDomainJoinServiceAccountSecret(val *string) {
	if err := j.validateSetDomainJoinServiceAccountSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainJoinServiceAccountSecret",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetFileSystemAdministratorsGroup(val *string) {
	if err := j.validateSetFileSystemAdministratorsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemAdministratorsGroup",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetInternalValue(val *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetOrganizationalUnitDistinguishedName(val *string) {
	if err := j.validateSetOrganizationalUnitDistinguishedNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetPasswordWoVersion(val *float64) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetDomainJoinServiceAccountSecret() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainJoinServiceAccountSecret",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetFileSystemAdministratorsGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemAdministratorsGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_SelfManagedActiveDirectoryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

