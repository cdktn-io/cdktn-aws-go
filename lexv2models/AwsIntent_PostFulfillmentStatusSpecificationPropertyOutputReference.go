package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference interface {
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
	FailureConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalPropertyList
	// Experimental.
	FailureConditionalInput() interface{}
	// Experimental.
	FailureNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepPropertyList
	// Experimental.
	FailureNextStepInput() interface{}
	// Experimental.
	FailureResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponsePropertyList
	// Experimental.
	FailureResponseInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SuccessConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalPropertyList
	// Experimental.
	SuccessConditionalInput() interface{}
	// Experimental.
	SuccessNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepPropertyList
	// Experimental.
	SuccessNextStepInput() interface{}
	// Experimental.
	SuccessResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponsePropertyList
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
	TimeoutConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalPropertyList
	// Experimental.
	TimeoutConditionalInput() interface{}
	// Experimental.
	TimeoutNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepPropertyList
	// Experimental.
	TimeoutNextStepInput() interface{}
	// Experimental.
	TimeoutResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponsePropertyList
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

// The jsii proxy struct for AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference
type jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalPropertyList
	_jsii_.Get(
		j,
		"failureConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepPropertyList
	_jsii_.Get(
		j,
		"failureNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponsePropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponsePropertyList
	_jsii_.Get(
		j,
		"failureResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) FailureResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failureResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalPropertyList
	_jsii_.Get(
		j,
		"successConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepPropertyList
	_jsii_.Get(
		j,
		"successNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponsePropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponsePropertyList
	_jsii_.Get(
		j,
		"successResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) SuccessResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutConditional() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalPropertyList
	_jsii_.Get(
		j,
		"timeoutConditional",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutConditionalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConditionalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutNextStep() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepPropertyList
	_jsii_.Get(
		j,
		"timeoutNextStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutNextStepInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutNextStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutResponse() AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponsePropertyList {
	var returns AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponsePropertyList
	_jsii_.Get(
		j,
		"timeoutResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) TimeoutResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutResponseInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.PostFulfillmentStatusSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference_Override(a AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.PostFulfillmentStatusSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutFailureConditional(value interface{}) {
	if err := a.validatePutFailureConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutFailureNextStep(value interface{}) {
	if err := a.validatePutFailureNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutFailureResponse(value interface{}) {
	if err := a.validatePutFailureResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutSuccessConditional(value interface{}) {
	if err := a.validatePutSuccessConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutSuccessNextStep(value interface{}) {
	if err := a.validatePutSuccessNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutSuccessResponse(value interface{}) {
	if err := a.validatePutSuccessResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSuccessResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutTimeoutConditional(value interface{}) {
	if err := a.validatePutTimeoutConditionalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutConditional",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutTimeoutNextStep(value interface{}) {
	if err := a.validatePutTimeoutNextStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutNextStep",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) PutTimeoutResponse(value interface{}) {
	if err := a.validatePutTimeoutResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeoutResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetFailureConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetFailureNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetFailureResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetSuccessConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetSuccessNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetSuccessResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetTimeoutConditional() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutConditional",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetTimeoutNextStep() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutNextStep",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ResetTimeoutResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

