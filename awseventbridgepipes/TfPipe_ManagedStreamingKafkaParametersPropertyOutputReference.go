package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference interface {
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
	// Experimental.
	ConsumerGroupId() *string
	// Experimental.
	SetConsumerGroupId(val *string)
	// Experimental.
	ConsumerGroupIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Credentials() TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference
	// Experimental.
	CredentialsInput() *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipe_ManagedStreamingKafkaParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_ManagedStreamingKafkaParametersProperty)
	// Experimental.
	MaximumBatchingWindowInSeconds() *float64
	// Experimental.
	SetMaximumBatchingWindowInSeconds(val *float64)
	// Experimental.
	MaximumBatchingWindowInSecondsInput() *float64
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
	TopicName() *string
	// Experimental.
	SetTopicName(val *string)
	// Experimental.
	TopicNameInput() *string
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
	PutCredentials(value *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty)
	// Experimental.
	ResetBatchSize()
	// Experimental.
	ResetConsumerGroupId()
	// Experimental.
	ResetCredentials()
	// Experimental.
	ResetMaximumBatchingWindowInSeconds()
	// Experimental.
	ResetStartingPosition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference
type jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ConsumerGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ConsumerGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) Credentials() TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference {
	var returns TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) CredentialsInput() *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty {
	var returns *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) InternalValue() *TfPipe_ManagedStreamingKafkaParametersProperty {
	var returns *TfPipe_ManagedStreamingKafkaParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) TopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicNameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_ManagedStreamingKafkaParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_ManagedStreamingKafkaParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.ManagedStreamingKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_ManagedStreamingKafkaParametersPropertyOutputReference_Override(t TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.ManagedStreamingKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetConsumerGroupId(val *string) {
	if err := j.validateSetConsumerGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupId",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetInternalValue(val *TfPipe_ManagedStreamingKafkaParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference)SetTopicName(val *string) {
	if err := j.validateSetTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicName",
		val,
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) PutCredentials(value *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty) {
	if err := t.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ResetConsumerGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetConsumerGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_ManagedStreamingKafkaParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

