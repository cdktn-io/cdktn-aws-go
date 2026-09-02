package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	All() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternAllPropertyList
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
	ExcludedHeaders() *[]*string
	// Experimental.
	SetExcludedHeaders(val *[]*string)
	// Experimental.
	ExcludedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludedHeaders() *[]*string
	// Experimental.
	SetIncludedHeaders(val *[]*string)
	// Experimental.
	IncludedHeadersInput() *[]*string
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
	ResetExcludedHeaders()
	// Experimental.
	ResetIncludedHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) All() TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternAllPropertyList {
	var returns TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternAllPropertyList
	_jsii_.Get(
		j,
		"all",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) AllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ExcludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ExcludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) IncludedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) IncludedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference_Override(t TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetExcludedHeaders(val *[]*string) {
	if err := j.validateSetExcludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedHeaders",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetIncludedHeaders(val *[]*string) {
	if err := j.validateSetIncludedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaders",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) PutAll(value interface{}) {
	if err := t.validatePutAllParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAll",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ResetAll() {
	_jsii_.InvokeVoid(
		t,
		"resetAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ResetExcludedHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetExcludedHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ResetIncludedHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludedHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersMatchPatternPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

