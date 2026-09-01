package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerFlowDefinition_HumanLoopConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerFlowDefinition_HumanLoopConfigProperty)
	// Experimental.
	PublicWorkforceTaskPrice() AwsSagemakerFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference
	// Experimental.
	PublicWorkforceTaskPriceInput() *AwsSagemakerFlowDefinition_PublicWorkforceTaskPriceProperty
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
	PutPublicWorkforceTaskPrice(value *AwsSagemakerFlowDefinition_PublicWorkforceTaskPriceProperty)
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

// The jsii proxy struct for AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) HumanTaskUiArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTaskUiArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) HumanTaskUiArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanTaskUiArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) InternalValue() *AwsSagemakerFlowDefinition_HumanLoopConfigProperty {
	var returns *AwsSagemakerFlowDefinition_HumanLoopConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) PublicWorkforceTaskPrice() AwsSagemakerFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference {
	var returns AwsSagemakerFlowDefinition_PublicWorkforceTaskPricePropertyOutputReference
	_jsii_.Get(
		j,
		"publicWorkforceTaskPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) PublicWorkforceTaskPriceInput() *AwsSagemakerFlowDefinition_PublicWorkforceTaskPriceProperty {
	var returns *AwsSagemakerFlowDefinition_PublicWorkforceTaskPriceProperty
	_jsii_.Get(
		j,
		"publicWorkforceTaskPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskAvailabilityLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskAvailabilityLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTimeLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTimeLimitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TaskTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) WorkteamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) WorkteamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFlowDefinition.HumanLoopConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference_Override(a AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFlowDefinition.HumanLoopConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetHumanTaskUiArn(val *string) {
	if err := j.validateSetHumanTaskUiArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"humanTaskUiArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerFlowDefinition_HumanLoopConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskAvailabilityLifetimeInSeconds(val *float64) {
	if err := j.validateSetTaskAvailabilityLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAvailabilityLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskDescription(val *string) {
	if err := j.validateSetTaskDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDescription",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskKeywords(val *[]*string) {
	if err := j.validateSetTaskKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKeywords",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskTimeLimitInSeconds(val *float64) {
	if err := j.validateSetTaskTimeLimitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTimeLimitInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTaskTitle(val *string) {
	if err := j.validateSetTaskTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTitle",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference)SetWorkteamArn(val *string) {
	if err := j.validateSetWorkteamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workteamArn",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) PutPublicWorkforceTaskPrice(value *AwsSagemakerFlowDefinition_PublicWorkforceTaskPriceProperty) {
	if err := a.validatePutPublicWorkforceTaskPriceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPublicWorkforceTaskPrice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetPublicWorkforceTaskPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicWorkforceTaskPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskAvailabilityLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskAvailabilityLifetimeInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskKeywords() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskKeywords",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ResetTaskTimeLimitInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskTimeLimitInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerFlowDefinition_HumanLoopConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

