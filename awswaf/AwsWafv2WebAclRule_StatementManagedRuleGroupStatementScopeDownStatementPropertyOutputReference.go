package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AsnMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList
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
	GeoMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList
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
	XssMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList
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
	PutRegexMatchStatement(value interface{})
	// Experimental.
	PutRegexPatternSetReferenceStatement(value interface{})
	// Experimental.
	PutSizeConstraintStatement(value interface{})
	// Experimental.
	PutSqliMatchStatement(value interface{})
	// Experimental.
	PutXssMatchStatement(value interface{})
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
	ResetRegexMatchStatement()
	// Experimental.
	ResetRegexPatternSetReferenceStatement()
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

// The jsii proxy struct for AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference
type jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) AsnMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ByteMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GeoMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) LabelMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SqliMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) XssMatchStatement() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference_Override(a AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := a.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := a.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := a.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := a.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := a.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := a.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := a.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := a.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := a.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := a.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

