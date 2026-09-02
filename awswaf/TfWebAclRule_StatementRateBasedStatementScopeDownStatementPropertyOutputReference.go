package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AsnMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList
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
	GeoMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList
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
	XssMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) AsnMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ByteMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GeoMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) LabelMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SqliMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) XssMatchStatement() TfWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference_Override(t TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := t.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := t.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := t.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := t.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := t.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := t.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := t.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := t.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := t.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := t.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

