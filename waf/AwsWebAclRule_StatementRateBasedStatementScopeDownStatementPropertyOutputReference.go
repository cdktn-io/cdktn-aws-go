package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AsnMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList
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
	GeoMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList
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
	XssMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList
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

// The jsii proxy struct for AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference
type jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) AsnMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ByteMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GeoMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) LabelMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SqliMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) XssMatchStatement() AwsWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList {
	var returns AwsWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementRateBasedStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference_Override(a AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRule.StatementRateBasedStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := a.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := a.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := a.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := a.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := a.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := a.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := a.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := a.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := a.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := a.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

