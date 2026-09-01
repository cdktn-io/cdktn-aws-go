package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference interface {
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
	FieldToMatch() AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchPropertyList
	// Experimental.
	FieldToMatchInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RegexString() *string
	// Experimental.
	SetRegexString(val *string)
	// Experimental.
	RegexStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextTransformation() AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationPropertyList
	// Experimental.
	TextTransformationInput() interface{}
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
	PutFieldToMatch(value interface{})
	// Experimental.
	PutTextTransformation(value interface{})
	// Experimental.
	ResetFieldToMatch()
	// Experimental.
	ResetTextTransformation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference
type jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) FieldToMatch() AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchPropertyList {
	var returns AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchPropertyList
	_jsii_.Get(
		j,
		"fieldToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) FieldToMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) RegexString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) RegexStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) TextTransformation() AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationPropertyList {
	var returns AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementTextTransformationPropertyList
	_jsii_.Get(
		j,
		"textTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) TextTransformationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textTransformationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference_Override(a AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetRegexString(val *string) {
	if err := j.validateSetRegexStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexString",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) PutFieldToMatch(value interface{}) {
	if err := a.validatePutFieldToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFieldToMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) PutTextTransformation(value interface{}) {
	if err := a.validatePutTextTransformationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTextTransformation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ResetFieldToMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetFieldToMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ResetTextTransformation() {
	_jsii_.InvokeVoid(
		a,
		"resetTextTransformation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

