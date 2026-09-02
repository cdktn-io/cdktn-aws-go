package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGuardrail_RegexesConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() *string
	// Experimental.
	SetAction(val *string)
	// Experimental.
	ActionInput() *string
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
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
	Pattern() *string
	// Experimental.
	SetPattern(val *string)
	// Experimental.
	PatternInput() *string
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
	ResetDescription()
	// Experimental.
	ResetInputAction()
	// Experimental.
	ResetInputEnabled()
	// Experimental.
	ResetOutputAction()
	// Experimental.
	ResetOutputEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGuardrail_RegexesConfigPropertyOutputReference
type jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) OutputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) OutputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) OutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) OutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGuardrail_RegexesConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGuardrail_RegexesConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGuardrail_RegexesConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.RegexesConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGuardrail_RegexesConfigPropertyOutputReference_Override(t TfGuardrail_RegexesConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.RegexesConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetInputAction(val *string) {
	if err := j.validateSetInputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetInputEnabled(val interface{}) {
	if err := j.validateSetInputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetOutputAction(val *string) {
	if err := j.validateSetOutputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetOutputEnabled(val interface{}) {
	if err := j.validateSetOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ResetInputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ResetInputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetInputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ResetOutputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ResetOutputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGuardrail_RegexesConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

