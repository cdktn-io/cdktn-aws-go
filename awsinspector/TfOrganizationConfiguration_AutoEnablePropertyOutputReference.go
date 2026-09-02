package awsinspector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsinspector/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsinspector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOrganizationConfiguration_AutoEnablePropertyOutputReference interface {
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
	InternalValue() *TfOrganizationConfiguration_AutoEnableProperty
	// Experimental.
	SetInternalValue(val *TfOrganizationConfiguration_AutoEnableProperty)
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

// The jsii proxy struct for TfOrganizationConfiguration_AutoEnablePropertyOutputReference
type jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) CodeRepository() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) CodeRepositoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Ec2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Ec2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Ecr() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) EcrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ecrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) InternalValue() *TfOrganizationConfiguration_AutoEnableProperty {
	var returns *TfOrganizationConfiguration_AutoEnableProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Lambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaCode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOrganizationConfiguration_AutoEnablePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOrganizationConfiguration_AutoEnablePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOrganizationConfiguration_AutoEnablePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-inspector.TfOrganizationConfiguration.AutoEnablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOrganizationConfiguration_AutoEnablePropertyOutputReference_Override(t TfOrganizationConfiguration_AutoEnablePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-inspector.TfOrganizationConfiguration.AutoEnablePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetCodeRepository(val interface{}) {
	if err := j.validateSetCodeRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeRepository",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetEc2(val interface{}) {
	if err := j.validateSetEc2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ec2",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetEcr(val interface{}) {
	if err := j.validateSetEcrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecr",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetInternalValue(val *TfOrganizationConfiguration_AutoEnableProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetLambda(val interface{}) {
	if err := j.validateSetLambdaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambda",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetLambdaCode(val interface{}) {
	if err := j.validateSetLambdaCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaCode",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ResetLambdaCode() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOrganizationConfiguration_AutoEnablePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

