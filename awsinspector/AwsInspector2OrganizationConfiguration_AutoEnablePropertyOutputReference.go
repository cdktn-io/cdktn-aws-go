package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsinspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsinspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodeRepository() interface{}
	// Experimental.
	SetCodeRepository(val interface{})
	// Experimental.
	CodeRepositoryInput() interface{}
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
	Ec2() interface{}
	// Experimental.
	SetEc2(val interface{})
	// Experimental.
	Ec2Input() interface{}
	// Experimental.
	Ecr() interface{}
	// Experimental.
	SetEcr(val interface{})
	// Experimental.
	EcrInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsInspector2OrganizationConfiguration_AutoEnableProperty
	// Experimental.
	SetInternalValue(val *AwsInspector2OrganizationConfiguration_AutoEnableProperty)
	// Experimental.
	Lambda() interface{}
	// Experimental.
	SetLambda(val interface{})
	// Experimental.
	LambdaCode() interface{}
	// Experimental.
	SetLambdaCode(val interface{})
	// Experimental.
	LambdaCodeInput() interface{}
	// Experimental.
	LambdaInput() interface{}
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
	ResetCodeRepository()
	// Experimental.
	ResetLambda()
	// Experimental.
	ResetLambdaCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference
type jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) CodeRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Ec2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Ec2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Ecr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) EcrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) InternalValue() *AwsInspector2OrganizationConfiguration_AutoEnableProperty {
	var returns *AwsInspector2OrganizationConfiguration_AutoEnableProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Lambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaCode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsInspector2OrganizationConfiguration.AutoEnablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference_Override(a AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.AwsInspector2OrganizationConfiguration.AutoEnablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetCodeRepository(val interface{}) {
	if err := j.validateSetCodeRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeRepository",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetEc2(val interface{}) {
	if err := j.validateSetEc2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetEcr(val interface{}) {
	if err := j.validateSetEcrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecr",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetInternalValue(val *AwsInspector2OrganizationConfiguration_AutoEnableProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetLambda(val interface{}) {
	if err := j.validateSetLambdaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambda",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetLambdaCode(val interface{}) {
	if err := j.validateSetLambdaCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaCode",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ResetLambdaCode() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsInspector2OrganizationConfiguration_AutoEnablePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

