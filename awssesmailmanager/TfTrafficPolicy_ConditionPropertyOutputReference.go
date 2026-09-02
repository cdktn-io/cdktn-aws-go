package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrafficPolicy_ConditionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BooleanExpression() TfTrafficPolicy_BooleanExpressionPropertyList
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
	IpExpression() TfTrafficPolicy_IpExpressionPropertyList
	// Experimental.
	IpExpressionInput() interface{}
	// Experimental.
	Ipv6Expression() TfTrafficPolicy_Ipv6ExpressionPropertyList
	// Experimental.
	Ipv6ExpressionInput() interface{}
	// Experimental.
	StringExpression() TfTrafficPolicy_StringExpressionPropertyList
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
	TlsExpression() TfTrafficPolicy_TlsExpressionPropertyList
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

// The jsii proxy struct for TfTrafficPolicy_ConditionPropertyOutputReference
type jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) BooleanExpression() TfTrafficPolicy_BooleanExpressionPropertyList {
	var returns TfTrafficPolicy_BooleanExpressionPropertyList
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) IpExpression() TfTrafficPolicy_IpExpressionPropertyList {
	var returns TfTrafficPolicy_IpExpressionPropertyList
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) Ipv6Expression() TfTrafficPolicy_Ipv6ExpressionPropertyList {
	var returns TfTrafficPolicy_Ipv6ExpressionPropertyList
	_jsii_.Get(
		j,
		"ipv6Expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) Ipv6ExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6ExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) StringExpression() TfTrafficPolicy_StringExpressionPropertyList {
	var returns TfTrafficPolicy_StringExpressionPropertyList
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) TlsExpression() TfTrafficPolicy_TlsExpressionPropertyList {
	var returns TfTrafficPolicy_TlsExpressionPropertyList
	_jsii_.Get(
		j,
		"tlsExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) TlsExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsExpressionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTrafficPolicy_ConditionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTrafficPolicy_ConditionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTrafficPolicy_ConditionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfTrafficPolicy.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTrafficPolicy_ConditionPropertyOutputReference_Override(t TfTrafficPolicy_ConditionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfTrafficPolicy.ConditionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) PutBooleanExpression(value interface{}) {
	if err := t.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) PutIpExpression(value interface{}) {
	if err := t.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) PutIpv6Expression(value interface{}) {
	if err := t.validatePutIpv6ExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpv6Expression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) PutStringExpression(value interface{}) {
	if err := t.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) PutTlsExpression(value interface{}) {
	if err := t.validatePutTlsExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTlsExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ResetIpv6Expression() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6Expression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ResetTlsExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTrafficPolicy_ConditionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

