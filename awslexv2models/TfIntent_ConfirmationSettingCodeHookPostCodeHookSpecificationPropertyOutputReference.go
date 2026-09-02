package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference interface {
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
	FailureConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList
	// Experimental.
	FailureConditionalInput() interface{}
	// Experimental.
	FailureNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList
	// Experimental.
	FailureNextStepInput() interface{}
	// Experimental.
	FailureResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList
	// Experimental.
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SuccessConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList
	// Experimental.
	SuccessConditionalInput() interface{}
	// Experimental.
	SuccessNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList
	// Experimental.
	SuccessNextStepInput() interface{}
	// Experimental.
	SuccessResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList
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
	TimeoutConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList
	// Experimental.
	TimeoutConditionalInput() interface{}
	// Experimental.
	TimeoutNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList
	// Experimental.
	TimeoutNextStepInput() interface{}
	// Experimental.
	TimeoutResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList
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

// The jsii proxy struct for TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference
type jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureConditionalPropertyList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureNextStepPropertyList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationFailureResponsePropertyList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessConditionalPropertyList
	_jsii_.Get(
		j,
		"successConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyList
	_jsii_.Get(
		j,
		"successNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationSuccessResponsePropertyList
	_jsii_.Get(
		j,
		"successResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) SuccessResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutConditional() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutConditionalPropertyList
	_jsii_.Get(
		j,
		"timeoutConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutNextStep() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutNextStepPropertyList
	_jsii_.Get(
		j,
		"timeoutNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutResponse() TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList {
	var returns TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationTimeoutResponsePropertyList
	_jsii_.Get(
		j,
		"timeoutResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) TimeoutResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutResponseInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference_Override(t TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureConditional(value interface{}) {
	if err := t.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureNextStep(value interface{}) {
	if err := t.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutFailureResponse(value interface{}) {
	if err := t.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessConditional(value interface{}) {
	if err := t.validatePutSuccessConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSuccessConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessNextStep(value interface{}) {
	if err := t.validatePutSuccessNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSuccessNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutSuccessResponse(value interface{}) {
	if err := t.validatePutSuccessResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSuccessResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutConditional(value interface{}) {
	if err := t.validatePutTimeoutConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeoutConditional",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutNextStep(value interface{}) {
	if err := t.validatePutTimeoutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeoutNextStep",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) PutTimeoutResponse(value interface{}) {
	if err := t.validatePutTimeoutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeoutResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetSuccessResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutConditional() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutConditional",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutNextStep() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutNextStep",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ResetTimeoutResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_ConfirmationSettingCodeHookPostCodeHookSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

