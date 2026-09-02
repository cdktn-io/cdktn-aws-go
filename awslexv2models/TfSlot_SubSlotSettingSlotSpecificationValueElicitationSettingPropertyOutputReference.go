package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference interface {
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
	DefaultValueSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList
	// Experimental.
	DefaultValueSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PromptSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList
	// Experimental.
	PromptSpecificationInput() interface{}
	// Experimental.
	SampleUtterance() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList
	// Experimental.
	SampleUtteranceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WaitAndContinueSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList
	// Experimental.
	WaitAndContinueSpecificationInput() interface{}
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
	PutDefaultValueSpecification(value interface{})
	// Experimental.
	PutPromptSpecification(value interface{})
	// Experimental.
	PutSampleUtterance(value interface{})
	// Experimental.
	PutWaitAndContinueSpecification(value interface{})
	// Experimental.
	ResetDefaultValueSpecification()
	// Experimental.
	ResetPromptSpecification()
	// Experimental.
	ResetSampleUtterance()
	// Experimental.
	ResetWaitAndContinueSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference
type jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) DefaultValueSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList {
	var returns TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList
	_jsii_.Get(
		j,
		"defaultValueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) DefaultValueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValueSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PromptSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList {
	var returns TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) SampleUtterance() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList {
	var returns TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecification() TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList {
	var returns TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList
	_jsii_.Get(
		j,
		"waitAndContinueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitAndContinueSpecificationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference_Override(t TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutDefaultValueSpecification(value interface{}) {
	if err := t.validatePutDefaultValueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultValueSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutPromptSpecification(value interface{}) {
	if err := t.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutSampleUtterance(value interface{}) {
	if err := t.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutWaitAndContinueSpecification(value interface{}) {
	if err := t.validatePutWaitAndContinueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWaitAndContinueSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetDefaultValueSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultValueSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetWaitAndContinueSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitAndContinueSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

