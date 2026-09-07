package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_ConfirmationSettingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Active() interface{}
	// Experimental.
	SetActive(val interface{})
	// Experimental.
	ActiveInput() interface{}
	// Experimental.
	CodeHook() AwsIntent_ConfirmationSettingCodeHookPropertyList
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
	ConfirmationConditional() AwsIntent_ConfirmationConditionalPropertyList
	// Experimental.
	ConfirmationConditionalInput() interface{}
	// Experimental.
	ConfirmationNextStep() AwsIntent_ConfirmationNextStepPropertyList
	// Experimental.
	ConfirmationNextStepInput() interface{}
	// Experimental.
	ConfirmationResponse() AwsIntent_ConfirmationResponsePropertyList
	// Experimental.
	ConfirmationResponseInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeclinationConditional() AwsIntent_DeclinationConditionalPropertyList
	// Experimental.
	DeclinationConditionalInput() interface{}
	// Experimental.
	DeclinationNextStep() AwsIntent_DeclinationNextStepPropertyList
	// Experimental.
	DeclinationNextStepInput() interface{}
	// Experimental.
	DeclinationResponse() AwsIntent_DeclinationResponsePropertyList
	// Experimental.
	DeclinationResponseInput() interface{}
	// Experimental.
	ElicitationCodeHook() AwsIntent_ElicitationCodeHookPropertyList
	// Experimental.
	ElicitationCodeHookInput() interface{}
	// Experimental.
	FailureConditional() AwsIntent_ConfirmationSettingFailureConditionalPropertyList
	// Experimental.
	FailureConditionalInput() interface{}
	// Experimental.
	FailureNextStep() AwsIntent_ConfirmationSettingFailureNextStepPropertyList
	// Experimental.
	FailureNextStepInput() interface{}
	// Experimental.
	FailureResponse() AwsIntent_ConfirmationSettingFailureResponsePropertyList
	// Experimental.
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PromptSpecification() AwsIntent_PromptSpecificationPropertyList
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

// The jsii proxy struct for AwsIntent_ConfirmationSettingPropertyOutputReference
type jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) Active() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"active",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ActiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) CodeHook() AwsIntent_ConfirmationSettingCodeHookPropertyList {
	var returns AwsIntent_ConfirmationSettingCodeHookPropertyList
	_jsii_.Get(
		j,
		"codeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) CodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationConditional() AwsIntent_ConfirmationConditionalPropertyList {
	var returns AwsIntent_ConfirmationConditionalPropertyList
	_jsii_.Get(
		j,
		"confirmationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationNextStep() AwsIntent_ConfirmationNextStepPropertyList {
	var returns AwsIntent_ConfirmationNextStepPropertyList
	_jsii_.Get(
		j,
		"confirmationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationResponse() AwsIntent_ConfirmationResponsePropertyList {
	var returns AwsIntent_ConfirmationResponsePropertyList
	_jsii_.Get(
		j,
		"confirmationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ConfirmationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationConditional() AwsIntent_DeclinationConditionalPropertyList {
	var returns AwsIntent_DeclinationConditionalPropertyList
	_jsii_.Get(
		j,
		"declinationConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationNextStep() AwsIntent_DeclinationNextStepPropertyList {
	var returns AwsIntent_DeclinationNextStepPropertyList
	_jsii_.Get(
		j,
		"declinationNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationResponse() AwsIntent_DeclinationResponsePropertyList {
	var returns AwsIntent_DeclinationResponsePropertyList
	_jsii_.Get(
		j,
		"declinationResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) DeclinationResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"declinationResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ElicitationCodeHook() AwsIntent_ElicitationCodeHookPropertyList {
	var returns AwsIntent_ElicitationCodeHookPropertyList
	_jsii_.Get(
		j,
		"elicitationCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ElicitationCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elicitationCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureConditional() AwsIntent_ConfirmationSettingFailureConditionalPropertyList {
	var returns AwsIntent_ConfirmationSettingFailureConditionalPropertyList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureNextStep() AwsIntent_ConfirmationSettingFailureNextStepPropertyList {
	var returns AwsIntent_ConfirmationSettingFailureNextStepPropertyList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureResponse() AwsIntent_ConfirmationSettingFailureResponsePropertyList {
	var returns AwsIntent_ConfirmationSettingFailureResponsePropertyList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PromptSpecification() AwsIntent_PromptSpecificationPropertyList {
	var returns AwsIntent_PromptSpecificationPropertyList
	_jsii_.Get(
		j,
		"promptSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PromptSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_ConfirmationSettingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsIntent_ConfirmationSettingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIntent_ConfirmationSettingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_ConfirmationSettingPropertyOutputReference_Override(a AwsIntent_ConfirmationSettingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetActive(val interface{}) {
	if err := j.validateSetActiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"active",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutCodeHook(value interface{}) {
	if err := a.validatePutCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationConditional(value interface{}) {
	if err := a.validatePutConfirmationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfirmationConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationNextStep(value interface{}) {
	if err := a.validatePutConfirmationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfirmationNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutConfirmationResponse(value interface{}) {
	if err := a.validatePutConfirmationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfirmationResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationConditional(value interface{}) {
	if err := a.validatePutDeclinationConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeclinationConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationNextStep(value interface{}) {
	if err := a.validatePutDeclinationNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeclinationNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutDeclinationResponse(value interface{}) {
	if err := a.validatePutDeclinationResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeclinationResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutElicitationCodeHook(value interface{}) {
	if err := a.validatePutElicitationCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElicitationCodeHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutFailureConditional(value interface{}) {
	if err := a.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutFailureNextStep(value interface{}) {
	if err := a.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutFailureResponse(value interface{}) {
	if err := a.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) PutPromptSpecification(value interface{}) {
	if err := a.validatePutPromptSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPromptSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetActive() {
	_jsii_.InvokeVoid(
		a,
		"resetActive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetCodeHook() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetConfirmationConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetConfirmationNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetConfirmationResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetConfirmationResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetDeclinationConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetDeclinationNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetDeclinationResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetDeclinationResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetElicitationCodeHook() {
	_jsii_.InvokeVoid(
		a,
		"resetElicitationCodeHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ResetPromptSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

