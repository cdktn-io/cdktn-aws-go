package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	All() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternAllPropertyList
	// Experimental.
	AllInput() interface{}
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
	ExcludedCookies() *[]*string
	// Experimental.
	SetExcludedCookies(val *[]*string)
	// Experimental.
	ExcludedCookiesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludedCookies() *[]*string
	// Experimental.
	SetIncludedCookies(val *[]*string)
	// Experimental.
	IncludedCookiesInput() *[]*string
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
	PutAll(value interface{})
	// Experimental.
	ResetAll()
	// Experimental.
	ResetExcludedCookies()
	// Experimental.
	ResetIncludedCookies()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) All() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternAllPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternAllPropertyList
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) AllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ExcludedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ExcludedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) IncludedCookies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) IncludedCookiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference_Override(t TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetExcludedCookies(val *[]*string) {
	if err := j.validateSetExcludedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedCookies",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetIncludedCookies(val *[]*string) {
	if err := j.validateSetIncludedCookiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedCookies",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) PutAll(value interface{}) {
	if err := t.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAll",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		t,
		"resetAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ResetExcludedCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludedCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ResetIncludedCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludedCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementFieldToMatchCookiesMatchPatternPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

