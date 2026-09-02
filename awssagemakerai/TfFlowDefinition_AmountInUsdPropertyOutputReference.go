package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlowDefinition_AmountInUsdPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Cents() *float64
	// Experimental.
	SetCents(val *float64)
	// Experimental.
	CentsInput() *float64
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
	Dollars() *float64
	// Experimental.
	SetDollars(val *float64)
	// Experimental.
	DollarsInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFlowDefinition_AmountInUsdProperty
	// Experimental.
	SetInternalValue(val *TfFlowDefinition_AmountInUsdProperty)
	// Experimental.
	TenthFractionsOfACent() *float64
	// Experimental.
	SetTenthFractionsOfACent(val *float64)
	// Experimental.
	TenthFractionsOfACentInput() *float64
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
	ResetCents()
	// Experimental.
	ResetDollars()
	// Experimental.
	ResetTenthFractionsOfACent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlowDefinition_AmountInUsdPropertyOutputReference
type jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) Cents() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) CentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"centsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) Dollars() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dollars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) DollarsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dollarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) InternalValue() *TfFlowDefinition_AmountInUsdProperty {
	var returns *TfFlowDefinition_AmountInUsdProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) TenthFractionsOfACent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tenthFractionsOfACent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) TenthFractionsOfACentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tenthFractionsOfACentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlowDefinition_AmountInUsdPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlowDefinition_AmountInUsdPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlowDefinition_AmountInUsdPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFlowDefinition.AmountInUsdPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlowDefinition_AmountInUsdPropertyOutputReference_Override(t TfFlowDefinition_AmountInUsdPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFlowDefinition.AmountInUsdPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetCents(val *float64) {
	if err := j.validateSetCentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cents",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetDollars(val *float64) {
	if err := j.validateSetDollarsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dollars",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetInternalValue(val *TfFlowDefinition_AmountInUsdProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetTenthFractionsOfACent(val *float64) {
	if err := j.validateSetTenthFractionsOfACentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenthFractionsOfACent",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ResetCents() {
	_jsii_.InvokeVoid(
		t,
		"resetCents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ResetDollars() {
	_jsii_.InvokeVoid(
		t,
		"resetDollars",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ResetTenthFractionsOfACent() {
	_jsii_.InvokeVoid(
		t,
		"resetTenthFractionsOfACent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlowDefinition_AmountInUsdPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

