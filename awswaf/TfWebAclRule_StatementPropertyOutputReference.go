package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AndStatement() TfWebAclRule_AndStatementPropertyList
	// Experimental.
	AndStatementInput() interface{}
	// Experimental.
	AsnMatchStatement() TfWebAclRule_StatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() TfWebAclRule_StatementByteMatchStatementPropertyList
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
	GeoMatchStatement() TfWebAclRule_StatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() TfWebAclRule_StatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() TfWebAclRule_StatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	ManagedRuleGroupStatement() TfWebAclRule_ManagedRuleGroupStatementPropertyList
	// Experimental.
	ManagedRuleGroupStatementInput() interface{}
	// Experimental.
	NotStatement() TfWebAclRule_NotStatementPropertyList
	// Experimental.
	NotStatementInput() interface{}
	// Experimental.
	OrStatement() TfWebAclRule_OrStatementPropertyList
	// Experimental.
	OrStatementInput() interface{}
	// Experimental.
	RateBasedStatement() TfWebAclRule_RateBasedStatementPropertyList
	// Experimental.
	RateBasedStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() TfWebAclRule_StatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() TfWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	RuleGroupReferenceStatement() TfWebAclRule_RuleGroupReferenceStatementPropertyList
	// Experimental.
	RuleGroupReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() TfWebAclRule_StatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() TfWebAclRule_StatementSqliMatchStatementPropertyList
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
	XssMatchStatement() TfWebAclRule_StatementXssMatchStatementPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) AndStatement() TfWebAclRule_AndStatementPropertyList {
	var returns TfWebAclRule_AndStatementPropertyList
	_jsii_.Get(
		j,
		"andStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) AndStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) AsnMatchStatement() TfWebAclRule_StatementAsnMatchStatementPropertyList {
	var returns TfWebAclRule_StatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ByteMatchStatement() TfWebAclRule_StatementByteMatchStatementPropertyList {
	var returns TfWebAclRule_StatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GeoMatchStatement() TfWebAclRule_StatementGeoMatchStatementPropertyList {
	var returns TfWebAclRule_StatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) IpSetReferenceStatement() TfWebAclRule_StatementIpSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) LabelMatchStatement() TfWebAclRule_StatementLabelMatchStatementPropertyList {
	var returns TfWebAclRule_StatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ManagedRuleGroupStatement() TfWebAclRule_ManagedRuleGroupStatementPropertyList {
	var returns TfWebAclRule_ManagedRuleGroupStatementPropertyList
	_jsii_.Get(
		j,
		"managedRuleGroupStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ManagedRuleGroupStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedRuleGroupStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) NotStatement() TfWebAclRule_NotStatementPropertyList {
	var returns TfWebAclRule_NotStatementPropertyList
	_jsii_.Get(
		j,
		"notStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) NotStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) OrStatement() TfWebAclRule_OrStatementPropertyList {
	var returns TfWebAclRule_OrStatementPropertyList
	_jsii_.Get(
		j,
		"orStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) OrStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RateBasedStatement() TfWebAclRule_RateBasedStatementPropertyList {
	var returns TfWebAclRule_RateBasedStatementPropertyList
	_jsii_.Get(
		j,
		"rateBasedStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RateBasedStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rateBasedStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RegexMatchStatement() TfWebAclRule_StatementRegexMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RegexPatternSetReferenceStatement() TfWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RuleGroupReferenceStatement() TfWebAclRule_RuleGroupReferenceStatementPropertyList {
	var returns TfWebAclRule_RuleGroupReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) RuleGroupReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleGroupReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) SizeConstraintStatement() TfWebAclRule_StatementSizeConstraintStatementPropertyList {
	var returns TfWebAclRule_StatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) SqliMatchStatement() TfWebAclRule_StatementSqliMatchStatementPropertyList {
	var returns TfWebAclRule_StatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) XssMatchStatement() TfWebAclRule_StatementXssMatchStatementPropertyList {
	var returns TfWebAclRule_StatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementPropertyOutputReference_Override(t TfWebAclRule_StatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutAndStatement(value interface{}) {
	if err := t.validatePutAndStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAndStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := t.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := t.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := t.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := t.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := t.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutManagedRuleGroupStatement(value interface{}) {
	if err := t.validatePutManagedRuleGroupStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedRuleGroupStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutNotStatement(value interface{}) {
	if err := t.validatePutNotStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNotStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutOrStatement(value interface{}) {
	if err := t.validatePutOrStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOrStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutRateBasedStatement(value interface{}) {
	if err := t.validatePutRateBasedStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRateBasedStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := t.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := t.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutRuleGroupReferenceStatement(value interface{}) {
	if err := t.validatePutRuleGroupReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleGroupReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := t.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := t.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := t.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetAndStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetAndStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetManagedRuleGroupStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedRuleGroupStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetNotStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetNotStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetOrStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetOrStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetRateBasedStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRateBasedStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetRuleGroupReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRuleGroupReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

