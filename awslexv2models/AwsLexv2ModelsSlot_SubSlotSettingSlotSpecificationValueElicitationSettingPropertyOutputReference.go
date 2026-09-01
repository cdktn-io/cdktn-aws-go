package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference interface {
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
	DefaultValueSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList
	// Experimental.
	DefaultValueSpecificationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PromptSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList
	// Experimental.
	PromptSpecificationInput() interface{}
	// Experimental.
	SampleUtterance() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList
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
	WaitAndContinueSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList
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

// The jsii proxy struct for AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference
type jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) DefaultValueSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList {
	var returns AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecificationPropertyList
	_jsii_.Get(
		j,
		"defaultValueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) DefaultValueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultValueSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PromptSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList {
	var returns AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecificationPropertyList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) SampleUtterance() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList {
	var returns AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterancePropertyList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecification() AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList {
	var returns AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecificationPropertyList
	_jsii_.Get(
		j,
		"waitAndContinueSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) WaitAndContinueSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitAndContinueSpecificationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference_Override(a AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsSlot.SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutDefaultValueSpecification(value interface{}) {
	if err := a.validatePutDefaultValueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultValueSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutPromptSpecification(value interface{}) {
	if err := a.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutSampleUtterance(value interface{}) {
	if err := a.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) PutWaitAndContinueSpecification(value interface{}) {
	if err := a.validatePutWaitAndContinueSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWaitAndContinueSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetDefaultValueSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultValueSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ResetWaitAndContinueSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitAndContinueSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_SubSlotSettingSlotSpecificationValueElicitationSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

