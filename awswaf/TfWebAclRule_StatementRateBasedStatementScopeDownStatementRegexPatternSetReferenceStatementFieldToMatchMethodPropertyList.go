package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList interface {
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
	Get(index *float64) TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList
type jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList_Override(t TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) Get(index *float64) TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

