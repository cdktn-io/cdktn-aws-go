package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsguardduty/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsguardduty/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFilter_CriterionPropertyOutputReference interface {
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
	EqualTo() *[]*string
	// Experimental.
	SetEqualTo(val *[]*string)
	// Experimental.
	EqualToInput() *[]*string
	// Experimental.
	Field() *string
	// Experimental.
	SetField(val *string)
	// Experimental.
	FieldInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GreaterThan() *string
	// Experimental.
	SetGreaterThan(val *string)
	// Experimental.
	GreaterThanInput() *string
	// Experimental.
	GreaterThanOrEqual() *string
	// Experimental.
	SetGreaterThanOrEqual(val *string)
	// Experimental.
	GreaterThanOrEqualInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LessThan() *string
	// Experimental.
	SetLessThan(val *string)
	// Experimental.
	LessThanInput() *string
	// Experimental.
	LessThanOrEqual() *string
	// Experimental.
	SetLessThanOrEqual(val *string)
	// Experimental.
	LessThanOrEqualInput() *string
	// Experimental.
	Matches() *[]*string
	// Experimental.
	SetMatches(val *[]*string)
	// Experimental.
	MatchesInput() *[]*string
	// Experimental.
	NotEquals() *[]*string
	// Experimental.
	SetNotEquals(val *[]*string)
	// Experimental.
	NotEqualsInput() *[]*string
	// Experimental.
	NotMatches() *[]*string
	// Experimental.
	SetNotMatches(val *[]*string)
	// Experimental.
	NotMatchesInput() *[]*string
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
	ResetEqualTo()
	// Experimental.
	ResetGreaterThan()
	// Experimental.
	ResetGreaterThanOrEqual()
	// Experimental.
	ResetLessThan()
	// Experimental.
	ResetLessThanOrEqual()
	// Experimental.
	ResetMatches()
	// Experimental.
	ResetNotEquals()
	// Experimental.
	ResetNotMatches()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFilter_CriterionPropertyOutputReference
type jsiiProxy_TfFilter_CriterionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GreaterThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GreaterThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GreaterThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) LessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) LessThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) LessThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) LessThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) Matches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) MatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) NotMatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) NotMatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFilter_CriterionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfFilter_CriterionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFilter_CriterionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFilter_CriterionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-guardduty.TfFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFilter_CriterionPropertyOutputReference_Override(t TfFilter_CriterionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-guardduty.TfFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetGreaterThan(val *string) {
	if err := j.validateSetGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThan",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetGreaterThanOrEqual(val *string) {
	if err := j.validateSetGreaterThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetLessThan(val *string) {
	if err := j.validateSetLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThan",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetLessThanOrEqual(val *string) {
	if err := j.validateSetLessThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetMatches(val *[]*string) {
	if err := j.validateSetMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matches",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetNotMatches(val *[]*string) {
	if err := j.validateSetNotMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMatches",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFilter_CriterionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		t,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetGreaterThan() {
	_jsii_.InvokeVoid(
		t,
		"resetGreaterThan",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetGreaterThanOrEqual() {
	_jsii_.InvokeVoid(
		t,
		"resetGreaterThanOrEqual",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetLessThan() {
	_jsii_.InvokeVoid(
		t,
		"resetLessThan",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetLessThanOrEqual() {
	_jsii_.InvokeVoid(
		t,
		"resetLessThanOrEqual",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetMatches() {
	_jsii_.InvokeVoid(
		t,
		"resetMatches",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		t,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ResetNotMatches() {
	_jsii_.InvokeVoid(
		t,
		"resetNotMatches",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFilter_CriterionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

