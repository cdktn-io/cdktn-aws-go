package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_ConfirmationSettingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Active() interface{}
	// Experimental.
	SetActive(val interface{})
	// Experimental.
	ActiveInput() interface{}
	// Experimental.
	CodeHook() TfIntent_ConfirmationSettingCodeHookPropertyList
	// Experimental.
	CodeHookInput() interface{}
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
	// Experimental.
	ConfirmationConditional() TfIntent_ConfirmationConditionalPropertyList
	// Experimental.
	ConfirmationConditionalInput() interface{}
	// Experimental.
	ConfirmationNextStep() TfIntent_ConfirmationNextStepPropertyList
	// Experimental.
	ConfirmationNextStepInput() interface{}
	// Experimental.
	ConfirmationResponse() TfIntent_ConfirmationResponsePropertyList
	// Experimental.
	ConfirmationResponseInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeclinationConditional() TfIntent_DeclinationConditionalPropertyList
	// Experimental.
	DeclinationConditionalInput() interface{}
	// Experimental.
	DeclinationNextStep() TfIntent_DeclinationNextStepPropertyList
	// Experimental.
	DeclinationNextStepInput() interface{}
	// Experimental.
	DeclinationResponse() TfIntent_DeclinationResponsePropertyList
	// Experimental.
	DeclinationResponseInput() interface{}
	// Experimental.
	ElicitationCodeHook() TfIntent_ElicitationCodeHookPropertyList
	// Experimental.
	ElicitationCodeHookInput() interface{}
	// Experimental.
	FailureConditional() TfIntent_ConfirmationSettingFailureConditionalPropertyList
	// Experimental.
	FailureConditionalInput() interface{}
	// Experimental.
	FailureNextStep() TfIntent_ConfirmationSettingFailureNextStepPropertyList
	// Experimental.
	FailureNextStepInput() interface{}
	// Experimental.
	FailureResponse() TfIntent_ConfirmationSettingFailureResponsePropertyList
	// Experimental.
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PromptSpecification() TfIntent_PromptSpecificationPropertyList
	// Experimental.
	PromptSpecificationInput() interface{}
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
	PutCodeHook(value interface{})
	// Experimental.
	PutConfirmationConditional(value interface{})
	// Experimental.
	PutConfirmationNextStep(value interface{})
	// Experimental.
	PutConfirmationResponse(value interface{})
	// Experimental.
	PutDeclinationConditional(value interface{})
	// Experimental.
	PutDeclinationNextStep(value interface{})
	// Experimental.
	PutDeclinationResponse(value interface{})
	// Experimental.
	PutElicitationCodeHook(value interface{})
	// Experimental.
	PutFailureConditional(value interface{})
	// Experimental.
	PutFailureNextStep(value interface{})
	// Experimental.
	PutFailureResponse(value interface{})
	// Experimental.
	PutPromptSpecification(value interface{})
	// Experimental.
	ResetActive()
	// Experimental.
	ResetCodeHook()
	// Experimental.
	ResetConfirmationConditional()
	// Experimental.
	ResetConfirmationNextStep()
	// Experimental.
	ResetConfirmationResponse()
	// Experimental.
	ResetDeclinationConditional()
	// Experimental.
	ResetDeclinationNextStep()
	// Experimental.
	ResetDeclinationResponse()
	// Experimental.
	ResetElicitationCodeHook()
	// Experimental.
	ResetFailureConditional()
	// Experimental.
	ResetFailureNextStep()
	// Experimental.
	ResetFailureResponse()
	// Experimental.
	ResetPromptSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_ConfirmationSettingPropertyOutputReference
type jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) CodeHook() TfIntent_ConfirmationSettingCodeHookPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPropertyList
	_jsii_.Get(
		j,
		"codeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) CodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationConditional() TfIntent_ConfirmationConditionalPropertyList {
	var returns TfIntent_ConfirmationConditionalPropertyList
	_jsii_.Get(
		j,
		"confirmationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationNextStep() TfIntent_ConfirmationNextStepPropertyList {
	var returns TfIntent_ConfirmationNextStepPropertyList
	_jsii_.Get(
		j,
		"confirmationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationResponse() TfIntent_ConfirmationResponsePropertyList {
	var returns TfIntent_ConfirmationResponsePropertyList
	_jsii_.Get(
		j,
		"confirmationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ConfirmationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationConditional() TfIntent_DeclinationConditionalPropertyList {
	var returns TfIntent_DeclinationConditionalPropertyList
	_jsii_.Get(
		j,
		"declinationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationNextStep() TfIntent_DeclinationNextStepPropertyList {
	var returns TfIntent_DeclinationNextStepPropertyList
	_jsii_.Get(
		j,
		"declinationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationResponse() TfIntent_DeclinationResponsePropertyList {
	var returns TfIntent_DeclinationResponsePropertyList
	_jsii_.Get(
		j,
		"declinationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) DeclinationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ElicitationCodeHook() TfIntent_ElicitationCodeHookPropertyList {
	var returns TfIntent_ElicitationCodeHookPropertyList
	_jsii_.Get(
		j,
		"elicitationCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ElicitationCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elicitationCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureConditional() TfIntent_ConfirmationSettingFailureConditionalPropertyList {
	var returns TfIntent_ConfirmationSettingFailureConditionalPropertyList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureNextStep() TfIntent_ConfirmationSettingFailureNextStepPropertyList {
	var returns TfIntent_ConfirmationSettingFailureNextStepPropertyList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureResponse() TfIntent_ConfirmationSettingFailureResponsePropertyList {
	var returns TfIntent_ConfirmationSettingFailureResponsePropertyList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PromptSpecification() TfIntent_PromptSpecificationPropertyList {
	var returns TfIntent_PromptSpecificationPropertyList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_ConfirmationSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_ConfirmationSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_ConfirmationSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_ConfirmationSettingPropertyOutputReference_Override(t TfIntent_ConfirmationSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutCodeHook(value interface{}) {
	if err := t.validatePutCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeHook",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationConditional(value interface{}) {
	if err := t.validatePutConfirmationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfirmationConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationNextStep(value interface{}) {
	if err := t.validatePutConfirmationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfirmationNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationResponse(value interface{}) {
	if err := t.validatePutConfirmationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfirmationResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationConditional(value interface{}) {
	if err := t.validatePutDeclinationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeclinationConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationNextStep(value interface{}) {
	if err := t.validatePutDeclinationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeclinationNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationResponse(value interface{}) {
	if err := t.validatePutDeclinationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeclinationResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutElicitationCodeHook(value interface{}) {
	if err := t.validatePutElicitationCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElicitationCodeHook",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutFailureConditional(value interface{}) {
	if err := t.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutFailureNextStep(value interface{}) {
	if err := t.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutFailureResponse(value interface{}) {
	if err := t.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) PutPromptSpecification(value interface{}) {
	if err := t.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetActive() {
	_jsii_.InvokeVoid(
		t,
		"resetActive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetCodeHook() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeHook",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetConfirmationConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetConfirmationNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetConfirmationResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetDeclinationConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetDeclinationNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetDeclinationResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetElicitationCodeHook() {
	_jsii_.InvokeVoid(
		t,
		"resetElicitationCodeHook",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

