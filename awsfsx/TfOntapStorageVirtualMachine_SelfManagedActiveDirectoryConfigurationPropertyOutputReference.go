package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty)
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
	ResetFileSystemAdministratorsGroup()
	// Experimental.
	ResetOrganizationalUnitDistinguishedName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference
type jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) DnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) DnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) FileSystemAdministratorsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) FileSystemAdministratorsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemAdministratorsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) InternalValue() *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty {
	var returns *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) OrganizationalUnitDistinguishedName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) OrganizationalUnitDistinguishedNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitDistinguishedNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference_Override(t TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapStorageVirtualMachine.SelfManagedActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetDnsIps(val *[]*string) {
	if err := j.validateSetDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsIps",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetFileSystemAdministratorsGroup(val *string) {
	if err := j.validateSetFileSystemAdministratorsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemAdministratorsGroup",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetInternalValue(val *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetOrganizationalUnitDistinguishedName(val *string) {
	if err := j.validateSetOrganizationalUnitDistinguishedNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitDistinguishedName",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ResetFileSystemAdministratorsGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemAdministratorsGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ResetOrganizationalUnitDistinguishedName() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationalUnitDistinguishedName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

