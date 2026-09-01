package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMacie2FindingsFilter_CriterionPropertyOutputReference interface {
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
	Eq() *[]*string
	// Experimental.
	SetEq(val *[]*string)
	// Experimental.
	EqExactMatch() *[]*string
	// Experimental.
	SetEqExactMatch(val *[]*string)
	// Experimental.
	EqExactMatchInput() *[]*string
	// Experimental.
	EqInput() *[]*string
	// Experimental.
	Field() *string
	// Experimental.
	SetField(val *string)
	// Experimental.
	FieldInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Gt() *string
	// Experimental.
	SetGt(val *string)
	// Experimental.
	Gte() *string
	// Experimental.
	SetGte(val *string)
	// Experimental.
	GteInput() *string
	// Experimental.
	GtInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Lt() *string
	// Experimental.
	SetLt(val *string)
	// Experimental.
	Lte() *string
	// Experimental.
	SetLte(val *string)
	// Experimental.
	LteInput() *string
	// Experimental.
	LtInput() *string
	// Experimental.
	Neq() *[]*string
	// Experimental.
	SetNeq(val *[]*string)
	// Experimental.
	NeqInput() *[]*string
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
	ResetEq()
	// Experimental.
	ResetEqExactMatch()
	// Experimental.
	ResetGt()
	// Experimental.
	ResetGte()
	// Experimental.
	ResetLt()
	// Experimental.
	ResetLte()
	// Experimental.
	ResetNeq()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMacie2FindingsFilter_CriterionPropertyOutputReference
type jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Eq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) EqExactMatch() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqExactMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) EqExactMatchInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqExactMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) EqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"eqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Field() *string {
	var returns *string
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) FieldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Gt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Gte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Lt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Lte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) LteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) LtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Neq() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) NeqInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"neqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMacie2FindingsFilter_CriterionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMacie2FindingsFilter_CriterionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMacie2FindingsFilter_CriterionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2FindingsFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMacie2FindingsFilter_CriterionPropertyOutputReference_Override(a AwsMacie2FindingsFilter_CriterionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.AwsMacie2FindingsFilter.CriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetEq(val *[]*string) {
	if err := j.validateSetEqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eq",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetEqExactMatch(val *[]*string) {
	if err := j.validateSetEqExactMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eqExactMatch",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetField(val *string) {
	if err := j.validateSetFieldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"field",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetGt(val *string) {
	if err := j.validateSetGtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gt",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetGte(val *string) {
	if err := j.validateSetGteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gte",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetLt(val *string) {
	if err := j.validateSetLtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lt",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetLte(val *string) {
	if err := j.validateSetLteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lte",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetNeq(val *[]*string) {
	if err := j.validateSetNeqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"neq",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetEq() {
	_jsii_.InvokeVoid(
		a,
		"resetEq",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetEqExactMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetEqExactMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetGt() {
	_jsii_.InvokeVoid(
		a,
		"resetGt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetGte() {
	_jsii_.InvokeVoid(
		a,
		"resetGte",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetLt() {
	_jsii_.InvokeVoid(
		a,
		"resetLt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetLte() {
	_jsii_.InvokeVoid(
		a,
		"resetLte",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ResetNeq() {
	_jsii_.InvokeVoid(
		a,
		"resetNeq",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMacie2FindingsFilter_CriterionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

