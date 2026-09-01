package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference interface {
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
	InsertHeader() AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingInsertHeaderPropertyList
	// Experimental.
	InsertHeaderInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutInsertHeader(value interface{})
	// Experimental.
	ResetInsertHeader()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference
type jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) InsertHeader() AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingInsertHeaderPropertyList {
	var returns AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingInsertHeaderPropertyList
	_jsii_.Get(
		j,
		"insertHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) InsertHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insertHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference_Override(a AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) PutInsertHeader(value interface{}) {
	if err := a.validatePutInsertHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInsertHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) ResetInsertHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetInsertHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

