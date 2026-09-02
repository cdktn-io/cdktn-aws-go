package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_SelfManagedKafkaParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalBootstrapServers() *[]*string
	// Experimental.
	SetAdditionalBootstrapServers(val *[]*string)
	// Experimental.
	AdditionalBootstrapServersInput() *[]*string
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
	Credentials() TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference
	// Experimental.
	CredentialsInput() *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPipe_SelfManagedKafkaParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_SelfManagedKafkaParametersProperty)
	// Experimental.
	MaximumBatchingWindowInSeconds() *float64
	// Experimental.
	SetMaximumBatchingWindowInSeconds(val *float64)
	// Experimental.
	MaximumBatchingWindowInSecondsInput() *float64
	// Experimental.
	ServerRootCaCertificate() *string
	// Experimental.
	SetServerRootCaCertificate(val *string)
	// Experimental.
	ServerRootCaCertificateInput() *string
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
	Vpc() TfPipe_VpcPropertyOutputReference
	// Experimental.
	VpcInput() *TfPipe_VpcProperty
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
	PutCredentials(value *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty)
	// Experimental.
	PutVpc(value *TfPipe_VpcProperty)
	// Experimental.
	ResetAdditionalBootstrapServers()
	// Experimental.
	ResetBatchSize()
	// Experimental.
	ResetConsumerGroupId()
	// Experimental.
	ResetCredentials()
	// Experimental.
	ResetMaximumBatchingWindowInSeconds()
	// Experimental.
	ResetServerRootCaCertificate()
	// Experimental.
	ResetStartingPosition()
	// Experimental.
	ResetVpc()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_SelfManagedKafkaParametersPropertyOutputReference
type jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) AdditionalBootstrapServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalBootstrapServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) AdditionalBootstrapServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalBootstrapServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ConsumerGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ConsumerGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) Credentials() TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference {
	var returns TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) CredentialsInput() *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty {
	var returns *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) InternalValue() *TfPipe_SelfManagedKafkaParametersProperty {
	var returns *TfPipe_SelfManagedKafkaParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ServerRootCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ServerRootCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) TopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) TopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) Vpc() TfPipe_VpcPropertyOutputReference {
	var returns TfPipe_VpcPropertyOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) VpcInput() *TfPipe_VpcProperty {
	var returns *TfPipe_VpcProperty
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_SelfManagedKafkaParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_SelfManagedKafkaParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_SelfManagedKafkaParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SelfManagedKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_SelfManagedKafkaParametersPropertyOutputReference_Override(t TfPipe_SelfManagedKafkaParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SelfManagedKafkaParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetAdditionalBootstrapServers(val *[]*string) {
	if err := j.validateSetAdditionalBootstrapServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalBootstrapServers",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetConsumerGroupId(val *string) {
	if err := j.validateSetConsumerGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerGroupId",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetInternalValue(val *TfPipe_SelfManagedKafkaParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetServerRootCaCertificate(val *string) {
	if err := j.validateSetServerRootCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverRootCaCertificate",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference)SetTopicName(val *string) {
	if err := j.validateSetTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topicName",
		val,
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) PutCredentials(value *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty) {
	if err := t.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) PutVpc(value *TfPipe_VpcProperty) {
	if err := t.validatePutVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetAdditionalBootstrapServers() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalBootstrapServers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetBatchSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetConsumerGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetConsumerGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetServerRootCaCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetServerRootCaCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ResetVpc() {
	_jsii_.InvokeVoid(
		t,
		"resetVpc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_SelfManagedKafkaParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

