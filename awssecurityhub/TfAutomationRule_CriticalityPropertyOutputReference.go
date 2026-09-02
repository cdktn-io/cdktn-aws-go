package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAutomationRule_CriticalityPropertyOutputReference interface {
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
	Eq() *float64
	// Experimental.
	SetEq(val *float64)
	// Experimental.
	EqInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	Gt() *float64
	// Experimental.
	SetGt(val *float64)
	// Experimental.
	Gte() *float64
	// Experimental.
	SetGte(val *float64)
	// Experimental.
	GteInput() *float64
	// Experimental.
	GtInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Lt() *float64
	// Experimental.
	SetLt(val *float64)
	// Experimental.
	Lte() *float64
	// Experimental.
	SetLte(val *float64)
	// Experimental.
	LteInput() *float64
	// Experimental.
	LtInput() *float64
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
	ResetGt()
	// Experimental.
	ResetGte()
	// Experimental.
	ResetLt()
	// Experimental.
	ResetLte()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAutomationRule_CriticalityPropertyOutputReference
type jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Eq() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) EqInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Gt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Gte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Lt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Lte() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) LteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) LtInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAutomationRule_CriticalityPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAutomationRule_CriticalityPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAutomationRule_CriticalityPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.CriticalityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAutomationRule_CriticalityPropertyOutputReference_Override(t TfAutomationRule_CriticalityPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.CriticalityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetEq(val *float64) {
	if err := j.validateSetEqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eq",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetGt(val *float64) {
	if err := j.validateSetGtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gt",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetGte(val *float64) {
	if err := j.validateSetGteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gte",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetLt(val *float64) {
	if err := j.validateSetLtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lt",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetLte(val *float64) {
	if err := j.validateSetLteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lte",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ResetEq() {
	_jsii_.InvokeVoid(
		t,
		"resetEq",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ResetGt() {
	_jsii_.InvokeVoid(
		t,
		"resetGt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ResetGte() {
	_jsii_.InvokeVoid(
		t,
		"resetGte",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ResetLt() {
	_jsii_.InvokeVoid(
		t,
		"resetLt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ResetLte() {
	_jsii_.InvokeVoid(
		t,
		"resetLte",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAutomationRule_CriticalityPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

