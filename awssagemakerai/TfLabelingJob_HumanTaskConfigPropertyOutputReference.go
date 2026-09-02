package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLabelingJob_HumanTaskConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AnnotationConsolidationConfig() TfLabelingJob_AnnotationConsolidationConfigPropertyList
	// Experimental.
	AnnotationConsolidationConfigInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxConcurrentTaskCount() *float64
	// Experimental.
	SetMaxConcurrentTaskCount(val *float64)
	// Experimental.
	MaxConcurrentTaskCountInput() *float64
	// Experimental.
	NumberOfHumanWorkersPerDataObject() *float64
	// Experimental.
	SetNumberOfHumanWorkersPerDataObject(val *float64)
	// Experimental.
	NumberOfHumanWorkersPerDataObjectInput() *float64
	// Experimental.
	PreHumanTaskLambdaArn() *string
	// Experimental.
	SetPreHumanTaskLambdaArn(val *string)
	// Experimental.
	PreHumanTaskLambdaArnInput() *string
	// Experimental.
	PublicWorkforceTaskPrice() TfLabelingJob_PublicWorkforceTaskPricePropertyList
	// Experimental.
	PublicWorkforceTaskPriceInput() interface{}
	// Experimental.
	TaskAvailabilityLifetimeInSeconds() *float64
	// Experimental.
	SetTaskAvailabilityLifetimeInSeconds(val *float64)
	// Experimental.
	TaskAvailabilityLifetimeInSecondsInput() *float64
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
	UiConfig() TfLabelingJob_UiConfigPropertyList
	// Experimental.
	UiConfigInput() interface{}
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
	PutAnnotationConsolidationConfig(value interface{})
	// Experimental.
	PutPublicWorkforceTaskPrice(value interface{})
	// Experimental.
	PutUiConfig(value interface{})
	// Experimental.
	ResetAnnotationConsolidationConfig()
	// Experimental.
	ResetMaxConcurrentTaskCount()
	// Experimental.
	ResetPreHumanTaskLambdaArn()
	// Experimental.
	ResetPublicWorkforceTaskPrice()
	// Experimental.
	ResetTaskAvailabilityLifetimeInSeconds()
	// Experimental.
	ResetTaskKeywords()
	// Experimental.
	ResetUiConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLabelingJob_HumanTaskConfigPropertyOutputReference
type jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) AnnotationConsolidationConfig() TfLabelingJob_AnnotationConsolidationConfigPropertyList {
	var returns TfLabelingJob_AnnotationConsolidationConfigPropertyList
	_jsii_.Get(
		j,
		"annotationConsolidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) AnnotationConsolidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"annotationConsolidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) MaxConcurrentTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) MaxConcurrentTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) NumberOfHumanWorkersPerDataObject() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfHumanWorkersPerDataObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) NumberOfHumanWorkersPerDataObjectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfHumanWorkersPerDataObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PreHumanTaskLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preHumanTaskLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PreHumanTaskLambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preHumanTaskLambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PublicWorkforceTaskPrice() TfLabelingJob_PublicWorkforceTaskPricePropertyList {
	var returns TfLabelingJob_PublicWorkforceTaskPricePropertyList
	_jsii_.Get(
		j,
		"publicWorkforceTaskPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PublicWorkforceTaskPriceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicWorkforceTaskPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskAvailabilityLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskAvailabilityLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskAvailabilityLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskKeywords() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskKeywordsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"taskKeywordsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskTimeLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskTimeLimitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskTimeLimitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TaskTitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskTitleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) UiConfig() TfLabelingJob_UiConfigPropertyList {
	var returns TfLabelingJob_UiConfigPropertyList
	_jsii_.Get(
		j,
		"uiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) UiConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) WorkteamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) WorkteamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workteamArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLabelingJob_HumanTaskConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfLabelingJob_HumanTaskConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLabelingJob_HumanTaskConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob.HumanTaskConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLabelingJob_HumanTaskConfigPropertyOutputReference_Override(t TfLabelingJob_HumanTaskConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfLabelingJob.HumanTaskConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetMaxConcurrentTaskCount(val *float64) {
	if err := j.validateSetMaxConcurrentTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentTaskCount",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetNumberOfHumanWorkersPerDataObject(val *float64) {
	if err := j.validateSetNumberOfHumanWorkersPerDataObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfHumanWorkersPerDataObject",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetPreHumanTaskLambdaArn(val *string) {
	if err := j.validateSetPreHumanTaskLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preHumanTaskLambdaArn",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTaskAvailabilityLifetimeInSeconds(val *float64) {
	if err := j.validateSetTaskAvailabilityLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskAvailabilityLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTaskDescription(val *string) {
	if err := j.validateSetTaskDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDescription",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTaskKeywords(val *[]*string) {
	if err := j.validateSetTaskKeywordsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskKeywords",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTaskTimeLimitInSeconds(val *float64) {
	if err := j.validateSetTaskTimeLimitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTimeLimitInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTaskTitle(val *string) {
	if err := j.validateSetTaskTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskTitle",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference)SetWorkteamArn(val *string) {
	if err := j.validateSetWorkteamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workteamArn",
		val,
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PutAnnotationConsolidationConfig(value interface{}) {
	if err := t.validatePutAnnotationConsolidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnnotationConsolidationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PutPublicWorkforceTaskPrice(value interface{}) {
	if err := t.validatePutPublicWorkforceTaskPriceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPublicWorkforceTaskPrice",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) PutUiConfig(value interface{}) {
	if err := t.validatePutUiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUiConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetAnnotationConsolidationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAnnotationConsolidationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetMaxConcurrentTaskCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrentTaskCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetPreHumanTaskLambdaArn() {
	_jsii_.InvokeVoid(
		t,
		"resetPreHumanTaskLambdaArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetPublicWorkforceTaskPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetPublicWorkforceTaskPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetTaskAvailabilityLifetimeInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskAvailabilityLifetimeInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetTaskKeywords() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskKeywords",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ResetUiConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUiConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLabelingJob_HumanTaskConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

