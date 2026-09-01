package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference interface {
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
	FailureConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList
	// Experimental.
	FailureConditionalInput() interface{}
	// Experimental.
	FailureNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList
	// Experimental.
	FailureNextStepInput() interface{}
	// Experimental.
	FailureResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList
	// Experimental.
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SuccessConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList
	// Experimental.
	SuccessConditionalInput() interface{}
	// Experimental.
	SuccessNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList
	// Experimental.
	SuccessNextStepInput() interface{}
	// Experimental.
	SuccessResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList
	// Experimental.
	SuccessResponseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList
	// Experimental.
	TimeoutConditionalInput() interface{}
	// Experimental.
	TimeoutNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList
	// Experimental.
	TimeoutNextStepInput() interface{}
	// Experimental.
	TimeoutResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList
	// Experimental.
	TimeoutResponseInput() interface{}
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
	PutFailureConditional(value interface{})
	// Experimental.
	PutFailureNextStep(value interface{})
	// Experimental.
	PutFailureResponse(value interface{})
	// Experimental.
	PutSuccessConditional(value interface{})
	// Experimental.
	PutSuccessNextStep(value interface{})
	// Experimental.
	PutSuccessResponse(value interface{})
	// Experimental.
	PutTimeoutConditional(value interface{})
	// Experimental.
	PutTimeoutNextStep(value interface{})
	// Experimental.
	PutTimeoutResponse(value interface{})
	// Experimental.
	ResetFailureConditional()
	// Experimental.
	ResetFailureNextStep()
	// Experimental.
	ResetFailureResponse()
	// Experimental.
	ResetSuccessConditional()
	// Experimental.
	ResetSuccessNextStep()
	// Experimental.
	ResetSuccessResponse()
	// Experimental.
	ResetTimeoutConditional()
	// Experimental.
	ResetTimeoutNextStep()
	// Experimental.
	ResetTimeoutResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference
type jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList
	_jsii_.Get(
		j,
		"successConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList
	_jsii_.Get(
		j,
		"successNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList
	_jsii_.Get(
		j,
		"successResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutConditional() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList
	_jsii_.Get(
		j,
		"timeoutConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutNextStep() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList
	_jsii_.Get(
		j,
		"timeoutNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutResponse() AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList {
	var returns AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList
	_jsii_.Get(
		j,
		"timeoutResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutResponseInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference_Override(a AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureConditional(value interface{}) {
	if err := a.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureNextStep(value interface{}) {
	if err := a.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureResponse(value interface{}) {
	if err := a.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessConditional(value interface{}) {
	if err := a.validatePutSuccessConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessNextStep(value interface{}) {
	if err := a.validatePutSuccessNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessResponse(value interface{}) {
	if err := a.validatePutSuccessResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutConditional(value interface{}) {
	if err := a.validatePutTimeoutConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutNextStep(value interface{}) {
	if err := a.validatePutTimeoutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutResponse(value interface{}) {
	if err := a.validatePutTimeoutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

