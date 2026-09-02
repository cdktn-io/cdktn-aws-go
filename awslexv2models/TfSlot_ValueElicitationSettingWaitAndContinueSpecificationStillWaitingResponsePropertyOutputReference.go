package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowInterrupt() interface{}
	// Experimental.
	SetAllowInterrupt(val interface{})
	// Experimental.
	AllowInterruptInput() interface{}
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
	FrequencyInSeconds() *float64
	// Experimental.
	SetFrequencyInSeconds(val *float64)
	// Experimental.
	FrequencyInSecondsInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MessageGroup() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupPropertyList
	// Experimental.
	MessageGroupInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutInSeconds() *float64
	// Experimental.
	SetTimeoutInSeconds(val *float64)
	// Experimental.
	TimeoutInSecondsInput() *float64
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
	PutMessageGroup(value interface{})
	// Experimental.
	ResetAllowInterrupt()
	// Experimental.
	ResetMessageGroup()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference
type jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) AllowInterrupt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterrupt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) AllowInterruptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInterruptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) FrequencyInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) FrequencyInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frequencyInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) MessageGroup() TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupPropertyList {
	var returns TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponseMessageGroupPropertyList
	_jsii_.Get(
		j,
		"messageGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) MessageGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference_Override(t TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfSlot.ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetAllowInterrupt(val interface{}) {
	if err := j.validateSetAllowInterruptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowInterrupt",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetFrequencyInSeconds(val *float64) {
	if err := j.validateSetFrequencyInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"frequencyInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) PutMessageGroup(value interface{}) {
	if err := t.validatePutMessageGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMessageGroup",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ResetAllowInterrupt() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowInterrupt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ResetMessageGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetMessageGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationStillWaitingResponsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

