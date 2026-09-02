package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlowDefinition_HumanLoopConfigPropertyOutputReference interface {
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
	HumanTaskUiArn() *string
	// Experimental.
	SetHumanTaskUiArn(val *string)
	// Experimental.
	HumanTaskUiArnInput() *string
	// Experimental.
	InternalValue() *TfFlowDefinition_HumanLoopConfigProperty
	// Experimental.
	SetInternalValue(val *TfFlowDefinition_HumanLoopConfigProperty)
	// Experimental.
	PublicWorkforceTaskPrice() TfFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference
	// Experimental.
	PublicWorkforceTaskPriceInput() *TfFlowDefinition_PublicWorkforceTaskPriceProperty
	// Experimental.
	TaskAvailabilityLifetimeInSeconds() *float64
	// Experimental.
	SetTaskAvailabilityLifetimeInSeconds(val *float64)
	// Experimental.
	TaskAvailabilityLifetimeInSecondsInput() *float64
	// Experimental.
	TaskCount() *float64
	// Experimental.
	SetTaskCount(val *float64)
	// Experimental.
	TaskCountInput() *float64
	// Experimental.
	TaskDescription() *string
	// Experimental.
	SetTaskDescription(val *string)
	// Experimental.
	TaskDescriptionInput() *string
	// Experimental.
	TaskKeywords() *[]*string
	// Experimental.
	SetTaskKeywords(val *[]*string)
	// Experimental.
	TaskKeywordsInput() *[]*string
	// Experimental.
	TaskTimeLimitInSeconds() *float64
	// Experimental.
	SetTaskTimeLimitInSeconds(val *float64)
	// Experimental.
	TaskTimeLimitInSecondsInput() *float64
	// Experimental.
	TaskTitle() *string
	// Experimental.
	SetTaskTitle(val *string)
	// Experimental.
	TaskTitleInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WorkteamArn() *string
	// Experimental.
	SetWorkteamArn(val *string)
	// Experimental.
	WorkteamArnInput() *string
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
	PutPublicWorkforceTaskPrice(value *TfFlowDefinition_PublicWorkforceTaskPriceProperty)
	// Experimental.
	ResetPublicWorkforceTaskPrice()
	// Experimental.
	ResetTaskAvailabilityLifetimeInSeconds()
	// Experimental.
	ResetTaskKeywords()
	// Experimental.
	ResetTaskTimeLimitInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlowDefinition_HumanLoopConfigPropertyOutputReference
type jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) HumanTaskUiArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTaskUiArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) HumanTaskUiArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTaskUiArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) InternalValue() *TfFlowDefinition_HumanLoopConfigProperty {
	var returns *TfFlowDefinition_HumanLoopConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) PublicWorkforceTaskPrice() TfFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference {
	var returns TfFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference
	_jsii_.Get(
		j,
		"publicWorkforceTaskPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) PublicWorkforceTaskPriceInput() *TfFlowDefinition_PublicWorkforceTaskPriceProperty {
	var returns *TfFlowDefinition_PublicWorkforceTaskPriceProperty
	_jsii_.Get(
		j,
		"publicWorkforceTaskPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskAvailabilityLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskAvailabilityLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTimeLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTimeLimitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) WorkteamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) WorkteamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlowDefinition_HumanLoopConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlowDefinition_HumanLoopConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlowDefinition_HumanLoopConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFlowDefinition.HumanLoopConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlowDefinition_HumanLoopConfigPropertyOutputReference_Override(t TfFlowDefinition_HumanLoopConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFlowDefinition.HumanLoopConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetHumanTaskUiArn(val *string) {
	if err := j.validateSetHumanTaskUiArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"humanTaskUiArn",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetInternalValue(val *TfFlowDefinition_HumanLoopConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskAvailabilityLifetimeInSeconds(val *float64) {
	if err := j.validateSetTaskAvailabilityLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAvailabilityLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskDescription(val *string) {
	if err := j.validateSetTaskDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDescription",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskKeywords(val *[]*string) {
	if err := j.validateSetTaskKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKeywords",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskTimeLimitInSeconds(val *float64) {
	if err := j.validateSetTaskTimeLimitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTimeLimitInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskTitle(val *string) {
	if err := j.validateSetTaskTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTitle",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference)SetWorkteamArn(val *string) {
	if err := j.validateSetWorkteamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workteamArn",
		val,
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) PutPublicWorkforceTaskPrice(value *TfFlowDefinition_PublicWorkforceTaskPriceProperty) {
	if err := t.validatePutPublicWorkforceTaskPriceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPublicWorkforceTaskPrice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetPublicWorkforceTaskPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetPublicWorkforceTaskPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskAvailabilityLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskAvailabilityLifetimeInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskKeywords() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskKeywords",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskTimeLimitInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskTimeLimitInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlowDefinition_HumanLoopConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

