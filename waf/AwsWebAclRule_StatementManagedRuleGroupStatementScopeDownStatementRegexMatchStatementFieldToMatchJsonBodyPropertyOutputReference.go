package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	InvalidFallbackBehavior() *string
	// Experimental.
	SetInvalidFallbackBehavior(val *string)
	// Experimental.
	InvalidFallbackBehaviorInput() *string
	// Experimental.
	MatchPattern() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternPropertyList
	// Experimental.
	MatchPatternInput() interface{}
	// Experimental.
	MatchScope() *string
	// Experimental.
	SetMatchScope(val *string)
	// Experimental.
	MatchScopeInput() *string
	// Experimental.
	OversizeHandling() *string
	// Experimental.
	SetOversizeHandling(val *string)
	// Experimental.
	OversizeHandlingInput() *string
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
	PutMatchPattern(value interface{})
	// Experimental.
	ResetInvalidFallbackBehavior()
	// Experimental.
	ResetMatchPattern()
	// Experimental.
	ResetOversizeHandling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference
type jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) InvalidFallbackBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invalidFallbackBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) InvalidFallbackBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invalidFallbackBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) MatchPattern() AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternPropertyList {
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternPropertyList
	_jsii_.Get(
		j,
		"matchPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) MatchPatternInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) MatchScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) MatchScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) OversizeHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oversizeHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) OversizeHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oversizeHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference_Override(a AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetInvalidFallbackBehavior(val *string) {
	if err := j.validateSetInvalidFallbackBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invalidFallbackBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetMatchScope(val *string) {
	if err := j.validateSetMatchScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchScope",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetOversizeHandling(val *string) {
	if err := j.validateSetOversizeHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oversizeHandling",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) PutMatchPattern(value interface{}) {
	if err := a.validatePutMatchPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMatchPattern",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ResetInvalidFallbackBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInvalidFallbackBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ResetMatchPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetMatchPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ResetOversizeHandling() {
	_jsii_.InvokeVoid(
		a,
		"resetOversizeHandling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchJsonBodyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

