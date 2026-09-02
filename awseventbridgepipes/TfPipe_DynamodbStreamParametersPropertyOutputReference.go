package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_DynamodbStreamParametersPropertyOutputReference interface {
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
	DeadLetterConfig() TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipe_DynamodbStreamParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_DynamodbStreamParametersProperty)
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
	PutDeadLetterConfig(value *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_DynamodbStreamParametersPropertyOutputReference
type jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) DeadLetterConfig() TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigPropertyOutputReference {
	var returns TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) DeadLetterConfigInput() *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty {
	var returns *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) InternalValue() *TfPipe_DynamodbStreamParametersProperty {
	var returns *TfPipe_DynamodbStreamParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) OnPartialBatchItemFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) OnPartialBatchItemFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPartialBatchItemFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_DynamodbStreamParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_DynamodbStreamParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_DynamodbStreamParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.DynamodbStreamParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_DynamodbStreamParametersPropertyOutputReference_Override(t TfPipe_DynamodbStreamParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.DynamodbStreamParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetInternalValue(val *TfPipe_DynamodbStreamParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetOnPartialBatchItemFailure(val *string) {
	if err := j.validateSetOnPartialBatchItemFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPartialBatchItemFailure",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) PutDeadLetterConfig(value *TfPipe_SourceParametersDynamodbStreamParametersDeadLetterConfigProperty) {
	if err := t.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetOnPartialBatchItemFailure() {
	_jsii_.InvokeVoid(
		t,
		"resetOnPartialBatchItemFailure",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_DynamodbStreamParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

