package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_StatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AndStatement() AwsWebAclRule_AndStatementPropertyList
	// Experimental.
	AndStatementInput() interface{}
	// Experimental.
	AsnMatchStatement() AwsWebAclRule_StatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() AwsWebAclRule_StatementByteMatchStatementPropertyList
	// Experimental.
	ByteMatchStatementInput() interface{}
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
	GeoMatchStatement() AwsWebAclRule_StatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() AwsWebAclRule_StatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() AwsWebAclRule_StatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	ManagedRuleGroupStatement() AwsWebAclRule_ManagedRuleGroupStatementPropertyList
	// Experimental.
	ManagedRuleGroupStatementInput() interface{}
	// Experimental.
	NotStatement() AwsWebAclRule_NotStatementPropertyList
	// Experimental.
	NotStatementInput() interface{}
	// Experimental.
	OrStatement() AwsWebAclRule_OrStatementPropertyList
	// Experimental.
	OrStatementInput() interface{}
	// Experimental.
	RateBasedStatement() AwsWebAclRule_RateBasedStatementPropertyList
	// Experimental.
	RateBasedStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() AwsWebAclRule_StatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() AwsWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	RuleGroupReferenceStatement() AwsWebAclRule_RuleGroupReferenceStatementPropertyList
	// Experimental.
	RuleGroupReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() AwsWebAclRule_StatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() AwsWebAclRule_StatementSqliMatchStatementPropertyList
	// Experimental.
	SqliMatchStatementInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	XssMatchStatement() AwsWebAclRule_StatementXssMatchStatementPropertyList
	// Experimental.
	XssMatchStatementInput() interface{}
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
	PutAndStatement(value interface{})
	// Experimental.
	PutAsnMatchStatement(value interface{})
	// Experimental.
	PutByteMatchStatement(value interface{})
	// Experimental.
	PutGeoMatchStatement(value interface{})
	// Experimental.
	PutIpSetReferenceStatement(value interface{})
	// Experimental.
	PutLabelMatchStatement(value interface{})
	// Experimental.
	PutManagedRuleGroupStatement(value interface{})
	// Experimental.
	PutNotStatement(value interface{})
	// Experimental.
	PutOrStatement(value interface{})
	// Experimental.
	PutRateBasedStatement(value interface{})
	// Experimental.
	PutRegexMatchStatement(value interface{})
	// Experimental.
	PutRegexPatternSetReferenceStatement(value interface{})
	// Experimental.
	PutRuleGroupReferenceStatement(value interface{})
	// Experimental.
	PutSizeConstraintStatement(value interface{})
	// Experimental.
	PutSqliMatchStatement(value interface{})
	// Experimental.
	PutXssMatchStatement(value interface{})
	// Experimental.
	ResetAndStatement()
	// Experimental.
	ResetAsnMatchStatement()
	// Experimental.
	ResetByteMatchStatement()
	// Experimental.
	ResetGeoMatchStatement()
	// Experimental.
	ResetIpSetReferenceStatement()
	// Experimental.
	ResetLabelMatchStatement()
	// Experimental.
	ResetManagedRuleGroupStatement()
	// Experimental.
	ResetNotStatement()
	// Experimental.
	ResetOrStatement()
	// Experimental.
	ResetRateBasedStatement()
	// Experimental.
	ResetRegexMatchStatement()
	// Experimental.
	ResetRegexPatternSetReferenceStatement()
	// Experimental.
	ResetRuleGroupReferenceStatement()
	// Experimental.
	ResetSizeConstraintStatement()
	// Experimental.
	ResetSqliMatchStatement()
	// Experimental.
	ResetXssMatchStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWebAclRule_StatementPropertyOutputReference
type jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) AndStatement() AwsWebAclRule_AndStatementPropertyList {
	var returns AwsWebAclRule_AndStatementPropertyList
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) AndStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) AsnMatchStatement() AwsWebAclRule_StatementAsnMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ByteMatchStatement() AwsWebAclRule_StatementByteMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GeoMatchStatement() AwsWebAclRule_StatementGeoMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) IpSetReferenceStatement() AwsWebAclRule_StatementIpSetReferenceStatementPropertyList {
	var returns AwsWebAclRule_StatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) LabelMatchStatement() AwsWebAclRule_StatementLabelMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ManagedRuleGroupStatement() AwsWebAclRule_ManagedRuleGroupStatementPropertyList {
	var returns AwsWebAclRule_ManagedRuleGroupStatementPropertyList
	_jsii_.Get(
		j,
		"managedRuleGroupStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ManagedRuleGroupStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedRuleGroupStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) NotStatement() AwsWebAclRule_NotStatementPropertyList {
	var returns AwsWebAclRule_NotStatementPropertyList
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) NotStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) OrStatement() AwsWebAclRule_OrStatementPropertyList {
	var returns AwsWebAclRule_OrStatementPropertyList
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) OrStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RateBasedStatement() AwsWebAclRule_RateBasedStatementPropertyList {
	var returns AwsWebAclRule_RateBasedStatementPropertyList
	_jsii_.Get(
		j,
		"rateBasedStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RateBasedStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateBasedStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RegexMatchStatement() AwsWebAclRule_StatementRegexMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RegexPatternSetReferenceStatement() AwsWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList {
	var returns AwsWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RuleGroupReferenceStatement() AwsWebAclRule_RuleGroupReferenceStatementPropertyList {
	var returns AwsWebAclRule_RuleGroupReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) RuleGroupReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) SizeConstraintStatement() AwsWebAclRule_StatementSizeConstraintStatementPropertyList {
	var returns AwsWebAclRule_StatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) SqliMatchStatement() AwsWebAclRule_StatementSqliMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) XssMatchStatement() AwsWebAclRule_StatementXssMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_StatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_StatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_StatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_StatementPropertyOutputReference_Override(a AwsWebAclRule_StatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutAndStatement(value interface{}) {
	if err := a.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := a.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := a.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := a.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := a.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := a.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutManagedRuleGroupStatement(value interface{}) {
	if err := a.validatePutManagedRuleGroupStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedRuleGroupStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutNotStatement(value interface{}) {
	if err := a.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutOrStatement(value interface{}) {
	if err := a.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutRateBasedStatement(value interface{}) {
	if err := a.validatePutRateBasedStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRateBasedStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := a.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := a.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutRuleGroupReferenceStatement(value interface{}) {
	if err := a.validatePutRuleGroupReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuleGroupReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := a.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := a.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := a.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetManagedRuleGroupStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedRuleGroupStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetRateBasedStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRateBasedStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetRuleGroupReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleGroupReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

