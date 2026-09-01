package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsguardduty/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsguardduty/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGuarddutyFilter_CriterionPropertyOutputReference interface {
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

// The jsii proxy struct for AwsGuarddutyFilter_CriterionPropertyOutputReference
type jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) EqualTo() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) EqualToInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"equalToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GreaterThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GreaterThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GreaterThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GreaterThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"greaterThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) LessThan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) LessThanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) LessThanOrEqual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) LessThanOrEqualInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lessThanOrEqualInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) Matches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) MatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) NotEquals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) NotEqualsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) NotMatches() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) NotMatchesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notMatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGuarddutyFilter_CriterionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsGuarddutyFilter_CriterionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGuarddutyFilter_CriterionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-guardduty.AwsGuarddutyFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGuarddutyFilter_CriterionPropertyOutputReference_Override(a AwsGuarddutyFilter_CriterionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-guardduty.AwsGuarddutyFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetEqualTo(val *[]*string) {
	if err := j.validateSetEqualToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"equalTo",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetGreaterThan(val *string) {
	if err := j.validateSetGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThan",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetGreaterThanOrEqual(val *string) {
	if err := j.validateSetGreaterThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"greaterThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetLessThan(val *string) {
	if err := j.validateSetLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThan",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetLessThanOrEqual(val *string) {
	if err := j.validateSetLessThanOrEqualParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lessThanOrEqual",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetMatches(val *[]*string) {
	if err := j.validateSetMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matches",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetNotEquals(val *[]*string) {
	if err := j.validateSetNotEqualsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notEquals",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetNotMatches(val *[]*string) {
	if err := j.validateSetNotMatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notMatches",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetEqualTo() {
	_jsii_.InvokeVoid(
		a,
		"resetEqualTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetGreaterThan() {
	_jsii_.InvokeVoid(
		a,
		"resetGreaterThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetGreaterThanOrEqual() {
	_jsii_.InvokeVoid(
		a,
		"resetGreaterThanOrEqual",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetLessThan() {
	_jsii_.InvokeVoid(
		a,
		"resetLessThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetLessThanOrEqual() {
	_jsii_.InvokeVoid(
		a,
		"resetLessThanOrEqual",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetMatches() {
	_jsii_.InvokeVoid(
		a,
		"resetMatches",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetNotEquals() {
	_jsii_.InvokeVoid(
		a,
		"resetNotEquals",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ResetNotMatches() {
	_jsii_.InvokeVoid(
		a,
		"resetNotMatches",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGuarddutyFilter_CriterionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

