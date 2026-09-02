package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InputAction() *string
	// Experimental.
	SetInputAction(val *string)
	// Experimental.
	InputActionInput() *string
	// Experimental.
	InputEnabled() interface{}
	// Experimental.
	SetInputEnabled(val interface{})
	// Experimental.
	InputEnabledInput() interface{}
	// Experimental.
	InputModalities() *[]*string
	// Experimental.
	SetInputModalities(val *[]*string)
	// Experimental.
	InputModalitiesInput() *[]*string
	// Experimental.
	InputStrength() *string
	// Experimental.
	SetInputStrength(val *string)
	// Experimental.
	InputStrengthInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputAction() *string
	// Experimental.
	SetOutputAction(val *string)
	// Experimental.
	OutputActionInput() *string
	// Experimental.
	OutputEnabled() interface{}
	// Experimental.
	SetOutputEnabled(val interface{})
	// Experimental.
	OutputEnabledInput() interface{}
	// Experimental.
	OutputModalities() *[]*string
	// Experimental.
	SetOutputModalities(val *[]*string)
	// Experimental.
	OutputModalitiesInput() *[]*string
	// Experimental.
	OutputStrength() *string
	// Experimental.
	SetOutputStrength(val *string)
	// Experimental.
	OutputStrengthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	ResetInputAction()
	// Experimental.
	ResetInputEnabled()
	// Experimental.
	ResetInputModalities()
	// Experimental.
	ResetOutputAction()
	// Experimental.
	ResetOutputEnabled()
	// Experimental.
	ResetOutputModalities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference
type jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputModalities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputModalitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputModalitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) OutputStrengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputStrengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.ContentPolicyConfigFiltersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference_Override(t TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.ContentPolicyConfigFiltersConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputAction(val *string) {
	if err := j.validateSetInputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputEnabled(val interface{}) {
	if err := j.validateSetInputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputModalities(val *[]*string) {
	if err := j.validateSetInputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputModalities",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInputStrength(val *string) {
	if err := j.validateSetInputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputStrength",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputAction(val *string) {
	if err := j.validateSetOutputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputEnabled(val interface{}) {
	if err := j.validateSetOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputModalities(val *[]*string) {
	if err := j.validateSetOutputModalitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputModalities",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetOutputStrength(val *string) {
	if err := j.validateSetOutputStrengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputStrength",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetInputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetInputModalities() {
	_jsii_.InvokeVoid(
		t,
		"resetInputModalities",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ResetOutputModalities() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputModalities",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGuardrail_ContentPolicyConfigFiltersConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

