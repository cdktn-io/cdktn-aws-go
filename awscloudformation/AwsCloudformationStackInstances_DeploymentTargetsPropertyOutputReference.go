package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountFilterType() *string
	// Experimental.
	SetAccountFilterType(val *string)
	// Experimental.
	AccountFilterTypeInput() *string
	// Experimental.
	Accounts() *[]*string
	// Experimental.
	SetAccounts(val *[]*string)
	// Experimental.
	AccountsInput() *[]*string
	// Experimental.
	AccountsUrl() *string
	// Experimental.
	SetAccountsUrl(val *string)
	// Experimental.
	AccountsUrlInput() *string
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCloudformationStackInstances_DeploymentTargetsProperty
	// Experimental.
	SetInternalValue(val *AwsCloudformationStackInstances_DeploymentTargetsProperty)
	// Experimental.
	OrganizationalUnitIds() *[]*string
	// Experimental.
	SetOrganizationalUnitIds(val *[]*string)
	// Experimental.
	OrganizationalUnitIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetAccountFilterType()
	// Experimental.
	ResetAccounts()
	// Experimental.
	ResetAccountsUrl()
	// Experimental.
	ResetOrganizationalUnitIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference
type jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) AccountFilterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) AccountFilterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountFilterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) AccountsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) AccountsUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountsUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) InternalValue() *AwsCloudformationStackInstances_DeploymentTargetsProperty {
	var returns *AwsCloudformationStackInstances_DeploymentTargetsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) OrganizationalUnitIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) OrganizationalUnitIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudformation.AwsCloudformationStackInstances.DeploymentTargetsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference_Override(a AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudformation.AwsCloudformationStackInstances.DeploymentTargetsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetAccountFilterType(val *string) {
	if err := j.validateSetAccountFilterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountFilterType",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetAccountsUrl(val *string) {
	if err := j.validateSetAccountsUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountsUrl",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetInternalValue(val *AwsCloudformationStackInstances_DeploymentTargetsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetOrganizationalUnitIds(val *[]*string) {
	if err := j.validateSetOrganizationalUnitIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitIds",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ResetAccountFilterType() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountFilterType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ResetAccountsUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountsUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ResetOrganizationalUnitIds() {
	_jsii_.InvokeVoid(
		a,
		"resetOrganizationalUnitIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudformationStackInstances_DeploymentTargetsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

