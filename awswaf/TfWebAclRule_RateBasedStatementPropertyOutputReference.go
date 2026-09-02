package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_RateBasedStatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AggregateKeyType() *string
	// Experimental.
	SetAggregateKeyType(val *string)
	// Experimental.
	AggregateKeyTypeInput() *string
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
	CustomKeys() TfWebAclRule_CustomKeysPropertyList
	// Experimental.
	CustomKeysInput() interface{}
	// Experimental.
	EvaluationWindowSec() *float64
	// Experimental.
	SetEvaluationWindowSec(val *float64)
	// Experimental.
	EvaluationWindowSecInput() *float64
	// Experimental.
	ForwardedIpConfig() TfWebAclRule_StatementRateBasedStatementForwardedIpConfigPropertyList
	// Experimental.
	ForwardedIpConfigInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Limit() *float64
	// Experimental.
	SetLimit(val *float64)
	// Experimental.
	LimitInput() *float64
	// Experimental.
	ScopeDownStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyList
	// Experimental.
	ScopeDownStatementInput() interface{}
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
	PutCustomKeys(value interface{})
	// Experimental.
	PutForwardedIpConfig(value interface{})
	// Experimental.
	PutScopeDownStatement(value interface{})
	// Experimental.
	ResetCustomKeys()
	// Experimental.
	ResetEvaluationWindowSec()
	// Experimental.
	ResetForwardedIpConfig()
	// Experimental.
	ResetScopeDownStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_RateBasedStatementPropertyOutputReference
type jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) AggregateKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregateKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) AggregateKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregateKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) CustomKeys() TfWebAclRule_CustomKeysPropertyList {
	var returns TfWebAclRule_CustomKeysPropertyList
	_jsii_.Get(
		j,
		"customKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) CustomKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) EvaluationWindowSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationWindowSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) EvaluationWindowSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationWindowSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ForwardedIpConfig() TfWebAclRule_StatementRateBasedStatementForwardedIpConfigPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementForwardedIpConfigPropertyList
	_jsii_.Get(
		j,
		"forwardedIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ForwardedIpConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardedIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) LimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ScopeDownStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyList
	_jsii_.Get(
		j,
		"scopeDownStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ScopeDownStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeDownStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_RateBasedStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_RateBasedStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_RateBasedStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.RateBasedStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_RateBasedStatementPropertyOutputReference_Override(t TfWebAclRule_RateBasedStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.RateBasedStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetAggregateKeyType(val *string) {
	if err := j.validateSetAggregateKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregateKeyType",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetEvaluationWindowSec(val *float64) {
	if err := j.validateSetEvaluationWindowSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationWindowSec",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetLimit(val *float64) {
	if err := j.validateSetLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limit",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) PutCustomKeys(value interface{}) {
	if err := t.validatePutCustomKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomKeys",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) PutForwardedIpConfig(value interface{}) {
	if err := t.validatePutForwardedIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForwardedIpConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) PutScopeDownStatement(value interface{}) {
	if err := t.validatePutScopeDownStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScopeDownStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ResetCustomKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ResetEvaluationWindowSec() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluationWindowSec",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ResetForwardedIpConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetForwardedIpConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ResetScopeDownStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetScopeDownStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_RateBasedStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

