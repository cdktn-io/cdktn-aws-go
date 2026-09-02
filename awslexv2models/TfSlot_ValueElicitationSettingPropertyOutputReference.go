package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_ValueElicitationSettingPropertyOutputReference interface {
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
	DefaultValueSpecification() TfSlot_ValueElicitationSettingDefaultValueSpecificationPropertyList
	// Experimental.
	DefaultValueSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PromptSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPropertyList
	// Experimental.
	PromptSpecificationInput() interface{}
	// Experimental.
	SampleUtterance() TfSlot_ValueElicitationSettingSampleUtterancePropertyList
	// Experimental.
	SampleUtteranceInput() interface{}
	// Experimental.
	SlotConstraint() *string
	// Experimental.
	SetSlotConstraint(val *string)
	// Experimental.
	SlotConstraintInput() *string
	// Experimental.
	SlotResolutionSetting() TfSlot_SlotResolutionSettingPropertyList
	// Experimental.
	SlotResolutionSettingInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WaitAndContinueSpecification() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationPropertyList
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
	PutSlotResolutionSetting(value interface{})
	// Experimental.
	PutWaitAndContinueSpecification(value interface{})
	// Experimental.
	ResetDefaultValueSpecification()
	// Experimental.
	ResetPromptSpecification()
	// Experimental.
	ResetSampleUtterance()
	// Experimental.
	ResetSlotResolutionSetting()
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

// The jsii proxy struct for TfSlot_ValueElicitationSettingPropertyOutputReference
type jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) DefaultValueSpecification() TfSlot_ValueElicitationSettingDefaultValueSpecificationPropertyList {
	var returns TfSlot_ValueElicitationSettingDefaultValueSpecificationPropertyList
	_jsii_.Get(
		j,
		"defaultValueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) DefaultValueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValueSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PromptSpecification() TfSlot_ValueElicitationSettingPromptSpecificationPropertyList {
	var returns TfSlot_ValueElicitationSettingPromptSpecificationPropertyList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SampleUtterance() TfSlot_ValueElicitationSettingSampleUtterancePropertyList {
	var returns TfSlot_ValueElicitationSettingSampleUtterancePropertyList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SlotConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SlotConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SlotResolutionSetting() TfSlot_SlotResolutionSettingPropertyList {
	var returns TfSlot_SlotResolutionSettingPropertyList
	_jsii_.Get(
		j,
		"slotResolutionSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) SlotResolutionSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotResolutionSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecification() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationPropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationPropertyList
	_jsii_.Get(
		j,
		"waitAndContinueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitAndContinueSpecificationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_ValueElicitationSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_ValueElicitationSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_ValueElicitationSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_ValueElicitationSettingPropertyOutputReference_Override(t TfSlot_ValueElicitationSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetSlotConstraint(val *string) {
	if err := j.validateSetSlotConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotConstraint",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PutDefaultValueSpecification(value interface{}) {
	if err := t.validatePutDefaultValueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultValueSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PutPromptSpecification(value interface{}) {
	if err := t.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PutSampleUtterance(value interface{}) {
	if err := t.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PutSlotResolutionSetting(value interface{}) {
	if err := t.validatePutSlotResolutionSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlotResolutionSetting",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) PutWaitAndContinueSpecification(value interface{}) {
	if err := t.validatePutWaitAndContinueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWaitAndContinueSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ResetDefaultValueSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultValueSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ResetSlotResolutionSetting() {
	_jsii_.InvokeVoid(
		t,
		"resetSlotResolutionSetting",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ResetWaitAndContinueSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitAndContinueSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

