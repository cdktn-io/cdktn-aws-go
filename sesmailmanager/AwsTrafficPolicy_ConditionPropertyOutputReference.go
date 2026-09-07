package sesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTrafficPolicy_ConditionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BooleanExpression() AwsTrafficPolicy_BooleanExpressionPropertyList
	// Experimental.
	BooleanExpressionInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpExpression() AwsTrafficPolicy_IpExpressionPropertyList
	// Experimental.
	IpExpressionInput() interface{}
	// Experimental.
	Ipv6Expression() AwsTrafficPolicy_Ipv6ExpressionPropertyList
	// Experimental.
	Ipv6ExpressionInput() interface{}
	// Experimental.
	StringExpression() AwsTrafficPolicy_StringExpressionPropertyList
	// Experimental.
	StringExpressionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TlsExpression() AwsTrafficPolicy_TlsExpressionPropertyList
	// Experimental.
	TlsExpressionInput() interface{}
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
	PutBooleanExpression(value interface{})
	// Experimental.
	PutIpExpression(value interface{})
	// Experimental.
	PutIpv6Expression(value interface{})
	// Experimental.
	PutStringExpression(value interface{})
	// Experimental.
	PutTlsExpression(value interface{})
	// Experimental.
	ResetBooleanExpression()
	// Experimental.
	ResetIpExpression()
	// Experimental.
	ResetIpv6Expression()
	// Experimental.
	ResetStringExpression()
	// Experimental.
	ResetTlsExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTrafficPolicy_ConditionPropertyOutputReference
type jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) BooleanExpression() AwsTrafficPolicy_BooleanExpressionPropertyList {
	var returns AwsTrafficPolicy_BooleanExpressionPropertyList
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) IpExpression() AwsTrafficPolicy_IpExpressionPropertyList {
	var returns AwsTrafficPolicy_IpExpressionPropertyList
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) Ipv6Expression() AwsTrafficPolicy_Ipv6ExpressionPropertyList {
	var returns AwsTrafficPolicy_Ipv6ExpressionPropertyList
	_jsii_.Get(
		j,
		"ipv6Expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) Ipv6ExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6ExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) StringExpression() AwsTrafficPolicy_StringExpressionPropertyList {
	var returns AwsTrafficPolicy_StringExpressionPropertyList
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) TlsExpression() AwsTrafficPolicy_TlsExpressionPropertyList {
	var returns AwsTrafficPolicy_TlsExpressionPropertyList
	_jsii_.Get(
		j,
		"tlsExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) TlsExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsExpressionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTrafficPolicy_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTrafficPolicy_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTrafficPolicy_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsTrafficPolicy.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTrafficPolicy_ConditionPropertyOutputReference_Override(a AwsTrafficPolicy_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsTrafficPolicy.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) PutBooleanExpression(value interface{}) {
	if err := a.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) PutIpExpression(value interface{}) {
	if err := a.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) PutIpv6Expression(value interface{}) {
	if err := a.validatePutIpv6ExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpv6Expression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) PutStringExpression(value interface{}) {
	if err := a.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) PutTlsExpression(value interface{}) {
	if err := a.validatePutTlsExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTlsExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ResetIpv6Expression() {
	_jsii_.InvokeVoid(
		a,
		"resetIpv6Expression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ResetTlsExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTrafficPolicy_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

