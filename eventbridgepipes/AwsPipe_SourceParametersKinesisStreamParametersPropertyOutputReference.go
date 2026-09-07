package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchSize() *float64
	// Experimental.
	SetBatchSize(val *float64)
	// Experimental.
	BatchSizeInput() *float64
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
	DeadLetterConfig() AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPipe_SourceParametersKinesisStreamParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_SourceParametersKinesisStreamParametersProperty)
	// Experimental.
	MaximumBatchingWindowInSeconds() *float64
	// Experimental.
	SetMaximumBatchingWindowInSeconds(val *float64)
	// Experimental.
	MaximumBatchingWindowInSecondsInput() *float64
	// Experimental.
	MaximumRecordAgeInSeconds() *float64
	// Experimental.
	SetMaximumRecordAgeInSeconds(val *float64)
	// Experimental.
	MaximumRecordAgeInSecondsInput() *float64
	// Experimental.
	MaximumRetryAttempts() *float64
	// Experimental.
	SetMaximumRetryAttempts(val *float64)
	// Experimental.
	MaximumRetryAttemptsInput() *float64
	// Experimental.
	OnPartialBatchItemFailure() *string
	// Experimental.
	SetOnPartialBatchItemFailure(val *string)
	// Experimental.
	OnPartialBatchItemFailureInput() *string
	// Experimental.
	ParallelizationFactor() *float64
	// Experimental.
	SetParallelizationFactor(val *float64)
	// Experimental.
	ParallelizationFactorInput() *float64
	// Experimental.
	StartingPosition() *string
	// Experimental.
	SetStartingPosition(val *string)
	// Experimental.
	StartingPositionInput() *string
	// Experimental.
	StartingPositionTimestamp() *string
	// Experimental.
	SetStartingPositionTimestamp(val *string)
	// Experimental.
	StartingPositionTimestampInput() *string
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
	PutDeadLetterConfig(value *AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigProperty)
	// Experimental.
	ResetBatchSize()
	// Experimental.
	ResetDeadLetterConfig()
	// Experimental.
	ResetMaximumBatchingWindowInSeconds()
	// Experimental.
	ResetMaximumRecordAgeInSeconds()
	// Experimental.
	ResetMaximumRetryAttempts()
	// Experimental.
	ResetOnPartialBatchItemFailure()
	// Experimental.
	ResetParallelizationFactor()
	// Experimental.
	ResetStartingPositionTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference
type jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) DeadLetterConfig() AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigPropertyOutputReference {
	var returns AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) DeadLetterConfigInput() *AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigProperty {
	var returns *AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) InternalValue() *AwsPipe_SourceParametersKinesisStreamParametersProperty {
	var returns *AwsPipe_SourceParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) OnPartialBatchItemFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) OnPartialBatchItemFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) StartingPositionTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) StartingPositionTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersKinesisStreamParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference_Override(a AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersKinesisStreamParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetInternalValue(val *AwsPipe_SourceParametersKinesisStreamParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetOnPartialBatchItemFailure(val *string) {
	if err := j.validateSetOnPartialBatchItemFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPartialBatchItemFailure",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetStartingPositionTimestamp(val *string) {
	if err := j.validateSetStartingPositionTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) PutDeadLetterConfig(value *AwsPipe_SourceParametersKinesisStreamParametersDeadLetterConfigProperty) {
	if err := a.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetOnPartialBatchItemFailure() {
	_jsii_.InvokeVoid(
		a,
		"resetOnPartialBatchItemFailure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ResetStartingPositionTimestamp() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPositionTimestamp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersKinesisStreamParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

