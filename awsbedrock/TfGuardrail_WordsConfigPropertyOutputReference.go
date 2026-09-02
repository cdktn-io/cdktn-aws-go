package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGuardrail_WordsConfigPropertyOutputReference interface {
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Text() *string
	// Experimental.
	SetText(val *string)
	// Experimental.
	TextInput() *string
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

// The jsii proxy struct for TfGuardrail_WordsConfigPropertyOutputReference
type jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) OutputAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) OutputActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) OutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) OutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) TextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGuardrail_WordsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGuardrail_WordsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGuardrail_WordsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.WordsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGuardrail_WordsConfigPropertyOutputReference_Override(t TfGuardrail_WordsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfGuardrail.WordsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetInputAction(val *string) {
	if err := j.validateSetInputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetInputEnabled(val interface{}) {
	if err := j.validateSetInputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetOutputAction(val *string) {
	if err := j.validateSetOutputActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputAction",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetOutputEnabled(val interface{}) {
	if err := j.validateSetOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputEnabled",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference)SetText(val *string) {
	if err := j.validateSetTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ResetInputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ResetInputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetInputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ResetOutputAction() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ResetOutputEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGuardrail_WordsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

