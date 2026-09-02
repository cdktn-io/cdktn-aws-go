package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList interface {
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
	Get(index *float64) TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList
type jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList_Override(t TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) Get(index *float64) TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyMatchPatternAllPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

