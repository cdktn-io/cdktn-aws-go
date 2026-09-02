package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AsnMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList
	// Experimental.
	AsnMatchStatementInput() interface{}
	// Experimental.
	ByteMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList
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
	GeoMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList
	// Experimental.
	GeoMatchStatementInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpSetReferenceStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList
	// Experimental.
	IpSetReferenceStatementInput() interface{}
	// Experimental.
	LabelMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList
	// Experimental.
	LabelMatchStatementInput() interface{}
	// Experimental.
	RegexMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList
	// Experimental.
	RegexMatchStatementInput() interface{}
	// Experimental.
	RegexPatternSetReferenceStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	// Experimental.
	RegexPatternSetReferenceStatementInput() interface{}
	// Experimental.
	SizeConstraintStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList
	// Experimental.
	SizeConstraintStatementInput() interface{}
	// Experimental.
	SqliMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList
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
	XssMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList
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

// The jsii proxy struct for TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) AsnMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementAsnMatchStatementPropertyList
	_jsii_.Get(
		j,
		"asnMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) AsnMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asnMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ByteMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementByteMatchStatementPropertyList
	_jsii_.Get(
		j,
		"byteMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ByteMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"byteMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GeoMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementGeoMatchStatementPropertyList
	_jsii_.Get(
		j,
		"geoMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GeoMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"ipSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) IpSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) LabelMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementLabelMatchStatementPropertyList
	_jsii_.Get(
		j,
		"labelMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) LabelMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexMatchStatementPropertyList
	_jsii_.Get(
		j,
		"regexMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementRegexPatternSetReferenceStatementPropertyList
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) RegexPatternSetReferenceStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regexPatternSetReferenceStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSizeConstraintStatementPropertyList
	_jsii_.Get(
		j,
		"sizeConstraintStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SizeConstraintStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sizeConstraintStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SqliMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementPropertyList
	_jsii_.Get(
		j,
		"sqliMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) SqliMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqliMatchStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) XssMatchStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementXssMatchStatementPropertyList
	_jsii_.Get(
		j,
		"xssMatchStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) XssMatchStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xssMatchStatementInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference_Override(t TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutAsnMatchStatement(value interface{}) {
	if err := t.validatePutAsnMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAsnMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutByteMatchStatement(value interface{}) {
	if err := t.validatePutByteMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putByteMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutGeoMatchStatement(value interface{}) {
	if err := t.validatePutGeoMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGeoMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutIpSetReferenceStatement(value interface{}) {
	if err := t.validatePutIpSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutLabelMatchStatement(value interface{}) {
	if err := t.validatePutLabelMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLabelMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutRegexMatchStatement(value interface{}) {
	if err := t.validatePutRegexMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutRegexPatternSetReferenceStatement(value interface{}) {
	if err := t.validatePutRegexPatternSetReferenceStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRegexPatternSetReferenceStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutSizeConstraintStatement(value interface{}) {
	if err := t.validatePutSizeConstraintStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSizeConstraintStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutSqliMatchStatement(value interface{}) {
	if err := t.validatePutSqliMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqliMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) PutXssMatchStatement(value interface{}) {
	if err := t.validatePutXssMatchStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXssMatchStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetAsnMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetAsnMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetByteMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetByteMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetGeoMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetGeoMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetIpSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetIpSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetLabelMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetLabelMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetRegexMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetRegexPatternSetReferenceStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRegexPatternSetReferenceStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetSizeConstraintStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSizeConstraintStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetSqliMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetSqliMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ResetXssMatchStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetXssMatchStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

