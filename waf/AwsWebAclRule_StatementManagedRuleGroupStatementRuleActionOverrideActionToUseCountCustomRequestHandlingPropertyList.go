package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList interface {
	cdktn.ComplexList
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList
type jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList_Override(a AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) Get(index *float64) AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementManagedRuleGroupStatementRuleActionOverrideActionToUseCountCustomRequestHandlingPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

